package gotdbot

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestUpdateIdentifiersDoNotConfuseQueriesAndMessages(t *testing.T) {
	message := &Message{ChatId: 42, Id: 77, SenderId: &MessageSenderChat{ChatId: -123}}
	for _, tc := range []struct {
		update  TlObject
		message int64
		sender  int64
	}{
		{nil, 0, 0}, {(*UpdateNewCallbackQuery)(nil), 0, 0}, {&GetMessage{MessageId: 77}, 0, 0},
		{&UpdateNewMessage{Message: message}, 77, -123},
		{&UpdateNewBusinessMessage{Message: &BusinessMessage{Message: message}}, 77, -123},
		{&UpdateBusinessMessageEdited{Message: &BusinessMessage{Message: message}}, 77, -123},
		{&UpdateNewCallbackQuery{Id: 9007199254740993, MessageId: 77, SenderUserId: 123}, 77, 123},
		{&UpdateNewInlineQuery{Id: 9007199254740993, SenderUserId: 123}, 0, 123},
		{&UpdateNewShippingQuery{Id: 9007199254740993, SenderUserId: 123}, 0, 123},
		{&UpdateNewPreCheckoutQuery{Id: 9007199254740993, SenderUserId: 123}, 0, 123},
		{&UpdateNewChatJoinRequest{Request: &ChatJoinRequest{UserId: 123}}, 0, 123},
		{&UpdateNewChatJoinRequest{}, 0, 0},
		{&UpdateNewBusinessCallbackQuery{SenderUserId: 123, Message: &BusinessMessage{Message: message}}, 77, 123},
		{&UpdateNewBusinessCallbackQuery{Message: &BusinessMessage{Message: message}}, 77, 0},
		{&UpdateChatAction{SenderId: &MessageSenderChat{ChatId: -123}}, 0, -123},
		{&UpdateChatAction{SenderId: (*MessageSenderUser)(nil)}, 0, 0},
		{&UpdateNewGuestQuery{Message: message}, 77, -123},
		{&UpdateMessageSendSucceeded{OldMessageId: -1, Message: message}, 77, -123},
		{&UpdateMessageSendFailed{OldMessageId: -1, Message: message}, 77, -123},
		{&UpdateDeleteMessages{MessageIds: []int64{77, 78}}, 0, 0},
	} {
		if got := new(Client).UpdateMessageID(tc.update); got != tc.message {
			t.Errorf("%T message ID = %d, want %d", tc.update, got, tc.message)
		}
		if got := new(Client).UpdateSenderID(tc.update); got != tc.sender {
			t.Errorf("%T sender ID = %d, want %d", tc.update, got, tc.sender)
		}
	}
}

func TestMessageHandlerDoesNotHandleMessageCarryingQueriesOrSendEvents(t *testing.T) {
	message := testMessage(42, "hello")
	handler := NewMessageHandler(func(*Client, *Message) error { return nil }, nil)
	for _, update := range []TlObject{
		&UpdateNewGuestQuery{Message: message},
		&UpdateNewBusinessCallbackQuery{Message: &BusinessMessage{Message: message}},
		&UpdateMessageSendSucceeded{Message: message},
		&UpdateMessageSendFailed{Message: message},
	} {
		if got, ok := messageFromUpdate(update); !ok || got != message {
			t.Fatalf("%T full message extraction failed", update)
		}
		if handler.CheckUpdate(nil, update) {
			t.Fatalf("message handler accepted %T", update)
		}
	}
}

func TestUpdateMessageGettersPreserveIdentity(t *testing.T) {
	message := &Message{ChatId: 42, Id: 77}
	getters := []func() *Message{
		(&UpdateNewMessage{Message: message}).GetMessage,
		(&UpdateNewBusinessMessage{Message: &BusinessMessage{Message: message}}).GetMessage,
		(&UpdateBusinessMessageEdited{Message: &BusinessMessage{Message: message}}).GetMessage,
		(&UpdateMessageSendSucceeded{Message: message}).GetMessage,
		(&UpdateMessageSendFailed{Message: message}).GetMessage,
	}
	for _, get := range getters {
		if get() != message {
			t.Fatal("getter copied or lost the message")
		}
	}
	var absent *UpdateNewMessage
	if absent.GetMessage() != nil || (&UpdateNewBusinessMessage{}).GetMessage() != nil || (&UpdateBusinessMessageEdited{}).GetMessage() != nil {
		t.Fatal("missing message was not nil")
	}
}

func TestEditUpdateFetchesCurrentMessageAndReturnsErrors(t *testing.T) {
	for _, fail := range []bool{false, true} {
		for _, content := range []bool{false, true} {
			t.Run(map[bool]string{false: "edit", true: "content"}[content]+"/"+map[bool]string{false: "success", true: "failure"}[fail], func(t *testing.T) {
				c := newTestClient(t)
				message := &Message{ChatId: 42, Id: 77}
				failure := &Error{Code: 403, Message: "message inaccessible"}
				calls := 0
				c.sendRequest = func(_ int, data []byte) {
					calls++
					request := decodeTestRequest(t, data)
					if request.Type != "getMessage" || request.ChatID != 42 || request.MessageID != 77 {
						t.Fatalf("wrong message lookup: %+v", request)
					}
					var result TlObject = message
					if fail {
						result = failure
					}
					c.deliverResponse(request.Extra, result)
				}
				var got *Message
				var err error
				if content {
					got, err = (&UpdateMessageContent{ChatId: 42, MessageId: 77}).GetMessage(c)
				} else {
					got, err = (&UpdateMessageEdited{ChatId: 42, MessageId: 77}).GetMessage(c)
				}
				if calls != 1 {
					t.Fatalf("lookup sent %d requests", calls)
				}
				if fail {
					if got != nil || !errors.Is(err, failure) {
						t.Fatalf("failure = (%v, %v)", got, err)
					}
				} else if got != message || err != nil {
					t.Fatalf("lookup = (%v, %v)", got, err)
				}
				assertNoPendingCalls(t, c)
			})
		}
	}
}

type queryAnswerWire struct {
	Type               string           `json:"@type"`
	Extra              string           `json:"@extra"`
	InlineQueryId      int64            `json:"inline_query_id,string"`
	PreCheckoutQueryId int64            `json:"pre_checkout_query_id,string"`
	ShippingQueryId    int64            `json:"shipping_query_id,string"`
	ChatId             int64            `json:"chat_id"`
	UserId             int64            `json:"user_id"`
	CacheTime          int32            `json:"cache_time"`
	NextOffset         string           `json:"next_offset"`
	IsPersonal         bool             `json:"is_personal"`
	ErrorMessage       string           `json:"error_message"`
	ShippingOptions    []ShippingOption `json:"shipping_options"`
	Approve            bool             `json:"approve"`
}

func TestQueryActionsUseUpdateIdentifiersAndDecision(t *testing.T) {
	const queryID int64 = 9007199254740993
	shipping := []ShippingOption{{Id: "fast", Title: "Fast", PriceParts: []LabeledPricePart{{Label: "Delivery", Amount: 100}}}}
	for _, action := range []string{"inline", "checkout approve", "checkout reject", "shipping", "shipping reject", "join approve", "join decline"} {
		t.Run(action, func(t *testing.T) {
			c := newTestClient(t)
			var wire queryAnswerWire
			calls := 0
			c.sendRequest = func(_ int, data []byte) {
				calls++
				if err := json.Unmarshal(data, &wire); err != nil {
					t.Fatal(err)
				}
				c.deliverResponse(wire.Extra, &Ok{})
			}
			var err error
			switch action {
			case "inline":
				err = (&UpdateNewInlineQuery{Id: queryID}).Answer(c, nil, &AnswerInlineQueryOpts{NextOffset: "next", CacheTime: 60, IsPersonal: true})
				if wire.Type != "answerInlineQuery" || wire.InlineQueryId != queryID || wire.CacheTime != 60 || wire.NextOffset != "next" || !wire.IsPersonal {
					t.Fatalf("inline answer lost identifiers/options: %+v", wire)
				}
			case "checkout approve", "checkout reject":
				message := ""
				if action == "checkout reject" {
					message = "unavailable"
				}
				err = (&UpdateNewPreCheckoutQuery{Id: queryID}).Answer(c, message)
				if wire.Type != "answerPreCheckoutQuery" || wire.PreCheckoutQueryId != queryID || wire.ErrorMessage != message {
					t.Fatalf("checkout decision lost: %+v", wire)
				}
			case "shipping", "shipping reject":
				options := shipping
				var opts *AnswerShippingQueryOpts
				if action == "shipping reject" {
					options = nil
					opts = &AnswerShippingQueryOpts{ErrorMessage: "unavailable"}
				}
				err = (&UpdateNewShippingQuery{Id: queryID}).Answer(c, options, opts)
				if wire.Type != "answerShippingQuery" || wire.ShippingQueryId != queryID || !reflect.DeepEqual(wire.ShippingOptions, options) || (action == "shipping reject" && wire.ErrorMessage != opts.ErrorMessage) {
					t.Fatalf("shipping options lost: %+v", wire)
				}
			case "join approve", "join decline":
				update := &UpdateNewChatJoinRequest{ChatId: 42, Request: &ChatJoinRequest{UserId: 123}}
				if action == "join approve" {
					err = update.Approve(c)
				} else {
					err = update.Decline(c)
				}
				if wire.Type != "processChatJoinRequest" || wire.ChatId != 42 || wire.UserId != 123 || wire.Approve != (action == "join approve") {
					t.Fatalf("join decision lost: %+v", wire)
				}
			}
			if err != nil || calls != 1 {
				t.Fatalf("action = %v; requests = %d", err, calls)
			}
		})
	}
}

func TestNilUpdateActionsDoNotSendRequests(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(int, []byte) { t.Fatal("nil update sent a request") }
	var inline *UpdateNewInlineQuery
	var checkout *UpdateNewPreCheckoutQuery
	var shipping *UpdateNewShippingQuery
	var join *UpdateNewChatJoinRequest
	var edited *UpdateMessageEdited
	var content *UpdateMessageContent
	_, editErr := edited.GetMessage(c)
	_, contentErr := content.GetMessage(c)
	for _, err := range []error{inline.Answer(c, nil, nil), checkout.Answer(c, ""), shipping.Answer(c, nil, nil), join.Approve(c), join.Decline(c), editErr, contentErr} {
		if !errors.Is(err, ErrNilUpdate) {
			t.Fatalf("nil action error = %v", err)
		}
	}
	if err := (&UpdateNewChatJoinRequest{}).Approve(c); err == nil {
		t.Fatal("missing join request accepted")
	}
}
