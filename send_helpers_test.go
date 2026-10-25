package gotdbot

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestSendHelpersExposeEphemeralDeliveryOptions(t *testing.T) {
	for _, tc := range []struct {
		name string
		send func(*Client) error
	}{
		{"contact", func(c *Client) error {
			_, err := c.SendContact(42, &Contact{}, &SendContactOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"venue", func(c *Client) error {
			_, err := c.SendVenue(42, &Venue{}, &SendVenueOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"checklist", func(c *Client) error {
			_, err := c.SendChecklist(42, &InputChecklist{}, &SendChecklistOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"dice", func(c *Client) error {
			_, err := c.SendDice(42, "🎲", &SendDiceOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"game", func(c *Client) error {
			_, err := c.SendGame(42, 7, "game", &SendGameOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"poll", func(c *Client) error {
			_, err := c.SendPoll(42, "question", nil, &SendPollOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"stakeDice", func(c *Client) error {
			_, err := c.SendStakeDice(42, 1, "state", &SendStakeDiceOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"story", func(c *Client) error {
			_, err := c.SendStory(42, 7, 3, &SendStoryOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"invoice", func(c *Client) error {
			_, err := c.SendInvoice(42, &Invoice{}, "title", "description", nil, &SendInvoiceOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"paidMedia", func(c *Client) error {
			_, err := c.SendPaidMedia(42, 1, nil, &SendPaidMediaOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"forwarded", func(c *Client) error {
			_, err := c.SendForwarded(42, 12, 20, &SendForwardedOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"location", func(c *Client) error {
			_, err := c.SendLocation(42, &Location{}, &SendLocationOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"copy", func(c *Client) error {
			_, err := c.CopyMessage(42, 12, 20, &CopyMessageOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
		{"forward", func(c *Client) error {
			_, err := c.ForwardMessage(42, 12, 20, &ForwardMessageOpts{CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true, ProtectContent: true, OnlyPreview: true, SendingId: 7})
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			requests := captureHelperRequests(t, c)
			if err := tc.send(c); err != nil {
				t.Fatal(err)
			}
			if len(*requests) != 1 {
				t.Fatalf("requests = %+v", *requests)
			}
			req := (*requests)[0]
			if req.Type != "sendEphemeralMessage" || req.ChatId != 42 || req.CallbackQueryId != 22 || req.ReceiverUserId != 44 || req.SendingId != 7 || !req.ProtectContent || !req.OnlyPreview || !req.ReplaceCallbackQueryMessage {
				t.Fatalf("lost ephemeral delivery: %+v", req)
			}
		})
	}
}

func TestRegularPollAndQuizBuildTheirTypesWithoutMutatingOptions(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	options := []InputPollOption{{Text: &FormattedText{Text: "A"}}, {Text: &FormattedText{Text: "B"}}}
	pollOpts := &SendPollOpts{AllowAddingOptions: true}
	if _, err := c.SendPoll(42, "question", options, pollOpts); err != nil {
		t.Fatal(err)
	}
	poll := helperRequestContent(t, (*requests)[0]).(*InputMessagePoll)
	regular, ok := poll.Type.(*InputPollTypeRegular)
	if !ok || !regular.AllowAddingOptions || pollOpts.Type != nil {
		t.Fatal("default poll type or caller options changed")
	}
	originalType := &InputPollTypeRegular{}
	opts := &SendQuizOpts{SendPollOpts: SendPollOpts{Type: originalType, ProtectContent: true}, Explanation: "B is correct", ExplanationEntities: []TextEntity{{Offset: 0, Length: 1, Type: &TextEntityTypeBold{}}}}
	before := *opts
	if _, err := (&Message{ChatId: 42, Id: 20}).ReplyQuiz(c, "question", options, []int32{1}, opts); err != nil {
		t.Fatal(err)
	}
	req := (*requests)[1]
	poll = helperRequestContent(t, req).(*InputMessagePoll)
	quiz, ok := poll.Type.(*InputPollTypeQuiz)
	if !ok || !reflect.DeepEqual(quiz.CorrectOptionIds, []int32{1}) || quiz.Explanation.Text != "B is correct" || !reflect.DeepEqual(quiz.Explanation.Entities, opts.ExplanationEntities) {
		t.Fatal("quiz type or explanation lost")
	}
	if helperReplyId(t, req) != 20 || !req.Options.ProtectContent || !reflect.DeepEqual(*opts, before) {
		t.Fatal("quiz reply/options changed")
	}
	for _, ids := range [][]int32{nil, {-1}, {2}, {1, 0}, {1, 1}} {
		if _, err := c.SendQuiz(42, "q", options, ids, nil); err == nil {
			t.Fatalf("accepted invalid answer IDs %v", ids)
		}
	}
	if len(*requests) != 2 {
		t.Fatal("invalid quiz sent a request")
	}
}

func TestForwardAndCopyPreserveExtraDeliveryAndContentOptions(t *testing.T) {
	for _, name := range []string{"forward", "send forwarded", "copy"} {
		t.Run(name, func(t *testing.T) {
			c := newTestClient(t)
			c.config.ParseMode = "HTML"
			requests := captureHelperRequests(t, c)
			reply := &InputMessageReplyToMessage{MessageId: 88}
			var err error
			switch name {
			case "forward":
				_, err = c.ForwardMessage(42, 12, 20, &ForwardMessageOpts{InGameShare: true, ReplyTo: reply, AllowPaidBroadcast: true, SuggestedPostInfo: &InputSuggestedPostInfo{}, ReplaceVideoStartTimestamp: true, NewVideoStartTimestamp: 7})
			case "send forwarded":
				_, err = c.SendForwarded(42, 12, 20, &SendForwardedOpts{InGameShare: true, ReplyTo: reply, AllowPaidBroadcast: true, SuggestedPostInfo: &InputSuggestedPostInfo{}, ReplaceVideoStartTimestamp: true, NewVideoStartTimestamp: 7})
			case "copy":
				_, err = c.CopyMessage(42, 12, 20, &CopyMessageOpts{InGameShare: true, ReplyTo: reply, AllowPaidBroadcast: true, SuggestedPostInfo: &InputSuggestedPostInfo{}, ReplaceVideoStartTimestamp: true, NewVideoStartTimestamp: 7, NewCaption: "ignored"})
			}
			if err != nil || len(*requests) != 1 {
				t.Fatalf("send = %v, requests = %+v", err, *requests)
			}
			req := (*requests)[0]
			content := helperRequestContent(t, req).(*InputMessageForwarded)
			if content.FromChatId != 12 || content.MessageId != 20 || !content.InGameShare || !content.ReplaceVideoStartTimestamp || content.NewVideoStartTimestamp != 7 {
				t.Fatal("lost forwarded content options")
			}
			if helperReplyId(t, req) != 88 || !req.Options.AllowPaidBroadcast || req.Options.SuggestedPostInfo == nil {
				t.Fatal("lost delivery options")
			}
			if name == "copy" && (content.CopyOptions == nil || !content.CopyOptions.SendCopy || content.CopyOptions.NewCaption != nil) {
				t.Fatal("copy processed an unreplaced caption")
			}
		})
	}
}

func TestLivePhotoPreservesMediaCaptionAndReplyOptions(t *testing.T) {
	for _, reply := range []bool{false, true} {
		t.Run(map[bool]string{false: "send", true: "reply"}[reply], func(t *testing.T) {
			c := newTestClient(t)
			requests := captureHelperRequests(t, c)
			photo := &InputFileLocal{Path: "photo.jpg"}
			video := &InputFileLocal{Path: "live.mp4"}
			opts := &SendPhotoOpts{
				Caption: "Hello 😀", CaptionEntities: []TextEntity{{Offset: 6, Length: 2, Type: &TextEntityTypeBold{}}},
				Video: &InputFileLocal{Path: "previous.mp4"}, Width: 640, Height: 480,
				HasSpoiler: true, ShowCaptionAboveMedia: true, DisableNotification: true, ProtectContent: true,
			}
			before := *opts
			var err error
			if reply {
				_, err = (&Message{ChatId: 12, Id: 20}).ReplyLivePhoto(c, photo, video, opts)
			} else {
				_, err = c.SendLivePhoto(12, photo, video, opts)
			}
			if err != nil || len(*requests) != 1 {
				t.Fatalf("send = %v, requests: %+v", err, *requests)
			}
			req := (*requests)[0]
			content := helperRequestContent(t, req).(*InputMessagePhoto)
			if req.Type != "sendMessage" || req.ChatId != 12 || !reflect.DeepEqual(content.Photo.Photo, photo) || !reflect.DeepEqual(content.Photo.Video, video) {
				t.Fatalf("live photo files lost: %+v", content.Photo)
			}
			if content.Photo.Width != 640 || content.Photo.Height != 480 || !content.HasSpoiler || !content.ShowCaptionAboveMedia || !reflect.DeepEqual(content.Caption, &FormattedText{Text: opts.Caption, Entities: opts.CaptionEntities}) {
				t.Fatalf("photo options lost: %+v", content)
			}
			wantReply := int64(0)
			if reply {
				wantReply = 20
			}
			if helperReplyId(t, req) != wantReply || !req.Options.DisableNotification || !req.Options.ProtectContent {
				t.Fatal("reply/delivery options lost")
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("live photo mutated caller options")
			}
		})
	}
}

func TestLiveLocationPreservesNativeStateAndExplicitReply(t *testing.T) {
	for _, explicit := range []bool{false, true} {
		t.Run(map[bool]string{false: "default reply", true: "explicit reply"}[explicit], func(t *testing.T) {
			c := newTestClient(t)
			requests := captureHelperRequests(t, c)
			location := &LiveLocation{
				Location:   &Location{Latitude: 12.34, Longitude: 56.78, HorizontalAccuracy: 10},
				LivePeriod: 0x7fffffff, Heading: 180, ProximityAlertRadius: 100,
			}
			opts := &SendLocationOpts{LivePeriod: 60, Heading: 90, ProximityAlertRadius: 50, ProtectContent: true}
			wantReply := int64(20)
			if explicit {
				opts.ReplyTo = &InputMessageReplyToMessage{MessageId: 30}
				wantReply = 30
			}
			before := *opts
			if _, err := (&Message{ChatId: 12, Id: 20}).ReplyLiveLocation(c, location, opts); err != nil {
				t.Fatal(err)
			}
			req := (*requests)[0]
			content := helperRequestContent(t, req).(*InputMessageLiveLocation)
			if req.Type != "sendMessage" || req.ChatId != 12 || !reflect.DeepEqual(content.Location, location) || helperReplyId(t, req) != wantReply || !req.Options.ProtectContent {
				t.Fatalf("live location state/reply lost: %+v", req)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("live location mutated options")
			}
		})
	}
}

func TestRichSendUsesEphemeralRouting(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	_, err := c.SendRichMessageHTML(12, "<h1>Hello</h1>", &SendRichTextOpts{SendMessageOpts: SendMessageOpts{
		CallbackQueryID: 22, ReceiverUserID: 44, ReplaceCallbackQueryMessage: true,
		ProtectContent: true, OnlyPreview: true, SendingId: 7,
	}})
	if err != nil || len(*requests) != 1 {
		t.Fatalf("ephemeral send = %v, requests: %+v", err, *requests)
	}
	req := (*requests)[0]
	if req.Type != "sendEphemeralMessage" || req.ChatId != 12 || req.ReceiverUserId != 44 {
		t.Fatalf("ephemeral routing lost: %+v", req)
	}
	if req.CallbackQueryId != 22 || !req.ReplaceCallbackQueryMessage || !req.ProtectContent || !req.OnlyPreview || req.SendingId != 7 {
		t.Fatalf("ephemeral delivery options lost: %+v", req)
	}
	if _, ok := helperRequestContent(t, req).(*InputMessageRichMessage); !ok {
		t.Fatal("ephemeral rich content lost")
	}
}

func TestReplyHelpersPreserveExplicitReplyObjects(t *testing.T) {
	for _, name := range []string{"text", "photo", "voice", "location", "poll", "forwarded"} {
		t.Run(name, func(t *testing.T) {
			c := newTestClient(t)
			requests := captureHelperRequests(t, c)
			message := &Message{ChatId: 42, Id: 20}
			reply := &InputMessageReplyToMessage{MessageId: 99, Quote: &InputTextQuote{Text: &FormattedText{Text: "quote"}}}
			var err error
			switch name {
			case "text":
				_, err = message.ReplyText(c, "hello", &SendMessageOpts{ReplyTo: reply})
			case "photo":
				_, err = message.ReplyPhoto(c, &InputFileId{Id: 1}, &SendPhotoOpts{ReplyTo: reply})
			case "voice":
				_, err = message.ReplyVoice(c, &InputFileId{Id: 1}, &SendVoiceOpts{ReplyTo: reply})
			case "location":
				_, err = message.ReplyLocation(c, &Location{}, &SendLocationOpts{ReplyTo: reply})
			case "poll":
				_, err = message.ReplyPoll(c, "question", nil, &SendPollOpts{ReplyTo: reply})
			case "forwarded":
				_, err = message.ReplyForwarded(c, 42, 1, &SendForwardedOpts{ReplyTo: reply})
			}
			if err != nil || len(*requests) != 1 {
				t.Fatalf("reply = %v; requests: %+v", err, *requests)
			}
			got, err := UnmarshalInputMessageReplyTo((*requests)[0].ReplyTo)
			if err != nil || !reflect.DeepEqual(got, reply) || reply.MessageId != 99 {
				t.Fatalf("explicit reply object changed: (%v, %v)", got, err)
			}
		})
	}
}

func TestSendPrefersExplicitReplyAndPropagatesLoadFailure(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(map[bool]string{false: "loaded", true: "load failure"}[fail], func(t *testing.T) {
			c := newTestClient(t)
			c.config.LoadMessagesBeforeReply = true
			failure := &Error{Code: 400, Message: "reply unavailable"}
			var types []string
			c.sendRequest = func(_ int, data []byte) {
				req := decodeTestRequest(t, data)
				types = append(types, req.Type)
				var response TlObject = &Message{ChatId: 42, Id: 99}
				if req.Type == "getMessage" {
					if req.MessageID != 88 {
						t.Fatalf("loaded reply %d, want explicit 88", req.MessageID)
					}
					if fail {
						response = failure
					}
				} else {
					var send helperWireRequest
					if err := json.Unmarshal(data, &send); err != nil {
						t.Fatal(err)
					}
					if helperReplyId(t, send) != 88 {
						t.Fatal("explicit ReplyTo was overwritten")
					}
				}
				c.deliverResponse(req.Extra, response)
			}
			reply := &InputMessageReplyToMessage{MessageId: 88}
			opts := &SendMessageOpts{ReplyTo: reply, ReplyToMessageID: 77}
			before := *opts
			_, err := c.SendMessage(42, "hello", opts)
			want := []string{"getMessage", "sendMessage"}
			if fail {
				want = want[:1]
				if !errors.Is(err, failure) {
					t.Fatalf("error = %v", err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(types, want) || !reflect.DeepEqual(*opts, before) {
				t.Fatalf("requests/options: %v / %+v", types, opts)
			}
		})
	}
}
