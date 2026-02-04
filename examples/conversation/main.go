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
		c := ctx.Client

		// Ask Name
		_, err := msg.ReplyText(c, "What is your name?", nil)
		if err != nil {
			return err
		}

		// Wait for response (Name)
		nameMsg, err := ctx.WaitForMessage(chatId, filters.Text.And(filters.SenderID(msg.FromID())), 30*time.Second)
		if err != nil {
			_, _ = msg.ReplyText(c, "Timed out waiting for name.", nil)
			return nil
		}

		name := nameMsg.Text()
		// Ask Age
		_, err = msg.ReplyText(c, fmt.Sprintf("Nice to meet you, %s! How old are you?", name), nil)
		if err != nil {
			return err
		}

		// Wait for response (Age)
		ageMsg, err := ctx.WaitForMessage(chatId, filters.Text.And(filters.SenderID(msg.FromID())), 30*time.Second)
		if err != nil {
			_, _ = msg.ReplyText(c, "Timed out waiting for age.", nil)
			return nil
		}

		age := ageMsg.Text()

		// Ask Fevorite Color
		_, err = msg.ReplyText(c, fmt.Sprintf("Great! Finally, what is your favorite color, %s?", name), &gotdbot.SendTextMessageOpts{ReplyMarkup: &gotdbot.ReplyMarkupForceReply{InputFieldPlaceholder: "Type your favorite color"}})
		if err != nil {
			return err
		}

		// Wait for response (Color)
		colorMsg, err := ctx.WaitForMessage(chatId, filters.Text.And(filters.SenderID(msg.FromID())), 30*time.Second)
		if err != nil {
			_, _ = msg.ReplyText(c, "Timed out waiting for favorite color.", nil)
			return nil
		}

		// Finish
		color := colorMsg.Text()
		summary := fmt.Sprintf("Thank you for completing the survey, %s!\nAge: %s\nFavorite Color: %s", name, age, color)
		_, err = msg.ReplyText(c, summary, nil)
		return err
	}))

	dispatcher.Start()

	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}
