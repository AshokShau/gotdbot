package gotdbot

import (
	"sync"
	"testing"
)

func TestAcceptResponseRefreshesMeBeforeHandlers(t *testing.T) {
	c := newTestClient(t)
	old := &User{Id: 7, FirstName: "Old", Usernames: &Usernames{ActiveUsernames: []string{"old_bot"}}}
	c.me = old
	updated := &User{Id: 7, FirstName: "New", Usernames: &Usernames{ActiveUsernames: []string{"new_bot"}}}
	command := NewCommandHandler("start", nil)
	seen := make(chan bool, 1)
	c.OnUpdateUser(func(client *Client, u *UpdateUser) error {
		refreshed := client.Me() == u.User
		seen <- refreshed && command.matchesBot(client, "new_bot") && !command.matchesBot(client, "old_bot")
		return nil
	}, nil)
	startTestProcessors(c, 1)
	c.acceptResponse(&UpdateUser{User: updated}, "", nil)
	if !receiveTest(t, seen) {
		t.Fatal("handler did not see refreshed identity and usernames")
	}
	if old.FirstName != "Old" || old.Usernames.ActiveUsernames[0] != "old_bot" {
		t.Fatal("refresh mutated the previous user object")
	}
}

func TestAcceptResponseUserIdentity(t *testing.T) {
	for _, tc := range []struct {
		name     string
		me, user *User
	}{
		{"other user", &User{Id: 7}, &User{Id: 8}},
		{"uninitialized identity", &User{}, &User{Id: 7}},
		{"missing user", &User{Id: 7}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.me = tc.me
			update := &UpdateUser{User: tc.user}
			c.acceptResponse(update, "", nil)
			if c.Me() != tc.me {
				t.Fatal("unrelated update changed the current account")
			}
			if receiveTest(t, c.updates.memory) != update {
				t.Fatal("user update was not queued")
			}
		})
	}
}

func TestAcceptResponseRefreshMeConcurrentReaders(t *testing.T) {
	c := newTestClient(t)
	c.me = &User{Id: 7, Type: &UserTypeBot{}, Usernames: &Usernames{ActiveUsernames: []string{"bot"}}}
	c.sendRequest = func(_ int, data []byte) {
		c.deliverResponse(decodeTestRequest(t, data).Extra, &User{Id: 7})
	}
	managerStop := make(chan struct{})
	close(managerStop)
	command := NewCommandHandler("start", nil)
	var readers sync.WaitGroup
	readers.Go(func() {
		for range 100 {
			if _, err := c.GetMe(); err != nil {
				t.Errorf("request during identity refresh: %v", err)
			}
		}
	})
	readers.Go(func() {
		for range 1000 {
			if id, err := c.currentBotID(); err != nil || id != 7 || !command.matchesBot(c, "bot") {
				t.Error("cached identity unavailable during refresh")
				return
			}
		}
	})
	var latest *User
	for range 1000 {
		latest = &User{Id: 7, Type: &UserTypeBot{}, Usernames: &Usernames{ActiveUsernames: []string{"bot"}}}
		c.acceptResponse(&UpdateUser{User: latest}, "", managerStop)
	}
	readers.Wait()
	if c.Me() != latest {
		t.Fatal("identity was not refreshed when handler delivery stopped")
	}
}
