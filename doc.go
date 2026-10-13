// Package gotdbot provides a Go client for Telegram bots and user accounts,
// combining TDLib's typed API with bot-style send, reply, edit, and handler methods.
// It requires a TDLib shared library and Telegram API credentials.
//
// # Accounts
//
// NewClient accepts an API ID, API hash, and either a bot token or phone number.
// ClientOpts configures the shared library path, account storage, authentication,
// logging, request timeouts, and update workers. Each account needs its own
// database directory. A nil configuration uses DefaultClientConfig.
//
// Register handlers before calling Client.Start, which waits for authentication.
// Client.Me contains the account identity after Start succeeds. Client.Idle waits
// for a termination signal, and Client.Close releases the account's resources.
// ClientManager runs multiple accounts through a shared TDLib receiver.
//
// # Messages
//
// Client.SendMessage and the typed media senders create messages. Message.ReplyText
// and other Reply methods use the message's chat and reply destination. Chat send
// methods use the chat's identifier. Optional settings use pointer structs; nil
// selects defaults, and helpers leave caller-owned options unchanged.
//
// Text and captions support HTML, Markdown, MarkdownV2, and explicit entities.
// FormattedText.Format renders entities as markup; User.Mention formats a user link.
// Rich messages support structured blocks, rich Markdown, and rich HTML.
// Message edit methods accept text, captions, files, or InputMessageContent.
// InputFileLocal uploads a local file; InputFileId references a TDLib file.
// Message.GetFile exposes attached media, and Message.Download downloads it.
// Chat, message, and file identifiers are TDLib identifiers.
//
// # Handlers
//
// Client.OnCommand registers a command without its prefix, such as "start".
// CommandHandler.SetTriggers selects accepted prefix characters, such as "/!.".
// Client.OnMessage handles new messages and business message edits. Client.OnUpdate
// accepts any update, while generated OnUpdate methods handle specific update types.
// Reusable predicates are available in filters/message and filters/callbackquery;
// the filters package combines predicates with AllFilters, AnyFilters, and NotFilter.
//
// Handler groups run in ascending order, with handlers in registration order.
// A matching handler normally ends its group. ContinueHandlers advances to another
// handler, ContinueGroups advances to another group, and EndGroups ends processing.
// Configure handler options, filters, and triggers before processing updates.
//
// # Updates and requests
//
// Typed updates expose answer, message lookup, and edit methods where applicable.
// Client.GetUpdateMessage resolves a message carried or identified by an update.
// Client.UpdateChatID, UpdateMessageID, and UpdateSenderID extract update identifiers.
// Client.Send and SendWithContext execute typed TDLib requests directly.
// Native methods group booleans, nullable values, and supported integer and string
// defaults in Opts structs. Functions with one argument accept it directly.
// Client.Init starts the client for requests that do not require authentication.
// Bindings can be regenerated with go generate; the installed TDLib shared library
// must support the schema used to generate the bindings.
package gotdbot
