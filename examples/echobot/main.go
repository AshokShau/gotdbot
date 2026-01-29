package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	botToken := "YOUR_BOT_TOKEN"

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	dispatcher.AddHandler(handlers.NewCommand("start", func(ctx *ext.Context) error {
		_, err := ctx.Reply("Hello! I am an echo bot powered by gotdbot.", nil)
		return err
	}))

	dispatcher.AddHandler(handlers.NewCommand("go", func(ctx *ext.Context) error {
		reply := fmt.Sprintf("Hello! There are currently %d goroutines running.", runtime.NumGoroutine())
		_, err := ctx.Reply(reply, nil)
		return err
	}))

	dispatcher.AddHandler(handlers.NewMessage(filters.Text.And(filters.Incoming), func(ctx *ext.Context) error {
		_, err := ctx.Client.ForwardMessages(ctx.EffectiveChatId, ctx.EffectiveChatId, []int64{ctx.EffectiveMessage.Id}, true, false, nil)
		return err
	}))

	dispatcher.Start()

	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me := bot.Me()
	if me != nil {
		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		log.Printf("Logged in as @%s (%d)", username, me.Id)
	}

	bot.Idle()
}
