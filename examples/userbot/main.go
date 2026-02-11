package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/mdp/qrterminal/v3"
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"

	bot := gotdbot.NewClient(apiID, apiHash, "", &gotdbot.ClientConfig{
		LibraryPath:           "./libtdjson.so.1.8.61",
		IsUser:                true,
		UseFileDatabase:       gotdbot.Bool(true),
		FilesDirectory:        "UserBot",
		AuthorizationTimeout:  2 * time.Minute,
		DatabaseEncryptionKey: "my_secret_key_29", // Optional: Set a key to encrypt the local database
	})

	gotdbot.SetTdlibLogStreamEmpty()
	dispatcher := ext.NewDispatcher(bot)

	// Auth handler
	dispatcher.AddHandler(handlers.NewUpdateAuthorizationState(nil, func(ctx *ext.Context) error {
		authState := ctx.Update.UpdateAuthorizationState.AuthorizationState

		switch authState.Type() {
		case "authorizationStateWaitPhoneNumber":
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("Enter your phone number, or type 'qr' to log in using a QR code: ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "" {
					continue
				}

				if strings.ToLower(input) == "qr" {
					err := ctx.Client.RequestQrCodeAuthentication(nil)
					if err != nil {
						fmt.Printf("Error requesting QR code: %v\n", err)
						continue
					}
				} else {
					err := ctx.Client.SetAuthenticationPhoneNumber(input, nil)
					if err != nil {
						fmt.Printf("Error setting phone number: %v\n", err)
						continue
					}
				}
				break
			}

		case "authorizationStateWaitOtherDeviceConfirmation":
			link := authState.(*gotdbot.AuthorizationStateWaitOtherDeviceConfirmation).Link
			fmt.Println("\nScan this QR code using your Telegram mobile app.")
			fmt.Println("Go to Settings -> Devices -> Link Desktop Device.")
			fmt.Println("\nIf you prefer, you can also open this link on your mobile browser:")
			fmt.Printf("%s\n\n", link)

			qrterminal.GenerateHalfBlock(link, qrterminal.L, os.Stdout)
			fmt.Println()

		case "authorizationStateWaitCode":
			codeInfo := authState.(*gotdbot.AuthorizationStateWaitCode).CodeInfo
			codeType := codeInfo.TypeField.Type()
			codeType = strings.TrimPrefix(codeType, "authenticationCodeType")
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Printf("Enter the code received via %s: ", codeType)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "" {
					continue
				}

				err := ctx.Client.CheckAuthenticationCode(input)
				if err != nil {
					fmt.Printf("Error checking code: %v\n", err)
					continue
				}
				break
			}

		case "authorizationStateWaitPassword":
			hint := authState.(*gotdbot.AuthorizationStateWaitPassword).PasswordHint
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Printf("Enter your 2FA password (hint: %s): ", hint)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "" {
					continue
				}

				err := ctx.Client.CheckAuthenticationPassword(input)
				if err != nil {
					fmt.Printf("Error checking password: %v\n", err)
					continue
				}
				break
			}

		case "authorizationStateWaitPremiumPurchase":
			fmt.Println("Your account requires Telegram Premium to log in. Please purchase Telegram Premium and try again.")
			os.Exit(1)

		case "authorizationStateReady":
			fmt.Println("Authorization successful!")
		}
		return nil
	}))

	dispatcher.AddHandler(handlers.NewCommand("hi", func(ctx *ext.Context) error {
		_, err := ctx.EffectiveMessage.ReplyText(ctx.Client, "Hi, this is from gotdbot!", nil)
		return err
	}))

	dispatcher.Start()
	log.Println("Starting userbot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	me := bot.Me()
	if me != nil {
		fmt.Printf("Current user: %s (ID: %d)\n", me.FirstName, me.Id)
	}

	bot.Idle()
}
