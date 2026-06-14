package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"log"

	"github.com/AshokShau/gotdbot"
)

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := ""

	bot, err := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientOpts{LibraryPath: "./libtdjson.so.1.8.65"})
	if err != nil {
		panic(err)
	}

	bot.OnCommand("start", func(c *gotdbot.Client, u *gotdbot.Message) error {
		_, err := c.SendTextMessage(u.ChatId, "Hello! I am a conversation bot powered by gotdbot "+gotdbot.Version+". What's your name?", &gotdbot.SendTextMessageOpts{ReplyToMessageID: u.Id})
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return err
		}

		res, err := c.Ask(u.ChatId, &gotdbot.WaitMessageOpts{
			Filter: func(m *gotdbot.Message) bool {
				_, ok := m.SenderId.(*gotdbot.MessageSenderUser)
				return ok
			},
		})

		if err != nil {
			log.Printf("Error asking: %v", err)
			return err
		}

		text := res.GetText()
		_, err = c.SendTextMessage(u.ChatId, "Hello "+text+"!", &gotdbot.SendTextMessageOpts{ReplyToMessageID: res.Id})
		return err
	})

	err = bot.Start()
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me, _ := bot.GetMe()
	username := ""
	if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
		username = me.Usernames.ActiveUsernames[0]
	}
	bot.Logger.Info("Logged in", "username", username, "id", me.Id)
	bot.Idle()
}
