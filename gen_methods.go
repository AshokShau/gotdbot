package gotdbot

// AcceptCall Accepts an incoming call @call_id Call identifier @protocol The call protocols supported by the application
func (c *Client) AcceptCall(callId int32, protocol *CallProtocol) (*Ok, error) {
	req := &AcceptCall{
		CallId:   callId,
		Protocol: protocol,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AcceptTermsOfService Accepts Telegram terms of services @terms_of_service_id Terms of service identifier
func (c *Client) AcceptTermsOfService(termsOfServiceId string) (*Ok, error) {
	req := &AcceptTermsOfService{
		TermsOfServiceId: termsOfServiceId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ActivateStoryStealthMode Activates stealth mode for stories, which hides all views of stories from the current user in the last "story_stealth_mode_past_period" seconds
func (c *Client) ActivateStoryStealthMode() (*Ok, error) {
	req := &ActivateStoryStealthMode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddBotMediaPreview Adds a new media preview to the beginning of the list of media previews of a bot. Returns the added preview after addition is completed server-side. The total number of previews must not exceed getOption("bot_media_preview_count_max") for the given language
func (c *Client) AddBotMediaPreview(botUserId int64, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	req := &AddBotMediaPreview{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Content:      content,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotMediaPreview), nil
}

// AddChatFolderByInviteLink Adds a chat folder by an invite link @invite_link Invite link for the chat folder @chat_ids Identifiers of the chats added to the chat folder. The chats are automatically joined if they aren't joined yet
func (c *Client) AddChatFolderByInviteLink(inviteLink string, chatIds []int64) (*Ok, error) {
	req := &AddChatFolderByInviteLink{
		InviteLink: inviteLink,
		ChatIds:    chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddChatMember Adds a new member to a chat; requires can_invite_users member right. Members can't be added to private or secret chats. Returns information about members that weren't added
func (c *Client) AddChatMember(chatId int64, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	req := &AddChatMember{
		ChatId:       chatId,
		UserId:       userId,
		ForwardLimit: forwardLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FailedToAddMembers), nil
}

// AddChatMembers Adds multiple new members to a chat; requires can_invite_users member right. Currently, this method is only available for supergroups and channels.
func (c *Client) AddChatMembers(chatId int64, userIds []int64) (*FailedToAddMembers, error) {
	req := &AddChatMembers{
		ChatId:  chatId,
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FailedToAddMembers), nil
}

// AddChatToList Adds a chat to a chat list. A chat can't be simultaneously in Main and Archive chat lists, so it is automatically removed from another one if needed
func (c *Client) AddChatToList(chatId int64, chatList *ChatList) (*Ok, error) {
	req := &AddChatToList{
		ChatId:   chatId,
		ChatList: chatList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddChecklistTasks Adds tasks to a checklist in a message
func (c *Client) AddChecklistTasks(chatId int64, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	req := &AddChecklistTasks{
		ChatId:    chatId,
		MessageId: messageId,
		Tasks:     tasks,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddContact Adds a user to the contact list or edits an existing contact by their user identifier
func (c *Client) AddContact(userId int64, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	req := &AddContact{
		UserId:           userId,
		Contact:          contact,
		SharePhoneNumber: sharePhoneNumber,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddCustomServerLanguagePack Adds a custom server language pack to the list of installed language packs in current localization target. Can be called before authorization @language_pack_id Identifier of a language pack to be added
func (c *Client) AddCustomServerLanguagePack(languagePackId string) (*Ok, error) {
	req := &AddCustomServerLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first.
func (c *Client) AddFavoriteSticker(sticker *InputFile) (*Ok, error) {
	req := &AddFavoriteSticker{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddFileToDownloads Adds a file from a message to the list of file downloads. Download progress and completion of the download will be notified through updateFile updates.
func (c *Client) AddFileToDownloads(fileId int32, chatId int64, messageId int64, priority int32) (*File, error) {
	req := &AddFileToDownloads{
		FileId:    fileId,
		ChatId:    chatId,
		MessageId: messageId,
		Priority:  priority,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// AddGiftCollectionGifts Adds gifts to the beginning of a previously created collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) AddGiftCollectionGifts(ownerId *MessageSender, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	req := &AddGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollection), nil
}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
func (c *Client) AddLocalMessage(chatId int64, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	req := &AddLocalMessage{
		ChatId:              chatId,
		SenderId:            senderId,
		DisableNotification: disableNotification,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// AddLoginPasskey Adds a passkey allowed to be used for the login by the current user and returns the added passkey. Call getPasskeyParameters to get parameters for creating of the passkey
func (c *Client) AddLoginPasskey(clientData string, attestationObject string) (*Passkey, error) {
	req := &AddLoginPasskey{
		ClientData:        clientData,
		AttestationObject: attestationObject,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Passkey), nil
}

// AddLogMessage Adds a message to TDLib internal log. Can be called synchronously
func (c *Client) AddLogMessage(verbosityLevel int32, text string) (*Ok, error) {
	req := &AddLogMessage{
		VerbosityLevel: verbosityLevel,
		Text:           text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddMessageReaction Adds a reaction or a tag to a message. Use getMessageAvailableReactions to receive the list of available reactions for the message
func (c *Client) AddMessageReaction(chatId int64, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	req := &AddMessageReaction{
		ChatId:                chatId,
		MessageId:             messageId,
		ReactionType:          reactionType,
		IsBig:                 isBig,
		UpdateRecentReactions: updateRecentReactions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization @entry The network statistics entry with the data to be added to statistics
func (c *Client) AddNetworkStatistics(entry *NetworkStatisticsEntry) (*Ok, error) {
	req := &AddNetworkStatistics{
		Entry: entry,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddOffer Sends a suggested post based on a previously sent message in a channel direct messages chat. Can be also used to suggest price or time change for an existing suggested post.
func (c *Client) AddOffer(chatId int64, messageId int64, options *MessageSendOptions) (*Message, error) {
	req := &AddOffer{
		ChatId:    chatId,
		MessageId: messageId,
		Options:   options,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// AddPendingLiveStoryReaction Adds pending paid reaction in a live story group call. Can't be used in live stories posted by the current user.
func (c *Client) AddPendingLiveStoryReaction(groupCallId int32, starCount int64) (*Ok, error) {
	req := &AddPendingLiveStoryReaction{
		GroupCallId: groupCallId,
		StarCount:   starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddPendingPaidMessageReaction Adds the paid message reaction to a message. Use getMessageAvailableReactions to check whether the reaction is available for the message
func (c *Client) AddPendingPaidMessageReaction(chatId int64, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	req := &AddPendingPaidMessageReaction{
		ChatId:    chatId,
		MessageId: messageId,
		StarCount: starCount,
	}
	if opts != nil {
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddProfileAudio Adds an audio file to the beginning of the profile audio files of the current user
func (c *Client) AddProfileAudio(fileId int32) (*Ok, error) {
	req := &AddProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
func (c *Client) AddProxy(server string, port int32, enable bool, typeField *ProxyType) (*Proxy, error) {
	req := &AddProxy{
		Server:    server,
		Port:      port,
		Enable:    enable,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Proxy), nil
}

// AddQuickReplyShortcutInlineQueryResultMessage Adds a message to a quick reply shortcut via inline bot. If shortcut doesn't exist and there are less than getOption("quick_reply_shortcut_count_max") shortcuts, then a new shortcut is created.
func (c *Client) AddQuickReplyShortcutInlineQueryResultMessage(shortcutName string, replyToMessageId int64, queryId string, resultId string, hideViaBot bool) (*QuickReplyMessage, error) {
	req := &AddQuickReplyShortcutInlineQueryResultMessage{
		ShortcutName:     shortcutName,
		ReplyToMessageId: replyToMessageId,
		QueryId:          queryId,
		ResultId:         resultId,
		HideViaBot:       hideViaBot,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*QuickReplyMessage), nil
}

// AddQuickReplyShortcutMessage Adds a message to a quick reply shortcut. If shortcut doesn't exist and there are less than getOption("quick_reply_shortcut_count_max") shortcuts, then a new shortcut is created.
func (c *Client) AddQuickReplyShortcutMessage(shortcutName string, replyToMessageId int64, inputMessageContent *InputMessageContent) (*QuickReplyMessage, error) {
	req := &AddQuickReplyShortcutMessage{
		ShortcutName:        shortcutName,
		ReplyToMessageId:    replyToMessageId,
		InputMessageContent: inputMessageContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*QuickReplyMessage), nil
}

// AddQuickReplyShortcutMessageAlbum Adds 2-10 messages grouped together into an album to a quick reply shortcut. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) AddQuickReplyShortcutMessageAlbum(shortcutName string, replyToMessageId int64, inputMessageContents []*InputMessageContent) (*QuickReplyMessages, error) {
	req := &AddQuickReplyShortcutMessageAlbum{
		ShortcutName:         shortcutName,
		ReplyToMessageId:     replyToMessageId,
		InputMessageContents: inputMessageContents,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*QuickReplyMessages), nil
}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first @chat_id Identifier of the chat to add
func (c *Client) AddRecentlyFoundChat(chatId int64) (*Ok, error) {
	req := &AddRecentlyFoundChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddRecentSticker Manually adds a new sticker to the list of recently used stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first.
func (c *Client) AddRecentSticker(isAttached bool, sticker *InputFile) (*Stickers, error) {
	req := &AddRecentSticker{
		IsAttached: isAttached,
		Sticker:    sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first.
func (c *Client) AddSavedAnimation(animation *InputFile) (*Ok, error) {
	req := &AddSavedAnimation{
		Animation: animation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddSavedNotificationSound Adds a new notification sound to the list of saved notification sounds. The new notification sound is added to the top of the list. If it is already in the list, its position isn't changed @sound Notification sound file to add
func (c *Client) AddSavedNotificationSound(sound *InputFile) (*NotificationSound, error) {
	req := &AddSavedNotificationSound{
		Sound: sound,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*NotificationSound), nil
}

// AddStickerToSet Adds a new sticker to a set
func (c *Client) AddStickerToSet(userId int64, name string, sticker *InputSticker) (*Ok, error) {
	req := &AddStickerToSet{
		UserId:  userId,
		Name:    name,
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AddStoryAlbumStories Adds stories to the beginning of a previously created story album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) AddStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	req := &AddStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbum), nil
}

// AllowBotToSendMessages Allows the specified bot to send messages to the user @bot_user_id Identifier of the target bot
func (c *Client) AllowBotToSendMessages(botUserId int64) (*Ok, error) {
	req := &AllowBotToSendMessages{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AllowUnpaidMessagesFromUser Allows the specified user to send unpaid private messages to the current user by adding a rule to userPrivacySettingAllowUnpaidMessages
func (c *Client) AllowUnpaidMessagesFromUser(userId int64, refundPayments bool) (*Ok, error) {
	req := &AllowUnpaidMessagesFromUser{
		UserId:         userId,
		RefundPayments: refundPayments,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
func (c *Client) AnswerCallbackQuery(callbackQueryId string, text string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	req := &AnswerCallbackQuery{
		CallbackQueryId: callbackQueryId,
		Text:            text,
		ShowAlert:       showAlert,
		Url:             url,
		CacheTime:       cacheTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerCustomQuery Answers a custom query; for bots only @custom_query_id Identifier of a custom query @data JSON-serialized answer to the query
func (c *Client) AnswerCustomQuery(customQueryId string, data string) (*Ok, error) {
	req := &AnswerCustomQuery{
		CustomQueryId: customQueryId,
		Data:          data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerInlineQuery Sets the result of an inline query; for bots only
func (c *Client) AnswerInlineQuery(inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, nextOffset string, opts *AnswerInlineQueryOpts) (*Ok, error) {
	req := &AnswerInlineQuery{
		InlineQueryId: inlineQueryId,
		IsPersonal:    isPersonal,
		Results:       results,
		CacheTime:     cacheTime,
		NextOffset:    nextOffset,
	}
	if opts != nil {
		req.Button = opts.Button
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only @pre_checkout_query_id Identifier of the pre-checkout query @error_message An error message, empty on success
func (c *Client) AnswerPreCheckoutQuery(preCheckoutQueryId string, errorMessage string) (*Ok, error) {
	req := &AnswerPreCheckoutQuery{
		PreCheckoutQueryId: preCheckoutQueryId,
		ErrorMessage:       errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerShippingQuery Sets the result of a shipping query; for bots only @shipping_query_id Identifier of the shipping query @shipping_options Available shipping options @error_message An error message, empty on success
func (c *Client) AnswerShippingQuery(shippingQueryId string, shippingOptions []*ShippingOption, errorMessage string) (*Ok, error) {
	req := &AnswerShippingQuery{
		ShippingQueryId: shippingQueryId,
		ShippingOptions: shippingOptions,
		ErrorMessage:    errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AnswerWebAppQuery Sets the result of interaction with a Web App and sends corresponding message on behalf of the user to the chat from which the query originated; for bots only
func (c *Client) AnswerWebAppQuery(webAppQueryId string, result *InputInlineQueryResult) (*SentWebAppMessage, error) {
	req := &AnswerWebAppQuery{
		WebAppQueryId: webAppQueryId,
		Result:        result,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SentWebAppMessage), nil
}

// ApplyPremiumGiftCode Applies a Telegram Premium gift code @code The code to apply
func (c *Client) ApplyPremiumGiftCode(code string) (*Ok, error) {
	req := &ApplyPremiumGiftCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ApproveSuggestedPost Approves a suggested post in a channel direct messages chat
func (c *Client) ApproveSuggestedPost(chatId int64, messageId int64, sendDate int32) (*Ok, error) {
	req := &ApproveSuggestedPost{
		ChatId:    chatId,
		MessageId: messageId,
		SendDate:  sendDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// AssignStoreTransaction Informs server about an in-store purchase. For official applications only @transaction Information about the transaction @purpose Transaction purpose
func (c *Client) AssignStoreTransaction(transaction *StoreTransaction, purpose *StorePaymentPurpose) (*Ok, error) {
	req := &AssignStoreTransaction{
		Transaction: transaction,
		Purpose:     purpose,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// BanChatMember Bans a member in a chat; requires can_restrict_members administrator right. Members can't be banned in private or secret chats. In supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first
func (c *Client) BanChatMember(chatId int64, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	req := &BanChatMember{
		ChatId:          chatId,
		MemberId:        memberId,
		BannedUntilDate: bannedUntilDate,
		RevokeMessages:  revokeMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// BanGroupCallParticipants Bans users from a group call not bound to a chat; requires groupCall.is_owned. Only the owner of the group call can invite the banned users back
func (c *Client) BanGroupCallParticipants(groupCallId int32, userIds []string) (*Ok, error) {
	req := &BanGroupCallParticipants{
		GroupCallId: groupCallId,
		UserIds:     userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// BlockMessageSenderFromReplies Blocks an original sender of a message in the Replies chat
func (c *Client) BlockMessageSenderFromReplies(messageId int64, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	req := &BlockMessageSenderFromReplies{
		MessageId:         messageId,
		DeleteMessage:     deleteMessage,
		DeleteAllMessages: deleteAllMessages,
		ReportSpam:        reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// BoostChat Boosts a chat and returns the list of available chat boost slots for the current user after the boost
func (c *Client) BoostChat(chatId int64, slotIds []int32) (*ChatBoostSlots, error) {
	req := &BoostChat{
		ChatId:  chatId,
		SlotIds: slotIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostSlots), nil
}

// BuyGiftUpgrade Pays for upgrade of a regular gift that is owned by another user or channel chat
func (c *Client) BuyGiftUpgrade(ownerId *MessageSender, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	req := &BuyGiftUpgrade{
		OwnerId:            ownerId,
		PrepaidUpgradeHash: prepaidUpgradeHash,
		StarCount:          starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CanBotSendMessages Checks whether the specified bot can send messages to the user. Returns a 404 error if can't and the access can be granted by call to allowBotToSendMessages @bot_user_id Identifier of the target bot
func (c *Client) CanBotSendMessages(botUserId int64) (*Ok, error) {
	req := &CanBotSendMessages{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing @file_id Identifier of a file to stop downloading @only_if_pending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (c *Client) CancelDownloadFile(fileId int32, onlyIfPending bool) (*Ok, error) {
	req := &CancelDownloadFile{
		FileId:        fileId,
		OnlyIfPending: onlyIfPending,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CancelPasswordReset Cancels reset of 2-step verification password. The method can be called if passwordState.pending_reset_date > 0
func (c *Client) CancelPasswordReset() (*Ok, error) {
	req := &CancelPasswordReset{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CancelPreliminaryUploadFile Stops the preliminary uploading of a file. Supported only for files uploaded by using preliminaryUploadFile @file_id Identifier of the file to stop uploading
func (c *Client) CancelPreliminaryUploadFile(fileId int32) (*Ok, error) {
	req := &CancelPreliminaryUploadFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CancelRecoveryEmailAddressVerification Cancels verification of the 2-step verification recovery email address
func (c *Client) CancelRecoveryEmailAddressVerification() (*PasswordState, error) {
	req := &CancelRecoveryEmailAddressVerification{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// CanPostStory Checks whether the current user can post a story on behalf of a chat; requires can_post_stories administrator right for supergroup and channel chats
func (c *Client) CanPostStory(chatId int64) (*CanPostStoryResult, error) {
	req := &CanPostStory{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CanPostStoryResult), nil
}

// CanPurchaseFromStore Checks whether an in-store purchase is possible. Must be called before any in-store purchase. For official applications only @purpose Transaction purpose
func (c *Client) CanPurchaseFromStore(purpose *StorePaymentPurpose) (*Ok, error) {
	req := &CanPurchaseFromStore{
		Purpose: purpose,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CanSendGift Checks whether a gift with next_send_date in the future can be sent already
func (c *Client) CanSendGift(giftId string) (*CanSendGiftResult, error) {
	req := &CanSendGift{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CanSendGiftResult), nil
}

// CanSendMessageToUser Checks whether the current user can message another user or try to create a chat with them
func (c *Client) CanSendMessageToUser(userId int64, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	req := &CanSendMessageToUser{
		UserId:    userId,
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CanSendMessageToUserResult), nil
}

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (c *Client) CanTransferOwnership() (*CanTransferOwnershipResult, error) {
	req := &CanTransferOwnership{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CanTransferOwnershipResult), nil
}

// ChangeImportedContacts Changes imported contacts using the list of contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts.
func (c *Client) ChangeImportedContacts(contacts []*ImportedContact) (*ImportedContacts, error) {
	req := &ChangeImportedContacts{
		Contacts: contacts,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ImportedContacts), nil
}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set @set_id Identifier of the sticker set @is_installed The new value of is_installed @is_archived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (c *Client) ChangeStickerSet(setId string, isInstalled bool, isArchived bool) (*Ok, error) {
	req := &ChangeStickerSet{
		SetId:       setId,
		IsInstalled: isInstalled,
		IsArchived:  isArchived,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in @token The bot token
func (c *Client) CheckAuthenticationBotToken(token string) (*Ok, error) {
	req := &CheckAuthenticationBotToken{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode @code Authentication code to check
func (c *Client) CheckAuthenticationCode(code string) (*Ok, error) {
	req := &CheckAuthenticationCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationEmailCode Checks the authentication of an email address. Works only when the current authorization state is authorizationStateWaitEmailCode @code Email address authentication to check
func (c *Client) CheckAuthenticationEmailCode(code *EmailAddressAuthentication) (*Ok, error) {
	req := &CheckAuthenticationEmailCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationPasskey Checks a passkey to log in to the corresponding account. Call getAuthenticationPasskeyParameters to get parameters for the passkey. Works only when the current authorization state is
func (c *Client) CheckAuthenticationPasskey(credentialId string, clientData string, authenticatorData string, signature string, userHandle string) (*Ok, error) {
	req := &CheckAuthenticationPasskey{
		CredentialId:      credentialId,
		ClientData:        clientData,
		AuthenticatorData: authenticatorData,
		Signature:         signature,
		UserHandle:        userHandle,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationPassword Checks the 2-step verification password for correctness. Works only when the current authorization state is authorizationStateWaitPassword @password The 2-step verification password to check
func (c *Client) CheckAuthenticationPassword(password string) (*Ok, error) {
	req := &CheckAuthenticationPassword{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid. Works only when the current authorization state is authorizationStateWaitPassword @recovery_code Recovery code to check
func (c *Client) CheckAuthenticationPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	req := &CheckAuthenticationPasswordRecoveryCode{
		RecoveryCode: recoveryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckAuthenticationPremiumPurchase Checks whether an in-store purchase of Telegram Premium is possible before authorization. Works only when the current authorization state is authorizationStateWaitPremiumPurchase
func (c *Client) CheckAuthenticationPremiumPurchase(currency string, amount int64) (*Ok, error) {
	req := &CheckAuthenticationPremiumPurchase{
		Currency: currency,
		Amount:   amount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckChatFolderInviteLink Checks the validity of an invite link for a chat folder and returns information about the corresponding chat folder @invite_link Invite link to be checked
func (c *Client) CheckChatFolderInviteLink(inviteLink string) (*ChatFolderInviteLinkInfo, error) {
	req := &CheckChatFolderInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInviteLinkInfo), nil
}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat @invite_link Invite link to be checked
func (c *Client) CheckChatInviteLink(inviteLink string) (*ChatInviteLinkInfo, error) {
	req := &CheckChatInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLinkInfo), nil
}

// CheckChatUsername Checks whether a username can be set for a chat @chat_id Chat identifier; must be identifier of a supergroup chat, or a channel chat, or a private chat with self, or 0 if the chat is being created @username Username to be checked
func (c *Client) CheckChatUsername(chatId int64, username string) (*CheckChatUsernameResult, error) {
	req := &CheckChatUsername{
		ChatId:   chatId,
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CheckChatUsernameResult), nil
}

// CheckCreatedPublicChatsLimit Checks whether the maximum number of owned public chats has been reached. Returns corresponding error if the limit was reached. The limit can be increased with Telegram Premium @type Type of the public chats, for which to check the limit
func (c *Client) CheckCreatedPublicChatsLimit(typeField *PublicChatType) (*Ok, error) {
	req := &CheckCreatedPublicChatsLimit{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport @code Verification code to check
func (c *Client) CheckEmailAddressVerificationCode(code string) (*Ok, error) {
	req := &CheckEmailAddressVerificationCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckLoginEmailAddressCode Checks the login email address authentication @code Email address authentication to check
func (c *Client) CheckLoginEmailAddressCode(code *EmailAddressAuthentication) (*Ok, error) {
	req := &CheckLoginEmailAddressCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid @recovery_code Recovery code to check
func (c *Client) CheckPasswordRecoveryCode(recoveryCode string) (*Ok, error) {
	req := &CheckPasswordRecoveryCode{
		RecoveryCode: recoveryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckPhoneNumberCode Checks the authentication code and completes the request for which the code was sent if appropriate @code Authentication code to check
func (c *Client) CheckPhoneNumberCode(code string) (*Ok, error) {
	req := &CheckPhoneNumberCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckPremiumGiftCode Returns information about a Telegram Premium gift code @code The code to check
func (c *Client) CheckPremiumGiftCode(code string) (*PremiumGiftCodeInfo, error) {
	req := &CheckPremiumGiftCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumGiftCodeInfo), nil
}

// CheckQuickReplyShortcutName Checks validness of a name for a quick reply shortcut. Can be called synchronously @name The name of the shortcut; 1-32 characters
func (c *Client) CheckQuickReplyShortcutName(name string) (*Ok, error) {
	req := &CheckQuickReplyShortcutName{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CheckRecoveryEmailAddressCode Checks the 2-step verification recovery email address verification code @code Verification code to check
func (c *Client) CheckRecoveryEmailAddressCode(code string) (*PasswordState, error) {
	req := &CheckRecoveryEmailAddressCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// CheckStickerSetName Checks whether a name can be used for a new sticker set @name Name to be checked
func (c *Client) CheckStickerSetName(name string) (*CheckStickerSetNameResult, error) {
	req := &CheckStickerSetName{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CheckStickerSetNameResult), nil
}

// CheckWebAppFileDownload Checks whether a file can be downloaded and saved locally by Web App request
func (c *Client) CheckWebAppFileDownload(botUserId int64, fileName string, url string) (*Ok, error) {
	req := &CheckWebAppFileDownload{
		BotUserId: botUserId,
		FileName:  fileName,
		Url:       url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CleanFileName Removes potentially dangerous characters from the name of a file. Returns an empty string on failure. Can be called synchronously @file_name File name or path to the file
func (c *Client) CleanFileName(fileName string) (*Text, error) {
	req := &CleanFileName{
		FileName: fileName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// ClearAllDraftMessages Clears message drafts in all chats @exclude_secret_chats Pass true to keep local message drafts in secret chats
func (c *Client) ClearAllDraftMessages(excludeSecretChats bool) (*Ok, error) {
	req := &ClearAllDraftMessages{
		ExcludeSecretChats: excludeSecretChats,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearAutosaveSettingsExceptions Clears the list of all autosave settings exceptions. The method is guaranteed to work only after at least one call to getAutosaveSettings
func (c *Client) ClearAutosaveSettingsExceptions() (*Ok, error) {
	req := &ClearAutosaveSettingsExceptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearImportedContacts Clears all imported contacts, contact list remains unchanged
func (c *Client) ClearImportedContacts() (*Ok, error) {
	req := &ClearImportedContacts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearRecentEmojiStatuses Clears the list of recently used emoji statuses for self status
func (c *Client) ClearRecentEmojiStatuses() (*Ok, error) {
	req := &ClearRecentEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (c *Client) ClearRecentlyFoundChats() (*Ok, error) {
	req := &ClearRecentlyFoundChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearRecentReactions Clears the list of recently used reactions
func (c *Client) ClearRecentReactions() (*Ok, error) {
	req := &ClearRecentReactions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearRecentStickers Clears the list of recently used stickers @is_attached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (c *Client) ClearRecentStickers(isAttached bool) (*Ok, error) {
	req := &ClearRecentStickers{
		IsAttached: isAttached,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClearSearchedForTags Clears the list of recently searched for hashtags or cashtags @clear_cashtags Pass true to clear the list of recently searched for cashtags; otherwise, the list of recently searched for hashtags will be cleared
func (c *Client) ClearSearchedForTags(clearCashtags bool) (*Ok, error) {
	req := &ClearSearchedForTags{
		ClearCashtags: clearCashtags,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClickAnimatedEmojiMessage Informs TDLib that a message with an animated emoji was clicked by the user. Returns a big animated sticker to be played or a 404 error if usual animation needs to be played @chat_id Chat identifier of the message @message_id Identifier of the clicked message
func (c *Client) ClickAnimatedEmojiMessage(chatId int64, messageId int64) (*Sticker, error) {
	req := &ClickAnimatedEmojiMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Sticker), nil
}

// ClickChatSponsoredMessage Informs TDLib that the user opened the sponsored chat via the button, the name, the chat photo, a mention in the sponsored message text, or the media in the sponsored message
func (c *Client) ClickChatSponsoredMessage(chatId int64, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	req := &ClickChatSponsoredMessage{
		ChatId:         chatId,
		MessageId:      messageId,
		IsMediaClick:   isMediaClick,
		FromFullscreen: fromFullscreen,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClickPremiumSubscriptionButton Informs TDLib that the user clicked Premium subscription button on the Premium features screen
func (c *Client) ClickPremiumSubscriptionButton() (*Ok, error) {
	req := &ClickPremiumSubscriptionButton{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ClickVideoMessageAdvertisement Informs TDLib that the user clicked a video message advertisement @advertisement_unique_id Unique identifier of the advertisement
func (c *Client) ClickVideoMessageAdvertisement(advertisementUniqueId int64) (*Ok, error) {
	req := &ClickVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (c *Client) Close() (*Ok, error) {
	req := &Close{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CloseChat Informs TDLib that the chat is closed by the user. Many useful activities depend on the chat being opened or closed @chat_id Chat identifier
func (c *Client) CloseChat(chatId int64) (*Ok, error) {
	req := &CloseChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CloseGiftAuction Informs TDLib that a gift auction was closed by the user @gift_id Identifier of the gift, which auction was closed
func (c *Client) CloseGiftAuction(giftId string) (*Ok, error) {
	req := &CloseGiftAuction{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CloseSecretChat Closes a secret chat, effectively transferring its state to secretChatStateClosed @secret_chat_id Secret chat identifier
func (c *Client) CloseSecretChat(secretChatId int32) (*Ok, error) {
	req := &CloseSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CloseStory Informs TDLib that a story is closed by the user
func (c *Client) CloseStory(storyPosterChatId int64, storyId int32) (*Ok, error) {
	req := &CloseStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CloseWebApp Informs TDLib that a previously opened Web App was closed @web_app_launch_id Identifier of Web App launch, received from openWebApp
func (c *Client) CloseWebApp(webAppLaunchId string) (*Ok, error) {
	req := &CloseWebApp{
		WebAppLaunchId: webAppLaunchId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CommitPendingLiveStoryReactions Applies all pending paid reactions in a live story group call @group_call_id Group call identifier
func (c *Client) CommitPendingLiveStoryReactions(groupCallId int32) (*Ok, error) {
	req := &CommitPendingLiveStoryReactions{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// CommitPendingPaidMessageReactions Applies all pending paid reactions on a message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message
func (c *Client) CommitPendingPaidMessageReactions(chatId int64, messageId int64) (*Ok, error) {
	req := &CommitPendingPaidMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success @link A link from a QR code. The link must be scanned by the in-app camera
func (c *Client) ConfirmQrCodeAuthentication(link string) (*Session, error) {
	req := &ConfirmQrCodeAuthentication{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Session), nil
}

// ConfirmSession Confirms an unconfirmed session of the current user from another device @session_id Session identifier
func (c *Client) ConfirmSession(sessionId string) (*Ok, error) {
	req := &ConfirmSession{
		SessionId: sessionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ConnectAffiliateProgram Connects an affiliate program to the given affiliate. Returns information about the connected affiliate program
func (c *Client) ConnectAffiliateProgram(affiliate *AffiliateType, botUserId int64) (*ConnectedAffiliateProgram, error) {
	req := &ConnectAffiliateProgram{
		Affiliate: affiliate,
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectedAffiliateProgram), nil
}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group @basic_group_id Basic group identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreateBasicGroupChat(basicGroupId int64, force bool) (*Chat, error) {
	req := &CreateBasicGroupChat{
		BasicGroupId: basicGroupId,
		Force:        force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreateBusinessChatLink Creates a business chat link for the current account. Requires Telegram Business subscription. There can be up to getOption("business_chat_link_count_max") links created. Returns the created link
func (c *Client) CreateBusinessChatLink(linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	req := &CreateBusinessChatLink{
		LinkInfo: linkInfo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessChatLink), nil
}

// CreateCall Creates a new call
func (c *Client) CreateCall(userId int64, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	req := &CreateCall{
		UserId:   userId,
		Protocol: protocol,
		IsVideo:  isVideo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CallId), nil
}

// CreateChatFolder Creates new chat folder. Returns information about the created chat folder. There can be up to getOption("chat_folder_count_max") chat folders, but the limit can be increased with Telegram Premium @folder The new chat folder
func (c *Client) CreateChatFolder(folder *ChatFolder) (*ChatFolderInfo, error) {
	req := &CreateChatFolder{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInfo), nil
}

// CreateChatFolderInviteLink Creates a new invite link for a chat folder. A link can be created for a chat folder if it has only pinned and included chats
func (c *Client) CreateChatFolderInviteLink(chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	req := &CreateChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		Name:         name,
		ChatIds:      chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInviteLink), nil
}

// CreateChatInviteLink Creates a new invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat
func (c *Client) CreateChatInviteLink(chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	req := &CreateChatInviteLink{
		ChatId:             chatId,
		Name:               name,
		ExpirationDate:     expirationDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// CreateChatSubscriptionInviteLink Creates a new subscription invite link for a channel chat. Requires can_invite_users right in the chat
func (c *Client) CreateChatSubscriptionInviteLink(chatId int64, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	req := &CreateChatSubscriptionInviteLink{
		ChatId:              chatId,
		Name:                name,
		SubscriptionPricing: subscriptionPricing,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// CreateForumTopic Creates a topic in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator or can_create_topics member right in the supergroup
func (c *Client) CreateForumTopic(chatId int64, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	req := &CreateForumTopic{
		ChatId:         chatId,
		Name:           name,
		IsNameImplicit: isNameImplicit,
		Icon:           icon,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ForumTopicInfo), nil
}

// CreateGiftCollection Creates a collection from gifts on the current user's or a channel's profile page; requires can_post_messages administrator right in the channel chat.
func (c *Client) CreateGiftCollection(ownerId *MessageSender, name string, receivedGiftIds []string) (*GiftCollection, error) {
	req := &CreateGiftCollection{
		OwnerId:         ownerId,
		Name:            name,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollection), nil
}

// CreateGroupCall Creates a new group call that isn't bound to a chat @join_parameters Parameters to join the call; pass null to only create call link without joining the call
func (c *Client) CreateGroupCall(joinParameters *GroupCallJoinParameters) (*GroupCallInfo, error) {
	req := &CreateGroupCall{
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallInfo), nil
}

// CreateInvoiceLink Creates a link for the given invoice; for bots only
func (c *Client) CreateInvoiceLink(businessConnectionId string, invoice *InputMessageContent) (*HttpUrl, error) {
	req := &CreateInvoiceLink{
		BusinessConnectionId: businessConnectionId,
		Invoice:              invoice,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns information about the newly created chat
func (c *Client) CreateNewBasicGroupChat(userIds []int64, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	req := &CreateNewBasicGroupChat{
		UserIds:               userIds,
		Title:                 title,
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CreatedBasicGroupChat), nil
}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat @user_id Identifier of the target user
func (c *Client) CreateNewSecretChat(userId int64) (*Chat, error) {
	req := &CreateNewSecretChat{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreateNewStickerSet Creates a new sticker set. Returns the newly created sticker set
func (c *Client) CreateNewStickerSet(userId int64, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	req := &CreateNewStickerSet{
		UserId:          userId,
		Title:           title,
		Name:            name,
		StickerType:     stickerType,
		NeedsRepainting: needsRepainting,
		Stickers:        stickers,
		Source:          source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSet), nil
}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
func (c *Client) CreateNewSupergroupChat(title string, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	req := &CreateNewSupergroupChat{
		Title:                 title,
		IsForum:               isForum,
		IsChannel:             isChannel,
		Description:           description,
		MessageAutoDeleteTime: messageAutoDeleteTime,
		ForImport:             forImport,
	}
	if opts != nil {
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreatePrivateChat Returns an existing chat corresponding to a given user @user_id User identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreatePrivateChat(userId int64, force bool) (*Chat, error) {
	req := &CreatePrivateChat{
		UserId: userId,
		Force:  force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat @secret_chat_id Secret chat identifier
func (c *Client) CreateSecretChat(secretChatId int32) (*Chat, error) {
	req := &CreateSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreateStoryAlbum Creates an album of stories; requires can_edit_stories administrator right for supergroup and channel chats
func (c *Client) CreateStoryAlbum(storyPosterChatId int64, name string, storyIds []int32) (*StoryAlbum, error) {
	req := &CreateStoryAlbum{
		StoryPosterChatId: storyPosterChatId,
		Name:              name,
		StoryIds:          storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbum), nil
}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel @supergroup_id Supergroup or channel identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreateSupergroupChat(supergroupId int64, force bool) (*Chat, error) {
	req := &CreateSupergroupChat{
		SupergroupId: supergroupId,
		Force:        force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// CreateTemporaryPassword Creates a new temporary password for processing payments @password The 2-step verification password of the current user @valid_for Time during which the temporary password will be valid, in seconds; must be between 60 and 86400
func (c *Client) CreateTemporaryPassword(password string, validFor int32) (*TemporaryPasswordState, error) {
	req := &CreateTemporaryPassword{
		Password: password,
		ValidFor: validFor,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TemporaryPasswordState), nil
}

// CreateVideoChat Creates a video chat (a group call bound to a chat). Available only for basic groups, supergroups and channels; requires can_manage_video_chats administrator right
func (c *Client) CreateVideoChat(chatId int64, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	req := &CreateVideoChat{
		ChatId:       chatId,
		Title:        title,
		StartDate:    startDate,
		IsRtmpStream: isRtmpStream,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallId), nil
}

// DeclineGroupCallInvitation Declines an invitation to an active group call via messageGroupCall. Can be called both by the sender and the receiver of the invitation
func (c *Client) DeclineGroupCallInvitation(chatId int64, messageId int64) (*Ok, error) {
	req := &DeclineGroupCallInvitation{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeclineSuggestedPost Declines a suggested post in a channel direct messages chat
func (c *Client) DeclineSuggestedPost(chatId int64, messageId int64, comment string) (*Ok, error) {
	req := &DeclineSuggestedPost{
		ChatId:    chatId,
		MessageId: messageId,
		Comment:   comment,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DecryptGroupCallData Decrypts group call data received by tgcalls
func (c *Client) DecryptGroupCallData(groupCallId int32, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	req := &DecryptGroupCallData{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		Data:          data,
	}
	if opts != nil {
		req.DataChannel = opts.DataChannel
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Data), nil
}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account.
func (c *Client) DeleteAccount(reason string, password string) (*Ok, error) {
	req := &DeleteAccount{
		Reason:   reason,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteAllCallMessages Deletes all call messages @revoke Pass true to delete the messages for all users
func (c *Client) DeleteAllCallMessages(revoke bool) (*Ok, error) {
	req := &DeleteAllCallMessages{
		Revoke: revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteAllRevokedChatInviteLinks Deletes all revoked chat invite links created by a given chat administrator. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) DeleteAllRevokedChatInviteLinks(chatId int64, creatorUserId int64) (*Ok, error) {
	req := &DeleteAllRevokedChatInviteLinks{
		ChatId:        chatId,
		CreatorUserId: creatorUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteBotMediaPreviews Deletes media previews from the list of media previews of a bot
func (c *Client) DeleteBotMediaPreviews(botUserId int64, languageCode string, fileIds []int32) (*Ok, error) {
	req := &DeleteBotMediaPreviews{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileIds:      fileIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteBusinessChatLink Deletes a business chat link of the current account @link The link to delete
func (c *Client) DeleteBusinessChatLink(link string) (*Ok, error) {
	req := &DeleteBusinessChatLink{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteBusinessConnectedBot Deletes the business bot that is connected to the current user account @bot_user_id Unique user identifier for the bot
func (c *Client) DeleteBusinessConnectedBot(botUserId int64) (*Ok, error) {
	req := &DeleteBusinessConnectedBot{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteBusinessMessages Deletes messages on behalf of a business account; for bots only
func (c *Client) DeleteBusinessMessages(businessConnectionId string, messageIds []int64) (*Ok, error) {
	req := &DeleteBusinessMessages{
		BusinessConnectionId: businessConnectionId,
		MessageIds:           messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteBusinessStory Deletes a story posted by the bot on behalf of a business account; for bots only
func (c *Client) DeleteBusinessStory(businessConnectionId string, storyId int32) (*Ok, error) {
	req := &DeleteBusinessStory{
		BusinessConnectionId: businessConnectionId,
		StoryId:              storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChat Deletes a chat along with all messages in the corresponding chat for all chat members. For group chats this will release the usernames and remove all members.
func (c *Client) DeleteChat(chatId int64) (*Ok, error) {
	req := &DeleteChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatBackground Deletes background in a specific chat
func (c *Client) DeleteChatBackground(chatId int64, restorePrevious bool) (*Ok, error) {
	req := &DeleteChatBackground{
		ChatId:          chatId,
		RestorePrevious: restorePrevious,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatFolder Deletes existing chat folder @chat_folder_id Chat folder identifier @leave_chat_ids Identifiers of the chats to leave. The chats must be pinned or always included in the folder
func (c *Client) DeleteChatFolder(chatFolderId int32, leaveChatIds []int64) (*Ok, error) {
	req := &DeleteChatFolder{
		ChatFolderId: chatFolderId,
		LeaveChatIds: leaveChatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatFolderInviteLink Deletes an invite link for a chat folder
func (c *Client) DeleteChatFolderInviteLink(chatFolderId int32, inviteLink string) (*Ok, error) {
	req := &DeleteChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		InviteLink:   inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatHistory Deletes all messages in the chat. Use chat.can_be_deleted_only_for_self and chat.can_be_deleted_for_all_users fields to find whether and how the method can be applied to the chat
func (c *Client) DeleteChatHistory(chatId int64, removeFromChatList bool, revoke bool) (*Ok, error) {
	req := &DeleteChatHistory{
		ChatId:             chatId,
		RemoveFromChatList: removeFromChatList,
		Revoke:             revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatMessagesByDate Deletes all messages between the specified dates in a chat. Supported only for private chats and basic groups. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteChatMessagesByDate(chatId int64, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	req := &DeleteChatMessagesByDate{
		ChatId:  chatId,
		MinDate: minDate,
		MaxDate: maxDate,
		Revoke:  revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatMessagesBySender Deletes all messages sent by the specified message sender in a chat. Supported only for supergroups; requires can_delete_messages administrator right @chat_id Chat identifier @sender_id Identifier of the sender of messages to delete
func (c *Client) DeleteChatMessagesBySender(chatId int64, senderId *MessageSender) (*Ok, error) {
	req := &DeleteChatMessagesBySender{
		ChatId:   chatId,
		SenderId: senderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a replyMarkupForceReply reply markup has been used or dismissed
func (c *Client) DeleteChatReplyMarkup(chatId int64, messageId int64) (*Ok, error) {
	req := &DeleteChatReplyMarkup{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteCommands Deletes commands supported by the bot for the given user scope and language; for bots only
func (c *Client) DeleteCommands(languageCode string, opts *DeleteCommandsOpts) (*Ok, error) {
	req := &DeleteCommands{
		LanguageCode: languageCode,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteDefaultBackground Deletes default background for chats @for_dark_theme Pass true if the background is deleted for a dark theme
func (c *Client) DeleteDefaultBackground(forDarkTheme bool) (*Ok, error) {
	req := &DeleteDefaultBackground{
		ForDarkTheme: forDarkTheme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteDirectMessagesChatTopicHistory Deletes all messages in the topic in a channel direct messages chat administered by the current user
func (c *Client) DeleteDirectMessagesChatTopicHistory(chatId int64, topicId int64) (*Ok, error) {
	req := &DeleteDirectMessagesChatTopicHistory{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteDirectMessagesChatTopicMessagesByDate Deletes all messages between the specified dates in the topic in a channel direct messages chat administered by the current user. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteDirectMessagesChatTopicMessagesByDate(chatId int64, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	req := &DeleteDirectMessagesChatTopicMessagesByDate{
		ChatId:  chatId,
		TopicId: topicId,
		MinDate: minDate,
		MaxDate: maxDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteFile Deletes a file from the TDLib file cache @file_id Identifier of the file to delete
func (c *Client) DeleteFile(fileId int32) (*Ok, error) {
	req := &DeleteFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteForumTopic Deletes all messages from a topic in a forum supergroup chat or a chat with a bot with topics; requires can_delete_messages administrator right in the supergroup
func (c *Client) DeleteForumTopic(chatId int64, forumTopicId int32) (*Ok, error) {
	req := &DeleteForumTopic{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteGiftCollection Deletes a gift collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat
func (c *Client) DeleteGiftCollection(ownerId *MessageSender, collectionId int32) (*Ok, error) {
	req := &DeleteGiftCollection{
		OwnerId:      ownerId,
		CollectionId: collectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteGroupCallMessages Deletes messages in a group call; for live story calls only. Requires groupCallMessage.can_be_deleted right
func (c *Client) DeleteGroupCallMessages(groupCallId int32, messageIds []int32, reportSpam bool) (*Ok, error) {
	req := &DeleteGroupCallMessages{
		GroupCallId: groupCallId,
		MessageIds:  messageIds,
		ReportSpam:  reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteGroupCallMessagesBySender Deletes all messages sent by the specified message sender in a group call; for live story calls only. Requires groupCall.can_delete_messages right
func (c *Client) DeleteGroupCallMessagesBySender(groupCallId int32, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	req := &DeleteGroupCallMessagesBySender{
		GroupCallId: groupCallId,
		SenderId:    senderId,
		ReportSpam:  reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack which is currently in use (including base language pack) or is being synchronized can't be deleted.
func (c *Client) DeleteLanguagePack(languagePackId string) (*Ok, error) {
	req := &DeleteLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteMessages Deletes messages
func (c *Client) DeleteMessages(chatId int64, messageIds []int64, revoke bool) (*Ok, error) {
	req := &DeleteMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
		Revoke:     revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeletePassportElement Deletes a Telegram Passport element @type Element type
func (c *Client) DeletePassportElement(typeField *PassportElementType) (*Ok, error) {
	req := &DeletePassportElement{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteProfilePhoto Deletes a profile photo @profile_photo_id Identifier of the profile photo to delete
func (c *Client) DeleteProfilePhoto(profilePhotoId string) (*Ok, error) {
	req := &DeleteProfilePhoto{
		ProfilePhotoId: profilePhotoId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteQuickReplyShortcut Deletes a quick reply shortcut @shortcut_id Unique identifier of the quick reply shortcut
func (c *Client) DeleteQuickReplyShortcut(shortcutId int32) (*Ok, error) {
	req := &DeleteQuickReplyShortcut{
		ShortcutId: shortcutId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteQuickReplyShortcutMessages Deletes specified quick reply messages
func (c *Client) DeleteQuickReplyShortcutMessages(shortcutId int32, messageIds []int64) (*Ok, error) {
	req := &DeleteQuickReplyShortcutMessages{
		ShortcutId: shortcutId,
		MessageIds: messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteRevokedChatInviteLink Deletes revoked chat invite links. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links @chat_id Chat identifier @invite_link Invite link to revoke
func (c *Client) DeleteRevokedChatInviteLink(chatId int64, inviteLink string) (*Ok, error) {
	req := &DeleteRevokedChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (c *Client) DeleteSavedCredentials() (*Ok, error) {
	req := &DeleteSavedCredentials{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteSavedMessagesTopicHistory Deletes all messages in a Saved Messages topic @saved_messages_topic_id Identifier of Saved Messages topic which messages will be deleted
func (c *Client) DeleteSavedMessagesTopicHistory(savedMessagesTopicId int64) (*Ok, error) {
	req := &DeleteSavedMessagesTopicHistory{
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteSavedMessagesTopicMessagesByDate Deletes all messages between the specified dates in a Saved Messages topic. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteSavedMessagesTopicMessagesByDate(savedMessagesTopicId int64, minDate int32, maxDate int32) (*Ok, error) {
	req := &DeleteSavedMessagesTopicMessagesByDate{
		SavedMessagesTopicId: savedMessagesTopicId,
		MinDate:              minDate,
		MaxDate:              maxDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteSavedOrderInfo Deletes saved order information
func (c *Client) DeleteSavedOrderInfo() (*Ok, error) {
	req := &DeleteSavedOrderInfo{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteStickerSet Completely deletes a sticker set @name Sticker set name. The sticker set must be owned by the current user
func (c *Client) DeleteStickerSet(name string) (*Ok, error) {
	req := &DeleteStickerSet{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteStory Deletes a previously posted story. Can be called only if story.can_be_deleted == true
func (c *Client) DeleteStory(storyPosterChatId int64, storyId int32) (*Ok, error) {
	req := &DeleteStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DeleteStoryAlbum Deletes a story album. If the album is owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat
func (c *Client) DeleteStoryAlbum(chatId int64, storyAlbumId int32) (*Ok, error) {
	req := &DeleteStoryAlbum{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed.
func (c *Client) Destroy() (*Ok, error) {
	req := &Destroy{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DisableAllSupergroupUsernames Disables all active non-editable usernames of a supergroup or channel, requires owner privileges in the supergroup or channel @supergroup_id Identifier of the supergroup or channel
func (c *Client) DisableAllSupergroupUsernames(supergroupId int64) (*Ok, error) {
	req := &DisableAllSupergroupUsernames{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (c *Client) DisableProxy() (*Ok, error) {
	req := &DisableProxy{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DiscardCall Discards a call
func (c *Client) DiscardCall(callId int32, isDisconnected bool, inviteLink string, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	req := &DiscardCall{
		CallId:         callId,
		IsDisconnected: isDisconnected,
		InviteLink:     inviteLink,
		Duration:       duration,
		IsVideo:        isVideo,
		ConnectionId:   connectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DisconnectAffiliateProgram Disconnects an affiliate program from the given affiliate and immediately deactivates its referral link. Returns updated information about the disconnected affiliate program
func (c *Client) DisconnectAffiliateProgram(affiliate *AffiliateType, url string) (*ConnectedAffiliateProgram, error) {
	req := &DisconnectAffiliateProgram{
		Affiliate: affiliate,
		Url:       url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectedAffiliateProgram), nil
}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (c *Client) DisconnectAllWebsites() (*Ok, error) {
	req := &DisconnectAllWebsites{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DisconnectWebsite Disconnects website from the current user's Telegram account @website_id Website identifier
func (c *Client) DisconnectWebsite(websiteId string) (*Ok, error) {
	req := &DisconnectWebsite{
		WebsiteId: websiteId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
func (c *Client) DownloadFile(fileId int32, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	req := &DownloadFile{
		FileId:      fileId,
		Priority:    priority,
		Offset:      offset,
		Limit:       limit,
		Synchronous: synchronous,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// DropGiftOriginalDetails Drops original details for an upgraded gift
func (c *Client) DropGiftOriginalDetails(receivedGiftId string, starCount int64) (*Ok, error) {
	req := &DropGiftOriginalDetails{
		ReceivedGiftId: receivedGiftId,
		StarCount:      starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditBotMediaPreview Replaces media preview in the list of media previews of a bot. Returns the new preview after edit is completed server-side
func (c *Client) EditBotMediaPreview(botUserId int64, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	req := &EditBotMediaPreview{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileId:       fileId,
		Content:      content,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotMediaPreview), nil
}

// EditBusinessChatLink Edits a business chat link of the current account. Requires Telegram Business subscription. Returns the edited link
func (c *Client) EditBusinessChatLink(link string, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	req := &EditBusinessChatLink{
		Link:     link,
		LinkInfo: linkInfo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessChatLink), nil
}

// EditBusinessMessageCaption Edits the caption of a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageCaption(businessConnectionId string, chatId int64, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageCaption{
		BusinessConnectionId:  businessConnectionId,
		ChatId:                chatId,
		MessageId:             messageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessMessageChecklist Edits the content of a checklist in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageChecklist(businessConnectionId string, chatId int64, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageChecklist{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		Checklist:            checklist,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessMessageLiveLocation Edits the content of a live location in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageLiveLocation(businessConnectionId string, chatId int64, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageLiveLocation{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessMessageMedia Edits the media content of a message with a text, an animation, an audio, a document, a photo or a video in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageMedia(businessConnectionId string, chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageMedia{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessMessageReplyMarkup Edits the reply markup of a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageReplyMarkup(businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageReplyMarkup{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessMessageText Edits the text of a text or game message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageText(businessConnectionId string, chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	req := &EditBusinessMessageText{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// EditBusinessStory Changes a story posted by the bot on behalf of a business account; for bots only
func (c *Client) EditBusinessStory(storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	req := &EditBusinessStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		Content:           content,
		Areas:             areas,
		Caption:           caption,
		PrivacySettings:   privacySettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Story), nil
}

// EditChatFolder Edits existing chat folder. Returns information about the edited chat folder @chat_folder_id Chat folder identifier @folder The edited chat folder
func (c *Client) EditChatFolder(chatFolderId int32, folder *ChatFolder) (*ChatFolderInfo, error) {
	req := &EditChatFolder{
		ChatFolderId: chatFolderId,
		Folder:       folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInfo), nil
}

// EditChatFolderInviteLink Edits an invite link for a chat folder
func (c *Client) EditChatFolderInviteLink(chatFolderId int32, inviteLink string, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	req := &EditChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		InviteLink:   inviteLink,
		Name:         name,
		ChatIds:      chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInviteLink), nil
}

// EditChatInviteLink Edits a non-primary invite link for a chat. Available for basic groups, supergroups, and channels.
func (c *Client) EditChatInviteLink(chatId int64, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	req := &EditChatInviteLink{
		ChatId:             chatId,
		InviteLink:         inviteLink,
		Name:               name,
		ExpirationDate:     expirationDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// EditChatSubscriptionInviteLink Edits a subscription invite link for a channel chat. Requires can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) EditChatSubscriptionInviteLink(chatId int64, inviteLink string, name string) (*ChatInviteLink, error) {
	req := &EditChatSubscriptionInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Name:       name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// EditCustomLanguagePackInfo Edits information about a custom local language pack in the current localization target. Can be called before authorization @info New information about the custom local language pack
func (c *Client) EditCustomLanguagePackInfo(info *LanguagePackInfo) (*Ok, error) {
	req := &EditCustomLanguagePackInfo{
		Info: info,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditForumTopic Edits title and icon of a topic in a forum supergroup chat or a chat with a bot with topics; for supergroup chats requires can_manage_topics administrator right
func (c *Client) EditForumTopic(chatId int64, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	req := &EditForumTopic{
		ChatId:              chatId,
		ForumTopicId:        forumTopicId,
		Name:                name,
		EditIconCustomEmoji: editIconCustomEmoji,
		IconCustomEmojiId:   iconCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageCaption(inlineMessageId string, showCaptionAboveMedia bool, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	req := &EditInlineMessageCaption{
		InlineMessageId:       inlineMessageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageLiveLocation(inlineMessageId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditInlineMessageLiveLocationOpts) (*Ok, error) {
	req := &EditInlineMessageLiveLocation{
		InlineMessageId:      inlineMessageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditInlineMessageMedia Edits the media content of a message with a text, an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageMedia(inlineMessageId string, inputMessageContent *InputMessageContent, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	req := &EditInlineMessageMedia{
		InlineMessageId:     inlineMessageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageReplyMarkup(inlineMessageId string, opts *EditInlineMessageReplyMarkupOpts) (*Ok, error) {
	req := &EditInlineMessageReplyMarkup{
		InlineMessageId: inlineMessageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
func (c *Client) EditInlineMessageText(inlineMessageId string, inputMessageContent *InputMessageContent, opts *EditInlineMessageTextOpts) (*Ok, error) {
	req := &EditInlineMessageText{
		InlineMessageId:     inlineMessageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageCaption(chatId int64, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	req := &EditMessageCaption{
		ChatId:                chatId,
		MessageId:             messageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditMessageChecklist Edits the message content of a checklist. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageChecklist(chatId int64, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	req := &EditMessageChecklist{
		ChatId:    chatId,
		MessageId: messageId,
		Checklist: checklist,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location.
func (c *Client) EditMessageLiveLocation(chatId int64, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	req := &EditMessageLiveLocation{
		ChatId:               chatId,
		MessageId:            messageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditMessageMedia Edits the media content of a message, including message caption. If only the caption needs to be edited, use editMessageCaption instead.
func (c *Client) EditMessageMedia(chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	req := &EditMessageMedia{
		ChatId:              chatId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageReplyMarkup(chatId int64, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	req := &EditMessageReplyMarkup{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditMessageSchedulingState Edits the time when a scheduled message will be sent. Scheduling state of all messages in the same album or forwarded together with the message will be also changed
func (c *Client) EditMessageSchedulingState(chatId int64, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	req := &EditMessageSchedulingState{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.SchedulingState = opts.SchedulingState
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageText(chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	req := &EditMessageText{
		ChatId:              chatId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
func (c *Client) EditProxy(proxyId int32, server string, port int32, enable bool, typeField *ProxyType) (*Proxy, error) {
	req := &EditProxy{
		ProxyId:   proxyId,
		Server:    server,
		Port:      port,
		Enable:    enable,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Proxy), nil
}

// EditQuickReplyMessage Asynchronously edits the text, media or caption of a quick reply message. Use quickReplyMessage.can_be_edited to check whether a message can be edited.
func (c *Client) EditQuickReplyMessage(shortcutId int32, messageId int64, inputMessageContent *InputMessageContent) (*Ok, error) {
	req := &EditQuickReplyMessage{
		ShortcutId:          shortcutId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditStarSubscription Cancels or re-enables Telegram Star subscription
func (c *Client) EditStarSubscription(subscriptionId string, isCanceled bool) (*Ok, error) {
	req := &EditStarSubscription{
		SubscriptionId: subscriptionId,
		IsCanceled:     isCanceled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditStory Changes content and caption of a story. Can be called only if story.can_be_edited == true
func (c *Client) EditStory(storyPosterChatId int64, storyId int32, opts *EditStoryOpts) (*Ok, error) {
	req := &EditStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	if opts != nil {
		req.Content = opts.Content
		req.Areas = opts.Areas
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditStoryCover Changes cover of a video story. Can be called only if story.can_be_edited == true and the story isn't being edited now
func (c *Client) EditStoryCover(storyPosterChatId int64, storyId int32, coverFrameTimestamp float64) (*Ok, error) {
	req := &EditStoryCover{
		StoryPosterChatId:   storyPosterChatId,
		StoryId:             storyId,
		CoverFrameTimestamp: coverFrameTimestamp,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EditUserStarSubscription Cancels or re-enables Telegram Star subscription for a user; for bots only
func (c *Client) EditUserStarSubscription(userId int64, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	req := &EditUserStarSubscription{
		UserId:                  userId,
		TelegramPaymentChargeId: telegramPaymentChargeId,
		IsCanceled:              isCanceled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) EnableProxy(proxyId int32) (*Ok, error) {
	req := &EnableProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EncryptGroupCallData Encrypts group call data before sending them over network using tgcalls
func (c *Client) EncryptGroupCallData(groupCallId int32, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	req := &EncryptGroupCallData{
		GroupCallId:           groupCallId,
		DataChannel:           dataChannel,
		Data:                  data,
		UnencryptedPrefixSize: unencryptedPrefixSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Data), nil
}

// EndGroupCall Ends a group call. Requires groupCall.can_be_managed right for video chats and live stories or groupCall.is_owned otherwise @group_call_id Group call identifier
func (c *Client) EndGroupCall(groupCallId int32) (*Ok, error) {
	req := &EndGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EndGroupCallRecording Ends recording of an active group call; for video chats only. Requires groupCall.can_be_managed right @group_call_id Group call identifier
func (c *Client) EndGroupCallRecording(groupCallId int32) (*Ok, error) {
	req := &EndGroupCallRecording{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// EndGroupCallScreenSharing Ends screen sharing in a joined group call; not supported in live stories @group_call_id Group call identifier
func (c *Client) EndGroupCallScreenSharing(groupCallId int32) (*Ok, error) {
	req := &EndGroupCallScreenSharing{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// FinishFileGeneration Finishes the file generation
func (c *Client) FinishFileGeneration(generationId string, opts *FinishFileGenerationOpts) (*Ok, error) {
	req := &FinishFileGeneration{
		GenerationId: generationId,
	}
	if opts != nil {
		req.Error = opts.Error
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
func (c *Client) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	req := &ForwardMessages{
		ChatId:        chatId,
		FromChatId:    fromChatId,
		MessageIds:    messageIds,
		SendCopy:      sendCopy,
		RemoveCaption: removeCaption,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetAccountTtl Returns the period of inactivity after which the account of the current user will automatically be deleted
func (c *Client) GetAccountTtl() (*AccountTtl, error) {
	req := &GetAccountTtl{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AccountTtl), nil
}

// GetActiveSessions Returns all active sessions of the current user
func (c *Client) GetActiveSessions() (*Sessions, error) {
	req := &GetActiveSessions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Sessions), nil
}

// GetAllPassportElements Returns all available Telegram Passport elements @password The 2-step verification password of the current user
func (c *Client) GetAllPassportElements(password string) (*PassportElements, error) {
	req := &GetAllPassportElements{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PassportElements), nil
}

// GetAllStickerEmojis Returns unique emoji that correspond to stickers to be found by the getStickers(sticker_type, query, 1000000, chat_id)
func (c *Client) GetAllStickerEmojis(stickerType *StickerType, query string, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	req := &GetAllStickerEmojis{
		StickerType:         stickerType,
		Query:               query,
		ChatId:              chatId,
		ReturnOnlyMainEmoji: returnOnlyMainEmoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Emojis), nil
}

// GetAnimatedEmoji Returns an animated emoji corresponding to a given emoji. Returns a 404 error if the emoji has no animated emoji @emoji The emoji
func (c *Client) GetAnimatedEmoji(emoji string) (*AnimatedEmoji, error) {
	req := &GetAnimatedEmoji{
		Emoji: emoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AnimatedEmoji), nil
}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (c *Client) GetApplicationConfig() (*JsonValue, error) {
	req := &GetApplicationConfig{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*JsonValue), nil
}

// GetApplicationDownloadLink Returns the link for downloading official Telegram application to be used when the current user invites friends to Telegram
func (c *Client) GetApplicationDownloadLink() (*HttpUrl, error) {
	req := &GetApplicationDownloadLink{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetArchiveChatListSettings Returns settings for automatic moving of chats to and from the Archive chat lists
func (c *Client) GetArchiveChatListSettings() (*ArchiveChatListSettings, error) {
	req := &GetArchiveChatListSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ArchiveChatListSettings), nil
}

// GetArchivedStickerSets Returns a list of archived sticker sets
func (c *Client) GetArchivedStickerSets(stickerType *StickerType, offsetStickerSetId string, limit int32) (*StickerSets, error) {
	req := &GetArchivedStickerSets{
		StickerType:        stickerType,
		OffsetStickerSetId: offsetStickerSetId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file, including regular, mask, and emoji sticker sets. Currently, only animations, photos, and videos can have attached sticker sets @file_id File identifier
func (c *Client) GetAttachedStickerSets(fileId int32) (*StickerSets, error) {
	req := &GetAttachedStickerSets{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// GetAttachmentMenuBot Returns information about a bot that can be added to attachment or side menu @bot_user_id Bot's user identifier
func (c *Client) GetAttachmentMenuBot(botUserId int64) (*AttachmentMenuBot, error) {
	req := &GetAttachmentMenuBot{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AttachmentMenuBot), nil
}

// GetAuthenticationPasskeyParameters Returns parameters for authentication using a passkey as JSON-serialized string
func (c *Client) GetAuthenticationPasskeyParameters() (*Text, error) {
	req := &GetAuthenticationPasskeyParameters{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetAuthorizationState Returns the current authorization state. This is an offline method. For informational purposes only. Use updateAuthorizationState instead to maintain the current authorization state. Can be called before initialization
func (c *Client) GetAuthorizationState() (*AuthorizationState, error) {
	req := &GetAuthorizationState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AuthorizationState), nil
}

// GetAutoDownloadSettingsPresets Returns auto-download settings presets for the current user
func (c *Client) GetAutoDownloadSettingsPresets() (*AutoDownloadSettingsPresets, error) {
	req := &GetAutoDownloadSettingsPresets{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AutoDownloadSettingsPresets), nil
}

// GetAutosaveSettings Returns autosave settings for the current user
func (c *Client) GetAutosaveSettings() (*AutosaveSettings, error) {
	req := &GetAutosaveSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AutosaveSettings), nil
}

// GetAvailableChatBoostSlots Returns the list of available chat boost slots for the current user
func (c *Client) GetAvailableChatBoostSlots() (*ChatBoostSlots, error) {
	req := &GetAvailableChatBoostSlots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostSlots), nil
}

// GetAvailableGifts Returns gifts that can be sent to other users and channel chats
func (c *Client) GetAvailableGifts() (*AvailableGifts, error) {
	req := &GetAvailableGifts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AvailableGifts), nil
}

// GetBackgroundUrl Constructs a persistent HTTP URL for a background @name Background name @type Background type; backgroundTypeChatTheme isn't supported
func (c *Client) GetBackgroundUrl(name string, typeField *BackgroundType) (*HttpUrl, error) {
	req := &GetBackgroundUrl{
		Name:      name,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetBankCardInfo Returns information about a bank card @bank_card_number The bank card number
func (c *Client) GetBankCardInfo(bankCardNumber string) (*BankCardInfo, error) {
	req := &GetBankCardInfo{
		BankCardNumber: bankCardNumber,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BankCardInfo), nil
}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline method if the current user is not a bot @basic_group_id Basic group identifier
func (c *Client) GetBasicGroup(basicGroupId int64) (*BasicGroup, error) {
	req := &GetBasicGroup{
		BasicGroupId: basicGroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BasicGroup), nil
}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier @basic_group_id Basic group identifier
func (c *Client) GetBasicGroupFullInfo(basicGroupId int64) (*BasicGroupFullInfo, error) {
	req := &GetBasicGroupFullInfo{
		BasicGroupId: basicGroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BasicGroupFullInfo), nil
}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
func (c *Client) GetBlockedMessageSenders(blockList *BlockList, offset int32, limit int32) (*MessageSenders, error) {
	req := &GetBlockedMessageSenders{
		BlockList: blockList,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageSenders), nil
}

// GetBotInfoDescription Returns the text shown in the chat with a bot if the chat is empty in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotInfoDescription(botUserId int64, languageCode string) (*Text, error) {
	req := &GetBotInfoDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetBotInfoShortDescription Returns the text shown on a bot's profile page and sent together with the link when users share the bot in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotInfoShortDescription(botUserId int64, languageCode string) (*Text, error) {
	req := &GetBotInfoShortDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetBotMediaPreviewInfo Returns the list of media previews for the given language and the list of languages for which the bot has dedicated previews
func (c *Client) GetBotMediaPreviewInfo(botUserId int64, languageCode string) (*BotMediaPreviewInfo, error) {
	req := &GetBotMediaPreviewInfo{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotMediaPreviewInfo), nil
}

// GetBotMediaPreviews Returns the list of media previews of a bot @bot_user_id Identifier of the target bot. The bot must have the main Web App
func (c *Client) GetBotMediaPreviews(botUserId int64) (*BotMediaPreviews, error) {
	req := &GetBotMediaPreviews{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotMediaPreviews), nil
}

// GetBotName Returns the name of a bot in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotName(botUserId int64, languageCode string) (*Text, error) {
	req := &GetBotName{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetBotSimilarBotCount Returns approximate number of bots similar to the given bot
func (c *Client) GetBotSimilarBotCount(botUserId int64, returnLocal bool) (*Count, error) {
	req := &GetBotSimilarBotCount{
		BotUserId:   botUserId,
		ReturnLocal: returnLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetBotSimilarBots Returns a list of bots similar to the given bot @bot_user_id User identifier of the target bot
func (c *Client) GetBotSimilarBots(botUserId int64) (*Users, error) {
	req := &GetBotSimilarBots{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// GetBusinessAccountStarAmount Returns the Telegram Star amount owned by a business account; for bots only @business_connection_id Unique identifier of business connection
func (c *Client) GetBusinessAccountStarAmount(businessConnectionId string) (*StarAmount, error) {
	req := &GetBusinessAccountStarAmount{
		BusinessConnectionId: businessConnectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarAmount), nil
}

// GetBusinessChatLinkInfo Returns information about a business chat link @link_name Name of the link
func (c *Client) GetBusinessChatLinkInfo(linkName string) (*BusinessChatLinkInfo, error) {
	req := &GetBusinessChatLinkInfo{
		LinkName: linkName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessChatLinkInfo), nil
}

// GetBusinessChatLinks Returns business chat links created for the current account
func (c *Client) GetBusinessChatLinks() (*BusinessChatLinks, error) {
	req := &GetBusinessChatLinks{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessChatLinks), nil
}

// GetBusinessConnectedBot Returns the business bot that is connected to the current user account. Returns a 404 error if there is no connected bot
func (c *Client) GetBusinessConnectedBot() (*BusinessConnectedBot, error) {
	req := &GetBusinessConnectedBot{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessConnectedBot), nil
}

// GetBusinessConnection Returns information about a business connection by its identifier; for bots only @connection_id Identifier of the business connection to return
func (c *Client) GetBusinessConnection(connectionId string) (*BusinessConnection, error) {
	req := &GetBusinessConnection{
		ConnectionId: connectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessConnection), nil
}

// GetBusinessFeatures Returns information about features, available to Business users @source Source of the request; pass null if the method is called from settings or some non-standard source
func (c *Client) GetBusinessFeatures(source *BusinessFeature) (*BusinessFeatures, error) {
	req := &GetBusinessFeatures{
		Source: source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessFeatures), nil
}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
func (c *Client) GetCallbackQueryAnswer(chatId int64, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	req := &GetCallbackQueryAnswer{
		ChatId:    chatId,
		MessageId: messageId,
		Payload:   payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CallbackQueryAnswer), nil
}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only @chat_id Identifier of the chat the message belongs to @message_id Message identifier @callback_query_id Identifier of the callback query
func (c *Client) GetCallbackQueryMessage(chatId int64, messageId int64, callbackQueryId string) (*Message, error) {
	req := &GetCallbackQueryMessage{
		ChatId:          chatId,
		MessageId:       messageId,
		CallbackQueryId: callbackQueryId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetChat Returns information about a chat by its identifier. This is an offline method if the current user is not a bot @chat_id Chat identifier
func (c *Client) GetChat(chatId int64) (*Chat, error) {
	req := &GetChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// GetChatActiveStories Returns the list of active stories posted by the given chat @chat_id Chat identifier
func (c *Client) GetChatActiveStories(chatId int64) (*ChatActiveStories, error) {
	req := &GetChatActiveStories{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatActiveStories), nil
}

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles @chat_id Chat identifier
func (c *Client) GetChatAdministrators(chatId int64) (*ChatAdministrators, error) {
	req := &GetChatAdministrators{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatAdministrators), nil
}

// GetChatArchivedStories Returns the list of all stories posted by the given chat; requires can_edit_stories administrator right in the chat.
func (c *Client) GetChatArchivedStories(chatId int64, fromStoryId int32, limit int32) (*Stories, error) {
	req := &GetChatArchivedStories{
		ChatId:      chatId,
		FromStoryId: fromStoryId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stories), nil
}

// GetChatAvailableMessageSenders Returns the list of message sender identifiers, which can be used to send messages in a chat @chat_id Chat identifier
func (c *Client) GetChatAvailableMessageSenders(chatId int64) (*ChatMessageSenders, error) {
	req := &GetChatAvailableMessageSenders{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatMessageSenders), nil
}

// GetChatAvailablePaidMessageReactionSenders Returns the list of message sender identifiers, which can be used to send a paid reaction in a chat @chat_id Chat identifier
func (c *Client) GetChatAvailablePaidMessageReactionSenders(chatId int64) (*MessageSenders, error) {
	req := &GetChatAvailablePaidMessageReactionSenders{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageSenders), nil
}

// GetChatBoostFeatures Returns the list of features available for different chat boost levels. This is an offline method
func (c *Client) GetChatBoostFeatures(isChannel bool) (*ChatBoostFeatures, error) {
	req := &GetChatBoostFeatures{
		IsChannel: isChannel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostFeatures), nil
}

// GetChatBoostLevelFeatures Returns the list of features available on the specific chat boost level. This is an offline method
func (c *Client) GetChatBoostLevelFeatures(isChannel bool, level int32) (*ChatBoostLevelFeatures, error) {
	req := &GetChatBoostLevelFeatures{
		IsChannel: isChannel,
		Level:     level,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostLevelFeatures), nil
}

// GetChatBoostLink Returns an HTTPS link to boost the specified supergroup or channel chat @chat_id Identifier of the chat
func (c *Client) GetChatBoostLink(chatId int64) (*ChatBoostLink, error) {
	req := &GetChatBoostLink{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostLink), nil
}

// GetChatBoostLinkInfo Returns information about a link to boost a chat. Can be called for any internal link of the type internalLinkTypeChatBoost @url The link to boost a chat
func (c *Client) GetChatBoostLinkInfo(url string) (*ChatBoostLinkInfo, error) {
	req := &GetChatBoostLinkInfo{
		Url: url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostLinkInfo), nil
}

// GetChatBoosts Returns the list of boosts applied to a chat; requires administrator rights in the chat
func (c *Client) GetChatBoosts(chatId int64, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	req := &GetChatBoosts{
		ChatId:        chatId,
		OnlyGiftCodes: onlyGiftCodes,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundChatBoosts), nil
}

// GetChatBoostStatus Returns the current boost status for a supergroup or a channel chat @chat_id Identifier of the chat
func (c *Client) GetChatBoostStatus(chatId int64) (*ChatBoostStatus, error) {
	req := &GetChatBoostStatus{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatBoostStatus), nil
}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i.e., in order of decreasing event_id)
func (c *Client) GetChatEventLog(chatId int64, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	req := &GetChatEventLog{
		ChatId:      chatId,
		Query:       query,
		FromEventId: fromEventId,
		Limit:       limit,
		UserIds:     userIds,
	}
	if opts != nil {
		req.Filters = opts.Filters
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatEvents), nil
}

// GetChatFolder Returns information about a chat folder by its identifier @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolder(chatFolderId int32) (*ChatFolder, error) {
	req := &GetChatFolder{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolder), nil
}

// GetChatFolderChatCount Returns approximate number of chats in a being created chat folder. Main and archive chat lists must be fully preloaded for this function to work correctly @folder The new chat folder
func (c *Client) GetChatFolderChatCount(folder *ChatFolder) (*Count, error) {
	req := &GetChatFolderChatCount{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetChatFolderChatsToLeave Returns identifiers of pinned or always included chats from a chat folder, which are suggested to be left when the chat folder is deleted @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderChatsToLeave(chatFolderId int32) (*Chats, error) {
	req := &GetChatFolderChatsToLeave{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatFolderDefaultIconName Returns default icon name for a folder. Can be called synchronously @folder Chat folder
func (c *Client) GetChatFolderDefaultIconName(folder *ChatFolder) (*ChatFolderIcon, error) {
	req := &GetChatFolderDefaultIconName{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderIcon), nil
}

// GetChatFolderInviteLinks Returns invite links created by the current user for a shareable chat folder @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderInviteLinks(chatFolderId int32) (*ChatFolderInviteLinks, error) {
	req := &GetChatFolderInviteLinks{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatFolderInviteLinks), nil
}

// GetChatFolderNewChats Returns new chats added to a shareable chat folder by its owner. The method must be called at most once in getOption("chat_folder_new_chats_update_period") for the given chat folder @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderNewChats(chatFolderId int32) (*Chats, error) {
	req := &GetChatFolderNewChats{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatHistory Returns messages in a chat. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id).
func (c *Client) GetChatHistory(chatId int64, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	req := &GetChatHistory{
		ChatId:        chatId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
		OnlyLocal:     onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetChatInviteLink Returns information about an invite link. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
func (c *Client) GetChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLink, error) {
	req := &GetChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// GetChatInviteLinkCounts Returns the list of chat administrators with number of their invite links. Requires owner privileges in the chat @chat_id Chat identifier
func (c *Client) GetChatInviteLinkCounts(chatId int64) (*ChatInviteLinkCounts, error) {
	req := &GetChatInviteLinkCounts{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLinkCounts), nil
}

// GetChatInviteLinkMembers Returns chat members joined a chat via an invite link. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) GetChatInviteLinkMembers(chatId int64, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	req := &GetChatInviteLinkMembers{
		ChatId:                      chatId,
		InviteLink:                  inviteLink,
		OnlyWithExpiredSubscription: onlyWithExpiredSubscription,
		Limit:                       limit,
	}
	if opts != nil {
		req.OffsetMember = opts.OffsetMember
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLinkMembers), nil
}

// GetChatInviteLinks Returns invite links for a chat created by specified administrator. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
func (c *Client) GetChatInviteLinks(chatId int64, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	req := &GetChatInviteLinks{
		ChatId:           chatId,
		CreatorUserId:    creatorUserId,
		IsRevoked:        isRevoked,
		OffsetDate:       offsetDate,
		OffsetInviteLink: offsetInviteLink,
		Limit:            limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLinks), nil
}

// GetChatJoinRequests Returns pending join requests in a chat
func (c *Client) GetChatJoinRequests(chatId int64, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	req := &GetChatJoinRequests{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Query:      query,
		Limit:      limit,
	}
	if opts != nil {
		req.OffsetRequest = opts.OffsetRequest
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatJoinRequests), nil
}

// GetChatListsToAddChat Returns chat lists to which the chat can be added. This is an offline method @chat_id Chat identifier
func (c *Client) GetChatListsToAddChat(chatId int64) (*ChatLists, error) {
	req := &GetChatListsToAddChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatLists), nil
}

// GetChatMember Returns information about a single member of a chat @chat_id Chat identifier @member_id Member identifier
func (c *Client) GetChatMember(chatId int64, memberId *MessageSender) (*ChatMember, error) {
	req := &GetChatMember{
		ChatId:   chatId,
		MemberId: memberId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatMember), nil
}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date. Returns a 404 error if such message doesn't exist
func (c *Client) GetChatMessageByDate(chatId int64, date int32) (*Message, error) {
	req := &GetChatMessageByDate{
		ChatId: chatId,
		Date:   date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetChatMessageCalendar Returns information about the next messages of the specified type in the chat split by days. Returns the results in reverse chronological order. Can return partial result for the last returned day. Behavior of this method depends on the value of the option "utc_time_offset"
func (c *Client) GetChatMessageCalendar(chatId int64, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	req := &GetChatMessageCalendar{
		ChatId:        chatId,
		Filter:        filter,
		FromMessageId: fromMessageId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageCalendar), nil
}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat or its topic
func (c *Client) GetChatMessageCount(chatId int64, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	req := &GetChatMessageCount{
		ChatId:      chatId,
		Filter:      filter,
		ReturnLocal: returnLocal,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetChatMessagePosition Returns approximate 1-based position of a message among messages, which can be found by the specified filter in the chat and topic. Cannot be used in secret chats
func (c *Client) GetChatMessagePosition(chatId int64, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	req := &GetChatMessagePosition{
		ChatId:    chatId,
		Filter:    filter,
		MessageId: messageId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetChatNotificationSettingsExceptions Returns the list of chats with non-default notification settings for new messages
func (c *Client) GetChatNotificationSettingsExceptions(compareSound bool, opts *GetChatNotificationSettingsExceptionsOpts) (*Chats, error) {
	req := &GetChatNotificationSettingsExceptions{
		CompareSound: compareSound,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat. Returns a 404 error if the message doesn't exist @chat_id Identifier of the chat the message belongs to
func (c *Client) GetChatPinnedMessage(chatId int64) (*Message, error) {
	req := &GetChatPinnedMessage{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetChatPostedToChatPageStories Returns the list of stories that posted by the given chat to its chat page. If from_story_id == 0, then pinned stories are returned first.
func (c *Client) GetChatPostedToChatPageStories(chatId int64, fromStoryId int32, limit int32) (*Stories, error) {
	req := &GetChatPostedToChatPageStories{
		ChatId:      chatId,
		FromStoryId: fromStoryId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stories), nil
}

// GetChatRevenueStatistics Returns detailed revenue statistics about a chat. Currently, this method can be used only
func (c *Client) GetChatRevenueStatistics(chatId int64, isDark bool) (*ChatRevenueStatistics, error) {
	req := &GetChatRevenueStatistics{
		ChatId: chatId,
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatRevenueStatistics), nil
}

// GetChatRevenueTransactions Returns the list of revenue transactions for a chat. Currently, this method can be used only
func (c *Client) GetChatRevenueTransactions(chatId int64, offset string, limit int32) (*ChatRevenueTransactions, error) {
	req := &GetChatRevenueTransactions{
		ChatId: chatId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatRevenueTransactions), nil
}

// GetChatRevenueWithdrawalUrl Returns a URL for chat revenue withdrawal; requires owner privileges in the channel chat or the bot. Currently, this method can be used only
func (c *Client) GetChatRevenueWithdrawalUrl(chatId int64, password string) (*HttpUrl, error) {
	req := &GetChatRevenueWithdrawalUrl{
		ChatId:   chatId,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetChats Returns an ordered list of chats from the beginning of a chat list. For informational purposes only. Use loadChats and updates processing instead to maintain chat lists in a consistent state
func (c *Client) GetChats(limit int32, opts *GetChatsOpts) (*Chats, error) {
	req := &GetChats{
		Limit: limit,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id) @chat_id Chat identifier
func (c *Client) GetChatScheduledMessages(chatId int64) (*Messages, error) {
	req := &GetChatScheduledMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetChatsForChatFolderInviteLink Returns identifiers of chats from a chat folder, suitable for adding to a chat folder invite link @chat_folder_id Chat folder identifier
func (c *Client) GetChatsForChatFolderInviteLink(chatFolderId int32) (*Chats, error) {
	req := &GetChatsForChatFolderInviteLink{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatSimilarChatCount Returns approximate number of chats similar to the given chat
func (c *Client) GetChatSimilarChatCount(chatId int64, returnLocal bool) (*Count, error) {
	req := &GetChatSimilarChatCount{
		ChatId:      chatId,
		ReturnLocal: returnLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetChatSimilarChats Returns a list of chats similar to the given chat @chat_id Identifier of the target chat; must be an identifier of a channel chat
func (c *Client) GetChatSimilarChats(chatId int64) (*Chats, error) {
	req := &GetChatSimilarChats{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatSparseMessagePositions Returns sparse positions of messages of the specified type in the chat to be used for shared media scroll implementation. Returns the results in reverse chronological order (i.e., in order of decreasing message_id).
func (c *Client) GetChatSparseMessagePositions(chatId int64, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	req := &GetChatSparseMessagePositions{
		ChatId:               chatId,
		Filter:               filter,
		FromMessageId:        fromMessageId,
		Limit:                limit,
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessagePositions), nil
}

// GetChatSponsoredMessages Returns sponsored messages to be shown in a chat; for channel chats and chats with bots only @chat_id Identifier of the chat
func (c *Client) GetChatSponsoredMessages(chatId int64) (*SponsoredMessages, error) {
	req := &GetChatSponsoredMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SponsoredMessages), nil
}

// GetChatStatistics Returns detailed statistics about a chat. Currently, this method can be used only for supergroups and channels. Can be used only if supergroupFullInfo.can_get_statistics == true @chat_id Chat identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetChatStatistics(chatId int64, isDark bool) (*ChatStatistics, error) {
	req := &GetChatStatistics{
		ChatId: chatId,
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatStatistics), nil
}

// GetChatsToPostStories Returns supergroup and channel chats in which the current user has the right to post stories. The chats must be rechecked with canPostStory before actually trying to post a story there
func (c *Client) GetChatsToPostStories() (*Chats, error) {
	req := &GetChatsToPostStories{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetChatStoryAlbums Returns the list of story albums owned by the given chat @chat_id Chat identifier
func (c *Client) GetChatStoryAlbums(chatId int64) (*StoryAlbums, error) {
	req := &GetChatStoryAlbums{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbums), nil
}

// GetChatStoryInteractions Returns interactions with a story posted in a chat. Can be used only if story is posted on behalf of a chat and the user is an administrator in the chat
func (c *Client) GetChatStoryInteractions(storyPosterChatId int64, storyId int32, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	req := &GetChatStoryInteractions{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		PreferForwards:    preferForwards,
		Offset:            offset,
		Limit:             limit,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryInteractions), nil
}

// GetCloseFriends Returns all close friends of the current user
func (c *Client) GetCloseFriends() (*Users, error) {
	req := &GetCloseFriends{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// GetCollectibleItemInfo Returns information about a given collectible item that was purchased at https://fragment.com
func (c *Client) GetCollectibleItemInfo(typeField *CollectibleItemType) (*CollectibleItemInfo, error) {
	req := &GetCollectibleItemInfo{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CollectibleItemInfo), nil
}

// GetCommands Returns the list of commands supported by the bot for the given user scope and language; for bots only
func (c *Client) GetCommands(languageCode string, opts *GetCommandsOpts) (*BotCommands, error) {
	req := &GetCommands{
		LanguageCode: languageCode,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotCommands), nil
}

// GetConnectedAffiliateProgram Returns an affiliate program that was connected to the given affiliate by identifier of the bot that created the program
func (c *Client) GetConnectedAffiliateProgram(affiliate *AffiliateType, botUserId int64) (*ConnectedAffiliateProgram, error) {
	req := &GetConnectedAffiliateProgram{
		Affiliate: affiliate,
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectedAffiliateProgram), nil
}

// GetConnectedAffiliatePrograms Returns affiliate programs that were connected to the given affiliate
func (c *Client) GetConnectedAffiliatePrograms(affiliate *AffiliateType, offset string, limit int32) (*ConnectedAffiliatePrograms, error) {
	req := &GetConnectedAffiliatePrograms{
		Affiliate: affiliate,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectedAffiliatePrograms), nil
}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (c *Client) GetConnectedWebsites() (*ConnectedWebsites, error) {
	req := &GetConnectedWebsites{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectedWebsites), nil
}

// GetContacts Returns all contacts of the user
func (c *Client) GetContacts() (*Users, error) {
	req := &GetContacts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// GetCountries Returns information about existing countries. Can be called before authorization
func (c *Client) GetCountries() (*Countries, error) {
	req := &GetCountries{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Countries), nil
}

// GetCountryCode Uses the current IP address to find the current country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (c *Client) GetCountryCode() (*Text, error) {
	req := &GetCountryCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetCountryFlagEmoji Returns an emoji for the given country. Returns an empty string on failure. Can be called synchronously @country_code A two-letter ISO 3166-1 alpha-2 country code as received from getCountries
func (c *Client) GetCountryFlagEmoji(countryCode string) (*Text, error) {
	req := &GetCountryFlagEmoji{
		CountryCode: countryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user @type Type of the public chats to return
func (c *Client) GetCreatedPublicChats(typeField *PublicChatType) (*Chats, error) {
	req := &GetCreatedPublicChats{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetCurrentState Returns all updates needed to restore current TDLib state, i.e. all actual updateAuthorizationState/updateUser/updateNewChat and others. This is especially useful if TDLib is run in a separate process. Can be called before initialization
func (c *Client) GetCurrentState() (*Updates, error) {
	req := &GetCurrentState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Updates), nil
}

// GetCurrentWeather Returns the current weather in the given location @location The location
func (c *Client) GetCurrentWeather(location *Location) (*CurrentWeather, error) {
	req := &GetCurrentWeather{
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CurrentWeather), nil
}

// GetCustomEmojiReactionAnimations Returns TGS stickers with generic animations for custom emoji reactions
func (c *Client) GetCustomEmojiReactionAnimations() (*Stickers, error) {
	req := &GetCustomEmojiReactionAnimations{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetCustomEmojiStickers Returns the list of custom emoji stickers by their identifiers. Stickers are returned in arbitrary order. Only found stickers are returned
func (c *Client) GetCustomEmojiStickers(customEmojiIds []string) (*Stickers, error) {
	req := &GetCustomEmojiStickers{
		CustomEmojiIds: customEmojiIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetDatabaseStatistics Returns database statistics
func (c *Client) GetDatabaseStatistics() (*DatabaseStatistics, error) {
	req := &GetDatabaseStatistics{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseStatistics), nil
}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization @link The link
func (c *Client) GetDeepLinkInfo(link string) (*DeepLinkInfo, error) {
	req := &GetDeepLinkInfo{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*DeepLinkInfo), nil
}

// GetDefaultBackgroundCustomEmojiStickers Returns default list of custom emoji stickers for reply background
func (c *Client) GetDefaultBackgroundCustomEmojiStickers() (*Stickers, error) {
	req := &GetDefaultBackgroundCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetDefaultChatEmojiStatuses Returns default emoji statuses for chats
func (c *Client) GetDefaultChatEmojiStatuses() (*EmojiStatusCustomEmojis, error) {
	req := &GetDefaultChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatusCustomEmojis), nil
}

// GetDefaultChatPhotoCustomEmojiStickers Returns default list of custom emoji stickers for placing on a chat photo
func (c *Client) GetDefaultChatPhotoCustomEmojiStickers() (*Stickers, error) {
	req := &GetDefaultChatPhotoCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetDefaultEmojiStatuses Returns default emoji statuses for self status
func (c *Client) GetDefaultEmojiStatuses() (*EmojiStatusCustomEmojis, error) {
	req := &GetDefaultEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatusCustomEmojis), nil
}

// GetDefaultMessageAutoDeleteTime Returns default message auto-delete time setting for new chats
func (c *Client) GetDefaultMessageAutoDeleteTime() (*MessageAutoDeleteTime, error) {
	req := &GetDefaultMessageAutoDeleteTime{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageAutoDeleteTime), nil
}

// GetDefaultProfilePhotoCustomEmojiStickers Returns default list of custom emoji stickers for placing on a profile photo
func (c *Client) GetDefaultProfilePhotoCustomEmojiStickers() (*Stickers, error) {
	req := &GetDefaultProfilePhotoCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetDirectMessagesChatTopic Returns information about the topic in a channel direct messages chat administered by the current user
func (c *Client) GetDirectMessagesChatTopic(chatId int64, topicId int64) (*DirectMessagesChatTopic, error) {
	req := &GetDirectMessagesChatTopic{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*DirectMessagesChatTopic), nil
}

// GetDirectMessagesChatTopicHistory Returns messages in the topic in a channel direct messages chat administered by the current user. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id)
func (c *Client) GetDirectMessagesChatTopicHistory(chatId int64, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	req := &GetDirectMessagesChatTopicHistory{
		ChatId:        chatId,
		TopicId:       topicId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetDirectMessagesChatTopicMessageByDate Returns the last message sent in the topic in a channel direct messages chat administered by the current user no later than the specified date
func (c *Client) GetDirectMessagesChatTopicMessageByDate(chatId int64, topicId int64, date int32) (*Message, error) {
	req := &GetDirectMessagesChatTopicMessageByDate{
		ChatId:  chatId,
		TopicId: topicId,
		Date:    date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetDirectMessagesChatTopicRevenue Returns the total number of Telegram Stars received by the channel chat for direct messages from the given topic
func (c *Client) GetDirectMessagesChatTopicRevenue(chatId int64, topicId int64) (*StarCount, error) {
	req := &GetDirectMessagesChatTopicRevenue{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarCount), nil
}

// GetDisallowedChatEmojiStatuses Returns the list of emoji statuses, which can't be used as chat emoji status, even if they are from a sticker set with is_allowed_as_chat_emoji_status == true
func (c *Client) GetDisallowedChatEmojiStatuses() (*EmojiStatusCustomEmojis, error) {
	req := &GetDisallowedChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatusCustomEmojis), nil
}

// GetEmojiCategories Returns available emoji categories @type Type of emoji categories to return; pass null to get default emoji categories
func (c *Client) GetEmojiCategories(typeField *EmojiCategoryType) (*EmojiCategories, error) {
	req := &GetEmojiCategories{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiCategories), nil
}

// GetEmojiReaction Returns information about an emoji reaction. Returns a 404 error if the reaction is not found @emoji Text representation of the reaction
func (c *Client) GetEmojiReaction(emoji string) (*EmojiReaction, error) {
	req := &GetEmojiReaction{
		Emoji: emoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiReaction), nil
}

// GetEmojiSuggestionsUrl Returns an HTTP URL which can be used to automatically log in to the translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
func (c *Client) GetEmojiSuggestionsUrl(languageCode string) (*HttpUrl, error) {
	req := &GetEmojiSuggestionsUrl{
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetExternalLink Returns an HTTP URL which can be used to automatically authorize the current user on a website after clicking an HTTP link. Use the method getExternalLinkInfo to find whether a prior user confirmation is needed
func (c *Client) GetExternalLink(link string, allowWriteAccess bool) (*HttpUrl, error) {
	req := &GetExternalLink{
		Link:             link,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetExternalLinkInfo Returns information about an action to be done when the current user clicks an external link. Don't use this method for links from secret chats if link preview is disabled in secret chats @link The link
func (c *Client) GetExternalLinkInfo(link string) (*LoginUrlInfo, error) {
	req := &GetExternalLinkInfo{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LoginUrlInfo), nil
}

// GetFavoriteStickers Returns favorite stickers
func (c *Client) GetFavoriteStickers() (*Stickers, error) {
	req := &GetFavoriteStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetFile Returns information about a file. This is an offline method @file_id Identifier of the file to get
func (c *Client) GetFile(fileId int32) (*File, error) {
	req := &GetFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset, in bytes @file_id Identifier of the file @offset Offset from which downloaded prefix size needs to be calculated
func (c *Client) GetFileDownloadedPrefixSize(fileId int32, offset int64) (*FileDownloadedPrefixSize, error) {
	req := &GetFileDownloadedPrefixSize{
		FileId: fileId,
		Offset: offset,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileDownloadedPrefixSize), nil
}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously @mime_type The MIME type of the file
func (c *Client) GetFileExtension(mimeType string) (*Text, error) {
	req := &GetFileExtension{
		MimeType: mimeType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously @file_name The name of the file or path to the file
func (c *Client) GetFileMimeType(fileName string) (*Text, error) {
	req := &GetFileMimeType{
		FileName: fileName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetForumTopic Returns information about a topic in a forum supergroup chat or a chat with a bot with topics
func (c *Client) GetForumTopic(chatId int64, forumTopicId int32) (*ForumTopic, error) {
	req := &GetForumTopic{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ForumTopic), nil
}

// GetForumTopicDefaultIcons Returns the list of custom emoji, which can be used as forum topic icon by all users
func (c *Client) GetForumTopicDefaultIcons() (*Stickers, error) {
	req := &GetForumTopicDefaultIcons{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetForumTopicHistory Returns messages in a topic in a forum supergroup chat or a chat with a bot with topics. The messages are returned in reverse chronological order
func (c *Client) GetForumTopicHistory(chatId int64, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	req := &GetForumTopicHistory{
		ChatId:        chatId,
		ForumTopicId:  forumTopicId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetForumTopicLink Returns an HTTPS link to a topic in a forum supergroup chat. This is an offline method @chat_id Identifier of the chat @forum_topic_id Forum topic identifier
func (c *Client) GetForumTopicLink(chatId int64, forumTopicId int32) (*MessageLink, error) {
	req := &GetForumTopicLink{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageLink), nil
}

// GetForumTopics Returns found forum topics in a forum supergroup chat or a chat with a bot with topics. This is a temporary method for getting information about topic list from the server
func (c *Client) GetForumTopics(chatId int64, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	req := &GetForumTopics{
		ChatId:             chatId,
		Query:              query,
		OffsetDate:         offsetDate,
		OffsetMessageId:    offsetMessageId,
		OffsetForumTopicId: offsetForumTopicId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ForumTopics), nil
}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only @chat_id The chat that contains the message with the game @message_id Identifier of the message @user_id User identifier
func (c *Client) GetGameHighScores(chatId int64, messageId int64, userId int64) (*GameHighScores, error) {
	req := &GetGameHighScores{
		ChatId:    chatId,
		MessageId: messageId,
		UserId:    userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GameHighScores), nil
}

// GetGiftAuctionAcquiredGifts Returns the gifts that were acquired by the current user on a gift auction @gift_id Identifier of the auctioned gift
func (c *Client) GetGiftAuctionAcquiredGifts(giftId string) (*GiftAuctionAcquiredGifts, error) {
	req := &GetGiftAuctionAcquiredGifts{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftAuctionAcquiredGifts), nil
}

// GetGiftAuctionState Returns auction state for a gift @auction_id Unique identifier of the auction
func (c *Client) GetGiftAuctionState(auctionId string) (*GiftAuctionState, error) {
	req := &GetGiftAuctionState{
		AuctionId: auctionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftAuctionState), nil
}

// GetGiftChatThemes Returns available to the current user gift chat themes
func (c *Client) GetGiftChatThemes(offset string, limit int32) (*GiftChatThemes, error) {
	req := &GetGiftChatThemes{
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftChatThemes), nil
}

// GetGiftCollections Returns collections of gifts owned by the given user or chat
func (c *Client) GetGiftCollections(ownerId *MessageSender) (*GiftCollections, error) {
	req := &GetGiftCollections{
		OwnerId: ownerId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollections), nil
}

// GetGiftUpgradePreview Returns examples of possible upgraded gifts for a regular gift @gift_id Identifier of the gift
func (c *Client) GetGiftUpgradePreview(giftId string) (*GiftUpgradePreview, error) {
	req := &GetGiftUpgradePreview{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftUpgradePreview), nil
}

// GetGiftUpgradeVariants Returns all possible variants of upgraded gifts for a regular gift @gift_id Identifier of the gift
func (c *Client) GetGiftUpgradeVariants(giftId string) (*GiftUpgradeVariants, error) {
	req := &GetGiftUpgradeVariants{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftUpgradeVariants), nil
}

// GetGiveawayInfo Returns information about a giveaway
func (c *Client) GetGiveawayInfo(chatId int64, messageId int64) (*GiveawayInfo, error) {
	req := &GetGiveawayInfo{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiveawayInfo), nil
}

// GetGreetingStickers Returns greeting stickers from regular sticker sets that can be used for the start page of other users
func (c *Client) GetGreetingStickers() (*Stickers, error) {
	req := &GetGreetingStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetGrossingWebAppBots Returns the most grossing Web App bots
func (c *Client) GetGrossingWebAppBots(offset string, limit int32) (*FoundUsers, error) {
	req := &GetGrossingWebAppBots{
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundUsers), nil
}

// GetGroupCall Returns information about a group call @group_call_id Group call identifier
func (c *Client) GetGroupCall(groupCallId int32) (*GroupCall, error) {
	req := &GetGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCall), nil
}

// GetGroupCallParticipants Returns information about participants of a non-joined group call that is not bound to a chat
func (c *Client) GetGroupCallParticipants(inputGroupCall *InputGroupCall, limit int32) (*GroupCallParticipants, error) {
	req := &GetGroupCallParticipants{
		InputGroupCall: inputGroupCall,
		Limit:          limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallParticipants), nil
}

// GetGroupCallStreams Returns information about available streams in a video chat or a live story @group_call_id Group call identifier
func (c *Client) GetGroupCallStreams(groupCallId int32) (*GroupCallStreams, error) {
	req := &GetGroupCallStreams{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallStreams), nil
}

// GetGroupCallStreamSegment Returns a file with a segment of a video chat or live story in a modified OGG format for audio or MPEG-4 format for video
func (c *Client) GetGroupCallStreamSegment(groupCallId int32, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	req := &GetGroupCallStreamSegment{
		GroupCallId: groupCallId,
		TimeOffset:  timeOffset,
		Scale:       scale,
		ChannelId:   channelId,
	}
	if opts != nil {
		req.VideoQuality = opts.VideoQuality
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Data), nil
}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
func (c *Client) GetGroupsInCommon(userId int64, offsetChatId int64, limit int32) (*Chats, error) {
	req := &GetGroupsInCommon{
		UserId:       userId,
		OffsetChatId: offsetChatId,
		Limit:        limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetImportedContactCount Returns the total number of imported contacts
func (c *Client) GetImportedContactCount() (*Count, error) {
	req := &GetImportedContactCount{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Count), nil
}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives the error "CHANNELS_TOO_MUCH". Also, the limit can be increased with Telegram Premium
func (c *Client) GetInactiveSupergroupChats() (*Chats, error) {
	req := &GetInactiveSupergroupChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only @inline_message_id Inline message identifier @user_id User identifier
func (c *Client) GetInlineGameHighScores(inlineMessageId string, userId int64) (*GameHighScores, error) {
	req := &GetInlineGameHighScores{
		InlineMessageId: inlineMessageId,
		UserId:          userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GameHighScores), nil
}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
func (c *Client) GetInlineQueryResults(botUserId int64, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	req := &GetInlineQueryResults{
		BotUserId: botUserId,
		ChatId:    chatId,
		Query:     query,
		Offset:    offset,
	}
	if opts != nil {
		req.UserLocation = opts.UserLocation
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*InlineQueryResults), nil
}

// GetInstalledBackgrounds Returns backgrounds installed by the user @for_dark_theme Pass true to order returned backgrounds for a dark theme
func (c *Client) GetInstalledBackgrounds(forDarkTheme bool) (*Backgrounds, error) {
	req := &GetInstalledBackgrounds{
		ForDarkTheme: forDarkTheme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Backgrounds), nil
}

// GetInstalledStickerSets Returns a list of installed sticker sets @sticker_type Type of the sticker sets to return
func (c *Client) GetInstalledStickerSets(stickerType *StickerType) (*StickerSets, error) {
	req := &GetInstalledStickerSets{
		StickerType: stickerType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// GetInternalLink Returns an HTTPS or a tg: link with the given type. Can be called before authorization @type Expected type of the link @is_http Pass true to create an HTTPS link (only available for some link types); pass false to create a tg: link
func (c *Client) GetInternalLink(typeField *InternalLinkType, isHttp bool) (*HttpUrl, error) {
	req := &GetInternalLink{
		TypeField: typeField,
		IsHttp:    isHttp,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetInternalLinkType Returns information about the type of internal link. Returns a 404 error if the link is not internal. Can be called before authorization @link The link
func (c *Client) GetInternalLinkType(link string) (*InternalLinkType, error) {
	req := &GetInternalLinkType{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*InternalLinkType), nil
}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously @json_value The JsonValue object
func (c *Client) GetJsonString(jsonValue *JsonValue) (*Text, error) {
	req := &GetJsonString{
		JsonValue: jsonValue,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously @json The JSON-serialized string
func (c *Client) GetJsonValue(json string) (*JsonValue, error) {
	req := &GetJsonValue{
		Json: json,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*JsonValue), nil
}

// GetKeywordEmojis Returns emojis matching the keyword. Supported only if the file database is enabled. Order of results is unspecified
func (c *Client) GetKeywordEmojis(text string, inputLanguageCodes []string) (*Emojis, error) {
	req := &GetKeywordEmojis{
		Text:               text,
		InputLanguageCodes: inputLanguageCodes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Emojis), nil
}

// GetLanguagePackInfo Returns information about a language pack. Returned language pack identifier may be different from a provided one. Can be called before authorization @language_pack_id Language pack identifier
func (c *Client) GetLanguagePackInfo(languagePackId string) (*LanguagePackInfo, error) {
	req := &GetLanguagePackInfo{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LanguagePackInfo), nil
}

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. Can be called synchronously
func (c *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackId string, key string) (*LanguagePackStringValue, error) {
	req := &GetLanguagePackString{
		LanguagePackDatabasePath: languagePackDatabasePath,
		LocalizationTarget:       localizationTarget,
		LanguagePackId:           languagePackId,
		Key:                      key,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LanguagePackStringValue), nil
}

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
func (c *Client) GetLanguagePackStrings(languagePackId string, keys []string) (*LanguagePackStrings, error) {
	req := &GetLanguagePackStrings{
		LanguagePackId: languagePackId,
		Keys:           keys,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LanguagePackStrings), nil
}

// GetLinkPreview Returns a link preview by the text of a message. Do not call this function too often. Returns a 404 error if the text has no link preview
func (c *Client) GetLinkPreview(text *FormattedText, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	req := &GetLinkPreview{
		Text: text,
	}
	if opts != nil {
		req.LinkPreviewOptions = opts.LinkPreviewOptions
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LinkPreview), nil
}

// GetLiveStoryAvailableMessageSenders Returns the list of message sender identifiers, on whose behalf messages can be sent to a live story @group_call_id Group call identifier
func (c *Client) GetLiveStoryAvailableMessageSenders(groupCallId int32) (*ChatMessageSenders, error) {
	req := &GetLiveStoryAvailableMessageSenders{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatMessageSenders), nil
}

// GetLiveStoryRtmpUrl Returns RTMP URL for streaming to a live story; requires can_post_stories administrator right for channel chats @chat_id Chat identifier
func (c *Client) GetLiveStoryRtmpUrl(chatId int64) (*RtmpUrl, error) {
	req := &GetLiveStoryRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RtmpUrl), nil
}

// GetLiveStoryStreamer Returns information about the user or the chat that streams to a live story; for live stories that aren't an RTMP stream only @group_call_id Group call identifier
func (c *Client) GetLiveStoryStreamer(groupCallId int32) (*GroupCallParticipant, error) {
	req := &GetLiveStoryStreamer{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallParticipant), nil
}

// GetLiveStoryTopDonors Returns the list of top live story donors @group_call_id Group call identifier of the live story
func (c *Client) GetLiveStoryTopDonors(groupCallId int32) (*LiveStoryDonors, error) {
	req := &GetLiveStoryTopDonors{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LiveStoryDonors), nil
}

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline method if only_local is true. Can be called before authorization @only_local Pass true to get only locally available information without sending network requests
func (c *Client) GetLocalizationTargetInfo(onlyLocal bool) (*LocalizationTargetInfo, error) {
	req := &GetLocalizationTargetInfo{
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LocalizationTargetInfo), nil
}

// GetLoginPasskeys Returns the list of passkeys allowed to be used for the login by the current user
func (c *Client) GetLoginPasskeys() (*Passkeys, error) {
	req := &GetLoginPasskeys{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Passkeys), nil
}

// GetLoginUrl Returns an HTTP URL which can be used to automatically authorize the user on a website after clicking an inline button of type inlineKeyboardButtonTypeLoginUrl.
func (c *Client) GetLoginUrl(chatId int64, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	req := &GetLoginUrl{
		ChatId:           chatId,
		MessageId:        messageId,
		ButtonId:         buttonId,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetLoginUrlInfo Returns information about a button of type inlineKeyboardButtonTypeLoginUrl. The method needs to be called when the user presses the button
func (c *Client) GetLoginUrlInfo(chatId int64, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	req := &GetLoginUrlInfo{
		ChatId:    chatId,
		MessageId: messageId,
		ButtonId:  buttonId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LoginUrlInfo), nil
}

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (c *Client) GetLogStream() (*LogStream, error) {
	req := &GetLogStream{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LogStream), nil
}

// GetLogTags Returns the list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (c *Client) GetLogTags() (*LogTags, error) {
	req := &GetLogTags{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LogTags), nil
}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously @tag Logging tag to change verbosity level
func (c *Client) GetLogTagVerbosityLevel(tag string) (*LogVerbosityLevel, error) {
	req := &GetLogTagVerbosityLevel{
		Tag: tag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LogVerbosityLevel), nil
}

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (c *Client) GetLogVerbosityLevel() (*LogVerbosityLevel, error) {
	req := &GetLogVerbosityLevel{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*LogVerbosityLevel), nil
}

// GetMainWebApp Returns information needed to open the main Web App of a bot
func (c *Client) GetMainWebApp(chatId int64, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	req := &GetMainWebApp{
		ChatId:         chatId,
		BotUserId:      botUserId,
		StartParameter: startParameter,
		Parameters:     parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MainWebApp), nil
}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
func (c *Client) GetMapThumbnailFile(location *Location, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	req := &GetMapThumbnailFile{
		Location: location,
		Zoom:     zoom,
		Width:    width,
		Height:   height,
		Scale:    scale,
		ChatId:   chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously @text The text
func (c *Client) GetMarkdownText(text *FormattedText) (*FormattedText, error) {
	req := &GetMarkdownText{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// GetMe Returns the current user
func (c *Client) GetMe() (*User, error) {
	req := &GetMe{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// GetMenuButton Returns menu button set by the bot for the given user; for bots only @user_id Identifier of the user or 0 to get the default menu button
func (c *Client) GetMenuButton(userId int64) (*BotMenuButton, error) {
	req := &GetMenuButton{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BotMenuButton), nil
}

// GetMessage Returns information about a message. Returns a 404 error if the message doesn't exist
func (c *Client) GetMessage(chatId int64, messageId int64) (*Message, error) {
	req := &GetMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetMessageAddedReactions Returns reactions added for a message, along with their sender
func (c *Client) GetMessageAddedReactions(chatId int64, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	req := &GetMessageAddedReactions{
		ChatId:    chatId,
		MessageId: messageId,
		Offset:    offset,
		Limit:     limit,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AddedReactions), nil
}

// GetMessageAuthor Returns information about actual author of a message sent on behalf of a channel. The method can be called if messageProperties.can_get_author == true
func (c *Client) GetMessageAuthor(chatId int64, messageId int64) (*User, error) {
	req := &GetMessageAuthor{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// GetMessageAvailableReactions Returns reactions, which can be added to a message. The list can change after updateActiveEmojiReactions, updateChatAvailableReactions for the chat, or updateMessageInteractionInfo for the message
func (c *Client) GetMessageAvailableReactions(chatId int64, messageId int64, rowSize int32) (*AvailableReactions, error) {
	req := &GetMessageAvailableReactions{
		ChatId:    chatId,
		MessageId: messageId,
		RowSize:   rowSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AvailableReactions), nil
}

// GetMessageEffect Returns information about a message effect. Returns a 404 error if the effect is not found @effect_id Unique identifier of the effect
func (c *Client) GetMessageEffect(effectId string) (*MessageEffect, error) {
	req := &GetMessageEffect{
		EffectId: effectId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageEffect), nil
}

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only if messageProperties.can_get_embedding_code
func (c *Client) GetMessageEmbeddingCode(chatId int64, messageId int64, forAlbum bool) (*Text, error) {
	req := &GetMessageEmbeddingCode{
		ChatId:    chatId,
		MessageId: messageId,
		ForAlbum:  forAlbum,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetMessageFileType Returns information about a file with messages exported from another application @message_file_head Beginning of the message file; up to 100 first lines
func (c *Client) GetMessageFileType(messageFileHead string) (*MessageFileType, error) {
	req := &GetMessageFileType{
		MessageFileHead: messageFileHead,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageFileType), nil
}

// GetMessageImportConfirmationText Returns a confirmation text to be shown to the user before starting message import
func (c *Client) GetMessageImportConfirmationText(chatId int64) (*Text, error) {
	req := &GetMessageImportConfirmationText{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only if messageProperties.can_get_link, or if messageProperties.can_get_media_timestamp_links and a media timestamp link is generated. This is an offline method
func (c *Client) GetMessageLink(chatId int64, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	req := &GetMessageLink{
		ChatId:          chatId,
		MessageId:       messageId,
		MediaTimestamp:  mediaTimestamp,
		ForAlbum:        forAlbum,
		InMessageThread: inMessageThread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageLink), nil
}

// GetMessageLinkInfo Returns information about a public or private message link. Can be called for any internal link of the type internalLinkTypeMessage @url The message link
func (c *Client) GetMessageLinkInfo(url string) (*MessageLinkInfo, error) {
	req := &GetMessageLinkInfo{
		Url: url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageLinkInfo), nil
}

// GetMessageLocally Returns information about a message, if it is available without sending network request. Returns a 404 error if message isn't available locally. This is an offline method
func (c *Client) GetMessageLocally(chatId int64, messageId int64) (*Message, error) {
	req := &GetMessageLocally{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetMessageProperties Returns properties of a message. This is an offline method @chat_id Chat identifier @message_id Identifier of the message
func (c *Client) GetMessageProperties(chatId int64, messageId int64) (*MessageProperties, error) {
	req := &GetMessageProperties{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageProperties), nil
}

// GetMessagePublicForwards Returns forwarded copies of a channel message to different public channels and public reposts as a story. Can be used only if messageProperties.can_get_statistics == true. For optimal performance, the number of returned messages and stories is chosen by TDLib
func (c *Client) GetMessagePublicForwards(chatId int64, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	req := &GetMessagePublicForwards{
		ChatId:    chatId,
		MessageId: messageId,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicForwards), nil
}

// GetMessageReadDate Returns read date of a recent outgoing message in a private chat. The method can be called if messageProperties.can_get_read_date == true
func (c *Client) GetMessageReadDate(chatId int64, messageId int64) (*MessageReadDate, error) {
	req := &GetMessageReadDate{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageReadDate), nil
}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result @chat_id Identifier of the chat the messages belong to @message_ids Identifiers of the messages to get
func (c *Client) GetMessages(chatId int64, messageIds []int64) (*Messages, error) {
	req := &GetMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if messageProperties.can_get_statistics == true @chat_id Chat identifier @message_id Message identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetMessageStatistics(chatId int64, messageId int64, isDark bool) (*MessageStatistics, error) {
	req := &GetMessageStatistics{
		ChatId:    chatId,
		MessageId: messageId,
		IsDark:    isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageStatistics), nil
}

// GetMessageThread Returns information about a message thread. Can be used only if messageProperties.can_get_message_thread == true @chat_id Chat identifier @message_id Identifier of the message
func (c *Client) GetMessageThread(chatId int64, messageId int64) (*MessageThreadInfo, error) {
	req := &GetMessageThread{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageThreadInfo), nil
}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if messageProperties.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup.
func (c *Client) GetMessageThreadHistory(chatId int64, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	req := &GetMessageThreadHistory{
		ChatId:        chatId,
		MessageId:     messageId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetMessageViewers Returns viewers of a recent outgoing message in a basic group or a supergroup chat. For video notes and voice notes only users, opened content of the message, are returned. The method can be called if messageProperties.can_get_viewers == true
func (c *Client) GetMessageViewers(chatId int64, messageId int64) (*MessageViewers, error) {
	req := &GetMessageViewers{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageViewers), nil
}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization @only_current Pass true to get statistics only for the current library launch
func (c *Client) GetNetworkStatistics(onlyCurrent bool) (*NetworkStatistics, error) {
	req := &GetNetworkStatistics{
		OnlyCurrent: onlyCurrent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*NetworkStatistics), nil
}

// GetNewChatPrivacySettings Returns privacy settings for new chat creation
func (c *Client) GetNewChatPrivacySettings() (*NewChatPrivacySettings, error) {
	req := &GetNewChatPrivacySettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*NewChatPrivacySettings), nil
}

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization. Can be called synchronously for options "version" and "commit_hash"
func (c *Client) GetOption(name string) (*OptionValue, error) {
	req := &GetOption{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*OptionValue), nil
}

// GetOwnedBots Returns the list of bots owned by the current user
func (c *Client) GetOwnedBots() (*Users, error) {
	req := &GetOwnedBots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// GetOwnedStickerSets Returns sticker sets owned by the current user
func (c *Client) GetOwnedStickerSets(offsetStickerSetId string, limit int32) (*StickerSets, error) {
	req := &GetOwnedStickerSets{
		OffsetStickerSetId: offsetStickerSetId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// GetPaidMessageRevenue Returns the total number of Telegram Stars received by the current user for paid messages from the given user @user_id Identifier of the user
func (c *Client) GetPaidMessageRevenue(userId int64) (*StarCount, error) {
	req := &GetPaidMessageRevenue{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarCount), nil
}

// GetPasskeyParameters Returns parameters for creating of a new passkey as JSON-serialized string
func (c *Client) GetPasskeyParameters() (*Text, error) {
	req := &GetPasskeyParameters{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetPassportAuthorizationForm Returns a Telegram Passport authorization form for sharing data with a service
func (c *Client) GetPassportAuthorizationForm(botUserId int64, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	req := &GetPassportAuthorizationForm{
		BotUserId: botUserId,
		Scope:     scope,
		PublicKey: publicKey,
		Nonce:     nonce,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PassportAuthorizationForm), nil
}

// GetPassportAuthorizationFormAvailableElements Returns already available Telegram Passport elements suitable for completing a Telegram Passport authorization form. Result can be received only once for each authorization form
func (c *Client) GetPassportAuthorizationFormAvailableElements(authorizationFormId int32, password string) (*PassportElementsWithErrors, error) {
	req := &GetPassportAuthorizationFormAvailableElements{
		AuthorizationFormId: authorizationFormId,
		Password:            password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PassportElementsWithErrors), nil
}

// GetPassportElement Returns one of the available Telegram Passport elements @type Telegram Passport element type @password The 2-step verification password of the current user
func (c *Client) GetPassportElement(typeField *PassportElementType, password string) (*PassportElement, error) {
	req := &GetPassportElement{
		TypeField: typeField,
		Password:  password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PassportElement), nil
}

// GetPasswordState Returns the current state of 2-step verification
func (c *Client) GetPasswordState() (*PasswordState, error) {
	req := &GetPasswordState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// GetPaymentForm Returns an invoice payment form. This method must be called when the user presses inline button of the type inlineKeyboardButtonTypeBuy, or wants to buy access to media in a messagePaidMedia message
func (c *Client) GetPaymentForm(inputInvoice *InputInvoice, opts *GetPaymentFormOpts) (*PaymentForm, error) {
	req := &GetPaymentForm{
		InputInvoice: inputInvoice,
	}
	if opts != nil {
		req.Theme = opts.Theme
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PaymentForm), nil
}

// GetPaymentReceipt Returns information about a successful payment @chat_id Chat identifier of the messagePaymentSuccessful message @message_id Message identifier
func (c *Client) GetPaymentReceipt(chatId int64, messageId int64) (*PaymentReceipt, error) {
	req := &GetPaymentReceipt{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PaymentReceipt), nil
}

// GetPhoneNumberInfo Returns information about a phone number by its prefix. Can be called before authorization @phone_number_prefix The phone number prefix
func (c *Client) GetPhoneNumberInfo(phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	req := &GetPhoneNumberInfo{
		PhoneNumberPrefix: phoneNumberPrefix,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PhoneNumberInfo), nil
}

// GetPhoneNumberInfoSync Returns information about a phone number by its prefix synchronously. getCountries must be called at least once after changing localization to the specified language if properly localized country information is expected. Can be called synchronously
func (c *Client) GetPhoneNumberInfoSync(languageCode string, phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	req := &GetPhoneNumberInfoSync{
		LanguageCode:      languageCode,
		PhoneNumberPrefix: phoneNumberPrefix,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PhoneNumberInfo), nil
}

// GetPollVoters Returns message senders voted for the specified option in a non-anonymous polls. For optimal performance, the number of returned users is chosen by TDLib
func (c *Client) GetPollVoters(chatId int64, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	req := &GetPollVoters{
		ChatId:    chatId,
		MessageId: messageId,
		OptionId:  optionId,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageSenders), nil
}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which must be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown @country_code A two-letter ISO 3166-1 alpha-2 country code
func (c *Client) GetPreferredCountryLanguage(countryCode string) (*Text, error) {
	req := &GetPreferredCountryLanguage{
		CountryCode: countryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetPremiumFeatures Returns information about features, available to Premium users @source Source of the request; pass null if the method is called from some non-standard source
func (c *Client) GetPremiumFeatures(source *PremiumSource) (*PremiumFeatures, error) {
	req := &GetPremiumFeatures{
		Source: source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumFeatures), nil
}

// GetPremiumGiftPaymentOptions Returns available options for gifting Telegram Premium to a user
func (c *Client) GetPremiumGiftPaymentOptions() (*PremiumGiftPaymentOptions, error) {
	req := &GetPremiumGiftPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumGiftPaymentOptions), nil
}

// GetPremiumGiveawayPaymentOptions Returns available options for creating of Telegram Premium giveaway or manual distribution of Telegram Premium among chat members
func (c *Client) GetPremiumGiveawayPaymentOptions(boostedChatId int64) (*PremiumGiveawayPaymentOptions, error) {
	req := &GetPremiumGiveawayPaymentOptions{
		BoostedChatId: boostedChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumGiveawayPaymentOptions), nil
}

// GetPremiumInfoSticker Returns the sticker to be used as representation of the Telegram Premium subscription @month_count Number of months the Telegram Premium subscription will be active
func (c *Client) GetPremiumInfoSticker(monthCount int32) (*Sticker, error) {
	req := &GetPremiumInfoSticker{
		MonthCount: monthCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Sticker), nil
}

// GetPremiumLimit Returns information about a limit, increased for Premium users. Returns a 404 error if the limit is unknown @limit_type Type of the limit
func (c *Client) GetPremiumLimit(limitType *PremiumLimitType) (*PremiumLimit, error) {
	req := &GetPremiumLimit{
		LimitType: limitType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumLimit), nil
}

// GetPremiumState Returns state of Telegram Premium subscription and promotion videos for Premium features
func (c *Client) GetPremiumState() (*PremiumState, error) {
	req := &GetPremiumState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PremiumState), nil
}

// GetPremiumStickerExamples Returns examples of premium stickers for demonstration purposes
func (c *Client) GetPremiumStickerExamples() (*Stickers, error) {
	req := &GetPremiumStickerExamples{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetPremiumStickers Returns premium stickers from regular sticker sets @limit The maximum number of stickers to be returned; 0-100
func (c *Client) GetPremiumStickers(limit int32) (*Stickers, error) {
	req := &GetPremiumStickers{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetPreparedInlineMessage Saves an inline message to be sent by the given user
func (c *Client) GetPreparedInlineMessage(botUserId int64, preparedMessageId string) (*PreparedInlineMessage, error) {
	req := &GetPreparedInlineMessage{
		BotUserId:         botUserId,
		PreparedMessageId: preparedMessageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PreparedInlineMessage), nil
}

// GetProxies Returns the list of proxies that are currently set up. Can be called before authorization
func (c *Client) GetProxies() (*Proxies, error) {
	req := &GetProxies{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Proxies), nil
}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) GetProxyLink(proxyId int32) (*HttpUrl, error) {
	req := &GetProxyLink{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetPublicPostSearchLimits Checks public post search limits without actually performing the search @query Query that will be searched for
func (c *Client) GetPublicPostSearchLimits(query string) (*PublicPostSearchLimits, error) {
	req := &GetPublicPostSearchLimits{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicPostSearchLimits), nil
}

// GetPushReceiverId Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously @payload JSON-encoded push notification payload
func (c *Client) GetPushReceiverId(payload string) (*PushReceiverId, error) {
	req := &GetPushReceiverId{
		Payload: payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PushReceiverId), nil
}

// GetReadDatePrivacySettings Returns privacy settings for message read date
func (c *Client) GetReadDatePrivacySettings() (*ReadDatePrivacySettings, error) {
	req := &GetReadDatePrivacySettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReadDatePrivacySettings), nil
}

// GetReceivedGift Returns information about a received gift @received_gift_id Identifier of the gift
func (c *Client) GetReceivedGift(receivedGiftId string) (*ReceivedGift, error) {
	req := &GetReceivedGift{
		ReceivedGiftId: receivedGiftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReceivedGift), nil
}

// GetReceivedGifts Returns gifts received by the given user or chat
func (c *Client) GetReceivedGifts(businessConnectionId string, ownerId *MessageSender, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	req := &GetReceivedGifts{
		BusinessConnectionId: businessConnectionId,
		OwnerId:              ownerId,
		CollectionId:         collectionId,
		ExcludeUnsaved:       excludeUnsaved,
		ExcludeSaved:         excludeSaved,
		ExcludeUnlimited:     excludeUnlimited,
		ExcludeUpgradable:    excludeUpgradable,
		ExcludeNonUpgradable: excludeNonUpgradable,
		ExcludeUpgraded:      excludeUpgraded,
		ExcludeWithoutColors: excludeWithoutColors,
		ExcludeHosted:        excludeHosted,
		SortByPrice:          sortByPrice,
		Offset:               offset,
		Limit:                limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReceivedGifts), nil
}

// GetRecentEmojiStatuses Returns recent emoji statuses for self status
func (c *Client) GetRecentEmojiStatuses() (*EmojiStatuses, error) {
	req := &GetRecentEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatuses), nil
}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (c *Client) GetRecentInlineBots() (*Users, error) {
	req := &GetRecentInlineBots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// GetRecentlyOpenedChats Returns recently opened chats. This is an offline method. Returns chats in the order of last opening @limit The maximum number of chats to be returned
func (c *Client) GetRecentlyOpenedChats(limit int32) (*Chats, error) {
	req := &GetRecentlyOpenedChats{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetRecentlyVisitedTMeUrls Returns t.me URLs recently visited by a newly registered user @referrer Google Play referrer to identify the user
func (c *Client) GetRecentlyVisitedTMeUrls(referrer string) (*TMeUrls, error) {
	req := &GetRecentlyVisitedTMeUrls{
		Referrer: referrer,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TMeUrls), nil
}

// GetRecentStickers Returns a list of recently used stickers @is_attached Pass true to return stickers and masks that were recently attached to photos or video files; pass false to return recently sent stickers
func (c *Client) GetRecentStickers(isAttached bool) (*Stickers, error) {
	req := &GetRecentStickers{
		IsAttached: isAttached,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetRecommendedChatFolders Returns recommended chat folders for the current user
func (c *Client) GetRecommendedChatFolders() (*RecommendedChatFolders, error) {
	req := &GetRecommendedChatFolders{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RecommendedChatFolders), nil
}

// GetRecommendedChats Returns a list of channel chats recommended to the current user
func (c *Client) GetRecommendedChats() (*Chats, error) {
	req := &GetRecommendedChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetRecoveryEmailAddress Returns a 2-step verification recovery email address that was previously set up. This method can be used to verify a password provided by the user @password The 2-step verification password for the current user
func (c *Client) GetRecoveryEmailAddress(password string) (*RecoveryEmailAddress, error) {
	req := &GetRecoveryEmailAddress{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RecoveryEmailAddress), nil
}

// GetRemoteFile Returns information about a file by its remote identifier. This is an offline method. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user.
func (c *Client) GetRemoteFile(remoteFileId string, opts *GetRemoteFileOpts) (*File, error) {
	req := &GetRemoteFile{
		RemoteFileId: remoteFileId,
	}
	if opts != nil {
		req.FileType = opts.FileType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// GetRepliedMessage Returns information about a non-bundled message that is replied by a given message. Also, returns the pinned message for messagePinMessage,
func (c *Client) GetRepliedMessage(chatId int64, messageId int64) (*Message, error) {
	req := &GetRepliedMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetSavedAnimations Returns saved animations
func (c *Client) GetSavedAnimations() (*Animations, error) {
	req := &GetSavedAnimations{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Animations), nil
}

// GetSavedMessagesTags Returns tags used in Saved Messages or a Saved Messages topic
func (c *Client) GetSavedMessagesTags(savedMessagesTopicId int64) (*SavedMessagesTags, error) {
	req := &GetSavedMessagesTags{
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SavedMessagesTags), nil
}

// GetSavedMessagesTopicHistory Returns messages in a Saved Messages topic. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id)
func (c *Client) GetSavedMessagesTopicHistory(savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	req := &GetSavedMessagesTopicHistory{
		SavedMessagesTopicId: savedMessagesTopicId,
		FromMessageId:        fromMessageId,
		Offset:               offset,
		Limit:                limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// GetSavedMessagesTopicMessageByDate Returns the last message sent in a Saved Messages topic no later than the specified date
func (c *Client) GetSavedMessagesTopicMessageByDate(savedMessagesTopicId int64, date int32) (*Message, error) {
	req := &GetSavedMessagesTopicMessageByDate{
		SavedMessagesTopicId: savedMessagesTopicId,
		Date:                 date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// GetSavedNotificationSound Returns saved notification sound by its identifier. Returns a 404 error if there is no saved notification sound with the specified identifier @notification_sound_id Identifier of the notification sound
func (c *Client) GetSavedNotificationSound(notificationSoundId string) (*NotificationSounds, error) {
	req := &GetSavedNotificationSound{
		NotificationSoundId: notificationSoundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*NotificationSounds), nil
}

// GetSavedNotificationSounds Returns the list of saved notification sounds. If a sound isn't in the list, then default sound needs to be used
func (c *Client) GetSavedNotificationSounds() (*NotificationSounds, error) {
	req := &GetSavedNotificationSounds{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*NotificationSounds), nil
}

// GetSavedOrderInfo Returns saved order information. Returns a 404 error if there is no saved order information
func (c *Client) GetSavedOrderInfo() (*OrderInfo, error) {
	req := &GetSavedOrderInfo{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*OrderInfo), nil
}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type @scope Types of chats for which to return the notification settings information
func (c *Client) GetScopeNotificationSettings(scope *NotificationSettingsScope) (*ScopeNotificationSettings, error) {
	req := &GetScopeNotificationSettings{
		Scope: scope,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopeNotificationSettings), nil
}

// GetSearchedForTags Returns recently searched for hashtags or cashtags by their prefix @tag_prefix Prefix of hashtags or cashtags to return @limit The maximum number of items to be returned
func (c *Client) GetSearchedForTags(tagPrefix string, limit int32) (*Hashtags, error) {
	req := &GetSearchedForTags{
		TagPrefix: tagPrefix,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Hashtags), nil
}

// GetSearchSponsoredChats Returns sponsored chats to be shown in the search results @query Query the user searches for
func (c *Client) GetSearchSponsoredChats(query string) (*SponsoredChats, error) {
	req := &GetSearchSponsoredChats{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SponsoredChats), nil
}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline method @secret_chat_id Secret chat identifier
func (c *Client) GetSecretChat(secretChatId int32) (*SecretChat, error) {
	req := &GetSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SecretChat), nil
}

// GetStakeDiceState Returns the current state of stake dice
func (c *Client) GetStakeDiceState() (*StakeDiceState, error) {
	req := &GetStakeDiceState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StakeDiceState), nil
}

// GetStarAdAccountUrl Returns a URL for a Telegram Ad platform account that can be used to set up advertisements for the chat paid in the owned Telegram Stars
func (c *Client) GetStarAdAccountUrl(ownerId *MessageSender) (*HttpUrl, error) {
	req := &GetStarAdAccountUrl{
		OwnerId: ownerId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetStarGiftPaymentOptions Returns available options for Telegram Stars gifting @user_id Identifier of the user that will receive Telegram Stars; pass 0 to get options for an unspecified user
func (c *Client) GetStarGiftPaymentOptions(userId int64) (*StarPaymentOptions, error) {
	req := &GetStarGiftPaymentOptions{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarPaymentOptions), nil
}

// GetStarGiveawayPaymentOptions Returns available options for Telegram Star giveaway creation
func (c *Client) GetStarGiveawayPaymentOptions() (*StarGiveawayPaymentOptions, error) {
	req := &GetStarGiveawayPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarGiveawayPaymentOptions), nil
}

// GetStarPaymentOptions Returns available options for Telegram Stars purchase
func (c *Client) GetStarPaymentOptions() (*StarPaymentOptions, error) {
	req := &GetStarPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarPaymentOptions), nil
}

// GetStarRevenueStatistics Returns detailed Telegram Star revenue statistics
func (c *Client) GetStarRevenueStatistics(ownerId *MessageSender, isDark bool) (*StarRevenueStatistics, error) {
	req := &GetStarRevenueStatistics{
		OwnerId: ownerId,
		IsDark:  isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarRevenueStatistics), nil
}

// GetStarSubscriptions Returns the list of Telegram Star subscriptions for the current user
func (c *Client) GetStarSubscriptions(onlyExpiring bool, offset string) (*StarSubscriptions, error) {
	req := &GetStarSubscriptions{
		OnlyExpiring: onlyExpiring,
		Offset:       offset,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarSubscriptions), nil
}

// GetStarTransactions Returns the list of Telegram Star transactions for the specified owner
func (c *Client) GetStarTransactions(ownerId *MessageSender, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	req := &GetStarTransactions{
		OwnerId:        ownerId,
		SubscriptionId: subscriptionId,
		Offset:         offset,
		Limit:          limit,
	}
	if opts != nil {
		req.Direction = opts.Direction
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StarTransactions), nil
}

// GetStarWithdrawalUrl Returns a URL for Telegram Star withdrawal
func (c *Client) GetStarWithdrawalUrl(ownerId *MessageSender, starCount int64, password string) (*HttpUrl, error) {
	req := &GetStarWithdrawalUrl{
		OwnerId:   ownerId,
		StarCount: starCount,
		Password:  password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph @chat_id Chat identifier @token The token for graph loading @x X-value for zoomed in graph or 0 otherwise
func (c *Client) GetStatisticalGraph(chatId int64, token string, x int64) (*StatisticalGraph, error) {
	req := &GetStatisticalGraph{
		ChatId: chatId,
		Token:  token,
		X:      x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StatisticalGraph), nil
}

// GetStickerEmojis Returns emoji corresponding to a sticker. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object @sticker Sticker file identifier
func (c *Client) GetStickerEmojis(sticker *InputFile) (*Emojis, error) {
	req := &GetStickerEmojis{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Emojis), nil
}

// GetStickerOutline Returns outline of a sticker. This is an offline method. Returns a 404 error if the outline isn't known
func (c *Client) GetStickerOutline(stickerFileId int32, forAnimatedEmoji bool, forClickedAnimatedEmojiMessage bool) (*Outline, error) {
	req := &GetStickerOutline{
		StickerFileId:                  stickerFileId,
		ForAnimatedEmoji:               forAnimatedEmoji,
		ForClickedAnimatedEmojiMessage: forClickedAnimatedEmojiMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Outline), nil
}

// GetStickerOutlineSvgPath Returns outline of a sticker as an SVG path. This is an offline method. Returns an empty string if the outline isn't known
func (c *Client) GetStickerOutlineSvgPath(stickerFileId int32, forAnimatedEmoji bool, forClickedAnimatedEmojiMessage bool) (*Text, error) {
	req := &GetStickerOutlineSvgPath{
		StickerFileId:                  stickerFileId,
		ForAnimatedEmoji:               forAnimatedEmoji,
		ForClickedAnimatedEmojiMessage: forClickedAnimatedEmojiMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetStickers Returns stickers from the installed sticker sets that correspond to any of the given emoji or can be found by sticker-specific keywords. If the query is non-empty, then favorite, recently used or trending stickers may also be returned
func (c *Client) GetStickers(stickerType *StickerType, query string, limit int32, chatId int64) (*Stickers, error) {
	req := &GetStickers{
		StickerType: stickerType,
		Query:       query,
		Limit:       limit,
		ChatId:      chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// GetStickerSet Returns information about a sticker set by its identifier @set_id Identifier of the sticker set
func (c *Client) GetStickerSet(setId string) (*StickerSet, error) {
	req := &GetStickerSet{
		SetId: setId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSet), nil
}

// GetStickerSetName Returns name of a sticker set by its identifier @set_id Identifier of the sticker set
func (c *Client) GetStickerSetName(setId string) (*Text, error) {
	req := &GetStickerSetName{
		SetId: setId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetStorageStatistics Returns storage usage statistics. Can be called before authorization
func (c *Client) GetStorageStatistics(chatLimit int32) (*StorageStatistics, error) {
	req := &GetStorageStatistics{
		ChatLimit: chatLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StorageStatistics), nil
}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics. Can be called before authorization
func (c *Client) GetStorageStatisticsFast() (*StorageStatisticsFast, error) {
	req := &GetStorageStatisticsFast{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StorageStatisticsFast), nil
}

// GetStory Returns a story
func (c *Client) GetStory(storyPosterChatId int64, storyId int32, onlyLocal bool) (*Story, error) {
	req := &GetStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		OnlyLocal:         onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Story), nil
}

// GetStoryAlbumStories Returns the list of stories added to the given story album. For optimal performance, the number of returned stories is chosen by TDLib
func (c *Client) GetStoryAlbumStories(chatId int64, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	req := &GetStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		Offset:       offset,
		Limit:        limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stories), nil
}

// GetStoryAvailableReactions Returns reactions, which can be chosen for a story @row_size Number of reaction per row, 5-25
func (c *Client) GetStoryAvailableReactions(rowSize int32) (*AvailableReactions, error) {
	req := &GetStoryAvailableReactions{
		RowSize: rowSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AvailableReactions), nil
}

// GetStoryInteractions Returns interactions with a story. The method can be called only for stories posted on behalf of the current user
func (c *Client) GetStoryInteractions(storyId int32, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	req := &GetStoryInteractions{
		StoryId:            storyId,
		Query:              query,
		OnlyContacts:       onlyContacts,
		PreferForwards:     preferForwards,
		PreferWithReaction: preferWithReaction,
		Offset:             offset,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryInteractions), nil
}

// GetStoryNotificationSettingsExceptions Returns the list of chats with non-default notification settings for stories
func (c *Client) GetStoryNotificationSettingsExceptions() (*Chats, error) {
	req := &GetStoryNotificationSettingsExceptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetStoryPublicForwards Returns forwards of a story as a message to public chats and reposts by public channels. Can be used only if the story is posted on behalf of the current user or story.can_get_statistics == true.
func (c *Client) GetStoryPublicForwards(storyPosterChatId int64, storyId int32, offset string, limit int32) (*PublicForwards, error) {
	req := &GetStoryPublicForwards{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		Offset:            offset,
		Limit:             limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicForwards), nil
}

// GetStoryStatistics Returns detailed statistics about a story. Can be used only if story.can_get_statistics == true @chat_id Chat identifier @story_id Story identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetStoryStatistics(chatId int64, storyId int32, isDark bool) (*StoryStatistics, error) {
	req := &GetStoryStatistics{
		ChatId:  chatId,
		StoryId: storyId,
		IsDark:  isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryStatistics), nil
}

// GetSuggestedFileName Returns suggested name for saving a file in a given directory @file_id Identifier of the file @directory Directory in which the file is expected to be saved
func (c *Client) GetSuggestedFileName(fileId int32, directory string) (*Text, error) {
	req := &GetSuggestedFileName{
		FileId:    fileId,
		Directory: directory,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetSuggestedStickerSetName Returns a suggested name for a new sticker set with a given title @title Sticker set title; 1-64 characters
func (c *Client) GetSuggestedStickerSetName(title string) (*Text, error) {
	req := &GetSuggestedStickerSetName{
		Title: title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group.
func (c *Client) GetSuitableDiscussionChats() (*Chats, error) {
	req := &GetSuitableDiscussionChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetSuitablePersonalChats Returns a list of channel chats, which can be used as a personal chat
func (c *Client) GetSuitablePersonalChats() (*Chats, error) {
	req := &GetSuitablePersonalChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline method if the current user is not a bot @supergroup_id Supergroup or channel identifier
func (c *Client) GetSupergroup(supergroupId int64) (*Supergroup, error) {
	req := &GetSupergroup{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Supergroup), nil
}

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute @supergroup_id Supergroup or channel identifier
func (c *Client) GetSupergroupFullInfo(supergroupId int64) (*SupergroupFullInfo, error) {
	req := &GetSupergroupFullInfo{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SupergroupFullInfo), nil
}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if supergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
func (c *Client) GetSupergroupMembers(supergroupId int64, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	req := &GetSupergroupMembers{
		SupergroupId: supergroupId,
		Offset:       offset,
		Limit:        limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatMembers), nil
}

// GetSupportName Returns localized name of the Telegram support user; for Telegram support only
func (c *Client) GetSupportName() (*Text, error) {
	req := &GetSupportName{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetSupportUser Returns a user that can be contacted to get support
func (c *Client) GetSupportUser() (*User, error) {
	req := &GetSupportUser{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// GetTemporaryPasswordState Returns information about the current temporary password
func (c *Client) GetTemporaryPasswordState() (*TemporaryPasswordState, error) {
	req := &GetTemporaryPasswordState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TemporaryPasswordState), nil
}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) found in the text. Can be called synchronously @text The text in which to look for entities
func (c *Client) GetTextEntities(text string) (*TextEntities, error) {
	req := &GetTextEntities{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TextEntities), nil
}

// GetThemedChatEmojiStatuses Returns up to 8 emoji statuses, which must be shown in the emoji status list for chats
func (c *Client) GetThemedChatEmojiStatuses() (*EmojiStatusCustomEmojis, error) {
	req := &GetThemedChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatusCustomEmojis), nil
}

// GetThemedEmojiStatuses Returns up to 8 emoji statuses, which must be shown right after the default Premium Badge in the emoji status list for self status
func (c *Client) GetThemedEmojiStatuses() (*EmojiStatusCustomEmojis, error) {
	req := &GetThemedEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatusCustomEmojis), nil
}

// GetThemeParametersJsonString Converts a themeParameters object to corresponding JSON-serialized string. Can be called synchronously @theme Theme parameters to convert to JSON
func (c *Client) GetThemeParametersJsonString(theme *ThemeParameters) (*Text, error) {
	req := &GetThemeParametersJsonString{
		Theme: theme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// GetTimeZones Returns the list of supported time zones
func (c *Client) GetTimeZones() (*TimeZones, error) {
	req := &GetTimeZones{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TimeZones), nil
}

// GetTonRevenueStatistics Returns detailed Toncoin revenue statistics of the current user @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetTonRevenueStatistics(isDark bool) (*TonRevenueStatistics, error) {
	req := &GetTonRevenueStatistics{
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TonRevenueStatistics), nil
}

// GetTonTransactions Returns the list of Toncoin transactions of the current user
func (c *Client) GetTonTransactions(offset string, limit int32, opts *GetTonTransactionsOpts) (*TonTransactions, error) {
	req := &GetTonTransactions{
		Offset: offset,
		Limit:  limit,
	}
	if opts != nil {
		req.Direction = opts.Direction
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TonTransactions), nil
}

// GetTonWithdrawalUrl Returns a URL for Toncoin withdrawal from the current user's account. The user must have at least 10 toncoins to withdraw
func (c *Client) GetTonWithdrawalUrl(password string) (*HttpUrl, error) {
	req := &GetTonWithdrawalUrl{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetTopChats Returns a list of frequently used chats @category Category of chats to be returned @limit The maximum number of chats to be returned; up to 30
func (c *Client) GetTopChats(category *TopChatCategory, limit int32) (*Chats, error) {
	req := &GetTopChats{
		Category: category,
		Limit:    limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// GetTrendingStickerSets Returns a list of trending sticker sets. For optimal performance, the number of returned sticker sets is chosen by TDLib
func (c *Client) GetTrendingStickerSets(stickerType *StickerType, offset int32, limit int32) (*TrendingStickerSets, error) {
	req := &GetTrendingStickerSets{
		StickerType: stickerType,
		Offset:      offset,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TrendingStickerSets), nil
}

// GetUpgradedGift Returns information about an upgraded gift by its name @name Unique name of the upgraded gift
func (c *Client) GetUpgradedGift(name string) (*UpgradedGift, error) {
	req := &GetUpgradedGift{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UpgradedGift), nil
}

// GetUpgradedGiftEmojiStatuses Returns available upgraded gift emoji statuses for self status
func (c *Client) GetUpgradedGiftEmojiStatuses() (*EmojiStatuses, error) {
	req := &GetUpgradedGiftEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiStatuses), nil
}

// GetUpgradedGiftsPromotionalAnimation Returns promotional anumation for upgraded gifts
func (c *Client) GetUpgradedGiftsPromotionalAnimation() (*Animation, error) {
	req := &GetUpgradedGiftsPromotionalAnimation{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Animation), nil
}

// GetUpgradedGiftValueInfo Returns information about value of an upgraded gift by its name @name Unique name of the upgraded gift
func (c *Client) GetUpgradedGiftValueInfo(name string) (*UpgradedGiftValueInfo, error) {
	req := &GetUpgradedGiftValueInfo{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UpgradedGiftValueInfo), nil
}

// GetUpgradedGiftWithdrawalUrl Returns a URL for upgraded gift withdrawal in the TON blockchain as an NFT; requires owner privileges for gifts owned by a chat
func (c *Client) GetUpgradedGiftWithdrawalUrl(receivedGiftId string, password string) (*HttpUrl, error) {
	req := &GetUpgradedGiftWithdrawalUrl{
		ReceivedGiftId: receivedGiftId,
		Password:       password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetUser Returns information about a user by their identifier. This is an offline method if the current user is not a bot @user_id User identifier
func (c *Client) GetUser(userId int64) (*User, error) {
	req := &GetUser{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// GetUserChatBoosts Returns the list of boosts applied to a chat by a given user; requires administrator rights in the chat; for bots only
func (c *Client) GetUserChatBoosts(chatId int64, userId int64) (*FoundChatBoosts, error) {
	req := &GetUserChatBoosts{
		ChatId: chatId,
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundChatBoosts), nil
}

// GetUserFullInfo Returns full information about a user by their identifier @user_id User identifier
func (c *Client) GetUserFullInfo(userId int64) (*UserFullInfo, error) {
	req := &GetUserFullInfo{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserFullInfo), nil
}

// GetUserLink Returns an HTTPS link, which can be used to get information about the current user
func (c *Client) GetUserLink() (*UserLink, error) {
	req := &GetUserLink{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserLink), nil
}

// GetUserPrivacySettingRules Returns the current privacy settings @setting The privacy setting
func (c *Client) GetUserPrivacySettingRules(setting *UserPrivacySetting) (*UserPrivacySettingRules, error) {
	req := &GetUserPrivacySettingRules{
		Setting: setting,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserPrivacySettingRules), nil
}

// GetUserProfileAudios Returns the list of profile audio files of a user
func (c *Client) GetUserProfileAudios(userId int64, offset int32, limit int32) (*Audios, error) {
	req := &GetUserProfileAudios{
		UserId: userId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Audios), nil
}

// GetUserProfilePhotos Returns the profile photos of a user. Personal and public photo aren't returned
func (c *Client) GetUserProfilePhotos(userId int64, offset int32, limit int32) (*ChatPhotos, error) {
	req := &GetUserProfilePhotos{
		UserId: userId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatPhotos), nil
}

// GetUserSupportInfo Returns support information for the given user; for Telegram support only @user_id User identifier
func (c *Client) GetUserSupportInfo(userId int64) (*UserSupportInfo, error) {
	req := &GetUserSupportInfo{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserSupportInfo), nil
}

// GetVideoChatAvailableParticipants Returns the list of participant identifiers, on whose behalf a video chat in the chat can be joined @chat_id Chat identifier
func (c *Client) GetVideoChatAvailableParticipants(chatId int64) (*MessageSenders, error) {
	req := &GetVideoChatAvailableParticipants{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageSenders), nil
}

// GetVideoChatInviteLink Returns invite link to a video chat in a public chat
func (c *Client) GetVideoChatInviteLink(groupCallId int32, canSelfUnmute bool) (*HttpUrl, error) {
	req := &GetVideoChatInviteLink{
		GroupCallId:   groupCallId,
		CanSelfUnmute: canSelfUnmute,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetVideoChatRtmpUrl Returns RTMP URL for streaming to the video chat of a chat; requires can_manage_video_chats administrator right @chat_id Chat identifier
func (c *Client) GetVideoChatRtmpUrl(chatId int64) (*RtmpUrl, error) {
	req := &GetVideoChatRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RtmpUrl), nil
}

// GetVideoMessageAdvertisements Returns advertisements to be shown while a video from a message is watched. Available only if messageProperties.can_get_video_advertisements
func (c *Client) GetVideoMessageAdvertisements(chatId int64, messageId int64) (*VideoMessageAdvertisements, error) {
	req := &GetVideoMessageAdvertisements{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*VideoMessageAdvertisements), nil
}

// GetWebAppLinkUrl Returns an HTTPS URL of a Web App to open after a link of the type internalLinkTypeWebApp is clicked
func (c *Client) GetWebAppLinkUrl(chatId int64, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	req := &GetWebAppLinkUrl{
		ChatId:           chatId,
		BotUserId:        botUserId,
		WebAppShortName:  webAppShortName,
		StartParameter:   startParameter,
		AllowWriteAccess: allowWriteAccess,
		Parameters:       parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetWebAppPlaceholder Returns a default placeholder for Web Apps of a bot. This is an offline method. Returns a 404 error if the placeholder isn't known @bot_user_id Identifier of the target bot
func (c *Client) GetWebAppPlaceholder(botUserId int64) (*Outline, error) {
	req := &GetWebAppPlaceholder{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Outline), nil
}

// GetWebAppUrl Returns an HTTPS URL of a Web App to open from the side menu, a keyboardButtonTypeWebApp button, or an inlineQueryResultsButtonTypeWebApp button
func (c *Client) GetWebAppUrl(botUserId int64, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	req := &GetWebAppUrl{
		BotUserId:  botUserId,
		Url:        url,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*HttpUrl), nil
}

// GetWebPageInstantView Returns an instant view version of a web page if available. This is an offline method if only_local is true. Returns a 404 error if the web page has no instant view page
func (c *Client) GetWebPageInstantView(url string, onlyLocal bool) (*WebPageInstantView, error) {
	req := &GetWebPageInstantView{
		Url:       url,
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebPageInstantView), nil
}

// GiftPremiumWithStars Allows to buy a Telegram Premium subscription for another user with payment in Telegram Stars; for bots only
func (c *Client) GiftPremiumWithStars(userId int64, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	req := &GiftPremiumWithStars{
		UserId:     userId,
		StarCount:  starCount,
		MonthCount: monthCount,
		Text:       text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// HideContactCloseBirthdays Hides the list of contacts that have close birthdays for 24 hours
func (c *Client) HideContactCloseBirthdays() (*Ok, error) {
	req := &HideContactCloseBirthdays{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// HideSuggestedAction Hides a suggested action @action Suggested action to hide
func (c *Client) HideSuggestedAction(action *SuggestedAction) (*Ok, error) {
	req := &HideSuggestedAction{
		Action: action,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
func (c *Client) ImportContacts(contacts []*ImportedContact) (*ImportedContacts, error) {
	req := &ImportContacts{
		Contacts: contacts,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ImportedContacts), nil
}

// ImportMessages Imports messages exported from another app
func (c *Client) ImportMessages(chatId int64, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	req := &ImportMessages{
		ChatId:        chatId,
		MessageFile:   messageFile,
		AttachedFiles: attachedFiles,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// IncreaseGiftAuctionBid Increases a bid for an auction gift without changing gift text and receiver
func (c *Client) IncreaseGiftAuctionBid(giftId string, starCount int64) (*Ok, error) {
	req := &IncreaseGiftAuctionBid{
		GiftId:    giftId,
		StarCount: starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// InviteGroupCallParticipant Invites a user to an active group call; for group calls not bound to a chat only. Sends a service message of the type messageGroupCall.
func (c *Client) InviteGroupCallParticipant(groupCallId int32, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	req := &InviteGroupCallParticipant{
		GroupCallId: groupCallId,
		UserId:      userId,
		IsVideo:     isVideo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*InviteGroupCallParticipantResult), nil
}

// InviteVideoChatParticipants Invites users to an active video chat. Sends a service message of the type messageInviteVideoChatParticipants to the chat bound to the group call
func (c *Client) InviteVideoChatParticipants(groupCallId int32, userIds []int64) (*Ok, error) {
	req := &InviteVideoChatParticipants{
		GroupCallId: groupCallId,
		UserIds:     userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// IsLoginEmailAddressRequired Checks whether the current user is required to set login email address
func (c *Client) IsLoginEmailAddressRequired() (*Ok, error) {
	req := &IsLoginEmailAddressRequired{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// IsProfileAudio Checks whether a file is in the profile audio files of the current user. Returns a 404 error if it isn't @file_id Identifier of the audio file to check
func (c *Client) IsProfileAudio(fileId int32) (*Ok, error) {
	req := &IsProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// JoinChat Adds the current user as a new member to a chat. Private and secret chats can't be joined using this method. May return an error with a message "INVITE_REQUEST_SENT" if only a join request was created @chat_id Chat identifier
func (c *Client) JoinChat(chatId int64) (*Ok, error) {
	req := &JoinChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible. May return an error with a message "INVITE_REQUEST_SENT" if only a join request was created @invite_link Invite link to use
func (c *Client) JoinChatByInviteLink(inviteLink string) (*Chat, error) {
	req := &JoinChatByInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// JoinGroupCall Joins a regular group call that is not bound to a chat @input_group_call The group call to join @join_parameters Parameters to join the call
func (c *Client) JoinGroupCall(inputGroupCall *InputGroupCall, joinParameters *GroupCallJoinParameters) (*GroupCallInfo, error) {
	req := &JoinGroupCall{
		InputGroupCall: inputGroupCall,
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCallInfo), nil
}

// JoinLiveStory Joins a group call of an active live story. Returns join response payload for tgcalls
func (c *Client) JoinLiveStory(groupCallId int32, joinParameters *GroupCallJoinParameters) (*Text, error) {
	req := &JoinLiveStory{
		GroupCallId:    groupCallId,
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// JoinVideoChat Joins an active video chat. Returns join response payload for tgcalls
func (c *Client) JoinVideoChat(groupCallId int32, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	req := &JoinVideoChat{
		GroupCallId:    groupCallId,
		JoinParameters: joinParameters,
		InviteHash:     inviteHash,
	}
	if opts != nil {
		req.ParticipantId = opts.ParticipantId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// LaunchPrepaidGiveaway Launches a prepaid giveaway
func (c *Client) LaunchPrepaidGiveaway(giveawayId string, parameters *GiveawayParameters, winnerCount int32, starCount int64) (*Ok, error) {
	req := &LaunchPrepaidGiveaway{
		GiveawayId:  giveawayId,
		Parameters:  parameters,
		WinnerCount: winnerCount,
		StarCount:   starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LeaveChat Removes the current user from chat members. Private and secret chats can't be left using this method @chat_id Chat identifier
func (c *Client) LeaveChat(chatId int64) (*Ok, error) {
	req := &LeaveChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LeaveGroupCall Leaves a group call @group_call_id Group call identifier
func (c *Client) LeaveGroupCall(groupCallId int32) (*Ok, error) {
	req := &LeaveGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadActiveStories Loads more active stories from a story list. The loaded stories will be sent through updates. Active stories are sorted by
func (c *Client) LoadActiveStories(storyList *StoryList) (*Ok, error) {
	req := &LoadActiveStories{
		StoryList: storyList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadChats Loads more chats from a chat list. The loaded chats and their positions in the chat list will be sent through updates. Chats are sorted by the pair (chat.position.order, chat.id) in descending order. Returns a 404 error if all chats have been loaded
func (c *Client) LoadChats(limit int32, opts *LoadChatsOpts) (*Ok, error) {
	req := &LoadChats{
		Limit: limit,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadDirectMessagesChatTopics Loads more topics in a channel direct messages chat administered by the current user. The loaded topics will be sent through updateDirectMessagesChatTopic.
func (c *Client) LoadDirectMessagesChatTopics(chatId int64, limit int32) (*Ok, error) {
	req := &LoadDirectMessagesChatTopics{
		ChatId: chatId,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadGroupCallParticipants Loads more participants of a group call; not supported in live stories. The loaded participants will be received through updates.
func (c *Client) LoadGroupCallParticipants(groupCallId int32, limit int32) (*Ok, error) {
	req := &LoadGroupCallParticipants{
		GroupCallId: groupCallId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadQuickReplyShortcutMessages Loads quick reply messages that can be sent by a given quick reply shortcut. The loaded messages will be sent through updateQuickReplyShortcutMessages
func (c *Client) LoadQuickReplyShortcutMessages(shortcutId int32) (*Ok, error) {
	req := &LoadQuickReplyShortcutMessages{
		ShortcutId: shortcutId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadQuickReplyShortcuts Loads quick reply shortcuts created by the current user. The loaded data will be sent through updateQuickReplyShortcut and updateQuickReplyShortcuts
func (c *Client) LoadQuickReplyShortcuts() (*Ok, error) {
	req := &LoadQuickReplyShortcuts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LoadSavedMessagesTopics Loads more Saved Messages topics. The loaded topics will be sent through updateSavedMessagesTopic. Topics are sorted by their topic.order in descending order. Returns a 404 error if all topics have been loaded
func (c *Client) LoadSavedMessagesTopics(limit int32) (*Ok, error) {
	req := &LoadSavedMessagesTopics{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (c *Client) LogOut() (*Ok, error) {
	req := &LogOut{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// MarkChecklistTasksAsDone Adds tasks of a checklist in a message as done or not done
func (c *Client) MarkChecklistTasksAsDone(chatId int64, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	req := &MarkChecklistTasksAsDone{
		ChatId:                 chatId,
		MessageId:              messageId,
		MarkedAsDoneTaskIds:    markedAsDoneTaskIds,
		MarkedAsNotDoneTaskIds: markedAsNotDoneTaskIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenBotSimilarBot Informs TDLib that a bot was opened from the list of similar bots
func (c *Client) OpenBotSimilarBot(botUserId int64, openedBotUserId int64) (*Ok, error) {
	req := &OpenBotSimilarBot{
		BotUserId:       botUserId,
		OpenedBotUserId: openedBotUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenChat Informs TDLib that the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats) @chat_id Chat identifier
func (c *Client) OpenChat(chatId int64) (*Ok, error) {
	req := &OpenChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenChatSimilarChat Informs TDLib that a chat was opened from the list of similar chats. The method is independent of openChat and closeChat methods
func (c *Client) OpenChatSimilarChat(chatId int64, openedChatId int64) (*Ok, error) {
	req := &OpenChatSimilarChat{
		ChatId:       chatId,
		OpenedChatId: openedChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenGiftAuction Informs TDLib that a gift auction was opened by the user @gift_id Identifier of the gift, which auction was opened
func (c *Client) OpenGiftAuction(giftId string) (*Ok, error) {
	req := &OpenGiftAuction{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenMessageContent Informs TDLib that the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message).
func (c *Client) OpenMessageContent(chatId int64, messageId int64) (*Ok, error) {
	req := &OpenMessageContent{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenSponsoredChat Informs TDLib that the user opened a sponsored chat @sponsored_chat_unique_id Unique identifier of the sponsored chat
func (c *Client) OpenSponsoredChat(sponsoredChatUniqueId int64) (*Ok, error) {
	req := &OpenSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenStory Informs TDLib that a story is opened and is being viewed by the user
func (c *Client) OpenStory(storyPosterChatId int64, storyId int32) (*Ok, error) {
	req := &OpenStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// OpenWebApp Informs TDLib that a Web App is being opened from the attachment menu, a botMenuButton button, an internalLinkTypeAttachmentMenuBot link, or an inlineKeyboardButtonTypeWebApp button.
func (c *Client) OpenWebApp(chatId int64, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	req := &OpenWebApp{
		ChatId:     chatId,
		BotUserId:  botUserId,
		Url:        url,
		Parameters: parameters,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebAppInfo), nil
}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
func (c *Client) OptimizeStorage(size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	req := &OptimizeStorage{
		Size:                        size,
		Ttl:                         ttl,
		Count:                       count,
		ImmunityDelay:               immunityDelay,
		FileTypes:                   fileTypes,
		ChatIds:                     chatIds,
		ExcludeChatIds:              excludeChatIds,
		ReturnDeletedFileStatistics: returnDeletedFileStatistics,
		ChatLimit:                   chatLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StorageStatistics), nil
}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
func (c *Client) ParseMarkdown(text *FormattedText) (*FormattedText, error) {
	req := &ParseMarkdown{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Spoiler, CustomEmoji, BlockQuote, ExpandableBlockQuote, Code, Pre, PreCode, TextUrl
func (c *Client) ParseTextEntities(text string, parseMode *TextParseMode) (*FormattedText, error) {
	req := &ParseTextEntities{
		Text:      text,
		ParseMode: parseMode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// PinChatMessage Pins a message in a chat. A message can be pinned only if messageProperties.can_be_pinned
func (c *Client) PinChatMessage(chatId int64, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	req := &PinChatMessage{
		ChatId:              chatId,
		MessageId:           messageId,
		DisableNotification: disableNotification,
		OnlyForSelf:         onlyForSelf,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization @proxy_id Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (c *Client) PingProxy(proxyId int32) (*Seconds, error) {
	req := &PingProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Seconds), nil
}

// PlaceGiftAuctionBid Places a bid on an auction gift
func (c *Client) PlaceGiftAuctionBid(giftId string, starCount int64, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	req := &PlaceGiftAuctionBid{
		GiftId:    giftId,
		StarCount: starCount,
		UserId:    userId,
		Text:      text,
		IsPrivate: isPrivate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// PostStory Posts a new story on behalf of a chat; requires can_post_stories administrator right for supergroup and channel chats. Returns a temporary story
func (c *Client) PostStory(chatId int64, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	req := &PostStory{
		ChatId:             chatId,
		Content:            content,
		PrivacySettings:    privacySettings,
		AlbumIds:           albumIds,
		ActivePeriod:       activePeriod,
		IsPostedToChatPage: isPostedToChatPage,
		ProtectContent:     protectContent,
	}
	if opts != nil {
		req.Areas = opts.Areas
		req.Caption = opts.Caption
		req.FromStoryFullId = opts.FromStoryFullId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Story), nil
}

// PreliminaryUploadFile Preliminarily uploads a file to the cloud before sending it in a message, which can be useful for uploading of being recorded voice and video notes.
func (c *Client) PreliminaryUploadFile(file *InputFile, priority int32, opts *PreliminaryUploadFileOpts) (*File, error) {
	req := &PreliminaryUploadFile{
		File:     file,
		Priority: priority,
	}
	if opts != nil {
		req.FileType = opts.FileType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// ProcessChatFolderNewChats Process new chats added to a shareable chat folder by its owner @chat_folder_id Chat folder identifier @added_chat_ids Identifiers of the new chats, which are added to the chat folder. The chats are automatically joined if they aren't joined yet
func (c *Client) ProcessChatFolderNewChats(chatFolderId int32, addedChatIds []int64) (*Ok, error) {
	req := &ProcessChatFolderNewChats{
		ChatFolderId: chatFolderId,
		AddedChatIds: addedChatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ProcessChatJoinRequest Handles a pending join request in a chat @chat_id Chat identifier @user_id Identifier of the user that sent the request @approve Pass true to approve the request; pass false to decline it
func (c *Client) ProcessChatJoinRequest(chatId int64, userId int64, approve bool) (*Ok, error) {
	req := &ProcessChatJoinRequest{
		ChatId:  chatId,
		UserId:  userId,
		Approve: approve,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ProcessChatJoinRequests Handles all pending join requests for a given link in a chat
func (c *Client) ProcessChatJoinRequests(chatId int64, inviteLink string, approve bool) (*Ok, error) {
	req := &ProcessChatJoinRequests{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Approve:    approve,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ProcessGiftPurchaseOffer Handles a pending gift purchase offer
func (c *Client) ProcessGiftPurchaseOffer(messageId int64, accept bool) (*Ok, error) {
	req := &ProcessGiftPurchaseOffer{
		MessageId: messageId,
		Accept:    accept,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ProcessPushNotification Handles a push notification. Returns error with code 406 if the push notification is not supported and connection to the server is required to fetch new data. Can be called before authorization
func (c *Client) ProcessPushNotification(payload string) (*Ok, error) {
	req := &ProcessPushNotification{
		Payload: payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RateSpeechRecognition Rates recognized speech in a video note or a voice note message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message @is_good Pass true if the speech recognition is good
func (c *Client) RateSpeechRecognition(chatId int64, messageId int64, isGood bool) (*Ok, error) {
	req := &RateSpeechRecognition{
		ChatId:    chatId,
		MessageId: messageId,
		IsGood:    isGood,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadAllChatMentions Marks all mentions in a chat as read @chat_id Chat identifier
func (c *Client) ReadAllChatMentions(chatId int64) (*Ok, error) {
	req := &ReadAllChatMentions{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadAllChatReactions Marks all reactions in a chat as read @chat_id Chat identifier
func (c *Client) ReadAllChatReactions(chatId int64) (*Ok, error) {
	req := &ReadAllChatReactions{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadAllDirectMessagesChatTopicReactions Removes all unread reactions in the topic in a channel direct messages chat administered by the current user
func (c *Client) ReadAllDirectMessagesChatTopicReactions(chatId int64, topicId int64) (*Ok, error) {
	req := &ReadAllDirectMessagesChatTopicReactions{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadAllForumTopicMentions Marks all mentions in a topic in a forum supergroup chat as read
func (c *Client) ReadAllForumTopicMentions(chatId int64, forumTopicId int32) (*Ok, error) {
	req := &ReadAllForumTopicMentions{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadAllForumTopicReactions Marks all reactions in a topic in a forum supergroup chat or a chat with a bot with topics as read
func (c *Client) ReadAllForumTopicReactions(chatId int64, forumTopicId int32) (*Ok, error) {
	req := &ReadAllForumTopicReactions{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadBusinessMessage Reads a message on behalf of a business account; for bots only
func (c *Client) ReadBusinessMessage(businessConnectionId string, chatId int64, messageId int64) (*Ok, error) {
	req := &ReadBusinessMessage{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReadChatList Traverses all chats in a chat list and marks all messages in the chats as read @chat_list Chat list in which to mark all chats as read
func (c *Client) ReadChatList(chatList *ChatList) (*Ok, error) {
	req := &ReadChatList{
		ChatList: chatList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReaddQuickReplyShortcutMessages Readds quick reply messages which failed to add. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
func (c *Client) ReaddQuickReplyShortcutMessages(shortcutName string, messageIds []int64) (*QuickReplyMessages, error) {
	req := &ReaddQuickReplyShortcutMessages{
		ShortcutName: shortcutName,
		MessageIds:   messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*QuickReplyMessages), nil
}

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
func (c *Client) ReadFilePart(fileId int32, offset int64, count int64) (*Data, error) {
	req := &ReadFilePart{
		FileId: fileId,
		Offset: offset,
		Count:  count,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Data), nil
}

// RecognizeSpeech Recognizes speech in a video note or a voice note message
func (c *Client) RecognizeSpeech(chatId int64, messageId int64) (*Ok, error) {
	req := &RecognizeSpeech{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RecoverAuthenticationPassword Recovers the 2-step verification password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (c *Client) RecoverAuthenticationPassword(recoveryCode string, newPassword string, newHint string) (*Ok, error) {
	req := &RecoverAuthenticationPassword{
		RecoveryCode: recoveryCode,
		NewPassword:  newPassword,
		NewHint:      newHint,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RecoverPassword Recovers the 2-step verification password using a recovery code sent to an email address that was previously set up
func (c *Client) RecoverPassword(recoveryCode string, newPassword string, newHint string) (*PasswordState, error) {
	req := &RecoverPassword{
		RecoveryCode: recoveryCode,
		NewPassword:  newPassword,
		NewHint:      newHint,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// RefundStarPayment Refunds a previously done payment in Telegram Stars; for bots only
func (c *Client) RefundStarPayment(userId int64, telegramPaymentChargeId string) (*Ok, error) {
	req := &RefundStarPayment{
		UserId:                  userId,
		TelegramPaymentChargeId: telegramPaymentChargeId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription @device_token Device token @other_user_ids List of user identifiers of other users currently using the application
func (c *Client) RegisterDevice(deviceToken *DeviceToken, otherUserIds []int64) (*PushReceiverId, error) {
	req := &RegisterDevice{
		DeviceToken:  deviceToken,
		OtherUserIds: otherUserIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PushReceiverId), nil
}

// RegisterUser Finishes user registration. Works only when the current authorization state is authorizationStateWaitRegistration
func (c *Client) RegisterUser(firstName string, lastName string, disableNotification bool) (*Ok, error) {
	req := &RegisterUser{
		FirstName:           firstName,
		LastName:            lastName,
		DisableNotification: disableNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveAllFilesFromDownloads Removes all files from the file download list
func (c *Client) RemoveAllFilesFromDownloads(onlyActive bool, onlyCompleted bool, deleteFromCache bool) (*Ok, error) {
	req := &RemoveAllFilesFromDownloads{
		OnlyActive:      onlyActive,
		OnlyCompleted:   onlyCompleted,
		DeleteFromCache: deleteFromCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveBusinessConnectedBotFromChat Removes the connected business bot from a specific chat by adding the chat to businessRecipients.excluded_chat_ids @chat_id Chat identifier
func (c *Client) RemoveBusinessConnectedBotFromChat(chatId int64) (*Ok, error) {
	req := &RemoveBusinessConnectedBotFromChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveChatActionBar Removes a chat action bar without any other action @chat_id Chat identifier
func (c *Client) RemoveChatActionBar(chatId int64) (*Ok, error) {
	req := &RemoveChatActionBar{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveContacts Removes users from the contact list @user_ids Identifiers of users to be deleted
func (c *Client) RemoveContacts(userIds []int64) (*Ok, error) {
	req := &RemoveContacts{
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers @sticker Sticker file to delete from the list
func (c *Client) RemoveFavoriteSticker(sticker *InputFile) (*Ok, error) {
	req := &RemoveFavoriteSticker{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveFileFromDownloads Removes a file from the file download list @file_id Identifier of the downloaded file @delete_from_cache Pass true to delete the file from the TDLib file cache
func (c *Client) RemoveFileFromDownloads(fileId int32, deleteFromCache bool) (*Ok, error) {
	req := &RemoveFileFromDownloads{
		FileId:          fileId,
		DeleteFromCache: deleteFromCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveGiftCollectionGifts Removes gifts from a collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) RemoveGiftCollectionGifts(ownerId *MessageSender, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	req := &RemoveGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollection), nil
}

// RemoveInstalledBackground Removes background from the list of installed backgrounds @background_id The background identifier
func (c *Client) RemoveInstalledBackground(backgroundId string) (*Ok, error) {
	req := &RemoveInstalledBackground{
		BackgroundId: backgroundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveLoginPasskey Removes a passkey from the list of passkeys allowed to be used for the login by the current user @passkey_id Unique identifier of the passkey to remove
func (c *Client) RemoveLoginPasskey(passkeyId string) (*Ok, error) {
	req := &RemoveLoginPasskey{
		PasskeyId: passkeyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveMessageReaction Removes a reaction from a message. A chosen reaction can always be removed
func (c *Client) RemoveMessageReaction(chatId int64, messageId int64, reactionType *ReactionType) (*Ok, error) {
	req := &RemoveMessageReaction{
		ChatId:       chatId,
		MessageId:    messageId,
		ReactionType: reactionType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveMessageSenderBotVerification Removes the verification status of a user or a chat by an owned bot
func (c *Client) RemoveMessageSenderBotVerification(botUserId int64, verifiedId *MessageSender) (*Ok, error) {
	req := &RemoveMessageSenderBotVerification{
		BotUserId:  botUserId,
		VerifiedId: verifiedId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveNotification Removes an active notification from notification list. Needs to be called only if the notification is removed by the current user @notification_group_id Identifier of notification group to which the notification belongs @notification_id Identifier of removed notification
func (c *Client) RemoveNotification(notificationGroupId int32, notificationId int32) (*Ok, error) {
	req := &RemoveNotification{
		NotificationGroupId: notificationGroupId,
		NotificationId:      notificationId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveNotificationGroup Removes a group of active notifications. Needs to be called only if the notification group is removed by the current user @notification_group_id Notification group identifier @max_notification_id The maximum identifier of removed notifications
func (c *Client) RemoveNotificationGroup(notificationGroupId int32, maxNotificationId int32) (*Ok, error) {
	req := &RemoveNotificationGroup{
		NotificationGroupId: notificationGroupId,
		MaxNotificationId:   maxNotificationId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemovePendingLiveStoryReactions Removes all pending paid reactions in a live story group call @group_call_id Group call identifier
func (c *Client) RemovePendingLiveStoryReactions(groupCallId int32) (*Ok, error) {
	req := &RemovePendingLiveStoryReactions{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemovePendingPaidMessageReactions Removes all pending paid reactions on a message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message
func (c *Client) RemovePendingPaidMessageReactions(chatId int64, messageId int64) (*Ok, error) {
	req := &RemovePendingPaidMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveProfileAudio Removes an audio file from the profile audio files of the current user @file_id Identifier of the audio file to be removed
func (c *Client) RemoveProfileAudio(fileId int32) (*Ok, error) {
	req := &RemoveProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveProxy Removes a proxy server. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) RemoveProxy(proxyId int32) (*Ok, error) {
	req := &RemoveProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags @hashtag Hashtag to delete
func (c *Client) RemoveRecentHashtag(hashtag string) (*Ok, error) {
	req := &RemoveRecentHashtag{
		Hashtag: hashtag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats @chat_id Identifier of the chat to be removed
func (c *Client) RemoveRecentlyFoundChat(chatId int64) (*Ok, error) {
	req := &RemoveRecentlyFoundChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers @is_attached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers @sticker Sticker file to delete
func (c *Client) RemoveRecentSticker(isAttached bool, sticker *InputFile) (*Ok, error) {
	req := &RemoveRecentSticker{
		IsAttached: isAttached,
		Sticker:    sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveSavedAnimation Removes an animation from the list of saved animations @animation Animation file to be removed
func (c *Client) RemoveSavedAnimation(animation *InputFile) (*Ok, error) {
	req := &RemoveSavedAnimation{
		Animation: animation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveSavedNotificationSound Removes a notification sound from the list of saved notification sounds @notification_sound_id Identifier of the notification sound
func (c *Client) RemoveSavedNotificationSound(notificationSoundId string) (*Ok, error) {
	req := &RemoveSavedNotificationSound{
		NotificationSoundId: notificationSoundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveSearchedForTag Removes a hashtag or a cashtag from the list of recently searched for hashtags or cashtags @tag Hashtag or cashtag to delete
func (c *Client) RemoveSearchedForTag(tag string) (*Ok, error) {
	req := &RemoveSearchedForTag{
		Tag: tag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs. The sticker set must be owned by the current user @sticker Sticker to remove from the set
func (c *Client) RemoveStickerFromSet(sticker *InputFile) (*Ok, error) {
	req := &RemoveStickerFromSet{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RemoveStoryAlbumStories Removes stories from an album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) RemoveStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	req := &RemoveStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbum), nil
}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled @category Category of frequently used chats @chat_id Chat identifier
func (c *Client) RemoveTopChat(category *TopChatCategory, chatId int64) (*Ok, error) {
	req := &RemoveTopChat{
		Category: category,
		ChatId:   chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderActiveUsernames Changes order of active usernames of the current user @usernames The new order of active usernames. All currently active usernames must be specified
func (c *Client) ReorderActiveUsernames(usernames []string) (*Ok, error) {
	req := &ReorderActiveUsernames{
		Usernames: usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderBotActiveUsernames Changes order of active usernames of a bot. Can be called only if userTypeBot.can_be_edited == true @bot_user_id Identifier of the target bot @usernames The new order of active usernames. All currently active usernames must be specified
func (c *Client) ReorderBotActiveUsernames(botUserId int64, usernames []string) (*Ok, error) {
	req := &ReorderBotActiveUsernames{
		BotUserId: botUserId,
		Usernames: usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderBotMediaPreviews Changes order of media previews in the list of media previews of a bot
func (c *Client) ReorderBotMediaPreviews(botUserId int64, languageCode string, fileIds []int32) (*Ok, error) {
	req := &ReorderBotMediaPreviews{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileIds:      fileIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderChatFolders Changes the order of chat folders @chat_folder_ids Identifiers of chat folders in the new correct order @main_chat_list_position Position of the main chat list among chat folders, 0-based. Can be non-zero only for Premium users
func (c *Client) ReorderChatFolders(chatFolderIds []int32, mainChatListPosition int32) (*Ok, error) {
	req := &ReorderChatFolders{
		ChatFolderIds:        chatFolderIds,
		MainChatListPosition: mainChatListPosition,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderGiftCollectionGifts Changes order of gifts in a collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) ReorderGiftCollectionGifts(ownerId *MessageSender, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	req := &ReorderGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollection), nil
}

// ReorderGiftCollections Changes order of gift collections. If the collections are owned by a channel chat, then requires can_post_messages administrator right in the channel chat
func (c *Client) ReorderGiftCollections(ownerId *MessageSender, collectionIds []int32) (*Ok, error) {
	req := &ReorderGiftCollections{
		OwnerId:       ownerId,
		CollectionIds: collectionIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderInstalledStickerSets Changes the order of installed sticker sets @sticker_type Type of the sticker sets to reorder @sticker_set_ids Identifiers of installed sticker sets in the new correct order
func (c *Client) ReorderInstalledStickerSets(stickerType *StickerType, stickerSetIds []string) (*Ok, error) {
	req := &ReorderInstalledStickerSets{
		StickerType:   stickerType,
		StickerSetIds: stickerSetIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderQuickReplyShortcuts Changes the order of quick reply shortcuts @shortcut_ids The new order of quick reply shortcuts
func (c *Client) ReorderQuickReplyShortcuts(shortcutIds []int32) (*Ok, error) {
	req := &ReorderQuickReplyShortcuts{
		ShortcutIds: shortcutIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderStoryAlbums Changes order of story albums. If the albums are owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat
func (c *Client) ReorderStoryAlbums(chatId int64, storyAlbumIds []int32) (*Ok, error) {
	req := &ReorderStoryAlbums{
		ChatId:        chatId,
		StoryAlbumIds: storyAlbumIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReorderStoryAlbumStories Changes order of stories in an album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) ReorderStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	req := &ReorderStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbum), nil
}

// ReorderSupergroupActiveUsernames Changes order of active usernames of a supergroup or channel, requires owner privileges in the supergroup or channel
func (c *Client) ReorderSupergroupActiveUsernames(supergroupId int64, usernames []string) (*Ok, error) {
	req := &ReorderSupergroupActiveUsernames{
		SupergroupId: supergroupId,
		Usernames:    usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReplaceLiveStoryRtmpUrl Replaces the current RTMP URL for streaming to a live story; requires owner privileges for channel chats @chat_id Chat identifier
func (c *Client) ReplaceLiveStoryRtmpUrl(chatId int64) (*RtmpUrl, error) {
	req := &ReplaceLiveStoryRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RtmpUrl), nil
}

// ReplacePrimaryChatInviteLink Replaces current primary invite link for a chat with a new primary invite link. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right @chat_id Chat identifier
func (c *Client) ReplacePrimaryChatInviteLink(chatId int64) (*ChatInviteLink, error) {
	req := &ReplacePrimaryChatInviteLink{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLink), nil
}

// ReplaceStickerInSet Replaces existing sticker in a set. The function is equivalent to removeStickerFromSet, then addStickerToSet, then setStickerPositionInSet
func (c *Client) ReplaceStickerInSet(userId int64, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	req := &ReplaceStickerInSet{
		UserId:     userId,
		Name:       name,
		OldSticker: oldSticker,
		NewSticker: newSticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReplaceVideoChatRtmpUrl Replaces the current RTMP URL for streaming to the video chat of a chat; requires owner privileges in the chat @chat_id Chat identifier
func (c *Client) ReplaceVideoChatRtmpUrl(chatId int64) (*RtmpUrl, error) {
	req := &ReplaceVideoChatRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RtmpUrl), nil
}

// ReportAuthenticationCodeMissing Reports that authentication code wasn't delivered via SMS; for official mobile applications only. Works only when the current authorization state is authorizationStateWaitCode @mobile_network_code Current mobile network code
func (c *Client) ReportAuthenticationCodeMissing(mobileNetworkCode string) (*Ok, error) {
	req := &ReportAuthenticationCodeMissing{
		MobileNetworkCode: mobileNetworkCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportChat Reports a chat to the Telegram moderators. A chat can be reported only from the chat action bar, or if chat.can_be_reported
func (c *Client) ReportChat(chatId int64, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	req := &ReportChat{
		ChatId:     chatId,
		OptionId:   optionId,
		MessageIds: messageIds,
		Text:       text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReportChatResult), nil
}

// ReportChatPhoto Reports a chat photo to the Telegram moderators. A chat photo can be reported only if chat.can_be_reported
func (c *Client) ReportChatPhoto(chatId int64, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	req := &ReportChatPhoto{
		ChatId: chatId,
		FileId: fileId,
		Reason: reason,
		Text:   text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportChatSponsoredMessage Reports a sponsored message to Telegram moderators
func (c *Client) ReportChatSponsoredMessage(chatId int64, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	req := &ReportChatSponsoredMessage{
		ChatId:    chatId,
		MessageId: messageId,
		OptionId:  optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReportSponsoredResult), nil
}

// ReportMessageReactions Reports reactions set on a message to the Telegram moderators. Reactions on a message can be reported only if messageProperties.can_report_reactions
func (c *Client) ReportMessageReactions(chatId int64, messageId int64, senderId *MessageSender) (*Ok, error) {
	req := &ReportMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
		SenderId:  senderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportPhoneNumberCodeMissing Reports that authentication code wasn't delivered via SMS to the specified phone number; for official mobile applications only @mobile_network_code Current mobile network code
func (c *Client) ReportPhoneNumberCodeMissing(mobileNetworkCode string) (*Ok, error) {
	req := &ReportPhoneNumberCodeMissing{
		MobileNetworkCode: mobileNetworkCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportSponsoredChat Reports a sponsored chat to Telegram moderators
func (c *Client) ReportSponsoredChat(sponsoredChatUniqueId int64, optionId string) (*ReportSponsoredResult, error) {
	req := &ReportSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
		OptionId:              optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReportSponsoredResult), nil
}

// ReportStory Reports a story to the Telegram moderators
func (c *Client) ReportStory(storyPosterChatId int64, storyId int32, optionId string, text string) (*ReportStoryResult, error) {
	req := &ReportStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		OptionId:          optionId,
		Text:              text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReportStoryResult), nil
}

// ReportSupergroupAntiSpamFalsePositive Reports a false deletion of a message by aggressive anti-spam checks; requires administrator rights in the supergroup. Can be called only for messages from chatEventMessageDeleted with can_report_anti_spam_false_positive == true
func (c *Client) ReportSupergroupAntiSpamFalsePositive(supergroupId int64, messageId int64) (*Ok, error) {
	req := &ReportSupergroupAntiSpamFalsePositive{
		SupergroupId: supergroupId,
		MessageId:    messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportSupergroupSpam Reports messages in a supergroup as spam; requires administrator rights in the supergroup
func (c *Client) ReportSupergroupSpam(supergroupId int64, messageIds []int64) (*Ok, error) {
	req := &ReportSupergroupSpam{
		SupergroupId: supergroupId,
		MessageIds:   messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ReportVideoMessageAdvertisement Reports a video message advertisement to Telegram moderators
func (c *Client) ReportVideoMessageAdvertisement(advertisementUniqueId int64, optionId string) (*ReportSponsoredResult, error) {
	req := &ReportVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
		OptionId:              optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ReportSponsoredResult), nil
}

// RequestAuthenticationPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (c *Client) RequestAuthenticationPasswordRecovery() (*Ok, error) {
	req := &RequestAuthenticationPasswordRecovery{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RequestPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up
func (c *Client) RequestPasswordRecovery() (*EmailAddressAuthenticationCodeInfo, error) {
	req := &RequestPasswordRecovery{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmailAddressAuthenticationCodeInfo), nil
}

// RequestQrCodeAuthentication Requests QR code authentication by scanning a QR code on another logged in device. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
func (c *Client) RequestQrCodeAuthentication(otherUserIds []int64) (*Ok, error) {
	req := &RequestQrCodeAuthentication{
		OtherUserIds: otherUserIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResendAuthenticationCode Resends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode, the next_code_type of the result is not null
func (c *Client) ResendAuthenticationCode(opts *ResendAuthenticationCodeOpts) (*Ok, error) {
	req := &ResendAuthenticationCode{}
	if opts != nil {
		req.Reason = opts.Reason
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResendEmailAddressVerificationCode Resends the code to verify an email address to be added to a user's Telegram Passport
func (c *Client) ResendEmailAddressVerificationCode() (*EmailAddressAuthenticationCodeInfo, error) {
	req := &ResendEmailAddressVerificationCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmailAddressAuthenticationCodeInfo), nil
}

// ResendLoginEmailAddressCode Resends the login email address verification code
func (c *Client) ResendLoginEmailAddressCode() (*EmailAddressAuthenticationCodeInfo, error) {
	req := &ResendLoginEmailAddressCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmailAddressAuthenticationCodeInfo), nil
}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
func (c *Client) ResendMessages(chatId int64, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	req := &ResendMessages{
		ChatId:               chatId,
		MessageIds:           messageIds,
		PaidMessageStarCount: paidMessageStarCount,
	}
	if opts != nil {
		req.Quote = opts.Quote
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// ResendPhoneNumberCode Resends the authentication code sent to a phone number. Works only if the previously received authenticationCodeInfo next_code_type was not null and the server-specified timeout has passed
func (c *Client) ResendPhoneNumberCode(opts *ResendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	req := &ResendPhoneNumberCode{}
	if opts != nil {
		req.Reason = opts.Reason
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AuthenticationCodeInfo), nil
}

// ResendRecoveryEmailAddressCode Resends the 2-step verification recovery email address verification code
func (c *Client) ResendRecoveryEmailAddressCode() (*PasswordState, error) {
	req := &ResendRecoveryEmailAddressCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// ResetAllNotificationSettings Resets all chat and scope notification settings to their default values. By default, all chats are unmuted and message previews are shown
func (c *Client) ResetAllNotificationSettings() (*Ok, error) {
	req := &ResetAllNotificationSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResetAuthenticationEmailAddress Resets the login email address. May return an error with a message "TASK_ALREADY_EXISTS" if reset is still pending.
func (c *Client) ResetAuthenticationEmailAddress() (*Ok, error) {
	req := &ResetAuthenticationEmailAddress{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResetInstalledBackgrounds Resets list of installed backgrounds to its default value
func (c *Client) ResetInstalledBackgrounds() (*Ok, error) {
	req := &ResetInstalledBackgrounds{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (c *Client) ResetNetworkStatistics() (*Ok, error) {
	req := &ResetNetworkStatistics{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ResetPassword Removes 2-step verification password without previous password and access to recovery email address. The password can't be reset immediately and the request needs to be repeated after the specified time
func (c *Client) ResetPassword() (*ResetPasswordResult, error) {
	req := &ResetPassword{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ResetPasswordResult), nil
}

// ReuseStarSubscription Reuses an active Telegram Star subscription to a channel chat and joins the chat again @subscription_id Identifier of the subscription
func (c *Client) ReuseStarSubscription(subscriptionId string) (*Ok, error) {
	req := &ReuseStarSubscription{
		SubscriptionId: subscriptionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// RevokeChatInviteLink Revokes invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links.
func (c *Client) RevokeChatInviteLink(chatId int64, inviteLink string) (*ChatInviteLinks, error) {
	req := &RevokeChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatInviteLinks), nil
}

// RevokeGroupCallInviteLink Revokes invite link for a group call. Requires groupCall.can_be_managed right for video chats or groupCall.is_owned otherwise @group_call_id Group call identifier
func (c *Client) RevokeGroupCallInviteLink(groupCallId int32) (*Ok, error) {
	req := &RevokeGroupCallInviteLink{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SaveApplicationLogEvent Saves application log event on the server. Can be called before authorization @type Event type @chat_id Optional chat identifier, associated with the event @data The log event data
func (c *Client) SaveApplicationLogEvent(typeField string, chatId int64, data *JsonValue) (*Ok, error) {
	req := &SaveApplicationLogEvent{
		TypeField: typeField,
		ChatId:    chatId,
		Data:      data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SavePreparedInlineMessage Saves an inline message to be sent by the given user; for bots only
func (c *Client) SavePreparedInlineMessage(userId int64, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	req := &SavePreparedInlineMessage{
		UserId:    userId,
		Result:    result,
		ChatTypes: chatTypes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PreparedInlineMessageId), nil
}

// SearchAffiliatePrograms Searches affiliate programs that can be connected to the given affiliate
func (c *Client) SearchAffiliatePrograms(affiliate *AffiliateType, sortOrder *AffiliateProgramSortOrder, offset string, limit int32) (*FoundAffiliatePrograms, error) {
	req := &SearchAffiliatePrograms{
		Affiliate: affiliate,
		SortOrder: sortOrder,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundAffiliatePrograms), nil
}

// SearchBackground Searches for a background by its name @name The name of the background
func (c *Client) SearchBackground(name string) (*Background, error) {
	req := &SearchBackground{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Background), nil
}

// SearchCallMessages Searches for call and group call messages. Returns the results in reverse chronological order (i.e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib
func (c *Client) SearchCallMessages(offset string, limit int32, onlyMissed bool) (*FoundMessages, error) {
	req := &SearchCallMessages{
		Offset:     offset,
		Limit:      limit,
		OnlyMissed: onlyMissed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundMessages), nil
}

// SearchChatAffiliateProgram Searches a chat with an affiliate program. Returns the chat if found and the program is active
func (c *Client) SearchChatAffiliateProgram(username string, referrer string) (*Chat, error) {
	req := &SearchChatAffiliateProgram{
		Username: username,
		Referrer: referrer,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// SearchChatMembers Searches for a specified query in the first name, last name and usernames of the members of a specified chat. Requires administrator rights if the chat is a channel
func (c *Client) SearchChatMembers(chatId int64, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	req := &SearchChatMembers{
		ChatId: chatId,
		Query:  query,
		Limit:  limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ChatMembers), nil
}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query
func (c *Client) SearchChatMessages(chatId int64, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	req := &SearchChatMessages{
		ChatId:        chatId,
		Query:         query,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.SenderId = opts.SenderId
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundChatMessages), nil
}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user @chat_id Chat identifier @limit The maximum number of messages to be returned
func (c *Client) SearchChatRecentLocationMessages(chatId int64, limit int32) (*Messages, error) {
	req := &SearchChatRecentLocationMessages{
		ChatId: chatId,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// SearchChats Searches for the specified query in the title and username of already known chats. This is an offline method. Returns chats in the order seen in the main chat list
func (c *Client) SearchChats(query string, limit int32) (*Chats, error) {
	req := &SearchChats{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list @query Query to search for @limit The maximum number of chats to be returned
func (c *Client) SearchChatsOnServer(query string, limit int32) (*Chats, error) {
	req := &SearchChatsOnServer{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
func (c *Client) SearchContacts(query string, limit int32) (*Users, error) {
	req := &SearchContacts{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Users), nil
}

// SearchEmojis Searches for emojis by keywords. Supported only if the file database is enabled. Order of results is unspecified
func (c *Client) SearchEmojis(text string, inputLanguageCodes []string) (*EmojiKeywords, error) {
	req := &SearchEmojis{
		Text:               text,
		InputLanguageCodes: inputLanguageCodes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmojiKeywords), nil
}

// SearchFileDownloads Searches for files in the file download list or recently downloaded files from the list
func (c *Client) SearchFileDownloads(query string, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	req := &SearchFileDownloads{
		Query:         query,
		OnlyActive:    onlyActive,
		OnlyCompleted: onlyCompleted,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundFileDownloads), nil
}

// SearchGiftsForResale Returns upgraded gifts that can be bought from other owners using sendResoldGift
func (c *Client) SearchGiftsForResale(giftId string, order *GiftForResaleOrder, attributes []*UpgradedGiftAttributeId, offset string, limit int32) (*GiftsForResale, error) {
	req := &SearchGiftsForResale{
		GiftId:     giftId,
		Order:      order,
		Attributes: attributes,
		Offset:     offset,
		Limit:      limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftsForResale), nil
}

// SearchHashtags Searches for recently used hashtags by their prefix @prefix Hashtag prefix to search for @limit The maximum number of hashtags to be returned
func (c *Client) SearchHashtags(prefix string, limit int32) (*Hashtags, error) {
	req := &SearchHashtags{
		Prefix: prefix,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Hashtags), nil
}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name @sticker_type Type of the sticker sets to search for @query Query to search for @limit The maximum number of sticker sets to return
func (c *Client) SearchInstalledStickerSets(stickerType *StickerType, query string, limit int32) (*StickerSets, error) {
	req := &SearchInstalledStickerSets{
		StickerType: stickerType,
		Query:       query,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)).
func (c *Client) SearchMessages(query string, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	req := &SearchMessages{
		Query:   query,
		Offset:  offset,
		Limit:   limit,
		MinDate: minDate,
		MaxDate: maxDate,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
		req.Filter = opts.Filter
		req.ChatTypeFilter = opts.ChatTypeFilter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundMessages), nil
}

// SearchOutgoingDocumentMessages Searches for outgoing messages with content of the type messageDocument in all chats except secret chats. Returns the results in reverse chronological order
func (c *Client) SearchOutgoingDocumentMessages(query string, limit int32) (*FoundMessages, error) {
	req := &SearchOutgoingDocumentMessages{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundMessages), nil
}

// SearchPublicChat Searches a public chat by its username. Currently, only private chats, supergroups and channels can be public. Returns the chat if found; otherwise, an error is returned @username Username to be resolved
func (c *Client) SearchPublicChat(username string) (*Chat, error) {
	req := &SearchPublicChat{
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently, only private chats, supergroups and channels can be public. Returns a meaningful number of results.
func (c *Client) SearchPublicChats(query string) (*Chats, error) {
	req := &SearchPublicChats{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// SearchPublicMessagesByTag Searches for public channel posts containing the given hashtag or cashtag. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicMessagesByTag(tag string, offset string, limit int32) (*FoundMessages, error) {
	req := &SearchPublicMessagesByTag{
		Tag:    tag,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundMessages), nil
}

// SearchPublicPosts Searches for public channel posts using the given query. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicPosts(query string, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	req := &SearchPublicPosts{
		Query:     query,
		Offset:    offset,
		Limit:     limit,
		StarCount: starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundPublicPosts), nil
}

// SearchPublicStoriesByLocation Searches for public stories by the given address location. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByLocation(address *LocationAddress, offset string, limit int32) (*FoundStories, error) {
	req := &SearchPublicStoriesByLocation{
		Address: address,
		Offset:  offset,
		Limit:   limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundStories), nil
}

// SearchPublicStoriesByTag Searches for public stories containing the given hashtag or cashtag. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByTag(storyPosterChatId int64, tag string, offset string, limit int32) (*FoundStories, error) {
	req := &SearchPublicStoriesByTag{
		StoryPosterChatId: storyPosterChatId,
		Tag:               tag,
		Offset:            offset,
		Limit:             limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundStories), nil
}

// SearchPublicStoriesByVenue Searches for public stories from the given venue. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByVenue(venueProvider string, venueId string, offset string, limit int32) (*FoundStories, error) {
	req := &SearchPublicStoriesByVenue{
		VenueProvider: venueProvider,
		VenueId:       venueId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundStories), nil
}

// SearchQuote Searches for a given quote in a text. Returns found quote start position in UTF-16 code units. Returns a 404 error if the quote is not found. Can be called synchronously
func (c *Client) SearchQuote(text *FormattedText, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	req := &SearchQuote{
		Text:          text,
		Quote:         quote,
		QuotePosition: quotePosition,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundPosition), nil
}

// SearchRecentlyFoundChats Searches for the specified query in the title and username of up to 50 recently found chats. This is an offline method
func (c *Client) SearchRecentlyFoundChats(query string, limit int32) (*Chats, error) {
	req := &SearchRecentlyFoundChats{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chats), nil
}

// SearchSavedMessages Searches for messages tagged by the given reaction and with the given words in the Saved Messages chat; for Telegram Premium users only.
func (c *Client) SearchSavedMessages(savedMessagesTopicId int64, query string, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	req := &SearchSavedMessages{
		SavedMessagesTopicId: savedMessagesTopicId,
		Query:                query,
		FromMessageId:        fromMessageId,
		Offset:               offset,
		Limit:                limit,
	}
	if opts != nil {
		req.Tag = opts.Tag
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundChatMessages), nil
}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance, the number of returned messages is chosen by TDLib
func (c *Client) SearchSecretMessages(chatId int64, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	req := &SearchSecretMessages{
		ChatId: chatId,
		Query:  query,
		Offset: offset,
		Limit:  limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundMessages), nil
}

// SearchStickers Searches for stickers from public sticker sets that correspond to any of the given emoji
func (c *Client) SearchStickers(stickerType *StickerType, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	req := &SearchStickers{
		StickerType:        stickerType,
		Emojis:             emojis,
		Query:              query,
		InputLanguageCodes: inputLanguageCodes,
		Offset:             offset,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Stickers), nil
}

// SearchStickerSet Searches for a sticker set by its name @name Name of the sticker set @ignore_cache Pass true to ignore local cache of sticker sets and always send a network request
func (c *Client) SearchStickerSet(name string, ignoreCache bool) (*StickerSet, error) {
	req := &SearchStickerSet{
		Name:        name,
		IgnoreCache: ignoreCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSet), nil
}

// SearchStickerSets Searches for sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
func (c *Client) SearchStickerSets(stickerType *StickerType, query string) (*StickerSets, error) {
	req := &SearchStickerSets{
		StickerType: stickerType,
		Query:       query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StickerSets), nil
}

// SearchStringsByPrefix Searches specified query by word prefixes in the provided strings. Returns 0-based positions of strings that matched. Can be called synchronously
func (c *Client) SearchStringsByPrefix(strings []string, query string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	req := &SearchStringsByPrefix{
		Strings:                 strings,
		Query:                   query,
		Limit:                   limit,
		ReturnNoneForEmptyQuery: returnNoneForEmptyQuery,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundPositions), nil
}

// SearchUserByPhoneNumber Searches a user by their phone number. Returns a 404 error if the user can't be found
func (c *Client) SearchUserByPhoneNumber(phoneNumber string, onlyLocal bool) (*User, error) {
	req := &SearchUserByPhoneNumber{
		PhoneNumber: phoneNumber,
		OnlyLocal:   onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// SearchUserByToken Searches a user by a token from the user's link @token Token to search for
func (c *Client) SearchUserByToken(token string) (*User, error) {
	req := &SearchUserByToken{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// SearchWebApp Returns information about a Web App by its short name. Returns a 404 error if the Web App is not found
func (c *Client) SearchWebApp(botUserId int64, webAppShortName string) (*FoundWebApp, error) {
	req := &SearchWebApp{
		BotUserId:       botUserId,
		WebAppShortName: webAppShortName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FoundWebApp), nil
}

// SellGift Sells a gift for Telegram Stars; requires owner privileges for gifts owned by a chat
func (c *Client) SellGift(businessConnectionId string, receivedGiftId string) (*Ok, error) {
	req := &SellGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendAuthenticationFirebaseSms Sends Firebase Authentication SMS to the phone number of the user. Works only when the current authorization state is authorizationStateWaitCode and the server returned code of the type authenticationCodeTypeFirebaseAndroid or authenticationCodeTypeFirebaseIos
func (c *Client) SendAuthenticationFirebaseSms(token string) (*Ok, error) {
	req := &SendAuthenticationFirebaseSms{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command; requires can_invite_users member right. Bots can't be invited to a private chat other than the chat with the bot.
func (c *Client) SendBotStartMessage(botUserId int64, chatId int64, parameter string) (*Message, error) {
	req := &SendBotStartMessage{
		BotUserId: botUserId,
		ChatId:    chatId,
		Parameter: parameter,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// SendBusinessMessage Sends a message on behalf of a business account; for bots only. Returns the message after it was sent
func (c *Client) SendBusinessMessage(businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	req := &SendBusinessMessage{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		DisableNotification:  disableNotification,
		ProtectContent:       protectContent,
		EffectId:             effectId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// SendBusinessMessageAlbum Sends 2-10 messages grouped together into an album on behalf of a business account; for bots only. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) SendBusinessMessageAlbum(businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	req := &SendBusinessMessageAlbum{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		DisableNotification:  disableNotification,
		ProtectContent:       protectContent,
		EffectId:             effectId,
		InputMessageContents: inputMessageContents,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessages), nil
}

// SendCallDebugInformation Sends debug information for a call to Telegram servers @call_id Call identifier @debug_information Debug information in application-specific format
func (c *Client) SendCallDebugInformation(callId int32, debugInformation string) (*Ok, error) {
	req := &SendCallDebugInformation{
		CallId:           callId,
		DebugInformation: debugInformation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendCallLog Sends log file for a call to Telegram servers @call_id Call identifier @log_file Call log file. Only inputFileLocal and inputFileGenerated are supported
func (c *Client) SendCallLog(callId int32, logFile *InputFile) (*Ok, error) {
	req := &SendCallLog{
		CallId:  callId,
		LogFile: logFile,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendCallRating Sends a call rating
func (c *Client) SendCallRating(callId int32, rating int32, comment string, problems []*CallProblem) (*Ok, error) {
	req := &SendCallRating{
		CallId:   callId,
		Rating:   rating,
		Comment:  comment,
		Problems: problems,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendCallSignalingData Sends call signaling data @call_id Call identifier @data The data
func (c *Client) SendCallSignalingData(callId int32, data string) (*Ok, error) {
	req := &SendCallSignalingData{
		CallId: callId,
		Data:   data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendChatAction Sends a notification about user activity in a chat
func (c *Client) SendChatAction(chatId int64, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	req := &SendChatAction{
		ChatId:               chatId,
		TopicId:              topicId,
		BusinessConnectionId: businessConnectionId,
	}
	if opts != nil {
		req.Action = opts.Action
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendCustomRequest Sends a custom request; for bots only @method The method name @parameters JSON-serialized method parameters
func (c *Client) SendCustomRequest(method string, parameters string) (*CustomRequestResult, error) {
	req := &SendCustomRequest{
		Method:     method,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CustomRequestResult), nil
}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport @email_address Email address
func (c *Client) SendEmailAddressVerificationCode(emailAddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	req := &SendEmailAddressVerificationCode{
		EmailAddress: emailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmailAddressAuthenticationCodeInfo), nil
}

// SendGift Sends a gift to another user or channel chat. May return an error with a message "STARGIFT_USAGE_LIMITED" if the gift was sold out
func (c *Client) SendGift(giftId string, ownerId *MessageSender, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	req := &SendGift{
		GiftId:        giftId,
		OwnerId:       ownerId,
		Text:          text,
		IsPrivate:     isPrivate,
		PayForUpgrade: payForUpgrade,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendGiftPurchaseOffer Sends an offer to purchase an upgraded gift
func (c *Client) SendGiftPurchaseOffer(ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	req := &SendGiftPurchaseOffer{
		OwnerId:              ownerId,
		GiftName:             giftName,
		Price:                price,
		Duration:             duration,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendGroupCallMessage Sends a message to other participants of a group call. Requires groupCall.can_send_messages right
func (c *Client) SendGroupCallMessage(groupCallId int32, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	req := &SendGroupCallMessage{
		GroupCallId:          groupCallId,
		Text:                 text,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
func (c *Client) SendInlineQueryResultMessage(chatId int64, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	req := &SendInlineQueryResultMessage{
		ChatId:     chatId,
		QueryId:    queryId,
		ResultId:   resultId,
		HideViaBot: hideViaBot,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// SendMessage Sends a message. Returns the sent message
func (c *Client) SendMessage(chatId int64, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	req := &SendMessage{
		ChatId:              chatId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// SendMessageAlbum Sends 2-10 messages grouped together into an album. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) SendMessageAlbum(chatId int64, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	req := &SendMessageAlbum{
		ChatId:               chatId,
		InputMessageContents: inputMessageContents,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service. This method must be called after getPassportAuthorizationFormAvailableElements if some previously available elements are going to be reused
func (c *Client) SendPassportAuthorizationForm(authorizationFormId int32, types []*PassportElementType) (*Ok, error) {
	req := &SendPassportAuthorizationForm{
		AuthorizationFormId: authorizationFormId,
		Types:               types,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
func (c *Client) SendPaymentForm(inputInvoice *InputInvoice, paymentFormId string, orderInfoId string, shippingOptionId string, tipAmount int64, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	req := &SendPaymentForm{
		InputInvoice:     inputInvoice,
		PaymentFormId:    paymentFormId,
		OrderInfoId:      orderInfoId,
		ShippingOptionId: shippingOptionId,
		TipAmount:        tipAmount,
	}
	if opts != nil {
		req.Credentials = opts.Credentials
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PaymentResult), nil
}

// SendPhoneNumberCode Sends a code to the specified phone number. Aborts previous phone number verification if there was one. On success, returns information about the sent code
func (c *Client) SendPhoneNumberCode(phoneNumber string, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	req := &SendPhoneNumberCode{
		PhoneNumber: phoneNumber,
		TypeField:   typeField,
	}
	if opts != nil {
		req.Settings = opts.Settings
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*AuthenticationCodeInfo), nil
}

// SendPhoneNumberFirebaseSms Sends Firebase Authentication SMS to the specified phone number. Works only when received a code of the type authenticationCodeTypeFirebaseAndroid or authenticationCodeTypeFirebaseIos
func (c *Client) SendPhoneNumberFirebaseSms(token string) (*Ok, error) {
	req := &SendPhoneNumberFirebaseSms{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendQuickReplyShortcutMessages Sends messages from a quick reply shortcut. Requires Telegram Business subscription. Can't be used to send paid messages
func (c *Client) SendQuickReplyShortcutMessages(chatId int64, shortcutId int32, sendingId int32) (*Messages, error) {
	req := &SendQuickReplyShortcutMessages{
		ChatId:     chatId,
		ShortcutId: shortcutId,
		SendingId:  sendingId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Messages), nil
}

// SendResoldGift Sends an upgraded gift that is available for resale to another user or channel chat; gifts already owned by the current user
func (c *Client) SendResoldGift(giftName string, ownerId *MessageSender, price *GiftResalePrice) (*GiftResaleResult, error) {
	req := &SendResoldGift{
		GiftName: giftName,
		OwnerId:  ownerId,
		Price:    price,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftResaleResult), nil
}

// SendTextMessageDraft Sends a draft for a being generated text message; for bots only
func (c *Client) SendTextMessageDraft(chatId int64, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	req := &SendTextMessageDraft{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		DraftId:      draftId,
		Text:         text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SendWebAppCustomRequest Sends a custom request from a Web App
func (c *Client) SendWebAppCustomRequest(botUserId int64, method string, parameters string) (*CustomRequestResult, error) {
	req := &SendWebAppCustomRequest{
		BotUserId:  botUserId,
		Method:     method,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*CustomRequestResult), nil
}

// SendWebAppData Sends data received from a keyboardButtonTypeWebApp Web App to a bot
func (c *Client) SendWebAppData(botUserId int64, buttonText string, data string) (*Ok, error) {
	req := &SendWebAppData{
		BotUserId:  botUserId,
		ButtonText: buttonText,
		Data:       data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAccentColor Changes accent color and background custom emoji for the current user; for Telegram Premium users only
func (c *Client) SetAccentColor(accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	req := &SetAccentColor{
		AccentColorId:           accentColorId,
		BackgroundCustomEmojiId: backgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAccountTtl Changes the period of inactivity after which the account of the current user will automatically be deleted @ttl New account TTL
func (c *Client) SetAccountTtl(ttl *AccountTtl) (*Ok, error) {
	req := &SetAccountTtl{
		Ttl: ttl,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before initialization @seconds Number of seconds before the function returns
func (c *Client) SetAlarm(seconds float64) (*Ok, error) {
	req := &SetAlarm{
		Seconds: seconds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetApplicationVerificationToken Informs TDLib that application or reCAPTCHA verification has been completed. Can be called before authorization
func (c *Client) SetApplicationVerificationToken(verificationId int64, token string) (*Ok, error) {
	req := &SetApplicationVerificationToken{
		VerificationId: verificationId,
		Token:          token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetArchiveChatListSettings Changes settings for automatic moving of chats to and from the Archive chat lists @settings New settings
func (c *Client) SetArchiveChatListSettings(settings *ArchiveChatListSettings) (*Ok, error) {
	req := &SetArchiveChatListSettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAuthenticationEmailAddress Sets the email address of the user and sends an authentication code to the email address. Works only when the current authorization state is authorizationStateWaitEmailAddress @email_address The email address of the user
func (c *Client) SetAuthenticationEmailAddress(emailAddress string) (*Ok, error) {
	req := &SetAuthenticationEmailAddress{
		EmailAddress: emailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
func (c *Client) SetAuthenticationPhoneNumber(phoneNumber string, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	req := &SetAuthenticationPhoneNumber{
		PhoneNumber: phoneNumber,
	}
	if opts != nil {
		req.Settings = opts.Settings
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAuthenticationPremiumPurchaseTransaction Informs server about an in-store purchase of Telegram Premium before authorization. Works only when the current authorization state is authorizationStateWaitPremiumPurchase
func (c *Client) SetAuthenticationPremiumPurchaseTransaction(transaction *StoreTransaction, isRestore bool, currency string, amount int64) (*Ok, error) {
	req := &SetAuthenticationPremiumPurchaseTransaction{
		Transaction: transaction,
		IsRestore:   isRestore,
		Currency:    currency,
		Amount:      amount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAutoDownloadSettings Sets auto-download settings @settings New user auto-download settings @type Type of the network for which the new settings are relevant
func (c *Client) SetAutoDownloadSettings(settings *AutoDownloadSettings, typeField *NetworkType) (*Ok, error) {
	req := &SetAutoDownloadSettings{
		Settings:  settings,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetAutosaveSettings Sets autosave settings for the given scope. The method is guaranteed to work only after at least one call to getAutosaveSettings @scope Autosave settings scope @settings New autosave settings for the scope; pass null to set autosave settings to default
func (c *Client) SetAutosaveSettings(scope *AutosaveSettingsScope, settings *ScopeAutosaveSettings) (*Ok, error) {
	req := &SetAutosaveSettings{
		Scope:    scope,
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBio Changes the bio of the current user @bio The new value of the user bio; 0-getOption("bio_length_max") characters without line feeds
func (c *Client) SetBio(bio string) (*Ok, error) {
	req := &SetBio{
		Bio: bio,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBirthdate Changes the birthdate of the current user @birthdate The new value of the current user's birthdate; pass null to remove the birthdate
func (c *Client) SetBirthdate(birthdate *Birthdate) (*Ok, error) {
	req := &SetBirthdate{
		Birthdate: birthdate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBotInfoDescription Sets the text shown in the chat with a bot if the chat is empty. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotInfoDescription(botUserId int64, languageCode string, description string) (*Ok, error) {
	req := &SetBotInfoDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Description:  description,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBotInfoShortDescription Sets the text shown on a bot's profile page and sent together with the link when users share the bot. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotInfoShortDescription(botUserId int64, languageCode string, shortDescription string) (*Ok, error) {
	req := &SetBotInfoShortDescription{
		BotUserId:        botUserId,
		LanguageCode:     languageCode,
		ShortDescription: shortDescription,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBotName Sets the name of a bot. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotName(botUserId int64, languageCode string, name string) (*Ok, error) {
	req := &SetBotName{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBotProfilePhoto Changes a profile photo for a bot @bot_user_id Identifier of the target bot @photo Profile photo to set; pass null to delete the chat photo
func (c *Client) SetBotProfilePhoto(botUserId int64, photo *InputChatPhoto) (*Ok, error) {
	req := &SetBotProfilePhoto{
		BotUserId: botUserId,
		Photo:     photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only @pending_update_count The number of pending updates @error_message The last error message
func (c *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*Ok, error) {
	req := &SetBotUpdatesStatus{
		PendingUpdateCount: pendingUpdateCount,
		ErrorMessage:       errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAccountBio Changes the bio of a business account; for bots only
func (c *Client) SetBusinessAccountBio(businessConnectionId string, bio string) (*Ok, error) {
	req := &SetBusinessAccountBio{
		BusinessConnectionId: businessConnectionId,
		Bio:                  bio,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAccountGiftSettings Changes settings for gift receiving of a business account; for bots only
func (c *Client) SetBusinessAccountGiftSettings(businessConnectionId string, settings *GiftSettings) (*Ok, error) {
	req := &SetBusinessAccountGiftSettings{
		BusinessConnectionId: businessConnectionId,
		Settings:             settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAccountName Changes the first and last name of a business account; for bots only
func (c *Client) SetBusinessAccountName(businessConnectionId string, firstName string, lastName string) (*Ok, error) {
	req := &SetBusinessAccountName{
		BusinessConnectionId: businessConnectionId,
		FirstName:            firstName,
		LastName:             lastName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAccountProfilePhoto Changes a profile photo of a business account; for bots only
func (c *Client) SetBusinessAccountProfilePhoto(businessConnectionId string, isPublic bool, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	req := &SetBusinessAccountProfilePhoto{
		BusinessConnectionId: businessConnectionId,
		IsPublic:             isPublic,
	}
	if opts != nil {
		req.Photo = opts.Photo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAccountUsername Changes the editable username of a business account; for bots only
func (c *Client) SetBusinessAccountUsername(businessConnectionId string, username string) (*Ok, error) {
	req := &SetBusinessAccountUsername{
		BusinessConnectionId: businessConnectionId,
		Username:             username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessAwayMessageSettings Changes the business away message settings of the current user. Requires Telegram Business subscription @away_message_settings The new settings for the away message of the business; pass null to disable the away message
func (c *Client) SetBusinessAwayMessageSettings(awayMessageSettings *BusinessAwayMessageSettings) (*Ok, error) {
	req := &SetBusinessAwayMessageSettings{
		AwayMessageSettings: awayMessageSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessConnectedBot Adds or changes business bot that is connected to the current user account @bot Connection settings for the bot
func (c *Client) SetBusinessConnectedBot(bot *BusinessConnectedBot) (*Ok, error) {
	req := &SetBusinessConnectedBot{
		Bot: bot,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessGreetingMessageSettings Changes the business greeting message settings of the current user. Requires Telegram Business subscription @greeting_message_settings The new settings for the greeting message of the business; pass null to disable the greeting message
func (c *Client) SetBusinessGreetingMessageSettings(greetingMessageSettings *BusinessGreetingMessageSettings) (*Ok, error) {
	req := &SetBusinessGreetingMessageSettings{
		GreetingMessageSettings: greetingMessageSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessLocation Changes the business location of the current user. Requires Telegram Business subscription @location The new location of the business; pass null to remove the location
func (c *Client) SetBusinessLocation(location *BusinessLocation) (*Ok, error) {
	req := &SetBusinessLocation{
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessMessageIsPinned Pins or unpins a message sent on behalf of a business account; for bots only
func (c *Client) SetBusinessMessageIsPinned(businessConnectionId string, chatId int64, messageId int64, isPinned bool) (*Ok, error) {
	req := &SetBusinessMessageIsPinned{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		IsPinned:             isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessOpeningHours Changes the business opening hours of the current user. Requires Telegram Business subscription
func (c *Client) SetBusinessOpeningHours(opts *SetBusinessOpeningHoursOpts) (*Ok, error) {
	req := &SetBusinessOpeningHours{}
	if opts != nil {
		req.OpeningHours = opts.OpeningHours
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetBusinessStartPage Changes the business start page of the current user. Requires Telegram Business subscription @start_page The new start page of the business; pass null to remove custom start page
func (c *Client) SetBusinessStartPage(startPage *InputBusinessStartPage) (*Ok, error) {
	req := &SetBusinessStartPage{
		StartPage: startPage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatAccentColor Changes accent color and background custom emoji of a channel chat. Requires can_change_info administrator right
func (c *Client) SetChatAccentColor(chatId int64, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	req := &SetChatAccentColor{
		ChatId:                  chatId,
		AccentColorId:           accentColorId,
		BackgroundCustomEmojiId: backgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatActiveStoriesList Changes story list in which stories from the chat are shown @chat_id Identifier of the chat that posted stories @story_list New list for active stories posted by the chat
func (c *Client) SetChatActiveStoriesList(chatId int64, storyList *StoryList) (*Ok, error) {
	req := &SetChatActiveStoriesList{
		ChatId:    chatId,
		StoryList: storyList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatAffiliateProgram Changes affiliate program for a bot
func (c *Client) SetChatAffiliateProgram(chatId int64, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	req := &SetChatAffiliateProgram{
		ChatId: chatId,
	}
	if opts != nil {
		req.Parameters = opts.Parameters
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatAvailableReactions Changes reactions, available in a chat. Available for basic groups, supergroups, and channels. Requires can_change_info member right
func (c *Client) SetChatAvailableReactions(chatId int64, availableReactions *ChatAvailableReactions) (*Ok, error) {
	req := &SetChatAvailableReactions{
		ChatId:             chatId,
		AvailableReactions: availableReactions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatBackground Sets the background in a specific chat. Supported only in private and secret chats with non-deleted users, and in chats with sufficient boost level and can_change_info administrator right
func (c *Client) SetChatBackground(chatId int64, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	req := &SetChatBackground{
		ChatId:           chatId,
		DarkThemeDimming: darkThemeDimming,
		OnlyForSelf:      onlyForSelf,
	}
	if opts != nil {
		req.Background = opts.Background
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatClientData Changes application-specific data associated with a chat @chat_id Chat identifier @client_data New value of client_data
func (c *Client) SetChatClientData(chatId int64, clientData string) (*Ok, error) {
	req := &SetChatClientData{
		ChatId:     chatId,
		ClientData: clientData,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatDescription Changes information about a chat. Available for basic groups, supergroups, and channels. Requires can_change_info member right @chat_id Identifier of the chat @param_description New chat description; 0-255 characters
func (c *Client) SetChatDescription(chatId int64, description string) (*Ok, error) {
	req := &SetChatDescription{
		ChatId:      chatId,
		Description: description,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatDirectMessagesGroup Changes direct messages group settings for a channel chat; requires owner privileges in the chat
func (c *Client) SetChatDirectMessagesGroup(chatId int64, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	req := &SetChatDirectMessagesGroup{
		ChatId:               chatId,
		IsEnabled:            isEnabled,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatDiscussionGroup Changes the discussion group of a channel chat; requires can_change_info administrator right in the channel if it is specified
func (c *Client) SetChatDiscussionGroup(chatId int64, discussionChatId int64) (*Ok, error) {
	req := &SetChatDiscussionGroup{
		ChatId:           chatId,
		DiscussionChatId: discussionChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatDraftMessage Changes the draft message in a chat or a topic
func (c *Client) SetChatDraftMessage(chatId int64, opts *SetChatDraftMessageOpts) (*Ok, error) {
	req := &SetChatDraftMessage{
		ChatId: chatId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.DraftMessage = opts.DraftMessage
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatEmojiStatus Changes the emoji status of a chat. Use chatBoostLevelFeatures.can_set_emoji_status to check whether an emoji status can be set. Requires can_change_info administrator right
func (c *Client) SetChatEmojiStatus(chatId int64, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	req := &SetChatEmojiStatus{
		ChatId: chatId,
	}
	if opts != nil {
		req.EmojiStatus = opts.EmojiStatus
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatLocation Changes the location of a chat. Available only for some location-based supergroups, use supergroupFullInfo.can_set_location to check whether the method is allowed to use @chat_id Chat identifier @location New location for the chat; must be valid and not null
func (c *Client) SetChatLocation(chatId int64, location *ChatLocation) (*Ok, error) {
	req := &SetChatLocation{
		ChatId:   chatId,
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatMemberStatus Changes the status of a chat member; requires can_invite_users member right to add a chat member, can_promote_members administrator right to change administrator rights of the member,
func (c *Client) SetChatMemberStatus(chatId int64, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	req := &SetChatMemberStatus{
		ChatId:   chatId,
		MemberId: memberId,
		Status:   status,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatMessageAutoDeleteTime Changes the message auto-delete or self-destruct (for secret chats) time in a chat. Requires change_info administrator right in basic groups, supergroups and channels.
func (c *Client) SetChatMessageAutoDeleteTime(chatId int64, messageAutoDeleteTime int32) (*Ok, error) {
	req := &SetChatMessageAutoDeleteTime{
		ChatId:                chatId,
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatMessageSender Selects a message sender to send messages in a chat @chat_id Chat identifier @message_sender_id New message sender for the chat
func (c *Client) SetChatMessageSender(chatId int64, messageSenderId *MessageSender) (*Ok, error) {
	req := &SetChatMessageSender{
		ChatId:          chatId,
		MessageSenderId: messageSenderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatNotificationSettings Changes the notification settings of a chat. Notification settings of a chat with the current user (Saved Messages) can't be changed
func (c *Client) SetChatNotificationSettings(chatId int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	req := &SetChatNotificationSettings{
		ChatId:               chatId,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatPaidMessageStarCount Changes the Telegram Star amount that must be paid to send a message to a supergroup chat; requires can_restrict_members administrator right and supergroupFullInfo.can_enable_paid_messages
func (c *Client) SetChatPaidMessageStarCount(chatId int64, paidMessageStarCount int64) (*Ok, error) {
	req := &SetChatPaidMessageStarCount{
		ChatId:               chatId,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatPermissions Changes the chat members permissions. Supported only for basic groups and supergroups. Requires can_restrict_members administrator right
func (c *Client) SetChatPermissions(chatId int64, permissions *ChatPermissions) (*Ok, error) {
	req := &SetChatPermissions{
		ChatId:      chatId,
		Permissions: permissions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires can_change_info member right
func (c *Client) SetChatPhoto(chatId int64, opts *SetChatPhotoOpts) (*Ok, error) {
	req := &SetChatPhoto{
		ChatId: chatId,
	}
	if opts != nil {
		req.Photo = opts.Photo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatPinnedStories Changes the list of pinned stories on a chat page; requires can_edit_stories administrator right in the chat
func (c *Client) SetChatPinnedStories(chatId int64, storyIds []int32) (*Ok, error) {
	req := &SetChatPinnedStories{
		ChatId:   chatId,
		StoryIds: storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatProfileAccentColor Changes accent color and background custom emoji for profile of a supergroup or channel chat. Requires can_change_info administrator right
func (c *Client) SetChatProfileAccentColor(chatId int64, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	req := &SetChatProfileAccentColor{
		ChatId:                         chatId,
		ProfileAccentColorId:           profileAccentColorId,
		ProfileBackgroundCustomEmojiId: profileBackgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatSlowModeDelay Changes the slow mode delay of a chat. Available only for supergroups; requires can_restrict_members administrator right @chat_id Chat identifier @slow_mode_delay New slow mode delay for the chat, in seconds; must be one of 0, 5, 10, 30, 60, 300, 900, 3600
func (c *Client) SetChatSlowModeDelay(chatId int64, slowModeDelay int32) (*Ok, error) {
	req := &SetChatSlowModeDelay{
		ChatId:        chatId,
		SlowModeDelay: slowModeDelay,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatTheme Changes the chat theme. Supported only in private and secret chats @chat_id Chat identifier @theme New chat theme; pass null to return the default theme
func (c *Client) SetChatTheme(chatId int64, theme *InputChatTheme) (*Ok, error) {
	req := &SetChatTheme{
		ChatId: chatId,
		Theme:  theme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires can_change_info member right
func (c *Client) SetChatTitle(chatId int64, title string) (*Ok, error) {
	req := &SetChatTitle{
		ChatId: chatId,
		Title:  title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetCloseFriends Changes the list of close friends of the current user @user_ids User identifiers of close friends; the users must be contacts of the current user
func (c *Client) SetCloseFriends(userIds []int64) (*Ok, error) {
	req := &SetCloseFriends{
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetCommands Sets the list of commands supported by the bot for the given user scope and language; for bots only
func (c *Client) SetCommands(languageCode string, commands []*BotCommand, opts *SetCommandsOpts) (*Ok, error) {
	req := &SetCommands{
		LanguageCode: languageCode,
		Commands:     commands,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetCustomEmojiStickerSetThumbnail Sets a custom emoji sticker set thumbnail
func (c *Client) SetCustomEmojiStickerSetThumbnail(name string, customEmojiId string) (*Ok, error) {
	req := &SetCustomEmojiStickerSetThumbnail{
		Name:          name,
		CustomEmojiId: customEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetCustomLanguagePack Adds or changes a custom local language pack to the current localization target
func (c *Client) SetCustomLanguagePack(info *LanguagePackInfo, strings []*LanguagePackString) (*Ok, error) {
	req := &SetCustomLanguagePack{
		Info:    info,
		Strings: strings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom local language pack. Can be called before authorization @language_pack_id Identifier of a previously added custom local language pack in the current localization target @new_string New language pack string
func (c *Client) SetCustomLanguagePackString(languagePackId string, newString *LanguagePackString) (*Ok, error) {
	req := &SetCustomLanguagePackString{
		LanguagePackId: languagePackId,
		NewString:      newString,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain @new_encryption_key New encryption key
func (c *Client) SetDatabaseEncryptionKey(newEncryptionKey string) (*Ok, error) {
	req := &SetDatabaseEncryptionKey{
		NewEncryptionKey: newEncryptionKey,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDefaultBackground Sets default background for chats; adds the background to the list of installed backgrounds
func (c *Client) SetDefaultBackground(forDarkTheme bool, opts *SetDefaultBackgroundOpts) (*Background, error) {
	req := &SetDefaultBackground{
		ForDarkTheme: forDarkTheme,
	}
	if opts != nil {
		req.Background = opts.Background
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Background), nil
}

// SetDefaultChannelAdministratorRights Sets default administrator rights for adding the bot to channel chats; for bots only @default_channel_administrator_rights Default administrator rights for adding the bot to channels; pass null to remove default rights
func (c *Client) SetDefaultChannelAdministratorRights(defaultChannelAdministratorRights *ChatAdministratorRights) (*Ok, error) {
	req := &SetDefaultChannelAdministratorRights{
		DefaultChannelAdministratorRights: defaultChannelAdministratorRights,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDefaultGroupAdministratorRights Sets default administrator rights for adding the bot to basic group and supergroup chats; for bots only @default_group_administrator_rights Default administrator rights for adding the bot to basic group and supergroup chats; pass null to remove default rights
func (c *Client) SetDefaultGroupAdministratorRights(defaultGroupAdministratorRights *ChatAdministratorRights) (*Ok, error) {
	req := &SetDefaultGroupAdministratorRights{
		DefaultGroupAdministratorRights: defaultGroupAdministratorRights,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDefaultMessageAutoDeleteTime Changes the default message auto-delete time for new chats @message_auto_delete_time New default message auto-delete time; must be from 0 up to 365 * 86400 and be divisible by 86400. If 0, then messages aren't deleted automatically
func (c *Client) SetDefaultMessageAutoDeleteTime(messageAutoDeleteTime *MessageAutoDeleteTime) (*Ok, error) {
	req := &SetDefaultMessageAutoDeleteTime{
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDefaultReactionType Changes type of default reaction for the current user @reaction_type New type of the default reaction. The paid reaction can't be set as default
func (c *Client) SetDefaultReactionType(reactionType *ReactionType) (*Ok, error) {
	req := &SetDefaultReactionType{
		ReactionType: reactionType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetDirectMessagesChatTopicIsMarkedAsUnread Changes the marked as unread state of the topic in a channel direct messages chat administered by the current user
func (c *Client) SetDirectMessagesChatTopicIsMarkedAsUnread(chatId int64, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	req := &SetDirectMessagesChatTopicIsMarkedAsUnread{
		ChatId:           chatId,
		TopicId:          topicId,
		IsMarkedAsUnread: isMarkedAsUnread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetEmojiStatus Changes the emoji status of the current user; for Telegram Premium users only @emoji_status New emoji status; pass null to switch to the default badge
func (c *Client) SetEmojiStatus(emojiStatus *EmojiStatus) (*Ok, error) {
	req := &SetEmojiStatus{
		EmojiStatus: emojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetFileGenerationProgress Informs TDLib on a file generation progress
func (c *Client) SetFileGenerationProgress(generationId string, expectedSize int64, localPrefixSize int64) (*Ok, error) {
	req := &SetFileGenerationProgress{
		GenerationId:    generationId,
		ExpectedSize:    expectedSize,
		LocalPrefixSize: localPrefixSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetForumTopicNotificationSettings Changes the notification settings of a forum topic in a forum supergroup chat or a chat with a bot with topics
func (c *Client) SetForumTopicNotificationSettings(chatId int64, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	req := &SetForumTopicNotificationSettings{
		ChatId:               chatId,
		ForumTopicId:         forumTopicId,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetGameScore Updates the game score of the specified user in the game; for bots only
func (c *Client) SetGameScore(chatId int64, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	req := &SetGameScore{
		ChatId:      chatId,
		MessageId:   messageId,
		EditMessage: editMessage,
		UserId:      userId,
		Score:       score,
		Force:       force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Message), nil
}

// SetGiftCollectionName Changes name of a gift collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) SetGiftCollectionName(ownerId *MessageSender, collectionId int32, name string) (*GiftCollection, error) {
	req := &SetGiftCollectionName{
		OwnerId:      ownerId,
		CollectionId: collectionId,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*GiftCollection), nil
}

// SetGiftResalePrice Changes resale price of a unique gift owned by the current user
func (c *Client) SetGiftResalePrice(receivedGiftId string, opts *SetGiftResalePriceOpts) (*Ok, error) {
	req := &SetGiftResalePrice{
		ReceivedGiftId: receivedGiftId,
	}
	if opts != nil {
		req.Price = opts.Price
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetGiftSettings Changes settings for gift receiving for the current user @settings The new settings
func (c *Client) SetGiftSettings(settings *GiftSettings) (*Ok, error) {
	req := &SetGiftSettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetGroupCallPaidMessageStarCount Changes the minimum number of Telegram Stars that must be paid by general participant for each sent message to a live story call. Requires groupCall.can_be_managed right
func (c *Client) SetGroupCallPaidMessageStarCount(groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	req := &SetGroupCallPaidMessageStarCount{
		GroupCallId:          groupCallId,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetGroupCallParticipantIsSpeaking Informs TDLib that speaking state of a participant of an active group call has changed. Returns identifier of the participant if it is found
func (c *Client) SetGroupCallParticipantIsSpeaking(groupCallId int32, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	req := &SetGroupCallParticipantIsSpeaking{
		GroupCallId: groupCallId,
		AudioSource: audioSource,
		IsSpeaking:  isSpeaking,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*MessageSender), nil
}

// SetGroupCallParticipantVolumeLevel Changes volume level of a participant of an active group call; not supported for live stories. If the current user can manage the group call or is the owner of the group call,
func (c *Client) SetGroupCallParticipantVolumeLevel(groupCallId int32, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	req := &SetGroupCallParticipantVolumeLevel{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		VolumeLevel:   volumeLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetInactiveSessionTtl Changes the period of inactivity after which sessions will automatically be terminated @inactive_session_ttl_days New number of days of inactivity before sessions will be automatically terminated; 1-366 days
func (c *Client) SetInactiveSessionTtl(inactiveSessionTtlDays int32) (*Ok, error) {
	req := &SetInactiveSessionTtl{
		InactiveSessionTtlDays: inactiveSessionTtlDays,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
func (c *Client) SetInlineGameScore(inlineMessageId string, editMessage bool, userId int64, score int32, force bool) (*Ok, error) {
	req := &SetInlineGameScore{
		InlineMessageId: inlineMessageId,
		EditMessage:     editMessage,
		UserId:          userId,
		Score:           score,
		Force:           force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetLiveStoryMessageSender Selects a message sender to send messages in a live story call
func (c *Client) SetLiveStoryMessageSender(groupCallId int32, messageSenderId *MessageSender) (*Ok, error) {
	req := &SetLiveStoryMessageSender{
		GroupCallId:     groupCallId,
		MessageSenderId: messageSenderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetLoginEmailAddress Changes the login email address of the user. The email address can be changed only if the current user already has login email and passwordState.login_email_address_pattern is non-empty,
func (c *Client) SetLoginEmailAddress(newLoginEmailAddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	req := &SetLoginEmailAddress{
		NewLoginEmailAddress: newLoginEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*EmailAddressAuthenticationCodeInfo), nil
}

// SetLogStream Sets new log stream for internal logging of TDLib. Can be called synchronously @log_stream New log stream
func (c *Client) SetLogStream(logStream *LogStream) (*Ok, error) {
	req := &SetLogStream{
		LogStream: logStream,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetLogTagVerbosityLevel Sets the verbosity level for a specified TDLib internal log tag. Can be called synchronously
func (c *Client) SetLogTagVerbosityLevel(tag string, newVerbosityLevel int32) (*Ok, error) {
	req := &SetLogTagVerbosityLevel{
		Tag:               tag,
		NewVerbosityLevel: newVerbosityLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib. Can be called synchronously
func (c *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*Ok, error) {
	req := &SetLogVerbosityLevel{
		NewVerbosityLevel: newVerbosityLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMainProfileTab Changes the main profile tab of the current user @main_profile_tab The new value of the main profile tab
func (c *Client) SetMainProfileTab(mainProfileTab *ProfileTab) (*Ok, error) {
	req := &SetMainProfileTab{
		MainProfileTab: mainProfileTab,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMenuButton Sets menu button for the given user or for all users; for bots only
func (c *Client) SetMenuButton(userId int64, menuButton *BotMenuButton) (*Ok, error) {
	req := &SetMenuButton{
		UserId:     userId,
		MenuButton: menuButton,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMessageFactCheck Changes the fact-check of a message. Can be only used if messageProperties.can_set_fact_check == true
func (c *Client) SetMessageFactCheck(chatId int64, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	req := &SetMessageFactCheck{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.Text = opts.Text
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMessageReactions Sets reactions on a message; for bots only
func (c *Client) SetMessageReactions(chatId int64, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	req := &SetMessageReactions{
		ChatId:        chatId,
		MessageId:     messageId,
		ReactionTypes: reactionTypes,
		IsBig:         isBig,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMessageSenderBlockList Changes the block list of a message sender. Currently, only users and supergroup chats can be blocked
func (c *Client) SetMessageSenderBlockList(senderId *MessageSender, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	req := &SetMessageSenderBlockList{
		SenderId: senderId,
	}
	if opts != nil {
		req.BlockList = opts.BlockList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetMessageSenderBotVerification Changes the verification status of a user or a chat by an owned bot
func (c *Client) SetMessageSenderBotVerification(botUserId int64, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	req := &SetMessageSenderBotVerification{
		BotUserId:         botUserId,
		VerifiedId:        verifiedId,
		CustomDescription: customDescription,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetName Changes the first and last name of the current user @first_name The new value of the first name for the current user; 1-64 characters @last_name The new value of the optional last name for the current user; 0-64 characters
func (c *Client) SetName(firstName string, lastName string) (*Ok, error) {
	req := &SetName{
		FirstName: firstName,
		LastName:  lastName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks,
func (c *Client) SetNetworkType(opts *SetNetworkTypeOpts) (*Ok, error) {
	req := &SetNetworkType{}
	if opts != nil {
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetNewChatPrivacySettings Changes privacy settings for new chat creation; can be used only if getOption("can_set_new_chat_privacy_settings") @settings New settings
func (c *Client) SetNewChatPrivacySettings(settings *NewChatPrivacySettings) (*Ok, error) {
	req := &SetNewChatPrivacySettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
func (c *Client) SetOption(name string, opts *SetOptionOpts) (*Ok, error) {
	req := &SetOption{
		Name: name,
	}
	if opts != nil {
		req.Value = opts.Value
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPaidMessageReactionType Changes type of paid message reaction of the current user on a message. The message must have paid reaction added by the current user
func (c *Client) SetPaidMessageReactionType(chatId int64, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	req := &SetPaidMessageReactionType{
		ChatId:    chatId,
		MessageId: messageId,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
func (c *Client) SetPassportElement(element *InputPassportElement, password string) (*PassportElement, error) {
	req := &SetPassportElement{
		Element:  element,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PassportElement), nil
}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed @user_id User identifier @errors The errors
func (c *Client) SetPassportElementErrors(userId int64, errors []*InputPassportElementError) (*Ok, error) {
	req := &SetPassportElementErrors{
		UserId: userId,
		Errors: errors,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPassword Changes the 2-step verification password for the current user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed
func (c *Client) SetPassword(oldPassword string, newPassword string, newHint string, setRecoveryEmailAddress bool, newRecoveryEmailAddress string) (*PasswordState, error) {
	req := &SetPassword{
		OldPassword:             oldPassword,
		NewPassword:             newPassword,
		NewHint:                 newHint,
		SetRecoveryEmailAddress: setRecoveryEmailAddress,
		NewRecoveryEmailAddress: newRecoveryEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// SetPersonalChat Changes the personal chat of the current user @chat_id Identifier of the new personal chat; pass 0 to remove the chat. Use getSuitablePersonalChats to get suitable chats
func (c *Client) SetPersonalChat(chatId int64) (*Ok, error) {
	req := &SetPersonalChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPinnedChats Changes the order of pinned chats @chat_list Chat list in which to change the order of pinned chats @chat_ids The new list of pinned chats
func (c *Client) SetPinnedChats(chatList *ChatList, chatIds []int64) (*Ok, error) {
	req := &SetPinnedChats{
		ChatList: chatList,
		ChatIds:  chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPinnedForumTopics Changes the order of pinned topics in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator right in the supergroup
func (c *Client) SetPinnedForumTopics(chatId int64, forumTopicIds []int32) (*Ok, error) {
	req := &SetPinnedForumTopics{
		ChatId:        chatId,
		ForumTopicIds: forumTopicIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPinnedGifts Changes the list of pinned gifts on the current user's or the channel's profile page; requires can_post_messages administrator right in the channel chat
func (c *Client) SetPinnedGifts(ownerId *MessageSender, receivedGiftIds []string) (*Ok, error) {
	req := &SetPinnedGifts{
		OwnerId:         ownerId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPinnedSavedMessagesTopics Changes the order of pinned Saved Messages topics @saved_messages_topic_ids Identifiers of the new pinned Saved Messages topics
func (c *Client) SetPinnedSavedMessagesTopics(savedMessagesTopicIds []int64) (*Ok, error) {
	req := &SetPinnedSavedMessagesTopics{
		SavedMessagesTopicIds: savedMessagesTopicIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetPollAnswer Changes the user answer to a poll. A poll in quiz mode can be answered only once
func (c *Client) SetPollAnswer(chatId int64, messageId int64, optionIds []int32) (*Ok, error) {
	req := &SetPollAnswer{
		ChatId:    chatId,
		MessageId: messageId,
		OptionIds: optionIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetProfileAccentColor Changes accent color and background custom emoji for profile of the current user; for Telegram Premium users only
func (c *Client) SetProfileAccentColor(profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	req := &SetProfileAccentColor{
		ProfileAccentColorId:           profileAccentColorId,
		ProfileBackgroundCustomEmojiId: profileBackgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetProfileAudioPosition Changes position of an audio file in the profile audio files of the current user
func (c *Client) SetProfileAudioPosition(fileId int32, afterFileId int32) (*Ok, error) {
	req := &SetProfileAudioPosition{
		FileId:      fileId,
		AfterFileId: afterFileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetProfilePhoto Changes a profile photo for the current user
func (c *Client) SetProfilePhoto(photo *InputChatPhoto, isPublic bool) (*Ok, error) {
	req := &SetProfilePhoto{
		Photo:    photo,
		IsPublic: isPublic,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetQuickReplyShortcutName Changes name of a quick reply shortcut @shortcut_id Unique identifier of the quick reply shortcut @name New name for the shortcut. Use checkQuickReplyShortcutName to check its validness
func (c *Client) SetQuickReplyShortcutName(shortcutId int32, name string) (*Ok, error) {
	req := &SetQuickReplyShortcutName{
		ShortcutId: shortcutId,
		Name:       name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetReactionNotificationSettings Changes notification settings for reactions @notification_settings The new notification settings for reactions
func (c *Client) SetReactionNotificationSettings(notificationSettings *ReactionNotificationSettings) (*Ok, error) {
	req := &SetReactionNotificationSettings{
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetReadDatePrivacySettings Changes privacy settings for message read date @settings New settings
func (c *Client) SetReadDatePrivacySettings(settings *ReadDatePrivacySettings) (*Ok, error) {
	req := &SetReadDatePrivacySettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetRecoveryEmailAddress Changes the 2-step verification recovery email address of the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed.
func (c *Client) SetRecoveryEmailAddress(password string, newRecoveryEmailAddress string) (*PasswordState, error) {
	req := &SetRecoveryEmailAddress{
		Password:                password,
		NewRecoveryEmailAddress: newRecoveryEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*PasswordState), nil
}

// SetSavedMessagesTagLabel Changes label of a Saved Messages tag; for Telegram Premium users only @tag The tag which label will be changed @label New label for the tag; 0-12 characters
func (c *Client) SetSavedMessagesTagLabel(tag *ReactionType, label string) (*Ok, error) {
	req := &SetSavedMessagesTagLabel{
		Tag:   tag,
		Label: label,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetScopeNotificationSettings Changes notification settings for chats of a given type @scope Types of chats for which to change the notification settings @notification_settings The new notification settings for the given scope
func (c *Client) SetScopeNotificationSettings(scope *NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) (*Ok, error) {
	req := &SetScopeNotificationSettings{
		Scope:                scope,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerEmojis Changes the list of emojis corresponding to a sticker. The sticker must belong to a regular or custom emoji sticker set that is owned by the current user
func (c *Client) SetStickerEmojis(sticker *InputFile, emojis string) (*Ok, error) {
	req := &SetStickerEmojis{
		Sticker: sticker,
		Emojis:  emojis,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerKeywords Changes the list of keywords of a sticker. The sticker must belong to a regular or custom emoji sticker set that is owned by the current user
func (c *Client) SetStickerKeywords(sticker *InputFile, keywords []string) (*Ok, error) {
	req := &SetStickerKeywords{
		Sticker:  sticker,
		Keywords: keywords,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerMaskPosition Changes the mask position of a mask sticker. The sticker must belong to a mask sticker set that is owned by the current user
func (c *Client) SetStickerMaskPosition(sticker *InputFile, opts *SetStickerMaskPositionOpts) (*Ok, error) {
	req := &SetStickerMaskPosition{
		Sticker: sticker,
	}
	if opts != nil {
		req.MaskPosition = opts.MaskPosition
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs. The sticker set must be owned by the current user
func (c *Client) SetStickerPositionInSet(sticker *InputFile, position int32) (*Ok, error) {
	req := &SetStickerPositionInSet{
		Sticker:  sticker,
		Position: position,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerSetThumbnail Sets a sticker set thumbnail
func (c *Client) SetStickerSetThumbnail(userId int64, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	req := &SetStickerSetThumbnail{
		UserId: userId,
		Name:   name,
	}
	if opts != nil {
		req.Thumbnail = opts.Thumbnail
		req.Format = opts.Format
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStickerSetTitle Sets a sticker set title @name Sticker set name. The sticker set must be owned by the current user @title New sticker set title
func (c *Client) SetStickerSetTitle(name string, title string) (*Ok, error) {
	req := &SetStickerSetTitle{
		Name:  name,
		Title: title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStoryAlbumName Changes name of an album of stories. If the album is owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat. Returns the changed album
func (c *Client) SetStoryAlbumName(chatId int64, storyAlbumId int32, name string) (*StoryAlbum, error) {
	req := &SetStoryAlbumName{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StoryAlbum), nil
}

// SetStoryPrivacySettings Changes privacy settings of a story. The method can be called only for stories posted on behalf of the current user and if story.can_set_privacy_settings == true
func (c *Client) SetStoryPrivacySettings(storyId int32, privacySettings *StoryPrivacySettings) (*Ok, error) {
	req := &SetStoryPrivacySettings{
		StoryId:         storyId,
		PrivacySettings: privacySettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetStoryReaction Changes chosen reaction on a story that has already been sent; not supported for live stories
func (c *Client) SetStoryReaction(storyPosterChatId int64, storyId int32, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	req := &SetStoryReaction{
		StoryPosterChatId:     storyPosterChatId,
		StoryId:               storyId,
		UpdateRecentReactions: updateRecentReactions,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetSupergroupCustomEmojiStickerSet Changes the custom emoji sticker set of a supergroup; requires can_change_info administrator right. The chat must have at least chatBoostFeatures.min_custom_emoji_sticker_set_boost_level boost level to pass the corresponding color
func (c *Client) SetSupergroupCustomEmojiStickerSet(supergroupId int64, customEmojiStickerSetId string) (*Ok, error) {
	req := &SetSupergroupCustomEmojiStickerSet{
		SupergroupId:            supergroupId,
		CustomEmojiStickerSetId: customEmojiStickerSetId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetSupergroupMainProfileTab Changes the main profile tab of the channel; requires can_change_info administrator right
func (c *Client) SetSupergroupMainProfileTab(supergroupId int64, mainProfileTab *ProfileTab) (*Ok, error) {
	req := &SetSupergroupMainProfileTab{
		SupergroupId:   supergroupId,
		MainProfileTab: mainProfileTab,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires can_change_info administrator right @supergroup_id Identifier of the supergroup @sticker_set_id New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (c *Client) SetSupergroupStickerSet(supergroupId int64, stickerSetId string) (*Ok, error) {
	req := &SetSupergroupStickerSet{
		SupergroupId: supergroupId,
		StickerSetId: stickerSetId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetSupergroupUnrestrictBoostCount Changes the number of times the supergroup must be boosted by a user to ignore slow mode and chat permission restrictions; requires can_restrict_members administrator right
func (c *Client) SetSupergroupUnrestrictBoostCount(supergroupId int64, unrestrictBoostCount int32) (*Ok, error) {
	req := &SetSupergroupUnrestrictBoostCount{
		SupergroupId:         supergroupId,
		UnrestrictBoostCount: unrestrictBoostCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetSupergroupUsername Changes the editable username of a supergroup or channel, requires owner privileges in the supergroup or channel
func (c *Client) SetSupergroupUsername(supergroupId int64, username string) (*Ok, error) {
	req := &SetSupergroupUsername{
		SupergroupId: supergroupId,
		Username:     username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
func (c *Client) SetTdlibParameters(useTestDc bool, databaseDirectory string, filesDirectory string, databaseEncryptionKey string, useFileDatabase bool, useChatInfoDatabase bool, useMessageDatabase bool, useSecretChats bool, apiId int32, apiHash string, systemLanguageCode string, deviceModel string, systemVersion string, applicationVersion string) (*Ok, error) {
	req := &SetTdlibParameters{
		UseTestDc:             useTestDc,
		DatabaseDirectory:     databaseDirectory,
		FilesDirectory:        filesDirectory,
		DatabaseEncryptionKey: databaseEncryptionKey,
		UseFileDatabase:       useFileDatabase,
		UseChatInfoDatabase:   useChatInfoDatabase,
		UseMessageDatabase:    useMessageDatabase,
		UseSecretChats:        useSecretChats,
		ApiId:                 apiId,
		ApiHash:               apiHash,
		SystemLanguageCode:    systemLanguageCode,
		DeviceModel:           deviceModel,
		SystemVersion:         systemVersion,
		ApplicationVersion:    applicationVersion,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUpgradedGiftColors Changes color scheme for the current user based on an owned or a hosted upgraded gift; for Telegram Premium users only
func (c *Client) SetUpgradedGiftColors(upgradedGiftColorsId string) (*Ok, error) {
	req := &SetUpgradedGiftColors{
		UpgradedGiftColorsId: upgradedGiftColorsId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUserEmojiStatus Changes the emoji status of a user; for bots only @user_id Identifier of the user @emoji_status New emoji status; pass null to switch to the default badge
func (c *Client) SetUserEmojiStatus(userId int64, emojiStatus *EmojiStatus) (*Ok, error) {
	req := &SetUserEmojiStatus{
		UserId:      userId,
		EmojiStatus: emojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUsername Changes the editable username of the current user
func (c *Client) SetUsername(username string) (*Ok, error) {
	req := &SetUsername{
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUserNote Changes a note of a contact user
func (c *Client) SetUserNote(userId int64, note *FormattedText) (*Ok, error) {
	req := &SetUserNote{
		UserId: userId,
		Note:   note,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUserPersonalProfilePhoto Changes a personal profile photo of a contact user @user_id User identifier @photo Profile photo to set; pass null to delete the photo; inputChatPhotoPrevious isn't supported in this function
func (c *Client) SetUserPersonalProfilePhoto(userId int64, photo *InputChatPhoto) (*Ok, error) {
	req := &SetUserPersonalProfilePhoto{
		UserId: userId,
		Photo:  photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUserPrivacySettingRules Changes user privacy settings @setting The privacy setting @rules The new privacy rules
func (c *Client) SetUserPrivacySettingRules(setting *UserPrivacySetting, rules *UserPrivacySettingRules) (*Ok, error) {
	req := &SetUserPrivacySettingRules{
		Setting: setting,
		Rules:   rules,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetUserSupportInfo Sets support information for the given user; for Telegram support only @user_id User identifier @message New information message
func (c *Client) SetUserSupportInfo(userId int64, message *FormattedText) (*UserSupportInfo, error) {
	req := &SetUserSupportInfo{
		UserId:  userId,
		Message: message,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UserSupportInfo), nil
}

// SetVideoChatDefaultParticipant Changes default participant identifier, on whose behalf a video chat in the chat will be joined
func (c *Client) SetVideoChatDefaultParticipant(chatId int64, defaultParticipantId *MessageSender) (*Ok, error) {
	req := &SetVideoChatDefaultParticipant{
		ChatId:               chatId,
		DefaultParticipantId: defaultParticipantId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SetVideoChatTitle Sets title of a video chat; requires groupCall.can_be_managed right @group_call_id Group call identifier @title New group call title; 1-64 characters
func (c *Client) SetVideoChatTitle(groupCallId int32, title string) (*Ok, error) {
	req := &SetVideoChatTitle{
		GroupCallId: groupCallId,
		Title:       title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ShareChatWithBot Shares a chat after pressing a keyboardButtonTypeRequestChat button with the bot
func (c *Client) ShareChatWithBot(chatId int64, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	req := &ShareChatWithBot{
		ChatId:       chatId,
		MessageId:    messageId,
		ButtonId:     buttonId,
		SharedChatId: sharedChatId,
		OnlyCheck:    onlyCheck,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SharePhoneNumber Shares the phone number of the current user with a mutual contact. Supposed to be called when the user clicks on chatActionBarSharePhoneNumber
func (c *Client) SharePhoneNumber(userId int64) (*Ok, error) {
	req := &SharePhoneNumber{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ShareUsersWithBot Shares users after pressing a keyboardButtonTypeRequestUsers button with the bot
func (c *Client) ShareUsersWithBot(chatId int64, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	req := &ShareUsersWithBot{
		ChatId:        chatId,
		MessageId:     messageId,
		ButtonId:      buttonId,
		SharedUserIds: sharedUserIds,
		OnlyCheck:     onlyCheck,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// StartGroupCallRecording Starts recording of an active group call; for video chats only. Requires groupCall.can_be_managed right
func (c *Client) StartGroupCallRecording(groupCallId int32, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	req := &StartGroupCallRecording{
		GroupCallId:            groupCallId,
		Title:                  title,
		RecordVideo:            recordVideo,
		UsePortraitOrientation: usePortraitOrientation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// StartGroupCallScreenSharing Starts screen sharing in a joined group call; not supported in live stories. Returns join response payload for tgcalls
func (c *Client) StartGroupCallScreenSharing(groupCallId int32, audioSourceId int32, payload string) (*Text, error) {
	req := &StartGroupCallScreenSharing{
		GroupCallId:   groupCallId,
		AudioSourceId: audioSourceId,
		Payload:       payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Text), nil
}

// StartLiveStory Starts a new live story on behalf of a chat; requires can_post_stories administrator right for channel chats
func (c *Client) StartLiveStory(chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	req := &StartLiveStory{
		ChatId:               chatId,
		PrivacySettings:      privacySettings,
		ProtectContent:       protectContent,
		IsRtmpStream:         isRtmpStream,
		EnableMessages:       enableMessages,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*StartLiveStoryResult), nil
}

// StartScheduledVideoChat Starts a scheduled video chat @group_call_id Group call identifier of the video chat
func (c *Client) StartScheduledVideoChat(groupCallId int32) (*Ok, error) {
	req := &StartScheduledVideoChat{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// StopBusinessPoll Stops a poll sent on behalf of a business account; for bots only
func (c *Client) StopBusinessPoll(businessConnectionId string, chatId int64, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	req := &StopBusinessPoll{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*BusinessMessage), nil
}

// StopPoll Stops a poll
func (c *Client) StopPoll(chatId int64, messageId int64, opts *StopPollOpts) (*Ok, error) {
	req := &StopPoll{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SuggestUserBirthdate Suggests a birthdate to another regular user with common messages and allowing non-paid messages
func (c *Client) SuggestUserBirthdate(userId int64, birthdate *Birthdate) (*Ok, error) {
	req := &SuggestUserBirthdate{
		UserId:    userId,
		Birthdate: birthdate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SuggestUserProfilePhoto Suggests a profile photo to another regular user with common messages and allowing non-paid messages
func (c *Client) SuggestUserProfilePhoto(userId int64, photo *InputChatPhoto) (*Ok, error) {
	req := &SuggestUserProfilePhoto{
		UserId: userId,
		Photo:  photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// SummarizeMessage Summarizes content of the message with non-empty summary_language_code
func (c *Client) SummarizeMessage(chatId int64, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	req := &SummarizeMessage{
		ChatId:                  chatId,
		MessageId:               messageId,
		TranslateToLanguageCode: translateToLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// SynchronizeLanguagePack Fetches the latest versions of all strings from a language pack in the current localization target from the server.
func (c *Client) SynchronizeLanguagePack(languagePackId string) (*Ok, error) {
	req := &SynchronizeLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (c *Client) TerminateAllOtherSessions() (*Ok, error) {
	req := &TerminateAllOtherSessions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TerminateSession Terminates a session of the current user @session_id Session identifier
func (c *Client) TerminateSession(sessionId string) (*Ok, error) {
	req := &TerminateSession{
		SessionId: sessionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization @x Bytes to return
func (c *Client) TestCallBytes(x string) (*TestBytes, error) {
	req := &TestCallBytes{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestBytes), nil
}

// TestCallEmpty Does nothing; for testing only. This is an offline method. Can be called before authorization
func (c *Client) TestCallEmpty() (*Ok, error) {
	req := &TestCallEmpty{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization @x String to return
func (c *Client) TestCallString(x string) (*TestString, error) {
	req := &TestCallString{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestString), nil
}

// TestCallVectorInt Returns the received vector of numbers; for testing only. This is an offline method. Can be called before authorization @x Vector of numbers to return
func (c *Client) TestCallVectorInt(x []int32) (*TestVectorInt, error) {
	req := &TestCallVectorInt{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestVectorInt), nil
}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only. This is an offline method. Can be called before authorization @x Vector of objects to return
func (c *Client) TestCallVectorIntObject(x []*TestInt) (*TestVectorIntObject, error) {
	req := &TestCallVectorIntObject{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestVectorIntObject), nil
}

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization @x Vector of strings to return
func (c *Client) TestCallVectorString(x []string) (*TestVectorString, error) {
	req := &TestCallVectorString{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestVectorString), nil
}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only. This is an offline method. Can be called before authorization @x Vector of objects to return
func (c *Client) TestCallVectorStringObject(x []*TestString) (*TestVectorStringObject, error) {
	req := &TestCallVectorStringObject{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestVectorStringObject), nil
}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (c *Client) TestGetDifference() (*Ok, error) {
	req := &TestGetDifference{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only. Can be called before authorization
func (c *Client) TestNetwork() (*Ok, error) {
	req := &TestNetwork{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TestProxy Sends a simple network request to the Telegram servers via proxy; for testing only. Can be called before authorization
func (c *Client) TestProxy(server string, port int32, typeField *ProxyType, dcId int32, timeout float64) (*Ok, error) {
	req := &TestProxy{
		Server:    server,
		Port:      port,
		TypeField: typeField,
		DcId:      dcId,
		Timeout:   timeout,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously @error The error to be returned
func (c *Client) TestReturnError(error *Error) (*Error, error) {
	req := &TestReturnError{
		Error: error,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Error), nil
}

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization @x Number to square
func (c *Client) TestSquareInt(x int32) (*TestInt, error) {
	req := &TestSquareInt{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*TestInt), nil
}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (c *Client) TestUseUpdate() (*Update, error) {
	req := &TestUseUpdate{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Update), nil
}

// ToggleAllDownloadsArePaused Changes pause state of all files in the file download list @are_paused Pass true to pause all downloads; pass false to unpause them
func (c *Client) ToggleAllDownloadsArePaused(arePaused bool) (*Ok, error) {
	req := &ToggleAllDownloadsArePaused{
		ArePaused: arePaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleBotCanManageEmojiStatus Toggles whether the bot can manage emoji status of the current user @bot_user_id User identifier of the bot @can_manage_emoji_status Pass true if the bot is allowed to change emoji status of the user; pass false otherwise
func (c *Client) ToggleBotCanManageEmojiStatus(botUserId int64, canManageEmojiStatus bool) (*Ok, error) {
	req := &ToggleBotCanManageEmojiStatus{
		BotUserId:            botUserId,
		CanManageEmojiStatus: canManageEmojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleBotIsAddedToAttachmentMenu Adds or removes a bot to attachment and side menu. Bot can be added to the menu, only if userTypeBot.can_be_added_to_attachment_menu == true
func (c *Client) ToggleBotIsAddedToAttachmentMenu(botUserId int64, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	req := &ToggleBotIsAddedToAttachmentMenu{
		BotUserId:        botUserId,
		IsAdded:          isAdded,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleBotUsernameIsActive Changes active state for a username of a bot. The editable username can be disabled only if there are other active usernames.
func (c *Client) ToggleBotUsernameIsActive(botUserId int64, username string, isActive bool) (*Ok, error) {
	req := &ToggleBotUsernameIsActive{
		BotUserId: botUserId,
		Username:  username,
		IsActive:  isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleBusinessConnectedBotChatIsPaused Pauses or resumes the connected business bot in a specific chat @chat_id Chat identifier @is_paused Pass true to pause the connected bot in the chat; pass false to resume the bot
func (c *Client) ToggleBusinessConnectedBotChatIsPaused(chatId int64, isPaused bool) (*Ok, error) {
	req := &ToggleBusinessConnectedBotChatIsPaused{
		ChatId:   chatId,
		IsPaused: isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat @chat_id Chat identifier @default_disable_notification New value of default_disable_notification
func (c *Client) ToggleChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) (*Ok, error) {
	req := &ToggleChatDefaultDisableNotification{
		ChatId:                     chatId,
		DefaultDisableNotification: defaultDisableNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatFolderTags Toggles whether chat folder tags are enabled @are_tags_enabled Pass true to enable folder tags; pass false to disable them
func (c *Client) ToggleChatFolderTags(areTagsEnabled bool) (*Ok, error) {
	req := &ToggleChatFolderTags{
		AreTagsEnabled: areTagsEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatGiftNotifications Toggles whether notifications for new gifts received by a channel chat are sent to the current user; requires can_post_messages administrator right in the chat
func (c *Client) ToggleChatGiftNotifications(chatId int64, areEnabled bool) (*Ok, error) {
	req := &ToggleChatGiftNotifications{
		ChatId:     chatId,
		AreEnabled: areEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatHasProtectedContent Changes the ability of users to save, forward, or copy chat content. Supported only for basic groups, supergroups and channels. Requires owner privileges
func (c *Client) ToggleChatHasProtectedContent(chatId int64, hasProtectedContent bool) (*Ok, error) {
	req := &ToggleChatHasProtectedContent{
		ChatId:              chatId,
		HasProtectedContent: hasProtectedContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat @chat_id Chat identifier @is_marked_as_unread New value of is_marked_as_unread
func (c *Client) ToggleChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) (*Ok, error) {
	req := &ToggleChatIsMarkedAsUnread{
		ChatId:           chatId,
		IsMarkedAsUnread: isMarkedAsUnread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatIsPinned Changes the pinned state of a chat. There can be up to getOption("pinned_chat_count_max")/getOption("pinned_archived_chat_count_max") pinned non-secret chats and the same number of secret chats in the main/archive chat list. The limit can be increased with Telegram Premium
func (c *Client) ToggleChatIsPinned(chatList *ChatList, chatId int64, isPinned bool) (*Ok, error) {
	req := &ToggleChatIsPinned{
		ChatList: chatList,
		ChatId:   chatId,
		IsPinned: isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatIsTranslatable Changes the translatable state of a chat @chat_id Chat identifier @is_translatable New value of is_translatable
func (c *Client) ToggleChatIsTranslatable(chatId int64, isTranslatable bool) (*Ok, error) {
	req := &ToggleChatIsTranslatable{
		ChatId:         chatId,
		IsTranslatable: isTranslatable,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleChatViewAsTopics Changes the view_as_topics setting of a forum chat or Saved Messages @chat_id Chat identifier @view_as_topics New value of view_as_topics
func (c *Client) ToggleChatViewAsTopics(chatId int64, viewAsTopics bool) (*Ok, error) {
	req := &ToggleChatViewAsTopics{
		ChatId:       chatId,
		ViewAsTopics: viewAsTopics,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages Allows to send unpaid messages to the given topic of the channel direct messages chat administered by the current user
func (c *Client) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(chatId int64, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	req := &ToggleDirectMessagesChatTopicCanSendUnpaidMessages{
		ChatId:                chatId,
		TopicId:               topicId,
		CanSendUnpaidMessages: canSendUnpaidMessages,
		RefundPayments:        refundPayments,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleDownloadIsPaused Changes pause state of a file in the file download list
func (c *Client) ToggleDownloadIsPaused(fileId int32, isPaused bool) (*Ok, error) {
	req := &ToggleDownloadIsPaused{
		FileId:   fileId,
		IsPaused: isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleForumTopicIsClosed Toggles whether a topic is closed in a forum supergroup chat; requires can_manage_topics administrator right in the supergroup unless the user is creator of the topic
func (c *Client) ToggleForumTopicIsClosed(chatId int64, forumTopicId int32, isClosed bool) (*Ok, error) {
	req := &ToggleForumTopicIsClosed{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		IsClosed:     isClosed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleForumTopicIsPinned Changes the pinned state of a topic in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator right in the supergroup.
func (c *Client) ToggleForumTopicIsPinned(chatId int64, forumTopicId int32, isPinned bool) (*Ok, error) {
	req := &ToggleForumTopicIsPinned{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		IsPinned:     isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGeneralForumTopicIsHidden Toggles whether a General topic is hidden in a forum supergroup chat; requires can_manage_topics administrator right in the supergroup
func (c *Client) ToggleGeneralForumTopicIsHidden(chatId int64, isHidden bool) (*Ok, error) {
	req := &ToggleGeneralForumTopicIsHidden{
		ChatId:   chatId,
		IsHidden: isHidden,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGiftIsSaved Toggles whether a gift is shown on the current user's or the channel's profile page; requires can_post_messages administrator right in the channel chat
func (c *Client) ToggleGiftIsSaved(receivedGiftId string, isSaved bool) (*Ok, error) {
	req := &ToggleGiftIsSaved{
		ReceivedGiftId: receivedGiftId,
		IsSaved:        isSaved,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallAreMessagesAllowed Toggles whether participants of a group call can send messages there. Requires groupCall.can_toggle_are_messages_allowed right
func (c *Client) ToggleGroupCallAreMessagesAllowed(groupCallId int32, areMessagesAllowed bool) (*Ok, error) {
	req := &ToggleGroupCallAreMessagesAllowed{
		GroupCallId:        groupCallId,
		AreMessagesAllowed: areMessagesAllowed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallIsMyVideoEnabled Toggles whether current user's video is enabled @group_call_id Group call identifier @is_my_video_enabled Pass true if the current user's video is enabled
func (c *Client) ToggleGroupCallIsMyVideoEnabled(groupCallId int32, isMyVideoEnabled bool) (*Ok, error) {
	req := &ToggleGroupCallIsMyVideoEnabled{
		GroupCallId:      groupCallId,
		IsMyVideoEnabled: isMyVideoEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallIsMyVideoPaused Toggles whether current user's video is paused @group_call_id Group call identifier @is_my_video_paused Pass true if the current user's video is paused
func (c *Client) ToggleGroupCallIsMyVideoPaused(groupCallId int32, isMyVideoPaused bool) (*Ok, error) {
	req := &ToggleGroupCallIsMyVideoPaused{
		GroupCallId:     groupCallId,
		IsMyVideoPaused: isMyVideoPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallParticipantIsHandRaised Toggles whether a group call participant hand is rased; for video chats only
func (c *Client) ToggleGroupCallParticipantIsHandRaised(groupCallId int32, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	req := &ToggleGroupCallParticipantIsHandRaised{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		IsHandRaised:  isHandRaised,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallParticipantIsMuted Toggles whether a participant of an active group call is muted, unmuted, or allowed to unmute themselves; not supported for live stories
func (c *Client) ToggleGroupCallParticipantIsMuted(groupCallId int32, participantId *MessageSender, isMuted bool) (*Ok, error) {
	req := &ToggleGroupCallParticipantIsMuted{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		IsMuted:       isMuted,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleGroupCallScreenSharingIsPaused Pauses or unpauses screen sharing in a joined group call; not supported in live stories @group_call_id Group call identifier @is_paused Pass true to pause screen sharing; pass false to unpause it
func (c *Client) ToggleGroupCallScreenSharingIsPaused(groupCallId int32, isPaused bool) (*Ok, error) {
	req := &ToggleGroupCallScreenSharingIsPaused{
		GroupCallId: groupCallId,
		IsPaused:    isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleHasSponsoredMessagesEnabled Toggles whether the current user has sponsored messages enabled. The setting has no effect for users without Telegram Premium for which sponsored messages are always enabled
func (c *Client) ToggleHasSponsoredMessagesEnabled(hasSponsoredMessagesEnabled bool) (*Ok, error) {
	req := &ToggleHasSponsoredMessagesEnabled{
		HasSponsoredMessagesEnabled: hasSponsoredMessagesEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSavedMessagesTopicIsPinned Changes the pinned state of a Saved Messages topic. There can be up to getOption("pinned_saved_messages_topic_count_max") pinned topics. The limit can be increased with Telegram Premium
func (c *Client) ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId int64, isPinned bool) (*Ok, error) {
	req := &ToggleSavedMessagesTopicIsPinned{
		SavedMessagesTopicId: savedMessagesTopicId,
		IsPinned:             isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSessionCanAcceptCalls Toggles whether a session can accept incoming calls @session_id Session identifier @can_accept_calls Pass true to allow accepting incoming calls by the session; pass false otherwise
func (c *Client) ToggleSessionCanAcceptCalls(sessionId string, canAcceptCalls bool) (*Ok, error) {
	req := &ToggleSessionCanAcceptCalls{
		SessionId:      sessionId,
		CanAcceptCalls: canAcceptCalls,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSessionCanAcceptSecretChats Toggles whether a session can accept incoming secret chats @session_id Session identifier @can_accept_secret_chats Pass true to allow accepting secret chats by the session; pass false otherwise
func (c *Client) ToggleSessionCanAcceptSecretChats(sessionId string, canAcceptSecretChats bool) (*Ok, error) {
	req := &ToggleSessionCanAcceptSecretChats{
		SessionId:            sessionId,
		CanAcceptSecretChats: canAcceptSecretChats,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleStoryIsPostedToChatPage Toggles whether a story is accessible after expiration. Can be called only if story.can_toggle_is_posted_to_chat_page == true
func (c *Client) ToggleStoryIsPostedToChatPage(storyPosterChatId int64, storyId int32, isPostedToChatPage bool) (*Ok, error) {
	req := &ToggleStoryIsPostedToChatPage{
		StoryPosterChatId:  storyPosterChatId,
		StoryId:            storyId,
		IsPostedToChatPage: isPostedToChatPage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupCanHaveSponsoredMessages Toggles whether sponsored messages are shown in the channel chat; requires owner privileges in the channel. The chat must have at least chatBoostFeatures.min_sponsored_message_disable_boost_level boost level to disable sponsored messages
func (c *Client) ToggleSupergroupCanHaveSponsoredMessages(supergroupId int64, canHaveSponsoredMessages bool) (*Ok, error) {
	req := &ToggleSupergroupCanHaveSponsoredMessages{
		SupergroupId:             supergroupId,
		CanHaveSponsoredMessages: canHaveSponsoredMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled Toggles whether aggressive anti-spam checks are enabled in the supergroup. Can be called only if supergroupFullInfo.can_toggle_aggressive_anti_spam == true
func (c *Client) ToggleSupergroupHasAggressiveAntiSpamEnabled(supergroupId int64, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	req := &ToggleSupergroupHasAggressiveAntiSpamEnabled{
		SupergroupId:                 supergroupId,
		HasAggressiveAntiSpamEnabled: hasAggressiveAntiSpamEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupHasAutomaticTranslation Toggles whether messages are automatically translated in the channel chat; requires can_change_info administrator right in the channel.
func (c *Client) ToggleSupergroupHasAutomaticTranslation(supergroupId int64, hasAutomaticTranslation bool) (*Ok, error) {
	req := &ToggleSupergroupHasAutomaticTranslation{
		SupergroupId:            supergroupId,
		HasAutomaticTranslation: hasAutomaticTranslation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupHasHiddenMembers Toggles whether non-administrators can receive only administrators and bots using getSupergroupMembers or searchChatMembers. Can be called only if supergroupFullInfo.can_hide_members == true
func (c *Client) ToggleSupergroupHasHiddenMembers(supergroupId int64, hasHiddenMembers bool) (*Ok, error) {
	req := &ToggleSupergroupHasHiddenMembers{
		SupergroupId:     supergroupId,
		HasHiddenMembers: hasHiddenMembers,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires can_change_info member right @supergroup_id The identifier of the supergroup @is_all_history_available The new value of is_all_history_available
func (c *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupId int64, isAllHistoryAvailable bool) (*Ok, error) {
	req := &ToggleSupergroupIsAllHistoryAvailable{
		SupergroupId:          supergroupId,
		IsAllHistoryAvailable: isAllHistoryAvailable,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupIsBroadcastGroup Upgrades supergroup to a broadcast group; requires owner privileges in the supergroup @supergroup_id Identifier of the supergroup
func (c *Client) ToggleSupergroupIsBroadcastGroup(supergroupId int64) (*Ok, error) {
	req := &ToggleSupergroupIsBroadcastGroup{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupIsForum Toggles whether the supergroup is a forum; requires owner privileges in the supergroup. Discussion supergroups can't be converted to forums
func (c *Client) ToggleSupergroupIsForum(supergroupId int64, isForum bool, hasForumTabs bool) (*Ok, error) {
	req := &ToggleSupergroupIsForum{
		SupergroupId: supergroupId,
		IsForum:      isForum,
		HasForumTabs: hasForumTabs,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupJoinByRequest Toggles whether all users directly joining the supergroup need to be approved by supergroup administrators; requires can_restrict_members administrator right
func (c *Client) ToggleSupergroupJoinByRequest(supergroupId int64, joinByRequest bool) (*Ok, error) {
	req := &ToggleSupergroupJoinByRequest{
		SupergroupId:  supergroupId,
		JoinByRequest: joinByRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupJoinToSendMessages Toggles whether joining is mandatory to send messages to a discussion supergroup; requires can_restrict_members administrator right
func (c *Client) ToggleSupergroupJoinToSendMessages(supergroupId int64, joinToSendMessages bool) (*Ok, error) {
	req := &ToggleSupergroupJoinToSendMessages{
		SupergroupId:       supergroupId,
		JoinToSendMessages: joinToSendMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupSignMessages Toggles whether sender signature or link to the account is added to sent messages in a channel; requires can_change_info member right
func (c *Client) ToggleSupergroupSignMessages(supergroupId int64, signMessages bool, showMessageSender bool) (*Ok, error) {
	req := &ToggleSupergroupSignMessages{
		SupergroupId:      supergroupId,
		SignMessages:      signMessages,
		ShowMessageSender: showMessageSender,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleSupergroupUsernameIsActive Changes active state for a username of a supergroup or channel, requires owner privileges in the supergroup or channel. The editable username can't be disabled.
func (c *Client) ToggleSupergroupUsernameIsActive(supergroupId int64, username string, isActive bool) (*Ok, error) {
	req := &ToggleSupergroupUsernameIsActive{
		SupergroupId: supergroupId,
		Username:     username,
		IsActive:     isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleUsernameIsActive Changes active state for a username of the current user. The editable username can't be disabled. May return an error with a message "USERNAMES_ACTIVE_TOO_MUCH" if the maximum number of active usernames has been reached
func (c *Client) ToggleUsernameIsActive(username string, isActive bool) (*Ok, error) {
	req := &ToggleUsernameIsActive{
		Username: username,
		IsActive: isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleVideoChatEnabledStartNotification Toggles whether the current user will receive a notification when the video chat starts; for scheduled video chats only
func (c *Client) ToggleVideoChatEnabledStartNotification(groupCallId int32, enabledStartNotification bool) (*Ok, error) {
	req := &ToggleVideoChatEnabledStartNotification{
		GroupCallId:              groupCallId,
		EnabledStartNotification: enabledStartNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ToggleVideoChatMuteNewParticipants Toggles whether new participants of a video chat can be unmuted only by administrators of the video chat. Requires groupCall.can_toggle_mute_new_participants right
func (c *Client) ToggleVideoChatMuteNewParticipants(groupCallId int32, muteNewParticipants bool) (*Ok, error) {
	req := &ToggleVideoChatMuteNewParticipants{
		GroupCallId:         groupCallId,
		MuteNewParticipants: muteNewParticipants,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TransferBusinessAccountStars Transfers Telegram Stars from the business account to the business bot; for bots only
func (c *Client) TransferBusinessAccountStars(businessConnectionId string, starCount int64) (*Ok, error) {
	req := &TransferBusinessAccountStars{
		BusinessConnectionId: businessConnectionId,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TransferChatOwnership Changes the owner of a chat; requires owner privileges in the chat. Use the method canTransferOwnership to check whether the ownership can be transferred from the current session. Available only for supergroups and channel chats
func (c *Client) TransferChatOwnership(chatId int64, userId int64, password string) (*Ok, error) {
	req := &TransferChatOwnership{
		ChatId:   chatId,
		UserId:   userId,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TransferGift Sends an upgraded gift to another user or channel chat
func (c *Client) TransferGift(businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	req := &TransferGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
		NewOwnerId:           newOwnerId,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// TranslateMessageText Extracts text or caption of the given message and translates it to the given language. If the current user is a Telegram Premium user, then text formatting is preserved
func (c *Client) TranslateMessageText(chatId int64, messageId int64, toLanguageCode string) (*FormattedText, error) {
	req := &TranslateMessageText{
		ChatId:         chatId,
		MessageId:      messageId,
		ToLanguageCode: toLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// TranslateText Translates a text to the given language. If the current user is a Telegram Premium user, then text formatting is preserved
func (c *Client) TranslateText(text *FormattedText, toLanguageCode string) (*FormattedText, error) {
	req := &TranslateText{
		Text:           text,
		ToLanguageCode: toLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*FormattedText), nil
}

// UnpinAllChatMessages Removes all pinned messages from a chat; requires can_pin_messages member right if the chat is a basic group or supergroup, or can_edit_messages administrator right if the chat is a channel @chat_id Identifier of the chat
func (c *Client) UnpinAllChatMessages(chatId int64) (*Ok, error) {
	req := &UnpinAllChatMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// UnpinAllDirectMessagesChatTopicMessages Removes all pinned messages from the topic in a channel direct messages chat administered by the current user
func (c *Client) UnpinAllDirectMessagesChatTopicMessages(chatId int64, topicId int64) (*Ok, error) {
	req := &UnpinAllDirectMessagesChatTopicMessages{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// UnpinAllForumTopicMessages Removes all pinned messages from a topic in a forum supergroup chat or a chat with a bot with topics; requires can_pin_messages member right in the supergroup
func (c *Client) UnpinAllForumTopicMessages(chatId int64, forumTopicId int32) (*Ok, error) {
	req := &UnpinAllForumTopicMessages{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// UnpinChatMessage Removes a pinned message from a chat; requires can_pin_messages member right if the chat is a basic group or supergroup, or can_edit_messages administrator right if the chat is a channel @chat_id Identifier of the chat @message_id Identifier of the removed pinned message
func (c *Client) UnpinChatMessage(chatId int64, messageId int64) (*Ok, error) {
	req := &UnpinChatMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires owner privileges. Deactivates the original basic group @chat_id Identifier of the chat to upgrade
func (c *Client) UpgradeBasicGroupChatToSupergroupChat(chatId int64) (*Chat, error) {
	req := &UpgradeBasicGroupChatToSupergroupChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Chat), nil
}

// UpgradeGift Upgrades a regular gift
func (c *Client) UpgradeGift(businessConnectionId string, receivedGiftId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	req := &UpgradeGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
		KeepOriginalDetails:  keepOriginalDetails,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*UpgradeGiftResult), nil
}

// UploadStickerFile Uploads a file with a sticker; returns the uploaded file
func (c *Client) UploadStickerFile(userId int64, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	req := &UploadStickerFile{
		UserId:        userId,
		StickerFormat: stickerFormat,
		Sticker:       sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*File), nil
}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
func (c *Client) ValidateOrderInfo(inputInvoice *InputInvoice, allowSave bool, opts *ValidateOrderInfoOpts) (*ValidatedOrderInfo, error) {
	req := &ValidateOrderInfo{
		InputInvoice: inputInvoice,
		AllowSave:    allowSave,
	}
	if opts != nil {
		req.OrderInfo = opts.OrderInfo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*ValidatedOrderInfo), nil
}

// ViewMessages Informs TDLib that messages are being viewed by the user. Sponsored messages must be marked as viewed only when the entire text of the message is shown on the screen (excluding the button).
func (c *Client) ViewMessages(chatId int64, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	req := &ViewMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
		ForceRead:  forceRead,
	}
	if opts != nil {
		req.Source = opts.Source
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ViewPremiumFeature Informs TDLib that the user viewed detailed information about a Premium feature on the Premium features screen @feature The viewed premium feature
func (c *Client) ViewPremiumFeature(feature *PremiumFeature) (*Ok, error) {
	req := &ViewPremiumFeature{
		Feature: feature,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ViewSponsoredChat Informs TDLib that the user fully viewed a sponsored chat @sponsored_chat_unique_id Unique identifier of the sponsored chat
func (c *Client) ViewSponsoredChat(sponsoredChatUniqueId int64) (*Ok, error) {
	req := &ViewSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user @sticker_set_ids Identifiers of viewed trending sticker sets
func (c *Client) ViewTrendingStickerSets(stickerSetIds []string) (*Ok, error) {
	req := &ViewTrendingStickerSets{
		StickerSetIds: stickerSetIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// ViewVideoMessageAdvertisement Informs TDLib that the user viewed a video message advertisement @advertisement_unique_id Unique identifier of the advertisement
func (c *Client) ViewVideoMessageAdvertisement(advertisementUniqueId int64) (*Ok, error) {
	req := &ViewVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}

// WriteGeneratedFilePart Writes a part of a generated file. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct write to the destination file
func (c *Client) WriteGeneratedFilePart(generationId string, offset int64, data string) (*Ok, error) {
	req := &WriteGeneratedFilePart{
		GenerationId: generationId,
		Offset:       offset,
		Data:         data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ok), nil
}
