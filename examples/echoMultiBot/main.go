package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers"
)

var (
	apiID   = int32(6)
	apiHash = "API_HASH"
	tokens  = "7501107071:ABC, 7609688364:DEF" // Comma-separated list of bot tokens to clone

	manager    *gotdbot.ClientManager
	tokenRegex = regexp.MustCompile(`\d{8,11}:[A-Za-z0-9_-]{35}`)
)

func main() {
	manager = gotdbot.NewClientManager("./libtdjson.so.1.8.63")
	dispatcher := gotdbot.NewDispatcher(nil)

	dispatcher.AddHandler(handlers.NewUpdateNewMessage(func(u *gotdbot.UpdateNewMessage) bool {
		msg := u.Message
		if msg == nil || msg.ForwardInfo == nil {
			return false
		}

		if !msg.IsPrivate() {
			return false
		}

		origin := msg.ForwardInfo.Origin
		if origin == nil {
			return false
		}

		var senderID int64
		switch o := origin.(type) {
		case *gotdbot.MessageOriginUser:
			senderID = o.SenderUserId
		}

		return senderID == 93372553
	}, cloneHandler))

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

	fmt.Println("Bots are running. Press Ctrl+C to stop.")
	manager.Idle()
}

func cloneHandler(c *gotdbot.Client, ctx *gotdbot.Context) error {
	msg := ctx.EffectiveMessage
	text := msg.GetText()
	match := tokenRegex.FindString(text)
	if match == "" {
		_, _ = msg.ReplyText(c, "No valid bot token found in the forwarded message. Please forward a message that contains your bot token.", nil)
		return gotdbot.EndGroups
	}
	botToken := match

	go func() {
		reply, err := msg.ReplyText(c, "Cloning "+match, nil)
		if err != nil {
			return
		}

		clientConfig := gotdbot.DefaultClientConfig()
		clientConfig.Dispatcher = c.Dispatcher
		clientConfig.DatabaseDirectory = "db_" + strings.Split(botToken, ":")[0]

		newBot, err := manager.RegisterClient(apiID, apiHash, botToken, clientConfig)
		if err != nil {
			log.Printf("Failed to register new bot: %v", err)
			_, _ = msg.EditText(c, "Failed to register your bot. Is the token valid?", nil)
			return
		}

		me, err := newBot.GetMe()
		if err != nil {
			log.Printf("Failed to get me for new bot: %v", err)
			_, _ = reply.EditText(c, "Bot started, but failed to retrieve its information.", nil)
			return
		}

		_, _ = reply.EditText(c, me.FirstName+" has been successfully started", nil)
	}()

	return nil
}
