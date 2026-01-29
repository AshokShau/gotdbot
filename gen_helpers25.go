package gotdbot

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpgradedGiftSymbol) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, u.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (u UpgradedGiftSymbol) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(u.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (u UpgradedGiftSymbol) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(u.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (u UpgradedGiftSymbol) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(u.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (u UpgradedGiftSymbol) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(u.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpgradedGiftSymbol) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, u.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (u UpgradedGiftSymbol) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(u.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (u UpgradedGiftSymbol) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(u.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (u UpgradedGiftSymbol) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, u.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (u UpgradedGiftSymbol) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(u.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpgradedGiftSymbol) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, u.Name)
}

// SetOption is a helper method for Client.SetOption
func (u UpgradedGiftSymbol) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(u.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpgradedGiftSymbol) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, u.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpgradedGiftSymbol) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, u.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpgradedGiftSymbol) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(u.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpgradedGiftSymbol) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, u.Name)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (u UpgradedGiftValueInfo) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(u.Currency, amount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (u UpgradedGiftValueInfo) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, u.Currency, amount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (u UpgradeGiftResult) DropGiftOriginalDetails(client *Client, starCount int64) (*Ok, error) {
	return client.DropGiftOriginalDetails(u.ReceivedGiftId, starCount)
}

// GetReceivedGift is a helper method for Client.GetReceivedGift
func (u UpgradeGiftResult) GetReceivedGift(client *Client) (*ReceivedGift, error) {
	return client.GetReceivedGift(u.ReceivedGiftId)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (u UpgradeGiftResult) GetUpgradedGiftWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(u.ReceivedGiftId, password)
}

// SellGift is a helper method for Client.SellGift
func (u UpgradeGiftResult) SellGift(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SellGift(businessConnectionId, u.ReceivedGiftId)
}

// SetGiftResalePrice is a helper method for Client.SetGiftResalePrice
func (u UpgradeGiftResult) SetGiftResalePrice(client *Client, opts *SetGiftResalePriceOpts) (*Ok, error) {
	return client.SetGiftResalePrice(u.ReceivedGiftId, opts)
}

// ToggleGiftIsSaved is a helper method for Client.ToggleGiftIsSaved
func (u UpgradeGiftResult) ToggleGiftIsSaved(client *Client) (*Ok, error) {
	return client.ToggleGiftIsSaved(u.ReceivedGiftId, u.IsSaved)
}

// TransferGift is a helper method for Client.TransferGift
func (u UpgradeGiftResult) TransferGift(client *Client, businessConnectionId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	return client.TransferGift(businessConnectionId, u.ReceivedGiftId, newOwnerId, starCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (u UpgradeGiftResult) UpgradeGift(client *Client, businessConnectionId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, u.ReceivedGiftId, keepOriginalDetails, starCount)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (u User) AddBotMediaPreview(client *Client, botUserId int64, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(botUserId, u.LanguageCode, content)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u User) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, u.Id, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (u User) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(u.Id, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u User) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(u.Id, name, sticker)
}

// AllowUnpaidMessagesFrom is a helper method for Client.AllowUnpaidMessagesFromUser
func (u User) AllowUnpaidMessagesFrom(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(u.Id, refundPayments)
}

// CanSendMessageTo is a helper method for Client.CanSendMessageToUser
func (u User) CanSendMessageTo(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(u.Id, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (u User) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(u.Id, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (u User) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(u.Id)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u User) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(u.Id, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (u User) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(u.Id, force)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (u User) DeleteBotMediaPreviews(client *Client, botUserId int64, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(botUserId, u.LanguageCode, fileIds)
}

// DeleteCommands is a helper method for Client.DeleteCommands
func (u User) DeleteCommands(client *Client, opts *DeleteCommandsOpts) (*Ok, error) {
	return client.DeleteCommands(u.LanguageCode, opts)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (u User) EditBotMediaPreview(client *Client, botUserId int64, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(botUserId, u.LanguageCode, fileId, content)
}

// EditStarSubscription is a helper method for Client.EditUserStarSubscription
func (u User) EditStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(u.Id, telegramPaymentChargeId, isCanceled)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (u User) GetBotInfoDescription(client *Client, botUserId int64) (*Text, error) {
	return client.GetBotInfoDescription(botUserId, u.LanguageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (u User) GetBotInfoShortDescription(client *Client, botUserId int64) (*Text, error) {
	return client.GetBotInfoShortDescription(botUserId, u.LanguageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (u User) GetBotMediaPreviewInfo(client *Client, botUserId int64) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(botUserId, u.LanguageCode)
}

// GetBotName is a helper method for Client.GetBotName
func (u User) GetBotName(client *Client, botUserId int64) (*Text, error) {
	return client.GetBotName(botUserId, u.LanguageCode)
}

// GetCommands is a helper method for Client.GetCommands
func (u User) GetCommands(client *Client, opts *GetCommandsOpts) (*BotCommands, error) {
	return client.GetCommands(u.LanguageCode, opts)
}

// GetEmojiSuggestionsUrl is a helper method for Client.GetEmojiSuggestionsUrl
func (u User) GetEmojiSuggestionsUrl(client *Client) (*HttpUrl, error) {
	return client.GetEmojiSuggestionsUrl(u.LanguageCode)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u User) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, u.Id)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (u User) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(u.Id, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (u User) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, u.Id)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (u User) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(u.Id)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (u User) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(u.Id)
}

// GetPhoneNumberInfoSync is a helper method for Client.GetPhoneNumberInfoSync
func (u User) GetPhoneNumberInfoSync(client *Client, phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	return client.GetPhoneNumberInfoSync(u.LanguageCode, phoneNumberPrefix)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (u User) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(u.Id)
}

// Get is a helper method for Client.GetUser
func (u User) Get(client *Client) (*User, error) {
	return client.GetUser(u.Id)
}

// GetChatBoosts is a helper method for Client.GetUserChatBoosts
func (u User) GetChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, u.Id)
}

// GetFullInfo is a helper method for Client.GetUserFullInfo
func (u User) GetFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(u.Id)
}

// GetProfileAudios is a helper method for Client.GetUserProfileAudios
func (u User) GetProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(u.Id, offset, limit)
}

// GetProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (u User) GetProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(u.Id, offset, limit)
}

// GetSupportInfo is a helper method for Client.GetUserSupportInfo
func (u User) GetSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(u.Id)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u User) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(u.Id, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u User) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, u.Id, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u User) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, u.Id, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u User) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, u.Id, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (u User) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(u.Id, telegramPaymentChargeId)
}

// Register is a helper method for Client.RegisterUser
func (u User) Register(client *Client, disableNotification bool) (*Ok, error) {
	return client.RegisterUser(u.FirstName, u.LastName, disableNotification)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (u User) ReorderBotMediaPreviews(client *Client, botUserId int64, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(botUserId, u.LanguageCode, fileIds)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u User) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(u.Id, name, oldSticker, newSticker)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u User) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, u.PaidMessageStarCount, opts)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (u User) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(u.Id, result, chatTypes)
}

// SearchByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (u User) SearchByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(u.PhoneNumber, onlyLocal)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u User) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, u.PaidMessageStarCount)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u User) SendGroupCallMessage(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, u.PaidMessageStarCount)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (u User) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(u.PhoneNumber, typeField, opts)
}

// SetAccentColor is a helper method for Client.SetAccentColor
func (u User) SetAccentColor(client *Client) (*Ok, error) {
	return client.SetAccentColor(u.AccentColorId, u.BackgroundCustomEmojiId)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (u User) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(u.PhoneNumber, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (u User) SetBotInfoDescription(client *Client, botUserId int64, description string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, u.LanguageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (u User) SetBotInfoShortDescription(client *Client, botUserId int64, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(botUserId, u.LanguageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (u User) SetBotName(client *Client, botUserId int64, name string) (*Ok, error) {
	return client.SetBotName(botUserId, u.LanguageCode, name)
}

// SetBusinessAccountName is a helper method for Client.SetBusinessAccountName
func (u User) SetBusinessAccountName(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountName(businessConnectionId, u.FirstName, u.LastName)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u User) SetChatAccentColor(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatAccentColor(chatId, u.AccentColorId, u.BackgroundCustomEmojiId)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u User) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, u.PaidMessageStarCount)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u User) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, u.PaidMessageStarCount)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u User) SetChatProfileAccentColor(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatProfileAccentColor(chatId, u.ProfileAccentColorId, u.ProfileBackgroundCustomEmojiId)
}

// SetCommands is a helper method for Client.SetCommands
func (u User) SetCommands(client *Client, commands []*BotCommand, opts *SetCommandsOpts) (*Ok, error) {
	return client.SetCommands(u.LanguageCode, commands, opts)
}

// SetEmojiStatus is a helper method for Client.SetEmojiStatus
func (u User) SetEmojiStatus(client *Client) (*Ok, error) {
	return client.SetEmojiStatus(u.EmojiStatus)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u User) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, u.Id, score, force)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u User) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, u.PaidMessageStarCount)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (u User) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, u.Id, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (u User) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(u.Id, menuButton)
}

// SetName is a helper method for Client.SetName
func (u User) SetName(client *Client) (*Ok, error) {
	return client.SetName(u.FirstName, u.LastName)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (u User) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(u.Id, errors)
}

// SetProfileAccentColor is a helper method for Client.SetProfileAccentColor
func (u User) SetProfileAccentColor(client *Client) (*Ok, error) {
	return client.SetProfileAccentColor(u.ProfileAccentColorId, u.ProfileBackgroundCustomEmojiId)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u User) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(u.Id, name, opts)
}

// SetEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (u User) SetEmojiStatus(client *Client) (*Ok, error) {
	return client.SetUserEmojiStatus(u.Id, u.EmojiStatus)
}

// SetNote is a helper method for Client.SetUserNote
func (u User) SetNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(u.Id, note)
}

// SetPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (u User) SetPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(u.Id, photo)
}

// SetSupportInfo is a helper method for Client.SetUserSupportInfo
func (u User) SetSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(u.Id, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (u User) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(u.Id)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u User) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, isRtmpStream, enableMessages, u.PaidMessageStarCount)
}

// SuggestBirthdate is a helper method for Client.SuggestUserBirthdate
func (u User) SuggestBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(u.Id, birthdate)
}

// SuggestProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (u User) SuggestProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(u.Id, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u User) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, u.Id, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (u User) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(u.Id, stickerFormat, sticker)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (u UserAuctionBid) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UserAuctionBid) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, u.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UserAuctionBid) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, u.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (u UserAuctionBid) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(u.OwnerId, prepaidUpgradeHash, u.StarCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UserAuctionBid) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(u.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (u UserAuctionBid) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(u.OwnerId, collectionId)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (u UserAuctionBid) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, u.StarCount)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (u UserAuctionBid) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(u.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (u UserAuctionBid) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, u.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (u UserAuctionBid) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(u.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (u UserAuctionBid) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(u.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (u UserAuctionBid) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(u.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (u UserAuctionBid) GetStarWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(u.OwnerId, u.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UserAuctionBid) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, u.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (u UserAuctionBid) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, u.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (u UserAuctionBid) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, u.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UserAuctionBid) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, u.StarCount, userId, text, isPrivate)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (u UserAuctionBid) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (u UserAuctionBid) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (u UserAuctionBid) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(u.OwnerId, collectionIds)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (u UserAuctionBid) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, u.StarCount)
}

// SendGift is a helper method for Client.SendGift
func (u UserAuctionBid) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, u.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u UserAuctionBid) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(u.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (u UserAuctionBid) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, u.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UserAuctionBid) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(u.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (u UserAuctionBid) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(u.OwnerId, receivedGiftIds)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (u UserAuctionBid) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, u.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (u UserAuctionBid) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, u.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (u UserAuctionBid) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, u.StarCount)
}

// GetBlockedMessageSenders is a helper method for Client.GetBlockedMessageSenders
func (u UserFullInfo) GetBlockedMessageSenders(client *Client, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetBlockedMessageSenders(u.BlockList, offset, limit)
}

// SetBirthdate is a helper method for Client.SetBirthdate
func (u UserFullInfo) SetBirthdate(client *Client) (*Ok, error) {
	return client.SetBirthdate(u.Birthdate)
}

// SetMainProfileTab is a helper method for Client.SetMainProfileTab
func (u UserFullInfo) SetMainProfileTab(client *Client) (*Ok, error) {
	return client.SetMainProfileTab(u.MainProfileTab)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (u UserFullInfo) SetSupergroupMainProfileTab(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(supergroupId, u.MainProfileTab)
}

// SetUserNote is a helper method for Client.SetUserNote
func (u UserFullInfo) SetUserNote(client *Client, userId int64) (*Ok, error) {
	return client.SetUserNote(userId, u.Note)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (u UserFullInfo) SuggestUserBirthdate(client *Client, userId int64) (*Ok, error) {
	return client.SuggestUserBirthdate(userId, u.Birthdate)
}

// ToggleHasSponsoredMessagesEnabled is a helper method for Client.ToggleHasSponsoredMessagesEnabled
func (u UserFullInfo) ToggleHasSponsoredMessagesEnabled(client *Client) (*Ok, error) {
	return client.ToggleHasSponsoredMessagesEnabled(u.HasSponsoredMessagesEnabled)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (u UserLink) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, u.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (u UserLink) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, u.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (u UserLink) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, u.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (u UserLink) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(u.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (u UserLink) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(u.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (u UserLink) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, u.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (u UserLink) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(u.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UserLink) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, u.Url, parameters, opts)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (u UserPrivacySettingRuleAllowChatMembers) AddChatFolderByInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.AddChatFolderByInviteLink(inviteLink, u.ChatIds)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UserPrivacySettingRuleAllowChatMembers) CreateChatFolderInviteLink(client *Client, chatFolderId int32, name string) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, name, u.ChatIds)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UserPrivacySettingRuleAllowChatMembers) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, name string) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, name, u.ChatIds)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (u UserPrivacySettingRuleAllowChatMembers) OptimizeStorage(client *Client, size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, count, immunityDelay, fileTypes, u.ChatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (u UserPrivacySettingRuleAllowChatMembers) SetPinnedChats(client *Client, chatList *ChatList) (*Ok, error) {
	return client.SetPinnedChats(chatList, u.ChatIds)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UserPrivacySettingRuleAllowUsers) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, u.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (u UserPrivacySettingRuleAllowUsers) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(u.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UserPrivacySettingRuleAllowUsers) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, u.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UserPrivacySettingRuleAllowUsers) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, u.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (u UserPrivacySettingRuleAllowUsers) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(u.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (u UserPrivacySettingRuleAllowUsers) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(u.UserIds)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (u UserPrivacySettingRuleRestrictChatMembers) AddChatFolderByInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.AddChatFolderByInviteLink(inviteLink, u.ChatIds)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UserPrivacySettingRuleRestrictChatMembers) CreateChatFolderInviteLink(client *Client, chatFolderId int32, name string) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, name, u.ChatIds)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UserPrivacySettingRuleRestrictChatMembers) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, name string) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, name, u.ChatIds)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (u UserPrivacySettingRuleRestrictChatMembers) OptimizeStorage(client *Client, size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, count, immunityDelay, fileTypes, u.ChatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (u UserPrivacySettingRuleRestrictChatMembers) SetPinnedChats(client *Client, chatList *ChatList) (*Ok, error) {
	return client.SetPinnedChats(chatList, u.ChatIds)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UserPrivacySettingRuleRestrictUsers) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, u.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (u UserPrivacySettingRuleRestrictUsers) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(u.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UserPrivacySettingRuleRestrictUsers) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, u.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UserPrivacySettingRuleRestrictUsers) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, u.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (u UserPrivacySettingRuleRestrictUsers) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(u.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (u UserPrivacySettingRuleRestrictUsers) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(u.UserIds)
}

// GetChatBoostLevelFeatures is a helper method for Client.GetChatBoostLevelFeatures
func (u UserRating) GetChatBoostLevelFeatures(client *Client, isChannel bool) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(isChannel, u.Level)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u Users) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, u.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (u Users) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(u.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u Users) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, u.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u Users) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, u.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (u Users) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(u.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (u Users) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(u.UserIds)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UserSupportInfo) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, u.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UserSupportInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, u.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (u UserSupportInfo) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, u.Date)
}

// Set is a helper method for Client.SetUserSupportInfo
func (u UserSupportInfo) Set(client *Client, userId int64) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(userId, u.Message)
}

// AnswerShippingQuery is a helper method for Client.AnswerShippingQuery
func (v ValidatedOrderInfo) AnswerShippingQuery(client *Client, shippingQueryId string, errorMessage string) (*Ok, error) {
	return client.AnswerShippingQuery(shippingQueryId, v.ShippingOptions, errorMessage)
}

// SendPaymentForm is a helper method for Client.SendPaymentForm
func (v ValidatedOrderInfo) SendPaymentForm(client *Client, inputInvoice *InputInvoice, paymentFormId string, shippingOptionId string, tipAmount int64, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	return client.SendPaymentForm(inputInvoice, paymentFormId, v.OrderInfoId, shippingOptionId, tipAmount, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (v Venue) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, v.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (v Venue) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, v.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (v Venue) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(v.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (v Venue) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, v.Title, startDate, isRtmpStream)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (v Venue) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(v.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (v Venue) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(v.Location, zoom, width, height, scale, chatId)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (v Venue) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(v.Title)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (v Venue) SaveApplicationLogEvent(client *Client, chatId int64, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(v.Type, chatId, data)
}

// SearchPublicStoriesBy is a helper method for Client.SearchPublicStoriesByVenue
func (v Venue) SearchPublicStoriesBy(client *Client, venueProvider string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByVenue(venueProvider, v.Id, offset, limit)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (v Venue) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, v.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (v Venue) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, v.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (v Venue) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, v.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (v Venue) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, v.Title, recordVideo, usePortraitOrientation)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (v Video) CheckWebAppFileDownload(client *Client, botUserId int64, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, v.FileName, url)
}

// CleanFileName is a helper method for Client.CleanFileName
func (v Video) CleanFileName(client *Client) (*Text, error) {
	return client.CleanFileName(v.FileName)
}

// DiscardCall is a helper method for Client.DiscardCall
func (v Video) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, v.Duration, isVideo, connectionId)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (v Video) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(v.MimeType)
}

// GetFileMimeType is a helper method for Client.GetFileMimeType
func (v Video) GetFileMimeType(client *Client) (*Text, error) {
	return client.GetFileMimeType(v.FileName)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (v Video) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, v.Width, v.Height, scale, chatId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (v Video) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, v.Duration, paidMessageStarCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (v VideoChat) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(v.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (v VideoChat) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(v.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (v VideoChat) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(v.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (v VideoChat) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(v.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (v VideoChat) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(v.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (v VideoChat) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(v.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (v VideoChat) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(v.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (v VideoChat) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(v.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (v VideoChat) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(v.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (v VideoChat) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(v.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (v VideoChat) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(v.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (v VideoChat) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(v.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (v VideoChat) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(v.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (v VideoChat) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(v.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (v VideoChat) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(v.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (v VideoChat) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(v.GroupCallId)
}

// GetInviteLink is a helper method for Client.GetVideoChatInviteLink
func (v VideoChat) GetInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(v.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (v VideoChat) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(v.GroupCallId, userId, isVideo)
}

// InviteParticipants is a helper method for Client.InviteVideoChatParticipants
func (v VideoChat) InviteParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(v.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (v VideoChat) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(v.GroupCallId, joinParameters)
}

// Join is a helper method for Client.JoinVideoChat
func (v VideoChat) Join(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(v.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (v VideoChat) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(v.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (v VideoChat) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(v.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (v VideoChat) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(v.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (v VideoChat) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(v.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (v VideoChat) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(v.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (v VideoChat) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(v.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (v VideoChat) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(v.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (v VideoChat) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(v.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (v VideoChat) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(v.GroupCallId, messageSenderId)
}

// SetDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (v VideoChat) SetDefaultParticipant(client *Client, chatId int64) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(chatId, v.DefaultParticipantId)
}

// SetTitle is a helper method for Client.SetVideoChatTitle
func (v VideoChat) SetTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(v.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (v VideoChat) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(v.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (v VideoChat) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(v.GroupCallId, audioSourceId, payload)
}

// StartScheduled is a helper method for Client.StartScheduledVideoChat
func (v VideoChat) StartScheduled(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(v.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (v VideoChat) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(v.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (v VideoChat) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(v.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (v VideoChat) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(v.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (v VideoChat) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(v.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (v VideoChat) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(v.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (v VideoChat) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(v.GroupCallId, isPaused)
}

// ToggleEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (v VideoChat) ToggleEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(v.GroupCallId, enabledStartNotification)
}

// ToggleMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (v VideoChat) ToggleMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(v.GroupCallId, muteNewParticipants)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (v VideoMessageAdvertisement) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, v.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (v VideoMessageAdvertisement) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, v.Text, showAlert, url, cacheTime)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (v VideoMessageAdvertisement) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, v.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (v VideoMessageAdvertisement) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, v.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (v VideoMessageAdvertisement) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(v.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (v VideoMessageAdvertisement) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, v.Title, startDate, isRtmpStream)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (v VideoMessageAdvertisement) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(v.Text, inputLanguageCodes)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (v VideoMessageAdvertisement) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(v.Title)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (v VideoMessageAdvertisement) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(v.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (v VideoMessageAdvertisement) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(v.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (v VideoMessageAdvertisement) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, v.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (v VideoMessageAdvertisement) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, v.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (v VideoMessageAdvertisement) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, v.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (v VideoMessageAdvertisement) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(v.Text, inputLanguageCodes)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (v VideoMessageAdvertisement) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, v.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (v VideoMessageAdvertisement) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, v.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (v VideoMessageAdvertisement) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, v.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (v VideoMessageAdvertisement) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, v.Title, recordVideo, usePortraitOrientation)
}

// DiscardCall is a helper method for Client.DiscardCall
func (v VideoNote) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, v.Duration, isVideo, connectionId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (v VideoNote) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, v.Duration, paidMessageStarCount)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (v VideoStoryboard) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, v.Width, v.Height, scale, chatId)
}

// DiscardCall is a helper method for Client.DiscardCall
func (v VoiceNote) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, v.Duration, isVideo, connectionId)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (v VoiceNote) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(v.MimeType)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (v VoiceNote) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, v.Duration, paidMessageStarCount)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (w WebApp) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, w.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (w WebApp) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, w.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (w WebApp) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(w.Title, isForum, isChannel, w.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (w WebApp) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, w.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (w WebApp) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(w.Title)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (w WebApp) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, w.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (w WebApp) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, w.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (w WebApp) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, w.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (w WebApp) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, w.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (w WebApp) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, w.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (w WebApp) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, w.Title, recordVideo, usePortraitOrientation)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (w WebAppInfo) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, w.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (w WebAppInfo) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, w.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (w WebAppInfo) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, w.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (w WebAppInfo) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(w.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (w WebAppInfo) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(w.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (w WebAppInfo) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, w.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (w WebAppInfo) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(w.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (w WebAppInfo) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, w.Url, parameters, opts)
}

// GetThemeParametersJsonString is a helper method for Client.GetThemeParametersJsonString
func (w WebAppOpenParameters) GetThemeParametersJsonString(client *Client) (*Text, error) {
	return client.GetThemeParametersJsonString(w.Theme)
}
