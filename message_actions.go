package gotdbot

import (
	"fmt"
)

// GetLink returns the message link.
func (m *Message) GetLink(c *Client) (*MessageLink, error) {
	return c.GetMessageLink(m.ChatId, m.Id, nil)
}

// Delete deletes the message.
func (m *Message) Delete(c *Client, opts *DeleteMessagesOpts) error {
	return c.DeleteMessages(m.ChatId, []int64{m.Id}, opts)
}

// Pin pins the message.
func (m *Message) Pin(c *Client, opts *PinChatMessageOpts) error {
	return c.PinChatMessage(m.ChatId, m.Id, opts)
}

// Unpin unpins the message.
func (m *Message) Unpin(c *Client) error {
	return c.UnpinChatMessage(m.ChatId, m.Id)
}

// GetChat returns information about the chat where the message was sent.
func (m *Message) GetChat(c *Client) (*Chat, error) {
	return c.GetChat(m.ChatId)
}

// GetUser returns the user sender, or nil when the sender is a chat or absent.
func (m *Message) GetUser(c *Client) (*User, error) {
	userId := m.SenderID()
	if userId <= 0 {
		return nil, nil
	}
	switch m.SenderId.(type) {
	case *MessageSenderUser, MessageSenderUser:
	default:
		return nil, nil
	}
	return c.GetUser(userId)
}

// LeaveChat leaves the chat where the message was sent.
func (m *Message) LeaveChat(c *Client) error {
	return c.LeaveChat(m.ChatId)
}

// Mention returns an escaped text mention for a user sender, or the title of a
// chat sender. It returns an error if the sender is absent or cannot be fetched.
func (m *Message) Mention(c *Client, parseMode string) (string, error) {
	senderId := m.SenderID()
	if senderId == 0 {
		return "", fmt.Errorf("message sender is absent")
	}
	switch m.SenderId.(type) {
	case *MessageSenderUser, MessageSenderUser:
		user, err := c.GetUser(senderId)
		if err != nil {
			return "", err
		}
		return user.Mention(parseMode), nil
	case *MessageSenderChat, MessageSenderChat:
		chat, err := c.GetChat(senderId)
		if err != nil {
			return "", err
		}
		return chat.Title, nil
	default:
		return "", fmt.Errorf("message sender is absent")
	}
}

// GetMessageProperties returns the message properties.
func (m *Message) GetMessageProperties(c *Client) (*MessageProperties, error) {
	return c.GetMessageProperties(m.ChatId, m.Id)
}

// GetMessageLink returns the message link.
func (m *Message) GetMessageLink(c *Client, opts *GetMessageLinkOpts) (*MessageLink, error) {
	return c.GetMessageLink(m.ChatId, m.Id, opts)
}

// GetRepliedMessage returns the replied message.
func (m *Message) GetRepliedMessage(c *Client) (*Message, error) {
	return c.GetRepliedMessage(m.ChatId, m.Id)
}

// GetChatMember returns member info in the current chat.
func (m *Message) GetChatMember(c *Client) (*ChatMember, error) {
	return c.GetChatMember(m.ChatId, m.SenderId)
}

// SetChatMemberStatus sets chat member status.
func (m *Message) SetChatMemberStatus(c *Client, status ChatMemberStatus) error {
	return c.SetChatMemberStatus(m.ChatId, m.SenderId, status)
}

// Ban bans the message sender.
func (m *Message) Ban(c *Client, opts *BanChatMemberOpts) error {
	return c.BanChatMember(m.ChatId, m.SenderId, opts)
}

// Kick kicks the message sender.
func (m *Message) Kick(c *Client) error {
	return m.SetChatMemberStatus(c, &ChatMemberStatusLeft{})
}

// Restrict restricts the message sender.
func (m *Message) Restrict(c *Client, permissions *ChatPermissions, opts *RestrictChatMemberOpts) error {
	return c.restrictChatMember(m.ChatId, m.SenderId, permissions, opts)
}

// React reacts to the current message.
func (m *Message) React(c *Client, reactionTypes []ReactionType, opts *SetMessageReactionsOpts) error {
	return c.SetMessageReactions(m.ChatId, m.Id, reactionTypes, opts)
}

// ReactEmoji adds an emoji reaction without replacing other chosen reactions.
func (m *Message) ReactEmoji(c *Client, emoji string, opts *AddMessageReactionOpts) error {
	return c.AddMessageReaction(m.ChatId, m.Id, &ReactionTypeEmoji{Emoji: emoji}, opts)
}

// ReactCustomEmoji adds a reaction using a custom emoji identifier.
func (m *Message) ReactCustomEmoji(c *Client, customEmojiId int64, opts *AddMessageReactionOpts) error {
	return c.AddMessageReaction(m.ChatId, m.Id, &ReactionTypeCustomEmoji{CustomEmojiId: customEmojiId}, opts)
}

// Action sends an action in this message's chat. The default action is typing.
func (m *Message) Action(c *Client, opts *SendChatActionOpts) error {
	copyOpts := SendChatActionOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.Action == nil {
		copyOpts.Action = &ChatActionTyping{}
	}
	return c.SendChatAction(m.ChatId, &copyOpts)
}

// Copy copies message to chat.
func (m *Message) Copy(c *Client, chatId int64, opts *CopyMessageOpts) (*Message, error) {
	return c.CopyMessage(chatId, m.ChatId, m.Id, opts)
}

// Forward forwards message to chat.
func (m *Message) Forward(c *Client, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return c.ForwardMessage(chatId, m.ChatId, m.Id, opts)
}
