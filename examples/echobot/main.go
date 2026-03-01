package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	botToken := "BOT_TOKEN"

	bot, err := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.62"})
	if err != nil {
		panic(err)
	}

	dispatcher := bot.Dispatcher
	gotdbot.SetTdlibLogVerbosityLevel(2)
	var startTime = time.Now()

	dispatcher.AddHandler(handlers.NewCommand("start", func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		kb := &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{
					{
						Text: "GoTDBot GitHub",
						TypeField: &gotdbot.InlineKeyboardButtonTypeUrl{
							Url: "https://github.com/AshokShau/gotdbot",
						},
					},
				},
			},
		}

		content := &gotdbot.InputMessageText{
			Text: &gotdbot.FormattedText{
				Text: "Hello! I am an echo bot powered by gotdbot " + gotdbot.Version,
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyTo: &gotdbot.InputMessageReplyToMessage{
				MessageId: ctx.EffectiveMessage.Id,
			},
			ReplyMarkup: kb,
		}

		_, err := c.SendMessage(ctx.EffectiveChatId, content, opts)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
		return nil
	}))

	dispatcher.AddHandler(handlers.NewUpdateDeleteMessages(nil, func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		update := ctx.Update.UpdateDeleteMessages
		log.Printf("Messages deleted: %v (ChatID %d)", update.MessageIds, update.ChatId)
		return nil
	}))

	dispatcher.AddHandler(handlers.NewCommand("go", func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		uptime := time.Since(startTime).Round(time.Second)
		reply := fmt.Sprintf(
			"🟢 Go runtime stats\n\n"+
				"• Goroutines : %d\n"+
				"• CPUs       : %d\n"+
				"• GOMAXPROCS : %d\n"+
				"• Uptime     : %s\n\n"+
				"🧠 Memory\n"+
				"• Alloc      : %.2f MB\n"+
				"• HeapAlloc  : %.2f MB\n"+
				"• Sys        : %.2f MB\n"+
				"• GC cycles  : %d",
			runtime.NumGoroutine(),
			runtime.NumCPU(),
			runtime.GOMAXPROCS(0),
			uptime,
			float64(m.Alloc)/1024/1024,
			float64(m.HeapAlloc)/1024/1024,
			float64(m.Sys)/1024/1024,
			m.NumGC,
		)

		_, err := c.SendTextMessage(ctx.EffectiveChatId, reply, &gotdbot.SendTextMessageOpts{ReplyToMessageID: ctx.EffectiveMessage.Id})
		return err
	}))

	dispatcher.AddHandler(handlers.NewMessage(filters.Incoming, func(c *gotdbot.Client, ctx *gotdbot.Context) error {
		_, err := c.ForwardMessages(ctx.EffectiveChatId, ctx.EffectiveChatId, []int64{ctx.EffectiveMessage.Id}, &gotdbot.ForwardMessagesOpts{SendCopy: true})
		return err
	}))

	err = bot.Start()
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me := bot.Me()
	username := ""
	if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
		username = me.Usernames.ActiveUsernames[0]
	}
	bot.Logger.Info("Logged in", "username", username, "id", me.Id)
	bot.Idle()
}
