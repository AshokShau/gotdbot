# gotdbot

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AshokShau/gotdbot)
[![Go Reference](https://pkg.go.dev/badge/github.com/AshokShau/gotdbot.svg)](https://pkg.go.dev/github.com/AshokShau/gotdbot)
[![License](https://img.shields.io/github/license/AshokShau/gotdbot)](LICENSE)

**gotdbot** is a powerful, pure Go wrapper for [TDLib](https://github.com/tdlib/td) (Telegram Database Library). It provides a simple and idiomatic way to build Telegram clients and bots using Go.

Current Version: **v0.5.0** and TDLib **v1.8.62**.

---

## Features

- **Pure Go**: No CGO mess in your application code (uses `purego` to load `libtdjson`).
- **High Performance**: Direct binding to TDLib's JSON interface.
- **Dispatcher System**: Built-in update dispatcher with middleware-style handlers.
- **Filters**: Powerful and composable filters for message matching.
- **Type-Safe**: Fully generated Go structs for all TDLib types and methods.
- **Context Aware**: Easy access to effective messages, chats, and users.

---

## Requirements

* **Go**: version 1.22 or newer
* **TDLib**: the compiled `libtdjson` shared library

  * Follow the official [TDLib build instructions](https://tdlib.github.io/td/build.html?language=Go)
  * Make sure the library is accessible on your system:

    * `libtdjson.so` (Linux)
    * `libtdjson.dylib` (macOS)
    * `tdjson.dll` (Windows)
  * **Alternative**: run

    ```bash
    go run github.com/AshokShau/gotdbot/scripts/tools@latest
    ```

    to automatically download the latest precompiled TDLib binaries.

---

## Installation

```bash
go get github.com/AshokShau/gotdbot
```

---

## Examples

Sample bots can be found in the [examples](./examples) directory.

- [Echo Bot](./examples/echobot): A simple bot that echoes text messages.
- [User Bot](./examples/userbot): Example of running a user account automation.

---

## Contributing

Contributions are welcome! Please open issues for bugs or feature requests.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Inspired by

- [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
- [pytdbot](https://github.com/pytdbot/client)
