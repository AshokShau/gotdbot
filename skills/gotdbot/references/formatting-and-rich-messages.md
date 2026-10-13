# Formatting and rich messages

Contents: GetFormattedText and ParseText; parse precedence; plain text; escaping; entities; rendering; rich sources; drafts.

Source anchors: `formatting_parse.go`, `formatting.go`, `message_formatting.go`, `send_helpers.go`, `edit_messages.go`, `gen_types.go`, `formatting_edge_test.go`, `rich_message_helpers_test.go`, `message_draft_test.go`.

## GetFormattedText and ParseText

These methods prepare or retrieve a `*FormattedText`, containing plain `Text` and `[]TextEntity`. Preparing formatted text does not send a message. Distinguish the Client method from the Message accessor with the same name:

| Method | Purpose | Empty mode / existing entities |
| --- | --- | --- |
| `Client.GetFormattedText(text string, entities []TextEntity, parseMode string) (*FormattedText, error)` | Prepare text using the same precedence as send/edit helpers. | Non-empty entities bypass parsing; otherwise an empty mode inherits `ClientOpts.ParseMode`. |
| `Client.ParseText(text string, parseMode string) (*FormattedText, error)` | Parse using the supplied mode, independently of the client default. | Empty or unrecognized mode returns literal text without entities. |
| `Client.ParseTextEntities(parseMode TextParseMode, text string) (*FormattedText, error)` | Native parser with a typed mode such as `TextParseModeHTML` or `TextParseModeMarkdown`. | Pass an appropriate native mode; argument order differs from ParseText. |
| `Message.GetFormattedText() (*FormattedText, error)` | Read formatted text already present in a `MessageText`. | Does not parse; returns `ErrNoFormattedText` for content without ordinary text. |
| `Message.GetFormattedCaption() (*FormattedText, error)` | Read a supported media caption. | Does not parse; returns `ErrNoFormattedText` if no formatted caption is present. |

Use GetFormattedText when the application accepts either explicit entities or a parse mode and should honor the client default. Use ParseText when choosing a specific mode or forcing literal text. Use the Message accessors when inspecting received content or preserving its existing entities.

```go
package recipes

import "github.com/AshokShau/gotdbot"

// If entities are non-empty, markup is not parsed. Otherwise this inherits
// the client's default parse mode because the mode argument is empty.
func PrepareWithDefaults(c *gotdbot.Client, text string, entities []gotdbot.TextEntity) (*gotdbot.FormattedText, error) {
	return c.GetFormattedText(text, entities, "")
}

// Parse HTML explicitly, regardless of the client's default parse mode.
// The result contains text "Hello" and a bold entity.
func ParseHTML(c *gotdbot.Client) (*gotdbot.FormattedText, error) {
	return c.ParseText("<b>Hello</b>", gotdbot.ParseModeHTML)
}

// ParseModeNone is empty: ParseText leaves any markup in text literal.
// Send native content so the send helper does not reapply a client default.
func SendLiteralText(c *gotdbot.Client, chatID int64, text string) (*gotdbot.Message, error) {
	formatted, err := c.ParseText(text, gotdbot.ParseModeNone)
	if err != nil {
		return nil, err
	}
	return c.SendMessageContent(chatID, &gotdbot.InputMessageText{Text: formatted}, nil)
}

// GetFormattedText is also useful when constructing native media content.
func SendPhotoWithParsedCaption(c *gotdbot.Client, chatID int64, path string) (*gotdbot.Message, error) {
	caption, err := c.GetFormattedText("<b>Photo</b>", nil, gotdbot.ParseModeHTML)
	if err != nil {
		return nil, err
	}
	return c.SendMessageContent(chatID, &gotdbot.InputMessagePhoto{
		Photo:   &gotdbot.InputPhoto{Photo: &gotdbot.InputFileLocal{Path: path}},
		Caption: caption,
	}, nil)
}
```

Recognized markup modes dispatch native ParseTextEntities requests, so initialize the client/receiver before parsing and propagate malformed-markup errors. An unrecognized mode silently produces plain text; validate user-configurable mode names if a typo should be an error. GetFormattedText with non-empty entities and ParseText with an empty mode do not dispatch parser requests.

GetFormattedText does not validate supplied entity ranges or deep-copy their slice. Preserve UTF-16 offsets and avoid modifying shared entities while another operation uses them. Reading Message.GetFormattedText likewise returns the existing formatted-text object; copy it if the application needs to modify it independently.

## Formatting precedence in helpers

`SendMessageOpts.Entities`, `SendPhotoOpts.CaptionEntities`, and corresponding edit/caption fields contain `[]TextEntity`. Helpers resolve formatting in this order:

1. If the supplied entities have nonzero length, return that text/entities pair without parsing markup.
2. Otherwise, use the non-empty helper `ParseMode` or inherit `ClientOpts.ParseMode`.
3. Recognized modes call native `ParseTextEntities`; unknown modes return ordinary plain text.

Modes are case-insensitive: HTML, Markdown (version 1), MarkdownV2 (version 2). Constants are `ParseModeHTML`, `ParseModeMarkdown`, `ParseModeMarkdownV2`, and `ParseModeNone`. `ParseModeNone` is an empty string, so in helpers it still inherits the client's default; it is not a forced-plain override. A zero-length entities slice also does not suppress that default.

For literal text independent of client defaults, send `InputMessageText{Text: &FormattedText{Text: text}}` through `SendMessageContent`/`ReplyContent`. Alternatively, `ParseText(text, "")` itself returns plain formatted text; `GetFormattedText` performs default resolution. ParseText's recognized markup modes dispatch a request through the client, so initialize the client/receiver before using them. Pure rendering with `FormattedText.Format` does not need TDLib.

## Escape interpolated text

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"html"
)

func Greet(c *gotdbot.Client, msg *gotdbot.Message, name string) (*gotdbot.Message, error) {
	return msg.ReplyText(c, "Hello <b>"+html.EscapeString(name)+"</b>", &gotdbot.SendMessageOpts{
		ParseMode: gotdbot.ParseModeHTML,
	})
}
```

For MarkdownV2, rendering a plain `FormattedText` with `.Format(ParseModeMarkdownV2)` escapes supported special characters. Do not treat arbitrary user input as markup. `User.Mention(HTML)` returns an escaped `tg://user?id=...` link; other modes use the library's MarkdownV2 escaping, so choose HTML or MarkdownV2 consistently when sending it. `Message.Mention(c, mode)` fetches the sender user/chat; user senders get links, chat senders get their title.

## Entities and UTF-16

TDLib entity offsets/lengths are in UTF-16 code units, not UTF-8 byte indexes or Go rune counts. Astral characters such as many emoji use two UTF-16 units. Preserve source entities when forwarding formatted text; compute UTF-16 lengths when building new entities.

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"unicode/utf16"
)

func BoldText(text string) *gotdbot.FormattedText {
	if text == "" {
		return &gotdbot.FormattedText{}
	}
	return &gotdbot.FormattedText{
		Text: text,
		Entities: []gotdbot.TextEntity{{
			Offset: 0,
			Length: int32(len(utf16.Encode([]rune(text)))),
			Type:   &gotdbot.TextEntityTypeBold{},
		}},
	}
}
```

`FormattedText.Format(mode)` renders entities back into markup and escapes text. `Message.OriginalHTML`, `OriginalMD`, `OriginalMDV2`, and their `OriginalCaption*` counterparts use it. These are rendered representations, not byte-for-byte recovery of the original markup. Native limitations still apply when sending the rendered markup back to Telegram.

## Rich messages are a distinct content model

Rich messages use `InputRichMessage` and `InputMessageRichMessage`, not `FormattedText` with a richer parse mode. Source types are:

| Source | Structure | Convenience route |
| --- | --- | --- |
| `RichMessageSourceMarkdown` | Text plus media references | `SendRichMessageMarkdown`, `ReplyRichMessageMarkdown` |
| `RichMessageSourceHtml` | Rich HTML plus media references | `SendRichMessageHTML`, `ReplyRichMessageHTML` |
| `RichMessageSourceBlocks` | `[]InputPageBlock` | `SendRichMessage` or `ReplyRichMessage` |

Rich text helpers are documented for bots only. Use native type comments for supported elements/block/media types instead of ordinary Telegram HTML or Markdown rules. `InputRichMessage.IsRtl` controls direction; `DetectAutomaticBlocks` enables URL/email and other automatic block detection.

For structured content, a paragraph block takes native `RichText`, not `FormattedText`:

```go
package recipes

import "github.com/AshokShau/gotdbot"

func SendRichBlocks(c *gotdbot.Client, chatID int64, text string) (*gotdbot.Message, error) {
	return c.SendRichMessage(chatID, &gotdbot.InputRichMessage{
		Source: &gotdbot.RichMessageSourceBlocks{
			Blocks: []gotdbot.InputPageBlock{
				&gotdbot.InputPageBlockParagraph{Text: &gotdbot.RichTextPlain{Text: text}},
			},
		},
	}, nil)
}
```

```go
package recipes

import "github.com/AshokShau/gotdbot"

func PublishRich(c *gotdbot.Client, chatID int64) (*gotdbot.Message, error) {
	sent, err := c.SendRichMessageMarkdown(chatID, "# Update\n\nFirst version.", &gotdbot.SendRichTextOpts{
		SendMessageOpts:       gotdbot.SendMessageOpts{DisableNotification: true},
		DetectAutomaticBlocks: true,
	})
	if err != nil {
		return nil, err
	}
	return sent.EditRichMessageHTML(c, "<h1>Update</h1><p>Revised version.</p>", &gotdbot.EditRichTextOpts{
		DetectAutomaticBlocks: true,
	})
}
```

`SendRichTextOpts` embeds `SendMessageOpts`; promoted delivery fields must be set in the embedded struct in composite literals, as shown. `Media` is `[]InputRichMessageMedia`, supplying referenced media IDs in source text. `EditRichTextOpts` has its own media/direction/detection/markup fields.

Client edits include `EditRichMessage`, `EditRichMessageMarkdown`, `EditRichMessageHTML`, their `EditInlineRichMessage*` counterparts, and `EditEphemeralRichMessage*`. Message, callback, business callback, and chosen-inline-result helpers reuse the correct identifier/connection. Inline edits require already-uploaded files. Do not substitute a local file upload into those requests.

## Streamed drafts

Drafts are partial bot responses. `SendMessageDraft(chatID, draftID, text, *SendMessageDraftOpts)` uses ordinary formatting. Reuse a nonzero draft ID for updates to the same response and send the final message normally. Allocate IDs per active response to avoid overlapping streams.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func StreamReply(c *gotdbot.Client, chatID, draftID int64, partial, final string) (*gotdbot.Message, error) {
	if err := c.SendMessageDraft(chatID, draftID, partial, &gotdbot.SendMessageDraftOpts{
		CanStop: true, KeepOnStop: true,
	}); err != nil {
		return nil, err
	}
	return c.SendMessage(chatID, final, nil)
}
```

This example assumes a plain client default and receives an application-generated draft ID. Adapt the final send's topic/reply settings to the original destination. Draft `ForumTopicId` is an int32; ordinary send routing uses a `MessageTopic` union, so the structs differ.

Rich streams use `SendRichMessageDraftMarkdown`/`SendRichMessageDraftHTML` with `SendRichMessageDraftTextOpts`, which embeds native `SendRichMessageDraftOpts` and adds media/direction/detection. Native `SendRichMessageDraft` accepts an `InputRichMessage`. All draft media must already be uploaded. `CanStop`/`KeepOnStop` configure native generation controls; wire relevant stop updates into the application's generator cancellation if implementing streaming. The helper does not cancel an external model/task by itself.
