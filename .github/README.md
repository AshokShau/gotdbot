# gotdbot

A high-performance, type-safe Telegram Bot and Client framework for Go built on TDLib.

## Installation

```bash
go get github.com/AshokShau/gotdbot
```

Requires `libtdjson` (TDLib shared library).

## Quick Start Example

```go
package main

import (
	"log"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters/message"
)

func main() {
	bot, err := gotdbot.NewClient(12345, "your_api_hash", "your_bot_token", &gotdbot.ClientOpts{
		LibraryPath: "./libtdjson.so",
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := c.SendTextMessage(msg.ChatId, "Hello from gotdbot!", nil)
		return err
	}).AddFilter(message.Private)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}

	bot.Idle()
}
```

## Commands & Filters

Register commands and chain filters cleanly:

```go
bot.OnCommand("admin", adminHandler).
AddFilter(message.Group).
AddFilter(customFilter)
```

### Custom Filters

```go
func SudoOnly(msg *gotdbot.Message) bool {
return msg.SenderID() == 12345678
}

bot.OnCommand("sudo", sudoHandler).AddFilter(SudoOnly)
```

## Handling Updates

```go
// Listen to specific update types
bot.OnUpdateMessageEdited(func(c *gotdbot.Client, u *gotdbot.UpdateMessageEdited) error {
log.Printf("Message %d edited", u.MessageId)
return nil
}, nil)

// Raw message handler
bot.OnMessage(func(c *gotdbot.Client, msg *gotdbot.Message) error {
log.Println("New message:", msg.GetText())
return nil
}, message.Private)
```

## Code Generation

To regenerate TDLib bindings, run:

```bash
go generate ./...
```
