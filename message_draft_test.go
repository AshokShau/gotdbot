package gotdbot

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type draftWireRequest struct {
	Type         string            `json:"@type"`
	Extra        string            `json:"@extra"`
	ChatId       int64             `json:"chat_id"`
	DraftId      int64             `json:"draft_id,string"`
	ForumTopicId int32             `json:"forum_topic_id"`
	CanStop      bool              `json:"can_stop"`
	KeepOnStop   bool              `json:"keep_on_stop"`
	Text         *FormattedText    `json:"text"`
	Message      *InputRichMessage `json:"message"`
}

func TestTextDraftFormattingControlsAndParseFailure(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(map[bool]string{false: "formatted", true: "parse failure"}[fail], func(t *testing.T) {
			c := newTestClient(t)
			var calls []string
			formatted := &FormattedText{Text: "Hello 😀", Entities: []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeBold{}}}}
			failure := &Error{Code: 400, Message: "invalid formatting"}
			c.sendRequest = func(_ int, data []byte) {
				req := decodeTestRequest(t, data)
				calls = append(calls, req.Type)
				var response TlObject = &Ok{}
				if req.Type == "parseTextEntities" {
					response = formatted
					if fail {
						response = failure
					}
				} else {
					var draft draftWireRequest
					if err := json.Unmarshal(data, &draft); err != nil {
						t.Fatal(err)
					}
					if draft.ChatId != 42 || draft.DraftId != 9007199254740993 || draft.ForumTopicId != 7 || !draft.CanStop || !draft.KeepOnStop || !reflect.DeepEqual(draft.Text, formatted) {
						t.Fatalf("draft lost formatting or controls: %+v", draft)
					}
				}
				c.deliverResponse(req.Extra, response)
			}
			err := (&Chat{Id: 42}).SendMessageDraft(c, 9007199254740993, "<b>Hello</b> 😀", &SendMessageDraftOpts{
				ParseMode: "HTML", ForumTopicId: 7, CanStop: true, KeepOnStop: true,
			})
			want := []string{"parseTextEntities", "sendTextMessageDraft"}
			if fail {
				want = []string{"parseTextEntities"}
				if !errors.Is(err, failure) {
					t.Fatalf("parse error = %v", err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(calls, want) {
				t.Fatalf("requests = %v, want %v", calls, want)
			}
			assertNoPendingCalls(t, c)
		})
	}
}

func TestRichDraftPreservesSourceWithoutTextParsing(t *testing.T) {
	for _, html := range []bool{false, true} {
		t.Run(map[bool]string{false: "markdown", true: "html"}[html], func(t *testing.T) {
			c := newTestClient(t)
			c.config.ParseMode = "HTML"
			opts := &SendRichMessageDraftTextOpts{SendRichMessageDraftOpts: SendRichMessageDraftOpts{ForumTopicId: 7, CanStop: true, KeepOnStop: true},
				IsRtl: true, DetectAutomaticBlocks: true,
				Media: []InputRichMessageMedia{{Id: "photo", Media: &InputMessagePhoto{Photo: &InputPhoto{Photo: &InputFileId{Id: 1}}}}},
			}
			before := *opts
			count := 0
			var source RichMessageSource = &RichMessageSourceMarkdown{Text: "rich draft", Media: opts.Media}
			if html {
				source = &RichMessageSourceHtml{Text: "rich draft", Media: opts.Media}
			}
			c.sendRequest = func(_ int, data []byte) {
				count++
				var req draftWireRequest
				if err := json.Unmarshal(data, &req); err != nil {
					t.Fatal(err)
				}
				wantMessage := &InputRichMessage{Source: source, IsRtl: true, DetectAutomaticBlocks: true}
				if req.Type != "sendRichMessageDraft" || req.ChatId != 42 || req.DraftId != 11 || req.ForumTopicId != 7 || !req.CanStop || !req.KeepOnStop || !reflect.DeepEqual(req.Message, wantMessage) {
					t.Fatalf("rich draft lost source or controls: %+v", req)
				}
				c.deliverResponse(req.Extra, &Ok{})
			}
			var err error
			chat := &Chat{Id: 42}
			if html {
				err = chat.SendRichMessageDraftHTML(c, 11, "rich draft", opts)
			} else {
				err = chat.SendRichMessageDraftMarkdown(c, 11, "rich draft", opts)
			}
			if err != nil || count != 1 || !reflect.DeepEqual(*opts, before) {
				t.Fatalf("draft error/count/options: %v/%d/%+v", err, count, opts)
			}
		})
	}
}
