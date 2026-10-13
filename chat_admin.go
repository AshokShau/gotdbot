package gotdbot

import (
	"fmt"
)

// GetSupergroupId returns the supergroup ID from a chat ID.
func (c *Client) GetSupergroupId(chatId int64) (int64, error) {
	chat, err := c.GetChat(chatId)
	if err != nil {
		return 0, err
	}

	if chat.Type == nil {
		return 0, nil
	}

	if ct, ok := chat.Type.(*ChatTypeSupergroup); ok {
		return ct.SupergroupId, nil
	}

	return 0, nil
}

// ApproveChatJoinRequest approves a user's pending join request.
func (c *Client) ApproveChatJoinRequest(chatId, userId int64) error {
	return c.ProcessChatJoinRequest(chatId, userId, &ProcessChatJoinRequestOpts{Approve: true})
}

// DeclineChatJoinRequest declines a user's pending join request.
func (c *Client) DeclineChatJoinRequest(chatId, userId int64) error {
	return c.ProcessChatJoinRequest(chatId, userId, &ProcessChatJoinRequestOpts{Approve: false})
}

// UnbanChatMemberOpts controls how unbanning affects existing members.
type UnbanChatMemberOpts struct {
	// OnlyIfBanned leaves members who are not banned untouched.
	OnlyIfBanned bool
}

// UnbanChatMember removes a ban. Unless OnlyIfBanned is set, it also removes
// existing members from the chat.
func (c *Client) UnbanChatMember(chatId, userId int64, opts *UnbanChatMemberOpts) error {
	sender := &MessageSenderUser{UserId: userId}
	if opts != nil && opts.OnlyIfBanned {
		member, err := c.GetChatMember(chatId, sender)
		if err != nil {
			return err
		}
		if _, banned := member.Status.(*ChatMemberStatusBanned); !banned {
			return nil
		}
	}
	return c.SetChatMemberStatus(chatId, sender, &ChatMemberStatusLeft{})
}

// RestrictChatMemberOpts controls the duration of a member's restrictions.
type RestrictChatMemberOpts struct {
	// RestrictedUntilDate is a Unix timestamp; zero restricts indefinitely.
	RestrictedUntilDate int32
}

// RestrictChatMember changes permissions while preserving current membership.
func (c *Client) RestrictChatMember(chatId, userId int64, permissions *ChatPermissions, opts *RestrictChatMemberOpts) error {
	return c.restrictChatMember(chatId, &MessageSenderUser{UserId: userId}, permissions, opts)
}

func (c *Client) restrictChatMember(chatId int64, sender MessageSender, permissions *ChatPermissions, opts *RestrictChatMemberOpts) error {
	if permissions == nil {
		return fmt.Errorf("chat permissions required")
	}
	if opts == nil {
		opts = &RestrictChatMemberOpts{}
	}
	member, err := c.GetChatMember(chatId, sender)
	if err != nil {
		return err
	}
	isMember := false
	switch status := member.Status.(type) {
	case *ChatMemberStatusMember, *ChatMemberStatusAdministrator:
		isMember = true
	case *ChatMemberStatusCreator:
		isMember = status.IsMember
	case *ChatMemberStatusRestricted:
		isMember = status.IsMember
	}
	return c.SetChatMemberStatus(chatId, sender, &ChatMemberStatusRestricted{IsMember: isMember, Permissions: permissions, RestrictedUntilDate: opts.RestrictedUntilDate})
}

// PromoteChatMember sets administrator rights. To demote a member, pass nil.
// Custom titles are managed separately with SetChatMemberTag.
func (c *Client) PromoteChatMember(chatId, userId int64, rights *ChatAdministratorRights) error {
	var status ChatMemberStatus = &ChatMemberStatusMember{}
	if rights != nil {
		status = &ChatMemberStatusAdministrator{Rights: rights}
	}
	return c.SetChatMemberStatus(chatId, &MessageSenderUser{UserId: userId}, status)
}

// BanChatSenderChat prevents a channel from sending in a group.
func (c *Client) BanChatSenderChat(chatId, senderChatId int64, opts *BanChatMemberOpts) error {
	return c.BanChatMember(chatId, &MessageSenderChat{ChatId: senderChatId}, opts)
}

// UnbanChatSenderChat removes a sender-chat ban.
func (c *Client) UnbanChatSenderChat(chatId, senderChatId int64) error {
	return c.SetChatMemberStatus(chatId, &MessageSenderChat{ChatId: senderChatId}, &ChatMemberStatusLeft{})
}

// GetChatMemberCount resolves TDLib's chat type before fetching the member count.
func (c *Client) GetChatMemberCount(chatId int64) (int32, error) {
	chat, err := c.GetChat(chatId)
	if err != nil {
		return 0, err
	}
	switch kind := chat.Type.(type) {
	case *ChatTypeBasicGroup:
		group, err := c.GetBasicGroup(kind.BasicGroupId)
		if err != nil {
			return 0, err
		}
		return group.MemberCount, nil
	case *ChatTypeSupergroup:
		group, err := c.GetSupergroupFullInfo(kind.SupergroupId)
		if err != nil {
			return 0, err
		}
		return group.MemberCount, nil
	default:
		return 0, fmt.Errorf("member count requires a group or channel")
	}
}

// ExportChatInviteLink generates a new primary link, revoking the previous one.
func (c *Client) ExportChatInviteLink(chatId int64) (string, error) {
	link, err := c.ReplacePrimaryChatInviteLink(chatId)
	if err != nil {
		return "", err
	}
	return link.InviteLink, nil
}

// GetChatMenuButton gets the menu for a private chat. Pass 0 for the default.
func (c *Client) GetChatMenuButton(chatId int64) (*BotMenuButton, error) {
	if chatId < 0 {
		return nil, fmt.Errorf("menu buttons require a private chat or 0")
	}
	return c.GetMenuButton(chatId)
}

// SetChatMenuButton sets the menu for a private chat. Pass 0 for the default.
func (c *Client) SetChatMenuButton(chatId int64, button *BotMenuButton) error {
	if chatId < 0 || button == nil {
		return fmt.Errorf("private chat or 0 and menu button required")
	}
	return c.SetMenuButton(button, chatId)
}
