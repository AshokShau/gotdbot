package gotdbot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AshokShau/gotdbot/internal/qrcode"
)

// sendAuthError publishes the first authorization result to all Start callers.
func (c *Client) sendAuthError(err error) {
	c.authOnce.Do(func() { c.authErr = err; close(c.authDone) })
}

func (c *Client) authHandler(authState *UpdateAuthorizationState) {
	c.Logger.Debug("Authorization state update", "state", authState.AuthorizationState.GetType())

	switch state := authState.AuthorizationState.(type) {
	case *AuthorizationStateWaitTdlibParameters:
		if c.config.TDLibOptions != nil {
			c.config.TDLibOptions.forEachSet(func(k string, v any) {
				if opt := toOptionValue(v); opt != nil {
					err := c.SetOption(k, &SetOptionOpts{Value: opt})
					if err != nil {
						c.Logger.Error("Error setting option", "option", k, "error", err)
					}
				}
			})
		}

		err := c.SetTdlibParameters(
			c.apiHash,
			c.apiID,
			c.config.ApplicationVersion,
			[]byte(c.config.DatabaseEncryptionKey),
			c.config.DeviceModel,
			c.config.SystemLanguageCode,
			&SetTdlibParametersOpts{
				DatabaseDirectory:   c.config.DatabaseDirectory,
				FilesDirectory:      c.config.FilesDirectory,
				SystemVersion:       c.config.SystemVersion,
				UseChatInfoDatabase: *c.config.UseChatInfoDatabase,
				UseFileDatabase:     *c.config.UseFileDatabase,
				UseMessageDatabase:  *c.config.UseMessageDatabase,
				UseSecretChats:      c.config.UseSecretChats,
				UseTestDc:           c.config.UseTestDC,
			},
		)
		if err != nil {
			c.Logger.Error("Error setting tdlib parameters", "error", err)
			c.sendAuthError(err)
		}

	case *AuthorizationStateWaitPhoneNumber:
		if c.botToken != "" {
			err := c.CheckAuthenticationBotToken(c.botToken)
			if err != nil {
				c.Logger.Error("Error checking bot token", "error", err)
				c.sendAuthError(err)
			}
		} else {
			if c.config.QrMode {
				err := c.RequestQrCodeAuthentication(nil)
				if err != nil {
					c.Logger.Error("Error requesting QR code", "error", err)
					c.sendAuthError(err)
				}
			} else if c.phoneNumber != "" {
				err := c.SetAuthenticationPhoneNumber(c.phoneNumber, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.sendAuthError(err)
				}
			} else {
				fmt.Print("Enter phone number: ")
				reader := bufio.NewReader(os.Stdin)
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)
				err := c.SetAuthenticationPhoneNumber(phone, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.sendAuthError(err)
				}
			}
		}

	case *AuthorizationStateWaitOtherDeviceConfirmation:
		link := state.Link
		fmt.Printf("Scan the QR code below: or open the link in Telegram: %s\n", link)
		q, err := qrcode.NewQRCode(link)
		if err != nil {
			fmt.Printf("Failed to generate QR: %v\nLink: %s\n", err, link)
		} else {
			fmt.Println(q.ToSmallString(false))
		}

	case *AuthorizationStateWaitCode:
		codeInfo := state.CodeInfo
		codeType := strings.TrimPrefix(codeInfo.Type.GetType(), "authenticationCodeType")
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("Enter the code received via %s: ", codeType)
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)
			if code == "" {
				continue
			}
			err := c.CheckAuthenticationCode(code)
			if err != nil {
				fmt.Printf("Error checking code: %v\n", err)
				continue
			}
			break
		}

	case *AuthorizationStateWaitPassword:
		hint := state.PasswordHint
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("Enter your 2FA password (hint: %s): ", hint)
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password == "" {
				continue
			}
			err := c.CheckAuthenticationPassword(password)
			if err != nil {
				fmt.Printf("Error checking password: %v\n", err)
				continue
			}
			break
		}

	case *AuthorizationStateWaitRegistration:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter first name: ")
		firstName, _ := reader.ReadString('\n')
		firstName = strings.TrimSpace(firstName)
		fmt.Print("Enter last name: ")
		lastName, _ := reader.ReadString('\n')
		lastName = strings.TrimSpace(lastName)

		err := c.RegisterUser(firstName, &RegisterUserOpts{LastName: lastName})
		if err != nil {
			c.Logger.Error("Error registering user", "error", err)
			c.sendAuthError(err)
		}

	case *AuthorizationStateWaitPremiumPurchase:
		c.Logger.Warn("Account is limited and requires Telegram Premium.")
		c.sendAuthError(WaitPremiumPurchase)

	case *AuthorizationStateReady:
		me, err := c.GetMe()
		if err != nil {
			c.Logger.Error("Failed to get me", "error", err)
			c.sendAuthError(err)
			return
		}

		c.meMu.Lock()
		c.Me = me
		c.meMu.Unlock()
		c.isAuthorized.Store(true)

		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		c.Logger.Info("Logged in", "user_id", me.Id, "username", username)

		c.sendAuthError(nil)

	}
	return
}

func toOptionValue(v any) OptionValue {
	switch val := v.(type) {
	case bool:
		return &OptionValueBoolean{Value: val}
	case int:
		return &OptionValueInteger{Value: int64(val)}
	case int32:
		return &OptionValueInteger{Value: int64(val)}
	case int64:
		return &OptionValueInteger{Value: val}
	case string:
		return &OptionValueString{Value: val}
	case nil:
		return &OptionValueEmpty{}
	default:
		return nil
	}
}
