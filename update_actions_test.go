package gotdbot

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestCallbackAnswersForwardOptionsForAllCallbackTypes(t *testing.T) {
	for _, target := range []string{"chat", "inline", "business"} {
		for _, configured := range []bool{false, true} {
			t.Run(target+map[bool]string{false: "/defaults", true: "/configured"}[configured], func(t *testing.T) {
				c := newTestClient(t)
				var request struct {
					Type      string `json:"@type"`
					Extra     string `json:"@extra"`
					QueryID   int64  `json:"callback_query_id,string"`
					Text, Url string
					CacheTime int32 `json:"cache_time"`
					ShowAlert bool  `json:"show_alert"`
				}
				calls := 0
				c.sendRequest = func(_ int, data []byte) {
					calls++
					if err := json.Unmarshal(data, &request); err != nil {
						t.Fatal(err)
					}
					c.deliverResponse(request.Extra, &Ok{})
				}
				var opts *CallbackAnswerOpts
				if configured {
					opts = &CallbackAnswerOpts{CacheTime: 60, ShowAlert: true, Url: "https://t.me/example_bot?start=next"}
				}
				var err error
				const queryID int64 = 9007199254740993
				switch target {
				case "chat":
					err = (&UpdateNewCallbackQuery{Id: queryID}).Answer(c, "done", opts)
				case "inline":
					err = (&UpdateNewInlineCallbackQuery{Id: queryID}).Answer(c, "done", opts)
				case "business":
					err = (&UpdateNewBusinessCallbackQuery{Id: queryID}).Answer(c, "done", opts)
				}
				if err != nil || calls != 1 || request.Type != "answerCallbackQuery" || request.QueryID != queryID || request.Text != "done" || request.ShowAlert != configured {
					t.Fatalf("answer = %v, request = %+v", err, request)
				}
				if configured && (request.CacheTime != 60 || request.Url != opts.Url) {
					t.Fatal("lost callback options")
				}
			})
		}
	}
}

func TestCallbackPayloadsAcceptValuesAndTypedNil(t *testing.T) {
	for _, payload := range []CallbackQueryPayload{CallbackQueryPayloadData{Data: []byte("data")}, &CallbackQueryPayloadData{Data: []byte("data")}, CallbackQueryPayloadDataWithPassword{Data: []byte("data"), Password: "password"}, &CallbackQueryPayloadDataWithPassword{Data: []byte("data"), Password: "password"}} {
		chat := &UpdateNewCallbackQuery{Payload: payload}
		inline := &UpdateNewInlineCallbackQuery{Payload: payload}
		business := &UpdateNewBusinessCallbackQuery{Payload: payload}
		if chat.DataString() != "data" || inline.DataString() != "data" || business.DataString() != "data" {
			t.Fatalf("lost payload %T", payload)
		}
	}
	for _, payload := range []CallbackQueryPayload{nil, (*CallbackQueryPayloadData)(nil), (*CallbackQueryPayloadDataWithPassword)(nil), (*CallbackQueryPayloadGame)(nil)} {
		chat := &UpdateNewCallbackQuery{Payload: payload}
		inline := &UpdateNewInlineCallbackQuery{Payload: payload}
		business := &UpdateNewBusinessCallbackQuery{Payload: payload}
		if chat.DataString() != "" || inline.DataString() != "" || business.DataString() != "" || chat.GameShortName() != "" || business.GameShortName() != "" {
			t.Fatalf("typed nil payload %T produced data", payload)
		}
	}
	inline := &UpdateNewInlineCallbackQuery{Payload: CallbackQueryPayloadDataWithPassword{Data: []byte("data"), Password: "password"}}
	if data, password := inline.CallbackDataWithPassword(); string(data) != "data" || password != "password" {
		t.Fatal("lost inline password")
	}
}

func TestGetMessageFromUpdateUsesCarriedMessageOrFetchesIDs(t *testing.T) {
	c := newTestClient(t)
	message := &Message{ChatId: 42, Id: 20}
	var requests []testWireRequest
	c.sendRequest = func(_ int, data []byte) {
		req := decodeTestRequest(t, data)
		requests = append(requests, req)
		c.deliverResponse(req.Extra, message)
	}
	for _, update := range []Update{&UpdateNewMessage{Message: message}, &UpdateNewGuestQuery{Message: message}, &UpdateNewBusinessCallbackQuery{Message: &BusinessMessage{Message: message}}, &UpdateMessageSendSucceeded{Message: message, OldMessageId: 19}} {
		got, err := c.GetUpdateMessage(update)
		if err != nil || got != message || len(requests) != 0 {
			t.Fatalf("carried message = %v, %v", got, err)
		}
	}
	got, err := c.GetUpdateMessage(&UpdateMessageIsPinned{ChatId: 42, MessageId: 20})
	if err != nil || got != message || len(requests) != 1 || requests[0].Type != "getMessage" || requests[0].ChatID != 42 || requests[0].MessageID != 20 {
		t.Fatalf("ID lookup = %v, %v, %v", got, err, requests)
	}
	if _, err := c.GetUpdateMessage(&UpdateDeleteMessages{ChatId: 42, MessageIds: []int64{20}}); !errors.Is(err, ErrUpdateHasNoMessage) {
		t.Fatalf("deletion error = %v", err)
	}
	if _, err := c.GetUpdateMessage((*UpdateMessageEdited)(nil)); !errors.Is(err, ErrNilUpdate) {
		t.Fatalf("nil error = %v", err)
	}
	if len(requests) != 1 {
		t.Fatal("unavailable message sent a request")
	}
}

func TestBusinessCallbackEditsUseConnectionAndCachedMessage(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	message := &Message{ChatId: 42, Id: 20}
	query := &UpdateNewBusinessCallbackQuery{ConnectionId: "connection", Message: &BusinessMessage{Message: message}}
	if got, err := query.GetMessage(); err != nil || got != message {
		t.Fatal("lost cached business message")
	}
	markup := &ReplyMarkupInlineKeyboard{}
	if _, err := query.EditMessageText(c, "text", &EditMessageTextOpts{ReplyMarkup: markup}); err != nil {
		t.Fatal(err)
	}
	if _, err := query.EditRichMessageHTML(c, "<h1>rich</h1>", &EditRichTextOpts{ReplyMarkup: markup, DetectAutomaticBlocks: true}); err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 2 {
		t.Fatalf("requests = %+v", *requests)
	}
	for _, req := range *requests {
		if req.Type != "editBusinessMessageText" || req.BusinessConnectionId != "connection" || req.ChatId != 42 || req.MessageId != 20 || !reflect.DeepEqual(helperRequestMarkup(t, req), markup) {
			t.Fatalf("business target/options lost: %+v", req)
		}
	}
	rich := helperRequestContent(t, (*requests)[1]).(*InputMessageRichMessage)
	if rich.Message.Source.(*RichMessageSourceHtml).Text != "<h1>rich</h1>" || !rich.Message.DetectAutomaticBlocks {
		t.Fatal("business rich source/options lost")
	}
}

func TestQueryHelpersRejectMissingUpdatesAndMessages(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(int, []byte) { t.Fatal("invalid update sent a request") }
	var chat *UpdateNewCallbackQuery
	var inline *UpdateNewInlineCallbackQuery
	var business *UpdateNewBusinessCallbackQuery
	var guest *UpdateNewGuestQuery
	var chosen *UpdateNewChosenInlineResult
	_, businessErr := business.EditMessageText(c, "text", nil)
	_, chatErr := chat.EditMessagePhoto(c, &InputFileId{Id: 1}, nil)
	_, guestErr := guest.Answer(c, &InputInlineQueryResultArticle{})
	for _, err := range []error{chat.Answer(c, "", nil), inline.Answer(c, "", nil), business.Answer(c, "", nil), inline.EditMessageText(c, "", nil), chosen.EditText(c, "", nil), businessErr, chatErr, guestErr} {
		if !errors.Is(err, ErrNilUpdate) {
			t.Fatalf("nil update error = %v", err)
		}
	}
	if err := (&UpdateNewChosenInlineResult{}).EditText(c, "", nil); !errors.Is(err, ErrUpdateHasNoMessage) {
		t.Fatal(err)
	}
	if _, err := (&UpdateNewBusinessCallbackQuery{}).EditMessageText(c, "", nil); !errors.Is(err, ErrUpdateHasNoMessage) {
		t.Fatal(err)
	}
}

func TestGuestQueryAnswerUsesItsIDAndNativeResult(t *testing.T) {
	c := newTestClient(t)
	var request struct {
		Type   string `json:"@type"`
		Extra  string `json:"@extra"`
		ID     int64  `json:"guest_query_id,string"`
		Result json.RawMessage
	}
	c.sendRequest = func(_ int, data []byte) {
		if err := json.Unmarshal(data, &request); err != nil {
			t.Fatal(err)
		}
		c.deliverResponse(request.Extra, &InlineMessageId{})
	}
	message := &Message{ChatId: 42, Id: 20}
	query := &UpdateNewGuestQuery{Id: 9007199254740993, Message: message}
	result := &InputInlineQueryResultArticle{Id: "answer", Title: "Title", InputMessageContent: &InputMessageText{Text: &FormattedText{Text: "hello"}}}
	if _, err := query.Answer(c, result); err != nil {
		t.Fatal(err)
	}
	got, err := UnmarshalInputInlineQueryResult(request.Result)
	if err != nil || !reflect.DeepEqual(got, result) || request.Type != "answerGuestQuery" || request.ID != query.Id || query.GetMessage() != message {
		t.Fatalf("guest answer = %+v, %v", request, err)
	}
}

func TestBusinessCallbackGetMessageReportsMissingData(t *testing.T) {
	var missing *UpdateNewBusinessCallbackQuery
	if _, err := missing.GetMessage(); !errors.Is(err, ErrNilUpdate) {
		t.Fatal(err)
	}
	for _, query := range []*UpdateNewBusinessCallbackQuery{{}, {Message: &BusinessMessage{}}} {
		if _, err := query.GetMessage(); !errors.Is(err, ErrUpdateHasNoMessage) {
			t.Fatal(err)
		}
	}
	message := &Message{ChatId: 42, Id: 20}
	if got, err := (&UpdateNewBusinessCallbackQuery{Message: &BusinessMessage{Message: message}}).GetMessage(); err != nil || got != message {
		t.Fatalf("business message = %v, %v", got, err)
	}
}
