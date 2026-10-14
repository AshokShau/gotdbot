# Handlers and conversations

Contents: registration; matching; commands; filters; group control; typed updates; conversations; concurrency.

Source anchors: `handlers.go`, `command_handler.go`, `message_handler.go`, `update_handler.go`, `dispatcher.go`, `gen_events.go`, `gen_filters.go`, `filters/`, `conversation.go`, `waiters.go`.

## Registration and matching

| Registration | Callback | Filter | Explicit group |
| --- | --- | --- | --- |
| `OnCommand(name, callback)` | `func(*Client, *Message) error` | Chain `AddFilter` | `OnCommandGroup(name, callback, group)` |
| `OnMessage(callback, filter)` | `func(*Client, *Message) error` | `MessageFilter` | `OnMessageGroup(callback, filter, group)` |
| `OnUpdate(callback, filter)` | `func(*Client, Update) error` | `UpdateFilter` | `OnUpdateGroup(callback, filter, group)` |
| `OnUpdateNewMessage(callback, filter)` and other generated methods | A pointer to the specific update type | Matching generated filter type | Inspect the generated group method signature |
| `AddHandler(handler)` | Custom `Handler` implementation | `CheckUpdate` | `AddHandlerGroup(handler, group)` |

Command/message handlers process `UpdateNewMessage`, `UpdateNewBusinessMessage`, and `UpdateBusinessMessageEdited`. An ordinary `UpdateMessageEdited` or `UpdateMessageContent` does not trigger `OnMessage`; register those typed updates and fetch the current message instead.

Default `MatchingOptions` accepts incoming messages, edited messages, and private/basic-group/supergroup/channel/secret chat categories. It rejects outgoing messages. These checks run before a supplied predicate; `message.Outgoing` alone cannot enable outgoing processing. Begin with `DefaultMatchingOptions()`, change `AllowOutgoing`/`AllowIncoming`, then call `SetOptions`. A literal zero-valued `MatchingOptions` rejects all categories.

`SetOptions`, `SetTriggers`, `AddAliases`, and `AddFilter` mutate handler objects. Configure them before starting updates. Handler registration uses an internal snapshot, but that does not make live field mutation safe.

## Commands and arguments

Pass `"start"`, not `"/start"`, to `OnCommand`. Commands default to `/`; `SetTriggers("/!.")` accepts each of those characters separately. Empty triggers disable matching. Matching is case-insensitive, trims surrounding whitespace, and recognizes a command as the first whitespace-separated token.

Chain `AddAliases("help", "info")` to accept several command names with one callback, for example `c.OnCommand("start", startHandler).AddAliases("help", "info")`. Names have no trigger prefix; surrounding whitespace is trimmed and empty aliases are ignored. Repeated calls append aliases. Aliases share triggers, matching options, filters, and bot-username checks with the primary command. This also works with `NewCommandHandler` and `OnCommandGroup`.

`/start@MyBot` matches only an active username on the cached current account. Command matching reads text/captions; it does not require a bot-command entity. By comparison, `message.Command` uses `Message.IsCommand`, which requires a `TextEntityTypeBotCommand` at offset zero. A custom `!start` trigger is not necessarily identified by that filter. Put explicit command handlers before broad handlers and use an explicit predicate for custom-trigger exclusions when needed.

`ArgsList()` uses `strings.Fields` on message **text**, omits the command token, and is not a shell/quoted-argument parser. `Args()` rejoins those arguments with spaces; `Arg(index)` returns an argument or empty string. Caption commands need caption-specific argument parsing.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func RegisterCommands(c *gotdbot.Client) {
	c.OnCommand("say", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		text := msg.Args()
		if text == "" {
			text = "Usage: /say some text"
		}
		_, err := msg.ReplyText(c, text, nil)
		return err
	}).SetTriggers("/!.")
}
```

This example assumes the client's default parse mode is plain text. With a markup default, escape user input or send native `FormattedText` as described in the formatting guide.

## Filters

Import `filters`, `filters/message`, and `filters/callbackquery` by their full module paths. `message` predicates include `Private`, `Group` (basic groups only), `Supergroup` (excluding channels), `Channel`, `Secret`, `Incoming`, `Outgoing`, `Edited`, `Command`, `Prefix`, `Suffix`, `Equal`, `FromUserID`, `ChatID`, and regex filters.

- `filters.AllFilters(...)` ANDs non-nil predicates and short-circuits on false; an empty list matches everything.
- `filters.AnyFilters(...)` ORs non-nil predicates and short-circuits on true; an empty list matches nothing.
- `filters.NotFilter(nil)` matches nothing.
- `MessageFilter.And`, `.Or`, `.Not` wrap the same composition behavior.
- `CompileRegex(pattern)` returns `(filter, error)`. `Regex(pattern)` returns a predicate that matches nothing on an invalid pattern; use `CompileRegex` for configurable patterns.
- `MessageFilter.ForNewMessage()` adapts a message predicate to `UpdateNewMessageFilter`. A nil source filter matches nothing.

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters"
	"github.com/AshokShau/gotdbot/filters/message"
)

func RegisterTextHandler(c *gotdbot.Client, pattern string) error {
	matches, err := message.CompileRegex(pattern)
	if err != nil {
		return err
	}
	c.OnMessage(func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, "Matched", nil)
		return err
	}, filters.AllFilters(message.Private, message.Incoming, matches))
	return nil
}
```

## Group control

Groups run in ascending numeric order, including negative groups. Within a group, handlers run in registration order until one matches. Its returned error is interpreted as follows:

| Result | Current group | Later groups |
| --- | --- | --- |
| `nil` | Stop scanning | Continue |
| `ContinueHandlers` | Scan further handlers | Continue when group finishes |
| `ContinueGroups` | Stop scanning | Continue |
| `EndGroups` | Stop | Stop |
| Other error | Invoke configured ErrorHandler, or log | Interpret ErrorHandler result using the same rules |

Thus returning `nil` does not end all processing. A general observer that shares a group with a command may consume that group's match; put it in a separate group or return `ContinueHandlers`. Sentinels can be wrapped because dispatch uses `errors.Is`.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func RegisterGroups(c *gotdbot.Client, allowed func(*gotdbot.Message) bool) {
	c.OnMessageGroup(func(_ *gotdbot.Client, msg *gotdbot.Message) error {
		if !allowed(msg) {
			return gotdbot.EndGroups
		}
		return gotdbot.ContinueGroups
	}, nil, -1)
	c.OnCommandGroup("ping", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, "pong", nil)
		return err
	}, 0)
	c.OnMessageGroup(func(c *gotdbot.Client, msg *gotdbot.Message) error {
		c.Logger.Debug("processed message", "chat_id", msg.ChatId, "message_id", msg.Id)
		return nil
	}, nil, 1)
}
```

The gate covers message updates accepted by its MatchingOptions. Authorize callbacks and other non-message events in their own handlers; this gate is not global access control.

## Conversations and waiters

`Ask(chatID, *WaitMessageOpts)` waits for `UpdateNewMessage` in that chat. Options contain `Filter`, `CancellationFilter`, and `Timeout`; missing/nonpositive timeout defaults to one minute. CancellationFilter is checked before the normal Filter and produces `ConversationCancelled` after a matching cancellation message arrives.

In a group, constrain both normal and cancellation predicates to the initiating sender. A chat-only filter lets any member answer or cancel. Concurrent waiters can receive the same update; enforce one active conversation per account/chat/sender in application state when needed.

```go
package recipes

import (
	"errors"
	"github.com/AshokShau/gotdbot"
	"time"
)

// Run with spare update workers, or invoke from an application-owned worker.
func AskName(c *gotdbot.Client, msg *gotdbot.Message) error {
	sender, ok := msg.SenderId.(*gotdbot.MessageSenderUser)
	if !ok || sender == nil {
		return nil
	}
	_, err := msg.ReplyText(c, "What is your name? /cancel to stop.", nil)
	if err != nil {
		return err
	}
	sameUser := func(m *gotdbot.Message) bool {
		s, ok := m.SenderId.(*gotdbot.MessageSenderUser)
		return ok && s != nil && s.UserId == sender.UserId && !m.IsOutgoing
	}
	answer, err := c.Ask(msg.ChatId, &gotdbot.WaitMessageOpts{
		Timeout: 30 * time.Second,
		Filter:  func(m *gotdbot.Message) bool { return sameUser(m) && m.Text() != "" },
		CancellationFilter: func(m *gotdbot.Message) bool {
			return sameUser(m) && m.Text() == "/cancel"
		},
	})
	if errors.Is(err, gotdbot.ConversationCancelled) || errors.Is(err, gotdbot.ConversationTimeout) {
		_, replyErr := msg.ReplyText(c, "Conversation ended.", nil)
		return replyErr
	}
	if err != nil {
		return err
	}
	_, err = answer.ReplyText(c, "Thanks!", nil)
	return err
}
```

`WaitFor(filter, timeout)` is global; `WaitForChat(chatID, filter, timeout)` narrows candidates; `WaitForContext(ctx, chatID, filter, timeout)` supports cancellation. Filters take `func(*Client, TlObject) bool`, not `UpdateFilter`. A chat ID of zero selects global waiters. Unlike Ask, the lower-level timeout is used directly; pass a positive duration. Filters must be non-nil.

Waiters run on processor workers before application handler groups and do not consume an update. The update may still trigger other handlers. Waiters are registered when the wait call starts; a message arriving between sending a prompt and registering Ask may be missed. If that race matters, design explicit registration/state handling rather than assuming atomic prompt-and-wait behavior.

## Concurrency and error callbacks

`UpdateWorkers: 1` serializes processing but cannot support a blocking Ask inside that sole worker. The default two workers support one such wait only while the other remains available; simultaneous blocking conversations can exhaust both. For many conversations, use application-owned jobs with cancellation and a bounded concurrency policy, or an update-driven state machine.

Protect application maps/state shared by workers. Processor filters and callbacks should avoid long blocking work unless the worker capacity accounts for it. Panic recovery applies to processor execution; goroutines launched by application handlers need their own error/panic management. ErrorHandler receives the original update and can translate an error to a group-control sentinel.
