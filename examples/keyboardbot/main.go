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

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	// /start - Send welcome message with inline keyboard
	dispatcher.AddHandler(handlers.NewCommand("start", func(ctx *ext.Context) error {
		ctx.Client.Logger.Info("Received /start command", "user_id", ctx.EffectiveSenderId.MessageSenderUser.UserId)

		user, err := ctx.Client.GetUser(ctx.EffectiveSenderId.MessageSenderUser.UserId)
		userName := "User"
		if err == nil {
			userName = user.FirstName
		}

		text := fmt.Sprintf("Hello %s!\nHere are some bot commands:\n\n- /keyboard - show keyboard\n- /inline - show inline keyboard\n- /remove - remove keyboard\n- /force - force reply", userName)

		kb := &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]*gotdbot.InlineKeyboardButton{
				{
					{
						Text: "GitHub",
						TypeField: &gotdbot.InlineKeyboardButtonType{
							InlineKeyboardButtonTypeUrl: &gotdbot.InlineKeyboardButtonTypeUrl{
								Url: "https://github.com/AshokShau/gotdbot",
							},
						},
					},
				},
			},
		}

		content := &gotdbot.InputMessageContent{
			InputMessageText: &gotdbot.InputMessageText{
				Text: &gotdbot.FormattedText{
					Text: text,
				},
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkup{
				ReplyMarkupInlineKeyboard: kb,
			},
		}

		_, err = ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// /inline - Send message with inline keyboard buttons
	dispatcher.AddHandler(handlers.NewCommand("inline", func(ctx *ext.Context) error {
		kb := &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]*gotdbot.InlineKeyboardButton{
				{
					{
						Text: "OwO",
						TypeField: &gotdbot.InlineKeyboardButtonType{
							InlineKeyboardButtonTypeCallback: &gotdbot.InlineKeyboardButtonTypeCallback{
								Data: []byte("OwO"),
							},
						},
					},
					{
						Text: "UwU",
						TypeField: &gotdbot.InlineKeyboardButtonType{
							InlineKeyboardButtonTypeCallback: &gotdbot.InlineKeyboardButtonTypeCallback{
								Data: []byte("UwU"),
							},
						},
					},
				},
			},
		}

		text, err := ctx.Client.ParseText("This is a Inline keyboard", "Markdown")
		if err != nil {
			ctx.Client.Logger.Error("Failed to parse text", "err", err)
			return err
		}

		content := &gotdbot.InputMessageContent{InputMessageText: &gotdbot.InputMessageText{Text: text}}
		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkup{
				ReplyMarkupInlineKeyboard: kb,
			},
		}

		_, err = ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// /keyboard - Send message with reply keyboard
	dispatcher.AddHandler(handlers.NewCommand("keyboard", func(ctx *ext.Context) error {
		kb := &gotdbot.ReplyMarkupShowKeyboard{
			Rows: [][]*gotdbot.KeyboardButton{
				{
					{
						Text: "OwO",
						TypeField: &gotdbot.KeyboardButtonType{
							KeyboardButtonTypeText: &gotdbot.KeyboardButtonTypeText{},
						},
					},
					{
						Text: "UwU",
						TypeField: &gotdbot.KeyboardButtonType{
							KeyboardButtonTypeText: &gotdbot.KeyboardButtonTypeText{},
						},
					},
				},
			},
			ResizeKeyboard: true,
			OneTime:        true,
		}

		content := &gotdbot.InputMessageContent{
			InputMessageText: &gotdbot.InputMessageText{
				Text: &gotdbot.FormattedText{
					Text: "This is a keyboard",
				},
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkup{
				ReplyMarkupShowKeyboard: kb,
			},
		}

		_, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// /remove - Remove keyboard
	dispatcher.AddHandler(handlers.NewCommand("remove", func(ctx *ext.Context) error {
		content := &gotdbot.InputMessageContent{
			InputMessageText: &gotdbot.InputMessageText{
				Text: &gotdbot.FormattedText{
					Text: "Keyboards removed",
				},
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkup{
				ReplyMarkupRemoveKeyboard: &gotdbot.ReplyMarkupRemoveKeyboard{},
			},
		}

		_, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// /force - Force reply
	dispatcher.AddHandler(handlers.NewCommand("force", func(ctx *ext.Context) error {
		content := &gotdbot.InputMessageContent{
			InputMessageText: &gotdbot.InputMessageText{
				Text: &gotdbot.FormattedText{
					Text: "This is a force reply",
				},
			},
		}

		opts := &gotdbot.SendMessageOpts{
			ReplyMarkup: &gotdbot.ReplyMarkup{
				ReplyMarkupForceReply: &gotdbot.ReplyMarkupForceReply{},
			},
		}

		_, err := ctx.Client.SendMessage(ctx.EffectiveChatId, content, opts)
		return err
	}))

	// CallbackQuery Handler
	dispatcher.AddHandler(handlers.NewCallbackQuery(nil, func(ctx *ext.Context) error {
		update := ctx.RawUpdate.(*gotdbot.UpdateNewCallbackQuery)

		var data string
		if update.Payload != nil && update.Payload.CallbackQueryPayloadData != nil {
			data = string(update.Payload.CallbackQueryPayloadData.Data)
		}

		if data != "" {
			kb := &gotdbot.ReplyMarkupInlineKeyboard{
				Rows: [][]*gotdbot.InlineKeyboardButton{
					{
						{
							Text: "GitHub",
							TypeField: &gotdbot.InlineKeyboardButtonType{
								InlineKeyboardButtonTypeUrl: &gotdbot.InlineKeyboardButtonTypeUrl{
									Url: "https://github.com/AshokShau/gotdbot",
								},
							},
						},
					},
				},
			}

			inputContent := &gotdbot.InputMessageContent{
				InputMessageText: &gotdbot.InputMessageText{
					Text: &gotdbot.FormattedText{
						Text: fmt.Sprintf("You pressed %s", data),
					},
				},
			}

			_, err := ctx.Client.EditMessageText(update.ChatId, update.MessageId, inputContent, &gotdbot.EditMessageTextOpts{
				ReplyMarkup: &gotdbot.ReplyMarkup{
					ReplyMarkupInlineKeyboard: kb,
				},
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
