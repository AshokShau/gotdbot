package ext

import (
	"github.com/AshokShau/gotdbot"
)

type Context struct {
	// Update contains pointers to all possible update types.
	Update *Updates
	// Client is the gotdbot client.
	Client *gotdbot.Client

	// Data is a map for storing data across handlers.
	Data map[string]interface{}

	// Helper fields derived from the update
	EffectiveMessage  *gotdbot.Message
	EffectiveChatId   int64
	EffectiveSenderId gotdbot.MessageSender
}

func NewContext(client *gotdbot.Client, update gotdbot.TlObject) *Context {
	ctx := &Context{
		Update:    NewUpdates(update),
		Client:    client,
		Data:      make(map[string]interface{}),
	}
	ctx.extractEffectiveFields(update)
	return ctx
}

func (c *Context) extractEffectiveFields(u gotdbot.TlObject) {
	switch u := u.(type) {
	// Message updates
	case *gotdbot.UpdateNewMessage:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
			c.EffectiveSenderId = u.Message.SenderId
		}
	case *gotdbot.UpdateMessageContent:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageEdited:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageSendSucceeded:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
			c.EffectiveSenderId = u.Message.SenderId
		}
	case *gotdbot.UpdateMessageSendFailed:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
			c.EffectiveSenderId = u.Message.SenderId
		}
	case *gotdbot.UpdateMessageContentOpened:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageIsPinned:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageInteractionInfo:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageLiveLocationViewed:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageReaction:
		c.EffectiveChatId = u.ChatId
		c.EffectiveSenderId = u.ActorId
	case *gotdbot.UpdateMessageReactions:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateMessageFactCheck:
		c.EffectiveChatId = u.ChatId

	// Callbacks and Queries
	case *gotdbot.UpdateNewCallbackQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateNewInlineCallbackQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
	case *gotdbot.UpdateNewBusinessCallbackQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
	case *gotdbot.UpdateNewInlineQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
	case *gotdbot.UpdateNewPreCheckoutQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
	case *gotdbot.UpdateNewShippingQuery:
		c.EffectiveSenderId = &gotdbot.MessageSenderUser{
			UserId: u.SenderUserId,
		}
	case *gotdbot.UpdateNewCustomQuery:
		// No sender ID in structure directly, only data

	// Chat updates
	case *gotdbot.UpdateNewChat:
		if u.Chat != nil {
			c.EffectiveChatId = u.Chat.Id
		}
	case *gotdbot.UpdateChatTitle:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatPhoto:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatPermissions:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatLastMessage:
		c.EffectiveChatId = u.ChatId
		if u.LastMessage != nil {
			c.EffectiveMessage = u.LastMessage
			c.EffectiveSenderId = u.LastMessage.SenderId
		}
	case *gotdbot.UpdateChatPosition:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatIsMarkedAsUnread:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatBlockList:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatDraftMessage:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatNotificationSettings:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatAvailableReactions:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatBackground:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatTheme:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatUnreadMentionCount:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatUnreadReactionCount:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatVideoChat:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatDefaultDisableNotification:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatHasProtectedContent:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatIsTranslatable:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatMember:
		c.EffectiveChatId = u.ChatId
	case *gotdbot.UpdateChatAction:
		c.EffectiveChatId = u.ChatId
		c.EffectiveSenderId = u.SenderId
	case *gotdbot.UpdateNewChatJoinRequest:
		c.EffectiveChatId = u.ChatId
	}
	extractGeneratedEffectiveFields(u, c)
}
