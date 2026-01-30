package ext

import "github.com/AshokShau/gotdbot"

func extractGeneratedEffectiveFields(u gotdbot.TlObject, c *Context) {
	switch u := u.(type) {
	case *gotdbot.UpdateAnimatedEmojiMessageClicked:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateBusinessMessagesDeleted:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatAccentColors:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatActionBar:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatAddedToList:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatBoost:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatBusinessBotManageBar:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatEmojiStatus:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatHasScheduledMessages:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatMessageAutoDeleteTime:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatMessageSender:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatOnlineMemberCount:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatPendingJoinRequests:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatReadInbox:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatReadOutbox:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatRemovedFromList:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatReplyMarkup:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatRevenueAmount:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatViewAsTopics:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateDeleteMessages:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateForumTopic:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageMentionRead:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageSendAcknowledged:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageSuggestedPostInfo:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageUnreadReactions:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateNewGroupCallPaidReaction:
		c.EffectiveSenderId = u.SenderId
	case *gotdbot.UpdateNotificationGroup:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdatePendingTextMessage:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateTopicMessageCount:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateVideoPublished:
		c.EffectiveChatId = u.ChatId
	}
}
