---
name: gotdbot
description: Build, explain, debug, or extend Go Telegram bots and user clients using github.com/AshokShau/gotdbot and TDLib. Use for GotdBot authentication, messages, media, formatting, commands, filters, callbacks, conversations, multi-account clients, native TDLib requests, and binding generation.
---

# GotdBot

Use this skill to implement working GotdBot applications or maintain the library. The import path is `github.com/AshokShau/gotdbot`. It wraps TDLib through a dynamically loaded shared library and exposes both convenience helpers and native typed requests.

## Establish the target

Determine whether the task changes an application that imports GotdBot or the GotdBot library itself. Preserve the project's Go module, account type, storage paths, and chosen dependency version. For a new application, start with [assets/bot/main.go](assets/bot/main.go); it reads credentials from environment variables and registers a command and filtered echo handler.

This skill was checked against library `v0.9.6`, TDLib `v1.8.68`, and a module requiring Go `1.26.0`, including shared compact button markup with `ParseModeNone`, native button types, rich formatting renderers and the reply-based formatting example. These are the bundled reference baseline, not a requirement to upgrade an existing project. Verify the actual dependency and native-library versions before adapting examples.

```sh
go list -m -json github.com/AshokShau/gotdbot
go doc github.com/AshokShau/gotdbot.Client.SendMessage
go doc github.com/AshokShau/gotdbot.SendMessageOpts
```

In a library checkout, read `go.mod`, `consts.go`, and the relevant source. In a consuming application, find the downloaded source with `go list -m -f '{{.Dir}}' github.com/AshokShau/gotdbot`. All library source paths in the references are relative to that module directory, not to this installed skill. Read only the topic guides needed for the current task.

## Choose the relevant guide

| Task | Read |
| --- | --- |
| Install TDLib, authenticate a bot/user, configure timeouts, run several accounts | [Setup and lifecycle](references/setup-and-lifecycle.md) |
| Commands, filters, group control, typed handlers, Ask and waiters | [Handlers and conversations](references/handlers-and-conversations.md) |
| Send/reply/edit/copy/forward messages, upload/download files, polls, albums | [Messages and media](references/messages-and-media.md) |
| `GetFormattedText`, `ParseText`, parse modes, entities, mentions, rich blocks, streamed drafts | [Formatting and rich messages](references/formatting-and-rich-messages.md) |
| Keyboards, callbacks, inline/business/guest queries, update lookup, payments | [Callbacks and updates](references/callbacks-and-updates.md) |
| Membership, moderation, topics, invite links, bot profile/menu/commands | [Chats and administration](references/chats-and-administration.md) |
| Any generated TDLib capability, raw request structs, JSON, signature discovery | [Native API](references/native-api.md) |
| Modify this library, regenerate bindings, diagnose failures, validate changes | [Development and troubleshooting](references/development-and-troubleshooting.md) |

The native API guide is the discovery path for capabilities beyond the examples, including stories, stickers, searches, business features, calls, proxies, downloads, gifts, and account settings. Search the installed generated bindings instead of assuming a method exists or copying another library's signature.

## Working rules

1. **Keep TDLib's types and identifiers.** Messages use `ChatId` and `Id`; `SenderId` is a `MessageSender` union. Use IDs returned by TDLib with the same account. Never assume Bot API message IDs or file IDs can be substituted.
2. **Match the API layer.** `Client.SendMessage(chatID, text, *SendMessageOpts)` is the text helper. `SendMessageContent` takes native `InputMessageContent` and `*SendMessageContentOpts`; its options nest `*MessageSendOptions`. `SendWithContext` takes a pointer to a native request, such as `&GetMessage{...}`. Verify argument order with `go doc` or source.
3. **Respect client lifecycle.** `NewClient` loads TDLib immediately; register handlers before `Start`/`StartContext`, check startup errors, then keep the application alive. `Me` is available after successful startup and is library-owned. Use one persistent database directory per account. Close clients on application exit.
4. **Choose formatting deliberately.** `Client.GetFormattedText(text, entities, mode)` preserves non-empty entities or resolves the client's default mode before parsing. `Client.ParseText(text, mode)` uses only the supplied mode; an empty mode produces plain text. `Message.GetFormattedText()` extracts existing message text and does not parse it. Ordinary HTML/Markdown and rich HTML/Markdown use different APIs. Escape interpolated text, and count entity offsets in UTF-16 code units.
5. **Treat handler options and concurrency as behavior.** Default handlers reject outgoing messages. One matching handler runs per group unless it returns `ContinueHandlers`; later groups still run unless processing ends. Configure handler fields before updates begin. Waiters need an available processor worker, so a blocking conversation cannot occupy the only worker.
6. **Use options without mutating shared inputs.** Helper options may be nil. Replies preserve an explicit `ReplyTo` object; otherwise they fill the source message ID. Native and convenience option structs are not interchangeable. Preserve caller-owned options when extending helpers.
7. **Handle errors at the right boundary.** Return request errors from handlers. Use `errors.Is` for sentinels/context cancellation and `errors.As` for `*gotdbot.Error`. A send timeout can occur after Telegram accepted a message; do not blindly resend it. Retry policy is bounded by the request deadline.
8. **Verify account restrictions.** Read the native method/type comments for bot-only, user-only, Premium, ownership, or administrator requirements. A Go method's existence does not establish permission to use it. Validate application authorization before moderation, paid actions, or processing sensitive callback payloads.
9. **Preserve generated ownership.** Files beginning `gen_` are generated. Change `internal/gen` or handwritten helpers as appropriate, then regenerate against the intended schema; ordinary application work does not require generation.

## Finish with a usable result

Implement the requested behavior, compile it, and run checks relevant to the change. Library unit tests need neither Telegram credentials nor an installed TDLib shared library; live startup does. For documentation/examples, compile the examples without logging in. Report what works, how to configure credentials and TDLib, what was checked, and any runtime behavior that remains unverified.

## Share or install this skill

Copy the entire `gotdbot` folder, including `references`, `assets`, and `agents`; `SKILL.md` links to these resources. For Codex, install that folder under `${CODEX_HOME:-$HOME/.codex}/skills/gotdbot`, or under a project's `.agents/skills/gotdbot` for project discovery. Other agents can load this `SKILL.md` and follow its relative links using their own skill mechanism.

Example invocation: `Use $gotdbot to add a /start command, a callback keyboard, and a photo download handler to this Go project.`
