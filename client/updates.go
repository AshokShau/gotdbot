package client

import "github.com/AshokShau/gotdbot/types"

// OnUpdateAuthorizationState The user authorization state has changed @authorization_state New authorization state
func (c *Client) OnUpdateAuthorizationState(fn func(client *Client, update *types.UpdateAuthorizationState) error, filter FilterFunc, position int) {
	c.AddHandler("updateAuthorizationState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAuthorizationState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewMessage A new message was received; can also be an outgoing message @message The new message
func (c *Client) OnUpdateNewMessage(fn func(client *Client, update *types.UpdateNewMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageSendAcknowledged A request to send a message has reached the Telegram server. This doesn't mean that the message will be sent successfully.
func (c *Client) OnUpdateMessageSendAcknowledged(fn func(client *Client, update *types.UpdateMessageSendAcknowledged) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageSendAcknowledged", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageSendAcknowledged); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageSendSucceeded A message has been successfully sent
func (c *Client) OnUpdateMessageSendSucceeded(fn func(client *Client, update *types.UpdateMessageSendSucceeded) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageSendSucceeded", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageSendSucceeded); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageSendFailed A message failed to send. Be aware that some messages being sent can be irrecoverably deleted, in which case updateDeleteMessages will be received instead of this update
func (c *Client) OnUpdateMessageSendFailed(fn func(client *Client, update *types.UpdateMessageSendFailed) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageSendFailed", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageSendFailed); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageContent The message content has changed @chat_id Chat identifier @message_id Message identifier @new_content New message content
func (c *Client) OnUpdateMessageContent(fn func(client *Client, update *types.UpdateMessageContent) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageContent", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageContent); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageEdited A message was edited. Changes in the message content will come in a separate updateMessageContent
func (c *Client) OnUpdateMessageEdited(fn func(client *Client, update *types.UpdateMessageEdited) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageEdited", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageEdited); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageIsPinned The message pinned state was changed @chat_id Chat identifier @message_id The message identifier @is_pinned True, if the message is pinned
func (c *Client) OnUpdateMessageIsPinned(fn func(client *Client, update *types.UpdateMessageIsPinned) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageIsPinned", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageIsPinned); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageInteractionInfo The information about interactions with a message has changed @chat_id Chat identifier @message_id Message identifier @interaction_info New information about interactions with the message; may be null
func (c *Client) OnUpdateMessageInteractionInfo(fn func(client *Client, update *types.UpdateMessageInteractionInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageInteractionInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageInteractionInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageContentOpened The message content was opened. Updates voice note messages to "listened", video note messages to "viewed" and starts the self-destruct timer @chat_id Chat identifier @message_id Message identifier
func (c *Client) OnUpdateMessageContentOpened(fn func(client *Client, update *types.UpdateMessageContentOpened) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageContentOpened", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageContentOpened); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageMentionRead A message with an unread mention was read @chat_id Chat identifier @message_id Message identifier @unread_mention_count The new number of unread mention messages left in the chat
func (c *Client) OnUpdateMessageMentionRead(fn func(client *Client, update *types.UpdateMessageMentionRead) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageMentionRead", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageMentionRead); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageUnreadReactions The list of unread reactions added to a message was changed
func (c *Client) OnUpdateMessageUnreadReactions(fn func(client *Client, update *types.UpdateMessageUnreadReactions) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageUnreadReactions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageUnreadReactions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageFactCheck A fact-check added to a message was changed
func (c *Client) OnUpdateMessageFactCheck(fn func(client *Client, update *types.UpdateMessageFactCheck) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageFactCheck", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageFactCheck); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageSuggestedPostInfo Information about suggested post of a message was changed
func (c *Client) OnUpdateMessageSuggestedPostInfo(fn func(client *Client, update *types.UpdateMessageSuggestedPostInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageSuggestedPostInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageSuggestedPostInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageLiveLocationViewed A message with a live location was viewed. When the update is received, the application is expected to update the live location
func (c *Client) OnUpdateMessageLiveLocationViewed(fn func(client *Client, update *types.UpdateMessageLiveLocationViewed) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageLiveLocationViewed", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageLiveLocationViewed); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateVideoPublished An automatically scheduled message with video has been successfully sent after conversion
func (c *Client) OnUpdateVideoPublished(fn func(client *Client, update *types.UpdateVideoPublished) error, filter FilterFunc, position int) {
	c.AddHandler("updateVideoPublished", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateVideoPublished); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewChat A new chat has been loaded/created. This update is guaranteed to come before the chat identifier is returned to the application. The chat field changes will be reported through separate updates @chat The chat
func (c *Client) OnUpdateNewChat(fn func(client *Client, update *types.UpdateNewChat) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewChat", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewChat); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatTitle The title of a chat was changed @chat_id Chat identifier @title The new chat title
func (c *Client) OnUpdateChatTitle(fn func(client *Client, update *types.UpdateChatTitle) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatTitle", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatTitle); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatPhoto A chat photo was changed @chat_id Chat identifier @photo The new chat photo; may be null
func (c *Client) OnUpdateChatPhoto(fn func(client *Client, update *types.UpdateChatPhoto) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatPhoto", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatPhoto); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatAccentColors Chat accent colors have changed
func (c *Client) OnUpdateChatAccentColors(fn func(client *Client, update *types.UpdateChatAccentColors) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatAccentColors", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatAccentColors); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatPermissions Chat permissions were changed @chat_id Chat identifier @permissions The new chat permissions
func (c *Client) OnUpdateChatPermissions(fn func(client *Client, update *types.UpdateChatPermissions) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatPermissions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatPermissions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatLastMessage The last message of a chat was changed
func (c *Client) OnUpdateChatLastMessage(fn func(client *Client, update *types.UpdateChatLastMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatLastMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatLastMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatPosition The position of a chat in a chat list has changed. An updateChatLastMessage or updateChatDraftMessage update might be sent instead of the update
func (c *Client) OnUpdateChatPosition(fn func(client *Client, update *types.UpdateChatPosition) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatPosition", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatPosition); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatAddedToList A chat was added to a chat list @chat_id Chat identifier @chat_list The chat list to which the chat was added
func (c *Client) OnUpdateChatAddedToList(fn func(client *Client, update *types.UpdateChatAddedToList) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatAddedToList", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatAddedToList); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatRemovedFromList A chat was removed from a chat list @chat_id Chat identifier @chat_list The chat list from which the chat was removed
func (c *Client) OnUpdateChatRemovedFromList(fn func(client *Client, update *types.UpdateChatRemovedFromList) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatRemovedFromList", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatRemovedFromList); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatReadInbox Incoming messages were read or the number of unread messages has been changed @chat_id Chat identifier @last_read_inbox_message_id Identifier of the last read incoming message @unread_count The number of unread messages left in the chat
func (c *Client) OnUpdateChatReadInbox(fn func(client *Client, update *types.UpdateChatReadInbox) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatReadInbox", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatReadInbox); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatReadOutbox Outgoing messages were read @chat_id Chat identifier @last_read_outbox_message_id Identifier of last read outgoing message
func (c *Client) OnUpdateChatReadOutbox(fn func(client *Client, update *types.UpdateChatReadOutbox) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatReadOutbox", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatReadOutbox); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatActionBar The chat action bar was changed @chat_id Chat identifier @action_bar The new value of the action bar; may be null
func (c *Client) OnUpdateChatActionBar(fn func(client *Client, update *types.UpdateChatActionBar) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatActionBar", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatActionBar); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatBusinessBotManageBar The bar for managing business bot was changed in a chat @chat_id Chat identifier @business_bot_manage_bar The new value of the business bot manage bar; may be null
func (c *Client) OnUpdateChatBusinessBotManageBar(fn func(client *Client, update *types.UpdateChatBusinessBotManageBar) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatBusinessBotManageBar", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatBusinessBotManageBar); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatAvailableReactions The chat available reactions were changed @chat_id Chat identifier @available_reactions The new reactions, available in the chat
func (c *Client) OnUpdateChatAvailableReactions(fn func(client *Client, update *types.UpdateChatAvailableReactions) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatAvailableReactions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatAvailableReactions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatDraftMessage A chat draft has changed. Be aware that the update may come in the currently opened chat but with old content of the draft. If the user has changed the content of the draft, this update mustn't be applied
func (c *Client) OnUpdateChatDraftMessage(fn func(client *Client, update *types.UpdateChatDraftMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatDraftMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatDraftMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatEmojiStatus Chat emoji status has changed
func (c *Client) OnUpdateChatEmojiStatus(fn func(client *Client, update *types.UpdateChatEmojiStatus) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatEmojiStatus", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatEmojiStatus); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatMessageSender The message sender that is selected to send messages in a chat has changed @chat_id Chat identifier @message_sender_id New value of message_sender_id; may be null if the user can't change message sender
func (c *Client) OnUpdateChatMessageSender(fn func(client *Client, update *types.UpdateChatMessageSender) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatMessageSender", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatMessageSender); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatMessageAutoDeleteTime The message auto-delete or self-destruct timer setting for a chat was changed @chat_id Chat identifier @message_auto_delete_time New value of message_auto_delete_time
func (c *Client) OnUpdateChatMessageAutoDeleteTime(fn func(client *Client, update *types.UpdateChatMessageAutoDeleteTime) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatMessageAutoDeleteTime", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatMessageAutoDeleteTime); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatNotificationSettings Notification settings for a chat were changed @chat_id Chat identifier @notification_settings The new notification settings
func (c *Client) OnUpdateChatNotificationSettings(fn func(client *Client, update *types.UpdateChatNotificationSettings) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatNotificationSettings", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatNotificationSettings); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatPendingJoinRequests The chat pending join requests were changed @chat_id Chat identifier @pending_join_requests The new data about pending join requests; may be null
func (c *Client) OnUpdateChatPendingJoinRequests(fn func(client *Client, update *types.UpdateChatPendingJoinRequests) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatPendingJoinRequests", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatPendingJoinRequests); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatReplyMarkup The default chat reply markup was changed. Can occur because new messages with reply markup were received or because an old reply markup was hidden by the user
func (c *Client) OnUpdateChatReplyMarkup(fn func(client *Client, update *types.UpdateChatReplyMarkup) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatReplyMarkup", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatReplyMarkup); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatBackground The chat background was changed @chat_id Chat identifier @background The new chat background; may be null if background was reset to default
func (c *Client) OnUpdateChatBackground(fn func(client *Client, update *types.UpdateChatBackground) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatBackground", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatBackground); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatTheme The chat theme was changed @chat_id Chat identifier @theme The new theme of the chat; may be null if theme was reset to default
func (c *Client) OnUpdateChatTheme(fn func(client *Client, update *types.UpdateChatTheme) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatTheme", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatTheme); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatUnreadMentionCount The chat unread_mention_count has changed @chat_id Chat identifier @unread_mention_count The number of unread mention messages left in the chat
func (c *Client) OnUpdateChatUnreadMentionCount(fn func(client *Client, update *types.UpdateChatUnreadMentionCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatUnreadMentionCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatUnreadMentionCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatUnreadReactionCount The chat unread_reaction_count has changed @chat_id Chat identifier @unread_reaction_count The number of messages with unread reactions left in the chat
func (c *Client) OnUpdateChatUnreadReactionCount(fn func(client *Client, update *types.UpdateChatUnreadReactionCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatUnreadReactionCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatUnreadReactionCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatVideoChat A chat video chat state has changed @chat_id Chat identifier @video_chat New value of video_chat
func (c *Client) OnUpdateChatVideoChat(fn func(client *Client, update *types.UpdateChatVideoChat) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatVideoChat", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatVideoChat); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatDefaultDisableNotification The value of the default disable_notification parameter, used when a message is sent to the chat, was changed @chat_id Chat identifier @default_disable_notification The new default_disable_notification value
func (c *Client) OnUpdateChatDefaultDisableNotification(fn func(client *Client, update *types.UpdateChatDefaultDisableNotification) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatDefaultDisableNotification", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatDefaultDisableNotification); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatHasProtectedContent A chat content was allowed or restricted for saving @chat_id Chat identifier @has_protected_content New value of has_protected_content
func (c *Client) OnUpdateChatHasProtectedContent(fn func(client *Client, update *types.UpdateChatHasProtectedContent) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatHasProtectedContent", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatHasProtectedContent); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatIsTranslatable Translation of chat messages was enabled or disabled @chat_id Chat identifier @is_translatable New value of is_translatable
func (c *Client) OnUpdateChatIsTranslatable(fn func(client *Client, update *types.UpdateChatIsTranslatable) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatIsTranslatable", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatIsTranslatable); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatIsMarkedAsUnread A chat was marked as unread or was read @chat_id Chat identifier @is_marked_as_unread New value of is_marked_as_unread
func (c *Client) OnUpdateChatIsMarkedAsUnread(fn func(client *Client, update *types.UpdateChatIsMarkedAsUnread) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatIsMarkedAsUnread", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatIsMarkedAsUnread); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatViewAsTopics A chat default appearance has changed @chat_id Chat identifier @view_as_topics New value of view_as_topics
func (c *Client) OnUpdateChatViewAsTopics(fn func(client *Client, update *types.UpdateChatViewAsTopics) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatViewAsTopics", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatViewAsTopics); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatBlockList A chat was blocked or unblocked @chat_id Chat identifier @block_list Block list to which the chat is added; may be null if none
func (c *Client) OnUpdateChatBlockList(fn func(client *Client, update *types.UpdateChatBlockList) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatBlockList", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatBlockList); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatHasScheduledMessages A chat's has_scheduled_messages field has changed @chat_id Chat identifier @has_scheduled_messages New value of has_scheduled_messages
func (c *Client) OnUpdateChatHasScheduledMessages(fn func(client *Client, update *types.UpdateChatHasScheduledMessages) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatHasScheduledMessages", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatHasScheduledMessages); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatFolders The list of chat folders or a chat folder has changed
func (c *Client) OnUpdateChatFolders(fn func(client *Client, update *types.UpdateChatFolders) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatFolders", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatFolders); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatOnlineMemberCount The number of online group members has changed. This update with non-zero number of online group members is sent only for currently opened chats.
func (c *Client) OnUpdateChatOnlineMemberCount(fn func(client *Client, update *types.UpdateChatOnlineMemberCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatOnlineMemberCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatOnlineMemberCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSavedMessagesTopic Basic information about a Saved Messages topic has changed. This update is guaranteed to come before the topic identifier is returned to the application
func (c *Client) OnUpdateSavedMessagesTopic(fn func(client *Client, update *types.UpdateSavedMessagesTopic) error, filter FilterFunc, position int) {
	c.AddHandler("updateSavedMessagesTopic", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSavedMessagesTopic); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSavedMessagesTopicCount Number of Saved Messages topics has changed @topic_count Approximate total number of Saved Messages topics
func (c *Client) OnUpdateSavedMessagesTopicCount(fn func(client *Client, update *types.UpdateSavedMessagesTopicCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateSavedMessagesTopicCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSavedMessagesTopicCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDirectMessagesChatTopic Basic information about a topic in a channel direct messages chat administered by the current user has changed. This update is guaranteed to come before the topic identifier is returned to the application
func (c *Client) OnUpdateDirectMessagesChatTopic(fn func(client *Client, update *types.UpdateDirectMessagesChatTopic) error, filter FilterFunc, position int) {
	c.AddHandler("updateDirectMessagesChatTopic", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDirectMessagesChatTopic); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateTopicMessageCount Number of messages in a topic has changed; for Saved Messages and channel direct messages chat topics only
func (c *Client) OnUpdateTopicMessageCount(fn func(client *Client, update *types.UpdateTopicMessageCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateTopicMessageCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateTopicMessageCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateQuickReplyShortcut Basic information about a quick reply shortcut has changed. This update is guaranteed to come before the quick shortcut name is returned to the application
func (c *Client) OnUpdateQuickReplyShortcut(fn func(client *Client, update *types.UpdateQuickReplyShortcut) error, filter FilterFunc, position int) {
	c.AddHandler("updateQuickReplyShortcut", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateQuickReplyShortcut); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateQuickReplyShortcutDeleted A quick reply shortcut and all its messages were deleted @shortcut_id The identifier of the deleted shortcut
func (c *Client) OnUpdateQuickReplyShortcutDeleted(fn func(client *Client, update *types.UpdateQuickReplyShortcutDeleted) error, filter FilterFunc, position int) {
	c.AddHandler("updateQuickReplyShortcutDeleted", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateQuickReplyShortcutDeleted); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateQuickReplyShortcuts The list of quick reply shortcuts has changed @shortcut_ids The new list of identifiers of quick reply shortcuts
func (c *Client) OnUpdateQuickReplyShortcuts(fn func(client *Client, update *types.UpdateQuickReplyShortcuts) error, filter FilterFunc, position int) {
	c.AddHandler("updateQuickReplyShortcuts", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateQuickReplyShortcuts); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateQuickReplyShortcutMessages The list of quick reply shortcut messages has changed
func (c *Client) OnUpdateQuickReplyShortcutMessages(fn func(client *Client, update *types.UpdateQuickReplyShortcutMessages) error, filter FilterFunc, position int) {
	c.AddHandler("updateQuickReplyShortcutMessages", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateQuickReplyShortcutMessages); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateForumTopicInfo Basic information about a topic in a forum chat was changed @info New information about the topic
func (c *Client) OnUpdateForumTopicInfo(fn func(client *Client, update *types.UpdateForumTopicInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateForumTopicInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateForumTopicInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateForumTopic Information about a topic in a forum chat was changed
func (c *Client) OnUpdateForumTopic(fn func(client *Client, update *types.UpdateForumTopic) error, filter FilterFunc, position int) {
	c.AddHandler("updateForumTopic", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateForumTopic); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateScopeNotificationSettings Notification settings for some type of chats were updated @scope Types of chats for which notification settings were updated @notification_settings The new notification settings
func (c *Client) OnUpdateScopeNotificationSettings(fn func(client *Client, update *types.UpdateScopeNotificationSettings) error, filter FilterFunc, position int) {
	c.AddHandler("updateScopeNotificationSettings", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateScopeNotificationSettings); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateReactionNotificationSettings Notification settings for reactions were updated @notification_settings The new notification settings
func (c *Client) OnUpdateReactionNotificationSettings(fn func(client *Client, update *types.UpdateReactionNotificationSettings) error, filter FilterFunc, position int) {
	c.AddHandler("updateReactionNotificationSettings", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateReactionNotificationSettings); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNotification A notification was changed @notification_group_id Unique notification group identifier @notification Changed notification
func (c *Client) OnUpdateNotification(fn func(client *Client, update *types.UpdateNotification) error, filter FilterFunc, position int) {
	c.AddHandler("updateNotification", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNotification); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNotificationGroup A list of active notifications in a notification group has changed
func (c *Client) OnUpdateNotificationGroup(fn func(client *Client, update *types.UpdateNotificationGroup) error, filter FilterFunc, position int) {
	c.AddHandler("updateNotificationGroup", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNotificationGroup); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateActiveNotifications Contains active notifications that were shown on previous application launches. This update is sent only if the message database is used. In that case it comes once before any updateNotification and updateNotificationGroup update @groups Lists of active notification groups
func (c *Client) OnUpdateActiveNotifications(fn func(client *Client, update *types.UpdateActiveNotifications) error, filter FilterFunc, position int) {
	c.AddHandler("updateActiveNotifications", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateActiveNotifications); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateHavePendingNotifications Describes whether there are some pending notification updates. Can be used to prevent application from killing, while there are some pending notifications
func (c *Client) OnUpdateHavePendingNotifications(fn func(client *Client, update *types.UpdateHavePendingNotifications) error, filter FilterFunc, position int) {
	c.AddHandler("updateHavePendingNotifications", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateHavePendingNotifications); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDeleteMessages Some messages were deleted
func (c *Client) OnUpdateDeleteMessages(fn func(client *Client, update *types.UpdateDeleteMessages) error, filter FilterFunc, position int) {
	c.AddHandler("updateDeleteMessages", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDeleteMessages); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatAction A message sender activity in the chat has changed
func (c *Client) OnUpdateChatAction(fn func(client *Client, update *types.UpdateChatAction) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatAction", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatAction); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdatePendingTextMessage A new pending text message was received in a chat with a bot. The message must be shown in the chat for at most getOption("pending_text_message_period") seconds,
func (c *Client) OnUpdatePendingTextMessage(fn func(client *Client, update *types.UpdatePendingTextMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updatePendingTextMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdatePendingTextMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUserStatus The user went online or offline @user_id User identifier @status New status of the user
func (c *Client) OnUpdateUserStatus(fn func(client *Client, update *types.UpdateUserStatus) error, filter FilterFunc, position int) {
	c.AddHandler("updateUserStatus", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUserStatus); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUser Some data of a user has changed. This update is guaranteed to come before the user identifier is returned to the application @user New data about the user
func (c *Client) OnUpdateUser(fn func(client *Client, update *types.UpdateUser) error, filter FilterFunc, position int) {
	c.AddHandler("updateUser", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUser); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateBasicGroup Some data of a basic group has changed. This update is guaranteed to come before the basic group identifier is returned to the application @basic_group New data about the group
func (c *Client) OnUpdateBasicGroup(fn func(client *Client, update *types.UpdateBasicGroup) error, filter FilterFunc, position int) {
	c.AddHandler("updateBasicGroup", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateBasicGroup); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSupergroup Some data of a supergroup or a channel has changed. This update is guaranteed to come before the supergroup identifier is returned to the application @supergroup New data about the supergroup
func (c *Client) OnUpdateSupergroup(fn func(client *Client, update *types.UpdateSupergroup) error, filter FilterFunc, position int) {
	c.AddHandler("updateSupergroup", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSupergroup); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSecretChat Some data of a secret chat has changed. This update is guaranteed to come before the secret chat identifier is returned to the application @secret_chat New data about the secret chat
func (c *Client) OnUpdateSecretChat(fn func(client *Client, update *types.UpdateSecretChat) error, filter FilterFunc, position int) {
	c.AddHandler("updateSecretChat", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSecretChat); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUserFullInfo Some data in userFullInfo has been changed @user_id User identifier @user_full_info New full information about the user
func (c *Client) OnUpdateUserFullInfo(fn func(client *Client, update *types.UpdateUserFullInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateUserFullInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUserFullInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateBasicGroupFullInfo Some data in basicGroupFullInfo has been changed @basic_group_id Identifier of a basic group @basic_group_full_info New full information about the group
func (c *Client) OnUpdateBasicGroupFullInfo(fn func(client *Client, update *types.UpdateBasicGroupFullInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateBasicGroupFullInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateBasicGroupFullInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSupergroupFullInfo Some data in supergroupFullInfo has been changed @supergroup_id Identifier of the supergroup or channel @supergroup_full_info New full information about the supergroup
func (c *Client) OnUpdateSupergroupFullInfo(fn func(client *Client, update *types.UpdateSupergroupFullInfo) error, filter FilterFunc, position int) {
	c.AddHandler("updateSupergroupFullInfo", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSupergroupFullInfo); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateServiceNotification A service notification from the server was received. Upon receiving this the application must show a popup with the content of the notification
func (c *Client) OnUpdateServiceNotification(fn func(client *Client, update *types.UpdateServiceNotification) error, filter FilterFunc, position int) {
	c.AddHandler("updateServiceNotification", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateServiceNotification); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFile Information about a file was updated @file New data about the file
func (c *Client) OnUpdateFile(fn func(client *Client, update *types.UpdateFile) error, filter FilterFunc, position int) {
	c.AddHandler("updateFile", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFile); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileGenerationStart The file generation process needs to be started by the application. Use setFileGenerationProgress and finishFileGeneration to generate the file
func (c *Client) OnUpdateFileGenerationStart(fn func(client *Client, update *types.UpdateFileGenerationStart) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileGenerationStart", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileGenerationStart); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileGenerationStop File generation is no longer needed @generation_id Unique identifier for the generation process
func (c *Client) OnUpdateFileGenerationStop(fn func(client *Client, update *types.UpdateFileGenerationStop) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileGenerationStop", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileGenerationStop); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileDownloads The state of the file download list has changed
func (c *Client) OnUpdateFileDownloads(fn func(client *Client, update *types.UpdateFileDownloads) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileDownloads", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileDownloads); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileAddedToDownloads A file was added to the file download list. This update is sent only after file download list is loaded for the first time @file_download The added file download @counts New number of being downloaded and recently downloaded files found
func (c *Client) OnUpdateFileAddedToDownloads(fn func(client *Client, update *types.UpdateFileAddedToDownloads) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileAddedToDownloads", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileAddedToDownloads); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileDownload A file download was changed. This update is sent only after file download list is loaded for the first time
func (c *Client) OnUpdateFileDownload(fn func(client *Client, update *types.UpdateFileDownload) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileDownload", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileDownload); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFileRemovedFromDownloads A file was removed from the file download list. This update is sent only after file download list is loaded for the first time @file_id File identifier @counts New number of being downloaded and recently downloaded files found
func (c *Client) OnUpdateFileRemovedFromDownloads(fn func(client *Client, update *types.UpdateFileRemovedFromDownloads) error, filter FilterFunc, position int) {
	c.AddHandler("updateFileRemovedFromDownloads", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFileRemovedFromDownloads); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateApplicationVerificationRequired A request can't be completed unless application verification is performed; for official mobile applications only.
func (c *Client) OnUpdateApplicationVerificationRequired(fn func(client *Client, update *types.UpdateApplicationVerificationRequired) error, filter FilterFunc, position int) {
	c.AddHandler("updateApplicationVerificationRequired", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateApplicationVerificationRequired); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateApplicationRecaptchaVerificationRequired A request can't be completed unless reCAPTCHA verification is performed; for official mobile applications only.
func (c *Client) OnUpdateApplicationRecaptchaVerificationRequired(fn func(client *Client, update *types.UpdateApplicationRecaptchaVerificationRequired) error, filter FilterFunc, position int) {
	c.AddHandler("updateApplicationRecaptchaVerificationRequired", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateApplicationRecaptchaVerificationRequired); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateCall New call was created or information about a call was updated @call New data about a call
func (c *Client) OnUpdateCall(fn func(client *Client, update *types.UpdateCall) error, filter FilterFunc, position int) {
	c.AddHandler("updateCall", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateCall); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCall Information about a group call was updated @group_call New data about the group call
func (c *Client) OnUpdateGroupCall(fn func(client *Client, update *types.UpdateGroupCall) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCall", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCall); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallParticipant Information about a group call participant was changed. The updates are sent only after the group call is received through getGroupCall and only if the call is joined or being joined
func (c *Client) OnUpdateGroupCallParticipant(fn func(client *Client, update *types.UpdateGroupCallParticipant) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallParticipant", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallParticipant); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallParticipants The list of group call participants that can send and receive encrypted call data has changed; for group calls not bound to a chat only
func (c *Client) OnUpdateGroupCallParticipants(fn func(client *Client, update *types.UpdateGroupCallParticipants) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallParticipants", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallParticipants); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallVerificationState The verification state of an encrypted group call has changed; for group calls not bound to a chat only
func (c *Client) OnUpdateGroupCallVerificationState(fn func(client *Client, update *types.UpdateGroupCallVerificationState) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallVerificationState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallVerificationState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewGroupCallMessage A new message was received in a group call @group_call_id Identifier of the group call @message The message
func (c *Client) OnUpdateNewGroupCallMessage(fn func(client *Client, update *types.UpdateNewGroupCallMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewGroupCallMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewGroupCallMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewGroupCallPaidReaction A new paid reaction was received in a live story group call
func (c *Client) OnUpdateNewGroupCallPaidReaction(fn func(client *Client, update *types.UpdateNewGroupCallPaidReaction) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewGroupCallPaidReaction", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewGroupCallPaidReaction); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallMessageSendFailed A group call message failed to send
func (c *Client) OnUpdateGroupCallMessageSendFailed(fn func(client *Client, update *types.UpdateGroupCallMessageSendFailed) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallMessageSendFailed", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallMessageSendFailed); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallMessagesDeleted Some group call messages were deleted
func (c *Client) OnUpdateGroupCallMessagesDeleted(fn func(client *Client, update *types.UpdateGroupCallMessagesDeleted) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallMessagesDeleted", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallMessagesDeleted); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateLiveStoryTopDonors The list of top donors in live story group call has changed @group_call_id Identifier of the group call @donors New list of live story donors
func (c *Client) OnUpdateLiveStoryTopDonors(fn func(client *Client, update *types.UpdateLiveStoryTopDonors) error, filter FilterFunc, position int) {
	c.AddHandler("updateLiveStoryTopDonors", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateLiveStoryTopDonors); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewCallSignalingData New call signaling data arrived @call_id The call identifier @data The data
func (c *Client) OnUpdateNewCallSignalingData(fn func(client *Client, update *types.UpdateNewCallSignalingData) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewCallSignalingData", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewCallSignalingData); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGiftAuctionState State of a gift auction was updated @state New state of the auction
func (c *Client) OnUpdateGiftAuctionState(fn func(client *Client, update *types.UpdateGiftAuctionState) error, filter FilterFunc, position int) {
	c.AddHandler("updateGiftAuctionState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGiftAuctionState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateActiveGiftAuctions The list of auctions in which participate the current user has changed @states New states of the auctions
func (c *Client) OnUpdateActiveGiftAuctions(fn func(client *Client, update *types.UpdateActiveGiftAuctions) error, filter FilterFunc, position int) {
	c.AddHandler("updateActiveGiftAuctions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateActiveGiftAuctions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUserPrivacySettingRules Some privacy setting rules have been changed @setting The privacy setting @rules New privacy rules
func (c *Client) OnUpdateUserPrivacySettingRules(fn func(client *Client, update *types.UpdateUserPrivacySettingRules) error, filter FilterFunc, position int) {
	c.AddHandler("updateUserPrivacySettingRules", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUserPrivacySettingRules); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUnreadMessageCount Number of unread messages in a chat list has changed. This update is sent only if the message database is used
func (c *Client) OnUpdateUnreadMessageCount(fn func(client *Client, update *types.UpdateUnreadMessageCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateUnreadMessageCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUnreadMessageCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUnreadChatCount Number of unread chats, i.e. with unread messages or marked as unread, has changed. This update is sent only if the message database is used
func (c *Client) OnUpdateUnreadChatCount(fn func(client *Client, update *types.UpdateUnreadChatCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateUnreadChatCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUnreadChatCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStory A story was changed @story The new information about the story
func (c *Client) OnUpdateStory(fn func(client *Client, update *types.UpdateStory) error, filter FilterFunc, position int) {
	c.AddHandler("updateStory", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStory); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStoryDeleted A story became inaccessible @story_poster_chat_id Identifier of the chat that posted the story @story_id Story identifier
func (c *Client) OnUpdateStoryDeleted(fn func(client *Client, update *types.UpdateStoryDeleted) error, filter FilterFunc, position int) {
	c.AddHandler("updateStoryDeleted", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStoryDeleted); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStoryPostSucceeded A story has been successfully posted @story The posted story @old_story_id The previous temporary story identifier
func (c *Client) OnUpdateStoryPostSucceeded(fn func(client *Client, update *types.UpdateStoryPostSucceeded) error, filter FilterFunc, position int) {
	c.AddHandler("updateStoryPostSucceeded", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStoryPostSucceeded); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStoryPostFailed A story failed to post. If the story posting is canceled, then updateStoryDeleted will be received instead of this update
func (c *Client) OnUpdateStoryPostFailed(fn func(client *Client, update *types.UpdateStoryPostFailed) error, filter FilterFunc, position int) {
	c.AddHandler("updateStoryPostFailed", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStoryPostFailed); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatActiveStories The list of active stories posted by a specific chat has changed
func (c *Client) OnUpdateChatActiveStories(fn func(client *Client, update *types.UpdateChatActiveStories) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatActiveStories", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatActiveStories); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStoryListChatCount Number of chats in a story list has changed @story_list The story list @chat_count Approximate total number of chats with active stories in the list
func (c *Client) OnUpdateStoryListChatCount(fn func(client *Client, update *types.UpdateStoryListChatCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateStoryListChatCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStoryListChatCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStoryStealthMode Story stealth mode settings have changed
func (c *Client) OnUpdateStoryStealthMode(fn func(client *Client, update *types.UpdateStoryStealthMode) error, filter FilterFunc, position int) {
	c.AddHandler("updateStoryStealthMode", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStoryStealthMode); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateTrustedMiniAppBots Lists of bots which Mini Apps must be allowed to read text from clipboard and must be opened without a warning
func (c *Client) OnUpdateTrustedMiniAppBots(fn func(client *Client, update *types.UpdateTrustedMiniAppBots) error, filter FilterFunc, position int) {
	c.AddHandler("updateTrustedMiniAppBots", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateTrustedMiniAppBots); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateOption An option changed its value @name The option name @value The new option value
func (c *Client) OnUpdateOption(fn func(client *Client, update *types.UpdateOption) error, filter FilterFunc, position int) {
	c.AddHandler("updateOption", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateOption); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStickerSet A sticker set has changed @sticker_set The sticker set
func (c *Client) OnUpdateStickerSet(fn func(client *Client, update *types.UpdateStickerSet) error, filter FilterFunc, position int) {
	c.AddHandler("updateStickerSet", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStickerSet); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateInstalledStickerSets The list of installed sticker sets was updated @sticker_type Type of the affected stickers @sticker_set_ids The new list of installed ordinary sticker sets
func (c *Client) OnUpdateInstalledStickerSets(fn func(client *Client, update *types.UpdateInstalledStickerSets) error, filter FilterFunc, position int) {
	c.AddHandler("updateInstalledStickerSets", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateInstalledStickerSets); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed @sticker_type Type of the affected stickers @sticker_sets The prefix of the list of trending sticker sets with the newest trending sticker sets
func (c *Client) OnUpdateTrendingStickerSets(fn func(client *Client, update *types.UpdateTrendingStickerSets) error, filter FilterFunc, position int) {
	c.AddHandler("updateTrendingStickerSets", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateTrendingStickerSets); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateRecentStickers The list of recently used stickers was updated @is_attached True, if the list of stickers attached to photo or video files was updated; otherwise, the list of sent stickers is updated @sticker_ids The new list of file identifiers of recently used stickers
func (c *Client) OnUpdateRecentStickers(fn func(client *Client, update *types.UpdateRecentStickers) error, filter FilterFunc, position int) {
	c.AddHandler("updateRecentStickers", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateRecentStickers); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFavoriteStickers The list of favorite stickers was updated @sticker_ids The new list of file identifiers of favorite stickers
func (c *Client) OnUpdateFavoriteStickers(fn func(client *Client, update *types.UpdateFavoriteStickers) error, filter FilterFunc, position int) {
	c.AddHandler("updateFavoriteStickers", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFavoriteStickers); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSavedAnimations The list of saved animations was updated @animation_ids The new list of file identifiers of saved animations
func (c *Client) OnUpdateSavedAnimations(fn func(client *Client, update *types.UpdateSavedAnimations) error, filter FilterFunc, position int) {
	c.AddHandler("updateSavedAnimations", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSavedAnimations); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSavedNotificationSounds The list of saved notification sounds was updated. This update may not be sent until information about a notification sound was requested for the first time @notification_sound_ids The new list of identifiers of saved notification sounds
func (c *Client) OnUpdateSavedNotificationSounds(fn func(client *Client, update *types.UpdateSavedNotificationSounds) error, filter FilterFunc, position int) {
	c.AddHandler("updateSavedNotificationSounds", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSavedNotificationSounds); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDefaultBackground The default background has changed @for_dark_theme True, if default background for dark theme has changed @background The new default background; may be null
func (c *Client) OnUpdateDefaultBackground(fn func(client *Client, update *types.UpdateDefaultBackground) error, filter FilterFunc, position int) {
	c.AddHandler("updateDefaultBackground", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDefaultBackground); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateEmojiChatThemes The list of available emoji chat themes has changed @chat_themes The new list of emoji chat themes
func (c *Client) OnUpdateEmojiChatThemes(fn func(client *Client, update *types.UpdateEmojiChatThemes) error, filter FilterFunc, position int) {
	c.AddHandler("updateEmojiChatThemes", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateEmojiChatThemes); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAccentColors The list of supported accent colors has changed
func (c *Client) OnUpdateAccentColors(fn func(client *Client, update *types.UpdateAccentColors) error, filter FilterFunc, position int) {
	c.AddHandler("updateAccentColors", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAccentColors); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateProfileAccentColors The list of supported accent colors for user profiles has changed
func (c *Client) OnUpdateProfileAccentColors(fn func(client *Client, update *types.UpdateProfileAccentColors) error, filter FilterFunc, position int) {
	c.AddHandler("updateProfileAccentColors", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateProfileAccentColors); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateLanguagePackStrings Some language pack strings have been updated @localization_target Localization target to which the language pack belongs @language_pack_id Identifier of the updated language pack @strings List of changed language pack strings; empty if all strings have changed
func (c *Client) OnUpdateLanguagePackStrings(fn func(client *Client, update *types.UpdateLanguagePackStrings) error, filter FilterFunc, position int) {
	c.AddHandler("updateLanguagePackStrings", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateLanguagePackStrings); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateConnectionState The connection state has changed. This update must be used only to show a human-readable description of the connection state @state The new connection state
func (c *Client) OnUpdateConnectionState(fn func(client *Client, update *types.UpdateConnectionState) error, filter FilterFunc, position int) {
	c.AddHandler("updateConnectionState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateConnectionState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateFreezeState The freeze state of the current user's account has changed
func (c *Client) OnUpdateFreezeState(fn func(client *Client, update *types.UpdateFreezeState) error, filter FilterFunc, position int) {
	c.AddHandler("updateFreezeState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateFreezeState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAgeVerificationParameters The parameters for age verification of the current user's account has changed
func (c *Client) OnUpdateAgeVerificationParameters(fn func(client *Client, update *types.UpdateAgeVerificationParameters) error, filter FilterFunc, position int) {
	c.AddHandler("updateAgeVerificationParameters", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAgeVerificationParameters); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateTermsOfService New terms of service must be accepted by the user. If the terms of service are declined, then the deleteAccount method must be called with the reason "Decline ToS update" @terms_of_service_id Identifier of the terms of service @terms_of_service The new terms of service
func (c *Client) OnUpdateTermsOfService(fn func(client *Client, update *types.UpdateTermsOfService) error, filter FilterFunc, position int) {
	c.AddHandler("updateTermsOfService", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateTermsOfService); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateUnconfirmedSession The first unconfirmed session has changed @session The unconfirmed session; may be null if none
func (c *Client) OnUpdateUnconfirmedSession(fn func(client *Client, update *types.UpdateUnconfirmedSession) error, filter FilterFunc, position int) {
	c.AddHandler("updateUnconfirmedSession", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateUnconfirmedSession); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAttachmentMenuBots The list of bots added to attachment or side menu has changed @bots The new list of bots. The bots must not be shown on scheduled messages screen
func (c *Client) OnUpdateAttachmentMenuBots(fn func(client *Client, update *types.UpdateAttachmentMenuBots) error, filter FilterFunc, position int) {
	c.AddHandler("updateAttachmentMenuBots", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAttachmentMenuBots); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateWebAppMessageSent A message was sent by an opened Web App, so the Web App needs to be closed @web_app_launch_id Identifier of Web App launch
func (c *Client) OnUpdateWebAppMessageSent(fn func(client *Client, update *types.UpdateWebAppMessageSent) error, filter FilterFunc, position int) {
	c.AddHandler("updateWebAppMessageSent", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateWebAppMessageSent); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateActiveEmojiReactions The list of active emoji reactions has changed @emojis The new list of active emoji reactions
func (c *Client) OnUpdateActiveEmojiReactions(fn func(client *Client, update *types.UpdateActiveEmojiReactions) error, filter FilterFunc, position int) {
	c.AddHandler("updateActiveEmojiReactions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateActiveEmojiReactions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAvailableMessageEffects The list of available message effects has changed
func (c *Client) OnUpdateAvailableMessageEffects(fn func(client *Client, update *types.UpdateAvailableMessageEffects) error, filter FilterFunc, position int) {
	c.AddHandler("updateAvailableMessageEffects", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAvailableMessageEffects); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDefaultReactionType The type of default reaction has changed @reaction_type The new type of the default reaction
func (c *Client) OnUpdateDefaultReactionType(fn func(client *Client, update *types.UpdateDefaultReactionType) error, filter FilterFunc, position int) {
	c.AddHandler("updateDefaultReactionType", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDefaultReactionType); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDefaultPaidReactionType The type of default paid reaction has changed @type The new type of the default paid reaction
func (c *Client) OnUpdateDefaultPaidReactionType(fn func(client *Client, update *types.UpdateDefaultPaidReactionType) error, filter FilterFunc, position int) {
	c.AddHandler("updateDefaultPaidReactionType", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDefaultPaidReactionType); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSavedMessagesTags Tags used in Saved Messages or a Saved Messages topic have changed
func (c *Client) OnUpdateSavedMessagesTags(fn func(client *Client, update *types.UpdateSavedMessagesTags) error, filter FilterFunc, position int) {
	c.AddHandler("updateSavedMessagesTags", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSavedMessagesTags); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateActiveLiveLocationMessages The list of messages with active live location that need to be updated by the application has changed. The list is persistent across application restarts only if the message database is used
func (c *Client) OnUpdateActiveLiveLocationMessages(fn func(client *Client, update *types.UpdateActiveLiveLocationMessages) error, filter FilterFunc, position int) {
	c.AddHandler("updateActiveLiveLocationMessages", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateActiveLiveLocationMessages); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateOwnedStarCount The number of Telegram Stars owned by the current user has changed @star_amount The new amount of owned Telegram Stars
func (c *Client) OnUpdateOwnedStarCount(fn func(client *Client, update *types.UpdateOwnedStarCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateOwnedStarCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateOwnedStarCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateOwnedTonCount The number of Toncoins owned by the current user has changed @ton_amount The new amount of owned Toncoins; in the smallest units of the cryptocurrency
func (c *Client) OnUpdateOwnedTonCount(fn func(client *Client, update *types.UpdateOwnedTonCount) error, filter FilterFunc, position int) {
	c.AddHandler("updateOwnedTonCount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateOwnedTonCount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatRevenueAmount The revenue earned from sponsored messages in a chat has changed. If chat revenue screen is opened, then getChatRevenueTransactions may be called to fetch new transactions
func (c *Client) OnUpdateChatRevenueAmount(fn func(client *Client, update *types.UpdateChatRevenueAmount) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatRevenueAmount", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatRevenueAmount); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStarRevenueStatus The Telegram Star revenue earned by a user or a chat has changed. If Telegram Star transaction screen of the chat is opened, then getStarTransactions may be called to fetch new transactions
func (c *Client) OnUpdateStarRevenueStatus(fn func(client *Client, update *types.UpdateStarRevenueStatus) error, filter FilterFunc, position int) {
	c.AddHandler("updateStarRevenueStatus", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStarRevenueStatus); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateTonRevenueStatus The Toncoin revenue earned by the current user has changed. If Toncoin transaction screen of the chat is opened, then getTonTransactions may be called to fetch new transactions
func (c *Client) OnUpdateTonRevenueStatus(fn func(client *Client, update *types.UpdateTonRevenueStatus) error, filter FilterFunc, position int) {
	c.AddHandler("updateTonRevenueStatus", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateTonRevenueStatus); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSpeechRecognitionTrial The parameters of speech recognition without Telegram Premium subscription has changed
func (c *Client) OnUpdateSpeechRecognitionTrial(fn func(client *Client, update *types.UpdateSpeechRecognitionTrial) error, filter FilterFunc, position int) {
	c.AddHandler("updateSpeechRecognitionTrial", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSpeechRecognitionTrial); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateGroupCallMessageLevels The levels of live story group call messages have changed @levels New description of the levels in decreasing order of groupCallMessageLevel.min_star_count
func (c *Client) OnUpdateGroupCallMessageLevels(fn func(client *Client, update *types.UpdateGroupCallMessageLevels) error, filter FilterFunc, position int) {
	c.AddHandler("updateGroupCallMessageLevels", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateGroupCallMessageLevels); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateDiceEmojis The list of supported dice emojis has changed @emojis The new list of supported dice emojis
func (c *Client) OnUpdateDiceEmojis(fn func(client *Client, update *types.UpdateDiceEmojis) error, filter FilterFunc, position int) {
	c.AddHandler("updateDiceEmojis", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateDiceEmojis); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateStakeDiceState The stake dice state has changed @state The new state. The state can be used only if it was received recently enough. Otherwise, a new state must be requested using getStakeDiceState
func (c *Client) OnUpdateStakeDiceState(fn func(client *Client, update *types.UpdateStakeDiceState) error, filter FilterFunc, position int) {
	c.AddHandler("updateStakeDiceState", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateStakeDiceState); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAnimatedEmojiMessageClicked Some animated emoji message was clicked and a big animated sticker must be played if the message is visible on the screen. chatActionWatchingAnimations with the text of the message needs to be sent if the sticker is played
func (c *Client) OnUpdateAnimatedEmojiMessageClicked(fn func(client *Client, update *types.UpdateAnimatedEmojiMessageClicked) error, filter FilterFunc, position int) {
	c.AddHandler("updateAnimatedEmojiMessageClicked", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAnimatedEmojiMessageClicked); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAnimationSearchParameters The parameters of animation search through getOption("animation_search_bot_username") bot has changed @provider Name of the animation search provider @emojis The new list of emojis suggested for searching
func (c *Client) OnUpdateAnimationSearchParameters(fn func(client *Client, update *types.UpdateAnimationSearchParameters) error, filter FilterFunc, position int) {
	c.AddHandler("updateAnimationSearchParameters", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAnimationSearchParameters); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSuggestedActions The list of suggested to the user actions has changed @added_actions Added suggested actions @removed_actions Removed suggested actions
func (c *Client) OnUpdateSuggestedActions(fn func(client *Client, update *types.UpdateSuggestedActions) error, filter FilterFunc, position int) {
	c.AddHandler("updateSuggestedActions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSuggestedActions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateSpeedLimitNotification Download or upload file speed for the user was limited, but it can be restored by subscription to Telegram Premium. The notification can be postponed until a being downloaded or uploaded file is visible to the user.
func (c *Client) OnUpdateSpeedLimitNotification(fn func(client *Client, update *types.UpdateSpeedLimitNotification) error, filter FilterFunc, position int) {
	c.AddHandler("updateSpeedLimitNotification", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateSpeedLimitNotification); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateContactCloseBirthdays The list of contacts that had birthdays recently or will have birthday soon has changed @close_birthday_users List of contact users with close birthday
func (c *Client) OnUpdateContactCloseBirthdays(fn func(client *Client, update *types.UpdateContactCloseBirthdays) error, filter FilterFunc, position int) {
	c.AddHandler("updateContactCloseBirthdays", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateContactCloseBirthdays); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateAutosaveSettings Autosave settings for some type of chats were updated @scope Type of chats for which autosave settings were updated @settings The new autosave settings; may be null if the settings are reset to default
func (c *Client) OnUpdateAutosaveSettings(fn func(client *Client, update *types.UpdateAutosaveSettings) error, filter FilterFunc, position int) {
	c.AddHandler("updateAutosaveSettings", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateAutosaveSettings); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateBusinessConnection A business connection has changed; for bots only @connection New data about the connection
func (c *Client) OnUpdateBusinessConnection(fn func(client *Client, update *types.UpdateBusinessConnection) error, filter FilterFunc, position int) {
	c.AddHandler("updateBusinessConnection", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateBusinessConnection); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewBusinessMessage A new message was added to a business account; for bots only @connection_id Unique identifier of the business connection @message The new message
func (c *Client) OnUpdateNewBusinessMessage(fn func(client *Client, update *types.UpdateNewBusinessMessage) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewBusinessMessage", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewBusinessMessage); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateBusinessMessageEdited A message in a business account was edited; for bots only @connection_id Unique identifier of the business connection @message The edited message
func (c *Client) OnUpdateBusinessMessageEdited(fn func(client *Client, update *types.UpdateBusinessMessageEdited) error, filter FilterFunc, position int) {
	c.AddHandler("updateBusinessMessageEdited", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateBusinessMessageEdited); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateBusinessMessagesDeleted Messages in a business account were deleted; for bots only
func (c *Client) OnUpdateBusinessMessagesDeleted(fn func(client *Client, update *types.UpdateBusinessMessagesDeleted) error, filter FilterFunc, position int) {
	c.AddHandler("updateBusinessMessagesDeleted", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateBusinessMessagesDeleted); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewInlineQuery A new incoming inline query; for bots only
func (c *Client) OnUpdateNewInlineQuery(fn func(client *Client, update *types.UpdateNewInlineQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewInlineQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewInlineQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only
func (c *Client) OnUpdateNewChosenInlineResult(fn func(client *Client, update *types.UpdateNewChosenInlineResult) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewChosenInlineResult", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewChosenInlineResult); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewCallbackQuery A new incoming callback query; for bots only
func (c *Client) OnUpdateNewCallbackQuery(fn func(client *Client, update *types.UpdateNewCallbackQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewCallbackQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewCallbackQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewInlineCallbackQuery A new incoming callback query from a message sent via a bot; for bots only
func (c *Client) OnUpdateNewInlineCallbackQuery(fn func(client *Client, update *types.UpdateNewInlineCallbackQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewInlineCallbackQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewInlineCallbackQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewBusinessCallbackQuery A new incoming callback query from a business message; for bots only
func (c *Client) OnUpdateNewBusinessCallbackQuery(fn func(client *Client, update *types.UpdateNewBusinessCallbackQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewBusinessCallbackQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewBusinessCallbackQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewShippingQuery A new incoming shipping query; for bots only. Only for invoices with flexible price
func (c *Client) OnUpdateNewShippingQuery(fn func(client *Client, update *types.UpdateNewShippingQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewShippingQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewShippingQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout
func (c *Client) OnUpdateNewPreCheckoutQuery(fn func(client *Client, update *types.UpdateNewPreCheckoutQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewPreCheckoutQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewPreCheckoutQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewCustomEvent A new incoming event; for bots only @event A JSON-serialized event
func (c *Client) OnUpdateNewCustomEvent(fn func(client *Client, update *types.UpdateNewCustomEvent) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewCustomEvent", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewCustomEvent); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewCustomQuery A new incoming query; for bots only @id The query identifier @data JSON-serialized query data @timeout Query timeout
func (c *Client) OnUpdateNewCustomQuery(fn func(client *Client, update *types.UpdateNewCustomQuery) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewCustomQuery", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewCustomQuery); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdatePoll A poll was updated; for bots only @poll New data about the poll
func (c *Client) OnUpdatePoll(fn func(client *Client, update *types.UpdatePoll) error, filter FilterFunc, position int) {
	c.AddHandler("updatePoll", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdatePoll); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdatePollAnswer A user changed the answer to a poll; for bots only
func (c *Client) OnUpdatePollAnswer(fn func(client *Client, update *types.UpdatePollAnswer) error, filter FilterFunc, position int) {
	c.AddHandler("updatePollAnswer", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdatePollAnswer); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatMember User rights changed in a chat; for bots only
func (c *Client) OnUpdateChatMember(fn func(client *Client, update *types.UpdateChatMember) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatMember", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatMember); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateNewChatJoinRequest A user sent a join request to a chat; for bots only
func (c *Client) OnUpdateNewChatJoinRequest(fn func(client *Client, update *types.UpdateNewChatJoinRequest) error, filter FilterFunc, position int) {
	c.AddHandler("updateNewChatJoinRequest", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateNewChatJoinRequest); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateChatBoost A chat boost has changed; for bots only
func (c *Client) OnUpdateChatBoost(fn func(client *Client, update *types.UpdateChatBoost) error, filter FilterFunc, position int) {
	c.AddHandler("updateChatBoost", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateChatBoost); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageReaction User changed its reactions on a message with public reactions; for bots only
func (c *Client) OnUpdateMessageReaction(fn func(client *Client, update *types.UpdateMessageReaction) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageReaction", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageReaction); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdateMessageReactions Reactions added to a message with anonymous reactions have changed; for bots only
func (c *Client) OnUpdateMessageReactions(fn func(client *Client, update *types.UpdateMessageReactions) error, filter FilterFunc, position int) {
	c.AddHandler("updateMessageReactions", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdateMessageReactions); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}

// OnUpdatePaidMediaPurchased Paid media were purchased by a user; for bots only
func (c *Client) OnUpdatePaidMediaPurchased(fn func(client *Client, update *types.UpdatePaidMediaPurchased) error, filter FilterFunc, position int) {
	c.AddHandler("updatePaidMediaPurchased", func(cl *Client, u types.TlObject) error {
		if up, ok := u.(*types.UpdatePaidMediaPurchased); ok {
			return fn(cl, up)
		}
		return nil
	}, filter, position)
}
