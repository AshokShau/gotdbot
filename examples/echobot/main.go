package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AshokShau/gotdbot/client"
	"github.com/AshokShau/gotdbot/types"
)

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := "YOUR_BOT_TOKEN"

	bot := client.NewClient(apiID, apiHash, botToken, &client.ClientConfig{LibraryPath: "./libtdjson.so"})
	client.SetLogVerbosityLevel(2) // 1-5

	bot.OnUpdateAuthorizationState(func(c *client.Client, u *types.UpdateAuthorizationState) error {
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

	bot.OnUpdateNewMessage(func(c *client.Client, u *types.UpdateNewMessage) error {
		msg := u.Message
		_, err := c.ForwardMessages(msg.ChatId, msg.ChatId, []int64{msg.Id}, true, false, nil)
		if err != nil {
			return err
		}
		return nil
	}, nil, 0)

	bot.Run()
	
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Stopping...")
	bot.Stop()
}
