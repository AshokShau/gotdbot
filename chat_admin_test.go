package gotdbot

import (
	"encoding/json"
	"testing"
)

func TestMemberHelpersPropagateLookupErrors(t *testing.T) {
	for _, message := range []string{"USER_NOT_PARTICIPANT", "CHAT_ADMIN_REQUIRED"} {
		for _, restrict := range []bool{false, true} {
			t.Run(message+"/restrict="+map[bool]string{false: "false", true: "true"}[restrict], func(t *testing.T) {
				c := newTestClient(t)
				failure := &Error{Code: 400, Message: message}
				calls := 0
				c.sendRequest = func(_ int, data []byte) {
					request := decodeTestRequest(t, data)
					calls++
					if request.Type != "getChatMember" {
						t.Fatalf("unexpected request after lookup error: %s", request.Type)
					}
					c.deliverResponse(request.Extra, failure)
				}
				var err error
				if restrict {
					err = c.RestrictChatMember(-42, 7, &ChatPermissions{}, nil)
				} else {
					err = c.UnbanChatMember(-42, 7, &UnbanChatMemberOpts{OnlyIfBanned: true})
				}
				if err != failure || calls != 1 {
					t.Fatalf("error/calls = %v/%d, want original error/1", err, calls)
				}
				assertNoPendingCalls(t, c)
			})
		}
	}
}

func TestRestrictChatMemberPreservesMembership(t *testing.T) {
	for _, tc := range []struct {
		name     string
		status   ChatMemberStatus
		isMember bool
	}{
		{"left", &ChatMemberStatusLeft{}, false},
		{"banned", &ChatMemberStatusBanned{}, false},
		{"member", &ChatMemberStatusMember{}, true},
		{"restricted member", &ChatMemberStatusRestricted{IsMember: true}, true},
		{"restricted nonmember", &ChatMemberStatusRestricted{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			var status *ChatMemberStatusRestricted
			c.sendRequest = func(_ int, data []byte) {
				request := decodeTestRequest(t, data)
				switch request.Type {
				case "getChatMember":
					c.deliverResponse(request.Extra, &ChatMember{Status: tc.status})
				case "setChatMemberStatus":
					var wire struct {
						Status *ChatMemberStatusRestricted `json:"status"`
					}
					if err := json.Unmarshal(data, &wire); err != nil {
						t.Fatal(err)
					}
					status = wire.Status
					c.deliverResponse(request.Extra, &Ok{})
				default:
					t.Fatalf("unexpected request: %s", request.Type)
				}
			}
			if err := c.RestrictChatMember(-42, 7, &ChatPermissions{CanSendBasicMessages: true}, &RestrictChatMemberOpts{RestrictedUntilDate: 123}); err != nil {
				t.Fatal(err)
			}
			if status == nil || status.IsMember != tc.isMember || status.RestrictedUntilDate != 123 || !status.Permissions.CanSendBasicMessages {
				t.Fatalf("wrong restricted status: %+v", status)
			}
		})
	}
}

func TestUnbanChatMemberOnlyIfBanned(t *testing.T) {
	for _, banned := range []bool{false, true} {
		c := newTestClient(t)
		sets := 0
		c.sendRequest = func(_ int, data []byte) {
			request := decodeTestRequest(t, data)
			switch request.Type {
			case "getChatMember":
				var status ChatMemberStatus = &ChatMemberStatusLeft{}
				if banned {
					status = &ChatMemberStatusBanned{}
				}
				c.deliverResponse(request.Extra, &ChatMember{Status: status})
			case "setChatMemberStatus":
				sets++
				c.deliverResponse(request.Extra, &Ok{})
			default:
				t.Fatalf("unexpected request: %s", request.Type)
			}
		}
		if err := c.UnbanChatMember(-42, 7, &UnbanChatMemberOpts{OnlyIfBanned: true}); err != nil {
			t.Fatal(err)
		}
		if (sets == 1) != banned {
			t.Fatalf("banned=%v: status writes=%d", banned, sets)
		}
	}
}
