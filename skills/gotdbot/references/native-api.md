# Native API and generated bindings

Contents: source discovery; method/options conventions; context-aware requests; complete API map; JSON and unions; errors and retries.

Source anchors: `gen_classes.go`, `gen_types.go`, `gen_functions.go`, `gen_methods.go`, `gen_options.go`, `gen_helpers.go`, `gen_events.go`, `gen_filters.go`, `gen_tdlib_options.go`, `client_requests.go`, `request_lifecycle.go`, `json_types.go`, `internal/gen/parameters.go`, `internal/gen/api_names.go`.

## Discover the actual installed API

The generated bindings are the full native capability layer. This skill gives a workflow for **any** native feature without copying thousands of generated symbols into its instructions. Determine the dependency version and search that version's module directory. In a GotdBot checkout:

```sh
rg -n '^func .* (GetChatHistory|SearchChatMessages|SendStory)\(' gen_methods.go
rg -n '^type (GetChatHistoryOpts|SearchChatMessagesOpts|InputStoryContent)' gen_options.go gen_types.go gen_classes.go
rg -n '^func .*OnUpdateNewCallbackQuery' gen_events.go
rg -n '^type UpdateNewCallbackQueryFilter' gen_filters.go
```

In an application, use `go doc github.com/AshokShau/gotdbot.Client.GetChatHistory` and `go doc github.com/AshokShau/gotdbot.GetChatHistoryOpts`, or run those searches in the directory reported by `go list -m -f '{{.Dir}}' github.com/AshokShau/gotdbot`.

Read the request method, its option struct, nested union implementations, response type, and relevant updates together. The generated comments carry bot/user restrictions and native constraints. If local comments are incomplete, consult the authoritative TDLib docs matching the schema; do not assume a newer schema is compatible with this binary.

## Method and options conventions

Native requests are exported structs in `gen_functions.go`; Client methods in `gen_methods.go` construct those requests and call `Send`. Required identifiers/data are usually positional. For requests with multiple parameters, nullable fields, booleans, and selected zero-default integer/string parameters move into an `Opts` struct. A one-parameter native function accepts its argument directly. There is no universal options parameter on every method.

The baseline generator sorts schema parameter names alphabetically before partitioning them into positional arguments and options. Argument order comes from that generated source, not a presumed Telegram Bot API order. Examples include `GetChatHistory(chatID, limit, opts)` and `SearchChatMessages(chatID, limit, query, opts)`. Read their signatures before generating code.

Name collisions are deliberately mapped:

| Native request struct | Generated Client method | Convenience helper |
| --- | --- | --- |
| `SendMessage` | `SendMessageContent` | `SendMessage(chatID, text, *SendMessageOpts)` |
| `EditMessageText` | `EditMessageTextContent` | `EditMessageText(chatID, messageID, text, *EditMessageTextOpts)` |
| `EditMessageCaption` | `EditMessageCaptionText` | `EditMessageCaption(chatID, messageID, caption, *EditMessageCaptionOpts)` |
| `Close` | Managed handwritten `Close()` | Use the managed lifecycle for normal shutdown. |

Native `SendMessageContentOpts` nests `Options *MessageSendOptions`, `ReplyMarkup`, `ReplyTo`, and `TopicId`. Convenience `SendMessageOpts` flattens delivery and formatting fields. Do not interchange them even when names are similar.

## Context-aware native requests

`SendWithContext(ctx, request)` requires a pointer to a generated **function** struct. Data objects may satisfy `TlObject` but are rejected as requests. Values of function structs are rejected too because the function marker is on the pointer. Nil/typed-nil requests and nil contexts produce errors.

```go
package recipes

import (
	"context"
	"fmt"
	"github.com/AshokShau/gotdbot"
)

func ReadMessage(ctx context.Context, c *gotdbot.Client, chatID, messageID int64) (*gotdbot.Message, error) {
	obj, err := c.SendWithContext(ctx, &gotdbot.GetMessage{ChatId: chatID, MessageId: messageID})
	if err != nil {
		return nil, err
	}
	msg, ok := obj.(*gotdbot.Message)
	if !ok || msg == nil {
		return nil, fmt.Errorf("unexpected response %T", obj)
	}
	return msg, nil
}
```

Generated Client convenience methods do not generally take a context. To cancel an individual native send, construct `&SendMessage{ChatId: ..., InputMessageContent: ..., Options: ...}` and call SendWithContext. Do not invent `SendMessageWithContext`/`GetMessageContext` methods. Native request structs use native field spelling (`ChatId`, `MessageId`), while handwritten convenience settings may use `ReplyToMessageID` or `CallbackQueryID`.

The transport inserts correlation `@extra` without mutating the request struct. `Send` is SendWithContext with a background context. Both apply the configured RequestTimeout, including retries, lookup requests, and send completion.

Outgoing send operations await `UpdateMessageSendSucceeded`/`UpdateMessageSendFailed` when TDLib returns a pending message. The returned successful message can have a different final ID from the temporary ID. Read/edit requests do not wait for another send's completion. Send results can include a message alongside an error; check the error before using the result as a confirmed send.

## API discovery map

Use these search terms as starting points, not guaranteed support for every account:

| Domain | Methods/types/updates to search |
| --- | --- |
| Chats and navigation | `GetChat`, `GetChats`, `LoadChats`, `ChatList`, `ChatPosition`, `UpdateNewChat` |
| History and search | `GetChatHistory`, `SearchChatMessages`, `SearchMessages`, `SearchMessagesFilter`, `UpdateMessageContent` |
| Members and permissions | `GetChatMember`, `GetSupergroupMembers`, `ChatMemberStatus`, `ChatAdministratorRights` |
| Forum/direct/saved topics | `ForumTopic`, `MessageTopic`, `GetForumTopics`, `DirectMessagesChatTopic`, `SavedMessagesTopic` |
| Files and transfer progress | `DownloadFile`, `UploadFile`, `UpdateFile`, `UpdateFileGenerationStart`, `FinishFileGeneration` |
| Stickers and emoji | `GetStickerSet`, `SearchStickers`, `ReactionType`, `CustomEmoji` |
| Stories | `SendStory`, `InputStoryContent`, `StoryPrivacySettings`, `UpdateStory` |
| Business and quick replies | `BusinessConnection`, `SendBusinessMessage`, `QuickReplyShortcut`, `UpdateNewBusinessMessage` |
| Calls and group calls | `Call`, `CallProtocol`, `GroupCall`, `UpdateCall`, `UpdateGroupCall` |
| Network and proxies | `AddProxy`, `ProxyType`, `PingProxy`, `UpdateConnectionState` |
| Account, contacts, and privacy | `GetMe`, `GetUser`, `AddContact`, `UserPrivacySetting`, `SetOption` |
| Payments, gifts, Stars | `Invoice`, `SendPaymentForm`, `SendGift`, `GetStarTransactions`, `UpdateNewPreCheckoutQuery` |
| Formatting, rich text, translation | `GetFormattedText`, `ParseText`, `ParseTextEntities`, `InputRichMessage`, `TranslateText`, `ComposeTextWithAi` |
| Logging/configuration | `SetLogStream`, `SetLogVerbosityLevel`, `TDLibOptions`, `OptionValue` |

TDLib may require explicit loading/pagination. For example, an initial chat list is not guaranteed complete: read GetChats/LoadChats semantics and observe native list updates. History/search responses use their own cursors and limits. Avoid treating a single page as the entire account dataset or treating the Go client as a high-level call/media engine.

## Unions, JSON, and IDs

`TlObject` supplies `GetType()`. Generated interfaces such as `Update`, `MessageContent`, `InputMessageContent`, `MessageSender`, `InputFile`, and `ReplyMarkup` represent native unions. Use their concrete exported implementations and type switches; custom outside types cannot implement private union markers.

Generated structs provide JSON marshalling with `@type` and custom decoding of union fields. `Unmarshal(data)` returns `(TlObject, extra, error)`; `UnmarshalWithClient(data)` adds the native client identifier for routing. `UnmarshalMessageContent`, `UnmarshalMessageSender`, and other union helpers select the concrete type. Decoders return pointers. Unknown types indicate unsupported input or schema/binary mismatch; do not silently discard required semantics.

Go types retain the distinction between TDLib `int32`, `int53`, and `int64`: `int53` maps to Go int64 but is normally a JSON number, while true `int64` fields use string encoding. `Int64Slice` handles vector<int64>, accepting numeric/quoted integer JSON and marshalling quoted decimal strings. Bytes use `[]byte` and standard JSON base64 encoding. Preserve these codecs; do not round identifiers through float64 or replace binary payloads with untyped maps.

## Errors and bounded retries

```go
package recipes

import (
	"context"
	"errors"
	"github.com/AshokShau/gotdbot"
)

func ErrorKind(err error) string {
	if err == nil {
		return "ok"
	}
	if errors.Is(err, gotdbot.ErrClientClosed) {
		return "closed"
	}
	if errors.Is(err, gotdbot.SendTimeout) || errors.Is(err, context.DeadlineExceeded) {
		return "timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "cancelled"
	}
	var tdErr *gotdbot.Error
	if errors.As(err, &tdErr) {
		if tdErr.Code == 429 {
			return "rate-limited"
		}
		return "tdlib-error"
	}
	return "application-error"
}
```

`Error.Code` and `.Message` describe native errors; native code 406 comments forbid processing or displaying its message. Avoid exposing raw error details to users without considering native rules and application context.

`AutoRetry` retries explicit rejections only:

- `ChatNotFound`: one GetChat lookup for the matching 400 rejection, then retry if lookup succeeds.
- `MessageNotFound`: one GetMessage lookup for the matching 400 rejection, then retry if lookup succeeds.
- `MaxFloodWait`: at most one parsed 429 retry-after wait within the configured maximum and current request deadline.

`Error.GetRetryAfter()` parses the expected retry-after error format and can return zero. Retries are not a guarantee that every native error spelling can be recovered. Transport timeouts are not automatically retried: an outgoing mutation may already have happened. Determine the actual state before an application-level retry that could duplicate a send or other mutation.
