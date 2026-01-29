package client

import "github.com/AshokShau/gotdbot/types"

// GetAuthorizationState Returns the current authorization state. This is an offline method. For informational purposes only. Use updateAuthorizationState instead to maintain the current authorization state. Can be called before initialization
func (c *Client) GetAuthorizationState() (*types.AuthorizationState, error) {
	req := &types.GetAuthorizationState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AuthorizationState), nil
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
func (c *Client) SetTdlibParameters(useTestDc bool, databaseDirectory string, filesDirectory string, databaseEncryptionKey string, useFileDatabase bool, useChatInfoDatabase bool, useMessageDatabase bool, useSecretChats bool, apiId int32, apiHash string, systemLanguageCode string, deviceModel string, systemVersion string, applicationVersion string) (*types.Ok, error) {
	req := &types.SetTdlibParameters{
		UseTestDc:             useTestDc,
		DatabaseDirectory:     databaseDirectory,
		FilesDirectory:        filesDirectory,
		DatabaseEncryptionKey: databaseEncryptionKey,
		UseFileDatabase:       useFileDatabase,
		UseChatInfoDatabase:   useChatInfoDatabase,
		UseMessageDatabase:    useMessageDatabase,
		UseSecretChats:        useSecretChats,
		ApiId:                 apiId,
		ApiHash:               apiHash,
		SystemLanguageCode:    systemLanguageCode,
		DeviceModel:           deviceModel,
		SystemVersion:         systemVersion,
		ApplicationVersion:    applicationVersion,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
func (c *Client) SetAuthenticationPhoneNumber(phoneNumber string, opts *types.SetAuthenticationPhoneNumberOpts) (*types.Ok, error) {
	req := &types.SetAuthenticationPhoneNumber{
		PhoneNumber: phoneNumber,
	}
	if opts != nil {
		req.Settings = opts.Settings
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationPremiumPurchase Checks whether an in-store purchase of Telegram Premium is possible before authorization. Works only when the current authorization state is authorizationStateWaitPremiumPurchase
func (c *Client) CheckAuthenticationPremiumPurchase(currency string, amount int64) (*types.Ok, error) {
	req := &types.CheckAuthenticationPremiumPurchase{
		Currency: currency,
		Amount:   amount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAuthenticationPremiumPurchaseTransaction Informs server about an in-store purchase of Telegram Premium before authorization. Works only when the current authorization state is authorizationStateWaitPremiumPurchase
func (c *Client) SetAuthenticationPremiumPurchaseTransaction(transaction *types.StoreTransaction, isRestore bool, currency string, amount int64) (*types.Ok, error) {
	req := &types.SetAuthenticationPremiumPurchaseTransaction{
		Transaction: transaction,
		IsRestore:   isRestore,
		Currency:    currency,
		Amount:      amount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAuthenticationEmailAddress Sets the email address of the user and sends an authentication code to the email address. Works only when the current authorization state is authorizationStateWaitEmailAddress @email_address The email address of the user
func (c *Client) SetAuthenticationEmailAddress(emailAddress string) (*types.Ok, error) {
	req := &types.SetAuthenticationEmailAddress{
		EmailAddress: emailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResendAuthenticationCode Resends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode, the next_code_type of the result is not null
func (c *Client) ResendAuthenticationCode(opts *types.ResendAuthenticationCodeOpts) (*types.Ok, error) {
	req := &types.ResendAuthenticationCode{}
	if opts != nil {
		req.Reason = opts.Reason
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationEmailCode Checks the authentication of an email address. Works only when the current authorization state is authorizationStateWaitEmailCode @code Email address authentication to check
func (c *Client) CheckAuthenticationEmailCode(code *types.EmailAddressAuthentication) (*types.Ok, error) {
	req := &types.CheckAuthenticationEmailCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode @code Authentication code to check
func (c *Client) CheckAuthenticationCode(code string) (*types.Ok, error) {
	req := &types.CheckAuthenticationCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RequestQrCodeAuthentication Requests QR code authentication by scanning a QR code on another logged in device. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
func (c *Client) RequestQrCodeAuthentication(otherUserIds []int64) (*types.Ok, error) {
	req := &types.RequestQrCodeAuthentication{
		OtherUserIds: otherUserIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetAuthenticationPasskeyParameters Returns parameters for authentication using a passkey as JSON-serialized string
func (c *Client) GetAuthenticationPasskeyParameters() (*types.Text, error) {
	req := &types.GetAuthenticationPasskeyParameters{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// CheckAuthenticationPasskey Checks a passkey to log in to the corresponding account. Call getAuthenticationPasskeyParameters to get parameters for the passkey. Works only when the current authorization state is
func (c *Client) CheckAuthenticationPasskey(credentialId string, clientData string, authenticatorData string, signature string, userHandle string) (*types.Ok, error) {
	req := &types.CheckAuthenticationPasskey{
		CredentialId:      credentialId,
		ClientData:        clientData,
		AuthenticatorData: authenticatorData,
		Signature:         signature,
		UserHandle:        userHandle,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RegisterUser Finishes user registration. Works only when the current authorization state is authorizationStateWaitRegistration
func (c *Client) RegisterUser(firstName string, lastName string, disableNotification bool) (*types.Ok, error) {
	req := &types.RegisterUser{
		FirstName:           firstName,
		LastName:            lastName,
		DisableNotification: disableNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResetAuthenticationEmailAddress Resets the login email address. May return an error with a message "TASK_ALREADY_EXISTS" if reset is still pending.
func (c *Client) ResetAuthenticationEmailAddress() (*types.Ok, error) {
	req := &types.ResetAuthenticationEmailAddress{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationPassword Checks the 2-step verification password for correctness. Works only when the current authorization state is authorizationStateWaitPassword @password The 2-step verification password to check
func (c *Client) CheckAuthenticationPassword(password string) (*types.Ok, error) {
	req := &types.CheckAuthenticationPassword{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RequestAuthenticationPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (c *Client) RequestAuthenticationPasswordRecovery() (*types.Ok, error) {
	req := &types.RequestAuthenticationPasswordRecovery{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid. Works only when the current authorization state is authorizationStateWaitPassword @recovery_code Recovery code to check
func (c *Client) CheckAuthenticationPasswordRecoveryCode(recoveryCode string) (*types.Ok, error) {
	req := &types.CheckAuthenticationPasswordRecoveryCode{
		RecoveryCode: recoveryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RecoverAuthenticationPassword Recovers the 2-step verification password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (c *Client) RecoverAuthenticationPassword(recoveryCode string, newPassword string, newHint string) (*types.Ok, error) {
	req := &types.RecoverAuthenticationPassword{
		RecoveryCode: recoveryCode,
		NewPassword:  newPassword,
		NewHint:      newHint,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendAuthenticationFirebaseSms Sends Firebase Authentication SMS to the phone number of the user. Works only when the current authorization state is authorizationStateWaitCode and the server returned code of the type authenticationCodeTypeFirebaseAndroid or authenticationCodeTypeFirebaseIos
func (c *Client) SendAuthenticationFirebaseSms(token string) (*types.Ok, error) {
	req := &types.SendAuthenticationFirebaseSms{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportAuthenticationCodeMissing Reports that authentication code wasn't delivered via SMS; for official mobile applications only. Works only when the current authorization state is authorizationStateWaitCode @mobile_network_code Current mobile network code
func (c *Client) ReportAuthenticationCodeMissing(mobileNetworkCode string) (*types.Ok, error) {
	req := &types.ReportAuthenticationCodeMissing{
		MobileNetworkCode: mobileNetworkCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in @token The bot token
func (c *Client) CheckAuthenticationBotToken(token string) (*types.Ok, error) {
	req := &types.CheckAuthenticationBotToken{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (c *Client) LogOut() (*types.Ok, error) {
	req := &types.LogOut{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (c *Client) Close() (*types.Ok, error) {
	req := &types.Close{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed.
func (c *Client) Destroy() (*types.Ok, error) {
	req := &types.Destroy{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success @link A link from a QR code. The link must be scanned by the in-app camera
func (c *Client) ConfirmQrCodeAuthentication(link string) (*types.Session, error) {
	req := &types.ConfirmQrCodeAuthentication{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Session), nil
}

// GetCurrentState Returns all updates needed to restore current TDLib state, i.e. all actual updateAuthorizationState/updateUser/updateNewChat and others. This is especially useful if TDLib is run in a separate process. Can be called before initialization
func (c *Client) GetCurrentState() (*types.Updates, error) {
	req := &types.GetCurrentState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Updates), nil
}

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain @new_encryption_key New encryption key
func (c *Client) SetDatabaseEncryptionKey(newEncryptionKey string) (*types.Ok, error) {
	req := &types.SetDatabaseEncryptionKey{
		NewEncryptionKey: newEncryptionKey,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPasswordState Returns the current state of 2-step verification
func (c *Client) GetPasswordState() (*types.PasswordState, error) {
	req := &types.GetPasswordState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// SetPassword Changes the 2-step verification password for the current user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed
func (c *Client) SetPassword(oldPassword string, newPassword string, newHint string, setRecoveryEmailAddress bool, newRecoveryEmailAddress string) (*types.PasswordState, error) {
	req := &types.SetPassword{
		OldPassword:             oldPassword,
		NewPassword:             newPassword,
		NewHint:                 newHint,
		SetRecoveryEmailAddress: setRecoveryEmailAddress,
		NewRecoveryEmailAddress: newRecoveryEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// IsLoginEmailAddressRequired Checks whether the current user is required to set login email address
func (c *Client) IsLoginEmailAddressRequired() (*types.Ok, error) {
	req := &types.IsLoginEmailAddressRequired{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetLoginEmailAddress Changes the login email address of the user. The email address can be changed only if the current user already has login email and passwordState.login_email_address_pattern is non-empty,
func (c *Client) SetLoginEmailAddress(newLoginEmailAddress string) (*types.EmailAddressAuthenticationCodeInfo, error) {
	req := &types.SetLoginEmailAddress{
		NewLoginEmailAddress: newLoginEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmailAddressAuthenticationCodeInfo), nil
}

// ResendLoginEmailAddressCode Resends the login email address verification code
func (c *Client) ResendLoginEmailAddressCode() (*types.EmailAddressAuthenticationCodeInfo, error) {
	req := &types.ResendLoginEmailAddressCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmailAddressAuthenticationCodeInfo), nil
}

// CheckLoginEmailAddressCode Checks the login email address authentication @code Email address authentication to check
func (c *Client) CheckLoginEmailAddressCode(code *types.EmailAddressAuthentication) (*types.Ok, error) {
	req := &types.CheckLoginEmailAddressCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetRecoveryEmailAddress Returns a 2-step verification recovery email address that was previously set up. This method can be used to verify a password provided by the user @password The 2-step verification password for the current user
func (c *Client) GetRecoveryEmailAddress(password string) (*types.RecoveryEmailAddress, error) {
	req := &types.GetRecoveryEmailAddress{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RecoveryEmailAddress), nil
}

// SetRecoveryEmailAddress Changes the 2-step verification recovery email address of the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed.
func (c *Client) SetRecoveryEmailAddress(password string, newRecoveryEmailAddress string) (*types.PasswordState, error) {
	req := &types.SetRecoveryEmailAddress{
		Password:                password,
		NewRecoveryEmailAddress: newRecoveryEmailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// CheckRecoveryEmailAddressCode Checks the 2-step verification recovery email address verification code @code Verification code to check
func (c *Client) CheckRecoveryEmailAddressCode(code string) (*types.PasswordState, error) {
	req := &types.CheckRecoveryEmailAddressCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// ResendRecoveryEmailAddressCode Resends the 2-step verification recovery email address verification code
func (c *Client) ResendRecoveryEmailAddressCode() (*types.PasswordState, error) {
	req := &types.ResendRecoveryEmailAddressCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// CancelRecoveryEmailAddressVerification Cancels verification of the 2-step verification recovery email address
func (c *Client) CancelRecoveryEmailAddressVerification() (*types.PasswordState, error) {
	req := &types.CancelRecoveryEmailAddressVerification{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// RequestPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up
func (c *Client) RequestPasswordRecovery() (*types.EmailAddressAuthenticationCodeInfo, error) {
	req := &types.RequestPasswordRecovery{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmailAddressAuthenticationCodeInfo), nil
}

// CheckPasswordRecoveryCode Checks whether a 2-step verification password recovery code sent to an email address is valid @recovery_code Recovery code to check
func (c *Client) CheckPasswordRecoveryCode(recoveryCode string) (*types.Ok, error) {
	req := &types.CheckPasswordRecoveryCode{
		RecoveryCode: recoveryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RecoverPassword Recovers the 2-step verification password using a recovery code sent to an email address that was previously set up
func (c *Client) RecoverPassword(recoveryCode string, newPassword string, newHint string) (*types.PasswordState, error) {
	req := &types.RecoverPassword{
		RecoveryCode: recoveryCode,
		NewPassword:  newPassword,
		NewHint:      newHint,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PasswordState), nil
}

// ResetPassword Removes 2-step verification password without previous password and access to recovery email address. The password can't be reset immediately and the request needs to be repeated after the specified time
func (c *Client) ResetPassword() (*types.ResetPasswordResult, error) {
	req := &types.ResetPassword{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ResetPasswordResult), nil
}

// CancelPasswordReset Cancels reset of 2-step verification password. The method can be called if passwordState.pending_reset_date > 0
func (c *Client) CancelPasswordReset() (*types.Ok, error) {
	req := &types.CancelPasswordReset{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CreateTemporaryPassword Creates a new temporary password for processing payments @password The 2-step verification password of the current user @valid_for Time during which the temporary password will be valid, in seconds; must be between 60 and 86400
func (c *Client) CreateTemporaryPassword(password string, validFor int32) (*types.TemporaryPasswordState, error) {
	req := &types.CreateTemporaryPassword{
		Password: password,
		ValidFor: validFor,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TemporaryPasswordState), nil
}

// GetTemporaryPasswordState Returns information about the current temporary password
func (c *Client) GetTemporaryPasswordState() (*types.TemporaryPasswordState, error) {
	req := &types.GetTemporaryPasswordState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TemporaryPasswordState), nil
}

// GetMe Returns the current user
func (c *Client) GetMe() (*types.User, error) {
	req := &types.GetMe{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// GetUser Returns information about a user by their identifier. This is an offline method if the current user is not a bot @user_id User identifier
func (c *Client) GetUser(userId int64) (*types.User, error) {
	req := &types.GetUser{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// GetUserFullInfo Returns full information about a user by their identifier @user_id User identifier
func (c *Client) GetUserFullInfo(userId int64) (*types.UserFullInfo, error) {
	req := &types.GetUserFullInfo{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UserFullInfo), nil
}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline method if the current user is not a bot @basic_group_id Basic group identifier
func (c *Client) GetBasicGroup(basicGroupId int64) (*types.BasicGroup, error) {
	req := &types.GetBasicGroup{
		BasicGroupId: basicGroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BasicGroup), nil
}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier @basic_group_id Basic group identifier
func (c *Client) GetBasicGroupFullInfo(basicGroupId int64) (*types.BasicGroupFullInfo, error) {
	req := &types.GetBasicGroupFullInfo{
		BasicGroupId: basicGroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BasicGroupFullInfo), nil
}

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline method if the current user is not a bot @supergroup_id Supergroup or channel identifier
func (c *Client) GetSupergroup(supergroupId int64) (*types.Supergroup, error) {
	req := &types.GetSupergroup{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Supergroup), nil
}

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute @supergroup_id Supergroup or channel identifier
func (c *Client) GetSupergroupFullInfo(supergroupId int64) (*types.SupergroupFullInfo, error) {
	req := &types.GetSupergroupFullInfo{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SupergroupFullInfo), nil
}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline method @secret_chat_id Secret chat identifier
func (c *Client) GetSecretChat(secretChatId int32) (*types.SecretChat, error) {
	req := &types.GetSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SecretChat), nil
}

// GetChat Returns information about a chat by its identifier. This is an offline method if the current user is not a bot @chat_id Chat identifier
func (c *Client) GetChat(chatId int64) (*types.Chat, error) {
	req := &types.GetChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// GetMessage Returns information about a message. Returns a 404 error if the message doesn't exist
func (c *Client) GetMessage(chatId int64, messageId int64) (*types.Message, error) {
	req := &types.GetMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetMessageLocally Returns information about a message, if it is available without sending network request. Returns a 404 error if message isn't available locally. This is an offline method
func (c *Client) GetMessageLocally(chatId int64, messageId int64) (*types.Message, error) {
	req := &types.GetMessageLocally{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetRepliedMessage Returns information about a non-bundled message that is replied by a given message. Also, returns the pinned message for messagePinMessage,
func (c *Client) GetRepliedMessage(chatId int64, messageId int64) (*types.Message, error) {
	req := &types.GetRepliedMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat. Returns a 404 error if the message doesn't exist @chat_id Identifier of the chat the message belongs to
func (c *Client) GetChatPinnedMessage(chatId int64) (*types.Message, error) {
	req := &types.GetChatPinnedMessage{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only @chat_id Identifier of the chat the message belongs to @message_id Message identifier @callback_query_id Identifier of the callback query
func (c *Client) GetCallbackQueryMessage(chatId int64, messageId int64, callbackQueryId string) (*types.Message, error) {
	req := &types.GetCallbackQueryMessage{
		ChatId:          chatId,
		MessageId:       messageId,
		CallbackQueryId: callbackQueryId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result @chat_id Identifier of the chat the messages belong to @message_ids Identifiers of the messages to get
func (c *Client) GetMessages(chatId int64, messageIds []int64) (*types.Messages, error) {
	req := &types.GetMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetMessageProperties Returns properties of a message. This is an offline method @chat_id Chat identifier @message_id Identifier of the message
func (c *Client) GetMessageProperties(chatId int64, messageId int64) (*types.MessageProperties, error) {
	req := &types.GetMessageProperties{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageProperties), nil
}

// GetMessageThread Returns information about a message thread. Can be used only if messageProperties.can_get_message_thread == true @chat_id Chat identifier @message_id Identifier of the message
func (c *Client) GetMessageThread(chatId int64, messageId int64) (*types.MessageThreadInfo, error) {
	req := &types.GetMessageThread{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageThreadInfo), nil
}

// GetMessageReadDate Returns read date of a recent outgoing message in a private chat. The method can be called if messageProperties.can_get_read_date == true
func (c *Client) GetMessageReadDate(chatId int64, messageId int64) (*types.MessageReadDate, error) {
	req := &types.GetMessageReadDate{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageReadDate), nil
}

// GetMessageViewers Returns viewers of a recent outgoing message in a basic group or a supergroup chat. For video notes and voice notes only users, opened content of the message, are returned. The method can be called if messageProperties.can_get_viewers == true
func (c *Client) GetMessageViewers(chatId int64, messageId int64) (*types.MessageViewers, error) {
	req := &types.GetMessageViewers{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageViewers), nil
}

// GetMessageAuthor Returns information about actual author of a message sent on behalf of a channel. The method can be called if messageProperties.can_get_author == true
func (c *Client) GetMessageAuthor(chatId int64, messageId int64) (*types.User, error) {
	req := &types.GetMessageAuthor{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// GetFile Returns information about a file. This is an offline method @file_id Identifier of the file to get
func (c *Client) GetFile(fileId int32) (*types.File, error) {
	req := &types.GetFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// GetRemoteFile Returns information about a file by its remote identifier. This is an offline method. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user.
func (c *Client) GetRemoteFile(remoteFileId string, opts *types.GetRemoteFileOpts) (*types.File, error) {
	req := &types.GetRemoteFile{
		RemoteFileId: remoteFileId,
	}
	if opts != nil {
		req.FileType = opts.FileType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// LoadChats Loads more chats from a chat list. The loaded chats and their positions in the chat list will be sent through updates. Chats are sorted by the pair (chat.position.order, chat.id) in descending order. Returns a 404 error if all chats have been loaded
func (c *Client) LoadChats(limit int32, opts *types.LoadChatsOpts) (*types.Ok, error) {
	req := &types.LoadChats{
		Limit: limit,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChats Returns an ordered list of chats from the beginning of a chat list. For informational purposes only. Use loadChats and updates processing instead to maintain chat lists in a consistent state
func (c *Client) GetChats(limit int32, opts *types.GetChatsOpts) (*types.Chats, error) {
	req := &types.GetChats{
		Limit: limit,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// SearchPublicChat Searches a public chat by its username. Currently, only private chats, supergroups and channels can be public. Returns the chat if found; otherwise, an error is returned @username Username to be resolved
func (c *Client) SearchPublicChat(username string) (*types.Chat, error) {
	req := &types.SearchPublicChat{
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently, only private chats, supergroups and channels can be public. Returns a meaningful number of results.
func (c *Client) SearchPublicChats(query string) (*types.Chats, error) {
	req := &types.SearchPublicChats{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// SearchChats Searches for the specified query in the title and username of already known chats. This is an offline method. Returns chats in the order seen in the main chat list
func (c *Client) SearchChats(query string, limit int32) (*types.Chats, error) {
	req := &types.SearchChats{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list @query Query to search for @limit The maximum number of chats to be returned
func (c *Client) SearchChatsOnServer(query string, limit int32) (*types.Chats, error) {
	req := &types.SearchChatsOnServer{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetRecommendedChats Returns a list of channel chats recommended to the current user
func (c *Client) GetRecommendedChats() (*types.Chats, error) {
	req := &types.GetRecommendedChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetChatSimilarChats Returns a list of chats similar to the given chat @chat_id Identifier of the target chat; must be an identifier of a channel chat
func (c *Client) GetChatSimilarChats(chatId int64) (*types.Chats, error) {
	req := &types.GetChatSimilarChats{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetChatSimilarChatCount Returns approximate number of chats similar to the given chat
func (c *Client) GetChatSimilarChatCount(chatId int64, returnLocal bool) (*types.Count, error) {
	req := &types.GetChatSimilarChatCount{
		ChatId:      chatId,
		ReturnLocal: returnLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// OpenChatSimilarChat Informs TDLib that a chat was opened from the list of similar chats. The method is independent of openChat and closeChat methods
func (c *Client) OpenChatSimilarChat(chatId int64, openedChatId int64) (*types.Ok, error) {
	req := &types.OpenChatSimilarChat{
		ChatId:       chatId,
		OpenedChatId: openedChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBotSimilarBots Returns a list of bots similar to the given bot @bot_user_id User identifier of the target bot
func (c *Client) GetBotSimilarBots(botUserId int64) (*types.Users, error) {
	req := &types.GetBotSimilarBots{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// GetBotSimilarBotCount Returns approximate number of bots similar to the given bot
func (c *Client) GetBotSimilarBotCount(botUserId int64, returnLocal bool) (*types.Count, error) {
	req := &types.GetBotSimilarBotCount{
		BotUserId:   botUserId,
		ReturnLocal: returnLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// OpenBotSimilarBot Informs TDLib that a bot was opened from the list of similar bots
func (c *Client) OpenBotSimilarBot(botUserId int64, openedBotUserId int64) (*types.Ok, error) {
	req := &types.OpenBotSimilarBot{
		BotUserId:       botUserId,
		OpenedBotUserId: openedBotUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetTopChats Returns a list of frequently used chats @category Category of chats to be returned @limit The maximum number of chats to be returned; up to 30
func (c *Client) GetTopChats(category *types.TopChatCategory, limit int32) (*types.Chats, error) {
	req := &types.GetTopChats{
		Category: category,
		Limit:    limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled @category Category of frequently used chats @chat_id Chat identifier
func (c *Client) RemoveTopChat(category *types.TopChatCategory, chatId int64) (*types.Ok, error) {
	req := &types.RemoveTopChat{
		Category: category,
		ChatId:   chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchRecentlyFoundChats Searches for the specified query in the title and username of up to 50 recently found chats. This is an offline method
func (c *Client) SearchRecentlyFoundChats(query string, limit int32) (*types.Chats, error) {
	req := &types.SearchRecentlyFoundChats{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first @chat_id Identifier of the chat to add
func (c *Client) AddRecentlyFoundChat(chatId int64) (*types.Ok, error) {
	req := &types.AddRecentlyFoundChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats @chat_id Identifier of the chat to be removed
func (c *Client) RemoveRecentlyFoundChat(chatId int64) (*types.Ok, error) {
	req := &types.RemoveRecentlyFoundChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (c *Client) ClearRecentlyFoundChats() (*types.Ok, error) {
	req := &types.ClearRecentlyFoundChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetRecentlyOpenedChats Returns recently opened chats. This is an offline method. Returns chats in the order of last opening @limit The maximum number of chats to be returned
func (c *Client) GetRecentlyOpenedChats(limit int32) (*types.Chats, error) {
	req := &types.GetRecentlyOpenedChats{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// CheckChatUsername Checks whether a username can be set for a chat @chat_id Chat identifier; must be identifier of a supergroup chat, or a channel chat, or a private chat with self, or 0 if the chat is being created @username Username to be checked
func (c *Client) CheckChatUsername(chatId int64, username string) (*types.CheckChatUsernameResult, error) {
	req := &types.CheckChatUsername{
		ChatId:   chatId,
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CheckChatUsernameResult), nil
}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user @type Type of the public chats to return
func (c *Client) GetCreatedPublicChats(typeField *types.PublicChatType) (*types.Chats, error) {
	req := &types.GetCreatedPublicChats{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// CheckCreatedPublicChatsLimit Checks whether the maximum number of owned public chats has been reached. Returns corresponding error if the limit was reached. The limit can be increased with Telegram Premium @type Type of the public chats, for which to check the limit
func (c *Client) CheckCreatedPublicChatsLimit(typeField *types.PublicChatType) (*types.Ok, error) {
	req := &types.CheckCreatedPublicChatsLimit{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group.
func (c *Client) GetSuitableDiscussionChats() (*types.Chats, error) {
	req := &types.GetSuitableDiscussionChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives the error "CHANNELS_TOO_MUCH". Also, the limit can be increased with Telegram Premium
func (c *Client) GetInactiveSupergroupChats() (*types.Chats, error) {
	req := &types.GetInactiveSupergroupChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetSuitablePersonalChats Returns a list of channel chats, which can be used as a personal chat
func (c *Client) GetSuitablePersonalChats() (*types.Chats, error) {
	req := &types.GetSuitablePersonalChats{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// LoadDirectMessagesChatTopics Loads more topics in a channel direct messages chat administered by the current user. The loaded topics will be sent through updateDirectMessagesChatTopic.
func (c *Client) LoadDirectMessagesChatTopics(chatId int64, limit int32) (*types.Ok, error) {
	req := &types.LoadDirectMessagesChatTopics{
		ChatId: chatId,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetDirectMessagesChatTopic Returns information about the topic in a channel direct messages chat administered by the current user
func (c *Client) GetDirectMessagesChatTopic(chatId int64, topicId int64) (*types.DirectMessagesChatTopic, error) {
	req := &types.GetDirectMessagesChatTopic{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.DirectMessagesChatTopic), nil
}

// GetDirectMessagesChatTopicHistory Returns messages in the topic in a channel direct messages chat administered by the current user. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id)
func (c *Client) GetDirectMessagesChatTopicHistory(chatId int64, topicId int64, fromMessageId int64, offset int32, limit int32) (*types.Messages, error) {
	req := &types.GetDirectMessagesChatTopicHistory{
		ChatId:        chatId,
		TopicId:       topicId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetDirectMessagesChatTopicMessageByDate Returns the last message sent in the topic in a channel direct messages chat administered by the current user no later than the specified date
func (c *Client) GetDirectMessagesChatTopicMessageByDate(chatId int64, topicId int64, date int32) (*types.Message, error) {
	req := &types.GetDirectMessagesChatTopicMessageByDate{
		ChatId:  chatId,
		TopicId: topicId,
		Date:    date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// DeleteDirectMessagesChatTopicHistory Deletes all messages in the topic in a channel direct messages chat administered by the current user
func (c *Client) DeleteDirectMessagesChatTopicHistory(chatId int64, topicId int64) (*types.Ok, error) {
	req := &types.DeleteDirectMessagesChatTopicHistory{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteDirectMessagesChatTopicMessagesByDate Deletes all messages between the specified dates in the topic in a channel direct messages chat administered by the current user. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteDirectMessagesChatTopicMessagesByDate(chatId int64, topicId int64, minDate int32, maxDate int32) (*types.Ok, error) {
	req := &types.DeleteDirectMessagesChatTopicMessagesByDate{
		ChatId:  chatId,
		TopicId: topicId,
		MinDate: minDate,
		MaxDate: maxDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetDirectMessagesChatTopicIsMarkedAsUnread Changes the marked as unread state of the topic in a channel direct messages chat administered by the current user
func (c *Client) SetDirectMessagesChatTopicIsMarkedAsUnread(chatId int64, topicId int64, isMarkedAsUnread bool) (*types.Ok, error) {
	req := &types.SetDirectMessagesChatTopicIsMarkedAsUnread{
		ChatId:           chatId,
		TopicId:          topicId,
		IsMarkedAsUnread: isMarkedAsUnread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// UnpinAllDirectMessagesChatTopicMessages Removes all pinned messages from the topic in a channel direct messages chat administered by the current user
func (c *Client) UnpinAllDirectMessagesChatTopicMessages(chatId int64, topicId int64) (*types.Ok, error) {
	req := &types.UnpinAllDirectMessagesChatTopicMessages{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadAllDirectMessagesChatTopicReactions Removes all unread reactions in the topic in a channel direct messages chat administered by the current user
func (c *Client) ReadAllDirectMessagesChatTopicReactions(chatId int64, topicId int64) (*types.Ok, error) {
	req := &types.ReadAllDirectMessagesChatTopicReactions{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetDirectMessagesChatTopicRevenue Returns the total number of Telegram Stars received by the channel chat for direct messages from the given topic
func (c *Client) GetDirectMessagesChatTopicRevenue(chatId int64, topicId int64) (*types.StarCount, error) {
	req := &types.GetDirectMessagesChatTopicRevenue{
		ChatId:  chatId,
		TopicId: topicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarCount), nil
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages Allows to send unpaid messages to the given topic of the channel direct messages chat administered by the current user
func (c *Client) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(chatId int64, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*types.Ok, error) {
	req := &types.ToggleDirectMessagesChatTopicCanSendUnpaidMessages{
		ChatId:                chatId,
		TopicId:               topicId,
		CanSendUnpaidMessages: canSendUnpaidMessages,
		RefundPayments:        refundPayments,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LoadSavedMessagesTopics Loads more Saved Messages topics. The loaded topics will be sent through updateSavedMessagesTopic. Topics are sorted by their topic.order in descending order. Returns a 404 error if all topics have been loaded
func (c *Client) LoadSavedMessagesTopics(limit int32) (*types.Ok, error) {
	req := &types.LoadSavedMessagesTopics{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSavedMessagesTopicHistory Returns messages in a Saved Messages topic. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id)
func (c *Client) GetSavedMessagesTopicHistory(savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32) (*types.Messages, error) {
	req := &types.GetSavedMessagesTopicHistory{
		SavedMessagesTopicId: savedMessagesTopicId,
		FromMessageId:        fromMessageId,
		Offset:               offset,
		Limit:                limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetSavedMessagesTopicMessageByDate Returns the last message sent in a Saved Messages topic no later than the specified date
func (c *Client) GetSavedMessagesTopicMessageByDate(savedMessagesTopicId int64, date int32) (*types.Message, error) {
	req := &types.GetSavedMessagesTopicMessageByDate{
		SavedMessagesTopicId: savedMessagesTopicId,
		Date:                 date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// DeleteSavedMessagesTopicHistory Deletes all messages in a Saved Messages topic @saved_messages_topic_id Identifier of Saved Messages topic which messages will be deleted
func (c *Client) DeleteSavedMessagesTopicHistory(savedMessagesTopicId int64) (*types.Ok, error) {
	req := &types.DeleteSavedMessagesTopicHistory{
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteSavedMessagesTopicMessagesByDate Deletes all messages between the specified dates in a Saved Messages topic. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteSavedMessagesTopicMessagesByDate(savedMessagesTopicId int64, minDate int32, maxDate int32) (*types.Ok, error) {
	req := &types.DeleteSavedMessagesTopicMessagesByDate{
		SavedMessagesTopicId: savedMessagesTopicId,
		MinDate:              minDate,
		MaxDate:              maxDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSavedMessagesTopicIsPinned Changes the pinned state of a Saved Messages topic. There can be up to getOption("pinned_saved_messages_topic_count_max") pinned topics. The limit can be increased with Telegram Premium
func (c *Client) ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId int64, isPinned bool) (*types.Ok, error) {
	req := &types.ToggleSavedMessagesTopicIsPinned{
		SavedMessagesTopicId: savedMessagesTopicId,
		IsPinned:             isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPinnedSavedMessagesTopics Changes the order of pinned Saved Messages topics @saved_messages_topic_ids Identifiers of the new pinned Saved Messages topics
func (c *Client) SetPinnedSavedMessagesTopics(savedMessagesTopicIds []int64) (*types.Ok, error) {
	req := &types.SetPinnedSavedMessagesTopics{
		SavedMessagesTopicIds: savedMessagesTopicIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
func (c *Client) GetGroupsInCommon(userId int64, offsetChatId int64, limit int32) (*types.Chats, error) {
	req := &types.GetGroupsInCommon{
		UserId:       userId,
		OffsetChatId: offsetChatId,
		Limit:        limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetChatHistory Returns messages in a chat. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id).
func (c *Client) GetChatHistory(chatId int64, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*types.Messages, error) {
	req := &types.GetChatHistory{
		ChatId:        chatId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
		OnlyLocal:     onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if messageProperties.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup.
func (c *Client) GetMessageThreadHistory(chatId int64, messageId int64, fromMessageId int64, offset int32, limit int32) (*types.Messages, error) {
	req := &types.GetMessageThreadHistory{
		ChatId:        chatId,
		MessageId:     messageId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// DeleteChatHistory Deletes all messages in the chat. Use chat.can_be_deleted_only_for_self and chat.can_be_deleted_for_all_users fields to find whether and how the method can be applied to the chat
func (c *Client) DeleteChatHistory(chatId int64, removeFromChatList bool, revoke bool) (*types.Ok, error) {
	req := &types.DeleteChatHistory{
		ChatId:             chatId,
		RemoveFromChatList: removeFromChatList,
		Revoke:             revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteChat Deletes a chat along with all messages in the corresponding chat for all chat members. For group chats this will release the usernames and remove all members.
func (c *Client) DeleteChat(chatId int64) (*types.Ok, error) {
	req := &types.DeleteChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query
func (c *Client) SearchChatMessages(chatId int64, query string, fromMessageId int64, offset int32, limit int32, opts *types.SearchChatMessagesOpts) (*types.FoundChatMessages, error) {
	req := &types.SearchChatMessages{
		ChatId:        chatId,
		Query:         query,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.SenderId = opts.SenderId
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundChatMessages), nil
}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)).
func (c *Client) SearchMessages(query string, offset string, limit int32, minDate int32, maxDate int32, opts *types.SearchMessagesOpts) (*types.FoundMessages, error) {
	req := &types.SearchMessages{
		Query:   query,
		Offset:  offset,
		Limit:   limit,
		MinDate: minDate,
		MaxDate: maxDate,
	}
	if opts != nil {
		req.ChatList = opts.ChatList
		req.Filter = opts.Filter
		req.ChatTypeFilter = opts.ChatTypeFilter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundMessages), nil
}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance, the number of returned messages is chosen by TDLib
func (c *Client) SearchSecretMessages(chatId int64, query string, offset string, limit int32, opts *types.SearchSecretMessagesOpts) (*types.FoundMessages, error) {
	req := &types.SearchSecretMessages{
		ChatId: chatId,
		Query:  query,
		Offset: offset,
		Limit:  limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundMessages), nil
}

// SearchSavedMessages Searches for messages tagged by the given reaction and with the given words in the Saved Messages chat; for Telegram Premium users only.
func (c *Client) SearchSavedMessages(savedMessagesTopicId int64, query string, fromMessageId int64, offset int32, limit int32, opts *types.SearchSavedMessagesOpts) (*types.FoundChatMessages, error) {
	req := &types.SearchSavedMessages{
		SavedMessagesTopicId: savedMessagesTopicId,
		Query:                query,
		FromMessageId:        fromMessageId,
		Offset:               offset,
		Limit:                limit,
	}
	if opts != nil {
		req.Tag = opts.Tag
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundChatMessages), nil
}

// SearchCallMessages Searches for call and group call messages. Returns the results in reverse chronological order (i.e., in order of decreasing message_id). For optimal performance, the number of returned messages is chosen by TDLib
func (c *Client) SearchCallMessages(offset string, limit int32, onlyMissed bool) (*types.FoundMessages, error) {
	req := &types.SearchCallMessages{
		Offset:     offset,
		Limit:      limit,
		OnlyMissed: onlyMissed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundMessages), nil
}

// SearchOutgoingDocumentMessages Searches for outgoing messages with content of the type messageDocument in all chats except secret chats. Returns the results in reverse chronological order
func (c *Client) SearchOutgoingDocumentMessages(query string, limit int32) (*types.FoundMessages, error) {
	req := &types.SearchOutgoingDocumentMessages{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundMessages), nil
}

// GetPublicPostSearchLimits Checks public post search limits without actually performing the search @query Query that will be searched for
func (c *Client) GetPublicPostSearchLimits(query string) (*types.PublicPostSearchLimits, error) {
	req := &types.GetPublicPostSearchLimits{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PublicPostSearchLimits), nil
}

// SearchPublicPosts Searches for public channel posts using the given query. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicPosts(query string, offset string, limit int32, starCount int64) (*types.FoundPublicPosts, error) {
	req := &types.SearchPublicPosts{
		Query:     query,
		Offset:    offset,
		Limit:     limit,
		StarCount: starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundPublicPosts), nil
}

// SearchPublicMessagesByTag Searches for public channel posts containing the given hashtag or cashtag. For optimal performance, the number of returned messages is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicMessagesByTag(tag string, offset string, limit int32) (*types.FoundMessages, error) {
	req := &types.SearchPublicMessagesByTag{
		Tag:    tag,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundMessages), nil
}

// SearchPublicStoriesByTag Searches for public stories containing the given hashtag or cashtag. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByTag(storyPosterChatId int64, tag string, offset string, limit int32) (*types.FoundStories, error) {
	req := &types.SearchPublicStoriesByTag{
		StoryPosterChatId: storyPosterChatId,
		Tag:               tag,
		Offset:            offset,
		Limit:             limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundStories), nil
}

// SearchPublicStoriesByLocation Searches for public stories by the given address location. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByLocation(address *types.LocationAddress, offset string, limit int32) (*types.FoundStories, error) {
	req := &types.SearchPublicStoriesByLocation{
		Address: address,
		Offset:  offset,
		Limit:   limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundStories), nil
}

// SearchPublicStoriesByVenue Searches for public stories from the given venue. For optimal performance, the number of returned stories is chosen by TDLib and can be smaller than the specified limit
func (c *Client) SearchPublicStoriesByVenue(venueProvider string, venueId string, offset string, limit int32) (*types.FoundStories, error) {
	req := &types.SearchPublicStoriesByVenue{
		VenueProvider: venueProvider,
		VenueId:       venueId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundStories), nil
}

// GetSearchedForTags Returns recently searched for hashtags or cashtags by their prefix @tag_prefix Prefix of hashtags or cashtags to return @limit The maximum number of items to be returned
func (c *Client) GetSearchedForTags(tagPrefix string, limit int32) (*types.Hashtags, error) {
	req := &types.GetSearchedForTags{
		TagPrefix: tagPrefix,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Hashtags), nil
}

// RemoveSearchedForTag Removes a hashtag or a cashtag from the list of recently searched for hashtags or cashtags @tag Hashtag or cashtag to delete
func (c *Client) RemoveSearchedForTag(tag string) (*types.Ok, error) {
	req := &types.RemoveSearchedForTag{
		Tag: tag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClearSearchedForTags Clears the list of recently searched for hashtags or cashtags @clear_cashtags Pass true to clear the list of recently searched for cashtags; otherwise, the list of recently searched for hashtags will be cleared
func (c *Client) ClearSearchedForTags(clearCashtags bool) (*types.Ok, error) {
	req := &types.ClearSearchedForTags{
		ClearCashtags: clearCashtags,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteAllCallMessages Deletes all call messages @revoke Pass true to delete the messages for all users
func (c *Client) DeleteAllCallMessages(revoke bool) (*types.Ok, error) {
	req := &types.DeleteAllCallMessages{
		Revoke: revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user @chat_id Chat identifier @limit The maximum number of messages to be returned
func (c *Client) SearchChatRecentLocationMessages(chatId int64, limit int32) (*types.Messages, error) {
	req := &types.SearchChatRecentLocationMessages{
		ChatId: chatId,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date. Returns a 404 error if such message doesn't exist
func (c *Client) GetChatMessageByDate(chatId int64, date int32) (*types.Message, error) {
	req := &types.GetChatMessageByDate{
		ChatId: chatId,
		Date:   date,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// GetChatSparseMessagePositions Returns sparse positions of messages of the specified type in the chat to be used for shared media scroll implementation. Returns the results in reverse chronological order (i.e., in order of decreasing message_id).
func (c *Client) GetChatSparseMessagePositions(chatId int64, filter *types.SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*types.MessagePositions, error) {
	req := &types.GetChatSparseMessagePositions{
		ChatId:               chatId,
		Filter:               filter,
		FromMessageId:        fromMessageId,
		Limit:                limit,
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessagePositions), nil
}

// GetChatMessageCalendar Returns information about the next messages of the specified type in the chat split by days. Returns the results in reverse chronological order. Can return partial result for the last returned day. Behavior of this method depends on the value of the option "utc_time_offset"
func (c *Client) GetChatMessageCalendar(chatId int64, filter *types.SearchMessagesFilter, fromMessageId int64, opts *types.GetChatMessageCalendarOpts) (*types.MessageCalendar, error) {
	req := &types.GetChatMessageCalendar{
		ChatId:        chatId,
		Filter:        filter,
		FromMessageId: fromMessageId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageCalendar), nil
}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat or its topic
func (c *Client) GetChatMessageCount(chatId int64, filter *types.SearchMessagesFilter, returnLocal bool, opts *types.GetChatMessageCountOpts) (*types.Count, error) {
	req := &types.GetChatMessageCount{
		ChatId:      chatId,
		Filter:      filter,
		ReturnLocal: returnLocal,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// GetChatMessagePosition Returns approximate 1-based position of a message among messages, which can be found by the specified filter in the chat and topic. Cannot be used in secret chats
func (c *Client) GetChatMessagePosition(chatId int64, filter *types.SearchMessagesFilter, messageId int64, opts *types.GetChatMessagePositionOpts) (*types.Count, error) {
	req := &types.GetChatMessagePosition{
		ChatId:    chatId,
		Filter:    filter,
		MessageId: messageId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in reverse chronological order (i.e., in order of decreasing message_id) @chat_id Chat identifier
func (c *Client) GetChatScheduledMessages(chatId int64) (*types.Messages, error) {
	req := &types.GetChatScheduledMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetChatSponsoredMessages Returns sponsored messages to be shown in a chat; for channel chats and chats with bots only @chat_id Identifier of the chat
func (c *Client) GetChatSponsoredMessages(chatId int64) (*types.SponsoredMessages, error) {
	req := &types.GetChatSponsoredMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SponsoredMessages), nil
}

// ClickChatSponsoredMessage Informs TDLib that the user opened the sponsored chat via the button, the name, the chat photo, a mention in the sponsored message text, or the media in the sponsored message
func (c *Client) ClickChatSponsoredMessage(chatId int64, messageId int64, isMediaClick bool, fromFullscreen bool) (*types.Ok, error) {
	req := &types.ClickChatSponsoredMessage{
		ChatId:         chatId,
		MessageId:      messageId,
		IsMediaClick:   isMediaClick,
		FromFullscreen: fromFullscreen,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportChatSponsoredMessage Reports a sponsored message to Telegram moderators
func (c *Client) ReportChatSponsoredMessage(chatId int64, messageId int64, optionId string) (*types.ReportSponsoredResult, error) {
	req := &types.ReportChatSponsoredMessage{
		ChatId:    chatId,
		MessageId: messageId,
		OptionId:  optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReportSponsoredResult), nil
}

// GetSearchSponsoredChats Returns sponsored chats to be shown in the search results @query Query the user searches for
func (c *Client) GetSearchSponsoredChats(query string) (*types.SponsoredChats, error) {
	req := &types.GetSearchSponsoredChats{
		Query: query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SponsoredChats), nil
}

// ViewSponsoredChat Informs TDLib that the user fully viewed a sponsored chat @sponsored_chat_unique_id Unique identifier of the sponsored chat
func (c *Client) ViewSponsoredChat(sponsoredChatUniqueId int64) (*types.Ok, error) {
	req := &types.ViewSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// OpenSponsoredChat Informs TDLib that the user opened a sponsored chat @sponsored_chat_unique_id Unique identifier of the sponsored chat
func (c *Client) OpenSponsoredChat(sponsoredChatUniqueId int64) (*types.Ok, error) {
	req := &types.OpenSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportSponsoredChat Reports a sponsored chat to Telegram moderators
func (c *Client) ReportSponsoredChat(sponsoredChatUniqueId int64, optionId string) (*types.ReportSponsoredResult, error) {
	req := &types.ReportSponsoredChat{
		SponsoredChatUniqueId: sponsoredChatUniqueId,
		OptionId:              optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReportSponsoredResult), nil
}

// GetVideoMessageAdvertisements Returns advertisements to be shown while a video from a message is watched. Available only if messageProperties.can_get_video_advertisements
func (c *Client) GetVideoMessageAdvertisements(chatId int64, messageId int64) (*types.VideoMessageAdvertisements, error) {
	req := &types.GetVideoMessageAdvertisements{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.VideoMessageAdvertisements), nil
}

// ViewVideoMessageAdvertisement Informs TDLib that the user viewed a video message advertisement @advertisement_unique_id Unique identifier of the advertisement
func (c *Client) ViewVideoMessageAdvertisement(advertisementUniqueId int64) (*types.Ok, error) {
	req := &types.ViewVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClickVideoMessageAdvertisement Informs TDLib that the user clicked a video message advertisement @advertisement_unique_id Unique identifier of the advertisement
func (c *Client) ClickVideoMessageAdvertisement(advertisementUniqueId int64) (*types.Ok, error) {
	req := &types.ClickVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportVideoMessageAdvertisement Reports a video message advertisement to Telegram moderators
func (c *Client) ReportVideoMessageAdvertisement(advertisementUniqueId int64, optionId string) (*types.ReportSponsoredResult, error) {
	req := &types.ReportVideoMessageAdvertisement{
		AdvertisementUniqueId: advertisementUniqueId,
		OptionId:              optionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReportSponsoredResult), nil
}

// RemoveNotification Removes an active notification from notification list. Needs to be called only if the notification is removed by the current user @notification_group_id Identifier of notification group to which the notification belongs @notification_id Identifier of removed notification
func (c *Client) RemoveNotification(notificationGroupId int32, notificationId int32) (*types.Ok, error) {
	req := &types.RemoveNotification{
		NotificationGroupId: notificationGroupId,
		NotificationId:      notificationId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveNotificationGroup Removes a group of active notifications. Needs to be called only if the notification group is removed by the current user @notification_group_id Notification group identifier @max_notification_id The maximum identifier of removed notifications
func (c *Client) RemoveNotificationGroup(notificationGroupId int32, maxNotificationId int32) (*types.Ok, error) {
	req := &types.RemoveNotificationGroup{
		NotificationGroupId: notificationGroupId,
		MaxNotificationId:   maxNotificationId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only if messageProperties.can_get_link, or if messageProperties.can_get_media_timestamp_links and a media timestamp link is generated. This is an offline method
func (c *Client) GetMessageLink(chatId int64, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*types.MessageLink, error) {
	req := &types.GetMessageLink{
		ChatId:          chatId,
		MessageId:       messageId,
		MediaTimestamp:  mediaTimestamp,
		ForAlbum:        forAlbum,
		InMessageThread: inMessageThread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageLink), nil
}

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only if messageProperties.can_get_embedding_code
func (c *Client) GetMessageEmbeddingCode(chatId int64, messageId int64, forAlbum bool) (*types.Text, error) {
	req := &types.GetMessageEmbeddingCode{
		ChatId:    chatId,
		MessageId: messageId,
		ForAlbum:  forAlbum,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetMessageLinkInfo Returns information about a public or private message link. Can be called for any internal link of the type internalLinkTypeMessage @url The message link
func (c *Client) GetMessageLinkInfo(url string) (*types.MessageLinkInfo, error) {
	req := &types.GetMessageLinkInfo{
		Url: url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageLinkInfo), nil
}

// TranslateText Translates a text to the given language. If the current user is a Telegram Premium user, then text formatting is preserved
func (c *Client) TranslateText(text *types.FormattedText, toLanguageCode string) (*types.FormattedText, error) {
	req := &types.TranslateText{
		Text:           text,
		ToLanguageCode: toLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// TranslateMessageText Extracts text or caption of the given message and translates it to the given language. If the current user is a Telegram Premium user, then text formatting is preserved
func (c *Client) TranslateMessageText(chatId int64, messageId int64, toLanguageCode string) (*types.FormattedText, error) {
	req := &types.TranslateMessageText{
		ChatId:         chatId,
		MessageId:      messageId,
		ToLanguageCode: toLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// SummarizeMessage Summarizes content of the message with non-empty summary_language_code
func (c *Client) SummarizeMessage(chatId int64, messageId int64, translateToLanguageCode string) (*types.FormattedText, error) {
	req := &types.SummarizeMessage{
		ChatId:                  chatId,
		MessageId:               messageId,
		TranslateToLanguageCode: translateToLanguageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// RecognizeSpeech Recognizes speech in a video note or a voice note message
func (c *Client) RecognizeSpeech(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.RecognizeSpeech{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RateSpeechRecognition Rates recognized speech in a video note or a voice note message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message @is_good Pass true if the speech recognition is good
func (c *Client) RateSpeechRecognition(chatId int64, messageId int64, isGood bool) (*types.Ok, error) {
	req := &types.RateSpeechRecognition{
		ChatId:    chatId,
		MessageId: messageId,
		IsGood:    isGood,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatAvailableMessageSenders Returns the list of message sender identifiers, which can be used to send messages in a chat @chat_id Chat identifier
func (c *Client) GetChatAvailableMessageSenders(chatId int64) (*types.ChatMessageSenders, error) {
	req := &types.GetChatAvailableMessageSenders{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatMessageSenders), nil
}

// SetChatMessageSender Selects a message sender to send messages in a chat @chat_id Chat identifier @message_sender_id New message sender for the chat
func (c *Client) SetChatMessageSender(chatId int64, messageSenderId *types.MessageSender) (*types.Ok, error) {
	req := &types.SetChatMessageSender{
		ChatId:          chatId,
		MessageSenderId: messageSenderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendMessage Sends a message. Returns the sent message
func (c *Client) SendMessage(chatId int64, inputMessageContent *types.InputMessageContent, opts *types.SendMessageOpts) (*types.Message, error) {
	req := &types.SendMessage{
		ChatId:              chatId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// SendMessageAlbum Sends 2-10 messages grouped together into an album. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) SendMessageAlbum(chatId int64, inputMessageContents []*types.InputMessageContent, opts *types.SendMessageAlbumOpts) (*types.Messages, error) {
	req := &types.SendMessageAlbum{
		ChatId:               chatId,
		InputMessageContents: inputMessageContents,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command; requires can_invite_users member right. Bots can't be invited to a private chat other than the chat with the bot.
func (c *Client) SendBotStartMessage(botUserId int64, chatId int64, parameter string) (*types.Message, error) {
	req := &types.SendBotStartMessage{
		BotUserId: botUserId,
		ChatId:    chatId,
		Parameter: parameter,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
func (c *Client) SendInlineQueryResultMessage(chatId int64, queryId string, resultId string, hideViaBot bool, opts *types.SendInlineQueryResultMessageOpts) (*types.Message, error) {
	req := &types.SendInlineQueryResultMessage{
		ChatId:     chatId,
		QueryId:    queryId,
		ResultId:   resultId,
		HideViaBot: hideViaBot,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
func (c *Client) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *types.ForwardMessagesOpts) (*types.Messages, error) {
	req := &types.ForwardMessages{
		ChatId:        chatId,
		FromChatId:    fromChatId,
		MessageIds:    messageIds,
		SendCopy:      sendCopy,
		RemoveCaption: removeCaption,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.Options = opts.Options
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// SendQuickReplyShortcutMessages Sends messages from a quick reply shortcut. Requires Telegram Business subscription. Can't be used to send paid messages
func (c *Client) SendQuickReplyShortcutMessages(chatId int64, shortcutId int32, sendingId int32) (*types.Messages, error) {
	req := &types.SendQuickReplyShortcutMessages{
		ChatId:     chatId,
		ShortcutId: shortcutId,
		SendingId:  sendingId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
func (c *Client) ResendMessages(chatId int64, messageIds []int64, paidMessageStarCount int64, opts *types.ResendMessagesOpts) (*types.Messages, error) {
	req := &types.ResendMessages{
		ChatId:               chatId,
		MessageIds:           messageIds,
		PaidMessageStarCount: paidMessageStarCount,
	}
	if opts != nil {
		req.Quote = opts.Quote
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
func (c *Client) AddLocalMessage(chatId int64, senderId *types.MessageSender, disableNotification bool, inputMessageContent *types.InputMessageContent, opts *types.AddLocalMessageOpts) (*types.Message, error) {
	req := &types.AddLocalMessage{
		ChatId:              chatId,
		SenderId:            senderId,
		DisableNotification: disableNotification,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// DeleteMessages Deletes messages
func (c *Client) DeleteMessages(chatId int64, messageIds []int64, revoke bool) (*types.Ok, error) {
	req := &types.DeleteMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
		Revoke:     revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteChatMessagesBySender Deletes all messages sent by the specified message sender in a chat. Supported only for supergroups; requires can_delete_messages administrator right @chat_id Chat identifier @sender_id Identifier of the sender of messages to delete
func (c *Client) DeleteChatMessagesBySender(chatId int64, senderId *types.MessageSender) (*types.Ok, error) {
	req := &types.DeleteChatMessagesBySender{
		ChatId:   chatId,
		SenderId: senderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteChatMessagesByDate Deletes all messages between the specified dates in a chat. Supported only for private chats and basic groups. Messages sent in the last 30 seconds will not be deleted
func (c *Client) DeleteChatMessagesByDate(chatId int64, minDate int32, maxDate int32, revoke bool) (*types.Ok, error) {
	req := &types.DeleteChatMessagesByDate{
		ChatId:  chatId,
		MinDate: minDate,
		MaxDate: maxDate,
		Revoke:  revoke,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageText(chatId int64, messageId int64, inputMessageContent *types.InputMessageContent, opts *types.EditMessageTextOpts) (*types.Message, error) {
	req := &types.EditMessageText{
		ChatId:              chatId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location.
func (c *Client) EditMessageLiveLocation(chatId int64, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *types.EditMessageLiveLocationOpts) (*types.Message, error) {
	req := &types.EditMessageLiveLocation{
		ChatId:               chatId,
		MessageId:            messageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditMessageChecklist Edits the message content of a checklist. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageChecklist(chatId int64, messageId int64, checklist *types.InputChecklist, opts *types.EditMessageChecklistOpts) (*types.Message, error) {
	req := &types.EditMessageChecklist{
		ChatId:    chatId,
		MessageId: messageId,
		Checklist: checklist,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditMessageMedia Edits the media content of a message, including message caption. If only the caption needs to be edited, use editMessageCaption instead.
func (c *Client) EditMessageMedia(chatId int64, messageId int64, inputMessageContent *types.InputMessageContent, opts *types.EditMessageMediaOpts) (*types.Message, error) {
	req := &types.EditMessageMedia{
		ChatId:              chatId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageCaption(chatId int64, messageId int64, showCaptionAboveMedia bool, opts *types.EditMessageCaptionOpts) (*types.Message, error) {
	req := &types.EditMessageCaption{
		ChatId:                chatId,
		MessageId:             messageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
func (c *Client) EditMessageReplyMarkup(chatId int64, messageId int64, opts *types.EditMessageReplyMarkupOpts) (*types.Message, error) {
	req := &types.EditMessageReplyMarkup{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
func (c *Client) EditInlineMessageText(inlineMessageId string, inputMessageContent *types.InputMessageContent, opts *types.EditInlineMessageTextOpts) (*types.Ok, error) {
	req := &types.EditInlineMessageText{
		InlineMessageId:     inlineMessageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageLiveLocation(inlineMessageId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *types.EditInlineMessageLiveLocationOpts) (*types.Ok, error) {
	req := &types.EditInlineMessageLiveLocation{
		InlineMessageId:      inlineMessageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditInlineMessageMedia Edits the media content of a message with a text, an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageMedia(inlineMessageId string, inputMessageContent *types.InputMessageContent, opts *types.EditInlineMessageMediaOpts) (*types.Ok, error) {
	req := &types.EditInlineMessageMedia{
		InlineMessageId:     inlineMessageId,
		InputMessageContent: inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageCaption(inlineMessageId string, showCaptionAboveMedia bool, opts *types.EditInlineMessageCaptionOpts) (*types.Ok, error) {
	req := &types.EditInlineMessageCaption{
		InlineMessageId:       inlineMessageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
func (c *Client) EditInlineMessageReplyMarkup(inlineMessageId string, opts *types.EditInlineMessageReplyMarkupOpts) (*types.Ok, error) {
	req := &types.EditInlineMessageReplyMarkup{
		InlineMessageId: inlineMessageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditMessageSchedulingState Edits the time when a scheduled message will be sent. Scheduling state of all messages in the same album or forwarded together with the message will be also changed
func (c *Client) EditMessageSchedulingState(chatId int64, messageId int64, opts *types.EditMessageSchedulingStateOpts) (*types.Ok, error) {
	req := &types.EditMessageSchedulingState{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.SchedulingState = opts.SchedulingState
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetMessageFactCheck Changes the fact-check of a message. Can be only used if messageProperties.can_set_fact_check == true
func (c *Client) SetMessageFactCheck(chatId int64, messageId int64, opts *types.SetMessageFactCheckOpts) (*types.Ok, error) {
	req := &types.SetMessageFactCheck{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.Text = opts.Text
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendBusinessMessage Sends a message on behalf of a business account; for bots only. Returns the message after it was sent
func (c *Client) SendBusinessMessage(businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContent *types.InputMessageContent, opts *types.SendBusinessMessageOpts) (*types.BusinessMessage, error) {
	req := &types.SendBusinessMessage{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		DisableNotification:  disableNotification,
		ProtectContent:       protectContent,
		EffectId:             effectId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// SendBusinessMessageAlbum Sends 2-10 messages grouped together into an album on behalf of a business account; for bots only. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) SendBusinessMessageAlbum(businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*types.InputMessageContent, opts *types.SendBusinessMessageAlbumOpts) (*types.BusinessMessages, error) {
	req := &types.SendBusinessMessageAlbum{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		DisableNotification:  disableNotification,
		ProtectContent:       protectContent,
		EffectId:             effectId,
		InputMessageContents: inputMessageContents,
	}
	if opts != nil {
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessages), nil
}

// EditBusinessMessageText Edits the text of a text or game message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageText(businessConnectionId string, chatId int64, messageId int64, inputMessageContent *types.InputMessageContent, opts *types.EditBusinessMessageTextOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageText{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// EditBusinessMessageLiveLocation Edits the content of a live location in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageLiveLocation(businessConnectionId string, chatId int64, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *types.EditBusinessMessageLiveLocationOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageLiveLocation{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		LivePeriod:           livePeriod,
		Heading:              heading,
		ProximityAlertRadius: proximityAlertRadius,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// EditBusinessMessageChecklist Edits the content of a checklist in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageChecklist(businessConnectionId string, chatId int64, messageId int64, checklist *types.InputChecklist, opts *types.EditBusinessMessageChecklistOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageChecklist{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		Checklist:            checklist,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// EditBusinessMessageMedia Edits the media content of a message with a text, an animation, an audio, a document, a photo or a video in a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageMedia(businessConnectionId string, chatId int64, messageId int64, inputMessageContent *types.InputMessageContent, opts *types.EditBusinessMessageMediaOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageMedia{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		InputMessageContent:  inputMessageContent,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// EditBusinessMessageCaption Edits the caption of a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageCaption(businessConnectionId string, chatId int64, messageId int64, showCaptionAboveMedia bool, opts *types.EditBusinessMessageCaptionOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageCaption{
		BusinessConnectionId:  businessConnectionId,
		ChatId:                chatId,
		MessageId:             messageId,
		ShowCaptionAboveMedia: showCaptionAboveMedia,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// EditBusinessMessageReplyMarkup Edits the reply markup of a message sent on behalf of a business account; for bots only
func (c *Client) EditBusinessMessageReplyMarkup(businessConnectionId string, chatId int64, messageId int64, opts *types.EditBusinessMessageReplyMarkupOpts) (*types.BusinessMessage, error) {
	req := &types.EditBusinessMessageReplyMarkup{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// StopBusinessPoll Stops a poll sent on behalf of a business account; for bots only
func (c *Client) StopBusinessPoll(businessConnectionId string, chatId int64, messageId int64, opts *types.StopBusinessPollOpts) (*types.BusinessMessage, error) {
	req := &types.StopBusinessPoll{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessMessage), nil
}

// SetBusinessMessageIsPinned Pins or unpins a message sent on behalf of a business account; for bots only
func (c *Client) SetBusinessMessageIsPinned(businessConnectionId string, chatId int64, messageId int64, isPinned bool) (*types.Ok, error) {
	req := &types.SetBusinessMessageIsPinned{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
		IsPinned:             isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadBusinessMessage Reads a message on behalf of a business account; for bots only
func (c *Client) ReadBusinessMessage(businessConnectionId string, chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.ReadBusinessMessage{
		BusinessConnectionId: businessConnectionId,
		ChatId:               chatId,
		MessageId:            messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteBusinessMessages Deletes messages on behalf of a business account; for bots only
func (c *Client) DeleteBusinessMessages(businessConnectionId string, messageIds []int64) (*types.Ok, error) {
	req := &types.DeleteBusinessMessages{
		BusinessConnectionId: businessConnectionId,
		MessageIds:           messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditBusinessStory Changes a story posted by the bot on behalf of a business account; for bots only
func (c *Client) EditBusinessStory(storyPosterChatId int64, storyId int32, content *types.InputStoryContent, areas *types.InputStoryAreas, caption *types.FormattedText, privacySettings *types.StoryPrivacySettings) (*types.Story, error) {
	req := &types.EditBusinessStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		Content:           content,
		Areas:             areas,
		Caption:           caption,
		PrivacySettings:   privacySettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Story), nil
}

// DeleteBusinessStory Deletes a story posted by the bot on behalf of a business account; for bots only
func (c *Client) DeleteBusinessStory(businessConnectionId string, storyId int32) (*types.Ok, error) {
	req := &types.DeleteBusinessStory{
		BusinessConnectionId: businessConnectionId,
		StoryId:              storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAccountName Changes the first and last name of a business account; for bots only
func (c *Client) SetBusinessAccountName(businessConnectionId string, firstName string, lastName string) (*types.Ok, error) {
	req := &types.SetBusinessAccountName{
		BusinessConnectionId: businessConnectionId,
		FirstName:            firstName,
		LastName:             lastName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAccountBio Changes the bio of a business account; for bots only
func (c *Client) SetBusinessAccountBio(businessConnectionId string, bio string) (*types.Ok, error) {
	req := &types.SetBusinessAccountBio{
		BusinessConnectionId: businessConnectionId,
		Bio:                  bio,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAccountProfilePhoto Changes a profile photo of a business account; for bots only
func (c *Client) SetBusinessAccountProfilePhoto(businessConnectionId string, isPublic bool, opts *types.SetBusinessAccountProfilePhotoOpts) (*types.Ok, error) {
	req := &types.SetBusinessAccountProfilePhoto{
		BusinessConnectionId: businessConnectionId,
		IsPublic:             isPublic,
	}
	if opts != nil {
		req.Photo = opts.Photo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAccountUsername Changes the editable username of a business account; for bots only
func (c *Client) SetBusinessAccountUsername(businessConnectionId string, username string) (*types.Ok, error) {
	req := &types.SetBusinessAccountUsername{
		BusinessConnectionId: businessConnectionId,
		Username:             username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAccountGiftSettings Changes settings for gift receiving of a business account; for bots only
func (c *Client) SetBusinessAccountGiftSettings(businessConnectionId string, settings *types.GiftSettings) (*types.Ok, error) {
	req := &types.SetBusinessAccountGiftSettings{
		BusinessConnectionId: businessConnectionId,
		Settings:             settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBusinessAccountStarAmount Returns the amount of Telegram Stars owned by a business account; for bots only @business_connection_id Unique identifier of business connection
func (c *Client) GetBusinessAccountStarAmount(businessConnectionId string) (*types.StarAmount, error) {
	req := &types.GetBusinessAccountStarAmount{
		BusinessConnectionId: businessConnectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarAmount), nil
}

// TransferBusinessAccountStars Transfers Telegram Stars from the business account to the business bot; for bots only
func (c *Client) TransferBusinessAccountStars(businessConnectionId string, starCount int64) (*types.Ok, error) {
	req := &types.TransferBusinessAccountStars{
		BusinessConnectionId: businessConnectionId,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckQuickReplyShortcutName Checks validness of a name for a quick reply shortcut. Can be called synchronously @name The name of the shortcut; 1-32 characters
func (c *Client) CheckQuickReplyShortcutName(name string) (*types.Ok, error) {
	req := &types.CheckQuickReplyShortcutName{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LoadQuickReplyShortcuts Loads quick reply shortcuts created by the current user. The loaded data will be sent through updateQuickReplyShortcut and updateQuickReplyShortcuts
func (c *Client) LoadQuickReplyShortcuts() (*types.Ok, error) {
	req := &types.LoadQuickReplyShortcuts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetQuickReplyShortcutName Changes name of a quick reply shortcut @shortcut_id Unique identifier of the quick reply shortcut @name New name for the shortcut. Use checkQuickReplyShortcutName to check its validness
func (c *Client) SetQuickReplyShortcutName(shortcutId int32, name string) (*types.Ok, error) {
	req := &types.SetQuickReplyShortcutName{
		ShortcutId: shortcutId,
		Name:       name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteQuickReplyShortcut Deletes a quick reply shortcut @shortcut_id Unique identifier of the quick reply shortcut
func (c *Client) DeleteQuickReplyShortcut(shortcutId int32) (*types.Ok, error) {
	req := &types.DeleteQuickReplyShortcut{
		ShortcutId: shortcutId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReorderQuickReplyShortcuts Changes the order of quick reply shortcuts @shortcut_ids The new order of quick reply shortcuts
func (c *Client) ReorderQuickReplyShortcuts(shortcutIds []int32) (*types.Ok, error) {
	req := &types.ReorderQuickReplyShortcuts{
		ShortcutIds: shortcutIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LoadQuickReplyShortcutMessages Loads quick reply messages that can be sent by a given quick reply shortcut. The loaded messages will be sent through updateQuickReplyShortcutMessages
func (c *Client) LoadQuickReplyShortcutMessages(shortcutId int32) (*types.Ok, error) {
	req := &types.LoadQuickReplyShortcutMessages{
		ShortcutId: shortcutId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteQuickReplyShortcutMessages Deletes specified quick reply messages
func (c *Client) DeleteQuickReplyShortcutMessages(shortcutId int32, messageIds []int64) (*types.Ok, error) {
	req := &types.DeleteQuickReplyShortcutMessages{
		ShortcutId: shortcutId,
		MessageIds: messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddQuickReplyShortcutMessage Adds a message to a quick reply shortcut. If shortcut doesn't exist and there are less than getOption("quick_reply_shortcut_count_max") shortcuts, then a new shortcut is created.
func (c *Client) AddQuickReplyShortcutMessage(shortcutName string, replyToMessageId int64, inputMessageContent *types.InputMessageContent) (*types.QuickReplyMessage, error) {
	req := &types.AddQuickReplyShortcutMessage{
		ShortcutName:        shortcutName,
		ReplyToMessageId:    replyToMessageId,
		InputMessageContent: inputMessageContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.QuickReplyMessage), nil
}

// AddQuickReplyShortcutInlineQueryResultMessage Adds a message to a quick reply shortcut via inline bot. If shortcut doesn't exist and there are less than getOption("quick_reply_shortcut_count_max") shortcuts, then a new shortcut is created.
func (c *Client) AddQuickReplyShortcutInlineQueryResultMessage(shortcutName string, replyToMessageId int64, queryId string, resultId string, hideViaBot bool) (*types.QuickReplyMessage, error) {
	req := &types.AddQuickReplyShortcutInlineQueryResultMessage{
		ShortcutName:     shortcutName,
		ReplyToMessageId: replyToMessageId,
		QueryId:          queryId,
		ResultId:         resultId,
		HideViaBot:       hideViaBot,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.QuickReplyMessage), nil
}

// AddQuickReplyShortcutMessageAlbum Adds 2-10 messages grouped together into an album to a quick reply shortcut. Currently, only audio, document, photo and video messages can be grouped into an album.
func (c *Client) AddQuickReplyShortcutMessageAlbum(shortcutName string, replyToMessageId int64, inputMessageContents []*types.InputMessageContent) (*types.QuickReplyMessages, error) {
	req := &types.AddQuickReplyShortcutMessageAlbum{
		ShortcutName:         shortcutName,
		ReplyToMessageId:     replyToMessageId,
		InputMessageContents: inputMessageContents,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.QuickReplyMessages), nil
}

// ReaddQuickReplyShortcutMessages Readds quick reply messages which failed to add. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
func (c *Client) ReaddQuickReplyShortcutMessages(shortcutName string, messageIds []int64) (*types.QuickReplyMessages, error) {
	req := &types.ReaddQuickReplyShortcutMessages{
		ShortcutName: shortcutName,
		MessageIds:   messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.QuickReplyMessages), nil
}

// EditQuickReplyMessage Asynchronously edits the text, media or caption of a quick reply message. Use quickReplyMessage.can_be_edited to check whether a message can be edited.
func (c *Client) EditQuickReplyMessage(shortcutId int32, messageId int64, inputMessageContent *types.InputMessageContent) (*types.Ok, error) {
	req := &types.EditQuickReplyMessage{
		ShortcutId:          shortcutId,
		MessageId:           messageId,
		InputMessageContent: inputMessageContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetForumTopicDefaultIcons Returns the list of custom emoji, which can be used as forum topic icon by all users
func (c *Client) GetForumTopicDefaultIcons() (*types.Stickers, error) {
	req := &types.GetForumTopicDefaultIcons{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// CreateForumTopic Creates a topic in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator or can_create_topics member right in the supergroup
func (c *Client) CreateForumTopic(chatId int64, name string, isNameImplicit bool, icon *types.ForumTopicIcon) (*types.ForumTopicInfo, error) {
	req := &types.CreateForumTopic{
		ChatId:         chatId,
		Name:           name,
		IsNameImplicit: isNameImplicit,
		Icon:           icon,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ForumTopicInfo), nil
}

// EditForumTopic Edits title and icon of a topic in a forum supergroup chat or a chat with a bot with topics; for supergroup chats requires can_manage_topics administrator right
func (c *Client) EditForumTopic(chatId int64, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*types.Ok, error) {
	req := &types.EditForumTopic{
		ChatId:              chatId,
		ForumTopicId:        forumTopicId,
		Name:                name,
		EditIconCustomEmoji: editIconCustomEmoji,
		IconCustomEmojiId:   iconCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetForumTopic Returns information about a topic in a forum supergroup chat or a chat with a bot with topics
func (c *Client) GetForumTopic(chatId int64, forumTopicId int32) (*types.ForumTopic, error) {
	req := &types.GetForumTopic{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ForumTopic), nil
}

// GetForumTopicHistory Returns messages in a topic in a forum supergroup chat or a chat with a bot with topics. The messages are returned in reverse chronological order
func (c *Client) GetForumTopicHistory(chatId int64, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*types.Messages, error) {
	req := &types.GetForumTopicHistory{
		ChatId:        chatId,
		ForumTopicId:  forumTopicId,
		FromMessageId: fromMessageId,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Messages), nil
}

// GetForumTopicLink Returns an HTTPS link to a topic in a forum supergroup chat. This is an offline method @chat_id Identifier of the chat @forum_topic_id Forum topic identifier
func (c *Client) GetForumTopicLink(chatId int64, forumTopicId int32) (*types.MessageLink, error) {
	req := &types.GetForumTopicLink{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageLink), nil
}

// GetForumTopics Returns found forum topics in a forum supergroup chat or a chat with a bot with topics. This is a temporary method for getting information about topic list from the server
func (c *Client) GetForumTopics(chatId int64, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*types.ForumTopics, error) {
	req := &types.GetForumTopics{
		ChatId:             chatId,
		Query:              query,
		OffsetDate:         offsetDate,
		OffsetMessageId:    offsetMessageId,
		OffsetForumTopicId: offsetForumTopicId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ForumTopics), nil
}

// SetForumTopicNotificationSettings Changes the notification settings of a forum topic in a forum supergroup chat or a chat with a bot with topics
func (c *Client) SetForumTopicNotificationSettings(chatId int64, forumTopicId int32, notificationSettings *types.ChatNotificationSettings) (*types.Ok, error) {
	req := &types.SetForumTopicNotificationSettings{
		ChatId:               chatId,
		ForumTopicId:         forumTopicId,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleForumTopicIsClosed Toggles whether a topic is closed in a forum supergroup chat; requires can_manage_topics administrator right in the supergroup unless the user is creator of the topic
func (c *Client) ToggleForumTopicIsClosed(chatId int64, forumTopicId int32, isClosed bool) (*types.Ok, error) {
	req := &types.ToggleForumTopicIsClosed{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		IsClosed:     isClosed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGeneralForumTopicIsHidden Toggles whether a General topic is hidden in a forum supergroup chat; requires can_manage_topics administrator right in the supergroup
func (c *Client) ToggleGeneralForumTopicIsHidden(chatId int64, isHidden bool) (*types.Ok, error) {
	req := &types.ToggleGeneralForumTopicIsHidden{
		ChatId:   chatId,
		IsHidden: isHidden,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleForumTopicIsPinned Changes the pinned state of a topic in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator right in the supergroup.
func (c *Client) ToggleForumTopicIsPinned(chatId int64, forumTopicId int32, isPinned bool) (*types.Ok, error) {
	req := &types.ToggleForumTopicIsPinned{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		IsPinned:     isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPinnedForumTopics Changes the order of pinned topics in a forum supergroup chat or a chat with a bot with topics; requires can_manage_topics administrator right in the supergroup
func (c *Client) SetPinnedForumTopics(chatId int64, forumTopicIds []int32) (*types.Ok, error) {
	req := &types.SetPinnedForumTopics{
		ChatId:        chatId,
		ForumTopicIds: forumTopicIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteForumTopic Deletes all messages from a topic in a forum supergroup chat or a chat with a bot with topics; requires can_delete_messages administrator right in the supergroup
func (c *Client) DeleteForumTopic(chatId int64, forumTopicId int32) (*types.Ok, error) {
	req := &types.DeleteForumTopic{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadAllForumTopicMentions Marks all mentions in a topic in a forum supergroup chat as read
func (c *Client) ReadAllForumTopicMentions(chatId int64, forumTopicId int32) (*types.Ok, error) {
	req := &types.ReadAllForumTopicMentions{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadAllForumTopicReactions Marks all reactions in a topic in a forum supergroup chat or a chat with a bot with topics as read
func (c *Client) ReadAllForumTopicReactions(chatId int64, forumTopicId int32) (*types.Ok, error) {
	req := &types.ReadAllForumTopicReactions{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// UnpinAllForumTopicMessages Removes all pinned messages from a topic in a forum supergroup chat or a chat with a bot with topics; requires can_pin_messages member right in the supergroup
func (c *Client) UnpinAllForumTopicMessages(chatId int64, forumTopicId int32) (*types.Ok, error) {
	req := &types.UnpinAllForumTopicMessages{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPasskeyParameters Returns parameters for creating of a new passkey as JSON-serialized string
func (c *Client) GetPasskeyParameters() (*types.Text, error) {
	req := &types.GetPasskeyParameters{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// AddLoginPasskey Adds a passkey allowed to be used for the login by the current user and returns the added passkey. Call getPasskeyParameters to get parameters for creating of the passkey
func (c *Client) AddLoginPasskey(clientData string, attestationObject string) (*types.Passkey, error) {
	req := &types.AddLoginPasskey{
		ClientData:        clientData,
		AttestationObject: attestationObject,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Passkey), nil
}

// GetLoginPasskeys Returns the list of passkeys allowed to be used for the login by the current user
func (c *Client) GetLoginPasskeys() (*types.Passkeys, error) {
	req := &types.GetLoginPasskeys{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Passkeys), nil
}

// RemoveLoginPasskey Removes a passkey from the list of passkeys allowed to be used for the login by the current user @passkey_id Unique identifier of the passkey to remove
func (c *Client) RemoveLoginPasskey(passkeyId string) (*types.Ok, error) {
	req := &types.RemoveLoginPasskey{
		PasskeyId: passkeyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetEmojiReaction Returns information about an emoji reaction. Returns a 404 error if the reaction is not found @emoji Text representation of the reaction
func (c *Client) GetEmojiReaction(emoji string) (*types.EmojiReaction, error) {
	req := &types.GetEmojiReaction{
		Emoji: emoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiReaction), nil
}

// GetCustomEmojiReactionAnimations Returns TGS stickers with generic animations for custom emoji reactions
func (c *Client) GetCustomEmojiReactionAnimations() (*types.Stickers, error) {
	req := &types.GetCustomEmojiReactionAnimations{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetMessageAvailableReactions Returns reactions, which can be added to a message. The list can change after updateActiveEmojiReactions, updateChatAvailableReactions for the chat, or updateMessageInteractionInfo for the message
func (c *Client) GetMessageAvailableReactions(chatId int64, messageId int64, rowSize int32) (*types.AvailableReactions, error) {
	req := &types.GetMessageAvailableReactions{
		ChatId:    chatId,
		MessageId: messageId,
		RowSize:   rowSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AvailableReactions), nil
}

// ClearRecentReactions Clears the list of recently used reactions
func (c *Client) ClearRecentReactions() (*types.Ok, error) {
	req := &types.ClearRecentReactions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddMessageReaction Adds a reaction or a tag to a message. Use getMessageAvailableReactions to receive the list of available reactions for the message
func (c *Client) AddMessageReaction(chatId int64, messageId int64, reactionType *types.ReactionType, isBig bool, updateRecentReactions bool) (*types.Ok, error) {
	req := &types.AddMessageReaction{
		ChatId:                chatId,
		MessageId:             messageId,
		ReactionType:          reactionType,
		IsBig:                 isBig,
		UpdateRecentReactions: updateRecentReactions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveMessageReaction Removes a reaction from a message. A chosen reaction can always be removed
func (c *Client) RemoveMessageReaction(chatId int64, messageId int64, reactionType *types.ReactionType) (*types.Ok, error) {
	req := &types.RemoveMessageReaction{
		ChatId:       chatId,
		MessageId:    messageId,
		ReactionType: reactionType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatAvailablePaidMessageReactionSenders Returns the list of message sender identifiers, which can be used to send a paid reaction in a chat @chat_id Chat identifier
func (c *Client) GetChatAvailablePaidMessageReactionSenders(chatId int64) (*types.MessageSenders, error) {
	req := &types.GetChatAvailablePaidMessageReactionSenders{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageSenders), nil
}

// AddPendingPaidMessageReaction Adds the paid message reaction to a message. Use getMessageAvailableReactions to check whether the reaction is available for the message
func (c *Client) AddPendingPaidMessageReaction(chatId int64, messageId int64, starCount int64, opts *types.AddPendingPaidMessageReactionOpts) (*types.Ok, error) {
	req := &types.AddPendingPaidMessageReaction{
		ChatId:    chatId,
		MessageId: messageId,
		StarCount: starCount,
	}
	if opts != nil {
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CommitPendingPaidMessageReactions Applies all pending paid reactions on a message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message
func (c *Client) CommitPendingPaidMessageReactions(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.CommitPendingPaidMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemovePendingPaidMessageReactions Removes all pending paid reactions on a message @chat_id Identifier of the chat to which the message belongs @message_id Identifier of the message
func (c *Client) RemovePendingPaidMessageReactions(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.RemovePendingPaidMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPaidMessageReactionType Changes type of paid message reaction of the current user on a message. The message must have paid reaction added by the current user
func (c *Client) SetPaidMessageReactionType(chatId int64, messageId int64, typeField *types.PaidReactionType) (*types.Ok, error) {
	req := &types.SetPaidMessageReactionType{
		ChatId:    chatId,
		MessageId: messageId,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetMessageReactions Sets reactions on a message; for bots only
func (c *Client) SetMessageReactions(chatId int64, messageId int64, reactionTypes []*types.ReactionType, isBig bool) (*types.Ok, error) {
	req := &types.SetMessageReactions{
		ChatId:        chatId,
		MessageId:     messageId,
		ReactionTypes: reactionTypes,
		IsBig:         isBig,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetMessageAddedReactions Returns reactions added for a message, along with their sender
func (c *Client) GetMessageAddedReactions(chatId int64, messageId int64, offset string, limit int32, opts *types.GetMessageAddedReactionsOpts) (*types.AddedReactions, error) {
	req := &types.GetMessageAddedReactions{
		ChatId:    chatId,
		MessageId: messageId,
		Offset:    offset,
		Limit:     limit,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AddedReactions), nil
}

// SetDefaultReactionType Changes type of default reaction for the current user @reaction_type New type of the default reaction. The paid reaction can't be set as default
func (c *Client) SetDefaultReactionType(reactionType *types.ReactionType) (*types.Ok, error) {
	req := &types.SetDefaultReactionType{
		ReactionType: reactionType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSavedMessagesTags Returns tags used in Saved Messages or a Saved Messages topic
func (c *Client) GetSavedMessagesTags(savedMessagesTopicId int64) (*types.SavedMessagesTags, error) {
	req := &types.GetSavedMessagesTags{
		SavedMessagesTopicId: savedMessagesTopicId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SavedMessagesTags), nil
}

// SetSavedMessagesTagLabel Changes label of a Saved Messages tag; for Telegram Premium users only @tag The tag which label will be changed @label New label for the tag; 0-12 characters
func (c *Client) SetSavedMessagesTagLabel(tag *types.ReactionType, label string) (*types.Ok, error) {
	req := &types.SetSavedMessagesTagLabel{
		Tag:   tag,
		Label: label,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetMessageEffect Returns information about a message effect. Returns a 404 error if the effect is not found @effect_id Unique identifier of the effect
func (c *Client) GetMessageEffect(effectId string) (*types.MessageEffect, error) {
	req := &types.GetMessageEffect{
		EffectId: effectId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageEffect), nil
}

// SearchQuote Searches for a given quote in a text. Returns found quote start position in UTF-16 code units. Returns a 404 error if the quote is not found. Can be called synchronously
func (c *Client) SearchQuote(text *types.FormattedText, quote *types.FormattedText, quotePosition int32) (*types.FoundPosition, error) {
	req := &types.SearchQuote{
		Text:          text,
		Quote:         quote,
		QuotePosition: quotePosition,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundPosition), nil
}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) found in the text. Can be called synchronously @text The text in which to look for entities
func (c *Client) GetTextEntities(text string) (*types.TextEntities, error) {
	req := &types.GetTextEntities{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TextEntities), nil
}

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Spoiler, CustomEmoji, BlockQuote, ExpandableBlockQuote, Code, Pre, PreCode, TextUrl
func (c *Client) ParseTextEntities(text string, parseMode *types.TextParseMode) (*types.FormattedText, error) {
	req := &types.ParseTextEntities{
		Text:      text,
		ParseMode: parseMode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
func (c *Client) ParseMarkdown(text *types.FormattedText) (*types.FormattedText, error) {
	req := &types.ParseMarkdown{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously @text The text
func (c *Client) GetMarkdownText(text *types.FormattedText) (*types.FormattedText, error) {
	req := &types.GetMarkdownText{
		Text: text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FormattedText), nil
}

// GetCountryFlagEmoji Returns an emoji for the given country. Returns an empty string on failure. Can be called synchronously @country_code A two-letter ISO 3166-1 alpha-2 country code as received from getCountries
func (c *Client) GetCountryFlagEmoji(countryCode string) (*types.Text, error) {
	req := &types.GetCountryFlagEmoji{
		CountryCode: countryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously @file_name The name of the file or path to the file
func (c *Client) GetFileMimeType(fileName string) (*types.Text, error) {
	req := &types.GetFileMimeType{
		FileName: fileName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously @mime_type The MIME type of the file
func (c *Client) GetFileExtension(mimeType string) (*types.Text, error) {
	req := &types.GetFileExtension{
		MimeType: mimeType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// CleanFileName Removes potentially dangerous characters from the name of a file. Returns an empty string on failure. Can be called synchronously @file_name File name or path to the file
func (c *Client) CleanFileName(fileName string) (*types.Text, error) {
	req := &types.CleanFileName{
		FileName: fileName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. Can be called synchronously
func (c *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackId string, key string) (*types.LanguagePackStringValue, error) {
	req := &types.GetLanguagePackString{
		LanguagePackDatabasePath: languagePackDatabasePath,
		LocalizationTarget:       localizationTarget,
		LanguagePackId:           languagePackId,
		Key:                      key,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LanguagePackStringValue), nil
}

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously @json The JSON-serialized string
func (c *Client) GetJsonValue(json string) (*types.JsonValue, error) {
	req := &types.GetJsonValue{
		Json: json,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.JsonValue), nil
}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously @json_value The JsonValue object
func (c *Client) GetJsonString(jsonValue *types.JsonValue) (*types.Text, error) {
	req := &types.GetJsonString{
		JsonValue: jsonValue,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetThemeParametersJsonString Converts a themeParameters object to corresponding JSON-serialized string. Can be called synchronously @theme Theme parameters to convert to JSON
func (c *Client) GetThemeParametersJsonString(theme *types.ThemeParameters) (*types.Text, error) {
	req := &types.GetThemeParametersJsonString{
		Theme: theme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SetPollAnswer Changes the user answer to a poll. A poll in quiz mode can be answered only once
func (c *Client) SetPollAnswer(chatId int64, messageId int64, optionIds []int32) (*types.Ok, error) {
	req := &types.SetPollAnswer{
		ChatId:    chatId,
		MessageId: messageId,
		OptionIds: optionIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPollVoters Returns message senders voted for the specified option in a non-anonymous polls. For optimal performance, the number of returned users is chosen by TDLib
func (c *Client) GetPollVoters(chatId int64, messageId int64, optionId int32, offset int32, limit int32) (*types.MessageSenders, error) {
	req := &types.GetPollVoters{
		ChatId:    chatId,
		MessageId: messageId,
		OptionId:  optionId,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageSenders), nil
}

// StopPoll Stops a poll
func (c *Client) StopPoll(chatId int64, messageId int64, opts *types.StopPollOpts) (*types.Ok, error) {
	req := &types.StopPoll{
		ChatId:    chatId,
		MessageId: messageId,
	}
	if opts != nil {
		req.ReplyMarkup = opts.ReplyMarkup
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddChecklistTasks Adds tasks to a checklist in a message
func (c *Client) AddChecklistTasks(chatId int64, messageId int64, tasks []*types.InputChecklistTask) (*types.Ok, error) {
	req := &types.AddChecklistTasks{
		ChatId:    chatId,
		MessageId: messageId,
		Tasks:     tasks,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// MarkChecklistTasksAsDone Adds tasks of a checklist in a message as done or not done
func (c *Client) MarkChecklistTasksAsDone(chatId int64, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*types.Ok, error) {
	req := &types.MarkChecklistTasksAsDone{
		ChatId:                 chatId,
		MessageId:              messageId,
		MarkedAsDoneTaskIds:    markedAsDoneTaskIds,
		MarkedAsNotDoneTaskIds: markedAsNotDoneTaskIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// HideSuggestedAction Hides a suggested action @action Suggested action to hide
func (c *Client) HideSuggestedAction(action *types.SuggestedAction) (*types.Ok, error) {
	req := &types.HideSuggestedAction{
		Action: action,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// HideContactCloseBirthdays Hides the list of contacts that have close birthdays for 24 hours
func (c *Client) HideContactCloseBirthdays() (*types.Ok, error) {
	req := &types.HideContactCloseBirthdays{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBusinessConnection Returns information about a business connection by its identifier; for bots only @connection_id Identifier of the business connection to return
func (c *Client) GetBusinessConnection(connectionId string) (*types.BusinessConnection, error) {
	req := &types.GetBusinessConnection{
		ConnectionId: connectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessConnection), nil
}

// GetLoginUrlInfo Returns information about a button of type inlineKeyboardButtonTypeLoginUrl. The method needs to be called when the user presses the button
func (c *Client) GetLoginUrlInfo(chatId int64, messageId int64, buttonId int64) (*types.LoginUrlInfo, error) {
	req := &types.GetLoginUrlInfo{
		ChatId:    chatId,
		MessageId: messageId,
		ButtonId:  buttonId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LoginUrlInfo), nil
}

// GetLoginUrl Returns an HTTP URL which can be used to automatically authorize the user on a website after clicking an inline button of type inlineKeyboardButtonTypeLoginUrl.
func (c *Client) GetLoginUrl(chatId int64, messageId int64, buttonId int64, allowWriteAccess bool) (*types.HttpUrl, error) {
	req := &types.GetLoginUrl{
		ChatId:           chatId,
		MessageId:        messageId,
		ButtonId:         buttonId,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// ShareUsersWithBot Shares users after pressing a keyboardButtonTypeRequestUsers button with the bot
func (c *Client) ShareUsersWithBot(chatId int64, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*types.Ok, error) {
	req := &types.ShareUsersWithBot{
		ChatId:        chatId,
		MessageId:     messageId,
		ButtonId:      buttonId,
		SharedUserIds: sharedUserIds,
		OnlyCheck:     onlyCheck,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ShareChatWithBot Shares a chat after pressing a keyboardButtonTypeRequestChat button with the bot
func (c *Client) ShareChatWithBot(chatId int64, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*types.Ok, error) {
	req := &types.ShareChatWithBot{
		ChatId:       chatId,
		MessageId:    messageId,
		ButtonId:     buttonId,
		SharedChatId: sharedChatId,
		OnlyCheck:    onlyCheck,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
func (c *Client) GetInlineQueryResults(botUserId int64, chatId int64, query string, offset string, opts *types.GetInlineQueryResultsOpts) (*types.InlineQueryResults, error) {
	req := &types.GetInlineQueryResults{
		BotUserId: botUserId,
		ChatId:    chatId,
		Query:     query,
		Offset:    offset,
	}
	if opts != nil {
		req.UserLocation = opts.UserLocation
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.InlineQueryResults), nil
}

// AnswerInlineQuery Sets the result of an inline query; for bots only
func (c *Client) AnswerInlineQuery(inlineQueryId string, isPersonal bool, results []*types.InputInlineQueryResult, cacheTime int32, nextOffset string, opts *types.AnswerInlineQueryOpts) (*types.Ok, error) {
	req := &types.AnswerInlineQuery{
		InlineQueryId: inlineQueryId,
		IsPersonal:    isPersonal,
		Results:       results,
		CacheTime:     cacheTime,
		NextOffset:    nextOffset,
	}
	if opts != nil {
		req.Button = opts.Button
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SavePreparedInlineMessage Saves an inline message to be sent by the given user; for bots only
func (c *Client) SavePreparedInlineMessage(userId int64, result *types.InputInlineQueryResult, chatTypes *types.TargetChatTypes) (*types.PreparedInlineMessageId, error) {
	req := &types.SavePreparedInlineMessage{
		UserId:    userId,
		Result:    result,
		ChatTypes: chatTypes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PreparedInlineMessageId), nil
}

// GetPreparedInlineMessage Saves an inline message to be sent by the given user
func (c *Client) GetPreparedInlineMessage(botUserId int64, preparedMessageId string) (*types.PreparedInlineMessage, error) {
	req := &types.GetPreparedInlineMessage{
		BotUserId:         botUserId,
		PreparedMessageId: preparedMessageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PreparedInlineMessage), nil
}

// GetGrossingWebAppBots Returns the most grossing Web App bots
func (c *Client) GetGrossingWebAppBots(offset string, limit int32) (*types.FoundUsers, error) {
	req := &types.GetGrossingWebAppBots{
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundUsers), nil
}

// SearchWebApp Returns information about a Web App by its short name. Returns a 404 error if the Web App is not found
func (c *Client) SearchWebApp(botUserId int64, webAppShortName string) (*types.FoundWebApp, error) {
	req := &types.SearchWebApp{
		BotUserId:       botUserId,
		WebAppShortName: webAppShortName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundWebApp), nil
}

// GetWebAppPlaceholder Returns a default placeholder for Web Apps of a bot. This is an offline method. Returns a 404 error if the placeholder isn't known @bot_user_id Identifier of the target bot
func (c *Client) GetWebAppPlaceholder(botUserId int64) (*types.Outline, error) {
	req := &types.GetWebAppPlaceholder{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Outline), nil
}

// GetWebAppLinkUrl Returns an HTTPS URL of a Web App to open after a link of the type internalLinkTypeWebApp is clicked
func (c *Client) GetWebAppLinkUrl(chatId int64, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *types.WebAppOpenParameters) (*types.HttpUrl, error) {
	req := &types.GetWebAppLinkUrl{
		ChatId:           chatId,
		BotUserId:        botUserId,
		WebAppShortName:  webAppShortName,
		StartParameter:   startParameter,
		AllowWriteAccess: allowWriteAccess,
		Parameters:       parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetMainWebApp Returns information needed to open the main Web App of a bot
func (c *Client) GetMainWebApp(chatId int64, botUserId int64, startParameter string, parameters *types.WebAppOpenParameters) (*types.MainWebApp, error) {
	req := &types.GetMainWebApp{
		ChatId:         chatId,
		BotUserId:      botUserId,
		StartParameter: startParameter,
		Parameters:     parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MainWebApp), nil
}

// GetWebAppUrl Returns an HTTPS URL of a Web App to open from the side menu, a keyboardButtonTypeWebApp button, or an inlineQueryResultsButtonTypeWebApp button
func (c *Client) GetWebAppUrl(botUserId int64, url string, parameters *types.WebAppOpenParameters) (*types.HttpUrl, error) {
	req := &types.GetWebAppUrl{
		BotUserId:  botUserId,
		Url:        url,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// SendWebAppData Sends data received from a keyboardButtonTypeWebApp Web App to a bot
func (c *Client) SendWebAppData(botUserId int64, buttonText string, data string) (*types.Ok, error) {
	req := &types.SendWebAppData{
		BotUserId:  botUserId,
		ButtonText: buttonText,
		Data:       data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// OpenWebApp Informs TDLib that a Web App is being opened from the attachment menu, a botMenuButton button, an internalLinkTypeAttachmentMenuBot link, or an inlineKeyboardButtonTypeWebApp button.
func (c *Client) OpenWebApp(chatId int64, botUserId int64, url string, parameters *types.WebAppOpenParameters, opts *types.OpenWebAppOpts) (*types.WebAppInfo, error) {
	req := &types.OpenWebApp{
		ChatId:     chatId,
		BotUserId:  botUserId,
		Url:        url,
		Parameters: parameters,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.ReplyTo = opts.ReplyTo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.WebAppInfo), nil
}

// CloseWebApp Informs TDLib that a previously opened Web App was closed @web_app_launch_id Identifier of Web App launch, received from openWebApp
func (c *Client) CloseWebApp(webAppLaunchId string) (*types.Ok, error) {
	req := &types.CloseWebApp{
		WebAppLaunchId: webAppLaunchId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AnswerWebAppQuery Sets the result of interaction with a Web App and sends corresponding message on behalf of the user to the chat from which the query originated; for bots only
func (c *Client) AnswerWebAppQuery(webAppQueryId string, result *types.InputInlineQueryResult) (*types.SentWebAppMessage, error) {
	req := &types.AnswerWebAppQuery{
		WebAppQueryId: webAppQueryId,
		Result:        result,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.SentWebAppMessage), nil
}

// CheckWebAppFileDownload Checks whether a file can be downloaded and saved locally by Web App request
func (c *Client) CheckWebAppFileDownload(botUserId int64, fileName string, url string) (*types.Ok, error) {
	req := &types.CheckWebAppFileDownload{
		BotUserId: botUserId,
		FileName:  fileName,
		Url:       url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
func (c *Client) GetCallbackQueryAnswer(chatId int64, messageId int64, payload *types.CallbackQueryPayload) (*types.CallbackQueryAnswer, error) {
	req := &types.GetCallbackQueryAnswer{
		ChatId:    chatId,
		MessageId: messageId,
		Payload:   payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CallbackQueryAnswer), nil
}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
func (c *Client) AnswerCallbackQuery(callbackQueryId string, text string, showAlert bool, url string, cacheTime int32) (*types.Ok, error) {
	req := &types.AnswerCallbackQuery{
		CallbackQueryId: callbackQueryId,
		Text:            text,
		ShowAlert:       showAlert,
		Url:             url,
		CacheTime:       cacheTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AnswerShippingQuery Sets the result of a shipping query; for bots only @shipping_query_id Identifier of the shipping query @shipping_options Available shipping options @error_message An error message, empty on success
func (c *Client) AnswerShippingQuery(shippingQueryId string, shippingOptions []*types.ShippingOption, errorMessage string) (*types.Ok, error) {
	req := &types.AnswerShippingQuery{
		ShippingQueryId: shippingQueryId,
		ShippingOptions: shippingOptions,
		ErrorMessage:    errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only @pre_checkout_query_id Identifier of the pre-checkout query @error_message An error message, empty on success
func (c *Client) AnswerPreCheckoutQuery(preCheckoutQueryId string, errorMessage string) (*types.Ok, error) {
	req := &types.AnswerPreCheckoutQuery{
		PreCheckoutQueryId: preCheckoutQueryId,
		ErrorMessage:       errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGameScore Updates the game score of the specified user in the game; for bots only
func (c *Client) SetGameScore(chatId int64, messageId int64, editMessage bool, userId int64, score int32, force bool) (*types.Message, error) {
	req := &types.SetGameScore{
		ChatId:      chatId,
		MessageId:   messageId,
		EditMessage: editMessage,
		UserId:      userId,
		Score:       score,
		Force:       force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
func (c *Client) SetInlineGameScore(inlineMessageId string, editMessage bool, userId int64, score int32, force bool) (*types.Ok, error) {
	req := &types.SetInlineGameScore{
		InlineMessageId: inlineMessageId,
		EditMessage:     editMessage,
		UserId:          userId,
		Score:           score,
		Force:           force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only @chat_id The chat that contains the message with the game @message_id Identifier of the message @user_id User identifier
func (c *Client) GetGameHighScores(chatId int64, messageId int64, userId int64) (*types.GameHighScores, error) {
	req := &types.GetGameHighScores{
		ChatId:    chatId,
		MessageId: messageId,
		UserId:    userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GameHighScores), nil
}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only @inline_message_id Inline message identifier @user_id User identifier
func (c *Client) GetInlineGameHighScores(inlineMessageId string, userId int64) (*types.GameHighScores, error) {
	req := &types.GetInlineGameHighScores{
		InlineMessageId: inlineMessageId,
		UserId:          userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GameHighScores), nil
}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a replyMarkupForceReply reply markup has been used or dismissed
func (c *Client) DeleteChatReplyMarkup(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.DeleteChatReplyMarkup{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendChatAction Sends a notification about user activity in a chat
func (c *Client) SendChatAction(chatId int64, topicId *types.MessageTopic, businessConnectionId string, opts *types.SendChatActionOpts) (*types.Ok, error) {
	req := &types.SendChatAction{
		ChatId:               chatId,
		TopicId:              topicId,
		BusinessConnectionId: businessConnectionId,
	}
	if opts != nil {
		req.Action = opts.Action
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendTextMessageDraft Sends a draft for a being generated text message; for bots only
func (c *Client) SendTextMessageDraft(chatId int64, forumTopicId int32, draftId string, text *types.FormattedText) (*types.Ok, error) {
	req := &types.SendTextMessageDraft{
		ChatId:       chatId,
		ForumTopicId: forumTopicId,
		DraftId:      draftId,
		Text:         text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// OpenChat Informs TDLib that the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats) @chat_id Chat identifier
func (c *Client) OpenChat(chatId int64) (*types.Ok, error) {
	req := &types.OpenChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CloseChat Informs TDLib that the chat is closed by the user. Many useful activities depend on the chat being opened or closed @chat_id Chat identifier
func (c *Client) CloseChat(chatId int64) (*types.Ok, error) {
	req := &types.CloseChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ViewMessages Informs TDLib that messages are being viewed by the user. Sponsored messages must be marked as viewed only when the entire text of the message is shown on the screen (excluding the button).
func (c *Client) ViewMessages(chatId int64, messageIds []int64, forceRead bool, opts *types.ViewMessagesOpts) (*types.Ok, error) {
	req := &types.ViewMessages{
		ChatId:     chatId,
		MessageIds: messageIds,
		ForceRead:  forceRead,
	}
	if opts != nil {
		req.Source = opts.Source
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// OpenMessageContent Informs TDLib that the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message).
func (c *Client) OpenMessageContent(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.OpenMessageContent{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClickAnimatedEmojiMessage Informs TDLib that a message with an animated emoji was clicked by the user. Returns a big animated sticker to be played or a 404 error if usual animation needs to be played @chat_id Chat identifier of the message @message_id Identifier of the clicked message
func (c *Client) ClickAnimatedEmojiMessage(chatId int64, messageId int64) (*types.Sticker, error) {
	req := &types.ClickAnimatedEmojiMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Sticker), nil
}

// GetInternalLink Returns an HTTPS or a tg: link with the given type. Can be called before authorization @type Expected type of the link @is_http Pass true to create an HTTPS link (only available for some link types); pass false to create a tg: link
func (c *Client) GetInternalLink(typeField *types.InternalLinkType, isHttp bool) (*types.HttpUrl, error) {
	req := &types.GetInternalLink{
		TypeField: typeField,
		IsHttp:    isHttp,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetInternalLinkType Returns information about the type of internal link. Returns a 404 error if the link is not internal. Can be called before authorization @link The link
func (c *Client) GetInternalLinkType(link string) (*types.InternalLinkType, error) {
	req := &types.GetInternalLinkType{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.InternalLinkType), nil
}

// GetExternalLinkInfo Returns information about an action to be done when the current user clicks an external link. Don't use this method for links from secret chats if link preview is disabled in secret chats @link The link
func (c *Client) GetExternalLinkInfo(link string) (*types.LoginUrlInfo, error) {
	req := &types.GetExternalLinkInfo{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LoginUrlInfo), nil
}

// GetExternalLink Returns an HTTP URL which can be used to automatically authorize the current user on a website after clicking an HTTP link. Use the method getExternalLinkInfo to find whether a prior user confirmation is needed
func (c *Client) GetExternalLink(link string, allowWriteAccess bool) (*types.HttpUrl, error) {
	req := &types.GetExternalLink{
		Link:             link,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// ReadAllChatMentions Marks all mentions in a chat as read @chat_id Chat identifier
func (c *Client) ReadAllChatMentions(chatId int64) (*types.Ok, error) {
	req := &types.ReadAllChatMentions{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadAllChatReactions Marks all reactions in a chat as read @chat_id Chat identifier
func (c *Client) ReadAllChatReactions(chatId int64) (*types.Ok, error) {
	req := &types.ReadAllChatReactions{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CreatePrivateChat Returns an existing chat corresponding to a given user @user_id User identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreatePrivateChat(userId int64, force bool) (*types.Chat, error) {
	req := &types.CreatePrivateChat{
		UserId: userId,
		Force:  force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group @basic_group_id Basic group identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreateBasicGroupChat(basicGroupId int64, force bool) (*types.Chat, error) {
	req := &types.CreateBasicGroupChat{
		BasicGroupId: basicGroupId,
		Force:        force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel @supergroup_id Supergroup or channel identifier @force Pass true to create the chat without a network request. In this case all information about the chat except its type, title and photo can be incorrect
func (c *Client) CreateSupergroupChat(supergroupId int64, force bool) (*types.Chat, error) {
	req := &types.CreateSupergroupChat{
		SupergroupId: supergroupId,
		Force:        force,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat @secret_chat_id Secret chat identifier
func (c *Client) CreateSecretChat(secretChatId int32) (*types.Chat, error) {
	req := &types.CreateSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns information about the newly created chat
func (c *Client) CreateNewBasicGroupChat(userIds []int64, title string, messageAutoDeleteTime int32) (*types.CreatedBasicGroupChat, error) {
	req := &types.CreateNewBasicGroupChat{
		UserIds:               userIds,
		Title:                 title,
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CreatedBasicGroupChat), nil
}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
func (c *Client) CreateNewSupergroupChat(title string, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *types.CreateNewSupergroupChatOpts) (*types.Chat, error) {
	req := &types.CreateNewSupergroupChat{
		Title:                 title,
		IsForum:               isForum,
		IsChannel:             isChannel,
		Description:           description,
		MessageAutoDeleteTime: messageAutoDeleteTime,
		ForImport:             forImport,
	}
	if opts != nil {
		req.Location = opts.Location
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat @user_id Identifier of the target user
func (c *Client) CreateNewSecretChat(userId int64) (*types.Chat, error) {
	req := &types.CreateNewSecretChat{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires owner privileges. Deactivates the original basic group @chat_id Identifier of the chat to upgrade
func (c *Client) UpgradeBasicGroupChatToSupergroupChat(chatId int64) (*types.Chat, error) {
	req := &types.UpgradeBasicGroupChatToSupergroupChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// GetChatListsToAddChat Returns chat lists to which the chat can be added. This is an offline method @chat_id Chat identifier
func (c *Client) GetChatListsToAddChat(chatId int64) (*types.ChatLists, error) {
	req := &types.GetChatListsToAddChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatLists), nil
}

// AddChatToList Adds a chat to a chat list. A chat can't be simultaneously in Main and Archive chat lists, so it is automatically removed from another one if needed
func (c *Client) AddChatToList(chatId int64, chatList *types.ChatList) (*types.Ok, error) {
	req := &types.AddChatToList{
		ChatId:   chatId,
		ChatList: chatList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatFolder Returns information about a chat folder by its identifier @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolder(chatFolderId int32) (*types.ChatFolder, error) {
	req := &types.GetChatFolder{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolder), nil
}

// CreateChatFolder Creates new chat folder. Returns information about the created chat folder. There can be up to getOption("chat_folder_count_max") chat folders, but the limit can be increased with Telegram Premium @folder The new chat folder
func (c *Client) CreateChatFolder(folder *types.ChatFolder) (*types.ChatFolderInfo, error) {
	req := &types.CreateChatFolder{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInfo), nil
}

// EditChatFolder Edits existing chat folder. Returns information about the edited chat folder @chat_folder_id Chat folder identifier @folder The edited chat folder
func (c *Client) EditChatFolder(chatFolderId int32, folder *types.ChatFolder) (*types.ChatFolderInfo, error) {
	req := &types.EditChatFolder{
		ChatFolderId: chatFolderId,
		Folder:       folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInfo), nil
}

// DeleteChatFolder Deletes existing chat folder @chat_folder_id Chat folder identifier @leave_chat_ids Identifiers of the chats to leave. The chats must be pinned or always included in the folder
func (c *Client) DeleteChatFolder(chatFolderId int32, leaveChatIds []int64) (*types.Ok, error) {
	req := &types.DeleteChatFolder{
		ChatFolderId: chatFolderId,
		LeaveChatIds: leaveChatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatFolderChatsToLeave Returns identifiers of pinned or always included chats from a chat folder, which are suggested to be left when the chat folder is deleted @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderChatsToLeave(chatFolderId int32) (*types.Chats, error) {
	req := &types.GetChatFolderChatsToLeave{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetChatFolderChatCount Returns approximate number of chats in a being created chat folder. Main and archive chat lists must be fully preloaded for this function to work correctly @folder The new chat folder
func (c *Client) GetChatFolderChatCount(folder *types.ChatFolder) (*types.Count, error) {
	req := &types.GetChatFolderChatCount{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// ReorderChatFolders Changes the order of chat folders @chat_folder_ids Identifiers of chat folders in the new correct order @main_chat_list_position Position of the main chat list among chat folders, 0-based. Can be non-zero only for Premium users
func (c *Client) ReorderChatFolders(chatFolderIds []int32, mainChatListPosition int32) (*types.Ok, error) {
	req := &types.ReorderChatFolders{
		ChatFolderIds:        chatFolderIds,
		MainChatListPosition: mainChatListPosition,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatFolderTags Toggles whether chat folder tags are enabled @are_tags_enabled Pass true to enable folder tags; pass false to disable them
func (c *Client) ToggleChatFolderTags(areTagsEnabled bool) (*types.Ok, error) {
	req := &types.ToggleChatFolderTags{
		AreTagsEnabled: areTagsEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetRecommendedChatFolders Returns recommended chat folders for the current user
func (c *Client) GetRecommendedChatFolders() (*types.RecommendedChatFolders, error) {
	req := &types.GetRecommendedChatFolders{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RecommendedChatFolders), nil
}

// GetChatFolderDefaultIconName Returns default icon name for a folder. Can be called synchronously @folder Chat folder
func (c *Client) GetChatFolderDefaultIconName(folder *types.ChatFolder) (*types.ChatFolderIcon, error) {
	req := &types.GetChatFolderDefaultIconName{
		Folder: folder,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderIcon), nil
}

// GetChatsForChatFolderInviteLink Returns identifiers of chats from a chat folder, suitable for adding to a chat folder invite link @chat_folder_id Chat folder identifier
func (c *Client) GetChatsForChatFolderInviteLink(chatFolderId int32) (*types.Chats, error) {
	req := &types.GetChatsForChatFolderInviteLink{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// CreateChatFolderInviteLink Creates a new invite link for a chat folder. A link can be created for a chat folder if it has only pinned and included chats
func (c *Client) CreateChatFolderInviteLink(chatFolderId int32, name string, chatIds []int64) (*types.ChatFolderInviteLink, error) {
	req := &types.CreateChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		Name:         name,
		ChatIds:      chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInviteLink), nil
}

// GetChatFolderInviteLinks Returns invite links created by the current user for a shareable chat folder @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderInviteLinks(chatFolderId int32) (*types.ChatFolderInviteLinks, error) {
	req := &types.GetChatFolderInviteLinks{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInviteLinks), nil
}

// EditChatFolderInviteLink Edits an invite link for a chat folder
func (c *Client) EditChatFolderInviteLink(chatFolderId int32, inviteLink string, name string, chatIds []int64) (*types.ChatFolderInviteLink, error) {
	req := &types.EditChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		InviteLink:   inviteLink,
		Name:         name,
		ChatIds:      chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInviteLink), nil
}

// DeleteChatFolderInviteLink Deletes an invite link for a chat folder
func (c *Client) DeleteChatFolderInviteLink(chatFolderId int32, inviteLink string) (*types.Ok, error) {
	req := &types.DeleteChatFolderInviteLink{
		ChatFolderId: chatFolderId,
		InviteLink:   inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckChatFolderInviteLink Checks the validity of an invite link for a chat folder and returns information about the corresponding chat folder @invite_link Invite link to be checked
func (c *Client) CheckChatFolderInviteLink(inviteLink string) (*types.ChatFolderInviteLinkInfo, error) {
	req := &types.CheckChatFolderInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatFolderInviteLinkInfo), nil
}

// AddChatFolderByInviteLink Adds a chat folder by an invite link @invite_link Invite link for the chat folder @chat_ids Identifiers of the chats added to the chat folder. The chats are automatically joined if they aren't joined yet
func (c *Client) AddChatFolderByInviteLink(inviteLink string, chatIds []int64) (*types.Ok, error) {
	req := &types.AddChatFolderByInviteLink{
		InviteLink: inviteLink,
		ChatIds:    chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatFolderNewChats Returns new chats added to a shareable chat folder by its owner. The method must be called at most once in getOption("chat_folder_new_chats_update_period") for the given chat folder @chat_folder_id Chat folder identifier
func (c *Client) GetChatFolderNewChats(chatFolderId int32) (*types.Chats, error) {
	req := &types.GetChatFolderNewChats{
		ChatFolderId: chatFolderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// ProcessChatFolderNewChats Process new chats added to a shareable chat folder by its owner @chat_folder_id Chat folder identifier @added_chat_ids Identifiers of the new chats, which are added to the chat folder. The chats are automatically joined if they aren't joined yet
func (c *Client) ProcessChatFolderNewChats(chatFolderId int32, addedChatIds []int64) (*types.Ok, error) {
	req := &types.ProcessChatFolderNewChats{
		ChatFolderId: chatFolderId,
		AddedChatIds: addedChatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetArchiveChatListSettings Returns settings for automatic moving of chats to and from the Archive chat lists
func (c *Client) GetArchiveChatListSettings() (*types.ArchiveChatListSettings, error) {
	req := &types.GetArchiveChatListSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ArchiveChatListSettings), nil
}

// SetArchiveChatListSettings Changes settings for automatic moving of chats to and from the Archive chat lists @settings New settings
func (c *Client) SetArchiveChatListSettings(settings *types.ArchiveChatListSettings) (*types.Ok, error) {
	req := &types.SetArchiveChatListSettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires can_change_info member right
func (c *Client) SetChatTitle(chatId int64, title string) (*types.Ok, error) {
	req := &types.SetChatTitle{
		ChatId: chatId,
		Title:  title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires can_change_info member right
func (c *Client) SetChatPhoto(chatId int64, opts *types.SetChatPhotoOpts) (*types.Ok, error) {
	req := &types.SetChatPhoto{
		ChatId: chatId,
	}
	if opts != nil {
		req.Photo = opts.Photo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatAccentColor Changes accent color and background custom emoji of a channel chat. Requires can_change_info administrator right
func (c *Client) SetChatAccentColor(chatId int64, accentColorId int32, backgroundCustomEmojiId string) (*types.Ok, error) {
	req := &types.SetChatAccentColor{
		ChatId:                  chatId,
		AccentColorId:           accentColorId,
		BackgroundCustomEmojiId: backgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatProfileAccentColor Changes accent color and background custom emoji for profile of a supergroup or channel chat. Requires can_change_info administrator right
func (c *Client) SetChatProfileAccentColor(chatId int64, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*types.Ok, error) {
	req := &types.SetChatProfileAccentColor{
		ChatId:                         chatId,
		ProfileAccentColorId:           profileAccentColorId,
		ProfileBackgroundCustomEmojiId: profileBackgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatMessageAutoDeleteTime Changes the message auto-delete or self-destruct (for secret chats) time in a chat. Requires change_info administrator right in basic groups, supergroups and channels.
func (c *Client) SetChatMessageAutoDeleteTime(chatId int64, messageAutoDeleteTime int32) (*types.Ok, error) {
	req := &types.SetChatMessageAutoDeleteTime{
		ChatId:                chatId,
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatEmojiStatus Changes the emoji status of a chat. Use chatBoostLevelFeatures.can_set_emoji_status to check whether an emoji status can be set. Requires can_change_info administrator right
func (c *Client) SetChatEmojiStatus(chatId int64, opts *types.SetChatEmojiStatusOpts) (*types.Ok, error) {
	req := &types.SetChatEmojiStatus{
		ChatId: chatId,
	}
	if opts != nil {
		req.EmojiStatus = opts.EmojiStatus
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatPermissions Changes the chat members permissions. Supported only for basic groups and supergroups. Requires can_restrict_members administrator right
func (c *Client) SetChatPermissions(chatId int64, permissions *types.ChatPermissions) (*types.Ok, error) {
	req := &types.SetChatPermissions{
		ChatId:      chatId,
		Permissions: permissions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatBackground Sets the background in a specific chat. Supported only in private and secret chats with non-deleted users, and in chats with sufficient boost level and can_change_info administrator right
func (c *Client) SetChatBackground(chatId int64, darkThemeDimming int32, onlyForSelf bool, opts *types.SetChatBackgroundOpts) (*types.Ok, error) {
	req := &types.SetChatBackground{
		ChatId:           chatId,
		DarkThemeDimming: darkThemeDimming,
		OnlyForSelf:      onlyForSelf,
	}
	if opts != nil {
		req.Background = opts.Background
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteChatBackground Deletes background in a specific chat
func (c *Client) DeleteChatBackground(chatId int64, restorePrevious bool) (*types.Ok, error) {
	req := &types.DeleteChatBackground{
		ChatId:          chatId,
		RestorePrevious: restorePrevious,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGiftChatThemes Returns available to the current user gift chat themes
func (c *Client) GetGiftChatThemes(offset string, limit int32) (*types.GiftChatThemes, error) {
	req := &types.GetGiftChatThemes{
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftChatThemes), nil
}

// SetChatTheme Changes the chat theme. Supported only in private and secret chats @chat_id Chat identifier @theme New chat theme; pass null to return the default theme
func (c *Client) SetChatTheme(chatId int64, theme *types.InputChatTheme) (*types.Ok, error) {
	req := &types.SetChatTheme{
		ChatId: chatId,
		Theme:  theme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatDraftMessage Changes the draft message in a chat or a topic
func (c *Client) SetChatDraftMessage(chatId int64, opts *types.SetChatDraftMessageOpts) (*types.Ok, error) {
	req := &types.SetChatDraftMessage{
		ChatId: chatId,
	}
	if opts != nil {
		req.TopicId = opts.TopicId
		req.DraftMessage = opts.DraftMessage
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatNotificationSettings Changes the notification settings of a chat. Notification settings of a chat with the current user (Saved Messages) can't be changed
func (c *Client) SetChatNotificationSettings(chatId int64, notificationSettings *types.ChatNotificationSettings) (*types.Ok, error) {
	req := &types.SetChatNotificationSettings{
		ChatId:               chatId,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatHasProtectedContent Changes the ability of users to save, forward, or copy chat content. Supported only for basic groups, supergroups and channels. Requires owner privileges
func (c *Client) ToggleChatHasProtectedContent(chatId int64, hasProtectedContent bool) (*types.Ok, error) {
	req := &types.ToggleChatHasProtectedContent{
		ChatId:              chatId,
		HasProtectedContent: hasProtectedContent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatViewAsTopics Changes the view_as_topics setting of a forum chat or Saved Messages @chat_id Chat identifier @view_as_topics New value of view_as_topics
func (c *Client) ToggleChatViewAsTopics(chatId int64, viewAsTopics bool) (*types.Ok, error) {
	req := &types.ToggleChatViewAsTopics{
		ChatId:       chatId,
		ViewAsTopics: viewAsTopics,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatIsTranslatable Changes the translatable state of a chat @chat_id Chat identifier @is_translatable New value of is_translatable
func (c *Client) ToggleChatIsTranslatable(chatId int64, isTranslatable bool) (*types.Ok, error) {
	req := &types.ToggleChatIsTranslatable{
		ChatId:         chatId,
		IsTranslatable: isTranslatable,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat @chat_id Chat identifier @is_marked_as_unread New value of is_marked_as_unread
func (c *Client) ToggleChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) (*types.Ok, error) {
	req := &types.ToggleChatIsMarkedAsUnread{
		ChatId:           chatId,
		IsMarkedAsUnread: isMarkedAsUnread,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat @chat_id Chat identifier @default_disable_notification New value of default_disable_notification
func (c *Client) ToggleChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) (*types.Ok, error) {
	req := &types.ToggleChatDefaultDisableNotification{
		ChatId:                     chatId,
		DefaultDisableNotification: defaultDisableNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatAvailableReactions Changes reactions, available in a chat. Available for basic groups, supergroups, and channels. Requires can_change_info member right
func (c *Client) SetChatAvailableReactions(chatId int64, availableReactions *types.ChatAvailableReactions) (*types.Ok, error) {
	req := &types.SetChatAvailableReactions{
		ChatId:             chatId,
		AvailableReactions: availableReactions,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatClientData Changes application-specific data associated with a chat @chat_id Chat identifier @client_data New value of client_data
func (c *Client) SetChatClientData(chatId int64, clientData string) (*types.Ok, error) {
	req := &types.SetChatClientData{
		ChatId:     chatId,
		ClientData: clientData,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatDescription Changes information about a chat. Available for basic groups, supergroups, and channels. Requires can_change_info member right @chat_id Identifier of the chat @param_description New chat description; 0-255 characters
func (c *Client) SetChatDescription(chatId int64, description string) (*types.Ok, error) {
	req := &types.SetChatDescription{
		ChatId:      chatId,
		Description: description,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatDiscussionGroup Changes the discussion group of a channel chat; requires can_change_info administrator right in the channel if it is specified
func (c *Client) SetChatDiscussionGroup(chatId int64, discussionChatId int64) (*types.Ok, error) {
	req := &types.SetChatDiscussionGroup{
		ChatId:           chatId,
		DiscussionChatId: discussionChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatDirectMessagesGroup Changes direct messages group settings for a channel chat; requires owner privileges in the chat
func (c *Client) SetChatDirectMessagesGroup(chatId int64, isEnabled bool, paidMessageStarCount int64) (*types.Ok, error) {
	req := &types.SetChatDirectMessagesGroup{
		ChatId:               chatId,
		IsEnabled:            isEnabled,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatLocation Changes the location of a chat. Available only for some location-based supergroups, use supergroupFullInfo.can_set_location to check whether the method is allowed to use @chat_id Chat identifier @location New location for the chat; must be valid and not null
func (c *Client) SetChatLocation(chatId int64, location *types.ChatLocation) (*types.Ok, error) {
	req := &types.SetChatLocation{
		ChatId:   chatId,
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatSlowModeDelay Changes the slow mode delay of a chat. Available only for supergroups; requires can_restrict_members administrator right @chat_id Chat identifier @slow_mode_delay New slow mode delay for the chat, in seconds; must be one of 0, 5, 10, 30, 60, 300, 900, 3600
func (c *Client) SetChatSlowModeDelay(chatId int64, slowModeDelay int32) (*types.Ok, error) {
	req := &types.SetChatSlowModeDelay{
		ChatId:        chatId,
		SlowModeDelay: slowModeDelay,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// PinChatMessage Pins a message in a chat. A message can be pinned only if messageProperties.can_be_pinned
func (c *Client) PinChatMessage(chatId int64, messageId int64, disableNotification bool, onlyForSelf bool) (*types.Ok, error) {
	req := &types.PinChatMessage{
		ChatId:              chatId,
		MessageId:           messageId,
		DisableNotification: disableNotification,
		OnlyForSelf:         onlyForSelf,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// UnpinChatMessage Removes a pinned message from a chat; requires can_pin_messages member right if the chat is a basic group or supergroup, or can_edit_messages administrator right if the chat is a channel @chat_id Identifier of the chat @message_id Identifier of the removed pinned message
func (c *Client) UnpinChatMessage(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.UnpinChatMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// UnpinAllChatMessages Removes all pinned messages from a chat; requires can_pin_messages member right if the chat is a basic group or supergroup, or can_edit_messages administrator right if the chat is a channel @chat_id Identifier of the chat
func (c *Client) UnpinAllChatMessages(chatId int64) (*types.Ok, error) {
	req := &types.UnpinAllChatMessages{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// JoinChat Adds the current user as a new member to a chat. Private and secret chats can't be joined using this method. May return an error with a message "INVITE_REQUEST_SENT" if only a join request was created @chat_id Chat identifier
func (c *Client) JoinChat(chatId int64) (*types.Ok, error) {
	req := &types.JoinChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LeaveChat Removes the current user from chat members. Private and secret chats can't be left using this method @chat_id Chat identifier
func (c *Client) LeaveChat(chatId int64) (*types.Ok, error) {
	req := &types.LeaveChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddChatMember Adds a new member to a chat; requires can_invite_users member right. Members can't be added to private or secret chats. Returns information about members that weren't added
func (c *Client) AddChatMember(chatId int64, userId int64, forwardLimit int32) (*types.FailedToAddMembers, error) {
	req := &types.AddChatMember{
		ChatId:       chatId,
		UserId:       userId,
		ForwardLimit: forwardLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FailedToAddMembers), nil
}

// AddChatMembers Adds multiple new members to a chat; requires can_invite_users member right. Currently, this method is only available for supergroups and channels.
func (c *Client) AddChatMembers(chatId int64, userIds []int64) (*types.FailedToAddMembers, error) {
	req := &types.AddChatMembers{
		ChatId:  chatId,
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FailedToAddMembers), nil
}

// SetChatMemberStatus Changes the status of a chat member; requires can_invite_users member right to add a chat member, can_promote_members administrator right to change administrator rights of the member,
func (c *Client) SetChatMemberStatus(chatId int64, memberId *types.MessageSender, status *types.ChatMemberStatus) (*types.Ok, error) {
	req := &types.SetChatMemberStatus{
		ChatId:   chatId,
		MemberId: memberId,
		Status:   status,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// BanChatMember Bans a member in a chat; requires can_restrict_members administrator right. Members can't be banned in private or secret chats. In supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first
func (c *Client) BanChatMember(chatId int64, memberId *types.MessageSender, bannedUntilDate int32, revokeMessages bool) (*types.Ok, error) {
	req := &types.BanChatMember{
		ChatId:          chatId,
		MemberId:        memberId,
		BannedUntilDate: bannedUntilDate,
		RevokeMessages:  revokeMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (c *Client) CanTransferOwnership() (*types.CanTransferOwnershipResult, error) {
	req := &types.CanTransferOwnership{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CanTransferOwnershipResult), nil
}

// TransferChatOwnership Changes the owner of a chat; requires owner privileges in the chat. Use the method canTransferOwnership to check whether the ownership can be transferred from the current session. Available only for supergroups and channel chats
func (c *Client) TransferChatOwnership(chatId int64, userId int64, password string) (*types.Ok, error) {
	req := &types.TransferChatOwnership{
		ChatId:   chatId,
		UserId:   userId,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatMember Returns information about a single member of a chat @chat_id Chat identifier @member_id Member identifier
func (c *Client) GetChatMember(chatId int64, memberId *types.MessageSender) (*types.ChatMember, error) {
	req := &types.GetChatMember{
		ChatId:   chatId,
		MemberId: memberId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatMember), nil
}

// SearchChatMembers Searches for a specified query in the first name, last name and usernames of the members of a specified chat. Requires administrator rights if the chat is a channel
func (c *Client) SearchChatMembers(chatId int64, query string, limit int32, opts *types.SearchChatMembersOpts) (*types.ChatMembers, error) {
	req := &types.SearchChatMembers{
		ChatId: chatId,
		Query:  query,
		Limit:  limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatMembers), nil
}

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles @chat_id Chat identifier
func (c *Client) GetChatAdministrators(chatId int64) (*types.ChatAdministrators, error) {
	req := &types.GetChatAdministrators{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatAdministrators), nil
}

// ClearAllDraftMessages Clears message drafts in all chats @exclude_secret_chats Pass true to keep local message drafts in secret chats
func (c *Client) ClearAllDraftMessages(excludeSecretChats bool) (*types.Ok, error) {
	req := &types.ClearAllDraftMessages{
		ExcludeSecretChats: excludeSecretChats,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStakeDiceState Returns the current state of stake dice
func (c *Client) GetStakeDiceState() (*types.StakeDiceState, error) {
	req := &types.GetStakeDiceState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StakeDiceState), nil
}

// GetSavedNotificationSound Returns saved notification sound by its identifier. Returns a 404 error if there is no saved notification sound with the specified identifier @notification_sound_id Identifier of the notification sound
func (c *Client) GetSavedNotificationSound(notificationSoundId string) (*types.NotificationSounds, error) {
	req := &types.GetSavedNotificationSound{
		NotificationSoundId: notificationSoundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.NotificationSounds), nil
}

// GetSavedNotificationSounds Returns the list of saved notification sounds. If a sound isn't in the list, then default sound needs to be used
func (c *Client) GetSavedNotificationSounds() (*types.NotificationSounds, error) {
	req := &types.GetSavedNotificationSounds{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.NotificationSounds), nil
}

// AddSavedNotificationSound Adds a new notification sound to the list of saved notification sounds. The new notification sound is added to the top of the list. If it is already in the list, its position isn't changed @sound Notification sound file to add
func (c *Client) AddSavedNotificationSound(sound *types.InputFile) (*types.NotificationSound, error) {
	req := &types.AddSavedNotificationSound{
		Sound: sound,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.NotificationSound), nil
}

// RemoveSavedNotificationSound Removes a notification sound from the list of saved notification sounds @notification_sound_id Identifier of the notification sound
func (c *Client) RemoveSavedNotificationSound(notificationSoundId string) (*types.Ok, error) {
	req := &types.RemoveSavedNotificationSound{
		NotificationSoundId: notificationSoundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatNotificationSettingsExceptions Returns the list of chats with non-default notification settings for new messages
func (c *Client) GetChatNotificationSettingsExceptions(compareSound bool, opts *types.GetChatNotificationSettingsExceptionsOpts) (*types.Chats, error) {
	req := &types.GetChatNotificationSettingsExceptions{
		CompareSound: compareSound,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type @scope Types of chats for which to return the notification settings information
func (c *Client) GetScopeNotificationSettings(scope *types.NotificationSettingsScope) (*types.ScopeNotificationSettings, error) {
	req := &types.GetScopeNotificationSettings{
		Scope: scope,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ScopeNotificationSettings), nil
}

// SetScopeNotificationSettings Changes notification settings for chats of a given type @scope Types of chats for which to change the notification settings @notification_settings The new notification settings for the given scope
func (c *Client) SetScopeNotificationSettings(scope *types.NotificationSettingsScope, notificationSettings *types.ScopeNotificationSettings) (*types.Ok, error) {
	req := &types.SetScopeNotificationSettings{
		Scope:                scope,
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetReactionNotificationSettings Changes notification settings for reactions @notification_settings The new notification settings for reactions
func (c *Client) SetReactionNotificationSettings(notificationSettings *types.ReactionNotificationSettings) (*types.Ok, error) {
	req := &types.SetReactionNotificationSettings{
		NotificationSettings: notificationSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResetAllNotificationSettings Resets all chat and scope notification settings to their default values. By default, all chats are unmuted and message previews are shown
func (c *Client) ResetAllNotificationSettings() (*types.Ok, error) {
	req := &types.ResetAllNotificationSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatIsPinned Changes the pinned state of a chat. There can be up to getOption("pinned_chat_count_max")/getOption("pinned_archived_chat_count_max") pinned non-secret chats and the same number of secret chats in the main/archive chat list. The limit can be increased with Telegram Premium
func (c *Client) ToggleChatIsPinned(chatList *types.ChatList, chatId int64, isPinned bool) (*types.Ok, error) {
	req := &types.ToggleChatIsPinned{
		ChatList: chatList,
		ChatId:   chatId,
		IsPinned: isPinned,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPinnedChats Changes the order of pinned chats @chat_list Chat list in which to change the order of pinned chats @chat_ids The new list of pinned chats
func (c *Client) SetPinnedChats(chatList *types.ChatList, chatIds []int64) (*types.Ok, error) {
	req := &types.SetPinnedChats{
		ChatList: chatList,
		ChatIds:  chatIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadChatList Traverses all chats in a chat list and marks all messages in the chats as read @chat_list Chat list in which to mark all chats as read
func (c *Client) ReadChatList(chatList *types.ChatList) (*types.Ok, error) {
	req := &types.ReadChatList{
		ChatList: chatList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetCurrentWeather Returns the current weather in the given location @location The location
func (c *Client) GetCurrentWeather(location *types.Location) (*types.CurrentWeather, error) {
	req := &types.GetCurrentWeather{
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CurrentWeather), nil
}

// GetStory Returns a story
func (c *Client) GetStory(storyPosterChatId int64, storyId int32, onlyLocal bool) (*types.Story, error) {
	req := &types.GetStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		OnlyLocal:         onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Story), nil
}

// GetChatsToPostStories Returns supergroup and channel chats in which the current user has the right to post stories. The chats must be rechecked with canPostStory before actually trying to post a story there
func (c *Client) GetChatsToPostStories() (*types.Chats, error) {
	req := &types.GetChatsToPostStories{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// CanPostStory Checks whether the current user can post a story on behalf of a chat; requires can_post_stories administrator right for supergroup and channel chats
func (c *Client) CanPostStory(chatId int64) (*types.CanPostStoryResult, error) {
	req := &types.CanPostStory{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CanPostStoryResult), nil
}

// PostStory Posts a new story on behalf of a chat; requires can_post_stories administrator right for supergroup and channel chats. Returns a temporary story
func (c *Client) PostStory(chatId int64, content *types.InputStoryContent, privacySettings *types.StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *types.PostStoryOpts) (*types.Story, error) {
	req := &types.PostStory{
		ChatId:             chatId,
		Content:            content,
		PrivacySettings:    privacySettings,
		AlbumIds:           albumIds,
		ActivePeriod:       activePeriod,
		IsPostedToChatPage: isPostedToChatPage,
		ProtectContent:     protectContent,
	}
	if opts != nil {
		req.Areas = opts.Areas
		req.Caption = opts.Caption
		req.FromStoryFullId = opts.FromStoryFullId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Story), nil
}

// StartLiveStory Starts a new live story on behalf of a chat; requires can_post_stories administrator right for channel chats
func (c *Client) StartLiveStory(chatId int64, privacySettings *types.StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*types.StartLiveStoryResult, error) {
	req := &types.StartLiveStory{
		ChatId:               chatId,
		PrivacySettings:      privacySettings,
		ProtectContent:       protectContent,
		IsRtmpStream:         isRtmpStream,
		EnableMessages:       enableMessages,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StartLiveStoryResult), nil
}

// EditStory Changes content and caption of a story. Can be called only if story.can_be_edited == true
func (c *Client) EditStory(storyPosterChatId int64, storyId int32, opts *types.EditStoryOpts) (*types.Ok, error) {
	req := &types.EditStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	if opts != nil {
		req.Content = opts.Content
		req.Areas = opts.Areas
		req.Caption = opts.Caption
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditStoryCover Changes cover of a video story. Can be called only if story.can_be_edited == true and the story isn't being edited now
func (c *Client) EditStoryCover(storyPosterChatId int64, storyId int32, coverFrameTimestamp float64) (*types.Ok, error) {
	req := &types.EditStoryCover{
		StoryPosterChatId:   storyPosterChatId,
		StoryId:             storyId,
		CoverFrameTimestamp: coverFrameTimestamp,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStoryPrivacySettings Changes privacy settings of a story. The method can be called only for stories posted on behalf of the current user and if story.can_set_privacy_settings == true
func (c *Client) SetStoryPrivacySettings(storyId int32, privacySettings *types.StoryPrivacySettings) (*types.Ok, error) {
	req := &types.SetStoryPrivacySettings{
		StoryId:         storyId,
		PrivacySettings: privacySettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleStoryIsPostedToChatPage Toggles whether a story is accessible after expiration. Can be called only if story.can_toggle_is_posted_to_chat_page == true
func (c *Client) ToggleStoryIsPostedToChatPage(storyPosterChatId int64, storyId int32, isPostedToChatPage bool) (*types.Ok, error) {
	req := &types.ToggleStoryIsPostedToChatPage{
		StoryPosterChatId:  storyPosterChatId,
		StoryId:            storyId,
		IsPostedToChatPage: isPostedToChatPage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteStory Deletes a previously posted story. Can be called only if story.can_be_deleted == true
func (c *Client) DeleteStory(storyPosterChatId int64, storyId int32) (*types.Ok, error) {
	req := &types.DeleteStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStoryNotificationSettingsExceptions Returns the list of chats with non-default notification settings for stories
func (c *Client) GetStoryNotificationSettingsExceptions() (*types.Chats, error) {
	req := &types.GetStoryNotificationSettingsExceptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chats), nil
}

// LoadActiveStories Loads more active stories from a story list. The loaded stories will be sent through updates. Active stories are sorted by
func (c *Client) LoadActiveStories(storyList *types.StoryList) (*types.Ok, error) {
	req := &types.LoadActiveStories{
		StoryList: storyList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatActiveStoriesList Changes story list in which stories from the chat are shown @chat_id Identifier of the chat that posted stories @story_list New list for active stories posted by the chat
func (c *Client) SetChatActiveStoriesList(chatId int64, storyList *types.StoryList) (*types.Ok, error) {
	req := &types.SetChatActiveStoriesList{
		ChatId:    chatId,
		StoryList: storyList,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatActiveStories Returns the list of active stories posted by the given chat @chat_id Chat identifier
func (c *Client) GetChatActiveStories(chatId int64) (*types.ChatActiveStories, error) {
	req := &types.GetChatActiveStories{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatActiveStories), nil
}

// GetChatPostedToChatPageStories Returns the list of stories that posted by the given chat to its chat page. If from_story_id == 0, then pinned stories are returned first.
func (c *Client) GetChatPostedToChatPageStories(chatId int64, fromStoryId int32, limit int32) (*types.Stories, error) {
	req := &types.GetChatPostedToChatPageStories{
		ChatId:      chatId,
		FromStoryId: fromStoryId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stories), nil
}

// GetChatArchivedStories Returns the list of all stories posted by the given chat; requires can_edit_stories administrator right in the chat.
func (c *Client) GetChatArchivedStories(chatId int64, fromStoryId int32, limit int32) (*types.Stories, error) {
	req := &types.GetChatArchivedStories{
		ChatId:      chatId,
		FromStoryId: fromStoryId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stories), nil
}

// SetChatPinnedStories Changes the list of pinned stories on a chat page; requires can_edit_stories administrator right in the chat
func (c *Client) SetChatPinnedStories(chatId int64, storyIds []int32) (*types.Ok, error) {
	req := &types.SetChatPinnedStories{
		ChatId:   chatId,
		StoryIds: storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// OpenStory Informs TDLib that a story is opened and is being viewed by the user
func (c *Client) OpenStory(storyPosterChatId int64, storyId int32) (*types.Ok, error) {
	req := &types.OpenStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CloseStory Informs TDLib that a story is closed by the user
func (c *Client) CloseStory(storyPosterChatId int64, storyId int32) (*types.Ok, error) {
	req := &types.CloseStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStoryAvailableReactions Returns reactions, which can be chosen for a story @row_size Number of reaction per row, 5-25
func (c *Client) GetStoryAvailableReactions(rowSize int32) (*types.AvailableReactions, error) {
	req := &types.GetStoryAvailableReactions{
		RowSize: rowSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AvailableReactions), nil
}

// SetStoryReaction Changes chosen reaction on a story that has already been sent; not supported for live stories
func (c *Client) SetStoryReaction(storyPosterChatId int64, storyId int32, updateRecentReactions bool, opts *types.SetStoryReactionOpts) (*types.Ok, error) {
	req := &types.SetStoryReaction{
		StoryPosterChatId:     storyPosterChatId,
		StoryId:               storyId,
		UpdateRecentReactions: updateRecentReactions,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStoryInteractions Returns interactions with a story. The method can be called only for stories posted on behalf of the current user
func (c *Client) GetStoryInteractions(storyId int32, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*types.StoryInteractions, error) {
	req := &types.GetStoryInteractions{
		StoryId:            storyId,
		Query:              query,
		OnlyContacts:       onlyContacts,
		PreferForwards:     preferForwards,
		PreferWithReaction: preferWithReaction,
		Offset:             offset,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryInteractions), nil
}

// GetChatStoryInteractions Returns interactions with a story posted in a chat. Can be used only if story is posted on behalf of a chat and the user is an administrator in the chat
func (c *Client) GetChatStoryInteractions(storyPosterChatId int64, storyId int32, preferForwards bool, offset string, limit int32, opts *types.GetChatStoryInteractionsOpts) (*types.StoryInteractions, error) {
	req := &types.GetChatStoryInteractions{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		PreferForwards:    preferForwards,
		Offset:            offset,
		Limit:             limit,
	}
	if opts != nil {
		req.ReactionType = opts.ReactionType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryInteractions), nil
}

// ReportStory Reports a story to the Telegram moderators
func (c *Client) ReportStory(storyPosterChatId int64, storyId int32, optionId string, text string) (*types.ReportStoryResult, error) {
	req := &types.ReportStory{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		OptionId:          optionId,
		Text:              text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReportStoryResult), nil
}

// ActivateStoryStealthMode Activates stealth mode for stories, which hides all views of stories from the current user in the last "story_stealth_mode_past_period" seconds
func (c *Client) ActivateStoryStealthMode() (*types.Ok, error) {
	req := &types.ActivateStoryStealthMode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStoryPublicForwards Returns forwards of a story as a message to public chats and reposts by public channels. Can be used only if the story is posted on behalf of the current user or story.can_get_statistics == true.
func (c *Client) GetStoryPublicForwards(storyPosterChatId int64, storyId int32, offset string, limit int32) (*types.PublicForwards, error) {
	req := &types.GetStoryPublicForwards{
		StoryPosterChatId: storyPosterChatId,
		StoryId:           storyId,
		Offset:            offset,
		Limit:             limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PublicForwards), nil
}

// GetChatStoryAlbums Returns the list of story albums owned by the given chat @chat_id Chat identifier
func (c *Client) GetChatStoryAlbums(chatId int64) (*types.StoryAlbums, error) {
	req := &types.GetChatStoryAlbums{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbums), nil
}

// GetStoryAlbumStories Returns the list of stories added to the given story album. For optimal performance, the number of returned stories is chosen by TDLib
func (c *Client) GetStoryAlbumStories(chatId int64, storyAlbumId int32, offset int32, limit int32) (*types.Stories, error) {
	req := &types.GetStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		Offset:       offset,
		Limit:        limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stories), nil
}

// CreateStoryAlbum Creates an album of stories; requires can_edit_stories administrator right for supergroup and channel chats
func (c *Client) CreateStoryAlbum(storyPosterChatId int64, name string, storyIds []int32) (*types.StoryAlbum, error) {
	req := &types.CreateStoryAlbum{
		StoryPosterChatId: storyPosterChatId,
		Name:              name,
		StoryIds:          storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbum), nil
}

// ReorderStoryAlbums Changes order of story albums. If the albums are owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat
func (c *Client) ReorderStoryAlbums(chatId int64, storyAlbumIds []int32) (*types.Ok, error) {
	req := &types.ReorderStoryAlbums{
		ChatId:        chatId,
		StoryAlbumIds: storyAlbumIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteStoryAlbum Deletes a story album. If the album is owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat
func (c *Client) DeleteStoryAlbum(chatId int64, storyAlbumId int32) (*types.Ok, error) {
	req := &types.DeleteStoryAlbum{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStoryAlbumName Changes name of an album of stories. If the album is owned by a supergroup or a channel chat, then requires can_edit_stories administrator right in the chat. Returns the changed album
func (c *Client) SetStoryAlbumName(chatId int64, storyAlbumId int32, name string) (*types.StoryAlbum, error) {
	req := &types.SetStoryAlbumName{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbum), nil
}

// AddStoryAlbumStories Adds stories to the beginning of a previously created story album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) AddStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*types.StoryAlbum, error) {
	req := &types.AddStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbum), nil
}

// RemoveStoryAlbumStories Removes stories from an album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) RemoveStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*types.StoryAlbum, error) {
	req := &types.RemoveStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbum), nil
}

// ReorderStoryAlbumStories Changes order of stories in an album. If the album is owned by a supergroup or a channel chat, then
func (c *Client) ReorderStoryAlbumStories(chatId int64, storyAlbumId int32, storyIds []int32) (*types.StoryAlbum, error) {
	req := &types.ReorderStoryAlbumStories{
		ChatId:       chatId,
		StoryAlbumId: storyAlbumId,
		StoryIds:     storyIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryAlbum), nil
}

// GetChatBoostLevelFeatures Returns the list of features available on the specific chat boost level. This is an offline method
func (c *Client) GetChatBoostLevelFeatures(isChannel bool, level int32) (*types.ChatBoostLevelFeatures, error) {
	req := &types.GetChatBoostLevelFeatures{
		IsChannel: isChannel,
		Level:     level,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostLevelFeatures), nil
}

// GetChatBoostFeatures Returns the list of features available for different chat boost levels. This is an offline method
func (c *Client) GetChatBoostFeatures(isChannel bool) (*types.ChatBoostFeatures, error) {
	req := &types.GetChatBoostFeatures{
		IsChannel: isChannel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostFeatures), nil
}

// GetAvailableChatBoostSlots Returns the list of available chat boost slots for the current user
func (c *Client) GetAvailableChatBoostSlots() (*types.ChatBoostSlots, error) {
	req := &types.GetAvailableChatBoostSlots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostSlots), nil
}

// GetChatBoostStatus Returns the current boost status for a supergroup or a channel chat @chat_id Identifier of the chat
func (c *Client) GetChatBoostStatus(chatId int64) (*types.ChatBoostStatus, error) {
	req := &types.GetChatBoostStatus{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostStatus), nil
}

// BoostChat Boosts a chat and returns the list of available chat boost slots for the current user after the boost
func (c *Client) BoostChat(chatId int64, slotIds []int32) (*types.ChatBoostSlots, error) {
	req := &types.BoostChat{
		ChatId:  chatId,
		SlotIds: slotIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostSlots), nil
}

// GetChatBoostLink Returns an HTTPS link to boost the specified supergroup or channel chat @chat_id Identifier of the chat
func (c *Client) GetChatBoostLink(chatId int64) (*types.ChatBoostLink, error) {
	req := &types.GetChatBoostLink{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostLink), nil
}

// GetChatBoostLinkInfo Returns information about a link to boost a chat. Can be called for any internal link of the type internalLinkTypeChatBoost @url The link to boost a chat
func (c *Client) GetChatBoostLinkInfo(url string) (*types.ChatBoostLinkInfo, error) {
	req := &types.GetChatBoostLinkInfo{
		Url: url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatBoostLinkInfo), nil
}

// GetChatBoosts Returns the list of boosts applied to a chat; requires administrator rights in the chat
func (c *Client) GetChatBoosts(chatId int64, onlyGiftCodes bool, offset string, limit int32) (*types.FoundChatBoosts, error) {
	req := &types.GetChatBoosts{
		ChatId:        chatId,
		OnlyGiftCodes: onlyGiftCodes,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundChatBoosts), nil
}

// GetUserChatBoosts Returns the list of boosts applied to a chat by a given user; requires administrator rights in the chat; for bots only
func (c *Client) GetUserChatBoosts(chatId int64, userId int64) (*types.FoundChatBoosts, error) {
	req := &types.GetUserChatBoosts{
		ChatId: chatId,
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundChatBoosts), nil
}

// GetAttachmentMenuBot Returns information about a bot that can be added to attachment or side menu @bot_user_id Bot's user identifier
func (c *Client) GetAttachmentMenuBot(botUserId int64) (*types.AttachmentMenuBot, error) {
	req := &types.GetAttachmentMenuBot{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AttachmentMenuBot), nil
}

// ToggleBotIsAddedToAttachmentMenu Adds or removes a bot to attachment and side menu. Bot can be added to the menu, only if userTypeBot.can_be_added_to_attachment_menu == true
func (c *Client) ToggleBotIsAddedToAttachmentMenu(botUserId int64, isAdded bool, allowWriteAccess bool) (*types.Ok, error) {
	req := &types.ToggleBotIsAddedToAttachmentMenu{
		BotUserId:        botUserId,
		IsAdded:          isAdded,
		AllowWriteAccess: allowWriteAccess,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetThemedEmojiStatuses Returns up to 8 emoji statuses, which must be shown right after the default Premium Badge in the emoji status list for self status
func (c *Client) GetThemedEmojiStatuses() (*types.EmojiStatusCustomEmojis, error) {
	req := &types.GetThemedEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatusCustomEmojis), nil
}

// GetRecentEmojiStatuses Returns recent emoji statuses for self status
func (c *Client) GetRecentEmojiStatuses() (*types.EmojiStatuses, error) {
	req := &types.GetRecentEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatuses), nil
}

// GetUpgradedGiftEmojiStatuses Returns available upgraded gift emoji statuses for self status
func (c *Client) GetUpgradedGiftEmojiStatuses() (*types.EmojiStatuses, error) {
	req := &types.GetUpgradedGiftEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatuses), nil
}

// GetDefaultEmojiStatuses Returns default emoji statuses for self status
func (c *Client) GetDefaultEmojiStatuses() (*types.EmojiStatusCustomEmojis, error) {
	req := &types.GetDefaultEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatusCustomEmojis), nil
}

// ClearRecentEmojiStatuses Clears the list of recently used emoji statuses for self status
func (c *Client) ClearRecentEmojiStatuses() (*types.Ok, error) {
	req := &types.ClearRecentEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetThemedChatEmojiStatuses Returns up to 8 emoji statuses, which must be shown in the emoji status list for chats
func (c *Client) GetThemedChatEmojiStatuses() (*types.EmojiStatusCustomEmojis, error) {
	req := &types.GetThemedChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatusCustomEmojis), nil
}

// GetDefaultChatEmojiStatuses Returns default emoji statuses for chats
func (c *Client) GetDefaultChatEmojiStatuses() (*types.EmojiStatusCustomEmojis, error) {
	req := &types.GetDefaultChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatusCustomEmojis), nil
}

// GetDisallowedChatEmojiStatuses Returns the list of emoji statuses, which can't be used as chat emoji status, even if they are from a sticker set with is_allowed_as_chat_emoji_status == true
func (c *Client) GetDisallowedChatEmojiStatuses() (*types.EmojiStatusCustomEmojis, error) {
	req := &types.GetDisallowedChatEmojiStatuses{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiStatusCustomEmojis), nil
}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
func (c *Client) DownloadFile(fileId int32, priority int32, offset int64, limit int64, synchronous bool) (*types.File, error) {
	req := &types.DownloadFile{
		FileId:      fileId,
		Priority:    priority,
		Offset:      offset,
		Limit:       limit,
		Synchronous: synchronous,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset, in bytes @file_id Identifier of the file @offset Offset from which downloaded prefix size needs to be calculated
func (c *Client) GetFileDownloadedPrefixSize(fileId int32, offset int64) (*types.FileDownloadedPrefixSize, error) {
	req := &types.GetFileDownloadedPrefixSize{
		FileId: fileId,
		Offset: offset,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FileDownloadedPrefixSize), nil
}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing @file_id Identifier of a file to stop downloading @only_if_pending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (c *Client) CancelDownloadFile(fileId int32, onlyIfPending bool) (*types.Ok, error) {
	req := &types.CancelDownloadFile{
		FileId:        fileId,
		OnlyIfPending: onlyIfPending,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSuggestedFileName Returns suggested name for saving a file in a given directory @file_id Identifier of the file @directory Directory in which the file is expected to be saved
func (c *Client) GetSuggestedFileName(fileId int32, directory string) (*types.Text, error) {
	req := &types.GetSuggestedFileName{
		FileId:    fileId,
		Directory: directory,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// PreliminaryUploadFile Preliminarily uploads a file to the cloud before sending it in a message, which can be useful for uploading of being recorded voice and video notes.
func (c *Client) PreliminaryUploadFile(file *types.InputFile, priority int32, opts *types.PreliminaryUploadFileOpts) (*types.File, error) {
	req := &types.PreliminaryUploadFile{
		File:     file,
		Priority: priority,
	}
	if opts != nil {
		req.FileType = opts.FileType
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// CancelPreliminaryUploadFile Stops the preliminary uploading of a file. Supported only for files uploaded by using preliminaryUploadFile @file_id Identifier of the file to stop uploading
func (c *Client) CancelPreliminaryUploadFile(fileId int32) (*types.Ok, error) {
	req := &types.CancelPreliminaryUploadFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// WriteGeneratedFilePart Writes a part of a generated file. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct write to the destination file
func (c *Client) WriteGeneratedFilePart(generationId string, offset int64, data string) (*types.Ok, error) {
	req := &types.WriteGeneratedFilePart{
		GenerationId: generationId,
		Offset:       offset,
		Data:         data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetFileGenerationProgress Informs TDLib on a file generation progress
func (c *Client) SetFileGenerationProgress(generationId string, expectedSize int64, localPrefixSize int64) (*types.Ok, error) {
	req := &types.SetFileGenerationProgress{
		GenerationId:    generationId,
		ExpectedSize:    expectedSize,
		LocalPrefixSize: localPrefixSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// FinishFileGeneration Finishes the file generation
func (c *Client) FinishFileGeneration(generationId string, opts *types.FinishFileGenerationOpts) (*types.Ok, error) {
	req := &types.FinishFileGeneration{
		GenerationId: generationId,
	}
	if opts != nil {
		req.Error = opts.Error
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
func (c *Client) ReadFilePart(fileId int32, offset int64, count int64) (*types.Data, error) {
	req := &types.ReadFilePart{
		FileId: fileId,
		Offset: offset,
		Count:  count,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Data), nil
}

// DeleteFile Deletes a file from the TDLib file cache @file_id Identifier of the file to delete
func (c *Client) DeleteFile(fileId int32) (*types.Ok, error) {
	req := &types.DeleteFile{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddFileToDownloads Adds a file from a message to the list of file downloads. Download progress and completion of the download will be notified through updateFile updates.
func (c *Client) AddFileToDownloads(fileId int32, chatId int64, messageId int64, priority int32) (*types.File, error) {
	req := &types.AddFileToDownloads{
		FileId:    fileId,
		ChatId:    chatId,
		MessageId: messageId,
		Priority:  priority,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// ToggleDownloadIsPaused Changes pause state of a file in the file download list
func (c *Client) ToggleDownloadIsPaused(fileId int32, isPaused bool) (*types.Ok, error) {
	req := &types.ToggleDownloadIsPaused{
		FileId:   fileId,
		IsPaused: isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleAllDownloadsArePaused Changes pause state of all files in the file download list @are_paused Pass true to pause all downloads; pass false to unpause them
func (c *Client) ToggleAllDownloadsArePaused(arePaused bool) (*types.Ok, error) {
	req := &types.ToggleAllDownloadsArePaused{
		ArePaused: arePaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveFileFromDownloads Removes a file from the file download list @file_id Identifier of the downloaded file @delete_from_cache Pass true to delete the file from the TDLib file cache
func (c *Client) RemoveFileFromDownloads(fileId int32, deleteFromCache bool) (*types.Ok, error) {
	req := &types.RemoveFileFromDownloads{
		FileId:          fileId,
		DeleteFromCache: deleteFromCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveAllFilesFromDownloads Removes all files from the file download list
func (c *Client) RemoveAllFilesFromDownloads(onlyActive bool, onlyCompleted bool, deleteFromCache bool) (*types.Ok, error) {
	req := &types.RemoveAllFilesFromDownloads{
		OnlyActive:      onlyActive,
		OnlyCompleted:   onlyCompleted,
		DeleteFromCache: deleteFromCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchFileDownloads Searches for files in the file download list or recently downloaded files from the list
func (c *Client) SearchFileDownloads(query string, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*types.FoundFileDownloads, error) {
	req := &types.SearchFileDownloads{
		Query:         query,
		OnlyActive:    onlyActive,
		OnlyCompleted: onlyCompleted,
		Offset:        offset,
		Limit:         limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundFileDownloads), nil
}

// SetApplicationVerificationToken Informs TDLib that application or reCAPTCHA verification has been completed. Can be called before authorization
func (c *Client) SetApplicationVerificationToken(verificationId int64, token string) (*types.Ok, error) {
	req := &types.SetApplicationVerificationToken{
		VerificationId: verificationId,
		Token:          token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetMessageFileType Returns information about a file with messages exported from another application @message_file_head Beginning of the message file; up to 100 first lines
func (c *Client) GetMessageFileType(messageFileHead string) (*types.MessageFileType, error) {
	req := &types.GetMessageFileType{
		MessageFileHead: messageFileHead,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageFileType), nil
}

// GetMessageImportConfirmationText Returns a confirmation text to be shown to the user before starting message import
func (c *Client) GetMessageImportConfirmationText(chatId int64) (*types.Text, error) {
	req := &types.GetMessageImportConfirmationText{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// ImportMessages Imports messages exported from another app
func (c *Client) ImportMessages(chatId int64, messageFile *types.InputFile, attachedFiles []*types.InputFile) (*types.Ok, error) {
	req := &types.ImportMessages{
		ChatId:        chatId,
		MessageFile:   messageFile,
		AttachedFiles: attachedFiles,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReplacePrimaryChatInviteLink Replaces current primary invite link for a chat with a new primary invite link. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right @chat_id Chat identifier
func (c *Client) ReplacePrimaryChatInviteLink(chatId int64) (*types.ChatInviteLink, error) {
	req := &types.ReplacePrimaryChatInviteLink{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// CreateChatInviteLink Creates a new invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat
func (c *Client) CreateChatInviteLink(chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*types.ChatInviteLink, error) {
	req := &types.CreateChatInviteLink{
		ChatId:             chatId,
		Name:               name,
		ExpirationDate:     expirationDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// CreateChatSubscriptionInviteLink Creates a new subscription invite link for a channel chat. Requires can_invite_users right in the chat
func (c *Client) CreateChatSubscriptionInviteLink(chatId int64, name string, subscriptionPricing *types.StarSubscriptionPricing) (*types.ChatInviteLink, error) {
	req := &types.CreateChatSubscriptionInviteLink{
		ChatId:              chatId,
		Name:                name,
		SubscriptionPricing: subscriptionPricing,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// EditChatInviteLink Edits a non-primary invite link for a chat. Available for basic groups, supergroups, and channels.
func (c *Client) EditChatInviteLink(chatId int64, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*types.ChatInviteLink, error) {
	req := &types.EditChatInviteLink{
		ChatId:             chatId,
		InviteLink:         inviteLink,
		Name:               name,
		ExpirationDate:     expirationDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// EditChatSubscriptionInviteLink Edits a subscription invite link for a channel chat. Requires can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) EditChatSubscriptionInviteLink(chatId int64, inviteLink string, name string) (*types.ChatInviteLink, error) {
	req := &types.EditChatSubscriptionInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Name:       name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// GetChatInviteLink Returns information about an invite link. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
func (c *Client) GetChatInviteLink(chatId int64, inviteLink string) (*types.ChatInviteLink, error) {
	req := &types.GetChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLink), nil
}

// GetChatInviteLinkCounts Returns the list of chat administrators with number of their invite links. Requires owner privileges in the chat @chat_id Chat identifier
func (c *Client) GetChatInviteLinkCounts(chatId int64) (*types.ChatInviteLinkCounts, error) {
	req := &types.GetChatInviteLinkCounts{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLinkCounts), nil
}

// GetChatInviteLinks Returns invite links for a chat created by specified administrator. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
func (c *Client) GetChatInviteLinks(chatId int64, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*types.ChatInviteLinks, error) {
	req := &types.GetChatInviteLinks{
		ChatId:           chatId,
		CreatorUserId:    creatorUserId,
		IsRevoked:        isRevoked,
		OffsetDate:       offsetDate,
		OffsetInviteLink: offsetInviteLink,
		Limit:            limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLinks), nil
}

// GetChatInviteLinkMembers Returns chat members joined a chat via an invite link. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) GetChatInviteLinkMembers(chatId int64, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *types.GetChatInviteLinkMembersOpts) (*types.ChatInviteLinkMembers, error) {
	req := &types.GetChatInviteLinkMembers{
		ChatId:                      chatId,
		InviteLink:                  inviteLink,
		OnlyWithExpiredSubscription: onlyWithExpiredSubscription,
		Limit:                       limit,
	}
	if opts != nil {
		req.OffsetMember = opts.OffsetMember
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLinkMembers), nil
}

// RevokeChatInviteLink Revokes invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links.
func (c *Client) RevokeChatInviteLink(chatId int64, inviteLink string) (*types.ChatInviteLinks, error) {
	req := &types.RevokeChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLinks), nil
}

// DeleteRevokedChatInviteLink Deletes revoked chat invite links. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links @chat_id Chat identifier @invite_link Invite link to revoke
func (c *Client) DeleteRevokedChatInviteLink(chatId int64, inviteLink string) (*types.Ok, error) {
	req := &types.DeleteRevokedChatInviteLink{
		ChatId:     chatId,
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteAllRevokedChatInviteLinks Deletes all revoked chat invite links created by a given chat administrator. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
func (c *Client) DeleteAllRevokedChatInviteLinks(chatId int64, creatorUserId int64) (*types.Ok, error) {
	req := &types.DeleteAllRevokedChatInviteLinks{
		ChatId:        chatId,
		CreatorUserId: creatorUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat @invite_link Invite link to be checked
func (c *Client) CheckChatInviteLink(inviteLink string) (*types.ChatInviteLinkInfo, error) {
	req := &types.CheckChatInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatInviteLinkInfo), nil
}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible. May return an error with a message "INVITE_REQUEST_SENT" if only a join request was created @invite_link Invite link to use
func (c *Client) JoinChatByInviteLink(inviteLink string) (*types.Chat, error) {
	req := &types.JoinChatByInviteLink{
		InviteLink: inviteLink,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// GetChatJoinRequests Returns pending join requests in a chat
func (c *Client) GetChatJoinRequests(chatId int64, inviteLink string, query string, limit int32, opts *types.GetChatJoinRequestsOpts) (*types.ChatJoinRequests, error) {
	req := &types.GetChatJoinRequests{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Query:      query,
		Limit:      limit,
	}
	if opts != nil {
		req.OffsetRequest = opts.OffsetRequest
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatJoinRequests), nil
}

// ProcessChatJoinRequest Handles a pending join request in a chat @chat_id Chat identifier @user_id Identifier of the user that sent the request @approve Pass true to approve the request; pass false to decline it
func (c *Client) ProcessChatJoinRequest(chatId int64, userId int64, approve bool) (*types.Ok, error) {
	req := &types.ProcessChatJoinRequest{
		ChatId:  chatId,
		UserId:  userId,
		Approve: approve,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ProcessChatJoinRequests Handles all pending join requests for a given link in a chat
func (c *Client) ProcessChatJoinRequests(chatId int64, inviteLink string, approve bool) (*types.Ok, error) {
	req := &types.ProcessChatJoinRequests{
		ChatId:     chatId,
		InviteLink: inviteLink,
		Approve:    approve,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ApproveSuggestedPost Approves a suggested post in a channel direct messages chat
func (c *Client) ApproveSuggestedPost(chatId int64, messageId int64, sendDate int32) (*types.Ok, error) {
	req := &types.ApproveSuggestedPost{
		ChatId:    chatId,
		MessageId: messageId,
		SendDate:  sendDate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeclineSuggestedPost Declines a suggested post in a channel direct messages chat
func (c *Client) DeclineSuggestedPost(chatId int64, messageId int64, comment string) (*types.Ok, error) {
	req := &types.DeclineSuggestedPost{
		ChatId:    chatId,
		MessageId: messageId,
		Comment:   comment,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddOffer Sends a suggested post based on a previously sent message in a channel direct messages chat. Can be also used to suggest price or time change for an existing suggested post.
func (c *Client) AddOffer(chatId int64, messageId int64, options *types.MessageSendOptions) (*types.Message, error) {
	req := &types.AddOffer{
		ChatId:    chatId,
		MessageId: messageId,
		Options:   options,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Message), nil
}

// CreateCall Creates a new call
func (c *Client) CreateCall(userId int64, protocol *types.CallProtocol, isVideo bool) (*types.CallId, error) {
	req := &types.CreateCall{
		UserId:   userId,
		Protocol: protocol,
		IsVideo:  isVideo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CallId), nil
}

// AcceptCall Accepts an incoming call @call_id Call identifier @protocol The call protocols supported by the application
func (c *Client) AcceptCall(callId int32, protocol *types.CallProtocol) (*types.Ok, error) {
	req := &types.AcceptCall{
		CallId:   callId,
		Protocol: protocol,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendCallSignalingData Sends call signaling data @call_id Call identifier @data The data
func (c *Client) SendCallSignalingData(callId int32, data string) (*types.Ok, error) {
	req := &types.SendCallSignalingData{
		CallId: callId,
		Data:   data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DiscardCall Discards a call
func (c *Client) DiscardCall(callId int32, isDisconnected bool, inviteLink string, duration int32, isVideo bool, connectionId string) (*types.Ok, error) {
	req := &types.DiscardCall{
		CallId:         callId,
		IsDisconnected: isDisconnected,
		InviteLink:     inviteLink,
		Duration:       duration,
		IsVideo:        isVideo,
		ConnectionId:   connectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendCallRating Sends a call rating
func (c *Client) SendCallRating(callId int32, rating int32, comment string, problems []*types.CallProblem) (*types.Ok, error) {
	req := &types.SendCallRating{
		CallId:   callId,
		Rating:   rating,
		Comment:  comment,
		Problems: problems,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendCallDebugInformation Sends debug information for a call to Telegram servers @call_id Call identifier @debug_information Debug information in application-specific format
func (c *Client) SendCallDebugInformation(callId int32, debugInformation string) (*types.Ok, error) {
	req := &types.SendCallDebugInformation{
		CallId:           callId,
		DebugInformation: debugInformation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendCallLog Sends log file for a call to Telegram servers @call_id Call identifier @log_file Call log file. Only inputFileLocal and inputFileGenerated are supported
func (c *Client) SendCallLog(callId int32, logFile *types.InputFile) (*types.Ok, error) {
	req := &types.SendCallLog{
		CallId:  callId,
		LogFile: logFile,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetVideoChatAvailableParticipants Returns the list of participant identifiers, on whose behalf a video chat in the chat can be joined @chat_id Chat identifier
func (c *Client) GetVideoChatAvailableParticipants(chatId int64) (*types.MessageSenders, error) {
	req := &types.GetVideoChatAvailableParticipants{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageSenders), nil
}

// SetVideoChatDefaultParticipant Changes default participant identifier, on whose behalf a video chat in the chat will be joined
func (c *Client) SetVideoChatDefaultParticipant(chatId int64, defaultParticipantId *types.MessageSender) (*types.Ok, error) {
	req := &types.SetVideoChatDefaultParticipant{
		ChatId:               chatId,
		DefaultParticipantId: defaultParticipantId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CreateVideoChat Creates a video chat (a group call bound to a chat). Available only for basic groups, supergroups and channels; requires can_manage_video_chats administrator right
func (c *Client) CreateVideoChat(chatId int64, title string, startDate int32, isRtmpStream bool) (*types.GroupCallId, error) {
	req := &types.CreateVideoChat{
		ChatId:       chatId,
		Title:        title,
		StartDate:    startDate,
		IsRtmpStream: isRtmpStream,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallId), nil
}

// CreateGroupCall Creates a new group call that isn't bound to a chat @join_parameters Parameters to join the call; pass null to only create call link without joining the call
func (c *Client) CreateGroupCall(joinParameters *types.GroupCallJoinParameters) (*types.GroupCallInfo, error) {
	req := &types.CreateGroupCall{
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallInfo), nil
}

// GetVideoChatRtmpUrl Returns RTMP URL for streaming to the video chat of a chat; requires can_manage_video_chats administrator right @chat_id Chat identifier
func (c *Client) GetVideoChatRtmpUrl(chatId int64) (*types.RtmpUrl, error) {
	req := &types.GetVideoChatRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RtmpUrl), nil
}

// ReplaceVideoChatRtmpUrl Replaces the current RTMP URL for streaming to the video chat of a chat; requires owner privileges in the chat @chat_id Chat identifier
func (c *Client) ReplaceVideoChatRtmpUrl(chatId int64) (*types.RtmpUrl, error) {
	req := &types.ReplaceVideoChatRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RtmpUrl), nil
}

// GetLiveStoryRtmpUrl Returns RTMP URL for streaming to a live story; requires can_post_stories administrator right for channel chats @chat_id Chat identifier
func (c *Client) GetLiveStoryRtmpUrl(chatId int64) (*types.RtmpUrl, error) {
	req := &types.GetLiveStoryRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RtmpUrl), nil
}

// ReplaceLiveStoryRtmpUrl Replaces the current RTMP URL for streaming to a live story; requires owner privileges for channel chats @chat_id Chat identifier
func (c *Client) ReplaceLiveStoryRtmpUrl(chatId int64) (*types.RtmpUrl, error) {
	req := &types.ReplaceLiveStoryRtmpUrl{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.RtmpUrl), nil
}

// GetGroupCall Returns information about a group call @group_call_id Group call identifier
func (c *Client) GetGroupCall(groupCallId int32) (*types.GroupCall, error) {
	req := &types.GetGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCall), nil
}

// StartScheduledVideoChat Starts a scheduled video chat @group_call_id Group call identifier of the video chat
func (c *Client) StartScheduledVideoChat(groupCallId int32) (*types.Ok, error) {
	req := &types.StartScheduledVideoChat{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleVideoChatEnabledStartNotification Toggles whether the current user will receive a notification when the video chat starts; for scheduled video chats only
func (c *Client) ToggleVideoChatEnabledStartNotification(groupCallId int32, enabledStartNotification bool) (*types.Ok, error) {
	req := &types.ToggleVideoChatEnabledStartNotification{
		GroupCallId:              groupCallId,
		EnabledStartNotification: enabledStartNotification,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// JoinGroupCall Joins a regular group call that is not bound to a chat @input_group_call The group call to join @join_parameters Parameters to join the call
func (c *Client) JoinGroupCall(inputGroupCall *types.InputGroupCall, joinParameters *types.GroupCallJoinParameters) (*types.GroupCallInfo, error) {
	req := &types.JoinGroupCall{
		InputGroupCall: inputGroupCall,
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallInfo), nil
}

// JoinVideoChat Joins an active video chat. Returns join response payload for tgcalls
func (c *Client) JoinVideoChat(groupCallId int32, joinParameters *types.GroupCallJoinParameters, inviteHash string, opts *types.JoinVideoChatOpts) (*types.Text, error) {
	req := &types.JoinVideoChat{
		GroupCallId:    groupCallId,
		JoinParameters: joinParameters,
		InviteHash:     inviteHash,
	}
	if opts != nil {
		req.ParticipantId = opts.ParticipantId
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// JoinLiveStory Joins a group call of an active live story. Returns join response payload for tgcalls
func (c *Client) JoinLiveStory(groupCallId int32, joinParameters *types.GroupCallJoinParameters) (*types.Text, error) {
	req := &types.JoinLiveStory{
		GroupCallId:    groupCallId,
		JoinParameters: joinParameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// StartGroupCallScreenSharing Starts screen sharing in a joined group call; not supported in live stories. Returns join response payload for tgcalls
func (c *Client) StartGroupCallScreenSharing(groupCallId int32, audioSourceId int32, payload string) (*types.Text, error) {
	req := &types.StartGroupCallScreenSharing{
		GroupCallId:   groupCallId,
		AudioSourceId: audioSourceId,
		Payload:       payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// ToggleGroupCallScreenSharingIsPaused Pauses or unpauses screen sharing in a joined group call; not supported in live stories @group_call_id Group call identifier @is_paused Pass true to pause screen sharing; pass false to unpause it
func (c *Client) ToggleGroupCallScreenSharingIsPaused(groupCallId int32, isPaused bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallScreenSharingIsPaused{
		GroupCallId: groupCallId,
		IsPaused:    isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EndGroupCallScreenSharing Ends screen sharing in a joined group call; not supported in live stories @group_call_id Group call identifier
func (c *Client) EndGroupCallScreenSharing(groupCallId int32) (*types.Ok, error) {
	req := &types.EndGroupCallScreenSharing{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetVideoChatTitle Sets title of a video chat; requires groupCall.can_be_managed right @group_call_id Group call identifier @title New group call title; 1-64 characters
func (c *Client) SetVideoChatTitle(groupCallId int32, title string) (*types.Ok, error) {
	req := &types.SetVideoChatTitle{
		GroupCallId: groupCallId,
		Title:       title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleVideoChatMuteNewParticipants Toggles whether new participants of a video chat can be unmuted only by administrators of the video chat. Requires groupCall.can_toggle_mute_new_participants right
func (c *Client) ToggleVideoChatMuteNewParticipants(groupCallId int32, muteNewParticipants bool) (*types.Ok, error) {
	req := &types.ToggleVideoChatMuteNewParticipants{
		GroupCallId:         groupCallId,
		MuteNewParticipants: muteNewParticipants,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGroupCallAreMessagesAllowed Toggles whether participants of a group call can send messages there. Requires groupCall.can_toggle_are_messages_allowed right
func (c *Client) ToggleGroupCallAreMessagesAllowed(groupCallId int32, areMessagesAllowed bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallAreMessagesAllowed{
		GroupCallId:        groupCallId,
		AreMessagesAllowed: areMessagesAllowed,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLiveStoryStreamer Returns information about the user or the chat that streams to a live story; for live stories that aren't an RTMP stream only @group_call_id Group call identifier
func (c *Client) GetLiveStoryStreamer(groupCallId int32) (*types.GroupCallParticipant, error) {
	req := &types.GetLiveStoryStreamer{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallParticipant), nil
}

// GetLiveStoryAvailableMessageSenders Returns the list of message sender identifiers, on whose behalf messages can be sent to a live story @group_call_id Group call identifier
func (c *Client) GetLiveStoryAvailableMessageSenders(groupCallId int32) (*types.ChatMessageSenders, error) {
	req := &types.GetLiveStoryAvailableMessageSenders{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatMessageSenders), nil
}

// SetLiveStoryMessageSender Selects a message sender to send messages in a live story call
func (c *Client) SetLiveStoryMessageSender(groupCallId int32, messageSenderId *types.MessageSender) (*types.Ok, error) {
	req := &types.SetLiveStoryMessageSender{
		GroupCallId:     groupCallId,
		MessageSenderId: messageSenderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendGroupCallMessage Sends a message to other participants of a group call. Requires groupCall.can_send_messages right
func (c *Client) SendGroupCallMessage(groupCallId int32, text *types.FormattedText, paidMessageStarCount int64) (*types.Ok, error) {
	req := &types.SendGroupCallMessage{
		GroupCallId:          groupCallId,
		Text:                 text,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddPendingLiveStoryReaction Adds pending paid reaction in a live story group call. Can't be used in live stories posted by the current user.
func (c *Client) AddPendingLiveStoryReaction(groupCallId int32, starCount int64) (*types.Ok, error) {
	req := &types.AddPendingLiveStoryReaction{
		GroupCallId: groupCallId,
		StarCount:   starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CommitPendingLiveStoryReactions Applies all pending paid reactions in a live story group call @group_call_id Group call identifier
func (c *Client) CommitPendingLiveStoryReactions(groupCallId int32) (*types.Ok, error) {
	req := &types.CommitPendingLiveStoryReactions{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemovePendingLiveStoryReactions Removes all pending paid reactions in a live story group call @group_call_id Group call identifier
func (c *Client) RemovePendingLiveStoryReactions(groupCallId int32) (*types.Ok, error) {
	req := &types.RemovePendingLiveStoryReactions{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteGroupCallMessages Deletes messages in a group call; for live story calls only. Requires groupCallMessage.can_be_deleted right
func (c *Client) DeleteGroupCallMessages(groupCallId int32, messageIds []int32, reportSpam bool) (*types.Ok, error) {
	req := &types.DeleteGroupCallMessages{
		GroupCallId: groupCallId,
		MessageIds:  messageIds,
		ReportSpam:  reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteGroupCallMessagesBySender Deletes all messages sent by the specified message sender in a group call; for live story calls only. Requires groupCall.can_delete_messages right
func (c *Client) DeleteGroupCallMessagesBySender(groupCallId int32, senderId *types.MessageSender, reportSpam bool) (*types.Ok, error) {
	req := &types.DeleteGroupCallMessagesBySender{
		GroupCallId: groupCallId,
		SenderId:    senderId,
		ReportSpam:  reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLiveStoryTopDonors Returns the list of top live story donors @group_call_id Group call identifier of the live story
func (c *Client) GetLiveStoryTopDonors(groupCallId int32) (*types.LiveStoryDonors, error) {
	req := &types.GetLiveStoryTopDonors{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LiveStoryDonors), nil
}

// InviteGroupCallParticipant Invites a user to an active group call; for group calls not bound to a chat only. Sends a service message of the type messageGroupCall.
func (c *Client) InviteGroupCallParticipant(groupCallId int32, userId int64, isVideo bool) (*types.InviteGroupCallParticipantResult, error) {
	req := &types.InviteGroupCallParticipant{
		GroupCallId: groupCallId,
		UserId:      userId,
		IsVideo:     isVideo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.InviteGroupCallParticipantResult), nil
}

// DeclineGroupCallInvitation Declines an invitation to an active group call via messageGroupCall. Can be called both by the sender and the receiver of the invitation
func (c *Client) DeclineGroupCallInvitation(chatId int64, messageId int64) (*types.Ok, error) {
	req := &types.DeclineGroupCallInvitation{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// BanGroupCallParticipants Bans users from a group call not bound to a chat; requires groupCall.is_owned. Only the owner of the group call can invite the banned users back
func (c *Client) BanGroupCallParticipants(groupCallId int32, userIds []string) (*types.Ok, error) {
	req := &types.BanGroupCallParticipants{
		GroupCallId: groupCallId,
		UserIds:     userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// InviteVideoChatParticipants Invites users to an active video chat. Sends a service message of the type messageInviteVideoChatParticipants to the chat bound to the group call
func (c *Client) InviteVideoChatParticipants(groupCallId int32, userIds []int64) (*types.Ok, error) {
	req := &types.InviteVideoChatParticipants{
		GroupCallId: groupCallId,
		UserIds:     userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetVideoChatInviteLink Returns invite link to a video chat in a public chat
func (c *Client) GetVideoChatInviteLink(groupCallId int32, canSelfUnmute bool) (*types.HttpUrl, error) {
	req := &types.GetVideoChatInviteLink{
		GroupCallId:   groupCallId,
		CanSelfUnmute: canSelfUnmute,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// RevokeGroupCallInviteLink Revokes invite link for a group call. Requires groupCall.can_be_managed right for video chats or groupCall.is_owned otherwise @group_call_id Group call identifier
func (c *Client) RevokeGroupCallInviteLink(groupCallId int32) (*types.Ok, error) {
	req := &types.RevokeGroupCallInviteLink{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// StartGroupCallRecording Starts recording of an active group call; for video chats only. Requires groupCall.can_be_managed right
func (c *Client) StartGroupCallRecording(groupCallId int32, title string, recordVideo bool, usePortraitOrientation bool) (*types.Ok, error) {
	req := &types.StartGroupCallRecording{
		GroupCallId:            groupCallId,
		Title:                  title,
		RecordVideo:            recordVideo,
		UsePortraitOrientation: usePortraitOrientation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EndGroupCallRecording Ends recording of an active group call; for video chats only. Requires groupCall.can_be_managed right @group_call_id Group call identifier
func (c *Client) EndGroupCallRecording(groupCallId int32) (*types.Ok, error) {
	req := &types.EndGroupCallRecording{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGroupCallIsMyVideoPaused Toggles whether current user's video is paused @group_call_id Group call identifier @is_my_video_paused Pass true if the current user's video is paused
func (c *Client) ToggleGroupCallIsMyVideoPaused(groupCallId int32, isMyVideoPaused bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallIsMyVideoPaused{
		GroupCallId:     groupCallId,
		IsMyVideoPaused: isMyVideoPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGroupCallIsMyVideoEnabled Toggles whether current user's video is enabled @group_call_id Group call identifier @is_my_video_enabled Pass true if the current user's video is enabled
func (c *Client) ToggleGroupCallIsMyVideoEnabled(groupCallId int32, isMyVideoEnabled bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallIsMyVideoEnabled{
		GroupCallId:      groupCallId,
		IsMyVideoEnabled: isMyVideoEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGroupCallPaidMessageStarCount Changes the minimum number of Telegram Stars that must be paid by general participant for each sent message to a live story call. Requires groupCall.can_be_managed right
func (c *Client) SetGroupCallPaidMessageStarCount(groupCallId int32, paidMessageStarCount int64) (*types.Ok, error) {
	req := &types.SetGroupCallPaidMessageStarCount{
		GroupCallId:          groupCallId,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGroupCallParticipantIsSpeaking Informs TDLib that speaking state of a participant of an active group call has changed. Returns identifier of the participant if it is found
func (c *Client) SetGroupCallParticipantIsSpeaking(groupCallId int32, audioSource int32, isSpeaking bool) (*types.MessageSender, error) {
	req := &types.SetGroupCallParticipantIsSpeaking{
		GroupCallId: groupCallId,
		AudioSource: audioSource,
		IsSpeaking:  isSpeaking,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageSender), nil
}

// ToggleGroupCallParticipantIsMuted Toggles whether a participant of an active group call is muted, unmuted, or allowed to unmute themselves; not supported for live stories
func (c *Client) ToggleGroupCallParticipantIsMuted(groupCallId int32, participantId *types.MessageSender, isMuted bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallParticipantIsMuted{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		IsMuted:       isMuted,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGroupCallParticipantVolumeLevel Changes volume level of a participant of an active group call; not supported for live stories. If the current user can manage the group call or is the owner of the group call,
func (c *Client) SetGroupCallParticipantVolumeLevel(groupCallId int32, participantId *types.MessageSender, volumeLevel int32) (*types.Ok, error) {
	req := &types.SetGroupCallParticipantVolumeLevel{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		VolumeLevel:   volumeLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGroupCallParticipantIsHandRaised Toggles whether a group call participant hand is rased; for video chats only
func (c *Client) ToggleGroupCallParticipantIsHandRaised(groupCallId int32, participantId *types.MessageSender, isHandRaised bool) (*types.Ok, error) {
	req := &types.ToggleGroupCallParticipantIsHandRaised{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		IsHandRaised:  isHandRaised,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGroupCallParticipants Returns information about participants of a non-joined group call that is not bound to a chat
func (c *Client) GetGroupCallParticipants(inputGroupCall *types.InputGroupCall, limit int32) (*types.GroupCallParticipants, error) {
	req := &types.GetGroupCallParticipants{
		InputGroupCall: inputGroupCall,
		Limit:          limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallParticipants), nil
}

// LoadGroupCallParticipants Loads more participants of a group call; not supported in live stories. The loaded participants will be received through updates.
func (c *Client) LoadGroupCallParticipants(groupCallId int32, limit int32) (*types.Ok, error) {
	req := &types.LoadGroupCallParticipants{
		GroupCallId: groupCallId,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LeaveGroupCall Leaves a group call @group_call_id Group call identifier
func (c *Client) LeaveGroupCall(groupCallId int32) (*types.Ok, error) {
	req := &types.LeaveGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EndGroupCall Ends a group call. Requires groupCall.can_be_managed right for video chats and live stories or groupCall.is_owned otherwise @group_call_id Group call identifier
func (c *Client) EndGroupCall(groupCallId int32) (*types.Ok, error) {
	req := &types.EndGroupCall{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGroupCallStreams Returns information about available streams in a video chat or a live story @group_call_id Group call identifier
func (c *Client) GetGroupCallStreams(groupCallId int32) (*types.GroupCallStreams, error) {
	req := &types.GetGroupCallStreams{
		GroupCallId: groupCallId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GroupCallStreams), nil
}

// GetGroupCallStreamSegment Returns a file with a segment of a video chat or live story in a modified OGG format for audio or MPEG-4 format for video
func (c *Client) GetGroupCallStreamSegment(groupCallId int32, timeOffset int64, scale int32, channelId int32, opts *types.GetGroupCallStreamSegmentOpts) (*types.Data, error) {
	req := &types.GetGroupCallStreamSegment{
		GroupCallId: groupCallId,
		TimeOffset:  timeOffset,
		Scale:       scale,
		ChannelId:   channelId,
	}
	if opts != nil {
		req.VideoQuality = opts.VideoQuality
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Data), nil
}

// EncryptGroupCallData Encrypts group call data before sending them over network using tgcalls
func (c *Client) EncryptGroupCallData(groupCallId int32, dataChannel *types.GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*types.Data, error) {
	req := &types.EncryptGroupCallData{
		GroupCallId:           groupCallId,
		DataChannel:           dataChannel,
		Data:                  data,
		UnencryptedPrefixSize: unencryptedPrefixSize,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Data), nil
}

// DecryptGroupCallData Decrypts group call data received by tgcalls
func (c *Client) DecryptGroupCallData(groupCallId int32, participantId *types.MessageSender, data string, opts *types.DecryptGroupCallDataOpts) (*types.Data, error) {
	req := &types.DecryptGroupCallData{
		GroupCallId:   groupCallId,
		ParticipantId: participantId,
		Data:          data,
	}
	if opts != nil {
		req.DataChannel = opts.DataChannel
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Data), nil
}

// SetMessageSenderBlockList Changes the block list of a message sender. Currently, only users and supergroup chats can be blocked
func (c *Client) SetMessageSenderBlockList(senderId *types.MessageSender, opts *types.SetMessageSenderBlockListOpts) (*types.Ok, error) {
	req := &types.SetMessageSenderBlockList{
		SenderId: senderId,
	}
	if opts != nil {
		req.BlockList = opts.BlockList
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// BlockMessageSenderFromReplies Blocks an original sender of a message in the Replies chat
func (c *Client) BlockMessageSenderFromReplies(messageId int64, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*types.Ok, error) {
	req := &types.BlockMessageSenderFromReplies{
		MessageId:         messageId,
		DeleteMessage:     deleteMessage,
		DeleteAllMessages: deleteAllMessages,
		ReportSpam:        reportSpam,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
func (c *Client) GetBlockedMessageSenders(blockList *types.BlockList, offset int32, limit int32) (*types.MessageSenders, error) {
	req := &types.GetBlockedMessageSenders{
		BlockList: blockList,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageSenders), nil
}

// AddContact Adds a user to the contact list or edits an existing contact by their user identifier
func (c *Client) AddContact(userId int64, contact *types.ImportedContact, sharePhoneNumber bool) (*types.Ok, error) {
	req := &types.AddContact{
		UserId:           userId,
		Contact:          contact,
		SharePhoneNumber: sharePhoneNumber,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
func (c *Client) ImportContacts(contacts []*types.ImportedContact) (*types.ImportedContacts, error) {
	req := &types.ImportContacts{
		Contacts: contacts,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ImportedContacts), nil
}

// GetContacts Returns all contacts of the user
func (c *Client) GetContacts() (*types.Users, error) {
	req := &types.GetContacts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
func (c *Client) SearchContacts(query string, limit int32) (*types.Users, error) {
	req := &types.SearchContacts{
		Query: query,
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// RemoveContacts Removes users from the contact list @user_ids Identifiers of users to be deleted
func (c *Client) RemoveContacts(userIds []int64) (*types.Ok, error) {
	req := &types.RemoveContacts{
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetImportedContactCount Returns the total number of imported contacts
func (c *Client) GetImportedContactCount() (*types.Count, error) {
	req := &types.GetImportedContactCount{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Count), nil
}

// ChangeImportedContacts Changes imported contacts using the list of contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts.
func (c *Client) ChangeImportedContacts(contacts []*types.ImportedContact) (*types.ImportedContacts, error) {
	req := &types.ChangeImportedContacts{
		Contacts: contacts,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ImportedContacts), nil
}

// ClearImportedContacts Clears all imported contacts, contact list remains unchanged
func (c *Client) ClearImportedContacts() (*types.Ok, error) {
	req := &types.ClearImportedContacts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetCloseFriends Changes the list of close friends of the current user @user_ids User identifiers of close friends; the users must be contacts of the current user
func (c *Client) SetCloseFriends(userIds []int64) (*types.Ok, error) {
	req := &types.SetCloseFriends{
		UserIds: userIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetCloseFriends Returns all close friends of the current user
func (c *Client) GetCloseFriends() (*types.Users, error) {
	req := &types.GetCloseFriends{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// SetUserPersonalProfilePhoto Changes a personal profile photo of a contact user @user_id User identifier @photo Profile photo to set; pass null to delete the photo; inputChatPhotoPrevious isn't supported in this function
func (c *Client) SetUserPersonalProfilePhoto(userId int64, photo *types.InputChatPhoto) (*types.Ok, error) {
	req := &types.SetUserPersonalProfilePhoto{
		UserId: userId,
		Photo:  photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetUserNote Changes a note of a contact user
func (c *Client) SetUserNote(userId int64, note *types.FormattedText) (*types.Ok, error) {
	req := &types.SetUserNote{
		UserId: userId,
		Note:   note,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SuggestUserProfilePhoto Suggests a profile photo to another regular user with common messages and allowing non-paid messages
func (c *Client) SuggestUserProfilePhoto(userId int64, photo *types.InputChatPhoto) (*types.Ok, error) {
	req := &types.SuggestUserProfilePhoto{
		UserId: userId,
		Photo:  photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SuggestUserBirthdate Suggests a birthdate to another regular user with common messages and allowing non-paid messages
func (c *Client) SuggestUserBirthdate(userId int64, birthdate *types.Birthdate) (*types.Ok, error) {
	req := &types.SuggestUserBirthdate{
		UserId:    userId,
		Birthdate: birthdate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleBotCanManageEmojiStatus Toggles whether the bot can manage emoji status of the current user @bot_user_id User identifier of the bot @can_manage_emoji_status Pass true if the bot is allowed to change emoji status of the user; pass false otherwise
func (c *Client) ToggleBotCanManageEmojiStatus(botUserId int64, canManageEmojiStatus bool) (*types.Ok, error) {
	req := &types.ToggleBotCanManageEmojiStatus{
		BotUserId:            botUserId,
		CanManageEmojiStatus: canManageEmojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetUserEmojiStatus Changes the emoji status of a user; for bots only @user_id Identifier of the user @emoji_status New emoji status; pass null to switch to the default badge
func (c *Client) SetUserEmojiStatus(userId int64, emojiStatus *types.EmojiStatus) (*types.Ok, error) {
	req := &types.SetUserEmojiStatus{
		UserId:      userId,
		EmojiStatus: emojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchUserByPhoneNumber Searches a user by their phone number. Returns a 404 error if the user can't be found
func (c *Client) SearchUserByPhoneNumber(phoneNumber string, onlyLocal bool) (*types.User, error) {
	req := &types.SearchUserByPhoneNumber{
		PhoneNumber: phoneNumber,
		OnlyLocal:   onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// SharePhoneNumber Shares the phone number of the current user with a mutual contact. Supposed to be called when the user clicks on chatActionBarSharePhoneNumber
func (c *Client) SharePhoneNumber(userId int64) (*types.Ok, error) {
	req := &types.SharePhoneNumber{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetUserProfilePhotos Returns the profile photos of a user. Personal and public photo aren't returned
func (c *Client) GetUserProfilePhotos(userId int64, offset int32, limit int32) (*types.ChatPhotos, error) {
	req := &types.GetUserProfilePhotos{
		UserId: userId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatPhotos), nil
}

// GetUserProfileAudios Returns the list of profile audio files of a user
func (c *Client) GetUserProfileAudios(userId int64, offset int32, limit int32) (*types.Audios, error) {
	req := &types.GetUserProfileAudios{
		UserId: userId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Audios), nil
}

// IsProfileAudio Checks whether a file is in the profile audio files of the current user. Returns a 404 error if it isn't @file_id Identifier of the audio file to check
func (c *Client) IsProfileAudio(fileId int32) (*types.Ok, error) {
	req := &types.IsProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddProfileAudio Adds an audio file to the beginning of the profile audio files of the current user
func (c *Client) AddProfileAudio(fileId int32) (*types.Ok, error) {
	req := &types.AddProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetProfileAudioPosition Changes position of an audio file in the profile audio files of the current user
func (c *Client) SetProfileAudioPosition(fileId int32, afterFileId int32) (*types.Ok, error) {
	req := &types.SetProfileAudioPosition{
		FileId:      fileId,
		AfterFileId: afterFileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveProfileAudio Removes an audio file from the profile audio files of the current user @file_id Identifier of the audio file to be removed
func (c *Client) RemoveProfileAudio(fileId int32) (*types.Ok, error) {
	req := &types.RemoveProfileAudio{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStickerOutline Returns outline of a sticker. This is an offline method. Returns a 404 error if the outline isn't known
func (c *Client) GetStickerOutline(stickerFileId int32, forAnimatedEmoji bool, forClickedAnimatedEmojiMessage bool) (*types.Outline, error) {
	req := &types.GetStickerOutline{
		StickerFileId:                  stickerFileId,
		ForAnimatedEmoji:               forAnimatedEmoji,
		ForClickedAnimatedEmojiMessage: forClickedAnimatedEmojiMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Outline), nil
}

// GetStickerOutlineSvgPath Returns outline of a sticker as an SVG path. This is an offline method. Returns an empty string if the outline isn't known
func (c *Client) GetStickerOutlineSvgPath(stickerFileId int32, forAnimatedEmoji bool, forClickedAnimatedEmojiMessage bool) (*types.Text, error) {
	req := &types.GetStickerOutlineSvgPath{
		StickerFileId:                  stickerFileId,
		ForAnimatedEmoji:               forAnimatedEmoji,
		ForClickedAnimatedEmojiMessage: forClickedAnimatedEmojiMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetStickers Returns stickers from the installed sticker sets that correspond to any of the given emoji or can be found by sticker-specific keywords. If the query is non-empty, then favorite, recently used or trending stickers may also be returned
func (c *Client) GetStickers(stickerType *types.StickerType, query string, limit int32, chatId int64) (*types.Stickers, error) {
	req := &types.GetStickers{
		StickerType: stickerType,
		Query:       query,
		Limit:       limit,
		ChatId:      chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetAllStickerEmojis Returns unique emoji that correspond to stickers to be found by the getStickers(sticker_type, query, 1000000, chat_id)
func (c *Client) GetAllStickerEmojis(stickerType *types.StickerType, query string, chatId int64, returnOnlyMainEmoji bool) (*types.Emojis, error) {
	req := &types.GetAllStickerEmojis{
		StickerType:         stickerType,
		Query:               query,
		ChatId:              chatId,
		ReturnOnlyMainEmoji: returnOnlyMainEmoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Emojis), nil
}

// SearchStickers Searches for stickers from public sticker sets that correspond to any of the given emoji
func (c *Client) SearchStickers(stickerType *types.StickerType, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*types.Stickers, error) {
	req := &types.SearchStickers{
		StickerType:        stickerType,
		Emojis:             emojis,
		Query:              query,
		InputLanguageCodes: inputLanguageCodes,
		Offset:             offset,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetGreetingStickers Returns greeting stickers from regular sticker sets that can be used for the start page of other users
func (c *Client) GetGreetingStickers() (*types.Stickers, error) {
	req := &types.GetGreetingStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetPremiumStickers Returns premium stickers from regular sticker sets @limit The maximum number of stickers to be returned; 0-100
func (c *Client) GetPremiumStickers(limit int32) (*types.Stickers, error) {
	req := &types.GetPremiumStickers{
		Limit: limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetInstalledStickerSets Returns a list of installed sticker sets @sticker_type Type of the sticker sets to return
func (c *Client) GetInstalledStickerSets(stickerType *types.StickerType) (*types.StickerSets, error) {
	req := &types.GetInstalledStickerSets{
		StickerType: stickerType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// GetArchivedStickerSets Returns a list of archived sticker sets
func (c *Client) GetArchivedStickerSets(stickerType *types.StickerType, offsetStickerSetId string, limit int32) (*types.StickerSets, error) {
	req := &types.GetArchivedStickerSets{
		StickerType:        stickerType,
		OffsetStickerSetId: offsetStickerSetId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// GetTrendingStickerSets Returns a list of trending sticker sets. For optimal performance, the number of returned sticker sets is chosen by TDLib
func (c *Client) GetTrendingStickerSets(stickerType *types.StickerType, offset int32, limit int32) (*types.TrendingStickerSets, error) {
	req := &types.GetTrendingStickerSets{
		StickerType: stickerType,
		Offset:      offset,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TrendingStickerSets), nil
}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file, including regular, mask, and emoji sticker sets. Currently, only animations, photos, and videos can have attached sticker sets @file_id File identifier
func (c *Client) GetAttachedStickerSets(fileId int32) (*types.StickerSets, error) {
	req := &types.GetAttachedStickerSets{
		FileId: fileId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// GetStickerSet Returns information about a sticker set by its identifier @set_id Identifier of the sticker set
func (c *Client) GetStickerSet(setId string) (*types.StickerSet, error) {
	req := &types.GetStickerSet{
		SetId: setId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSet), nil
}

// GetStickerSetName Returns name of a sticker set by its identifier @set_id Identifier of the sticker set
func (c *Client) GetStickerSetName(setId string) (*types.Text, error) {
	req := &types.GetStickerSetName{
		SetId: setId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SearchStickerSet Searches for a sticker set by its name @name Name of the sticker set @ignore_cache Pass true to ignore local cache of sticker sets and always send a network request
func (c *Client) SearchStickerSet(name string, ignoreCache bool) (*types.StickerSet, error) {
	req := &types.SearchStickerSet{
		Name:        name,
		IgnoreCache: ignoreCache,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSet), nil
}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name @sticker_type Type of the sticker sets to search for @query Query to search for @limit The maximum number of sticker sets to return
func (c *Client) SearchInstalledStickerSets(stickerType *types.StickerType, query string, limit int32) (*types.StickerSets, error) {
	req := &types.SearchInstalledStickerSets{
		StickerType: stickerType,
		Query:       query,
		Limit:       limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// SearchStickerSets Searches for sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
func (c *Client) SearchStickerSets(stickerType *types.StickerType, query string) (*types.StickerSets, error) {
	req := &types.SearchStickerSets{
		StickerType: stickerType,
		Query:       query,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set @set_id Identifier of the sticker set @is_installed The new value of is_installed @is_archived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (c *Client) ChangeStickerSet(setId string, isInstalled bool, isArchived bool) (*types.Ok, error) {
	req := &types.ChangeStickerSet{
		SetId:       setId,
		IsInstalled: isInstalled,
		IsArchived:  isArchived,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user @sticker_set_ids Identifiers of viewed trending sticker sets
func (c *Client) ViewTrendingStickerSets(stickerSetIds []string) (*types.Ok, error) {
	req := &types.ViewTrendingStickerSets{
		StickerSetIds: stickerSetIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReorderInstalledStickerSets Changes the order of installed sticker sets @sticker_type Type of the sticker sets to reorder @sticker_set_ids Identifiers of installed sticker sets in the new correct order
func (c *Client) ReorderInstalledStickerSets(stickerType *types.StickerType, stickerSetIds []string) (*types.Ok, error) {
	req := &types.ReorderInstalledStickerSets{
		StickerType:   stickerType,
		StickerSetIds: stickerSetIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetRecentStickers Returns a list of recently used stickers @is_attached Pass true to return stickers and masks that were recently attached to photos or video files; pass false to return recently sent stickers
func (c *Client) GetRecentStickers(isAttached bool) (*types.Stickers, error) {
	req := &types.GetRecentStickers{
		IsAttached: isAttached,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// AddRecentSticker Manually adds a new sticker to the list of recently used stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first.
func (c *Client) AddRecentSticker(isAttached bool, sticker *types.InputFile) (*types.Stickers, error) {
	req := &types.AddRecentSticker{
		IsAttached: isAttached,
		Sticker:    sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers @is_attached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers @sticker Sticker file to delete
func (c *Client) RemoveRecentSticker(isAttached bool, sticker *types.InputFile) (*types.Ok, error) {
	req := &types.RemoveRecentSticker{
		IsAttached: isAttached,
		Sticker:    sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClearRecentStickers Clears the list of recently used stickers @is_attached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (c *Client) ClearRecentStickers(isAttached bool) (*types.Ok, error) {
	req := &types.ClearRecentStickers{
		IsAttached: isAttached,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetFavoriteStickers Returns favorite stickers
func (c *Client) GetFavoriteStickers() (*types.Stickers, error) {
	req := &types.GetFavoriteStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first.
func (c *Client) AddFavoriteSticker(sticker *types.InputFile) (*types.Ok, error) {
	req := &types.AddFavoriteSticker{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers @sticker Sticker file to delete from the list
func (c *Client) RemoveFavoriteSticker(sticker *types.InputFile) (*types.Ok, error) {
	req := &types.RemoveFavoriteSticker{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetStickerEmojis Returns emoji corresponding to a sticker. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object @sticker Sticker file identifier
func (c *Client) GetStickerEmojis(sticker *types.InputFile) (*types.Emojis, error) {
	req := &types.GetStickerEmojis{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Emojis), nil
}

// SearchEmojis Searches for emojis by keywords. Supported only if the file database is enabled. Order of results is unspecified
func (c *Client) SearchEmojis(text string, inputLanguageCodes []string) (*types.EmojiKeywords, error) {
	req := &types.SearchEmojis{
		Text:               text,
		InputLanguageCodes: inputLanguageCodes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiKeywords), nil
}

// GetKeywordEmojis Returns emojis matching the keyword. Supported only if the file database is enabled. Order of results is unspecified
func (c *Client) GetKeywordEmojis(text string, inputLanguageCodes []string) (*types.Emojis, error) {
	req := &types.GetKeywordEmojis{
		Text:               text,
		InputLanguageCodes: inputLanguageCodes,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Emojis), nil
}

// GetEmojiCategories Returns available emoji categories @type Type of emoji categories to return; pass null to get default emoji categories
func (c *Client) GetEmojiCategories(typeField *types.EmojiCategoryType) (*types.EmojiCategories, error) {
	req := &types.GetEmojiCategories{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmojiCategories), nil
}

// GetAnimatedEmoji Returns an animated emoji corresponding to a given emoji. Returns a 404 error if the emoji has no animated emoji @emoji The emoji
func (c *Client) GetAnimatedEmoji(emoji string) (*types.AnimatedEmoji, error) {
	req := &types.GetAnimatedEmoji{
		Emoji: emoji,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AnimatedEmoji), nil
}

// GetEmojiSuggestionsUrl Returns an HTTP URL which can be used to automatically log in to the translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
func (c *Client) GetEmojiSuggestionsUrl(languageCode string) (*types.HttpUrl, error) {
	req := &types.GetEmojiSuggestionsUrl{
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetCustomEmojiStickers Returns the list of custom emoji stickers by their identifiers. Stickers are returned in arbitrary order. Only found stickers are returned
func (c *Client) GetCustomEmojiStickers(customEmojiIds []string) (*types.Stickers, error) {
	req := &types.GetCustomEmojiStickers{
		CustomEmojiIds: customEmojiIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetDefaultChatPhotoCustomEmojiStickers Returns default list of custom emoji stickers for placing on a chat photo
func (c *Client) GetDefaultChatPhotoCustomEmojiStickers() (*types.Stickers, error) {
	req := &types.GetDefaultChatPhotoCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetDefaultProfilePhotoCustomEmojiStickers Returns default list of custom emoji stickers for placing on a profile photo
func (c *Client) GetDefaultProfilePhotoCustomEmojiStickers() (*types.Stickers, error) {
	req := &types.GetDefaultProfilePhotoCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetDefaultBackgroundCustomEmojiStickers Returns default list of custom emoji stickers for reply background
func (c *Client) GetDefaultBackgroundCustomEmojiStickers() (*types.Stickers, error) {
	req := &types.GetDefaultBackgroundCustomEmojiStickers{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetSavedAnimations Returns saved animations
func (c *Client) GetSavedAnimations() (*types.Animations, error) {
	req := &types.GetSavedAnimations{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Animations), nil
}

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first.
func (c *Client) AddSavedAnimation(animation *types.InputFile) (*types.Ok, error) {
	req := &types.AddSavedAnimation{
		Animation: animation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveSavedAnimation Removes an animation from the list of saved animations @animation Animation file to be removed
func (c *Client) RemoveSavedAnimation(animation *types.InputFile) (*types.Ok, error) {
	req := &types.RemoveSavedAnimation{
		Animation: animation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (c *Client) GetRecentInlineBots() (*types.Users, error) {
	req := &types.GetRecentInlineBots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// GetOwnedBots Returns the list of bots owned by the current user
func (c *Client) GetOwnedBots() (*types.Users, error) {
	req := &types.GetOwnedBots{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Users), nil
}

// SearchHashtags Searches for recently used hashtags by their prefix @prefix Hashtag prefix to search for @limit The maximum number of hashtags to be returned
func (c *Client) SearchHashtags(prefix string, limit int32) (*types.Hashtags, error) {
	req := &types.SearchHashtags{
		Prefix: prefix,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Hashtags), nil
}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags @hashtag Hashtag to delete
func (c *Client) RemoveRecentHashtag(hashtag string) (*types.Ok, error) {
	req := &types.RemoveRecentHashtag{
		Hashtag: hashtag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLinkPreview Returns a link preview by the text of a message. Do not call this function too often. Returns a 404 error if the text has no link preview
func (c *Client) GetLinkPreview(text *types.FormattedText, opts *types.GetLinkPreviewOpts) (*types.LinkPreview, error) {
	req := &types.GetLinkPreview{
		Text: text,
	}
	if opts != nil {
		req.LinkPreviewOptions = opts.LinkPreviewOptions
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LinkPreview), nil
}

// GetWebPageInstantView Returns an instant view version of a web page if available. This is an offline method if only_local is true. Returns a 404 error if the web page has no instant view page
func (c *Client) GetWebPageInstantView(url string, onlyLocal bool) (*types.WebPageInstantView, error) {
	req := &types.GetWebPageInstantView{
		Url:       url,
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.WebPageInstantView), nil
}

// SetProfilePhoto Changes a profile photo for the current user
func (c *Client) SetProfilePhoto(photo *types.InputChatPhoto, isPublic bool) (*types.Ok, error) {
	req := &types.SetProfilePhoto{
		Photo:    photo,
		IsPublic: isPublic,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteProfilePhoto Deletes a profile photo @profile_photo_id Identifier of the profile photo to delete
func (c *Client) DeleteProfilePhoto(profilePhotoId string) (*types.Ok, error) {
	req := &types.DeleteProfilePhoto{
		ProfilePhotoId: profilePhotoId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAccentColor Changes accent color and background custom emoji for the current user; for Telegram Premium users only
func (c *Client) SetAccentColor(accentColorId int32, backgroundCustomEmojiId string) (*types.Ok, error) {
	req := &types.SetAccentColor{
		AccentColorId:           accentColorId,
		BackgroundCustomEmojiId: backgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetUpgradedGiftColors Changes color scheme for the current user based on an owned or a hosted upgraded gift; for Telegram Premium users only
func (c *Client) SetUpgradedGiftColors(upgradedGiftColorsId string) (*types.Ok, error) {
	req := &types.SetUpgradedGiftColors{
		UpgradedGiftColorsId: upgradedGiftColorsId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetProfileAccentColor Changes accent color and background custom emoji for profile of the current user; for Telegram Premium users only
func (c *Client) SetProfileAccentColor(profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*types.Ok, error) {
	req := &types.SetProfileAccentColor{
		ProfileAccentColorId:           profileAccentColorId,
		ProfileBackgroundCustomEmojiId: profileBackgroundCustomEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetName Changes the first and last name of the current user @first_name The new value of the first name for the current user; 1-64 characters @last_name The new value of the optional last name for the current user; 0-64 characters
func (c *Client) SetName(firstName string, lastName string) (*types.Ok, error) {
	req := &types.SetName{
		FirstName: firstName,
		LastName:  lastName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBio Changes the bio of the current user @bio The new value of the user bio; 0-getOption("bio_length_max") characters without line feeds
func (c *Client) SetBio(bio string) (*types.Ok, error) {
	req := &types.SetBio{
		Bio: bio,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetUsername Changes the editable username of the current user
func (c *Client) SetUsername(username string) (*types.Ok, error) {
	req := &types.SetUsername{
		Username: username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleUsernameIsActive Changes active state for a username of the current user. The editable username can't be disabled. May return an error with a message "USERNAMES_ACTIVE_TOO_MUCH" if the maximum number of active usernames has been reached
func (c *Client) ToggleUsernameIsActive(username string, isActive bool) (*types.Ok, error) {
	req := &types.ToggleUsernameIsActive{
		Username: username,
		IsActive: isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReorderActiveUsernames Changes order of active usernames of the current user @usernames The new order of active usernames. All currently active usernames must be specified
func (c *Client) ReorderActiveUsernames(usernames []string) (*types.Ok, error) {
	req := &types.ReorderActiveUsernames{
		Usernames: usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBirthdate Changes the birthdate of the current user @birthdate The new value of the current user's birthdate; pass null to remove the birthdate
func (c *Client) SetBirthdate(birthdate *types.Birthdate) (*types.Ok, error) {
	req := &types.SetBirthdate{
		Birthdate: birthdate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetMainProfileTab Changes the main profile tab of the current user @main_profile_tab The new value of the main profile tab
func (c *Client) SetMainProfileTab(mainProfileTab *types.ProfileTab) (*types.Ok, error) {
	req := &types.SetMainProfileTab{
		MainProfileTab: mainProfileTab,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPersonalChat Changes the personal chat of the current user @chat_id Identifier of the new personal chat; pass 0 to remove the chat. Use getSuitablePersonalChats to get suitable chats
func (c *Client) SetPersonalChat(chatId int64) (*types.Ok, error) {
	req := &types.SetPersonalChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetEmojiStatus Changes the emoji status of the current user; for Telegram Premium users only @emoji_status New emoji status; pass null to switch to the default badge
func (c *Client) SetEmojiStatus(emojiStatus *types.EmojiStatus) (*types.Ok, error) {
	req := &types.SetEmojiStatus{
		EmojiStatus: emojiStatus,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleHasSponsoredMessagesEnabled Toggles whether the current user has sponsored messages enabled. The setting has no effect for users without Telegram Premium for which sponsored messages are always enabled
func (c *Client) ToggleHasSponsoredMessagesEnabled(hasSponsoredMessagesEnabled bool) (*types.Ok, error) {
	req := &types.ToggleHasSponsoredMessagesEnabled{
		HasSponsoredMessagesEnabled: hasSponsoredMessagesEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessLocation Changes the business location of the current user. Requires Telegram Business subscription @location The new location of the business; pass null to remove the location
func (c *Client) SetBusinessLocation(location *types.BusinessLocation) (*types.Ok, error) {
	req := &types.SetBusinessLocation{
		Location: location,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessOpeningHours Changes the business opening hours of the current user. Requires Telegram Business subscription
func (c *Client) SetBusinessOpeningHours(opts *types.SetBusinessOpeningHoursOpts) (*types.Ok, error) {
	req := &types.SetBusinessOpeningHours{}
	if opts != nil {
		req.OpeningHours = opts.OpeningHours
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessGreetingMessageSettings Changes the business greeting message settings of the current user. Requires Telegram Business subscription @greeting_message_settings The new settings for the greeting message of the business; pass null to disable the greeting message
func (c *Client) SetBusinessGreetingMessageSettings(greetingMessageSettings *types.BusinessGreetingMessageSettings) (*types.Ok, error) {
	req := &types.SetBusinessGreetingMessageSettings{
		GreetingMessageSettings: greetingMessageSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessAwayMessageSettings Changes the business away message settings of the current user. Requires Telegram Business subscription @away_message_settings The new settings for the away message of the business; pass null to disable the away message
func (c *Client) SetBusinessAwayMessageSettings(awayMessageSettings *types.BusinessAwayMessageSettings) (*types.Ok, error) {
	req := &types.SetBusinessAwayMessageSettings{
		AwayMessageSettings: awayMessageSettings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBusinessStartPage Changes the business start page of the current user. Requires Telegram Business subscription @start_page The new start page of the business; pass null to remove custom start page
func (c *Client) SetBusinessStartPage(startPage *types.InputBusinessStartPage) (*types.Ok, error) {
	req := &types.SetBusinessStartPage{
		StartPage: startPage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendPhoneNumberCode Sends a code to the specified phone number. Aborts previous phone number verification if there was one. On success, returns information about the sent code
func (c *Client) SendPhoneNumberCode(phoneNumber string, typeField *types.PhoneNumberCodeType, opts *types.SendPhoneNumberCodeOpts) (*types.AuthenticationCodeInfo, error) {
	req := &types.SendPhoneNumberCode{
		PhoneNumber: phoneNumber,
		TypeField:   typeField,
	}
	if opts != nil {
		req.Settings = opts.Settings
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AuthenticationCodeInfo), nil
}

// SendPhoneNumberFirebaseSms Sends Firebase Authentication SMS to the specified phone number. Works only when received a code of the type authenticationCodeTypeFirebaseAndroid or authenticationCodeTypeFirebaseIos
func (c *Client) SendPhoneNumberFirebaseSms(token string) (*types.Ok, error) {
	req := &types.SendPhoneNumberFirebaseSms{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportPhoneNumberCodeMissing Reports that authentication code wasn't delivered via SMS to the specified phone number; for official mobile applications only @mobile_network_code Current mobile network code
func (c *Client) ReportPhoneNumberCodeMissing(mobileNetworkCode string) (*types.Ok, error) {
	req := &types.ReportPhoneNumberCodeMissing{
		MobileNetworkCode: mobileNetworkCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResendPhoneNumberCode Resends the authentication code sent to a phone number. Works only if the previously received authenticationCodeInfo next_code_type was not null and the server-specified timeout has passed
func (c *Client) ResendPhoneNumberCode(opts *types.ResendPhoneNumberCodeOpts) (*types.AuthenticationCodeInfo, error) {
	req := &types.ResendPhoneNumberCode{}
	if opts != nil {
		req.Reason = opts.Reason
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AuthenticationCodeInfo), nil
}

// CheckPhoneNumberCode Checks the authentication code and completes the request for which the code was sent if appropriate @code Authentication code to check
func (c *Client) CheckPhoneNumberCode(code string) (*types.Ok, error) {
	req := &types.CheckPhoneNumberCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBusinessConnectedBot Returns the business bot that is connected to the current user account. Returns a 404 error if there is no connected bot
func (c *Client) GetBusinessConnectedBot() (*types.BusinessConnectedBot, error) {
	req := &types.GetBusinessConnectedBot{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessConnectedBot), nil
}

// SetBusinessConnectedBot Adds or changes business bot that is connected to the current user account @bot Connection settings for the bot
func (c *Client) SetBusinessConnectedBot(bot *types.BusinessConnectedBot) (*types.Ok, error) {
	req := &types.SetBusinessConnectedBot{
		Bot: bot,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteBusinessConnectedBot Deletes the business bot that is connected to the current user account @bot_user_id Unique user identifier for the bot
func (c *Client) DeleteBusinessConnectedBot(botUserId int64) (*types.Ok, error) {
	req := &types.DeleteBusinessConnectedBot{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleBusinessConnectedBotChatIsPaused Pauses or resumes the connected business bot in a specific chat @chat_id Chat identifier @is_paused Pass true to pause the connected bot in the chat; pass false to resume the bot
func (c *Client) ToggleBusinessConnectedBotChatIsPaused(chatId int64, isPaused bool) (*types.Ok, error) {
	req := &types.ToggleBusinessConnectedBotChatIsPaused{
		ChatId:   chatId,
		IsPaused: isPaused,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveBusinessConnectedBotFromChat Removes the connected business bot from a specific chat by adding the chat to businessRecipients.excluded_chat_ids @chat_id Chat identifier
func (c *Client) RemoveBusinessConnectedBotFromChat(chatId int64) (*types.Ok, error) {
	req := &types.RemoveBusinessConnectedBotFromChat{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBusinessChatLinks Returns business chat links created for the current account
func (c *Client) GetBusinessChatLinks() (*types.BusinessChatLinks, error) {
	req := &types.GetBusinessChatLinks{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessChatLinks), nil
}

// CreateBusinessChatLink Creates a business chat link for the current account. Requires Telegram Business subscription. There can be up to getOption("business_chat_link_count_max") links created. Returns the created link
func (c *Client) CreateBusinessChatLink(linkInfo *types.InputBusinessChatLink) (*types.BusinessChatLink, error) {
	req := &types.CreateBusinessChatLink{
		LinkInfo: linkInfo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessChatLink), nil
}

// EditBusinessChatLink Edits a business chat link of the current account. Requires Telegram Business subscription. Returns the edited link
func (c *Client) EditBusinessChatLink(link string, linkInfo *types.InputBusinessChatLink) (*types.BusinessChatLink, error) {
	req := &types.EditBusinessChatLink{
		Link:     link,
		LinkInfo: linkInfo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessChatLink), nil
}

// DeleteBusinessChatLink Deletes a business chat link of the current account @link The link to delete
func (c *Client) DeleteBusinessChatLink(link string) (*types.Ok, error) {
	req := &types.DeleteBusinessChatLink{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBusinessChatLinkInfo Returns information about a business chat link @link_name Name of the link
func (c *Client) GetBusinessChatLinkInfo(linkName string) (*types.BusinessChatLinkInfo, error) {
	req := &types.GetBusinessChatLinkInfo{
		LinkName: linkName,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessChatLinkInfo), nil
}

// GetUserLink Returns an HTTPS link, which can be used to get information about the current user
func (c *Client) GetUserLink() (*types.UserLink, error) {
	req := &types.GetUserLink{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UserLink), nil
}

// SearchUserByToken Searches a user by a token from the user's link @token Token to search for
func (c *Client) SearchUserByToken(token string) (*types.User, error) {
	req := &types.SearchUserByToken{
		Token: token,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// SetCommands Sets the list of commands supported by the bot for the given user scope and language; for bots only
func (c *Client) SetCommands(languageCode string, commands []*types.BotCommand, opts *types.SetCommandsOpts) (*types.Ok, error) {
	req := &types.SetCommands{
		LanguageCode: languageCode,
		Commands:     commands,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteCommands Deletes commands supported by the bot for the given user scope and language; for bots only
func (c *Client) DeleteCommands(languageCode string, opts *types.DeleteCommandsOpts) (*types.Ok, error) {
	req := &types.DeleteCommands{
		LanguageCode: languageCode,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetCommands Returns the list of commands supported by the bot for the given user scope and language; for bots only
func (c *Client) GetCommands(languageCode string, opts *types.GetCommandsOpts) (*types.BotCommands, error) {
	req := &types.GetCommands{
		LanguageCode: languageCode,
	}
	if opts != nil {
		req.Scope = opts.Scope
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotCommands), nil
}

// SetMenuButton Sets menu button for the given user or for all users; for bots only
func (c *Client) SetMenuButton(userId int64, menuButton *types.BotMenuButton) (*types.Ok, error) {
	req := &types.SetMenuButton{
		UserId:     userId,
		MenuButton: menuButton,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetMenuButton Returns menu button set by the bot for the given user; for bots only @user_id Identifier of the user or 0 to get the default menu button
func (c *Client) GetMenuButton(userId int64) (*types.BotMenuButton, error) {
	req := &types.GetMenuButton{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotMenuButton), nil
}

// SetDefaultGroupAdministratorRights Sets default administrator rights for adding the bot to basic group and supergroup chats; for bots only @default_group_administrator_rights Default administrator rights for adding the bot to basic group and supergroup chats; pass null to remove default rights
func (c *Client) SetDefaultGroupAdministratorRights(defaultGroupAdministratorRights *types.ChatAdministratorRights) (*types.Ok, error) {
	req := &types.SetDefaultGroupAdministratorRights{
		DefaultGroupAdministratorRights: defaultGroupAdministratorRights,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetDefaultChannelAdministratorRights Sets default administrator rights for adding the bot to channel chats; for bots only @default_channel_administrator_rights Default administrator rights for adding the bot to channels; pass null to remove default rights
func (c *Client) SetDefaultChannelAdministratorRights(defaultChannelAdministratorRights *types.ChatAdministratorRights) (*types.Ok, error) {
	req := &types.SetDefaultChannelAdministratorRights{
		DefaultChannelAdministratorRights: defaultChannelAdministratorRights,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CanBotSendMessages Checks whether the specified bot can send messages to the user. Returns a 404 error if can't and the access can be granted by call to allowBotToSendMessages @bot_user_id Identifier of the target bot
func (c *Client) CanBotSendMessages(botUserId int64) (*types.Ok, error) {
	req := &types.CanBotSendMessages{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AllowBotToSendMessages Allows the specified bot to send messages to the user @bot_user_id Identifier of the target bot
func (c *Client) AllowBotToSendMessages(botUserId int64) (*types.Ok, error) {
	req := &types.AllowBotToSendMessages{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendWebAppCustomRequest Sends a custom request from a Web App
func (c *Client) SendWebAppCustomRequest(botUserId int64, method string, parameters string) (*types.CustomRequestResult, error) {
	req := &types.SendWebAppCustomRequest{
		BotUserId:  botUserId,
		Method:     method,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CustomRequestResult), nil
}

// GetBotMediaPreviews Returns the list of media previews of a bot @bot_user_id Identifier of the target bot. The bot must have the main Web App
func (c *Client) GetBotMediaPreviews(botUserId int64) (*types.BotMediaPreviews, error) {
	req := &types.GetBotMediaPreviews{
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotMediaPreviews), nil
}

// GetBotMediaPreviewInfo Returns the list of media previews for the given language and the list of languages for which the bot has dedicated previews
func (c *Client) GetBotMediaPreviewInfo(botUserId int64, languageCode string) (*types.BotMediaPreviewInfo, error) {
	req := &types.GetBotMediaPreviewInfo{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotMediaPreviewInfo), nil
}

// AddBotMediaPreview Adds a new media preview to the beginning of the list of media previews of a bot. Returns the added preview after addition is completed server-side. The total number of previews must not exceed getOption("bot_media_preview_count_max") for the given language
func (c *Client) AddBotMediaPreview(botUserId int64, languageCode string, content *types.InputStoryContent) (*types.BotMediaPreview, error) {
	req := &types.AddBotMediaPreview{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Content:      content,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotMediaPreview), nil
}

// EditBotMediaPreview Replaces media preview in the list of media previews of a bot. Returns the new preview after edit is completed server-side
func (c *Client) EditBotMediaPreview(botUserId int64, languageCode string, fileId int32, content *types.InputStoryContent) (*types.BotMediaPreview, error) {
	req := &types.EditBotMediaPreview{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileId:       fileId,
		Content:      content,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BotMediaPreview), nil
}

// ReorderBotMediaPreviews Changes order of media previews in the list of media previews of a bot
func (c *Client) ReorderBotMediaPreviews(botUserId int64, languageCode string, fileIds []int32) (*types.Ok, error) {
	req := &types.ReorderBotMediaPreviews{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileIds:      fileIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteBotMediaPreviews Deletes media previews from the list of media previews of a bot
func (c *Client) DeleteBotMediaPreviews(botUserId int64, languageCode string, fileIds []int32) (*types.Ok, error) {
	req := &types.DeleteBotMediaPreviews{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		FileIds:      fileIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBotName Sets the name of a bot. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotName(botUserId int64, languageCode string, name string) (*types.Ok, error) {
	req := &types.SetBotName{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBotName Returns the name of a bot in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotName(botUserId int64, languageCode string) (*types.Text, error) {
	req := &types.GetBotName{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SetBotProfilePhoto Changes a profile photo for a bot @bot_user_id Identifier of the target bot @photo Profile photo to set; pass null to delete the chat photo
func (c *Client) SetBotProfilePhoto(botUserId int64, photo *types.InputChatPhoto) (*types.Ok, error) {
	req := &types.SetBotProfilePhoto{
		BotUserId: botUserId,
		Photo:     photo,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleBotUsernameIsActive Changes active state for a username of a bot. The editable username can be disabled only if there are other active usernames.
func (c *Client) ToggleBotUsernameIsActive(botUserId int64, username string, isActive bool) (*types.Ok, error) {
	req := &types.ToggleBotUsernameIsActive{
		BotUserId: botUserId,
		Username:  username,
		IsActive:  isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReorderBotActiveUsernames Changes order of active usernames of a bot. Can be called only if userTypeBot.can_be_edited == true @bot_user_id Identifier of the target bot @usernames The new order of active usernames. All currently active usernames must be specified
func (c *Client) ReorderBotActiveUsernames(botUserId int64, usernames []string) (*types.Ok, error) {
	req := &types.ReorderBotActiveUsernames{
		BotUserId: botUserId,
		Usernames: usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBotInfoDescription Sets the text shown in the chat with a bot if the chat is empty. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotInfoDescription(botUserId int64, languageCode string, description string) (*types.Ok, error) {
	req := &types.SetBotInfoDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
		Description:  description,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBotInfoDescription Returns the text shown in the chat with a bot if the chat is empty in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotInfoDescription(botUserId int64, languageCode string) (*types.Text, error) {
	req := &types.GetBotInfoDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SetBotInfoShortDescription Sets the text shown on a bot's profile page and sent together with the link when users share the bot. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) SetBotInfoShortDescription(botUserId int64, languageCode string, shortDescription string) (*types.Ok, error) {
	req := &types.SetBotInfoShortDescription{
		BotUserId:        botUserId,
		LanguageCode:     languageCode,
		ShortDescription: shortDescription,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBotInfoShortDescription Returns the text shown on a bot's profile page and sent together with the link when users share the bot in the given language. Can be called only if userTypeBot.can_be_edited == true
func (c *Client) GetBotInfoShortDescription(botUserId int64, languageCode string) (*types.Text, error) {
	req := &types.GetBotInfoShortDescription{
		BotUserId:    botUserId,
		LanguageCode: languageCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SetMessageSenderBotVerification Changes the verification status of a user or a chat by an owned bot
func (c *Client) SetMessageSenderBotVerification(botUserId int64, verifiedId *types.MessageSender, customDescription string) (*types.Ok, error) {
	req := &types.SetMessageSenderBotVerification{
		BotUserId:         botUserId,
		VerifiedId:        verifiedId,
		CustomDescription: customDescription,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveMessageSenderBotVerification Removes the verification status of a user or a chat by an owned bot
func (c *Client) RemoveMessageSenderBotVerification(botUserId int64, verifiedId *types.MessageSender) (*types.Ok, error) {
	req := &types.RemoveMessageSenderBotVerification{
		BotUserId:  botUserId,
		VerifiedId: verifiedId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetActiveSessions Returns all active sessions of the current user
func (c *Client) GetActiveSessions() (*types.Sessions, error) {
	req := &types.GetActiveSessions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Sessions), nil
}

// TerminateSession Terminates a session of the current user @session_id Session identifier
func (c *Client) TerminateSession(sessionId string) (*types.Ok, error) {
	req := &types.TerminateSession{
		SessionId: sessionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (c *Client) TerminateAllOtherSessions() (*types.Ok, error) {
	req := &types.TerminateAllOtherSessions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ConfirmSession Confirms an unconfirmed session of the current user from another device @session_id Session identifier
func (c *Client) ConfirmSession(sessionId string) (*types.Ok, error) {
	req := &types.ConfirmSession{
		SessionId: sessionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSessionCanAcceptCalls Toggles whether a session can accept incoming calls @session_id Session identifier @can_accept_calls Pass true to allow accepting incoming calls by the session; pass false otherwise
func (c *Client) ToggleSessionCanAcceptCalls(sessionId string, canAcceptCalls bool) (*types.Ok, error) {
	req := &types.ToggleSessionCanAcceptCalls{
		SessionId:      sessionId,
		CanAcceptCalls: canAcceptCalls,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSessionCanAcceptSecretChats Toggles whether a session can accept incoming secret chats @session_id Session identifier @can_accept_secret_chats Pass true to allow accepting secret chats by the session; pass false otherwise
func (c *Client) ToggleSessionCanAcceptSecretChats(sessionId string, canAcceptSecretChats bool) (*types.Ok, error) {
	req := &types.ToggleSessionCanAcceptSecretChats{
		SessionId:            sessionId,
		CanAcceptSecretChats: canAcceptSecretChats,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetInactiveSessionTtl Changes the period of inactivity after which sessions will automatically be terminated @inactive_session_ttl_days New number of days of inactivity before sessions will be automatically terminated; 1-366 days
func (c *Client) SetInactiveSessionTtl(inactiveSessionTtlDays int32) (*types.Ok, error) {
	req := &types.SetInactiveSessionTtl{
		InactiveSessionTtlDays: inactiveSessionTtlDays,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (c *Client) GetConnectedWebsites() (*types.ConnectedWebsites, error) {
	req := &types.GetConnectedWebsites{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ConnectedWebsites), nil
}

// DisconnectWebsite Disconnects website from the current user's Telegram account @website_id Website identifier
func (c *Client) DisconnectWebsite(websiteId string) (*types.Ok, error) {
	req := &types.DisconnectWebsite{
		WebsiteId: websiteId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (c *Client) DisconnectAllWebsites() (*types.Ok, error) {
	req := &types.DisconnectAllWebsites{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetSupergroupUsername Changes the editable username of a supergroup or channel, requires owner privileges in the supergroup or channel
func (c *Client) SetSupergroupUsername(supergroupId int64, username string) (*types.Ok, error) {
	req := &types.SetSupergroupUsername{
		SupergroupId: supergroupId,
		Username:     username,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupUsernameIsActive Changes active state for a username of a supergroup or channel, requires owner privileges in the supergroup or channel. The editable username can't be disabled.
func (c *Client) ToggleSupergroupUsernameIsActive(supergroupId int64, username string, isActive bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupUsernameIsActive{
		SupergroupId: supergroupId,
		Username:     username,
		IsActive:     isActive,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DisableAllSupergroupUsernames Disables all active non-editable usernames of a supergroup or channel, requires owner privileges in the supergroup or channel @supergroup_id Identifier of the supergroup or channel
func (c *Client) DisableAllSupergroupUsernames(supergroupId int64) (*types.Ok, error) {
	req := &types.DisableAllSupergroupUsernames{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReorderSupergroupActiveUsernames Changes order of active usernames of a supergroup or channel, requires owner privileges in the supergroup or channel
func (c *Client) ReorderSupergroupActiveUsernames(supergroupId int64, usernames []string) (*types.Ok, error) {
	req := &types.ReorderSupergroupActiveUsernames{
		SupergroupId: supergroupId,
		Usernames:    usernames,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires can_change_info administrator right @supergroup_id Identifier of the supergroup @sticker_set_id New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (c *Client) SetSupergroupStickerSet(supergroupId int64, stickerSetId string) (*types.Ok, error) {
	req := &types.SetSupergroupStickerSet{
		SupergroupId: supergroupId,
		StickerSetId: stickerSetId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetSupergroupCustomEmojiStickerSet Changes the custom emoji sticker set of a supergroup; requires can_change_info administrator right. The chat must have at least chatBoostFeatures.min_custom_emoji_sticker_set_boost_level boost level to pass the corresponding color
func (c *Client) SetSupergroupCustomEmojiStickerSet(supergroupId int64, customEmojiStickerSetId string) (*types.Ok, error) {
	req := &types.SetSupergroupCustomEmojiStickerSet{
		SupergroupId:            supergroupId,
		CustomEmojiStickerSetId: customEmojiStickerSetId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetSupergroupUnrestrictBoostCount Changes the number of times the supergroup must be boosted by a user to ignore slow mode and chat permission restrictions; requires can_restrict_members administrator right
func (c *Client) SetSupergroupUnrestrictBoostCount(supergroupId int64, unrestrictBoostCount int32) (*types.Ok, error) {
	req := &types.SetSupergroupUnrestrictBoostCount{
		SupergroupId:         supergroupId,
		UnrestrictBoostCount: unrestrictBoostCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetSupergroupMainProfileTab Changes the main profile tab of the channel; requires can_change_info administrator right
func (c *Client) SetSupergroupMainProfileTab(supergroupId int64, mainProfileTab *types.ProfileTab) (*types.Ok, error) {
	req := &types.SetSupergroupMainProfileTab{
		SupergroupId:   supergroupId,
		MainProfileTab: mainProfileTab,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupSignMessages Toggles whether sender signature or link to the account is added to sent messages in a channel; requires can_change_info member right
func (c *Client) ToggleSupergroupSignMessages(supergroupId int64, signMessages bool, showMessageSender bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupSignMessages{
		SupergroupId:      supergroupId,
		SignMessages:      signMessages,
		ShowMessageSender: showMessageSender,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupJoinToSendMessages Toggles whether joining is mandatory to send messages to a discussion supergroup; requires can_restrict_members administrator right
func (c *Client) ToggleSupergroupJoinToSendMessages(supergroupId int64, joinToSendMessages bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupJoinToSendMessages{
		SupergroupId:       supergroupId,
		JoinToSendMessages: joinToSendMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupJoinByRequest Toggles whether all users directly joining the supergroup need to be approved by supergroup administrators; requires can_restrict_members administrator right
func (c *Client) ToggleSupergroupJoinByRequest(supergroupId int64, joinByRequest bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupJoinByRequest{
		SupergroupId:  supergroupId,
		JoinByRequest: joinByRequest,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires can_change_info member right @supergroup_id The identifier of the supergroup @is_all_history_available The new value of is_all_history_available
func (c *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupId int64, isAllHistoryAvailable bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupIsAllHistoryAvailable{
		SupergroupId:          supergroupId,
		IsAllHistoryAvailable: isAllHistoryAvailable,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupCanHaveSponsoredMessages Toggles whether sponsored messages are shown in the channel chat; requires owner privileges in the channel. The chat must have at least chatBoostFeatures.min_sponsored_message_disable_boost_level boost level to disable sponsored messages
func (c *Client) ToggleSupergroupCanHaveSponsoredMessages(supergroupId int64, canHaveSponsoredMessages bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupCanHaveSponsoredMessages{
		SupergroupId:             supergroupId,
		CanHaveSponsoredMessages: canHaveSponsoredMessages,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupHasAutomaticTranslation Toggles whether messages are automatically translated in the channel chat; requires can_change_info administrator right in the channel.
func (c *Client) ToggleSupergroupHasAutomaticTranslation(supergroupId int64, hasAutomaticTranslation bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupHasAutomaticTranslation{
		SupergroupId:            supergroupId,
		HasAutomaticTranslation: hasAutomaticTranslation,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupHasHiddenMembers Toggles whether non-administrators can receive only administrators and bots using getSupergroupMembers or searchChatMembers. Can be called only if supergroupFullInfo.can_hide_members == true
func (c *Client) ToggleSupergroupHasHiddenMembers(supergroupId int64, hasHiddenMembers bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupHasHiddenMembers{
		SupergroupId:     supergroupId,
		HasHiddenMembers: hasHiddenMembers,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled Toggles whether aggressive anti-spam checks are enabled in the supergroup. Can be called only if supergroupFullInfo.can_toggle_aggressive_anti_spam == true
func (c *Client) ToggleSupergroupHasAggressiveAntiSpamEnabled(supergroupId int64, hasAggressiveAntiSpamEnabled bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupHasAggressiveAntiSpamEnabled{
		SupergroupId:                 supergroupId,
		HasAggressiveAntiSpamEnabled: hasAggressiveAntiSpamEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupIsForum Toggles whether the supergroup is a forum; requires owner privileges in the supergroup. Discussion supergroups can't be converted to forums
func (c *Client) ToggleSupergroupIsForum(supergroupId int64, isForum bool, hasForumTabs bool) (*types.Ok, error) {
	req := &types.ToggleSupergroupIsForum{
		SupergroupId: supergroupId,
		IsForum:      isForum,
		HasForumTabs: hasForumTabs,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleSupergroupIsBroadcastGroup Upgrades supergroup to a broadcast group; requires owner privileges in the supergroup @supergroup_id Identifier of the supergroup
func (c *Client) ToggleSupergroupIsBroadcastGroup(supergroupId int64) (*types.Ok, error) {
	req := &types.ToggleSupergroupIsBroadcastGroup{
		SupergroupId: supergroupId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportSupergroupSpam Reports messages in a supergroup as spam; requires administrator rights in the supergroup
func (c *Client) ReportSupergroupSpam(supergroupId int64, messageIds []int64) (*types.Ok, error) {
	req := &types.ReportSupergroupSpam{
		SupergroupId: supergroupId,
		MessageIds:   messageIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportSupergroupAntiSpamFalsePositive Reports a false deletion of a message by aggressive anti-spam checks; requires administrator rights in the supergroup. Can be called only for messages from chatEventMessageDeleted with can_report_anti_spam_false_positive == true
func (c *Client) ReportSupergroupAntiSpamFalsePositive(supergroupId int64, messageId int64) (*types.Ok, error) {
	req := &types.ReportSupergroupAntiSpamFalsePositive{
		SupergroupId: supergroupId,
		MessageId:    messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if supergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
func (c *Client) GetSupergroupMembers(supergroupId int64, offset int32, limit int32, opts *types.GetSupergroupMembersOpts) (*types.ChatMembers, error) {
	req := &types.GetSupergroupMembers{
		SupergroupId: supergroupId,
		Offset:       offset,
		Limit:        limit,
	}
	if opts != nil {
		req.Filter = opts.Filter
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatMembers), nil
}

// CloseSecretChat Closes a secret chat, effectively transferring its state to secretChatStateClosed @secret_chat_id Secret chat identifier
func (c *Client) CloseSecretChat(secretChatId int32) (*types.Ok, error) {
	req := &types.CloseSecretChat{
		SecretChatId: secretChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i.e., in order of decreasing event_id)
func (c *Client) GetChatEventLog(chatId int64, query string, fromEventId string, limit int32, userIds []int64, opts *types.GetChatEventLogOpts) (*types.ChatEvents, error) {
	req := &types.GetChatEventLog{
		ChatId:      chatId,
		Query:       query,
		FromEventId: fromEventId,
		Limit:       limit,
		UserIds:     userIds,
	}
	if opts != nil {
		req.Filters = opts.Filters
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatEvents), nil
}

// GetTimeZones Returns the list of supported time zones
func (c *Client) GetTimeZones() (*types.TimeZones, error) {
	req := &types.GetTimeZones{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TimeZones), nil
}

// GetPaymentForm Returns an invoice payment form. This method must be called when the user presses inline button of the type inlineKeyboardButtonTypeBuy, or wants to buy access to media in a messagePaidMedia message
func (c *Client) GetPaymentForm(inputInvoice *types.InputInvoice, opts *types.GetPaymentFormOpts) (*types.PaymentForm, error) {
	req := &types.GetPaymentForm{
		InputInvoice: inputInvoice,
	}
	if opts != nil {
		req.Theme = opts.Theme
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PaymentForm), nil
}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
func (c *Client) ValidateOrderInfo(inputInvoice *types.InputInvoice, allowSave bool, opts *types.ValidateOrderInfoOpts) (*types.ValidatedOrderInfo, error) {
	req := &types.ValidateOrderInfo{
		InputInvoice: inputInvoice,
		AllowSave:    allowSave,
	}
	if opts != nil {
		req.OrderInfo = opts.OrderInfo
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ValidatedOrderInfo), nil
}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
func (c *Client) SendPaymentForm(inputInvoice *types.InputInvoice, paymentFormId string, orderInfoId string, shippingOptionId string, tipAmount int64, opts *types.SendPaymentFormOpts) (*types.PaymentResult, error) {
	req := &types.SendPaymentForm{
		InputInvoice:     inputInvoice,
		PaymentFormId:    paymentFormId,
		OrderInfoId:      orderInfoId,
		ShippingOptionId: shippingOptionId,
		TipAmount:        tipAmount,
	}
	if opts != nil {
		req.Credentials = opts.Credentials
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PaymentResult), nil
}

// GetPaymentReceipt Returns information about a successful payment @chat_id Chat identifier of the messagePaymentSuccessful message @message_id Message identifier
func (c *Client) GetPaymentReceipt(chatId int64, messageId int64) (*types.PaymentReceipt, error) {
	req := &types.GetPaymentReceipt{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PaymentReceipt), nil
}

// GetSavedOrderInfo Returns saved order information. Returns a 404 error if there is no saved order information
func (c *Client) GetSavedOrderInfo() (*types.OrderInfo, error) {
	req := &types.GetSavedOrderInfo{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.OrderInfo), nil
}

// DeleteSavedOrderInfo Deletes saved order information
func (c *Client) DeleteSavedOrderInfo() (*types.Ok, error) {
	req := &types.DeleteSavedOrderInfo{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (c *Client) DeleteSavedCredentials() (*types.Ok, error) {
	req := &types.DeleteSavedCredentials{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGiftSettings Changes settings for gift receiving for the current user @settings The new settings
func (c *Client) SetGiftSettings(settings *types.GiftSettings) (*types.Ok, error) {
	req := &types.SetGiftSettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetAvailableGifts Returns gifts that can be sent to other users and channel chats
func (c *Client) GetAvailableGifts() (*types.AvailableGifts, error) {
	req := &types.GetAvailableGifts{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AvailableGifts), nil
}

// CanSendGift Checks whether a gift with next_send_date in the future can be sent already
func (c *Client) CanSendGift(giftId string) (*types.CanSendGiftResult, error) {
	req := &types.CanSendGift{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CanSendGiftResult), nil
}

// SendGift Sends a gift to another user or channel chat. May return an error with a message "STARGIFT_USAGE_LIMITED" if the gift was sold out
func (c *Client) SendGift(giftId string, ownerId *types.MessageSender, text *types.FormattedText, isPrivate bool, payForUpgrade bool) (*types.Ok, error) {
	req := &types.SendGift{
		GiftId:        giftId,
		OwnerId:       ownerId,
		Text:          text,
		IsPrivate:     isPrivate,
		PayForUpgrade: payForUpgrade,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGiftAuctionState Returns auction state for a gift @auction_id Unique identifier of the auction
func (c *Client) GetGiftAuctionState(auctionId string) (*types.GiftAuctionState, error) {
	req := &types.GetGiftAuctionState{
		AuctionId: auctionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftAuctionState), nil
}

// GetGiftAuctionAcquiredGifts Returns the gifts that were acquired by the current user on a gift auction @gift_id Identifier of the auctioned gift
func (c *Client) GetGiftAuctionAcquiredGifts(giftId string) (*types.GiftAuctionAcquiredGifts, error) {
	req := &types.GetGiftAuctionAcquiredGifts{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftAuctionAcquiredGifts), nil
}

// OpenGiftAuction Informs TDLib that a gift auction was opened by the user @gift_id Identifier of the gift, which auction was opened
func (c *Client) OpenGiftAuction(giftId string) (*types.Ok, error) {
	req := &types.OpenGiftAuction{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CloseGiftAuction Informs TDLib that a gift auction was closed by the user @gift_id Identifier of the gift, which auction was closed
func (c *Client) CloseGiftAuction(giftId string) (*types.Ok, error) {
	req := &types.CloseGiftAuction{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// PlaceGiftAuctionBid Places a bid on an auction gift
func (c *Client) PlaceGiftAuctionBid(giftId string, starCount int64, userId int64, text *types.FormattedText, isPrivate bool) (*types.Ok, error) {
	req := &types.PlaceGiftAuctionBid{
		GiftId:    giftId,
		StarCount: starCount,
		UserId:    userId,
		Text:      text,
		IsPrivate: isPrivate,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// IncreaseGiftAuctionBid Increases a bid for an auction gift without changing gift text and receiver
func (c *Client) IncreaseGiftAuctionBid(giftId string, starCount int64) (*types.Ok, error) {
	req := &types.IncreaseGiftAuctionBid{
		GiftId:    giftId,
		StarCount: starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SellGift Sells a gift for Telegram Stars; requires owner privileges for gifts owned by a chat
func (c *Client) SellGift(businessConnectionId string, receivedGiftId string) (*types.Ok, error) {
	req := &types.SellGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleGiftIsSaved Toggles whether a gift is shown on the current user's or the channel's profile page; requires can_post_messages administrator right in the channel chat
func (c *Client) ToggleGiftIsSaved(receivedGiftId string, isSaved bool) (*types.Ok, error) {
	req := &types.ToggleGiftIsSaved{
		ReceivedGiftId: receivedGiftId,
		IsSaved:        isSaved,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPinnedGifts Changes the list of pinned gifts on the current user's or the channel's profile page; requires can_post_messages administrator right in the channel chat
func (c *Client) SetPinnedGifts(ownerId *types.MessageSender, receivedGiftIds []string) (*types.Ok, error) {
	req := &types.SetPinnedGifts{
		OwnerId:         ownerId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ToggleChatGiftNotifications Toggles whether notifications for new gifts received by a channel chat are sent to the current user; requires can_post_messages administrator right in the chat
func (c *Client) ToggleChatGiftNotifications(chatId int64, areEnabled bool) (*types.Ok, error) {
	req := &types.ToggleChatGiftNotifications{
		ChatId:     chatId,
		AreEnabled: areEnabled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGiftUpgradePreview Returns examples of possible upgraded gifts for a regular gift @gift_id Identifier of the gift
func (c *Client) GetGiftUpgradePreview(giftId string) (*types.GiftUpgradePreview, error) {
	req := &types.GetGiftUpgradePreview{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftUpgradePreview), nil
}

// GetGiftUpgradeVariants Returns all possible variants of upgraded gifts for a regular gift @gift_id Identifier of the gift
func (c *Client) GetGiftUpgradeVariants(giftId string) (*types.GiftUpgradeVariants, error) {
	req := &types.GetGiftUpgradeVariants{
		GiftId: giftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftUpgradeVariants), nil
}

// UpgradeGift Upgrades a regular gift
func (c *Client) UpgradeGift(businessConnectionId string, receivedGiftId string, keepOriginalDetails bool, starCount int64) (*types.UpgradeGiftResult, error) {
	req := &types.UpgradeGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
		KeepOriginalDetails:  keepOriginalDetails,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UpgradeGiftResult), nil
}

// BuyGiftUpgrade Pays for upgrade of a regular gift that is owned by another user or channel chat
func (c *Client) BuyGiftUpgrade(ownerId *types.MessageSender, prepaidUpgradeHash string, starCount int64) (*types.Ok, error) {
	req := &types.BuyGiftUpgrade{
		OwnerId:            ownerId,
		PrepaidUpgradeHash: prepaidUpgradeHash,
		StarCount:          starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TransferGift Sends an upgraded gift to another user or channel chat
func (c *Client) TransferGift(businessConnectionId string, receivedGiftId string, newOwnerId *types.MessageSender, starCount int64) (*types.Ok, error) {
	req := &types.TransferGift{
		BusinessConnectionId: businessConnectionId,
		ReceivedGiftId:       receivedGiftId,
		NewOwnerId:           newOwnerId,
		StarCount:            starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DropGiftOriginalDetails Drops original details for an upgraded gift
func (c *Client) DropGiftOriginalDetails(receivedGiftId string, starCount int64) (*types.Ok, error) {
	req := &types.DropGiftOriginalDetails{
		ReceivedGiftId: receivedGiftId,
		StarCount:      starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SendResoldGift Sends an upgraded gift that is available for resale to another user or channel chat; gifts already owned by the current user
func (c *Client) SendResoldGift(giftName string, ownerId *types.MessageSender, price *types.GiftResalePrice) (*types.GiftResaleResult, error) {
	req := &types.SendResoldGift{
		GiftName: giftName,
		OwnerId:  ownerId,
		Price:    price,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftResaleResult), nil
}

// SendGiftPurchaseOffer Sends an offer to purchase an upgraded gift
func (c *Client) SendGiftPurchaseOffer(ownerId *types.MessageSender, giftName string, price *types.GiftResalePrice, duration int32, paidMessageStarCount int64) (*types.Ok, error) {
	req := &types.SendGiftPurchaseOffer{
		OwnerId:              ownerId,
		GiftName:             giftName,
		Price:                price,
		Duration:             duration,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ProcessGiftPurchaseOffer Handles a pending gift purchase offer
func (c *Client) ProcessGiftPurchaseOffer(messageId int64, accept bool) (*types.Ok, error) {
	req := &types.ProcessGiftPurchaseOffer{
		MessageId: messageId,
		Accept:    accept,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetReceivedGifts Returns gifts received by the given user or chat
func (c *Client) GetReceivedGifts(businessConnectionId string, ownerId *types.MessageSender, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*types.ReceivedGifts, error) {
	req := &types.GetReceivedGifts{
		BusinessConnectionId: businessConnectionId,
		OwnerId:              ownerId,
		CollectionId:         collectionId,
		ExcludeUnsaved:       excludeUnsaved,
		ExcludeSaved:         excludeSaved,
		ExcludeUnlimited:     excludeUnlimited,
		ExcludeUpgradable:    excludeUpgradable,
		ExcludeNonUpgradable: excludeNonUpgradable,
		ExcludeUpgraded:      excludeUpgraded,
		ExcludeWithoutColors: excludeWithoutColors,
		ExcludeHosted:        excludeHosted,
		SortByPrice:          sortByPrice,
		Offset:               offset,
		Limit:                limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReceivedGifts), nil
}

// GetReceivedGift Returns information about a received gift @received_gift_id Identifier of the gift
func (c *Client) GetReceivedGift(receivedGiftId string) (*types.ReceivedGift, error) {
	req := &types.GetReceivedGift{
		ReceivedGiftId: receivedGiftId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReceivedGift), nil
}

// GetUpgradedGift Returns information about an upgraded gift by its name @name Unique name of the upgraded gift
func (c *Client) GetUpgradedGift(name string) (*types.UpgradedGift, error) {
	req := &types.GetUpgradedGift{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UpgradedGift), nil
}

// GetUpgradedGiftValueInfo Returns information about value of an upgraded gift by its name @name Unique name of the upgraded gift
func (c *Client) GetUpgradedGiftValueInfo(name string) (*types.UpgradedGiftValueInfo, error) {
	req := &types.GetUpgradedGiftValueInfo{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UpgradedGiftValueInfo), nil
}

// GetUpgradedGiftWithdrawalUrl Returns a URL for upgraded gift withdrawal in the TON blockchain as an NFT; requires owner privileges for gifts owned by a chat
func (c *Client) GetUpgradedGiftWithdrawalUrl(receivedGiftId string, password string) (*types.HttpUrl, error) {
	req := &types.GetUpgradedGiftWithdrawalUrl{
		ReceivedGiftId: receivedGiftId,
		Password:       password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetUpgradedGiftsPromotionalAnimation Returns promotional anumation for upgraded gifts
func (c *Client) GetUpgradedGiftsPromotionalAnimation() (*types.Animation, error) {
	req := &types.GetUpgradedGiftsPromotionalAnimation{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Animation), nil
}

// SetGiftResalePrice Changes resale price of a unique gift owned by the current user
func (c *Client) SetGiftResalePrice(receivedGiftId string, opts *types.SetGiftResalePriceOpts) (*types.Ok, error) {
	req := &types.SetGiftResalePrice{
		ReceivedGiftId: receivedGiftId,
	}
	if opts != nil {
		req.Price = opts.Price
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchGiftsForResale Returns upgraded gifts that can be bought from other owners using sendResoldGift
func (c *Client) SearchGiftsForResale(giftId string, order *types.GiftForResaleOrder, attributes []*types.UpgradedGiftAttributeId, offset string, limit int32) (*types.GiftsForResale, error) {
	req := &types.SearchGiftsForResale{
		GiftId:     giftId,
		Order:      order,
		Attributes: attributes,
		Offset:     offset,
		Limit:      limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftsForResale), nil
}

// GetGiftCollections Returns collections of gifts owned by the given user or chat
func (c *Client) GetGiftCollections(ownerId *types.MessageSender) (*types.GiftCollections, error) {
	req := &types.GetGiftCollections{
		OwnerId: ownerId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollections), nil
}

// CreateGiftCollection Creates a collection from gifts on the current user's or a channel's profile page; requires can_post_messages administrator right in the channel chat.
func (c *Client) CreateGiftCollection(ownerId *types.MessageSender, name string, receivedGiftIds []string) (*types.GiftCollection, error) {
	req := &types.CreateGiftCollection{
		OwnerId:         ownerId,
		Name:            name,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollection), nil
}

// ReorderGiftCollections Changes order of gift collections. If the collections are owned by a channel chat, then requires can_post_messages administrator right in the channel chat
func (c *Client) ReorderGiftCollections(ownerId *types.MessageSender, collectionIds []int32) (*types.Ok, error) {
	req := &types.ReorderGiftCollections{
		OwnerId:       ownerId,
		CollectionIds: collectionIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteGiftCollection Deletes a gift collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat
func (c *Client) DeleteGiftCollection(ownerId *types.MessageSender, collectionId int32) (*types.Ok, error) {
	req := &types.DeleteGiftCollection{
		OwnerId:      ownerId,
		CollectionId: collectionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetGiftCollectionName Changes name of a gift collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) SetGiftCollectionName(ownerId *types.MessageSender, collectionId int32, name string) (*types.GiftCollection, error) {
	req := &types.SetGiftCollectionName{
		OwnerId:      ownerId,
		CollectionId: collectionId,
		Name:         name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollection), nil
}

// AddGiftCollectionGifts Adds gifts to the beginning of a previously created collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) AddGiftCollectionGifts(ownerId *types.MessageSender, collectionId int32, receivedGiftIds []string) (*types.GiftCollection, error) {
	req := &types.AddGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollection), nil
}

// RemoveGiftCollectionGifts Removes gifts from a collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) RemoveGiftCollectionGifts(ownerId *types.MessageSender, collectionId int32, receivedGiftIds []string) (*types.GiftCollection, error) {
	req := &types.RemoveGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollection), nil
}

// ReorderGiftCollectionGifts Changes order of gifts in a collection. If the collection is owned by a channel chat, then requires can_post_messages administrator right in the channel chat. Returns the changed collection
func (c *Client) ReorderGiftCollectionGifts(ownerId *types.MessageSender, collectionId int32, receivedGiftIds []string) (*types.GiftCollection, error) {
	req := &types.ReorderGiftCollectionGifts{
		OwnerId:         ownerId,
		CollectionId:    collectionId,
		ReceivedGiftIds: receivedGiftIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiftCollection), nil
}

// CreateInvoiceLink Creates a link for the given invoice; for bots only
func (c *Client) CreateInvoiceLink(businessConnectionId string, invoice *types.InputMessageContent) (*types.HttpUrl, error) {
	req := &types.CreateInvoiceLink{
		BusinessConnectionId: businessConnectionId,
		Invoice:              invoice,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// RefundStarPayment Refunds a previously done payment in Telegram Stars; for bots only
func (c *Client) RefundStarPayment(userId int64, telegramPaymentChargeId string) (*types.Ok, error) {
	req := &types.RefundStarPayment{
		UserId:                  userId,
		TelegramPaymentChargeId: telegramPaymentChargeId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetSupportUser Returns a user that can be contacted to get support
func (c *Client) GetSupportUser() (*types.User, error) {
	req := &types.GetSupportUser{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.User), nil
}

// GetBackgroundUrl Constructs a persistent HTTP URL for a background @name Background name @type Background type; backgroundTypeChatTheme isn't supported
func (c *Client) GetBackgroundUrl(name string, typeField *types.BackgroundType) (*types.HttpUrl, error) {
	req := &types.GetBackgroundUrl{
		Name:      name,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// SearchBackground Searches for a background by its name @name The name of the background
func (c *Client) SearchBackground(name string) (*types.Background, error) {
	req := &types.SearchBackground{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Background), nil
}

// SetDefaultBackground Sets default background for chats; adds the background to the list of installed backgrounds
func (c *Client) SetDefaultBackground(forDarkTheme bool, opts *types.SetDefaultBackgroundOpts) (*types.Background, error) {
	req := &types.SetDefaultBackground{
		ForDarkTheme: forDarkTheme,
	}
	if opts != nil {
		req.Background = opts.Background
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Background), nil
}

// DeleteDefaultBackground Deletes default background for chats @for_dark_theme Pass true if the background is deleted for a dark theme
func (c *Client) DeleteDefaultBackground(forDarkTheme bool) (*types.Ok, error) {
	req := &types.DeleteDefaultBackground{
		ForDarkTheme: forDarkTheme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetInstalledBackgrounds Returns backgrounds installed by the user @for_dark_theme Pass true to order returned backgrounds for a dark theme
func (c *Client) GetInstalledBackgrounds(forDarkTheme bool) (*types.Backgrounds, error) {
	req := &types.GetInstalledBackgrounds{
		ForDarkTheme: forDarkTheme,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Backgrounds), nil
}

// RemoveInstalledBackground Removes background from the list of installed backgrounds @background_id The background identifier
func (c *Client) RemoveInstalledBackground(backgroundId string) (*types.Ok, error) {
	req := &types.RemoveInstalledBackground{
		BackgroundId: backgroundId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResetInstalledBackgrounds Resets list of installed backgrounds to its default value
func (c *Client) ResetInstalledBackgrounds() (*types.Ok, error) {
	req := &types.ResetInstalledBackgrounds{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline method if only_local is true. Can be called before authorization @only_local Pass true to get only locally available information without sending network requests
func (c *Client) GetLocalizationTargetInfo(onlyLocal bool) (*types.LocalizationTargetInfo, error) {
	req := &types.GetLocalizationTargetInfo{
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LocalizationTargetInfo), nil
}

// GetLanguagePackInfo Returns information about a language pack. Returned language pack identifier may be different from a provided one. Can be called before authorization @language_pack_id Language pack identifier
func (c *Client) GetLanguagePackInfo(languagePackId string) (*types.LanguagePackInfo, error) {
	req := &types.GetLanguagePackInfo{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LanguagePackInfo), nil
}

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
func (c *Client) GetLanguagePackStrings(languagePackId string, keys []string) (*types.LanguagePackStrings, error) {
	req := &types.GetLanguagePackStrings{
		LanguagePackId: languagePackId,
		Keys:           keys,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LanguagePackStrings), nil
}

// SynchronizeLanguagePack Fetches the latest versions of all strings from a language pack in the current localization target from the server.
func (c *Client) SynchronizeLanguagePack(languagePackId string) (*types.Ok, error) {
	req := &types.SynchronizeLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AddCustomServerLanguagePack Adds a custom server language pack to the list of installed language packs in current localization target. Can be called before authorization @language_pack_id Identifier of a language pack to be added
func (c *Client) AddCustomServerLanguagePack(languagePackId string) (*types.Ok, error) {
	req := &types.AddCustomServerLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetCustomLanguagePack Adds or changes a custom local language pack to the current localization target
func (c *Client) SetCustomLanguagePack(info *types.LanguagePackInfo, strings []*types.LanguagePackString) (*types.Ok, error) {
	req := &types.SetCustomLanguagePack{
		Info:    info,
		Strings: strings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditCustomLanguagePackInfo Edits information about a custom local language pack in the current localization target. Can be called before authorization @info New information about the custom local language pack
func (c *Client) EditCustomLanguagePackInfo(info *types.LanguagePackInfo) (*types.Ok, error) {
	req := &types.EditCustomLanguagePackInfo{
		Info: info,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom local language pack. Can be called before authorization @language_pack_id Identifier of a previously added custom local language pack in the current localization target @new_string New language pack string
func (c *Client) SetCustomLanguagePackString(languagePackId string, newString *types.LanguagePackString) (*types.Ok, error) {
	req := &types.SetCustomLanguagePackString{
		LanguagePackId: languagePackId,
		NewString:      newString,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack which is currently in use (including base language pack) or is being synchronized can't be deleted.
func (c *Client) DeleteLanguagePack(languagePackId string) (*types.Ok, error) {
	req := &types.DeleteLanguagePack{
		LanguagePackId: languagePackId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription @device_token Device token @other_user_ids List of user identifiers of other users currently using the application
func (c *Client) RegisterDevice(deviceToken *types.DeviceToken, otherUserIds []int64) (*types.PushReceiverId, error) {
	req := &types.RegisterDevice{
		DeviceToken:  deviceToken,
		OtherUserIds: otherUserIds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PushReceiverId), nil
}

// ProcessPushNotification Handles a push notification. Returns error with code 406 if the push notification is not supported and connection to the server is required to fetch new data. Can be called before authorization
func (c *Client) ProcessPushNotification(payload string) (*types.Ok, error) {
	req := &types.ProcessPushNotification{
		Payload: payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPushReceiverId Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously @payload JSON-encoded push notification payload
func (c *Client) GetPushReceiverId(payload string) (*types.PushReceiverId, error) {
	req := &types.GetPushReceiverId{
		Payload: payload,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PushReceiverId), nil
}

// GetRecentlyVisitedTMeUrls Returns t.me URLs recently visited by a newly registered user @referrer Google Play referrer to identify the user
func (c *Client) GetRecentlyVisitedTMeUrls(referrer string) (*types.TMeUrls, error) {
	req := &types.GetRecentlyVisitedTMeUrls{
		Referrer: referrer,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TMeUrls), nil
}

// SetUserPrivacySettingRules Changes user privacy settings @setting The privacy setting @rules The new privacy rules
func (c *Client) SetUserPrivacySettingRules(setting *types.UserPrivacySetting, rules *types.UserPrivacySettingRules) (*types.Ok, error) {
	req := &types.SetUserPrivacySettingRules{
		Setting: setting,
		Rules:   rules,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetUserPrivacySettingRules Returns the current privacy settings @setting The privacy setting
func (c *Client) GetUserPrivacySettingRules(setting *types.UserPrivacySetting) (*types.UserPrivacySettingRules, error) {
	req := &types.GetUserPrivacySettingRules{
		Setting: setting,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UserPrivacySettingRules), nil
}

// SetReadDatePrivacySettings Changes privacy settings for message read date @settings New settings
func (c *Client) SetReadDatePrivacySettings(settings *types.ReadDatePrivacySettings) (*types.Ok, error) {
	req := &types.SetReadDatePrivacySettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetReadDatePrivacySettings Returns privacy settings for message read date
func (c *Client) GetReadDatePrivacySettings() (*types.ReadDatePrivacySettings, error) {
	req := &types.GetReadDatePrivacySettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReadDatePrivacySettings), nil
}

// SetNewChatPrivacySettings Changes privacy settings for new chat creation; can be used only if getOption("can_set_new_chat_privacy_settings") @settings New settings
func (c *Client) SetNewChatPrivacySettings(settings *types.NewChatPrivacySettings) (*types.Ok, error) {
	req := &types.SetNewChatPrivacySettings{
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetNewChatPrivacySettings Returns privacy settings for new chat creation
func (c *Client) GetNewChatPrivacySettings() (*types.NewChatPrivacySettings, error) {
	req := &types.GetNewChatPrivacySettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.NewChatPrivacySettings), nil
}

// GetPaidMessageRevenue Returns the total number of Telegram Stars received by the current user for paid messages from the given user @user_id Identifier of the user
func (c *Client) GetPaidMessageRevenue(userId int64) (*types.StarCount, error) {
	req := &types.GetPaidMessageRevenue{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarCount), nil
}

// AllowUnpaidMessagesFromUser Allows the specified user to send unpaid private messages to the current user by adding a rule to userPrivacySettingAllowUnpaidMessages
func (c *Client) AllowUnpaidMessagesFromUser(userId int64, refundPayments bool) (*types.Ok, error) {
	req := &types.AllowUnpaidMessagesFromUser{
		UserId:         userId,
		RefundPayments: refundPayments,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatPaidMessageStarCount Changes the amount of Telegram Stars that must be paid to send a message to a supergroup chat; requires can_restrict_members administrator right and supergroupFullInfo.can_enable_paid_messages
func (c *Client) SetChatPaidMessageStarCount(chatId int64, paidMessageStarCount int64) (*types.Ok, error) {
	req := &types.SetChatPaidMessageStarCount{
		ChatId:               chatId,
		PaidMessageStarCount: paidMessageStarCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// CanSendMessageToUser Checks whether the current user can message another user or try to create a chat with them
func (c *Client) CanSendMessageToUser(userId int64, onlyLocal bool) (*types.CanSendMessageToUserResult, error) {
	req := &types.CanSendMessageToUser{
		UserId:    userId,
		OnlyLocal: onlyLocal,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CanSendMessageToUserResult), nil
}

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization. Can be called synchronously for options "version" and "commit_hash"
func (c *Client) GetOption(name string) (*types.OptionValue, error) {
	req := &types.GetOption{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.OptionValue), nil
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
func (c *Client) SetOption(name string, opts *types.SetOptionOpts) (*types.Ok, error) {
	req := &types.SetOption{
		Name: name,
	}
	if opts != nil {
		req.Value = opts.Value
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAccountTtl Changes the period of inactivity after which the account of the current user will automatically be deleted @ttl New account TTL
func (c *Client) SetAccountTtl(ttl *types.AccountTtl) (*types.Ok, error) {
	req := &types.SetAccountTtl{
		Ttl: ttl,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetAccountTtl Returns the period of inactivity after which the account of the current user will automatically be deleted
func (c *Client) GetAccountTtl() (*types.AccountTtl, error) {
	req := &types.GetAccountTtl{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AccountTtl), nil
}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account.
func (c *Client) DeleteAccount(reason string, password string) (*types.Ok, error) {
	req := &types.DeleteAccount{
		Reason:   reason,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetDefaultMessageAutoDeleteTime Changes the default message auto-delete time for new chats @message_auto_delete_time New default message auto-delete time; must be from 0 up to 365 * 86400 and be divisible by 86400. If 0, then messages aren't deleted automatically
func (c *Client) SetDefaultMessageAutoDeleteTime(messageAutoDeleteTime *types.MessageAutoDeleteTime) (*types.Ok, error) {
	req := &types.SetDefaultMessageAutoDeleteTime{
		MessageAutoDeleteTime: messageAutoDeleteTime,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetDefaultMessageAutoDeleteTime Returns default message auto-delete time setting for new chats
func (c *Client) GetDefaultMessageAutoDeleteTime() (*types.MessageAutoDeleteTime, error) {
	req := &types.GetDefaultMessageAutoDeleteTime{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageAutoDeleteTime), nil
}

// RemoveChatActionBar Removes a chat action bar without any other action @chat_id Chat identifier
func (c *Client) RemoveChatActionBar(chatId int64) (*types.Ok, error) {
	req := &types.RemoveChatActionBar{
		ChatId: chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportChat Reports a chat to the Telegram moderators. A chat can be reported only from the chat action bar, or if chat.can_be_reported
func (c *Client) ReportChat(chatId int64, optionId string, messageIds []int64, text string) (*types.ReportChatResult, error) {
	req := &types.ReportChat{
		ChatId:     chatId,
		OptionId:   optionId,
		MessageIds: messageIds,
		Text:       text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ReportChatResult), nil
}

// ReportChatPhoto Reports a chat photo to the Telegram moderators. A chat photo can be reported only if chat.can_be_reported
func (c *Client) ReportChatPhoto(chatId int64, fileId int32, reason *types.ReportReason, text string) (*types.Ok, error) {
	req := &types.ReportChatPhoto{
		ChatId: chatId,
		FileId: fileId,
		Reason: reason,
		Text:   text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReportMessageReactions Reports reactions set on a message to the Telegram moderators. Reactions on a message can be reported only if messageProperties.can_report_reactions
func (c *Client) ReportMessageReactions(chatId int64, messageId int64, senderId *types.MessageSender) (*types.Ok, error) {
	req := &types.ReportMessageReactions{
		ChatId:    chatId,
		MessageId: messageId,
		SenderId:  senderId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetChatRevenueStatistics Returns detailed revenue statistics about a chat. Currently, this method can be used only
func (c *Client) GetChatRevenueStatistics(chatId int64, isDark bool) (*types.ChatRevenueStatistics, error) {
	req := &types.GetChatRevenueStatistics{
		ChatId: chatId,
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatRevenueStatistics), nil
}

// GetChatRevenueWithdrawalUrl Returns a URL for chat revenue withdrawal; requires owner privileges in the channel chat or the bot. Currently, this method can be used only
func (c *Client) GetChatRevenueWithdrawalUrl(chatId int64, password string) (*types.HttpUrl, error) {
	req := &types.GetChatRevenueWithdrawalUrl{
		ChatId:   chatId,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetChatRevenueTransactions Returns the list of revenue transactions for a chat. Currently, this method can be used only
func (c *Client) GetChatRevenueTransactions(chatId int64, offset string, limit int32) (*types.ChatRevenueTransactions, error) {
	req := &types.GetChatRevenueTransactions{
		ChatId: chatId,
		Offset: offset,
		Limit:  limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatRevenueTransactions), nil
}

// GetTonTransactions Returns the list of Toncoin transactions of the current user
func (c *Client) GetTonTransactions(offset string, limit int32, opts *types.GetTonTransactionsOpts) (*types.TonTransactions, error) {
	req := &types.GetTonTransactions{
		Offset: offset,
		Limit:  limit,
	}
	if opts != nil {
		req.Direction = opts.Direction
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TonTransactions), nil
}

// GetStarRevenueStatistics Returns detailed Telegram Star revenue statistics
func (c *Client) GetStarRevenueStatistics(ownerId *types.MessageSender, isDark bool) (*types.StarRevenueStatistics, error) {
	req := &types.GetStarRevenueStatistics{
		OwnerId: ownerId,
		IsDark:  isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarRevenueStatistics), nil
}

// GetStarWithdrawalUrl Returns a URL for Telegram Star withdrawal
func (c *Client) GetStarWithdrawalUrl(ownerId *types.MessageSender, starCount int64, password string) (*types.HttpUrl, error) {
	req := &types.GetStarWithdrawalUrl{
		OwnerId:   ownerId,
		StarCount: starCount,
		Password:  password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetStarAdAccountUrl Returns a URL for a Telegram Ad platform account that can be used to set up advertisements for the chat paid in the owned Telegram Stars
func (c *Client) GetStarAdAccountUrl(ownerId *types.MessageSender) (*types.HttpUrl, error) {
	req := &types.GetStarAdAccountUrl{
		OwnerId: ownerId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetTonRevenueStatistics Returns detailed Toncoin revenue statistics of the current user @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetTonRevenueStatistics(isDark bool) (*types.TonRevenueStatistics, error) {
	req := &types.GetTonRevenueStatistics{
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TonRevenueStatistics), nil
}

// GetTonWithdrawalUrl Returns a URL for Toncoin withdrawal from the current user's account. The user must have at least 10 toncoins to withdraw
func (c *Client) GetTonWithdrawalUrl(password string) (*types.HttpUrl, error) {
	req := &types.GetTonWithdrawalUrl{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// GetChatStatistics Returns detailed statistics about a chat. Currently, this method can be used only for supergroups and channels. Can be used only if supergroupFullInfo.can_get_statistics == true @chat_id Chat identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetChatStatistics(chatId int64, isDark bool) (*types.ChatStatistics, error) {
	req := &types.GetChatStatistics{
		ChatId: chatId,
		IsDark: isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ChatStatistics), nil
}

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if messageProperties.can_get_statistics == true @chat_id Chat identifier @message_id Message identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetMessageStatistics(chatId int64, messageId int64, isDark bool) (*types.MessageStatistics, error) {
	req := &types.GetMessageStatistics{
		ChatId:    chatId,
		MessageId: messageId,
		IsDark:    isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.MessageStatistics), nil
}

// GetMessagePublicForwards Returns forwarded copies of a channel message to different public channels and public reposts as a story. Can be used only if messageProperties.can_get_statistics == true. For optimal performance, the number of returned messages and stories is chosen by TDLib
func (c *Client) GetMessagePublicForwards(chatId int64, messageId int64, offset string, limit int32) (*types.PublicForwards, error) {
	req := &types.GetMessagePublicForwards{
		ChatId:    chatId,
		MessageId: messageId,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PublicForwards), nil
}

// GetStoryStatistics Returns detailed statistics about a story. Can be used only if story.can_get_statistics == true @chat_id Chat identifier @story_id Story identifier @is_dark Pass true if a dark theme is used by the application
func (c *Client) GetStoryStatistics(chatId int64, storyId int32, isDark bool) (*types.StoryStatistics, error) {
	req := &types.GetStoryStatistics{
		ChatId:  chatId,
		StoryId: storyId,
		IsDark:  isDark,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StoryStatistics), nil
}

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph @chat_id Chat identifier @token The token for graph loading @x X-value for zoomed in graph or 0 otherwise
func (c *Client) GetStatisticalGraph(chatId int64, token string, x int64) (*types.StatisticalGraph, error) {
	req := &types.GetStatisticalGraph{
		ChatId: chatId,
		Token:  token,
		X:      x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StatisticalGraph), nil
}

// GetStorageStatistics Returns storage usage statistics. Can be called before authorization
func (c *Client) GetStorageStatistics(chatLimit int32) (*types.StorageStatistics, error) {
	req := &types.GetStorageStatistics{
		ChatLimit: chatLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StorageStatistics), nil
}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics. Can be called before authorization
func (c *Client) GetStorageStatisticsFast() (*types.StorageStatisticsFast, error) {
	req := &types.GetStorageStatisticsFast{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StorageStatisticsFast), nil
}

// GetDatabaseStatistics Returns database statistics
func (c *Client) GetDatabaseStatistics() (*types.DatabaseStatistics, error) {
	req := &types.GetDatabaseStatistics{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.DatabaseStatistics), nil
}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
func (c *Client) OptimizeStorage(size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*types.FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*types.StorageStatistics, error) {
	req := &types.OptimizeStorage{
		Size:                        size,
		Ttl:                         ttl,
		Count:                       count,
		ImmunityDelay:               immunityDelay,
		FileTypes:                   fileTypes,
		ChatIds:                     chatIds,
		ExcludeChatIds:              excludeChatIds,
		ReturnDeletedFileStatistics: returnDeletedFileStatistics,
		ChatLimit:                   chatLimit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StorageStatistics), nil
}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks,
func (c *Client) SetNetworkType(opts *types.SetNetworkTypeOpts) (*types.Ok, error) {
	req := &types.SetNetworkType{}
	if opts != nil {
		req.TypeField = opts.TypeField
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization @only_current Pass true to get statistics only for the current library launch
func (c *Client) GetNetworkStatistics(onlyCurrent bool) (*types.NetworkStatistics, error) {
	req := &types.GetNetworkStatistics{
		OnlyCurrent: onlyCurrent,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.NetworkStatistics), nil
}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization @entry The network statistics entry with the data to be added to statistics
func (c *Client) AddNetworkStatistics(entry *types.NetworkStatisticsEntry) (*types.Ok, error) {
	req := &types.AddNetworkStatistics{
		Entry: entry,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (c *Client) ResetNetworkStatistics() (*types.Ok, error) {
	req := &types.ResetNetworkStatistics{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetAutoDownloadSettingsPresets Returns auto-download settings presets for the current user
func (c *Client) GetAutoDownloadSettingsPresets() (*types.AutoDownloadSettingsPresets, error) {
	req := &types.GetAutoDownloadSettingsPresets{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AutoDownloadSettingsPresets), nil
}

// SetAutoDownloadSettings Sets auto-download settings @settings New user auto-download settings @type Type of the network for which the new settings are relevant
func (c *Client) SetAutoDownloadSettings(settings *types.AutoDownloadSettings, typeField *types.NetworkType) (*types.Ok, error) {
	req := &types.SetAutoDownloadSettings{
		Settings:  settings,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetAutosaveSettings Returns autosave settings for the current user
func (c *Client) GetAutosaveSettings() (*types.AutosaveSettings, error) {
	req := &types.GetAutosaveSettings{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.AutosaveSettings), nil
}

// SetAutosaveSettings Sets autosave settings for the given scope. The method is guaranteed to work only after at least one call to getAutosaveSettings @scope Autosave settings scope @settings New autosave settings for the scope; pass null to set autosave settings to default
func (c *Client) SetAutosaveSettings(scope *types.AutosaveSettingsScope, settings *types.ScopeAutosaveSettings) (*types.Ok, error) {
	req := &types.SetAutosaveSettings{
		Scope:    scope,
		Settings: settings,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClearAutosaveSettingsExceptions Clears the list of all autosave settings exceptions. The method is guaranteed to work only after at least one call to getAutosaveSettings
func (c *Client) ClearAutosaveSettingsExceptions() (*types.Ok, error) {
	req := &types.ClearAutosaveSettingsExceptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetBankCardInfo Returns information about a bank card @bank_card_number The bank card number
func (c *Client) GetBankCardInfo(bankCardNumber string) (*types.BankCardInfo, error) {
	req := &types.GetBankCardInfo{
		BankCardNumber: bankCardNumber,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BankCardInfo), nil
}

// GetPassportElement Returns one of the available Telegram Passport elements @type Telegram Passport element type @password The 2-step verification password of the current user
func (c *Client) GetPassportElement(typeField *types.PassportElementType, password string) (*types.PassportElement, error) {
	req := &types.GetPassportElement{
		TypeField: typeField,
		Password:  password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PassportElement), nil
}

// GetAllPassportElements Returns all available Telegram Passport elements @password The 2-step verification password of the current user
func (c *Client) GetAllPassportElements(password string) (*types.PassportElements, error) {
	req := &types.GetAllPassportElements{
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PassportElements), nil
}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
func (c *Client) SetPassportElement(element *types.InputPassportElement, password string) (*types.PassportElement, error) {
	req := &types.SetPassportElement{
		Element:  element,
		Password: password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PassportElement), nil
}

// DeletePassportElement Deletes a Telegram Passport element @type Element type
func (c *Client) DeletePassportElement(typeField *types.PassportElementType) (*types.Ok, error) {
	req := &types.DeletePassportElement{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed @user_id User identifier @errors The errors
func (c *Client) SetPassportElementErrors(userId int64, errors []*types.InputPassportElementError) (*types.Ok, error) {
	req := &types.SetPassportElementErrors{
		UserId: userId,
		Errors: errors,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which must be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown @country_code A two-letter ISO 3166-1 alpha-2 country code
func (c *Client) GetPreferredCountryLanguage(countryCode string) (*types.Text, error) {
	req := &types.GetPreferredCountryLanguage{
		CountryCode: countryCode,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport @email_address Email address
func (c *Client) SendEmailAddressVerificationCode(emailAddress string) (*types.EmailAddressAuthenticationCodeInfo, error) {
	req := &types.SendEmailAddressVerificationCode{
		EmailAddress: emailAddress,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmailAddressAuthenticationCodeInfo), nil
}

// ResendEmailAddressVerificationCode Resends the code to verify an email address to be added to a user's Telegram Passport
func (c *Client) ResendEmailAddressVerificationCode() (*types.EmailAddressAuthenticationCodeInfo, error) {
	req := &types.ResendEmailAddressVerificationCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.EmailAddressAuthenticationCodeInfo), nil
}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport @code Verification code to check
func (c *Client) CheckEmailAddressVerificationCode(code string) (*types.Ok, error) {
	req := &types.CheckEmailAddressVerificationCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPassportAuthorizationForm Returns a Telegram Passport authorization form for sharing data with a service
func (c *Client) GetPassportAuthorizationForm(botUserId int64, scope string, publicKey string, nonce string) (*types.PassportAuthorizationForm, error) {
	req := &types.GetPassportAuthorizationForm{
		BotUserId: botUserId,
		Scope:     scope,
		PublicKey: publicKey,
		Nonce:     nonce,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PassportAuthorizationForm), nil
}

// GetPassportAuthorizationFormAvailableElements Returns already available Telegram Passport elements suitable for completing a Telegram Passport authorization form. Result can be received only once for each authorization form
func (c *Client) GetPassportAuthorizationFormAvailableElements(authorizationFormId int32, password string) (*types.PassportElementsWithErrors, error) {
	req := &types.GetPassportAuthorizationFormAvailableElements{
		AuthorizationFormId: authorizationFormId,
		Password:            password,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PassportElementsWithErrors), nil
}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service. This method must be called after getPassportAuthorizationFormAvailableElements if some previously available elements are going to be reused
func (c *Client) SendPassportAuthorizationForm(authorizationFormId int32, ty []*types.PassportElementType) (*types.Ok, error) {
	req := &types.SendPassportAuthorizationForm{
		AuthorizationFormId: authorizationFormId,
		Types:               ty,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only @pending_update_count The number of pending updates @error_message The last error message
func (c *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*types.Ok, error) {
	req := &types.SetBotUpdatesStatus{
		PendingUpdateCount: pendingUpdateCount,
		ErrorMessage:       errorMessage,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// UploadStickerFile Uploads a file with a sticker; returns the uploaded file
func (c *Client) UploadStickerFile(userId int64, stickerFormat *types.StickerFormat, sticker *types.InputFile) (*types.File, error) {
	req := &types.UploadStickerFile{
		UserId:        userId,
		StickerFormat: stickerFormat,
		Sticker:       sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// GetSuggestedStickerSetName Returns a suggested name for a new sticker set with a given title @title Sticker set title; 1-64 characters
func (c *Client) GetSuggestedStickerSetName(title string) (*types.Text, error) {
	req := &types.GetSuggestedStickerSetName{
		Title: title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// CheckStickerSetName Checks whether a name can be used for a new sticker set @name Name to be checked
func (c *Client) CheckStickerSetName(name string) (*types.CheckStickerSetNameResult, error) {
	req := &types.CheckStickerSetName{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CheckStickerSetNameResult), nil
}

// CreateNewStickerSet Creates a new sticker set. Returns the newly created sticker set
func (c *Client) CreateNewStickerSet(userId int64, title string, name string, stickerType *types.StickerType, needsRepainting bool, stickers []*types.InputSticker, source string) (*types.StickerSet, error) {
	req := &types.CreateNewStickerSet{
		UserId:          userId,
		Title:           title,
		Name:            name,
		StickerType:     stickerType,
		NeedsRepainting: needsRepainting,
		Stickers:        stickers,
		Source:          source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSet), nil
}

// AddStickerToSet Adds a new sticker to a set
func (c *Client) AddStickerToSet(userId int64, name string, sticker *types.InputSticker) (*types.Ok, error) {
	req := &types.AddStickerToSet{
		UserId:  userId,
		Name:    name,
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReplaceStickerInSet Replaces existing sticker in a set. The function is equivalent to removeStickerFromSet, then addStickerToSet, then setStickerPositionInSet
func (c *Client) ReplaceStickerInSet(userId int64, name string, oldSticker *types.InputFile, newSticker *types.InputSticker) (*types.Ok, error) {
	req := &types.ReplaceStickerInSet{
		UserId:     userId,
		Name:       name,
		OldSticker: oldSticker,
		NewSticker: newSticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerSetThumbnail Sets a sticker set thumbnail
func (c *Client) SetStickerSetThumbnail(userId int64, name string, opts *types.SetStickerSetThumbnailOpts) (*types.Ok, error) {
	req := &types.SetStickerSetThumbnail{
		UserId: userId,
		Name:   name,
	}
	if opts != nil {
		req.Thumbnail = opts.Thumbnail
		req.Format = opts.Format
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetCustomEmojiStickerSetThumbnail Sets a custom emoji sticker set thumbnail
func (c *Client) SetCustomEmojiStickerSetThumbnail(name string, customEmojiId string) (*types.Ok, error) {
	req := &types.SetCustomEmojiStickerSetThumbnail{
		Name:          name,
		CustomEmojiId: customEmojiId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerSetTitle Sets a sticker set title @name Sticker set name. The sticker set must be owned by the current user @title New sticker set title
func (c *Client) SetStickerSetTitle(name string, title string) (*types.Ok, error) {
	req := &types.SetStickerSetTitle{
		Name:  name,
		Title: title,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DeleteStickerSet Completely deletes a sticker set @name Sticker set name. The sticker set must be owned by the current user
func (c *Client) DeleteStickerSet(name string) (*types.Ok, error) {
	req := &types.DeleteStickerSet{
		Name: name,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs. The sticker set must be owned by the current user
func (c *Client) SetStickerPositionInSet(sticker *types.InputFile, position int32) (*types.Ok, error) {
	req := &types.SetStickerPositionInSet{
		Sticker:  sticker,
		Position: position,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs. The sticker set must be owned by the current user @sticker Sticker to remove from the set
func (c *Client) RemoveStickerFromSet(sticker *types.InputFile) (*types.Ok, error) {
	req := &types.RemoveStickerFromSet{
		Sticker: sticker,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerEmojis Changes the list of emojis corresponding to a sticker. The sticker must belong to a regular or custom emoji sticker set that is owned by the current user
func (c *Client) SetStickerEmojis(sticker *types.InputFile, emojis string) (*types.Ok, error) {
	req := &types.SetStickerEmojis{
		Sticker: sticker,
		Emojis:  emojis,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerKeywords Changes the list of keywords of a sticker. The sticker must belong to a regular or custom emoji sticker set that is owned by the current user
func (c *Client) SetStickerKeywords(sticker *types.InputFile, keywords []string) (*types.Ok, error) {
	req := &types.SetStickerKeywords{
		Sticker:  sticker,
		Keywords: keywords,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetStickerMaskPosition Changes the mask position of a mask sticker. The sticker must belong to a mask sticker set that is owned by the current user
func (c *Client) SetStickerMaskPosition(sticker *types.InputFile, opts *types.SetStickerMaskPositionOpts) (*types.Ok, error) {
	req := &types.SetStickerMaskPosition{
		Sticker: sticker,
	}
	if opts != nil {
		req.MaskPosition = opts.MaskPosition
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetOwnedStickerSets Returns sticker sets owned by the current user
func (c *Client) GetOwnedStickerSets(offsetStickerSetId string, limit int32) (*types.StickerSets, error) {
	req := &types.GetOwnedStickerSets{
		OffsetStickerSetId: offsetStickerSetId,
		Limit:              limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StickerSets), nil
}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
func (c *Client) GetMapThumbnailFile(location *types.Location, zoom int32, width int32, height int32, scale int32, chatId int64) (*types.File, error) {
	req := &types.GetMapThumbnailFile{
		Location: location,
		Zoom:     zoom,
		Width:    width,
		Height:   height,
		Scale:    scale,
		ChatId:   chatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.File), nil
}

// GetPremiumLimit Returns information about a limit, increased for Premium users. Returns a 404 error if the limit is unknown @limit_type Type of the limit
func (c *Client) GetPremiumLimit(limitType *types.PremiumLimitType) (*types.PremiumLimit, error) {
	req := &types.GetPremiumLimit{
		LimitType: limitType,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumLimit), nil
}

// GetPremiumFeatures Returns information about features, available to Premium users @source Source of the request; pass null if the method is called from some non-standard source
func (c *Client) GetPremiumFeatures(source *types.PremiumSource) (*types.PremiumFeatures, error) {
	req := &types.GetPremiumFeatures{
		Source: source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumFeatures), nil
}

// GetPremiumStickerExamples Returns examples of premium stickers for demonstration purposes
func (c *Client) GetPremiumStickerExamples() (*types.Stickers, error) {
	req := &types.GetPremiumStickerExamples{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Stickers), nil
}

// GetPremiumInfoSticker Returns the sticker to be used as representation of the Telegram Premium subscription @month_count Number of months the Telegram Premium subscription will be active
func (c *Client) GetPremiumInfoSticker(monthCount int32) (*types.Sticker, error) {
	req := &types.GetPremiumInfoSticker{
		MonthCount: monthCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Sticker), nil
}

// ViewPremiumFeature Informs TDLib that the user viewed detailed information about a Premium feature on the Premium features screen @feature The viewed premium feature
func (c *Client) ViewPremiumFeature(feature *types.PremiumFeature) (*types.Ok, error) {
	req := &types.ViewPremiumFeature{
		Feature: feature,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ClickPremiumSubscriptionButton Informs TDLib that the user clicked Premium subscription button on the Premium features screen
func (c *Client) ClickPremiumSubscriptionButton() (*types.Ok, error) {
	req := &types.ClickPremiumSubscriptionButton{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetPremiumState Returns state of Telegram Premium subscription and promotion videos for Premium features
func (c *Client) GetPremiumState() (*types.PremiumState, error) {
	req := &types.GetPremiumState{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumState), nil
}

// GetPremiumGiftPaymentOptions Returns available options for gifting Telegram Premium to a user
func (c *Client) GetPremiumGiftPaymentOptions() (*types.PremiumGiftPaymentOptions, error) {
	req := &types.GetPremiumGiftPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumGiftPaymentOptions), nil
}

// GetPremiumGiveawayPaymentOptions Returns available options for creating of Telegram Premium giveaway or manual distribution of Telegram Premium among chat members
func (c *Client) GetPremiumGiveawayPaymentOptions(boostedChatId int64) (*types.PremiumGiveawayPaymentOptions, error) {
	req := &types.GetPremiumGiveawayPaymentOptions{
		BoostedChatId: boostedChatId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumGiveawayPaymentOptions), nil
}

// CheckPremiumGiftCode Returns information about a Telegram Premium gift code @code The code to check
func (c *Client) CheckPremiumGiftCode(code string) (*types.PremiumGiftCodeInfo, error) {
	req := &types.CheckPremiumGiftCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PremiumGiftCodeInfo), nil
}

// ApplyPremiumGiftCode Applies a Telegram Premium gift code @code The code to apply
func (c *Client) ApplyPremiumGiftCode(code string) (*types.Ok, error) {
	req := &types.ApplyPremiumGiftCode{
		Code: code,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GiftPremiumWithStars Allows to buy a Telegram Premium subscription for another user with payment in Telegram Stars; for bots only
func (c *Client) GiftPremiumWithStars(userId int64, starCount int64, monthCount int32, text *types.FormattedText) (*types.Ok, error) {
	req := &types.GiftPremiumWithStars{
		UserId:     userId,
		StarCount:  starCount,
		MonthCount: monthCount,
		Text:       text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// LaunchPrepaidGiveaway Launches a prepaid giveaway
func (c *Client) LaunchPrepaidGiveaway(giveawayId string, parameters *types.GiveawayParameters, winnerCount int32, starCount int64) (*types.Ok, error) {
	req := &types.LaunchPrepaidGiveaway{
		GiveawayId:  giveawayId,
		Parameters:  parameters,
		WinnerCount: winnerCount,
		StarCount:   starCount,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetGiveawayInfo Returns information about a giveaway
func (c *Client) GetGiveawayInfo(chatId int64, messageId int64) (*types.GiveawayInfo, error) {
	req := &types.GetGiveawayInfo{
		ChatId:    chatId,
		MessageId: messageId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.GiveawayInfo), nil
}

// GetStarPaymentOptions Returns available options for Telegram Stars purchase
func (c *Client) GetStarPaymentOptions() (*types.StarPaymentOptions, error) {
	req := &types.GetStarPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarPaymentOptions), nil
}

// GetStarGiftPaymentOptions Returns available options for Telegram Stars gifting @user_id Identifier of the user that will receive Telegram Stars; pass 0 to get options for an unspecified user
func (c *Client) GetStarGiftPaymentOptions(userId int64) (*types.StarPaymentOptions, error) {
	req := &types.GetStarGiftPaymentOptions{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarPaymentOptions), nil
}

// GetStarGiveawayPaymentOptions Returns available options for Telegram Star giveaway creation
func (c *Client) GetStarGiveawayPaymentOptions() (*types.StarGiveawayPaymentOptions, error) {
	req := &types.GetStarGiveawayPaymentOptions{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarGiveawayPaymentOptions), nil
}

// GetStarTransactions Returns the list of Telegram Star transactions for the specified owner
func (c *Client) GetStarTransactions(ownerId *types.MessageSender, subscriptionId string, offset string, limit int32, opts *types.GetStarTransactionsOpts) (*types.StarTransactions, error) {
	req := &types.GetStarTransactions{
		OwnerId:        ownerId,
		SubscriptionId: subscriptionId,
		Offset:         offset,
		Limit:          limit,
	}
	if opts != nil {
		req.Direction = opts.Direction
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarTransactions), nil
}

// GetStarSubscriptions Returns the list of Telegram Star subscriptions for the current user
func (c *Client) GetStarSubscriptions(onlyExpiring bool, offset string) (*types.StarSubscriptions, error) {
	req := &types.GetStarSubscriptions{
		OnlyExpiring: onlyExpiring,
		Offset:       offset,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.StarSubscriptions), nil
}

// CanPurchaseFromStore Checks whether an in-store purchase is possible. Must be called before any in-store purchase. For official applications only @purpose Transaction purpose
func (c *Client) CanPurchaseFromStore(purpose *types.StorePaymentPurpose) (*types.Ok, error) {
	req := &types.CanPurchaseFromStore{
		Purpose: purpose,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// AssignStoreTransaction Informs server about an in-store purchase. For official applications only @transaction Information about the transaction @purpose Transaction purpose
func (c *Client) AssignStoreTransaction(transaction *types.StoreTransaction, purpose *types.StorePaymentPurpose) (*types.Ok, error) {
	req := &types.AssignStoreTransaction{
		Transaction: transaction,
		Purpose:     purpose,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditStarSubscription Cancels or re-enables Telegram Star subscription
func (c *Client) EditStarSubscription(subscriptionId string, isCanceled bool) (*types.Ok, error) {
	req := &types.EditStarSubscription{
		SubscriptionId: subscriptionId,
		IsCanceled:     isCanceled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// EditUserStarSubscription Cancels or re-enables Telegram Star subscription for a user; for bots only
func (c *Client) EditUserStarSubscription(userId int64, telegramPaymentChargeId string, isCanceled bool) (*types.Ok, error) {
	req := &types.EditUserStarSubscription{
		UserId:                  userId,
		TelegramPaymentChargeId: telegramPaymentChargeId,
		IsCanceled:              isCanceled,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// ReuseStarSubscription Reuses an active Telegram Star subscription to a channel chat and joins the chat again @subscription_id Identifier of the subscription
func (c *Client) ReuseStarSubscription(subscriptionId string) (*types.Ok, error) {
	req := &types.ReuseStarSubscription{
		SubscriptionId: subscriptionId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetChatAffiliateProgram Changes affiliate program for a bot
func (c *Client) SetChatAffiliateProgram(chatId int64, opts *types.SetChatAffiliateProgramOpts) (*types.Ok, error) {
	req := &types.SetChatAffiliateProgram{
		ChatId: chatId,
	}
	if opts != nil {
		req.Parameters = opts.Parameters
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchChatAffiliateProgram Searches a chat with an affiliate program. Returns the chat if found and the program is active
func (c *Client) SearchChatAffiliateProgram(username string, referrer string) (*types.Chat, error) {
	req := &types.SearchChatAffiliateProgram{
		Username: username,
		Referrer: referrer,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Chat), nil
}

// SearchAffiliatePrograms Searches affiliate programs that can be connected to the given affiliate
func (c *Client) SearchAffiliatePrograms(affiliate *types.AffiliateType, sortOrder *types.AffiliateProgramSortOrder, offset string, limit int32) (*types.FoundAffiliatePrograms, error) {
	req := &types.SearchAffiliatePrograms{
		Affiliate: affiliate,
		SortOrder: sortOrder,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundAffiliatePrograms), nil
}

// ConnectAffiliateProgram Connects an affiliate program to the given affiliate. Returns information about the connected affiliate program
func (c *Client) ConnectAffiliateProgram(affiliate *types.AffiliateType, botUserId int64) (*types.ConnectedAffiliateProgram, error) {
	req := &types.ConnectAffiliateProgram{
		Affiliate: affiliate,
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ConnectedAffiliateProgram), nil
}

// DisconnectAffiliateProgram Disconnects an affiliate program from the given affiliate and immediately deactivates its referral link. Returns updated information about the disconnected affiliate program
func (c *Client) DisconnectAffiliateProgram(affiliate *types.AffiliateType, url string) (*types.ConnectedAffiliateProgram, error) {
	req := &types.DisconnectAffiliateProgram{
		Affiliate: affiliate,
		Url:       url,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ConnectedAffiliateProgram), nil
}

// GetConnectedAffiliateProgram Returns an affiliate program that were connected to the given affiliate by identifier of the bot that created the program
func (c *Client) GetConnectedAffiliateProgram(affiliate *types.AffiliateType, botUserId int64) (*types.ConnectedAffiliateProgram, error) {
	req := &types.GetConnectedAffiliateProgram{
		Affiliate: affiliate,
		BotUserId: botUserId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ConnectedAffiliateProgram), nil
}

// GetConnectedAffiliatePrograms Returns affiliate programs that were connected to the given affiliate
func (c *Client) GetConnectedAffiliatePrograms(affiliate *types.AffiliateType, offset string, limit int32) (*types.ConnectedAffiliatePrograms, error) {
	req := &types.GetConnectedAffiliatePrograms{
		Affiliate: affiliate,
		Offset:    offset,
		Limit:     limit,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.ConnectedAffiliatePrograms), nil
}

// GetBusinessFeatures Returns information about features, available to Business users @source Source of the request; pass null if the method is called from settings or some non-standard source
func (c *Client) GetBusinessFeatures(source *types.BusinessFeature) (*types.BusinessFeatures, error) {
	req := &types.GetBusinessFeatures{
		Source: source,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.BusinessFeatures), nil
}

// AcceptTermsOfService Accepts Telegram terms of services @terms_of_service_id Terms of service identifier
func (c *Client) AcceptTermsOfService(termsOfServiceId string) (*types.Ok, error) {
	req := &types.AcceptTermsOfService{
		TermsOfServiceId: termsOfServiceId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SearchStringsByPrefix Searches specified query by word prefixes in the provided strings. Returns 0-based positions of strings that matched. Can be called synchronously
func (c *Client) SearchStringsByPrefix(strings []string, query string, limit int32, returnNoneForEmptyQuery bool) (*types.FoundPositions, error) {
	req := &types.SearchStringsByPrefix{
		Strings:                 strings,
		Query:                   query,
		Limit:                   limit,
		ReturnNoneForEmptyQuery: returnNoneForEmptyQuery,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.FoundPositions), nil
}

// SendCustomRequest Sends a custom request; for bots only @method The method name @parameters JSON-serialized method parameters
func (c *Client) SendCustomRequest(method string, parameters string) (*types.CustomRequestResult, error) {
	req := &types.SendCustomRequest{
		Method:     method,
		Parameters: parameters,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CustomRequestResult), nil
}

// AnswerCustomQuery Answers a custom query; for bots only @custom_query_id Identifier of a custom query @data JSON-serialized answer to the query
func (c *Client) AnswerCustomQuery(customQueryId string, data string) (*types.Ok, error) {
	req := &types.AnswerCustomQuery{
		CustomQueryId: customQueryId,
		Data:          data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before initialization @seconds Number of seconds before the function returns
func (c *Client) SetAlarm(seconds float64) (*types.Ok, error) {
	req := &types.SetAlarm{
		Seconds: seconds,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetCountries Returns information about existing countries. Can be called before authorization
func (c *Client) GetCountries() (*types.Countries, error) {
	req := &types.GetCountries{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Countries), nil
}

// GetCountryCode Uses the current IP address to find the current country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (c *Client) GetCountryCode() (*types.Text, error) {
	req := &types.GetCountryCode{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// GetPhoneNumberInfo Returns information about a phone number by its prefix. Can be called before authorization @phone_number_prefix The phone number prefix
func (c *Client) GetPhoneNumberInfo(phoneNumberPrefix string) (*types.PhoneNumberInfo, error) {
	req := &types.GetPhoneNumberInfo{
		PhoneNumberPrefix: phoneNumberPrefix,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PhoneNumberInfo), nil
}

// GetPhoneNumberInfoSync Returns information about a phone number by its prefix synchronously. getCountries must be called at least once after changing localization to the specified language if properly localized country information is expected. Can be called synchronously
func (c *Client) GetPhoneNumberInfoSync(languageCode string, phoneNumberPrefix string) (*types.PhoneNumberInfo, error) {
	req := &types.GetPhoneNumberInfoSync{
		LanguageCode:      languageCode,
		PhoneNumberPrefix: phoneNumberPrefix,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.PhoneNumberInfo), nil
}

// GetCollectibleItemInfo Returns information about a given collectible item that was purchased at https://fragment.com
func (c *Client) GetCollectibleItemInfo(typeField *types.CollectibleItemType) (*types.CollectibleItemInfo, error) {
	req := &types.GetCollectibleItemInfo{
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.CollectibleItemInfo), nil
}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization @link The link
func (c *Client) GetDeepLinkInfo(link string) (*types.DeepLinkInfo, error) {
	req := &types.GetDeepLinkInfo{
		Link: link,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.DeepLinkInfo), nil
}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (c *Client) GetApplicationConfig() (*types.JsonValue, error) {
	req := &types.GetApplicationConfig{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.JsonValue), nil
}

// SaveApplicationLogEvent Saves application log event on the server. Can be called before authorization @type Event type @chat_id Optional chat identifier, associated with the event @data The log event data
func (c *Client) SaveApplicationLogEvent(typeField string, chatId int64, data *types.JsonValue) (*types.Ok, error) {
	req := &types.SaveApplicationLogEvent{
		TypeField: typeField,
		ChatId:    chatId,
		Data:      data,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetApplicationDownloadLink Returns the link for downloading official Telegram application to be used when the current user invites friends to Telegram
func (c *Client) GetApplicationDownloadLink() (*types.HttpUrl, error) {
	req := &types.GetApplicationDownloadLink{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
func (c *Client) AddProxy(server string, port int32, enable bool, typeField *types.ProxyType) (*types.Proxy, error) {
	req := &types.AddProxy{
		Server:    server,
		Port:      port,
		Enable:    enable,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Proxy), nil
}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
func (c *Client) EditProxy(proxyId int32, server string, port int32, enable bool, typeField *types.ProxyType) (*types.Proxy, error) {
	req := &types.EditProxy{
		ProxyId:   proxyId,
		Server:    server,
		Port:      port,
		Enable:    enable,
		TypeField: typeField,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Proxy), nil
}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) EnableProxy(proxyId int32) (*types.Ok, error) {
	req := &types.EnableProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (c *Client) DisableProxy() (*types.Ok, error) {
	req := &types.DisableProxy{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// RemoveProxy Removes a proxy server. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) RemoveProxy(proxyId int32) (*types.Ok, error) {
	req := &types.RemoveProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetProxies Returns the list of proxies that are currently set up. Can be called before authorization
func (c *Client) GetProxies() (*types.Proxies, error) {
	req := &types.GetProxies{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Proxies), nil
}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization @proxy_id Proxy identifier
func (c *Client) GetProxyLink(proxyId int32) (*types.HttpUrl, error) {
	req := &types.GetProxyLink{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.HttpUrl), nil
}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization @proxy_id Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (c *Client) PingProxy(proxyId int32) (*types.Seconds, error) {
	req := &types.PingProxy{
		ProxyId: proxyId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Seconds), nil
}

// SetLogStream Sets new log stream for internal logging of TDLib. Can be called synchronously @log_stream New log stream
func (c *Client) SetLogStream(logStream *types.LogStream) (*types.Ok, error) {
	req := &types.SetLogStream{
		LogStream: logStream,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (c *Client) GetLogStream() (*types.LogStream, error) {
	req := &types.GetLogStream{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LogStream), nil
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib. Can be called synchronously
func (c *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*types.Ok, error) {
	req := &types.SetLogVerbosityLevel{
		NewVerbosityLevel: newVerbosityLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (c *Client) GetLogVerbosityLevel() (*types.LogVerbosityLevel, error) {
	req := &types.GetLogVerbosityLevel{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LogVerbosityLevel), nil
}

// GetLogTags Returns the list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (c *Client) GetLogTags() (*types.LogTags, error) {
	req := &types.GetLogTags{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LogTags), nil
}

// SetLogTagVerbosityLevel Sets the verbosity level for a specified TDLib internal log tag. Can be called synchronously
func (c *Client) SetLogTagVerbosityLevel(tag string, newVerbosityLevel int32) (*types.Ok, error) {
	req := &types.SetLogTagVerbosityLevel{
		Tag:               tag,
		NewVerbosityLevel: newVerbosityLevel,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously @tag Logging tag to change verbosity level
func (c *Client) GetLogTagVerbosityLevel(tag string) (*types.LogVerbosityLevel, error) {
	req := &types.GetLogTagVerbosityLevel{
		Tag: tag,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.LogVerbosityLevel), nil
}

// AddLogMessage Adds a message to TDLib internal log. Can be called synchronously
func (c *Client) AddLogMessage(verbosityLevel int32, text string) (*types.Ok, error) {
	req := &types.AddLogMessage{
		VerbosityLevel: verbosityLevel,
		Text:           text,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// GetUserSupportInfo Returns support information for the given user; for Telegram support only @user_id User identifier
func (c *Client) GetUserSupportInfo(userId int64) (*types.UserSupportInfo, error) {
	req := &types.GetUserSupportInfo{
		UserId: userId,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UserSupportInfo), nil
}

// SetUserSupportInfo Sets support information for the given user; for Telegram support only @user_id User identifier @message New information message
func (c *Client) SetUserSupportInfo(userId int64, message *types.FormattedText) (*types.UserSupportInfo, error) {
	req := &types.SetUserSupportInfo{
		UserId:  userId,
		Message: message,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.UserSupportInfo), nil
}

// GetSupportName Returns localized name of the Telegram support user; for Telegram support only
func (c *Client) GetSupportName() (*types.Text, error) {
	req := &types.GetSupportName{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Text), nil
}

// TestCallEmpty Does nothing; for testing only. This is an offline method. Can be called before authorization
func (c *Client) TestCallEmpty() (*types.Ok, error) {
	req := &types.TestCallEmpty{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization @x String to return
func (c *Client) TestCallString(x string) (*types.TestString, error) {
	req := &types.TestCallString{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestString), nil
}

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization @x Bytes to return
func (c *Client) TestCallBytes(x string) (*types.TestBytes, error) {
	req := &types.TestCallBytes{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestBytes), nil
}

// TestCallVectorInt Returns the received vector of numbers; for testing only. This is an offline method. Can be called before authorization @x Vector of numbers to return
func (c *Client) TestCallVectorInt(x []int32) (*types.TestVectorInt, error) {
	req := &types.TestCallVectorInt{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestVectorInt), nil
}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only. This is an offline method. Can be called before authorization @x Vector of objects to return
func (c *Client) TestCallVectorIntObject(x []*types.TestInt) (*types.TestVectorIntObject, error) {
	req := &types.TestCallVectorIntObject{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestVectorIntObject), nil
}

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization @x Vector of strings to return
func (c *Client) TestCallVectorString(x []string) (*types.TestVectorString, error) {
	req := &types.TestCallVectorString{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestVectorString), nil
}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only. This is an offline method. Can be called before authorization @x Vector of objects to return
func (c *Client) TestCallVectorStringObject(x []*types.TestString) (*types.TestVectorStringObject, error) {
	req := &types.TestCallVectorStringObject{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestVectorStringObject), nil
}

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization @x Number to square
func (c *Client) TestSquareInt(x int32) (*types.TestInt, error) {
	req := &types.TestSquareInt{
		X: x,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.TestInt), nil
}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only. Can be called before authorization
func (c *Client) TestNetwork() (*types.Ok, error) {
	req := &types.TestNetwork{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TestProxy Sends a simple network request to the Telegram servers via proxy; for testing only. Can be called before authorization
func (c *Client) TestProxy(server string, port int32, typeField *types.ProxyType, dcId int32, timeout float64) (*types.Ok, error) {
	req := &types.TestProxy{
		Server:    server,
		Port:      port,
		TypeField: typeField,
		DcId:      dcId,
		Timeout:   timeout,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (c *Client) TestGetDifference() (*types.Ok, error) {
	req := &types.TestGetDifference{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Ok), nil
}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (c *Client) TestUseUpdate() (*types.Update, error) {
	req := &types.TestUseUpdate{}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Update), nil
}

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously @error The error to be returned
func (c *Client) TestReturnError(error *types.Error) (*types.Error, error) {
	req := &types.TestReturnError{
		Error: error,
	}
	resp, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return resp.(*types.Error), nil
}
