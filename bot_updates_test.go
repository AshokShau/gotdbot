package gotdbot

import (
	"reflect"
	"testing"
	"time"
)

func TestBotRejoinFiltersCommandsBeforeWorkers(t *testing.T) {
	for _, basic := range []bool{false, true} {
		name := "supergroup"
		chatID := supergroupChatIDOffset - 42
		if basic {
			name, chatID = "basic group", -42
		}
		t.Run(name, func(t *testing.T) {
			c := newTestClient(t)
			c.botToken = "7:test"
			now := int32(time.Now().Unix())
			membership := func(date int32, status ChatMemberStatus) {
				if basic {
					c.acceptResponse(&UpdateBasicGroup{BasicGroup: &BasicGroup{Id: 42, Status: status}}, "", nil)
				} else {
					c.acceptResponse(&UpdateSupergroup{Supergroup: &Supergroup{Id: 42, Date: date, Status: status}}, "", nil)
				}
			}
			message := func(id int64, date int32, content MessageContent) {
				c.acceptResponse(&UpdateNewMessage{Message: &Message{Id: id, ChatId: chatID, Date: date, Content: content}}, "", nil)
			}
			command := &MessageText{Text: &FormattedText{Text: "/ping"}}
			var commands []int64
			c.OnCommand("ping", func(_ *Client, m *Message) error {
				commands = append(commands, m.Id)
				return nil
			})
			var messages []int64
			c.AddUpdateNewMessageHandlerGroup(func(_ *Client, u *UpdateNewMessage) error {
				messages = append(messages, u.Message.Id)
				return nil
			}, nil, 1)
			done := make(chan struct{}, 1)
			c.OnUpdateOption(func(_ *Client, u *UpdateOption) error {
				if u.Name == "test_barrier" {
					done <- struct{}{}
				}
				return nil
			}, nil)

			// Queue the full sequence before starting even one worker. Correctness
			// must depend on receive order, not application membership handlers.
			membership(now-100, &ChatMemberStatusMember{})
			message(1, now-50, command)
			membership(now-40, &ChatMemberStatusBanned{})
			message(2, now-40, &MessageChatDeleteMember{UserId: 7})
			message(3, now-35, &MessageChatDeleteMember{UserId: 8})
			message(4, now-30, command)
			membership(now-10, &ChatMemberStatusAdministrator{})
			if basic {
				c.acceptResponse(&UpdateChatMember{
					ChatId: chatID, Date: now - 10,
					OldChatMember: &ChatMember{MemberId: &MessageSenderUser{UserId: 7}, Status: &ChatMemberStatusBanned{}},
					NewChatMember: &ChatMember{MemberId: &MessageSenderUser{UserId: 7}, Status: &ChatMemberStatusAdministrator{}},
				}, "", nil)
			}
			message(5, now-30, command) // History delivered after re-adding.
			message(6, now-10, command) // Exact join second must be accepted.
			message(7, now, command)
			c.acceptResponse(&UpdateOption{Name: "test_barrier"}, "", nil)
			startTestProcessors(c, 1)
			receiveTest(t, done)
			if want := []int64{1, 6, 7}; !reflect.DeepEqual(commands, want) {
				t.Errorf("commands = %v, want %v", commands, want)
			}
			if want := []int64{1, 2, 6, 7}; !reflect.DeepEqual(messages, want) {
				t.Errorf("native messages = %v, want %v", messages, want)
			}
		})
	}
}

func TestBotMessageDateFiltering(t *testing.T) {
	now := time.Unix(1800000000, 0)
	chatID := supergroupChatIDOffset - 42
	for _, tc := range []struct {
		name    string
		message Message
		want    bool
	}{
		{"before join", Message{ChatId: chatID, Date: 1799999979}, false},
		{"at join", Message{ChatId: chatID, Date: 1799999980}, true},
		{"other supergroup before authorization", Message{ChatId: chatID - 1, Date: 1799999969}, false},
		{"at authorization", Message{ChatId: chatID - 1, Date: 1799999970}, true},
		{"unknown supergroup", Message{ChatId: chatID - 1, Date: 1799999990}, true},
		{"private before authorization", Message{ChatId: 123, Date: 1799999969}, true},
		{"private exactly one day old", Message{ChatId: 123, Date: 1800000000 - 86400}, false},
		{"private less than one day old", Message{ChatId: 123, Date: 1800000000 - 86399}, true},
		{"recent edit of old message", Message{ChatId: chatID, Date: 1700000000, EditDate: 1799999990}, true},
		{"edit before join", Message{ChatId: chatID, Date: 1700000000, EditDate: 1799999979}, false},
		{"outgoing lifecycle", Message{ChatId: chatID, IsOutgoing: true}, true},
		{"pending send", Message{ChatId: chatID, SendingState: &MessageSendingStatePending{}}, true},
		{"scheduled", Message{ChatId: chatID, SchedulingState: &MessageSchedulingStateSendAtDate{}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.botToken = "7:test"
			c.acceptBotUpdate(&UpdateOption{Name: "authorization_date", Value: &OptionValueInteger{Value: 1799999970}}, now)
			c.acceptBotUpdate(&UpdateSupergroup{Supergroup: &Supergroup{Id: 42, Date: 1799999980, Status: &ChatMemberStatusMember{}}}, now)
			if got := c.acceptBotUpdate(&UpdateNewMessage{Message: &tc.message}, now); got != tc.want {
				t.Errorf("accepted = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestBotFilteringUsesServerTime(t *testing.T) {
	c := newTestClient(t)
	now := time.Unix(1800000000, 0)
	serverTime := now.Unix() - 86400
	c.acceptBotUpdate(&UpdateOption{Name: "unix_time", Value: &OptionValueInteger{Value: serverTime}}, now)
	message := &UpdateNewMessage{Message: &Message{ChatId: 123, Date: int32(serverTime - 86399)}}
	if !c.acceptBotUpdate(message, now) {
		t.Fatal("fresh message rejected because local clock is ahead of server")
	}
	if c.acceptBotUpdate(message, now.Add(time.Second)) {
		t.Fatal("server time did not advance between option updates")
	}
}

func TestBotPromotionDoesNotChangeJoinCutoff(t *testing.T) {
	c := newTestClient(t)
	c.botToken = "7:test"
	now := time.Unix(1800000000, 0)
	chatID := supergroupChatIDOffset - 42
	c.acceptBotUpdate(&UpdateSupergroup{Supergroup: &Supergroup{Id: 42, Date: 1799999900, Status: &ChatMemberStatusMember{}}}, now)
	for _, userID := range []int64{7, 8} {
		c.acceptBotUpdate(&UpdateChatMember{
			ChatId: chatID, Date: 1799999990,
			OldChatMember: &ChatMember{MemberId: &MessageSenderUser{UserId: userID}, Status: &ChatMemberStatusMember{}},
			NewChatMember: &ChatMember{MemberId: &MessageSenderUser{UserId: userID}, Status: &ChatMemberStatusAdministrator{}, JoinedChatDate: 1799999990},
		}, now)
	}
	if !c.acceptBotUpdate(&UpdateNewMessage{Message: &Message{ChatId: chatID, Date: 1799999950}}, now) {
		t.Fatal("promotion incorrectly discarded a command sent after joining")
	}
	// Another member's departure must not mark the bot as absent.
	c.acceptBotUpdate(&UpdateChatMember{ChatId: chatID, NewChatMember: &ChatMember{
		MemberId: &MessageSenderUser{UserId: 8}, Status: &ChatMemberStatusLeft{},
	}}, now)
	if !c.acceptBotUpdate(&UpdateNewMessage{Message: &Message{ChatId: chatID, Date: 1799999995}}, now) {
		t.Fatal("another member's departure suppressed the bot's updates")
	}
	c.acceptBotUpdate(&UpdateChatMember{ChatId: chatID, Date: 1799999800, NewChatMember: &ChatMember{
		MemberId: &MessageSenderUser{UserId: 7}, Status: &ChatMemberStatusBanned{},
	}}, now)
	if !c.acceptBotUpdate(&UpdateNewMessage{Message: &Message{ChatId: chatID, Date: 1799999995}}, now) {
		t.Fatal("historical membership event overwrote the current join")
	}
}

func TestBotFilteringScope(t *testing.T) {
	for _, tc := range []struct {
		name  string
		token string
		raw   bool
		want  int
	}{
		{"bot default", "7:test", false, 0},
		{"bot raw updates", "7:test", true, 1},
		{"user account", "", false, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.botToken, c.config.ReceiveOldMessages = tc.token, tc.raw
			message := &Message{ChatId: 123, Date: int32(time.Now().Unix() - 90000)}
			c.acceptResponse(&UpdateNewMessage{Message: message}, "", nil)
			if got := len(c.updates.memory); got != tc.want {
				t.Fatalf("queued messages = %d, want %d", got, tc.want)
			}
			// Send completion must bypass filtering even when it carries an old message.
			c.acceptResponse(&UpdateMessageSendSucceeded{Message: message, OldMessageId: 9}, "", nil)
			if got := len(c.updates.memory); got != tc.want+1 {
				t.Fatalf("send completion was filtered: queued = %d", got)
			}
		})
	}
}
