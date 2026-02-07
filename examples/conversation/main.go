package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	botToken := "BOT_TOKEN"

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.61"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	dispatcher.AddHandler(handlers.NewCommand("start", func(ctx *ext.Context) error {
		msg := ctx.EffectiveMessage
		c := ctx.Client
		_, err := msg.ReplyText(c, "Welcome! Use /survey to start the survey.\nSend /cancel to stop talking to me", nil)
		return err
	}))

	dispatcher.AddHandler(handlers.NewCommand("survey", func(ctx *ext.Context) error {
		chatId := ctx.EffectiveChatId
		msg := ctx.EffectiveMessage
		c := ctx.Client

		timeOut := 30 * time.Second
		stopFilter := filters.Text.And(filters.SenderID(msg.FromID())).And(filters.Command("cancel"))

		_, err := msg.ReplyText(c, "What is your name?", nil)
		if err != nil {
			return err
		}

		nameMsg, err := ctx.Ask(chatId, &ext.WaitMessageOpts{Timeout: timeOut, Filter: filters.Text.And(filters.SenderID(msg.FromID())), CancellationFilter: stopFilter})
		if err != nil {
			_, _ = msg.ReplyText(c, err.Error(), nil)
			return nil
		}

		_, err = msg.ReplyText(c, fmt.Sprintf("I see! Please send me a photo of yourself, %s.", nameMsg.Text()), &gotdbot.SendTextMessageOpts{ReplyMarkup: &gotdbot.ReplyMarkupForceReply{InputFieldPlaceholder: "Send a picture"}})
		if err != nil {
			return err
		}

		picMsg, err := ctx.Ask(chatId, &ext.WaitMessageOpts{Timeout: timeOut, Filter: filters.Photo.And(filters.SenderID(msg.FromID())), CancellationFilter: stopFilter})
		if err != nil {
			if errors.Is(err, ext.ConversationCancelled) {
				_, _ = msg.ReplyText(c, "Survey cancelled. Send /survey to start again.", nil)
				return nil
			}

			_, _ = msg.ReplyText(c, "Timeout !", nil)
			return nil
		}

		_, err = msg.ReplyPhoto(c, picMsg.RemoteFileID(), &gotdbot.SendPhotoOpts{Caption: fmt.Sprintf("Nice to meet you, %s!", nameMsg.Text())})
		if err != nil {
			return err
		}

		return nil
	}))

	dispatcher.Start()
	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}
