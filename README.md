# gotdbot

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AshokShau/gotdbot)
[![Go Reference](https://pkg.go.dev/badge/github.com/AshokShau/gotdbot.svg)](https://pkg.go.dev/github.com/AshokShau/gotdbot)
[![License](https://img.shields.io/github/license/AshokShau/gotdbot)](LICENSE)

**gotdbot** is a powerful, pure Go wrapper for [TDLib](https://github.com/tdlib/td) (Telegram Database Library). It provides a simple and idiomatic way to build Telegram clients and bots using Go.

Current Version: **v0.1.0**

## Features

- 🚀 **Pure Go**: No CGO mess in your application code (uses `purego` to load `libtdjson`).
- ⚡ **High Performance**: Direct binding to TDLib's JSON interface.
- 🛠 **Dispatcher System**: Built-in update dispatcher with middleware-style handlers.
- 🔍 **Filters**: Powerful and composable filters for message matching.
- 📦 **Type-Safe**: Fully generated Go structs for all TDLib types and methods.
- 🔄 **Context Aware**: Easy access to effective messages, chats, and users.

## Requirements

- **Go**: 1.22+
- **TDLib**: You need the compiled `libtdjson` shared library.
    - [Build TDLib instructions](https://tdlib.github.io/td/build.html?language=Go)
    - Ensure `libtdjson.so` (Linux), `libtdjson.dylib` (macOS), or `tdjson.dll` (Windows) is available.

## Installation

```bash
go get github.com/AshokShau/gotdbot
```

## Quick Start

Here is a simple echo bot example using the high-level Dispatcher.

```go
package main

import (
	"log"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

func main() {
	// Configuration
	apiID := int32(123456) // Your API ID
	apiHash := "your_api_hash"
	botToken := "your_bot_token"
	
	// Initialize Client
	// Point LibraryPath to your compiled libtdjson library
	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{
		LibraryPath: "./libtdjson.so", 
	})

	// Optional: specific logging
	gotdbot.SetTdlibLogVerbosityLevel(1)

	// Create Dispatcher
	dispatcher := ext.NewDispatcher(bot)

	// Add Handler: Command /start
	dispatcher.AddHandler(handlers.NewCommand("start", func(ctx *ext.Context) error {
      _, err := ctx.Client.SendTextMessage(ctx.EffectiveChatId, "Hello! I am a bot running on gotdbot " + gotdbot.Version, &gotdbot.SendTextMessageOpts{ReplyToMessageID: ctx.EffectiveMessage.Id})
		return err
	}))

	// Add Handler: Echo all incoming text messages
	// Uses filters to match only incoming messages that contain text
	dispatcher.AddHandler(handlers.NewMessage(filters.Text.And(filters.Incoming), func(ctx *ext.Context) error {
		// Forward the message back to the sender
		_, err := ctx.Client.ForwardMessages(ctx.EffectiveChatId, ctx.EffectiveChatId, []int64{ctx.EffectiveMessage.Id}, true, false, nil)
		return err
	}))

	// Start the Dispatcher (background routing)
	dispatcher.Start()

	// Start the Client (TDLib connection)
	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	// Get bot info
	me, err := bot.GetMe()
	if err == nil {
		log.Printf("Logged in as @%s (%d)", me.Usernames.ActiveUsernames[0], me.Id)
	}

	// Wait for interrupt signal to gracefully shutdown
	bot.Idle()
}
```

## Core Concepts

### Client
The `Client` struct is the main entry point. It wraps the TDLib instance and handles the low-level communication (send/receive) with Telegram.

### Dispatcher
The `Dispatcher` organizes how updates are processed. It allows you to register `Handlers` that match specific criteria.

### Handlers & Filters
- **Handlers**: Wrappers that execute a callback function when an update matches. Common handlers include `Command`, `Message`, `CallbackQuery`.
- **Filters**: Logic to determine if a handler should run. Example: `filters.Text`, `filters.Photo`, `filters.Chat(12345)`.

## Contributing

Contributions are welcome! Please open issues for bugs or feature requests.

## License

This project is licensed under the [MIT License](LICENSE).

## Inspired by

- [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
- [pytdbot](https://github.com/pytdbot/client)
