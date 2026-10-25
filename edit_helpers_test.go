package gotdbot

import (
	"errors"
	"reflect"
	"testing"
)

func TestMediaEditsPreserveContentAndRouteByUpdate(t *testing.T) {
	file := &InputFileId{Id: 1}
	entities := []TextEntity{{Offset: 3, Length: 2, Type: &TextEntityTypeBold{}}}
	markup := &ReplyMarkupInlineKeyboard{}
	message := &Message{ChatId: 42, Id: 20}
	callback := &UpdateNewCallbackQuery{ChatId: 42, MessageId: 20}
	inline := &UpdateNewInlineCallbackQuery{InlineMessageId: "inline-id"}
	business := &UpdateNewBusinessCallbackQuery{ConnectionId: "business-id", Message: &BusinessMessage{Message: message}}
	chosen := &UpdateNewChosenInlineResult{InlineMessageId: "inline-id"}
	for _, tc := range []struct {
		name, contentType string
		invoke            func(*testing.T, *Client, string) error
	}{
		{name: "photo", contentType: "inputMessagePhoto", invoke: func(t *testing.T, c *Client, target string) error {
			opts := &EditMessagePhotoOpts{Caption: "hi 😀", CaptionEntities: entities, ReplyMarkup: markup, Width: 640, Height: 480, AddedStickerFileIds: []int32{7}, Video: &InputFileId{Id: 2}, HasSpoiler: true, ShowCaptionAboveMedia: true}
			before := *opts
			var err error
			switch target {
			case "message":
				_, err = message.EditPhoto(c, file, opts)
			case "callback":
				_, err = callback.EditMessagePhoto(c, file, opts)
			case "inline":
				err = inline.EditMessagePhoto(c, file, opts)
			case "business":
				_, err = business.EditMessagePhoto(c, file, opts)
			case "chosen":
				err = chosen.EditPhoto(c, file, opts)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("media edit mutated caller options")
			}
			return err
		}},

		{name: "video", contentType: "inputMessageVideo", invoke: func(t *testing.T, c *Client, target string) error {
			opts := &EditMessageVideoOpts{Caption: "hi 😀", CaptionEntities: entities, ReplyMarkup: markup, Width: 640, Height: 480, Duration: 12, StartTimestamp: 3, SupportsStreaming: true, Cover: &InputFileId{Id: 2}, HasSpoiler: true, ShowCaptionAboveMedia: true}
			before := *opts
			var err error
			switch target {
			case "message":
				_, err = message.EditVideo(c, file, opts)
			case "callback":
				_, err = callback.EditMessageVideo(c, file, opts)
			case "inline":
				err = inline.EditMessageVideo(c, file, opts)
			case "business":
				_, err = business.EditMessageVideo(c, file, opts)
			case "chosen":
				err = chosen.EditVideo(c, file, opts)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("media edit mutated caller options")
			}
			return err
		}},

		{name: "animation", contentType: "inputMessageAnimation", invoke: func(t *testing.T, c *Client, target string) error {
			opts := &EditMessageAnimationOpts{Caption: "hi 😀", CaptionEntities: entities, ReplyMarkup: markup, Width: 640, Height: 480, Duration: 12, AddedStickerFileIds: []int32{7}, HasSpoiler: true, ShowCaptionAboveMedia: true}
			before := *opts
			var err error
			switch target {
			case "message":
				_, err = message.EditAnimation(c, file, opts)
			case "callback":
				_, err = callback.EditMessageAnimation(c, file, opts)
			case "inline":
				err = inline.EditMessageAnimation(c, file, opts)
			case "business":
				_, err = business.EditMessageAnimation(c, file, opts)
			case "chosen":
				err = chosen.EditAnimation(c, file, opts)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("media edit mutated caller options")
			}
			return err
		}},

		{name: "audio", contentType: "inputMessageAudio", invoke: func(t *testing.T, c *Client, target string) error {
			opts := &EditMessageAudioOpts{Caption: "hi 😀", CaptionEntities: entities, ReplyMarkup: markup, Duration: 12, Title: "Title", Performer: "Artist", AlbumCoverThumbnail: &InputThumbnail{Thumbnail: &InputFileId{Id: 2}, Width: 40, Height: 40}}
			before := *opts
			var err error
			switch target {
			case "message":
				_, err = message.EditAudio(c, file, opts)
			case "callback":
				_, err = callback.EditMessageAudio(c, file, opts)
			case "inline":
				err = inline.EditMessageAudio(c, file, opts)
			case "business":
				_, err = business.EditMessageAudio(c, file, opts)
			case "chosen":
				err = chosen.EditAudio(c, file, opts)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("media edit mutated caller options")
			}
			return err
		}},

		{name: "document", contentType: "inputMessageDocument", invoke: func(t *testing.T, c *Client, target string) error {
			opts := &EditMessageDocumentOpts{Caption: "hi 😀", CaptionEntities: entities, ReplyMarkup: markup, DisableContentTypeDetection: true, Thumbnail: &InputThumbnail{Thumbnail: &InputFileId{Id: 2}, Width: 40, Height: 40}}
			before := *opts
			var err error
			switch target {
			case "message":
				_, err = message.EditDocument(c, file, opts)
			case "callback":
				_, err = callback.EditMessageDocument(c, file, opts)
			case "inline":
				err = inline.EditMessageDocument(c, file, opts)
			case "business":
				_, err = business.EditMessageDocument(c, file, opts)
			case "chosen":
				err = chosen.EditDocument(c, file, opts)
			}
			if !reflect.DeepEqual(*opts, before) {
				t.Fatal("media edit mutated caller options")
			}
			return err
		}},
	} {
		for _, target := range []string{"message", "callback", "inline", "business", "chosen"} {
			t.Run(tc.name+"/"+target, func(t *testing.T) {
				c := newTestClient(t)
				c.config.ParseMode = "HTML"
				requests := captureHelperRequests(t, c)
				if err := tc.invoke(t, c, target); err != nil {
					t.Fatal(err)
				}
				if len(*requests) != 1 {
					t.Fatalf("requests = %+v", *requests)
				}
				req := (*requests)[0]
				expectedType := "editMessageMedia"
				if target == "inline" || target == "chosen" {
					expectedType = "editInlineMessageMedia"
					if req.InlineMessageId != "inline-id" {
						t.Fatal("lost inline ID")
					}
				} else {
					if req.ChatId != 42 || req.MessageId != 20 {
						t.Fatal("lost chat/message IDs")
					}
				}
				if target == "business" {
					expectedType = "editBusinessMessageMedia"
					if req.BusinessConnectionId != "business-id" {
						t.Fatal("lost business connection")
					}
				}
				if req.Type != expectedType {
					t.Fatalf("endpoint = %s, want %s", req.Type, expectedType)
				}
				if !reflect.DeepEqual(helperRequestMarkup(t, req), markup) {
					t.Fatal("lost keyboard")
				}
				content := helperRequestContent(t, req)
				if content.GetType() != tc.contentType {
					t.Fatalf("content = %T", content)
				}
				var caption *FormattedText
				switch v := content.(type) {
				case *InputMessagePhoto:
					caption = v.Caption
					if !reflect.DeepEqual(v.Photo.Photo, file) || v.Photo.Width != 640 || v.Photo.Height != 480 || v.Photo.Video == nil || !v.HasSpoiler || !v.ShowCaptionAboveMedia {
						t.Fatal("lost photo options")
					}
				case *InputMessageVideo:
					caption = v.Caption
					if !reflect.DeepEqual(v.Video.Video, file) || v.Video.Duration != 12 || v.Video.StartTimestamp != 3 || !v.Video.SupportsStreaming || v.Video.Cover == nil || !v.HasSpoiler {
						t.Fatal("lost video options")
					}
				case *InputMessageAnimation:
					caption = v.Caption
					if !reflect.DeepEqual(v.Animation.Animation, file) || v.Animation.Duration != 12 || !reflect.DeepEqual(v.Animation.AddedStickerFileIds, []int32{7}) || !v.HasSpoiler {
						t.Fatal("lost animation options")
					}
				case *InputMessageAudio:
					caption = v.Caption
					if !reflect.DeepEqual(v.Audio.Audio, file) || v.Audio.Title != "Title" || v.Audio.Performer != "Artist" || v.Audio.AlbumCoverThumbnail == nil {
						t.Fatal("lost audio options")
					}
				case *InputMessageDocument:
					caption = v.Caption
					if !reflect.DeepEqual(v.Document.Document, file) || !v.Document.DisableContentTypeDetection || v.Document.Thumbnail == nil {
						t.Fatal("lost document options")
					}
				}
				if caption == nil || caption.Text != "hi 😀" || !reflect.DeepEqual(caption.Entities, entities) {
					t.Fatal("lost caption/entities")
				}
			})
		}
	}
}

func TestContentEditsAcceptValuesAndRejectUnsupportedInputs(t *testing.T) {
	for _, tc := range []struct {
		name, endpoint string
		content        InputMessageContent
	}{
		{"text", "Text", InputMessageText{Text: &FormattedText{Text: "new"}}},
		{"rich", "Text", InputMessageRichMessage{Message: &InputRichMessage{Source: &RichMessageSourceMarkdown{Text: "new"}}}},
		{"photo", "Media", InputMessagePhoto{Photo: &InputPhoto{Photo: &InputFileId{Id: 1}}}},
		{"live location", "LiveLocation", InputMessageLiveLocation{Location: &LiveLocation{LivePeriod: 60}}},
		{"checklist", "Checklist", InputMessageChecklist{Checklist: &InputChecklist{Title: &FormattedText{Text: "Tasks"}}}},
	} {
		for _, target := range []string{"chat", "inline", "business"} {
			t.Run(tc.name+"/"+target, func(t *testing.T) {
				c := newTestClient(t)
				requests := captureHelperRequests(t, c)
				markup := &ReplyMarkupInlineKeyboard{}
				var err error
				var prefix string
				switch target {
				case "chat":
					_, err = c.EditContent(42, 20, tc.content, &EditMessageMediaOpts{ReplyMarkup: markup})
					prefix = "editMessage"
				case "inline":
					err = c.EditInlineContent("inline-id", tc.content, &EditInlineMessageMediaOpts{ReplyMarkup: markup})
					prefix = "editInlineMessage"
				case "business":
					_, err = c.EditBusinessContent("business-id", 42, 20, tc.content, &EditBusinessMessageMediaOpts{ReplyMarkup: markup})
					prefix = "editBusinessMessage"
				}
				if target == "inline" && tc.endpoint == "Checklist" {
					if err == nil || len(*requests) != 0 {
						t.Fatal("inline checklist accepted")
					}
					return
				}
				if err != nil || len(*requests) != 1 {
					t.Fatalf("edit = %v; requests = %+v", err, *requests)
				}
				if (*requests)[0].Type != prefix+tc.endpoint {
					t.Fatalf("wrong endpoint: %+v", (*requests)[0])
				}
				if !reflect.DeepEqual(helperRequestMarkup(t, (*requests)[0]), markup) {
					t.Fatal("lost reply markup")
				}
			})
		}
	}
	for _, content := range []InputMessageContent{nil, (*InputMessagePhoto)(nil), (*InputMessageLiveLocation)(nil), (*InputMessageChecklist)(nil), InputMessageDice{Emoji: "🎲"}} {
		c := newTestClient(t)
		c.sendRequest = func(int, []byte) { t.Fatal("unsupported content sent a request") }
		if _, err := c.EditContent(42, 20, content, nil); err == nil {
			t.Fatal("chat accepted invalid edit")
		}
		if err := c.EditInlineContent("inline-id", content, nil); err == nil {
			t.Fatal("inline accepted invalid edit")
		}
		if _, err := c.EditBusinessContent("business-id", 42, 20, content, nil); err == nil {
			t.Fatal("business accepted invalid edit")
		}
	}
}

func TestMediaEditParseFailureStopsBeforeEdit(t *testing.T) {
	c := newTestClient(t)
	failure := &Error{Code: 400, Message: "invalid formatting"}
	var types []string
	c.sendRequest = func(_ int, data []byte) {
		req := decodeTestRequest(t, data)
		types = append(types, req.Type)
		c.deliverResponse(req.Extra, failure)
	}
	_, err := c.EditMessagePhoto(42, 20, &InputFileId{Id: 1}, &EditMessagePhotoOpts{Caption: "<bad>", ParseMode: "HTML"})
	if !errors.Is(err, failure) || !reflect.DeepEqual(types, []string{"parseTextEntities"}) {
		t.Fatalf("error/requests = %v / %v", err, types)
	}
}
