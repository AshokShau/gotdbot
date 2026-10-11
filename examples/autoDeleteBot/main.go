package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/logger"
)

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := ""
	const deleteAfter = 30 * time.Second

	bot, err := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientOpts{
		LibraryPath: "./libtdjson.so.1.8.68",
		Logger:      logger.New(logger.WithLevel(slog.LevelDebug)),
	})
	if err != nil {
		log.Fatal(err)
	}

	var mu sync.Mutex
	timers := make(map[[2]int64]*time.Timer)
	schedule := func(c *gotdbot.Client, msg *gotdbot.Message) {
		// Wait for a final ID; never schedule incoming, pending, or failed messages.
		if msg == nil || !msg.IsOutgoing || msg.SendingState != nil {
			return
		}

		key := [2]int64{msg.ChatId, msg.Id}
		mu.Lock()
		defer mu.Unlock()
		if _, exists := timers[key]; exists {
			return
		}

		timers[key] = time.AfterFunc(deleteAfter, func() {
			mu.Lock()
			if _, exists := timers[key]; !exists {
				mu.Unlock()
				return // A delete update cancelled this timer.
			}
			delete(timers, key)
			mu.Unlock()

			if err := c.DeleteMessages(key[0], []int64{key[1]}, &gotdbot.DeleteMessagesOpts{Revoke: true}); err != nil {
				c.Logger.Debugf("Delete failed: chat=%d message=%d: %v", key[0], key[1], err)
				return
			}
			c.Logger.Debugf("Deleted: chat=%d message=%d", key[0], key[1])
		})
	}

	// Successful sends cover command replies and sends from background tasks.
	bot.OnUpdateMessageSendSucceeded(func(c *gotdbot.Client, u *gotdbot.UpdateMessageSendSucceeded) error {
		schedule(c, u.Message)
		return nil
	}, nil)

	// A separate group lets command handlers receive the same new-message update.
	bot.AddUpdateNewMessageHandlerGroup(func(c *gotdbot.Client, u *gotdbot.UpdateNewMessage) error {
		schedule(c, u.Message)
		return nil
	}, nil, -1)

	bot.OnUpdateDeleteMessages(func(c *gotdbot.Client, u *gotdbot.UpdateDeleteMessages) error {
		if u.FromCache || !u.IsPermanent {
			return nil
		}

		mu.Lock()
		defer mu.Unlock()
		for _, messageID := range u.MessageIds {
			key := [2]int64{u.ChatId, messageID}
			if timer, exists := timers[key]; exists {
				timer.Stop()
				delete(timers, key)
				c.Logger.Debugf("Already deleted; cancelled timer: chat=%d message=%d", u.ChatId, messageID)
			}
		}
		return nil
	}, nil)

	bot.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := c.SendMessage(msg.ChatId, "This message will be deleted after 30 seconds.", nil)
		return err
	})

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
	bot.Idle()

	mu.Lock()
	defer mu.Unlock()
	for key, timer := range timers {
		timer.Stop()
		delete(timers, key)
	}
}
