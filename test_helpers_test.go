package gotdbot

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/AshokShau/gotdbot/logger"
)

// Build a client at the transport boundary: tests require neither a shared
// TDLib library nor network access, and exercise the real request/worker code.
func newTestClient(t testing.TB) *Client {
	t.Helper()
	c := &Client{
		clientID: 7, config: DefaultClientConfig(), Logger: logger.Disabled(),
		updates: newUpdateQueue(8), stop: make(chan struct{}), closed: make(chan struct{}),
		authDone: make(chan struct{}), pendingRequests: make(map[string]*pendingCall),
		pendingMessages: make(map[[2]int64]chan TlObject),
		me:              &User{},
	}
	c.config.RequestTimeout = 3 * time.Second
	c.config.AuthorizationTimeout = 3 * time.Second
	t.Cleanup(func() {
		c.stopOnce.Do(func() { close(c.stop) })
		if err := c.updates.close(); err != nil {
			t.Errorf("queue cleanup: %v", err)
		}
		done := make(chan struct{})
		go func() { c.wg.Wait(); close(done) }()
		receiveTest(t, done)
	})
	return c
}

func receiveTest[T any](t testing.TB, ch <-chan T) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for test goroutine")
		var zero T
		return zero
	}
}

func eventuallyTest(t *testing.T, condition func() bool) {
	t.Helper()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	tick := time.NewTicker(time.Millisecond)
	defer tick.Stop()
	for !condition() {
		select {
		case <-timer.C:
			t.Fatal("test condition never became true")
		case <-tick.C:
		}
	}
}

type testWireRequest struct {
	Type      string `json:"@type"`
	Extra     string `json:"@extra"`
	ChatID    int64  `json:"chat_id"`
	MessageID int64  `json:"message_id"`
}

func decodeTestRequest(t *testing.T, data []byte) testWireRequest {
	t.Helper()
	var request testWireRequest
	if err := json.Unmarshal(data, &request); err != nil {
		t.Errorf("invalid wire request %q: %v", data, err)
	}
	return request
}

func assertNoPendingCalls(t *testing.T, c *Client) {
	t.Helper()
	c.requestMu.Lock()
	defer c.requestMu.Unlock()
	if len(c.pendingRequests) != 0 || len(c.pendingMessages) != 0 {
		t.Errorf("leaked requests/messages: %d/%d", len(c.pendingRequests), len(c.pendingMessages))
	}
}

func startTestProcessors(c *Client, count int) {
	c.wg.Add(count)
	for range count {
		go c.processor()
	}
}

func testMessage(chatID int64, text string) *Message {
	return &Message{ChatId: chatID, Content: &MessageText{Text: &FormattedText{Text: text}}}
}

type helperWireRequest struct {
	Type                        string              `json:"@type"`
	Extra                       string              `json:"@extra"`
	ChatId                      int64               `json:"chat_id"`
	MessageId                   int64               `json:"message_id"`
	InlineMessageId             string              `json:"inline_message_id"`
	BusinessConnectionId        string              `json:"business_connection_id"`
	Location                    *LiveLocation       `json:"location"`
	Checklist                   *InputChecklist     `json:"checklist"`
	EphemeralMessageId          int32               `json:"ephemeral_message_id"`
	ReceiverUserId              int64               `json:"receiver_user_id"`
	CallbackQueryId             int64               `json:"callback_query_id,string"`
	ReplaceCallbackQueryMessage bool                `json:"replace_callback_query_message"`
	ProtectContent              bool                `json:"protect_content"`
	OnlyPreview                 bool                `json:"only_preview"`
	SendingId                   int32               `json:"sending_id"`
	InputMessageContent         json.RawMessage     `json:"input_message_content"`
	ReplyTo                     json.RawMessage     `json:"reply_to"`
	ReplyMarkup                 json.RawMessage     `json:"reply_markup"`
	TopicId                     json.RawMessage     `json:"topic_id"`
	Options                     *MessageSendOptions `json:"options"`
}

func captureHelperRequests(t *testing.T, c *Client) *[]helperWireRequest {
	t.Helper()
	requests := []helperWireRequest{}
	c.sendRequest = func(_ int, data []byte) {
		var req helperWireRequest
		if err := json.Unmarshal(data, &req); err != nil {
			t.Fatal(err)
		}
		requests = append(requests, req)
		var response TlObject = &Message{ChatId: req.ChatId, Id: 99}
		switch {
		case strings.HasPrefix(req.Type, "editInlineMessage"), req.Type == "editEphemeralMessage":
			response = &Ok{}
		case strings.HasPrefix(req.Type, "editBusinessMessage"):
			response = &BusinessMessage{Message: &Message{ChatId: req.ChatId, Id: req.MessageId}}
		}
		c.deliverResponse(req.Extra, response)
	}
	return &requests
}

func helperRequestContent(t *testing.T, req helperWireRequest) InputMessageContent {
	t.Helper()
	content, err := UnmarshalInputMessageContent(req.InputMessageContent)
	if err != nil {
		t.Fatal(err)
	}
	return content
}

func helperRequestMarkup(t *testing.T, req helperWireRequest) ReplyMarkup {
	t.Helper()
	markup, err := UnmarshalReplyMarkup(req.ReplyMarkup)
	if err != nil {
		t.Fatal(err)
	}
	return markup
}

func helperReplyId(t *testing.T, req helperWireRequest) int64 {
	t.Helper()
	if len(req.ReplyTo) == 0 {
		return 0
	}
	reply, err := UnmarshalInputMessageReplyTo(req.ReplyTo)
	if err != nil {
		t.Fatal(err)
	}
	if reply == nil {
		return 0
	}
	return reply.(*InputMessageReplyToMessage).MessageId
}
