package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	tokens := os.Getenv("TOKENS") // Comma-separated bot tokens
	if tokens == "" {
		log.Fatal("TOKENS environment variable is empty")
	}

	manager := gotdbot.NewClientManager("./libtdjson.so.1.8.62")
	dispatcher := gotdbot.NewDispatcher(nil)

	dispatcher.AddHandler(handlers.NewCommand("start", func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		_, err := c.SendTextMessage(ctx.EffectiveChatId, "Hello! I am a bot managed by ClientManager.", nil)
		return err
	}))

	dispatcher.AddHandler(handlers.NewCommand("stop", func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		_, err := c.SendTextMessage(ctx.EffectiveChatId, "Stopping this bot...", nil)
		if err != nil {
			return err
		}
		go c.Close()
		return nil
	}))

	dispatcher.AddHandler(handlers.NewCommand("stopall", func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		_, err := c.SendTextMessage(ctx.EffectiveChatId, "Stopping all bots...", nil)
		if err != nil {
			return err
		}
		go manager.Stop()
		return nil
	}))

	dispatcher.AddHandler(handlers.NewUpdateNewMessage(nil, func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		msg := ctx.EffectiveMessage
		_, err := msg.Copy(c, ctx.EffectiveChatId, &gotdbot.SendCopyOpts{
			ReplyMarkup: msg.ReplyMarkup,
		})
		return err
	}))

	splitTokens := strings.Split(tokens, ",")

	for _, token := range splitTokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		config := gotdbot.DefaultClientConfig()
		config.Dispatcher = dispatcher
		config.DatabaseDirectory = "db_" + strings.Split(token, ":")[0]
		_, err := manager.RegisterClient(apiID, apiHash, token, config)
		if err != nil {
			log.Printf("Failed to add client for token %s: %v", token, err)
			continue
		}
	}

	gotdbot.SetTdlibLogVerbosityLevel(2)
	manager.Start()
	fmt.Println("Bots are running. Press Ctrl+C to stop.")
	manager.Idle()
}
