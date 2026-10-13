# Setup and lifecycle

Contents: prerequisites; runnable starter; configuration; user login; lifecycle; multiple accounts; logging.

Source anchors: `go.mod`, `consts.go`, `client.go`, `client_opts.go`, `client_auth.go`, `manager.go`, `internal/tdjson/tdjson.go`, `scripts/tools/get_tdjson.go`, `logger/logger.go`.

## Prerequisites and starter

The reference baseline uses Go 1.25+, GotdBot v0.9.6, and TDLib v1.8.68. Install the application dependency with `go get github.com/AshokShau/gotdbot@<chosen-version>`, preserving an existing application's version. The native TDLib schema must match the generated bindings. Read `TDLibVersion` in the source rather than selecting an unrelated latest native build.

The downloader in this snapshot targets v1.8.68:

```sh
go run github.com/AshokShau/gotdbot/scripts/tools
```

It downloads and extracts native libraries into the current directory. To use another version, obtain or build a matching TDLib library. An archive for an OS/architecture is not proof that the Go runtime loader supports that target; check the selected purego version and library runtime dependencies when deploying to a new platform.

`NewClient` loads the library immediately. Prefer an explicit `LibraryPath` (an absolute path avoids working-directory surprises). Defaults are `tdjson.dll` on Windows, `libtdjson.dylib` on macOS, and a local `libtdjson.so` or matching `libtdjson.so.*` on Linux-like systems. This snapshot loads one library per process; later configurations do not switch the already-loaded library.

Get an API ID/hash from Telegram's application registration and a bot token from BotFather. Use configuration or environment variables, not committed secrets. The [starter](../assets/bot/main.go) expects `TELEGRAM_API_ID`, `TELEGRAM_API_HASH`, `TELEGRAM_BOT_TOKEN`, and `TDLIB_LIBRARY`; optional `GOTDBOT_DATA_DIR` selects its persistent storage parent.

Copy it into a new application's `main.go`, initialize a module if needed, add GotdBot, and run `go build .` before `go run .`. From this repository, it can be built with `go build -o /tmp/gotdbot-starter ./skills/gotdbot/assets/bot`.

## Configuration semantics

`DefaultClientConfig()` returns a fresh config. `NewClient` shallow-copies a supplied config and fills selected missing defaults; it does not turn every zero boolean into the default value. Configure nested pointers before startup and do not mutate them concurrently.

| Setting | Baseline behavior and use |
| --- | --- |
| `DatabaseDirectory` | Defaults to `database`. Set a distinct, persistent path per account; preserve it across restarts. |
| `FilesDirectory` | Empty in the Go default; TDLib resolves its native default. Set an explicit path when needed. |
| `DatabaseEncryptionKey` | String passed as bytes to TDLib. Retain the key when reopening an encrypted database. |
| `UseFileDatabase`, `UseChatInfoDatabase`, `UseMessageDatabase` | `*bool`; nil defaults to true. Set `gotdbot.Bool(false)` to disable. |
| `UseSecretChats`, `UseTestDC` | Ordinary booleans default false. Secret chats require appropriate account/API support. |
| `AuthorizationTimeout` | 60 seconds when missing/nonpositive. Interactive login may need more time. |
| `RequestTimeout` | 30 seconds when missing/nonpositive; bounds request retries and send completion together. |
| `UpdateWorkers` | Defaults to 2 when nonpositive. Use 1 for serial processing; blocking conversations need spare workers. |
| `UpdateQueueSize` | Defaults to 1000 when nonpositive. Overflow spills to temporary files; this is not durable restart recovery. |
| `ParseMode` | Empty means plain text by default. Recognized modes are HTML, Markdown, MarkdownV2. |
| `LoadMessagesBeforeReply` | Optional message lookup before replying. Read reply helpers when diagnosing lookup failures. |
| `AutoRetry` | Defaults to no enabled retry features. Options are `ChatNotFound`, `MessageNotFound`, `MaxFloodWait`. |
| `TDLibOptions` | Generated typed option fields applied during parameter setup. Pointer fields distinguish unset from zero. |
| `PanicHandler`, `ErrorHandler` | Application update callbacks. ErrorHandler can return group-control sentinels. |

`Logger` is the Go application logger. `LogStream`/`LogVerbosityLevel` are separate configuration fields; inspect their actual call sites before promising TDLib logging behavior. The runtime loader in this baseline explicitly disables TDLib's native log stream on initialization.

## Bot and user authentication

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"time"
)

func NewUser(apiID int32, apiHash, phone, libraryPath string) (*gotdbot.Client, error) {
	cfg := gotdbot.DefaultClientConfig()
	cfg.LibraryPath = libraryPath
	cfg.DatabaseDirectory = "./data/user"
	cfg.FilesDirectory = "./data/user-files"
	cfg.AuthorizationTimeout = 5 * time.Minute
	return gotdbot.NewClient(apiID, apiHash, phone, cfg)
}
```

The token detector expects digits, a colon, and a token body. Other strings are treated as phone numbers. Validate environment values so a missing token cannot silently start an interactive phone login. User authentication reads the code and two-factor password from stdin; registration also prompts for names. `QrMode: true` requests QR authentication for a user and prints the confirmation link/QR.

Do not invent an `AuthCallback` or password field: this version does not expose either on `ClientOpts`. Headless first-login workflows need a supported design around the real authorization implementation or a library change. Persisted authorized sessions avoid repeated first-login prompts. Telegram remains responsible for account authorization rules.

## Lifecycle

Use `NewClient` → register handlers → `Start` (or `StartContext`) → `Idle` or an application-owned lifetime → `Close`.

- `Start` waits for authorization and caches `Me`. Check its error before using identity-dependent methods.
- `StartContext` cancellation stops waiting; it does not close the running instance. Close it if aborting startup. Cancellation also does not cancel the interactive stdin reader.
- `Init` starts the receiver/processors without waiting for auth. Use it only for functions documented as callable before auth; it does not authorize normal sends.
- `Me` is library-owned, refreshed by updates, and must not be modified by the application. Fetch identity with `GetMe` when a fresh value is needed.
- `Idle` waits for SIGINT/SIGTERM or closure, then closes the client. An existing HTTP/service host can manage shutdown instead.
- `Close` is idempotent, can be called from a handler, and performs bounded native shutdown. It does not wait for in-flight application handlers; drain application jobs separately when needed.
- Closure unblocks pending requests/waiters with `ErrClientClosed`. A closed client is not a restartable instance.

## Multiple accounts

There is one active TDLib receiver manager per process. Use a shared `ClientManager`; do not create competing receive loops. `GetDefaultManager` reuses the receiver owner when present. A client cannot belong to two managers.

`RegisterClient` creates, attaches, and starts a client immediately. Use `NewClient` + handler registration + `AddClient` + `Start` when handlers must be ready before auth. Supply an application account key for storage; do not put tokens in directory names or logs.

```go
package recipes

import (
	"github.com/AshokShau/gotdbot"
	"path/filepath"
)

func AddBot(m *gotdbot.ClientManager, apiID int32, apiHash, token, accountDir string) (*gotdbot.Client, error) {
	cfg := gotdbot.DefaultClientConfig()
	cfg.LibraryPath = m.LibraryPath
	cfg.DatabaseDirectory = filepath.Join(accountDir, "database")
	cfg.FilesDirectory = filepath.Join(accountDir, "files")
	c, err := gotdbot.NewClient(apiID, apiHash, token, cfg)
	if err != nil {
		return nil, err
	}
	c.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, "Ready", nil)
		return err
	})
	if err := m.AddClient(c); err != nil {
		c.Close()
		return nil, err
	}
	if err := c.Start(); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}
```

Create the manager with `NewClientManager(libraryPath)` and defer `Stop` in the owning application. `manager.Idle()` handles signal shutdown. `RemoveClient` only detaches; it does not close the native instance. `GetClients` returns a snapshot in unspecified order. Keep each account's messages/files associated with its client.

## Logging

Use `logger.New(logger.WithLevel(slog.LevelDebug), logger.WithOutput(writer))`, `logger.From(existingSlogLogger)`, or `logger.Disabled()`. Structured calls such as `c.Logger.Info("sent", "chat_id", chatID)` are preferred for metadata. Never log full tokens, authentication codes, passwords, or payment credentials.
