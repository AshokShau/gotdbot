package gotdbot

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatStatisticsInviterInfo) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatStatisticsInviterInfo) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatStatisticsInviterInfo) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatStatisticsInviterInfo) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatStatisticsInviterInfo) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatStatisticsInviterInfo) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatStatisticsInviterInfo) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatStatisticsInviterInfo) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatStatisticsInviterInfo) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatStatisticsInviterInfo) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatStatisticsInviterInfo) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatStatisticsInviterInfo) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatStatisticsInviterInfo) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatStatisticsInviterInfo) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatStatisticsInviterInfo) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatStatisticsInviterInfo) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatStatisticsInviterInfo) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatStatisticsInviterInfo) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatStatisticsMessageSenderInfo) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatStatisticsMessageSenderInfo) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatStatisticsMessageSenderInfo) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatStatisticsMessageSenderInfo) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatStatisticsMessageSenderInfo) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatStatisticsMessageSenderInfo) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatStatisticsMessageSenderInfo) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatStatisticsMessageSenderInfo) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatStatisticsMessageSenderInfo) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatStatisticsMessageSenderInfo) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatStatisticsMessageSenderInfo) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatStatisticsMessageSenderInfo) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatStatisticsMessageSenderInfo) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatStatisticsMessageSenderInfo) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatStatisticsMessageSenderInfo) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatStatisticsMessageSenderInfo) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatStatisticsMessageSenderInfo) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatStatisticsMessageSenderInfo) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatStatisticsMessageSenderInfo) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatStatisticsMessageSenderInfo) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatStatisticsMessageSenderInfo) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatStatisticsMessageSenderInfo) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatStatisticsMessageSenderInfo) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatStatisticsMessageSenderInfo) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatStatisticsMessageSenderInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatStatisticsMessageSenderInfo) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatStatisticsMessageSenderInfo) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatStatisticsMessageSenderInfo) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatStatisticsMessageSenderInfo) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatStatisticsMessageSenderInfo) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatStatisticsMessageSenderInfo) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatStatisticsMessageSenderInfo) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatStatisticsMessageSenderInfo) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatStatisticsMessageSenderInfo) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatStatisticsMessageSenderInfo) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatStatisticsMessageSenderInfo) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatStatisticsMessageSenderInfo) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatStatisticsMessageSenderInfo) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatStatisticsMessageSenderInfo) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatStatisticsMessageSenderInfo) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatStatisticsMessageSenderInfo) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatStatisticsMessageSenderInfo) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatStatisticsMessageSenderInfo) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c ChatStatisticsObjectTypeMessage) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, c.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c ChatStatisticsObjectTypeMessage) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, c.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c ChatStatisticsObjectTypeMessage) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, c.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c ChatStatisticsObjectTypeMessage) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, c.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c ChatStatisticsObjectTypeMessage) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, c.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c ChatStatisticsObjectTypeMessage) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, c.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (c ChatStatisticsObjectTypeMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(c.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c ChatStatisticsObjectTypeMessage) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, c.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c ChatStatisticsObjectTypeMessage) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, c.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c ChatStatisticsObjectTypeMessage) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, c.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c ChatStatisticsObjectTypeMessage) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, c.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c ChatStatisticsObjectTypeMessage) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, c.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c ChatStatisticsObjectTypeMessage) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, c.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, c.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, c.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, c.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, c.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, c.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c ChatStatisticsObjectTypeMessage) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, c.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c ChatStatisticsObjectTypeMessage) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, c.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c ChatStatisticsObjectTypeMessage) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, c.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c ChatStatisticsObjectTypeMessage) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, c.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c ChatStatisticsObjectTypeMessage) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, c.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c ChatStatisticsObjectTypeMessage) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, c.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c ChatStatisticsObjectTypeMessage) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, c.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c ChatStatisticsObjectTypeMessage) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, c.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (c ChatStatisticsObjectTypeMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, c.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c ChatStatisticsObjectTypeMessage) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, c.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c ChatStatisticsObjectTypeMessage) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, c.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (c ChatStatisticsObjectTypeMessage) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, c.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatStatisticsObjectTypeMessage) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, c.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c ChatStatisticsObjectTypeMessage) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, c.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c ChatStatisticsObjectTypeMessage) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, c.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c ChatStatisticsObjectTypeMessage) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, c.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (c ChatStatisticsObjectTypeMessage) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, c.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c ChatStatisticsObjectTypeMessage) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, c.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c ChatStatisticsObjectTypeMessage) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, c.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c ChatStatisticsObjectTypeMessage) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, c.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c ChatStatisticsObjectTypeMessage) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, c.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c ChatStatisticsObjectTypeMessage) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, c.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c ChatStatisticsObjectTypeMessage) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, c.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c ChatStatisticsObjectTypeMessage) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, c.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c ChatStatisticsObjectTypeMessage) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, c.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c ChatStatisticsObjectTypeMessage) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, c.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c ChatStatisticsObjectTypeMessage) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, c.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c ChatStatisticsObjectTypeMessage) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, c.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c ChatStatisticsObjectTypeMessage) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, c.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c ChatStatisticsObjectTypeMessage) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, c.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c ChatStatisticsObjectTypeMessage) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, c.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c ChatStatisticsObjectTypeMessage) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, c.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c ChatStatisticsObjectTypeMessage) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, c.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c ChatStatisticsObjectTypeMessage) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, c.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c ChatStatisticsObjectTypeMessage) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, c.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c ChatStatisticsObjectTypeMessage) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, c.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (c ChatStatisticsObjectTypeMessage) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, c.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (c ChatStatisticsObjectTypeMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(c.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c ChatStatisticsObjectTypeMessage) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, c.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c ChatStatisticsObjectTypeMessage) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, c.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c ChatStatisticsObjectTypeMessage) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, c.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c ChatStatisticsObjectTypeMessage) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, c.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c ChatStatisticsObjectTypeMessage) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, c.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c ChatStatisticsObjectTypeMessage) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, c.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c ChatStatisticsObjectTypeMessage) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, c.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (c ChatStatisticsObjectTypeMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, c.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c ChatStatisticsObjectTypeMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, c.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatStatisticsObjectTypeMessage) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, c.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c ChatStatisticsObjectTypeMessage) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, c.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c ChatStatisticsObjectTypeMessage) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, c.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c ChatStatisticsObjectTypeMessage) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, c.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c ChatStatisticsObjectTypeMessage) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, c.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (c ChatStatisticsObjectTypeMessage) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, c.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c ChatStatisticsObjectTypeMessage) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, c.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c ChatStatisticsObjectTypeMessage) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, c.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c ChatStatisticsObjectTypeMessage) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, c.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c ChatStatisticsObjectTypeMessage) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, c.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c ChatStatisticsObjectTypeMessage) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, c.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (c ChatStatisticsObjectTypeMessage) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, c.MessageId)
}

// CloseStory is a helper method for Client.CloseStory
func (c ChatStatisticsObjectTypeStory) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, c.StoryId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (c ChatStatisticsObjectTypeStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, c.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (c ChatStatisticsObjectTypeStory) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, c.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (c ChatStatisticsObjectTypeStory) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, c.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (c ChatStatisticsObjectTypeStory) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, c.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (c ChatStatisticsObjectTypeStory) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, c.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (c ChatStatisticsObjectTypeStory) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, c.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (c ChatStatisticsObjectTypeStory) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, c.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (c ChatStatisticsObjectTypeStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(c.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (c ChatStatisticsObjectTypeStory) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, c.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c ChatStatisticsObjectTypeStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, c.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (c ChatStatisticsObjectTypeStory) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, c.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (c ChatStatisticsObjectTypeStory) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, c.StoryId, optionId, text)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (c ChatStatisticsObjectTypeStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(c.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (c ChatStatisticsObjectTypeStory) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, c.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (c ChatStatisticsObjectTypeStory) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, c.StoryId, isPostedToChatPage)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatThemeEmoji) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, c.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (c ChatThemeEmoji) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(c.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (c ChatThemeEmoji) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(c.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c ChatThemeEmoji) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, c.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatThemeEmoji) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatThemeEmoji) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, c.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatThemeEmoji) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, c.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (c ChatThemeEmoji) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, c.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatThemeEmoji) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, c.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (c ChatThemeEmoji) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, c.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (c ChatThemeEmoji) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(c.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c ChatThemeEmoji) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, c.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatThemeEmoji) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatThemeEmoji) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, c.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatThemeEmoji) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, c.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (c ChatThemeEmoji) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(c.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (c ChatThemeEmoji) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(c.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (c ChatThemeEmoji) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(c.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (c ChatThemeEmoji) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(c.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatThemeEmoji) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, c.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (c ChatThemeEmoji) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(c.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (c ChatThemeEmoji) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(c.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (c ChatThemeEmoji) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, c.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c ChatThemeEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(c.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (c ChatThemeEmoji) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, c.Name)
}

// SetOption is a helper method for Client.SetOption
func (c ChatThemeEmoji) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(c.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (c ChatThemeEmoji) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, c.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatThemeEmoji) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, c.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatThemeEmoji) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(c.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatThemeEmoji) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, c.Name)
}

// CreateBasicGroupChat is a helper method for Client.CreateBasicGroupChat
func (c ChatTypeBasicGroup) CreateBasicGroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateBasicGroupChat(c.BasicGroupId, force)
}

// GetBasicGroup is a helper method for Client.GetBasicGroup
func (c ChatTypeBasicGroup) GetBasicGroup(client *Client) (*BasicGroup, error) {
	return client.GetBasicGroup(c.BasicGroupId)
}

// GetBasicGroupFullInfo is a helper method for Client.GetBasicGroupFullInfo
func (c ChatTypeBasicGroup) GetBasicGroupFullInfo(client *Client) (*BasicGroupFullInfo, error) {
	return client.GetBasicGroupFullInfo(c.BasicGroupId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatTypePrivate) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatTypePrivate) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatTypePrivate) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatTypePrivate) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatTypePrivate) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatTypePrivate) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatTypePrivate) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatTypePrivate) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatTypePrivate) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatTypePrivate) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatTypePrivate) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatTypePrivate) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatTypePrivate) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatTypePrivate) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatTypePrivate) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatTypePrivate) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatTypePrivate) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatTypePrivate) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatTypePrivate) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatTypePrivate) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatTypePrivate) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatTypePrivate) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatTypePrivate) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatTypePrivate) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatTypePrivate) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatTypePrivate) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatTypePrivate) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatTypePrivate) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatTypePrivate) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatTypePrivate) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatTypePrivate) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatTypePrivate) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatTypePrivate) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatTypePrivate) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatTypePrivate) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatTypePrivate) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatTypePrivate) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatTypePrivate) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatTypePrivate) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatTypePrivate) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatTypePrivate) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatTypePrivate) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatTypePrivate) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatTypeSecret) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatTypeSecret) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatTypeSecret) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatTypeSecret) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatTypeSecret) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CloseSecretChat is a helper method for Client.CloseSecretChat
func (c ChatTypeSecret) CloseSecretChat(client *Client) (*Ok, error) {
	return client.CloseSecretChat(c.SecretChatId)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatTypeSecret) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatTypeSecret) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatTypeSecret) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatTypeSecret) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// CreateSecretChat is a helper method for Client.CreateSecretChat
func (c ChatTypeSecret) CreateSecretChat(client *Client) (*Chat, error) {
	return client.CreateSecretChat(c.SecretChatId)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatTypeSecret) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatTypeSecret) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatTypeSecret) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatTypeSecret) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatTypeSecret) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatTypeSecret) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetSecretChat is a helper method for Client.GetSecretChat
func (c ChatTypeSecret) GetSecretChat(client *Client) (*SecretChat, error) {
	return client.GetSecretChat(c.SecretChatId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatTypeSecret) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatTypeSecret) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatTypeSecret) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatTypeSecret) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatTypeSecret) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatTypeSecret) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatTypeSecret) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatTypeSecret) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatTypeSecret) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatTypeSecret) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatTypeSecret) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatTypeSecret) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatTypeSecret) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatTypeSecret) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatTypeSecret) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatTypeSecret) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatTypeSecret) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatTypeSecret) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatTypeSecret) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatTypeSecret) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatTypeSecret) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatTypeSecret) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatTypeSecret) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatTypeSecret) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatTypeSecret) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatTypeSecret) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatTypeSecret) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatTypeSecret) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (c ChatTypeSupergroup) CreateNewSupergroupChat(client *Client, title string, isForum bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, isForum, c.IsChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateSupergroupChat is a helper method for Client.CreateSupergroupChat
func (c ChatTypeSupergroup) CreateSupergroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(c.SupergroupId, force)
}

// DisableAllSupergroupUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (c ChatTypeSupergroup) DisableAllSupergroupUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(c.SupergroupId)
}

// GetChatBoostFeatures is a helper method for Client.GetChatBoostFeatures
func (c ChatTypeSupergroup) GetChatBoostFeatures(client *Client) (*ChatBoostFeatures, error) {
	return client.GetChatBoostFeatures(c.IsChannel)
}

// GetChatBoostLevelFeatures is a helper method for Client.GetChatBoostLevelFeatures
func (c ChatTypeSupergroup) GetChatBoostLevelFeatures(client *Client, level int32) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(c.IsChannel, level)
}

// GetSupergroup is a helper method for Client.GetSupergroup
func (c ChatTypeSupergroup) GetSupergroup(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(c.SupergroupId)
}

// GetSupergroupFullInfo is a helper method for Client.GetSupergroupFullInfo
func (c ChatTypeSupergroup) GetSupergroupFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(c.SupergroupId)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (c ChatTypeSupergroup) GetSupergroupMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(c.SupergroupId, offset, limit, opts)
}

// ReorderSupergroupActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (c ChatTypeSupergroup) ReorderSupergroupActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(c.SupergroupId, usernames)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (c ChatTypeSupergroup) ReportSupergroupAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(c.SupergroupId, messageId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (c ChatTypeSupergroup) ReportSupergroupSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(c.SupergroupId, messageIds)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (c ChatTypeSupergroup) SetSupergroupCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(c.SupergroupId, customEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (c ChatTypeSupergroup) SetSupergroupMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(c.SupergroupId, mainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (c ChatTypeSupergroup) SetSupergroupStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(c.SupergroupId, stickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (c ChatTypeSupergroup) SetSupergroupUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(c.SupergroupId, unrestrictBoostCount)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (c ChatTypeSupergroup) SetSupergroupUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(c.SupergroupId, username)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (c ChatTypeSupergroup) ToggleSupergroupCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(c.SupergroupId, canHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (c ChatTypeSupergroup) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(c.SupergroupId, hasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (c ChatTypeSupergroup) ToggleSupergroupHasAutomaticTranslation(client *Client, hasAutomaticTranslation bool) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(c.SupergroupId, hasAutomaticTranslation)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (c ChatTypeSupergroup) ToggleSupergroupHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(c.SupergroupId, hasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (c ChatTypeSupergroup) ToggleSupergroupIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(c.SupergroupId, isAllHistoryAvailable)
}

// ToggleSupergroupIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (c ChatTypeSupergroup) ToggleSupergroupIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(c.SupergroupId)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (c ChatTypeSupergroup) ToggleSupergroupIsForum(client *Client, isForum bool, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(c.SupergroupId, isForum, hasForumTabs)
}

// ToggleSupergroupJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (c ChatTypeSupergroup) ToggleSupergroupJoinByRequest(client *Client, joinByRequest bool) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(c.SupergroupId, joinByRequest)
}

// ToggleSupergroupJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (c ChatTypeSupergroup) ToggleSupergroupJoinToSendMessages(client *Client, joinToSendMessages bool) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(c.SupergroupId, joinToSendMessages)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (c ChatTypeSupergroup) ToggleSupergroupSignMessages(client *Client, signMessages bool, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(c.SupergroupId, signMessages, showMessageSender)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (c ChatTypeSupergroup) ToggleSupergroupUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(c.SupergroupId, username, isActive)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (c ChecklistTask) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(c.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (c ChecklistTask) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(c.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChecklistTask) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, c.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (c ChecklistTask) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(c.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChecklistTask) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, c.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (c ChecklistTask) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(c.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (c ChecklistTask) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, c.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (c ChecklistTask) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, c.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c ChecklistTask) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, c.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (c ChecklistTask) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(c.Text, toLanguageCode)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c CloseBirthdayUser) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c CloseBirthdayUser) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c CloseBirthdayUser) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c CloseBirthdayUser) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c CloseBirthdayUser) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c CloseBirthdayUser) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c CloseBirthdayUser) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c CloseBirthdayUser) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c CloseBirthdayUser) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c CloseBirthdayUser) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c CloseBirthdayUser) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c CloseBirthdayUser) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c CloseBirthdayUser) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c CloseBirthdayUser) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c CloseBirthdayUser) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c CloseBirthdayUser) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c CloseBirthdayUser) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c CloseBirthdayUser) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c CloseBirthdayUser) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c CloseBirthdayUser) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c CloseBirthdayUser) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c CloseBirthdayUser) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c CloseBirthdayUser) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c CloseBirthdayUser) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c CloseBirthdayUser) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c CloseBirthdayUser) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c CloseBirthdayUser) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c CloseBirthdayUser) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c CloseBirthdayUser) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetBirthdate is a helper method for Client.SetBirthdate
func (c CloseBirthdayUser) SetBirthdate(client *Client) (*Ok, error) {
	return client.SetBirthdate(c.Birthdate)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c CloseBirthdayUser) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c CloseBirthdayUser) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c CloseBirthdayUser) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c CloseBirthdayUser) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c CloseBirthdayUser) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c CloseBirthdayUser) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c CloseBirthdayUser) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c CloseBirthdayUser) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c CloseBirthdayUser) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c CloseBirthdayUser) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c CloseBirthdayUser) SuggestUserBirthdate(client *Client) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, c.Birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c CloseBirthdayUser) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c CloseBirthdayUser) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c CloseBirthdayUser) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (c CollectibleItemInfo) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, c.Url, cacheTime)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (c CollectibleItemInfo) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(c.Currency, c.Amount)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (c CollectibleItemInfo) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, c.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (c CollectibleItemInfo) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, c.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (c CollectibleItemInfo) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(c.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (c CollectibleItemInfo) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(c.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (c CollectibleItemInfo) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, c.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (c CollectibleItemInfo) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(c.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c CollectibleItemInfo) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, c.Url, parameters, opts)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (c CollectibleItemInfo) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, c.Currency, c.Amount)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (c CollectibleItemTypePhoneNumber) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(c.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (c CollectibleItemTypePhoneNumber) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(c.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (c CollectibleItemTypePhoneNumber) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(c.PhoneNumber, opts)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c CollectibleItemTypeUsername) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, c.Username)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (c CollectibleItemTypeUsername) SearchChatAffiliateProgram(client *Client, referrer string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(c.Username, referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (c CollectibleItemTypeUsername) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(c.Username)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (c CollectibleItemTypeUsername) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, c.Username)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (c CollectibleItemTypeUsername) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, c.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (c CollectibleItemTypeUsername) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(c.Username)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (c CollectibleItemTypeUsername) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, c.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (c CollectibleItemTypeUsername) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, c.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (c CollectibleItemTypeUsername) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(c.Username, isActive)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (c ConnectedAffiliateProgram) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(c.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (c ConnectedAffiliateProgram) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(c.BotUserId)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (c ConnectedAffiliateProgram) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, c.Url, cacheTime)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (c ConnectedAffiliateProgram) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(c.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (c ConnectedAffiliateProgram) CheckWebAppFileDownload(client *Client, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(c.BotUserId, fileName, c.Url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (c ConnectedAffiliateProgram) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, c.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (c ConnectedAffiliateProgram) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(c.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (c ConnectedAffiliateProgram) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(c.BotUserId)
}

// DiscardCall is a helper method for Client.DiscardCall
func (c ConnectedAffiliateProgram) DiscardCall(client *Client, callId int32, inviteLink string, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, c.IsDisconnected, inviteLink, duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (c ConnectedAffiliateProgram) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, c.Url)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (c ConnectedAffiliateProgram) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(c.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (c ConnectedAffiliateProgram) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(c.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (c ConnectedAffiliateProgram) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(c.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (c ConnectedAffiliateProgram) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(c.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (c ConnectedAffiliateProgram) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(c.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (c ConnectedAffiliateProgram) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(c.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (c ConnectedAffiliateProgram) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(c.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (c ConnectedAffiliateProgram) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(c.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (c ConnectedAffiliateProgram) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(c.BotUserId)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (c ConnectedAffiliateProgram) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(c.Url)
}

// Get is a helper method for Client.GetConnectedAffiliateProgram
func (c ConnectedAffiliateProgram) Get(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, c.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c ConnectedAffiliateProgram) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(c.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c ConnectedAffiliateProgram) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, c.BotUserId, startParameter, parameters)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (c ConnectedAffiliateProgram) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(c.Url)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (c ConnectedAffiliateProgram) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(c.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (c ConnectedAffiliateProgram) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(c.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c ConnectedAffiliateProgram) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, c.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (c ConnectedAffiliateProgram) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(c.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (c ConnectedAffiliateProgram) GetWebAppUrl(client *Client, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(c.BotUserId, c.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (c ConnectedAffiliateProgram) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(c.Url, onlyLocal)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (c ConnectedAffiliateProgram) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(c.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c ConnectedAffiliateProgram) OpenWebApp(client *Client, chatId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, c.BotUserId, c.Url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (c ConnectedAffiliateProgram) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(c.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (c ConnectedAffiliateProgram) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(c.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (c ConnectedAffiliateProgram) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(c.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (c ConnectedAffiliateProgram) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(c.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c ConnectedAffiliateProgram) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(c.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (c ConnectedAffiliateProgram) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(c.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (c ConnectedAffiliateProgram) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(c.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (c ConnectedAffiliateProgram) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(c.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (c ConnectedAffiliateProgram) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(c.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (c ConnectedAffiliateProgram) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(c.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (c ConnectedAffiliateProgram) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(c.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (c ConnectedAffiliateProgram) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(c.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (c ConnectedAffiliateProgram) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(c.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (c ConnectedAffiliateProgram) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(c.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (c ConnectedAffiliateProgram) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(c.BotUserId, username, isActive)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (c ConnectedAffiliatePrograms) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, c.NextOffset, opts)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (c ConnectedWebsite) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(c.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (c ConnectedWebsite) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(c.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (c ConnectedWebsite) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(c.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (c ConnectedWebsite) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(c.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (c ConnectedWebsite) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, c.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (c ConnectedWebsite) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(c.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (c ConnectedWebsite) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(c.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (c ConnectedWebsite) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(c.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (c ConnectedWebsite) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(c.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (c ConnectedWebsite) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(c.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (c ConnectedWebsite) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(c.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (c ConnectedWebsite) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(c.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (c ConnectedWebsite) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(c.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (c ConnectedWebsite) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(c.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (c ConnectedWebsite) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(c.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (c ConnectedWebsite) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(c.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (c ConnectedWebsite) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, c.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c ConnectedWebsite) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(c.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c ConnectedWebsite) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, c.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (c ConnectedWebsite) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(c.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (c ConnectedWebsite) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(c.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c ConnectedWebsite) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, c.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (c ConnectedWebsite) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(c.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (c ConnectedWebsite) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(c.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (c ConnectedWebsite) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(c.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c ConnectedWebsite) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, c.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (c ConnectedWebsite) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(c.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (c ConnectedWebsite) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(c.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (c ConnectedWebsite) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(c.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (c ConnectedWebsite) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(c.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c ConnectedWebsite) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(c.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (c ConnectedWebsite) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(c.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (c ConnectedWebsite) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(c.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (c ConnectedWebsite) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(c.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (c ConnectedWebsite) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(c.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (c ConnectedWebsite) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(c.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (c ConnectedWebsite) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(c.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (c ConnectedWebsite) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(c.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (c ConnectedWebsite) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(c.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (c ConnectedWebsite) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(c.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (c ConnectedWebsite) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(c.BotUserId, username, isActive)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c Contact) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// Add is a helper method for Client.AddContact
func (c Contact) Add(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c Contact) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c Contact) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c Contact) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c Contact) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c Contact) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c Contact) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c Contact) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c Contact) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c Contact) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c Contact) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c Contact) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c Contact) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c Contact) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c Contact) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c Contact) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c Contact) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c Contact) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c Contact) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c Contact) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c Contact) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c Contact) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c Contact) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c Contact) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c Contact) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c Contact) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// RegisterUser is a helper method for Client.RegisterUser
func (c Contact) RegisterUser(client *Client, disableNotification bool) (*Ok, error) {
	return client.RegisterUser(c.FirstName, c.LastName, disableNotification)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c Contact) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c Contact) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (c Contact) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(c.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (c Contact) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(c.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (c Contact) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(c.PhoneNumber, opts)
}

// SetBusinessAccountName is a helper method for Client.SetBusinessAccountName
func (c Contact) SetBusinessAccountName(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountName(businessConnectionId, c.FirstName, c.LastName)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c Contact) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c Contact) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c Contact) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetName is a helper method for Client.SetName
func (c Contact) SetName(client *Client) (*Ok, error) {
	return client.SetName(c.FirstName, c.LastName)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c Contact) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c Contact) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c Contact) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c Contact) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c Contact) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c Contact) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c Contact) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c Contact) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c Contact) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c Contact) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c Contact) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (c Count) OptimizeStorage(client *Client, size int64, ttl int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, c.Count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c CountryInfo) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, c.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (c CountryInfo) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(c.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (c CountryInfo) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(c.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c CountryInfo) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, c.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c CountryInfo) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c CountryInfo) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, c.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c CountryInfo) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, c.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (c CountryInfo) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, c.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c CountryInfo) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, c.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (c CountryInfo) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, c.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (c CountryInfo) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(c.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c CountryInfo) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, c.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c CountryInfo) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c CountryInfo) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, c.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c CountryInfo) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, c.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (c CountryInfo) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(c.Name, typeField)
}

// GetCountryFlagEmoji is a helper method for Client.GetCountryFlagEmoji
func (c CountryInfo) GetCountryFlagEmoji(client *Client) (*Text, error) {
	return client.GetCountryFlagEmoji(c.CountryCode)
}

// GetOption is a helper method for Client.GetOption
func (c CountryInfo) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(c.Name)
}

// GetPreferredCountryLanguage is a helper method for Client.GetPreferredCountryLanguage
func (c CountryInfo) GetPreferredCountryLanguage(client *Client) (*Text, error) {
	return client.GetPreferredCountryLanguage(c.CountryCode)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (c CountryInfo) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(c.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (c CountryInfo) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(c.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c CountryInfo) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, c.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (c CountryInfo) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(c.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (c CountryInfo) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(c.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (c CountryInfo) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, c.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c CountryInfo) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(c.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (c CountryInfo) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, c.Name)
}

// SetOption is a helper method for Client.SetOption
func (c CountryInfo) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(c.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (c CountryInfo) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, c.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c CountryInfo) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, c.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c CountryInfo) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(c.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c CountryInfo) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, c.Name)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c CountryInfo) ToggleGeneralForumTopicIsHidden(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(chatId, c.IsHidden)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c CreatedBasicGroupChat) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(c.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (c CreatedBasicGroupChat) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(c.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (c CreatedBasicGroupChat) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(c.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c CreatedBasicGroupChat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(c.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c CreatedBasicGroupChat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, c.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (c CreatedBasicGroupChat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(c.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c CreatedBasicGroupChat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(c.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c CreatedBasicGroupChat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(c.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c CreatedBasicGroupChat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(c.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (c CreatedBasicGroupChat) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(c.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (c CreatedBasicGroupChat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c CreatedBasicGroupChat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(c.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c CreatedBasicGroupChat) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(c.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (c CreatedBasicGroupChat) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(c.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (c CreatedBasicGroupChat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(c.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c CreatedBasicGroupChat) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(c.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c CreatedBasicGroupChat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(c.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c CreatedBasicGroupChat) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(c.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (c CreatedBasicGroupChat) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(c.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c CreatedBasicGroupChat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(c.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c CreatedBasicGroupChat) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(c.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c CreatedBasicGroupChat) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(c.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c CreatedBasicGroupChat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(c.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c CreatedBasicGroupChat) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(c.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c CreatedBasicGroupChat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(c.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c CreatedBasicGroupChat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(c.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c CreatedBasicGroupChat) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(c.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (c CreatedBasicGroupChat) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(c.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (c CreatedBasicGroupChat) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(c.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (c CreatedBasicGroupChat) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(c.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (c CreatedBasicGroupChat) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(c.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (c CreatedBasicGroupChat) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(c.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c CreatedBasicGroupChat) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(c.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (c CreatedBasicGroupChat) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(c.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (c CreatedBasicGroupChat) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(c.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (c CreatedBasicGroupChat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(c.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (c CreatedBasicGroupChat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(c.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c CreatedBasicGroupChat) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(c.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (c CreatedBasicGroupChat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(c.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c CreatedBasicGroupChat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c CreatedBasicGroupChat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, c.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c CreatedBasicGroupChat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c CreatedBasicGroupChat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c CreatedBasicGroupChat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, c.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c CreatedBasicGroupChat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c CreatedBasicGroupChat) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(c.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c CreatedBasicGroupChat) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(c.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c CreatedBasicGroupChat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(c.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c CreatedBasicGroupChat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c CreatedBasicGroupChat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(c.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c CreatedBasicGroupChat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c CreatedBasicGroupChat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(c.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c CreatedBasicGroupChat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(c.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c CreatedBasicGroupChat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(c.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c CreatedBasicGroupChat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(c.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (c CreatedBasicGroupChat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(c.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (c CreatedBasicGroupChat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, c.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c CreatedBasicGroupChat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(c.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c CreatedBasicGroupChat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(c.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (c CreatedBasicGroupChat) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(c.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (c CreatedBasicGroupChat) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(c.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (c CreatedBasicGroupChat) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(c.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (c CreatedBasicGroupChat) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(c.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (c CreatedBasicGroupChat) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(c.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (c CreatedBasicGroupChat) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(c.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (c CreatedBasicGroupChat) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(c.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (c CreatedBasicGroupChat) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(c.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (c CreatedBasicGroupChat) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(c.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (c CreatedBasicGroupChat) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(c.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (c CreatedBasicGroupChat) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(c.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c CreatedBasicGroupChat) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(c.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (c CreatedBasicGroupChat) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(c.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c CreatedBasicGroupChat) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(c.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (c CreatedBasicGroupChat) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(c.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c CreatedBasicGroupChat) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(c.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (c CreatedBasicGroupChat) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(c.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c CreatedBasicGroupChat) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(c.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c CreatedBasicGroupChat) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(c.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (c CreatedBasicGroupChat) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(c.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (c CreatedBasicGroupChat) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(c.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (c CreatedBasicGroupChat) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(c.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (c CreatedBasicGroupChat) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(c.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (c CreatedBasicGroupChat) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(c.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (c CreatedBasicGroupChat) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(c.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (c CreatedBasicGroupChat) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(c.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c CreatedBasicGroupChat) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(c.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (c CreatedBasicGroupChat) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(c.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (c CreatedBasicGroupChat) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(c.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (c CreatedBasicGroupChat) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(c.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (c CreatedBasicGroupChat) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(c.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (c CreatedBasicGroupChat) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(c.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (c CreatedBasicGroupChat) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(c.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (c CreatedBasicGroupChat) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(c.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (c CreatedBasicGroupChat) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(c.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (c CreatedBasicGroupChat) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(c.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c CreatedBasicGroupChat) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(c.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (c CreatedBasicGroupChat) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(c.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (c CreatedBasicGroupChat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(c.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (c CreatedBasicGroupChat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(c.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (c CreatedBasicGroupChat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(c.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (c CreatedBasicGroupChat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(c.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c CreatedBasicGroupChat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(c.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c CreatedBasicGroupChat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(c.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c CreatedBasicGroupChat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, c.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (c CreatedBasicGroupChat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(c.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c CreatedBasicGroupChat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(c.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c CreatedBasicGroupChat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(c.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c CreatedBasicGroupChat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(c.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c CreatedBasicGroupChat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, c.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (c CreatedBasicGroupChat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(c.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c CreatedBasicGroupChat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(c.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c CreatedBasicGroupChat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(c.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c CreatedBasicGroupChat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(c.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c CreatedBasicGroupChat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(c.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (c CreatedBasicGroupChat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(c.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c CreatedBasicGroupChat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(c.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c CreatedBasicGroupChat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(c.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c CreatedBasicGroupChat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(c.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c CreatedBasicGroupChat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(c.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c CreatedBasicGroupChat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(c.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (c CreatedBasicGroupChat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(c.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c CreatedBasicGroupChat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(c.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c CreatedBasicGroupChat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(c.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c CreatedBasicGroupChat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(c.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c CreatedBasicGroupChat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(c.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c CreatedBasicGroupChat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(c.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c CreatedBasicGroupChat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(c.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c CreatedBasicGroupChat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(c.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (c CreatedBasicGroupChat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(c.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (c CreatedBasicGroupChat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, c.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (c CreatedBasicGroupChat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(c.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c CreatedBasicGroupChat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(c.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c CreatedBasicGroupChat) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(c.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (c CreatedBasicGroupChat) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(c.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (c CreatedBasicGroupChat) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(c.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c CreatedBasicGroupChat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(c.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c CreatedBasicGroupChat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(c.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (c CreatedBasicGroupChat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(c.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (c CreatedBasicGroupChat) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(c.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (c CreatedBasicGroupChat) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(c.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (c CreatedBasicGroupChat) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(c.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c CreatedBasicGroupChat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(c.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (c CreatedBasicGroupChat) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(c.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (c CreatedBasicGroupChat) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(c.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c CreatedBasicGroupChat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(c.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c CreatedBasicGroupChat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(c.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (c CreatedBasicGroupChat) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(c.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (c CreatedBasicGroupChat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(c.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c CreatedBasicGroupChat) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(c.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c CreatedBasicGroupChat) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(c.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c CreatedBasicGroupChat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(c.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (c CreatedBasicGroupChat) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(c.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (c CreatedBasicGroupChat) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(c.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (c CreatedBasicGroupChat) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(c.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (c CreatedBasicGroupChat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(c.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (c CreatedBasicGroupChat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(c.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c CreatedBasicGroupChat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, c.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c CreatedBasicGroupChat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(c.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (c CreatedBasicGroupChat) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(c.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (c CreatedBasicGroupChat) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(c.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c CreatedBasicGroupChat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(c.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c CreatedBasicGroupChat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(c.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (c CreatedBasicGroupChat) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(c.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (c CreatedBasicGroupChat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (c CreatedBasicGroupChat) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, c.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (c CreatedBasicGroupChat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(c.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (c CreatedBasicGroupChat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (c CreatedBasicGroupChat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(c.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (c CreatedBasicGroupChat) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(c.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (c CreatedBasicGroupChat) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(c.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (c CreatedBasicGroupChat) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(c.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c CreatedBasicGroupChat) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(c.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c CreatedBasicGroupChat) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(c.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c CreatedBasicGroupChat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(c.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (c CreatedBasicGroupChat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(c.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c CreatedBasicGroupChat) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(c.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c CreatedBasicGroupChat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, c.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (c CreatedBasicGroupChat) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(c.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (c CreatedBasicGroupChat) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(c.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (c CreatedBasicGroupChat) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(c.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (c CreatedBasicGroupChat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(c.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c CreatedBasicGroupChat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, c.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (c CreatedBasicGroupChat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (c CreatedBasicGroupChat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (c CreatedBasicGroupChat) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(c.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (c CreatedBasicGroupChat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(c.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (c CreatedBasicGroupChat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(c.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (c CreatedBasicGroupChat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(c.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (c CreatedBasicGroupChat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(c.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c CreatedBasicGroupChat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(c.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c CreatedBasicGroupChat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, c.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (c CreatedBasicGroupChat) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(c.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (c CreatedBasicGroupChat) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(c.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (c CreatedBasicGroupChat) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(c.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (c CreatedBasicGroupChat) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(c.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (c CreatedBasicGroupChat) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(c.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (c CreatedBasicGroupChat) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(c.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (c CreatedBasicGroupChat) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(c.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (c CreatedBasicGroupChat) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(c.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (c CreatedBasicGroupChat) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(c.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (c CreatedBasicGroupChat) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(c.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (c CreatedBasicGroupChat) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(c.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (c CreatedBasicGroupChat) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(c.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c CreatedBasicGroupChat) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(c.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (c CreatedBasicGroupChat) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(c.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (c CreatedBasicGroupChat) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(c.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (c CreatedBasicGroupChat) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(c.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (c CreatedBasicGroupChat) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(c.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (c CreatedBasicGroupChat) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(c.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (c CreatedBasicGroupChat) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(c.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (c CreatedBasicGroupChat) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(c.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (c CreatedBasicGroupChat) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(c.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (c CreatedBasicGroupChat) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(c.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (c CreatedBasicGroupChat) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(c.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (c CreatedBasicGroupChat) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(c.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (c CreatedBasicGroupChat) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(c.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (c CreatedBasicGroupChat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(c.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c CreatedBasicGroupChat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(c.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c CreatedBasicGroupChat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(c.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c CreatedBasicGroupChat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(c.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c CreatedBasicGroupChat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(c.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (c CreatedBasicGroupChat) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(c.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (c CreatedBasicGroupChat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(c.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c CreatedBasicGroupChat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(c.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c CreatedBasicGroupChat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(c.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (c CreatedBasicGroupChat) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(c.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (c CreatedBasicGroupChat) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(c.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c CreatedBasicGroupChat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(c.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (c CreatedBasicGroupChat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(c.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c CreatedBasicGroupChat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, c.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c CreatedBasicGroupChat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(c.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c CreatedBasicGroupChat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(c.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (c CreatedBasicGroupChat) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(c.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (c CreatedBasicGroupChat) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(c.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (c CreatedBasicGroupChat) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(c.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c CreatedBasicGroupChat) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(c.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (c CreatedBasicGroupChat) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(c.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (c CreatedBasicGroupChat) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, c.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (c CreatedBasicGroupChat) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(c.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (c CreatedBasicGroupChat) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(c.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (c CreatedBasicGroupChat) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(c.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (c CreatedBasicGroupChat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(c.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c CreatedBasicGroupChat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(c.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c CreatedBasicGroupChat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(c.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c CreatedBasicGroupChat) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(c.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c CreatedBasicGroupChat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(c.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (c CreatedBasicGroupChat) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(c.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (c CreatedBasicGroupChat) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(c.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (c CreatedBasicGroupChat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(c.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (c CreatedBasicGroupChat) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(c.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (c CreatedBasicGroupChat) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(c.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (c CreatedBasicGroupChat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(c.ChatId, messageIds, forceRead, opts)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (c CurrentWeather) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(c.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (c CurrentWeather) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(c.Emoji)
}

// DecryptGroupCall is a helper method for Client.DecryptGroupCallData
func (d Data) DecryptGroupCall(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, d.Data, opts)
}

// EncryptGroupCall is a helper method for Client.EncryptGroupCallData
func (d Data) EncryptGroupCall(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, d.Data, unencryptedPrefixSize)
}

// SendCallSignaling is a helper method for Client.SendCallSignalingData
func (d Data) SendCallSignaling(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, d.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (d Data) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, d.Data)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (d DatedFile) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, d.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (d DatedFile) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, d.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (d DatedFile) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, d.Date)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (d DateRange) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, d.StartDate, isRtmpStream)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (d DeepLinkInfo) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(d.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (d DeepLinkInfo) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(d.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (d DeepLinkInfo) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, d.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (d DeepLinkInfo) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(d.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (d DeepLinkInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, d.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (d DeepLinkInfo) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(d.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (d DeepLinkInfo) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, d.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (d DeepLinkInfo) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, d.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (d DeepLinkInfo) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, d.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (d DeepLinkInfo) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(d.Text, toLanguageCode)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (d DeviceTokenBlackBerryPush) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(d.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (d DeviceTokenBlackBerryPush) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, d.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (d DeviceTokenBlackBerryPush) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(d.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (d DeviceTokenBlackBerryPush) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(d.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (d DeviceTokenBlackBerryPush) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(d.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (d DeviceTokenBlackBerryPush) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, d.Token)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (d DeviceTokenFirebaseCloudMessaging) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(d.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (d DeviceTokenFirebaseCloudMessaging) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, d.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (d DeviceTokenFirebaseCloudMessaging) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(d.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (d DeviceTokenFirebaseCloudMessaging) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(d.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (d DeviceTokenFirebaseCloudMessaging) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(d.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (d DeviceTokenFirebaseCloudMessaging) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, d.Token)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (d DeviceTokenHuaweiPush) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(d.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (d DeviceTokenHuaweiPush) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, d.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (d DeviceTokenHuaweiPush) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(d.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (d DeviceTokenHuaweiPush) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(d.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (d DeviceTokenHuaweiPush) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(d.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (d DeviceTokenHuaweiPush) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, d.Token)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (d DeviceTokenUbuntuPush) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(d.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (d DeviceTokenUbuntuPush) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, d.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (d DeviceTokenUbuntuPush) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(d.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (d DeviceTokenUbuntuPush) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(d.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (d DeviceTokenUbuntuPush) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(d.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (d DeviceTokenUbuntuPush) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, d.Token)
}

// AddChatMember is a helper method for Client.AddChatMember
func (d DirectMessagesChatTopic) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(d.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (d DirectMessagesChatTopic) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(d.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (d DirectMessagesChatTopic) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(d.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (d DirectMessagesChatTopic) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(d.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (d DirectMessagesChatTopic) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, d.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (d DirectMessagesChatTopic) AddLocalMessage(client *Client, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(d.ChatId, d.SenderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (d DirectMessagesChatTopic) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(d.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (d DirectMessagesChatTopic) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(d.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (d DirectMessagesChatTopic) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(d.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (d DirectMessagesChatTopic) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(d.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (d DirectMessagesChatTopic) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(d.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (d DirectMessagesChatTopic) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(d.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (d DirectMessagesChatTopic) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(d.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (d DirectMessagesChatTopic) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(d.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (d DirectMessagesChatTopic) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(d.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (d DirectMessagesChatTopic) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(d.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (d DirectMessagesChatTopic) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(d.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (d DirectMessagesChatTopic) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(d.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (d DirectMessagesChatTopic) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(d.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (d DirectMessagesChatTopic) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(d.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (d DirectMessagesChatTopic) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(d.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (d DirectMessagesChatTopic) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(d.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (d DirectMessagesChatTopic) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(d.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (d DirectMessagesChatTopic) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(d.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (d DirectMessagesChatTopic) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(d.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (d DirectMessagesChatTopic) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(d.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (d DirectMessagesChatTopic) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(d.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (d DirectMessagesChatTopic) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(d.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (d DirectMessagesChatTopic) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(d.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (d DirectMessagesChatTopic) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(d.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (d DirectMessagesChatTopic) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(d.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (d DirectMessagesChatTopic) DeleteChatMessagesBySender(client *Client) (*Ok, error) {
	return client.DeleteChatMessagesBySender(d.ChatId, d.SenderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (d DirectMessagesChatTopic) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(d.ChatId, messageId)
}

// DeleteHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (d DirectMessagesChatTopic) DeleteHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(d.ChatId, topicId)
}

// DeleteMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (d DirectMessagesChatTopic) DeleteMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(d.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (d DirectMessagesChatTopic) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(d.ChatId, forumTopicId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (d DirectMessagesChatTopic) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, d.SenderId, reportSpam)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (d DirectMessagesChatTopic) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(d.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (d DirectMessagesChatTopic) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(d.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (d DirectMessagesChatTopic) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(d.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (d DirectMessagesChatTopic) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, d.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (d DirectMessagesChatTopic) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, d.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (d DirectMessagesChatTopic) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, d.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (d DirectMessagesChatTopic) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, d.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (d DirectMessagesChatTopic) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, d.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (d DirectMessagesChatTopic) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, d.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (d DirectMessagesChatTopic) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(d.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (d DirectMessagesChatTopic) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(d.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (d DirectMessagesChatTopic) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(d.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (d DirectMessagesChatTopic) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(d.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (d DirectMessagesChatTopic) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(d.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (d DirectMessagesChatTopic) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(d.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (d DirectMessagesChatTopic) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(d.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (d DirectMessagesChatTopic) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(d.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (d DirectMessagesChatTopic) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(d.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (d DirectMessagesChatTopic) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(d.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (d DirectMessagesChatTopic) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(d.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (d DirectMessagesChatTopic) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, d.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (d DirectMessagesChatTopic) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(d.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (d DirectMessagesChatTopic) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(d.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (d DirectMessagesChatTopic) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(d.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (d DirectMessagesChatTopic) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(d.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (d DirectMessagesChatTopic) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(d.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (d DirectMessagesChatTopic) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(d.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (d DirectMessagesChatTopic) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(d.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (d DirectMessagesChatTopic) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(d.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (d DirectMessagesChatTopic) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(d.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (d DirectMessagesChatTopic) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(d.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (d DirectMessagesChatTopic) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(d.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (d DirectMessagesChatTopic) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(d.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (d DirectMessagesChatTopic) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(d.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (d DirectMessagesChatTopic) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(d.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (d DirectMessagesChatTopic) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(d.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (d DirectMessagesChatTopic) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(d.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (d DirectMessagesChatTopic) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(d.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (d DirectMessagesChatTopic) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(d.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (d DirectMessagesChatTopic) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(d.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (d DirectMessagesChatTopic) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(d.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (d DirectMessagesChatTopic) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(d.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (d DirectMessagesChatTopic) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(d.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (d DirectMessagesChatTopic) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(d.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (d DirectMessagesChatTopic) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(d.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (d DirectMessagesChatTopic) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(d.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (d DirectMessagesChatTopic) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(d.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (d DirectMessagesChatTopic) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(d.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (d DirectMessagesChatTopic) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(d.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (d DirectMessagesChatTopic) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(d.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (d DirectMessagesChatTopic) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(d.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (d DirectMessagesChatTopic) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(d.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (d DirectMessagesChatTopic) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(d.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (d DirectMessagesChatTopic) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(d.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (d DirectMessagesChatTopic) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(d.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (d DirectMessagesChatTopic) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(d.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (d DirectMessagesChatTopic) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(d.ChatId)
}

// Get is a helper method for Client.GetDirectMessagesChatTopic
func (d DirectMessagesChatTopic) Get(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(d.ChatId, topicId)
}

// GetHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (d DirectMessagesChatTopic) GetHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(d.ChatId, topicId, fromMessageId, offset, limit)
}

// GetMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (d DirectMessagesChatTopic) GetMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(d.ChatId, topicId, date)
}

// GetRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (d DirectMessagesChatTopic) GetRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(d.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (d DirectMessagesChatTopic) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(d.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (d DirectMessagesChatTopic) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(d.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (d DirectMessagesChatTopic) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(d.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (d DirectMessagesChatTopic) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(d.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (d DirectMessagesChatTopic) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(d.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (d DirectMessagesChatTopic) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(d.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (d DirectMessagesChatTopic) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, d.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (d DirectMessagesChatTopic) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(d.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (d DirectMessagesChatTopic) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(d.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (d DirectMessagesChatTopic) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(d.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (d DirectMessagesChatTopic) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(d.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (d DirectMessagesChatTopic) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, d.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (d DirectMessagesChatTopic) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(d.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (d DirectMessagesChatTopic) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(d.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (d DirectMessagesChatTopic) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(d.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (d DirectMessagesChatTopic) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(d.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (d DirectMessagesChatTopic) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(d.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (d DirectMessagesChatTopic) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(d.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (d DirectMessagesChatTopic) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(d.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (d DirectMessagesChatTopic) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(d.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (d DirectMessagesChatTopic) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(d.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (d DirectMessagesChatTopic) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(d.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (d DirectMessagesChatTopic) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(d.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (d DirectMessagesChatTopic) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(d.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (d DirectMessagesChatTopic) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(d.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (d DirectMessagesChatTopic) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(d.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (d DirectMessagesChatTopic) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(d.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (d DirectMessagesChatTopic) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(d.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (d DirectMessagesChatTopic) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(d.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (d DirectMessagesChatTopic) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(d.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (d DirectMessagesChatTopic) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(d.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (d DirectMessagesChatTopic) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(d.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (d DirectMessagesChatTopic) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, d.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (d DirectMessagesChatTopic) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(d.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (d DirectMessagesChatTopic) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(d.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (d DirectMessagesChatTopic) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(d.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (d DirectMessagesChatTopic) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(d.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (d DirectMessagesChatTopic) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(d.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (d DirectMessagesChatTopic) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(d.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (d DirectMessagesChatTopic) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(d.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (d DirectMessagesChatTopic) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(d.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (d DirectMessagesChatTopic) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(d.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (d DirectMessagesChatTopic) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(d.ChatId)
}

// Loads is a helper method for Client.LoadDirectMessagesChatTopics
func (d DirectMessagesChatTopic) Loads(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(d.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (d DirectMessagesChatTopic) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(d.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (d DirectMessagesChatTopic) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(d.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (d DirectMessagesChatTopic) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(d.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (d DirectMessagesChatTopic) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(d.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (d DirectMessagesChatTopic) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(d.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (d DirectMessagesChatTopic) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(d.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (d DirectMessagesChatTopic) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(d.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (d DirectMessagesChatTopic) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(d.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (d DirectMessagesChatTopic) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(d.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (d DirectMessagesChatTopic) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(d.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (d DirectMessagesChatTopic) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(d.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (d DirectMessagesChatTopic) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(d.ChatId)
}

// ReadAllReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (d DirectMessagesChatTopic) ReadAllReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(d.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (d DirectMessagesChatTopic) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(d.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (d DirectMessagesChatTopic) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(d.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (d DirectMessagesChatTopic) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, d.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (d DirectMessagesChatTopic) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(d.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (d DirectMessagesChatTopic) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(d.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (d DirectMessagesChatTopic) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(d.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (d DirectMessagesChatTopic) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(d.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (d DirectMessagesChatTopic) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(d.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (d DirectMessagesChatTopic) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(d.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (d DirectMessagesChatTopic) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(d.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (d DirectMessagesChatTopic) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, d.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (d DirectMessagesChatTopic) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(d.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (d DirectMessagesChatTopic) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(d.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (d DirectMessagesChatTopic) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(d.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (d DirectMessagesChatTopic) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(d.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (d DirectMessagesChatTopic) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(d.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (d DirectMessagesChatTopic) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(d.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (d DirectMessagesChatTopic) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(d.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (d DirectMessagesChatTopic) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(d.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (d DirectMessagesChatTopic) ReportMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(d.ChatId, messageId, d.SenderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (d DirectMessagesChatTopic) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(d.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (d DirectMessagesChatTopic) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(d.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (d DirectMessagesChatTopic) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, d.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (d DirectMessagesChatTopic) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(d.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (d DirectMessagesChatTopic) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(d.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (d DirectMessagesChatTopic) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(d.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (d DirectMessagesChatTopic) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(d.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (d DirectMessagesChatTopic) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, d.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (d DirectMessagesChatTopic) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, d.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (d DirectMessagesChatTopic) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, d.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (d DirectMessagesChatTopic) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(d.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (d DirectMessagesChatTopic) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(d.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (d DirectMessagesChatTopic) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(d.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (d DirectMessagesChatTopic) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(d.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (d DirectMessagesChatTopic) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(d.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (d DirectMessagesChatTopic) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(d.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (d DirectMessagesChatTopic) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, d.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (d DirectMessagesChatTopic) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(d.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (d DirectMessagesChatTopic) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(d.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (d DirectMessagesChatTopic) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(d.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (d DirectMessagesChatTopic) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(d.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (d DirectMessagesChatTopic) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(d.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (d DirectMessagesChatTopic) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(d.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (d DirectMessagesChatTopic) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(d.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (d DirectMessagesChatTopic) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(d.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (d DirectMessagesChatTopic) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(d.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (d DirectMessagesChatTopic) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(d.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (d DirectMessagesChatTopic) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(d.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (d DirectMessagesChatTopic) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(d.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (d DirectMessagesChatTopic) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(d.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (d DirectMessagesChatTopic) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(d.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (d DirectMessagesChatTopic) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(d.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (d DirectMessagesChatTopic) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(d.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (d DirectMessagesChatTopic) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(d.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (d DirectMessagesChatTopic) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(d.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (d DirectMessagesChatTopic) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(d.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (d DirectMessagesChatTopic) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(d.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (d DirectMessagesChatTopic) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(d.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (d DirectMessagesChatTopic) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(d.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (d DirectMessagesChatTopic) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(d.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (d DirectMessagesChatTopic) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(d.ChatId, title)
}

// SetIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (d DirectMessagesChatTopic) SetIsMarkedAsUnread(client *Client, topicId int64) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(d.ChatId, topicId, d.IsMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (d DirectMessagesChatTopic) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(d.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (d DirectMessagesChatTopic) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(d.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (d DirectMessagesChatTopic) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(d.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (d DirectMessagesChatTopic) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(d.ChatId, messageId, reactionTypes, isBig)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (d DirectMessagesChatTopic) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(d.SenderId, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (d DirectMessagesChatTopic) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(d.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (d DirectMessagesChatTopic) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(d.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (d DirectMessagesChatTopic) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(d.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (d DirectMessagesChatTopic) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(d.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (d DirectMessagesChatTopic) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(d.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (d DirectMessagesChatTopic) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(d.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (d DirectMessagesChatTopic) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(d.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (d DirectMessagesChatTopic) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(d.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (d DirectMessagesChatTopic) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(d.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (d DirectMessagesChatTopic) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, d.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (d DirectMessagesChatTopic) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(d.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (d DirectMessagesChatTopic) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(d.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (d DirectMessagesChatTopic) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(d.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (d DirectMessagesChatTopic) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(d.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (d DirectMessagesChatTopic) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(d.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (d DirectMessagesChatTopic) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(d.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (d DirectMessagesChatTopic) ToggleChatIsMarkedAsUnread(client *Client) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(d.ChatId, d.IsMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (d DirectMessagesChatTopic) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, d.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (d DirectMessagesChatTopic) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(d.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (d DirectMessagesChatTopic) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(d.ChatId, viewAsTopics)
}

// ToggleCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (d DirectMessagesChatTopic) ToggleCanSendUnpaidMessages(client *Client, topicId int64, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(d.ChatId, topicId, d.CanSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (d DirectMessagesChatTopic) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(d.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (d DirectMessagesChatTopic) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(d.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (d DirectMessagesChatTopic) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(d.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (d DirectMessagesChatTopic) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(d.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (d DirectMessagesChatTopic) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(d.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (d DirectMessagesChatTopic) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(d.ChatId)
}

// UnpinAllMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (d DirectMessagesChatTopic) UnpinAllMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(d.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (d DirectMessagesChatTopic) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(d.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (d DirectMessagesChatTopic) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(d.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (d DirectMessagesChatTopic) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(d.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (d DirectMessagesChatTopic) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(d.ChatId, messageIds, forceRead, opts)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (d Document) CheckWebAppFileDownload(client *Client, botUserId int64, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, d.FileName, url)
}

// CleanFileName is a helper method for Client.CleanFileName
func (d Document) CleanFileName(client *Client) (*Text, error) {
	return client.CleanFileName(d.FileName)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (d Document) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(d.MimeType)
}

// GetFileMimeType is a helper method for Client.GetFileMimeType
func (d Document) GetFileMimeType(client *Client) (*Text, error) {
	return client.GetFileMimeType(d.FileName)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (d DraftMessage) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, d.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (d DraftMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, d.Date)
}

// GetMessageEffect is a helper method for Client.GetMessageEffect
func (d DraftMessage) GetMessageEffect(client *Client) (*MessageEffect, error) {
	return client.GetMessageEffect(d.EffectId)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (d DraftMessage) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, d.Date)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (d DraftMessage) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, d.EffectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (d DraftMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, chatId, disableNotification, protectContent, d.EffectId, inputMessageContents, opts)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (e EmailAddressAuthenticationAppleId) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(e.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (e EmailAddressAuthenticationAppleId) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, e.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (e EmailAddressAuthenticationAppleId) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(e.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (e EmailAddressAuthenticationAppleId) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(e.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (e EmailAddressAuthenticationAppleId) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(e.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (e EmailAddressAuthenticationAppleId) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, e.Token)
}

// ApplyPremiumGiftCode is a helper method for Client.ApplyPremiumGiftCode
func (e EmailAddressAuthenticationCode) ApplyPremiumGiftCode(client *Client) (*Ok, error) {
	return client.ApplyPremiumGiftCode(e.Code)
}

// CheckAuthenticationCode is a helper method for Client.CheckAuthenticationCode
func (e EmailAddressAuthenticationCode) CheckAuthenticationCode(client *Client) (*Ok, error) {
	return client.CheckAuthenticationCode(e.Code)
}

// CheckEmailAddressVerificationCode is a helper method for Client.CheckEmailAddressVerificationCode
func (e EmailAddressAuthenticationCode) CheckEmailAddressVerificationCode(client *Client) (*Ok, error) {
	return client.CheckEmailAddressVerificationCode(e.Code)
}

// CheckPhoneNumberCode is a helper method for Client.CheckPhoneNumberCode
func (e EmailAddressAuthenticationCode) CheckPhoneNumberCode(client *Client) (*Ok, error) {
	return client.CheckPhoneNumberCode(e.Code)
}

// CheckPremiumGiftCode is a helper method for Client.CheckPremiumGiftCode
func (e EmailAddressAuthenticationCode) CheckPremiumGiftCode(client *Client) (*PremiumGiftCodeInfo, error) {
	return client.CheckPremiumGiftCode(e.Code)
}

// CheckRecoveryEmailAddressCode is a helper method for Client.CheckRecoveryEmailAddressCode
func (e EmailAddressAuthenticationCode) CheckRecoveryEmailAddressCode(client *Client) (*PasswordState, error) {
	return client.CheckRecoveryEmailAddressCode(e.Code)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (e EmailAddressAuthenticationGoogleId) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(e.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (e EmailAddressAuthenticationGoogleId) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, e.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (e EmailAddressAuthenticationGoogleId) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(e.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (e EmailAddressAuthenticationGoogleId) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(e.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (e EmailAddressAuthenticationGoogleId) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(e.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (e EmailAddressAuthenticationGoogleId) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, e.Token)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (e EmojiCategory) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, e.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (e EmojiCategory) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(e.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (e EmojiCategory) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(e.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (e EmojiCategory) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, e.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (e EmojiCategory) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, e.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (e EmojiCategory) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, e.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (e EmojiCategory) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, e.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (e EmojiCategory) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, e.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (e EmojiCategory) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, e.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (e EmojiCategory) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, e.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (e EmojiCategory) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(e.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (e EmojiCategory) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, e.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (e EmojiCategory) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, e.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (e EmojiCategory) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, e.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (e EmojiCategory) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, e.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (e EmojiCategory) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(e.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (e EmojiCategory) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(e.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (e EmojiCategory) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(e.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (e EmojiCategory) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(e.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (e EmojiCategory) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, e.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (e EmojiCategory) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(e.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (e EmojiCategory) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(e.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (e EmojiCategory) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, e.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (e EmojiCategory) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(e.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (e EmojiCategory) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, e.Name)
}

// SetOption is a helper method for Client.SetOption
func (e EmojiCategory) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(e.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (e EmojiCategory) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, e.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (e EmojiCategory) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, e.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (e EmojiCategory) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(e.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (e EmojiCategory) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, e.Name)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (e EmojiChatTheme) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, e.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (e EmojiChatTheme) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(e.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (e EmojiChatTheme) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(e.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (e EmojiChatTheme) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, e.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (e EmojiChatTheme) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, e.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (e EmojiChatTheme) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, e.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (e EmojiChatTheme) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, e.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (e EmojiChatTheme) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, e.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (e EmojiChatTheme) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, e.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (e EmojiChatTheme) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, e.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (e EmojiChatTheme) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(e.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (e EmojiChatTheme) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, e.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (e EmojiChatTheme) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, e.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (e EmojiChatTheme) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, e.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (e EmojiChatTheme) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, e.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (e EmojiChatTheme) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(e.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (e EmojiChatTheme) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(e.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (e EmojiChatTheme) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(e.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (e EmojiChatTheme) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(e.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (e EmojiChatTheme) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, e.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (e EmojiChatTheme) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(e.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (e EmojiChatTheme) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(e.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (e EmojiChatTheme) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, e.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (e EmojiChatTheme) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(e.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (e EmojiChatTheme) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, e.Name)
}

// SetOption is a helper method for Client.SetOption
func (e EmojiChatTheme) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(e.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (e EmojiChatTheme) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, e.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (e EmojiChatTheme) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, e.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (e EmojiChatTheme) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(e.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (e EmojiChatTheme) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, e.Name)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (e EmojiKeyword) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(e.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (e EmojiKeyword) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(e.Emoji)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (e EmojiReaction) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, e.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (e EmojiReaction) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, e.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (e EmojiReaction) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(e.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (e EmojiReaction) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, e.Title, startDate, isRtmpStream)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (e EmojiReaction) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(e.Emoji)
}

// Get is a helper method for Client.GetEmojiReaction
func (e EmojiReaction) Get(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(e.Emoji)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (e EmojiReaction) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(e.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (e EmojiReaction) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, e.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (e EmojiReaction) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, e.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (e EmojiReaction) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, e.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (e EmojiReaction) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, e.Title, recordVideo, usePortraitOrientation)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (e EmojiReaction) ToggleBotUsernameIsActive(client *Client, botUserId int64, username string) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, username, e.IsActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (e EmojiReaction) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, username string) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, username, e.IsActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (e EmojiReaction) ToggleUsernameIsActive(client *Client, username string) (*Ok, error) {
	return client.ToggleUsernameIsActive(username, e.IsActive)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (e EmojiStatus) CreateChatInviteLink(client *Client, chatId int64, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, e.ExpirationDate, memberLimit, createsJoinRequest)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (e EmojiStatus) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, e.ExpirationDate, memberLimit, createsJoinRequest)
}

// GetCustomEmojiStickers is a helper method for Client.GetCustomEmojiStickers
func (e EmojiStatusCustomEmojis) GetCustomEmojiStickers(client *Client) (*Stickers, error) {
	return client.GetCustomEmojiStickers(e.CustomEmojiIds)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (e EmojiStatusTypeCustomEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, name string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(name, e.CustomEmojiId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (e EmojiStatusTypeUpgradedGift) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, e.GiftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (e EmojiStatusTypeUpgradedGift) SendResoldGift(client *Client, ownerId *MessageSender, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(e.GiftName, ownerId, price)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (e EncryptedCredentials) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, e.Data, opts)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (e EncryptedCredentials) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, e.Data, unencryptedPrefixSize)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (e EncryptedCredentials) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, e.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (e EncryptedCredentials) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, e.Data)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (e EncryptedPassportElement) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, e.Data, opts)
}

// DeletePassportElement is a helper method for Client.DeletePassportElement
func (e EncryptedPassportElement) DeletePassportElement(client *Client) (*Ok, error) {
	return client.DeletePassportElement(e.Type)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (e EncryptedPassportElement) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, e.Data, unencryptedPrefixSize)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (e EncryptedPassportElement) GetPassportElement(client *Client, password string) (*PassportElement, error) {
	return client.GetPassportElement(e.Type, password)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (e EncryptedPassportElement) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, e.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (e EncryptedPassportElement) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, e.Data)
}

// GetCountryFlagEmoji is a helper method for Client.GetCountryFlagEmoji
func (f FactCheck) GetCountryFlagEmoji(client *Client) (*Text, error) {
	return client.GetCountryFlagEmoji(f.CountryCode)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (f FactCheck) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(f.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (f FactCheck) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(f.Text)
}

// GetPreferredCountryLanguage is a helper method for Client.GetPreferredCountryLanguage
func (f FactCheck) GetPreferredCountryLanguage(client *Client) (*Text, error) {
	return client.GetPreferredCountryLanguage(f.CountryCode)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (f FactCheck) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, f.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (f FactCheck) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(f.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (f FactCheck) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, f.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (f FactCheck) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(f.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (f FactCheck) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, f.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (f FactCheck) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, f.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (f FactCheck) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, f.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (f FactCheck) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(f.Text, toLanguageCode)
}

// AddChatMember is a helper method for Client.AddChatMember
func (f FailedToAddMember) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, f.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (f FailedToAddMember) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(f.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (f FailedToAddMember) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(f.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (f FailedToAddMember) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(f.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (f FailedToAddMember) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(f.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (f FailedToAddMember) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(f.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (f FailedToAddMember) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(f.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (f FailedToAddMember) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(f.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (f FailedToAddMember) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(f.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (f FailedToAddMember) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(f.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (f FailedToAddMember) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, f.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (f FailedToAddMember) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(f.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (f FailedToAddMember) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, f.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (f FailedToAddMember) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(f.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (f FailedToAddMember) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(f.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (f FailedToAddMember) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(f.UserId)
}

// GetUser is a helper method for Client.GetUser
func (f FailedToAddMember) GetUser(client *Client) (*User, error) {
	return client.GetUser(f.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (f FailedToAddMember) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, f.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (f FailedToAddMember) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(f.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (f FailedToAddMember) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(f.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (f FailedToAddMember) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(f.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (f FailedToAddMember) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(f.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (f FailedToAddMember) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(f.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (f FailedToAddMember) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, f.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (f FailedToAddMember) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, f.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (f FailedToAddMember) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, f.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (f FailedToAddMember) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(f.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (f FailedToAddMember) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(f.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (f FailedToAddMember) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(f.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (f FailedToAddMember) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, f.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (f FailedToAddMember) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, f.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (f FailedToAddMember) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(f.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (f FailedToAddMember) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(f.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (f FailedToAddMember) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(f.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (f FailedToAddMember) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(f.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (f FailedToAddMember) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(f.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (f FailedToAddMember) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(f.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (f FailedToAddMember) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(f.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (f FailedToAddMember) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(f.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (f FailedToAddMember) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(f.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (f FailedToAddMember) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(f.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (f FailedToAddMember) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, f.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (f FailedToAddMember) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(f.UserId, stickerFormat, sticker)
}

// AddToDownloads is a helper method for Client.AddFileToDownloads
func (f File) AddToDownloads(client *Client, chatId int64, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(f.Id, chatId, messageId, priority)
}

// AddProfileAudio is a helper method for Client.AddProfileAudio
func (f File) AddProfileAudio(client *Client) (*Ok, error) {
	return client.AddProfileAudio(f.Id)
}

// CancelDownload is a helper method for Client.CancelDownloadFile
func (f File) CancelDownload(client *Client, onlyIfPending bool) (*Ok, error) {
	return client.CancelDownloadFile(f.Id, onlyIfPending)
}

// CancelPreliminaryUpload is a helper method for Client.CancelPreliminaryUploadFile
func (f File) CancelPreliminaryUpload(client *Client) (*Ok, error) {
	return client.CancelPreliminaryUploadFile(f.Id)
}

// Delete is a helper method for Client.DeleteFile
func (f File) Delete(client *Client) (*Ok, error) {
	return client.DeleteFile(f.Id)
}

// Download is a helper method for Client.DownloadFile
func (f File) Download(client *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	return client.DownloadFile(f.Id, priority, offset, limit, synchronous)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (f File) EditBotMediaPreview(client *Client, botUserId int64, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(botUserId, languageCode, f.Id, content)
}

// GetAttachedStickerSets is a helper method for Client.GetAttachedStickerSets
func (f File) GetAttachedStickerSets(client *Client) (*StickerSets, error) {
	return client.GetAttachedStickerSets(f.Id)
}

// Get is a helper method for Client.GetFile
func (f File) Get(client *Client) (*File, error) {
	return client.GetFile(f.Id)
}

// GetDownloadedPrefixSize is a helper method for Client.GetFileDownloadedPrefixSize
func (f File) GetDownloadedPrefixSize(client *Client, offset int64) (*FileDownloadedPrefixSize, error) {
	return client.GetFileDownloadedPrefixSize(f.Id, offset)
}

// GetSuggestedName is a helper method for Client.GetSuggestedFileName
func (f File) GetSuggestedName(client *Client, directory string) (*Text, error) {
	return client.GetSuggestedFileName(f.Id, directory)
}

// IsProfileAudio is a helper method for Client.IsProfileAudio
func (f File) IsProfileAudio(client *Client) (*Ok, error) {
	return client.IsProfileAudio(f.Id)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (f File) OptimizeStorage(client *Client, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(f.Size, ttl, count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// ReadPart is a helper method for Client.ReadFilePart
func (f File) ReadPart(client *Client, offset int64, count int64) (*Data, error) {
	return client.ReadFilePart(f.Id, offset, count)
}

// RemoveFromDownloads is a helper method for Client.RemoveFileFromDownloads
func (f File) RemoveFromDownloads(client *Client, deleteFromCache bool) (*Ok, error) {
	return client.RemoveFileFromDownloads(f.Id, deleteFromCache)
}

// RemoveProfileAudio is a helper method for Client.RemoveProfileAudio
func (f File) RemoveProfileAudio(client *Client) (*Ok, error) {
	return client.RemoveProfileAudio(f.Id)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (f File) ReportChatPhoto(client *Client, chatId int64, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(chatId, f.Id, reason, text)
}

// SetGenerationProgress is a helper method for Client.SetFileGenerationProgress
func (f File) SetGenerationProgress(client *Client, generationId string, localPrefixSize int64) (*Ok, error) {
	return client.SetFileGenerationProgress(generationId, f.ExpectedSize, localPrefixSize)
}

// SetProfileAudioPosition is a helper method for Client.SetProfileAudioPosition
func (f File) SetProfileAudioPosition(client *Client, afterFileId int32) (*Ok, error) {
	return client.SetProfileAudioPosition(f.Id, afterFileId)
}

// ToggleDownloadIsPaused is a helper method for Client.ToggleDownloadIsPaused
func (f File) ToggleDownloadIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleDownloadIsPaused(f.Id, isPaused)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (f FileDownload) AddFileToDownloads(client *Client, chatId int64, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(f.FileId, chatId, messageId, priority)
}

// AddProfileAudio is a helper method for Client.AddProfileAudio
func (f FileDownload) AddProfileAudio(client *Client) (*Ok, error) {
	return client.AddProfileAudio(f.FileId)
}

// CancelDownloadFile is a helper method for Client.CancelDownloadFile
func (f FileDownload) CancelDownloadFile(client *Client, onlyIfPending bool) (*Ok, error) {
	return client.CancelDownloadFile(f.FileId, onlyIfPending)
}

// CancelPreliminaryUploadFile is a helper method for Client.CancelPreliminaryUploadFile
func (f FileDownload) CancelPreliminaryUploadFile(client *Client) (*Ok, error) {
	return client.CancelPreliminaryUploadFile(f.FileId)
}

// DeleteFile is a helper method for Client.DeleteFile
func (f FileDownload) DeleteFile(client *Client) (*Ok, error) {
	return client.DeleteFile(f.FileId)
}

// DownloadFile is a helper method for Client.DownloadFile
func (f FileDownload) DownloadFile(client *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	return client.DownloadFile(f.FileId, priority, offset, limit, synchronous)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (f FileDownload) EditBotMediaPreview(client *Client, botUserId int64, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(botUserId, languageCode, f.FileId, content)
}

// GetAttachedStickerSets is a helper method for Client.GetAttachedStickerSets
func (f FileDownload) GetAttachedStickerSets(client *Client) (*StickerSets, error) {
	return client.GetAttachedStickerSets(f.FileId)
}

// GetFile is a helper method for Client.GetFile
func (f FileDownload) GetFile(client *Client) (*File, error) {
	return client.GetFile(f.FileId)
}

// GetedPrefixSize is a helper method for Client.GetFileDownloadedPrefixSize
func (f FileDownload) GetedPrefixSize(client *Client, offset int64) (*FileDownloadedPrefixSize, error) {
	return client.GetFileDownloadedPrefixSize(f.FileId, offset)
}

// GetSuggestedFileName is a helper method for Client.GetSuggestedFileName
func (f FileDownload) GetSuggestedFileName(client *Client, directory string) (*Text, error) {
	return client.GetSuggestedFileName(f.FileId, directory)
}

// IsProfileAudio is a helper method for Client.IsProfileAudio
func (f FileDownload) IsProfileAudio(client *Client) (*Ok, error) {
	return client.IsProfileAudio(f.FileId)
}

// ReadFilePart is a helper method for Client.ReadFilePart
func (f FileDownload) ReadFilePart(client *Client, offset int64, count int64) (*Data, error) {
	return client.ReadFilePart(f.FileId, offset, count)
}

// RemoveFileFromDownloads is a helper method for Client.RemoveFileFromDownloads
func (f FileDownload) RemoveFileFromDownloads(client *Client, deleteFromCache bool) (*Ok, error) {
	return client.RemoveFileFromDownloads(f.FileId, deleteFromCache)
}

// RemoveProfileAudio is a helper method for Client.RemoveProfileAudio
func (f FileDownload) RemoveProfileAudio(client *Client) (*Ok, error) {
	return client.RemoveProfileAudio(f.FileId)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (f FileDownload) ReportChatPhoto(client *Client, chatId int64, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(chatId, f.FileId, reason, text)
}

// SetProfileAudioPosition is a helper method for Client.SetProfileAudioPosition
func (f FileDownload) SetProfileAudioPosition(client *Client, afterFileId int32) (*Ok, error) {
	return client.SetProfileAudioPosition(f.FileId, afterFileId)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (f FileDownload) ToggleBusinessConnectedBotChatIsPaused(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(chatId, f.IsPaused)
}

// ToggleDownloadIsPaused is a helper method for Client.ToggleDownloadIsPaused
func (f FileDownload) ToggleDownloadIsPaused(client *Client) (*Ok, error) {
	return client.ToggleDownloadIsPaused(f.FileId, f.IsPaused)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (f FileDownload) ToggleGroupCallScreenSharingIsPaused(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(groupCallId, f.IsPaused)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (f FileDownloadedPrefixSize) OptimizeStorage(client *Client, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(f.Size, ttl, count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (f FirebaseDeviceVerificationParametersPlayIntegrity) GetPassportAuthorizationForm(client *Client, botUserId int64, scope string, publicKey string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(botUserId, scope, publicKey, f.Nonce)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (f FormattedText) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, f.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (f FormattedText) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, f.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (f FormattedText) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(f.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (f FormattedText) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(f.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (f FormattedText) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(f.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (f FormattedText) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, f.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (f FormattedText) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, f.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (f FormattedText) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, f.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (f FormattedText) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(f.Text, inputLanguageCodes)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (f ForumTopic) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, messageId int64) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, messageId, f.IsPinned)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (f ForumTopic) SetChatNotificationSettings(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatNotificationSettings(chatId, f.NotificationSettings)
}

// SetNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (f ForumTopic) SetNotificationSettings(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(chatId, forumTopicId, f.NotificationSettings)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (f ForumTopic) ToggleChatIsPinned(client *Client, chatList *ChatList, chatId int64) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, chatId, f.IsPinned)
}

// ToggleIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (f ForumTopic) ToggleIsPinned(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(chatId, forumTopicId, f.IsPinned)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (f ForumTopic) ToggleSavedMessagesTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, f.IsPinned)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (f ForumTopicIcon) SetCustomEmojiStickerSetThumbnail(client *Client, name string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(name, f.CustomEmojiId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (f ForumTopicInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(f.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (f ForumTopicInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(f.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (f ForumTopicInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(f.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (f ForumTopicInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(f.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (f ForumTopicInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, f.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (f ForumTopicInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(f.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (f ForumTopicInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(f.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (f ForumTopicInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(f.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (f ForumTopicInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(f.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (f ForumTopicInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(f.ChatId)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (f ForumTopicInfo) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, f.Name, sticker)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (f ForumTopicInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (f ForumTopicInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(f.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (f ForumTopicInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(f.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (f ForumTopicInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(f.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (f ForumTopicInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(f.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (f ForumTopicInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(f.ChatId, username)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (f ForumTopicInfo) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(f.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (f ForumTopicInfo) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(f.Name)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (f ForumTopicInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(f.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (f ForumTopicInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(f.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (f ForumTopicInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(f.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (f ForumTopicInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(f.ChatId, messageId)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (f ForumTopicInfo) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, f.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (f ForumTopicInfo) CreateChatInviteLink(client *Client, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(f.ChatId, f.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (f ForumTopicInfo) CreateChatSubscriptionInviteLink(client *Client, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(f.ChatId, f.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (f ForumTopicInfo) CreateForumTopic(client *Client) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(f.ChatId, f.Name, f.IsNameImplicit, f.Icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (f ForumTopicInfo) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, f.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (f ForumTopicInfo) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, f.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (f ForumTopicInfo) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, f.Name, storyIds)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (f ForumTopicInfo) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(f.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (f ForumTopicInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(f.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (f ForumTopicInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(f.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (f ForumTopicInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(f.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (f ForumTopicInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(f.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (f ForumTopicInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(f.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (f ForumTopicInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(f.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (f ForumTopicInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(f.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (f ForumTopicInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(f.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (f ForumTopicInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(f.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (f ForumTopicInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(f.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (f ForumTopicInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(f.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (f ForumTopicInfo) DeleteForumTopic(client *Client) (*Ok, error) {
	return client.DeleteForumTopic(f.ChatId, f.ForumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (f ForumTopicInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(f.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (f ForumTopicInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(f.ChatId, inviteLink)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (f ForumTopicInfo) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(f.Name)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (f ForumTopicInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(f.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (f ForumTopicInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, f.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (f ForumTopicInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, f.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (f ForumTopicInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, f.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (f ForumTopicInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, f.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (f ForumTopicInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, f.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (f ForumTopicInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, f.ChatId, messageId, inputMessageContent, opts)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (f ForumTopicInfo) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, f.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (f ForumTopicInfo) EditChatInviteLink(client *Client, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(f.ChatId, inviteLink, f.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (f ForumTopicInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(f.ChatId, inviteLink, f.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (f ForumTopicInfo) EditForumTopic(client *Client, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(f.ChatId, f.ForumTopicId, f.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (f ForumTopicInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(f.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (f ForumTopicInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(f.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (f ForumTopicInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(f.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (f ForumTopicInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(f.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (f ForumTopicInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(f.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (f ForumTopicInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(f.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (f ForumTopicInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(f.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (f ForumTopicInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(f.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (f ForumTopicInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, f.ChatId, returnOnlyMainEmoji)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (f ForumTopicInfo) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(f.Name, typeField)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (f ForumTopicInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(f.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (f ForumTopicInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(f.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (f ForumTopicInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(f.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (f ForumTopicInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(f.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (f ForumTopicInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(f.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (f ForumTopicInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(f.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (f ForumTopicInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(f.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (f ForumTopicInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(f.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (f ForumTopicInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(f.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (f ForumTopicInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(f.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (f ForumTopicInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(f.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (f ForumTopicInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(f.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (f ForumTopicInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(f.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (f ForumTopicInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(f.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (f ForumTopicInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(f.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (f ForumTopicInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(f.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (f ForumTopicInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(f.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (f ForumTopicInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(f.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (f ForumTopicInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(f.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (f ForumTopicInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(f.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (f ForumTopicInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(f.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (f ForumTopicInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(f.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (f ForumTopicInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(f.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (f ForumTopicInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(f.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (f ForumTopicInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(f.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (f ForumTopicInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(f.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (f ForumTopicInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(f.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (f ForumTopicInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(f.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (f ForumTopicInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(f.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (f ForumTopicInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(f.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (f ForumTopicInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(f.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (f ForumTopicInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(f.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (f ForumTopicInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(f.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (f ForumTopicInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(f.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (f ForumTopicInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(f.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (f ForumTopicInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(f.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (f ForumTopicInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(f.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (f ForumTopicInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(f.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (f ForumTopicInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(f.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (f ForumTopicInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(f.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (f ForumTopicInfo) GetForumTopic(client *Client) (*ForumTopic, error) {
	return client.GetForumTopic(f.ChatId, f.ForumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (f ForumTopicInfo) GetForumTopicHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(f.ChatId, f.ForumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (f ForumTopicInfo) GetForumTopicLink(client *Client) (*MessageLink, error) {
	return client.GetForumTopicLink(f.ChatId, f.ForumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (f ForumTopicInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(f.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (f ForumTopicInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(f.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (f ForumTopicInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(f.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (f ForumTopicInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, f.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (f ForumTopicInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(f.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (f ForumTopicInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(f.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (f ForumTopicInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(f.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (f ForumTopicInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(f.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (f ForumTopicInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, f.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (f ForumTopicInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(f.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (f ForumTopicInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(f.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (f ForumTopicInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(f.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (f ForumTopicInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(f.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (f ForumTopicInfo) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(f.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (f ForumTopicInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(f.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (f ForumTopicInfo) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(f.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (f ForumTopicInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(f.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (f ForumTopicInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(f.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (f ForumTopicInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(f.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (f ForumTopicInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(f.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (f ForumTopicInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(f.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (f ForumTopicInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(f.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (f ForumTopicInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(f.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (f ForumTopicInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(f.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (f ForumTopicInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(f.ChatId, messageId)
}

// GetOption is a helper method for Client.GetOption
func (f ForumTopicInfo) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(f.Name)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (f ForumTopicInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(f.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (f ForumTopicInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(f.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (f ForumTopicInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(f.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (f ForumTopicInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(f.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (f ForumTopicInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, f.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (f ForumTopicInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(f.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (f ForumTopicInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(f.ChatId, storyId, isDark)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (f ForumTopicInfo) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(f.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (f ForumTopicInfo) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(f.Name)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (f ForumTopicInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(f.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (f ForumTopicInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(f.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (f ForumTopicInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(f.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (f ForumTopicInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(f.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (f ForumTopicInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(f.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (f ForumTopicInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(f.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (f ForumTopicInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(f.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (f ForumTopicInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(f.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (f ForumTopicInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(f.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (f ForumTopicInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(f.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (f ForumTopicInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(f.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (f ForumTopicInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(f.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (f ForumTopicInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(f.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (f ForumTopicInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(f.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (f ForumTopicInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(f.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (f ForumTopicInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(f.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (f ForumTopicInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(f.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (f ForumTopicInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(f.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (f ForumTopicInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(f.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (f ForumTopicInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(f.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (f ForumTopicInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(f.ChatId)
}
