# Messages and media

Contents: identifiers; sending and replies; native content; media/files; editing; transfers; polls; actions.

Source anchors: `send_helpers.go`, `message_reply.go`, `message_edit.go`, `edit_messages.go`, `message_content.go`, `message_media.go`, `message_actions.go`, `message_transfer.go`, `chat_helpers.go`, `gen_methods.go`, `gen_options.go`.

## Message model

`Message.ChatId` and `Message.Id` are TDLib identifiers. Keep them together with the account that produced them. `SenderId` is a union of `MessageSenderUser` and `MessageSenderChat`; use a type switch before assuming a user. `SenderID()` returns either kind's ID; a numeric result alone does not establish sender type.

`Text()` reads ordinary message text. `Caption()` reads supported media captions. `GetText()` prefers a non-empty caption, then text. `Empty()` uses trimmed GetText, so an image without a caption is considered empty text. It does not mean the message has no content. `GetFormattedText()` retrieves text only; `GetFormattedCaption()` retrieves a caption. Missing formatted content returns `ErrNoFormattedText`.

`ChatID`, `ReplyToMessageID`, `GetEntities`, `Entities`, `CaptionEntities`, `StartsWith`, `EndsWith`, `TrimmedText`, `IsPrivate`, `IsGroup`, `IsSupergroup`, `IsChannel`, `IsSecret`, `IsOutgoingMessage`, and `IsEditedMessage` support inspection. Use `GetChat` and its `Chat.Type` when authoritative native chat-type information is required.

## Sending, replying, and options

Use `Client.SendMessage(chatID, text, *SendMessageOpts)` for text. `Message.ReplyText(client, text, opts)` fills its chat/reply destination. `Chat.SendMessage(client, text, opts)` uses `Chat.Id`. Optional pointers can be nil.

Replies shallow-copy options and preserve an explicit `ReplyTo`; otherwise they use the source message's ID. In direct sends, explicit `ReplyTo` takes precedence over `ReplyToMessageID`. Use `InputMessageReplyToMessage` for quotes/advanced reply targets and `MessageTopicForum` for forum routing. Do not assume the reply helper automatically selects a forum topic; specify `TopicId` when required.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func ReplyInTopic(c *gotdbot.Client, msg *gotdbot.Message, topicID int32) (*gotdbot.Message, error) {
	return msg.ReplyText(c, "Received", &gotdbot.SendMessageOpts{
		TopicId:             &gotdbot.MessageTopicForum{ForumTopicId: topicID},
		DisableNotification: true,
	})
}
```

Send helpers expose delivery fields such as `DisableNotification`, `ProtectContent`, `SchedulingState`, `EffectId`, `SendingId`, `OnlyPreview`, `FromBackground`, and supported paid-message settings. Read the specific option struct; media helpers add content metadata, captions, thumbnails, and spoilers. Do not enable paid settings as generic defaults.

Setting a nonzero `ReceiverUserID` or positive `CallbackQueryID` selects an ephemeral send. `ReplaceCallbackQueryMessage` controls replacement in that flow; both identifiers are passed to the native request when supplied. Inspect `SendEphemeralMessage`/`EditEphemeralMessage` for native restrictions before using this route; it is not a general private-message broadcast mechanism.

For native content use `SendMessageContent` or `Message.ReplyContent`. Its `SendMessageContentOpts` nests `Options: &MessageSendOptions{...}` rather than flattening delivery settings. This is also the route for plain formatted text when a client-wide parse mode should not apply.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func SendNative(c *gotdbot.Client, chatID int64, text string) (*gotdbot.Message, error) {
	return c.SendMessageContent(chatID, &gotdbot.InputMessageText{
		Text: &gotdbot.FormattedText{Text: text},
	}, &gotdbot.SendMessageContentOpts{
		Options: &gotdbot.MessageSendOptions{DisableNotification: true},
	})
}
```

## Media and files

| Content | Client helpers | Notes |
| --- | --- | --- |
| Photo/live photo | `SendPhoto`, `SendLivePhoto` | Live photo takes photo and video files. |
| Video/animation/round video | `SendVideo`, `SendAnimation`, `SendVideoNote` | Distinct input content/options. |
| Audio/voice/document/sticker | `SendAudio`, `SendVoice`, `SendDocument`, `SendSticker` | Voice is not an audio music message. |
| Albums | `SendMediaGroup`, native `SendMessageAlbum` | `[]InputMessageContent`, 2–10 supported items; consult native constraints. |
| Geo/contact | `SendLocation`, `SendLiveLocation`, `SendVenue`, `SendContact` | Live location takes native `*LiveLocation`. |
| Poll/quiz/checklist | `SendPoll`, `SendQuiz`, `SendChecklist` | Structured options, not `[]string`. |
| Dice/game/story | `SendDice`, `SendStakeDice`, `SendGame`, `SendStory` | Read native account/feature constraints. |
| Payments | `SendInvoice`, `SendPaidMedia` | Preserve payment units/payloads from the native model. |

Most senders have `Message.Reply*` counterparts; inspect `message_reply.go` or `go doc` for the exact name. Chat convenience methods cover a subset and generated native helpers; do not assume every Client method has a Chat wrapper.

- `InputFileLocal{Path: ...}` uploads a local file available to the running process.
- `InputFileId{Id: ...}` reuses a TDLib file ID associated with that client.
- `InputFileRemote{Id: ...}` identifies a remote file still accessible to the account and known to TDLib. Do not assume an arbitrary URL is a valid remote identifier.
- `InputFileGenerated` participates in TDLib file-generation updates. It does not automatically run an arbitrary converter; implement the native generation lifecycle if chosen.
- Files use native integer IDs. `Remote.Id` is a remote string ID and `Remote.UniqueId` identifies content; the unique ID is not an upload/download handle.

```go
package recipes

import (
	"fmt"
	"github.com/AshokShau/gotdbot"
)

func SendAndDownload(c *gotdbot.Client, chatID int64, path string) (string, error) {
	sent, err := c.SendPhoto(chatID, &gotdbot.InputFileLocal{Path: path}, &gotdbot.SendPhotoOpts{
		Caption: "<b>A photo</b>", ParseMode: gotdbot.ParseModeHTML,
	})
	if err != nil {
		return "", err
	}
	file, err := sent.Download(c, 16, &gotdbot.DownloadFileOpts{Synchronous: true})
	if err != nil {
		return "", err
	}
	if file.Local == nil || !file.Local.IsDownloadingCompleted {
		return "", fmt.Errorf("download has not completed")
	}
	return file.Local.Path, nil
}
```

`GetFile()` returns the primary supported attachment, including the largest photo size by area; unsupported content returns nil. `Download` returns `ErrNoMedia` for no supported file and uses the original client. A default download may be asynchronous; use `Synchronous: true` or `OnUpdateFile` to determine completion. `RemoteFileID`, `RemoteUniqueFileID`, `RemoteFileSize`, and `RemoteDuration` provide convenient metadata. Rich/paid/multi-file content may require walking its native structure rather than GetFile.

## Albums

```go
package recipes

import "github.com/AshokShau/gotdbot"

func SendAlbum(c *gotdbot.Client, chatID int64, first, second string) (*gotdbot.Messages, error) {
	return c.SendMediaGroup(chatID, []gotdbot.InputMessageContent{
		&gotdbot.InputMessagePhoto{Photo: &gotdbot.InputPhoto{Photo: &gotdbot.InputFileLocal{Path: first}}},
		&gotdbot.InputMessagePhoto{Photo: &gotdbot.InputPhoto{Photo: &gotdbot.InputFileLocal{Path: second}}},
	}, &gotdbot.SendMessageAlbumOpts{
		Options: &gotdbot.MessageSendOptions{DisableNotification: true},
	})
}
```

Albums return `*Messages`, whose `Messages` field is a slice of message values. This schema uses nested `InputPhoto` for native photo content; the high-level `SendPhoto` helper constructs it for you. TDLib restricts albums to supported audio/document/photo/video combinations. Inspect the native method for grouping rules; parse any captions into `FormattedText` before constructing album items.

## Editing and transfers

Message methods take the client first and reuse chat/message IDs:

- `EditText`, `EditCaption`, `EditReplyMarkup`, and `EditContent` dispatch appropriate edits.
- `EditPhoto`, `EditVideo`, `EditAnimation`, `EditAudio`, and `EditDocument` replace editable media with matching `EditMessage*Opts`.
- `StopLiveLocation` ends an active live location.
- Client methods use `EditMessageText`/`EditMessageCaption` convenience names; inline text/caption convenience names are `EditInlineText`/`EditInlineCaption`.
- Inline media edits use `EditInlineMessagePhoto` and siblings and require previously uploaded files. They return error only, not a chat Message.
- Native text/caption edit names are `EditMessageTextContent` and `EditMessageCaptionText`.

Some helpers fetch a message first when `UseMessageDatabase` is explicitly false. Handle that lookup's error rather than assuming edit methods always issue exactly one request.

`CopyMessage(chatID, fromChatID, messageID, *CopyMessageOpts)` constructs copied content; `ForwardMessage` and native `ForwardMessages` preserve forwarding by default. `Message.Copy`/`Forward` accept the destination ID. `CopyMessages` wraps the native multi-message forward-copy path and leaves caller-owned options unchanged. Forward/copy behavior remains subject to content protection and native account constraints.

## Polls and quizzes

`SendPoll` defaults to regular polls with optional anonymity/multiple answers. `SendQuiz` takes zero-based `[]int32` correct-answer IDs. Questions/options are native formatted structures; use `InputPollOption{Text: &FormattedText{Text: ...}}`.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func SendQuiz(c *gotdbot.Client, chatID int64) (*gotdbot.Message, error) {
	return c.SendQuiz(chatID, "Which answer is four?", []gotdbot.InputPollOption{
		{Text: &gotdbot.FormattedText{Text: "Three"}},
		{Text: &gotdbot.FormattedText{Text: "Four"}},
	}, []int32{1}, nil)
}
```

## Message actions

`Delete(c, &DeleteMessagesOpts{Revoke: true})` requests deletion for everyone. Nil does not set Revoke. `Pin`, `Unpin`, `GetLink`, `GetMessageLink`, `GetMessageProperties`, `GetRepliedMessage`, `GetChat`, `GetUser`, and `Mention` cover common operations. `GetUser` requires a user sender; handle sender-chat messages separately.

`React`, `ReactEmoji`, and `ReactCustomEmoji` use native reaction settings. `Action` takes `*SendChatActionOpts`; set its native `Action` field for typing/upload indicators. Sender moderation methods `Ban`, `Kick`, `Restrict`, and `SetChatMemberStatus` operate on the actual `SenderId`, including chat senders. See the administration guide before choosing those actions.
