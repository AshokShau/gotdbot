package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

func main() {
	apiID := int32(12345)
	apiHash := "YOUR_API_HASH"
	botToken := "YOUR_BOT_TOKEN"

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	// Register /survey command
	dispatcher.AddHandler(handlers.NewCommand("survey", func(ctx *ext.Context) error {
		chatId := ctx.EffectiveChatId
		msg := ctx.EffectiveMessage

		// Ask Name
		_, err := ctx.Client.SendTextMessage(chatId, "What is your name?", nil)
		if err != nil {
			return err
		}

		// Wait for response (Name)
		nameMsg, err := ctx.WaitForMessage(chatId, filters.Text.And(filters.SenderID(msg.FromID())), 30*time.Second)
		if err != nil {
			ctx.Client.SendTextMessage(chatId, "Timed out waiting for name.", nil)
			return nil
		}

		name := nameMsg.Text()
		// Ask Age
		_, err = ctx.Client.SendTextMessage(chatId, fmt.Sprintf("Nice to meet you, %s! How old are you?", name), nil)
		if err != nil {
			return err
		}

		// Wait for response (Age)
		ageMsg, err := ctx.WaitForMessage(chatId, filters.Text, 30*time.Second)
		if err != nil {
			ctx.Client.SendTextMessage(chatId, "Timed out waiting for age.", nil)
			return nil
		}

		age := ageMsg.Text()
		// Finish
		_, err = ctx.Client.SendTextMessage(chatId, fmt.Sprintf("Got it! Name: %s, Age: %s", name, age), nil)
		return err
	}))

	dispatcher.Start()

	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}
