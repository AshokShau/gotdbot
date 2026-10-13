# Callbacks and updates

Contents: keyboard recipe; callback kinds; payloads; message lookup; typed updates; inline/guest queries; payments.

Source anchors: `callback_helpers.go`, `update_helpers.go`, `gen_events.go`, `gen_filters.go`, `filters/callbackquery/callbackquery.go`, `gen_types.go`, `gen_options.go`.

## Inline keyboard and callback

Keyboard rows contain **values**: `[][]InlineKeyboardButton`. Each button uses a `Type` union, such as `InlineKeyboardButtonTypeCallback{Data: []byte(...)}` or `InlineKeyboardButtonTypeUrl{Url: ...}`. This is not a Bot API `CallbackData` field on the button.

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters/callbackquery"
)

func RegisterKeyboard(c *gotdbot.Client) {
	c.OnCommand("menu", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, "Choose an action", &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkupInlineKeyboard{
				Rows: [][]gotdbot.InlineKeyboardButton{{{
					Text: "Done",
					Type: &gotdbot.InlineKeyboardButtonTypeCallback{Data: []byte("menu:done")},
				}}},
			},
		})
		return err
	})
	c.OnUpdateNewCallbackQuery(func(c *gotdbot.Client, q *gotdbot.UpdateNewCallbackQuery) error {
		if err := q.Answer(c, "Done", nil); err != nil {
			return err
		}
		_, err := q.EditMessageText(c, "Completed", nil)
		return err
	}, callbackquery.Equal("menu:done"))
}
```

Answer callbacks promptly to acknowledge the user's action. `CallbackAnswerOpts` controls `CacheTime`, `ShowAlert`, and `Url`. For account-specific operations, validate `SenderUserId`, the selected resource, and application access before applying a callback's requested change. A data prefix is routing, not authorization.

## Match the callback kind

| Update | Identifiers/context | Message access | Edit result |
| --- | --- | --- | --- |
| `UpdateNewCallbackQuery` | Chat and message IDs | `GetMessage(c)` fetches | `*Message, error` |
| `UpdateNewInlineCallbackQuery` | `InlineMessageId` | No ordinary chat-message lookup | `error` |
| `UpdateNewBusinessCallbackQuery` | Business connection plus carried business message | `GetMessage()` extracts with error | `*BusinessMessage, error` |
| `UpdateNewChosenInlineResult` | Optional `InlineMessageId` | `MessageID()` validates | `error` via `EditText`, `EditContent`, and siblings |

Chat callback helpers include `Answer`, `EditMessageText`, `EditMessageCaption`, `EditMessageReplyMarkup`, `EditMessageMedia`, `DeleteMessage`, rich edits, and live-location helpers. Inline callbacks expose their applicable counterparts and cannot upload fresh files when editing. Business callbacks include `EditMessageContent` and media/rich helpers using the connection ID; `EditMessageReplyMarkup` takes business-specific opts. Confirm each method's signature before sharing options across contexts.

## Callback payloads and filters

`Payload` is a `CallbackQueryPayload` union. Prefer `CallbackData()`/`DataString()` over unsafe payload assertions. `CallbackDataWithPassword()` exposes data plus the password field for that variant; do not log the password. `GameShortName()` handles game payloads. No-data variants can return nil/empty values.

The callback filter package handles **chat callbacks** (`UpdateNewCallbackQueryFilter`): `Prefix`, `Suffix`, `Equal`, `Regex`, `CompileRegex`, `FromUserID`, `ChatInstance`, `GameName`, and `Private`. Write correctly typed predicates for inline/business callbacks rather than passing the chat filter type.

## Resolve a message from an update

| Helper | Behavior |
| --- | --- |
| `c.UpdateChatID(update)` | Extracts carried message/native `ChatId`; zero if absent. |
| `c.UpdateMessageID(update)` | Extracts one message ID; current ID for send-completion updates, zero for multi-message deletions. |
| `c.UpdateSenderID(update)` | Extracts sender user/chat/query/join-request ID; zero if absent. |
| `c.GetUpdateMessage(update)` | Returns a carried message or fetches it from chat/message IDs. |

`GetUpdateMessage` takes `Update`, reports `ErrNilUpdate` for nil and `ErrUpdateHasNoMessage` when the event has no single message. Deletion updates have `MessageIds`, not one MessageId. Chosen inline results and inline callbacks do not identify an ordinary chat message. Avoid synthesizing IDs or assuming every update can be replied to.

No-client `GetMessage()` returns a carried message on new-message, new-business-message, business-edited, guest-query, and send-completion updates. `UpdateMessageEdited.GetMessage(c)` and `UpdateMessageContent.GetMessage(c)` fetch the current message and may fail if it was deleted. Business callbacks' no-client `GetMessage` returns an error as well as a message, unlike the carried-message helpers.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func ObserveEdits(c *gotdbot.Client) {
	c.OnUpdateMessageEdited(func(c *gotdbot.Client, u *gotdbot.UpdateMessageEdited) error {
		msg, err := u.GetMessage(c)
		if err != nil {
			return err
		}
		c.Logger.Info("edited", "chat_id", msg.ChatId, "message_id", msg.Id)
		return nil
	}, nil)
	c.OnUpdateDeleteMessages(func(c *gotdbot.Client, u *gotdbot.UpdateDeleteMessages) error {
		c.Logger.Info("deleted", "chat_id", u.ChatId, "message_ids", u.MessageIds)
		return nil
	}, nil)
}
```

The fetched message represents current state, not necessarily an exact snapshot at the edit event's time.

## Inline and guest queries

`UpdateNewInlineQuery.Answer(c, []InputInlineQueryResult, *AnswerInlineQueryOpts)` supplies native results. Configure inline mode for the bot through Telegram's supported setup, and read result type constraints. Chosen inline results can be edited only when an InlineMessageId is available; `MessageID()` returns `ErrUpdateHasNoMessage` otherwise.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func RegisterInline(c *gotdbot.Client) {
	c.OnUpdateNewInlineQuery(func(c *gotdbot.Client, u *gotdbot.UpdateNewInlineQuery) error {
		return u.Answer(c, []gotdbot.InputInlineQueryResult{
			&gotdbot.InputInlineQueryResultArticle{
				Id: "hello", Title: "Hello",
				InputMessageContent: &gotdbot.InputMessageText{
					Text: &gotdbot.FormattedText{Text: "Hello from GotdBot"},
				},
			},
		}, nil)
	}, nil)
}
```

`UpdateNewGuestQuery.GetMessage()` exposes its carried message. `Answer(c, InputInlineQueryResult)` returns `*InlineMessageId, error`; it is not the same signature as an inline query answer.

## Payments and join requests

`UpdateNewPreCheckoutQuery.Answer(c, errorMessage)` accepts checkout only when the error string is empty. Validate the invoice payload, currency, amount, and application order before accepting. A successful pre-checkout answer is not proof of payment; inspect the native successful-payment message/update before fulfillment.

`UpdateNewShippingQuery.Answer(c, []ShippingOption, *AnswerShippingQueryOpts)` supplies shipping choices or a rejection explanation in `ErrorMessage`. `UpdateNewChatJoinRequest.Approve(c)`/`Decline(c)` use the actual request user/chat; ensure the application decision and bot rights are appropriate. Missing request objects produce an error.
