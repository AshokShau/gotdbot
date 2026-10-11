# GotdBot

A Go library for building Telegram bots and clients with
[TDLib](https://github.com/tdlib/td).

GotdBot provides generated Go bindings for TDLib's methods and types, along with
update handlers, filters, and helpers for everyday bot development. TDLib handles
connections, authentication, and file transfers, while GotdBot provides a familiar
Go interface with Bot API-style conveniences.

[Documentation](https://pkg.go.dev/github.com/AshokShau/gotdbot) ·
[Examples](./examples) ·
[Issues](https://github.com/AshokShau/gotdbot/issues)

## Features

- Generated TDLib bindings with typed parameters, responses, and native documentation.
- Message and media helpers, formatting, keyboards, and callback handling.
- Command handlers, composable filters, and handler groups.
- Concurrent update processing with configurable workers and panic recovery.
- File uploads, downloads, and caching through TDLib.
- Multiple accounts in one application.
- Runtime loading of the TDLib shared library through purego.

## Getting started

Install the library:

```sh
go get github.com/AshokShau/gotdbot
```

You will need:

- **Go 1.25 or later.**
- A **TDLib 1.8.68** shared library, matching the generated bindings.
- A Telegram API ID and hash from [my.telegram.org](https://my.telegram.org).
- A bot token from [@BotFather](https://t.me/BotFather) when using a bot account.

Install or [build TDLib](https://github.com/tdlib/td#building), then set the shared
library path and your credentials in your application. The examples below show
how to configure a client. Use a separate database directory for each account.

##### Download TDLib dynamic library

```sh
go run github.com/AshokShau/gotdbot/scripts/tools
```

### Example bots

The [examples directory](./examples) contains applications to build on:

- [Echo bot](./examples/echoBot): Commands, inline keyboards, and message filters.
- [Auto-delete bot](./examples/autoDeleteBot): Delete outgoing messages after 30 seconds and cancel timers when messages are permanently deleted.
- [Conversation bot](./examples/conversation): Ask a question and handle the reply.
- [Multiple bots](./examples/echoMultiBot): Run and manage several bot accounts together.

Set your API ID, API hash, bot token, and TDLib library path in an example before
running it from its directory:

```sh
cd examples/echoBot
go run .
```

## Documentation

The [Go package reference](https://pkg.go.dev/github.com/AshokShau/gotdbot)
documents the public API. Refer to the
[TDLib documentation](https://core.telegram.org/tdlib/docs/) for native types,
identifiers, and account permissions.

## Contributing

Bug reports, examples, documentation improvements, and pull requests are welcome.
Please include a reproducible example when reporting an issue.

To check your changes:

```sh
go test ./...
go test -race ./...
go vet ./...
```

Ordinary tests run without Telegram credentials or an installed TDLib library.
They cover request correlation and retries, send completion, shutdown, handlers,
conversations, memory/disk update queues, JSON bindings, filters, logging, and
QR encoding. Native transport tests use in-process stubs.

For repeated concurrency checks, fuzzing, and performance measurements:

```sh
go test -race -shuffle=on -count=20 -timeout=120s ./...
go test -run='^$' -fuzz=FuzzInt64SliceRoundTrip -fuzztime=10s .
go test -run='^$' -fuzz=FuzzInt64SliceDecode -fuzztime=10s .
go test -run='^$' -fuzz=FuzzUnmarshalWithClient -fuzztime=10s .
go test -run='^$' -fuzz=FuzzUnparseEntitiesUnicode -fuzztime=10s .
go test -run='^$' -bench=. -benchmem .
```

### Code generation

The bindings are generated from a TDLib JSON schema. To regenerate them from the
configured schema source, run this command from the repository root:

```sh
go generate .
```

To use a local schema snapshot:

```sh
go run ./internal/gen -schema /path/to/tdlib.json
```

Keep the schema version and installed TDLib library in sync. Use the same schema
snapshot when you need reproducible output, and make binding changes in the
generator rather than editing generated files directly.

## License

GotdBot is licensed under the [MIT License](./LICENSE).
