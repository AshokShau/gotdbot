package gotdbot

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type testHandler struct {
	check  func(*Client, TlObject) bool
	handle func(*Client, TlObject) error
}

func (h testHandler) CheckUpdate(c *Client, u TlObject) bool { return h.check == nil || h.check(c, u) }

func (h testHandler) HandleUpdate(c *Client, u TlObject) error {
	if h.handle == nil {
		return nil
	}
	return h.handle(c, u)
}

func TestProcessorGroupControl(t *testing.T) {
	for _, tc := range []struct {
		name   string
		action error
		want   []string
	}{
		{"default", nil, []string{"first", "next"}},
		{"continue handlers", ContinueHandlers, []string{"first", "second", "next"}},
		{"continue groups", ContinueGroups, []string{"first", "next"}},
		{"end groups", EndGroups, []string{"first"}},
		{"wrapped continue", fmt.Errorf("wrapped: %w", ContinueHandlers), []string{"first", "second", "next"}},
		{"wrapped end", fmt.Errorf("wrapped: %w", EndGroups), []string{"first"}},
		{"ordinary error", errors.New("failed"), []string{"first", "next"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			var calls []string
			barrier := make(chan struct{}, 1)
			c.AddHandlerGroup(testHandler{
				check:  func(_ *Client, u TlObject) bool { return u.(*UpdateNewMessage).Message.Id == 2 },
				handle: func(*Client, TlObject) error { barrier <- struct{}{}; return EndGroups },
			}, -100)
			c.AddHandlerGroup(testHandler{handle: func(*Client, TlObject) error { calls = append(calls, "next"); return nil }}, 5)
			c.AddHandlerGroup(testHandler{check: func(*Client, TlObject) bool { return false }, handle: func(*Client, TlObject) error { t.Error("non-matching handler ran"); return nil }}, -2)
			c.AddHandlerGroup(testHandler{handle: func(*Client, TlObject) error { calls = append(calls, "first"); return tc.action }}, -2)
			c.AddHandlerGroup(testHandler{handle: func(*Client, TlObject) error { calls = append(calls, "second"); return nil }}, -2)
			startTestProcessors(c, 1)
			if err := c.updates.push(queueTestUpdate(1)); err != nil {
				t.Fatal(err)
			}
			if err := c.updates.push(queueTestUpdate(2)); err != nil {
				t.Fatal(err)
			}
			receiveTest(t, barrier)
			if !reflect.DeepEqual(calls, tc.want) {
				t.Fatalf("calls = %v, want %v", calls, tc.want)
			}
		})
	}
}

func TestProcessorErrorAndPanicRecovery(t *testing.T) {
	c := newTestClient(t)
	boom := errors.New("handler failed")
	errorsSeen := make(chan error, 1)
	panicsSeen := make(chan any, 1)
	processed := make(chan int64, 2)
	c.errorHandler = func(_ *Client, _ TlObject, err error) error { errorsSeen <- err; return ContinueHandlers }
	c.panicHandler = func(_ *Client, _ TlObject, value any) { panicsSeen <- value }
	c.AddHandler(testHandler{handle: func(_ *Client, u TlObject) error {
		switch u.(*UpdateNewMessage).Message.Id {
		case 1:
			return boom
		case 2:
			panic("handler panic")
		}
		return ContinueHandlers
	}})
	c.AddHandler(testHandler{handle: func(_ *Client, u TlObject) error { processed <- u.(*UpdateNewMessage).Message.Id; return nil }})
	startTestProcessors(c, 1)
	for i := int64(1); i <= 3; i++ {
		if err := c.updates.push(queueTestUpdate(i)); err != nil {
			t.Fatal(err)
		}
	}
	if err := receiveTest(t, errorsSeen); err != boom {
		t.Fatal(err)
	}
	if p := receiveTest(t, panicsSeen); p != "handler panic" {
		t.Fatalf("panic = %v", p)
	}
	if id := receiveTest(t, processed); id != 1 {
		t.Fatalf("first ID = %d", id)
	}
	if id := receiveTest(t, processed); id != 3 {
		t.Fatalf("processor did not recover: %d", id)
	}
}

func TestConcurrentHandlerRegistrationPreservesSnapshots(t *testing.T) {
	c := newTestClient(t)
	c.AddHandler(nil)
	if c.handlers.Load() != nil {
		t.Fatal("nil registration changed handlers")
	}
	c.AddHandlerGroup(testHandler{}, 0)
	old := c.handlers.Load()
	const count = 64
	var wg sync.WaitGroup
	for i := range count {
		wg.Go(func() {
			c.AddHandlerGroup(testHandler{}, i%4-2)
			data := c.handlers.Load()
			for j := 1; j < len(data.groups); j++ {
				if data.groups[j-1] >= data.groups[j] {
					t.Error("groups not sorted")
				}
			}
			for _, group := range data.groups {
				if len(data.handlers[group]) == 0 {
					t.Error("published empty handler group")
				}
			}
		})
	}
	wg.Wait()
	if len(old.groups) != 1 || len(old.handlers[0]) != 1 {
		t.Fatal("registration mutated old snapshot")
	}
	data := c.handlers.Load()
	if !reflect.DeepEqual(data.groups, []int{-2, -1, 0, 1}) {
		t.Fatalf("groups = %v", data.groups)
	}
	total := 0
	for _, handlers := range data.handlers {
		total += len(handlers)
	}
	if total != count+1 {
		t.Fatalf("lost registrations: %d", total)
	}
}

func TestCommandMatching(t *testing.T) {
	c := newTestClient(t)
	c.me = &User{Usernames: &Usernames{ActiveUsernames: []string{"TestBot", "AliasBot"}}}
	h := NewCommandHandler(" START ", nil)
	for _, tc := range []struct {
		text string
		want bool
	}{
		{"/start", true}, {" /START args \n", true}, {"/start@TestBot", true}, {"/start@aliasBOT a", true},
		{"/start@otherbot", false}, {"/starting", false}, {"start", false}, {"/", false}, {"", false}, {"/help", false},
	} {
		t.Run(tc.text, func(t *testing.T) {
			if got := h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, tc.text)}); got != tc.want {
				t.Fatalf("command %q = %v", tc.text, got)
			}
		})
	}
	h.SetTriggers("!λ")
	if !h.matchCommand(c, "λstart hello") || !h.matchCommand(c, "!start") {
		t.Fatal("custom Unicode trigger failed")
	}
	if h.matchesBot(nil, "TestBot") || h.matchesBot(&Client{}, "TestBot") {
		t.Fatal("matched bot without identity")
	}
	var filters []int
	h.AddFilter(nil).AddFilter(func(*Message) bool { filters = append(filters, 1); return false }).AddFilter(func(*Message) bool { filters = append(filters, 2); return true })
	if h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, "!start")}) || !reflect.DeepEqual(filters, []int{1}) {
		t.Fatal("filters did not short-circuit")
	}
}

func TestCommandAliases(t *testing.T) {
	c := newTestClient(t)
	c.me = &User{Usernames: &Usernames{ActiveUsernames: []string{"TestBot", "AliasBot"}}}
	var calledWith *Message
	h := c.OnCommandGroup("start", func(_ *Client, msg *Message) error {
		calledWith = msg
		return EndGroups
	}, 7)
	if h.AddAliases(" HELP ", "info", "", " \t ").AddAliases("about") != h {
		t.Fatal("AddAliases did not return the handler")
	}
	if h.Command != "start" || !reflect.DeepEqual(h.Aliases, []string{"help", "info", "about"}) {
		t.Fatalf("command = %q, aliases = %v", h.Command, h.Aliases)
	}
	registered := c.handlers.Load().handlers[7]
	if len(registered) != 1 || registered[0] != h {
		t.Fatal("aliases should share one registered handler")
	}
	h.SetTriggers("/!λ")
	for _, tc := range []struct {
		text string
		want bool
	}{
		{"/start", true}, {"/help", true}, {" /HELP args \n", true},
		{"/info", true}, {"/about", true}, {"!help", true}, {"λinfo args", true},
		{"/help@TestBot", true}, {"/INFO@aliasBOT args", true},
		{"/help@otherbot", false}, {"/help@", false}, {"/info@TestBot@otherbot", false},
		{"/helper", false}, {"/unknown", false}, {"help", false},
		{".help", false}, {"/", false}, {"/@TestBot", false}, {"", false},
	} {
		t.Run(tc.text, func(t *testing.T) {
			msg := testMessage(42, tc.text)
			update := &UpdateNewMessage{Message: msg}
			if got := h.CheckUpdate(c, update); got != tc.want {
				t.Fatalf("command %q = %v, want %v", tc.text, got, tc.want)
			}
			if tc.want {
				calledWith = nil
				if err := h.HandleUpdate(c, update); err != EndGroups || calledWith != msg {
					t.Fatal("alias did not call the shared function with the original message")
				}
			}
		})
	}
	if h.matchCommand(nil, "/help@TestBot") {
		t.Fatal("alias matched a bot mention without a client")
	}
	h.SetTriggers("")
	if h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, "/help")}) {
		t.Fatal("empty triggers should disable aliases")
	}
	h.SetTriggers("/")
	options := DefaultMatchingOptions()
	options.AllowPrivate = false
	h.SetOptions(options)
	if h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, "/help")}) {
		t.Fatal("alias bypassed matching options")
	}
	h.SetOptions(DefaultMatchingOptions()).AddFilter(func(*Message) bool { return false })
	if h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, "/info")}) {
		t.Fatal("alias bypassed the filter")
	}
}

func TestMatchingOptions(t *testing.T) {
	for _, tc := range []struct {
		name    string
		msg     *Message
		disable func(*MatchingOptions)
	}{
		{"private", &Message{ChatId: 42}, func(o *MatchingOptions) { o.AllowPrivate = false }},
		{"group", &Message{ChatId: -42}, func(o *MatchingOptions) { o.AllowGroup = false }},
		{"supergroup", &Message{ChatId: supergroupChatIDOffset - 42}, func(o *MatchingOptions) { o.AllowSupergroup = false }},
		{"channel", &Message{ChatId: supergroupChatIDOffset - 42, IsChannelPost: true}, func(o *MatchingOptions) { o.AllowChannel = false }},
		{"secret", &Message{ChatId: secretChatIDOffset + 42}, func(o *MatchingOptions) { o.AllowSecret = false }},
		{"edited", &Message{ChatId: 42, EditDate: 1}, func(o *MatchingOptions) { o.AllowEdited = false }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			o := DefaultMatchingOptions()
			if !o.check(tc.msg) {
				t.Fatal("default rejected message")
			}
			tc.disable(&o)
			if o.check(tc.msg) {
				t.Fatal("disabled option accepted message")
			}
		})
	}
	o := DefaultMatchingOptions()
	if o.check(nil) || o.check(&Message{}) || o.check(&Message{ChatId: 42, IsOutgoing: true}) {
		t.Fatal("default accepted nil/unknown/outgoing")
	}
	o.AllowOutgoing = true
	if !o.check(&Message{ChatId: 42, IsOutgoing: true}) {
		t.Fatal("outgoing opt-in ignored")
	}
}

func TestUpdateMessageAndChatExtraction(t *testing.T) {
	m := testMessage(42, "hello")
	for _, update := range []TlObject{
		&UpdateNewMessage{Message: m},
		&UpdateNewBusinessMessage{Message: &BusinessMessage{Message: m}},
		&UpdateBusinessMessageEdited{Message: &BusinessMessage{Message: m}},
	} {
		if got, ok := getMessageFromUpdate(update); got != m || !ok {
			t.Fatalf("extract %T = (%v, %v)", update, got, ok)
		}
		if id := new(Client).UpdateChatID(update); id != 42 {
			t.Fatalf("chat %T = %d", update, id)
		}
	}
	for _, update := range []TlObject{nil, (*UpdateNewMessage)(nil), &UpdateNewMessage{}, &UpdateNewBusinessMessage{}, &UpdateBusinessMessageEdited{}, &Ok{}} {
		if msg, ok := getMessageFromUpdate(update); msg != nil || ok {
			t.Fatalf("invalid extraction %T", update)
		}
	}
	for _, update := range []TlObject{
		&UpdateNewCallbackQuery{ChatId: 42}, &UpdateMessageEdited{ChatId: 42}, &UpdateMessageContent{ChatId: 42},
		&UpdateMessageEphemeralContent{ChatId: 42}, &UpdateMessageLiveLocationViewed{ChatId: 42}, &UpdateMessageInteractionInfo{ChatId: 42},
		&UpdateMessageSendSucceeded{Message: m}, &UpdateMessageSendFailed{Message: m},
	} {
		if id := new(Client).UpdateChatID(update); id != 42 {
			t.Fatalf("chat %T = %d", update, id)
		}
	}
	called := false
	h := NewCommandHandler("start", func(_ *Client, msg *Message) error { called = msg == m; return EndGroups })
	if err := h.HandleUpdate(nil, &UpdateNewMessage{Message: m}); err != EndGroups || !called {
		t.Fatal("handler did not receive message")
	}
	if err := h.HandleUpdate(nil, &Ok{}); err != nil {
		t.Fatal(err)
	}
}

func TestCommandTriggerSetterMatchesOnePrefix(t *testing.T) {
	c := newTestClient(t)
	c.me = &User{Id: 7, Usernames: &Usernames{ActiveUsernames: []string{"TestBot"}}}
	h := NewCommandHandler("start", func(*Client, *Message) error { return nil })
	if h.SetTriggers("/!.λ") != h {
		t.Fatal("trigger setter did not return the handler")
	}
	for _, text := range []string{"/start", "!start", ".start", "λstart", "!START@TestBot args"} {
		if !h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, text)}) {
			t.Errorf("did not match %q", text)
		}
	}
	for _, text := range []string{"/!start", "!.start", "start", "!start@", "!start@other", "!start@TestBot@extra", "!"} {
		if h.CheckUpdate(c, &UpdateNewMessage{Message: testMessage(42, text)}) {
			t.Errorf("incorrectly matched %q", text)
		}
	}
	previous := h.Triggers
	h.SetTriggers("#")
	if string(previous) != "/!.λ" || !h.matchCommand(c, "#start") || h.matchCommand(c, "/start") {
		t.Fatal("trigger replacement mutated the old slice or retained old triggers")
	}
	h.SetTriggers("")
	if h.matchCommand(c, "#start") {
		t.Fatal("empty triggers did not disable matching")
	}
}

func TestGeneralUpdateHandlerGroupControl(t *testing.T) {
	c := newTestClient(t)
	seen := make(chan string, 2)
	c.OnUpdateGroup(func(_ *Client, u Update) error {
		seen <- "general:" + u.GetType()
		return ContinueHandlers
	}, func(u Update) bool { return new(Client).UpdateChatID(u) == 42 }, 0)
	c.OnUpdateDeleteMessages(func(_ *Client, u *UpdateDeleteMessages) error {
		seen <- "typed:" + u.GetType()
		return nil
	}, nil)
	startTestProcessors(c, 1)
	if err := c.updates.push(&UpdateDeleteMessages{ChatId: 42, MessageIds: []int64{1}}); err != nil {
		t.Fatal(err)
	}
	if got := receiveTest(t, seen); got != "general:updateDeleteMessages" {
		t.Fatal(got)
	}
	if got := receiveTest(t, seen); got != "typed:updateDeleteMessages" {
		t.Fatal(got)
	}
}

func TestGeneralUpdateHandlerRejectsNonUpdatesAndTypedNil(t *testing.T) {
	checks, calls := 0, 0
	h := NewUpdateHandler(func(*Client, Update) error { calls++; return nil }, func(Update) bool { checks++; return true })
	for _, object := range []TlObject{nil, (*UpdateNewMessage)(nil), &GetMe{}, &Message{}} {
		if h.CheckUpdate(nil, object) {
			t.Fatalf("accepted %T", object)
		}
		if err := h.HandleUpdate(nil, object); err != nil {
			t.Fatal(err)
		}
	}
	if checks != 0 || calls != 0 {
		t.Fatal("invoked callback for a non-update")
	}
	h.AddFilter(func(Update) bool { return false })
	if h.CheckUpdate(nil, &UpdateNewMessage{}) {
		t.Fatal("additional filter ignored")
	}
	c := newTestClient(t)
	c.OnUpdate(nil, nil)
	if c.handlers.Load() != nil {
		t.Fatal("nil response registered a handler")
	}
}

func TestUpdateHelpersRouteChatIDs(t *testing.T) {
	msg := &Message{ChatId: 42}
	for _, tc := range []struct {
		object TlObject
		id     int64
	}{
		{nil, 0}, {(*UpdateMessageEdited)(nil), 0}, {&GetMessage{ChatId: 42}, 0},
		{&UpdateNewMessage{Message: msg}, 42}, {&UpdateNewBusinessMessage{Message: &BusinessMessage{Message: msg}}, 42},
		{&UpdateBusinessMessageEdited{Message: &BusinessMessage{Message: msg}}, 42},
		{&UpdateDeleteMessages{ChatId: 42}, 42}, {&UpdateChatTitle{ChatId: 42}, 42},
		{&UpdateMessageSendSucceeded{Message: msg}, 42}, {&UpdateMessageSendFailed{Message: msg}, 42},
		{&UpdateMessageSendFailed{}, 0}, {&UpdateConnectionState{}, 0},
		{UpdateMessageEdited{ChatId: 42}, 42},
	} {
		if id := new(Client).UpdateChatID(tc.object); id != tc.id {
			t.Errorf("%T chat = %d, want %d", tc.object, id, tc.id)
		}
	}
	if got, ok := messageFromUpdate(&UpdateNewMessage{Message: msg}); !ok || got != msg {
		t.Fatal("message extraction lost identity")
	}
	if got, ok := messageFromUpdate(&UpdateMessageEdited{ChatId: 42}); ok || got != nil {
		t.Fatal("ID-only update produced a message")
	}
}

func TestMessageHandlerAdditionalFiltersAndCommandOptions(t *testing.T) {
	var calls []int
	h := NewMessageHandler(func(*Client, *Message) error { return nil }, func(*Message) bool { calls = append(calls, 1); return false })
	h.AddFilter(func(*Message) bool { calls = append(calls, 2); return true })
	if h.CheckUpdate(nil, &UpdateNewMessage{Message: testMessage(42, "hello")}) || !reflect.DeepEqual(calls, []int{1}) {
		t.Fatal("message filters did not short-circuit")
	}
	opts := DefaultMatchingOptions()
	opts.AllowPrivate = false
	command := NewCommandHandler("start", func(*Client, *Message) error { return nil }).SetOptions(opts)
	if command.CheckUpdate(nil, &UpdateNewMessage{Message: testMessage(42, "/start")}) {
		t.Fatal("command ignored configured matching options")
	}
}
