package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"os"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	botToken := "BOT_TOKEN"

	if envID := os.Getenv("API_ID"); envID != "" {
		fmt.Sscanf(envID, "%d", &apiID)
	}
	if envHash := os.Getenv("API_HASH"); envHash != "" {
		apiHash = envHash
	}
	if envToken := os.Getenv("BOT_TOKEN"); envToken != "" {
		botToken = envToken
	}

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.61"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	// /start - Send welcome message with inline keyboard
	dispatcher.AddHandler(handlers.NewCommand("start", func(ctx *ext.Context) error {
		msg := ctx.EffectiveMessage
		c := ctx.Client
		userId := msg.FromID()

		action, err := msg.Action(c, "typing", nil)
		if err != nil {
			ctx.Client.Logger.Error("Failed to create chat action", "err", err)
			return err
		}

		//action.Start() // Start sending action in loop
		//defer action.Stop()
		action.Send() // Send once

		ctx.Client.Logger.Info("Received /start command", "user_id", userId)

		user, err := ctx.Client.GetUser(userId)
		userName := "User"
		if err == nil {
			userName = user.FirstName
		}

		text := fmt.Sprintf("Hello %s! (gotdbot %s) <tg-emoji emoji-id='5346181118884331907'>🤖</tg-emoji>\nHere are some bot commands:\n\n- /keyboard - show keyboard\n- /inline - show inline keyboard\n- /remove - remove keyboard\n- /force - force reply", userName, gotdbot.Version)

		kb := &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{
					{
						Text: "GitHub",
						TypeField: &gotdbot.InlineKeyboardButtonTypeUrl{
							Url: "https://github.com/AshokShau/gotdbot",
						},
						IconCustomEmojiId: 5271604874419647061,
						Style:             &gotdbot.ButtonStylePrimary{},
					},
				},
			},
		}

		_, err = msg.ReplyText(c, text, &gotdbot.SendTextMessageOpts{
			ReplyMarkup: kb,
			ParseMode:   "HTML",
		})

		return err
	}))

	// /inline - Send message with inline keyboard buttons
	dispatcher.AddHandler(handlers.NewCommand("inline", func(ctx *ext.Context) error {
		kb := &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{
					{
						Text: "OwO",
						TypeField: &gotdbot.InlineKeyboardButtonTypeCallback{
							Data: []byte("OwO"),
						},
						Style: &gotdbot.ButtonStylePrimary{},
					},
					{
						Text: "UwU",
						TypeField: &gotdbot.InlineKeyboardButtonTypeCallback{
							Data: []byte("UwU"),
						},
						Style: &gotdbot.ButtonStyleDanger{},
					},
				},
			},
		}

		text, err := ctx.Client.ParseText("This is a Inline keyboard", "Markdown")
		if err != nil {
			ctx.Client.Logger.Error("Failed to parse text", "err", err)
			return err
		}

		content := &gotdbot.InputMessageText{Text: text}
		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: kb,
		}

		message, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		if err != nil {
			ctx.Client.Logger.Error("Failed to send message", "err", err)
			return err
		}
		ctx.Client.Logger.Info("Sent message with inline", "message_id", message.Id)

		return nil
	}))

	// /keyboard - Send message with reply keyboard
	dispatcher.AddHandler(handlers.NewCommand("keyboard", func(ctx *ext.Context) error {
		kb := &gotdbot.ReplyMarkupShowKeyboard{
			Rows: [][]gotdbot.KeyboardButton{
				{
					{
						Text:      "OwO",
						TypeField: &gotdbot.KeyboardButtonTypeText{},
					},
					{
						Text:      "UwU",
						TypeField: &gotdbot.KeyboardButtonTypeText{},
					},
				},
			},
			ResizeKeyboard: true,
			OneTime:        true,
		}

		content := &gotdbot.InputMessageText{
			Text: &gotdbot.FormattedText{
				Text: "This is a keyboard",
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: kb,
		}

		message, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		if err != nil {
			ctx.Client.Logger.Error("Failed to send message", "err", err)
			return err
		}
		ctx.Client.Logger.Info("Sent message with keyboard", "message_id", message.Id)
		return nil
	}))

	// /remove - Remove keyboard
	dispatcher.AddHandler(handlers.NewCommand("remove", func(ctx *ext.Context) error {
		content := &gotdbot.InputMessageText{
			Text: &gotdbot.FormattedText{
				Text: "Keyboards removed",
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkupRemoveKeyboard{},
		}

		_, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// /force - Force reply
	dispatcher.AddHandler(handlers.NewCommand("force", func(ctx *ext.Context) error {
		content := &gotdbot.InputMessageText{
			Text: &gotdbot.FormattedText{
				Text: "This is a force reply",
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkupForceReply{},
		}

		_, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)

		// _, err = ctx.EffectiveMessage.ReplyText(ctx.Client, "This is a force reply", &gotdbot.SendTextMessageOpts{ReplyMarkup: &gotdbot.ReplyMarkupForceReply{}})
		return err
	}))

	// CallbackQuery Handler
	dispatcher.AddHandler(handlers.NewUpdateNewCallbackQuery(nil, func(ctx *ext.Context) error {
		update := ctx.Update.UpdateNewCallbackQuery
		ctx.Client.Logger.Info("Received callback query", "message_id", update.MessageId, "chat_id", update.ChatId)
		var data string
		if update.Payload != nil {
			if p, ok := update.Payload.(*gotdbot.CallbackQueryPayloadData); ok {
				data = string(p.Data)
			}
		}

		if data != "" {
			kb := &gotdbot.ReplyMarkupInlineKeyboard{
				Rows: [][]gotdbot.InlineKeyboardButton{
					{
						{
							Text: "GitHub",
							TypeField: &gotdbot.InlineKeyboardButtonTypeUrl{
								Url: "https://github.com/AshokShau/gotdbot",
							},
							IconCustomEmojiId: 5330237710655306682,
							Style:             &gotdbot.ButtonStyleSuccess{},
						},
					},
				},
			}

			inputContent := &gotdbot.InputMessageText{
				Text: &gotdbot.FormattedText{
					Text: fmt.Sprintf("You pressed %s", data),
				},
			}

			_, err := ctx.Client.EditMessageText(update.ChatId, update.MessageId, inputContent, &gotdbot.EditMessageTextOpts{
				ReplyMarkup: kb,
			})
			if err != nil {
				ctx.Client.Logger.Error("Failed to edit message", "error", err)
			}
		}

		return nil
	}))

	dispatcher.Start()
	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me := bot.Me()
	if me != nil {
		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		bot.Logger.Info("Logged in", "username", username, "id", me.Id)
	}

	bot.Idle()
}
