package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"

	bot, err := gotdbot.NewClient(apiID, apiHash, "", &gotdbot.ClientConfig{
		LibraryPath:           "./libtdjson.so.1.8.61",
		UseFileDatabase:       gotdbot.Bool(true),
		AuthorizationTimeout:  2 * time.Minute,
		QrMode:                true,               // Enable QR code login
		DatabaseEncryptionKey: "my_secret_key_29", // Optional: Set a key to encrypt the local database
	})

	if err != nil {
		panic(err)
	}

	gotdbot.SetTdlibLogStreamEmpty()
	dispatcher := bot.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("hi", func(ctx *gotdbot.Context) error {
		_, err := ctx.EffectiveMessage.ReplyText(ctx.Client, "Hi, this is from gotdbot!", nil)
		return err
	}))

	err = bot.Start()
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me := bot.Me()
	if me != nil {
		fmt.Printf("Current user: %s (ID: %d)\n", me.FirstName, me.Id)
	}

	bot.Idle()
}
