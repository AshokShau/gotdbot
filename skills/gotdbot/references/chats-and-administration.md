# Chats and administration

Contents: native chat types; membership; moderation; topics; invite links; profiles; menus and commands.

Source anchors: `chat_admin.go`, `forum_helpers.go`, `bot_profile.go`, `message_actions.go`, `chat_helpers.go`, `gen_helpers.go`, `gen_methods.go`, `gen_types.go`, `gen_options.go`.

## Identify the correct target

`GetChat(chatID)` returns a Chat with a `ChatType` union: private, basic group, supergroup (including channel), or secret chat. `ChatTypeSupergroup.SupergroupId` is a **supergroup ID**, not the chat ID. Use `GetSupergroupId(chatID)` or the returned native type for methods expecting that ID; do not derive it by arbitrary numeric slicing.

`GetChatMember(chatID, MessageSender)` takes a sender union. Use `&MessageSenderUser{UserId: ...}` for a user or `&MessageSenderChat{ChatId: ...}` for a channel/chat sender. Native rights and account restrictions apply to all administrative calls. When implementing a moderation command, verify the command issuer's application authorization as well as the client's Telegram rights.

## Membership and moderation

| Helper | Behavior to preserve |
| --- | --- |
| `ApproveChatJoinRequest`, `DeclineChatJoinRequest` | Process the supplied chat/user request; update methods also expose Approve/Decline. |
| `BanChatMember` | Native sender union, optional ban-until timestamp in `BanChatMemberOpts`. |
| `UnbanChatMember` | Sets status Left; without `OnlyIfBanned`, can remove an existing member. |
| `RestrictChatMember` | Requires permissions; fetches current status to preserve membership. |
| `PromoteChatMember` | Applies `ChatAdministratorRights`; nil rights demote to member. |
| `BanChatSenderChat`, `UnbanChatSenderChat` | Use the chat-sender variant. |
| `GetChatMemberCount` | Resolves basic-group vs supergroup/channel info; rejects non-group chats. |
| `Message.Ban`, `Kick`, `Restrict` | Act on the message's actual SenderId, not a supplied third-party user. |

Use `UnbanChatMemberOpts{OnlyIfBanned: true}` when the intent is only to remove a ban. `RestrictChatMemberOpts.RestrictedUntilDate` is a Unix timestamp; zero means indefinite. Permissions are a complete native structure, so do not assume omitted fields preserve existing rights. Custom administrator titles/tags use `SetChatMemberTag` separately in this baseline.

```go
package recipes

import "github.com/AshokShau/gotdbot"

// Call after validating the moderator's authorization and target account.
func UnbanOnly(c *gotdbot.Client, chatID, userID int64) error {
	return c.UnbanChatMember(chatID, userID, &gotdbot.UnbanChatMemberOpts{OnlyIfBanned: true})
}
```

For pin/delete operations, use message properties when deciding whether to show an action. Deletion for everyone requires `DeleteMessagesOpts{Revoke: true}`; pinning options include silent pin/only-for-self where supported.

## Forum topics

Native forum topic IDs are int32. The General topic uses ID 1 in these helpers:

- `CloseForumTopic(chatID, topicID)` and `ReopenForumTopic`.
- `CloseGeneralForumTopic` and `ReopenGeneralForumTopic`.
- `HideGeneralForumTopic` and `UnhideGeneralForumTopic`.
- `UnpinAllGeneralForumTopicMessages`.

For creation, editing, history, deletion, and other topic operations, inspect the generated `CreateForumTopic`, `EditForumTopic`, `GetForumTopics`, and related methods/opts. Sending into a topic uses `SendMessageOpts.TopicId: &MessageTopicForum{ForumTopicId: topicID}`; draft options instead use `ForumTopicId` directly. Do not substitute a Bot API message thread ID without verifying the corresponding native identifier.

## Invite links

`ExportChatInviteLink(chatID)` replaces the primary invite link and revokes the previous primary link. It is a mutation, not a read-only export. For creating an additional link, inspect `CreateChatInviteLink` and its expiry/member-limit/join-request options. For inspecting existing links, use the appropriate native lookup methods. Preserve the user's intended link behavior.

## Bot profile

After `Start`, the current-account profile helpers validate a cached bot identity for the relevant calls:

- `GetMyName(languageCode)` / `SetMyName(name, *SetBotNameOpts)`.
- `GetMyDescription` / `SetMyDescription` with bot-info description opts.
- `GetMyShortDescription` / `SetMyShortDescription` with short-description opts.
- `SetMyProfilePhoto(InputChatPhoto)`; nil is rejected, so use `RemoveMyProfilePhoto()` to remove it.
- `SetMyDefaultAdministratorRights(rights, forChannels)`; nil clears the defaults. The boolean selects channel or group rights.
- `GetMyDefaultAdministratorRights(forChannels)` reads the corresponding bot info.

`SetMyProfilePhoto` takes `InputChatPhoto`, not `InputFile`; inspect `InputChatPhotoStatic` or another native variant before constructing it. Default administrator rights describe the rights requested on addition, not a promotion in a particular existing chat. Native bot-management methods may also require ownership; read their comments rather than assuming a user account can use every `SetMy*` method.

## Menus and command lists

`GetChatMenuButton(chatID)` / `SetChatMenuButton(chatID, *BotMenuButton)` accept private chat IDs or zero for the default; negative chat IDs are rejected. `BotMenuButton` is a struct in this baseline, not a Bot API type-name string.

`OnCommand` installs a Go handler. It does not publish the bot's menu commands. Native `SetCommands` publishes a list and scope/language options; `GetCommands`/`DeleteCommands` manage that list independently.

```go
package recipes

import "github.com/AshokShau/gotdbot"

func PublishCommands(c *gotdbot.Client) error {
	return c.SetCommands([]gotdbot.BotCommand{
		{Command: "start", Description: "Start the bot"},
		{Command: "help", Description: "Show help"},
	}, &gotdbot.SetCommandsOpts{
		Scope:        &gotdbot.BotCommandScopeDefault{},
		LanguageCode: "en",
	})
}
```

Keep published commands aligned with installed handlers. Explicitly choose language/scope instead of replacing unrelated localized lists.
