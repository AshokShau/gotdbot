package gotdbot

import (
	"time"
)

type Context struct {
	// RawUpdate is the original update object.
	RawUpdate TlObject
	// Update contains pointers to all possible update types.
	Update *ContextUpdates
	// Dispatcher is the dispatcher that created this context.
	Dispatcher *Dispatcher

	// Data is a map for storing data across handlers.
	Data map[string]interface{}

	// Helper fields derived from the update
	EffectiveMessage *Message
	EffectiveChatId  int64
}

func NewContext(update TlObject, dispatcher *Dispatcher) *Context {
	ctx := &Context{
		RawUpdate:  update,
		Update:     NewContextUpdates(update),
		Dispatcher: dispatcher,
		Data:       make(map[string]interface{}),
	}
	ctx.extractEffectiveFields(update)
	return ctx
}

// WaitFor waits for an update that matches the filter.
func (c *Context) WaitFor(filter UpdateFilter, timeout time.Duration) (TlObject, error) {
	return c.Dispatcher.WaitFor(filter, timeout)
}

func (c *Context) extractEffectiveFields(u TlObject) {
	switch u := u.(type) {
	// Message updates
	case *UpdateNewMessage:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
		}
	case *UpdateMessageContent:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageEdited:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageSendSucceeded:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
		}
	case *UpdateMessageSendFailed:
		if u.Message != nil {
			c.EffectiveMessage = u.Message
			c.EffectiveChatId = u.Message.ChatId
		}
	case *UpdateMessageContentOpened:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageIsPinned:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageInteractionInfo:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageLiveLocationViewed:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageReaction:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageReactions:
		c.EffectiveChatId = u.ChatId
	case *UpdateMessageFactCheck:
		c.EffectiveChatId = u.ChatId

	// Callbacks and Queries
	case *UpdateNewCallbackQuery:
		c.EffectiveChatId = u.ChatId
	case *UpdateNewInlineCallbackQuery:
		c.EffectiveChatId = u.SenderUserId
	case *UpdateNewBusinessCallbackQuery:
		c.EffectiveChatId = u.SenderUserId
	case *UpdateNewInlineQuery:
		c.EffectiveChatId = u.SenderUserId
	case *UpdateNewPreCheckoutQuery:
		c.EffectiveChatId = u.SenderUserId
	case *UpdateNewShippingQuery:
		c.EffectiveChatId = u.SenderUserId
	case *UpdateNewCustomQuery:
		//  only data

	// Chat updates
	case *UpdateNewChat:
		if u.Chat != nil {
			c.EffectiveChatId = u.Chat.Id
		}
	case *UpdateChatTitle:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatPhoto:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatPermissions:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatLastMessage:
		c.EffectiveChatId = u.ChatId
		if u.LastMessage != nil {
			c.EffectiveMessage = u.LastMessage
		}
	case *UpdateChatPosition:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatIsMarkedAsUnread:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatBlockList:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatDraftMessage:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatNotificationSettings:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatAvailableReactions:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatBackground:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatTheme:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatUnreadMentionCount:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatUnreadReactionCount:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatVideoChat:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatDefaultDisableNotification:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatHasProtectedContent:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatIsTranslatable:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatMember:
		c.EffectiveChatId = u.ChatId
	case *UpdateChatAction:
		c.EffectiveChatId = u.ChatId
	case *UpdateNewChatJoinRequest:
		c.EffectiveChatId = u.ChatId
	}
	extractGeneratedEffectiveFields(u, c)
}
