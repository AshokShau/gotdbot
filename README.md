# gotdbot

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AshokShau/gotdbot)
[![Go Reference](https://pkg.go.dev/badge/github.com/AshokShau/gotdbot.svg)](https://pkg.go.dev/github.com/AshokShau/gotdbot)
[![License](https://img.shields.io/github/license/AshokShau/gotdbot)](LICENSE)

**gotdbot** is a powerful, pure Go wrapper for [TDLib](https://github.com/tdlib/td) (Telegram Database Library). It provides a simple, idiomatic, and highly performant way to build Telegram clients and bots.

By using `purego`, **gotdbot** eliminates the need for CGO, making your build process cleaner and more portable.

---

## Features

- **Pure Go**: No CGO required. Uses `purego` to dynamically load `libtdjson`.
- **Type-Safe**: 100% generated Go structs and methods for all TDLib types and functions.
- **High Performance**: Direct binding to TDLib's JSON interface for minimal overhead.
- **Composable Filters**: Powerful and easy-to-use filters for update matching.
- **Conversations**: Built-in support for sequential interaction (Ask/Wait for messages).
- **Rich Helpers**: Convenient methods on `Message` and `Client` to simplify common tasks like replying, downloading files, or managing chat members.
- **User & Bot Support**: Handles both Bot API tokens and User authorization (Phone, QR code, 2FA).

---

## 📦 Installation

```bash
go get github.com/AshokShau/gotdbot
```

### Requirements

- **Go**: 1.22 or newer.
- **TDLib**: Compiled `libtdjson` shared library.

#### Quick TDLib Setup
You can use the built-in tool to automatically download the latest precompiled TDLib binaries for your platform:

```bash
go run github.com/AshokShau/gotdbot/scripts/tools
```

---

## Quick Start

Here's a simple Echo Bot to get you started:

```go
package main

import (
	"log"
	"github.com/AshokShau/gotdbot"
)

func main() {
	// Initialize the client
	bot, err := gotdbot.NewClient(12345, "YOUR_API_HASH", "YOUR_BOT_TOKEN", nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Register a command handler
	bot.OnCommand("start", func(c *gotdbot.Client, u *gotdbot.Message) error {
		_, err := u.ReplyText(c, "Hello! I am powered by gotdbot.", nil)
		return err
	})

	// Echo all other messages
	bot.OnMessage(func(c *gotdbot.Client, u *gotdbot.Message) error {
		_, err := c.ForwardMessages(u.ChatId, u.ChatId, []int64{u.Id}, &gotdbot.ForwardMessagesOpts{SendCopy: true})
		return err
	}, nil)

	// Start the bot
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}
```

---

## Advanced Usage

### Composable Filters
Filters allow you to precisely target which updates your handlers should process. They are located in the `filters` sub-packages.

```go
import (
    "github.com/AshokShau/gotdbot/filters/message"
)

// ...

// Only handle incoming text messages in private chats
bot.OnMessage(handlePrivate, message.And(message.Private, message.Incoming))
```

Available filter packages:
- `github.com/AshokShau/gotdbot/filters/message` - Filters for `gotdbot.Message`
- `github.com/AshokShau/gotdbot/filters/callbackquery` - Filters for `gotdbot.UpdateNewCallbackQuery`

### Conversations
The `Ask` method makes it easy to handle multi-step interactions.

```go
bot.OnCommand("rename", func(c *gotdbot.Client, u *gotdbot.Message) error {
    u.ReplyText(c, "What is your new name?", nil)
	
    res, err := c.Ask(u.ChatId, &gotdbot.WaitMessageOpts{
    Timeout: 30 * time.Second,
    })
    if err != nil {
    return err
    }
    
    u.ReplyText(c, "Nice to meet you, " + res.GetText(), nil)
    return nil
})
```

### Message Helpers
**gotdbot** adds many helpful methods directly to the `Message` struct to make your code more readable:

```go
msg.ReplyPhoto(c, gotdbot.InputFileLocal{Path: "image.png"}, nil)
msg.Delete(c, true)
msg.Pin(c, false, false)
msg.React(c, []gotdbot.ReactionType{&gotdbot.ReactionTypeEmoji{Emoji: "👍"}}, nil)
```

---

## Examples

Explore the [examples](./examples) directory for more detailed implementations:
- [**Echo Bot**](./examples/echoBot): Basic functionality.
- [**Conversation Bot**](./examples/conversation): Using the `Ask` API.
- [**Multi Bot**](./examples/echoMultiBot): Managing multiple clients.

---

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

Distributed under the MIT License. See `LICENSE` for more information.
