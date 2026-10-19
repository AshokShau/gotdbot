# GotdBot

A Go library for building Telegram bots and user clients with
[TDLib](https://github.com/tdlib/td).

GotdBot combines typed TDLib bindings with message helpers, commands, filters,
and update handlers. TDLib manages authentication, connections, and file
transfers. The library provides Bot API-style conveniences alongside access to
native TDLib methods and types.

[Go reference](https://pkg.go.dev/github.com/AshokShau/gotdbot) ·
[Examples](./examples) ·
[AI coding skill](./skills/gotdbot/SKILL.md) ·
[Issues](https://github.com/AshokShau/gotdbot/issues)

## Features

- Typed bindings for TDLib methods, updates, and message content.
- Sending, replying, editing, forwarding, and downloading messages and media.
- HTML, Markdown, entities, rich messages, and streamed message drafts.
- Commands, composable filters, typed update handlers, and handler groups.
- Concurrent update processing, request cancellation, and panic recovery.
- Bot and user authentication, persistent sessions, and multiple accounts.
- Runtime loading of TDLib through purego.

## Installation

```sh
go get github.com/AshokShau/gotdbot
```

You need **Go 1.26 or later** and a **TDLib 1.8.68** shared library matching the
bindings. Get a Telegram API ID and hash from [my.telegram.org](https://my.telegram.org),
and a bot token from [@BotFather](https://t.me/BotFather) for a bot account.

[Build TDLib](https://github.com/tdlib/td#building), or download the matching
prebuilt library for your platform:

```sh
go run github.com/AshokShau/gotdbot/scripts/tools
```

Set `ClientOpts.LibraryPath` to the shared library's path. Each account needs its
own database directory.

## Quick start

Set these environment variables:

| Variable | Value |
| --- | --- |
| `TELEGRAM_API_ID` | Your Telegram API ID |
| `TELEGRAM_API_HASH` | Your Telegram API hash |
| `TELEGRAM_BOT_TOKEN` | Your bot token |
| `TDLIB_LIBRARY` | Path to the TDLib shared library |

```go
package main

import (
    "log"
    "os"
    "strconv"

    "github.com/AshokShau/gotdbot"
)

func main() {
    apiID, err := strconv.ParseInt(os.Getenv("TELEGRAM_API_ID"), 10, 32)
    if err != nil || apiID <= 0 {
        log.Fatal("Set TELEGRAM_API_ID to your Telegram API ID")
    }

    apiHash := os.Getenv("TELEGRAM_API_HASH")
    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if apiHash == "" || botToken == "" {
        log.Fatal("Set TELEGRAM_API_HASH and TELEGRAM_BOT_TOKEN")
    }

    bot, err := gotdbot.NewClient(
        int32(apiID),
        apiHash,
        botToken,
        &gotdbot.ClientOpts{
            LibraryPath:       os.Getenv("TDLIB_LIBRARY"),
            DatabaseDirectory: "./data/bot",
            FilesDirectory:    "./data/files",
        },
    )
    if err != nil {
        log.Fatal(err)
    }
    defer bot.Close()

    bot.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
        _, err := msg.ReplyText(c, "Hello from GotdBot!", nil)
        return err
    })

    if err := bot.Start(); err != nil {
        log.Fatal(err)
    }
    bot.Idle()
}
```

Register handlers before calling `Start`. `Start` waits for authentication;
`Idle` keeps the client running until a termination signal arrives.
After `Start` succeeds, `bot.Me` contains the cached account identity. Use it
when you need the account's ID or username without another `GetMe` request.
For a user account, pass a phone number instead of a bot token to `NewClient`.

## Messages

Use IDs returned by TDLib, such as `msg.ChatId` and `msg.Id`. Client methods take
a target chat ID; `Message.Reply*` methods reply in the message's chat, and
`Chat.Send*` methods use that chat as the destination. Optional settings can be
`nil`.

```go
sent, err := bot.SendMessage(chatID, "<b>Hello</b>", &gotdbot.SendMessageOpts{
    ParseMode: "HTML",
})
if err != nil {
    return err
}
_, err = sent.EditText(bot, "<i>Updated</i>", &gotdbot.EditMessageTextOpts{
    ParseMode: "HTML",
})
```

Text and caption helpers accept `HTML`, `Markdown`, or `MarkdownV2` parse modes.
Explicit text entities take precedence over the parse mode. Set
`ClientOpts.ParseMode` to choose a default for the client.
`FormattedText.Format(parseMode)` renders stored text and entities as markup.
`User.Mention(parseMode)` creates a text mention from a user's full name.

Use `msg.OriginalTextHTML()`, `OriginalTextMD()`, or `OriginalTextMDV2()` to
convert message text or captions back to markup. Text-only `OriginalHTML/MD/MDV2`
and caption-only `OriginalCaptionHTML/MD/MDV2` helpers are also available.

| Content | Send helpers |
| --- | --- |
| Text and rich messages | `SendMessage`, `SendRichMessage`, `SendRichMessageMarkdown`, `SendRichMessageHTML` |
| Photos and video | `SendPhoto`, `SendLivePhoto`, `SendVideo`, `SendAnimation`, `SendVideoNote` |
| Audio, files, stickers | `SendAudio`, `SendVoice`, `SendDocument`, `SendSticker` |
| Locations and contacts | `SendLocation`, `SendLiveLocation`, `SendVenue`, `SendContact` |
| Interactive messages | `SendPoll`, `SendQuiz`, `SendChecklist`, `SendDice`, `SendStakeDice`, `SendGame` |
| Payments and sharing | `SendInvoice`, `SendPaidMedia`, `SendStory`, `SendForwarded`, `SendMediaGroup` |

Pass an `InputFileLocal` to upload a file, an `InputFileId` for a TDLib file ID,
or another supported `InputFile` type. Media captions use the same formatting
options as text. `SendLivePhoto` takes both the photo and its video;
`SendLiveLocation` takes a native `LiveLocation`.

`EditText`, `EditCaption`, and `EditReplyMarkup` update existing messages.
`EditContent` selects the appropriate edit operation for native content.
`Copy` and `Forward` send a message to another chat, and `Download` downloads
its media. `SendMessageContent` and `ReplyContent` accept any native
`InputMessageContent`.

`GetFile` returns the attached media file or `nil`. `Download` uses its TDLib file
ID with the client that received the message and returns `ErrNoMedia` when no
file is attached. `Mention` formats a user sender's name or returns a chat
sender's title.

Use `EditPhoto`, `EditVideo`, `EditAnimation`, `EditAudio`, or `EditDocument` to
replace media without constructing an `InputMessageContent`. Their matching
`EditMessage*Opts` structs control captions, formatting, media metadata, and
reply markup. Client methods use the `EditMessage*` and `EditInlineMessage*`
names. Inline edits require files that are already uploaded.

```go
sent, err := bot.SendPhoto(chatID, &gotdbot.InputFileLocal{Path: "photo.jpg"}, nil)
if err != nil {
    return err
}
if err := sent.Pin(bot, &gotdbot.PinChatMessageOpts{
    DisableNotification: true,
}); err != nil {
    return err
}
_, err = sent.EditPhoto(bot, &gotdbot.InputFileLocal{Path: "updated.jpg"},
    &gotdbot.EditMessagePhotoOpts{
        Caption:   "<b>Updated photo</b>",
        ParseMode: "HTML",
        HasSpoiler: true,
    })
```

Optional settings use pointer structs; pass `nil` for defaults. Helpers leave
caller-owned options unchanged, including when they fill a reply destination.
An explicit `ReplyTo` takes precedence over `ReplyToMessageID`. To delete a
message for everyone, use `msg.Delete(bot, &gotdbot.DeleteMessagesOpts{Revoke: true})`.
For an ephemeral message, set `ReceiverUserID` or `CallbackQueryID` in the send
options. `SendPoll` defaults to a regular poll; `SendQuiz` takes zero-based
correct answer IDs and optional explanation text, entities, and media.

### Buttons

`ParseTextWithButtons` returns formatted text and a normal inline keyboard:

```go
source := "*Choose*\n[Docs](buttonurl:https://core.telegram.org)\n[Done](buttoncallback:done:same)"
text, keyboard, err := bot.ParseTextWithButtons(source, gotdbot.ParseModeMarkdown)
if err != nil {
    return err
}
_, err = msg.ReplyContent(bot, &gotdbot.InputMessageText{Text: text},
    &gotdbot.SendMessageContentOpts{ReplyMarkup: keyboard})
```

For `ReplyText`, extract the buttons and pass the remaining markup:

```go
body, keyboard, err := gotdbot.ParseButtonMarkup(source, gotdbot.ParseModeMarkdown)
if err != nil {
    return err
}
_, err = msg.ReplyText(bot, body, &gotdbot.SendMessageOpts{
    ParseMode:   gotdbot.ParseModeMarkdown,
    ReplyMarkup: keyboard,
})
```

Use `buttonurl:`, `buttoncallback:` or `buttoncopy:`; `:same` joins the previous
row. Buttons use this same syntax with HTML, Markdown, MarkdownV2 or plain text
(`ParseModeNone`).
Add `|style=primary` or `|icon=123` inside the target for a style or custom emoji:
`[Docs](buttonurl:https://example.com|style=primary)`.

`msg.OriginalTextWithButtons(mode)` converts text or captions with their keyboard
back to markup. All button types, styles and custom emoji icons are preserved.
Try `/buttons` and `/format` in the [formatting example](./examples/formatting).

Button syntax is inspired by
[Paul Larsen's gotg_md2html](https://github.com/PaulSonOfLars/gotg_md2html).

### Rich messages and drafts

Rich messages support structured blocks and embedded media. Rich Markdown and
HTML use separate formatting rules from ordinary text parse modes and are
available to bots only.

```go
sent, err := bot.SendRichMessageMarkdown(chatID, "# Hello\n\nA rich message.",
    &gotdbot.SendRichTextOpts{
        SendMessageOpts: gotdbot.SendMessageOpts{DisableNotification: true},
        DetectAutomaticBlocks: true,
    })
if err != nil {
    return err
}
_, err = sent.EditRichMessageHTML(bot, "<h1>Updated</h1><p>Rich content.</p>", nil)
```

`msg.OriginalRichHTML()` and `msg.OriginalRichMD()` convert received rich
messages, ordinary text, or captions into rich markup and reusable media:

```go
source, media := msg.OriginalRichHTML()
_, err := msg.ReplyRichMessageHTML(bot, source, &gotdbot.SendRichTextOpts{
    Media: media,
    DetectAutomaticBlocks: true,
})
```

Fetch `GetFullRichMessage` when `IsFull` is false and preserve `IsRtl` when sending.
Ordinary media conversion uses the caption. Native blocks are also supported
through `RichMessageSourceBlocks`.

`SendMessageDraft`, `SendRichMessageDraftMarkdown`, and
`SendRichMessageDraftHTML` stream partial responses. Reuse a draft ID while
producing the response, then send the final message normally. Draft options
include forum topics, `CanStop`, and `KeepOnStop`. Rich draft media must already
be uploaded.

## Handlers and filters

`OnCommand` handles commands, and `OnMessage` handles new messages and business
message updates. Use generated typed handlers such as `OnUpdateNewCallbackQuery`,
`OnUpdateMessageEdited`, and `OnUpdateDeleteMessages` for specific updates.
`OnUpdate` accepts any native update.

Commands default to `/`. Use `SetTriggers` to accept other prefix characters:

```go
bot.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
    _, err := msg.ReplyText(c, "Hello!", nil)
    return err
}).SetTriggers("/!.")
```

Use `AddAliases` to send several commands to the same function:

```go
bot.OnCommand("start", startHandler).AddAliases("help", "info")
```

Here `/start`, `/help`, and `/info` all call `startHandler`. Aliases are
case-insensitive and share the handler's triggers, filters, and matching options.
Provide names without a prefix and configure aliases before starting the client.

Filter predicates are available in `filters/message` and `filters/callbackquery`.
Import `github.com/AshokShau/gotdbot/filters` and
`github.com/AshokShau/gotdbot/filters/message` to compose message predicates:

```go
bot.OnMessage(func(c *gotdbot.Client, msg *gotdbot.Message) error {
    _, err := msg.ReplyText(c, msg.GetText(), nil)
    return err
}, filters.AllFilters(
    message.Private,
    message.Incoming,
    filters.NotFilter(message.Command),
    func(msg *gotdbot.Message) bool { return !msg.Empty() },
))
```

`filters.AllFilters`, `filters.AnyFilters`, and `filters.NotFilter` compose
predicates. Message filters also expose `And`, `Or`, and `Not`.
AND and OR short-circuit and ignore nil predicates.
`CompileRegex` in either filter package returns an error for an
invalid pattern.

Handler groups run in ascending order; handlers within a group run in registration
order. The first matching handler in each group runs unless it returns
`ContinueHandlers`. Return `ContinueGroups` to advance to the next group or
`EndGroups` to stop processing that update. Use `OnCommandGroup`,
`OnMessageGroup`, or `OnUpdateGroup` to select a group. Configure handler fields
and filters before starting the client.

### Working with updates

Bots skip incoming messages sent/edited at least 24 hours ago, and group messages
from before joining or while absent. Supergroups/channels also enforce the session
authorization date. Set `ClientOpts.ReceiveOldMessages: true` to disable this filter.

`ClientOpts.DropPendingUpdates: true` skips application updates during initial sync,
including new arrivals. Internal handling continues; later reconnects are unaffected.

Typed message and guest-query updates expose their carried message through
`GetMessage`. Business callbacks return the carried message or an error.
`client.UpdateChatID(update)`, `UpdateMessageID`, and `UpdateSenderID` extract
common identifiers, returning `0` when unavailable.
`client.GetUpdateMessage(update)` returns a carried message or fetches it
from the update's chat and message IDs. Updates without a single message return
`ErrUpdateHasNoMessage`.

Edit and content updates expose `GetMessage(client)` to fetch the
current message:

```go
bot.OnUpdateMessageEdited(func(c *gotdbot.Client, update *gotdbot.UpdateMessageEdited) error {
    msg, err := update.GetMessage(c)
    if err != nil {
        return err
    }
    c.Logger.Info("Message edited", "chat_id", msg.ChatId, "message_id", msg.Id)
    return nil
}, nil)
```

Chat, inline, and business callback updates provide `Answer` and message edit
helpers. Use `query.Answer(client, "Done", nil)` for a notification, or pass
`CallbackAnswerOpts` to set the cache time, alert behavior, or URL. Business
callback edits use the callback's connection ID and return `BusinessMessage`.
Chosen inline results provide edit helpers when their `InlineMessageId` is known.
Guest queries expose their carried message and an `Answer` method that accepts
an `InputInlineQueryResult`.

Inline-query, shipping-query, and pre-checkout updates also expose `Answer`; join-request
updates provide `Approve` and `Decline`. An empty checkout error message accepts
the checkout; a non-empty message rejects it.

## Configuration and accounts

`ClientOpts` configures the library path, storage directories, formatting,
request timeout, update workers, retry behavior, and error/panic callbacks.
`DefaultClientConfig` provides the defaults. Use `UpdateWorkers: 1` when handlers
must run serially; the default is two workers.

`Ask` waits for a matching message in a chat using `WaitMessageOpts` filters and
a timeout. `WaitForContext` waits for updates with cancellation. A blocking
conversation occupies a handler worker; reserve workers for incoming updates
when running conversations inside handlers.

For multiple accounts, use `ClientManager` with a separate database directory for
each client. The [multiple-account example](./examples/echoMultiBot) shows setup
and shutdown.

## Examples and reference

- [Echo bot](./examples/echoBot): Commands, keyboards, callbacks, and filters.
- [Formatting bot](./examples/formatting): Reply to text, captions or rich messages
  with `/format` to compare HTML, Markdown, MarkdownV2 and rich conversions.
- [Auto-delete bot](./examples/autoDeleteBot): Outgoing-message timers and deletion updates.
- [Conversation bot](./examples/conversation): Questions and reply filters.
- [Multiple accounts](./examples/echoMultiBot): Client management and shared handlers.

The [Go reference](https://pkg.go.dev/github.com/AshokShau/gotdbot) documents public
methods and options. Refer to [TDLib documentation](https://core.telegram.org/tdlib/docs/)
for native identifiers, types, and account permissions.

## AI coding skill

The optional [GotdBot skill](./skills/gotdbot/SKILL.md) provides coding-agent
guides and a [bot starter](./skills/gotdbot/assets/bot/main.go).

## Development

Tests run without Telegram credentials or an installed TDLib library:

```sh
go test ./...
go test -race ./...
go vet ./...
```

Regenerate bindings from the current TDLib JSON schema:

```sh
go generate .
```

The generator uses
[`FallenProjects/tdlib-build`'s master schema](https://raw.githubusercontent.com/FallenProjects/tdlib-build/refs/heads/master/tdlib.json).
To use a local schema, run `go run ./internal/gen -schema /path/to/tdlib.json`.
Use a TDLib shared library that matches the generated bindings. Binding changes
belong in the generator.

## License

[MIT](./LICENSE).
