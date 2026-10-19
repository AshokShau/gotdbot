package gotdbot

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

type botUpdateState struct {
	mu                sync.Mutex
	chats             map[int64]botChatState
	authorizationDate int64
	serverTime        int64
	serverTimeAt      time.Time
}

type botChatState struct {
	status   ChatMemberStatus
	joinDate int32
}

func (s *botUpdateState) setChat(chatID int64, state botChatState) {
	if s.chats == nil {
		s.chats = make(map[int64]botChatState)
	}
	s.chats[chatID] = state
}

func (c *Client) botTokenUserID() int64 {
	id, _, _ := strings.Cut(c.botToken, ":")
	userID, _ := strconv.ParseInt(id, 10, 64)
	return userID
}

func isBotChatMember(status ChatMemberStatus) bool {
	switch status := status.(type) {
	case *ChatMemberStatusAdministrator, *ChatMemberStatusMember:
		return true
	case *ChatMemberStatusCreator:
		return status != nil && status.IsMember
	case *ChatMemberStatusRestricted:
		return status != nil && status.IsMember
	default:
		return false
	}
}

func (c *Client) acceptBotUpdate(update TlObject, now time.Time) bool {
	s := &c.botUpdates
	s.mu.Lock()
	defer s.mu.Unlock()

	switch u := update.(type) {
	case *UpdateOption:
		if u == nil {
			return true
		}
		switch u.Name {
		case "authorization_date":
			s.authorizationDate = 0
			if value, ok := u.Value.(*OptionValueInteger); ok && value != nil {
				s.authorizationDate = value.Value
			}
		case "unix_time":
			if value, ok := u.Value.(*OptionValueInteger); ok && value != nil {
				s.serverTime, s.serverTimeAt = value.Value, now
			}
		}
	case *UpdateSupergroup:
		if u != nil && u.Supergroup != nil {
			group := u.Supergroup
			s.setChat(supergroupChatIDOffset-group.Id, botChatState{status: group.Status, joinDate: group.Date})
		}
	case *UpdateBasicGroup:
		if u != nil && u.BasicGroup != nil {
			chatID := -u.BasicGroup.Id
			state := s.chats[chatID]
			state.status = u.BasicGroup.Status
			s.setChat(chatID, state)
		}
	case *UpdateChatMember:
		if u == nil || u.NewChatMember == nil {
			return true
		}
		member := u.NewChatMember
		sender, ok := member.MemberId.(*MessageSenderUser)
		if !ok || sender == nil || sender.UserId != c.botTokenUserID() {
			return true
		}
		state := s.chats[u.ChatId]
		if u.Date < state.joinDate {
			return true
		}
		state.status = member.Status
		if u.OldChatMember != nil && !isBotChatMember(u.OldChatMember.Status) && isBotChatMember(member.Status) {
			state.joinDate = u.Date
		}
		s.setChat(u.ChatId, state)
	case *UpdateNewMessage:
		if u == nil || u.Message == nil {
			return true
		}
		message := u.Message

		if message.IsOutgoing || message.SendingState != nil || message.SchedulingState != nil {
			return true
		}
		messageDate := int64(message.Date)
		if message.EditDate != 0 {
			messageDate = int64(message.EditDate)
		}
		unixTime := now.Unix()
		if !s.serverTimeAt.IsZero() {
			unixTime = s.serverTime + int64(now.Sub(s.serverTimeAt)/time.Second)
		}
		if messageDate <= unixTime-86400 {
			return false
		}
		state, known := s.chats[message.ChatId]
		if known && state.status != nil && !isBotChatMember(state.status) {
			removed, ok := message.Content.(*MessageChatDeleteMember)
			return ok && removed != nil && removed.UserId == c.botTokenUserID()
		}
		if known && int64(state.joinDate) > messageDate {
			return false
		}
		if isSupergroupOrChannelID(message.ChatId) && s.authorizationDate > messageDate {
			return false
		}
	}
	return true
}
