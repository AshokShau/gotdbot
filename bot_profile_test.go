package gotdbot

import (
	"encoding/json"
	"testing"
)

func TestBotProfileHelpersUseCachedIdentity(t *testing.T) {
	for _, tc := range []struct {
		name, request string
		response      TlObject
		call          func(*Client) error
	}{
		{"name", "getBotName", &Text{}, func(c *Client) error { _, err := c.GetMyName("en"); return err }},
		{"set name", "setBotName", &Ok{}, func(c *Client) error { return c.SetMyName("name", &SetBotNameOpts{LanguageCode: "en"}) }},
		{"description", "getBotInfoDescription", &Text{}, func(c *Client) error { _, err := c.GetMyDescription("en"); return err }},
		{"set description", "setBotInfoDescription", &Ok{}, func(c *Client) error {
			return c.SetMyDescription("description", &SetBotInfoDescriptionOpts{LanguageCode: "en"})
		}},
		{"short description", "getBotInfoShortDescription", &Text{}, func(c *Client) error { _, err := c.GetMyShortDescription("en"); return err }},
		{"set short description", "setBotInfoShortDescription", &Ok{}, func(c *Client) error {
			return c.SetMyShortDescription("short", &SetBotInfoShortDescriptionOpts{LanguageCode: "en"})
		}},
		{"set photo", "setBotProfilePhoto", &Ok{}, func(c *Client) error { return c.SetMyProfilePhoto(&InputChatPhotoPrevious{ChatPhotoId: 7}) }},
		{"remove photo", "setBotProfilePhoto", &Ok{}, func(c *Client) error { return c.RemoveMyProfilePhoto() }},
		{"group rights", "getUserFullInfo", &UserFullInfo{BotInfo: &BotInfo{}}, func(c *Client) error { _, err := c.GetMyDefaultAdministratorRights(false); return err }},
		{"channel rights", "getUserFullInfo", &UserFullInfo{BotInfo: &BotInfo{}}, func(c *Client) error { _, err := c.GetMyDefaultAdministratorRights(true); return err }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.me = &User{Id: 12345, Type: &UserTypeBot{}}
			calls := 0
			c.sendRequest = func(_ int, data []byte) {
				calls++
				var req struct {
					Type         string `json:"@type"`
					Extra        string `json:"@extra"`
					BotUserID    int64  `json:"bot_user_id"`
					UserID       int64  `json:"user_id"`
					LanguageCode string `json:"language_code"`
				}
				if err := json.Unmarshal(data, &req); err != nil {
					t.Fatal(err)
				}
				if req.Type != tc.request || (req.BotUserID != c.Me().Id && req.UserID != c.Me().Id) {
					t.Errorf("request = %s; want %s using cached user %d", data, tc.request, c.Me().Id)
				}
				if tc.request == "getBotName" || tc.request == "setBotName" || tc.request == "getBotInfoDescription" || tc.request == "setBotInfoDescription" || tc.request == "getBotInfoShortDescription" || tc.request == "setBotInfoShortDescription" {
					if req.LanguageCode != "en" {
						t.Fatalf("language lost: %s", data)
					}
				}
				c.deliverResponse(req.Extra, tc.response)
			}
			if err := tc.call(c); err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Fatalf("sent %d requests; want only the profile request", calls)
			}
		})
	}
}

func TestBotProfileHelpersRejectUnavailableOrNonBotIdentity(t *testing.T) {
	for _, tc := range []struct {
		name string
		me   *User
	}{
		{"unavailable", nil},
		{"user account", &User{Id: 42, Type: &UserTypeRegular{}}},
		{"missing type", &User{Id: 42}},
		{"nil bot type", &User{Id: 42, Type: (*UserTypeBot)(nil)}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.me = tc.me
			c.sendRequest = func(_ int, data []byte) {
				t.Errorf("invalid identity sent a request: %s", data)
				c.deliverResponse(decodeTestRequest(t, data).Extra, &Error{Code: 400, Message: "unexpected request"})
			}
			if _, err := c.GetMyName("en"); err == nil {
				t.Fatal("invalid cached identity was accepted")
			}
		})
	}
}

func TestAuthorizationPopulatesIdentityCache(t *testing.T) {
	c := newTestClient(t)
	me := &User{Id: 42, Type: &UserTypeBot{}}
	calls := 0
	c.sendRequest = func(_ int, data []byte) {
		calls++
		req := decodeTestRequest(t, data)
		if req.Type != "getMe" {
			t.Errorf("authorization sent %s, want getMe", req.Type)
		}
		c.deliverResponse(req.Extra, me)
	}
	c.authHandler(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateReady{}})
	if c.Me() != me || !c.isAuthorized.Load() || calls != 1 {
		t.Fatalf("authorization did not cache identity: Me=%v, authorized=%v, calls=%d", c.Me(), c.isAuthorized.Load(), calls)
	}
	if err := c.authErr; err != nil {
		t.Fatal(err)
	}
}
