package gotdbot

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestRichSendPreservesSourceAndDeliveryOptions(t *testing.T) {
	for _, format := range []string{"markdown", "html"} {
		t.Run(format, func(t *testing.T) {
			c := newTestClient(t)
			// Rich helpers must not send their source through ParseTextEntities.
			c.config.ParseMode = "HTML"
			requests := captureHelperRequests(t, c)
			opts := &SendRichTextOpts{
				SendMessageOpts: SendMessageOpts{
					ClearDraft: true, DisableNotification: true, ProtectContent: true,
					AllowPaidBroadcast: true, EffectId: 123, FromBackground: true,
					OnlyPreview: true, PaidMessageStarCount: 7, SendingId: 8,
					SchedulingState:   &MessageSchedulingStateSendAtDate{SendDate: 42},
					SuggestedPostInfo: &InputSuggestedPostInfo{}, UpdateOrderOfInstalledStickerSets: true,
					TopicId: &MessageTopicForum{ForumTopicId: 9}, ReplyToMessageID: 10,
					ReplyMarkup: &ReplyMarkupInlineKeyboard{}, Quote: &InputTextQuote{Text: &FormattedText{Text: "quote"}},
				},
				Media: []InputRichMessageMedia{{Id: "image", Media: &InputMessagePhoto{Photo: &InputPhoto{Photo: &InputFileLocal{Path: "photo.jpg"}}}}},
				IsRtl: true, DetectAutomaticBlocks: true,
			}
			before := *opts
			text := "# Hello 😀\n![photo](tg://photo?id=image)"
			var message *Message
			var err error
			var source RichMessageSource = &RichMessageSourceMarkdown{Text: text, Media: opts.Media}
			if format == "html" {
				text = "<h1>Hello 😀</h1><img src=\"tg://photo?id=image\">"
				source = &RichMessageSourceHtml{Text: text, Media: opts.Media}
				message, err = c.SendRichMessageHTML(12, text, opts)
			} else {
				message, err = c.SendRichMessageMarkdown(12, text, opts)
			}
			if err != nil || message.Id != 99 {
				t.Fatalf("send = (%v, %v)", message, err)
			}
			if len(*requests) != 1 || (*requests)[0].Type != "sendMessage" {
				t.Fatalf("unexpected requests: %+v", *requests)
			}
			req := (*requests)[0]
			content := helperRequestContent(t, req).(*InputMessageRichMessage)
			want := &InputRichMessage{Source: source, IsRtl: true, DetectAutomaticBlocks: true}
			if !content.ClearDraft || !reflect.DeepEqual(content.Message, want) {
				t.Fatalf("rich source lost: %+v", content)
			}
			wantOptions := &MessageSendOptions{
				DisableNotification: true, ProtectContent: true, AllowPaidBroadcast: true,
				EffectId: 123, FromBackground: true, OnlyPreview: true, PaidMessageStarCount: 7,
				SendingId: 8, SchedulingState: opts.SchedulingState, SuggestedPostInfo: opts.SuggestedPostInfo,
				UpdateOrderOfInstalledStickerSets: true,
			}
			if req.ChatId != 12 || helperReplyId(t, req) != 10 || !reflect.DeepEqual(req.Options, wantOptions) {
				t.Fatalf("delivery options lost: %+v", req)
			}
			if !reflect.DeepEqual(helperRequestMarkup(t, req), opts.ReplyMarkup) {
				t.Fatal("reply markup lost")
			}
			var topic MessageTopicForum
			var reply InputMessageReplyToMessage
			if err := json.Unmarshal(req.TopicId, &topic); err != nil || topic.ForumTopicId != 9 {
				t.Fatalf("topic lost: %s", req.TopicId)
			}
			if err := json.Unmarshal(req.ReplyTo, &reply); err != nil || reply.Quote.Text.Text != "quote" {
				t.Fatalf("quote lost: %s", req.ReplyTo)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("send mutated options")
			}
		})
	}
}

func TestRichRepliesPreserveExplicitTargets(t *testing.T) {
	for _, target := range []string{"default", "message id", "reply object", "nil options"} {
		for _, format := range []string{"markdown", "html", "blocks"} {
			t.Run(target+"/"+format, func(t *testing.T) {
				c := newTestClient(t)
				requests := captureHelperRequests(t, c)
				m := &Message{ChatId: 12, Id: 20}
				opts := &SendRichTextOpts{}
				wantId := int64(20)
				switch target {
				case "message id":
					opts.ReplyToMessageID, wantId = 30, 30
				case "reply object":
					opts.ReplyTo = &InputMessageReplyToMessage{MessageId: 40}
					wantId = 40
				case "nil options":
					opts = nil
				}
				var before SendRichTextOpts
				if opts != nil {
					before = *opts
				}
				var err error
				switch format {
				case "markdown":
					_, err = m.ReplyRichMessageMarkdown(c, "# Hello", opts)
				case "html":
					_, err = m.ReplyRichMessageHTML(c, "<h1>Hello</h1>", opts)
				case "blocks":
					var sendOpts *SendMessageOpts
					if opts != nil {
						sendOpts = &opts.SendMessageOpts
					}
					_, err = m.ReplyRichMessage(c, &InputRichMessage{Source: &RichMessageSourceBlocks{}}, sendOpts)
				}
				if err != nil {
					t.Fatal(err)
				}
				if len(*requests) != 1 || (*requests)[0].ChatId != 12 || helperReplyId(t, (*requests)[0]) != wantId {
					t.Fatalf("reply target lost: %+v", *requests)
				}
				if opts != nil && !reflect.DeepEqual(*opts, before) {
					t.Fatal("reply mutated options")
				}
			})
		}
	}
}

func TestRichEditsRouteChatInlineAndEphemeralMessages(t *testing.T) {
	for _, target := range []string{"chat", "message", "callback", "inline", "inline callback", "ephemeral"} {
		for _, format := range []string{"markdown", "html", "blocks"} {
			t.Run(target+"/"+format, func(t *testing.T) {
				c := newTestClient(t)
				requests := captureHelperRequests(t, c)
				markup := &ReplyMarkupInlineKeyboard{}
				opts := &EditRichTextOpts{IsRtl: true, DetectAutomaticBlocks: true, ReplyMarkup: markup,
					Media: []InputRichMessageMedia{{Id: "voice", Media: &InputMessageVoiceNote{VoiceNote: &InputVoiceNote{VoiceNote: &InputFileLocal{Path: "voice.ogg"}}}}},
				}
				before := *opts
				m := &Message{ChatId: 12, Id: 20}
				cb := &UpdateNewCallbackQuery{ChatId: 12, MessageId: 20}
				inline := &UpdateNewInlineCallbackQuery{InlineMessageId: "inline-id"}
				text := "rich 😀"
				var source RichMessageSource
				var err error
				if format == "blocks" {
					source = &RichMessageSourceBlocks{Blocks: []InputPageBlock{&InputPageBlockParagraph{Text: &RichTextPlain{Text: text}}}}
					rich := &InputRichMessage{Source: source, IsRtl: true, DetectAutomaticBlocks: true}
					richOpts := &EditRichMessageOpts{ReplyMarkup: markup}
					switch target {
					case "chat":
						_, err = c.EditRichMessage(12, 20, rich, richOpts)
					case "message":
						_, err = m.EditRichMessage(c, rich, richOpts)
					case "callback":
						_, err = cb.EditRichMessage(c, rich, richOpts)
					case "inline":
						err = c.EditInlineRichMessage("inline-id", rich, richOpts)
					case "inline callback":
						err = inline.EditRichMessage(c, rich, richOpts)
					case "ephemeral":
						err = c.EditEphemeralRichMessage(12, 3, 44, rich, richOpts)
					}
				} else if format == "markdown" {
					source = &RichMessageSourceMarkdown{Text: text, Media: opts.Media}
					switch target {
					case "chat":
						_, err = c.EditRichMessageMarkdown(12, 20, text, opts)
					case "message":
						_, err = m.EditRichMessageMarkdown(c, text, opts)
					case "callback":
						_, err = cb.EditRichMessageMarkdown(c, text, opts)
					case "inline":
						err = c.EditInlineRichMessageMarkdown("inline-id", text, opts)
					case "inline callback":
						err = inline.EditRichMessageMarkdown(c, text, opts)
					case "ephemeral":
						err = c.EditEphemeralRichMessageMarkdown(12, 3, 44, text, opts)
					}
				} else {
					source = &RichMessageSourceHtml{Text: text, Media: opts.Media}
					switch target {
					case "chat":
						_, err = c.EditRichMessageHTML(12, 20, text, opts)
					case "message":
						_, err = m.EditRichMessageHTML(c, text, opts)
					case "callback":
						_, err = cb.EditRichMessageHTML(c, text, opts)
					case "inline":
						err = c.EditInlineRichMessageHTML("inline-id", text, opts)
					case "inline callback":
						err = inline.EditRichMessageHTML(c, text, opts)
					case "ephemeral":
						err = c.EditEphemeralRichMessageHTML(12, 3, 44, text, opts)
					}
				}
				if err != nil || len(*requests) != 1 {
					t.Fatalf("edit = %v, requests: %+v", err, *requests)
				}
				req := (*requests)[0]
				wantType := "editMessageText"
				switch target {
				case "inline", "inline callback":
					wantType = "editInlineMessageText"
					if req.InlineMessageId != "inline-id" || req.ChatId != 0 {
						t.Fatalf("inline routing lost: %+v", req)
					}
				case "ephemeral":
					wantType = "editEphemeralMessage"
					if req.ChatId != 12 || req.EphemeralMessageId != 3 || req.ReceiverUserId != 44 {
						t.Fatalf("ephemeral routing lost: %+v", req)
					}
				default:
					if req.ChatId != 12 || req.MessageId != 20 {
						t.Fatalf("chat routing lost: %+v", req)
					}
				}
				content := helperRequestContent(t, req).(*InputMessageRichMessage)
				want := &InputRichMessage{Source: source, IsRtl: true, DetectAutomaticBlocks: true}
				if req.Type != wantType || !reflect.DeepEqual(content.Message, want) || !reflect.DeepEqual(helperRequestMarkup(t, req), markup) {
					t.Fatalf("rich edit lost source/markup: %+v", req)
				}
				if !reflect.DeepEqual(*opts, before) {
					t.Fatal("edit mutated options")
				}
			})
		}
	}
}

func TestRichEditLoadsMessageWithoutDatabase(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(map[bool]string{false: "success", true: "load error"}[fail], func(t *testing.T) {
			c := newTestClient(t)
			c.config.UseMessageDatabase = new(false)
			calls := []string{}
			failure := &Error{Code: 404, Message: "message not found"}
			c.sendRequest = func(_ int, data []byte) {
				req := decodeTestRequest(t, data)
				calls = append(calls, req.Type)
				if req.ChatID != 12 || req.MessageID != 20 {
					t.Fatal("load/edit targeted wrong message")
				}
				var response TlObject = &Message{ChatId: 12, Id: 20}
				if fail {
					response = failure
				}
				c.deliverResponse(req.Extra, response)
			}
			_, err := c.EditRichMessageHTML(12, 20, "<h1>Hello</h1>", nil)
			wantCalls := []string{"getMessage", "editMessageText"}
			if fail {
				wantCalls = []string{"getMessage"}
				if !errors.Is(err, failure) {
					t.Fatalf("load error not returned: %v", err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(calls, wantCalls) {
				t.Fatalf("calls = %v, want %v", calls, wantCalls)
			}
			assertNoPendingCalls(t, c)
		})
	}
}
