package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/AshokShau/gotdbot"
)

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := "BOT_TOKEN"

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2) // 1-5

	bot.OnUpdateAuthorizationState(func(c *gotdbot.Client, u *gotdbot.UpdateAuthorizationState) error {
		if u.AuthorizationState.Type() == "authorizationStateReady" {
			me, err := c.GetMe()
			if err != nil {
				log.Fatalf("Failed to get me: %v", err)
				return nil
			}
			fmt.Printf("Logged in as: %s %s (%d)\n", me.FirstName, me.LastName, me.Id)
		}
		return nil
	}, nil, 0)

	bot.OnUpdateNewMessage(func(c *gotdbot.Client, u *gotdbot.UpdateNewMessage) error {
		msg := u.Message
		if msg.Content.MessageText != nil && msg.Content.MessageText.Text != nil && msg.Content.MessageText.Text.Text == "/go" {
			reply := fmt.Sprintf("Hello! There are currently %d goroutines running.", runtime.NumGoroutine())
			_, err := c.SendMessage(msg.ChatId, &gotdbot.InputMessageContent{
				InputMessageText: &gotdbot.InputMessageText{
					Text: &gotdbot.FormattedText{
						Text: reply,
					},
				},
			}, nil)
			if err != nil {
				return err
			}
			return nil

		}

		_, err := c.ForwardMessages(msg.ChatId, msg.ChatId, []int64{msg.Id}, true, false, nil)
		if err != nil {
			return err
		}
		return nil
	}, nil, 0)

	bot.OnUpdateNewCallbackQuery(func(c *gotdbot.Client, msg *gotdbot.UpdateNewCallbackQuery) error {
		fmt.Printf("Callback query received: %s\n", msg.Payload.CallbackQueryPayloadData)
		return nil
	}, nil, 0)

	bot.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Stopping...")
	bot.Stop()
}
