package types

import "encoding/json"
import "fmt"

// TlObject is the interface that all TDLib types satisfy
type TlObject interface {
	Type() string
	SetExtra(extra string)
	GetExtra() string
}

// ActiveStoryState Describes state of active stories posted by a chat
type ActiveStoryState struct {
	TypeStr                string                  `json:"@type"`
	ActiveStoryStateLive   *ActiveStoryStateLive   `json:"activeStoryStateLive,omitempty"`
	ActiveStoryStateRead   *ActiveStoryStateRead   `json:"activeStoryStateRead,omitempty"`
	ActiveStoryStateUnread *ActiveStoryStateUnread `json:"activeStoryStateUnread,omitempty"`
}

func (t *ActiveStoryState) Type() string {
	return t.TypeStr
}

func (t *ActiveStoryState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ActiveStoryState) GetExtra() string {
	return ""
}

func (t *ActiveStoryState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "activeStoryStateLive":
		t.ActiveStoryStateLive = new(ActiveStoryStateLive)
		return json.Unmarshal(b, t.ActiveStoryStateLive)
	case "activeStoryStateRead":
		t.ActiveStoryStateRead = new(ActiveStoryStateRead)
		return json.Unmarshal(b, t.ActiveStoryStateRead)
	case "activeStoryStateUnread":
		t.ActiveStoryStateUnread = new(ActiveStoryStateUnread)
		return json.Unmarshal(b, t.ActiveStoryStateUnread)
	}
	return nil
}

func (t *ActiveStoryState) MarshalJSON() ([]byte, error) {
	if t.ActiveStoryStateLive != nil {
		return json.Marshal(t.ActiveStoryStateLive)
	}
	if t.ActiveStoryStateRead != nil {
		return json.Marshal(t.ActiveStoryStateRead)
	}
	if t.ActiveStoryStateUnread != nil {
		return json.Marshal(t.ActiveStoryStateUnread)
	}
	return nil, fmt.Errorf("no value set for ActiveStoryState")
}

// AffiliateProgramSortOrder Describes the order of the found affiliate programs
type AffiliateProgramSortOrder struct {
	TypeStr                                string                                  `json:"@type"`
	AffiliateProgramSortOrderCreationDate  *AffiliateProgramSortOrderCreationDate  `json:"affiliateProgramSortOrderCreationDate,omitempty"`
	AffiliateProgramSortOrderProfitability *AffiliateProgramSortOrderProfitability `json:"affiliateProgramSortOrderProfitability,omitempty"`
	AffiliateProgramSortOrderRevenue       *AffiliateProgramSortOrderRevenue       `json:"affiliateProgramSortOrderRevenue,omitempty"`
}

func (t *AffiliateProgramSortOrder) Type() string {
	return t.TypeStr
}

func (t *AffiliateProgramSortOrder) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AffiliateProgramSortOrder) GetExtra() string {
	return ""
}

func (t *AffiliateProgramSortOrder) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "affiliateProgramSortOrderCreationDate":
		t.AffiliateProgramSortOrderCreationDate = new(AffiliateProgramSortOrderCreationDate)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderCreationDate)
	case "affiliateProgramSortOrderProfitability":
		t.AffiliateProgramSortOrderProfitability = new(AffiliateProgramSortOrderProfitability)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderProfitability)
	case "affiliateProgramSortOrderRevenue":
		t.AffiliateProgramSortOrderRevenue = new(AffiliateProgramSortOrderRevenue)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderRevenue)
	}
	return nil
}

func (t *AffiliateProgramSortOrder) MarshalJSON() ([]byte, error) {
	if t.AffiliateProgramSortOrderCreationDate != nil {
		return json.Marshal(t.AffiliateProgramSortOrderCreationDate)
	}
	if t.AffiliateProgramSortOrderProfitability != nil {
		return json.Marshal(t.AffiliateProgramSortOrderProfitability)
	}
	if t.AffiliateProgramSortOrderRevenue != nil {
		return json.Marshal(t.AffiliateProgramSortOrderRevenue)
	}
	return nil, fmt.Errorf("no value set for AffiliateProgramSortOrder")
}

// AffiliateType Describes type of affiliate for an affiliate program
type AffiliateType struct {
	TypeStr                  string                    `json:"@type"`
	AffiliateTypeBot         *AffiliateTypeBot         `json:"affiliateTypeBot,omitempty"`
	AffiliateTypeChannel     *AffiliateTypeChannel     `json:"affiliateTypeChannel,omitempty"`
	AffiliateTypeCurrentUser *AffiliateTypeCurrentUser `json:"affiliateTypeCurrentUser,omitempty"`
}

func (t *AffiliateType) Type() string {
	return t.TypeStr
}

func (t *AffiliateType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AffiliateType) GetExtra() string {
	return ""
}

func (t *AffiliateType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "affiliateTypeBot":
		t.AffiliateTypeBot = new(AffiliateTypeBot)
		return json.Unmarshal(b, t.AffiliateTypeBot)
	case "affiliateTypeChannel":
		t.AffiliateTypeChannel = new(AffiliateTypeChannel)
		return json.Unmarshal(b, t.AffiliateTypeChannel)
	case "affiliateTypeCurrentUser":
		t.AffiliateTypeCurrentUser = new(AffiliateTypeCurrentUser)
		return json.Unmarshal(b, t.AffiliateTypeCurrentUser)
	}
	return nil
}

func (t *AffiliateType) MarshalJSON() ([]byte, error) {
	if t.AffiliateTypeBot != nil {
		return json.Marshal(t.AffiliateTypeBot)
	}
	if t.AffiliateTypeChannel != nil {
		return json.Marshal(t.AffiliateTypeChannel)
	}
	if t.AffiliateTypeCurrentUser != nil {
		return json.Marshal(t.AffiliateTypeCurrentUser)
	}
	return nil, fmt.Errorf("no value set for AffiliateType")
}

// AuctionState Describes state of an auction
type AuctionState struct {
	TypeStr              string                `json:"@type"`
	AuctionStateActive   *AuctionStateActive   `json:"auctionStateActive,omitempty"`
	AuctionStateFinished *AuctionStateFinished `json:"auctionStateFinished,omitempty"`
}

func (t *AuctionState) Type() string {
	return t.TypeStr
}

func (t *AuctionState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AuctionState) GetExtra() string {
	return ""
}

func (t *AuctionState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "auctionStateActive":
		t.AuctionStateActive = new(AuctionStateActive)
		return json.Unmarshal(b, t.AuctionStateActive)
	case "auctionStateFinished":
		t.AuctionStateFinished = new(AuctionStateFinished)
		return json.Unmarshal(b, t.AuctionStateFinished)
	}
	return nil
}

func (t *AuctionState) MarshalJSON() ([]byte, error) {
	if t.AuctionStateActive != nil {
		return json.Marshal(t.AuctionStateActive)
	}
	if t.AuctionStateFinished != nil {
		return json.Marshal(t.AuctionStateFinished)
	}
	return nil, fmt.Errorf("no value set for AuctionState")
}

// AuthenticationCodeType Provides information about the method by which an authentication code is delivered to the user
type AuthenticationCodeType struct {
	TypeStr                               string                                 `json:"@type"`
	AuthenticationCodeTypeCall            *AuthenticationCodeTypeCall            `json:"authenticationCodeTypeCall,omitempty"`
	AuthenticationCodeTypeFirebaseAndroid *AuthenticationCodeTypeFirebaseAndroid `json:"authenticationCodeTypeFirebaseAndroid,omitempty"`
	AuthenticationCodeTypeFirebaseIos     *AuthenticationCodeTypeFirebaseIos     `json:"authenticationCodeTypeFirebaseIos,omitempty"`
	AuthenticationCodeTypeFlashCall       *AuthenticationCodeTypeFlashCall       `json:"authenticationCodeTypeFlashCall,omitempty"`
	AuthenticationCodeTypeFragment        *AuthenticationCodeTypeFragment        `json:"authenticationCodeTypeFragment,omitempty"`
	AuthenticationCodeTypeMissedCall      *AuthenticationCodeTypeMissedCall      `json:"authenticationCodeTypeMissedCall,omitempty"`
	AuthenticationCodeTypeSms             *AuthenticationCodeTypeSms             `json:"authenticationCodeTypeSms,omitempty"`
	AuthenticationCodeTypeSmsPhrase       *AuthenticationCodeTypeSmsPhrase       `json:"authenticationCodeTypeSmsPhrase,omitempty"`
	AuthenticationCodeTypeSmsWord         *AuthenticationCodeTypeSmsWord         `json:"authenticationCodeTypeSmsWord,omitempty"`
	AuthenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage `json:"authenticationCodeTypeTelegramMessage,omitempty"`
}

func (t *AuthenticationCodeType) Type() string {
	return t.TypeStr
}

func (t *AuthenticationCodeType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AuthenticationCodeType) GetExtra() string {
	return ""
}

func (t *AuthenticationCodeType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "authenticationCodeTypeCall":
		t.AuthenticationCodeTypeCall = new(AuthenticationCodeTypeCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeCall)
	case "authenticationCodeTypeFirebaseAndroid":
		t.AuthenticationCodeTypeFirebaseAndroid = new(AuthenticationCodeTypeFirebaseAndroid)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFirebaseAndroid)
	case "authenticationCodeTypeFirebaseIos":
		t.AuthenticationCodeTypeFirebaseIos = new(AuthenticationCodeTypeFirebaseIos)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFirebaseIos)
	case "authenticationCodeTypeFlashCall":
		t.AuthenticationCodeTypeFlashCall = new(AuthenticationCodeTypeFlashCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFlashCall)
	case "authenticationCodeTypeFragment":
		t.AuthenticationCodeTypeFragment = new(AuthenticationCodeTypeFragment)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFragment)
	case "authenticationCodeTypeMissedCall":
		t.AuthenticationCodeTypeMissedCall = new(AuthenticationCodeTypeMissedCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeMissedCall)
	case "authenticationCodeTypeSms":
		t.AuthenticationCodeTypeSms = new(AuthenticationCodeTypeSms)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSms)
	case "authenticationCodeTypeSmsPhrase":
		t.AuthenticationCodeTypeSmsPhrase = new(AuthenticationCodeTypeSmsPhrase)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSmsPhrase)
	case "authenticationCodeTypeSmsWord":
		t.AuthenticationCodeTypeSmsWord = new(AuthenticationCodeTypeSmsWord)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSmsWord)
	case "authenticationCodeTypeTelegramMessage":
		t.AuthenticationCodeTypeTelegramMessage = new(AuthenticationCodeTypeTelegramMessage)
		return json.Unmarshal(b, t.AuthenticationCodeTypeTelegramMessage)
	}
	return nil
}

func (t *AuthenticationCodeType) MarshalJSON() ([]byte, error) {
	if t.AuthenticationCodeTypeCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeCall)
	}
	if t.AuthenticationCodeTypeFirebaseAndroid != nil {
		return json.Marshal(t.AuthenticationCodeTypeFirebaseAndroid)
	}
	if t.AuthenticationCodeTypeFirebaseIos != nil {
		return json.Marshal(t.AuthenticationCodeTypeFirebaseIos)
	}
	if t.AuthenticationCodeTypeFlashCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeFlashCall)
	}
	if t.AuthenticationCodeTypeFragment != nil {
		return json.Marshal(t.AuthenticationCodeTypeFragment)
	}
	if t.AuthenticationCodeTypeMissedCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeMissedCall)
	}
	if t.AuthenticationCodeTypeSms != nil {
		return json.Marshal(t.AuthenticationCodeTypeSms)
	}
	if t.AuthenticationCodeTypeSmsPhrase != nil {
		return json.Marshal(t.AuthenticationCodeTypeSmsPhrase)
	}
	if t.AuthenticationCodeTypeSmsWord != nil {
		return json.Marshal(t.AuthenticationCodeTypeSmsWord)
	}
	if t.AuthenticationCodeTypeTelegramMessage != nil {
		return json.Marshal(t.AuthenticationCodeTypeTelegramMessage)
	}
	return nil, fmt.Errorf("no value set for AuthenticationCodeType")
}

// AuthorizationState Represents the current authorization state of the TDLib client
type AuthorizationState struct {
	TypeStr                                       string                                         `json:"@type"`
	AuthorizationStateClosed                      *AuthorizationStateClosed                      `json:"authorizationStateClosed,omitempty"`
	AuthorizationStateClosing                     *AuthorizationStateClosing                     `json:"authorizationStateClosing,omitempty"`
	AuthorizationStateLoggingOut                  *AuthorizationStateLoggingOut                  `json:"authorizationStateLoggingOut,omitempty"`
	AuthorizationStateReady                       *AuthorizationStateReady                       `json:"authorizationStateReady,omitempty"`
	AuthorizationStateWaitCode                    *AuthorizationStateWaitCode                    `json:"authorizationStateWaitCode,omitempty"`
	AuthorizationStateWaitEmailAddress            *AuthorizationStateWaitEmailAddress            `json:"authorizationStateWaitEmailAddress,omitempty"`
	AuthorizationStateWaitEmailCode               *AuthorizationStateWaitEmailCode               `json:"authorizationStateWaitEmailCode,omitempty"`
	AuthorizationStateWaitOtherDeviceConfirmation *AuthorizationStateWaitOtherDeviceConfirmation `json:"authorizationStateWaitOtherDeviceConfirmation,omitempty"`
	AuthorizationStateWaitPassword                *AuthorizationStateWaitPassword                `json:"authorizationStateWaitPassword,omitempty"`
	AuthorizationStateWaitPhoneNumber             *AuthorizationStateWaitPhoneNumber             `json:"authorizationStateWaitPhoneNumber,omitempty"`
	AuthorizationStateWaitPremiumPurchase         *AuthorizationStateWaitPremiumPurchase         `json:"authorizationStateWaitPremiumPurchase,omitempty"`
	AuthorizationStateWaitRegistration            *AuthorizationStateWaitRegistration            `json:"authorizationStateWaitRegistration,omitempty"`
	AuthorizationStateWaitTdlibParameters         *AuthorizationStateWaitTdlibParameters         `json:"authorizationStateWaitTdlibParameters,omitempty"`
}

func (t *AuthorizationState) Type() string {
	return t.TypeStr
}

func (t *AuthorizationState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AuthorizationState) GetExtra() string {
	return ""
}

func (t *AuthorizationState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "authorizationStateClosed":
		t.AuthorizationStateClosed = new(AuthorizationStateClosed)
		return json.Unmarshal(b, t.AuthorizationStateClosed)
	case "authorizationStateClosing":
		t.AuthorizationStateClosing = new(AuthorizationStateClosing)
		return json.Unmarshal(b, t.AuthorizationStateClosing)
	case "authorizationStateLoggingOut":
		t.AuthorizationStateLoggingOut = new(AuthorizationStateLoggingOut)
		return json.Unmarshal(b, t.AuthorizationStateLoggingOut)
	case "authorizationStateReady":
		t.AuthorizationStateReady = new(AuthorizationStateReady)
		return json.Unmarshal(b, t.AuthorizationStateReady)
	case "authorizationStateWaitCode":
		t.AuthorizationStateWaitCode = new(AuthorizationStateWaitCode)
		return json.Unmarshal(b, t.AuthorizationStateWaitCode)
	case "authorizationStateWaitEmailAddress":
		t.AuthorizationStateWaitEmailAddress = new(AuthorizationStateWaitEmailAddress)
		return json.Unmarshal(b, t.AuthorizationStateWaitEmailAddress)
	case "authorizationStateWaitEmailCode":
		t.AuthorizationStateWaitEmailCode = new(AuthorizationStateWaitEmailCode)
		return json.Unmarshal(b, t.AuthorizationStateWaitEmailCode)
	case "authorizationStateWaitOtherDeviceConfirmation":
		t.AuthorizationStateWaitOtherDeviceConfirmation = new(AuthorizationStateWaitOtherDeviceConfirmation)
		return json.Unmarshal(b, t.AuthorizationStateWaitOtherDeviceConfirmation)
	case "authorizationStateWaitPassword":
		t.AuthorizationStateWaitPassword = new(AuthorizationStateWaitPassword)
		return json.Unmarshal(b, t.AuthorizationStateWaitPassword)
	case "authorizationStateWaitPhoneNumber":
		t.AuthorizationStateWaitPhoneNumber = new(AuthorizationStateWaitPhoneNumber)
		return json.Unmarshal(b, t.AuthorizationStateWaitPhoneNumber)
	case "authorizationStateWaitPremiumPurchase":
		t.AuthorizationStateWaitPremiumPurchase = new(AuthorizationStateWaitPremiumPurchase)
		return json.Unmarshal(b, t.AuthorizationStateWaitPremiumPurchase)
	case "authorizationStateWaitRegistration":
		t.AuthorizationStateWaitRegistration = new(AuthorizationStateWaitRegistration)
		return json.Unmarshal(b, t.AuthorizationStateWaitRegistration)
	case "authorizationStateWaitTdlibParameters":
		t.AuthorizationStateWaitTdlibParameters = new(AuthorizationStateWaitTdlibParameters)
		return json.Unmarshal(b, t.AuthorizationStateWaitTdlibParameters)
	}
	return nil
}

func (t *AuthorizationState) MarshalJSON() ([]byte, error) {
	if t.AuthorizationStateClosed != nil {
		return json.Marshal(t.AuthorizationStateClosed)
	}
	if t.AuthorizationStateClosing != nil {
		return json.Marshal(t.AuthorizationStateClosing)
	}
	if t.AuthorizationStateLoggingOut != nil {
		return json.Marshal(t.AuthorizationStateLoggingOut)
	}
	if t.AuthorizationStateReady != nil {
		return json.Marshal(t.AuthorizationStateReady)
	}
	if t.AuthorizationStateWaitCode != nil {
		return json.Marshal(t.AuthorizationStateWaitCode)
	}
	if t.AuthorizationStateWaitEmailAddress != nil {
		return json.Marshal(t.AuthorizationStateWaitEmailAddress)
	}
	if t.AuthorizationStateWaitEmailCode != nil {
		return json.Marshal(t.AuthorizationStateWaitEmailCode)
	}
	if t.AuthorizationStateWaitOtherDeviceConfirmation != nil {
		return json.Marshal(t.AuthorizationStateWaitOtherDeviceConfirmation)
	}
	if t.AuthorizationStateWaitPassword != nil {
		return json.Marshal(t.AuthorizationStateWaitPassword)
	}
	if t.AuthorizationStateWaitPhoneNumber != nil {
		return json.Marshal(t.AuthorizationStateWaitPhoneNumber)
	}
	if t.AuthorizationStateWaitPremiumPurchase != nil {
		return json.Marshal(t.AuthorizationStateWaitPremiumPurchase)
	}
	if t.AuthorizationStateWaitRegistration != nil {
		return json.Marshal(t.AuthorizationStateWaitRegistration)
	}
	if t.AuthorizationStateWaitTdlibParameters != nil {
		return json.Marshal(t.AuthorizationStateWaitTdlibParameters)
	}
	return nil, fmt.Errorf("no value set for AuthorizationState")
}

// AutosaveSettingsScope Describes scope of autosave settings
type AutosaveSettingsScope struct {
	TypeStr                           string                             `json:"@type"`
	AutosaveSettingsScopeChannelChats *AutosaveSettingsScopeChannelChats `json:"autosaveSettingsScopeChannelChats,omitempty"`
	AutosaveSettingsScopeChat         *AutosaveSettingsScopeChat         `json:"autosaveSettingsScopeChat,omitempty"`
	AutosaveSettingsScopeGroupChats   *AutosaveSettingsScopeGroupChats   `json:"autosaveSettingsScopeGroupChats,omitempty"`
	AutosaveSettingsScopePrivateChats *AutosaveSettingsScopePrivateChats `json:"autosaveSettingsScopePrivateChats,omitempty"`
}

func (t *AutosaveSettingsScope) Type() string {
	return t.TypeStr
}

func (t *AutosaveSettingsScope) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *AutosaveSettingsScope) GetExtra() string {
	return ""
}

func (t *AutosaveSettingsScope) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "autosaveSettingsScopeChannelChats":
		t.AutosaveSettingsScopeChannelChats = new(AutosaveSettingsScopeChannelChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopeChannelChats)
	case "autosaveSettingsScopeChat":
		t.AutosaveSettingsScopeChat = new(AutosaveSettingsScopeChat)
		return json.Unmarshal(b, t.AutosaveSettingsScopeChat)
	case "autosaveSettingsScopeGroupChats":
		t.AutosaveSettingsScopeGroupChats = new(AutosaveSettingsScopeGroupChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopeGroupChats)
	case "autosaveSettingsScopePrivateChats":
		t.AutosaveSettingsScopePrivateChats = new(AutosaveSettingsScopePrivateChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopePrivateChats)
	}
	return nil
}

func (t *AutosaveSettingsScope) MarshalJSON() ([]byte, error) {
	if t.AutosaveSettingsScopeChannelChats != nil {
		return json.Marshal(t.AutosaveSettingsScopeChannelChats)
	}
	if t.AutosaveSettingsScopeChat != nil {
		return json.Marshal(t.AutosaveSettingsScopeChat)
	}
	if t.AutosaveSettingsScopeGroupChats != nil {
		return json.Marshal(t.AutosaveSettingsScopeGroupChats)
	}
	if t.AutosaveSettingsScopePrivateChats != nil {
		return json.Marshal(t.AutosaveSettingsScopePrivateChats)
	}
	return nil, fmt.Errorf("no value set for AutosaveSettingsScope")
}

// BackgroundFill Describes a fill of a background
type BackgroundFill struct {
	TypeStr                        string                          `json:"@type"`
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient `json:"backgroundFillFreeformGradient,omitempty"`
	BackgroundFillGradient         *BackgroundFillGradient         `json:"backgroundFillGradient,omitempty"`
	BackgroundFillSolid            *BackgroundFillSolid            `json:"backgroundFillSolid,omitempty"`
}

func (t *BackgroundFill) Type() string {
	return t.TypeStr
}

func (t *BackgroundFill) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BackgroundFill) GetExtra() string {
	return ""
}

func (t *BackgroundFill) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "backgroundFillFreeformGradient":
		t.BackgroundFillFreeformGradient = new(BackgroundFillFreeformGradient)
		return json.Unmarshal(b, t.BackgroundFillFreeformGradient)
	case "backgroundFillGradient":
		t.BackgroundFillGradient = new(BackgroundFillGradient)
		return json.Unmarshal(b, t.BackgroundFillGradient)
	case "backgroundFillSolid":
		t.BackgroundFillSolid = new(BackgroundFillSolid)
		return json.Unmarshal(b, t.BackgroundFillSolid)
	}
	return nil
}

func (t *BackgroundFill) MarshalJSON() ([]byte, error) {
	if t.BackgroundFillFreeformGradient != nil {
		return json.Marshal(t.BackgroundFillFreeformGradient)
	}
	if t.BackgroundFillGradient != nil {
		return json.Marshal(t.BackgroundFillGradient)
	}
	if t.BackgroundFillSolid != nil {
		return json.Marshal(t.BackgroundFillSolid)
	}
	return nil, fmt.Errorf("no value set for BackgroundFill")
}

// BackgroundType Describes the type of background
type BackgroundType struct {
	TypeStr                 string                   `json:"@type"`
	BackgroundTypeChatTheme *BackgroundTypeChatTheme `json:"backgroundTypeChatTheme,omitempty"`
	BackgroundTypeFill      *BackgroundTypeFill      `json:"backgroundTypeFill,omitempty"`
	BackgroundTypePattern   *BackgroundTypePattern   `json:"backgroundTypePattern,omitempty"`
	BackgroundTypeWallpaper *BackgroundTypeWallpaper `json:"backgroundTypeWallpaper,omitempty"`
}

func (t *BackgroundType) Type() string {
	return t.TypeStr
}

func (t *BackgroundType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BackgroundType) GetExtra() string {
	return ""
}

func (t *BackgroundType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "backgroundTypeChatTheme":
		t.BackgroundTypeChatTheme = new(BackgroundTypeChatTheme)
		return json.Unmarshal(b, t.BackgroundTypeChatTheme)
	case "backgroundTypeFill":
		t.BackgroundTypeFill = new(BackgroundTypeFill)
		return json.Unmarshal(b, t.BackgroundTypeFill)
	case "backgroundTypePattern":
		t.BackgroundTypePattern = new(BackgroundTypePattern)
		return json.Unmarshal(b, t.BackgroundTypePattern)
	case "backgroundTypeWallpaper":
		t.BackgroundTypeWallpaper = new(BackgroundTypeWallpaper)
		return json.Unmarshal(b, t.BackgroundTypeWallpaper)
	}
	return nil
}

func (t *BackgroundType) MarshalJSON() ([]byte, error) {
	if t.BackgroundTypeChatTheme != nil {
		return json.Marshal(t.BackgroundTypeChatTheme)
	}
	if t.BackgroundTypeFill != nil {
		return json.Marshal(t.BackgroundTypeFill)
	}
	if t.BackgroundTypePattern != nil {
		return json.Marshal(t.BackgroundTypePattern)
	}
	if t.BackgroundTypeWallpaper != nil {
		return json.Marshal(t.BackgroundTypeWallpaper)
	}
	return nil, fmt.Errorf("no value set for BackgroundType")
}

// BlockList Describes type of block list
type BlockList struct {
	TypeStr          string            `json:"@type"`
	BlockListMain    *BlockListMain    `json:"blockListMain,omitempty"`
	BlockListStories *BlockListStories `json:"blockListStories,omitempty"`
}

func (t *BlockList) Type() string {
	return t.TypeStr
}

func (t *BlockList) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BlockList) GetExtra() string {
	return ""
}

func (t *BlockList) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "blockListMain":
		t.BlockListMain = new(BlockListMain)
		return json.Unmarshal(b, t.BlockListMain)
	case "blockListStories":
		t.BlockListStories = new(BlockListStories)
		return json.Unmarshal(b, t.BlockListStories)
	}
	return nil
}

func (t *BlockList) MarshalJSON() ([]byte, error) {
	if t.BlockListMain != nil {
		return json.Marshal(t.BlockListMain)
	}
	if t.BlockListStories != nil {
		return json.Marshal(t.BlockListStories)
	}
	return nil, fmt.Errorf("no value set for BlockList")
}

// BotCommandScope Represents the scope to which bot commands are relevant
type BotCommandScope struct {
	TypeStr                              string                                `json:"@type"`
	BotCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators `json:"botCommandScopeAllChatAdministrators,omitempty"`
	BotCommandScopeAllGroupChats         *BotCommandScopeAllGroupChats         `json:"botCommandScopeAllGroupChats,omitempty"`
	BotCommandScopeAllPrivateChats       *BotCommandScopeAllPrivateChats       `json:"botCommandScopeAllPrivateChats,omitempty"`
	BotCommandScopeChat                  *BotCommandScopeChat                  `json:"botCommandScopeChat,omitempty"`
	BotCommandScopeChatAdministrators    *BotCommandScopeChatAdministrators    `json:"botCommandScopeChatAdministrators,omitempty"`
	BotCommandScopeChatMember            *BotCommandScopeChatMember            `json:"botCommandScopeChatMember,omitempty"`
	BotCommandScopeDefault               *BotCommandScopeDefault               `json:"botCommandScopeDefault,omitempty"`
}

func (t *BotCommandScope) Type() string {
	return t.TypeStr
}

func (t *BotCommandScope) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BotCommandScope) GetExtra() string {
	return ""
}

func (t *BotCommandScope) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "botCommandScopeAllChatAdministrators":
		t.BotCommandScopeAllChatAdministrators = new(BotCommandScopeAllChatAdministrators)
		return json.Unmarshal(b, t.BotCommandScopeAllChatAdministrators)
	case "botCommandScopeAllGroupChats":
		t.BotCommandScopeAllGroupChats = new(BotCommandScopeAllGroupChats)
		return json.Unmarshal(b, t.BotCommandScopeAllGroupChats)
	case "botCommandScopeAllPrivateChats":
		t.BotCommandScopeAllPrivateChats = new(BotCommandScopeAllPrivateChats)
		return json.Unmarshal(b, t.BotCommandScopeAllPrivateChats)
	case "botCommandScopeChat":
		t.BotCommandScopeChat = new(BotCommandScopeChat)
		return json.Unmarshal(b, t.BotCommandScopeChat)
	case "botCommandScopeChatAdministrators":
		t.BotCommandScopeChatAdministrators = new(BotCommandScopeChatAdministrators)
		return json.Unmarshal(b, t.BotCommandScopeChatAdministrators)
	case "botCommandScopeChatMember":
		t.BotCommandScopeChatMember = new(BotCommandScopeChatMember)
		return json.Unmarshal(b, t.BotCommandScopeChatMember)
	case "botCommandScopeDefault":
		t.BotCommandScopeDefault = new(BotCommandScopeDefault)
		return json.Unmarshal(b, t.BotCommandScopeDefault)
	}
	return nil
}

func (t *BotCommandScope) MarshalJSON() ([]byte, error) {
	if t.BotCommandScopeAllChatAdministrators != nil {
		return json.Marshal(t.BotCommandScopeAllChatAdministrators)
	}
	if t.BotCommandScopeAllGroupChats != nil {
		return json.Marshal(t.BotCommandScopeAllGroupChats)
	}
	if t.BotCommandScopeAllPrivateChats != nil {
		return json.Marshal(t.BotCommandScopeAllPrivateChats)
	}
	if t.BotCommandScopeChat != nil {
		return json.Marshal(t.BotCommandScopeChat)
	}
	if t.BotCommandScopeChatAdministrators != nil {
		return json.Marshal(t.BotCommandScopeChatAdministrators)
	}
	if t.BotCommandScopeChatMember != nil {
		return json.Marshal(t.BotCommandScopeChatMember)
	}
	if t.BotCommandScopeDefault != nil {
		return json.Marshal(t.BotCommandScopeDefault)
	}
	return nil, fmt.Errorf("no value set for BotCommandScope")
}

// BotWriteAccessAllowReason Describes a reason why a bot was allowed to write messages to the current user
type BotWriteAccessAllowReason struct {
	TypeStr                                        string                                          `json:"@type"`
	BotWriteAccessAllowReasonAcceptedRequest       *BotWriteAccessAllowReasonAcceptedRequest       `json:"botWriteAccessAllowReasonAcceptedRequest,omitempty"`
	BotWriteAccessAllowReasonAddedToAttachmentMenu *BotWriteAccessAllowReasonAddedToAttachmentMenu `json:"botWriteAccessAllowReasonAddedToAttachmentMenu,omitempty"`
	BotWriteAccessAllowReasonConnectedWebsite      *BotWriteAccessAllowReasonConnectedWebsite      `json:"botWriteAccessAllowReasonConnectedWebsite,omitempty"`
	BotWriteAccessAllowReasonLaunchedWebApp        *BotWriteAccessAllowReasonLaunchedWebApp        `json:"botWriteAccessAllowReasonLaunchedWebApp,omitempty"`
}

func (t *BotWriteAccessAllowReason) Type() string {
	return t.TypeStr
}

func (t *BotWriteAccessAllowReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BotWriteAccessAllowReason) GetExtra() string {
	return ""
}

func (t *BotWriteAccessAllowReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "botWriteAccessAllowReasonAcceptedRequest":
		t.BotWriteAccessAllowReasonAcceptedRequest = new(BotWriteAccessAllowReasonAcceptedRequest)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonAcceptedRequest)
	case "botWriteAccessAllowReasonAddedToAttachmentMenu":
		t.BotWriteAccessAllowReasonAddedToAttachmentMenu = new(BotWriteAccessAllowReasonAddedToAttachmentMenu)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonAddedToAttachmentMenu)
	case "botWriteAccessAllowReasonConnectedWebsite":
		t.BotWriteAccessAllowReasonConnectedWebsite = new(BotWriteAccessAllowReasonConnectedWebsite)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonConnectedWebsite)
	case "botWriteAccessAllowReasonLaunchedWebApp":
		t.BotWriteAccessAllowReasonLaunchedWebApp = new(BotWriteAccessAllowReasonLaunchedWebApp)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonLaunchedWebApp)
	}
	return nil
}

func (t *BotWriteAccessAllowReason) MarshalJSON() ([]byte, error) {
	if t.BotWriteAccessAllowReasonAcceptedRequest != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonAcceptedRequest)
	}
	if t.BotWriteAccessAllowReasonAddedToAttachmentMenu != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonAddedToAttachmentMenu)
	}
	if t.BotWriteAccessAllowReasonConnectedWebsite != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonConnectedWebsite)
	}
	if t.BotWriteAccessAllowReasonLaunchedWebApp != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonLaunchedWebApp)
	}
	return nil, fmt.Errorf("no value set for BotWriteAccessAllowReason")
}

// BuiltInTheme Describes a built-in theme of an official app
type BuiltInTheme struct {
	TypeStr             string               `json:"@type"`
	BuiltInThemeArctic  *BuiltInThemeArctic  `json:"builtInThemeArctic,omitempty"`
	BuiltInThemeClassic *BuiltInThemeClassic `json:"builtInThemeClassic,omitempty"`
	BuiltInThemeDay     *BuiltInThemeDay     `json:"builtInThemeDay,omitempty"`
	BuiltInThemeNight   *BuiltInThemeNight   `json:"builtInThemeNight,omitempty"`
	BuiltInThemeTinted  *BuiltInThemeTinted  `json:"builtInThemeTinted,omitempty"`
}

func (t *BuiltInTheme) Type() string {
	return t.TypeStr
}

func (t *BuiltInTheme) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BuiltInTheme) GetExtra() string {
	return ""
}

func (t *BuiltInTheme) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "builtInThemeArctic":
		t.BuiltInThemeArctic = new(BuiltInThemeArctic)
		return json.Unmarshal(b, t.BuiltInThemeArctic)
	case "builtInThemeClassic":
		t.BuiltInThemeClassic = new(BuiltInThemeClassic)
		return json.Unmarshal(b, t.BuiltInThemeClassic)
	case "builtInThemeDay":
		t.BuiltInThemeDay = new(BuiltInThemeDay)
		return json.Unmarshal(b, t.BuiltInThemeDay)
	case "builtInThemeNight":
		t.BuiltInThemeNight = new(BuiltInThemeNight)
		return json.Unmarshal(b, t.BuiltInThemeNight)
	case "builtInThemeTinted":
		t.BuiltInThemeTinted = new(BuiltInThemeTinted)
		return json.Unmarshal(b, t.BuiltInThemeTinted)
	}
	return nil
}

func (t *BuiltInTheme) MarshalJSON() ([]byte, error) {
	if t.BuiltInThemeArctic != nil {
		return json.Marshal(t.BuiltInThemeArctic)
	}
	if t.BuiltInThemeClassic != nil {
		return json.Marshal(t.BuiltInThemeClassic)
	}
	if t.BuiltInThemeDay != nil {
		return json.Marshal(t.BuiltInThemeDay)
	}
	if t.BuiltInThemeNight != nil {
		return json.Marshal(t.BuiltInThemeNight)
	}
	if t.BuiltInThemeTinted != nil {
		return json.Marshal(t.BuiltInThemeTinted)
	}
	return nil, fmt.Errorf("no value set for BuiltInTheme")
}

// BusinessAwayMessageSchedule Describes conditions for sending of away messages by a Telegram Business account
type BusinessAwayMessageSchedule struct {
	TypeStr                                          string                                            `json:"@type"`
	BusinessAwayMessageScheduleAlways                *BusinessAwayMessageScheduleAlways                `json:"businessAwayMessageScheduleAlways,omitempty"`
	BusinessAwayMessageScheduleCustom                *BusinessAwayMessageScheduleCustom                `json:"businessAwayMessageScheduleCustom,omitempty"`
	BusinessAwayMessageScheduleOutsideOfOpeningHours *BusinessAwayMessageScheduleOutsideOfOpeningHours `json:"businessAwayMessageScheduleOutsideOfOpeningHours,omitempty"`
}

func (t *BusinessAwayMessageSchedule) Type() string {
	return t.TypeStr
}

func (t *BusinessAwayMessageSchedule) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BusinessAwayMessageSchedule) GetExtra() string {
	return ""
}

func (t *BusinessAwayMessageSchedule) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "businessAwayMessageScheduleAlways":
		t.BusinessAwayMessageScheduleAlways = new(BusinessAwayMessageScheduleAlways)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleAlways)
	case "businessAwayMessageScheduleCustom":
		t.BusinessAwayMessageScheduleCustom = new(BusinessAwayMessageScheduleCustom)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleCustom)
	case "businessAwayMessageScheduleOutsideOfOpeningHours":
		t.BusinessAwayMessageScheduleOutsideOfOpeningHours = new(BusinessAwayMessageScheduleOutsideOfOpeningHours)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleOutsideOfOpeningHours)
	}
	return nil
}

func (t *BusinessAwayMessageSchedule) MarshalJSON() ([]byte, error) {
	if t.BusinessAwayMessageScheduleAlways != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleAlways)
	}
	if t.BusinessAwayMessageScheduleCustom != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleCustom)
	}
	if t.BusinessAwayMessageScheduleOutsideOfOpeningHours != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleOutsideOfOpeningHours)
	}
	return nil, fmt.Errorf("no value set for BusinessAwayMessageSchedule")
}

// BusinessFeature Describes a feature available to Business user accounts
type BusinessFeature struct {
	TypeStr                        string                          `json:"@type"`
	BusinessFeatureAccountLinks    *BusinessFeatureAccountLinks    `json:"businessFeatureAccountLinks,omitempty"`
	BusinessFeatureAwayMessage     *BusinessFeatureAwayMessage     `json:"businessFeatureAwayMessage,omitempty"`
	BusinessFeatureBots            *BusinessFeatureBots            `json:"businessFeatureBots,omitempty"`
	BusinessFeatureChatFolderTags  *BusinessFeatureChatFolderTags  `json:"businessFeatureChatFolderTags,omitempty"`
	BusinessFeatureEmojiStatus     *BusinessFeatureEmojiStatus     `json:"businessFeatureEmojiStatus,omitempty"`
	BusinessFeatureGreetingMessage *BusinessFeatureGreetingMessage `json:"businessFeatureGreetingMessage,omitempty"`
	BusinessFeatureLocation        *BusinessFeatureLocation        `json:"businessFeatureLocation,omitempty"`
	BusinessFeatureOpeningHours    *BusinessFeatureOpeningHours    `json:"businessFeatureOpeningHours,omitempty"`
	BusinessFeatureQuickReplies    *BusinessFeatureQuickReplies    `json:"businessFeatureQuickReplies,omitempty"`
	BusinessFeatureStartPage       *BusinessFeatureStartPage       `json:"businessFeatureStartPage,omitempty"`
	BusinessFeatureUpgradedStories *BusinessFeatureUpgradedStories `json:"businessFeatureUpgradedStories,omitempty"`
}

func (t *BusinessFeature) Type() string {
	return t.TypeStr
}

func (t *BusinessFeature) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *BusinessFeature) GetExtra() string {
	return ""
}

func (t *BusinessFeature) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "businessFeatureAccountLinks":
		t.BusinessFeatureAccountLinks = new(BusinessFeatureAccountLinks)
		return json.Unmarshal(b, t.BusinessFeatureAccountLinks)
	case "businessFeatureAwayMessage":
		t.BusinessFeatureAwayMessage = new(BusinessFeatureAwayMessage)
		return json.Unmarshal(b, t.BusinessFeatureAwayMessage)
	case "businessFeatureBots":
		t.BusinessFeatureBots = new(BusinessFeatureBots)
		return json.Unmarshal(b, t.BusinessFeatureBots)
	case "businessFeatureChatFolderTags":
		t.BusinessFeatureChatFolderTags = new(BusinessFeatureChatFolderTags)
		return json.Unmarshal(b, t.BusinessFeatureChatFolderTags)
	case "businessFeatureEmojiStatus":
		t.BusinessFeatureEmojiStatus = new(BusinessFeatureEmojiStatus)
		return json.Unmarshal(b, t.BusinessFeatureEmojiStatus)
	case "businessFeatureGreetingMessage":
		t.BusinessFeatureGreetingMessage = new(BusinessFeatureGreetingMessage)
		return json.Unmarshal(b, t.BusinessFeatureGreetingMessage)
	case "businessFeatureLocation":
		t.BusinessFeatureLocation = new(BusinessFeatureLocation)
		return json.Unmarshal(b, t.BusinessFeatureLocation)
	case "businessFeatureOpeningHours":
		t.BusinessFeatureOpeningHours = new(BusinessFeatureOpeningHours)
		return json.Unmarshal(b, t.BusinessFeatureOpeningHours)
	case "businessFeatureQuickReplies":
		t.BusinessFeatureQuickReplies = new(BusinessFeatureQuickReplies)
		return json.Unmarshal(b, t.BusinessFeatureQuickReplies)
	case "businessFeatureStartPage":
		t.BusinessFeatureStartPage = new(BusinessFeatureStartPage)
		return json.Unmarshal(b, t.BusinessFeatureStartPage)
	case "businessFeatureUpgradedStories":
		t.BusinessFeatureUpgradedStories = new(BusinessFeatureUpgradedStories)
		return json.Unmarshal(b, t.BusinessFeatureUpgradedStories)
	}
	return nil
}

func (t *BusinessFeature) MarshalJSON() ([]byte, error) {
	if t.BusinessFeatureAccountLinks != nil {
		return json.Marshal(t.BusinessFeatureAccountLinks)
	}
	if t.BusinessFeatureAwayMessage != nil {
		return json.Marshal(t.BusinessFeatureAwayMessage)
	}
	if t.BusinessFeatureBots != nil {
		return json.Marshal(t.BusinessFeatureBots)
	}
	if t.BusinessFeatureChatFolderTags != nil {
		return json.Marshal(t.BusinessFeatureChatFolderTags)
	}
	if t.BusinessFeatureEmojiStatus != nil {
		return json.Marshal(t.BusinessFeatureEmojiStatus)
	}
	if t.BusinessFeatureGreetingMessage != nil {
		return json.Marshal(t.BusinessFeatureGreetingMessage)
	}
	if t.BusinessFeatureLocation != nil {
		return json.Marshal(t.BusinessFeatureLocation)
	}
	if t.BusinessFeatureOpeningHours != nil {
		return json.Marshal(t.BusinessFeatureOpeningHours)
	}
	if t.BusinessFeatureQuickReplies != nil {
		return json.Marshal(t.BusinessFeatureQuickReplies)
	}
	if t.BusinessFeatureStartPage != nil {
		return json.Marshal(t.BusinessFeatureStartPage)
	}
	if t.BusinessFeatureUpgradedStories != nil {
		return json.Marshal(t.BusinessFeatureUpgradedStories)
	}
	return nil, fmt.Errorf("no value set for BusinessFeature")
}

// CallbackQueryPayload Represents a payload of a callback query
type CallbackQueryPayload struct {
	TypeStr                              string                                `json:"@type"`
	CallbackQueryPayloadData             *CallbackQueryPayloadData             `json:"callbackQueryPayloadData,omitempty"`
	CallbackQueryPayloadDataWithPassword *CallbackQueryPayloadDataWithPassword `json:"callbackQueryPayloadDataWithPassword,omitempty"`
	CallbackQueryPayloadGame             *CallbackQueryPayloadGame             `json:"callbackQueryPayloadGame,omitempty"`
}

func (t *CallbackQueryPayload) Type() string {
	return t.TypeStr
}

func (t *CallbackQueryPayload) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CallbackQueryPayload) GetExtra() string {
	return ""
}

func (t *CallbackQueryPayload) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "callbackQueryPayloadData":
		t.CallbackQueryPayloadData = new(CallbackQueryPayloadData)
		return json.Unmarshal(b, t.CallbackQueryPayloadData)
	case "callbackQueryPayloadDataWithPassword":
		t.CallbackQueryPayloadDataWithPassword = new(CallbackQueryPayloadDataWithPassword)
		return json.Unmarshal(b, t.CallbackQueryPayloadDataWithPassword)
	case "callbackQueryPayloadGame":
		t.CallbackQueryPayloadGame = new(CallbackQueryPayloadGame)
		return json.Unmarshal(b, t.CallbackQueryPayloadGame)
	}
	return nil
}

func (t *CallbackQueryPayload) MarshalJSON() ([]byte, error) {
	if t.CallbackQueryPayloadData != nil {
		return json.Marshal(t.CallbackQueryPayloadData)
	}
	if t.CallbackQueryPayloadDataWithPassword != nil {
		return json.Marshal(t.CallbackQueryPayloadDataWithPassword)
	}
	if t.CallbackQueryPayloadGame != nil {
		return json.Marshal(t.CallbackQueryPayloadGame)
	}
	return nil, fmt.Errorf("no value set for CallbackQueryPayload")
}

// CallDiscardReason Describes the reason why a call was discarded
type CallDiscardReason struct {
	TypeStr                             string                               `json:"@type"`
	CallDiscardReasonDeclined           *CallDiscardReasonDeclined           `json:"callDiscardReasonDeclined,omitempty"`
	CallDiscardReasonDisconnected       *CallDiscardReasonDisconnected       `json:"callDiscardReasonDisconnected,omitempty"`
	CallDiscardReasonEmpty              *CallDiscardReasonEmpty              `json:"callDiscardReasonEmpty,omitempty"`
	CallDiscardReasonHungUp             *CallDiscardReasonHungUp             `json:"callDiscardReasonHungUp,omitempty"`
	CallDiscardReasonMissed             *CallDiscardReasonMissed             `json:"callDiscardReasonMissed,omitempty"`
	CallDiscardReasonUpgradeToGroupCall *CallDiscardReasonUpgradeToGroupCall `json:"callDiscardReasonUpgradeToGroupCall,omitempty"`
}

func (t *CallDiscardReason) Type() string {
	return t.TypeStr
}

func (t *CallDiscardReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CallDiscardReason) GetExtra() string {
	return ""
}

func (t *CallDiscardReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "callDiscardReasonDeclined":
		t.CallDiscardReasonDeclined = new(CallDiscardReasonDeclined)
		return json.Unmarshal(b, t.CallDiscardReasonDeclined)
	case "callDiscardReasonDisconnected":
		t.CallDiscardReasonDisconnected = new(CallDiscardReasonDisconnected)
		return json.Unmarshal(b, t.CallDiscardReasonDisconnected)
	case "callDiscardReasonEmpty":
		t.CallDiscardReasonEmpty = new(CallDiscardReasonEmpty)
		return json.Unmarshal(b, t.CallDiscardReasonEmpty)
	case "callDiscardReasonHungUp":
		t.CallDiscardReasonHungUp = new(CallDiscardReasonHungUp)
		return json.Unmarshal(b, t.CallDiscardReasonHungUp)
	case "callDiscardReasonMissed":
		t.CallDiscardReasonMissed = new(CallDiscardReasonMissed)
		return json.Unmarshal(b, t.CallDiscardReasonMissed)
	case "callDiscardReasonUpgradeToGroupCall":
		t.CallDiscardReasonUpgradeToGroupCall = new(CallDiscardReasonUpgradeToGroupCall)
		return json.Unmarshal(b, t.CallDiscardReasonUpgradeToGroupCall)
	}
	return nil
}

func (t *CallDiscardReason) MarshalJSON() ([]byte, error) {
	if t.CallDiscardReasonDeclined != nil {
		return json.Marshal(t.CallDiscardReasonDeclined)
	}
	if t.CallDiscardReasonDisconnected != nil {
		return json.Marshal(t.CallDiscardReasonDisconnected)
	}
	if t.CallDiscardReasonEmpty != nil {
		return json.Marshal(t.CallDiscardReasonEmpty)
	}
	if t.CallDiscardReasonHungUp != nil {
		return json.Marshal(t.CallDiscardReasonHungUp)
	}
	if t.CallDiscardReasonMissed != nil {
		return json.Marshal(t.CallDiscardReasonMissed)
	}
	if t.CallDiscardReasonUpgradeToGroupCall != nil {
		return json.Marshal(t.CallDiscardReasonUpgradeToGroupCall)
	}
	return nil, fmt.Errorf("no value set for CallDiscardReason")
}

// CallProblem Describes the exact type of problem with a call
type CallProblem struct {
	TypeStr                    string                      `json:"@type"`
	CallProblemDistortedSpeech *CallProblemDistortedSpeech `json:"callProblemDistortedSpeech,omitempty"`
	CallProblemDistortedVideo  *CallProblemDistortedVideo  `json:"callProblemDistortedVideo,omitempty"`
	CallProblemDropped         *CallProblemDropped         `json:"callProblemDropped,omitempty"`
	CallProblemEcho            *CallProblemEcho            `json:"callProblemEcho,omitempty"`
	CallProblemInterruptions   *CallProblemInterruptions   `json:"callProblemInterruptions,omitempty"`
	CallProblemNoise           *CallProblemNoise           `json:"callProblemNoise,omitempty"`
	CallProblemPixelatedVideo  *CallProblemPixelatedVideo  `json:"callProblemPixelatedVideo,omitempty"`
	CallProblemSilentLocal     *CallProblemSilentLocal     `json:"callProblemSilentLocal,omitempty"`
	CallProblemSilentRemote    *CallProblemSilentRemote    `json:"callProblemSilentRemote,omitempty"`
}

func (t *CallProblem) Type() string {
	return t.TypeStr
}

func (t *CallProblem) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CallProblem) GetExtra() string {
	return ""
}

func (t *CallProblem) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "callProblemDistortedSpeech":
		t.CallProblemDistortedSpeech = new(CallProblemDistortedSpeech)
		return json.Unmarshal(b, t.CallProblemDistortedSpeech)
	case "callProblemDistortedVideo":
		t.CallProblemDistortedVideo = new(CallProblemDistortedVideo)
		return json.Unmarshal(b, t.CallProblemDistortedVideo)
	case "callProblemDropped":
		t.CallProblemDropped = new(CallProblemDropped)
		return json.Unmarshal(b, t.CallProblemDropped)
	case "callProblemEcho":
		t.CallProblemEcho = new(CallProblemEcho)
		return json.Unmarshal(b, t.CallProblemEcho)
	case "callProblemInterruptions":
		t.CallProblemInterruptions = new(CallProblemInterruptions)
		return json.Unmarshal(b, t.CallProblemInterruptions)
	case "callProblemNoise":
		t.CallProblemNoise = new(CallProblemNoise)
		return json.Unmarshal(b, t.CallProblemNoise)
	case "callProblemPixelatedVideo":
		t.CallProblemPixelatedVideo = new(CallProblemPixelatedVideo)
		return json.Unmarshal(b, t.CallProblemPixelatedVideo)
	case "callProblemSilentLocal":
		t.CallProblemSilentLocal = new(CallProblemSilentLocal)
		return json.Unmarshal(b, t.CallProblemSilentLocal)
	case "callProblemSilentRemote":
		t.CallProblemSilentRemote = new(CallProblemSilentRemote)
		return json.Unmarshal(b, t.CallProblemSilentRemote)
	}
	return nil
}

func (t *CallProblem) MarshalJSON() ([]byte, error) {
	if t.CallProblemDistortedSpeech != nil {
		return json.Marshal(t.CallProblemDistortedSpeech)
	}
	if t.CallProblemDistortedVideo != nil {
		return json.Marshal(t.CallProblemDistortedVideo)
	}
	if t.CallProblemDropped != nil {
		return json.Marshal(t.CallProblemDropped)
	}
	if t.CallProblemEcho != nil {
		return json.Marshal(t.CallProblemEcho)
	}
	if t.CallProblemInterruptions != nil {
		return json.Marshal(t.CallProblemInterruptions)
	}
	if t.CallProblemNoise != nil {
		return json.Marshal(t.CallProblemNoise)
	}
	if t.CallProblemPixelatedVideo != nil {
		return json.Marshal(t.CallProblemPixelatedVideo)
	}
	if t.CallProblemSilentLocal != nil {
		return json.Marshal(t.CallProblemSilentLocal)
	}
	if t.CallProblemSilentRemote != nil {
		return json.Marshal(t.CallProblemSilentRemote)
	}
	return nil, fmt.Errorf("no value set for CallProblem")
}

// CallServerType Describes the type of call server
type CallServerType struct {
	TypeStr                         string                           `json:"@type"`
	CallServerTypeTelegramReflector *CallServerTypeTelegramReflector `json:"callServerTypeTelegramReflector,omitempty"`
	CallServerTypeWebrtc            *CallServerTypeWebrtc            `json:"callServerTypeWebrtc,omitempty"`
}

func (t *CallServerType) Type() string {
	return t.TypeStr
}

func (t *CallServerType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CallServerType) GetExtra() string {
	return ""
}

func (t *CallServerType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "callServerTypeTelegramReflector":
		t.CallServerTypeTelegramReflector = new(CallServerTypeTelegramReflector)
		return json.Unmarshal(b, t.CallServerTypeTelegramReflector)
	case "callServerTypeWebrtc":
		t.CallServerTypeWebrtc = new(CallServerTypeWebrtc)
		return json.Unmarshal(b, t.CallServerTypeWebrtc)
	}
	return nil
}

func (t *CallServerType) MarshalJSON() ([]byte, error) {
	if t.CallServerTypeTelegramReflector != nil {
		return json.Marshal(t.CallServerTypeTelegramReflector)
	}
	if t.CallServerTypeWebrtc != nil {
		return json.Marshal(t.CallServerTypeWebrtc)
	}
	return nil, fmt.Errorf("no value set for CallServerType")
}

// CallState Describes the current call state
type CallState struct {
	TypeStr                 string                   `json:"@type"`
	CallStateDiscarded      *CallStateDiscarded      `json:"callStateDiscarded,omitempty"`
	CallStateError          *CallStateError          `json:"callStateError,omitempty"`
	CallStateExchangingKeys *CallStateExchangingKeys `json:"callStateExchangingKeys,omitempty"`
	CallStateHangingUp      *CallStateHangingUp      `json:"callStateHangingUp,omitempty"`
	CallStatePending        *CallStatePending        `json:"callStatePending,omitempty"`
	CallStateReady          *CallStateReady          `json:"callStateReady,omitempty"`
}

func (t *CallState) Type() string {
	return t.TypeStr
}

func (t *CallState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CallState) GetExtra() string {
	return ""
}

func (t *CallState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "callStateDiscarded":
		t.CallStateDiscarded = new(CallStateDiscarded)
		return json.Unmarshal(b, t.CallStateDiscarded)
	case "callStateError":
		t.CallStateError = new(CallStateError)
		return json.Unmarshal(b, t.CallStateError)
	case "callStateExchangingKeys":
		t.CallStateExchangingKeys = new(CallStateExchangingKeys)
		return json.Unmarshal(b, t.CallStateExchangingKeys)
	case "callStateHangingUp":
		t.CallStateHangingUp = new(CallStateHangingUp)
		return json.Unmarshal(b, t.CallStateHangingUp)
	case "callStatePending":
		t.CallStatePending = new(CallStatePending)
		return json.Unmarshal(b, t.CallStatePending)
	case "callStateReady":
		t.CallStateReady = new(CallStateReady)
		return json.Unmarshal(b, t.CallStateReady)
	}
	return nil
}

func (t *CallState) MarshalJSON() ([]byte, error) {
	if t.CallStateDiscarded != nil {
		return json.Marshal(t.CallStateDiscarded)
	}
	if t.CallStateError != nil {
		return json.Marshal(t.CallStateError)
	}
	if t.CallStateExchangingKeys != nil {
		return json.Marshal(t.CallStateExchangingKeys)
	}
	if t.CallStateHangingUp != nil {
		return json.Marshal(t.CallStateHangingUp)
	}
	if t.CallStatePending != nil {
		return json.Marshal(t.CallStatePending)
	}
	if t.CallStateReady != nil {
		return json.Marshal(t.CallStateReady)
	}
	return nil, fmt.Errorf("no value set for CallState")
}

// CanPostStoryResult Represents result of checking whether the current user can post a story on behalf of the specific chat
type CanPostStoryResult struct {
	TypeStr                                    string                                      `json:"@type"`
	CanPostStoryResultActiveStoryLimitExceeded *CanPostStoryResultActiveStoryLimitExceeded `json:"canPostStoryResultActiveStoryLimitExceeded,omitempty"`
	CanPostStoryResultBoostNeeded              *CanPostStoryResultBoostNeeded              `json:"canPostStoryResultBoostNeeded,omitempty"`
	CanPostStoryResultLiveStoryIsActive        *CanPostStoryResultLiveStoryIsActive        `json:"canPostStoryResultLiveStoryIsActive,omitempty"`
	CanPostStoryResultMonthlyLimitExceeded     *CanPostStoryResultMonthlyLimitExceeded     `json:"canPostStoryResultMonthlyLimitExceeded,omitempty"`
	CanPostStoryResultOk                       *CanPostStoryResultOk                       `json:"canPostStoryResultOk,omitempty"`
	CanPostStoryResultPremiumNeeded            *CanPostStoryResultPremiumNeeded            `json:"canPostStoryResultPremiumNeeded,omitempty"`
	CanPostStoryResultWeeklyLimitExceeded      *CanPostStoryResultWeeklyLimitExceeded      `json:"canPostStoryResultWeeklyLimitExceeded,omitempty"`
}

func (t *CanPostStoryResult) Type() string {
	return t.TypeStr
}

func (t *CanPostStoryResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CanPostStoryResult) GetExtra() string {
	return ""
}

func (t *CanPostStoryResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "canPostStoryResultActiveStoryLimitExceeded":
		t.CanPostStoryResultActiveStoryLimitExceeded = new(CanPostStoryResultActiveStoryLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultActiveStoryLimitExceeded)
	case "canPostStoryResultBoostNeeded":
		t.CanPostStoryResultBoostNeeded = new(CanPostStoryResultBoostNeeded)
		return json.Unmarshal(b, t.CanPostStoryResultBoostNeeded)
	case "canPostStoryResultLiveStoryIsActive":
		t.CanPostStoryResultLiveStoryIsActive = new(CanPostStoryResultLiveStoryIsActive)
		return json.Unmarshal(b, t.CanPostStoryResultLiveStoryIsActive)
	case "canPostStoryResultMonthlyLimitExceeded":
		t.CanPostStoryResultMonthlyLimitExceeded = new(CanPostStoryResultMonthlyLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultMonthlyLimitExceeded)
	case "canPostStoryResultOk":
		t.CanPostStoryResultOk = new(CanPostStoryResultOk)
		return json.Unmarshal(b, t.CanPostStoryResultOk)
	case "canPostStoryResultPremiumNeeded":
		t.CanPostStoryResultPremiumNeeded = new(CanPostStoryResultPremiumNeeded)
		return json.Unmarshal(b, t.CanPostStoryResultPremiumNeeded)
	case "canPostStoryResultWeeklyLimitExceeded":
		t.CanPostStoryResultWeeklyLimitExceeded = new(CanPostStoryResultWeeklyLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultWeeklyLimitExceeded)
	}
	return nil
}

func (t *CanPostStoryResult) MarshalJSON() ([]byte, error) {
	if t.CanPostStoryResultActiveStoryLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultActiveStoryLimitExceeded)
	}
	if t.CanPostStoryResultBoostNeeded != nil {
		return json.Marshal(t.CanPostStoryResultBoostNeeded)
	}
	if t.CanPostStoryResultLiveStoryIsActive != nil {
		return json.Marshal(t.CanPostStoryResultLiveStoryIsActive)
	}
	if t.CanPostStoryResultMonthlyLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultMonthlyLimitExceeded)
	}
	if t.CanPostStoryResultOk != nil {
		return json.Marshal(t.CanPostStoryResultOk)
	}
	if t.CanPostStoryResultPremiumNeeded != nil {
		return json.Marshal(t.CanPostStoryResultPremiumNeeded)
	}
	if t.CanPostStoryResultWeeklyLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultWeeklyLimitExceeded)
	}
	return nil, fmt.Errorf("no value set for CanPostStoryResult")
}

// CanSendGiftResult Describes whether a gift can be sent now by the current user
type CanSendGiftResult struct {
	TypeStr               string                 `json:"@type"`
	CanSendGiftResultFail *CanSendGiftResultFail `json:"canSendGiftResultFail,omitempty"`
	CanSendGiftResultOk   *CanSendGiftResultOk   `json:"canSendGiftResultOk,omitempty"`
}

func (t *CanSendGiftResult) Type() string {
	return t.TypeStr
}

func (t *CanSendGiftResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CanSendGiftResult) GetExtra() string {
	return ""
}

func (t *CanSendGiftResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "canSendGiftResultFail":
		t.CanSendGiftResultFail = new(CanSendGiftResultFail)
		return json.Unmarshal(b, t.CanSendGiftResultFail)
	case "canSendGiftResultOk":
		t.CanSendGiftResultOk = new(CanSendGiftResultOk)
		return json.Unmarshal(b, t.CanSendGiftResultOk)
	}
	return nil
}

func (t *CanSendGiftResult) MarshalJSON() ([]byte, error) {
	if t.CanSendGiftResultFail != nil {
		return json.Marshal(t.CanSendGiftResultFail)
	}
	if t.CanSendGiftResultOk != nil {
		return json.Marshal(t.CanSendGiftResultOk)
	}
	return nil, fmt.Errorf("no value set for CanSendGiftResult")
}

// CanSendMessageToUserResult Describes result of canSendMessageToUser
type CanSendMessageToUserResult struct {
	TypeStr                                         string                                           `json:"@type"`
	CanSendMessageToUserResultOk                    *CanSendMessageToUserResultOk                    `json:"canSendMessageToUserResultOk,omitempty"`
	CanSendMessageToUserResultUserHasPaidMessages   *CanSendMessageToUserResultUserHasPaidMessages   `json:"canSendMessageToUserResultUserHasPaidMessages,omitempty"`
	CanSendMessageToUserResultUserIsDeleted         *CanSendMessageToUserResultUserIsDeleted         `json:"canSendMessageToUserResultUserIsDeleted,omitempty"`
	CanSendMessageToUserResultUserRestrictsNewChats *CanSendMessageToUserResultUserRestrictsNewChats `json:"canSendMessageToUserResultUserRestrictsNewChats,omitempty"`
}

func (t *CanSendMessageToUserResult) Type() string {
	return t.TypeStr
}

func (t *CanSendMessageToUserResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CanSendMessageToUserResult) GetExtra() string {
	return ""
}

func (t *CanSendMessageToUserResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "canSendMessageToUserResultOk":
		t.CanSendMessageToUserResultOk = new(CanSendMessageToUserResultOk)
		return json.Unmarshal(b, t.CanSendMessageToUserResultOk)
	case "canSendMessageToUserResultUserHasPaidMessages":
		t.CanSendMessageToUserResultUserHasPaidMessages = new(CanSendMessageToUserResultUserHasPaidMessages)
		return json.Unmarshal(b, t.CanSendMessageToUserResultUserHasPaidMessages)
	case "canSendMessageToUserResultUserIsDeleted":
		t.CanSendMessageToUserResultUserIsDeleted = new(CanSendMessageToUserResultUserIsDeleted)
		return json.Unmarshal(b, t.CanSendMessageToUserResultUserIsDeleted)
	case "canSendMessageToUserResultUserRestrictsNewChats":
		t.CanSendMessageToUserResultUserRestrictsNewChats = new(CanSendMessageToUserResultUserRestrictsNewChats)
		return json.Unmarshal(b, t.CanSendMessageToUserResultUserRestrictsNewChats)
	}
	return nil
}

func (t *CanSendMessageToUserResult) MarshalJSON() ([]byte, error) {
	if t.CanSendMessageToUserResultOk != nil {
		return json.Marshal(t.CanSendMessageToUserResultOk)
	}
	if t.CanSendMessageToUserResultUserHasPaidMessages != nil {
		return json.Marshal(t.CanSendMessageToUserResultUserHasPaidMessages)
	}
	if t.CanSendMessageToUserResultUserIsDeleted != nil {
		return json.Marshal(t.CanSendMessageToUserResultUserIsDeleted)
	}
	if t.CanSendMessageToUserResultUserRestrictsNewChats != nil {
		return json.Marshal(t.CanSendMessageToUserResultUserRestrictsNewChats)
	}
	return nil, fmt.Errorf("no value set for CanSendMessageToUserResult")
}

// CanTransferOwnershipResult Represents result of checking whether the current session can be used to transfer a chat ownership to another user
type CanTransferOwnershipResult struct {
	TypeStr                                    string                                      `json:"@type"`
	CanTransferOwnershipResultOk               *CanTransferOwnershipResultOk               `json:"canTransferOwnershipResultOk,omitempty"`
	CanTransferOwnershipResultPasswordNeeded   *CanTransferOwnershipResultPasswordNeeded   `json:"canTransferOwnershipResultPasswordNeeded,omitempty"`
	CanTransferOwnershipResultPasswordTooFresh *CanTransferOwnershipResultPasswordTooFresh `json:"canTransferOwnershipResultPasswordTooFresh,omitempty"`
	CanTransferOwnershipResultSessionTooFresh  *CanTransferOwnershipResultSessionTooFresh  `json:"canTransferOwnershipResultSessionTooFresh,omitempty"`
}

func (t *CanTransferOwnershipResult) Type() string {
	return t.TypeStr
}

func (t *CanTransferOwnershipResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CanTransferOwnershipResult) GetExtra() string {
	return ""
}

func (t *CanTransferOwnershipResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "canTransferOwnershipResultOk":
		t.CanTransferOwnershipResultOk = new(CanTransferOwnershipResultOk)
		return json.Unmarshal(b, t.CanTransferOwnershipResultOk)
	case "canTransferOwnershipResultPasswordNeeded":
		t.CanTransferOwnershipResultPasswordNeeded = new(CanTransferOwnershipResultPasswordNeeded)
		return json.Unmarshal(b, t.CanTransferOwnershipResultPasswordNeeded)
	case "canTransferOwnershipResultPasswordTooFresh":
		t.CanTransferOwnershipResultPasswordTooFresh = new(CanTransferOwnershipResultPasswordTooFresh)
		return json.Unmarshal(b, t.CanTransferOwnershipResultPasswordTooFresh)
	case "canTransferOwnershipResultSessionTooFresh":
		t.CanTransferOwnershipResultSessionTooFresh = new(CanTransferOwnershipResultSessionTooFresh)
		return json.Unmarshal(b, t.CanTransferOwnershipResultSessionTooFresh)
	}
	return nil
}

func (t *CanTransferOwnershipResult) MarshalJSON() ([]byte, error) {
	if t.CanTransferOwnershipResultOk != nil {
		return json.Marshal(t.CanTransferOwnershipResultOk)
	}
	if t.CanTransferOwnershipResultPasswordNeeded != nil {
		return json.Marshal(t.CanTransferOwnershipResultPasswordNeeded)
	}
	if t.CanTransferOwnershipResultPasswordTooFresh != nil {
		return json.Marshal(t.CanTransferOwnershipResultPasswordTooFresh)
	}
	if t.CanTransferOwnershipResultSessionTooFresh != nil {
		return json.Marshal(t.CanTransferOwnershipResultSessionTooFresh)
	}
	return nil, fmt.Errorf("no value set for CanTransferOwnershipResult")
}

// ChatAction Describes the different types of activity in a chat
type ChatAction struct {
	TypeStr                      string                        `json:"@type"`
	ChatActionCancel             *ChatActionCancel             `json:"chatActionCancel,omitempty"`
	ChatActionChoosingContact    *ChatActionChoosingContact    `json:"chatActionChoosingContact,omitempty"`
	ChatActionChoosingLocation   *ChatActionChoosingLocation   `json:"chatActionChoosingLocation,omitempty"`
	ChatActionChoosingSticker    *ChatActionChoosingSticker    `json:"chatActionChoosingSticker,omitempty"`
	ChatActionRecordingVideo     *ChatActionRecordingVideo     `json:"chatActionRecordingVideo,omitempty"`
	ChatActionRecordingVideoNote *ChatActionRecordingVideoNote `json:"chatActionRecordingVideoNote,omitempty"`
	ChatActionRecordingVoiceNote *ChatActionRecordingVoiceNote `json:"chatActionRecordingVoiceNote,omitempty"`
	ChatActionStartPlayingGame   *ChatActionStartPlayingGame   `json:"chatActionStartPlayingGame,omitempty"`
	ChatActionTyping             *ChatActionTyping             `json:"chatActionTyping,omitempty"`
	ChatActionUploadingDocument  *ChatActionUploadingDocument  `json:"chatActionUploadingDocument,omitempty"`
	ChatActionUploadingPhoto     *ChatActionUploadingPhoto     `json:"chatActionUploadingPhoto,omitempty"`
	ChatActionUploadingVideo     *ChatActionUploadingVideo     `json:"chatActionUploadingVideo,omitempty"`
	ChatActionUploadingVideoNote *ChatActionUploadingVideoNote `json:"chatActionUploadingVideoNote,omitempty"`
	ChatActionUploadingVoiceNote *ChatActionUploadingVoiceNote `json:"chatActionUploadingVoiceNote,omitempty"`
	ChatActionWatchingAnimations *ChatActionWatchingAnimations `json:"chatActionWatchingAnimations,omitempty"`
}

func (t *ChatAction) Type() string {
	return t.TypeStr
}

func (t *ChatAction) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatAction) GetExtra() string {
	return ""
}

func (t *ChatAction) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatActionCancel":
		t.ChatActionCancel = new(ChatActionCancel)
		return json.Unmarshal(b, t.ChatActionCancel)
	case "chatActionChoosingContact":
		t.ChatActionChoosingContact = new(ChatActionChoosingContact)
		return json.Unmarshal(b, t.ChatActionChoosingContact)
	case "chatActionChoosingLocation":
		t.ChatActionChoosingLocation = new(ChatActionChoosingLocation)
		return json.Unmarshal(b, t.ChatActionChoosingLocation)
	case "chatActionChoosingSticker":
		t.ChatActionChoosingSticker = new(ChatActionChoosingSticker)
		return json.Unmarshal(b, t.ChatActionChoosingSticker)
	case "chatActionRecordingVideo":
		t.ChatActionRecordingVideo = new(ChatActionRecordingVideo)
		return json.Unmarshal(b, t.ChatActionRecordingVideo)
	case "chatActionRecordingVideoNote":
		t.ChatActionRecordingVideoNote = new(ChatActionRecordingVideoNote)
		return json.Unmarshal(b, t.ChatActionRecordingVideoNote)
	case "chatActionRecordingVoiceNote":
		t.ChatActionRecordingVoiceNote = new(ChatActionRecordingVoiceNote)
		return json.Unmarshal(b, t.ChatActionRecordingVoiceNote)
	case "chatActionStartPlayingGame":
		t.ChatActionStartPlayingGame = new(ChatActionStartPlayingGame)
		return json.Unmarshal(b, t.ChatActionStartPlayingGame)
	case "chatActionTyping":
		t.ChatActionTyping = new(ChatActionTyping)
		return json.Unmarshal(b, t.ChatActionTyping)
	case "chatActionUploadingDocument":
		t.ChatActionUploadingDocument = new(ChatActionUploadingDocument)
		return json.Unmarshal(b, t.ChatActionUploadingDocument)
	case "chatActionUploadingPhoto":
		t.ChatActionUploadingPhoto = new(ChatActionUploadingPhoto)
		return json.Unmarshal(b, t.ChatActionUploadingPhoto)
	case "chatActionUploadingVideo":
		t.ChatActionUploadingVideo = new(ChatActionUploadingVideo)
		return json.Unmarshal(b, t.ChatActionUploadingVideo)
	case "chatActionUploadingVideoNote":
		t.ChatActionUploadingVideoNote = new(ChatActionUploadingVideoNote)
		return json.Unmarshal(b, t.ChatActionUploadingVideoNote)
	case "chatActionUploadingVoiceNote":
		t.ChatActionUploadingVoiceNote = new(ChatActionUploadingVoiceNote)
		return json.Unmarshal(b, t.ChatActionUploadingVoiceNote)
	case "chatActionWatchingAnimations":
		t.ChatActionWatchingAnimations = new(ChatActionWatchingAnimations)
		return json.Unmarshal(b, t.ChatActionWatchingAnimations)
	}
	return nil
}

func (t *ChatAction) MarshalJSON() ([]byte, error) {
	if t.ChatActionCancel != nil {
		return json.Marshal(t.ChatActionCancel)
	}
	if t.ChatActionChoosingContact != nil {
		return json.Marshal(t.ChatActionChoosingContact)
	}
	if t.ChatActionChoosingLocation != nil {
		return json.Marshal(t.ChatActionChoosingLocation)
	}
	if t.ChatActionChoosingSticker != nil {
		return json.Marshal(t.ChatActionChoosingSticker)
	}
	if t.ChatActionRecordingVideo != nil {
		return json.Marshal(t.ChatActionRecordingVideo)
	}
	if t.ChatActionRecordingVideoNote != nil {
		return json.Marshal(t.ChatActionRecordingVideoNote)
	}
	if t.ChatActionRecordingVoiceNote != nil {
		return json.Marshal(t.ChatActionRecordingVoiceNote)
	}
	if t.ChatActionStartPlayingGame != nil {
		return json.Marshal(t.ChatActionStartPlayingGame)
	}
	if t.ChatActionTyping != nil {
		return json.Marshal(t.ChatActionTyping)
	}
	if t.ChatActionUploadingDocument != nil {
		return json.Marshal(t.ChatActionUploadingDocument)
	}
	if t.ChatActionUploadingPhoto != nil {
		return json.Marshal(t.ChatActionUploadingPhoto)
	}
	if t.ChatActionUploadingVideo != nil {
		return json.Marshal(t.ChatActionUploadingVideo)
	}
	if t.ChatActionUploadingVideoNote != nil {
		return json.Marshal(t.ChatActionUploadingVideoNote)
	}
	if t.ChatActionUploadingVoiceNote != nil {
		return json.Marshal(t.ChatActionUploadingVoiceNote)
	}
	if t.ChatActionWatchingAnimations != nil {
		return json.Marshal(t.ChatActionWatchingAnimations)
	}
	return nil, fmt.Errorf("no value set for ChatAction")
}

// ChatActionBar Describes actions which must be possible to do through a chat action bar
type ChatActionBar struct {
	TypeStr                       string                         `json:"@type"`
	ChatActionBarAddContact       *ChatActionBarAddContact       `json:"chatActionBarAddContact,omitempty"`
	ChatActionBarInviteMembers    *ChatActionBarInviteMembers    `json:"chatActionBarInviteMembers,omitempty"`
	ChatActionBarJoinRequest      *ChatActionBarJoinRequest      `json:"chatActionBarJoinRequest,omitempty"`
	ChatActionBarReportAddBlock   *ChatActionBarReportAddBlock   `json:"chatActionBarReportAddBlock,omitempty"`
	ChatActionBarReportSpam       *ChatActionBarReportSpam       `json:"chatActionBarReportSpam,omitempty"`
	ChatActionBarSharePhoneNumber *ChatActionBarSharePhoneNumber `json:"chatActionBarSharePhoneNumber,omitempty"`
}

func (t *ChatActionBar) Type() string {
	return t.TypeStr
}

func (t *ChatActionBar) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatActionBar) GetExtra() string {
	return ""
}

func (t *ChatActionBar) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatActionBarAddContact":
		t.ChatActionBarAddContact = new(ChatActionBarAddContact)
		return json.Unmarshal(b, t.ChatActionBarAddContact)
	case "chatActionBarInviteMembers":
		t.ChatActionBarInviteMembers = new(ChatActionBarInviteMembers)
		return json.Unmarshal(b, t.ChatActionBarInviteMembers)
	case "chatActionBarJoinRequest":
		t.ChatActionBarJoinRequest = new(ChatActionBarJoinRequest)
		return json.Unmarshal(b, t.ChatActionBarJoinRequest)
	case "chatActionBarReportAddBlock":
		t.ChatActionBarReportAddBlock = new(ChatActionBarReportAddBlock)
		return json.Unmarshal(b, t.ChatActionBarReportAddBlock)
	case "chatActionBarReportSpam":
		t.ChatActionBarReportSpam = new(ChatActionBarReportSpam)
		return json.Unmarshal(b, t.ChatActionBarReportSpam)
	case "chatActionBarSharePhoneNumber":
		t.ChatActionBarSharePhoneNumber = new(ChatActionBarSharePhoneNumber)
		return json.Unmarshal(b, t.ChatActionBarSharePhoneNumber)
	}
	return nil
}

func (t *ChatActionBar) MarshalJSON() ([]byte, error) {
	if t.ChatActionBarAddContact != nil {
		return json.Marshal(t.ChatActionBarAddContact)
	}
	if t.ChatActionBarInviteMembers != nil {
		return json.Marshal(t.ChatActionBarInviteMembers)
	}
	if t.ChatActionBarJoinRequest != nil {
		return json.Marshal(t.ChatActionBarJoinRequest)
	}
	if t.ChatActionBarReportAddBlock != nil {
		return json.Marshal(t.ChatActionBarReportAddBlock)
	}
	if t.ChatActionBarReportSpam != nil {
		return json.Marshal(t.ChatActionBarReportSpam)
	}
	if t.ChatActionBarSharePhoneNumber != nil {
		return json.Marshal(t.ChatActionBarSharePhoneNumber)
	}
	return nil, fmt.Errorf("no value set for ChatActionBar")
}

// ChatAvailableReactions Describes reactions available in the chat
type ChatAvailableReactions struct {
	TypeStr                    string                      `json:"@type"`
	ChatAvailableReactionsAll  *ChatAvailableReactionsAll  `json:"chatAvailableReactionsAll,omitempty"`
	ChatAvailableReactionsSome *ChatAvailableReactionsSome `json:"chatAvailableReactionsSome,omitempty"`
}

func (t *ChatAvailableReactions) Type() string {
	return t.TypeStr
}

func (t *ChatAvailableReactions) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatAvailableReactions) GetExtra() string {
	return ""
}

func (t *ChatAvailableReactions) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatAvailableReactionsAll":
		t.ChatAvailableReactionsAll = new(ChatAvailableReactionsAll)
		return json.Unmarshal(b, t.ChatAvailableReactionsAll)
	case "chatAvailableReactionsSome":
		t.ChatAvailableReactionsSome = new(ChatAvailableReactionsSome)
		return json.Unmarshal(b, t.ChatAvailableReactionsSome)
	}
	return nil
}

func (t *ChatAvailableReactions) MarshalJSON() ([]byte, error) {
	if t.ChatAvailableReactionsAll != nil {
		return json.Marshal(t.ChatAvailableReactionsAll)
	}
	if t.ChatAvailableReactionsSome != nil {
		return json.Marshal(t.ChatAvailableReactionsSome)
	}
	return nil, fmt.Errorf("no value set for ChatAvailableReactions")
}

// ChatBoostSource Describes source of a chat boost
type ChatBoostSource struct {
	TypeStr                 string                   `json:"@type"`
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode `json:"chatBoostSourceGiftCode,omitempty"`
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway `json:"chatBoostSourceGiveaway,omitempty"`
	ChatBoostSourcePremium  *ChatBoostSourcePremium  `json:"chatBoostSourcePremium,omitempty"`
}

func (t *ChatBoostSource) Type() string {
	return t.TypeStr
}

func (t *ChatBoostSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatBoostSource) GetExtra() string {
	return ""
}

func (t *ChatBoostSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatBoostSourceGiftCode":
		t.ChatBoostSourceGiftCode = new(ChatBoostSourceGiftCode)
		return json.Unmarshal(b, t.ChatBoostSourceGiftCode)
	case "chatBoostSourceGiveaway":
		t.ChatBoostSourceGiveaway = new(ChatBoostSourceGiveaway)
		return json.Unmarshal(b, t.ChatBoostSourceGiveaway)
	case "chatBoostSourcePremium":
		t.ChatBoostSourcePremium = new(ChatBoostSourcePremium)
		return json.Unmarshal(b, t.ChatBoostSourcePremium)
	}
	return nil
}

func (t *ChatBoostSource) MarshalJSON() ([]byte, error) {
	if t.ChatBoostSourceGiftCode != nil {
		return json.Marshal(t.ChatBoostSourceGiftCode)
	}
	if t.ChatBoostSourceGiveaway != nil {
		return json.Marshal(t.ChatBoostSourceGiveaway)
	}
	if t.ChatBoostSourcePremium != nil {
		return json.Marshal(t.ChatBoostSourcePremium)
	}
	return nil, fmt.Errorf("no value set for ChatBoostSource")
}

// ChatEventAction Represents a chat event
type ChatEventAction struct {
	TypeStr                                         string                                           `json:"@type"`
	ChatEventAccentColorChanged                     *ChatEventAccentColorChanged                     `json:"chatEventAccentColorChanged,omitempty"`
	ChatEventActiveUsernamesChanged                 *ChatEventActiveUsernamesChanged                 `json:"chatEventActiveUsernamesChanged,omitempty"`
	ChatEventAutomaticTranslationToggled            *ChatEventAutomaticTranslationToggled            `json:"chatEventAutomaticTranslationToggled,omitempty"`
	ChatEventAvailableReactionsChanged              *ChatEventAvailableReactionsChanged              `json:"chatEventAvailableReactionsChanged,omitempty"`
	ChatEventBackgroundChanged                      *ChatEventBackgroundChanged                      `json:"chatEventBackgroundChanged,omitempty"`
	ChatEventCustomEmojiStickerSetChanged           *ChatEventCustomEmojiStickerSetChanged           `json:"chatEventCustomEmojiStickerSetChanged,omitempty"`
	ChatEventDescriptionChanged                     *ChatEventDescriptionChanged                     `json:"chatEventDescriptionChanged,omitempty"`
	ChatEventEmojiStatusChanged                     *ChatEventEmojiStatusChanged                     `json:"chatEventEmojiStatusChanged,omitempty"`
	ChatEventForumTopicCreated                      *ChatEventForumTopicCreated                      `json:"chatEventForumTopicCreated,omitempty"`
	ChatEventForumTopicDeleted                      *ChatEventForumTopicDeleted                      `json:"chatEventForumTopicDeleted,omitempty"`
	ChatEventForumTopicEdited                       *ChatEventForumTopicEdited                       `json:"chatEventForumTopicEdited,omitempty"`
	ChatEventForumTopicPinned                       *ChatEventForumTopicPinned                       `json:"chatEventForumTopicPinned,omitempty"`
	ChatEventForumTopicToggleIsClosed               *ChatEventForumTopicToggleIsClosed               `json:"chatEventForumTopicToggleIsClosed,omitempty"`
	ChatEventForumTopicToggleIsHidden               *ChatEventForumTopicToggleIsHidden               `json:"chatEventForumTopicToggleIsHidden,omitempty"`
	ChatEventHasAggressiveAntiSpamEnabledToggled    *ChatEventHasAggressiveAntiSpamEnabledToggled    `json:"chatEventHasAggressiveAntiSpamEnabledToggled,omitempty"`
	ChatEventHasProtectedContentToggled             *ChatEventHasProtectedContentToggled             `json:"chatEventHasProtectedContentToggled,omitempty"`
	ChatEventInviteLinkDeleted                      *ChatEventInviteLinkDeleted                      `json:"chatEventInviteLinkDeleted,omitempty"`
	ChatEventInviteLinkEdited                       *ChatEventInviteLinkEdited                       `json:"chatEventInviteLinkEdited,omitempty"`
	ChatEventInviteLinkRevoked                      *ChatEventInviteLinkRevoked                      `json:"chatEventInviteLinkRevoked,omitempty"`
	ChatEventInvitesToggled                         *ChatEventInvitesToggled                         `json:"chatEventInvitesToggled,omitempty"`
	ChatEventIsAllHistoryAvailableToggled           *ChatEventIsAllHistoryAvailableToggled           `json:"chatEventIsAllHistoryAvailableToggled,omitempty"`
	ChatEventIsForumToggled                         *ChatEventIsForumToggled                         `json:"chatEventIsForumToggled,omitempty"`
	ChatEventLinkedChatChanged                      *ChatEventLinkedChatChanged                      `json:"chatEventLinkedChatChanged,omitempty"`
	ChatEventLocationChanged                        *ChatEventLocationChanged                        `json:"chatEventLocationChanged,omitempty"`
	ChatEventMemberInvited                          *ChatEventMemberInvited                          `json:"chatEventMemberInvited,omitempty"`
	ChatEventMemberJoined                           *ChatEventMemberJoined                           `json:"chatEventMemberJoined,omitempty"`
	ChatEventMemberJoinedByInviteLink               *ChatEventMemberJoinedByInviteLink               `json:"chatEventMemberJoinedByInviteLink,omitempty"`
	ChatEventMemberJoinedByRequest                  *ChatEventMemberJoinedByRequest                  `json:"chatEventMemberJoinedByRequest,omitempty"`
	ChatEventMemberLeft                             *ChatEventMemberLeft                             `json:"chatEventMemberLeft,omitempty"`
	ChatEventMemberPromoted                         *ChatEventMemberPromoted                         `json:"chatEventMemberPromoted,omitempty"`
	ChatEventMemberRestricted                       *ChatEventMemberRestricted                       `json:"chatEventMemberRestricted,omitempty"`
	ChatEventMemberSubscriptionExtended             *ChatEventMemberSubscriptionExtended             `json:"chatEventMemberSubscriptionExtended,omitempty"`
	ChatEventMessageAutoDeleteTimeChanged           *ChatEventMessageAutoDeleteTimeChanged           `json:"chatEventMessageAutoDeleteTimeChanged,omitempty"`
	ChatEventMessageDeleted                         *ChatEventMessageDeleted                         `json:"chatEventMessageDeleted,omitempty"`
	ChatEventMessageEdited                          *ChatEventMessageEdited                          `json:"chatEventMessageEdited,omitempty"`
	ChatEventMessagePinned                          *ChatEventMessagePinned                          `json:"chatEventMessagePinned,omitempty"`
	ChatEventMessageUnpinned                        *ChatEventMessageUnpinned                        `json:"chatEventMessageUnpinned,omitempty"`
	ChatEventPermissionsChanged                     *ChatEventPermissionsChanged                     `json:"chatEventPermissionsChanged,omitempty"`
	ChatEventPhotoChanged                           *ChatEventPhotoChanged                           `json:"chatEventPhotoChanged,omitempty"`
	ChatEventPollStopped                            *ChatEventPollStopped                            `json:"chatEventPollStopped,omitempty"`
	ChatEventProfileAccentColorChanged              *ChatEventProfileAccentColorChanged              `json:"chatEventProfileAccentColorChanged,omitempty"`
	ChatEventShowMessageSenderToggled               *ChatEventShowMessageSenderToggled               `json:"chatEventShowMessageSenderToggled,omitempty"`
	ChatEventSignMessagesToggled                    *ChatEventSignMessagesToggled                    `json:"chatEventSignMessagesToggled,omitempty"`
	ChatEventSlowModeDelayChanged                   *ChatEventSlowModeDelayChanged                   `json:"chatEventSlowModeDelayChanged,omitempty"`
	ChatEventStickerSetChanged                      *ChatEventStickerSetChanged                      `json:"chatEventStickerSetChanged,omitempty"`
	ChatEventTitleChanged                           *ChatEventTitleChanged                           `json:"chatEventTitleChanged,omitempty"`
	ChatEventUsernameChanged                        *ChatEventUsernameChanged                        `json:"chatEventUsernameChanged,omitempty"`
	ChatEventVideoChatCreated                       *ChatEventVideoChatCreated                       `json:"chatEventVideoChatCreated,omitempty"`
	ChatEventVideoChatEnded                         *ChatEventVideoChatEnded                         `json:"chatEventVideoChatEnded,omitempty"`
	ChatEventVideoChatMuteNewParticipantsToggled    *ChatEventVideoChatMuteNewParticipantsToggled    `json:"chatEventVideoChatMuteNewParticipantsToggled,omitempty"`
	ChatEventVideoChatParticipantIsMutedToggled     *ChatEventVideoChatParticipantIsMutedToggled     `json:"chatEventVideoChatParticipantIsMutedToggled,omitempty"`
	ChatEventVideoChatParticipantVolumeLevelChanged *ChatEventVideoChatParticipantVolumeLevelChanged `json:"chatEventVideoChatParticipantVolumeLevelChanged,omitempty"`
}

func (t *ChatEventAction) Type() string {
	return t.TypeStr
}

func (t *ChatEventAction) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatEventAction) GetExtra() string {
	return ""
}

func (t *ChatEventAction) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatEventAccentColorChanged":
		t.ChatEventAccentColorChanged = new(ChatEventAccentColorChanged)
		return json.Unmarshal(b, t.ChatEventAccentColorChanged)
	case "chatEventActiveUsernamesChanged":
		t.ChatEventActiveUsernamesChanged = new(ChatEventActiveUsernamesChanged)
		return json.Unmarshal(b, t.ChatEventActiveUsernamesChanged)
	case "chatEventAutomaticTranslationToggled":
		t.ChatEventAutomaticTranslationToggled = new(ChatEventAutomaticTranslationToggled)
		return json.Unmarshal(b, t.ChatEventAutomaticTranslationToggled)
	case "chatEventAvailableReactionsChanged":
		t.ChatEventAvailableReactionsChanged = new(ChatEventAvailableReactionsChanged)
		return json.Unmarshal(b, t.ChatEventAvailableReactionsChanged)
	case "chatEventBackgroundChanged":
		t.ChatEventBackgroundChanged = new(ChatEventBackgroundChanged)
		return json.Unmarshal(b, t.ChatEventBackgroundChanged)
	case "chatEventCustomEmojiStickerSetChanged":
		t.ChatEventCustomEmojiStickerSetChanged = new(ChatEventCustomEmojiStickerSetChanged)
		return json.Unmarshal(b, t.ChatEventCustomEmojiStickerSetChanged)
	case "chatEventDescriptionChanged":
		t.ChatEventDescriptionChanged = new(ChatEventDescriptionChanged)
		return json.Unmarshal(b, t.ChatEventDescriptionChanged)
	case "chatEventEmojiStatusChanged":
		t.ChatEventEmojiStatusChanged = new(ChatEventEmojiStatusChanged)
		return json.Unmarshal(b, t.ChatEventEmojiStatusChanged)
	case "chatEventForumTopicCreated":
		t.ChatEventForumTopicCreated = new(ChatEventForumTopicCreated)
		return json.Unmarshal(b, t.ChatEventForumTopicCreated)
	case "chatEventForumTopicDeleted":
		t.ChatEventForumTopicDeleted = new(ChatEventForumTopicDeleted)
		return json.Unmarshal(b, t.ChatEventForumTopicDeleted)
	case "chatEventForumTopicEdited":
		t.ChatEventForumTopicEdited = new(ChatEventForumTopicEdited)
		return json.Unmarshal(b, t.ChatEventForumTopicEdited)
	case "chatEventForumTopicPinned":
		t.ChatEventForumTopicPinned = new(ChatEventForumTopicPinned)
		return json.Unmarshal(b, t.ChatEventForumTopicPinned)
	case "chatEventForumTopicToggleIsClosed":
		t.ChatEventForumTopicToggleIsClosed = new(ChatEventForumTopicToggleIsClosed)
		return json.Unmarshal(b, t.ChatEventForumTopicToggleIsClosed)
	case "chatEventForumTopicToggleIsHidden":
		t.ChatEventForumTopicToggleIsHidden = new(ChatEventForumTopicToggleIsHidden)
		return json.Unmarshal(b, t.ChatEventForumTopicToggleIsHidden)
	case "chatEventHasAggressiveAntiSpamEnabledToggled":
		t.ChatEventHasAggressiveAntiSpamEnabledToggled = new(ChatEventHasAggressiveAntiSpamEnabledToggled)
		return json.Unmarshal(b, t.ChatEventHasAggressiveAntiSpamEnabledToggled)
	case "chatEventHasProtectedContentToggled":
		t.ChatEventHasProtectedContentToggled = new(ChatEventHasProtectedContentToggled)
		return json.Unmarshal(b, t.ChatEventHasProtectedContentToggled)
	case "chatEventInviteLinkDeleted":
		t.ChatEventInviteLinkDeleted = new(ChatEventInviteLinkDeleted)
		return json.Unmarshal(b, t.ChatEventInviteLinkDeleted)
	case "chatEventInviteLinkEdited":
		t.ChatEventInviteLinkEdited = new(ChatEventInviteLinkEdited)
		return json.Unmarshal(b, t.ChatEventInviteLinkEdited)
	case "chatEventInviteLinkRevoked":
		t.ChatEventInviteLinkRevoked = new(ChatEventInviteLinkRevoked)
		return json.Unmarshal(b, t.ChatEventInviteLinkRevoked)
	case "chatEventInvitesToggled":
		t.ChatEventInvitesToggled = new(ChatEventInvitesToggled)
		return json.Unmarshal(b, t.ChatEventInvitesToggled)
	case "chatEventIsAllHistoryAvailableToggled":
		t.ChatEventIsAllHistoryAvailableToggled = new(ChatEventIsAllHistoryAvailableToggled)
		return json.Unmarshal(b, t.ChatEventIsAllHistoryAvailableToggled)
	case "chatEventIsForumToggled":
		t.ChatEventIsForumToggled = new(ChatEventIsForumToggled)
		return json.Unmarshal(b, t.ChatEventIsForumToggled)
	case "chatEventLinkedChatChanged":
		t.ChatEventLinkedChatChanged = new(ChatEventLinkedChatChanged)
		return json.Unmarshal(b, t.ChatEventLinkedChatChanged)
	case "chatEventLocationChanged":
		t.ChatEventLocationChanged = new(ChatEventLocationChanged)
		return json.Unmarshal(b, t.ChatEventLocationChanged)
	case "chatEventMemberInvited":
		t.ChatEventMemberInvited = new(ChatEventMemberInvited)
		return json.Unmarshal(b, t.ChatEventMemberInvited)
	case "chatEventMemberJoined":
		t.ChatEventMemberJoined = new(ChatEventMemberJoined)
		return json.Unmarshal(b, t.ChatEventMemberJoined)
	case "chatEventMemberJoinedByInviteLink":
		t.ChatEventMemberJoinedByInviteLink = new(ChatEventMemberJoinedByInviteLink)
		return json.Unmarshal(b, t.ChatEventMemberJoinedByInviteLink)
	case "chatEventMemberJoinedByRequest":
		t.ChatEventMemberJoinedByRequest = new(ChatEventMemberJoinedByRequest)
		return json.Unmarshal(b, t.ChatEventMemberJoinedByRequest)
	case "chatEventMemberLeft":
		t.ChatEventMemberLeft = new(ChatEventMemberLeft)
		return json.Unmarshal(b, t.ChatEventMemberLeft)
	case "chatEventMemberPromoted":
		t.ChatEventMemberPromoted = new(ChatEventMemberPromoted)
		return json.Unmarshal(b, t.ChatEventMemberPromoted)
	case "chatEventMemberRestricted":
		t.ChatEventMemberRestricted = new(ChatEventMemberRestricted)
		return json.Unmarshal(b, t.ChatEventMemberRestricted)
	case "chatEventMemberSubscriptionExtended":
		t.ChatEventMemberSubscriptionExtended = new(ChatEventMemberSubscriptionExtended)
		return json.Unmarshal(b, t.ChatEventMemberSubscriptionExtended)
	case "chatEventMessageAutoDeleteTimeChanged":
		t.ChatEventMessageAutoDeleteTimeChanged = new(ChatEventMessageAutoDeleteTimeChanged)
		return json.Unmarshal(b, t.ChatEventMessageAutoDeleteTimeChanged)
	case "chatEventMessageDeleted":
		t.ChatEventMessageDeleted = new(ChatEventMessageDeleted)
		return json.Unmarshal(b, t.ChatEventMessageDeleted)
	case "chatEventMessageEdited":
		t.ChatEventMessageEdited = new(ChatEventMessageEdited)
		return json.Unmarshal(b, t.ChatEventMessageEdited)
	case "chatEventMessagePinned":
		t.ChatEventMessagePinned = new(ChatEventMessagePinned)
		return json.Unmarshal(b, t.ChatEventMessagePinned)
	case "chatEventMessageUnpinned":
		t.ChatEventMessageUnpinned = new(ChatEventMessageUnpinned)
		return json.Unmarshal(b, t.ChatEventMessageUnpinned)
	case "chatEventPermissionsChanged":
		t.ChatEventPermissionsChanged = new(ChatEventPermissionsChanged)
		return json.Unmarshal(b, t.ChatEventPermissionsChanged)
	case "chatEventPhotoChanged":
		t.ChatEventPhotoChanged = new(ChatEventPhotoChanged)
		return json.Unmarshal(b, t.ChatEventPhotoChanged)
	case "chatEventPollStopped":
		t.ChatEventPollStopped = new(ChatEventPollStopped)
		return json.Unmarshal(b, t.ChatEventPollStopped)
	case "chatEventProfileAccentColorChanged":
		t.ChatEventProfileAccentColorChanged = new(ChatEventProfileAccentColorChanged)
		return json.Unmarshal(b, t.ChatEventProfileAccentColorChanged)
	case "chatEventShowMessageSenderToggled":
		t.ChatEventShowMessageSenderToggled = new(ChatEventShowMessageSenderToggled)
		return json.Unmarshal(b, t.ChatEventShowMessageSenderToggled)
	case "chatEventSignMessagesToggled":
		t.ChatEventSignMessagesToggled = new(ChatEventSignMessagesToggled)
		return json.Unmarshal(b, t.ChatEventSignMessagesToggled)
	case "chatEventSlowModeDelayChanged":
		t.ChatEventSlowModeDelayChanged = new(ChatEventSlowModeDelayChanged)
		return json.Unmarshal(b, t.ChatEventSlowModeDelayChanged)
	case "chatEventStickerSetChanged":
		t.ChatEventStickerSetChanged = new(ChatEventStickerSetChanged)
		return json.Unmarshal(b, t.ChatEventStickerSetChanged)
	case "chatEventTitleChanged":
		t.ChatEventTitleChanged = new(ChatEventTitleChanged)
		return json.Unmarshal(b, t.ChatEventTitleChanged)
	case "chatEventUsernameChanged":
		t.ChatEventUsernameChanged = new(ChatEventUsernameChanged)
		return json.Unmarshal(b, t.ChatEventUsernameChanged)
	case "chatEventVideoChatCreated":
		t.ChatEventVideoChatCreated = new(ChatEventVideoChatCreated)
		return json.Unmarshal(b, t.ChatEventVideoChatCreated)
	case "chatEventVideoChatEnded":
		t.ChatEventVideoChatEnded = new(ChatEventVideoChatEnded)
		return json.Unmarshal(b, t.ChatEventVideoChatEnded)
	case "chatEventVideoChatMuteNewParticipantsToggled":
		t.ChatEventVideoChatMuteNewParticipantsToggled = new(ChatEventVideoChatMuteNewParticipantsToggled)
		return json.Unmarshal(b, t.ChatEventVideoChatMuteNewParticipantsToggled)
	case "chatEventVideoChatParticipantIsMutedToggled":
		t.ChatEventVideoChatParticipantIsMutedToggled = new(ChatEventVideoChatParticipantIsMutedToggled)
		return json.Unmarshal(b, t.ChatEventVideoChatParticipantIsMutedToggled)
	case "chatEventVideoChatParticipantVolumeLevelChanged":
		t.ChatEventVideoChatParticipantVolumeLevelChanged = new(ChatEventVideoChatParticipantVolumeLevelChanged)
		return json.Unmarshal(b, t.ChatEventVideoChatParticipantVolumeLevelChanged)
	}
	return nil
}

func (t *ChatEventAction) MarshalJSON() ([]byte, error) {
	if t.ChatEventAccentColorChanged != nil {
		return json.Marshal(t.ChatEventAccentColorChanged)
	}
	if t.ChatEventActiveUsernamesChanged != nil {
		return json.Marshal(t.ChatEventActiveUsernamesChanged)
	}
	if t.ChatEventAutomaticTranslationToggled != nil {
		return json.Marshal(t.ChatEventAutomaticTranslationToggled)
	}
	if t.ChatEventAvailableReactionsChanged != nil {
		return json.Marshal(t.ChatEventAvailableReactionsChanged)
	}
	if t.ChatEventBackgroundChanged != nil {
		return json.Marshal(t.ChatEventBackgroundChanged)
	}
	if t.ChatEventCustomEmojiStickerSetChanged != nil {
		return json.Marshal(t.ChatEventCustomEmojiStickerSetChanged)
	}
	if t.ChatEventDescriptionChanged != nil {
		return json.Marshal(t.ChatEventDescriptionChanged)
	}
	if t.ChatEventEmojiStatusChanged != nil {
		return json.Marshal(t.ChatEventEmojiStatusChanged)
	}
	if t.ChatEventForumTopicCreated != nil {
		return json.Marshal(t.ChatEventForumTopicCreated)
	}
	if t.ChatEventForumTopicDeleted != nil {
		return json.Marshal(t.ChatEventForumTopicDeleted)
	}
	if t.ChatEventForumTopicEdited != nil {
		return json.Marshal(t.ChatEventForumTopicEdited)
	}
	if t.ChatEventForumTopicPinned != nil {
		return json.Marshal(t.ChatEventForumTopicPinned)
	}
	if t.ChatEventForumTopicToggleIsClosed != nil {
		return json.Marshal(t.ChatEventForumTopicToggleIsClosed)
	}
	if t.ChatEventForumTopicToggleIsHidden != nil {
		return json.Marshal(t.ChatEventForumTopicToggleIsHidden)
	}
	if t.ChatEventHasAggressiveAntiSpamEnabledToggled != nil {
		return json.Marshal(t.ChatEventHasAggressiveAntiSpamEnabledToggled)
	}
	if t.ChatEventHasProtectedContentToggled != nil {
		return json.Marshal(t.ChatEventHasProtectedContentToggled)
	}
	if t.ChatEventInviteLinkDeleted != nil {
		return json.Marshal(t.ChatEventInviteLinkDeleted)
	}
	if t.ChatEventInviteLinkEdited != nil {
		return json.Marshal(t.ChatEventInviteLinkEdited)
	}
	if t.ChatEventInviteLinkRevoked != nil {
		return json.Marshal(t.ChatEventInviteLinkRevoked)
	}
	if t.ChatEventInvitesToggled != nil {
		return json.Marshal(t.ChatEventInvitesToggled)
	}
	if t.ChatEventIsAllHistoryAvailableToggled != nil {
		return json.Marshal(t.ChatEventIsAllHistoryAvailableToggled)
	}
	if t.ChatEventIsForumToggled != nil {
		return json.Marshal(t.ChatEventIsForumToggled)
	}
	if t.ChatEventLinkedChatChanged != nil {
		return json.Marshal(t.ChatEventLinkedChatChanged)
	}
	if t.ChatEventLocationChanged != nil {
		return json.Marshal(t.ChatEventLocationChanged)
	}
	if t.ChatEventMemberInvited != nil {
		return json.Marshal(t.ChatEventMemberInvited)
	}
	if t.ChatEventMemberJoined != nil {
		return json.Marshal(t.ChatEventMemberJoined)
	}
	if t.ChatEventMemberJoinedByInviteLink != nil {
		return json.Marshal(t.ChatEventMemberJoinedByInviteLink)
	}
	if t.ChatEventMemberJoinedByRequest != nil {
		return json.Marshal(t.ChatEventMemberJoinedByRequest)
	}
	if t.ChatEventMemberLeft != nil {
		return json.Marshal(t.ChatEventMemberLeft)
	}
	if t.ChatEventMemberPromoted != nil {
		return json.Marshal(t.ChatEventMemberPromoted)
	}
	if t.ChatEventMemberRestricted != nil {
		return json.Marshal(t.ChatEventMemberRestricted)
	}
	if t.ChatEventMemberSubscriptionExtended != nil {
		return json.Marshal(t.ChatEventMemberSubscriptionExtended)
	}
	if t.ChatEventMessageAutoDeleteTimeChanged != nil {
		return json.Marshal(t.ChatEventMessageAutoDeleteTimeChanged)
	}
	if t.ChatEventMessageDeleted != nil {
		return json.Marshal(t.ChatEventMessageDeleted)
	}
	if t.ChatEventMessageEdited != nil {
		return json.Marshal(t.ChatEventMessageEdited)
	}
	if t.ChatEventMessagePinned != nil {
		return json.Marshal(t.ChatEventMessagePinned)
	}
	if t.ChatEventMessageUnpinned != nil {
		return json.Marshal(t.ChatEventMessageUnpinned)
	}
	if t.ChatEventPermissionsChanged != nil {
		return json.Marshal(t.ChatEventPermissionsChanged)
	}
	if t.ChatEventPhotoChanged != nil {
		return json.Marshal(t.ChatEventPhotoChanged)
	}
	if t.ChatEventPollStopped != nil {
		return json.Marshal(t.ChatEventPollStopped)
	}
	if t.ChatEventProfileAccentColorChanged != nil {
		return json.Marshal(t.ChatEventProfileAccentColorChanged)
	}
	if t.ChatEventShowMessageSenderToggled != nil {
		return json.Marshal(t.ChatEventShowMessageSenderToggled)
	}
	if t.ChatEventSignMessagesToggled != nil {
		return json.Marshal(t.ChatEventSignMessagesToggled)
	}
	if t.ChatEventSlowModeDelayChanged != nil {
		return json.Marshal(t.ChatEventSlowModeDelayChanged)
	}
	if t.ChatEventStickerSetChanged != nil {
		return json.Marshal(t.ChatEventStickerSetChanged)
	}
	if t.ChatEventTitleChanged != nil {
		return json.Marshal(t.ChatEventTitleChanged)
	}
	if t.ChatEventUsernameChanged != nil {
		return json.Marshal(t.ChatEventUsernameChanged)
	}
	if t.ChatEventVideoChatCreated != nil {
		return json.Marshal(t.ChatEventVideoChatCreated)
	}
	if t.ChatEventVideoChatEnded != nil {
		return json.Marshal(t.ChatEventVideoChatEnded)
	}
	if t.ChatEventVideoChatMuteNewParticipantsToggled != nil {
		return json.Marshal(t.ChatEventVideoChatMuteNewParticipantsToggled)
	}
	if t.ChatEventVideoChatParticipantIsMutedToggled != nil {
		return json.Marshal(t.ChatEventVideoChatParticipantIsMutedToggled)
	}
	if t.ChatEventVideoChatParticipantVolumeLevelChanged != nil {
		return json.Marshal(t.ChatEventVideoChatParticipantVolumeLevelChanged)
	}
	return nil, fmt.Errorf("no value set for ChatEventAction")
}

// ChatList Describes a list of chats
type ChatList struct {
	TypeStr         string           `json:"@type"`
	ChatListArchive *ChatListArchive `json:"chatListArchive,omitempty"`
	ChatListFolder  *ChatListFolder  `json:"chatListFolder,omitempty"`
	ChatListMain    *ChatListMain    `json:"chatListMain,omitempty"`
}

func (t *ChatList) Type() string {
	return t.TypeStr
}

func (t *ChatList) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatList) GetExtra() string {
	return ""
}

func (t *ChatList) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatListArchive":
		t.ChatListArchive = new(ChatListArchive)
		return json.Unmarshal(b, t.ChatListArchive)
	case "chatListFolder":
		t.ChatListFolder = new(ChatListFolder)
		return json.Unmarshal(b, t.ChatListFolder)
	case "chatListMain":
		t.ChatListMain = new(ChatListMain)
		return json.Unmarshal(b, t.ChatListMain)
	}
	return nil
}

func (t *ChatList) MarshalJSON() ([]byte, error) {
	if t.ChatListArchive != nil {
		return json.Marshal(t.ChatListArchive)
	}
	if t.ChatListFolder != nil {
		return json.Marshal(t.ChatListFolder)
	}
	if t.ChatListMain != nil {
		return json.Marshal(t.ChatListMain)
	}
	return nil, fmt.Errorf("no value set for ChatList")
}

// ChatMembersFilter Specifies the kind of chat members to return in searchChatMembers
type ChatMembersFilter struct {
	TypeStr                         string                           `json:"@type"`
	ChatMembersFilterAdministrators *ChatMembersFilterAdministrators `json:"chatMembersFilterAdministrators,omitempty"`
	ChatMembersFilterBanned         *ChatMembersFilterBanned         `json:"chatMembersFilterBanned,omitempty"`
	ChatMembersFilterBots           *ChatMembersFilterBots           `json:"chatMembersFilterBots,omitempty"`
	ChatMembersFilterContacts       *ChatMembersFilterContacts       `json:"chatMembersFilterContacts,omitempty"`
	ChatMembersFilterMembers        *ChatMembersFilterMembers        `json:"chatMembersFilterMembers,omitempty"`
	ChatMembersFilterMention        *ChatMembersFilterMention        `json:"chatMembersFilterMention,omitempty"`
	ChatMembersFilterRestricted     *ChatMembersFilterRestricted     `json:"chatMembersFilterRestricted,omitempty"`
}

func (t *ChatMembersFilter) Type() string {
	return t.TypeStr
}

func (t *ChatMembersFilter) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatMembersFilter) GetExtra() string {
	return ""
}

func (t *ChatMembersFilter) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatMembersFilterAdministrators":
		t.ChatMembersFilterAdministrators = new(ChatMembersFilterAdministrators)
		return json.Unmarshal(b, t.ChatMembersFilterAdministrators)
	case "chatMembersFilterBanned":
		t.ChatMembersFilterBanned = new(ChatMembersFilterBanned)
		return json.Unmarshal(b, t.ChatMembersFilterBanned)
	case "chatMembersFilterBots":
		t.ChatMembersFilterBots = new(ChatMembersFilterBots)
		return json.Unmarshal(b, t.ChatMembersFilterBots)
	case "chatMembersFilterContacts":
		t.ChatMembersFilterContacts = new(ChatMembersFilterContacts)
		return json.Unmarshal(b, t.ChatMembersFilterContacts)
	case "chatMembersFilterMembers":
		t.ChatMembersFilterMembers = new(ChatMembersFilterMembers)
		return json.Unmarshal(b, t.ChatMembersFilterMembers)
	case "chatMembersFilterMention":
		t.ChatMembersFilterMention = new(ChatMembersFilterMention)
		return json.Unmarshal(b, t.ChatMembersFilterMention)
	case "chatMembersFilterRestricted":
		t.ChatMembersFilterRestricted = new(ChatMembersFilterRestricted)
		return json.Unmarshal(b, t.ChatMembersFilterRestricted)
	}
	return nil
}

func (t *ChatMembersFilter) MarshalJSON() ([]byte, error) {
	if t.ChatMembersFilterAdministrators != nil {
		return json.Marshal(t.ChatMembersFilterAdministrators)
	}
	if t.ChatMembersFilterBanned != nil {
		return json.Marshal(t.ChatMembersFilterBanned)
	}
	if t.ChatMembersFilterBots != nil {
		return json.Marshal(t.ChatMembersFilterBots)
	}
	if t.ChatMembersFilterContacts != nil {
		return json.Marshal(t.ChatMembersFilterContacts)
	}
	if t.ChatMembersFilterMembers != nil {
		return json.Marshal(t.ChatMembersFilterMembers)
	}
	if t.ChatMembersFilterMention != nil {
		return json.Marshal(t.ChatMembersFilterMention)
	}
	if t.ChatMembersFilterRestricted != nil {
		return json.Marshal(t.ChatMembersFilterRestricted)
	}
	return nil, fmt.Errorf("no value set for ChatMembersFilter")
}

// ChatMemberStatus Provides information about the status of a member in a chat
type ChatMemberStatus struct {
	TypeStr                       string                         `json:"@type"`
	ChatMemberStatusAdministrator *ChatMemberStatusAdministrator `json:"chatMemberStatusAdministrator,omitempty"`
	ChatMemberStatusBanned        *ChatMemberStatusBanned        `json:"chatMemberStatusBanned,omitempty"`
	ChatMemberStatusCreator       *ChatMemberStatusCreator       `json:"chatMemberStatusCreator,omitempty"`
	ChatMemberStatusLeft          *ChatMemberStatusLeft          `json:"chatMemberStatusLeft,omitempty"`
	ChatMemberStatusMember        *ChatMemberStatusMember        `json:"chatMemberStatusMember,omitempty"`
	ChatMemberStatusRestricted    *ChatMemberStatusRestricted    `json:"chatMemberStatusRestricted,omitempty"`
}

func (t *ChatMemberStatus) Type() string {
	return t.TypeStr
}

func (t *ChatMemberStatus) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatMemberStatus) GetExtra() string {
	return ""
}

func (t *ChatMemberStatus) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatMemberStatusAdministrator":
		t.ChatMemberStatusAdministrator = new(ChatMemberStatusAdministrator)
		return json.Unmarshal(b, t.ChatMemberStatusAdministrator)
	case "chatMemberStatusBanned":
		t.ChatMemberStatusBanned = new(ChatMemberStatusBanned)
		return json.Unmarshal(b, t.ChatMemberStatusBanned)
	case "chatMemberStatusCreator":
		t.ChatMemberStatusCreator = new(ChatMemberStatusCreator)
		return json.Unmarshal(b, t.ChatMemberStatusCreator)
	case "chatMemberStatusLeft":
		t.ChatMemberStatusLeft = new(ChatMemberStatusLeft)
		return json.Unmarshal(b, t.ChatMemberStatusLeft)
	case "chatMemberStatusMember":
		t.ChatMemberStatusMember = new(ChatMemberStatusMember)
		return json.Unmarshal(b, t.ChatMemberStatusMember)
	case "chatMemberStatusRestricted":
		t.ChatMemberStatusRestricted = new(ChatMemberStatusRestricted)
		return json.Unmarshal(b, t.ChatMemberStatusRestricted)
	}
	return nil
}

func (t *ChatMemberStatus) MarshalJSON() ([]byte, error) {
	if t.ChatMemberStatusAdministrator != nil {
		return json.Marshal(t.ChatMemberStatusAdministrator)
	}
	if t.ChatMemberStatusBanned != nil {
		return json.Marshal(t.ChatMemberStatusBanned)
	}
	if t.ChatMemberStatusCreator != nil {
		return json.Marshal(t.ChatMemberStatusCreator)
	}
	if t.ChatMemberStatusLeft != nil {
		return json.Marshal(t.ChatMemberStatusLeft)
	}
	if t.ChatMemberStatusMember != nil {
		return json.Marshal(t.ChatMemberStatusMember)
	}
	if t.ChatMemberStatusRestricted != nil {
		return json.Marshal(t.ChatMemberStatusRestricted)
	}
	return nil, fmt.Errorf("no value set for ChatMemberStatus")
}

// ChatPhotoStickerType Describes type of sticker, which was used to create a chat photo
type ChatPhotoStickerType struct {
	TypeStr                           string                             `json:"@type"`
	ChatPhotoStickerTypeCustomEmoji   *ChatPhotoStickerTypeCustomEmoji   `json:"chatPhotoStickerTypeCustomEmoji,omitempty"`
	ChatPhotoStickerTypeRegularOrMask *ChatPhotoStickerTypeRegularOrMask `json:"chatPhotoStickerTypeRegularOrMask,omitempty"`
}

func (t *ChatPhotoStickerType) Type() string {
	return t.TypeStr
}

func (t *ChatPhotoStickerType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatPhotoStickerType) GetExtra() string {
	return ""
}

func (t *ChatPhotoStickerType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatPhotoStickerTypeCustomEmoji":
		t.ChatPhotoStickerTypeCustomEmoji = new(ChatPhotoStickerTypeCustomEmoji)
		return json.Unmarshal(b, t.ChatPhotoStickerTypeCustomEmoji)
	case "chatPhotoStickerTypeRegularOrMask":
		t.ChatPhotoStickerTypeRegularOrMask = new(ChatPhotoStickerTypeRegularOrMask)
		return json.Unmarshal(b, t.ChatPhotoStickerTypeRegularOrMask)
	}
	return nil
}

func (t *ChatPhotoStickerType) MarshalJSON() ([]byte, error) {
	if t.ChatPhotoStickerTypeCustomEmoji != nil {
		return json.Marshal(t.ChatPhotoStickerTypeCustomEmoji)
	}
	if t.ChatPhotoStickerTypeRegularOrMask != nil {
		return json.Marshal(t.ChatPhotoStickerTypeRegularOrMask)
	}
	return nil, fmt.Errorf("no value set for ChatPhotoStickerType")
}

// ChatRevenueTransactionType Describes type of transaction for revenue earned from sponsored messages in a chat
type ChatRevenueTransactionType struct {
	TypeStr                                            string                                              `json:"@type"`
	ChatRevenueTransactionTypeFragmentRefund           *ChatRevenueTransactionTypeFragmentRefund           `json:"chatRevenueTransactionTypeFragmentRefund,omitempty"`
	ChatRevenueTransactionTypeFragmentWithdrawal       *ChatRevenueTransactionTypeFragmentWithdrawal       `json:"chatRevenueTransactionTypeFragmentWithdrawal,omitempty"`
	ChatRevenueTransactionTypeSponsoredMessageEarnings *ChatRevenueTransactionTypeSponsoredMessageEarnings `json:"chatRevenueTransactionTypeSponsoredMessageEarnings,omitempty"`
	ChatRevenueTransactionTypeSuggestedPostEarnings    *ChatRevenueTransactionTypeSuggestedPostEarnings    `json:"chatRevenueTransactionTypeSuggestedPostEarnings,omitempty"`
	ChatRevenueTransactionTypeUnsupported              *ChatRevenueTransactionTypeUnsupported              `json:"chatRevenueTransactionTypeUnsupported,omitempty"`
}

func (t *ChatRevenueTransactionType) Type() string {
	return t.TypeStr
}

func (t *ChatRevenueTransactionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatRevenueTransactionType) GetExtra() string {
	return ""
}

func (t *ChatRevenueTransactionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatRevenueTransactionTypeFragmentRefund":
		t.ChatRevenueTransactionTypeFragmentRefund = new(ChatRevenueTransactionTypeFragmentRefund)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeFragmentRefund)
	case "chatRevenueTransactionTypeFragmentWithdrawal":
		t.ChatRevenueTransactionTypeFragmentWithdrawal = new(ChatRevenueTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeFragmentWithdrawal)
	case "chatRevenueTransactionTypeSponsoredMessageEarnings":
		t.ChatRevenueTransactionTypeSponsoredMessageEarnings = new(ChatRevenueTransactionTypeSponsoredMessageEarnings)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeSponsoredMessageEarnings)
	case "chatRevenueTransactionTypeSuggestedPostEarnings":
		t.ChatRevenueTransactionTypeSuggestedPostEarnings = new(ChatRevenueTransactionTypeSuggestedPostEarnings)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeSuggestedPostEarnings)
	case "chatRevenueTransactionTypeUnsupported":
		t.ChatRevenueTransactionTypeUnsupported = new(ChatRevenueTransactionTypeUnsupported)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeUnsupported)
	}
	return nil
}

func (t *ChatRevenueTransactionType) MarshalJSON() ([]byte, error) {
	if t.ChatRevenueTransactionTypeFragmentRefund != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeFragmentRefund)
	}
	if t.ChatRevenueTransactionTypeFragmentWithdrawal != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeFragmentWithdrawal)
	}
	if t.ChatRevenueTransactionTypeSponsoredMessageEarnings != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeSponsoredMessageEarnings)
	}
	if t.ChatRevenueTransactionTypeSuggestedPostEarnings != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeSuggestedPostEarnings)
	}
	if t.ChatRevenueTransactionTypeUnsupported != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeUnsupported)
	}
	return nil, fmt.Errorf("no value set for ChatRevenueTransactionType")
}

// ChatSource Describes a reason why an external chat is shown in a chat list
type ChatSource struct {
	TypeStr                             string                               `json:"@type"`
	ChatSourceMtprotoProxy              *ChatSourceMtprotoProxy              `json:"chatSourceMtprotoProxy,omitempty"`
	ChatSourcePublicServiceAnnouncement *ChatSourcePublicServiceAnnouncement `json:"chatSourcePublicServiceAnnouncement,omitempty"`
}

func (t *ChatSource) Type() string {
	return t.TypeStr
}

func (t *ChatSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatSource) GetExtra() string {
	return ""
}

func (t *ChatSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatSourceMtprotoProxy":
		t.ChatSourceMtprotoProxy = new(ChatSourceMtprotoProxy)
		return json.Unmarshal(b, t.ChatSourceMtprotoProxy)
	case "chatSourcePublicServiceAnnouncement":
		t.ChatSourcePublicServiceAnnouncement = new(ChatSourcePublicServiceAnnouncement)
		return json.Unmarshal(b, t.ChatSourcePublicServiceAnnouncement)
	}
	return nil
}

func (t *ChatSource) MarshalJSON() ([]byte, error) {
	if t.ChatSourceMtprotoProxy != nil {
		return json.Marshal(t.ChatSourceMtprotoProxy)
	}
	if t.ChatSourcePublicServiceAnnouncement != nil {
		return json.Marshal(t.ChatSourcePublicServiceAnnouncement)
	}
	return nil, fmt.Errorf("no value set for ChatSource")
}

// ChatStatistics Contains a detailed statistics about a chat
type ChatStatistics struct {
	TypeStr                  string                    `json:"@type"`
	ChatStatisticsChannel    *ChatStatisticsChannel    `json:"chatStatisticsChannel,omitempty"`
	ChatStatisticsSupergroup *ChatStatisticsSupergroup `json:"chatStatisticsSupergroup,omitempty"`
}

func (t *ChatStatistics) Type() string {
	return t.TypeStr
}

func (t *ChatStatistics) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatStatistics) GetExtra() string {
	return ""
}

func (t *ChatStatistics) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatStatisticsChannel":
		t.ChatStatisticsChannel = new(ChatStatisticsChannel)
		return json.Unmarshal(b, t.ChatStatisticsChannel)
	case "chatStatisticsSupergroup":
		t.ChatStatisticsSupergroup = new(ChatStatisticsSupergroup)
		return json.Unmarshal(b, t.ChatStatisticsSupergroup)
	}
	return nil
}

func (t *ChatStatistics) MarshalJSON() ([]byte, error) {
	if t.ChatStatisticsChannel != nil {
		return json.Marshal(t.ChatStatisticsChannel)
	}
	if t.ChatStatisticsSupergroup != nil {
		return json.Marshal(t.ChatStatisticsSupergroup)
	}
	return nil, fmt.Errorf("no value set for ChatStatistics")
}

// ChatStatisticsObjectType Describes type of object, for which statistics are provided
type ChatStatisticsObjectType struct {
	TypeStr                         string                           `json:"@type"`
	ChatStatisticsObjectTypeMessage *ChatStatisticsObjectTypeMessage `json:"chatStatisticsObjectTypeMessage,omitempty"`
	ChatStatisticsObjectTypeStory   *ChatStatisticsObjectTypeStory   `json:"chatStatisticsObjectTypeStory,omitempty"`
}

func (t *ChatStatisticsObjectType) Type() string {
	return t.TypeStr
}

func (t *ChatStatisticsObjectType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatStatisticsObjectType) GetExtra() string {
	return ""
}

func (t *ChatStatisticsObjectType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatStatisticsObjectTypeMessage":
		t.ChatStatisticsObjectTypeMessage = new(ChatStatisticsObjectTypeMessage)
		return json.Unmarshal(b, t.ChatStatisticsObjectTypeMessage)
	case "chatStatisticsObjectTypeStory":
		t.ChatStatisticsObjectTypeStory = new(ChatStatisticsObjectTypeStory)
		return json.Unmarshal(b, t.ChatStatisticsObjectTypeStory)
	}
	return nil
}

func (t *ChatStatisticsObjectType) MarshalJSON() ([]byte, error) {
	if t.ChatStatisticsObjectTypeMessage != nil {
		return json.Marshal(t.ChatStatisticsObjectTypeMessage)
	}
	if t.ChatStatisticsObjectTypeStory != nil {
		return json.Marshal(t.ChatStatisticsObjectTypeStory)
	}
	return nil, fmt.Errorf("no value set for ChatStatisticsObjectType")
}

// ChatTheme Describes a chat theme
type ChatTheme struct {
	TypeStr        string          `json:"@type"`
	ChatThemeEmoji *ChatThemeEmoji `json:"chatThemeEmoji,omitempty"`
	ChatThemeGift  *ChatThemeGift  `json:"chatThemeGift,omitempty"`
}

func (t *ChatTheme) Type() string {
	return t.TypeStr
}

func (t *ChatTheme) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatTheme) GetExtra() string {
	return ""
}

func (t *ChatTheme) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatThemeEmoji":
		t.ChatThemeEmoji = new(ChatThemeEmoji)
		return json.Unmarshal(b, t.ChatThemeEmoji)
	case "chatThemeGift":
		t.ChatThemeGift = new(ChatThemeGift)
		return json.Unmarshal(b, t.ChatThemeGift)
	}
	return nil
}

func (t *ChatTheme) MarshalJSON() ([]byte, error) {
	if t.ChatThemeEmoji != nil {
		return json.Marshal(t.ChatThemeEmoji)
	}
	if t.ChatThemeGift != nil {
		return json.Marshal(t.ChatThemeGift)
	}
	return nil, fmt.Errorf("no value set for ChatTheme")
}

// ChatType Describes the type of chat
type ChatType struct {
	TypeStr            string              `json:"@type"`
	ChatTypeBasicGroup *ChatTypeBasicGroup `json:"chatTypeBasicGroup,omitempty"`
	ChatTypePrivate    *ChatTypePrivate    `json:"chatTypePrivate,omitempty"`
	ChatTypeSecret     *ChatTypeSecret     `json:"chatTypeSecret,omitempty"`
	ChatTypeSupergroup *ChatTypeSupergroup `json:"chatTypeSupergroup,omitempty"`
}

func (t *ChatType) Type() string {
	return t.TypeStr
}

func (t *ChatType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ChatType) GetExtra() string {
	return ""
}

func (t *ChatType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "chatTypeBasicGroup":
		t.ChatTypeBasicGroup = new(ChatTypeBasicGroup)
		return json.Unmarshal(b, t.ChatTypeBasicGroup)
	case "chatTypePrivate":
		t.ChatTypePrivate = new(ChatTypePrivate)
		return json.Unmarshal(b, t.ChatTypePrivate)
	case "chatTypeSecret":
		t.ChatTypeSecret = new(ChatTypeSecret)
		return json.Unmarshal(b, t.ChatTypeSecret)
	case "chatTypeSupergroup":
		t.ChatTypeSupergroup = new(ChatTypeSupergroup)
		return json.Unmarshal(b, t.ChatTypeSupergroup)
	}
	return nil
}

func (t *ChatType) MarshalJSON() ([]byte, error) {
	if t.ChatTypeBasicGroup != nil {
		return json.Marshal(t.ChatTypeBasicGroup)
	}
	if t.ChatTypePrivate != nil {
		return json.Marshal(t.ChatTypePrivate)
	}
	if t.ChatTypeSecret != nil {
		return json.Marshal(t.ChatTypeSecret)
	}
	if t.ChatTypeSupergroup != nil {
		return json.Marshal(t.ChatTypeSupergroup)
	}
	return nil, fmt.Errorf("no value set for ChatType")
}

// CheckChatUsernameResult Represents result of checking whether a username can be set for a chat
type CheckChatUsernameResult struct {
	TypeStr                                        string                                          `json:"@type"`
	CheckChatUsernameResultOk                      *CheckChatUsernameResultOk                      `json:"checkChatUsernameResultOk,omitempty"`
	CheckChatUsernameResultPublicChatsTooMany      *CheckChatUsernameResultPublicChatsTooMany      `json:"checkChatUsernameResultPublicChatsTooMany,omitempty"`
	CheckChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable `json:"checkChatUsernameResultPublicGroupsUnavailable,omitempty"`
	CheckChatUsernameResultUsernameInvalid         *CheckChatUsernameResultUsernameInvalid         `json:"checkChatUsernameResultUsernameInvalid,omitempty"`
	CheckChatUsernameResultUsernameOccupied        *CheckChatUsernameResultUsernameOccupied        `json:"checkChatUsernameResultUsernameOccupied,omitempty"`
	CheckChatUsernameResultUsernamePurchasable     *CheckChatUsernameResultUsernamePurchasable     `json:"checkChatUsernameResultUsernamePurchasable,omitempty"`
}

func (t *CheckChatUsernameResult) Type() string {
	return t.TypeStr
}

func (t *CheckChatUsernameResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CheckChatUsernameResult) GetExtra() string {
	return ""
}

func (t *CheckChatUsernameResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "checkChatUsernameResultOk":
		t.CheckChatUsernameResultOk = new(CheckChatUsernameResultOk)
		return json.Unmarshal(b, t.CheckChatUsernameResultOk)
	case "checkChatUsernameResultPublicChatsTooMany":
		t.CheckChatUsernameResultPublicChatsTooMany = new(CheckChatUsernameResultPublicChatsTooMany)
		return json.Unmarshal(b, t.CheckChatUsernameResultPublicChatsTooMany)
	case "checkChatUsernameResultPublicGroupsUnavailable":
		t.CheckChatUsernameResultPublicGroupsUnavailable = new(CheckChatUsernameResultPublicGroupsUnavailable)
		return json.Unmarshal(b, t.CheckChatUsernameResultPublicGroupsUnavailable)
	case "checkChatUsernameResultUsernameInvalid":
		t.CheckChatUsernameResultUsernameInvalid = new(CheckChatUsernameResultUsernameInvalid)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernameInvalid)
	case "checkChatUsernameResultUsernameOccupied":
		t.CheckChatUsernameResultUsernameOccupied = new(CheckChatUsernameResultUsernameOccupied)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernameOccupied)
	case "checkChatUsernameResultUsernamePurchasable":
		t.CheckChatUsernameResultUsernamePurchasable = new(CheckChatUsernameResultUsernamePurchasable)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernamePurchasable)
	}
	return nil
}

func (t *CheckChatUsernameResult) MarshalJSON() ([]byte, error) {
	if t.CheckChatUsernameResultOk != nil {
		return json.Marshal(t.CheckChatUsernameResultOk)
	}
	if t.CheckChatUsernameResultPublicChatsTooMany != nil {
		return json.Marshal(t.CheckChatUsernameResultPublicChatsTooMany)
	}
	if t.CheckChatUsernameResultPublicGroupsUnavailable != nil {
		return json.Marshal(t.CheckChatUsernameResultPublicGroupsUnavailable)
	}
	if t.CheckChatUsernameResultUsernameInvalid != nil {
		return json.Marshal(t.CheckChatUsernameResultUsernameInvalid)
	}
	if t.CheckChatUsernameResultUsernameOccupied != nil {
		return json.Marshal(t.CheckChatUsernameResultUsernameOccupied)
	}
	if t.CheckChatUsernameResultUsernamePurchasable != nil {
		return json.Marshal(t.CheckChatUsernameResultUsernamePurchasable)
	}
	return nil, fmt.Errorf("no value set for CheckChatUsernameResult")
}

// CheckStickerSetNameResult Represents result of checking whether a name can be used for a new sticker set
type CheckStickerSetNameResult struct {
	TypeStr                               string                                 `json:"@type"`
	CheckStickerSetNameResultNameInvalid  *CheckStickerSetNameResultNameInvalid  `json:"checkStickerSetNameResultNameInvalid,omitempty"`
	CheckStickerSetNameResultNameOccupied *CheckStickerSetNameResultNameOccupied `json:"checkStickerSetNameResultNameOccupied,omitempty"`
	CheckStickerSetNameResultOk           *CheckStickerSetNameResultOk           `json:"checkStickerSetNameResultOk,omitempty"`
}

func (t *CheckStickerSetNameResult) Type() string {
	return t.TypeStr
}

func (t *CheckStickerSetNameResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CheckStickerSetNameResult) GetExtra() string {
	return ""
}

func (t *CheckStickerSetNameResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "checkStickerSetNameResultNameInvalid":
		t.CheckStickerSetNameResultNameInvalid = new(CheckStickerSetNameResultNameInvalid)
		return json.Unmarshal(b, t.CheckStickerSetNameResultNameInvalid)
	case "checkStickerSetNameResultNameOccupied":
		t.CheckStickerSetNameResultNameOccupied = new(CheckStickerSetNameResultNameOccupied)
		return json.Unmarshal(b, t.CheckStickerSetNameResultNameOccupied)
	case "checkStickerSetNameResultOk":
		t.CheckStickerSetNameResultOk = new(CheckStickerSetNameResultOk)
		return json.Unmarshal(b, t.CheckStickerSetNameResultOk)
	}
	return nil
}

func (t *CheckStickerSetNameResult) MarshalJSON() ([]byte, error) {
	if t.CheckStickerSetNameResultNameInvalid != nil {
		return json.Marshal(t.CheckStickerSetNameResultNameInvalid)
	}
	if t.CheckStickerSetNameResultNameOccupied != nil {
		return json.Marshal(t.CheckStickerSetNameResultNameOccupied)
	}
	if t.CheckStickerSetNameResultOk != nil {
		return json.Marshal(t.CheckStickerSetNameResultOk)
	}
	return nil, fmt.Errorf("no value set for CheckStickerSetNameResult")
}

// CollectibleItemType Describes a collectible item that can be purchased at https://fragment.com
type CollectibleItemType struct {
	TypeStr                        string                          `json:"@type"`
	CollectibleItemTypePhoneNumber *CollectibleItemTypePhoneNumber `json:"collectibleItemTypePhoneNumber,omitempty"`
	CollectibleItemTypeUsername    *CollectibleItemTypeUsername    `json:"collectibleItemTypeUsername,omitempty"`
}

func (t *CollectibleItemType) Type() string {
	return t.TypeStr
}

func (t *CollectibleItemType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *CollectibleItemType) GetExtra() string {
	return ""
}

func (t *CollectibleItemType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "collectibleItemTypePhoneNumber":
		t.CollectibleItemTypePhoneNumber = new(CollectibleItemTypePhoneNumber)
		return json.Unmarshal(b, t.CollectibleItemTypePhoneNumber)
	case "collectibleItemTypeUsername":
		t.CollectibleItemTypeUsername = new(CollectibleItemTypeUsername)
		return json.Unmarshal(b, t.CollectibleItemTypeUsername)
	}
	return nil
}

func (t *CollectibleItemType) MarshalJSON() ([]byte, error) {
	if t.CollectibleItemTypePhoneNumber != nil {
		return json.Marshal(t.CollectibleItemTypePhoneNumber)
	}
	if t.CollectibleItemTypeUsername != nil {
		return json.Marshal(t.CollectibleItemTypeUsername)
	}
	return nil, fmt.Errorf("no value set for CollectibleItemType")
}

// ConnectionState Describes the current state of the connection to Telegram servers
type ConnectionState struct {
	TypeStr                          string                            `json:"@type"`
	ConnectionStateConnecting        *ConnectionStateConnecting        `json:"connectionStateConnecting,omitempty"`
	ConnectionStateConnectingToProxy *ConnectionStateConnectingToProxy `json:"connectionStateConnectingToProxy,omitempty"`
	ConnectionStateReady             *ConnectionStateReady             `json:"connectionStateReady,omitempty"`
	ConnectionStateUpdating          *ConnectionStateUpdating          `json:"connectionStateUpdating,omitempty"`
	ConnectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork `json:"connectionStateWaitingForNetwork,omitempty"`
}

func (t *ConnectionState) Type() string {
	return t.TypeStr
}

func (t *ConnectionState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ConnectionState) GetExtra() string {
	return ""
}

func (t *ConnectionState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "connectionStateConnecting":
		t.ConnectionStateConnecting = new(ConnectionStateConnecting)
		return json.Unmarshal(b, t.ConnectionStateConnecting)
	case "connectionStateConnectingToProxy":
		t.ConnectionStateConnectingToProxy = new(ConnectionStateConnectingToProxy)
		return json.Unmarshal(b, t.ConnectionStateConnectingToProxy)
	case "connectionStateReady":
		t.ConnectionStateReady = new(ConnectionStateReady)
		return json.Unmarshal(b, t.ConnectionStateReady)
	case "connectionStateUpdating":
		t.ConnectionStateUpdating = new(ConnectionStateUpdating)
		return json.Unmarshal(b, t.ConnectionStateUpdating)
	case "connectionStateWaitingForNetwork":
		t.ConnectionStateWaitingForNetwork = new(ConnectionStateWaitingForNetwork)
		return json.Unmarshal(b, t.ConnectionStateWaitingForNetwork)
	}
	return nil
}

func (t *ConnectionState) MarshalJSON() ([]byte, error) {
	if t.ConnectionStateConnecting != nil {
		return json.Marshal(t.ConnectionStateConnecting)
	}
	if t.ConnectionStateConnectingToProxy != nil {
		return json.Marshal(t.ConnectionStateConnectingToProxy)
	}
	if t.ConnectionStateReady != nil {
		return json.Marshal(t.ConnectionStateReady)
	}
	if t.ConnectionStateUpdating != nil {
		return json.Marshal(t.ConnectionStateUpdating)
	}
	if t.ConnectionStateWaitingForNetwork != nil {
		return json.Marshal(t.ConnectionStateWaitingForNetwork)
	}
	return nil, fmt.Errorf("no value set for ConnectionState")
}

// DeviceToken Represents a data needed to subscribe for push notifications through registerDevice method.
type DeviceToken struct {
	TypeStr                           string                             `json:"@type"`
	DeviceTokenApplePush              *DeviceTokenApplePush              `json:"deviceTokenApplePush,omitempty"`
	DeviceTokenApplePushVoIP          *DeviceTokenApplePushVoIP          `json:"deviceTokenApplePushVoIP,omitempty"`
	DeviceTokenBlackBerryPush         *DeviceTokenBlackBerryPush         `json:"deviceTokenBlackBerryPush,omitempty"`
	DeviceTokenFirebaseCloudMessaging *DeviceTokenFirebaseCloudMessaging `json:"deviceTokenFirebaseCloudMessaging,omitempty"`
	DeviceTokenHuaweiPush             *DeviceTokenHuaweiPush             `json:"deviceTokenHuaweiPush,omitempty"`
	DeviceTokenMicrosoftPush          *DeviceTokenMicrosoftPush          `json:"deviceTokenMicrosoftPush,omitempty"`
	DeviceTokenMicrosoftPushVoIP      *DeviceTokenMicrosoftPushVoIP      `json:"deviceTokenMicrosoftPushVoIP,omitempty"`
	DeviceTokenSimplePush             *DeviceTokenSimplePush             `json:"deviceTokenSimplePush,omitempty"`
	DeviceTokenTizenPush              *DeviceTokenTizenPush              `json:"deviceTokenTizenPush,omitempty"`
	DeviceTokenUbuntuPush             *DeviceTokenUbuntuPush             `json:"deviceTokenUbuntuPush,omitempty"`
	DeviceTokenWebPush                *DeviceTokenWebPush                `json:"deviceTokenWebPush,omitempty"`
	DeviceTokenWindowsPush            *DeviceTokenWindowsPush            `json:"deviceTokenWindowsPush,omitempty"`
}

func (t *DeviceToken) Type() string {
	return t.TypeStr
}

func (t *DeviceToken) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *DeviceToken) GetExtra() string {
	return ""
}

func (t *DeviceToken) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "deviceTokenApplePush":
		t.DeviceTokenApplePush = new(DeviceTokenApplePush)
		return json.Unmarshal(b, t.DeviceTokenApplePush)
	case "deviceTokenApplePushVoIP":
		t.DeviceTokenApplePushVoIP = new(DeviceTokenApplePushVoIP)
		return json.Unmarshal(b, t.DeviceTokenApplePushVoIP)
	case "deviceTokenBlackBerryPush":
		t.DeviceTokenBlackBerryPush = new(DeviceTokenBlackBerryPush)
		return json.Unmarshal(b, t.DeviceTokenBlackBerryPush)
	case "deviceTokenFirebaseCloudMessaging":
		t.DeviceTokenFirebaseCloudMessaging = new(DeviceTokenFirebaseCloudMessaging)
		return json.Unmarshal(b, t.DeviceTokenFirebaseCloudMessaging)
	case "deviceTokenHuaweiPush":
		t.DeviceTokenHuaweiPush = new(DeviceTokenHuaweiPush)
		return json.Unmarshal(b, t.DeviceTokenHuaweiPush)
	case "deviceTokenMicrosoftPush":
		t.DeviceTokenMicrosoftPush = new(DeviceTokenMicrosoftPush)
		return json.Unmarshal(b, t.DeviceTokenMicrosoftPush)
	case "deviceTokenMicrosoftPushVoIP":
		t.DeviceTokenMicrosoftPushVoIP = new(DeviceTokenMicrosoftPushVoIP)
		return json.Unmarshal(b, t.DeviceTokenMicrosoftPushVoIP)
	case "deviceTokenSimplePush":
		t.DeviceTokenSimplePush = new(DeviceTokenSimplePush)
		return json.Unmarshal(b, t.DeviceTokenSimplePush)
	case "deviceTokenTizenPush":
		t.DeviceTokenTizenPush = new(DeviceTokenTizenPush)
		return json.Unmarshal(b, t.DeviceTokenTizenPush)
	case "deviceTokenUbuntuPush":
		t.DeviceTokenUbuntuPush = new(DeviceTokenUbuntuPush)
		return json.Unmarshal(b, t.DeviceTokenUbuntuPush)
	case "deviceTokenWebPush":
		t.DeviceTokenWebPush = new(DeviceTokenWebPush)
		return json.Unmarshal(b, t.DeviceTokenWebPush)
	case "deviceTokenWindowsPush":
		t.DeviceTokenWindowsPush = new(DeviceTokenWindowsPush)
		return json.Unmarshal(b, t.DeviceTokenWindowsPush)
	}
	return nil
}

func (t *DeviceToken) MarshalJSON() ([]byte, error) {
	if t.DeviceTokenApplePush != nil {
		return json.Marshal(t.DeviceTokenApplePush)
	}
	if t.DeviceTokenApplePushVoIP != nil {
		return json.Marshal(t.DeviceTokenApplePushVoIP)
	}
	if t.DeviceTokenBlackBerryPush != nil {
		return json.Marshal(t.DeviceTokenBlackBerryPush)
	}
	if t.DeviceTokenFirebaseCloudMessaging != nil {
		return json.Marshal(t.DeviceTokenFirebaseCloudMessaging)
	}
	if t.DeviceTokenHuaweiPush != nil {
		return json.Marshal(t.DeviceTokenHuaweiPush)
	}
	if t.DeviceTokenMicrosoftPush != nil {
		return json.Marshal(t.DeviceTokenMicrosoftPush)
	}
	if t.DeviceTokenMicrosoftPushVoIP != nil {
		return json.Marshal(t.DeviceTokenMicrosoftPushVoIP)
	}
	if t.DeviceTokenSimplePush != nil {
		return json.Marshal(t.DeviceTokenSimplePush)
	}
	if t.DeviceTokenTizenPush != nil {
		return json.Marshal(t.DeviceTokenTizenPush)
	}
	if t.DeviceTokenUbuntuPush != nil {
		return json.Marshal(t.DeviceTokenUbuntuPush)
	}
	if t.DeviceTokenWebPush != nil {
		return json.Marshal(t.DeviceTokenWebPush)
	}
	if t.DeviceTokenWindowsPush != nil {
		return json.Marshal(t.DeviceTokenWindowsPush)
	}
	return nil, fmt.Errorf("no value set for DeviceToken")
}

// DiceStickers Contains animated stickers which must be used for dice animation rendering
type DiceStickers struct {
	TypeStr                 string                   `json:"@type"`
	DiceStickersRegular     *DiceStickersRegular     `json:"diceStickersRegular,omitempty"`
	DiceStickersSlotMachine *DiceStickersSlotMachine `json:"diceStickersSlotMachine,omitempty"`
}

func (t *DiceStickers) Type() string {
	return t.TypeStr
}

func (t *DiceStickers) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *DiceStickers) GetExtra() string {
	return ""
}

func (t *DiceStickers) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "diceStickersRegular":
		t.DiceStickersRegular = new(DiceStickersRegular)
		return json.Unmarshal(b, t.DiceStickersRegular)
	case "diceStickersSlotMachine":
		t.DiceStickersSlotMachine = new(DiceStickersSlotMachine)
		return json.Unmarshal(b, t.DiceStickersSlotMachine)
	}
	return nil
}

func (t *DiceStickers) MarshalJSON() ([]byte, error) {
	if t.DiceStickersRegular != nil {
		return json.Marshal(t.DiceStickersRegular)
	}
	if t.DiceStickersSlotMachine != nil {
		return json.Marshal(t.DiceStickersSlotMachine)
	}
	return nil, fmt.Errorf("no value set for DiceStickers")
}

// EmailAddressAuthentication Contains authentication data for an email address
type EmailAddressAuthentication struct {
	TypeStr                            string                              `json:"@type"`
	EmailAddressAuthenticationAppleId  *EmailAddressAuthenticationAppleId  `json:"emailAddressAuthenticationAppleId,omitempty"`
	EmailAddressAuthenticationCode     *EmailAddressAuthenticationCode     `json:"emailAddressAuthenticationCode,omitempty"`
	EmailAddressAuthenticationGoogleId *EmailAddressAuthenticationGoogleId `json:"emailAddressAuthenticationGoogleId,omitempty"`
}

func (t *EmailAddressAuthentication) Type() string {
	return t.TypeStr
}

func (t *EmailAddressAuthentication) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *EmailAddressAuthentication) GetExtra() string {
	return ""
}

func (t *EmailAddressAuthentication) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "emailAddressAuthenticationAppleId":
		t.EmailAddressAuthenticationAppleId = new(EmailAddressAuthenticationAppleId)
		return json.Unmarshal(b, t.EmailAddressAuthenticationAppleId)
	case "emailAddressAuthenticationCode":
		t.EmailAddressAuthenticationCode = new(EmailAddressAuthenticationCode)
		return json.Unmarshal(b, t.EmailAddressAuthenticationCode)
	case "emailAddressAuthenticationGoogleId":
		t.EmailAddressAuthenticationGoogleId = new(EmailAddressAuthenticationGoogleId)
		return json.Unmarshal(b, t.EmailAddressAuthenticationGoogleId)
	}
	return nil
}

func (t *EmailAddressAuthentication) MarshalJSON() ([]byte, error) {
	if t.EmailAddressAuthenticationAppleId != nil {
		return json.Marshal(t.EmailAddressAuthenticationAppleId)
	}
	if t.EmailAddressAuthenticationCode != nil {
		return json.Marshal(t.EmailAddressAuthenticationCode)
	}
	if t.EmailAddressAuthenticationGoogleId != nil {
		return json.Marshal(t.EmailAddressAuthenticationGoogleId)
	}
	return nil, fmt.Errorf("no value set for EmailAddressAuthentication")
}

// EmailAddressResetState Describes reset state of an email address
type EmailAddressResetState struct {
	TypeStr                         string                           `json:"@type"`
	EmailAddressResetStateAvailable *EmailAddressResetStateAvailable `json:"emailAddressResetStateAvailable,omitempty"`
	EmailAddressResetStatePending   *EmailAddressResetStatePending   `json:"emailAddressResetStatePending,omitempty"`
}

func (t *EmailAddressResetState) Type() string {
	return t.TypeStr
}

func (t *EmailAddressResetState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *EmailAddressResetState) GetExtra() string {
	return ""
}

func (t *EmailAddressResetState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "emailAddressResetStateAvailable":
		t.EmailAddressResetStateAvailable = new(EmailAddressResetStateAvailable)
		return json.Unmarshal(b, t.EmailAddressResetStateAvailable)
	case "emailAddressResetStatePending":
		t.EmailAddressResetStatePending = new(EmailAddressResetStatePending)
		return json.Unmarshal(b, t.EmailAddressResetStatePending)
	}
	return nil
}

func (t *EmailAddressResetState) MarshalJSON() ([]byte, error) {
	if t.EmailAddressResetStateAvailable != nil {
		return json.Marshal(t.EmailAddressResetStateAvailable)
	}
	if t.EmailAddressResetStatePending != nil {
		return json.Marshal(t.EmailAddressResetStatePending)
	}
	return nil, fmt.Errorf("no value set for EmailAddressResetState")
}

// EmojiCategorySource Describes source of stickers for an emoji category
type EmojiCategorySource struct {
	TypeStr                    string                      `json:"@type"`
	EmojiCategorySourcePremium *EmojiCategorySourcePremium `json:"emojiCategorySourcePremium,omitempty"`
	EmojiCategorySourceSearch  *EmojiCategorySourceSearch  `json:"emojiCategorySourceSearch,omitempty"`
}

func (t *EmojiCategorySource) Type() string {
	return t.TypeStr
}

func (t *EmojiCategorySource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *EmojiCategorySource) GetExtra() string {
	return ""
}

func (t *EmojiCategorySource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "emojiCategorySourcePremium":
		t.EmojiCategorySourcePremium = new(EmojiCategorySourcePremium)
		return json.Unmarshal(b, t.EmojiCategorySourcePremium)
	case "emojiCategorySourceSearch":
		t.EmojiCategorySourceSearch = new(EmojiCategorySourceSearch)
		return json.Unmarshal(b, t.EmojiCategorySourceSearch)
	}
	return nil
}

func (t *EmojiCategorySource) MarshalJSON() ([]byte, error) {
	if t.EmojiCategorySourcePremium != nil {
		return json.Marshal(t.EmojiCategorySourcePremium)
	}
	if t.EmojiCategorySourceSearch != nil {
		return json.Marshal(t.EmojiCategorySourceSearch)
	}
	return nil, fmt.Errorf("no value set for EmojiCategorySource")
}

// EmojiCategoryType Describes type of emoji category
type EmojiCategoryType struct {
	TypeStr                          string                            `json:"@type"`
	EmojiCategoryTypeChatPhoto       *EmojiCategoryTypeChatPhoto       `json:"emojiCategoryTypeChatPhoto,omitempty"`
	EmojiCategoryTypeDefault         *EmojiCategoryTypeDefault         `json:"emojiCategoryTypeDefault,omitempty"`
	EmojiCategoryTypeEmojiStatus     *EmojiCategoryTypeEmojiStatus     `json:"emojiCategoryTypeEmojiStatus,omitempty"`
	EmojiCategoryTypeRegularStickers *EmojiCategoryTypeRegularStickers `json:"emojiCategoryTypeRegularStickers,omitempty"`
}

func (t *EmojiCategoryType) Type() string {
	return t.TypeStr
}

func (t *EmojiCategoryType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *EmojiCategoryType) GetExtra() string {
	return ""
}

func (t *EmojiCategoryType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "emojiCategoryTypeChatPhoto":
		t.EmojiCategoryTypeChatPhoto = new(EmojiCategoryTypeChatPhoto)
		return json.Unmarshal(b, t.EmojiCategoryTypeChatPhoto)
	case "emojiCategoryTypeDefault":
		t.EmojiCategoryTypeDefault = new(EmojiCategoryTypeDefault)
		return json.Unmarshal(b, t.EmojiCategoryTypeDefault)
	case "emojiCategoryTypeEmojiStatus":
		t.EmojiCategoryTypeEmojiStatus = new(EmojiCategoryTypeEmojiStatus)
		return json.Unmarshal(b, t.EmojiCategoryTypeEmojiStatus)
	case "emojiCategoryTypeRegularStickers":
		t.EmojiCategoryTypeRegularStickers = new(EmojiCategoryTypeRegularStickers)
		return json.Unmarshal(b, t.EmojiCategoryTypeRegularStickers)
	}
	return nil
}

func (t *EmojiCategoryType) MarshalJSON() ([]byte, error) {
	if t.EmojiCategoryTypeChatPhoto != nil {
		return json.Marshal(t.EmojiCategoryTypeChatPhoto)
	}
	if t.EmojiCategoryTypeDefault != nil {
		return json.Marshal(t.EmojiCategoryTypeDefault)
	}
	if t.EmojiCategoryTypeEmojiStatus != nil {
		return json.Marshal(t.EmojiCategoryTypeEmojiStatus)
	}
	if t.EmojiCategoryTypeRegularStickers != nil {
		return json.Marshal(t.EmojiCategoryTypeRegularStickers)
	}
	return nil, fmt.Errorf("no value set for EmojiCategoryType")
}

// EmojiStatusType Describes type of emoji status
type EmojiStatusType struct {
	TypeStr                     string                       `json:"@type"`
	EmojiStatusTypeCustomEmoji  *EmojiStatusTypeCustomEmoji  `json:"emojiStatusTypeCustomEmoji,omitempty"`
	EmojiStatusTypeUpgradedGift *EmojiStatusTypeUpgradedGift `json:"emojiStatusTypeUpgradedGift,omitempty"`
}

func (t *EmojiStatusType) Type() string {
	return t.TypeStr
}

func (t *EmojiStatusType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *EmojiStatusType) GetExtra() string {
	return ""
}

func (t *EmojiStatusType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "emojiStatusTypeCustomEmoji":
		t.EmojiStatusTypeCustomEmoji = new(EmojiStatusTypeCustomEmoji)
		return json.Unmarshal(b, t.EmojiStatusTypeCustomEmoji)
	case "emojiStatusTypeUpgradedGift":
		t.EmojiStatusTypeUpgradedGift = new(EmojiStatusTypeUpgradedGift)
		return json.Unmarshal(b, t.EmojiStatusTypeUpgradedGift)
	}
	return nil
}

func (t *EmojiStatusType) MarshalJSON() ([]byte, error) {
	if t.EmojiStatusTypeCustomEmoji != nil {
		return json.Marshal(t.EmojiStatusTypeCustomEmoji)
	}
	if t.EmojiStatusTypeUpgradedGift != nil {
		return json.Marshal(t.EmojiStatusTypeUpgradedGift)
	}
	return nil, fmt.Errorf("no value set for EmojiStatusType")
}

// FileType Represents the type of file
type FileType struct {
	TypeStr                          string                            `json:"@type"`
	FileTypeAnimation                *FileTypeAnimation                `json:"fileTypeAnimation,omitempty"`
	FileTypeAudio                    *FileTypeAudio                    `json:"fileTypeAudio,omitempty"`
	FileTypeDocument                 *FileTypeDocument                 `json:"fileTypeDocument,omitempty"`
	FileTypeNone                     *FileTypeNone                     `json:"fileTypeNone,omitempty"`
	FileTypeNotificationSound        *FileTypeNotificationSound        `json:"fileTypeNotificationSound,omitempty"`
	FileTypePhoto                    *FileTypePhoto                    `json:"fileTypePhoto,omitempty"`
	FileTypePhotoStory               *FileTypePhotoStory               `json:"fileTypePhotoStory,omitempty"`
	FileTypeProfilePhoto             *FileTypeProfilePhoto             `json:"fileTypeProfilePhoto,omitempty"`
	FileTypeSecret                   *FileTypeSecret                   `json:"fileTypeSecret,omitempty"`
	FileTypeSecretThumbnail          *FileTypeSecretThumbnail          `json:"fileTypeSecretThumbnail,omitempty"`
	FileTypeSecure                   *FileTypeSecure                   `json:"fileTypeSecure,omitempty"`
	FileTypeSelfDestructingPhoto     *FileTypeSelfDestructingPhoto     `json:"fileTypeSelfDestructingPhoto,omitempty"`
	FileTypeSelfDestructingVideo     *FileTypeSelfDestructingVideo     `json:"fileTypeSelfDestructingVideo,omitempty"`
	FileTypeSelfDestructingVideoNote *FileTypeSelfDestructingVideoNote `json:"fileTypeSelfDestructingVideoNote,omitempty"`
	FileTypeSelfDestructingVoiceNote *FileTypeSelfDestructingVoiceNote `json:"fileTypeSelfDestructingVoiceNote,omitempty"`
	FileTypeSticker                  *FileTypeSticker                  `json:"fileTypeSticker,omitempty"`
	FileTypeThumbnail                *FileTypeThumbnail                `json:"fileTypeThumbnail,omitempty"`
	FileTypeUnknown                  *FileTypeUnknown                  `json:"fileTypeUnknown,omitempty"`
	FileTypeVideo                    *FileTypeVideo                    `json:"fileTypeVideo,omitempty"`
	FileTypeVideoNote                *FileTypeVideoNote                `json:"fileTypeVideoNote,omitempty"`
	FileTypeVideoStory               *FileTypeVideoStory               `json:"fileTypeVideoStory,omitempty"`
	FileTypeVoiceNote                *FileTypeVoiceNote                `json:"fileTypeVoiceNote,omitempty"`
	FileTypeWallpaper                *FileTypeWallpaper                `json:"fileTypeWallpaper,omitempty"`
}

func (t *FileType) Type() string {
	return t.TypeStr
}

func (t *FileType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *FileType) GetExtra() string {
	return ""
}

func (t *FileType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "fileTypeAnimation":
		t.FileTypeAnimation = new(FileTypeAnimation)
		return json.Unmarshal(b, t.FileTypeAnimation)
	case "fileTypeAudio":
		t.FileTypeAudio = new(FileTypeAudio)
		return json.Unmarshal(b, t.FileTypeAudio)
	case "fileTypeDocument":
		t.FileTypeDocument = new(FileTypeDocument)
		return json.Unmarshal(b, t.FileTypeDocument)
	case "fileTypeNone":
		t.FileTypeNone = new(FileTypeNone)
		return json.Unmarshal(b, t.FileTypeNone)
	case "fileTypeNotificationSound":
		t.FileTypeNotificationSound = new(FileTypeNotificationSound)
		return json.Unmarshal(b, t.FileTypeNotificationSound)
	case "fileTypePhoto":
		t.FileTypePhoto = new(FileTypePhoto)
		return json.Unmarshal(b, t.FileTypePhoto)
	case "fileTypePhotoStory":
		t.FileTypePhotoStory = new(FileTypePhotoStory)
		return json.Unmarshal(b, t.FileTypePhotoStory)
	case "fileTypeProfilePhoto":
		t.FileTypeProfilePhoto = new(FileTypeProfilePhoto)
		return json.Unmarshal(b, t.FileTypeProfilePhoto)
	case "fileTypeSecret":
		t.FileTypeSecret = new(FileTypeSecret)
		return json.Unmarshal(b, t.FileTypeSecret)
	case "fileTypeSecretThumbnail":
		t.FileTypeSecretThumbnail = new(FileTypeSecretThumbnail)
		return json.Unmarshal(b, t.FileTypeSecretThumbnail)
	case "fileTypeSecure":
		t.FileTypeSecure = new(FileTypeSecure)
		return json.Unmarshal(b, t.FileTypeSecure)
	case "fileTypeSelfDestructingPhoto":
		t.FileTypeSelfDestructingPhoto = new(FileTypeSelfDestructingPhoto)
		return json.Unmarshal(b, t.FileTypeSelfDestructingPhoto)
	case "fileTypeSelfDestructingVideo":
		t.FileTypeSelfDestructingVideo = new(FileTypeSelfDestructingVideo)
		return json.Unmarshal(b, t.FileTypeSelfDestructingVideo)
	case "fileTypeSelfDestructingVideoNote":
		t.FileTypeSelfDestructingVideoNote = new(FileTypeSelfDestructingVideoNote)
		return json.Unmarshal(b, t.FileTypeSelfDestructingVideoNote)
	case "fileTypeSelfDestructingVoiceNote":
		t.FileTypeSelfDestructingVoiceNote = new(FileTypeSelfDestructingVoiceNote)
		return json.Unmarshal(b, t.FileTypeSelfDestructingVoiceNote)
	case "fileTypeSticker":
		t.FileTypeSticker = new(FileTypeSticker)
		return json.Unmarshal(b, t.FileTypeSticker)
	case "fileTypeThumbnail":
		t.FileTypeThumbnail = new(FileTypeThumbnail)
		return json.Unmarshal(b, t.FileTypeThumbnail)
	case "fileTypeUnknown":
		t.FileTypeUnknown = new(FileTypeUnknown)
		return json.Unmarshal(b, t.FileTypeUnknown)
	case "fileTypeVideo":
		t.FileTypeVideo = new(FileTypeVideo)
		return json.Unmarshal(b, t.FileTypeVideo)
	case "fileTypeVideoNote":
		t.FileTypeVideoNote = new(FileTypeVideoNote)
		return json.Unmarshal(b, t.FileTypeVideoNote)
	case "fileTypeVideoStory":
		t.FileTypeVideoStory = new(FileTypeVideoStory)
		return json.Unmarshal(b, t.FileTypeVideoStory)
	case "fileTypeVoiceNote":
		t.FileTypeVoiceNote = new(FileTypeVoiceNote)
		return json.Unmarshal(b, t.FileTypeVoiceNote)
	case "fileTypeWallpaper":
		t.FileTypeWallpaper = new(FileTypeWallpaper)
		return json.Unmarshal(b, t.FileTypeWallpaper)
	}
	return nil
}

func (t *FileType) MarshalJSON() ([]byte, error) {
	if t.FileTypeAnimation != nil {
		return json.Marshal(t.FileTypeAnimation)
	}
	if t.FileTypeAudio != nil {
		return json.Marshal(t.FileTypeAudio)
	}
	if t.FileTypeDocument != nil {
		return json.Marshal(t.FileTypeDocument)
	}
	if t.FileTypeNone != nil {
		return json.Marshal(t.FileTypeNone)
	}
	if t.FileTypeNotificationSound != nil {
		return json.Marshal(t.FileTypeNotificationSound)
	}
	if t.FileTypePhoto != nil {
		return json.Marshal(t.FileTypePhoto)
	}
	if t.FileTypePhotoStory != nil {
		return json.Marshal(t.FileTypePhotoStory)
	}
	if t.FileTypeProfilePhoto != nil {
		return json.Marshal(t.FileTypeProfilePhoto)
	}
	if t.FileTypeSecret != nil {
		return json.Marshal(t.FileTypeSecret)
	}
	if t.FileTypeSecretThumbnail != nil {
		return json.Marshal(t.FileTypeSecretThumbnail)
	}
	if t.FileTypeSecure != nil {
		return json.Marshal(t.FileTypeSecure)
	}
	if t.FileTypeSelfDestructingPhoto != nil {
		return json.Marshal(t.FileTypeSelfDestructingPhoto)
	}
	if t.FileTypeSelfDestructingVideo != nil {
		return json.Marshal(t.FileTypeSelfDestructingVideo)
	}
	if t.FileTypeSelfDestructingVideoNote != nil {
		return json.Marshal(t.FileTypeSelfDestructingVideoNote)
	}
	if t.FileTypeSelfDestructingVoiceNote != nil {
		return json.Marshal(t.FileTypeSelfDestructingVoiceNote)
	}
	if t.FileTypeSticker != nil {
		return json.Marshal(t.FileTypeSticker)
	}
	if t.FileTypeThumbnail != nil {
		return json.Marshal(t.FileTypeThumbnail)
	}
	if t.FileTypeUnknown != nil {
		return json.Marshal(t.FileTypeUnknown)
	}
	if t.FileTypeVideo != nil {
		return json.Marshal(t.FileTypeVideo)
	}
	if t.FileTypeVideoNote != nil {
		return json.Marshal(t.FileTypeVideoNote)
	}
	if t.FileTypeVideoStory != nil {
		return json.Marshal(t.FileTypeVideoStory)
	}
	if t.FileTypeVoiceNote != nil {
		return json.Marshal(t.FileTypeVoiceNote)
	}
	if t.FileTypeWallpaper != nil {
		return json.Marshal(t.FileTypeWallpaper)
	}
	return nil, fmt.Errorf("no value set for FileType")
}

// FirebaseAuthenticationSettings Contains settings for Firebase Authentication in the official applications
type FirebaseAuthenticationSettings struct {
	TypeStr                               string                                 `json:"@type"`
	FirebaseAuthenticationSettingsAndroid *FirebaseAuthenticationSettingsAndroid `json:"firebaseAuthenticationSettingsAndroid,omitempty"`
	FirebaseAuthenticationSettingsIos     *FirebaseAuthenticationSettingsIos     `json:"firebaseAuthenticationSettingsIos,omitempty"`
}

func (t *FirebaseAuthenticationSettings) Type() string {
	return t.TypeStr
}

func (t *FirebaseAuthenticationSettings) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *FirebaseAuthenticationSettings) GetExtra() string {
	return ""
}

func (t *FirebaseAuthenticationSettings) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "firebaseAuthenticationSettingsAndroid":
		t.FirebaseAuthenticationSettingsAndroid = new(FirebaseAuthenticationSettingsAndroid)
		return json.Unmarshal(b, t.FirebaseAuthenticationSettingsAndroid)
	case "firebaseAuthenticationSettingsIos":
		t.FirebaseAuthenticationSettingsIos = new(FirebaseAuthenticationSettingsIos)
		return json.Unmarshal(b, t.FirebaseAuthenticationSettingsIos)
	}
	return nil
}

func (t *FirebaseAuthenticationSettings) MarshalJSON() ([]byte, error) {
	if t.FirebaseAuthenticationSettingsAndroid != nil {
		return json.Marshal(t.FirebaseAuthenticationSettingsAndroid)
	}
	if t.FirebaseAuthenticationSettingsIos != nil {
		return json.Marshal(t.FirebaseAuthenticationSettingsIos)
	}
	return nil, fmt.Errorf("no value set for FirebaseAuthenticationSettings")
}

// FirebaseDeviceVerificationParameters Describes parameters to be used for device verification
type FirebaseDeviceVerificationParameters struct {
	TypeStr                                           string                                             `json:"@type"`
	FirebaseDeviceVerificationParametersPlayIntegrity *FirebaseDeviceVerificationParametersPlayIntegrity `json:"firebaseDeviceVerificationParametersPlayIntegrity,omitempty"`
	FirebaseDeviceVerificationParametersSafetyNet     *FirebaseDeviceVerificationParametersSafetyNet     `json:"firebaseDeviceVerificationParametersSafetyNet,omitempty"`
}

func (t *FirebaseDeviceVerificationParameters) Type() string {
	return t.TypeStr
}

func (t *FirebaseDeviceVerificationParameters) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *FirebaseDeviceVerificationParameters) GetExtra() string {
	return ""
}

func (t *FirebaseDeviceVerificationParameters) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "firebaseDeviceVerificationParametersPlayIntegrity":
		t.FirebaseDeviceVerificationParametersPlayIntegrity = new(FirebaseDeviceVerificationParametersPlayIntegrity)
		return json.Unmarshal(b, t.FirebaseDeviceVerificationParametersPlayIntegrity)
	case "firebaseDeviceVerificationParametersSafetyNet":
		t.FirebaseDeviceVerificationParametersSafetyNet = new(FirebaseDeviceVerificationParametersSafetyNet)
		return json.Unmarshal(b, t.FirebaseDeviceVerificationParametersSafetyNet)
	}
	return nil
}

func (t *FirebaseDeviceVerificationParameters) MarshalJSON() ([]byte, error) {
	if t.FirebaseDeviceVerificationParametersPlayIntegrity != nil {
		return json.Marshal(t.FirebaseDeviceVerificationParametersPlayIntegrity)
	}
	if t.FirebaseDeviceVerificationParametersSafetyNet != nil {
		return json.Marshal(t.FirebaseDeviceVerificationParametersSafetyNet)
	}
	return nil, fmt.Errorf("no value set for FirebaseDeviceVerificationParameters")
}

// GiftForResaleOrder Describes order in which upgraded gifts for resale will be sorted
type GiftForResaleOrder struct {
	TypeStr                           string                             `json:"@type"`
	GiftForResaleOrderNumber          *GiftForResaleOrderNumber          `json:"giftForResaleOrderNumber,omitempty"`
	GiftForResaleOrderPrice           *GiftForResaleOrderPrice           `json:"giftForResaleOrderPrice,omitempty"`
	GiftForResaleOrderPriceChangeDate *GiftForResaleOrderPriceChangeDate `json:"giftForResaleOrderPriceChangeDate,omitempty"`
}

func (t *GiftForResaleOrder) Type() string {
	return t.TypeStr
}

func (t *GiftForResaleOrder) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiftForResaleOrder) GetExtra() string {
	return ""
}

func (t *GiftForResaleOrder) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giftForResaleOrderNumber":
		t.GiftForResaleOrderNumber = new(GiftForResaleOrderNumber)
		return json.Unmarshal(b, t.GiftForResaleOrderNumber)
	case "giftForResaleOrderPrice":
		t.GiftForResaleOrderPrice = new(GiftForResaleOrderPrice)
		return json.Unmarshal(b, t.GiftForResaleOrderPrice)
	case "giftForResaleOrderPriceChangeDate":
		t.GiftForResaleOrderPriceChangeDate = new(GiftForResaleOrderPriceChangeDate)
		return json.Unmarshal(b, t.GiftForResaleOrderPriceChangeDate)
	}
	return nil
}

func (t *GiftForResaleOrder) MarshalJSON() ([]byte, error) {
	if t.GiftForResaleOrderNumber != nil {
		return json.Marshal(t.GiftForResaleOrderNumber)
	}
	if t.GiftForResaleOrderPrice != nil {
		return json.Marshal(t.GiftForResaleOrderPrice)
	}
	if t.GiftForResaleOrderPriceChangeDate != nil {
		return json.Marshal(t.GiftForResaleOrderPriceChangeDate)
	}
	return nil, fmt.Errorf("no value set for GiftForResaleOrder")
}

// GiftPurchaseOfferState Describes state of a gift purchase offer
type GiftPurchaseOfferState struct {
	TypeStr                        string                          `json:"@type"`
	GiftPurchaseOfferStateAccepted *GiftPurchaseOfferStateAccepted `json:"giftPurchaseOfferStateAccepted,omitempty"`
	GiftPurchaseOfferStatePending  *GiftPurchaseOfferStatePending  `json:"giftPurchaseOfferStatePending,omitempty"`
	GiftPurchaseOfferStateRejected *GiftPurchaseOfferStateRejected `json:"giftPurchaseOfferStateRejected,omitempty"`
}

func (t *GiftPurchaseOfferState) Type() string {
	return t.TypeStr
}

func (t *GiftPurchaseOfferState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiftPurchaseOfferState) GetExtra() string {
	return ""
}

func (t *GiftPurchaseOfferState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giftPurchaseOfferStateAccepted":
		t.GiftPurchaseOfferStateAccepted = new(GiftPurchaseOfferStateAccepted)
		return json.Unmarshal(b, t.GiftPurchaseOfferStateAccepted)
	case "giftPurchaseOfferStatePending":
		t.GiftPurchaseOfferStatePending = new(GiftPurchaseOfferStatePending)
		return json.Unmarshal(b, t.GiftPurchaseOfferStatePending)
	case "giftPurchaseOfferStateRejected":
		t.GiftPurchaseOfferStateRejected = new(GiftPurchaseOfferStateRejected)
		return json.Unmarshal(b, t.GiftPurchaseOfferStateRejected)
	}
	return nil
}

func (t *GiftPurchaseOfferState) MarshalJSON() ([]byte, error) {
	if t.GiftPurchaseOfferStateAccepted != nil {
		return json.Marshal(t.GiftPurchaseOfferStateAccepted)
	}
	if t.GiftPurchaseOfferStatePending != nil {
		return json.Marshal(t.GiftPurchaseOfferStatePending)
	}
	if t.GiftPurchaseOfferStateRejected != nil {
		return json.Marshal(t.GiftPurchaseOfferStateRejected)
	}
	return nil, fmt.Errorf("no value set for GiftPurchaseOfferState")
}

// GiftResalePrice Describes price of a resold gift
type GiftResalePrice struct {
	TypeStr             string               `json:"@type"`
	GiftResalePriceStar *GiftResalePriceStar `json:"giftResalePriceStar,omitempty"`
	GiftResalePriceTon  *GiftResalePriceTon  `json:"giftResalePriceTon,omitempty"`
}

func (t *GiftResalePrice) Type() string {
	return t.TypeStr
}

func (t *GiftResalePrice) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiftResalePrice) GetExtra() string {
	return ""
}

func (t *GiftResalePrice) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giftResalePriceStar":
		t.GiftResalePriceStar = new(GiftResalePriceStar)
		return json.Unmarshal(b, t.GiftResalePriceStar)
	case "giftResalePriceTon":
		t.GiftResalePriceTon = new(GiftResalePriceTon)
		return json.Unmarshal(b, t.GiftResalePriceTon)
	}
	return nil
}

func (t *GiftResalePrice) MarshalJSON() ([]byte, error) {
	if t.GiftResalePriceStar != nil {
		return json.Marshal(t.GiftResalePriceStar)
	}
	if t.GiftResalePriceTon != nil {
		return json.Marshal(t.GiftResalePriceTon)
	}
	return nil, fmt.Errorf("no value set for GiftResalePrice")
}

// GiftResaleResult Describes result of sending a resold gift
type GiftResaleResult struct {
	TypeStr                        string                          `json:"@type"`
	GiftResaleResultOk             *GiftResaleResultOk             `json:"giftResaleResultOk,omitempty"`
	GiftResaleResultPriceIncreased *GiftResaleResultPriceIncreased `json:"giftResaleResultPriceIncreased,omitempty"`
}

func (t *GiftResaleResult) Type() string {
	return t.TypeStr
}

func (t *GiftResaleResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiftResaleResult) GetExtra() string {
	return ""
}

func (t *GiftResaleResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giftResaleResultOk":
		t.GiftResaleResultOk = new(GiftResaleResultOk)
		return json.Unmarshal(b, t.GiftResaleResultOk)
	case "giftResaleResultPriceIncreased":
		t.GiftResaleResultPriceIncreased = new(GiftResaleResultPriceIncreased)
		return json.Unmarshal(b, t.GiftResaleResultPriceIncreased)
	}
	return nil
}

func (t *GiftResaleResult) MarshalJSON() ([]byte, error) {
	if t.GiftResaleResultOk != nil {
		return json.Marshal(t.GiftResaleResultOk)
	}
	if t.GiftResaleResultPriceIncreased != nil {
		return json.Marshal(t.GiftResaleResultPriceIncreased)
	}
	return nil, fmt.Errorf("no value set for GiftResaleResult")
}

// GiveawayInfo Contains information about a giveaway
type GiveawayInfo struct {
	TypeStr               string                 `json:"@type"`
	GiveawayInfoCompleted *GiveawayInfoCompleted `json:"giveawayInfoCompleted,omitempty"`
	GiveawayInfoOngoing   *GiveawayInfoOngoing   `json:"giveawayInfoOngoing,omitempty"`
}

func (t *GiveawayInfo) Type() string {
	return t.TypeStr
}

func (t *GiveawayInfo) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiveawayInfo) GetExtra() string {
	return ""
}

func (t *GiveawayInfo) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giveawayInfoCompleted":
		t.GiveawayInfoCompleted = new(GiveawayInfoCompleted)
		return json.Unmarshal(b, t.GiveawayInfoCompleted)
	case "giveawayInfoOngoing":
		t.GiveawayInfoOngoing = new(GiveawayInfoOngoing)
		return json.Unmarshal(b, t.GiveawayInfoOngoing)
	}
	return nil
}

func (t *GiveawayInfo) MarshalJSON() ([]byte, error) {
	if t.GiveawayInfoCompleted != nil {
		return json.Marshal(t.GiveawayInfoCompleted)
	}
	if t.GiveawayInfoOngoing != nil {
		return json.Marshal(t.GiveawayInfoOngoing)
	}
	return nil, fmt.Errorf("no value set for GiveawayInfo")
}

// GiveawayParticipantStatus Contains information about status of a user in a giveaway
type GiveawayParticipantStatus struct {
	TypeStr                                    string                                      `json:"@type"`
	GiveawayParticipantStatusAdministrator     *GiveawayParticipantStatusAdministrator     `json:"giveawayParticipantStatusAdministrator,omitempty"`
	GiveawayParticipantStatusAlreadyWasMember  *GiveawayParticipantStatusAlreadyWasMember  `json:"giveawayParticipantStatusAlreadyWasMember,omitempty"`
	GiveawayParticipantStatusDisallowedCountry *GiveawayParticipantStatusDisallowedCountry `json:"giveawayParticipantStatusDisallowedCountry,omitempty"`
	GiveawayParticipantStatusEligible          *GiveawayParticipantStatusEligible          `json:"giveawayParticipantStatusEligible,omitempty"`
	GiveawayParticipantStatusParticipating     *GiveawayParticipantStatusParticipating     `json:"giveawayParticipantStatusParticipating,omitempty"`
}

func (t *GiveawayParticipantStatus) Type() string {
	return t.TypeStr
}

func (t *GiveawayParticipantStatus) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiveawayParticipantStatus) GetExtra() string {
	return ""
}

func (t *GiveawayParticipantStatus) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giveawayParticipantStatusAdministrator":
		t.GiveawayParticipantStatusAdministrator = new(GiveawayParticipantStatusAdministrator)
		return json.Unmarshal(b, t.GiveawayParticipantStatusAdministrator)
	case "giveawayParticipantStatusAlreadyWasMember":
		t.GiveawayParticipantStatusAlreadyWasMember = new(GiveawayParticipantStatusAlreadyWasMember)
		return json.Unmarshal(b, t.GiveawayParticipantStatusAlreadyWasMember)
	case "giveawayParticipantStatusDisallowedCountry":
		t.GiveawayParticipantStatusDisallowedCountry = new(GiveawayParticipantStatusDisallowedCountry)
		return json.Unmarshal(b, t.GiveawayParticipantStatusDisallowedCountry)
	case "giveawayParticipantStatusEligible":
		t.GiveawayParticipantStatusEligible = new(GiveawayParticipantStatusEligible)
		return json.Unmarshal(b, t.GiveawayParticipantStatusEligible)
	case "giveawayParticipantStatusParticipating":
		t.GiveawayParticipantStatusParticipating = new(GiveawayParticipantStatusParticipating)
		return json.Unmarshal(b, t.GiveawayParticipantStatusParticipating)
	}
	return nil
}

func (t *GiveawayParticipantStatus) MarshalJSON() ([]byte, error) {
	if t.GiveawayParticipantStatusAdministrator != nil {
		return json.Marshal(t.GiveawayParticipantStatusAdministrator)
	}
	if t.GiveawayParticipantStatusAlreadyWasMember != nil {
		return json.Marshal(t.GiveawayParticipantStatusAlreadyWasMember)
	}
	if t.GiveawayParticipantStatusDisallowedCountry != nil {
		return json.Marshal(t.GiveawayParticipantStatusDisallowedCountry)
	}
	if t.GiveawayParticipantStatusEligible != nil {
		return json.Marshal(t.GiveawayParticipantStatusEligible)
	}
	if t.GiveawayParticipantStatusParticipating != nil {
		return json.Marshal(t.GiveawayParticipantStatusParticipating)
	}
	return nil, fmt.Errorf("no value set for GiveawayParticipantStatus")
}

// GiveawayPrize Contains information about a giveaway prize
type GiveawayPrize struct {
	TypeStr              string                `json:"@type"`
	GiveawayPrizePremium *GiveawayPrizePremium `json:"giveawayPrizePremium,omitempty"`
	GiveawayPrizeStars   *GiveawayPrizeStars   `json:"giveawayPrizeStars,omitempty"`
}

func (t *GiveawayPrize) Type() string {
	return t.TypeStr
}

func (t *GiveawayPrize) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GiveawayPrize) GetExtra() string {
	return ""
}

func (t *GiveawayPrize) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "giveawayPrizePremium":
		t.GiveawayPrizePremium = new(GiveawayPrizePremium)
		return json.Unmarshal(b, t.GiveawayPrizePremium)
	case "giveawayPrizeStars":
		t.GiveawayPrizeStars = new(GiveawayPrizeStars)
		return json.Unmarshal(b, t.GiveawayPrizeStars)
	}
	return nil
}

func (t *GiveawayPrize) MarshalJSON() ([]byte, error) {
	if t.GiveawayPrizePremium != nil {
		return json.Marshal(t.GiveawayPrizePremium)
	}
	if t.GiveawayPrizeStars != nil {
		return json.Marshal(t.GiveawayPrizeStars)
	}
	return nil, fmt.Errorf("no value set for GiveawayPrize")
}

// GroupCallDataChannel Describes data channel for a group call
type GroupCallDataChannel struct {
	TypeStr                           string                             `json:"@type"`
	GroupCallDataChannelMain          *GroupCallDataChannelMain          `json:"groupCallDataChannelMain,omitempty"`
	GroupCallDataChannelScreenSharing *GroupCallDataChannelScreenSharing `json:"groupCallDataChannelScreenSharing,omitempty"`
}

func (t *GroupCallDataChannel) Type() string {
	return t.TypeStr
}

func (t *GroupCallDataChannel) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GroupCallDataChannel) GetExtra() string {
	return ""
}

func (t *GroupCallDataChannel) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "groupCallDataChannelMain":
		t.GroupCallDataChannelMain = new(GroupCallDataChannelMain)
		return json.Unmarshal(b, t.GroupCallDataChannelMain)
	case "groupCallDataChannelScreenSharing":
		t.GroupCallDataChannelScreenSharing = new(GroupCallDataChannelScreenSharing)
		return json.Unmarshal(b, t.GroupCallDataChannelScreenSharing)
	}
	return nil
}

func (t *GroupCallDataChannel) MarshalJSON() ([]byte, error) {
	if t.GroupCallDataChannelMain != nil {
		return json.Marshal(t.GroupCallDataChannelMain)
	}
	if t.GroupCallDataChannelScreenSharing != nil {
		return json.Marshal(t.GroupCallDataChannelScreenSharing)
	}
	return nil, fmt.Errorf("no value set for GroupCallDataChannel")
}

// GroupCallVideoQuality Describes the quality of a group call video
type GroupCallVideoQuality struct {
	TypeStr                        string                          `json:"@type"`
	GroupCallVideoQualityFull      *GroupCallVideoQualityFull      `json:"groupCallVideoQualityFull,omitempty"`
	GroupCallVideoQualityMedium    *GroupCallVideoQualityMedium    `json:"groupCallVideoQualityMedium,omitempty"`
	GroupCallVideoQualityThumbnail *GroupCallVideoQualityThumbnail `json:"groupCallVideoQualityThumbnail,omitempty"`
}

func (t *GroupCallVideoQuality) Type() string {
	return t.TypeStr
}

func (t *GroupCallVideoQuality) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *GroupCallVideoQuality) GetExtra() string {
	return ""
}

func (t *GroupCallVideoQuality) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "groupCallVideoQualityFull":
		t.GroupCallVideoQualityFull = new(GroupCallVideoQualityFull)
		return json.Unmarshal(b, t.GroupCallVideoQualityFull)
	case "groupCallVideoQualityMedium":
		t.GroupCallVideoQualityMedium = new(GroupCallVideoQualityMedium)
		return json.Unmarshal(b, t.GroupCallVideoQualityMedium)
	case "groupCallVideoQualityThumbnail":
		t.GroupCallVideoQualityThumbnail = new(GroupCallVideoQualityThumbnail)
		return json.Unmarshal(b, t.GroupCallVideoQualityThumbnail)
	}
	return nil
}

func (t *GroupCallVideoQuality) MarshalJSON() ([]byte, error) {
	if t.GroupCallVideoQualityFull != nil {
		return json.Marshal(t.GroupCallVideoQualityFull)
	}
	if t.GroupCallVideoQualityMedium != nil {
		return json.Marshal(t.GroupCallVideoQualityMedium)
	}
	if t.GroupCallVideoQualityThumbnail != nil {
		return json.Marshal(t.GroupCallVideoQualityThumbnail)
	}
	return nil, fmt.Errorf("no value set for GroupCallVideoQuality")
}

// InlineKeyboardButtonType Describes the type of inline keyboard button
type InlineKeyboardButtonType struct {
	TypeStr                                      string                                        `json:"@type"`
	InlineKeyboardButtonTypeBuy                  *InlineKeyboardButtonTypeBuy                  `json:"inlineKeyboardButtonTypeBuy,omitempty"`
	InlineKeyboardButtonTypeCallback             *InlineKeyboardButtonTypeCallback             `json:"inlineKeyboardButtonTypeCallback,omitempty"`
	InlineKeyboardButtonTypeCallbackGame         *InlineKeyboardButtonTypeCallbackGame         `json:"inlineKeyboardButtonTypeCallbackGame,omitempty"`
	InlineKeyboardButtonTypeCallbackWithPassword *InlineKeyboardButtonTypeCallbackWithPassword `json:"inlineKeyboardButtonTypeCallbackWithPassword,omitempty"`
	InlineKeyboardButtonTypeCopyText             *InlineKeyboardButtonTypeCopyText             `json:"inlineKeyboardButtonTypeCopyText,omitempty"`
	InlineKeyboardButtonTypeLoginUrl             *InlineKeyboardButtonTypeLoginUrl             `json:"inlineKeyboardButtonTypeLoginUrl,omitempty"`
	InlineKeyboardButtonTypeSwitchInline         *InlineKeyboardButtonTypeSwitchInline         `json:"inlineKeyboardButtonTypeSwitchInline,omitempty"`
	InlineKeyboardButtonTypeUrl                  *InlineKeyboardButtonTypeUrl                  `json:"inlineKeyboardButtonTypeUrl,omitempty"`
	InlineKeyboardButtonTypeUser                 *InlineKeyboardButtonTypeUser                 `json:"inlineKeyboardButtonTypeUser,omitempty"`
	InlineKeyboardButtonTypeWebApp               *InlineKeyboardButtonTypeWebApp               `json:"inlineKeyboardButtonTypeWebApp,omitempty"`
}

func (t *InlineKeyboardButtonType) Type() string {
	return t.TypeStr
}

func (t *InlineKeyboardButtonType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InlineKeyboardButtonType) GetExtra() string {
	return ""
}

func (t *InlineKeyboardButtonType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inlineKeyboardButtonTypeBuy":
		t.InlineKeyboardButtonTypeBuy = new(InlineKeyboardButtonTypeBuy)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeBuy)
	case "inlineKeyboardButtonTypeCallback":
		t.InlineKeyboardButtonTypeCallback = new(InlineKeyboardButtonTypeCallback)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallback)
	case "inlineKeyboardButtonTypeCallbackGame":
		t.InlineKeyboardButtonTypeCallbackGame = new(InlineKeyboardButtonTypeCallbackGame)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallbackGame)
	case "inlineKeyboardButtonTypeCallbackWithPassword":
		t.InlineKeyboardButtonTypeCallbackWithPassword = new(InlineKeyboardButtonTypeCallbackWithPassword)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallbackWithPassword)
	case "inlineKeyboardButtonTypeCopyText":
		t.InlineKeyboardButtonTypeCopyText = new(InlineKeyboardButtonTypeCopyText)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCopyText)
	case "inlineKeyboardButtonTypeLoginUrl":
		t.InlineKeyboardButtonTypeLoginUrl = new(InlineKeyboardButtonTypeLoginUrl)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeLoginUrl)
	case "inlineKeyboardButtonTypeSwitchInline":
		t.InlineKeyboardButtonTypeSwitchInline = new(InlineKeyboardButtonTypeSwitchInline)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeSwitchInline)
	case "inlineKeyboardButtonTypeUrl":
		t.InlineKeyboardButtonTypeUrl = new(InlineKeyboardButtonTypeUrl)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeUrl)
	case "inlineKeyboardButtonTypeUser":
		t.InlineKeyboardButtonTypeUser = new(InlineKeyboardButtonTypeUser)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeUser)
	case "inlineKeyboardButtonTypeWebApp":
		t.InlineKeyboardButtonTypeWebApp = new(InlineKeyboardButtonTypeWebApp)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeWebApp)
	}
	return nil
}

func (t *InlineKeyboardButtonType) MarshalJSON() ([]byte, error) {
	if t.InlineKeyboardButtonTypeBuy != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeBuy)
	}
	if t.InlineKeyboardButtonTypeCallback != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallback)
	}
	if t.InlineKeyboardButtonTypeCallbackGame != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallbackGame)
	}
	if t.InlineKeyboardButtonTypeCallbackWithPassword != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallbackWithPassword)
	}
	if t.InlineKeyboardButtonTypeCopyText != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCopyText)
	}
	if t.InlineKeyboardButtonTypeLoginUrl != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeLoginUrl)
	}
	if t.InlineKeyboardButtonTypeSwitchInline != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeSwitchInline)
	}
	if t.InlineKeyboardButtonTypeUrl != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeUrl)
	}
	if t.InlineKeyboardButtonTypeUser != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeUser)
	}
	if t.InlineKeyboardButtonTypeWebApp != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for InlineKeyboardButtonType")
}

// InlineQueryResult Represents a single result of an inline query
type InlineQueryResult struct {
	TypeStr                    string                      `json:"@type"`
	InlineQueryResultAnimation *InlineQueryResultAnimation `json:"inlineQueryResultAnimation,omitempty"`
	InlineQueryResultArticle   *InlineQueryResultArticle   `json:"inlineQueryResultArticle,omitempty"`
	InlineQueryResultAudio     *InlineQueryResultAudio     `json:"inlineQueryResultAudio,omitempty"`
	InlineQueryResultContact   *InlineQueryResultContact   `json:"inlineQueryResultContact,omitempty"`
	InlineQueryResultDocument  *InlineQueryResultDocument  `json:"inlineQueryResultDocument,omitempty"`
	InlineQueryResultGame      *InlineQueryResultGame      `json:"inlineQueryResultGame,omitempty"`
	InlineQueryResultLocation  *InlineQueryResultLocation  `json:"inlineQueryResultLocation,omitempty"`
	InlineQueryResultPhoto     *InlineQueryResultPhoto     `json:"inlineQueryResultPhoto,omitempty"`
	InlineQueryResultSticker   *InlineQueryResultSticker   `json:"inlineQueryResultSticker,omitempty"`
	InlineQueryResultVenue     *InlineQueryResultVenue     `json:"inlineQueryResultVenue,omitempty"`
	InlineQueryResultVideo     *InlineQueryResultVideo     `json:"inlineQueryResultVideo,omitempty"`
	InlineQueryResultVoiceNote *InlineQueryResultVoiceNote `json:"inlineQueryResultVoiceNote,omitempty"`
}

func (t *InlineQueryResult) Type() string {
	return t.TypeStr
}

func (t *InlineQueryResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InlineQueryResult) GetExtra() string {
	return ""
}

func (t *InlineQueryResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inlineQueryResultAnimation":
		t.InlineQueryResultAnimation = new(InlineQueryResultAnimation)
		return json.Unmarshal(b, t.InlineQueryResultAnimation)
	case "inlineQueryResultArticle":
		t.InlineQueryResultArticle = new(InlineQueryResultArticle)
		return json.Unmarshal(b, t.InlineQueryResultArticle)
	case "inlineQueryResultAudio":
		t.InlineQueryResultAudio = new(InlineQueryResultAudio)
		return json.Unmarshal(b, t.InlineQueryResultAudio)
	case "inlineQueryResultContact":
		t.InlineQueryResultContact = new(InlineQueryResultContact)
		return json.Unmarshal(b, t.InlineQueryResultContact)
	case "inlineQueryResultDocument":
		t.InlineQueryResultDocument = new(InlineQueryResultDocument)
		return json.Unmarshal(b, t.InlineQueryResultDocument)
	case "inlineQueryResultGame":
		t.InlineQueryResultGame = new(InlineQueryResultGame)
		return json.Unmarshal(b, t.InlineQueryResultGame)
	case "inlineQueryResultLocation":
		t.InlineQueryResultLocation = new(InlineQueryResultLocation)
		return json.Unmarshal(b, t.InlineQueryResultLocation)
	case "inlineQueryResultPhoto":
		t.InlineQueryResultPhoto = new(InlineQueryResultPhoto)
		return json.Unmarshal(b, t.InlineQueryResultPhoto)
	case "inlineQueryResultSticker":
		t.InlineQueryResultSticker = new(InlineQueryResultSticker)
		return json.Unmarshal(b, t.InlineQueryResultSticker)
	case "inlineQueryResultVenue":
		t.InlineQueryResultVenue = new(InlineQueryResultVenue)
		return json.Unmarshal(b, t.InlineQueryResultVenue)
	case "inlineQueryResultVideo":
		t.InlineQueryResultVideo = new(InlineQueryResultVideo)
		return json.Unmarshal(b, t.InlineQueryResultVideo)
	case "inlineQueryResultVoiceNote":
		t.InlineQueryResultVoiceNote = new(InlineQueryResultVoiceNote)
		return json.Unmarshal(b, t.InlineQueryResultVoiceNote)
	}
	return nil
}

func (t *InlineQueryResult) MarshalJSON() ([]byte, error) {
	if t.InlineQueryResultAnimation != nil {
		return json.Marshal(t.InlineQueryResultAnimation)
	}
	if t.InlineQueryResultArticle != nil {
		return json.Marshal(t.InlineQueryResultArticle)
	}
	if t.InlineQueryResultAudio != nil {
		return json.Marshal(t.InlineQueryResultAudio)
	}
	if t.InlineQueryResultContact != nil {
		return json.Marshal(t.InlineQueryResultContact)
	}
	if t.InlineQueryResultDocument != nil {
		return json.Marshal(t.InlineQueryResultDocument)
	}
	if t.InlineQueryResultGame != nil {
		return json.Marshal(t.InlineQueryResultGame)
	}
	if t.InlineQueryResultLocation != nil {
		return json.Marshal(t.InlineQueryResultLocation)
	}
	if t.InlineQueryResultPhoto != nil {
		return json.Marshal(t.InlineQueryResultPhoto)
	}
	if t.InlineQueryResultSticker != nil {
		return json.Marshal(t.InlineQueryResultSticker)
	}
	if t.InlineQueryResultVenue != nil {
		return json.Marshal(t.InlineQueryResultVenue)
	}
	if t.InlineQueryResultVideo != nil {
		return json.Marshal(t.InlineQueryResultVideo)
	}
	if t.InlineQueryResultVoiceNote != nil {
		return json.Marshal(t.InlineQueryResultVoiceNote)
	}
	return nil, fmt.Errorf("no value set for InlineQueryResult")
}

// InlineQueryResultsButtonType Represents type of button in results of inline query
type InlineQueryResultsButtonType struct {
	TypeStr                              string                                `json:"@type"`
	InlineQueryResultsButtonTypeStartBot *InlineQueryResultsButtonTypeStartBot `json:"inlineQueryResultsButtonTypeStartBot,omitempty"`
	InlineQueryResultsButtonTypeWebApp   *InlineQueryResultsButtonTypeWebApp   `json:"inlineQueryResultsButtonTypeWebApp,omitempty"`
}

func (t *InlineQueryResultsButtonType) Type() string {
	return t.TypeStr
}

func (t *InlineQueryResultsButtonType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InlineQueryResultsButtonType) GetExtra() string {
	return ""
}

func (t *InlineQueryResultsButtonType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inlineQueryResultsButtonTypeStartBot":
		t.InlineQueryResultsButtonTypeStartBot = new(InlineQueryResultsButtonTypeStartBot)
		return json.Unmarshal(b, t.InlineQueryResultsButtonTypeStartBot)
	case "inlineQueryResultsButtonTypeWebApp":
		t.InlineQueryResultsButtonTypeWebApp = new(InlineQueryResultsButtonTypeWebApp)
		return json.Unmarshal(b, t.InlineQueryResultsButtonTypeWebApp)
	}
	return nil
}

func (t *InlineQueryResultsButtonType) MarshalJSON() ([]byte, error) {
	if t.InlineQueryResultsButtonTypeStartBot != nil {
		return json.Marshal(t.InlineQueryResultsButtonTypeStartBot)
	}
	if t.InlineQueryResultsButtonTypeWebApp != nil {
		return json.Marshal(t.InlineQueryResultsButtonTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for InlineQueryResultsButtonType")
}

// InputBackground Contains information about background to set
type InputBackground struct {
	TypeStr                 string                   `json:"@type"`
	InputBackgroundLocal    *InputBackgroundLocal    `json:"inputBackgroundLocal,omitempty"`
	InputBackgroundPrevious *InputBackgroundPrevious `json:"inputBackgroundPrevious,omitempty"`
	InputBackgroundRemote   *InputBackgroundRemote   `json:"inputBackgroundRemote,omitempty"`
}

func (t *InputBackground) Type() string {
	return t.TypeStr
}

func (t *InputBackground) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputBackground) GetExtra() string {
	return ""
}

func (t *InputBackground) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputBackgroundLocal":
		t.InputBackgroundLocal = new(InputBackgroundLocal)
		return json.Unmarshal(b, t.InputBackgroundLocal)
	case "inputBackgroundPrevious":
		t.InputBackgroundPrevious = new(InputBackgroundPrevious)
		return json.Unmarshal(b, t.InputBackgroundPrevious)
	case "inputBackgroundRemote":
		t.InputBackgroundRemote = new(InputBackgroundRemote)
		return json.Unmarshal(b, t.InputBackgroundRemote)
	}
	return nil
}

func (t *InputBackground) MarshalJSON() ([]byte, error) {
	if t.InputBackgroundLocal != nil {
		return json.Marshal(t.InputBackgroundLocal)
	}
	if t.InputBackgroundPrevious != nil {
		return json.Marshal(t.InputBackgroundPrevious)
	}
	if t.InputBackgroundRemote != nil {
		return json.Marshal(t.InputBackgroundRemote)
	}
	return nil, fmt.Errorf("no value set for InputBackground")
}

// InputChatPhoto Describes a photo to be set as a user profile or chat photo
type InputChatPhoto struct {
	TypeStr                 string                   `json:"@type"`
	InputChatPhotoAnimation *InputChatPhotoAnimation `json:"inputChatPhotoAnimation,omitempty"`
	InputChatPhotoPrevious  *InputChatPhotoPrevious  `json:"inputChatPhotoPrevious,omitempty"`
	InputChatPhotoStatic    *InputChatPhotoStatic    `json:"inputChatPhotoStatic,omitempty"`
	InputChatPhotoSticker   *InputChatPhotoSticker   `json:"inputChatPhotoSticker,omitempty"`
}

func (t *InputChatPhoto) Type() string {
	return t.TypeStr
}

func (t *InputChatPhoto) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputChatPhoto) GetExtra() string {
	return ""
}

func (t *InputChatPhoto) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputChatPhotoAnimation":
		t.InputChatPhotoAnimation = new(InputChatPhotoAnimation)
		return json.Unmarshal(b, t.InputChatPhotoAnimation)
	case "inputChatPhotoPrevious":
		t.InputChatPhotoPrevious = new(InputChatPhotoPrevious)
		return json.Unmarshal(b, t.InputChatPhotoPrevious)
	case "inputChatPhotoStatic":
		t.InputChatPhotoStatic = new(InputChatPhotoStatic)
		return json.Unmarshal(b, t.InputChatPhotoStatic)
	case "inputChatPhotoSticker":
		t.InputChatPhotoSticker = new(InputChatPhotoSticker)
		return json.Unmarshal(b, t.InputChatPhotoSticker)
	}
	return nil
}

func (t *InputChatPhoto) MarshalJSON() ([]byte, error) {
	if t.InputChatPhotoAnimation != nil {
		return json.Marshal(t.InputChatPhotoAnimation)
	}
	if t.InputChatPhotoPrevious != nil {
		return json.Marshal(t.InputChatPhotoPrevious)
	}
	if t.InputChatPhotoStatic != nil {
		return json.Marshal(t.InputChatPhotoStatic)
	}
	if t.InputChatPhotoSticker != nil {
		return json.Marshal(t.InputChatPhotoSticker)
	}
	return nil, fmt.Errorf("no value set for InputChatPhoto")
}

// InputChatTheme Describes a chat theme to set
type InputChatTheme struct {
	TypeStr             string               `json:"@type"`
	InputChatThemeEmoji *InputChatThemeEmoji `json:"inputChatThemeEmoji,omitempty"`
	InputChatThemeGift  *InputChatThemeGift  `json:"inputChatThemeGift,omitempty"`
}

func (t *InputChatTheme) Type() string {
	return t.TypeStr
}

func (t *InputChatTheme) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputChatTheme) GetExtra() string {
	return ""
}

func (t *InputChatTheme) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputChatThemeEmoji":
		t.InputChatThemeEmoji = new(InputChatThemeEmoji)
		return json.Unmarshal(b, t.InputChatThemeEmoji)
	case "inputChatThemeGift":
		t.InputChatThemeGift = new(InputChatThemeGift)
		return json.Unmarshal(b, t.InputChatThemeGift)
	}
	return nil
}

func (t *InputChatTheme) MarshalJSON() ([]byte, error) {
	if t.InputChatThemeEmoji != nil {
		return json.Marshal(t.InputChatThemeEmoji)
	}
	if t.InputChatThemeGift != nil {
		return json.Marshal(t.InputChatThemeGift)
	}
	return nil, fmt.Errorf("no value set for InputChatTheme")
}

// InputCredentials Contains information about the payment method chosen by the user
type InputCredentials struct {
	TypeStr                   string                     `json:"@type"`
	InputCredentialsApplePay  *InputCredentialsApplePay  `json:"inputCredentialsApplePay,omitempty"`
	InputCredentialsGooglePay *InputCredentialsGooglePay `json:"inputCredentialsGooglePay,omitempty"`
	InputCredentialsNew       *InputCredentialsNew       `json:"inputCredentialsNew,omitempty"`
	InputCredentialsSaved     *InputCredentialsSaved     `json:"inputCredentialsSaved,omitempty"`
}

func (t *InputCredentials) Type() string {
	return t.TypeStr
}

func (t *InputCredentials) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputCredentials) GetExtra() string {
	return ""
}

func (t *InputCredentials) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputCredentialsApplePay":
		t.InputCredentialsApplePay = new(InputCredentialsApplePay)
		return json.Unmarshal(b, t.InputCredentialsApplePay)
	case "inputCredentialsGooglePay":
		t.InputCredentialsGooglePay = new(InputCredentialsGooglePay)
		return json.Unmarshal(b, t.InputCredentialsGooglePay)
	case "inputCredentialsNew":
		t.InputCredentialsNew = new(InputCredentialsNew)
		return json.Unmarshal(b, t.InputCredentialsNew)
	case "inputCredentialsSaved":
		t.InputCredentialsSaved = new(InputCredentialsSaved)
		return json.Unmarshal(b, t.InputCredentialsSaved)
	}
	return nil
}

func (t *InputCredentials) MarshalJSON() ([]byte, error) {
	if t.InputCredentialsApplePay != nil {
		return json.Marshal(t.InputCredentialsApplePay)
	}
	if t.InputCredentialsGooglePay != nil {
		return json.Marshal(t.InputCredentialsGooglePay)
	}
	if t.InputCredentialsNew != nil {
		return json.Marshal(t.InputCredentialsNew)
	}
	if t.InputCredentialsSaved != nil {
		return json.Marshal(t.InputCredentialsSaved)
	}
	return nil, fmt.Errorf("no value set for InputCredentials")
}

// InputFile Points to a file
type InputFile struct {
	TypeStr            string              `json:"@type"`
	InputFileGenerated *InputFileGenerated `json:"inputFileGenerated,omitempty"`
	InputFileId        *InputFileId        `json:"inputFileId,omitempty"`
	InputFileLocal     *InputFileLocal     `json:"inputFileLocal,omitempty"`
	InputFileRemote    *InputFileRemote    `json:"inputFileRemote,omitempty"`
}

func (t *InputFile) Type() string {
	return t.TypeStr
}

func (t *InputFile) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputFile) GetExtra() string {
	return ""
}

func (t *InputFile) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputFileGenerated":
		t.InputFileGenerated = new(InputFileGenerated)
		return json.Unmarshal(b, t.InputFileGenerated)
	case "inputFileId":
		t.InputFileId = new(InputFileId)
		return json.Unmarshal(b, t.InputFileId)
	case "inputFileLocal":
		t.InputFileLocal = new(InputFileLocal)
		return json.Unmarshal(b, t.InputFileLocal)
	case "inputFileRemote":
		t.InputFileRemote = new(InputFileRemote)
		return json.Unmarshal(b, t.InputFileRemote)
	}
	return nil
}

func (t *InputFile) MarshalJSON() ([]byte, error) {
	if t.InputFileGenerated != nil {
		return json.Marshal(t.InputFileGenerated)
	}
	if t.InputFileId != nil {
		return json.Marshal(t.InputFileId)
	}
	if t.InputFileLocal != nil {
		return json.Marshal(t.InputFileLocal)
	}
	if t.InputFileRemote != nil {
		return json.Marshal(t.InputFileRemote)
	}
	return nil, fmt.Errorf("no value set for InputFile")
}

// InputGroupCall Describes a non-joined group call that isn't bound to a chat
type InputGroupCall struct {
	TypeStr               string                 `json:"@type"`
	InputGroupCallLink    *InputGroupCallLink    `json:"inputGroupCallLink,omitempty"`
	InputGroupCallMessage *InputGroupCallMessage `json:"inputGroupCallMessage,omitempty"`
}

func (t *InputGroupCall) Type() string {
	return t.TypeStr
}

func (t *InputGroupCall) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputGroupCall) GetExtra() string {
	return ""
}

func (t *InputGroupCall) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputGroupCallLink":
		t.InputGroupCallLink = new(InputGroupCallLink)
		return json.Unmarshal(b, t.InputGroupCallLink)
	case "inputGroupCallMessage":
		t.InputGroupCallMessage = new(InputGroupCallMessage)
		return json.Unmarshal(b, t.InputGroupCallMessage)
	}
	return nil
}

func (t *InputGroupCall) MarshalJSON() ([]byte, error) {
	if t.InputGroupCallLink != nil {
		return json.Marshal(t.InputGroupCallLink)
	}
	if t.InputGroupCallMessage != nil {
		return json.Marshal(t.InputGroupCallMessage)
	}
	return nil, fmt.Errorf("no value set for InputGroupCall")
}

// InputInlineQueryResult Represents a single result of an inline query; for bots only
type InputInlineQueryResult struct {
	TypeStr                         string                           `json:"@type"`
	InputInlineQueryResultAnimation *InputInlineQueryResultAnimation `json:"inputInlineQueryResultAnimation,omitempty"`
	InputInlineQueryResultArticle   *InputInlineQueryResultArticle   `json:"inputInlineQueryResultArticle,omitempty"`
	InputInlineQueryResultAudio     *InputInlineQueryResultAudio     `json:"inputInlineQueryResultAudio,omitempty"`
	InputInlineQueryResultContact   *InputInlineQueryResultContact   `json:"inputInlineQueryResultContact,omitempty"`
	InputInlineQueryResultDocument  *InputInlineQueryResultDocument  `json:"inputInlineQueryResultDocument,omitempty"`
	InputInlineQueryResultGame      *InputInlineQueryResultGame      `json:"inputInlineQueryResultGame,omitempty"`
	InputInlineQueryResultLocation  *InputInlineQueryResultLocation  `json:"inputInlineQueryResultLocation,omitempty"`
	InputInlineQueryResultPhoto     *InputInlineQueryResultPhoto     `json:"inputInlineQueryResultPhoto,omitempty"`
	InputInlineQueryResultSticker   *InputInlineQueryResultSticker   `json:"inputInlineQueryResultSticker,omitempty"`
	InputInlineQueryResultVenue     *InputInlineQueryResultVenue     `json:"inputInlineQueryResultVenue,omitempty"`
	InputInlineQueryResultVideo     *InputInlineQueryResultVideo     `json:"inputInlineQueryResultVideo,omitempty"`
	InputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote `json:"inputInlineQueryResultVoiceNote,omitempty"`
}

func (t *InputInlineQueryResult) Type() string {
	return t.TypeStr
}

func (t *InputInlineQueryResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputInlineQueryResult) GetExtra() string {
	return ""
}

func (t *InputInlineQueryResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputInlineQueryResultAnimation":
		t.InputInlineQueryResultAnimation = new(InputInlineQueryResultAnimation)
		return json.Unmarshal(b, t.InputInlineQueryResultAnimation)
	case "inputInlineQueryResultArticle":
		t.InputInlineQueryResultArticle = new(InputInlineQueryResultArticle)
		return json.Unmarshal(b, t.InputInlineQueryResultArticle)
	case "inputInlineQueryResultAudio":
		t.InputInlineQueryResultAudio = new(InputInlineQueryResultAudio)
		return json.Unmarshal(b, t.InputInlineQueryResultAudio)
	case "inputInlineQueryResultContact":
		t.InputInlineQueryResultContact = new(InputInlineQueryResultContact)
		return json.Unmarshal(b, t.InputInlineQueryResultContact)
	case "inputInlineQueryResultDocument":
		t.InputInlineQueryResultDocument = new(InputInlineQueryResultDocument)
		return json.Unmarshal(b, t.InputInlineQueryResultDocument)
	case "inputInlineQueryResultGame":
		t.InputInlineQueryResultGame = new(InputInlineQueryResultGame)
		return json.Unmarshal(b, t.InputInlineQueryResultGame)
	case "inputInlineQueryResultLocation":
		t.InputInlineQueryResultLocation = new(InputInlineQueryResultLocation)
		return json.Unmarshal(b, t.InputInlineQueryResultLocation)
	case "inputInlineQueryResultPhoto":
		t.InputInlineQueryResultPhoto = new(InputInlineQueryResultPhoto)
		return json.Unmarshal(b, t.InputInlineQueryResultPhoto)
	case "inputInlineQueryResultSticker":
		t.InputInlineQueryResultSticker = new(InputInlineQueryResultSticker)
		return json.Unmarshal(b, t.InputInlineQueryResultSticker)
	case "inputInlineQueryResultVenue":
		t.InputInlineQueryResultVenue = new(InputInlineQueryResultVenue)
		return json.Unmarshal(b, t.InputInlineQueryResultVenue)
	case "inputInlineQueryResultVideo":
		t.InputInlineQueryResultVideo = new(InputInlineQueryResultVideo)
		return json.Unmarshal(b, t.InputInlineQueryResultVideo)
	case "inputInlineQueryResultVoiceNote":
		t.InputInlineQueryResultVoiceNote = new(InputInlineQueryResultVoiceNote)
		return json.Unmarshal(b, t.InputInlineQueryResultVoiceNote)
	}
	return nil
}

func (t *InputInlineQueryResult) MarshalJSON() ([]byte, error) {
	if t.InputInlineQueryResultAnimation != nil {
		return json.Marshal(t.InputInlineQueryResultAnimation)
	}
	if t.InputInlineQueryResultArticle != nil {
		return json.Marshal(t.InputInlineQueryResultArticle)
	}
	if t.InputInlineQueryResultAudio != nil {
		return json.Marshal(t.InputInlineQueryResultAudio)
	}
	if t.InputInlineQueryResultContact != nil {
		return json.Marshal(t.InputInlineQueryResultContact)
	}
	if t.InputInlineQueryResultDocument != nil {
		return json.Marshal(t.InputInlineQueryResultDocument)
	}
	if t.InputInlineQueryResultGame != nil {
		return json.Marshal(t.InputInlineQueryResultGame)
	}
	if t.InputInlineQueryResultLocation != nil {
		return json.Marshal(t.InputInlineQueryResultLocation)
	}
	if t.InputInlineQueryResultPhoto != nil {
		return json.Marshal(t.InputInlineQueryResultPhoto)
	}
	if t.InputInlineQueryResultSticker != nil {
		return json.Marshal(t.InputInlineQueryResultSticker)
	}
	if t.InputInlineQueryResultVenue != nil {
		return json.Marshal(t.InputInlineQueryResultVenue)
	}
	if t.InputInlineQueryResultVideo != nil {
		return json.Marshal(t.InputInlineQueryResultVideo)
	}
	if t.InputInlineQueryResultVoiceNote != nil {
		return json.Marshal(t.InputInlineQueryResultVoiceNote)
	}
	return nil, fmt.Errorf("no value set for InputInlineQueryResult")
}

// InputInvoice Describes an invoice to process
type InputInvoice struct {
	TypeStr              string                `json:"@type"`
	InputInvoiceMessage  *InputInvoiceMessage  `json:"inputInvoiceMessage,omitempty"`
	InputInvoiceName     *InputInvoiceName     `json:"inputInvoiceName,omitempty"`
	InputInvoiceTelegram *InputInvoiceTelegram `json:"inputInvoiceTelegram,omitempty"`
}

func (t *InputInvoice) Type() string {
	return t.TypeStr
}

func (t *InputInvoice) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputInvoice) GetExtra() string {
	return ""
}

func (t *InputInvoice) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputInvoiceMessage":
		t.InputInvoiceMessage = new(InputInvoiceMessage)
		return json.Unmarshal(b, t.InputInvoiceMessage)
	case "inputInvoiceName":
		t.InputInvoiceName = new(InputInvoiceName)
		return json.Unmarshal(b, t.InputInvoiceName)
	case "inputInvoiceTelegram":
		t.InputInvoiceTelegram = new(InputInvoiceTelegram)
		return json.Unmarshal(b, t.InputInvoiceTelegram)
	}
	return nil
}

func (t *InputInvoice) MarshalJSON() ([]byte, error) {
	if t.InputInvoiceMessage != nil {
		return json.Marshal(t.InputInvoiceMessage)
	}
	if t.InputInvoiceName != nil {
		return json.Marshal(t.InputInvoiceName)
	}
	if t.InputInvoiceTelegram != nil {
		return json.Marshal(t.InputInvoiceTelegram)
	}
	return nil, fmt.Errorf("no value set for InputInvoice")
}

// InputMessageContent The content of a message to send
type InputMessageContent struct {
	TypeStr               string                 `json:"@type"`
	InputMessageAnimation *InputMessageAnimation `json:"inputMessageAnimation,omitempty"`
	InputMessageAudio     *InputMessageAudio     `json:"inputMessageAudio,omitempty"`
	InputMessageChecklist *InputMessageChecklist `json:"inputMessageChecklist,omitempty"`
	InputMessageContact   *InputMessageContact   `json:"inputMessageContact,omitempty"`
	InputMessageDice      *InputMessageDice      `json:"inputMessageDice,omitempty"`
	InputMessageDocument  *InputMessageDocument  `json:"inputMessageDocument,omitempty"`
	InputMessageForwarded *InputMessageForwarded `json:"inputMessageForwarded,omitempty"`
	InputMessageGame      *InputMessageGame      `json:"inputMessageGame,omitempty"`
	InputMessageInvoice   *InputMessageInvoice   `json:"inputMessageInvoice,omitempty"`
	InputMessageLocation  *InputMessageLocation  `json:"inputMessageLocation,omitempty"`
	InputMessagePaidMedia *InputMessagePaidMedia `json:"inputMessagePaidMedia,omitempty"`
	InputMessagePhoto     *InputMessagePhoto     `json:"inputMessagePhoto,omitempty"`
	InputMessagePoll      *InputMessagePoll      `json:"inputMessagePoll,omitempty"`
	InputMessageStakeDice *InputMessageStakeDice `json:"inputMessageStakeDice,omitempty"`
	InputMessageSticker   *InputMessageSticker   `json:"inputMessageSticker,omitempty"`
	InputMessageStory     *InputMessageStory     `json:"inputMessageStory,omitempty"`
	InputMessageText      *InputMessageText      `json:"inputMessageText,omitempty"`
	InputMessageVenue     *InputMessageVenue     `json:"inputMessageVenue,omitempty"`
	InputMessageVideo     *InputMessageVideo     `json:"inputMessageVideo,omitempty"`
	InputMessageVideoNote *InputMessageVideoNote `json:"inputMessageVideoNote,omitempty"`
	InputMessageVoiceNote *InputMessageVoiceNote `json:"inputMessageVoiceNote,omitempty"`
}

func (t *InputMessageContent) Type() string {
	return t.TypeStr
}

func (t *InputMessageContent) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputMessageContent) GetExtra() string {
	return ""
}

func (t *InputMessageContent) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputMessageAnimation":
		t.InputMessageAnimation = new(InputMessageAnimation)
		return json.Unmarshal(b, t.InputMessageAnimation)
	case "inputMessageAudio":
		t.InputMessageAudio = new(InputMessageAudio)
		return json.Unmarshal(b, t.InputMessageAudio)
	case "inputMessageChecklist":
		t.InputMessageChecklist = new(InputMessageChecklist)
		return json.Unmarshal(b, t.InputMessageChecklist)
	case "inputMessageContact":
		t.InputMessageContact = new(InputMessageContact)
		return json.Unmarshal(b, t.InputMessageContact)
	case "inputMessageDice":
		t.InputMessageDice = new(InputMessageDice)
		return json.Unmarshal(b, t.InputMessageDice)
	case "inputMessageDocument":
		t.InputMessageDocument = new(InputMessageDocument)
		return json.Unmarshal(b, t.InputMessageDocument)
	case "inputMessageForwarded":
		t.InputMessageForwarded = new(InputMessageForwarded)
		return json.Unmarshal(b, t.InputMessageForwarded)
	case "inputMessageGame":
		t.InputMessageGame = new(InputMessageGame)
		return json.Unmarshal(b, t.InputMessageGame)
	case "inputMessageInvoice":
		t.InputMessageInvoice = new(InputMessageInvoice)
		return json.Unmarshal(b, t.InputMessageInvoice)
	case "inputMessageLocation":
		t.InputMessageLocation = new(InputMessageLocation)
		return json.Unmarshal(b, t.InputMessageLocation)
	case "inputMessagePaidMedia":
		t.InputMessagePaidMedia = new(InputMessagePaidMedia)
		return json.Unmarshal(b, t.InputMessagePaidMedia)
	case "inputMessagePhoto":
		t.InputMessagePhoto = new(InputMessagePhoto)
		return json.Unmarshal(b, t.InputMessagePhoto)
	case "inputMessagePoll":
		t.InputMessagePoll = new(InputMessagePoll)
		return json.Unmarshal(b, t.InputMessagePoll)
	case "inputMessageStakeDice":
		t.InputMessageStakeDice = new(InputMessageStakeDice)
		return json.Unmarshal(b, t.InputMessageStakeDice)
	case "inputMessageSticker":
		t.InputMessageSticker = new(InputMessageSticker)
		return json.Unmarshal(b, t.InputMessageSticker)
	case "inputMessageStory":
		t.InputMessageStory = new(InputMessageStory)
		return json.Unmarshal(b, t.InputMessageStory)
	case "inputMessageText":
		t.InputMessageText = new(InputMessageText)
		return json.Unmarshal(b, t.InputMessageText)
	case "inputMessageVenue":
		t.InputMessageVenue = new(InputMessageVenue)
		return json.Unmarshal(b, t.InputMessageVenue)
	case "inputMessageVideo":
		t.InputMessageVideo = new(InputMessageVideo)
		return json.Unmarshal(b, t.InputMessageVideo)
	case "inputMessageVideoNote":
		t.InputMessageVideoNote = new(InputMessageVideoNote)
		return json.Unmarshal(b, t.InputMessageVideoNote)
	case "inputMessageVoiceNote":
		t.InputMessageVoiceNote = new(InputMessageVoiceNote)
		return json.Unmarshal(b, t.InputMessageVoiceNote)
	}
	return nil
}

func (t *InputMessageContent) MarshalJSON() ([]byte, error) {
	if t.InputMessageAnimation != nil {
		return json.Marshal(t.InputMessageAnimation)
	}
	if t.InputMessageAudio != nil {
		return json.Marshal(t.InputMessageAudio)
	}
	if t.InputMessageChecklist != nil {
		return json.Marshal(t.InputMessageChecklist)
	}
	if t.InputMessageContact != nil {
		return json.Marshal(t.InputMessageContact)
	}
	if t.InputMessageDice != nil {
		return json.Marshal(t.InputMessageDice)
	}
	if t.InputMessageDocument != nil {
		return json.Marshal(t.InputMessageDocument)
	}
	if t.InputMessageForwarded != nil {
		return json.Marshal(t.InputMessageForwarded)
	}
	if t.InputMessageGame != nil {
		return json.Marshal(t.InputMessageGame)
	}
	if t.InputMessageInvoice != nil {
		return json.Marshal(t.InputMessageInvoice)
	}
	if t.InputMessageLocation != nil {
		return json.Marshal(t.InputMessageLocation)
	}
	if t.InputMessagePaidMedia != nil {
		return json.Marshal(t.InputMessagePaidMedia)
	}
	if t.InputMessagePhoto != nil {
		return json.Marshal(t.InputMessagePhoto)
	}
	if t.InputMessagePoll != nil {
		return json.Marshal(t.InputMessagePoll)
	}
	if t.InputMessageStakeDice != nil {
		return json.Marshal(t.InputMessageStakeDice)
	}
	if t.InputMessageSticker != nil {
		return json.Marshal(t.InputMessageSticker)
	}
	if t.InputMessageStory != nil {
		return json.Marshal(t.InputMessageStory)
	}
	if t.InputMessageText != nil {
		return json.Marshal(t.InputMessageText)
	}
	if t.InputMessageVenue != nil {
		return json.Marshal(t.InputMessageVenue)
	}
	if t.InputMessageVideo != nil {
		return json.Marshal(t.InputMessageVideo)
	}
	if t.InputMessageVideoNote != nil {
		return json.Marshal(t.InputMessageVideoNote)
	}
	if t.InputMessageVoiceNote != nil {
		return json.Marshal(t.InputMessageVoiceNote)
	}
	return nil, fmt.Errorf("no value set for InputMessageContent")
}

// InputMessageReplyTo Contains information about the message or the story to be replied
type InputMessageReplyTo struct {
	TypeStr                            string                              `json:"@type"`
	InputMessageReplyToExternalMessage *InputMessageReplyToExternalMessage `json:"inputMessageReplyToExternalMessage,omitempty"`
	InputMessageReplyToMessage         *InputMessageReplyToMessage         `json:"inputMessageReplyToMessage,omitempty"`
	InputMessageReplyToStory           *InputMessageReplyToStory           `json:"inputMessageReplyToStory,omitempty"`
}

func (t *InputMessageReplyTo) Type() string {
	return t.TypeStr
}

func (t *InputMessageReplyTo) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputMessageReplyTo) GetExtra() string {
	return ""
}

func (t *InputMessageReplyTo) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputMessageReplyToExternalMessage":
		t.InputMessageReplyToExternalMessage = new(InputMessageReplyToExternalMessage)
		return json.Unmarshal(b, t.InputMessageReplyToExternalMessage)
	case "inputMessageReplyToMessage":
		t.InputMessageReplyToMessage = new(InputMessageReplyToMessage)
		return json.Unmarshal(b, t.InputMessageReplyToMessage)
	case "inputMessageReplyToStory":
		t.InputMessageReplyToStory = new(InputMessageReplyToStory)
		return json.Unmarshal(b, t.InputMessageReplyToStory)
	}
	return nil
}

func (t *InputMessageReplyTo) MarshalJSON() ([]byte, error) {
	if t.InputMessageReplyToExternalMessage != nil {
		return json.Marshal(t.InputMessageReplyToExternalMessage)
	}
	if t.InputMessageReplyToMessage != nil {
		return json.Marshal(t.InputMessageReplyToMessage)
	}
	if t.InputMessageReplyToStory != nil {
		return json.Marshal(t.InputMessageReplyToStory)
	}
	return nil, fmt.Errorf("no value set for InputMessageReplyTo")
}

// InputPaidMediaType Describes type of paid media to sent
type InputPaidMediaType struct {
	TypeStr                 string                   `json:"@type"`
	InputPaidMediaTypePhoto *InputPaidMediaTypePhoto `json:"inputPaidMediaTypePhoto,omitempty"`
	InputPaidMediaTypeVideo *InputPaidMediaTypeVideo `json:"inputPaidMediaTypeVideo,omitempty"`
}

func (t *InputPaidMediaType) Type() string {
	return t.TypeStr
}

func (t *InputPaidMediaType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputPaidMediaType) GetExtra() string {
	return ""
}

func (t *InputPaidMediaType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputPaidMediaTypePhoto":
		t.InputPaidMediaTypePhoto = new(InputPaidMediaTypePhoto)
		return json.Unmarshal(b, t.InputPaidMediaTypePhoto)
	case "inputPaidMediaTypeVideo":
		t.InputPaidMediaTypeVideo = new(InputPaidMediaTypeVideo)
		return json.Unmarshal(b, t.InputPaidMediaTypeVideo)
	}
	return nil
}

func (t *InputPaidMediaType) MarshalJSON() ([]byte, error) {
	if t.InputPaidMediaTypePhoto != nil {
		return json.Marshal(t.InputPaidMediaTypePhoto)
	}
	if t.InputPaidMediaTypeVideo != nil {
		return json.Marshal(t.InputPaidMediaTypeVideo)
	}
	return nil, fmt.Errorf("no value set for InputPaidMediaType")
}

// InputPassportElement Contains information about a Telegram Passport element to be saved
type InputPassportElement struct {
	TypeStr                                   string                                     `json:"@type"`
	InputPassportElementAddress               *InputPassportElementAddress               `json:"inputPassportElementAddress,omitempty"`
	InputPassportElementBankStatement         *InputPassportElementBankStatement         `json:"inputPassportElementBankStatement,omitempty"`
	InputPassportElementDriverLicense         *InputPassportElementDriverLicense         `json:"inputPassportElementDriverLicense,omitempty"`
	InputPassportElementEmailAddress          *InputPassportElementEmailAddress          `json:"inputPassportElementEmailAddress,omitempty"`
	InputPassportElementIdentityCard          *InputPassportElementIdentityCard          `json:"inputPassportElementIdentityCard,omitempty"`
	InputPassportElementInternalPassport      *InputPassportElementInternalPassport      `json:"inputPassportElementInternalPassport,omitempty"`
	InputPassportElementPassport              *InputPassportElementPassport              `json:"inputPassportElementPassport,omitempty"`
	InputPassportElementPassportRegistration  *InputPassportElementPassportRegistration  `json:"inputPassportElementPassportRegistration,omitempty"`
	InputPassportElementPersonalDetails       *InputPassportElementPersonalDetails       `json:"inputPassportElementPersonalDetails,omitempty"`
	InputPassportElementPhoneNumber           *InputPassportElementPhoneNumber           `json:"inputPassportElementPhoneNumber,omitempty"`
	InputPassportElementRentalAgreement       *InputPassportElementRentalAgreement       `json:"inputPassportElementRentalAgreement,omitempty"`
	InputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration `json:"inputPassportElementTemporaryRegistration,omitempty"`
	InputPassportElementUtilityBill           *InputPassportElementUtilityBill           `json:"inputPassportElementUtilityBill,omitempty"`
}

func (t *InputPassportElement) Type() string {
	return t.TypeStr
}

func (t *InputPassportElement) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputPassportElement) GetExtra() string {
	return ""
}

func (t *InputPassportElement) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputPassportElementAddress":
		t.InputPassportElementAddress = new(InputPassportElementAddress)
		return json.Unmarshal(b, t.InputPassportElementAddress)
	case "inputPassportElementBankStatement":
		t.InputPassportElementBankStatement = new(InputPassportElementBankStatement)
		return json.Unmarshal(b, t.InputPassportElementBankStatement)
	case "inputPassportElementDriverLicense":
		t.InputPassportElementDriverLicense = new(InputPassportElementDriverLicense)
		return json.Unmarshal(b, t.InputPassportElementDriverLicense)
	case "inputPassportElementEmailAddress":
		t.InputPassportElementEmailAddress = new(InputPassportElementEmailAddress)
		return json.Unmarshal(b, t.InputPassportElementEmailAddress)
	case "inputPassportElementIdentityCard":
		t.InputPassportElementIdentityCard = new(InputPassportElementIdentityCard)
		return json.Unmarshal(b, t.InputPassportElementIdentityCard)
	case "inputPassportElementInternalPassport":
		t.InputPassportElementInternalPassport = new(InputPassportElementInternalPassport)
		return json.Unmarshal(b, t.InputPassportElementInternalPassport)
	case "inputPassportElementPassport":
		t.InputPassportElementPassport = new(InputPassportElementPassport)
		return json.Unmarshal(b, t.InputPassportElementPassport)
	case "inputPassportElementPassportRegistration":
		t.InputPassportElementPassportRegistration = new(InputPassportElementPassportRegistration)
		return json.Unmarshal(b, t.InputPassportElementPassportRegistration)
	case "inputPassportElementPersonalDetails":
		t.InputPassportElementPersonalDetails = new(InputPassportElementPersonalDetails)
		return json.Unmarshal(b, t.InputPassportElementPersonalDetails)
	case "inputPassportElementPhoneNumber":
		t.InputPassportElementPhoneNumber = new(InputPassportElementPhoneNumber)
		return json.Unmarshal(b, t.InputPassportElementPhoneNumber)
	case "inputPassportElementRentalAgreement":
		t.InputPassportElementRentalAgreement = new(InputPassportElementRentalAgreement)
		return json.Unmarshal(b, t.InputPassportElementRentalAgreement)
	case "inputPassportElementTemporaryRegistration":
		t.InputPassportElementTemporaryRegistration = new(InputPassportElementTemporaryRegistration)
		return json.Unmarshal(b, t.InputPassportElementTemporaryRegistration)
	case "inputPassportElementUtilityBill":
		t.InputPassportElementUtilityBill = new(InputPassportElementUtilityBill)
		return json.Unmarshal(b, t.InputPassportElementUtilityBill)
	}
	return nil
}

func (t *InputPassportElement) MarshalJSON() ([]byte, error) {
	if t.InputPassportElementAddress != nil {
		return json.Marshal(t.InputPassportElementAddress)
	}
	if t.InputPassportElementBankStatement != nil {
		return json.Marshal(t.InputPassportElementBankStatement)
	}
	if t.InputPassportElementDriverLicense != nil {
		return json.Marshal(t.InputPassportElementDriverLicense)
	}
	if t.InputPassportElementEmailAddress != nil {
		return json.Marshal(t.InputPassportElementEmailAddress)
	}
	if t.InputPassportElementIdentityCard != nil {
		return json.Marshal(t.InputPassportElementIdentityCard)
	}
	if t.InputPassportElementInternalPassport != nil {
		return json.Marshal(t.InputPassportElementInternalPassport)
	}
	if t.InputPassportElementPassport != nil {
		return json.Marshal(t.InputPassportElementPassport)
	}
	if t.InputPassportElementPassportRegistration != nil {
		return json.Marshal(t.InputPassportElementPassportRegistration)
	}
	if t.InputPassportElementPersonalDetails != nil {
		return json.Marshal(t.InputPassportElementPersonalDetails)
	}
	if t.InputPassportElementPhoneNumber != nil {
		return json.Marshal(t.InputPassportElementPhoneNumber)
	}
	if t.InputPassportElementRentalAgreement != nil {
		return json.Marshal(t.InputPassportElementRentalAgreement)
	}
	if t.InputPassportElementTemporaryRegistration != nil {
		return json.Marshal(t.InputPassportElementTemporaryRegistration)
	}
	if t.InputPassportElementUtilityBill != nil {
		return json.Marshal(t.InputPassportElementUtilityBill)
	}
	return nil, fmt.Errorf("no value set for InputPassportElement")
}

// InputPassportElementErrorSource Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementErrorSource struct {
	TypeStr                                         string                                           `json:"@type"`
	InputPassportElementErrorSourceDataField        *InputPassportElementErrorSourceDataField        `json:"inputPassportElementErrorSourceDataField,omitempty"`
	InputPassportElementErrorSourceFile             *InputPassportElementErrorSourceFile             `json:"inputPassportElementErrorSourceFile,omitempty"`
	InputPassportElementErrorSourceFiles            *InputPassportElementErrorSourceFiles            `json:"inputPassportElementErrorSourceFiles,omitempty"`
	InputPassportElementErrorSourceFrontSide        *InputPassportElementErrorSourceFrontSide        `json:"inputPassportElementErrorSourceFrontSide,omitempty"`
	InputPassportElementErrorSourceReverseSide      *InputPassportElementErrorSourceReverseSide      `json:"inputPassportElementErrorSourceReverseSide,omitempty"`
	InputPassportElementErrorSourceSelfie           *InputPassportElementErrorSourceSelfie           `json:"inputPassportElementErrorSourceSelfie,omitempty"`
	InputPassportElementErrorSourceTranslationFile  *InputPassportElementErrorSourceTranslationFile  `json:"inputPassportElementErrorSourceTranslationFile,omitempty"`
	InputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles `json:"inputPassportElementErrorSourceTranslationFiles,omitempty"`
	InputPassportElementErrorSourceUnspecified      *InputPassportElementErrorSourceUnspecified      `json:"inputPassportElementErrorSourceUnspecified,omitempty"`
}

func (t *InputPassportElementErrorSource) Type() string {
	return t.TypeStr
}

func (t *InputPassportElementErrorSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputPassportElementErrorSource) GetExtra() string {
	return ""
}

func (t *InputPassportElementErrorSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputPassportElementErrorSourceDataField":
		t.InputPassportElementErrorSourceDataField = new(InputPassportElementErrorSourceDataField)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceDataField)
	case "inputPassportElementErrorSourceFile":
		t.InputPassportElementErrorSourceFile = new(InputPassportElementErrorSourceFile)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceFile)
	case "inputPassportElementErrorSourceFiles":
		t.InputPassportElementErrorSourceFiles = new(InputPassportElementErrorSourceFiles)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceFiles)
	case "inputPassportElementErrorSourceFrontSide":
		t.InputPassportElementErrorSourceFrontSide = new(InputPassportElementErrorSourceFrontSide)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceFrontSide)
	case "inputPassportElementErrorSourceReverseSide":
		t.InputPassportElementErrorSourceReverseSide = new(InputPassportElementErrorSourceReverseSide)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceReverseSide)
	case "inputPassportElementErrorSourceSelfie":
		t.InputPassportElementErrorSourceSelfie = new(InputPassportElementErrorSourceSelfie)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceSelfie)
	case "inputPassportElementErrorSourceTranslationFile":
		t.InputPassportElementErrorSourceTranslationFile = new(InputPassportElementErrorSourceTranslationFile)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceTranslationFile)
	case "inputPassportElementErrorSourceTranslationFiles":
		t.InputPassportElementErrorSourceTranslationFiles = new(InputPassportElementErrorSourceTranslationFiles)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceTranslationFiles)
	case "inputPassportElementErrorSourceUnspecified":
		t.InputPassportElementErrorSourceUnspecified = new(InputPassportElementErrorSourceUnspecified)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceUnspecified)
	}
	return nil
}

func (t *InputPassportElementErrorSource) MarshalJSON() ([]byte, error) {
	if t.InputPassportElementErrorSourceDataField != nil {
		return json.Marshal(t.InputPassportElementErrorSourceDataField)
	}
	if t.InputPassportElementErrorSourceFile != nil {
		return json.Marshal(t.InputPassportElementErrorSourceFile)
	}
	if t.InputPassportElementErrorSourceFiles != nil {
		return json.Marshal(t.InputPassportElementErrorSourceFiles)
	}
	if t.InputPassportElementErrorSourceFrontSide != nil {
		return json.Marshal(t.InputPassportElementErrorSourceFrontSide)
	}
	if t.InputPassportElementErrorSourceReverseSide != nil {
		return json.Marshal(t.InputPassportElementErrorSourceReverseSide)
	}
	if t.InputPassportElementErrorSourceSelfie != nil {
		return json.Marshal(t.InputPassportElementErrorSourceSelfie)
	}
	if t.InputPassportElementErrorSourceTranslationFile != nil {
		return json.Marshal(t.InputPassportElementErrorSourceTranslationFile)
	}
	if t.InputPassportElementErrorSourceTranslationFiles != nil {
		return json.Marshal(t.InputPassportElementErrorSourceTranslationFiles)
	}
	if t.InputPassportElementErrorSourceUnspecified != nil {
		return json.Marshal(t.InputPassportElementErrorSourceUnspecified)
	}
	return nil, fmt.Errorf("no value set for InputPassportElementErrorSource")
}

// InputStoryAreaType Describes type of clickable area on a story media to be added
type InputStoryAreaType struct {
	TypeStr                             string                               `json:"@type"`
	InputStoryAreaTypeFoundVenue        *InputStoryAreaTypeFoundVenue        `json:"inputStoryAreaTypeFoundVenue,omitempty"`
	InputStoryAreaTypeLink              *InputStoryAreaTypeLink              `json:"inputStoryAreaTypeLink,omitempty"`
	InputStoryAreaTypeLocation          *InputStoryAreaTypeLocation          `json:"inputStoryAreaTypeLocation,omitempty"`
	InputStoryAreaTypeMessage           *InputStoryAreaTypeMessage           `json:"inputStoryAreaTypeMessage,omitempty"`
	InputStoryAreaTypePreviousVenue     *InputStoryAreaTypePreviousVenue     `json:"inputStoryAreaTypePreviousVenue,omitempty"`
	InputStoryAreaTypeSuggestedReaction *InputStoryAreaTypeSuggestedReaction `json:"inputStoryAreaTypeSuggestedReaction,omitempty"`
	InputStoryAreaTypeUpgradedGift      *InputStoryAreaTypeUpgradedGift      `json:"inputStoryAreaTypeUpgradedGift,omitempty"`
	InputStoryAreaTypeWeather           *InputStoryAreaTypeWeather           `json:"inputStoryAreaTypeWeather,omitempty"`
}

func (t *InputStoryAreaType) Type() string {
	return t.TypeStr
}

func (t *InputStoryAreaType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputStoryAreaType) GetExtra() string {
	return ""
}

func (t *InputStoryAreaType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputStoryAreaTypeFoundVenue":
		t.InputStoryAreaTypeFoundVenue = new(InputStoryAreaTypeFoundVenue)
		return json.Unmarshal(b, t.InputStoryAreaTypeFoundVenue)
	case "inputStoryAreaTypeLink":
		t.InputStoryAreaTypeLink = new(InputStoryAreaTypeLink)
		return json.Unmarshal(b, t.InputStoryAreaTypeLink)
	case "inputStoryAreaTypeLocation":
		t.InputStoryAreaTypeLocation = new(InputStoryAreaTypeLocation)
		return json.Unmarshal(b, t.InputStoryAreaTypeLocation)
	case "inputStoryAreaTypeMessage":
		t.InputStoryAreaTypeMessage = new(InputStoryAreaTypeMessage)
		return json.Unmarshal(b, t.InputStoryAreaTypeMessage)
	case "inputStoryAreaTypePreviousVenue":
		t.InputStoryAreaTypePreviousVenue = new(InputStoryAreaTypePreviousVenue)
		return json.Unmarshal(b, t.InputStoryAreaTypePreviousVenue)
	case "inputStoryAreaTypeSuggestedReaction":
		t.InputStoryAreaTypeSuggestedReaction = new(InputStoryAreaTypeSuggestedReaction)
		return json.Unmarshal(b, t.InputStoryAreaTypeSuggestedReaction)
	case "inputStoryAreaTypeUpgradedGift":
		t.InputStoryAreaTypeUpgradedGift = new(InputStoryAreaTypeUpgradedGift)
		return json.Unmarshal(b, t.InputStoryAreaTypeUpgradedGift)
	case "inputStoryAreaTypeWeather":
		t.InputStoryAreaTypeWeather = new(InputStoryAreaTypeWeather)
		return json.Unmarshal(b, t.InputStoryAreaTypeWeather)
	}
	return nil
}

func (t *InputStoryAreaType) MarshalJSON() ([]byte, error) {
	if t.InputStoryAreaTypeFoundVenue != nil {
		return json.Marshal(t.InputStoryAreaTypeFoundVenue)
	}
	if t.InputStoryAreaTypeLink != nil {
		return json.Marshal(t.InputStoryAreaTypeLink)
	}
	if t.InputStoryAreaTypeLocation != nil {
		return json.Marshal(t.InputStoryAreaTypeLocation)
	}
	if t.InputStoryAreaTypeMessage != nil {
		return json.Marshal(t.InputStoryAreaTypeMessage)
	}
	if t.InputStoryAreaTypePreviousVenue != nil {
		return json.Marshal(t.InputStoryAreaTypePreviousVenue)
	}
	if t.InputStoryAreaTypeSuggestedReaction != nil {
		return json.Marshal(t.InputStoryAreaTypeSuggestedReaction)
	}
	if t.InputStoryAreaTypeUpgradedGift != nil {
		return json.Marshal(t.InputStoryAreaTypeUpgradedGift)
	}
	if t.InputStoryAreaTypeWeather != nil {
		return json.Marshal(t.InputStoryAreaTypeWeather)
	}
	return nil, fmt.Errorf("no value set for InputStoryAreaType")
}

// InputStoryContent The content of a story to post
type InputStoryContent struct {
	TypeStr                string                  `json:"@type"`
	InputStoryContentPhoto *InputStoryContentPhoto `json:"inputStoryContentPhoto,omitempty"`
	InputStoryContentVideo *InputStoryContentVideo `json:"inputStoryContentVideo,omitempty"`
}

func (t *InputStoryContent) Type() string {
	return t.TypeStr
}

func (t *InputStoryContent) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InputStoryContent) GetExtra() string {
	return ""
}

func (t *InputStoryContent) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inputStoryContentPhoto":
		t.InputStoryContentPhoto = new(InputStoryContentPhoto)
		return json.Unmarshal(b, t.InputStoryContentPhoto)
	case "inputStoryContentVideo":
		t.InputStoryContentVideo = new(InputStoryContentVideo)
		return json.Unmarshal(b, t.InputStoryContentVideo)
	}
	return nil
}

func (t *InputStoryContent) MarshalJSON() ([]byte, error) {
	if t.InputStoryContentPhoto != nil {
		return json.Marshal(t.InputStoryContentPhoto)
	}
	if t.InputStoryContentVideo != nil {
		return json.Marshal(t.InputStoryContentVideo)
	}
	return nil, fmt.Errorf("no value set for InputStoryContent")
}

// InternalLinkType Describes an internal https://t.me or tg: link, which must be processed by the application in a special way
type InternalLinkType struct {
	TypeStr                                               string                                                 `json:"@type"`
	InternalLinkTypeActiveSessions                        *InternalLinkTypeActiveSessions                        `json:"internalLinkTypeActiveSessions,omitempty"`
	InternalLinkTypeAttachmentMenuBot                     *InternalLinkTypeAttachmentMenuBot                     `json:"internalLinkTypeAttachmentMenuBot,omitempty"`
	InternalLinkTypeAuthenticationCode                    *InternalLinkTypeAuthenticationCode                    `json:"internalLinkTypeAuthenticationCode,omitempty"`
	InternalLinkTypeBackground                            *InternalLinkTypeBackground                            `json:"internalLinkTypeBackground,omitempty"`
	InternalLinkTypeBotAddToChannel                       *InternalLinkTypeBotAddToChannel                       `json:"internalLinkTypeBotAddToChannel,omitempty"`
	InternalLinkTypeBotStart                              *InternalLinkTypeBotStart                              `json:"internalLinkTypeBotStart,omitempty"`
	InternalLinkTypeBotStartInGroup                       *InternalLinkTypeBotStartInGroup                       `json:"internalLinkTypeBotStartInGroup,omitempty"`
	InternalLinkTypeBusinessChat                          *InternalLinkTypeBusinessChat                          `json:"internalLinkTypeBusinessChat,omitempty"`
	InternalLinkTypeBuyStars                              *InternalLinkTypeBuyStars                              `json:"internalLinkTypeBuyStars,omitempty"`
	InternalLinkTypeChangePhoneNumber                     *InternalLinkTypeChangePhoneNumber                     `json:"internalLinkTypeChangePhoneNumber,omitempty"`
	InternalLinkTypeChatAffiliateProgram                  *InternalLinkTypeChatAffiliateProgram                  `json:"internalLinkTypeChatAffiliateProgram,omitempty"`
	InternalLinkTypeChatBoost                             *InternalLinkTypeChatBoost                             `json:"internalLinkTypeChatBoost,omitempty"`
	InternalLinkTypeChatFolderInvite                      *InternalLinkTypeChatFolderInvite                      `json:"internalLinkTypeChatFolderInvite,omitempty"`
	InternalLinkTypeChatFolderSettings                    *InternalLinkTypeChatFolderSettings                    `json:"internalLinkTypeChatFolderSettings,omitempty"`
	InternalLinkTypeChatInvite                            *InternalLinkTypeChatInvite                            `json:"internalLinkTypeChatInvite,omitempty"`
	InternalLinkTypeDefaultMessageAutoDeleteTimerSettings *InternalLinkTypeDefaultMessageAutoDeleteTimerSettings `json:"internalLinkTypeDefaultMessageAutoDeleteTimerSettings,omitempty"`
	InternalLinkTypeDirectMessagesChat                    *InternalLinkTypeDirectMessagesChat                    `json:"internalLinkTypeDirectMessagesChat,omitempty"`
	InternalLinkTypeEditProfileSettings                   *InternalLinkTypeEditProfileSettings                   `json:"internalLinkTypeEditProfileSettings,omitempty"`
	InternalLinkTypeGame                                  *InternalLinkTypeGame                                  `json:"internalLinkTypeGame,omitempty"`
	InternalLinkTypeGiftAuction                           *InternalLinkTypeGiftAuction                           `json:"internalLinkTypeGiftAuction,omitempty"`
	InternalLinkTypeGiftCollection                        *InternalLinkTypeGiftCollection                        `json:"internalLinkTypeGiftCollection,omitempty"`
	InternalLinkTypeGroupCall                             *InternalLinkTypeGroupCall                             `json:"internalLinkTypeGroupCall,omitempty"`
	InternalLinkTypeInstantView                           *InternalLinkTypeInstantView                           `json:"internalLinkTypeInstantView,omitempty"`
	InternalLinkTypeInvoice                               *InternalLinkTypeInvoice                               `json:"internalLinkTypeInvoice,omitempty"`
	InternalLinkTypeLanguagePack                          *InternalLinkTypeLanguagePack                          `json:"internalLinkTypeLanguagePack,omitempty"`
	InternalLinkTypeLanguageSettings                      *InternalLinkTypeLanguageSettings                      `json:"internalLinkTypeLanguageSettings,omitempty"`
	InternalLinkTypeLiveStory                             *InternalLinkTypeLiveStory                             `json:"internalLinkTypeLiveStory,omitempty"`
	InternalLinkTypeLoginEmailSettings                    *InternalLinkTypeLoginEmailSettings                    `json:"internalLinkTypeLoginEmailSettings,omitempty"`
	InternalLinkTypeMainWebApp                            *InternalLinkTypeMainWebApp                            `json:"internalLinkTypeMainWebApp,omitempty"`
	InternalLinkTypeMessage                               *InternalLinkTypeMessage                               `json:"internalLinkTypeMessage,omitempty"`
	InternalLinkTypeMessageDraft                          *InternalLinkTypeMessageDraft                          `json:"internalLinkTypeMessageDraft,omitempty"`
	InternalLinkTypeMyStars                               *InternalLinkTypeMyStars                               `json:"internalLinkTypeMyStars,omitempty"`
	InternalLinkTypeMyToncoins                            *InternalLinkTypeMyToncoins                            `json:"internalLinkTypeMyToncoins,omitempty"`
	InternalLinkTypePassportDataRequest                   *InternalLinkTypePassportDataRequest                   `json:"internalLinkTypePassportDataRequest,omitempty"`
	InternalLinkTypePasswordSettings                      *InternalLinkTypePasswordSettings                      `json:"internalLinkTypePasswordSettings,omitempty"`
	InternalLinkTypePhoneNumberConfirmation               *InternalLinkTypePhoneNumberConfirmation               `json:"internalLinkTypePhoneNumberConfirmation,omitempty"`
	InternalLinkTypePhoneNumberPrivacySettings            *InternalLinkTypePhoneNumberPrivacySettings            `json:"internalLinkTypePhoneNumberPrivacySettings,omitempty"`
	InternalLinkTypePremiumFeatures                       *InternalLinkTypePremiumFeatures                       `json:"internalLinkTypePremiumFeatures,omitempty"`
	InternalLinkTypePremiumGift                           *InternalLinkTypePremiumGift                           `json:"internalLinkTypePremiumGift,omitempty"`
	InternalLinkTypePremiumGiftCode                       *InternalLinkTypePremiumGiftCode                       `json:"internalLinkTypePremiumGiftCode,omitempty"`
	InternalLinkTypePrivacyAndSecuritySettings            *InternalLinkTypePrivacyAndSecuritySettings            `json:"internalLinkTypePrivacyAndSecuritySettings,omitempty"`
	InternalLinkTypeProxy                                 *InternalLinkTypeProxy                                 `json:"internalLinkTypeProxy,omitempty"`
	InternalLinkTypePublicChat                            *InternalLinkTypePublicChat                            `json:"internalLinkTypePublicChat,omitempty"`
	InternalLinkTypeQrCodeAuthentication                  *InternalLinkTypeQrCodeAuthentication                  `json:"internalLinkTypeQrCodeAuthentication,omitempty"`
	InternalLinkTypeRestorePurchases                      *InternalLinkTypeRestorePurchases                      `json:"internalLinkTypeRestorePurchases,omitempty"`
	InternalLinkTypeSettings                              *InternalLinkTypeSettings                              `json:"internalLinkTypeSettings,omitempty"`
	InternalLinkTypeStickerSet                            *InternalLinkTypeStickerSet                            `json:"internalLinkTypeStickerSet,omitempty"`
	InternalLinkTypeStory                                 *InternalLinkTypeStory                                 `json:"internalLinkTypeStory,omitempty"`
	InternalLinkTypeStoryAlbum                            *InternalLinkTypeStoryAlbum                            `json:"internalLinkTypeStoryAlbum,omitempty"`
	InternalLinkTypeTheme                                 *InternalLinkTypeTheme                                 `json:"internalLinkTypeTheme,omitempty"`
	InternalLinkTypeThemeSettings                         *InternalLinkTypeThemeSettings                         `json:"internalLinkTypeThemeSettings,omitempty"`
	InternalLinkTypeUnknownDeepLink                       *InternalLinkTypeUnknownDeepLink                       `json:"internalLinkTypeUnknownDeepLink,omitempty"`
	InternalLinkTypeUnsupportedProxy                      *InternalLinkTypeUnsupportedProxy                      `json:"internalLinkTypeUnsupportedProxy,omitempty"`
	InternalLinkTypeUpgradedGift                          *InternalLinkTypeUpgradedGift                          `json:"internalLinkTypeUpgradedGift,omitempty"`
	InternalLinkTypeUserPhoneNumber                       *InternalLinkTypeUserPhoneNumber                       `json:"internalLinkTypeUserPhoneNumber,omitempty"`
	InternalLinkTypeUserToken                             *InternalLinkTypeUserToken                             `json:"internalLinkTypeUserToken,omitempty"`
	InternalLinkTypeVideoChat                             *InternalLinkTypeVideoChat                             `json:"internalLinkTypeVideoChat,omitempty"`
	InternalLinkTypeWebApp                                *InternalLinkTypeWebApp                                `json:"internalLinkTypeWebApp,omitempty"`
}

func (t *InternalLinkType) Type() string {
	return t.TypeStr
}

func (t *InternalLinkType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InternalLinkType) GetExtra() string {
	return ""
}

func (t *InternalLinkType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "internalLinkTypeActiveSessions":
		t.InternalLinkTypeActiveSessions = new(InternalLinkTypeActiveSessions)
		return json.Unmarshal(b, t.InternalLinkTypeActiveSessions)
	case "internalLinkTypeAttachmentMenuBot":
		t.InternalLinkTypeAttachmentMenuBot = new(InternalLinkTypeAttachmentMenuBot)
		return json.Unmarshal(b, t.InternalLinkTypeAttachmentMenuBot)
	case "internalLinkTypeAuthenticationCode":
		t.InternalLinkTypeAuthenticationCode = new(InternalLinkTypeAuthenticationCode)
		return json.Unmarshal(b, t.InternalLinkTypeAuthenticationCode)
	case "internalLinkTypeBackground":
		t.InternalLinkTypeBackground = new(InternalLinkTypeBackground)
		return json.Unmarshal(b, t.InternalLinkTypeBackground)
	case "internalLinkTypeBotAddToChannel":
		t.InternalLinkTypeBotAddToChannel = new(InternalLinkTypeBotAddToChannel)
		return json.Unmarshal(b, t.InternalLinkTypeBotAddToChannel)
	case "internalLinkTypeBotStart":
		t.InternalLinkTypeBotStart = new(InternalLinkTypeBotStart)
		return json.Unmarshal(b, t.InternalLinkTypeBotStart)
	case "internalLinkTypeBotStartInGroup":
		t.InternalLinkTypeBotStartInGroup = new(InternalLinkTypeBotStartInGroup)
		return json.Unmarshal(b, t.InternalLinkTypeBotStartInGroup)
	case "internalLinkTypeBusinessChat":
		t.InternalLinkTypeBusinessChat = new(InternalLinkTypeBusinessChat)
		return json.Unmarshal(b, t.InternalLinkTypeBusinessChat)
	case "internalLinkTypeBuyStars":
		t.InternalLinkTypeBuyStars = new(InternalLinkTypeBuyStars)
		return json.Unmarshal(b, t.InternalLinkTypeBuyStars)
	case "internalLinkTypeChangePhoneNumber":
		t.InternalLinkTypeChangePhoneNumber = new(InternalLinkTypeChangePhoneNumber)
		return json.Unmarshal(b, t.InternalLinkTypeChangePhoneNumber)
	case "internalLinkTypeChatAffiliateProgram":
		t.InternalLinkTypeChatAffiliateProgram = new(InternalLinkTypeChatAffiliateProgram)
		return json.Unmarshal(b, t.InternalLinkTypeChatAffiliateProgram)
	case "internalLinkTypeChatBoost":
		t.InternalLinkTypeChatBoost = new(InternalLinkTypeChatBoost)
		return json.Unmarshal(b, t.InternalLinkTypeChatBoost)
	case "internalLinkTypeChatFolderInvite":
		t.InternalLinkTypeChatFolderInvite = new(InternalLinkTypeChatFolderInvite)
		return json.Unmarshal(b, t.InternalLinkTypeChatFolderInvite)
	case "internalLinkTypeChatFolderSettings":
		t.InternalLinkTypeChatFolderSettings = new(InternalLinkTypeChatFolderSettings)
		return json.Unmarshal(b, t.InternalLinkTypeChatFolderSettings)
	case "internalLinkTypeChatInvite":
		t.InternalLinkTypeChatInvite = new(InternalLinkTypeChatInvite)
		return json.Unmarshal(b, t.InternalLinkTypeChatInvite)
	case "internalLinkTypeDefaultMessageAutoDeleteTimerSettings":
		t.InternalLinkTypeDefaultMessageAutoDeleteTimerSettings = new(InternalLinkTypeDefaultMessageAutoDeleteTimerSettings)
		return json.Unmarshal(b, t.InternalLinkTypeDefaultMessageAutoDeleteTimerSettings)
	case "internalLinkTypeDirectMessagesChat":
		t.InternalLinkTypeDirectMessagesChat = new(InternalLinkTypeDirectMessagesChat)
		return json.Unmarshal(b, t.InternalLinkTypeDirectMessagesChat)
	case "internalLinkTypeEditProfileSettings":
		t.InternalLinkTypeEditProfileSettings = new(InternalLinkTypeEditProfileSettings)
		return json.Unmarshal(b, t.InternalLinkTypeEditProfileSettings)
	case "internalLinkTypeGame":
		t.InternalLinkTypeGame = new(InternalLinkTypeGame)
		return json.Unmarshal(b, t.InternalLinkTypeGame)
	case "internalLinkTypeGiftAuction":
		t.InternalLinkTypeGiftAuction = new(InternalLinkTypeGiftAuction)
		return json.Unmarshal(b, t.InternalLinkTypeGiftAuction)
	case "internalLinkTypeGiftCollection":
		t.InternalLinkTypeGiftCollection = new(InternalLinkTypeGiftCollection)
		return json.Unmarshal(b, t.InternalLinkTypeGiftCollection)
	case "internalLinkTypeGroupCall":
		t.InternalLinkTypeGroupCall = new(InternalLinkTypeGroupCall)
		return json.Unmarshal(b, t.InternalLinkTypeGroupCall)
	case "internalLinkTypeInstantView":
		t.InternalLinkTypeInstantView = new(InternalLinkTypeInstantView)
		return json.Unmarshal(b, t.InternalLinkTypeInstantView)
	case "internalLinkTypeInvoice":
		t.InternalLinkTypeInvoice = new(InternalLinkTypeInvoice)
		return json.Unmarshal(b, t.InternalLinkTypeInvoice)
	case "internalLinkTypeLanguagePack":
		t.InternalLinkTypeLanguagePack = new(InternalLinkTypeLanguagePack)
		return json.Unmarshal(b, t.InternalLinkTypeLanguagePack)
	case "internalLinkTypeLanguageSettings":
		t.InternalLinkTypeLanguageSettings = new(InternalLinkTypeLanguageSettings)
		return json.Unmarshal(b, t.InternalLinkTypeLanguageSettings)
	case "internalLinkTypeLiveStory":
		t.InternalLinkTypeLiveStory = new(InternalLinkTypeLiveStory)
		return json.Unmarshal(b, t.InternalLinkTypeLiveStory)
	case "internalLinkTypeLoginEmailSettings":
		t.InternalLinkTypeLoginEmailSettings = new(InternalLinkTypeLoginEmailSettings)
		return json.Unmarshal(b, t.InternalLinkTypeLoginEmailSettings)
	case "internalLinkTypeMainWebApp":
		t.InternalLinkTypeMainWebApp = new(InternalLinkTypeMainWebApp)
		return json.Unmarshal(b, t.InternalLinkTypeMainWebApp)
	case "internalLinkTypeMessage":
		t.InternalLinkTypeMessage = new(InternalLinkTypeMessage)
		return json.Unmarshal(b, t.InternalLinkTypeMessage)
	case "internalLinkTypeMessageDraft":
		t.InternalLinkTypeMessageDraft = new(InternalLinkTypeMessageDraft)
		return json.Unmarshal(b, t.InternalLinkTypeMessageDraft)
	case "internalLinkTypeMyStars":
		t.InternalLinkTypeMyStars = new(InternalLinkTypeMyStars)
		return json.Unmarshal(b, t.InternalLinkTypeMyStars)
	case "internalLinkTypeMyToncoins":
		t.InternalLinkTypeMyToncoins = new(InternalLinkTypeMyToncoins)
		return json.Unmarshal(b, t.InternalLinkTypeMyToncoins)
	case "internalLinkTypePassportDataRequest":
		t.InternalLinkTypePassportDataRequest = new(InternalLinkTypePassportDataRequest)
		return json.Unmarshal(b, t.InternalLinkTypePassportDataRequest)
	case "internalLinkTypePasswordSettings":
		t.InternalLinkTypePasswordSettings = new(InternalLinkTypePasswordSettings)
		return json.Unmarshal(b, t.InternalLinkTypePasswordSettings)
	case "internalLinkTypePhoneNumberConfirmation":
		t.InternalLinkTypePhoneNumberConfirmation = new(InternalLinkTypePhoneNumberConfirmation)
		return json.Unmarshal(b, t.InternalLinkTypePhoneNumberConfirmation)
	case "internalLinkTypePhoneNumberPrivacySettings":
		t.InternalLinkTypePhoneNumberPrivacySettings = new(InternalLinkTypePhoneNumberPrivacySettings)
		return json.Unmarshal(b, t.InternalLinkTypePhoneNumberPrivacySettings)
	case "internalLinkTypePremiumFeatures":
		t.InternalLinkTypePremiumFeatures = new(InternalLinkTypePremiumFeatures)
		return json.Unmarshal(b, t.InternalLinkTypePremiumFeatures)
	case "internalLinkTypePremiumGift":
		t.InternalLinkTypePremiumGift = new(InternalLinkTypePremiumGift)
		return json.Unmarshal(b, t.InternalLinkTypePremiumGift)
	case "internalLinkTypePremiumGiftCode":
		t.InternalLinkTypePremiumGiftCode = new(InternalLinkTypePremiumGiftCode)
		return json.Unmarshal(b, t.InternalLinkTypePremiumGiftCode)
	case "internalLinkTypePrivacyAndSecuritySettings":
		t.InternalLinkTypePrivacyAndSecuritySettings = new(InternalLinkTypePrivacyAndSecuritySettings)
		return json.Unmarshal(b, t.InternalLinkTypePrivacyAndSecuritySettings)
	case "internalLinkTypeProxy":
		t.InternalLinkTypeProxy = new(InternalLinkTypeProxy)
		return json.Unmarshal(b, t.InternalLinkTypeProxy)
	case "internalLinkTypePublicChat":
		t.InternalLinkTypePublicChat = new(InternalLinkTypePublicChat)
		return json.Unmarshal(b, t.InternalLinkTypePublicChat)
	case "internalLinkTypeQrCodeAuthentication":
		t.InternalLinkTypeQrCodeAuthentication = new(InternalLinkTypeQrCodeAuthentication)
		return json.Unmarshal(b, t.InternalLinkTypeQrCodeAuthentication)
	case "internalLinkTypeRestorePurchases":
		t.InternalLinkTypeRestorePurchases = new(InternalLinkTypeRestorePurchases)
		return json.Unmarshal(b, t.InternalLinkTypeRestorePurchases)
	case "internalLinkTypeSettings":
		t.InternalLinkTypeSettings = new(InternalLinkTypeSettings)
		return json.Unmarshal(b, t.InternalLinkTypeSettings)
	case "internalLinkTypeStickerSet":
		t.InternalLinkTypeStickerSet = new(InternalLinkTypeStickerSet)
		return json.Unmarshal(b, t.InternalLinkTypeStickerSet)
	case "internalLinkTypeStory":
		t.InternalLinkTypeStory = new(InternalLinkTypeStory)
		return json.Unmarshal(b, t.InternalLinkTypeStory)
	case "internalLinkTypeStoryAlbum":
		t.InternalLinkTypeStoryAlbum = new(InternalLinkTypeStoryAlbum)
		return json.Unmarshal(b, t.InternalLinkTypeStoryAlbum)
	case "internalLinkTypeTheme":
		t.InternalLinkTypeTheme = new(InternalLinkTypeTheme)
		return json.Unmarshal(b, t.InternalLinkTypeTheme)
	case "internalLinkTypeThemeSettings":
		t.InternalLinkTypeThemeSettings = new(InternalLinkTypeThemeSettings)
		return json.Unmarshal(b, t.InternalLinkTypeThemeSettings)
	case "internalLinkTypeUnknownDeepLink":
		t.InternalLinkTypeUnknownDeepLink = new(InternalLinkTypeUnknownDeepLink)
		return json.Unmarshal(b, t.InternalLinkTypeUnknownDeepLink)
	case "internalLinkTypeUnsupportedProxy":
		t.InternalLinkTypeUnsupportedProxy = new(InternalLinkTypeUnsupportedProxy)
		return json.Unmarshal(b, t.InternalLinkTypeUnsupportedProxy)
	case "internalLinkTypeUpgradedGift":
		t.InternalLinkTypeUpgradedGift = new(InternalLinkTypeUpgradedGift)
		return json.Unmarshal(b, t.InternalLinkTypeUpgradedGift)
	case "internalLinkTypeUserPhoneNumber":
		t.InternalLinkTypeUserPhoneNumber = new(InternalLinkTypeUserPhoneNumber)
		return json.Unmarshal(b, t.InternalLinkTypeUserPhoneNumber)
	case "internalLinkTypeUserToken":
		t.InternalLinkTypeUserToken = new(InternalLinkTypeUserToken)
		return json.Unmarshal(b, t.InternalLinkTypeUserToken)
	case "internalLinkTypeVideoChat":
		t.InternalLinkTypeVideoChat = new(InternalLinkTypeVideoChat)
		return json.Unmarshal(b, t.InternalLinkTypeVideoChat)
	case "internalLinkTypeWebApp":
		t.InternalLinkTypeWebApp = new(InternalLinkTypeWebApp)
		return json.Unmarshal(b, t.InternalLinkTypeWebApp)
	}
	return nil
}

func (t *InternalLinkType) MarshalJSON() ([]byte, error) {
	if t.InternalLinkTypeActiveSessions != nil {
		return json.Marshal(t.InternalLinkTypeActiveSessions)
	}
	if t.InternalLinkTypeAttachmentMenuBot != nil {
		return json.Marshal(t.InternalLinkTypeAttachmentMenuBot)
	}
	if t.InternalLinkTypeAuthenticationCode != nil {
		return json.Marshal(t.InternalLinkTypeAuthenticationCode)
	}
	if t.InternalLinkTypeBackground != nil {
		return json.Marshal(t.InternalLinkTypeBackground)
	}
	if t.InternalLinkTypeBotAddToChannel != nil {
		return json.Marshal(t.InternalLinkTypeBotAddToChannel)
	}
	if t.InternalLinkTypeBotStart != nil {
		return json.Marshal(t.InternalLinkTypeBotStart)
	}
	if t.InternalLinkTypeBotStartInGroup != nil {
		return json.Marshal(t.InternalLinkTypeBotStartInGroup)
	}
	if t.InternalLinkTypeBusinessChat != nil {
		return json.Marshal(t.InternalLinkTypeBusinessChat)
	}
	if t.InternalLinkTypeBuyStars != nil {
		return json.Marshal(t.InternalLinkTypeBuyStars)
	}
	if t.InternalLinkTypeChangePhoneNumber != nil {
		return json.Marshal(t.InternalLinkTypeChangePhoneNumber)
	}
	if t.InternalLinkTypeChatAffiliateProgram != nil {
		return json.Marshal(t.InternalLinkTypeChatAffiliateProgram)
	}
	if t.InternalLinkTypeChatBoost != nil {
		return json.Marshal(t.InternalLinkTypeChatBoost)
	}
	if t.InternalLinkTypeChatFolderInvite != nil {
		return json.Marshal(t.InternalLinkTypeChatFolderInvite)
	}
	if t.InternalLinkTypeChatFolderSettings != nil {
		return json.Marshal(t.InternalLinkTypeChatFolderSettings)
	}
	if t.InternalLinkTypeChatInvite != nil {
		return json.Marshal(t.InternalLinkTypeChatInvite)
	}
	if t.InternalLinkTypeDefaultMessageAutoDeleteTimerSettings != nil {
		return json.Marshal(t.InternalLinkTypeDefaultMessageAutoDeleteTimerSettings)
	}
	if t.InternalLinkTypeDirectMessagesChat != nil {
		return json.Marshal(t.InternalLinkTypeDirectMessagesChat)
	}
	if t.InternalLinkTypeEditProfileSettings != nil {
		return json.Marshal(t.InternalLinkTypeEditProfileSettings)
	}
	if t.InternalLinkTypeGame != nil {
		return json.Marshal(t.InternalLinkTypeGame)
	}
	if t.InternalLinkTypeGiftAuction != nil {
		return json.Marshal(t.InternalLinkTypeGiftAuction)
	}
	if t.InternalLinkTypeGiftCollection != nil {
		return json.Marshal(t.InternalLinkTypeGiftCollection)
	}
	if t.InternalLinkTypeGroupCall != nil {
		return json.Marshal(t.InternalLinkTypeGroupCall)
	}
	if t.InternalLinkTypeInstantView != nil {
		return json.Marshal(t.InternalLinkTypeInstantView)
	}
	if t.InternalLinkTypeInvoice != nil {
		return json.Marshal(t.InternalLinkTypeInvoice)
	}
	if t.InternalLinkTypeLanguagePack != nil {
		return json.Marshal(t.InternalLinkTypeLanguagePack)
	}
	if t.InternalLinkTypeLanguageSettings != nil {
		return json.Marshal(t.InternalLinkTypeLanguageSettings)
	}
	if t.InternalLinkTypeLiveStory != nil {
		return json.Marshal(t.InternalLinkTypeLiveStory)
	}
	if t.InternalLinkTypeLoginEmailSettings != nil {
		return json.Marshal(t.InternalLinkTypeLoginEmailSettings)
	}
	if t.InternalLinkTypeMainWebApp != nil {
		return json.Marshal(t.InternalLinkTypeMainWebApp)
	}
	if t.InternalLinkTypeMessage != nil {
		return json.Marshal(t.InternalLinkTypeMessage)
	}
	if t.InternalLinkTypeMessageDraft != nil {
		return json.Marshal(t.InternalLinkTypeMessageDraft)
	}
	if t.InternalLinkTypeMyStars != nil {
		return json.Marshal(t.InternalLinkTypeMyStars)
	}
	if t.InternalLinkTypeMyToncoins != nil {
		return json.Marshal(t.InternalLinkTypeMyToncoins)
	}
	if t.InternalLinkTypePassportDataRequest != nil {
		return json.Marshal(t.InternalLinkTypePassportDataRequest)
	}
	if t.InternalLinkTypePasswordSettings != nil {
		return json.Marshal(t.InternalLinkTypePasswordSettings)
	}
	if t.InternalLinkTypePhoneNumberConfirmation != nil {
		return json.Marshal(t.InternalLinkTypePhoneNumberConfirmation)
	}
	if t.InternalLinkTypePhoneNumberPrivacySettings != nil {
		return json.Marshal(t.InternalLinkTypePhoneNumberPrivacySettings)
	}
	if t.InternalLinkTypePremiumFeatures != nil {
		return json.Marshal(t.InternalLinkTypePremiumFeatures)
	}
	if t.InternalLinkTypePremiumGift != nil {
		return json.Marshal(t.InternalLinkTypePremiumGift)
	}
	if t.InternalLinkTypePremiumGiftCode != nil {
		return json.Marshal(t.InternalLinkTypePremiumGiftCode)
	}
	if t.InternalLinkTypePrivacyAndSecuritySettings != nil {
		return json.Marshal(t.InternalLinkTypePrivacyAndSecuritySettings)
	}
	if t.InternalLinkTypeProxy != nil {
		return json.Marshal(t.InternalLinkTypeProxy)
	}
	if t.InternalLinkTypePublicChat != nil {
		return json.Marshal(t.InternalLinkTypePublicChat)
	}
	if t.InternalLinkTypeQrCodeAuthentication != nil {
		return json.Marshal(t.InternalLinkTypeQrCodeAuthentication)
	}
	if t.InternalLinkTypeRestorePurchases != nil {
		return json.Marshal(t.InternalLinkTypeRestorePurchases)
	}
	if t.InternalLinkTypeSettings != nil {
		return json.Marshal(t.InternalLinkTypeSettings)
	}
	if t.InternalLinkTypeStickerSet != nil {
		return json.Marshal(t.InternalLinkTypeStickerSet)
	}
	if t.InternalLinkTypeStory != nil {
		return json.Marshal(t.InternalLinkTypeStory)
	}
	if t.InternalLinkTypeStoryAlbum != nil {
		return json.Marshal(t.InternalLinkTypeStoryAlbum)
	}
	if t.InternalLinkTypeTheme != nil {
		return json.Marshal(t.InternalLinkTypeTheme)
	}
	if t.InternalLinkTypeThemeSettings != nil {
		return json.Marshal(t.InternalLinkTypeThemeSettings)
	}
	if t.InternalLinkTypeUnknownDeepLink != nil {
		return json.Marshal(t.InternalLinkTypeUnknownDeepLink)
	}
	if t.InternalLinkTypeUnsupportedProxy != nil {
		return json.Marshal(t.InternalLinkTypeUnsupportedProxy)
	}
	if t.InternalLinkTypeUpgradedGift != nil {
		return json.Marshal(t.InternalLinkTypeUpgradedGift)
	}
	if t.InternalLinkTypeUserPhoneNumber != nil {
		return json.Marshal(t.InternalLinkTypeUserPhoneNumber)
	}
	if t.InternalLinkTypeUserToken != nil {
		return json.Marshal(t.InternalLinkTypeUserToken)
	}
	if t.InternalLinkTypeVideoChat != nil {
		return json.Marshal(t.InternalLinkTypeVideoChat)
	}
	if t.InternalLinkTypeWebApp != nil {
		return json.Marshal(t.InternalLinkTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for InternalLinkType")
}

// InviteGroupCallParticipantResult Describes result of group call participant invitation
type InviteGroupCallParticipantResult struct {
	TypeStr                                                string                                                  `json:"@type"`
	InviteGroupCallParticipantResultSuccess                *InviteGroupCallParticipantResultSuccess                `json:"inviteGroupCallParticipantResultSuccess,omitempty"`
	InviteGroupCallParticipantResultUserAlreadyParticipant *InviteGroupCallParticipantResultUserAlreadyParticipant `json:"inviteGroupCallParticipantResultUserAlreadyParticipant,omitempty"`
	InviteGroupCallParticipantResultUserPrivacyRestricted  *InviteGroupCallParticipantResultUserPrivacyRestricted  `json:"inviteGroupCallParticipantResultUserPrivacyRestricted,omitempty"`
	InviteGroupCallParticipantResultUserWasBanned          *InviteGroupCallParticipantResultUserWasBanned          `json:"inviteGroupCallParticipantResultUserWasBanned,omitempty"`
}

func (t *InviteGroupCallParticipantResult) Type() string {
	return t.TypeStr
}

func (t *InviteGroupCallParticipantResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InviteGroupCallParticipantResult) GetExtra() string {
	return ""
}

func (t *InviteGroupCallParticipantResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inviteGroupCallParticipantResultSuccess":
		t.InviteGroupCallParticipantResultSuccess = new(InviteGroupCallParticipantResultSuccess)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultSuccess)
	case "inviteGroupCallParticipantResultUserAlreadyParticipant":
		t.InviteGroupCallParticipantResultUserAlreadyParticipant = new(InviteGroupCallParticipantResultUserAlreadyParticipant)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserAlreadyParticipant)
	case "inviteGroupCallParticipantResultUserPrivacyRestricted":
		t.InviteGroupCallParticipantResultUserPrivacyRestricted = new(InviteGroupCallParticipantResultUserPrivacyRestricted)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserPrivacyRestricted)
	case "inviteGroupCallParticipantResultUserWasBanned":
		t.InviteGroupCallParticipantResultUserWasBanned = new(InviteGroupCallParticipantResultUserWasBanned)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserWasBanned)
	}
	return nil
}

func (t *InviteGroupCallParticipantResult) MarshalJSON() ([]byte, error) {
	if t.InviteGroupCallParticipantResultSuccess != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultSuccess)
	}
	if t.InviteGroupCallParticipantResultUserAlreadyParticipant != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserAlreadyParticipant)
	}
	if t.InviteGroupCallParticipantResultUserPrivacyRestricted != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserPrivacyRestricted)
	}
	if t.InviteGroupCallParticipantResultUserWasBanned != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserWasBanned)
	}
	return nil, fmt.Errorf("no value set for InviteGroupCallParticipantResult")
}

// InviteLinkChatType Describes the type of chat to which points an invite link
type InviteLinkChatType struct {
	TypeStr                      string                        `json:"@type"`
	InviteLinkChatTypeBasicGroup *InviteLinkChatTypeBasicGroup `json:"inviteLinkChatTypeBasicGroup,omitempty"`
	InviteLinkChatTypeChannel    *InviteLinkChatTypeChannel    `json:"inviteLinkChatTypeChannel,omitempty"`
	InviteLinkChatTypeSupergroup *InviteLinkChatTypeSupergroup `json:"inviteLinkChatTypeSupergroup,omitempty"`
}

func (t *InviteLinkChatType) Type() string {
	return t.TypeStr
}

func (t *InviteLinkChatType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *InviteLinkChatType) GetExtra() string {
	return ""
}

func (t *InviteLinkChatType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "inviteLinkChatTypeBasicGroup":
		t.InviteLinkChatTypeBasicGroup = new(InviteLinkChatTypeBasicGroup)
		return json.Unmarshal(b, t.InviteLinkChatTypeBasicGroup)
	case "inviteLinkChatTypeChannel":
		t.InviteLinkChatTypeChannel = new(InviteLinkChatTypeChannel)
		return json.Unmarshal(b, t.InviteLinkChatTypeChannel)
	case "inviteLinkChatTypeSupergroup":
		t.InviteLinkChatTypeSupergroup = new(InviteLinkChatTypeSupergroup)
		return json.Unmarshal(b, t.InviteLinkChatTypeSupergroup)
	}
	return nil
}

func (t *InviteLinkChatType) MarshalJSON() ([]byte, error) {
	if t.InviteLinkChatTypeBasicGroup != nil {
		return json.Marshal(t.InviteLinkChatTypeBasicGroup)
	}
	if t.InviteLinkChatTypeChannel != nil {
		return json.Marshal(t.InviteLinkChatTypeChannel)
	}
	if t.InviteLinkChatTypeSupergroup != nil {
		return json.Marshal(t.InviteLinkChatTypeSupergroup)
	}
	return nil, fmt.Errorf("no value set for InviteLinkChatType")
}

// JsonValue Represents a JSON value
type JsonValue struct {
	TypeStr          string            `json:"@type"`
	JsonValueArray   *JsonValueArray   `json:"jsonValueArray,omitempty"`
	JsonValueBoolean *JsonValueBoolean `json:"jsonValueBoolean,omitempty"`
	JsonValueNull    *JsonValueNull    `json:"jsonValueNull,omitempty"`
	JsonValueNumber  *JsonValueNumber  `json:"jsonValueNumber,omitempty"`
	JsonValueObject  *JsonValueObject  `json:"jsonValueObject,omitempty"`
	JsonValueString  *JsonValueString  `json:"jsonValueString,omitempty"`
}

func (t *JsonValue) Type() string {
	return t.TypeStr
}

func (t *JsonValue) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *JsonValue) GetExtra() string {
	return ""
}

func (t *JsonValue) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "jsonValueArray":
		t.JsonValueArray = new(JsonValueArray)
		return json.Unmarshal(b, t.JsonValueArray)
	case "jsonValueBoolean":
		t.JsonValueBoolean = new(JsonValueBoolean)
		return json.Unmarshal(b, t.JsonValueBoolean)
	case "jsonValueNull":
		t.JsonValueNull = new(JsonValueNull)
		return json.Unmarshal(b, t.JsonValueNull)
	case "jsonValueNumber":
		t.JsonValueNumber = new(JsonValueNumber)
		return json.Unmarshal(b, t.JsonValueNumber)
	case "jsonValueObject":
		t.JsonValueObject = new(JsonValueObject)
		return json.Unmarshal(b, t.JsonValueObject)
	case "jsonValueString":
		t.JsonValueString = new(JsonValueString)
		return json.Unmarshal(b, t.JsonValueString)
	}
	return nil
}

func (t *JsonValue) MarshalJSON() ([]byte, error) {
	if t.JsonValueArray != nil {
		return json.Marshal(t.JsonValueArray)
	}
	if t.JsonValueBoolean != nil {
		return json.Marshal(t.JsonValueBoolean)
	}
	if t.JsonValueNull != nil {
		return json.Marshal(t.JsonValueNull)
	}
	if t.JsonValueNumber != nil {
		return json.Marshal(t.JsonValueNumber)
	}
	if t.JsonValueObject != nil {
		return json.Marshal(t.JsonValueObject)
	}
	if t.JsonValueString != nil {
		return json.Marshal(t.JsonValueString)
	}
	return nil, fmt.Errorf("no value set for JsonValue")
}

// KeyboardButtonType Describes a keyboard button type
type KeyboardButtonType struct {
	TypeStr                              string                                `json:"@type"`
	KeyboardButtonTypeRequestChat        *KeyboardButtonTypeRequestChat        `json:"keyboardButtonTypeRequestChat,omitempty"`
	KeyboardButtonTypeRequestLocation    *KeyboardButtonTypeRequestLocation    `json:"keyboardButtonTypeRequestLocation,omitempty"`
	KeyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber `json:"keyboardButtonTypeRequestPhoneNumber,omitempty"`
	KeyboardButtonTypeRequestPoll        *KeyboardButtonTypeRequestPoll        `json:"keyboardButtonTypeRequestPoll,omitempty"`
	KeyboardButtonTypeRequestUsers       *KeyboardButtonTypeRequestUsers       `json:"keyboardButtonTypeRequestUsers,omitempty"`
	KeyboardButtonTypeText               *KeyboardButtonTypeText               `json:"keyboardButtonTypeText,omitempty"`
	KeyboardButtonTypeWebApp             *KeyboardButtonTypeWebApp             `json:"keyboardButtonTypeWebApp,omitempty"`
}

func (t *KeyboardButtonType) Type() string {
	return t.TypeStr
}

func (t *KeyboardButtonType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *KeyboardButtonType) GetExtra() string {
	return ""
}

func (t *KeyboardButtonType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "keyboardButtonTypeRequestChat":
		t.KeyboardButtonTypeRequestChat = new(KeyboardButtonTypeRequestChat)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestChat)
	case "keyboardButtonTypeRequestLocation":
		t.KeyboardButtonTypeRequestLocation = new(KeyboardButtonTypeRequestLocation)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestLocation)
	case "keyboardButtonTypeRequestPhoneNumber":
		t.KeyboardButtonTypeRequestPhoneNumber = new(KeyboardButtonTypeRequestPhoneNumber)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestPhoneNumber)
	case "keyboardButtonTypeRequestPoll":
		t.KeyboardButtonTypeRequestPoll = new(KeyboardButtonTypeRequestPoll)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestPoll)
	case "keyboardButtonTypeRequestUsers":
		t.KeyboardButtonTypeRequestUsers = new(KeyboardButtonTypeRequestUsers)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestUsers)
	case "keyboardButtonTypeText":
		t.KeyboardButtonTypeText = new(KeyboardButtonTypeText)
		return json.Unmarshal(b, t.KeyboardButtonTypeText)
	case "keyboardButtonTypeWebApp":
		t.KeyboardButtonTypeWebApp = new(KeyboardButtonTypeWebApp)
		return json.Unmarshal(b, t.KeyboardButtonTypeWebApp)
	}
	return nil
}

func (t *KeyboardButtonType) MarshalJSON() ([]byte, error) {
	if t.KeyboardButtonTypeRequestChat != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestChat)
	}
	if t.KeyboardButtonTypeRequestLocation != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestLocation)
	}
	if t.KeyboardButtonTypeRequestPhoneNumber != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestPhoneNumber)
	}
	if t.KeyboardButtonTypeRequestPoll != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestPoll)
	}
	if t.KeyboardButtonTypeRequestUsers != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestUsers)
	}
	if t.KeyboardButtonTypeText != nil {
		return json.Marshal(t.KeyboardButtonTypeText)
	}
	if t.KeyboardButtonTypeWebApp != nil {
		return json.Marshal(t.KeyboardButtonTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for KeyboardButtonType")
}

// LanguagePackStringValue Represents the value of a string in a language pack
type LanguagePackStringValue struct {
	TypeStr                           string                             `json:"@type"`
	LanguagePackStringValueDeleted    *LanguagePackStringValueDeleted    `json:"languagePackStringValueDeleted,omitempty"`
	LanguagePackStringValueOrdinary   *LanguagePackStringValueOrdinary   `json:"languagePackStringValueOrdinary,omitempty"`
	LanguagePackStringValuePluralized *LanguagePackStringValuePluralized `json:"languagePackStringValuePluralized,omitempty"`
}

func (t *LanguagePackStringValue) Type() string {
	return t.TypeStr
}

func (t *LanguagePackStringValue) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *LanguagePackStringValue) GetExtra() string {
	return ""
}

func (t *LanguagePackStringValue) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "languagePackStringValueDeleted":
		t.LanguagePackStringValueDeleted = new(LanguagePackStringValueDeleted)
		return json.Unmarshal(b, t.LanguagePackStringValueDeleted)
	case "languagePackStringValueOrdinary":
		t.LanguagePackStringValueOrdinary = new(LanguagePackStringValueOrdinary)
		return json.Unmarshal(b, t.LanguagePackStringValueOrdinary)
	case "languagePackStringValuePluralized":
		t.LanguagePackStringValuePluralized = new(LanguagePackStringValuePluralized)
		return json.Unmarshal(b, t.LanguagePackStringValuePluralized)
	}
	return nil
}

func (t *LanguagePackStringValue) MarshalJSON() ([]byte, error) {
	if t.LanguagePackStringValueDeleted != nil {
		return json.Marshal(t.LanguagePackStringValueDeleted)
	}
	if t.LanguagePackStringValueOrdinary != nil {
		return json.Marshal(t.LanguagePackStringValueOrdinary)
	}
	if t.LanguagePackStringValuePluralized != nil {
		return json.Marshal(t.LanguagePackStringValuePluralized)
	}
	return nil, fmt.Errorf("no value set for LanguagePackStringValue")
}

// LinkPreviewAlbumMedia Describes a media from a link preview album
type LinkPreviewAlbumMedia struct {
	TypeStr                    string                      `json:"@type"`
	LinkPreviewAlbumMediaPhoto *LinkPreviewAlbumMediaPhoto `json:"linkPreviewAlbumMediaPhoto,omitempty"`
	LinkPreviewAlbumMediaVideo *LinkPreviewAlbumMediaVideo `json:"linkPreviewAlbumMediaVideo,omitempty"`
}

func (t *LinkPreviewAlbumMedia) Type() string {
	return t.TypeStr
}

func (t *LinkPreviewAlbumMedia) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *LinkPreviewAlbumMedia) GetExtra() string {
	return ""
}

func (t *LinkPreviewAlbumMedia) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "linkPreviewAlbumMediaPhoto":
		t.LinkPreviewAlbumMediaPhoto = new(LinkPreviewAlbumMediaPhoto)
		return json.Unmarshal(b, t.LinkPreviewAlbumMediaPhoto)
	case "linkPreviewAlbumMediaVideo":
		t.LinkPreviewAlbumMediaVideo = new(LinkPreviewAlbumMediaVideo)
		return json.Unmarshal(b, t.LinkPreviewAlbumMediaVideo)
	}
	return nil
}

func (t *LinkPreviewAlbumMedia) MarshalJSON() ([]byte, error) {
	if t.LinkPreviewAlbumMediaPhoto != nil {
		return json.Marshal(t.LinkPreviewAlbumMediaPhoto)
	}
	if t.LinkPreviewAlbumMediaVideo != nil {
		return json.Marshal(t.LinkPreviewAlbumMediaVideo)
	}
	return nil, fmt.Errorf("no value set for LinkPreviewAlbumMedia")
}

// LinkPreviewType Describes type of link preview
type LinkPreviewType struct {
	TypeStr                                string                                  `json:"@type"`
	LinkPreviewTypeAlbum                   *LinkPreviewTypeAlbum                   `json:"linkPreviewTypeAlbum,omitempty"`
	LinkPreviewTypeAnimation               *LinkPreviewTypeAnimation               `json:"linkPreviewTypeAnimation,omitempty"`
	LinkPreviewTypeApp                     *LinkPreviewTypeApp                     `json:"linkPreviewTypeApp,omitempty"`
	LinkPreviewTypeArticle                 *LinkPreviewTypeArticle                 `json:"linkPreviewTypeArticle,omitempty"`
	LinkPreviewTypeAudio                   *LinkPreviewTypeAudio                   `json:"linkPreviewTypeAudio,omitempty"`
	LinkPreviewTypeBackground              *LinkPreviewTypeBackground              `json:"linkPreviewTypeBackground,omitempty"`
	LinkPreviewTypeChannelBoost            *LinkPreviewTypeChannelBoost            `json:"linkPreviewTypeChannelBoost,omitempty"`
	LinkPreviewTypeChat                    *LinkPreviewTypeChat                    `json:"linkPreviewTypeChat,omitempty"`
	LinkPreviewTypeDirectMessagesChat      *LinkPreviewTypeDirectMessagesChat      `json:"linkPreviewTypeDirectMessagesChat,omitempty"`
	LinkPreviewTypeDocument                *LinkPreviewTypeDocument                `json:"linkPreviewTypeDocument,omitempty"`
	LinkPreviewTypeEmbeddedAnimationPlayer *LinkPreviewTypeEmbeddedAnimationPlayer `json:"linkPreviewTypeEmbeddedAnimationPlayer,omitempty"`
	LinkPreviewTypeEmbeddedAudioPlayer     *LinkPreviewTypeEmbeddedAudioPlayer     `json:"linkPreviewTypeEmbeddedAudioPlayer,omitempty"`
	LinkPreviewTypeEmbeddedVideoPlayer     *LinkPreviewTypeEmbeddedVideoPlayer     `json:"linkPreviewTypeEmbeddedVideoPlayer,omitempty"`
	LinkPreviewTypeExternalAudio           *LinkPreviewTypeExternalAudio           `json:"linkPreviewTypeExternalAudio,omitempty"`
	LinkPreviewTypeExternalVideo           *LinkPreviewTypeExternalVideo           `json:"linkPreviewTypeExternalVideo,omitempty"`
	LinkPreviewTypeGiftAuction             *LinkPreviewTypeGiftAuction             `json:"linkPreviewTypeGiftAuction,omitempty"`
	LinkPreviewTypeGiftCollection          *LinkPreviewTypeGiftCollection          `json:"linkPreviewTypeGiftCollection,omitempty"`
	LinkPreviewTypeGroupCall               *LinkPreviewTypeGroupCall               `json:"linkPreviewTypeGroupCall,omitempty"`
	LinkPreviewTypeInvoice                 *LinkPreviewTypeInvoice                 `json:"linkPreviewTypeInvoice,omitempty"`
	LinkPreviewTypeLiveStory               *LinkPreviewTypeLiveStory               `json:"linkPreviewTypeLiveStory,omitempty"`
	LinkPreviewTypeMessage                 *LinkPreviewTypeMessage                 `json:"linkPreviewTypeMessage,omitempty"`
	LinkPreviewTypePhoto                   *LinkPreviewTypePhoto                   `json:"linkPreviewTypePhoto,omitempty"`
	LinkPreviewTypePremiumGiftCode         *LinkPreviewTypePremiumGiftCode         `json:"linkPreviewTypePremiumGiftCode,omitempty"`
	LinkPreviewTypeShareableChatFolder     *LinkPreviewTypeShareableChatFolder     `json:"linkPreviewTypeShareableChatFolder,omitempty"`
	LinkPreviewTypeSticker                 *LinkPreviewTypeSticker                 `json:"linkPreviewTypeSticker,omitempty"`
	LinkPreviewTypeStickerSet              *LinkPreviewTypeStickerSet              `json:"linkPreviewTypeStickerSet,omitempty"`
	LinkPreviewTypeStory                   *LinkPreviewTypeStory                   `json:"linkPreviewTypeStory,omitempty"`
	LinkPreviewTypeStoryAlbum              *LinkPreviewTypeStoryAlbum              `json:"linkPreviewTypeStoryAlbum,omitempty"`
	LinkPreviewTypeSupergroupBoost         *LinkPreviewTypeSupergroupBoost         `json:"linkPreviewTypeSupergroupBoost,omitempty"`
	LinkPreviewTypeTheme                   *LinkPreviewTypeTheme                   `json:"linkPreviewTypeTheme,omitempty"`
	LinkPreviewTypeUnsupported             *LinkPreviewTypeUnsupported             `json:"linkPreviewTypeUnsupported,omitempty"`
	LinkPreviewTypeUpgradedGift            *LinkPreviewTypeUpgradedGift            `json:"linkPreviewTypeUpgradedGift,omitempty"`
	LinkPreviewTypeUser                    *LinkPreviewTypeUser                    `json:"linkPreviewTypeUser,omitempty"`
	LinkPreviewTypeVideo                   *LinkPreviewTypeVideo                   `json:"linkPreviewTypeVideo,omitempty"`
	LinkPreviewTypeVideoChat               *LinkPreviewTypeVideoChat               `json:"linkPreviewTypeVideoChat,omitempty"`
	LinkPreviewTypeVideoNote               *LinkPreviewTypeVideoNote               `json:"linkPreviewTypeVideoNote,omitempty"`
	LinkPreviewTypeVoiceNote               *LinkPreviewTypeVoiceNote               `json:"linkPreviewTypeVoiceNote,omitempty"`
	LinkPreviewTypeWebApp                  *LinkPreviewTypeWebApp                  `json:"linkPreviewTypeWebApp,omitempty"`
}

func (t *LinkPreviewType) Type() string {
	return t.TypeStr
}

func (t *LinkPreviewType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *LinkPreviewType) GetExtra() string {
	return ""
}

func (t *LinkPreviewType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "linkPreviewTypeAlbum":
		t.LinkPreviewTypeAlbum = new(LinkPreviewTypeAlbum)
		return json.Unmarshal(b, t.LinkPreviewTypeAlbum)
	case "linkPreviewTypeAnimation":
		t.LinkPreviewTypeAnimation = new(LinkPreviewTypeAnimation)
		return json.Unmarshal(b, t.LinkPreviewTypeAnimation)
	case "linkPreviewTypeApp":
		t.LinkPreviewTypeApp = new(LinkPreviewTypeApp)
		return json.Unmarshal(b, t.LinkPreviewTypeApp)
	case "linkPreviewTypeArticle":
		t.LinkPreviewTypeArticle = new(LinkPreviewTypeArticle)
		return json.Unmarshal(b, t.LinkPreviewTypeArticle)
	case "linkPreviewTypeAudio":
		t.LinkPreviewTypeAudio = new(LinkPreviewTypeAudio)
		return json.Unmarshal(b, t.LinkPreviewTypeAudio)
	case "linkPreviewTypeBackground":
		t.LinkPreviewTypeBackground = new(LinkPreviewTypeBackground)
		return json.Unmarshal(b, t.LinkPreviewTypeBackground)
	case "linkPreviewTypeChannelBoost":
		t.LinkPreviewTypeChannelBoost = new(LinkPreviewTypeChannelBoost)
		return json.Unmarshal(b, t.LinkPreviewTypeChannelBoost)
	case "linkPreviewTypeChat":
		t.LinkPreviewTypeChat = new(LinkPreviewTypeChat)
		return json.Unmarshal(b, t.LinkPreviewTypeChat)
	case "linkPreviewTypeDirectMessagesChat":
		t.LinkPreviewTypeDirectMessagesChat = new(LinkPreviewTypeDirectMessagesChat)
		return json.Unmarshal(b, t.LinkPreviewTypeDirectMessagesChat)
	case "linkPreviewTypeDocument":
		t.LinkPreviewTypeDocument = new(LinkPreviewTypeDocument)
		return json.Unmarshal(b, t.LinkPreviewTypeDocument)
	case "linkPreviewTypeEmbeddedAnimationPlayer":
		t.LinkPreviewTypeEmbeddedAnimationPlayer = new(LinkPreviewTypeEmbeddedAnimationPlayer)
		return json.Unmarshal(b, t.LinkPreviewTypeEmbeddedAnimationPlayer)
	case "linkPreviewTypeEmbeddedAudioPlayer":
		t.LinkPreviewTypeEmbeddedAudioPlayer = new(LinkPreviewTypeEmbeddedAudioPlayer)
		return json.Unmarshal(b, t.LinkPreviewTypeEmbeddedAudioPlayer)
	case "linkPreviewTypeEmbeddedVideoPlayer":
		t.LinkPreviewTypeEmbeddedVideoPlayer = new(LinkPreviewTypeEmbeddedVideoPlayer)
		return json.Unmarshal(b, t.LinkPreviewTypeEmbeddedVideoPlayer)
	case "linkPreviewTypeExternalAudio":
		t.LinkPreviewTypeExternalAudio = new(LinkPreviewTypeExternalAudio)
		return json.Unmarshal(b, t.LinkPreviewTypeExternalAudio)
	case "linkPreviewTypeExternalVideo":
		t.LinkPreviewTypeExternalVideo = new(LinkPreviewTypeExternalVideo)
		return json.Unmarshal(b, t.LinkPreviewTypeExternalVideo)
	case "linkPreviewTypeGiftAuction":
		t.LinkPreviewTypeGiftAuction = new(LinkPreviewTypeGiftAuction)
		return json.Unmarshal(b, t.LinkPreviewTypeGiftAuction)
	case "linkPreviewTypeGiftCollection":
		t.LinkPreviewTypeGiftCollection = new(LinkPreviewTypeGiftCollection)
		return json.Unmarshal(b, t.LinkPreviewTypeGiftCollection)
	case "linkPreviewTypeGroupCall":
		t.LinkPreviewTypeGroupCall = new(LinkPreviewTypeGroupCall)
		return json.Unmarshal(b, t.LinkPreviewTypeGroupCall)
	case "linkPreviewTypeInvoice":
		t.LinkPreviewTypeInvoice = new(LinkPreviewTypeInvoice)
		return json.Unmarshal(b, t.LinkPreviewTypeInvoice)
	case "linkPreviewTypeLiveStory":
		t.LinkPreviewTypeLiveStory = new(LinkPreviewTypeLiveStory)
		return json.Unmarshal(b, t.LinkPreviewTypeLiveStory)
	case "linkPreviewTypeMessage":
		t.LinkPreviewTypeMessage = new(LinkPreviewTypeMessage)
		return json.Unmarshal(b, t.LinkPreviewTypeMessage)
	case "linkPreviewTypePhoto":
		t.LinkPreviewTypePhoto = new(LinkPreviewTypePhoto)
		return json.Unmarshal(b, t.LinkPreviewTypePhoto)
	case "linkPreviewTypePremiumGiftCode":
		t.LinkPreviewTypePremiumGiftCode = new(LinkPreviewTypePremiumGiftCode)
		return json.Unmarshal(b, t.LinkPreviewTypePremiumGiftCode)
	case "linkPreviewTypeShareableChatFolder":
		t.LinkPreviewTypeShareableChatFolder = new(LinkPreviewTypeShareableChatFolder)
		return json.Unmarshal(b, t.LinkPreviewTypeShareableChatFolder)
	case "linkPreviewTypeSticker":
		t.LinkPreviewTypeSticker = new(LinkPreviewTypeSticker)
		return json.Unmarshal(b, t.LinkPreviewTypeSticker)
	case "linkPreviewTypeStickerSet":
		t.LinkPreviewTypeStickerSet = new(LinkPreviewTypeStickerSet)
		return json.Unmarshal(b, t.LinkPreviewTypeStickerSet)
	case "linkPreviewTypeStory":
		t.LinkPreviewTypeStory = new(LinkPreviewTypeStory)
		return json.Unmarshal(b, t.LinkPreviewTypeStory)
	case "linkPreviewTypeStoryAlbum":
		t.LinkPreviewTypeStoryAlbum = new(LinkPreviewTypeStoryAlbum)
		return json.Unmarshal(b, t.LinkPreviewTypeStoryAlbum)
	case "linkPreviewTypeSupergroupBoost":
		t.LinkPreviewTypeSupergroupBoost = new(LinkPreviewTypeSupergroupBoost)
		return json.Unmarshal(b, t.LinkPreviewTypeSupergroupBoost)
	case "linkPreviewTypeTheme":
		t.LinkPreviewTypeTheme = new(LinkPreviewTypeTheme)
		return json.Unmarshal(b, t.LinkPreviewTypeTheme)
	case "linkPreviewTypeUnsupported":
		t.LinkPreviewTypeUnsupported = new(LinkPreviewTypeUnsupported)
		return json.Unmarshal(b, t.LinkPreviewTypeUnsupported)
	case "linkPreviewTypeUpgradedGift":
		t.LinkPreviewTypeUpgradedGift = new(LinkPreviewTypeUpgradedGift)
		return json.Unmarshal(b, t.LinkPreviewTypeUpgradedGift)
	case "linkPreviewTypeUser":
		t.LinkPreviewTypeUser = new(LinkPreviewTypeUser)
		return json.Unmarshal(b, t.LinkPreviewTypeUser)
	case "linkPreviewTypeVideo":
		t.LinkPreviewTypeVideo = new(LinkPreviewTypeVideo)
		return json.Unmarshal(b, t.LinkPreviewTypeVideo)
	case "linkPreviewTypeVideoChat":
		t.LinkPreviewTypeVideoChat = new(LinkPreviewTypeVideoChat)
		return json.Unmarshal(b, t.LinkPreviewTypeVideoChat)
	case "linkPreviewTypeVideoNote":
		t.LinkPreviewTypeVideoNote = new(LinkPreviewTypeVideoNote)
		return json.Unmarshal(b, t.LinkPreviewTypeVideoNote)
	case "linkPreviewTypeVoiceNote":
		t.LinkPreviewTypeVoiceNote = new(LinkPreviewTypeVoiceNote)
		return json.Unmarshal(b, t.LinkPreviewTypeVoiceNote)
	case "linkPreviewTypeWebApp":
		t.LinkPreviewTypeWebApp = new(LinkPreviewTypeWebApp)
		return json.Unmarshal(b, t.LinkPreviewTypeWebApp)
	}
	return nil
}

func (t *LinkPreviewType) MarshalJSON() ([]byte, error) {
	if t.LinkPreviewTypeAlbum != nil {
		return json.Marshal(t.LinkPreviewTypeAlbum)
	}
	if t.LinkPreviewTypeAnimation != nil {
		return json.Marshal(t.LinkPreviewTypeAnimation)
	}
	if t.LinkPreviewTypeApp != nil {
		return json.Marshal(t.LinkPreviewTypeApp)
	}
	if t.LinkPreviewTypeArticle != nil {
		return json.Marshal(t.LinkPreviewTypeArticle)
	}
	if t.LinkPreviewTypeAudio != nil {
		return json.Marshal(t.LinkPreviewTypeAudio)
	}
	if t.LinkPreviewTypeBackground != nil {
		return json.Marshal(t.LinkPreviewTypeBackground)
	}
	if t.LinkPreviewTypeChannelBoost != nil {
		return json.Marshal(t.LinkPreviewTypeChannelBoost)
	}
	if t.LinkPreviewTypeChat != nil {
		return json.Marshal(t.LinkPreviewTypeChat)
	}
	if t.LinkPreviewTypeDirectMessagesChat != nil {
		return json.Marshal(t.LinkPreviewTypeDirectMessagesChat)
	}
	if t.LinkPreviewTypeDocument != nil {
		return json.Marshal(t.LinkPreviewTypeDocument)
	}
	if t.LinkPreviewTypeEmbeddedAnimationPlayer != nil {
		return json.Marshal(t.LinkPreviewTypeEmbeddedAnimationPlayer)
	}
	if t.LinkPreviewTypeEmbeddedAudioPlayer != nil {
		return json.Marshal(t.LinkPreviewTypeEmbeddedAudioPlayer)
	}
	if t.LinkPreviewTypeEmbeddedVideoPlayer != nil {
		return json.Marshal(t.LinkPreviewTypeEmbeddedVideoPlayer)
	}
	if t.LinkPreviewTypeExternalAudio != nil {
		return json.Marshal(t.LinkPreviewTypeExternalAudio)
	}
	if t.LinkPreviewTypeExternalVideo != nil {
		return json.Marshal(t.LinkPreviewTypeExternalVideo)
	}
	if t.LinkPreviewTypeGiftAuction != nil {
		return json.Marshal(t.LinkPreviewTypeGiftAuction)
	}
	if t.LinkPreviewTypeGiftCollection != nil {
		return json.Marshal(t.LinkPreviewTypeGiftCollection)
	}
	if t.LinkPreviewTypeGroupCall != nil {
		return json.Marshal(t.LinkPreviewTypeGroupCall)
	}
	if t.LinkPreviewTypeInvoice != nil {
		return json.Marshal(t.LinkPreviewTypeInvoice)
	}
	if t.LinkPreviewTypeLiveStory != nil {
		return json.Marshal(t.LinkPreviewTypeLiveStory)
	}
	if t.LinkPreviewTypeMessage != nil {
		return json.Marshal(t.LinkPreviewTypeMessage)
	}
	if t.LinkPreviewTypePhoto != nil {
		return json.Marshal(t.LinkPreviewTypePhoto)
	}
	if t.LinkPreviewTypePremiumGiftCode != nil {
		return json.Marshal(t.LinkPreviewTypePremiumGiftCode)
	}
	if t.LinkPreviewTypeShareableChatFolder != nil {
		return json.Marshal(t.LinkPreviewTypeShareableChatFolder)
	}
	if t.LinkPreviewTypeSticker != nil {
		return json.Marshal(t.LinkPreviewTypeSticker)
	}
	if t.LinkPreviewTypeStickerSet != nil {
		return json.Marshal(t.LinkPreviewTypeStickerSet)
	}
	if t.LinkPreviewTypeStory != nil {
		return json.Marshal(t.LinkPreviewTypeStory)
	}
	if t.LinkPreviewTypeStoryAlbum != nil {
		return json.Marshal(t.LinkPreviewTypeStoryAlbum)
	}
	if t.LinkPreviewTypeSupergroupBoost != nil {
		return json.Marshal(t.LinkPreviewTypeSupergroupBoost)
	}
	if t.LinkPreviewTypeTheme != nil {
		return json.Marshal(t.LinkPreviewTypeTheme)
	}
	if t.LinkPreviewTypeUnsupported != nil {
		return json.Marshal(t.LinkPreviewTypeUnsupported)
	}
	if t.LinkPreviewTypeUpgradedGift != nil {
		return json.Marshal(t.LinkPreviewTypeUpgradedGift)
	}
	if t.LinkPreviewTypeUser != nil {
		return json.Marshal(t.LinkPreviewTypeUser)
	}
	if t.LinkPreviewTypeVideo != nil {
		return json.Marshal(t.LinkPreviewTypeVideo)
	}
	if t.LinkPreviewTypeVideoChat != nil {
		return json.Marshal(t.LinkPreviewTypeVideoChat)
	}
	if t.LinkPreviewTypeVideoNote != nil {
		return json.Marshal(t.LinkPreviewTypeVideoNote)
	}
	if t.LinkPreviewTypeVoiceNote != nil {
		return json.Marshal(t.LinkPreviewTypeVoiceNote)
	}
	if t.LinkPreviewTypeWebApp != nil {
		return json.Marshal(t.LinkPreviewTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for LinkPreviewType")
}

// LoginUrlInfo Contains information about an inline button of type inlineKeyboardButtonTypeLoginUrl
type LoginUrlInfo struct {
	TypeStr                         string                           `json:"@type"`
	LoginUrlInfoOpen                *LoginUrlInfoOpen                `json:"loginUrlInfoOpen,omitempty"`
	LoginUrlInfoRequestConfirmation *LoginUrlInfoRequestConfirmation `json:"loginUrlInfoRequestConfirmation,omitempty"`
}

func (t *LoginUrlInfo) Type() string {
	return t.TypeStr
}

func (t *LoginUrlInfo) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *LoginUrlInfo) GetExtra() string {
	return ""
}

func (t *LoginUrlInfo) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "loginUrlInfoOpen":
		t.LoginUrlInfoOpen = new(LoginUrlInfoOpen)
		return json.Unmarshal(b, t.LoginUrlInfoOpen)
	case "loginUrlInfoRequestConfirmation":
		t.LoginUrlInfoRequestConfirmation = new(LoginUrlInfoRequestConfirmation)
		return json.Unmarshal(b, t.LoginUrlInfoRequestConfirmation)
	}
	return nil
}

func (t *LoginUrlInfo) MarshalJSON() ([]byte, error) {
	if t.LoginUrlInfoOpen != nil {
		return json.Marshal(t.LoginUrlInfoOpen)
	}
	if t.LoginUrlInfoRequestConfirmation != nil {
		return json.Marshal(t.LoginUrlInfoRequestConfirmation)
	}
	return nil, fmt.Errorf("no value set for LoginUrlInfo")
}

// LogStream Describes a stream to which TDLib internal log is written
type LogStream struct {
	TypeStr          string            `json:"@type"`
	LogStreamDefault *LogStreamDefault `json:"logStreamDefault,omitempty"`
	LogStreamEmpty   *LogStreamEmpty   `json:"logStreamEmpty,omitempty"`
	LogStreamFile    *LogStreamFile    `json:"logStreamFile,omitempty"`
}

func (t *LogStream) Type() string {
	return t.TypeStr
}

func (t *LogStream) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *LogStream) GetExtra() string {
	return ""
}

func (t *LogStream) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "logStreamDefault":
		t.LogStreamDefault = new(LogStreamDefault)
		return json.Unmarshal(b, t.LogStreamDefault)
	case "logStreamEmpty":
		t.LogStreamEmpty = new(LogStreamEmpty)
		return json.Unmarshal(b, t.LogStreamEmpty)
	case "logStreamFile":
		t.LogStreamFile = new(LogStreamFile)
		return json.Unmarshal(b, t.LogStreamFile)
	}
	return nil
}

func (t *LogStream) MarshalJSON() ([]byte, error) {
	if t.LogStreamDefault != nil {
		return json.Marshal(t.LogStreamDefault)
	}
	if t.LogStreamEmpty != nil {
		return json.Marshal(t.LogStreamEmpty)
	}
	if t.LogStreamFile != nil {
		return json.Marshal(t.LogStreamFile)
	}
	return nil, fmt.Errorf("no value set for LogStream")
}

// MaskPoint Part of the face, relative to which a mask is placed
type MaskPoint struct {
	TypeStr           string             `json:"@type"`
	MaskPointChin     *MaskPointChin     `json:"maskPointChin,omitempty"`
	MaskPointEyes     *MaskPointEyes     `json:"maskPointEyes,omitempty"`
	MaskPointForehead *MaskPointForehead `json:"maskPointForehead,omitempty"`
	MaskPointMouth    *MaskPointMouth    `json:"maskPointMouth,omitempty"`
}

func (t *MaskPoint) Type() string {
	return t.TypeStr
}

func (t *MaskPoint) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MaskPoint) GetExtra() string {
	return ""
}

func (t *MaskPoint) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "maskPointChin":
		t.MaskPointChin = new(MaskPointChin)
		return json.Unmarshal(b, t.MaskPointChin)
	case "maskPointEyes":
		t.MaskPointEyes = new(MaskPointEyes)
		return json.Unmarshal(b, t.MaskPointEyes)
	case "maskPointForehead":
		t.MaskPointForehead = new(MaskPointForehead)
		return json.Unmarshal(b, t.MaskPointForehead)
	case "maskPointMouth":
		t.MaskPointMouth = new(MaskPointMouth)
		return json.Unmarshal(b, t.MaskPointMouth)
	}
	return nil
}

func (t *MaskPoint) MarshalJSON() ([]byte, error) {
	if t.MaskPointChin != nil {
		return json.Marshal(t.MaskPointChin)
	}
	if t.MaskPointEyes != nil {
		return json.Marshal(t.MaskPointEyes)
	}
	if t.MaskPointForehead != nil {
		return json.Marshal(t.MaskPointForehead)
	}
	if t.MaskPointMouth != nil {
		return json.Marshal(t.MaskPointMouth)
	}
	return nil, fmt.Errorf("no value set for MaskPoint")
}

// MessageContent Contains the content of a message
type MessageContent struct {
	TypeStr                                  string                                    `json:"@type"`
	MessageAnimatedEmoji                     *MessageAnimatedEmoji                     `json:"messageAnimatedEmoji,omitempty"`
	MessageAnimation                         *MessageAnimation                         `json:"messageAnimation,omitempty"`
	MessageAudio                             *MessageAudio                             `json:"messageAudio,omitempty"`
	MessageBasicGroupChatCreate              *MessageBasicGroupChatCreate              `json:"messageBasicGroupChatCreate,omitempty"`
	MessageBotWriteAccessAllowed             *MessageBotWriteAccessAllowed             `json:"messageBotWriteAccessAllowed,omitempty"`
	MessageCall                              *MessageCall                              `json:"messageCall,omitempty"`
	MessageChatAddMembers                    *MessageChatAddMembers                    `json:"messageChatAddMembers,omitempty"`
	MessageChatBoost                         *MessageChatBoost                         `json:"messageChatBoost,omitempty"`
	MessageChatChangePhoto                   *MessageChatChangePhoto                   `json:"messageChatChangePhoto,omitempty"`
	MessageChatChangeTitle                   *MessageChatChangeTitle                   `json:"messageChatChangeTitle,omitempty"`
	MessageChatDeleteMember                  *MessageChatDeleteMember                  `json:"messageChatDeleteMember,omitempty"`
	MessageChatDeletePhoto                   *MessageChatDeletePhoto                   `json:"messageChatDeletePhoto,omitempty"`
	MessageChatJoinByLink                    *MessageChatJoinByLink                    `json:"messageChatJoinByLink,omitempty"`
	MessageChatJoinByRequest                 *MessageChatJoinByRequest                 `json:"messageChatJoinByRequest,omitempty"`
	MessageChatSetBackground                 *MessageChatSetBackground                 `json:"messageChatSetBackground,omitempty"`
	MessageChatSetMessageAutoDeleteTime      *MessageChatSetMessageAutoDeleteTime      `json:"messageChatSetMessageAutoDeleteTime,omitempty"`
	MessageChatSetTheme                      *MessageChatSetTheme                      `json:"messageChatSetTheme,omitempty"`
	MessageChatShared                        *MessageChatShared                        `json:"messageChatShared,omitempty"`
	MessageChatUpgradeFrom                   *MessageChatUpgradeFrom                   `json:"messageChatUpgradeFrom,omitempty"`
	MessageChatUpgradeTo                     *MessageChatUpgradeTo                     `json:"messageChatUpgradeTo,omitempty"`
	MessageChecklist                         *MessageChecklist                         `json:"messageChecklist,omitempty"`
	MessageChecklistTasksAdded               *MessageChecklistTasksAdded               `json:"messageChecklistTasksAdded,omitempty"`
	MessageChecklistTasksDone                *MessageChecklistTasksDone                `json:"messageChecklistTasksDone,omitempty"`
	MessageContact                           *MessageContact                           `json:"messageContact,omitempty"`
	MessageContactRegistered                 *MessageContactRegistered                 `json:"messageContactRegistered,omitempty"`
	MessageCustomServiceAction               *MessageCustomServiceAction               `json:"messageCustomServiceAction,omitempty"`
	MessageDice                              *MessageDice                              `json:"messageDice,omitempty"`
	MessageDirectMessagePriceChanged         *MessageDirectMessagePriceChanged         `json:"messageDirectMessagePriceChanged,omitempty"`
	MessageDocument                          *MessageDocument                          `json:"messageDocument,omitempty"`
	MessageExpiredPhoto                      *MessageExpiredPhoto                      `json:"messageExpiredPhoto,omitempty"`
	MessageExpiredVideo                      *MessageExpiredVideo                      `json:"messageExpiredVideo,omitempty"`
	MessageExpiredVideoNote                  *MessageExpiredVideoNote                  `json:"messageExpiredVideoNote,omitempty"`
	MessageExpiredVoiceNote                  *MessageExpiredVoiceNote                  `json:"messageExpiredVoiceNote,omitempty"`
	MessageForumTopicCreated                 *MessageForumTopicCreated                 `json:"messageForumTopicCreated,omitempty"`
	MessageForumTopicEdited                  *MessageForumTopicEdited                  `json:"messageForumTopicEdited,omitempty"`
	MessageForumTopicIsClosedToggled         *MessageForumTopicIsClosedToggled         `json:"messageForumTopicIsClosedToggled,omitempty"`
	MessageForumTopicIsHiddenToggled         *MessageForumTopicIsHiddenToggled         `json:"messageForumTopicIsHiddenToggled,omitempty"`
	MessageGame                              *MessageGame                              `json:"messageGame,omitempty"`
	MessageGameScore                         *MessageGameScore                         `json:"messageGameScore,omitempty"`
	MessageGift                              *MessageGift                              `json:"messageGift,omitempty"`
	MessageGiftedPremium                     *MessageGiftedPremium                     `json:"messageGiftedPremium,omitempty"`
	MessageGiftedStars                       *MessageGiftedStars                       `json:"messageGiftedStars,omitempty"`
	MessageGiftedTon                         *MessageGiftedTon                         `json:"messageGiftedTon,omitempty"`
	MessageGiveaway                          *MessageGiveaway                          `json:"messageGiveaway,omitempty"`
	MessageGiveawayCompleted                 *MessageGiveawayCompleted                 `json:"messageGiveawayCompleted,omitempty"`
	MessageGiveawayCreated                   *MessageGiveawayCreated                   `json:"messageGiveawayCreated,omitempty"`
	MessageGiveawayPrizeStars                *MessageGiveawayPrizeStars                `json:"messageGiveawayPrizeStars,omitempty"`
	MessageGiveawayWinners                   *MessageGiveawayWinners                   `json:"messageGiveawayWinners,omitempty"`
	MessageGroupCall                         *MessageGroupCall                         `json:"messageGroupCall,omitempty"`
	MessageInviteVideoChatParticipants       *MessageInviteVideoChatParticipants       `json:"messageInviteVideoChatParticipants,omitempty"`
	MessageInvoice                           *MessageInvoice                           `json:"messageInvoice,omitempty"`
	MessageLocation                          *MessageLocation                          `json:"messageLocation,omitempty"`
	MessagePaidMedia                         *MessagePaidMedia                         `json:"messagePaidMedia,omitempty"`
	MessagePaidMessagePriceChanged           *MessagePaidMessagePriceChanged           `json:"messagePaidMessagePriceChanged,omitempty"`
	MessagePaidMessagesRefunded              *MessagePaidMessagesRefunded              `json:"messagePaidMessagesRefunded,omitempty"`
	MessagePassportDataReceived              *MessagePassportDataReceived              `json:"messagePassportDataReceived,omitempty"`
	MessagePassportDataSent                  *MessagePassportDataSent                  `json:"messagePassportDataSent,omitempty"`
	MessagePaymentRefunded                   *MessagePaymentRefunded                   `json:"messagePaymentRefunded,omitempty"`
	MessagePaymentSuccessful                 *MessagePaymentSuccessful                 `json:"messagePaymentSuccessful,omitempty"`
	MessagePaymentSuccessfulBot              *MessagePaymentSuccessfulBot              `json:"messagePaymentSuccessfulBot,omitempty"`
	MessagePhoto                             *MessagePhoto                             `json:"messagePhoto,omitempty"`
	MessagePinMessage                        *MessagePinMessage                        `json:"messagePinMessage,omitempty"`
	MessagePoll                              *MessagePoll                              `json:"messagePoll,omitempty"`
	MessagePremiumGiftCode                   *MessagePremiumGiftCode                   `json:"messagePremiumGiftCode,omitempty"`
	MessageProximityAlertTriggered           *MessageProximityAlertTriggered           `json:"messageProximityAlertTriggered,omitempty"`
	MessageRefundedUpgradedGift              *MessageRefundedUpgradedGift              `json:"messageRefundedUpgradedGift,omitempty"`
	MessageScreenshotTaken                   *MessageScreenshotTaken                   `json:"messageScreenshotTaken,omitempty"`
	MessageStakeDice                         *MessageStakeDice                         `json:"messageStakeDice,omitempty"`
	MessageSticker                           *MessageSticker                           `json:"messageSticker,omitempty"`
	MessageStory                             *MessageStory                             `json:"messageStory,omitempty"`
	MessageSuggestBirthdate                  *MessageSuggestBirthdate                  `json:"messageSuggestBirthdate,omitempty"`
	MessageSuggestProfilePhoto               *MessageSuggestProfilePhoto               `json:"messageSuggestProfilePhoto,omitempty"`
	MessageSuggestedPostApprovalFailed       *MessageSuggestedPostApprovalFailed       `json:"messageSuggestedPostApprovalFailed,omitempty"`
	MessageSuggestedPostApproved             *MessageSuggestedPostApproved             `json:"messageSuggestedPostApproved,omitempty"`
	MessageSuggestedPostDeclined             *MessageSuggestedPostDeclined             `json:"messageSuggestedPostDeclined,omitempty"`
	MessageSuggestedPostPaid                 *MessageSuggestedPostPaid                 `json:"messageSuggestedPostPaid,omitempty"`
	MessageSuggestedPostRefunded             *MessageSuggestedPostRefunded             `json:"messageSuggestedPostRefunded,omitempty"`
	MessageSupergroupChatCreate              *MessageSupergroupChatCreate              `json:"messageSupergroupChatCreate,omitempty"`
	MessageText                              *MessageText                              `json:"messageText,omitempty"`
	MessageUnsupported                       *MessageUnsupported                       `json:"messageUnsupported,omitempty"`
	MessageUpgradedGift                      *MessageUpgradedGift                      `json:"messageUpgradedGift,omitempty"`
	MessageUpgradedGiftPurchaseOffer         *MessageUpgradedGiftPurchaseOffer         `json:"messageUpgradedGiftPurchaseOffer,omitempty"`
	MessageUpgradedGiftPurchaseOfferRejected *MessageUpgradedGiftPurchaseOfferRejected `json:"messageUpgradedGiftPurchaseOfferRejected,omitempty"`
	MessageUsersShared                       *MessageUsersShared                       `json:"messageUsersShared,omitempty"`
	MessageVenue                             *MessageVenue                             `json:"messageVenue,omitempty"`
	MessageVideo                             *MessageVideo                             `json:"messageVideo,omitempty"`
	MessageVideoChatEnded                    *MessageVideoChatEnded                    `json:"messageVideoChatEnded,omitempty"`
	MessageVideoChatScheduled                *MessageVideoChatScheduled                `json:"messageVideoChatScheduled,omitempty"`
	MessageVideoChatStarted                  *MessageVideoChatStarted                  `json:"messageVideoChatStarted,omitempty"`
	MessageVideoNote                         *MessageVideoNote                         `json:"messageVideoNote,omitempty"`
	MessageVoiceNote                         *MessageVoiceNote                         `json:"messageVoiceNote,omitempty"`
	MessageWebAppDataReceived                *MessageWebAppDataReceived                `json:"messageWebAppDataReceived,omitempty"`
	MessageWebAppDataSent                    *MessageWebAppDataSent                    `json:"messageWebAppDataSent,omitempty"`
}

func (t *MessageContent) Type() string {
	return t.TypeStr
}

func (t *MessageContent) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageContent) GetExtra() string {
	return ""
}

func (t *MessageContent) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageAnimatedEmoji":
		t.MessageAnimatedEmoji = new(MessageAnimatedEmoji)
		return json.Unmarshal(b, t.MessageAnimatedEmoji)
	case "messageAnimation":
		t.MessageAnimation = new(MessageAnimation)
		return json.Unmarshal(b, t.MessageAnimation)
	case "messageAudio":
		t.MessageAudio = new(MessageAudio)
		return json.Unmarshal(b, t.MessageAudio)
	case "messageBasicGroupChatCreate":
		t.MessageBasicGroupChatCreate = new(MessageBasicGroupChatCreate)
		return json.Unmarshal(b, t.MessageBasicGroupChatCreate)
	case "messageBotWriteAccessAllowed":
		t.MessageBotWriteAccessAllowed = new(MessageBotWriteAccessAllowed)
		return json.Unmarshal(b, t.MessageBotWriteAccessAllowed)
	case "messageCall":
		t.MessageCall = new(MessageCall)
		return json.Unmarshal(b, t.MessageCall)
	case "messageChatAddMembers":
		t.MessageChatAddMembers = new(MessageChatAddMembers)
		return json.Unmarshal(b, t.MessageChatAddMembers)
	case "messageChatBoost":
		t.MessageChatBoost = new(MessageChatBoost)
		return json.Unmarshal(b, t.MessageChatBoost)
	case "messageChatChangePhoto":
		t.MessageChatChangePhoto = new(MessageChatChangePhoto)
		return json.Unmarshal(b, t.MessageChatChangePhoto)
	case "messageChatChangeTitle":
		t.MessageChatChangeTitle = new(MessageChatChangeTitle)
		return json.Unmarshal(b, t.MessageChatChangeTitle)
	case "messageChatDeleteMember":
		t.MessageChatDeleteMember = new(MessageChatDeleteMember)
		return json.Unmarshal(b, t.MessageChatDeleteMember)
	case "messageChatDeletePhoto":
		t.MessageChatDeletePhoto = new(MessageChatDeletePhoto)
		return json.Unmarshal(b, t.MessageChatDeletePhoto)
	case "messageChatJoinByLink":
		t.MessageChatJoinByLink = new(MessageChatJoinByLink)
		return json.Unmarshal(b, t.MessageChatJoinByLink)
	case "messageChatJoinByRequest":
		t.MessageChatJoinByRequest = new(MessageChatJoinByRequest)
		return json.Unmarshal(b, t.MessageChatJoinByRequest)
	case "messageChatSetBackground":
		t.MessageChatSetBackground = new(MessageChatSetBackground)
		return json.Unmarshal(b, t.MessageChatSetBackground)
	case "messageChatSetMessageAutoDeleteTime":
		t.MessageChatSetMessageAutoDeleteTime = new(MessageChatSetMessageAutoDeleteTime)
		return json.Unmarshal(b, t.MessageChatSetMessageAutoDeleteTime)
	case "messageChatSetTheme":
		t.MessageChatSetTheme = new(MessageChatSetTheme)
		return json.Unmarshal(b, t.MessageChatSetTheme)
	case "messageChatShared":
		t.MessageChatShared = new(MessageChatShared)
		return json.Unmarshal(b, t.MessageChatShared)
	case "messageChatUpgradeFrom":
		t.MessageChatUpgradeFrom = new(MessageChatUpgradeFrom)
		return json.Unmarshal(b, t.MessageChatUpgradeFrom)
	case "messageChatUpgradeTo":
		t.MessageChatUpgradeTo = new(MessageChatUpgradeTo)
		return json.Unmarshal(b, t.MessageChatUpgradeTo)
	case "messageChecklist":
		t.MessageChecklist = new(MessageChecklist)
		return json.Unmarshal(b, t.MessageChecklist)
	case "messageChecklistTasksAdded":
		t.MessageChecklistTasksAdded = new(MessageChecklistTasksAdded)
		return json.Unmarshal(b, t.MessageChecklistTasksAdded)
	case "messageChecklistTasksDone":
		t.MessageChecklistTasksDone = new(MessageChecklistTasksDone)
		return json.Unmarshal(b, t.MessageChecklistTasksDone)
	case "messageContact":
		t.MessageContact = new(MessageContact)
		return json.Unmarshal(b, t.MessageContact)
	case "messageContactRegistered":
		t.MessageContactRegistered = new(MessageContactRegistered)
		return json.Unmarshal(b, t.MessageContactRegistered)
	case "messageCustomServiceAction":
		t.MessageCustomServiceAction = new(MessageCustomServiceAction)
		return json.Unmarshal(b, t.MessageCustomServiceAction)
	case "messageDice":
		t.MessageDice = new(MessageDice)
		return json.Unmarshal(b, t.MessageDice)
	case "messageDirectMessagePriceChanged":
		t.MessageDirectMessagePriceChanged = new(MessageDirectMessagePriceChanged)
		return json.Unmarshal(b, t.MessageDirectMessagePriceChanged)
	case "messageDocument":
		t.MessageDocument = new(MessageDocument)
		return json.Unmarshal(b, t.MessageDocument)
	case "messageExpiredPhoto":
		t.MessageExpiredPhoto = new(MessageExpiredPhoto)
		return json.Unmarshal(b, t.MessageExpiredPhoto)
	case "messageExpiredVideo":
		t.MessageExpiredVideo = new(MessageExpiredVideo)
		return json.Unmarshal(b, t.MessageExpiredVideo)
	case "messageExpiredVideoNote":
		t.MessageExpiredVideoNote = new(MessageExpiredVideoNote)
		return json.Unmarshal(b, t.MessageExpiredVideoNote)
	case "messageExpiredVoiceNote":
		t.MessageExpiredVoiceNote = new(MessageExpiredVoiceNote)
		return json.Unmarshal(b, t.MessageExpiredVoiceNote)
	case "messageForumTopicCreated":
		t.MessageForumTopicCreated = new(MessageForumTopicCreated)
		return json.Unmarshal(b, t.MessageForumTopicCreated)
	case "messageForumTopicEdited":
		t.MessageForumTopicEdited = new(MessageForumTopicEdited)
		return json.Unmarshal(b, t.MessageForumTopicEdited)
	case "messageForumTopicIsClosedToggled":
		t.MessageForumTopicIsClosedToggled = new(MessageForumTopicIsClosedToggled)
		return json.Unmarshal(b, t.MessageForumTopicIsClosedToggled)
	case "messageForumTopicIsHiddenToggled":
		t.MessageForumTopicIsHiddenToggled = new(MessageForumTopicIsHiddenToggled)
		return json.Unmarshal(b, t.MessageForumTopicIsHiddenToggled)
	case "messageGame":
		t.MessageGame = new(MessageGame)
		return json.Unmarshal(b, t.MessageGame)
	case "messageGameScore":
		t.MessageGameScore = new(MessageGameScore)
		return json.Unmarshal(b, t.MessageGameScore)
	case "messageGift":
		t.MessageGift = new(MessageGift)
		return json.Unmarshal(b, t.MessageGift)
	case "messageGiftedPremium":
		t.MessageGiftedPremium = new(MessageGiftedPremium)
		return json.Unmarshal(b, t.MessageGiftedPremium)
	case "messageGiftedStars":
		t.MessageGiftedStars = new(MessageGiftedStars)
		return json.Unmarshal(b, t.MessageGiftedStars)
	case "messageGiftedTon":
		t.MessageGiftedTon = new(MessageGiftedTon)
		return json.Unmarshal(b, t.MessageGiftedTon)
	case "messageGiveaway":
		t.MessageGiveaway = new(MessageGiveaway)
		return json.Unmarshal(b, t.MessageGiveaway)
	case "messageGiveawayCompleted":
		t.MessageGiveawayCompleted = new(MessageGiveawayCompleted)
		return json.Unmarshal(b, t.MessageGiveawayCompleted)
	case "messageGiveawayCreated":
		t.MessageGiveawayCreated = new(MessageGiveawayCreated)
		return json.Unmarshal(b, t.MessageGiveawayCreated)
	case "messageGiveawayPrizeStars":
		t.MessageGiveawayPrizeStars = new(MessageGiveawayPrizeStars)
		return json.Unmarshal(b, t.MessageGiveawayPrizeStars)
	case "messageGiveawayWinners":
		t.MessageGiveawayWinners = new(MessageGiveawayWinners)
		return json.Unmarshal(b, t.MessageGiveawayWinners)
	case "messageGroupCall":
		t.MessageGroupCall = new(MessageGroupCall)
		return json.Unmarshal(b, t.MessageGroupCall)
	case "messageInviteVideoChatParticipants":
		t.MessageInviteVideoChatParticipants = new(MessageInviteVideoChatParticipants)
		return json.Unmarshal(b, t.MessageInviteVideoChatParticipants)
	case "messageInvoice":
		t.MessageInvoice = new(MessageInvoice)
		return json.Unmarshal(b, t.MessageInvoice)
	case "messageLocation":
		t.MessageLocation = new(MessageLocation)
		return json.Unmarshal(b, t.MessageLocation)
	case "messagePaidMedia":
		t.MessagePaidMedia = new(MessagePaidMedia)
		return json.Unmarshal(b, t.MessagePaidMedia)
	case "messagePaidMessagePriceChanged":
		t.MessagePaidMessagePriceChanged = new(MessagePaidMessagePriceChanged)
		return json.Unmarshal(b, t.MessagePaidMessagePriceChanged)
	case "messagePaidMessagesRefunded":
		t.MessagePaidMessagesRefunded = new(MessagePaidMessagesRefunded)
		return json.Unmarshal(b, t.MessagePaidMessagesRefunded)
	case "messagePassportDataReceived":
		t.MessagePassportDataReceived = new(MessagePassportDataReceived)
		return json.Unmarshal(b, t.MessagePassportDataReceived)
	case "messagePassportDataSent":
		t.MessagePassportDataSent = new(MessagePassportDataSent)
		return json.Unmarshal(b, t.MessagePassportDataSent)
	case "messagePaymentRefunded":
		t.MessagePaymentRefunded = new(MessagePaymentRefunded)
		return json.Unmarshal(b, t.MessagePaymentRefunded)
	case "messagePaymentSuccessful":
		t.MessagePaymentSuccessful = new(MessagePaymentSuccessful)
		return json.Unmarshal(b, t.MessagePaymentSuccessful)
	case "messagePaymentSuccessfulBot":
		t.MessagePaymentSuccessfulBot = new(MessagePaymentSuccessfulBot)
		return json.Unmarshal(b, t.MessagePaymentSuccessfulBot)
	case "messagePhoto":
		t.MessagePhoto = new(MessagePhoto)
		return json.Unmarshal(b, t.MessagePhoto)
	case "messagePinMessage":
		t.MessagePinMessage = new(MessagePinMessage)
		return json.Unmarshal(b, t.MessagePinMessage)
	case "messagePoll":
		t.MessagePoll = new(MessagePoll)
		return json.Unmarshal(b, t.MessagePoll)
	case "messagePremiumGiftCode":
		t.MessagePremiumGiftCode = new(MessagePremiumGiftCode)
		return json.Unmarshal(b, t.MessagePremiumGiftCode)
	case "messageProximityAlertTriggered":
		t.MessageProximityAlertTriggered = new(MessageProximityAlertTriggered)
		return json.Unmarshal(b, t.MessageProximityAlertTriggered)
	case "messageRefundedUpgradedGift":
		t.MessageRefundedUpgradedGift = new(MessageRefundedUpgradedGift)
		return json.Unmarshal(b, t.MessageRefundedUpgradedGift)
	case "messageScreenshotTaken":
		t.MessageScreenshotTaken = new(MessageScreenshotTaken)
		return json.Unmarshal(b, t.MessageScreenshotTaken)
	case "messageStakeDice":
		t.MessageStakeDice = new(MessageStakeDice)
		return json.Unmarshal(b, t.MessageStakeDice)
	case "messageSticker":
		t.MessageSticker = new(MessageSticker)
		return json.Unmarshal(b, t.MessageSticker)
	case "messageStory":
		t.MessageStory = new(MessageStory)
		return json.Unmarshal(b, t.MessageStory)
	case "messageSuggestBirthdate":
		t.MessageSuggestBirthdate = new(MessageSuggestBirthdate)
		return json.Unmarshal(b, t.MessageSuggestBirthdate)
	case "messageSuggestProfilePhoto":
		t.MessageSuggestProfilePhoto = new(MessageSuggestProfilePhoto)
		return json.Unmarshal(b, t.MessageSuggestProfilePhoto)
	case "messageSuggestedPostApprovalFailed":
		t.MessageSuggestedPostApprovalFailed = new(MessageSuggestedPostApprovalFailed)
		return json.Unmarshal(b, t.MessageSuggestedPostApprovalFailed)
	case "messageSuggestedPostApproved":
		t.MessageSuggestedPostApproved = new(MessageSuggestedPostApproved)
		return json.Unmarshal(b, t.MessageSuggestedPostApproved)
	case "messageSuggestedPostDeclined":
		t.MessageSuggestedPostDeclined = new(MessageSuggestedPostDeclined)
		return json.Unmarshal(b, t.MessageSuggestedPostDeclined)
	case "messageSuggestedPostPaid":
		t.MessageSuggestedPostPaid = new(MessageSuggestedPostPaid)
		return json.Unmarshal(b, t.MessageSuggestedPostPaid)
	case "messageSuggestedPostRefunded":
		t.MessageSuggestedPostRefunded = new(MessageSuggestedPostRefunded)
		return json.Unmarshal(b, t.MessageSuggestedPostRefunded)
	case "messageSupergroupChatCreate":
		t.MessageSupergroupChatCreate = new(MessageSupergroupChatCreate)
		return json.Unmarshal(b, t.MessageSupergroupChatCreate)
	case "messageText":
		t.MessageText = new(MessageText)
		return json.Unmarshal(b, t.MessageText)
	case "messageUnsupported":
		t.MessageUnsupported = new(MessageUnsupported)
		return json.Unmarshal(b, t.MessageUnsupported)
	case "messageUpgradedGift":
		t.MessageUpgradedGift = new(MessageUpgradedGift)
		return json.Unmarshal(b, t.MessageUpgradedGift)
	case "messageUpgradedGiftPurchaseOffer":
		t.MessageUpgradedGiftPurchaseOffer = new(MessageUpgradedGiftPurchaseOffer)
		return json.Unmarshal(b, t.MessageUpgradedGiftPurchaseOffer)
	case "messageUpgradedGiftPurchaseOfferRejected":
		t.MessageUpgradedGiftPurchaseOfferRejected = new(MessageUpgradedGiftPurchaseOfferRejected)
		return json.Unmarshal(b, t.MessageUpgradedGiftPurchaseOfferRejected)
	case "messageUsersShared":
		t.MessageUsersShared = new(MessageUsersShared)
		return json.Unmarshal(b, t.MessageUsersShared)
	case "messageVenue":
		t.MessageVenue = new(MessageVenue)
		return json.Unmarshal(b, t.MessageVenue)
	case "messageVideo":
		t.MessageVideo = new(MessageVideo)
		return json.Unmarshal(b, t.MessageVideo)
	case "messageVideoChatEnded":
		t.MessageVideoChatEnded = new(MessageVideoChatEnded)
		return json.Unmarshal(b, t.MessageVideoChatEnded)
	case "messageVideoChatScheduled":
		t.MessageVideoChatScheduled = new(MessageVideoChatScheduled)
		return json.Unmarshal(b, t.MessageVideoChatScheduled)
	case "messageVideoChatStarted":
		t.MessageVideoChatStarted = new(MessageVideoChatStarted)
		return json.Unmarshal(b, t.MessageVideoChatStarted)
	case "messageVideoNote":
		t.MessageVideoNote = new(MessageVideoNote)
		return json.Unmarshal(b, t.MessageVideoNote)
	case "messageVoiceNote":
		t.MessageVoiceNote = new(MessageVoiceNote)
		return json.Unmarshal(b, t.MessageVoiceNote)
	case "messageWebAppDataReceived":
		t.MessageWebAppDataReceived = new(MessageWebAppDataReceived)
		return json.Unmarshal(b, t.MessageWebAppDataReceived)
	case "messageWebAppDataSent":
		t.MessageWebAppDataSent = new(MessageWebAppDataSent)
		return json.Unmarshal(b, t.MessageWebAppDataSent)
	}
	return nil
}

func (t *MessageContent) MarshalJSON() ([]byte, error) {
	if t.MessageAnimatedEmoji != nil {
		return json.Marshal(t.MessageAnimatedEmoji)
	}
	if t.MessageAnimation != nil {
		return json.Marshal(t.MessageAnimation)
	}
	if t.MessageAudio != nil {
		return json.Marshal(t.MessageAudio)
	}
	if t.MessageBasicGroupChatCreate != nil {
		return json.Marshal(t.MessageBasicGroupChatCreate)
	}
	if t.MessageBotWriteAccessAllowed != nil {
		return json.Marshal(t.MessageBotWriteAccessAllowed)
	}
	if t.MessageCall != nil {
		return json.Marshal(t.MessageCall)
	}
	if t.MessageChatAddMembers != nil {
		return json.Marshal(t.MessageChatAddMembers)
	}
	if t.MessageChatBoost != nil {
		return json.Marshal(t.MessageChatBoost)
	}
	if t.MessageChatChangePhoto != nil {
		return json.Marshal(t.MessageChatChangePhoto)
	}
	if t.MessageChatChangeTitle != nil {
		return json.Marshal(t.MessageChatChangeTitle)
	}
	if t.MessageChatDeleteMember != nil {
		return json.Marshal(t.MessageChatDeleteMember)
	}
	if t.MessageChatDeletePhoto != nil {
		return json.Marshal(t.MessageChatDeletePhoto)
	}
	if t.MessageChatJoinByLink != nil {
		return json.Marshal(t.MessageChatJoinByLink)
	}
	if t.MessageChatJoinByRequest != nil {
		return json.Marshal(t.MessageChatJoinByRequest)
	}
	if t.MessageChatSetBackground != nil {
		return json.Marshal(t.MessageChatSetBackground)
	}
	if t.MessageChatSetMessageAutoDeleteTime != nil {
		return json.Marshal(t.MessageChatSetMessageAutoDeleteTime)
	}
	if t.MessageChatSetTheme != nil {
		return json.Marshal(t.MessageChatSetTheme)
	}
	if t.MessageChatShared != nil {
		return json.Marshal(t.MessageChatShared)
	}
	if t.MessageChatUpgradeFrom != nil {
		return json.Marshal(t.MessageChatUpgradeFrom)
	}
	if t.MessageChatUpgradeTo != nil {
		return json.Marshal(t.MessageChatUpgradeTo)
	}
	if t.MessageChecklist != nil {
		return json.Marshal(t.MessageChecklist)
	}
	if t.MessageChecklistTasksAdded != nil {
		return json.Marshal(t.MessageChecklistTasksAdded)
	}
	if t.MessageChecklistTasksDone != nil {
		return json.Marshal(t.MessageChecklistTasksDone)
	}
	if t.MessageContact != nil {
		return json.Marshal(t.MessageContact)
	}
	if t.MessageContactRegistered != nil {
		return json.Marshal(t.MessageContactRegistered)
	}
	if t.MessageCustomServiceAction != nil {
		return json.Marshal(t.MessageCustomServiceAction)
	}
	if t.MessageDice != nil {
		return json.Marshal(t.MessageDice)
	}
	if t.MessageDirectMessagePriceChanged != nil {
		return json.Marshal(t.MessageDirectMessagePriceChanged)
	}
	if t.MessageDocument != nil {
		return json.Marshal(t.MessageDocument)
	}
	if t.MessageExpiredPhoto != nil {
		return json.Marshal(t.MessageExpiredPhoto)
	}
	if t.MessageExpiredVideo != nil {
		return json.Marshal(t.MessageExpiredVideo)
	}
	if t.MessageExpiredVideoNote != nil {
		return json.Marshal(t.MessageExpiredVideoNote)
	}
	if t.MessageExpiredVoiceNote != nil {
		return json.Marshal(t.MessageExpiredVoiceNote)
	}
	if t.MessageForumTopicCreated != nil {
		return json.Marshal(t.MessageForumTopicCreated)
	}
	if t.MessageForumTopicEdited != nil {
		return json.Marshal(t.MessageForumTopicEdited)
	}
	if t.MessageForumTopicIsClosedToggled != nil {
		return json.Marshal(t.MessageForumTopicIsClosedToggled)
	}
	if t.MessageForumTopicIsHiddenToggled != nil {
		return json.Marshal(t.MessageForumTopicIsHiddenToggled)
	}
	if t.MessageGame != nil {
		return json.Marshal(t.MessageGame)
	}
	if t.MessageGameScore != nil {
		return json.Marshal(t.MessageGameScore)
	}
	if t.MessageGift != nil {
		return json.Marshal(t.MessageGift)
	}
	if t.MessageGiftedPremium != nil {
		return json.Marshal(t.MessageGiftedPremium)
	}
	if t.MessageGiftedStars != nil {
		return json.Marshal(t.MessageGiftedStars)
	}
	if t.MessageGiftedTon != nil {
		return json.Marshal(t.MessageGiftedTon)
	}
	if t.MessageGiveaway != nil {
		return json.Marshal(t.MessageGiveaway)
	}
	if t.MessageGiveawayCompleted != nil {
		return json.Marshal(t.MessageGiveawayCompleted)
	}
	if t.MessageGiveawayCreated != nil {
		return json.Marshal(t.MessageGiveawayCreated)
	}
	if t.MessageGiveawayPrizeStars != nil {
		return json.Marshal(t.MessageGiveawayPrizeStars)
	}
	if t.MessageGiveawayWinners != nil {
		return json.Marshal(t.MessageGiveawayWinners)
	}
	if t.MessageGroupCall != nil {
		return json.Marshal(t.MessageGroupCall)
	}
	if t.MessageInviteVideoChatParticipants != nil {
		return json.Marshal(t.MessageInviteVideoChatParticipants)
	}
	if t.MessageInvoice != nil {
		return json.Marshal(t.MessageInvoice)
	}
	if t.MessageLocation != nil {
		return json.Marshal(t.MessageLocation)
	}
	if t.MessagePaidMedia != nil {
		return json.Marshal(t.MessagePaidMedia)
	}
	if t.MessagePaidMessagePriceChanged != nil {
		return json.Marshal(t.MessagePaidMessagePriceChanged)
	}
	if t.MessagePaidMessagesRefunded != nil {
		return json.Marshal(t.MessagePaidMessagesRefunded)
	}
	if t.MessagePassportDataReceived != nil {
		return json.Marshal(t.MessagePassportDataReceived)
	}
	if t.MessagePassportDataSent != nil {
		return json.Marshal(t.MessagePassportDataSent)
	}
	if t.MessagePaymentRefunded != nil {
		return json.Marshal(t.MessagePaymentRefunded)
	}
	if t.MessagePaymentSuccessful != nil {
		return json.Marshal(t.MessagePaymentSuccessful)
	}
	if t.MessagePaymentSuccessfulBot != nil {
		return json.Marshal(t.MessagePaymentSuccessfulBot)
	}
	if t.MessagePhoto != nil {
		return json.Marshal(t.MessagePhoto)
	}
	if t.MessagePinMessage != nil {
		return json.Marshal(t.MessagePinMessage)
	}
	if t.MessagePoll != nil {
		return json.Marshal(t.MessagePoll)
	}
	if t.MessagePremiumGiftCode != nil {
		return json.Marshal(t.MessagePremiumGiftCode)
	}
	if t.MessageProximityAlertTriggered != nil {
		return json.Marshal(t.MessageProximityAlertTriggered)
	}
	if t.MessageRefundedUpgradedGift != nil {
		return json.Marshal(t.MessageRefundedUpgradedGift)
	}
	if t.MessageScreenshotTaken != nil {
		return json.Marshal(t.MessageScreenshotTaken)
	}
	if t.MessageStakeDice != nil {
		return json.Marshal(t.MessageStakeDice)
	}
	if t.MessageSticker != nil {
		return json.Marshal(t.MessageSticker)
	}
	if t.MessageStory != nil {
		return json.Marshal(t.MessageStory)
	}
	if t.MessageSuggestBirthdate != nil {
		return json.Marshal(t.MessageSuggestBirthdate)
	}
	if t.MessageSuggestProfilePhoto != nil {
		return json.Marshal(t.MessageSuggestProfilePhoto)
	}
	if t.MessageSuggestedPostApprovalFailed != nil {
		return json.Marshal(t.MessageSuggestedPostApprovalFailed)
	}
	if t.MessageSuggestedPostApproved != nil {
		return json.Marshal(t.MessageSuggestedPostApproved)
	}
	if t.MessageSuggestedPostDeclined != nil {
		return json.Marshal(t.MessageSuggestedPostDeclined)
	}
	if t.MessageSuggestedPostPaid != nil {
		return json.Marshal(t.MessageSuggestedPostPaid)
	}
	if t.MessageSuggestedPostRefunded != nil {
		return json.Marshal(t.MessageSuggestedPostRefunded)
	}
	if t.MessageSupergroupChatCreate != nil {
		return json.Marshal(t.MessageSupergroupChatCreate)
	}
	if t.MessageText != nil {
		return json.Marshal(t.MessageText)
	}
	if t.MessageUnsupported != nil {
		return json.Marshal(t.MessageUnsupported)
	}
	if t.MessageUpgradedGift != nil {
		return json.Marshal(t.MessageUpgradedGift)
	}
	if t.MessageUpgradedGiftPurchaseOffer != nil {
		return json.Marshal(t.MessageUpgradedGiftPurchaseOffer)
	}
	if t.MessageUpgradedGiftPurchaseOfferRejected != nil {
		return json.Marshal(t.MessageUpgradedGiftPurchaseOfferRejected)
	}
	if t.MessageUsersShared != nil {
		return json.Marshal(t.MessageUsersShared)
	}
	if t.MessageVenue != nil {
		return json.Marshal(t.MessageVenue)
	}
	if t.MessageVideo != nil {
		return json.Marshal(t.MessageVideo)
	}
	if t.MessageVideoChatEnded != nil {
		return json.Marshal(t.MessageVideoChatEnded)
	}
	if t.MessageVideoChatScheduled != nil {
		return json.Marshal(t.MessageVideoChatScheduled)
	}
	if t.MessageVideoChatStarted != nil {
		return json.Marshal(t.MessageVideoChatStarted)
	}
	if t.MessageVideoNote != nil {
		return json.Marshal(t.MessageVideoNote)
	}
	if t.MessageVoiceNote != nil {
		return json.Marshal(t.MessageVoiceNote)
	}
	if t.MessageWebAppDataReceived != nil {
		return json.Marshal(t.MessageWebAppDataReceived)
	}
	if t.MessageWebAppDataSent != nil {
		return json.Marshal(t.MessageWebAppDataSent)
	}
	return nil, fmt.Errorf("no value set for MessageContent")
}

// MessageEffectType Describes type of emoji effect
type MessageEffectType struct {
	TypeStr                         string                           `json:"@type"`
	MessageEffectTypeEmojiReaction  *MessageEffectTypeEmojiReaction  `json:"messageEffectTypeEmojiReaction,omitempty"`
	MessageEffectTypePremiumSticker *MessageEffectTypePremiumSticker `json:"messageEffectTypePremiumSticker,omitempty"`
}

func (t *MessageEffectType) Type() string {
	return t.TypeStr
}

func (t *MessageEffectType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageEffectType) GetExtra() string {
	return ""
}

func (t *MessageEffectType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageEffectTypeEmojiReaction":
		t.MessageEffectTypeEmojiReaction = new(MessageEffectTypeEmojiReaction)
		return json.Unmarshal(b, t.MessageEffectTypeEmojiReaction)
	case "messageEffectTypePremiumSticker":
		t.MessageEffectTypePremiumSticker = new(MessageEffectTypePremiumSticker)
		return json.Unmarshal(b, t.MessageEffectTypePremiumSticker)
	}
	return nil
}

func (t *MessageEffectType) MarshalJSON() ([]byte, error) {
	if t.MessageEffectTypeEmojiReaction != nil {
		return json.Marshal(t.MessageEffectTypeEmojiReaction)
	}
	if t.MessageEffectTypePremiumSticker != nil {
		return json.Marshal(t.MessageEffectTypePremiumSticker)
	}
	return nil, fmt.Errorf("no value set for MessageEffectType")
}

// MessageFileType Contains information about a file with messages exported from another app
type MessageFileType struct {
	TypeStr                string                  `json:"@type"`
	MessageFileTypeGroup   *MessageFileTypeGroup   `json:"messageFileTypeGroup,omitempty"`
	MessageFileTypePrivate *MessageFileTypePrivate `json:"messageFileTypePrivate,omitempty"`
	MessageFileTypeUnknown *MessageFileTypeUnknown `json:"messageFileTypeUnknown,omitempty"`
}

func (t *MessageFileType) Type() string {
	return t.TypeStr
}

func (t *MessageFileType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageFileType) GetExtra() string {
	return ""
}

func (t *MessageFileType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageFileTypeGroup":
		t.MessageFileTypeGroup = new(MessageFileTypeGroup)
		return json.Unmarshal(b, t.MessageFileTypeGroup)
	case "messageFileTypePrivate":
		t.MessageFileTypePrivate = new(MessageFileTypePrivate)
		return json.Unmarshal(b, t.MessageFileTypePrivate)
	case "messageFileTypeUnknown":
		t.MessageFileTypeUnknown = new(MessageFileTypeUnknown)
		return json.Unmarshal(b, t.MessageFileTypeUnknown)
	}
	return nil
}

func (t *MessageFileType) MarshalJSON() ([]byte, error) {
	if t.MessageFileTypeGroup != nil {
		return json.Marshal(t.MessageFileTypeGroup)
	}
	if t.MessageFileTypePrivate != nil {
		return json.Marshal(t.MessageFileTypePrivate)
	}
	if t.MessageFileTypeUnknown != nil {
		return json.Marshal(t.MessageFileTypeUnknown)
	}
	return nil, fmt.Errorf("no value set for MessageFileType")
}

// MessageOrigin Contains information about the origin of a message
type MessageOrigin struct {
	TypeStr                 string                   `json:"@type"`
	MessageOriginChannel    *MessageOriginChannel    `json:"messageOriginChannel,omitempty"`
	MessageOriginChat       *MessageOriginChat       `json:"messageOriginChat,omitempty"`
	MessageOriginHiddenUser *MessageOriginHiddenUser `json:"messageOriginHiddenUser,omitempty"`
	MessageOriginUser       *MessageOriginUser       `json:"messageOriginUser,omitempty"`
}

func (t *MessageOrigin) Type() string {
	return t.TypeStr
}

func (t *MessageOrigin) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageOrigin) GetExtra() string {
	return ""
}

func (t *MessageOrigin) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageOriginChannel":
		t.MessageOriginChannel = new(MessageOriginChannel)
		return json.Unmarshal(b, t.MessageOriginChannel)
	case "messageOriginChat":
		t.MessageOriginChat = new(MessageOriginChat)
		return json.Unmarshal(b, t.MessageOriginChat)
	case "messageOriginHiddenUser":
		t.MessageOriginHiddenUser = new(MessageOriginHiddenUser)
		return json.Unmarshal(b, t.MessageOriginHiddenUser)
	case "messageOriginUser":
		t.MessageOriginUser = new(MessageOriginUser)
		return json.Unmarshal(b, t.MessageOriginUser)
	}
	return nil
}

func (t *MessageOrigin) MarshalJSON() ([]byte, error) {
	if t.MessageOriginChannel != nil {
		return json.Marshal(t.MessageOriginChannel)
	}
	if t.MessageOriginChat != nil {
		return json.Marshal(t.MessageOriginChat)
	}
	if t.MessageOriginHiddenUser != nil {
		return json.Marshal(t.MessageOriginHiddenUser)
	}
	if t.MessageOriginUser != nil {
		return json.Marshal(t.MessageOriginUser)
	}
	return nil, fmt.Errorf("no value set for MessageOrigin")
}

// MessageReadDate Describes read date of a recent outgoing message in a private chat
type MessageReadDate struct {
	TypeStr                              string                                `json:"@type"`
	MessageReadDateMyPrivacyRestricted   *MessageReadDateMyPrivacyRestricted   `json:"messageReadDateMyPrivacyRestricted,omitempty"`
	MessageReadDateRead                  *MessageReadDateRead                  `json:"messageReadDateRead,omitempty"`
	MessageReadDateTooOld                *MessageReadDateTooOld                `json:"messageReadDateTooOld,omitempty"`
	MessageReadDateUnread                *MessageReadDateUnread                `json:"messageReadDateUnread,omitempty"`
	MessageReadDateUserPrivacyRestricted *MessageReadDateUserPrivacyRestricted `json:"messageReadDateUserPrivacyRestricted,omitempty"`
}

func (t *MessageReadDate) Type() string {
	return t.TypeStr
}

func (t *MessageReadDate) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageReadDate) GetExtra() string {
	return ""
}

func (t *MessageReadDate) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageReadDateMyPrivacyRestricted":
		t.MessageReadDateMyPrivacyRestricted = new(MessageReadDateMyPrivacyRestricted)
		return json.Unmarshal(b, t.MessageReadDateMyPrivacyRestricted)
	case "messageReadDateRead":
		t.MessageReadDateRead = new(MessageReadDateRead)
		return json.Unmarshal(b, t.MessageReadDateRead)
	case "messageReadDateTooOld":
		t.MessageReadDateTooOld = new(MessageReadDateTooOld)
		return json.Unmarshal(b, t.MessageReadDateTooOld)
	case "messageReadDateUnread":
		t.MessageReadDateUnread = new(MessageReadDateUnread)
		return json.Unmarshal(b, t.MessageReadDateUnread)
	case "messageReadDateUserPrivacyRestricted":
		t.MessageReadDateUserPrivacyRestricted = new(MessageReadDateUserPrivacyRestricted)
		return json.Unmarshal(b, t.MessageReadDateUserPrivacyRestricted)
	}
	return nil
}

func (t *MessageReadDate) MarshalJSON() ([]byte, error) {
	if t.MessageReadDateMyPrivacyRestricted != nil {
		return json.Marshal(t.MessageReadDateMyPrivacyRestricted)
	}
	if t.MessageReadDateRead != nil {
		return json.Marshal(t.MessageReadDateRead)
	}
	if t.MessageReadDateTooOld != nil {
		return json.Marshal(t.MessageReadDateTooOld)
	}
	if t.MessageReadDateUnread != nil {
		return json.Marshal(t.MessageReadDateUnread)
	}
	if t.MessageReadDateUserPrivacyRestricted != nil {
		return json.Marshal(t.MessageReadDateUserPrivacyRestricted)
	}
	return nil, fmt.Errorf("no value set for MessageReadDate")
}

// MessageReplyTo Contains information about the message or the story a message is replying to
type MessageReplyTo struct {
	TypeStr               string                 `json:"@type"`
	MessageReplyToMessage *MessageReplyToMessage `json:"messageReplyToMessage,omitempty"`
	MessageReplyToStory   *MessageReplyToStory   `json:"messageReplyToStory,omitempty"`
}

func (t *MessageReplyTo) Type() string {
	return t.TypeStr
}

func (t *MessageReplyTo) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageReplyTo) GetExtra() string {
	return ""
}

func (t *MessageReplyTo) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageReplyToMessage":
		t.MessageReplyToMessage = new(MessageReplyToMessage)
		return json.Unmarshal(b, t.MessageReplyToMessage)
	case "messageReplyToStory":
		t.MessageReplyToStory = new(MessageReplyToStory)
		return json.Unmarshal(b, t.MessageReplyToStory)
	}
	return nil
}

func (t *MessageReplyTo) MarshalJSON() ([]byte, error) {
	if t.MessageReplyToMessage != nil {
		return json.Marshal(t.MessageReplyToMessage)
	}
	if t.MessageReplyToStory != nil {
		return json.Marshal(t.MessageReplyToStory)
	}
	return nil, fmt.Errorf("no value set for MessageReplyTo")
}

// MessageSchedulingState Contains information about the time when a scheduled message will be sent
type MessageSchedulingState struct {
	TypeStr                                      string                                        `json:"@type"`
	MessageSchedulingStateSendAtDate             *MessageSchedulingStateSendAtDate             `json:"messageSchedulingStateSendAtDate,omitempty"`
	MessageSchedulingStateSendWhenOnline         *MessageSchedulingStateSendWhenOnline         `json:"messageSchedulingStateSendWhenOnline,omitempty"`
	MessageSchedulingStateSendWhenVideoProcessed *MessageSchedulingStateSendWhenVideoProcessed `json:"messageSchedulingStateSendWhenVideoProcessed,omitempty"`
}

func (t *MessageSchedulingState) Type() string {
	return t.TypeStr
}

func (t *MessageSchedulingState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageSchedulingState) GetExtra() string {
	return ""
}

func (t *MessageSchedulingState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageSchedulingStateSendAtDate":
		t.MessageSchedulingStateSendAtDate = new(MessageSchedulingStateSendAtDate)
		return json.Unmarshal(b, t.MessageSchedulingStateSendAtDate)
	case "messageSchedulingStateSendWhenOnline":
		t.MessageSchedulingStateSendWhenOnline = new(MessageSchedulingStateSendWhenOnline)
		return json.Unmarshal(b, t.MessageSchedulingStateSendWhenOnline)
	case "messageSchedulingStateSendWhenVideoProcessed":
		t.MessageSchedulingStateSendWhenVideoProcessed = new(MessageSchedulingStateSendWhenVideoProcessed)
		return json.Unmarshal(b, t.MessageSchedulingStateSendWhenVideoProcessed)
	}
	return nil
}

func (t *MessageSchedulingState) MarshalJSON() ([]byte, error) {
	if t.MessageSchedulingStateSendAtDate != nil {
		return json.Marshal(t.MessageSchedulingStateSendAtDate)
	}
	if t.MessageSchedulingStateSendWhenOnline != nil {
		return json.Marshal(t.MessageSchedulingStateSendWhenOnline)
	}
	if t.MessageSchedulingStateSendWhenVideoProcessed != nil {
		return json.Marshal(t.MessageSchedulingStateSendWhenVideoProcessed)
	}
	return nil, fmt.Errorf("no value set for MessageSchedulingState")
}

// MessageSelfDestructType Describes when a message will be self-destructed
type MessageSelfDestructType struct {
	TypeStr                            string                              `json:"@type"`
	MessageSelfDestructTypeImmediately *MessageSelfDestructTypeImmediately `json:"messageSelfDestructTypeImmediately,omitempty"`
	MessageSelfDestructTypeTimer       *MessageSelfDestructTypeTimer       `json:"messageSelfDestructTypeTimer,omitempty"`
}

func (t *MessageSelfDestructType) Type() string {
	return t.TypeStr
}

func (t *MessageSelfDestructType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageSelfDestructType) GetExtra() string {
	return ""
}

func (t *MessageSelfDestructType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageSelfDestructTypeImmediately":
		t.MessageSelfDestructTypeImmediately = new(MessageSelfDestructTypeImmediately)
		return json.Unmarshal(b, t.MessageSelfDestructTypeImmediately)
	case "messageSelfDestructTypeTimer":
		t.MessageSelfDestructTypeTimer = new(MessageSelfDestructTypeTimer)
		return json.Unmarshal(b, t.MessageSelfDestructTypeTimer)
	}
	return nil
}

func (t *MessageSelfDestructType) MarshalJSON() ([]byte, error) {
	if t.MessageSelfDestructTypeImmediately != nil {
		return json.Marshal(t.MessageSelfDestructTypeImmediately)
	}
	if t.MessageSelfDestructTypeTimer != nil {
		return json.Marshal(t.MessageSelfDestructTypeTimer)
	}
	return nil, fmt.Errorf("no value set for MessageSelfDestructType")
}

// MessageSender Contains information about the sender of a message
type MessageSender struct {
	TypeStr           string             `json:"@type"`
	MessageSenderChat *MessageSenderChat `json:"messageSenderChat,omitempty"`
	MessageSenderUser *MessageSenderUser `json:"messageSenderUser,omitempty"`
}

func (t *MessageSender) Type() string {
	return t.TypeStr
}

func (t *MessageSender) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageSender) GetExtra() string {
	return ""
}

func (t *MessageSender) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageSenderChat":
		t.MessageSenderChat = new(MessageSenderChat)
		return json.Unmarshal(b, t.MessageSenderChat)
	case "messageSenderUser":
		t.MessageSenderUser = new(MessageSenderUser)
		return json.Unmarshal(b, t.MessageSenderUser)
	}
	return nil
}

func (t *MessageSender) MarshalJSON() ([]byte, error) {
	if t.MessageSenderChat != nil {
		return json.Marshal(t.MessageSenderChat)
	}
	if t.MessageSenderUser != nil {
		return json.Marshal(t.MessageSenderUser)
	}
	return nil, fmt.Errorf("no value set for MessageSender")
}

// MessageSendingState Contains information about the sending state of the message
type MessageSendingState struct {
	TypeStr                    string                      `json:"@type"`
	MessageSendingStateFailed  *MessageSendingStateFailed  `json:"messageSendingStateFailed,omitempty"`
	MessageSendingStatePending *MessageSendingStatePending `json:"messageSendingStatePending,omitempty"`
}

func (t *MessageSendingState) Type() string {
	return t.TypeStr
}

func (t *MessageSendingState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageSendingState) GetExtra() string {
	return ""
}

func (t *MessageSendingState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageSendingStateFailed":
		t.MessageSendingStateFailed = new(MessageSendingStateFailed)
		return json.Unmarshal(b, t.MessageSendingStateFailed)
	case "messageSendingStatePending":
		t.MessageSendingStatePending = new(MessageSendingStatePending)
		return json.Unmarshal(b, t.MessageSendingStatePending)
	}
	return nil
}

func (t *MessageSendingState) MarshalJSON() ([]byte, error) {
	if t.MessageSendingStateFailed != nil {
		return json.Marshal(t.MessageSendingStateFailed)
	}
	if t.MessageSendingStatePending != nil {
		return json.Marshal(t.MessageSendingStatePending)
	}
	return nil, fmt.Errorf("no value set for MessageSendingState")
}

// MessageSource Describes source of a message
type MessageSource struct {
	TypeStr                                     string                                       `json:"@type"`
	MessageSourceChatEventLog                   *MessageSourceChatEventLog                   `json:"messageSourceChatEventLog,omitempty"`
	MessageSourceChatHistory                    *MessageSourceChatHistory                    `json:"messageSourceChatHistory,omitempty"`
	MessageSourceChatList                       *MessageSourceChatList                       `json:"messageSourceChatList,omitempty"`
	MessageSourceDirectMessagesChatTopicHistory *MessageSourceDirectMessagesChatTopicHistory `json:"messageSourceDirectMessagesChatTopicHistory,omitempty"`
	MessageSourceForumTopicHistory              *MessageSourceForumTopicHistory              `json:"messageSourceForumTopicHistory,omitempty"`
	MessageSourceHistoryPreview                 *MessageSourceHistoryPreview                 `json:"messageSourceHistoryPreview,omitempty"`
	MessageSourceMessageThreadHistory           *MessageSourceMessageThreadHistory           `json:"messageSourceMessageThreadHistory,omitempty"`
	MessageSourceNotification                   *MessageSourceNotification                   `json:"messageSourceNotification,omitempty"`
	MessageSourceOther                          *MessageSourceOther                          `json:"messageSourceOther,omitempty"`
	MessageSourceScreenshot                     *MessageSourceScreenshot                     `json:"messageSourceScreenshot,omitempty"`
	MessageSourceSearch                         *MessageSourceSearch                         `json:"messageSourceSearch,omitempty"`
}

func (t *MessageSource) Type() string {
	return t.TypeStr
}

func (t *MessageSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageSource) GetExtra() string {
	return ""
}

func (t *MessageSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageSourceChatEventLog":
		t.MessageSourceChatEventLog = new(MessageSourceChatEventLog)
		return json.Unmarshal(b, t.MessageSourceChatEventLog)
	case "messageSourceChatHistory":
		t.MessageSourceChatHistory = new(MessageSourceChatHistory)
		return json.Unmarshal(b, t.MessageSourceChatHistory)
	case "messageSourceChatList":
		t.MessageSourceChatList = new(MessageSourceChatList)
		return json.Unmarshal(b, t.MessageSourceChatList)
	case "messageSourceDirectMessagesChatTopicHistory":
		t.MessageSourceDirectMessagesChatTopicHistory = new(MessageSourceDirectMessagesChatTopicHistory)
		return json.Unmarshal(b, t.MessageSourceDirectMessagesChatTopicHistory)
	case "messageSourceForumTopicHistory":
		t.MessageSourceForumTopicHistory = new(MessageSourceForumTopicHistory)
		return json.Unmarshal(b, t.MessageSourceForumTopicHistory)
	case "messageSourceHistoryPreview":
		t.MessageSourceHistoryPreview = new(MessageSourceHistoryPreview)
		return json.Unmarshal(b, t.MessageSourceHistoryPreview)
	case "messageSourceMessageThreadHistory":
		t.MessageSourceMessageThreadHistory = new(MessageSourceMessageThreadHistory)
		return json.Unmarshal(b, t.MessageSourceMessageThreadHistory)
	case "messageSourceNotification":
		t.MessageSourceNotification = new(MessageSourceNotification)
		return json.Unmarshal(b, t.MessageSourceNotification)
	case "messageSourceOther":
		t.MessageSourceOther = new(MessageSourceOther)
		return json.Unmarshal(b, t.MessageSourceOther)
	case "messageSourceScreenshot":
		t.MessageSourceScreenshot = new(MessageSourceScreenshot)
		return json.Unmarshal(b, t.MessageSourceScreenshot)
	case "messageSourceSearch":
		t.MessageSourceSearch = new(MessageSourceSearch)
		return json.Unmarshal(b, t.MessageSourceSearch)
	}
	return nil
}

func (t *MessageSource) MarshalJSON() ([]byte, error) {
	if t.MessageSourceChatEventLog != nil {
		return json.Marshal(t.MessageSourceChatEventLog)
	}
	if t.MessageSourceChatHistory != nil {
		return json.Marshal(t.MessageSourceChatHistory)
	}
	if t.MessageSourceChatList != nil {
		return json.Marshal(t.MessageSourceChatList)
	}
	if t.MessageSourceDirectMessagesChatTopicHistory != nil {
		return json.Marshal(t.MessageSourceDirectMessagesChatTopicHistory)
	}
	if t.MessageSourceForumTopicHistory != nil {
		return json.Marshal(t.MessageSourceForumTopicHistory)
	}
	if t.MessageSourceHistoryPreview != nil {
		return json.Marshal(t.MessageSourceHistoryPreview)
	}
	if t.MessageSourceMessageThreadHistory != nil {
		return json.Marshal(t.MessageSourceMessageThreadHistory)
	}
	if t.MessageSourceNotification != nil {
		return json.Marshal(t.MessageSourceNotification)
	}
	if t.MessageSourceOther != nil {
		return json.Marshal(t.MessageSourceOther)
	}
	if t.MessageSourceScreenshot != nil {
		return json.Marshal(t.MessageSourceScreenshot)
	}
	if t.MessageSourceSearch != nil {
		return json.Marshal(t.MessageSourceSearch)
	}
	return nil, fmt.Errorf("no value set for MessageSource")
}

// MessageTopic Describes a topic of messages in a chat
type MessageTopic struct {
	TypeStr                    string                      `json:"@type"`
	MessageTopicDirectMessages *MessageTopicDirectMessages `json:"messageTopicDirectMessages,omitempty"`
	MessageTopicForum          *MessageTopicForum          `json:"messageTopicForum,omitempty"`
	MessageTopicSavedMessages  *MessageTopicSavedMessages  `json:"messageTopicSavedMessages,omitempty"`
	MessageTopicThread         *MessageTopicThread         `json:"messageTopicThread,omitempty"`
}

func (t *MessageTopic) Type() string {
	return t.TypeStr
}

func (t *MessageTopic) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *MessageTopic) GetExtra() string {
	return ""
}

func (t *MessageTopic) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "messageTopicDirectMessages":
		t.MessageTopicDirectMessages = new(MessageTopicDirectMessages)
		return json.Unmarshal(b, t.MessageTopicDirectMessages)
	case "messageTopicForum":
		t.MessageTopicForum = new(MessageTopicForum)
		return json.Unmarshal(b, t.MessageTopicForum)
	case "messageTopicSavedMessages":
		t.MessageTopicSavedMessages = new(MessageTopicSavedMessages)
		return json.Unmarshal(b, t.MessageTopicSavedMessages)
	case "messageTopicThread":
		t.MessageTopicThread = new(MessageTopicThread)
		return json.Unmarshal(b, t.MessageTopicThread)
	}
	return nil
}

func (t *MessageTopic) MarshalJSON() ([]byte, error) {
	if t.MessageTopicDirectMessages != nil {
		return json.Marshal(t.MessageTopicDirectMessages)
	}
	if t.MessageTopicForum != nil {
		return json.Marshal(t.MessageTopicForum)
	}
	if t.MessageTopicSavedMessages != nil {
		return json.Marshal(t.MessageTopicSavedMessages)
	}
	if t.MessageTopicThread != nil {
		return json.Marshal(t.MessageTopicThread)
	}
	return nil, fmt.Errorf("no value set for MessageTopic")
}

// NetworkStatisticsEntry Contains statistics about network usage
type NetworkStatisticsEntry struct {
	TypeStr                    string                      `json:"@type"`
	NetworkStatisticsEntryCall *NetworkStatisticsEntryCall `json:"networkStatisticsEntryCall,omitempty"`
	NetworkStatisticsEntryFile *NetworkStatisticsEntryFile `json:"networkStatisticsEntryFile,omitempty"`
}

func (t *NetworkStatisticsEntry) Type() string {
	return t.TypeStr
}

func (t *NetworkStatisticsEntry) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *NetworkStatisticsEntry) GetExtra() string {
	return ""
}

func (t *NetworkStatisticsEntry) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "networkStatisticsEntryCall":
		t.NetworkStatisticsEntryCall = new(NetworkStatisticsEntryCall)
		return json.Unmarshal(b, t.NetworkStatisticsEntryCall)
	case "networkStatisticsEntryFile":
		t.NetworkStatisticsEntryFile = new(NetworkStatisticsEntryFile)
		return json.Unmarshal(b, t.NetworkStatisticsEntryFile)
	}
	return nil
}

func (t *NetworkStatisticsEntry) MarshalJSON() ([]byte, error) {
	if t.NetworkStatisticsEntryCall != nil {
		return json.Marshal(t.NetworkStatisticsEntryCall)
	}
	if t.NetworkStatisticsEntryFile != nil {
		return json.Marshal(t.NetworkStatisticsEntryFile)
	}
	return nil, fmt.Errorf("no value set for NetworkStatisticsEntry")
}

// NetworkType Represents the type of network
type NetworkType struct {
	TypeStr                  string                    `json:"@type"`
	NetworkTypeMobile        *NetworkTypeMobile        `json:"networkTypeMobile,omitempty"`
	NetworkTypeMobileRoaming *NetworkTypeMobileRoaming `json:"networkTypeMobileRoaming,omitempty"`
	NetworkTypeNone          *NetworkTypeNone          `json:"networkTypeNone,omitempty"`
	NetworkTypeOther         *NetworkTypeOther         `json:"networkTypeOther,omitempty"`
	NetworkTypeWiFi          *NetworkTypeWiFi          `json:"networkTypeWiFi,omitempty"`
}

func (t *NetworkType) Type() string {
	return t.TypeStr
}

func (t *NetworkType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *NetworkType) GetExtra() string {
	return ""
}

func (t *NetworkType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "networkTypeMobile":
		t.NetworkTypeMobile = new(NetworkTypeMobile)
		return json.Unmarshal(b, t.NetworkTypeMobile)
	case "networkTypeMobileRoaming":
		t.NetworkTypeMobileRoaming = new(NetworkTypeMobileRoaming)
		return json.Unmarshal(b, t.NetworkTypeMobileRoaming)
	case "networkTypeNone":
		t.NetworkTypeNone = new(NetworkTypeNone)
		return json.Unmarshal(b, t.NetworkTypeNone)
	case "networkTypeOther":
		t.NetworkTypeOther = new(NetworkTypeOther)
		return json.Unmarshal(b, t.NetworkTypeOther)
	case "networkTypeWiFi":
		t.NetworkTypeWiFi = new(NetworkTypeWiFi)
		return json.Unmarshal(b, t.NetworkTypeWiFi)
	}
	return nil
}

func (t *NetworkType) MarshalJSON() ([]byte, error) {
	if t.NetworkTypeMobile != nil {
		return json.Marshal(t.NetworkTypeMobile)
	}
	if t.NetworkTypeMobileRoaming != nil {
		return json.Marshal(t.NetworkTypeMobileRoaming)
	}
	if t.NetworkTypeNone != nil {
		return json.Marshal(t.NetworkTypeNone)
	}
	if t.NetworkTypeOther != nil {
		return json.Marshal(t.NetworkTypeOther)
	}
	if t.NetworkTypeWiFi != nil {
		return json.Marshal(t.NetworkTypeWiFi)
	}
	return nil, fmt.Errorf("no value set for NetworkType")
}

// NotificationGroupType Describes the type of notifications in a notification group
type NotificationGroupType struct {
	TypeStr                         string                           `json:"@type"`
	NotificationGroupTypeCalls      *NotificationGroupTypeCalls      `json:"notificationGroupTypeCalls,omitempty"`
	NotificationGroupTypeMentions   *NotificationGroupTypeMentions   `json:"notificationGroupTypeMentions,omitempty"`
	NotificationGroupTypeMessages   *NotificationGroupTypeMessages   `json:"notificationGroupTypeMessages,omitempty"`
	NotificationGroupTypeSecretChat *NotificationGroupTypeSecretChat `json:"notificationGroupTypeSecretChat,omitempty"`
}

func (t *NotificationGroupType) Type() string {
	return t.TypeStr
}

func (t *NotificationGroupType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *NotificationGroupType) GetExtra() string {
	return ""
}

func (t *NotificationGroupType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "notificationGroupTypeCalls":
		t.NotificationGroupTypeCalls = new(NotificationGroupTypeCalls)
		return json.Unmarshal(b, t.NotificationGroupTypeCalls)
	case "notificationGroupTypeMentions":
		t.NotificationGroupTypeMentions = new(NotificationGroupTypeMentions)
		return json.Unmarshal(b, t.NotificationGroupTypeMentions)
	case "notificationGroupTypeMessages":
		t.NotificationGroupTypeMessages = new(NotificationGroupTypeMessages)
		return json.Unmarshal(b, t.NotificationGroupTypeMessages)
	case "notificationGroupTypeSecretChat":
		t.NotificationGroupTypeSecretChat = new(NotificationGroupTypeSecretChat)
		return json.Unmarshal(b, t.NotificationGroupTypeSecretChat)
	}
	return nil
}

func (t *NotificationGroupType) MarshalJSON() ([]byte, error) {
	if t.NotificationGroupTypeCalls != nil {
		return json.Marshal(t.NotificationGroupTypeCalls)
	}
	if t.NotificationGroupTypeMentions != nil {
		return json.Marshal(t.NotificationGroupTypeMentions)
	}
	if t.NotificationGroupTypeMessages != nil {
		return json.Marshal(t.NotificationGroupTypeMessages)
	}
	if t.NotificationGroupTypeSecretChat != nil {
		return json.Marshal(t.NotificationGroupTypeSecretChat)
	}
	return nil, fmt.Errorf("no value set for NotificationGroupType")
}

// NotificationSettingsScope Describes the types of chats to which notification settings are relevant
type NotificationSettingsScope struct {
	TypeStr                               string                                 `json:"@type"`
	NotificationSettingsScopeChannelChats *NotificationSettingsScopeChannelChats `json:"notificationSettingsScopeChannelChats,omitempty"`
	NotificationSettingsScopeGroupChats   *NotificationSettingsScopeGroupChats   `json:"notificationSettingsScopeGroupChats,omitempty"`
	NotificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats `json:"notificationSettingsScopePrivateChats,omitempty"`
}

func (t *NotificationSettingsScope) Type() string {
	return t.TypeStr
}

func (t *NotificationSettingsScope) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *NotificationSettingsScope) GetExtra() string {
	return ""
}

func (t *NotificationSettingsScope) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "notificationSettingsScopeChannelChats":
		t.NotificationSettingsScopeChannelChats = new(NotificationSettingsScopeChannelChats)
		return json.Unmarshal(b, t.NotificationSettingsScopeChannelChats)
	case "notificationSettingsScopeGroupChats":
		t.NotificationSettingsScopeGroupChats = new(NotificationSettingsScopeGroupChats)
		return json.Unmarshal(b, t.NotificationSettingsScopeGroupChats)
	case "notificationSettingsScopePrivateChats":
		t.NotificationSettingsScopePrivateChats = new(NotificationSettingsScopePrivateChats)
		return json.Unmarshal(b, t.NotificationSettingsScopePrivateChats)
	}
	return nil
}

func (t *NotificationSettingsScope) MarshalJSON() ([]byte, error) {
	if t.NotificationSettingsScopeChannelChats != nil {
		return json.Marshal(t.NotificationSettingsScopeChannelChats)
	}
	if t.NotificationSettingsScopeGroupChats != nil {
		return json.Marshal(t.NotificationSettingsScopeGroupChats)
	}
	if t.NotificationSettingsScopePrivateChats != nil {
		return json.Marshal(t.NotificationSettingsScopePrivateChats)
	}
	return nil, fmt.Errorf("no value set for NotificationSettingsScope")
}

// NotificationType Contains detailed information about a notification
type NotificationType struct {
	TypeStr                        string                          `json:"@type"`
	NotificationTypeNewCall        *NotificationTypeNewCall        `json:"notificationTypeNewCall,omitempty"`
	NotificationTypeNewMessage     *NotificationTypeNewMessage     `json:"notificationTypeNewMessage,omitempty"`
	NotificationTypeNewPushMessage *NotificationTypeNewPushMessage `json:"notificationTypeNewPushMessage,omitempty"`
	NotificationTypeNewSecretChat  *NotificationTypeNewSecretChat  `json:"notificationTypeNewSecretChat,omitempty"`
}

func (t *NotificationType) Type() string {
	return t.TypeStr
}

func (t *NotificationType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *NotificationType) GetExtra() string {
	return ""
}

func (t *NotificationType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "notificationTypeNewCall":
		t.NotificationTypeNewCall = new(NotificationTypeNewCall)
		return json.Unmarshal(b, t.NotificationTypeNewCall)
	case "notificationTypeNewMessage":
		t.NotificationTypeNewMessage = new(NotificationTypeNewMessage)
		return json.Unmarshal(b, t.NotificationTypeNewMessage)
	case "notificationTypeNewPushMessage":
		t.NotificationTypeNewPushMessage = new(NotificationTypeNewPushMessage)
		return json.Unmarshal(b, t.NotificationTypeNewPushMessage)
	case "notificationTypeNewSecretChat":
		t.NotificationTypeNewSecretChat = new(NotificationTypeNewSecretChat)
		return json.Unmarshal(b, t.NotificationTypeNewSecretChat)
	}
	return nil
}

func (t *NotificationType) MarshalJSON() ([]byte, error) {
	if t.NotificationTypeNewCall != nil {
		return json.Marshal(t.NotificationTypeNewCall)
	}
	if t.NotificationTypeNewMessage != nil {
		return json.Marshal(t.NotificationTypeNewMessage)
	}
	if t.NotificationTypeNewPushMessage != nil {
		return json.Marshal(t.NotificationTypeNewPushMessage)
	}
	if t.NotificationTypeNewSecretChat != nil {
		return json.Marshal(t.NotificationTypeNewSecretChat)
	}
	return nil, fmt.Errorf("no value set for NotificationType")
}

// OptionValue Represents the value of an option
type OptionValue struct {
	TypeStr            string              `json:"@type"`
	OptionValueBoolean *OptionValueBoolean `json:"optionValueBoolean,omitempty"`
	OptionValueEmpty   *OptionValueEmpty   `json:"optionValueEmpty,omitempty"`
	OptionValueInteger *OptionValueInteger `json:"optionValueInteger,omitempty"`
	OptionValueString  *OptionValueString  `json:"optionValueString,omitempty"`
}

func (t *OptionValue) Type() string {
	return t.TypeStr
}

func (t *OptionValue) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *OptionValue) GetExtra() string {
	return ""
}

func (t *OptionValue) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "optionValueBoolean":
		t.OptionValueBoolean = new(OptionValueBoolean)
		return json.Unmarshal(b, t.OptionValueBoolean)
	case "optionValueEmpty":
		t.OptionValueEmpty = new(OptionValueEmpty)
		return json.Unmarshal(b, t.OptionValueEmpty)
	case "optionValueInteger":
		t.OptionValueInteger = new(OptionValueInteger)
		return json.Unmarshal(b, t.OptionValueInteger)
	case "optionValueString":
		t.OptionValueString = new(OptionValueString)
		return json.Unmarshal(b, t.OptionValueString)
	}
	return nil
}

func (t *OptionValue) MarshalJSON() ([]byte, error) {
	if t.OptionValueBoolean != nil {
		return json.Marshal(t.OptionValueBoolean)
	}
	if t.OptionValueEmpty != nil {
		return json.Marshal(t.OptionValueEmpty)
	}
	if t.OptionValueInteger != nil {
		return json.Marshal(t.OptionValueInteger)
	}
	if t.OptionValueString != nil {
		return json.Marshal(t.OptionValueString)
	}
	return nil, fmt.Errorf("no value set for OptionValue")
}

// PageBlock Describes a block of an instant view for a web page
type PageBlock struct {
	TypeStr                  string                    `json:"@type"`
	PageBlockAnchor          *PageBlockAnchor          `json:"pageBlockAnchor,omitempty"`
	PageBlockAnimation       *PageBlockAnimation       `json:"pageBlockAnimation,omitempty"`
	PageBlockAudio           *PageBlockAudio           `json:"pageBlockAudio,omitempty"`
	PageBlockAuthorDate      *PageBlockAuthorDate      `json:"pageBlockAuthorDate,omitempty"`
	PageBlockBlockQuote      *PageBlockBlockQuote      `json:"pageBlockBlockQuote,omitempty"`
	PageBlockChatLink        *PageBlockChatLink        `json:"pageBlockChatLink,omitempty"`
	PageBlockCollage         *PageBlockCollage         `json:"pageBlockCollage,omitempty"`
	PageBlockCover           *PageBlockCover           `json:"pageBlockCover,omitempty"`
	PageBlockDetails         *PageBlockDetails         `json:"pageBlockDetails,omitempty"`
	PageBlockDivider         *PageBlockDivider         `json:"pageBlockDivider,omitempty"`
	PageBlockEmbedded        *PageBlockEmbedded        `json:"pageBlockEmbedded,omitempty"`
	PageBlockEmbeddedPost    *PageBlockEmbeddedPost    `json:"pageBlockEmbeddedPost,omitempty"`
	PageBlockFooter          *PageBlockFooter          `json:"pageBlockFooter,omitempty"`
	PageBlockHeader          *PageBlockHeader          `json:"pageBlockHeader,omitempty"`
	PageBlockKicker          *PageBlockKicker          `json:"pageBlockKicker,omitempty"`
	PageBlockList            *PageBlockList            `json:"pageBlockList,omitempty"`
	PageBlockMap             *PageBlockMap             `json:"pageBlockMap,omitempty"`
	PageBlockParagraph       *PageBlockParagraph       `json:"pageBlockParagraph,omitempty"`
	PageBlockPhoto           *PageBlockPhoto           `json:"pageBlockPhoto,omitempty"`
	PageBlockPreformatted    *PageBlockPreformatted    `json:"pageBlockPreformatted,omitempty"`
	PageBlockPullQuote       *PageBlockPullQuote       `json:"pageBlockPullQuote,omitempty"`
	PageBlockRelatedArticles *PageBlockRelatedArticles `json:"pageBlockRelatedArticles,omitempty"`
	PageBlockSlideshow       *PageBlockSlideshow       `json:"pageBlockSlideshow,omitempty"`
	PageBlockSubheader       *PageBlockSubheader       `json:"pageBlockSubheader,omitempty"`
	PageBlockSubtitle        *PageBlockSubtitle        `json:"pageBlockSubtitle,omitempty"`
	PageBlockTable           *PageBlockTable           `json:"pageBlockTable,omitempty"`
	PageBlockTitle           *PageBlockTitle           `json:"pageBlockTitle,omitempty"`
	PageBlockVideo           *PageBlockVideo           `json:"pageBlockVideo,omitempty"`
	PageBlockVoiceNote       *PageBlockVoiceNote       `json:"pageBlockVoiceNote,omitempty"`
}

func (t *PageBlock) Type() string {
	return t.TypeStr
}

func (t *PageBlock) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PageBlock) GetExtra() string {
	return ""
}

func (t *PageBlock) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "pageBlockAnchor":
		t.PageBlockAnchor = new(PageBlockAnchor)
		return json.Unmarshal(b, t.PageBlockAnchor)
	case "pageBlockAnimation":
		t.PageBlockAnimation = new(PageBlockAnimation)
		return json.Unmarshal(b, t.PageBlockAnimation)
	case "pageBlockAudio":
		t.PageBlockAudio = new(PageBlockAudio)
		return json.Unmarshal(b, t.PageBlockAudio)
	case "pageBlockAuthorDate":
		t.PageBlockAuthorDate = new(PageBlockAuthorDate)
		return json.Unmarshal(b, t.PageBlockAuthorDate)
	case "pageBlockBlockQuote":
		t.PageBlockBlockQuote = new(PageBlockBlockQuote)
		return json.Unmarshal(b, t.PageBlockBlockQuote)
	case "pageBlockChatLink":
		t.PageBlockChatLink = new(PageBlockChatLink)
		return json.Unmarshal(b, t.PageBlockChatLink)
	case "pageBlockCollage":
		t.PageBlockCollage = new(PageBlockCollage)
		return json.Unmarshal(b, t.PageBlockCollage)
	case "pageBlockCover":
		t.PageBlockCover = new(PageBlockCover)
		return json.Unmarshal(b, t.PageBlockCover)
	case "pageBlockDetails":
		t.PageBlockDetails = new(PageBlockDetails)
		return json.Unmarshal(b, t.PageBlockDetails)
	case "pageBlockDivider":
		t.PageBlockDivider = new(PageBlockDivider)
		return json.Unmarshal(b, t.PageBlockDivider)
	case "pageBlockEmbedded":
		t.PageBlockEmbedded = new(PageBlockEmbedded)
		return json.Unmarshal(b, t.PageBlockEmbedded)
	case "pageBlockEmbeddedPost":
		t.PageBlockEmbeddedPost = new(PageBlockEmbeddedPost)
		return json.Unmarshal(b, t.PageBlockEmbeddedPost)
	case "pageBlockFooter":
		t.PageBlockFooter = new(PageBlockFooter)
		return json.Unmarshal(b, t.PageBlockFooter)
	case "pageBlockHeader":
		t.PageBlockHeader = new(PageBlockHeader)
		return json.Unmarshal(b, t.PageBlockHeader)
	case "pageBlockKicker":
		t.PageBlockKicker = new(PageBlockKicker)
		return json.Unmarshal(b, t.PageBlockKicker)
	case "pageBlockList":
		t.PageBlockList = new(PageBlockList)
		return json.Unmarshal(b, t.PageBlockList)
	case "pageBlockMap":
		t.PageBlockMap = new(PageBlockMap)
		return json.Unmarshal(b, t.PageBlockMap)
	case "pageBlockParagraph":
		t.PageBlockParagraph = new(PageBlockParagraph)
		return json.Unmarshal(b, t.PageBlockParagraph)
	case "pageBlockPhoto":
		t.PageBlockPhoto = new(PageBlockPhoto)
		return json.Unmarshal(b, t.PageBlockPhoto)
	case "pageBlockPreformatted":
		t.PageBlockPreformatted = new(PageBlockPreformatted)
		return json.Unmarshal(b, t.PageBlockPreformatted)
	case "pageBlockPullQuote":
		t.PageBlockPullQuote = new(PageBlockPullQuote)
		return json.Unmarshal(b, t.PageBlockPullQuote)
	case "pageBlockRelatedArticles":
		t.PageBlockRelatedArticles = new(PageBlockRelatedArticles)
		return json.Unmarshal(b, t.PageBlockRelatedArticles)
	case "pageBlockSlideshow":
		t.PageBlockSlideshow = new(PageBlockSlideshow)
		return json.Unmarshal(b, t.PageBlockSlideshow)
	case "pageBlockSubheader":
		t.PageBlockSubheader = new(PageBlockSubheader)
		return json.Unmarshal(b, t.PageBlockSubheader)
	case "pageBlockSubtitle":
		t.PageBlockSubtitle = new(PageBlockSubtitle)
		return json.Unmarshal(b, t.PageBlockSubtitle)
	case "pageBlockTable":
		t.PageBlockTable = new(PageBlockTable)
		return json.Unmarshal(b, t.PageBlockTable)
	case "pageBlockTitle":
		t.PageBlockTitle = new(PageBlockTitle)
		return json.Unmarshal(b, t.PageBlockTitle)
	case "pageBlockVideo":
		t.PageBlockVideo = new(PageBlockVideo)
		return json.Unmarshal(b, t.PageBlockVideo)
	case "pageBlockVoiceNote":
		t.PageBlockVoiceNote = new(PageBlockVoiceNote)
		return json.Unmarshal(b, t.PageBlockVoiceNote)
	}
	return nil
}

func (t *PageBlock) MarshalJSON() ([]byte, error) {
	if t.PageBlockAnchor != nil {
		return json.Marshal(t.PageBlockAnchor)
	}
	if t.PageBlockAnimation != nil {
		return json.Marshal(t.PageBlockAnimation)
	}
	if t.PageBlockAudio != nil {
		return json.Marshal(t.PageBlockAudio)
	}
	if t.PageBlockAuthorDate != nil {
		return json.Marshal(t.PageBlockAuthorDate)
	}
	if t.PageBlockBlockQuote != nil {
		return json.Marshal(t.PageBlockBlockQuote)
	}
	if t.PageBlockChatLink != nil {
		return json.Marshal(t.PageBlockChatLink)
	}
	if t.PageBlockCollage != nil {
		return json.Marshal(t.PageBlockCollage)
	}
	if t.PageBlockCover != nil {
		return json.Marshal(t.PageBlockCover)
	}
	if t.PageBlockDetails != nil {
		return json.Marshal(t.PageBlockDetails)
	}
	if t.PageBlockDivider != nil {
		return json.Marshal(t.PageBlockDivider)
	}
	if t.PageBlockEmbedded != nil {
		return json.Marshal(t.PageBlockEmbedded)
	}
	if t.PageBlockEmbeddedPost != nil {
		return json.Marshal(t.PageBlockEmbeddedPost)
	}
	if t.PageBlockFooter != nil {
		return json.Marshal(t.PageBlockFooter)
	}
	if t.PageBlockHeader != nil {
		return json.Marshal(t.PageBlockHeader)
	}
	if t.PageBlockKicker != nil {
		return json.Marshal(t.PageBlockKicker)
	}
	if t.PageBlockList != nil {
		return json.Marshal(t.PageBlockList)
	}
	if t.PageBlockMap != nil {
		return json.Marshal(t.PageBlockMap)
	}
	if t.PageBlockParagraph != nil {
		return json.Marshal(t.PageBlockParagraph)
	}
	if t.PageBlockPhoto != nil {
		return json.Marshal(t.PageBlockPhoto)
	}
	if t.PageBlockPreformatted != nil {
		return json.Marshal(t.PageBlockPreformatted)
	}
	if t.PageBlockPullQuote != nil {
		return json.Marshal(t.PageBlockPullQuote)
	}
	if t.PageBlockRelatedArticles != nil {
		return json.Marshal(t.PageBlockRelatedArticles)
	}
	if t.PageBlockSlideshow != nil {
		return json.Marshal(t.PageBlockSlideshow)
	}
	if t.PageBlockSubheader != nil {
		return json.Marshal(t.PageBlockSubheader)
	}
	if t.PageBlockSubtitle != nil {
		return json.Marshal(t.PageBlockSubtitle)
	}
	if t.PageBlockTable != nil {
		return json.Marshal(t.PageBlockTable)
	}
	if t.PageBlockTitle != nil {
		return json.Marshal(t.PageBlockTitle)
	}
	if t.PageBlockVideo != nil {
		return json.Marshal(t.PageBlockVideo)
	}
	if t.PageBlockVoiceNote != nil {
		return json.Marshal(t.PageBlockVoiceNote)
	}
	return nil, fmt.Errorf("no value set for PageBlock")
}

// PageBlockHorizontalAlignment Describes a horizontal alignment of a table cell content
type PageBlockHorizontalAlignment struct {
	TypeStr                            string                              `json:"@type"`
	PageBlockHorizontalAlignmentCenter *PageBlockHorizontalAlignmentCenter `json:"pageBlockHorizontalAlignmentCenter,omitempty"`
	PageBlockHorizontalAlignmentLeft   *PageBlockHorizontalAlignmentLeft   `json:"pageBlockHorizontalAlignmentLeft,omitempty"`
	PageBlockHorizontalAlignmentRight  *PageBlockHorizontalAlignmentRight  `json:"pageBlockHorizontalAlignmentRight,omitempty"`
}

func (t *PageBlockHorizontalAlignment) Type() string {
	return t.TypeStr
}

func (t *PageBlockHorizontalAlignment) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PageBlockHorizontalAlignment) GetExtra() string {
	return ""
}

func (t *PageBlockHorizontalAlignment) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "pageBlockHorizontalAlignmentCenter":
		t.PageBlockHorizontalAlignmentCenter = new(PageBlockHorizontalAlignmentCenter)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentCenter)
	case "pageBlockHorizontalAlignmentLeft":
		t.PageBlockHorizontalAlignmentLeft = new(PageBlockHorizontalAlignmentLeft)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentLeft)
	case "pageBlockHorizontalAlignmentRight":
		t.PageBlockHorizontalAlignmentRight = new(PageBlockHorizontalAlignmentRight)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentRight)
	}
	return nil
}

func (t *PageBlockHorizontalAlignment) MarshalJSON() ([]byte, error) {
	if t.PageBlockHorizontalAlignmentCenter != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentCenter)
	}
	if t.PageBlockHorizontalAlignmentLeft != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentLeft)
	}
	if t.PageBlockHorizontalAlignmentRight != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentRight)
	}
	return nil, fmt.Errorf("no value set for PageBlockHorizontalAlignment")
}

// PageBlockVerticalAlignment Describes a Vertical alignment of a table cell content
type PageBlockVerticalAlignment struct {
	TypeStr                          string                            `json:"@type"`
	PageBlockVerticalAlignmentBottom *PageBlockVerticalAlignmentBottom `json:"pageBlockVerticalAlignmentBottom,omitempty"`
	PageBlockVerticalAlignmentMiddle *PageBlockVerticalAlignmentMiddle `json:"pageBlockVerticalAlignmentMiddle,omitempty"`
	PageBlockVerticalAlignmentTop    *PageBlockVerticalAlignmentTop    `json:"pageBlockVerticalAlignmentTop,omitempty"`
}

func (t *PageBlockVerticalAlignment) Type() string {
	return t.TypeStr
}

func (t *PageBlockVerticalAlignment) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PageBlockVerticalAlignment) GetExtra() string {
	return ""
}

func (t *PageBlockVerticalAlignment) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "pageBlockVerticalAlignmentBottom":
		t.PageBlockVerticalAlignmentBottom = new(PageBlockVerticalAlignmentBottom)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentBottom)
	case "pageBlockVerticalAlignmentMiddle":
		t.PageBlockVerticalAlignmentMiddle = new(PageBlockVerticalAlignmentMiddle)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentMiddle)
	case "pageBlockVerticalAlignmentTop":
		t.PageBlockVerticalAlignmentTop = new(PageBlockVerticalAlignmentTop)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentTop)
	}
	return nil
}

func (t *PageBlockVerticalAlignment) MarshalJSON() ([]byte, error) {
	if t.PageBlockVerticalAlignmentBottom != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentBottom)
	}
	if t.PageBlockVerticalAlignmentMiddle != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentMiddle)
	}
	if t.PageBlockVerticalAlignmentTop != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentTop)
	}
	return nil, fmt.Errorf("no value set for PageBlockVerticalAlignment")
}

// PaidMedia Describes a paid media
type PaidMedia struct {
	TypeStr              string                `json:"@type"`
	PaidMediaPhoto       *PaidMediaPhoto       `json:"paidMediaPhoto,omitempty"`
	PaidMediaPreview     *PaidMediaPreview     `json:"paidMediaPreview,omitempty"`
	PaidMediaUnsupported *PaidMediaUnsupported `json:"paidMediaUnsupported,omitempty"`
	PaidMediaVideo       *PaidMediaVideo       `json:"paidMediaVideo,omitempty"`
}

func (t *PaidMedia) Type() string {
	return t.TypeStr
}

func (t *PaidMedia) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PaidMedia) GetExtra() string {
	return ""
}

func (t *PaidMedia) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "paidMediaPhoto":
		t.PaidMediaPhoto = new(PaidMediaPhoto)
		return json.Unmarshal(b, t.PaidMediaPhoto)
	case "paidMediaPreview":
		t.PaidMediaPreview = new(PaidMediaPreview)
		return json.Unmarshal(b, t.PaidMediaPreview)
	case "paidMediaUnsupported":
		t.PaidMediaUnsupported = new(PaidMediaUnsupported)
		return json.Unmarshal(b, t.PaidMediaUnsupported)
	case "paidMediaVideo":
		t.PaidMediaVideo = new(PaidMediaVideo)
		return json.Unmarshal(b, t.PaidMediaVideo)
	}
	return nil
}

func (t *PaidMedia) MarshalJSON() ([]byte, error) {
	if t.PaidMediaPhoto != nil {
		return json.Marshal(t.PaidMediaPhoto)
	}
	if t.PaidMediaPreview != nil {
		return json.Marshal(t.PaidMediaPreview)
	}
	if t.PaidMediaUnsupported != nil {
		return json.Marshal(t.PaidMediaUnsupported)
	}
	if t.PaidMediaVideo != nil {
		return json.Marshal(t.PaidMediaVideo)
	}
	return nil, fmt.Errorf("no value set for PaidMedia")
}

// PaidReactionType Describes type of paid message reaction
type PaidReactionType struct {
	TypeStr                   string                     `json:"@type"`
	PaidReactionTypeAnonymous *PaidReactionTypeAnonymous `json:"paidReactionTypeAnonymous,omitempty"`
	PaidReactionTypeChat      *PaidReactionTypeChat      `json:"paidReactionTypeChat,omitempty"`
	PaidReactionTypeRegular   *PaidReactionTypeRegular   `json:"paidReactionTypeRegular,omitempty"`
}

func (t *PaidReactionType) Type() string {
	return t.TypeStr
}

func (t *PaidReactionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PaidReactionType) GetExtra() string {
	return ""
}

func (t *PaidReactionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "paidReactionTypeAnonymous":
		t.PaidReactionTypeAnonymous = new(PaidReactionTypeAnonymous)
		return json.Unmarshal(b, t.PaidReactionTypeAnonymous)
	case "paidReactionTypeChat":
		t.PaidReactionTypeChat = new(PaidReactionTypeChat)
		return json.Unmarshal(b, t.PaidReactionTypeChat)
	case "paidReactionTypeRegular":
		t.PaidReactionTypeRegular = new(PaidReactionTypeRegular)
		return json.Unmarshal(b, t.PaidReactionTypeRegular)
	}
	return nil
}

func (t *PaidReactionType) MarshalJSON() ([]byte, error) {
	if t.PaidReactionTypeAnonymous != nil {
		return json.Marshal(t.PaidReactionTypeAnonymous)
	}
	if t.PaidReactionTypeChat != nil {
		return json.Marshal(t.PaidReactionTypeChat)
	}
	if t.PaidReactionTypeRegular != nil {
		return json.Marshal(t.PaidReactionTypeRegular)
	}
	return nil, fmt.Errorf("no value set for PaidReactionType")
}

// PassportElement Contains information about a Telegram Passport element
type PassportElement struct {
	TypeStr                              string                                `json:"@type"`
	PassportElementAddress               *PassportElementAddress               `json:"passportElementAddress,omitempty"`
	PassportElementBankStatement         *PassportElementBankStatement         `json:"passportElementBankStatement,omitempty"`
	PassportElementDriverLicense         *PassportElementDriverLicense         `json:"passportElementDriverLicense,omitempty"`
	PassportElementEmailAddress          *PassportElementEmailAddress          `json:"passportElementEmailAddress,omitempty"`
	PassportElementIdentityCard          *PassportElementIdentityCard          `json:"passportElementIdentityCard,omitempty"`
	PassportElementInternalPassport      *PassportElementInternalPassport      `json:"passportElementInternalPassport,omitempty"`
	PassportElementPassport              *PassportElementPassport              `json:"passportElementPassport,omitempty"`
	PassportElementPassportRegistration  *PassportElementPassportRegistration  `json:"passportElementPassportRegistration,omitempty"`
	PassportElementPersonalDetails       *PassportElementPersonalDetails       `json:"passportElementPersonalDetails,omitempty"`
	PassportElementPhoneNumber           *PassportElementPhoneNumber           `json:"passportElementPhoneNumber,omitempty"`
	PassportElementRentalAgreement       *PassportElementRentalAgreement       `json:"passportElementRentalAgreement,omitempty"`
	PassportElementTemporaryRegistration *PassportElementTemporaryRegistration `json:"passportElementTemporaryRegistration,omitempty"`
	PassportElementUtilityBill           *PassportElementUtilityBill           `json:"passportElementUtilityBill,omitempty"`
}

func (t *PassportElement) Type() string {
	return t.TypeStr
}

func (t *PassportElement) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PassportElement) GetExtra() string {
	return ""
}

func (t *PassportElement) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "passportElementAddress":
		t.PassportElementAddress = new(PassportElementAddress)
		return json.Unmarshal(b, t.PassportElementAddress)
	case "passportElementBankStatement":
		t.PassportElementBankStatement = new(PassportElementBankStatement)
		return json.Unmarshal(b, t.PassportElementBankStatement)
	case "passportElementDriverLicense":
		t.PassportElementDriverLicense = new(PassportElementDriverLicense)
		return json.Unmarshal(b, t.PassportElementDriverLicense)
	case "passportElementEmailAddress":
		t.PassportElementEmailAddress = new(PassportElementEmailAddress)
		return json.Unmarshal(b, t.PassportElementEmailAddress)
	case "passportElementIdentityCard":
		t.PassportElementIdentityCard = new(PassportElementIdentityCard)
		return json.Unmarshal(b, t.PassportElementIdentityCard)
	case "passportElementInternalPassport":
		t.PassportElementInternalPassport = new(PassportElementInternalPassport)
		return json.Unmarshal(b, t.PassportElementInternalPassport)
	case "passportElementPassport":
		t.PassportElementPassport = new(PassportElementPassport)
		return json.Unmarshal(b, t.PassportElementPassport)
	case "passportElementPassportRegistration":
		t.PassportElementPassportRegistration = new(PassportElementPassportRegistration)
		return json.Unmarshal(b, t.PassportElementPassportRegistration)
	case "passportElementPersonalDetails":
		t.PassportElementPersonalDetails = new(PassportElementPersonalDetails)
		return json.Unmarshal(b, t.PassportElementPersonalDetails)
	case "passportElementPhoneNumber":
		t.PassportElementPhoneNumber = new(PassportElementPhoneNumber)
		return json.Unmarshal(b, t.PassportElementPhoneNumber)
	case "passportElementRentalAgreement":
		t.PassportElementRentalAgreement = new(PassportElementRentalAgreement)
		return json.Unmarshal(b, t.PassportElementRentalAgreement)
	case "passportElementTemporaryRegistration":
		t.PassportElementTemporaryRegistration = new(PassportElementTemporaryRegistration)
		return json.Unmarshal(b, t.PassportElementTemporaryRegistration)
	case "passportElementUtilityBill":
		t.PassportElementUtilityBill = new(PassportElementUtilityBill)
		return json.Unmarshal(b, t.PassportElementUtilityBill)
	}
	return nil
}

func (t *PassportElement) MarshalJSON() ([]byte, error) {
	if t.PassportElementAddress != nil {
		return json.Marshal(t.PassportElementAddress)
	}
	if t.PassportElementBankStatement != nil {
		return json.Marshal(t.PassportElementBankStatement)
	}
	if t.PassportElementDriverLicense != nil {
		return json.Marshal(t.PassportElementDriverLicense)
	}
	if t.PassportElementEmailAddress != nil {
		return json.Marshal(t.PassportElementEmailAddress)
	}
	if t.PassportElementIdentityCard != nil {
		return json.Marshal(t.PassportElementIdentityCard)
	}
	if t.PassportElementInternalPassport != nil {
		return json.Marshal(t.PassportElementInternalPassport)
	}
	if t.PassportElementPassport != nil {
		return json.Marshal(t.PassportElementPassport)
	}
	if t.PassportElementPassportRegistration != nil {
		return json.Marshal(t.PassportElementPassportRegistration)
	}
	if t.PassportElementPersonalDetails != nil {
		return json.Marshal(t.PassportElementPersonalDetails)
	}
	if t.PassportElementPhoneNumber != nil {
		return json.Marshal(t.PassportElementPhoneNumber)
	}
	if t.PassportElementRentalAgreement != nil {
		return json.Marshal(t.PassportElementRentalAgreement)
	}
	if t.PassportElementTemporaryRegistration != nil {
		return json.Marshal(t.PassportElementTemporaryRegistration)
	}
	if t.PassportElementUtilityBill != nil {
		return json.Marshal(t.PassportElementUtilityBill)
	}
	return nil, fmt.Errorf("no value set for PassportElement")
}

// PassportElementErrorSource Contains the description of an error in a Telegram Passport element
type PassportElementErrorSource struct {
	TypeStr                                    string                                      `json:"@type"`
	PassportElementErrorSourceDataField        *PassportElementErrorSourceDataField        `json:"passportElementErrorSourceDataField,omitempty"`
	PassportElementErrorSourceFile             *PassportElementErrorSourceFile             `json:"passportElementErrorSourceFile,omitempty"`
	PassportElementErrorSourceFiles            *PassportElementErrorSourceFiles            `json:"passportElementErrorSourceFiles,omitempty"`
	PassportElementErrorSourceFrontSide        *PassportElementErrorSourceFrontSide        `json:"passportElementErrorSourceFrontSide,omitempty"`
	PassportElementErrorSourceReverseSide      *PassportElementErrorSourceReverseSide      `json:"passportElementErrorSourceReverseSide,omitempty"`
	PassportElementErrorSourceSelfie           *PassportElementErrorSourceSelfie           `json:"passportElementErrorSourceSelfie,omitempty"`
	PassportElementErrorSourceTranslationFile  *PassportElementErrorSourceTranslationFile  `json:"passportElementErrorSourceTranslationFile,omitempty"`
	PassportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles `json:"passportElementErrorSourceTranslationFiles,omitempty"`
	PassportElementErrorSourceUnspecified      *PassportElementErrorSourceUnspecified      `json:"passportElementErrorSourceUnspecified,omitempty"`
}

func (t *PassportElementErrorSource) Type() string {
	return t.TypeStr
}

func (t *PassportElementErrorSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PassportElementErrorSource) GetExtra() string {
	return ""
}

func (t *PassportElementErrorSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "passportElementErrorSourceDataField":
		t.PassportElementErrorSourceDataField = new(PassportElementErrorSourceDataField)
		return json.Unmarshal(b, t.PassportElementErrorSourceDataField)
	case "passportElementErrorSourceFile":
		t.PassportElementErrorSourceFile = new(PassportElementErrorSourceFile)
		return json.Unmarshal(b, t.PassportElementErrorSourceFile)
	case "passportElementErrorSourceFiles":
		t.PassportElementErrorSourceFiles = new(PassportElementErrorSourceFiles)
		return json.Unmarshal(b, t.PassportElementErrorSourceFiles)
	case "passportElementErrorSourceFrontSide":
		t.PassportElementErrorSourceFrontSide = new(PassportElementErrorSourceFrontSide)
		return json.Unmarshal(b, t.PassportElementErrorSourceFrontSide)
	case "passportElementErrorSourceReverseSide":
		t.PassportElementErrorSourceReverseSide = new(PassportElementErrorSourceReverseSide)
		return json.Unmarshal(b, t.PassportElementErrorSourceReverseSide)
	case "passportElementErrorSourceSelfie":
		t.PassportElementErrorSourceSelfie = new(PassportElementErrorSourceSelfie)
		return json.Unmarshal(b, t.PassportElementErrorSourceSelfie)
	case "passportElementErrorSourceTranslationFile":
		t.PassportElementErrorSourceTranslationFile = new(PassportElementErrorSourceTranslationFile)
		return json.Unmarshal(b, t.PassportElementErrorSourceTranslationFile)
	case "passportElementErrorSourceTranslationFiles":
		t.PassportElementErrorSourceTranslationFiles = new(PassportElementErrorSourceTranslationFiles)
		return json.Unmarshal(b, t.PassportElementErrorSourceTranslationFiles)
	case "passportElementErrorSourceUnspecified":
		t.PassportElementErrorSourceUnspecified = new(PassportElementErrorSourceUnspecified)
		return json.Unmarshal(b, t.PassportElementErrorSourceUnspecified)
	}
	return nil
}

func (t *PassportElementErrorSource) MarshalJSON() ([]byte, error) {
	if t.PassportElementErrorSourceDataField != nil {
		return json.Marshal(t.PassportElementErrorSourceDataField)
	}
	if t.PassportElementErrorSourceFile != nil {
		return json.Marshal(t.PassportElementErrorSourceFile)
	}
	if t.PassportElementErrorSourceFiles != nil {
		return json.Marshal(t.PassportElementErrorSourceFiles)
	}
	if t.PassportElementErrorSourceFrontSide != nil {
		return json.Marshal(t.PassportElementErrorSourceFrontSide)
	}
	if t.PassportElementErrorSourceReverseSide != nil {
		return json.Marshal(t.PassportElementErrorSourceReverseSide)
	}
	if t.PassportElementErrorSourceSelfie != nil {
		return json.Marshal(t.PassportElementErrorSourceSelfie)
	}
	if t.PassportElementErrorSourceTranslationFile != nil {
		return json.Marshal(t.PassportElementErrorSourceTranslationFile)
	}
	if t.PassportElementErrorSourceTranslationFiles != nil {
		return json.Marshal(t.PassportElementErrorSourceTranslationFiles)
	}
	if t.PassportElementErrorSourceUnspecified != nil {
		return json.Marshal(t.PassportElementErrorSourceUnspecified)
	}
	return nil, fmt.Errorf("no value set for PassportElementErrorSource")
}

// PassportElementType Contains the type of Telegram Passport element
type PassportElementType struct {
	TypeStr                                  string                                    `json:"@type"`
	PassportElementTypeAddress               *PassportElementTypeAddress               `json:"passportElementTypeAddress,omitempty"`
	PassportElementTypeBankStatement         *PassportElementTypeBankStatement         `json:"passportElementTypeBankStatement,omitempty"`
	PassportElementTypeDriverLicense         *PassportElementTypeDriverLicense         `json:"passportElementTypeDriverLicense,omitempty"`
	PassportElementTypeEmailAddress          *PassportElementTypeEmailAddress          `json:"passportElementTypeEmailAddress,omitempty"`
	PassportElementTypeIdentityCard          *PassportElementTypeIdentityCard          `json:"passportElementTypeIdentityCard,omitempty"`
	PassportElementTypeInternalPassport      *PassportElementTypeInternalPassport      `json:"passportElementTypeInternalPassport,omitempty"`
	PassportElementTypePassport              *PassportElementTypePassport              `json:"passportElementTypePassport,omitempty"`
	PassportElementTypePassportRegistration  *PassportElementTypePassportRegistration  `json:"passportElementTypePassportRegistration,omitempty"`
	PassportElementTypePersonalDetails       *PassportElementTypePersonalDetails       `json:"passportElementTypePersonalDetails,omitempty"`
	PassportElementTypePhoneNumber           *PassportElementTypePhoneNumber           `json:"passportElementTypePhoneNumber,omitempty"`
	PassportElementTypeRentalAgreement       *PassportElementTypeRentalAgreement       `json:"passportElementTypeRentalAgreement,omitempty"`
	PassportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration `json:"passportElementTypeTemporaryRegistration,omitempty"`
	PassportElementTypeUtilityBill           *PassportElementTypeUtilityBill           `json:"passportElementTypeUtilityBill,omitempty"`
}

func (t *PassportElementType) Type() string {
	return t.TypeStr
}

func (t *PassportElementType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PassportElementType) GetExtra() string {
	return ""
}

func (t *PassportElementType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "passportElementTypeAddress":
		t.PassportElementTypeAddress = new(PassportElementTypeAddress)
		return json.Unmarshal(b, t.PassportElementTypeAddress)
	case "passportElementTypeBankStatement":
		t.PassportElementTypeBankStatement = new(PassportElementTypeBankStatement)
		return json.Unmarshal(b, t.PassportElementTypeBankStatement)
	case "passportElementTypeDriverLicense":
		t.PassportElementTypeDriverLicense = new(PassportElementTypeDriverLicense)
		return json.Unmarshal(b, t.PassportElementTypeDriverLicense)
	case "passportElementTypeEmailAddress":
		t.PassportElementTypeEmailAddress = new(PassportElementTypeEmailAddress)
		return json.Unmarshal(b, t.PassportElementTypeEmailAddress)
	case "passportElementTypeIdentityCard":
		t.PassportElementTypeIdentityCard = new(PassportElementTypeIdentityCard)
		return json.Unmarshal(b, t.PassportElementTypeIdentityCard)
	case "passportElementTypeInternalPassport":
		t.PassportElementTypeInternalPassport = new(PassportElementTypeInternalPassport)
		return json.Unmarshal(b, t.PassportElementTypeInternalPassport)
	case "passportElementTypePassport":
		t.PassportElementTypePassport = new(PassportElementTypePassport)
		return json.Unmarshal(b, t.PassportElementTypePassport)
	case "passportElementTypePassportRegistration":
		t.PassportElementTypePassportRegistration = new(PassportElementTypePassportRegistration)
		return json.Unmarshal(b, t.PassportElementTypePassportRegistration)
	case "passportElementTypePersonalDetails":
		t.PassportElementTypePersonalDetails = new(PassportElementTypePersonalDetails)
		return json.Unmarshal(b, t.PassportElementTypePersonalDetails)
	case "passportElementTypePhoneNumber":
		t.PassportElementTypePhoneNumber = new(PassportElementTypePhoneNumber)
		return json.Unmarshal(b, t.PassportElementTypePhoneNumber)
	case "passportElementTypeRentalAgreement":
		t.PassportElementTypeRentalAgreement = new(PassportElementTypeRentalAgreement)
		return json.Unmarshal(b, t.PassportElementTypeRentalAgreement)
	case "passportElementTypeTemporaryRegistration":
		t.PassportElementTypeTemporaryRegistration = new(PassportElementTypeTemporaryRegistration)
		return json.Unmarshal(b, t.PassportElementTypeTemporaryRegistration)
	case "passportElementTypeUtilityBill":
		t.PassportElementTypeUtilityBill = new(PassportElementTypeUtilityBill)
		return json.Unmarshal(b, t.PassportElementTypeUtilityBill)
	}
	return nil
}

func (t *PassportElementType) MarshalJSON() ([]byte, error) {
	if t.PassportElementTypeAddress != nil {
		return json.Marshal(t.PassportElementTypeAddress)
	}
	if t.PassportElementTypeBankStatement != nil {
		return json.Marshal(t.PassportElementTypeBankStatement)
	}
	if t.PassportElementTypeDriverLicense != nil {
		return json.Marshal(t.PassportElementTypeDriverLicense)
	}
	if t.PassportElementTypeEmailAddress != nil {
		return json.Marshal(t.PassportElementTypeEmailAddress)
	}
	if t.PassportElementTypeIdentityCard != nil {
		return json.Marshal(t.PassportElementTypeIdentityCard)
	}
	if t.PassportElementTypeInternalPassport != nil {
		return json.Marshal(t.PassportElementTypeInternalPassport)
	}
	if t.PassportElementTypePassport != nil {
		return json.Marshal(t.PassportElementTypePassport)
	}
	if t.PassportElementTypePassportRegistration != nil {
		return json.Marshal(t.PassportElementTypePassportRegistration)
	}
	if t.PassportElementTypePersonalDetails != nil {
		return json.Marshal(t.PassportElementTypePersonalDetails)
	}
	if t.PassportElementTypePhoneNumber != nil {
		return json.Marshal(t.PassportElementTypePhoneNumber)
	}
	if t.PassportElementTypeRentalAgreement != nil {
		return json.Marshal(t.PassportElementTypeRentalAgreement)
	}
	if t.PassportElementTypeTemporaryRegistration != nil {
		return json.Marshal(t.PassportElementTypeTemporaryRegistration)
	}
	if t.PassportElementTypeUtilityBill != nil {
		return json.Marshal(t.PassportElementTypeUtilityBill)
	}
	return nil, fmt.Errorf("no value set for PassportElementType")
}

// PaymentFormType Describes type of payment form
type PaymentFormType struct {
	TypeStr                         string                           `json:"@type"`
	PaymentFormTypeRegular          *PaymentFormTypeRegular          `json:"paymentFormTypeRegular,omitempty"`
	PaymentFormTypeStarSubscription *PaymentFormTypeStarSubscription `json:"paymentFormTypeStarSubscription,omitempty"`
	PaymentFormTypeStars            *PaymentFormTypeStars            `json:"paymentFormTypeStars,omitempty"`
}

func (t *PaymentFormType) Type() string {
	return t.TypeStr
}

func (t *PaymentFormType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PaymentFormType) GetExtra() string {
	return ""
}

func (t *PaymentFormType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "paymentFormTypeRegular":
		t.PaymentFormTypeRegular = new(PaymentFormTypeRegular)
		return json.Unmarshal(b, t.PaymentFormTypeRegular)
	case "paymentFormTypeStarSubscription":
		t.PaymentFormTypeStarSubscription = new(PaymentFormTypeStarSubscription)
		return json.Unmarshal(b, t.PaymentFormTypeStarSubscription)
	case "paymentFormTypeStars":
		t.PaymentFormTypeStars = new(PaymentFormTypeStars)
		return json.Unmarshal(b, t.PaymentFormTypeStars)
	}
	return nil
}

func (t *PaymentFormType) MarshalJSON() ([]byte, error) {
	if t.PaymentFormTypeRegular != nil {
		return json.Marshal(t.PaymentFormTypeRegular)
	}
	if t.PaymentFormTypeStarSubscription != nil {
		return json.Marshal(t.PaymentFormTypeStarSubscription)
	}
	if t.PaymentFormTypeStars != nil {
		return json.Marshal(t.PaymentFormTypeStars)
	}
	return nil, fmt.Errorf("no value set for PaymentFormType")
}

// PaymentProvider Contains information about a payment provider
type PaymentProvider struct {
	TypeStr                    string                      `json:"@type"`
	PaymentProviderOther       *PaymentProviderOther       `json:"paymentProviderOther,omitempty"`
	PaymentProviderSmartGlocal *PaymentProviderSmartGlocal `json:"paymentProviderSmartGlocal,omitempty"`
	PaymentProviderStripe      *PaymentProviderStripe      `json:"paymentProviderStripe,omitempty"`
}

func (t *PaymentProvider) Type() string {
	return t.TypeStr
}

func (t *PaymentProvider) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PaymentProvider) GetExtra() string {
	return ""
}

func (t *PaymentProvider) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "paymentProviderOther":
		t.PaymentProviderOther = new(PaymentProviderOther)
		return json.Unmarshal(b, t.PaymentProviderOther)
	case "paymentProviderSmartGlocal":
		t.PaymentProviderSmartGlocal = new(PaymentProviderSmartGlocal)
		return json.Unmarshal(b, t.PaymentProviderSmartGlocal)
	case "paymentProviderStripe":
		t.PaymentProviderStripe = new(PaymentProviderStripe)
		return json.Unmarshal(b, t.PaymentProviderStripe)
	}
	return nil
}

func (t *PaymentProvider) MarshalJSON() ([]byte, error) {
	if t.PaymentProviderOther != nil {
		return json.Marshal(t.PaymentProviderOther)
	}
	if t.PaymentProviderSmartGlocal != nil {
		return json.Marshal(t.PaymentProviderSmartGlocal)
	}
	if t.PaymentProviderStripe != nil {
		return json.Marshal(t.PaymentProviderStripe)
	}
	return nil, fmt.Errorf("no value set for PaymentProvider")
}

// PaymentReceiptType Describes type of successful payment
type PaymentReceiptType struct {
	TypeStr                   string                     `json:"@type"`
	PaymentReceiptTypeRegular *PaymentReceiptTypeRegular `json:"paymentReceiptTypeRegular,omitempty"`
	PaymentReceiptTypeStars   *PaymentReceiptTypeStars   `json:"paymentReceiptTypeStars,omitempty"`
}

func (t *PaymentReceiptType) Type() string {
	return t.TypeStr
}

func (t *PaymentReceiptType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PaymentReceiptType) GetExtra() string {
	return ""
}

func (t *PaymentReceiptType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "paymentReceiptTypeRegular":
		t.PaymentReceiptTypeRegular = new(PaymentReceiptTypeRegular)
		return json.Unmarshal(b, t.PaymentReceiptTypeRegular)
	case "paymentReceiptTypeStars":
		t.PaymentReceiptTypeStars = new(PaymentReceiptTypeStars)
		return json.Unmarshal(b, t.PaymentReceiptTypeStars)
	}
	return nil
}

func (t *PaymentReceiptType) MarshalJSON() ([]byte, error) {
	if t.PaymentReceiptTypeRegular != nil {
		return json.Marshal(t.PaymentReceiptTypeRegular)
	}
	if t.PaymentReceiptTypeStars != nil {
		return json.Marshal(t.PaymentReceiptTypeStars)
	}
	return nil, fmt.Errorf("no value set for PaymentReceiptType")
}

// PhoneNumberCodeType Describes type of the request for which a code is sent to a phone number
type PhoneNumberCodeType struct {
	TypeStr                             string                               `json:"@type"`
	PhoneNumberCodeTypeChange           *PhoneNumberCodeTypeChange           `json:"phoneNumberCodeTypeChange,omitempty"`
	PhoneNumberCodeTypeConfirmOwnership *PhoneNumberCodeTypeConfirmOwnership `json:"phoneNumberCodeTypeConfirmOwnership,omitempty"`
	PhoneNumberCodeTypeVerify           *PhoneNumberCodeTypeVerify           `json:"phoneNumberCodeTypeVerify,omitempty"`
}

func (t *PhoneNumberCodeType) Type() string {
	return t.TypeStr
}

func (t *PhoneNumberCodeType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PhoneNumberCodeType) GetExtra() string {
	return ""
}

func (t *PhoneNumberCodeType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "phoneNumberCodeTypeChange":
		t.PhoneNumberCodeTypeChange = new(PhoneNumberCodeTypeChange)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeChange)
	case "phoneNumberCodeTypeConfirmOwnership":
		t.PhoneNumberCodeTypeConfirmOwnership = new(PhoneNumberCodeTypeConfirmOwnership)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeConfirmOwnership)
	case "phoneNumberCodeTypeVerify":
		t.PhoneNumberCodeTypeVerify = new(PhoneNumberCodeTypeVerify)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeVerify)
	}
	return nil
}

func (t *PhoneNumberCodeType) MarshalJSON() ([]byte, error) {
	if t.PhoneNumberCodeTypeChange != nil {
		return json.Marshal(t.PhoneNumberCodeTypeChange)
	}
	if t.PhoneNumberCodeTypeConfirmOwnership != nil {
		return json.Marshal(t.PhoneNumberCodeTypeConfirmOwnership)
	}
	if t.PhoneNumberCodeTypeVerify != nil {
		return json.Marshal(t.PhoneNumberCodeTypeVerify)
	}
	return nil, fmt.Errorf("no value set for PhoneNumberCodeType")
}

// PollType Describes the type of poll
type PollType struct {
	TypeStr         string           `json:"@type"`
	PollTypeQuiz    *PollTypeQuiz    `json:"pollTypeQuiz,omitempty"`
	PollTypeRegular *PollTypeRegular `json:"pollTypeRegular,omitempty"`
}

func (t *PollType) Type() string {
	return t.TypeStr
}

func (t *PollType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PollType) GetExtra() string {
	return ""
}

func (t *PollType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "pollTypeQuiz":
		t.PollTypeQuiz = new(PollTypeQuiz)
		return json.Unmarshal(b, t.PollTypeQuiz)
	case "pollTypeRegular":
		t.PollTypeRegular = new(PollTypeRegular)
		return json.Unmarshal(b, t.PollTypeRegular)
	}
	return nil
}

func (t *PollType) MarshalJSON() ([]byte, error) {
	if t.PollTypeQuiz != nil {
		return json.Marshal(t.PollTypeQuiz)
	}
	if t.PollTypeRegular != nil {
		return json.Marshal(t.PollTypeRegular)
	}
	return nil, fmt.Errorf("no value set for PollType")
}

// PremiumFeature Describes a feature available to Premium users
type PremiumFeature struct {
	TypeStr                               string                                 `json:"@type"`
	PremiumFeatureAccentColor             *PremiumFeatureAccentColor             `json:"premiumFeatureAccentColor,omitempty"`
	PremiumFeatureAdvancedChatManagement  *PremiumFeatureAdvancedChatManagement  `json:"premiumFeatureAdvancedChatManagement,omitempty"`
	PremiumFeatureAnimatedProfilePhoto    *PremiumFeatureAnimatedProfilePhoto    `json:"premiumFeatureAnimatedProfilePhoto,omitempty"`
	PremiumFeatureAppIcons                *PremiumFeatureAppIcons                `json:"premiumFeatureAppIcons,omitempty"`
	PremiumFeatureBackgroundForBoth       *PremiumFeatureBackgroundForBoth       `json:"premiumFeatureBackgroundForBoth,omitempty"`
	PremiumFeatureBusiness                *PremiumFeatureBusiness                `json:"premiumFeatureBusiness,omitempty"`
	PremiumFeatureChatBoost               *PremiumFeatureChatBoost               `json:"premiumFeatureChatBoost,omitempty"`
	PremiumFeatureChecklists              *PremiumFeatureChecklists              `json:"premiumFeatureChecklists,omitempty"`
	PremiumFeatureCustomEmoji             *PremiumFeatureCustomEmoji             `json:"premiumFeatureCustomEmoji,omitempty"`
	PremiumFeatureDisabledAds             *PremiumFeatureDisabledAds             `json:"premiumFeatureDisabledAds,omitempty"`
	PremiumFeatureEmojiStatus             *PremiumFeatureEmojiStatus             `json:"premiumFeatureEmojiStatus,omitempty"`
	PremiumFeatureForumTopicIcon          *PremiumFeatureForumTopicIcon          `json:"premiumFeatureForumTopicIcon,omitempty"`
	PremiumFeatureImprovedDownloadSpeed   *PremiumFeatureImprovedDownloadSpeed   `json:"premiumFeatureImprovedDownloadSpeed,omitempty"`
	PremiumFeatureIncreasedLimits         *PremiumFeatureIncreasedLimits         `json:"premiumFeatureIncreasedLimits,omitempty"`
	PremiumFeatureIncreasedUploadFileSize *PremiumFeatureIncreasedUploadFileSize `json:"premiumFeatureIncreasedUploadFileSize,omitempty"`
	PremiumFeatureLastSeenTimes           *PremiumFeatureLastSeenTimes           `json:"premiumFeatureLastSeenTimes,omitempty"`
	PremiumFeatureMessageEffects          *PremiumFeatureMessageEffects          `json:"premiumFeatureMessageEffects,omitempty"`
	PremiumFeatureMessagePrivacy          *PremiumFeatureMessagePrivacy          `json:"premiumFeatureMessagePrivacy,omitempty"`
	PremiumFeaturePaidMessages            *PremiumFeaturePaidMessages            `json:"premiumFeaturePaidMessages,omitempty"`
	PremiumFeatureProfileBadge            *PremiumFeatureProfileBadge            `json:"premiumFeatureProfileBadge,omitempty"`
	PremiumFeatureRealTimeChatTranslation *PremiumFeatureRealTimeChatTranslation `json:"premiumFeatureRealTimeChatTranslation,omitempty"`
	PremiumFeatureSavedMessagesTags       *PremiumFeatureSavedMessagesTags       `json:"premiumFeatureSavedMessagesTags,omitempty"`
	PremiumFeatureUniqueReactions         *PremiumFeatureUniqueReactions         `json:"premiumFeatureUniqueReactions,omitempty"`
	PremiumFeatureUniqueStickers          *PremiumFeatureUniqueStickers          `json:"premiumFeatureUniqueStickers,omitempty"`
	PremiumFeatureUpgradedStories         *PremiumFeatureUpgradedStories         `json:"premiumFeatureUpgradedStories,omitempty"`
	PremiumFeatureVoiceRecognition        *PremiumFeatureVoiceRecognition        `json:"premiumFeatureVoiceRecognition,omitempty"`
}

func (t *PremiumFeature) Type() string {
	return t.TypeStr
}

func (t *PremiumFeature) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PremiumFeature) GetExtra() string {
	return ""
}

func (t *PremiumFeature) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "premiumFeatureAccentColor":
		t.PremiumFeatureAccentColor = new(PremiumFeatureAccentColor)
		return json.Unmarshal(b, t.PremiumFeatureAccentColor)
	case "premiumFeatureAdvancedChatManagement":
		t.PremiumFeatureAdvancedChatManagement = new(PremiumFeatureAdvancedChatManagement)
		return json.Unmarshal(b, t.PremiumFeatureAdvancedChatManagement)
	case "premiumFeatureAnimatedProfilePhoto":
		t.PremiumFeatureAnimatedProfilePhoto = new(PremiumFeatureAnimatedProfilePhoto)
		return json.Unmarshal(b, t.PremiumFeatureAnimatedProfilePhoto)
	case "premiumFeatureAppIcons":
		t.PremiumFeatureAppIcons = new(PremiumFeatureAppIcons)
		return json.Unmarshal(b, t.PremiumFeatureAppIcons)
	case "premiumFeatureBackgroundForBoth":
		t.PremiumFeatureBackgroundForBoth = new(PremiumFeatureBackgroundForBoth)
		return json.Unmarshal(b, t.PremiumFeatureBackgroundForBoth)
	case "premiumFeatureBusiness":
		t.PremiumFeatureBusiness = new(PremiumFeatureBusiness)
		return json.Unmarshal(b, t.PremiumFeatureBusiness)
	case "premiumFeatureChatBoost":
		t.PremiumFeatureChatBoost = new(PremiumFeatureChatBoost)
		return json.Unmarshal(b, t.PremiumFeatureChatBoost)
	case "premiumFeatureChecklists":
		t.PremiumFeatureChecklists = new(PremiumFeatureChecklists)
		return json.Unmarshal(b, t.PremiumFeatureChecklists)
	case "premiumFeatureCustomEmoji":
		t.PremiumFeatureCustomEmoji = new(PremiumFeatureCustomEmoji)
		return json.Unmarshal(b, t.PremiumFeatureCustomEmoji)
	case "premiumFeatureDisabledAds":
		t.PremiumFeatureDisabledAds = new(PremiumFeatureDisabledAds)
		return json.Unmarshal(b, t.PremiumFeatureDisabledAds)
	case "premiumFeatureEmojiStatus":
		t.PremiumFeatureEmojiStatus = new(PremiumFeatureEmojiStatus)
		return json.Unmarshal(b, t.PremiumFeatureEmojiStatus)
	case "premiumFeatureForumTopicIcon":
		t.PremiumFeatureForumTopicIcon = new(PremiumFeatureForumTopicIcon)
		return json.Unmarshal(b, t.PremiumFeatureForumTopicIcon)
	case "premiumFeatureImprovedDownloadSpeed":
		t.PremiumFeatureImprovedDownloadSpeed = new(PremiumFeatureImprovedDownloadSpeed)
		return json.Unmarshal(b, t.PremiumFeatureImprovedDownloadSpeed)
	case "premiumFeatureIncreasedLimits":
		t.PremiumFeatureIncreasedLimits = new(PremiumFeatureIncreasedLimits)
		return json.Unmarshal(b, t.PremiumFeatureIncreasedLimits)
	case "premiumFeatureIncreasedUploadFileSize":
		t.PremiumFeatureIncreasedUploadFileSize = new(PremiumFeatureIncreasedUploadFileSize)
		return json.Unmarshal(b, t.PremiumFeatureIncreasedUploadFileSize)
	case "premiumFeatureLastSeenTimes":
		t.PremiumFeatureLastSeenTimes = new(PremiumFeatureLastSeenTimes)
		return json.Unmarshal(b, t.PremiumFeatureLastSeenTimes)
	case "premiumFeatureMessageEffects":
		t.PremiumFeatureMessageEffects = new(PremiumFeatureMessageEffects)
		return json.Unmarshal(b, t.PremiumFeatureMessageEffects)
	case "premiumFeatureMessagePrivacy":
		t.PremiumFeatureMessagePrivacy = new(PremiumFeatureMessagePrivacy)
		return json.Unmarshal(b, t.PremiumFeatureMessagePrivacy)
	case "premiumFeaturePaidMessages":
		t.PremiumFeaturePaidMessages = new(PremiumFeaturePaidMessages)
		return json.Unmarshal(b, t.PremiumFeaturePaidMessages)
	case "premiumFeatureProfileBadge":
		t.PremiumFeatureProfileBadge = new(PremiumFeatureProfileBadge)
		return json.Unmarshal(b, t.PremiumFeatureProfileBadge)
	case "premiumFeatureRealTimeChatTranslation":
		t.PremiumFeatureRealTimeChatTranslation = new(PremiumFeatureRealTimeChatTranslation)
		return json.Unmarshal(b, t.PremiumFeatureRealTimeChatTranslation)
	case "premiumFeatureSavedMessagesTags":
		t.PremiumFeatureSavedMessagesTags = new(PremiumFeatureSavedMessagesTags)
		return json.Unmarshal(b, t.PremiumFeatureSavedMessagesTags)
	case "premiumFeatureUniqueReactions":
		t.PremiumFeatureUniqueReactions = new(PremiumFeatureUniqueReactions)
		return json.Unmarshal(b, t.PremiumFeatureUniqueReactions)
	case "premiumFeatureUniqueStickers":
		t.PremiumFeatureUniqueStickers = new(PremiumFeatureUniqueStickers)
		return json.Unmarshal(b, t.PremiumFeatureUniqueStickers)
	case "premiumFeatureUpgradedStories":
		t.PremiumFeatureUpgradedStories = new(PremiumFeatureUpgradedStories)
		return json.Unmarshal(b, t.PremiumFeatureUpgradedStories)
	case "premiumFeatureVoiceRecognition":
		t.PremiumFeatureVoiceRecognition = new(PremiumFeatureVoiceRecognition)
		return json.Unmarshal(b, t.PremiumFeatureVoiceRecognition)
	}
	return nil
}

func (t *PremiumFeature) MarshalJSON() ([]byte, error) {
	if t.PremiumFeatureAccentColor != nil {
		return json.Marshal(t.PremiumFeatureAccentColor)
	}
	if t.PremiumFeatureAdvancedChatManagement != nil {
		return json.Marshal(t.PremiumFeatureAdvancedChatManagement)
	}
	if t.PremiumFeatureAnimatedProfilePhoto != nil {
		return json.Marshal(t.PremiumFeatureAnimatedProfilePhoto)
	}
	if t.PremiumFeatureAppIcons != nil {
		return json.Marshal(t.PremiumFeatureAppIcons)
	}
	if t.PremiumFeatureBackgroundForBoth != nil {
		return json.Marshal(t.PremiumFeatureBackgroundForBoth)
	}
	if t.PremiumFeatureBusiness != nil {
		return json.Marshal(t.PremiumFeatureBusiness)
	}
	if t.PremiumFeatureChatBoost != nil {
		return json.Marshal(t.PremiumFeatureChatBoost)
	}
	if t.PremiumFeatureChecklists != nil {
		return json.Marshal(t.PremiumFeatureChecklists)
	}
	if t.PremiumFeatureCustomEmoji != nil {
		return json.Marshal(t.PremiumFeatureCustomEmoji)
	}
	if t.PremiumFeatureDisabledAds != nil {
		return json.Marshal(t.PremiumFeatureDisabledAds)
	}
	if t.PremiumFeatureEmojiStatus != nil {
		return json.Marshal(t.PremiumFeatureEmojiStatus)
	}
	if t.PremiumFeatureForumTopicIcon != nil {
		return json.Marshal(t.PremiumFeatureForumTopicIcon)
	}
	if t.PremiumFeatureImprovedDownloadSpeed != nil {
		return json.Marshal(t.PremiumFeatureImprovedDownloadSpeed)
	}
	if t.PremiumFeatureIncreasedLimits != nil {
		return json.Marshal(t.PremiumFeatureIncreasedLimits)
	}
	if t.PremiumFeatureIncreasedUploadFileSize != nil {
		return json.Marshal(t.PremiumFeatureIncreasedUploadFileSize)
	}
	if t.PremiumFeatureLastSeenTimes != nil {
		return json.Marshal(t.PremiumFeatureLastSeenTimes)
	}
	if t.PremiumFeatureMessageEffects != nil {
		return json.Marshal(t.PremiumFeatureMessageEffects)
	}
	if t.PremiumFeatureMessagePrivacy != nil {
		return json.Marshal(t.PremiumFeatureMessagePrivacy)
	}
	if t.PremiumFeaturePaidMessages != nil {
		return json.Marshal(t.PremiumFeaturePaidMessages)
	}
	if t.PremiumFeatureProfileBadge != nil {
		return json.Marshal(t.PremiumFeatureProfileBadge)
	}
	if t.PremiumFeatureRealTimeChatTranslation != nil {
		return json.Marshal(t.PremiumFeatureRealTimeChatTranslation)
	}
	if t.PremiumFeatureSavedMessagesTags != nil {
		return json.Marshal(t.PremiumFeatureSavedMessagesTags)
	}
	if t.PremiumFeatureUniqueReactions != nil {
		return json.Marshal(t.PremiumFeatureUniqueReactions)
	}
	if t.PremiumFeatureUniqueStickers != nil {
		return json.Marshal(t.PremiumFeatureUniqueStickers)
	}
	if t.PremiumFeatureUpgradedStories != nil {
		return json.Marshal(t.PremiumFeatureUpgradedStories)
	}
	if t.PremiumFeatureVoiceRecognition != nil {
		return json.Marshal(t.PremiumFeatureVoiceRecognition)
	}
	return nil, fmt.Errorf("no value set for PremiumFeature")
}

// PremiumLimitType Describes type of limit, increased for Premium users
type PremiumLimitType struct {
	TypeStr                                         string                                           `json:"@type"`
	PremiumLimitTypeActiveStoryCount                *PremiumLimitTypeActiveStoryCount                `json:"premiumLimitTypeActiveStoryCount,omitempty"`
	PremiumLimitTypeBioLength                       *PremiumLimitTypeBioLength                       `json:"premiumLimitTypeBioLength,omitempty"`
	PremiumLimitTypeCaptionLength                   *PremiumLimitTypeCaptionLength                   `json:"premiumLimitTypeCaptionLength,omitempty"`
	PremiumLimitTypeChatFolderChosenChatCount       *PremiumLimitTypeChatFolderChosenChatCount       `json:"premiumLimitTypeChatFolderChosenChatCount,omitempty"`
	PremiumLimitTypeChatFolderCount                 *PremiumLimitTypeChatFolderCount                 `json:"premiumLimitTypeChatFolderCount,omitempty"`
	PremiumLimitTypeChatFolderInviteLinkCount       *PremiumLimitTypeChatFolderInviteLinkCount       `json:"premiumLimitTypeChatFolderInviteLinkCount,omitempty"`
	PremiumLimitTypeCreatedPublicChatCount          *PremiumLimitTypeCreatedPublicChatCount          `json:"premiumLimitTypeCreatedPublicChatCount,omitempty"`
	PremiumLimitTypeFavoriteStickerCount            *PremiumLimitTypeFavoriteStickerCount            `json:"premiumLimitTypeFavoriteStickerCount,omitempty"`
	PremiumLimitTypeMonthlyPostedStoryCount         *PremiumLimitTypeMonthlyPostedStoryCount         `json:"premiumLimitTypeMonthlyPostedStoryCount,omitempty"`
	PremiumLimitTypePinnedArchivedChatCount         *PremiumLimitTypePinnedArchivedChatCount         `json:"premiumLimitTypePinnedArchivedChatCount,omitempty"`
	PremiumLimitTypePinnedChatCount                 *PremiumLimitTypePinnedChatCount                 `json:"premiumLimitTypePinnedChatCount,omitempty"`
	PremiumLimitTypePinnedSavedMessagesTopicCount   *PremiumLimitTypePinnedSavedMessagesTopicCount   `json:"premiumLimitTypePinnedSavedMessagesTopicCount,omitempty"`
	PremiumLimitTypeSavedAnimationCount             *PremiumLimitTypeSavedAnimationCount             `json:"premiumLimitTypeSavedAnimationCount,omitempty"`
	PremiumLimitTypeShareableChatFolderCount        *PremiumLimitTypeShareableChatFolderCount        `json:"premiumLimitTypeShareableChatFolderCount,omitempty"`
	PremiumLimitTypeSimilarChatCount                *PremiumLimitTypeSimilarChatCount                `json:"premiumLimitTypeSimilarChatCount,omitempty"`
	PremiumLimitTypeStoryCaptionLength              *PremiumLimitTypeStoryCaptionLength              `json:"premiumLimitTypeStoryCaptionLength,omitempty"`
	PremiumLimitTypeStorySuggestedReactionAreaCount *PremiumLimitTypeStorySuggestedReactionAreaCount `json:"premiumLimitTypeStorySuggestedReactionAreaCount,omitempty"`
	PremiumLimitTypeSupergroupCount                 *PremiumLimitTypeSupergroupCount                 `json:"premiumLimitTypeSupergroupCount,omitempty"`
	PremiumLimitTypeWeeklyPostedStoryCount          *PremiumLimitTypeWeeklyPostedStoryCount          `json:"premiumLimitTypeWeeklyPostedStoryCount,omitempty"`
}

func (t *PremiumLimitType) Type() string {
	return t.TypeStr
}

func (t *PremiumLimitType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PremiumLimitType) GetExtra() string {
	return ""
}

func (t *PremiumLimitType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "premiumLimitTypeActiveStoryCount":
		t.PremiumLimitTypeActiveStoryCount = new(PremiumLimitTypeActiveStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeActiveStoryCount)
	case "premiumLimitTypeBioLength":
		t.PremiumLimitTypeBioLength = new(PremiumLimitTypeBioLength)
		return json.Unmarshal(b, t.PremiumLimitTypeBioLength)
	case "premiumLimitTypeCaptionLength":
		t.PremiumLimitTypeCaptionLength = new(PremiumLimitTypeCaptionLength)
		return json.Unmarshal(b, t.PremiumLimitTypeCaptionLength)
	case "premiumLimitTypeChatFolderChosenChatCount":
		t.PremiumLimitTypeChatFolderChosenChatCount = new(PremiumLimitTypeChatFolderChosenChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderChosenChatCount)
	case "premiumLimitTypeChatFolderCount":
		t.PremiumLimitTypeChatFolderCount = new(PremiumLimitTypeChatFolderCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderCount)
	case "premiumLimitTypeChatFolderInviteLinkCount":
		t.PremiumLimitTypeChatFolderInviteLinkCount = new(PremiumLimitTypeChatFolderInviteLinkCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderInviteLinkCount)
	case "premiumLimitTypeCreatedPublicChatCount":
		t.PremiumLimitTypeCreatedPublicChatCount = new(PremiumLimitTypeCreatedPublicChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeCreatedPublicChatCount)
	case "premiumLimitTypeFavoriteStickerCount":
		t.PremiumLimitTypeFavoriteStickerCount = new(PremiumLimitTypeFavoriteStickerCount)
		return json.Unmarshal(b, t.PremiumLimitTypeFavoriteStickerCount)
	case "premiumLimitTypeMonthlyPostedStoryCount":
		t.PremiumLimitTypeMonthlyPostedStoryCount = new(PremiumLimitTypeMonthlyPostedStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeMonthlyPostedStoryCount)
	case "premiumLimitTypePinnedArchivedChatCount":
		t.PremiumLimitTypePinnedArchivedChatCount = new(PremiumLimitTypePinnedArchivedChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedArchivedChatCount)
	case "premiumLimitTypePinnedChatCount":
		t.PremiumLimitTypePinnedChatCount = new(PremiumLimitTypePinnedChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedChatCount)
	case "premiumLimitTypePinnedSavedMessagesTopicCount":
		t.PremiumLimitTypePinnedSavedMessagesTopicCount = new(PremiumLimitTypePinnedSavedMessagesTopicCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedSavedMessagesTopicCount)
	case "premiumLimitTypeSavedAnimationCount":
		t.PremiumLimitTypeSavedAnimationCount = new(PremiumLimitTypeSavedAnimationCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSavedAnimationCount)
	case "premiumLimitTypeShareableChatFolderCount":
		t.PremiumLimitTypeShareableChatFolderCount = new(PremiumLimitTypeShareableChatFolderCount)
		return json.Unmarshal(b, t.PremiumLimitTypeShareableChatFolderCount)
	case "premiumLimitTypeSimilarChatCount":
		t.PremiumLimitTypeSimilarChatCount = new(PremiumLimitTypeSimilarChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSimilarChatCount)
	case "premiumLimitTypeStoryCaptionLength":
		t.PremiumLimitTypeStoryCaptionLength = new(PremiumLimitTypeStoryCaptionLength)
		return json.Unmarshal(b, t.PremiumLimitTypeStoryCaptionLength)
	case "premiumLimitTypeStorySuggestedReactionAreaCount":
		t.PremiumLimitTypeStorySuggestedReactionAreaCount = new(PremiumLimitTypeStorySuggestedReactionAreaCount)
		return json.Unmarshal(b, t.PremiumLimitTypeStorySuggestedReactionAreaCount)
	case "premiumLimitTypeSupergroupCount":
		t.PremiumLimitTypeSupergroupCount = new(PremiumLimitTypeSupergroupCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSupergroupCount)
	case "premiumLimitTypeWeeklyPostedStoryCount":
		t.PremiumLimitTypeWeeklyPostedStoryCount = new(PremiumLimitTypeWeeklyPostedStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeWeeklyPostedStoryCount)
	}
	return nil
}

func (t *PremiumLimitType) MarshalJSON() ([]byte, error) {
	if t.PremiumLimitTypeActiveStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeActiveStoryCount)
	}
	if t.PremiumLimitTypeBioLength != nil {
		return json.Marshal(t.PremiumLimitTypeBioLength)
	}
	if t.PremiumLimitTypeCaptionLength != nil {
		return json.Marshal(t.PremiumLimitTypeCaptionLength)
	}
	if t.PremiumLimitTypeChatFolderChosenChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderChosenChatCount)
	}
	if t.PremiumLimitTypeChatFolderCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderCount)
	}
	if t.PremiumLimitTypeChatFolderInviteLinkCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderInviteLinkCount)
	}
	if t.PremiumLimitTypeCreatedPublicChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeCreatedPublicChatCount)
	}
	if t.PremiumLimitTypeFavoriteStickerCount != nil {
		return json.Marshal(t.PremiumLimitTypeFavoriteStickerCount)
	}
	if t.PremiumLimitTypeMonthlyPostedStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeMonthlyPostedStoryCount)
	}
	if t.PremiumLimitTypePinnedArchivedChatCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedArchivedChatCount)
	}
	if t.PremiumLimitTypePinnedChatCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedChatCount)
	}
	if t.PremiumLimitTypePinnedSavedMessagesTopicCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedSavedMessagesTopicCount)
	}
	if t.PremiumLimitTypeSavedAnimationCount != nil {
		return json.Marshal(t.PremiumLimitTypeSavedAnimationCount)
	}
	if t.PremiumLimitTypeShareableChatFolderCount != nil {
		return json.Marshal(t.PremiumLimitTypeShareableChatFolderCount)
	}
	if t.PremiumLimitTypeSimilarChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeSimilarChatCount)
	}
	if t.PremiumLimitTypeStoryCaptionLength != nil {
		return json.Marshal(t.PremiumLimitTypeStoryCaptionLength)
	}
	if t.PremiumLimitTypeStorySuggestedReactionAreaCount != nil {
		return json.Marshal(t.PremiumLimitTypeStorySuggestedReactionAreaCount)
	}
	if t.PremiumLimitTypeSupergroupCount != nil {
		return json.Marshal(t.PremiumLimitTypeSupergroupCount)
	}
	if t.PremiumLimitTypeWeeklyPostedStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeWeeklyPostedStoryCount)
	}
	return nil, fmt.Errorf("no value set for PremiumLimitType")
}

// PremiumSource Describes a source from which the Premium features screen is opened
type PremiumSource struct {
	TypeStr                      string                        `json:"@type"`
	PremiumSourceBusinessFeature *PremiumSourceBusinessFeature `json:"premiumSourceBusinessFeature,omitempty"`
	PremiumSourceFeature         *PremiumSourceFeature         `json:"premiumSourceFeature,omitempty"`
	PremiumSourceLimitExceeded   *PremiumSourceLimitExceeded   `json:"premiumSourceLimitExceeded,omitempty"`
	PremiumSourceLink            *PremiumSourceLink            `json:"premiumSourceLink,omitempty"`
	PremiumSourceSettings        *PremiumSourceSettings        `json:"premiumSourceSettings,omitempty"`
	PremiumSourceStoryFeature    *PremiumSourceStoryFeature    `json:"premiumSourceStoryFeature,omitempty"`
}

func (t *PremiumSource) Type() string {
	return t.TypeStr
}

func (t *PremiumSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PremiumSource) GetExtra() string {
	return ""
}

func (t *PremiumSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "premiumSourceBusinessFeature":
		t.PremiumSourceBusinessFeature = new(PremiumSourceBusinessFeature)
		return json.Unmarshal(b, t.PremiumSourceBusinessFeature)
	case "premiumSourceFeature":
		t.PremiumSourceFeature = new(PremiumSourceFeature)
		return json.Unmarshal(b, t.PremiumSourceFeature)
	case "premiumSourceLimitExceeded":
		t.PremiumSourceLimitExceeded = new(PremiumSourceLimitExceeded)
		return json.Unmarshal(b, t.PremiumSourceLimitExceeded)
	case "premiumSourceLink":
		t.PremiumSourceLink = new(PremiumSourceLink)
		return json.Unmarshal(b, t.PremiumSourceLink)
	case "premiumSourceSettings":
		t.PremiumSourceSettings = new(PremiumSourceSettings)
		return json.Unmarshal(b, t.PremiumSourceSettings)
	case "premiumSourceStoryFeature":
		t.PremiumSourceStoryFeature = new(PremiumSourceStoryFeature)
		return json.Unmarshal(b, t.PremiumSourceStoryFeature)
	}
	return nil
}

func (t *PremiumSource) MarshalJSON() ([]byte, error) {
	if t.PremiumSourceBusinessFeature != nil {
		return json.Marshal(t.PremiumSourceBusinessFeature)
	}
	if t.PremiumSourceFeature != nil {
		return json.Marshal(t.PremiumSourceFeature)
	}
	if t.PremiumSourceLimitExceeded != nil {
		return json.Marshal(t.PremiumSourceLimitExceeded)
	}
	if t.PremiumSourceLink != nil {
		return json.Marshal(t.PremiumSourceLink)
	}
	if t.PremiumSourceSettings != nil {
		return json.Marshal(t.PremiumSourceSettings)
	}
	if t.PremiumSourceStoryFeature != nil {
		return json.Marshal(t.PremiumSourceStoryFeature)
	}
	return nil, fmt.Errorf("no value set for PremiumSource")
}

// PremiumStoryFeature Describes a story feature available to Premium users
type PremiumStoryFeature struct {
	TypeStr                                     string                                       `json:"@type"`
	PremiumStoryFeatureCustomExpirationDuration *PremiumStoryFeatureCustomExpirationDuration `json:"premiumStoryFeatureCustomExpirationDuration,omitempty"`
	PremiumStoryFeatureLinksAndFormatting       *PremiumStoryFeatureLinksAndFormatting       `json:"premiumStoryFeatureLinksAndFormatting,omitempty"`
	PremiumStoryFeaturePermanentViewsHistory    *PremiumStoryFeaturePermanentViewsHistory    `json:"premiumStoryFeaturePermanentViewsHistory,omitempty"`
	PremiumStoryFeaturePriorityOrder            *PremiumStoryFeaturePriorityOrder            `json:"premiumStoryFeaturePriorityOrder,omitempty"`
	PremiumStoryFeatureSaveStories              *PremiumStoryFeatureSaveStories              `json:"premiumStoryFeatureSaveStories,omitempty"`
	PremiumStoryFeatureStealthMode              *PremiumStoryFeatureStealthMode              `json:"premiumStoryFeatureStealthMode,omitempty"`
	PremiumStoryFeatureVideoQuality             *PremiumStoryFeatureVideoQuality             `json:"premiumStoryFeatureVideoQuality,omitempty"`
}

func (t *PremiumStoryFeature) Type() string {
	return t.TypeStr
}

func (t *PremiumStoryFeature) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PremiumStoryFeature) GetExtra() string {
	return ""
}

func (t *PremiumStoryFeature) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "premiumStoryFeatureCustomExpirationDuration":
		t.PremiumStoryFeatureCustomExpirationDuration = new(PremiumStoryFeatureCustomExpirationDuration)
		return json.Unmarshal(b, t.PremiumStoryFeatureCustomExpirationDuration)
	case "premiumStoryFeatureLinksAndFormatting":
		t.PremiumStoryFeatureLinksAndFormatting = new(PremiumStoryFeatureLinksAndFormatting)
		return json.Unmarshal(b, t.PremiumStoryFeatureLinksAndFormatting)
	case "premiumStoryFeaturePermanentViewsHistory":
		t.PremiumStoryFeaturePermanentViewsHistory = new(PremiumStoryFeaturePermanentViewsHistory)
		return json.Unmarshal(b, t.PremiumStoryFeaturePermanentViewsHistory)
	case "premiumStoryFeaturePriorityOrder":
		t.PremiumStoryFeaturePriorityOrder = new(PremiumStoryFeaturePriorityOrder)
		return json.Unmarshal(b, t.PremiumStoryFeaturePriorityOrder)
	case "premiumStoryFeatureSaveStories":
		t.PremiumStoryFeatureSaveStories = new(PremiumStoryFeatureSaveStories)
		return json.Unmarshal(b, t.PremiumStoryFeatureSaveStories)
	case "premiumStoryFeatureStealthMode":
		t.PremiumStoryFeatureStealthMode = new(PremiumStoryFeatureStealthMode)
		return json.Unmarshal(b, t.PremiumStoryFeatureStealthMode)
	case "premiumStoryFeatureVideoQuality":
		t.PremiumStoryFeatureVideoQuality = new(PremiumStoryFeatureVideoQuality)
		return json.Unmarshal(b, t.PremiumStoryFeatureVideoQuality)
	}
	return nil
}

func (t *PremiumStoryFeature) MarshalJSON() ([]byte, error) {
	if t.PremiumStoryFeatureCustomExpirationDuration != nil {
		return json.Marshal(t.PremiumStoryFeatureCustomExpirationDuration)
	}
	if t.PremiumStoryFeatureLinksAndFormatting != nil {
		return json.Marshal(t.PremiumStoryFeatureLinksAndFormatting)
	}
	if t.PremiumStoryFeaturePermanentViewsHistory != nil {
		return json.Marshal(t.PremiumStoryFeaturePermanentViewsHistory)
	}
	if t.PremiumStoryFeaturePriorityOrder != nil {
		return json.Marshal(t.PremiumStoryFeaturePriorityOrder)
	}
	if t.PremiumStoryFeatureSaveStories != nil {
		return json.Marshal(t.PremiumStoryFeatureSaveStories)
	}
	if t.PremiumStoryFeatureStealthMode != nil {
		return json.Marshal(t.PremiumStoryFeatureStealthMode)
	}
	if t.PremiumStoryFeatureVideoQuality != nil {
		return json.Marshal(t.PremiumStoryFeatureVideoQuality)
	}
	return nil, fmt.Errorf("no value set for PremiumStoryFeature")
}

// ProfileTab Describes a tab shown in a user or a chat profile
type ProfileTab struct {
	TypeStr         string           `json:"@type"`
	ProfileTabFiles *ProfileTabFiles `json:"profileTabFiles,omitempty"`
	ProfileTabGifs  *ProfileTabGifs  `json:"profileTabGifs,omitempty"`
	ProfileTabGifts *ProfileTabGifts `json:"profileTabGifts,omitempty"`
	ProfileTabLinks *ProfileTabLinks `json:"profileTabLinks,omitempty"`
	ProfileTabMedia *ProfileTabMedia `json:"profileTabMedia,omitempty"`
	ProfileTabMusic *ProfileTabMusic `json:"profileTabMusic,omitempty"`
	ProfileTabPosts *ProfileTabPosts `json:"profileTabPosts,omitempty"`
	ProfileTabVoice *ProfileTabVoice `json:"profileTabVoice,omitempty"`
}

func (t *ProfileTab) Type() string {
	return t.TypeStr
}

func (t *ProfileTab) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ProfileTab) GetExtra() string {
	return ""
}

func (t *ProfileTab) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "profileTabFiles":
		t.ProfileTabFiles = new(ProfileTabFiles)
		return json.Unmarshal(b, t.ProfileTabFiles)
	case "profileTabGifs":
		t.ProfileTabGifs = new(ProfileTabGifs)
		return json.Unmarshal(b, t.ProfileTabGifs)
	case "profileTabGifts":
		t.ProfileTabGifts = new(ProfileTabGifts)
		return json.Unmarshal(b, t.ProfileTabGifts)
	case "profileTabLinks":
		t.ProfileTabLinks = new(ProfileTabLinks)
		return json.Unmarshal(b, t.ProfileTabLinks)
	case "profileTabMedia":
		t.ProfileTabMedia = new(ProfileTabMedia)
		return json.Unmarshal(b, t.ProfileTabMedia)
	case "profileTabMusic":
		t.ProfileTabMusic = new(ProfileTabMusic)
		return json.Unmarshal(b, t.ProfileTabMusic)
	case "profileTabPosts":
		t.ProfileTabPosts = new(ProfileTabPosts)
		return json.Unmarshal(b, t.ProfileTabPosts)
	case "profileTabVoice":
		t.ProfileTabVoice = new(ProfileTabVoice)
		return json.Unmarshal(b, t.ProfileTabVoice)
	}
	return nil
}

func (t *ProfileTab) MarshalJSON() ([]byte, error) {
	if t.ProfileTabFiles != nil {
		return json.Marshal(t.ProfileTabFiles)
	}
	if t.ProfileTabGifs != nil {
		return json.Marshal(t.ProfileTabGifs)
	}
	if t.ProfileTabGifts != nil {
		return json.Marshal(t.ProfileTabGifts)
	}
	if t.ProfileTabLinks != nil {
		return json.Marshal(t.ProfileTabLinks)
	}
	if t.ProfileTabMedia != nil {
		return json.Marshal(t.ProfileTabMedia)
	}
	if t.ProfileTabMusic != nil {
		return json.Marshal(t.ProfileTabMusic)
	}
	if t.ProfileTabPosts != nil {
		return json.Marshal(t.ProfileTabPosts)
	}
	if t.ProfileTabVoice != nil {
		return json.Marshal(t.ProfileTabVoice)
	}
	return nil, fmt.Errorf("no value set for ProfileTab")
}

// ProxyType Describes the type of proxy server
type ProxyType struct {
	TypeStr          string            `json:"@type"`
	ProxyTypeHttp    *ProxyTypeHttp    `json:"proxyTypeHttp,omitempty"`
	ProxyTypeMtproto *ProxyTypeMtproto `json:"proxyTypeMtproto,omitempty"`
	ProxyTypeSocks5  *ProxyTypeSocks5  `json:"proxyTypeSocks5,omitempty"`
}

func (t *ProxyType) Type() string {
	return t.TypeStr
}

func (t *ProxyType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ProxyType) GetExtra() string {
	return ""
}

func (t *ProxyType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "proxyTypeHttp":
		t.ProxyTypeHttp = new(ProxyTypeHttp)
		return json.Unmarshal(b, t.ProxyTypeHttp)
	case "proxyTypeMtproto":
		t.ProxyTypeMtproto = new(ProxyTypeMtproto)
		return json.Unmarshal(b, t.ProxyTypeMtproto)
	case "proxyTypeSocks5":
		t.ProxyTypeSocks5 = new(ProxyTypeSocks5)
		return json.Unmarshal(b, t.ProxyTypeSocks5)
	}
	return nil
}

func (t *ProxyType) MarshalJSON() ([]byte, error) {
	if t.ProxyTypeHttp != nil {
		return json.Marshal(t.ProxyTypeHttp)
	}
	if t.ProxyTypeMtproto != nil {
		return json.Marshal(t.ProxyTypeMtproto)
	}
	if t.ProxyTypeSocks5 != nil {
		return json.Marshal(t.ProxyTypeSocks5)
	}
	return nil, fmt.Errorf("no value set for ProxyType")
}

// PublicChatType Describes type of public chat
type PublicChatType struct {
	TypeStr                       string                         `json:"@type"`
	PublicChatTypeHasUsername     *PublicChatTypeHasUsername     `json:"publicChatTypeHasUsername,omitempty"`
	PublicChatTypeIsLocationBased *PublicChatTypeIsLocationBased `json:"publicChatTypeIsLocationBased,omitempty"`
}

func (t *PublicChatType) Type() string {
	return t.TypeStr
}

func (t *PublicChatType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PublicChatType) GetExtra() string {
	return ""
}

func (t *PublicChatType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "publicChatTypeHasUsername":
		t.PublicChatTypeHasUsername = new(PublicChatTypeHasUsername)
		return json.Unmarshal(b, t.PublicChatTypeHasUsername)
	case "publicChatTypeIsLocationBased":
		t.PublicChatTypeIsLocationBased = new(PublicChatTypeIsLocationBased)
		return json.Unmarshal(b, t.PublicChatTypeIsLocationBased)
	}
	return nil
}

func (t *PublicChatType) MarshalJSON() ([]byte, error) {
	if t.PublicChatTypeHasUsername != nil {
		return json.Marshal(t.PublicChatTypeHasUsername)
	}
	if t.PublicChatTypeIsLocationBased != nil {
		return json.Marshal(t.PublicChatTypeIsLocationBased)
	}
	return nil, fmt.Errorf("no value set for PublicChatType")
}

// PublicForward Describes a public forward or repost of a story
type PublicForward struct {
	TypeStr              string                `json:"@type"`
	PublicForwardMessage *PublicForwardMessage `json:"publicForwardMessage,omitempty"`
	PublicForwardStory   *PublicForwardStory   `json:"publicForwardStory,omitempty"`
}

func (t *PublicForward) Type() string {
	return t.TypeStr
}

func (t *PublicForward) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PublicForward) GetExtra() string {
	return ""
}

func (t *PublicForward) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "publicForwardMessage":
		t.PublicForwardMessage = new(PublicForwardMessage)
		return json.Unmarshal(b, t.PublicForwardMessage)
	case "publicForwardStory":
		t.PublicForwardStory = new(PublicForwardStory)
		return json.Unmarshal(b, t.PublicForwardStory)
	}
	return nil
}

func (t *PublicForward) MarshalJSON() ([]byte, error) {
	if t.PublicForwardMessage != nil {
		return json.Marshal(t.PublicForwardMessage)
	}
	if t.PublicForwardStory != nil {
		return json.Marshal(t.PublicForwardStory)
	}
	return nil, fmt.Errorf("no value set for PublicForward")
}

// PushMessageContent Contains content of a push message notification
type PushMessageContent struct {
	TypeStr                                       string                                         `json:"@type"`
	PushMessageContentAnimation                   *PushMessageContentAnimation                   `json:"pushMessageContentAnimation,omitempty"`
	PushMessageContentAudio                       *PushMessageContentAudio                       `json:"pushMessageContentAudio,omitempty"`
	PushMessageContentBasicGroupChatCreate        *PushMessageContentBasicGroupChatCreate        `json:"pushMessageContentBasicGroupChatCreate,omitempty"`
	PushMessageContentChatAddMembers              *PushMessageContentChatAddMembers              `json:"pushMessageContentChatAddMembers,omitempty"`
	PushMessageContentChatChangePhoto             *PushMessageContentChatChangePhoto             `json:"pushMessageContentChatChangePhoto,omitempty"`
	PushMessageContentChatChangeTitle             *PushMessageContentChatChangeTitle             `json:"pushMessageContentChatChangeTitle,omitempty"`
	PushMessageContentChatDeleteMember            *PushMessageContentChatDeleteMember            `json:"pushMessageContentChatDeleteMember,omitempty"`
	PushMessageContentChatJoinByLink              *PushMessageContentChatJoinByLink              `json:"pushMessageContentChatJoinByLink,omitempty"`
	PushMessageContentChatJoinByRequest           *PushMessageContentChatJoinByRequest           `json:"pushMessageContentChatJoinByRequest,omitempty"`
	PushMessageContentChatSetBackground           *PushMessageContentChatSetBackground           `json:"pushMessageContentChatSetBackground,omitempty"`
	PushMessageContentChatSetTheme                *PushMessageContentChatSetTheme                `json:"pushMessageContentChatSetTheme,omitempty"`
	PushMessageContentChecklist                   *PushMessageContentChecklist                   `json:"pushMessageContentChecklist,omitempty"`
	PushMessageContentChecklistTasksAdded         *PushMessageContentChecklistTasksAdded         `json:"pushMessageContentChecklistTasksAdded,omitempty"`
	PushMessageContentChecklistTasksDone          *PushMessageContentChecklistTasksDone          `json:"pushMessageContentChecklistTasksDone,omitempty"`
	PushMessageContentContact                     *PushMessageContentContact                     `json:"pushMessageContentContact,omitempty"`
	PushMessageContentContactRegistered           *PushMessageContentContactRegistered           `json:"pushMessageContentContactRegistered,omitempty"`
	PushMessageContentDocument                    *PushMessageContentDocument                    `json:"pushMessageContentDocument,omitempty"`
	PushMessageContentGame                        *PushMessageContentGame                        `json:"pushMessageContentGame,omitempty"`
	PushMessageContentGameScore                   *PushMessageContentGameScore                   `json:"pushMessageContentGameScore,omitempty"`
	PushMessageContentGift                        *PushMessageContentGift                        `json:"pushMessageContentGift,omitempty"`
	PushMessageContentGiveaway                    *PushMessageContentGiveaway                    `json:"pushMessageContentGiveaway,omitempty"`
	PushMessageContentHidden                      *PushMessageContentHidden                      `json:"pushMessageContentHidden,omitempty"`
	PushMessageContentInviteVideoChatParticipants *PushMessageContentInviteVideoChatParticipants `json:"pushMessageContentInviteVideoChatParticipants,omitempty"`
	PushMessageContentInvoice                     *PushMessageContentInvoice                     `json:"pushMessageContentInvoice,omitempty"`
	PushMessageContentLocation                    *PushMessageContentLocation                    `json:"pushMessageContentLocation,omitempty"`
	PushMessageContentMediaAlbum                  *PushMessageContentMediaAlbum                  `json:"pushMessageContentMediaAlbum,omitempty"`
	PushMessageContentMessageForwards             *PushMessageContentMessageForwards             `json:"pushMessageContentMessageForwards,omitempty"`
	PushMessageContentPaidMedia                   *PushMessageContentPaidMedia                   `json:"pushMessageContentPaidMedia,omitempty"`
	PushMessageContentPhoto                       *PushMessageContentPhoto                       `json:"pushMessageContentPhoto,omitempty"`
	PushMessageContentPoll                        *PushMessageContentPoll                        `json:"pushMessageContentPoll,omitempty"`
	PushMessageContentPremiumGiftCode             *PushMessageContentPremiumGiftCode             `json:"pushMessageContentPremiumGiftCode,omitempty"`
	PushMessageContentProximityAlertTriggered     *PushMessageContentProximityAlertTriggered     `json:"pushMessageContentProximityAlertTriggered,omitempty"`
	PushMessageContentRecurringPayment            *PushMessageContentRecurringPayment            `json:"pushMessageContentRecurringPayment,omitempty"`
	PushMessageContentScreenshotTaken             *PushMessageContentScreenshotTaken             `json:"pushMessageContentScreenshotTaken,omitempty"`
	PushMessageContentSticker                     *PushMessageContentSticker                     `json:"pushMessageContentSticker,omitempty"`
	PushMessageContentStory                       *PushMessageContentStory                       `json:"pushMessageContentStory,omitempty"`
	PushMessageContentSuggestBirthdate            *PushMessageContentSuggestBirthdate            `json:"pushMessageContentSuggestBirthdate,omitempty"`
	PushMessageContentSuggestProfilePhoto         *PushMessageContentSuggestProfilePhoto         `json:"pushMessageContentSuggestProfilePhoto,omitempty"`
	PushMessageContentText                        *PushMessageContentText                        `json:"pushMessageContentText,omitempty"`
	PushMessageContentUpgradedGift                *PushMessageContentUpgradedGift                `json:"pushMessageContentUpgradedGift,omitempty"`
	PushMessageContentVideo                       *PushMessageContentVideo                       `json:"pushMessageContentVideo,omitempty"`
	PushMessageContentVideoChatEnded              *PushMessageContentVideoChatEnded              `json:"pushMessageContentVideoChatEnded,omitempty"`
	PushMessageContentVideoChatStarted            *PushMessageContentVideoChatStarted            `json:"pushMessageContentVideoChatStarted,omitempty"`
	PushMessageContentVideoNote                   *PushMessageContentVideoNote                   `json:"pushMessageContentVideoNote,omitempty"`
	PushMessageContentVoiceNote                   *PushMessageContentVoiceNote                   `json:"pushMessageContentVoiceNote,omitempty"`
}

func (t *PushMessageContent) Type() string {
	return t.TypeStr
}

func (t *PushMessageContent) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *PushMessageContent) GetExtra() string {
	return ""
}

func (t *PushMessageContent) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "pushMessageContentAnimation":
		t.PushMessageContentAnimation = new(PushMessageContentAnimation)
		return json.Unmarshal(b, t.PushMessageContentAnimation)
	case "pushMessageContentAudio":
		t.PushMessageContentAudio = new(PushMessageContentAudio)
		return json.Unmarshal(b, t.PushMessageContentAudio)
	case "pushMessageContentBasicGroupChatCreate":
		t.PushMessageContentBasicGroupChatCreate = new(PushMessageContentBasicGroupChatCreate)
		return json.Unmarshal(b, t.PushMessageContentBasicGroupChatCreate)
	case "pushMessageContentChatAddMembers":
		t.PushMessageContentChatAddMembers = new(PushMessageContentChatAddMembers)
		return json.Unmarshal(b, t.PushMessageContentChatAddMembers)
	case "pushMessageContentChatChangePhoto":
		t.PushMessageContentChatChangePhoto = new(PushMessageContentChatChangePhoto)
		return json.Unmarshal(b, t.PushMessageContentChatChangePhoto)
	case "pushMessageContentChatChangeTitle":
		t.PushMessageContentChatChangeTitle = new(PushMessageContentChatChangeTitle)
		return json.Unmarshal(b, t.PushMessageContentChatChangeTitle)
	case "pushMessageContentChatDeleteMember":
		t.PushMessageContentChatDeleteMember = new(PushMessageContentChatDeleteMember)
		return json.Unmarshal(b, t.PushMessageContentChatDeleteMember)
	case "pushMessageContentChatJoinByLink":
		t.PushMessageContentChatJoinByLink = new(PushMessageContentChatJoinByLink)
		return json.Unmarshal(b, t.PushMessageContentChatJoinByLink)
	case "pushMessageContentChatJoinByRequest":
		t.PushMessageContentChatJoinByRequest = new(PushMessageContentChatJoinByRequest)
		return json.Unmarshal(b, t.PushMessageContentChatJoinByRequest)
	case "pushMessageContentChatSetBackground":
		t.PushMessageContentChatSetBackground = new(PushMessageContentChatSetBackground)
		return json.Unmarshal(b, t.PushMessageContentChatSetBackground)
	case "pushMessageContentChatSetTheme":
		t.PushMessageContentChatSetTheme = new(PushMessageContentChatSetTheme)
		return json.Unmarshal(b, t.PushMessageContentChatSetTheme)
	case "pushMessageContentChecklist":
		t.PushMessageContentChecklist = new(PushMessageContentChecklist)
		return json.Unmarshal(b, t.PushMessageContentChecklist)
	case "pushMessageContentChecklistTasksAdded":
		t.PushMessageContentChecklistTasksAdded = new(PushMessageContentChecklistTasksAdded)
		return json.Unmarshal(b, t.PushMessageContentChecklistTasksAdded)
	case "pushMessageContentChecklistTasksDone":
		t.PushMessageContentChecklistTasksDone = new(PushMessageContentChecklistTasksDone)
		return json.Unmarshal(b, t.PushMessageContentChecklistTasksDone)
	case "pushMessageContentContact":
		t.PushMessageContentContact = new(PushMessageContentContact)
		return json.Unmarshal(b, t.PushMessageContentContact)
	case "pushMessageContentContactRegistered":
		t.PushMessageContentContactRegistered = new(PushMessageContentContactRegistered)
		return json.Unmarshal(b, t.PushMessageContentContactRegistered)
	case "pushMessageContentDocument":
		t.PushMessageContentDocument = new(PushMessageContentDocument)
		return json.Unmarshal(b, t.PushMessageContentDocument)
	case "pushMessageContentGame":
		t.PushMessageContentGame = new(PushMessageContentGame)
		return json.Unmarshal(b, t.PushMessageContentGame)
	case "pushMessageContentGameScore":
		t.PushMessageContentGameScore = new(PushMessageContentGameScore)
		return json.Unmarshal(b, t.PushMessageContentGameScore)
	case "pushMessageContentGift":
		t.PushMessageContentGift = new(PushMessageContentGift)
		return json.Unmarshal(b, t.PushMessageContentGift)
	case "pushMessageContentGiveaway":
		t.PushMessageContentGiveaway = new(PushMessageContentGiveaway)
		return json.Unmarshal(b, t.PushMessageContentGiveaway)
	case "pushMessageContentHidden":
		t.PushMessageContentHidden = new(PushMessageContentHidden)
		return json.Unmarshal(b, t.PushMessageContentHidden)
	case "pushMessageContentInviteVideoChatParticipants":
		t.PushMessageContentInviteVideoChatParticipants = new(PushMessageContentInviteVideoChatParticipants)
		return json.Unmarshal(b, t.PushMessageContentInviteVideoChatParticipants)
	case "pushMessageContentInvoice":
		t.PushMessageContentInvoice = new(PushMessageContentInvoice)
		return json.Unmarshal(b, t.PushMessageContentInvoice)
	case "pushMessageContentLocation":
		t.PushMessageContentLocation = new(PushMessageContentLocation)
		return json.Unmarshal(b, t.PushMessageContentLocation)
	case "pushMessageContentMediaAlbum":
		t.PushMessageContentMediaAlbum = new(PushMessageContentMediaAlbum)
		return json.Unmarshal(b, t.PushMessageContentMediaAlbum)
	case "pushMessageContentMessageForwards":
		t.PushMessageContentMessageForwards = new(PushMessageContentMessageForwards)
		return json.Unmarshal(b, t.PushMessageContentMessageForwards)
	case "pushMessageContentPaidMedia":
		t.PushMessageContentPaidMedia = new(PushMessageContentPaidMedia)
		return json.Unmarshal(b, t.PushMessageContentPaidMedia)
	case "pushMessageContentPhoto":
		t.PushMessageContentPhoto = new(PushMessageContentPhoto)
		return json.Unmarshal(b, t.PushMessageContentPhoto)
	case "pushMessageContentPoll":
		t.PushMessageContentPoll = new(PushMessageContentPoll)
		return json.Unmarshal(b, t.PushMessageContentPoll)
	case "pushMessageContentPremiumGiftCode":
		t.PushMessageContentPremiumGiftCode = new(PushMessageContentPremiumGiftCode)
		return json.Unmarshal(b, t.PushMessageContentPremiumGiftCode)
	case "pushMessageContentProximityAlertTriggered":
		t.PushMessageContentProximityAlertTriggered = new(PushMessageContentProximityAlertTriggered)
		return json.Unmarshal(b, t.PushMessageContentProximityAlertTriggered)
	case "pushMessageContentRecurringPayment":
		t.PushMessageContentRecurringPayment = new(PushMessageContentRecurringPayment)
		return json.Unmarshal(b, t.PushMessageContentRecurringPayment)
	case "pushMessageContentScreenshotTaken":
		t.PushMessageContentScreenshotTaken = new(PushMessageContentScreenshotTaken)
		return json.Unmarshal(b, t.PushMessageContentScreenshotTaken)
	case "pushMessageContentSticker":
		t.PushMessageContentSticker = new(PushMessageContentSticker)
		return json.Unmarshal(b, t.PushMessageContentSticker)
	case "pushMessageContentStory":
		t.PushMessageContentStory = new(PushMessageContentStory)
		return json.Unmarshal(b, t.PushMessageContentStory)
	case "pushMessageContentSuggestBirthdate":
		t.PushMessageContentSuggestBirthdate = new(PushMessageContentSuggestBirthdate)
		return json.Unmarshal(b, t.PushMessageContentSuggestBirthdate)
	case "pushMessageContentSuggestProfilePhoto":
		t.PushMessageContentSuggestProfilePhoto = new(PushMessageContentSuggestProfilePhoto)
		return json.Unmarshal(b, t.PushMessageContentSuggestProfilePhoto)
	case "pushMessageContentText":
		t.PushMessageContentText = new(PushMessageContentText)
		return json.Unmarshal(b, t.PushMessageContentText)
	case "pushMessageContentUpgradedGift":
		t.PushMessageContentUpgradedGift = new(PushMessageContentUpgradedGift)
		return json.Unmarshal(b, t.PushMessageContentUpgradedGift)
	case "pushMessageContentVideo":
		t.PushMessageContentVideo = new(PushMessageContentVideo)
		return json.Unmarshal(b, t.PushMessageContentVideo)
	case "pushMessageContentVideoChatEnded":
		t.PushMessageContentVideoChatEnded = new(PushMessageContentVideoChatEnded)
		return json.Unmarshal(b, t.PushMessageContentVideoChatEnded)
	case "pushMessageContentVideoChatStarted":
		t.PushMessageContentVideoChatStarted = new(PushMessageContentVideoChatStarted)
		return json.Unmarshal(b, t.PushMessageContentVideoChatStarted)
	case "pushMessageContentVideoNote":
		t.PushMessageContentVideoNote = new(PushMessageContentVideoNote)
		return json.Unmarshal(b, t.PushMessageContentVideoNote)
	case "pushMessageContentVoiceNote":
		t.PushMessageContentVoiceNote = new(PushMessageContentVoiceNote)
		return json.Unmarshal(b, t.PushMessageContentVoiceNote)
	}
	return nil
}

func (t *PushMessageContent) MarshalJSON() ([]byte, error) {
	if t.PushMessageContentAnimation != nil {
		return json.Marshal(t.PushMessageContentAnimation)
	}
	if t.PushMessageContentAudio != nil {
		return json.Marshal(t.PushMessageContentAudio)
	}
	if t.PushMessageContentBasicGroupChatCreate != nil {
		return json.Marshal(t.PushMessageContentBasicGroupChatCreate)
	}
	if t.PushMessageContentChatAddMembers != nil {
		return json.Marshal(t.PushMessageContentChatAddMembers)
	}
	if t.PushMessageContentChatChangePhoto != nil {
		return json.Marshal(t.PushMessageContentChatChangePhoto)
	}
	if t.PushMessageContentChatChangeTitle != nil {
		return json.Marshal(t.PushMessageContentChatChangeTitle)
	}
	if t.PushMessageContentChatDeleteMember != nil {
		return json.Marshal(t.PushMessageContentChatDeleteMember)
	}
	if t.PushMessageContentChatJoinByLink != nil {
		return json.Marshal(t.PushMessageContentChatJoinByLink)
	}
	if t.PushMessageContentChatJoinByRequest != nil {
		return json.Marshal(t.PushMessageContentChatJoinByRequest)
	}
	if t.PushMessageContentChatSetBackground != nil {
		return json.Marshal(t.PushMessageContentChatSetBackground)
	}
	if t.PushMessageContentChatSetTheme != nil {
		return json.Marshal(t.PushMessageContentChatSetTheme)
	}
	if t.PushMessageContentChecklist != nil {
		return json.Marshal(t.PushMessageContentChecklist)
	}
	if t.PushMessageContentChecklistTasksAdded != nil {
		return json.Marshal(t.PushMessageContentChecklistTasksAdded)
	}
	if t.PushMessageContentChecklistTasksDone != nil {
		return json.Marshal(t.PushMessageContentChecklistTasksDone)
	}
	if t.PushMessageContentContact != nil {
		return json.Marshal(t.PushMessageContentContact)
	}
	if t.PushMessageContentContactRegistered != nil {
		return json.Marshal(t.PushMessageContentContactRegistered)
	}
	if t.PushMessageContentDocument != nil {
		return json.Marshal(t.PushMessageContentDocument)
	}
	if t.PushMessageContentGame != nil {
		return json.Marshal(t.PushMessageContentGame)
	}
	if t.PushMessageContentGameScore != nil {
		return json.Marshal(t.PushMessageContentGameScore)
	}
	if t.PushMessageContentGift != nil {
		return json.Marshal(t.PushMessageContentGift)
	}
	if t.PushMessageContentGiveaway != nil {
		return json.Marshal(t.PushMessageContentGiveaway)
	}
	if t.PushMessageContentHidden != nil {
		return json.Marshal(t.PushMessageContentHidden)
	}
	if t.PushMessageContentInviteVideoChatParticipants != nil {
		return json.Marshal(t.PushMessageContentInviteVideoChatParticipants)
	}
	if t.PushMessageContentInvoice != nil {
		return json.Marshal(t.PushMessageContentInvoice)
	}
	if t.PushMessageContentLocation != nil {
		return json.Marshal(t.PushMessageContentLocation)
	}
	if t.PushMessageContentMediaAlbum != nil {
		return json.Marshal(t.PushMessageContentMediaAlbum)
	}
	if t.PushMessageContentMessageForwards != nil {
		return json.Marshal(t.PushMessageContentMessageForwards)
	}
	if t.PushMessageContentPaidMedia != nil {
		return json.Marshal(t.PushMessageContentPaidMedia)
	}
	if t.PushMessageContentPhoto != nil {
		return json.Marshal(t.PushMessageContentPhoto)
	}
	if t.PushMessageContentPoll != nil {
		return json.Marshal(t.PushMessageContentPoll)
	}
	if t.PushMessageContentPremiumGiftCode != nil {
		return json.Marshal(t.PushMessageContentPremiumGiftCode)
	}
	if t.PushMessageContentProximityAlertTriggered != nil {
		return json.Marshal(t.PushMessageContentProximityAlertTriggered)
	}
	if t.PushMessageContentRecurringPayment != nil {
		return json.Marshal(t.PushMessageContentRecurringPayment)
	}
	if t.PushMessageContentScreenshotTaken != nil {
		return json.Marshal(t.PushMessageContentScreenshotTaken)
	}
	if t.PushMessageContentSticker != nil {
		return json.Marshal(t.PushMessageContentSticker)
	}
	if t.PushMessageContentStory != nil {
		return json.Marshal(t.PushMessageContentStory)
	}
	if t.PushMessageContentSuggestBirthdate != nil {
		return json.Marshal(t.PushMessageContentSuggestBirthdate)
	}
	if t.PushMessageContentSuggestProfilePhoto != nil {
		return json.Marshal(t.PushMessageContentSuggestProfilePhoto)
	}
	if t.PushMessageContentText != nil {
		return json.Marshal(t.PushMessageContentText)
	}
	if t.PushMessageContentUpgradedGift != nil {
		return json.Marshal(t.PushMessageContentUpgradedGift)
	}
	if t.PushMessageContentVideo != nil {
		return json.Marshal(t.PushMessageContentVideo)
	}
	if t.PushMessageContentVideoChatEnded != nil {
		return json.Marshal(t.PushMessageContentVideoChatEnded)
	}
	if t.PushMessageContentVideoChatStarted != nil {
		return json.Marshal(t.PushMessageContentVideoChatStarted)
	}
	if t.PushMessageContentVideoNote != nil {
		return json.Marshal(t.PushMessageContentVideoNote)
	}
	if t.PushMessageContentVoiceNote != nil {
		return json.Marshal(t.PushMessageContentVoiceNote)
	}
	return nil, fmt.Errorf("no value set for PushMessageContent")
}

// ReactionNotificationSource Describes sources of reactions for which notifications will be shown
type ReactionNotificationSource struct {
	TypeStr                            string                              `json:"@type"`
	ReactionNotificationSourceAll      *ReactionNotificationSourceAll      `json:"reactionNotificationSourceAll,omitempty"`
	ReactionNotificationSourceContacts *ReactionNotificationSourceContacts `json:"reactionNotificationSourceContacts,omitempty"`
	ReactionNotificationSourceNone     *ReactionNotificationSourceNone     `json:"reactionNotificationSourceNone,omitempty"`
}

func (t *ReactionNotificationSource) Type() string {
	return t.TypeStr
}

func (t *ReactionNotificationSource) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReactionNotificationSource) GetExtra() string {
	return ""
}

func (t *ReactionNotificationSource) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reactionNotificationSourceAll":
		t.ReactionNotificationSourceAll = new(ReactionNotificationSourceAll)
		return json.Unmarshal(b, t.ReactionNotificationSourceAll)
	case "reactionNotificationSourceContacts":
		t.ReactionNotificationSourceContacts = new(ReactionNotificationSourceContacts)
		return json.Unmarshal(b, t.ReactionNotificationSourceContacts)
	case "reactionNotificationSourceNone":
		t.ReactionNotificationSourceNone = new(ReactionNotificationSourceNone)
		return json.Unmarshal(b, t.ReactionNotificationSourceNone)
	}
	return nil
}

func (t *ReactionNotificationSource) MarshalJSON() ([]byte, error) {
	if t.ReactionNotificationSourceAll != nil {
		return json.Marshal(t.ReactionNotificationSourceAll)
	}
	if t.ReactionNotificationSourceContacts != nil {
		return json.Marshal(t.ReactionNotificationSourceContacts)
	}
	if t.ReactionNotificationSourceNone != nil {
		return json.Marshal(t.ReactionNotificationSourceNone)
	}
	return nil, fmt.Errorf("no value set for ReactionNotificationSource")
}

// ReactionType Describes type of message reaction
type ReactionType struct {
	TypeStr                 string                   `json:"@type"`
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji `json:"reactionTypeCustomEmoji,omitempty"`
	ReactionTypeEmoji       *ReactionTypeEmoji       `json:"reactionTypeEmoji,omitempty"`
	ReactionTypePaid        *ReactionTypePaid        `json:"reactionTypePaid,omitempty"`
}

func (t *ReactionType) Type() string {
	return t.TypeStr
}

func (t *ReactionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReactionType) GetExtra() string {
	return ""
}

func (t *ReactionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reactionTypeCustomEmoji":
		t.ReactionTypeCustomEmoji = new(ReactionTypeCustomEmoji)
		return json.Unmarshal(b, t.ReactionTypeCustomEmoji)
	case "reactionTypeEmoji":
		t.ReactionTypeEmoji = new(ReactionTypeEmoji)
		return json.Unmarshal(b, t.ReactionTypeEmoji)
	case "reactionTypePaid":
		t.ReactionTypePaid = new(ReactionTypePaid)
		return json.Unmarshal(b, t.ReactionTypePaid)
	}
	return nil
}

func (t *ReactionType) MarshalJSON() ([]byte, error) {
	if t.ReactionTypeCustomEmoji != nil {
		return json.Marshal(t.ReactionTypeCustomEmoji)
	}
	if t.ReactionTypeEmoji != nil {
		return json.Marshal(t.ReactionTypeEmoji)
	}
	if t.ReactionTypePaid != nil {
		return json.Marshal(t.ReactionTypePaid)
	}
	return nil, fmt.Errorf("no value set for ReactionType")
}

// ReactionUnavailabilityReason Describes why the current user can't add reactions to the message, despite some other users can
type ReactionUnavailabilityReason struct {
	TypeStr                                            string                                              `json:"@type"`
	ReactionUnavailabilityReasonAnonymousAdministrator *ReactionUnavailabilityReasonAnonymousAdministrator `json:"reactionUnavailabilityReasonAnonymousAdministrator,omitempty"`
	ReactionUnavailabilityReasonGuest                  *ReactionUnavailabilityReasonGuest                  `json:"reactionUnavailabilityReasonGuest,omitempty"`
}

func (t *ReactionUnavailabilityReason) Type() string {
	return t.TypeStr
}

func (t *ReactionUnavailabilityReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReactionUnavailabilityReason) GetExtra() string {
	return ""
}

func (t *ReactionUnavailabilityReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reactionUnavailabilityReasonAnonymousAdministrator":
		t.ReactionUnavailabilityReasonAnonymousAdministrator = new(ReactionUnavailabilityReasonAnonymousAdministrator)
		return json.Unmarshal(b, t.ReactionUnavailabilityReasonAnonymousAdministrator)
	case "reactionUnavailabilityReasonGuest":
		t.ReactionUnavailabilityReasonGuest = new(ReactionUnavailabilityReasonGuest)
		return json.Unmarshal(b, t.ReactionUnavailabilityReasonGuest)
	}
	return nil
}

func (t *ReactionUnavailabilityReason) MarshalJSON() ([]byte, error) {
	if t.ReactionUnavailabilityReasonAnonymousAdministrator != nil {
		return json.Marshal(t.ReactionUnavailabilityReasonAnonymousAdministrator)
	}
	if t.ReactionUnavailabilityReasonGuest != nil {
		return json.Marshal(t.ReactionUnavailabilityReasonGuest)
	}
	return nil, fmt.Errorf("no value set for ReactionUnavailabilityReason")
}

// ReplyMarkup Contains a description of a custom keyboard and actions that can be done with it to quickly reply to bots
type ReplyMarkup struct {
	TypeStr                   string                     `json:"@type"`
	ReplyMarkupForceReply     *ReplyMarkupForceReply     `json:"replyMarkupForceReply,omitempty"`
	ReplyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard `json:"replyMarkupInlineKeyboard,omitempty"`
	ReplyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard `json:"replyMarkupRemoveKeyboard,omitempty"`
	ReplyMarkupShowKeyboard   *ReplyMarkupShowKeyboard   `json:"replyMarkupShowKeyboard,omitempty"`
}

func (t *ReplyMarkup) Type() string {
	return t.TypeStr
}

func (t *ReplyMarkup) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReplyMarkup) GetExtra() string {
	return ""
}

func (t *ReplyMarkup) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "replyMarkupForceReply":
		t.ReplyMarkupForceReply = new(ReplyMarkupForceReply)
		return json.Unmarshal(b, t.ReplyMarkupForceReply)
	case "replyMarkupInlineKeyboard":
		t.ReplyMarkupInlineKeyboard = new(ReplyMarkupInlineKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupInlineKeyboard)
	case "replyMarkupRemoveKeyboard":
		t.ReplyMarkupRemoveKeyboard = new(ReplyMarkupRemoveKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupRemoveKeyboard)
	case "replyMarkupShowKeyboard":
		t.ReplyMarkupShowKeyboard = new(ReplyMarkupShowKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupShowKeyboard)
	}
	return nil
}

func (t *ReplyMarkup) MarshalJSON() ([]byte, error) {
	if t.ReplyMarkupForceReply != nil {
		return json.Marshal(t.ReplyMarkupForceReply)
	}
	if t.ReplyMarkupInlineKeyboard != nil {
		return json.Marshal(t.ReplyMarkupInlineKeyboard)
	}
	if t.ReplyMarkupRemoveKeyboard != nil {
		return json.Marshal(t.ReplyMarkupRemoveKeyboard)
	}
	if t.ReplyMarkupShowKeyboard != nil {
		return json.Marshal(t.ReplyMarkupShowKeyboard)
	}
	return nil, fmt.Errorf("no value set for ReplyMarkup")
}

// ReportChatResult Describes result of chat report
type ReportChatResult struct {
	TypeStr                          string                            `json:"@type"`
	ReportChatResultMessagesRequired *ReportChatResultMessagesRequired `json:"reportChatResultMessagesRequired,omitempty"`
	ReportChatResultOk               *ReportChatResultOk               `json:"reportChatResultOk,omitempty"`
	ReportChatResultOptionRequired   *ReportChatResultOptionRequired   `json:"reportChatResultOptionRequired,omitempty"`
	ReportChatResultTextRequired     *ReportChatResultTextRequired     `json:"reportChatResultTextRequired,omitempty"`
}

func (t *ReportChatResult) Type() string {
	return t.TypeStr
}

func (t *ReportChatResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReportChatResult) GetExtra() string {
	return ""
}

func (t *ReportChatResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reportChatResultMessagesRequired":
		t.ReportChatResultMessagesRequired = new(ReportChatResultMessagesRequired)
		return json.Unmarshal(b, t.ReportChatResultMessagesRequired)
	case "reportChatResultOk":
		t.ReportChatResultOk = new(ReportChatResultOk)
		return json.Unmarshal(b, t.ReportChatResultOk)
	case "reportChatResultOptionRequired":
		t.ReportChatResultOptionRequired = new(ReportChatResultOptionRequired)
		return json.Unmarshal(b, t.ReportChatResultOptionRequired)
	case "reportChatResultTextRequired":
		t.ReportChatResultTextRequired = new(ReportChatResultTextRequired)
		return json.Unmarshal(b, t.ReportChatResultTextRequired)
	}
	return nil
}

func (t *ReportChatResult) MarshalJSON() ([]byte, error) {
	if t.ReportChatResultMessagesRequired != nil {
		return json.Marshal(t.ReportChatResultMessagesRequired)
	}
	if t.ReportChatResultOk != nil {
		return json.Marshal(t.ReportChatResultOk)
	}
	if t.ReportChatResultOptionRequired != nil {
		return json.Marshal(t.ReportChatResultOptionRequired)
	}
	if t.ReportChatResultTextRequired != nil {
		return json.Marshal(t.ReportChatResultTextRequired)
	}
	return nil, fmt.Errorf("no value set for ReportChatResult")
}

// ReportReason Describes the reason why a chat is reported
type ReportReason struct {
	TypeStr                       string                         `json:"@type"`
	ReportReasonChildAbuse        *ReportReasonChildAbuse        `json:"reportReasonChildAbuse,omitempty"`
	ReportReasonCopyright         *ReportReasonCopyright         `json:"reportReasonCopyright,omitempty"`
	ReportReasonCustom            *ReportReasonCustom            `json:"reportReasonCustom,omitempty"`
	ReportReasonFake              *ReportReasonFake              `json:"reportReasonFake,omitempty"`
	ReportReasonIllegalDrugs      *ReportReasonIllegalDrugs      `json:"reportReasonIllegalDrugs,omitempty"`
	ReportReasonPersonalDetails   *ReportReasonPersonalDetails   `json:"reportReasonPersonalDetails,omitempty"`
	ReportReasonPornography       *ReportReasonPornography       `json:"reportReasonPornography,omitempty"`
	ReportReasonSpam              *ReportReasonSpam              `json:"reportReasonSpam,omitempty"`
	ReportReasonUnrelatedLocation *ReportReasonUnrelatedLocation `json:"reportReasonUnrelatedLocation,omitempty"`
	ReportReasonViolence          *ReportReasonViolence          `json:"reportReasonViolence,omitempty"`
}

func (t *ReportReason) Type() string {
	return t.TypeStr
}

func (t *ReportReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReportReason) GetExtra() string {
	return ""
}

func (t *ReportReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reportReasonChildAbuse":
		t.ReportReasonChildAbuse = new(ReportReasonChildAbuse)
		return json.Unmarshal(b, t.ReportReasonChildAbuse)
	case "reportReasonCopyright":
		t.ReportReasonCopyright = new(ReportReasonCopyright)
		return json.Unmarshal(b, t.ReportReasonCopyright)
	case "reportReasonCustom":
		t.ReportReasonCustom = new(ReportReasonCustom)
		return json.Unmarshal(b, t.ReportReasonCustom)
	case "reportReasonFake":
		t.ReportReasonFake = new(ReportReasonFake)
		return json.Unmarshal(b, t.ReportReasonFake)
	case "reportReasonIllegalDrugs":
		t.ReportReasonIllegalDrugs = new(ReportReasonIllegalDrugs)
		return json.Unmarshal(b, t.ReportReasonIllegalDrugs)
	case "reportReasonPersonalDetails":
		t.ReportReasonPersonalDetails = new(ReportReasonPersonalDetails)
		return json.Unmarshal(b, t.ReportReasonPersonalDetails)
	case "reportReasonPornography":
		t.ReportReasonPornography = new(ReportReasonPornography)
		return json.Unmarshal(b, t.ReportReasonPornography)
	case "reportReasonSpam":
		t.ReportReasonSpam = new(ReportReasonSpam)
		return json.Unmarshal(b, t.ReportReasonSpam)
	case "reportReasonUnrelatedLocation":
		t.ReportReasonUnrelatedLocation = new(ReportReasonUnrelatedLocation)
		return json.Unmarshal(b, t.ReportReasonUnrelatedLocation)
	case "reportReasonViolence":
		t.ReportReasonViolence = new(ReportReasonViolence)
		return json.Unmarshal(b, t.ReportReasonViolence)
	}
	return nil
}

func (t *ReportReason) MarshalJSON() ([]byte, error) {
	if t.ReportReasonChildAbuse != nil {
		return json.Marshal(t.ReportReasonChildAbuse)
	}
	if t.ReportReasonCopyright != nil {
		return json.Marshal(t.ReportReasonCopyright)
	}
	if t.ReportReasonCustom != nil {
		return json.Marshal(t.ReportReasonCustom)
	}
	if t.ReportReasonFake != nil {
		return json.Marshal(t.ReportReasonFake)
	}
	if t.ReportReasonIllegalDrugs != nil {
		return json.Marshal(t.ReportReasonIllegalDrugs)
	}
	if t.ReportReasonPersonalDetails != nil {
		return json.Marshal(t.ReportReasonPersonalDetails)
	}
	if t.ReportReasonPornography != nil {
		return json.Marshal(t.ReportReasonPornography)
	}
	if t.ReportReasonSpam != nil {
		return json.Marshal(t.ReportReasonSpam)
	}
	if t.ReportReasonUnrelatedLocation != nil {
		return json.Marshal(t.ReportReasonUnrelatedLocation)
	}
	if t.ReportReasonViolence != nil {
		return json.Marshal(t.ReportReasonViolence)
	}
	return nil, fmt.Errorf("no value set for ReportReason")
}

// ReportSponsoredResult Describes result of sponsored message or chat report
type ReportSponsoredResult struct {
	TypeStr                              string                                `json:"@type"`
	ReportSponsoredResultAdsHidden       *ReportSponsoredResultAdsHidden       `json:"reportSponsoredResultAdsHidden,omitempty"`
	ReportSponsoredResultFailed          *ReportSponsoredResultFailed          `json:"reportSponsoredResultFailed,omitempty"`
	ReportSponsoredResultOk              *ReportSponsoredResultOk              `json:"reportSponsoredResultOk,omitempty"`
	ReportSponsoredResultOptionRequired  *ReportSponsoredResultOptionRequired  `json:"reportSponsoredResultOptionRequired,omitempty"`
	ReportSponsoredResultPremiumRequired *ReportSponsoredResultPremiumRequired `json:"reportSponsoredResultPremiumRequired,omitempty"`
}

func (t *ReportSponsoredResult) Type() string {
	return t.TypeStr
}

func (t *ReportSponsoredResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReportSponsoredResult) GetExtra() string {
	return ""
}

func (t *ReportSponsoredResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reportSponsoredResultAdsHidden":
		t.ReportSponsoredResultAdsHidden = new(ReportSponsoredResultAdsHidden)
		return json.Unmarshal(b, t.ReportSponsoredResultAdsHidden)
	case "reportSponsoredResultFailed":
		t.ReportSponsoredResultFailed = new(ReportSponsoredResultFailed)
		return json.Unmarshal(b, t.ReportSponsoredResultFailed)
	case "reportSponsoredResultOk":
		t.ReportSponsoredResultOk = new(ReportSponsoredResultOk)
		return json.Unmarshal(b, t.ReportSponsoredResultOk)
	case "reportSponsoredResultOptionRequired":
		t.ReportSponsoredResultOptionRequired = new(ReportSponsoredResultOptionRequired)
		return json.Unmarshal(b, t.ReportSponsoredResultOptionRequired)
	case "reportSponsoredResultPremiumRequired":
		t.ReportSponsoredResultPremiumRequired = new(ReportSponsoredResultPremiumRequired)
		return json.Unmarshal(b, t.ReportSponsoredResultPremiumRequired)
	}
	return nil
}

func (t *ReportSponsoredResult) MarshalJSON() ([]byte, error) {
	if t.ReportSponsoredResultAdsHidden != nil {
		return json.Marshal(t.ReportSponsoredResultAdsHidden)
	}
	if t.ReportSponsoredResultFailed != nil {
		return json.Marshal(t.ReportSponsoredResultFailed)
	}
	if t.ReportSponsoredResultOk != nil {
		return json.Marshal(t.ReportSponsoredResultOk)
	}
	if t.ReportSponsoredResultOptionRequired != nil {
		return json.Marshal(t.ReportSponsoredResultOptionRequired)
	}
	if t.ReportSponsoredResultPremiumRequired != nil {
		return json.Marshal(t.ReportSponsoredResultPremiumRequired)
	}
	return nil, fmt.Errorf("no value set for ReportSponsoredResult")
}

// ReportStoryResult Describes result of story report
type ReportStoryResult struct {
	TypeStr                         string                           `json:"@type"`
	ReportStoryResultOk             *ReportStoryResultOk             `json:"reportStoryResultOk,omitempty"`
	ReportStoryResultOptionRequired *ReportStoryResultOptionRequired `json:"reportStoryResultOptionRequired,omitempty"`
	ReportStoryResultTextRequired   *ReportStoryResultTextRequired   `json:"reportStoryResultTextRequired,omitempty"`
}

func (t *ReportStoryResult) Type() string {
	return t.TypeStr
}

func (t *ReportStoryResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ReportStoryResult) GetExtra() string {
	return ""
}

func (t *ReportStoryResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "reportStoryResultOk":
		t.ReportStoryResultOk = new(ReportStoryResultOk)
		return json.Unmarshal(b, t.ReportStoryResultOk)
	case "reportStoryResultOptionRequired":
		t.ReportStoryResultOptionRequired = new(ReportStoryResultOptionRequired)
		return json.Unmarshal(b, t.ReportStoryResultOptionRequired)
	case "reportStoryResultTextRequired":
		t.ReportStoryResultTextRequired = new(ReportStoryResultTextRequired)
		return json.Unmarshal(b, t.ReportStoryResultTextRequired)
	}
	return nil
}

func (t *ReportStoryResult) MarshalJSON() ([]byte, error) {
	if t.ReportStoryResultOk != nil {
		return json.Marshal(t.ReportStoryResultOk)
	}
	if t.ReportStoryResultOptionRequired != nil {
		return json.Marshal(t.ReportStoryResultOptionRequired)
	}
	if t.ReportStoryResultTextRequired != nil {
		return json.Marshal(t.ReportStoryResultTextRequired)
	}
	return nil, fmt.Errorf("no value set for ReportStoryResult")
}

// ResendCodeReason Describes the reason why a code needs to be re-sent
type ResendCodeReason struct {
	TypeStr                            string                              `json:"@type"`
	ResendCodeReasonUserRequest        *ResendCodeReasonUserRequest        `json:"resendCodeReasonUserRequest,omitempty"`
	ResendCodeReasonVerificationFailed *ResendCodeReasonVerificationFailed `json:"resendCodeReasonVerificationFailed,omitempty"`
}

func (t *ResendCodeReason) Type() string {
	return t.TypeStr
}

func (t *ResendCodeReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ResendCodeReason) GetExtra() string {
	return ""
}

func (t *ResendCodeReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "resendCodeReasonUserRequest":
		t.ResendCodeReasonUserRequest = new(ResendCodeReasonUserRequest)
		return json.Unmarshal(b, t.ResendCodeReasonUserRequest)
	case "resendCodeReasonVerificationFailed":
		t.ResendCodeReasonVerificationFailed = new(ResendCodeReasonVerificationFailed)
		return json.Unmarshal(b, t.ResendCodeReasonVerificationFailed)
	}
	return nil
}

func (t *ResendCodeReason) MarshalJSON() ([]byte, error) {
	if t.ResendCodeReasonUserRequest != nil {
		return json.Marshal(t.ResendCodeReasonUserRequest)
	}
	if t.ResendCodeReasonVerificationFailed != nil {
		return json.Marshal(t.ResendCodeReasonVerificationFailed)
	}
	return nil, fmt.Errorf("no value set for ResendCodeReason")
}

// ResetPasswordResult Represents result of 2-step verification password reset
type ResetPasswordResult struct {
	TypeStr                     string                       `json:"@type"`
	ResetPasswordResultDeclined *ResetPasswordResultDeclined `json:"resetPasswordResultDeclined,omitempty"`
	ResetPasswordResultOk       *ResetPasswordResultOk       `json:"resetPasswordResultOk,omitempty"`
	ResetPasswordResultPending  *ResetPasswordResultPending  `json:"resetPasswordResultPending,omitempty"`
}

func (t *ResetPasswordResult) Type() string {
	return t.TypeStr
}

func (t *ResetPasswordResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ResetPasswordResult) GetExtra() string {
	return ""
}

func (t *ResetPasswordResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "resetPasswordResultDeclined":
		t.ResetPasswordResultDeclined = new(ResetPasswordResultDeclined)
		return json.Unmarshal(b, t.ResetPasswordResultDeclined)
	case "resetPasswordResultOk":
		t.ResetPasswordResultOk = new(ResetPasswordResultOk)
		return json.Unmarshal(b, t.ResetPasswordResultOk)
	case "resetPasswordResultPending":
		t.ResetPasswordResultPending = new(ResetPasswordResultPending)
		return json.Unmarshal(b, t.ResetPasswordResultPending)
	}
	return nil
}

func (t *ResetPasswordResult) MarshalJSON() ([]byte, error) {
	if t.ResetPasswordResultDeclined != nil {
		return json.Marshal(t.ResetPasswordResultDeclined)
	}
	if t.ResetPasswordResultOk != nil {
		return json.Marshal(t.ResetPasswordResultOk)
	}
	if t.ResetPasswordResultPending != nil {
		return json.Marshal(t.ResetPasswordResultPending)
	}
	return nil, fmt.Errorf("no value set for ResetPasswordResult")
}

// RevenueWithdrawalState Describes state of a revenue withdrawal
type RevenueWithdrawalState struct {
	TypeStr                         string                           `json:"@type"`
	RevenueWithdrawalStateFailed    *RevenueWithdrawalStateFailed    `json:"revenueWithdrawalStateFailed,omitempty"`
	RevenueWithdrawalStatePending   *RevenueWithdrawalStatePending   `json:"revenueWithdrawalStatePending,omitempty"`
	RevenueWithdrawalStateSucceeded *RevenueWithdrawalStateSucceeded `json:"revenueWithdrawalStateSucceeded,omitempty"`
}

func (t *RevenueWithdrawalState) Type() string {
	return t.TypeStr
}

func (t *RevenueWithdrawalState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *RevenueWithdrawalState) GetExtra() string {
	return ""
}

func (t *RevenueWithdrawalState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "revenueWithdrawalStateFailed":
		t.RevenueWithdrawalStateFailed = new(RevenueWithdrawalStateFailed)
		return json.Unmarshal(b, t.RevenueWithdrawalStateFailed)
	case "revenueWithdrawalStatePending":
		t.RevenueWithdrawalStatePending = new(RevenueWithdrawalStatePending)
		return json.Unmarshal(b, t.RevenueWithdrawalStatePending)
	case "revenueWithdrawalStateSucceeded":
		t.RevenueWithdrawalStateSucceeded = new(RevenueWithdrawalStateSucceeded)
		return json.Unmarshal(b, t.RevenueWithdrawalStateSucceeded)
	}
	return nil
}

func (t *RevenueWithdrawalState) MarshalJSON() ([]byte, error) {
	if t.RevenueWithdrawalStateFailed != nil {
		return json.Marshal(t.RevenueWithdrawalStateFailed)
	}
	if t.RevenueWithdrawalStatePending != nil {
		return json.Marshal(t.RevenueWithdrawalStatePending)
	}
	if t.RevenueWithdrawalStateSucceeded != nil {
		return json.Marshal(t.RevenueWithdrawalStateSucceeded)
	}
	return nil, fmt.Errorf("no value set for RevenueWithdrawalState")
}

// RichText Describes a formatted text object
type RichText struct {
	TypeStr               string                 `json:"@type"`
	RichTextAnchor        *RichTextAnchor        `json:"richTextAnchor,omitempty"`
	RichTextAnchorLink    *RichTextAnchorLink    `json:"richTextAnchorLink,omitempty"`
	RichTextBold          *RichTextBold          `json:"richTextBold,omitempty"`
	RichTextEmailAddress  *RichTextEmailAddress  `json:"richTextEmailAddress,omitempty"`
	RichTextFixed         *RichTextFixed         `json:"richTextFixed,omitempty"`
	RichTextIcon          *RichTextIcon          `json:"richTextIcon,omitempty"`
	RichTextItalic        *RichTextItalic        `json:"richTextItalic,omitempty"`
	RichTextMarked        *RichTextMarked        `json:"richTextMarked,omitempty"`
	RichTextPhoneNumber   *RichTextPhoneNumber   `json:"richTextPhoneNumber,omitempty"`
	RichTextPlain         *RichTextPlain         `json:"richTextPlain,omitempty"`
	RichTextReference     *RichTextReference     `json:"richTextReference,omitempty"`
	RichTextStrikethrough *RichTextStrikethrough `json:"richTextStrikethrough,omitempty"`
	RichTextSubscript     *RichTextSubscript     `json:"richTextSubscript,omitempty"`
	RichTextSuperscript   *RichTextSuperscript   `json:"richTextSuperscript,omitempty"`
	RichTextUnderline     *RichTextUnderline     `json:"richTextUnderline,omitempty"`
	RichTextUrl           *RichTextUrl           `json:"richTextUrl,omitempty"`
	RichTexts             *RichTexts             `json:"richTexts,omitempty"`
}

func (t *RichText) Type() string {
	return t.TypeStr
}

func (t *RichText) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *RichText) GetExtra() string {
	return ""
}

func (t *RichText) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "richTextAnchor":
		t.RichTextAnchor = new(RichTextAnchor)
		return json.Unmarshal(b, t.RichTextAnchor)
	case "richTextAnchorLink":
		t.RichTextAnchorLink = new(RichTextAnchorLink)
		return json.Unmarshal(b, t.RichTextAnchorLink)
	case "richTextBold":
		t.RichTextBold = new(RichTextBold)
		return json.Unmarshal(b, t.RichTextBold)
	case "richTextEmailAddress":
		t.RichTextEmailAddress = new(RichTextEmailAddress)
		return json.Unmarshal(b, t.RichTextEmailAddress)
	case "richTextFixed":
		t.RichTextFixed = new(RichTextFixed)
		return json.Unmarshal(b, t.RichTextFixed)
	case "richTextIcon":
		t.RichTextIcon = new(RichTextIcon)
		return json.Unmarshal(b, t.RichTextIcon)
	case "richTextItalic":
		t.RichTextItalic = new(RichTextItalic)
		return json.Unmarshal(b, t.RichTextItalic)
	case "richTextMarked":
		t.RichTextMarked = new(RichTextMarked)
		return json.Unmarshal(b, t.RichTextMarked)
	case "richTextPhoneNumber":
		t.RichTextPhoneNumber = new(RichTextPhoneNumber)
		return json.Unmarshal(b, t.RichTextPhoneNumber)
	case "richTextPlain":
		t.RichTextPlain = new(RichTextPlain)
		return json.Unmarshal(b, t.RichTextPlain)
	case "richTextReference":
		t.RichTextReference = new(RichTextReference)
		return json.Unmarshal(b, t.RichTextReference)
	case "richTextStrikethrough":
		t.RichTextStrikethrough = new(RichTextStrikethrough)
		return json.Unmarshal(b, t.RichTextStrikethrough)
	case "richTextSubscript":
		t.RichTextSubscript = new(RichTextSubscript)
		return json.Unmarshal(b, t.RichTextSubscript)
	case "richTextSuperscript":
		t.RichTextSuperscript = new(RichTextSuperscript)
		return json.Unmarshal(b, t.RichTextSuperscript)
	case "richTextUnderline":
		t.RichTextUnderline = new(RichTextUnderline)
		return json.Unmarshal(b, t.RichTextUnderline)
	case "richTextUrl":
		t.RichTextUrl = new(RichTextUrl)
		return json.Unmarshal(b, t.RichTextUrl)
	case "richTexts":
		t.RichTexts = new(RichTexts)
		return json.Unmarshal(b, t.RichTexts)
	}
	return nil
}

func (t *RichText) MarshalJSON() ([]byte, error) {
	if t.RichTextAnchor != nil {
		return json.Marshal(t.RichTextAnchor)
	}
	if t.RichTextAnchorLink != nil {
		return json.Marshal(t.RichTextAnchorLink)
	}
	if t.RichTextBold != nil {
		return json.Marshal(t.RichTextBold)
	}
	if t.RichTextEmailAddress != nil {
		return json.Marshal(t.RichTextEmailAddress)
	}
	if t.RichTextFixed != nil {
		return json.Marshal(t.RichTextFixed)
	}
	if t.RichTextIcon != nil {
		return json.Marshal(t.RichTextIcon)
	}
	if t.RichTextItalic != nil {
		return json.Marshal(t.RichTextItalic)
	}
	if t.RichTextMarked != nil {
		return json.Marshal(t.RichTextMarked)
	}
	if t.RichTextPhoneNumber != nil {
		return json.Marshal(t.RichTextPhoneNumber)
	}
	if t.RichTextPlain != nil {
		return json.Marshal(t.RichTextPlain)
	}
	if t.RichTextReference != nil {
		return json.Marshal(t.RichTextReference)
	}
	if t.RichTextStrikethrough != nil {
		return json.Marshal(t.RichTextStrikethrough)
	}
	if t.RichTextSubscript != nil {
		return json.Marshal(t.RichTextSubscript)
	}
	if t.RichTextSuperscript != nil {
		return json.Marshal(t.RichTextSuperscript)
	}
	if t.RichTextUnderline != nil {
		return json.Marshal(t.RichTextUnderline)
	}
	if t.RichTextUrl != nil {
		return json.Marshal(t.RichTextUrl)
	}
	if t.RichTexts != nil {
		return json.Marshal(t.RichTexts)
	}
	return nil, fmt.Errorf("no value set for RichText")
}

// SavedMessagesTopicType Describes type of Saved Messages topic
type SavedMessagesTopicType struct {
	TypeStr                             string                               `json:"@type"`
	SavedMessagesTopicTypeAuthorHidden  *SavedMessagesTopicTypeAuthorHidden  `json:"savedMessagesTopicTypeAuthorHidden,omitempty"`
	SavedMessagesTopicTypeMyNotes       *SavedMessagesTopicTypeMyNotes       `json:"savedMessagesTopicTypeMyNotes,omitempty"`
	SavedMessagesTopicTypeSavedFromChat *SavedMessagesTopicTypeSavedFromChat `json:"savedMessagesTopicTypeSavedFromChat,omitempty"`
}

func (t *SavedMessagesTopicType) Type() string {
	return t.TypeStr
}

func (t *SavedMessagesTopicType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SavedMessagesTopicType) GetExtra() string {
	return ""
}

func (t *SavedMessagesTopicType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "savedMessagesTopicTypeAuthorHidden":
		t.SavedMessagesTopicTypeAuthorHidden = new(SavedMessagesTopicTypeAuthorHidden)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeAuthorHidden)
	case "savedMessagesTopicTypeMyNotes":
		t.SavedMessagesTopicTypeMyNotes = new(SavedMessagesTopicTypeMyNotes)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeMyNotes)
	case "savedMessagesTopicTypeSavedFromChat":
		t.SavedMessagesTopicTypeSavedFromChat = new(SavedMessagesTopicTypeSavedFromChat)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeSavedFromChat)
	}
	return nil
}

func (t *SavedMessagesTopicType) MarshalJSON() ([]byte, error) {
	if t.SavedMessagesTopicTypeAuthorHidden != nil {
		return json.Marshal(t.SavedMessagesTopicTypeAuthorHidden)
	}
	if t.SavedMessagesTopicTypeMyNotes != nil {
		return json.Marshal(t.SavedMessagesTopicTypeMyNotes)
	}
	if t.SavedMessagesTopicTypeSavedFromChat != nil {
		return json.Marshal(t.SavedMessagesTopicTypeSavedFromChat)
	}
	return nil, fmt.Errorf("no value set for SavedMessagesTopicType")
}

// SearchMessagesChatTypeFilter Represents a filter for type of the chats in which to search messages
type SearchMessagesChatTypeFilter struct {
	TypeStr                             string                               `json:"@type"`
	SearchMessagesChatTypeFilterChannel *SearchMessagesChatTypeFilterChannel `json:"searchMessagesChatTypeFilterChannel,omitempty"`
	SearchMessagesChatTypeFilterGroup   *SearchMessagesChatTypeFilterGroup   `json:"searchMessagesChatTypeFilterGroup,omitempty"`
	SearchMessagesChatTypeFilterPrivate *SearchMessagesChatTypeFilterPrivate `json:"searchMessagesChatTypeFilterPrivate,omitempty"`
}

func (t *SearchMessagesChatTypeFilter) Type() string {
	return t.TypeStr
}

func (t *SearchMessagesChatTypeFilter) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SearchMessagesChatTypeFilter) GetExtra() string {
	return ""
}

func (t *SearchMessagesChatTypeFilter) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "searchMessagesChatTypeFilterChannel":
		t.SearchMessagesChatTypeFilterChannel = new(SearchMessagesChatTypeFilterChannel)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterChannel)
	case "searchMessagesChatTypeFilterGroup":
		t.SearchMessagesChatTypeFilterGroup = new(SearchMessagesChatTypeFilterGroup)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterGroup)
	case "searchMessagesChatTypeFilterPrivate":
		t.SearchMessagesChatTypeFilterPrivate = new(SearchMessagesChatTypeFilterPrivate)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterPrivate)
	}
	return nil
}

func (t *SearchMessagesChatTypeFilter) MarshalJSON() ([]byte, error) {
	if t.SearchMessagesChatTypeFilterChannel != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterChannel)
	}
	if t.SearchMessagesChatTypeFilterGroup != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterGroup)
	}
	if t.SearchMessagesChatTypeFilterPrivate != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterPrivate)
	}
	return nil, fmt.Errorf("no value set for SearchMessagesChatTypeFilter")
}

// SearchMessagesFilter Represents a filter for message search results
type SearchMessagesFilter struct {
	TypeStr                               string                                 `json:"@type"`
	SearchMessagesFilterAnimation         *SearchMessagesFilterAnimation         `json:"searchMessagesFilterAnimation,omitempty"`
	SearchMessagesFilterAudio             *SearchMessagesFilterAudio             `json:"searchMessagesFilterAudio,omitempty"`
	SearchMessagesFilterChatPhoto         *SearchMessagesFilterChatPhoto         `json:"searchMessagesFilterChatPhoto,omitempty"`
	SearchMessagesFilterDocument          *SearchMessagesFilterDocument          `json:"searchMessagesFilterDocument,omitempty"`
	SearchMessagesFilterEmpty             *SearchMessagesFilterEmpty             `json:"searchMessagesFilterEmpty,omitempty"`
	SearchMessagesFilterFailedToSend      *SearchMessagesFilterFailedToSend      `json:"searchMessagesFilterFailedToSend,omitempty"`
	SearchMessagesFilterMention           *SearchMessagesFilterMention           `json:"searchMessagesFilterMention,omitempty"`
	SearchMessagesFilterPhoto             *SearchMessagesFilterPhoto             `json:"searchMessagesFilterPhoto,omitempty"`
	SearchMessagesFilterPhotoAndVideo     *SearchMessagesFilterPhotoAndVideo     `json:"searchMessagesFilterPhotoAndVideo,omitempty"`
	SearchMessagesFilterPinned            *SearchMessagesFilterPinned            `json:"searchMessagesFilterPinned,omitempty"`
	SearchMessagesFilterUnreadMention     *SearchMessagesFilterUnreadMention     `json:"searchMessagesFilterUnreadMention,omitempty"`
	SearchMessagesFilterUnreadReaction    *SearchMessagesFilterUnreadReaction    `json:"searchMessagesFilterUnreadReaction,omitempty"`
	SearchMessagesFilterUrl               *SearchMessagesFilterUrl               `json:"searchMessagesFilterUrl,omitempty"`
	SearchMessagesFilterVideo             *SearchMessagesFilterVideo             `json:"searchMessagesFilterVideo,omitempty"`
	SearchMessagesFilterVideoNote         *SearchMessagesFilterVideoNote         `json:"searchMessagesFilterVideoNote,omitempty"`
	SearchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote `json:"searchMessagesFilterVoiceAndVideoNote,omitempty"`
	SearchMessagesFilterVoiceNote         *SearchMessagesFilterVoiceNote         `json:"searchMessagesFilterVoiceNote,omitempty"`
}

func (t *SearchMessagesFilter) Type() string {
	return t.TypeStr
}

func (t *SearchMessagesFilter) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SearchMessagesFilter) GetExtra() string {
	return ""
}

func (t *SearchMessagesFilter) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "searchMessagesFilterAnimation":
		t.SearchMessagesFilterAnimation = new(SearchMessagesFilterAnimation)
		return json.Unmarshal(b, t.SearchMessagesFilterAnimation)
	case "searchMessagesFilterAudio":
		t.SearchMessagesFilterAudio = new(SearchMessagesFilterAudio)
		return json.Unmarshal(b, t.SearchMessagesFilterAudio)
	case "searchMessagesFilterChatPhoto":
		t.SearchMessagesFilterChatPhoto = new(SearchMessagesFilterChatPhoto)
		return json.Unmarshal(b, t.SearchMessagesFilterChatPhoto)
	case "searchMessagesFilterDocument":
		t.SearchMessagesFilterDocument = new(SearchMessagesFilterDocument)
		return json.Unmarshal(b, t.SearchMessagesFilterDocument)
	case "searchMessagesFilterEmpty":
		t.SearchMessagesFilterEmpty = new(SearchMessagesFilterEmpty)
		return json.Unmarshal(b, t.SearchMessagesFilterEmpty)
	case "searchMessagesFilterFailedToSend":
		t.SearchMessagesFilterFailedToSend = new(SearchMessagesFilterFailedToSend)
		return json.Unmarshal(b, t.SearchMessagesFilterFailedToSend)
	case "searchMessagesFilterMention":
		t.SearchMessagesFilterMention = new(SearchMessagesFilterMention)
		return json.Unmarshal(b, t.SearchMessagesFilterMention)
	case "searchMessagesFilterPhoto":
		t.SearchMessagesFilterPhoto = new(SearchMessagesFilterPhoto)
		return json.Unmarshal(b, t.SearchMessagesFilterPhoto)
	case "searchMessagesFilterPhotoAndVideo":
		t.SearchMessagesFilterPhotoAndVideo = new(SearchMessagesFilterPhotoAndVideo)
		return json.Unmarshal(b, t.SearchMessagesFilterPhotoAndVideo)
	case "searchMessagesFilterPinned":
		t.SearchMessagesFilterPinned = new(SearchMessagesFilterPinned)
		return json.Unmarshal(b, t.SearchMessagesFilterPinned)
	case "searchMessagesFilterUnreadMention":
		t.SearchMessagesFilterUnreadMention = new(SearchMessagesFilterUnreadMention)
		return json.Unmarshal(b, t.SearchMessagesFilterUnreadMention)
	case "searchMessagesFilterUnreadReaction":
		t.SearchMessagesFilterUnreadReaction = new(SearchMessagesFilterUnreadReaction)
		return json.Unmarshal(b, t.SearchMessagesFilterUnreadReaction)
	case "searchMessagesFilterUrl":
		t.SearchMessagesFilterUrl = new(SearchMessagesFilterUrl)
		return json.Unmarshal(b, t.SearchMessagesFilterUrl)
	case "searchMessagesFilterVideo":
		t.SearchMessagesFilterVideo = new(SearchMessagesFilterVideo)
		return json.Unmarshal(b, t.SearchMessagesFilterVideo)
	case "searchMessagesFilterVideoNote":
		t.SearchMessagesFilterVideoNote = new(SearchMessagesFilterVideoNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVideoNote)
	case "searchMessagesFilterVoiceAndVideoNote":
		t.SearchMessagesFilterVoiceAndVideoNote = new(SearchMessagesFilterVoiceAndVideoNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVoiceAndVideoNote)
	case "searchMessagesFilterVoiceNote":
		t.SearchMessagesFilterVoiceNote = new(SearchMessagesFilterVoiceNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVoiceNote)
	}
	return nil
}

func (t *SearchMessagesFilter) MarshalJSON() ([]byte, error) {
	if t.SearchMessagesFilterAnimation != nil {
		return json.Marshal(t.SearchMessagesFilterAnimation)
	}
	if t.SearchMessagesFilterAudio != nil {
		return json.Marshal(t.SearchMessagesFilterAudio)
	}
	if t.SearchMessagesFilterChatPhoto != nil {
		return json.Marshal(t.SearchMessagesFilterChatPhoto)
	}
	if t.SearchMessagesFilterDocument != nil {
		return json.Marshal(t.SearchMessagesFilterDocument)
	}
	if t.SearchMessagesFilterEmpty != nil {
		return json.Marshal(t.SearchMessagesFilterEmpty)
	}
	if t.SearchMessagesFilterFailedToSend != nil {
		return json.Marshal(t.SearchMessagesFilterFailedToSend)
	}
	if t.SearchMessagesFilterMention != nil {
		return json.Marshal(t.SearchMessagesFilterMention)
	}
	if t.SearchMessagesFilterPhoto != nil {
		return json.Marshal(t.SearchMessagesFilterPhoto)
	}
	if t.SearchMessagesFilterPhotoAndVideo != nil {
		return json.Marshal(t.SearchMessagesFilterPhotoAndVideo)
	}
	if t.SearchMessagesFilterPinned != nil {
		return json.Marshal(t.SearchMessagesFilterPinned)
	}
	if t.SearchMessagesFilterUnreadMention != nil {
		return json.Marshal(t.SearchMessagesFilterUnreadMention)
	}
	if t.SearchMessagesFilterUnreadReaction != nil {
		return json.Marshal(t.SearchMessagesFilterUnreadReaction)
	}
	if t.SearchMessagesFilterUrl != nil {
		return json.Marshal(t.SearchMessagesFilterUrl)
	}
	if t.SearchMessagesFilterVideo != nil {
		return json.Marshal(t.SearchMessagesFilterVideo)
	}
	if t.SearchMessagesFilterVideoNote != nil {
		return json.Marshal(t.SearchMessagesFilterVideoNote)
	}
	if t.SearchMessagesFilterVoiceAndVideoNote != nil {
		return json.Marshal(t.SearchMessagesFilterVoiceAndVideoNote)
	}
	if t.SearchMessagesFilterVoiceNote != nil {
		return json.Marshal(t.SearchMessagesFilterVoiceNote)
	}
	return nil, fmt.Errorf("no value set for SearchMessagesFilter")
}

// SecretChatState Describes the current secret chat state
type SecretChatState struct {
	TypeStr                string                  `json:"@type"`
	SecretChatStateClosed  *SecretChatStateClosed  `json:"secretChatStateClosed,omitempty"`
	SecretChatStatePending *SecretChatStatePending `json:"secretChatStatePending,omitempty"`
	SecretChatStateReady   *SecretChatStateReady   `json:"secretChatStateReady,omitempty"`
}

func (t *SecretChatState) Type() string {
	return t.TypeStr
}

func (t *SecretChatState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SecretChatState) GetExtra() string {
	return ""
}

func (t *SecretChatState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "secretChatStateClosed":
		t.SecretChatStateClosed = new(SecretChatStateClosed)
		return json.Unmarshal(b, t.SecretChatStateClosed)
	case "secretChatStatePending":
		t.SecretChatStatePending = new(SecretChatStatePending)
		return json.Unmarshal(b, t.SecretChatStatePending)
	case "secretChatStateReady":
		t.SecretChatStateReady = new(SecretChatStateReady)
		return json.Unmarshal(b, t.SecretChatStateReady)
	}
	return nil
}

func (t *SecretChatState) MarshalJSON() ([]byte, error) {
	if t.SecretChatStateClosed != nil {
		return json.Marshal(t.SecretChatStateClosed)
	}
	if t.SecretChatStatePending != nil {
		return json.Marshal(t.SecretChatStatePending)
	}
	if t.SecretChatStateReady != nil {
		return json.Marshal(t.SecretChatStateReady)
	}
	return nil, fmt.Errorf("no value set for SecretChatState")
}

// SentGift Represents content of a gift received by a user or a channel chat
type SentGift struct {
	TypeStr          string            `json:"@type"`
	SentGiftRegular  *SentGiftRegular  `json:"sentGiftRegular,omitempty"`
	SentGiftUpgraded *SentGiftUpgraded `json:"sentGiftUpgraded,omitempty"`
}

func (t *SentGift) Type() string {
	return t.TypeStr
}

func (t *SentGift) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SentGift) GetExtra() string {
	return ""
}

func (t *SentGift) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "sentGiftRegular":
		t.SentGiftRegular = new(SentGiftRegular)
		return json.Unmarshal(b, t.SentGiftRegular)
	case "sentGiftUpgraded":
		t.SentGiftUpgraded = new(SentGiftUpgraded)
		return json.Unmarshal(b, t.SentGiftUpgraded)
	}
	return nil
}

func (t *SentGift) MarshalJSON() ([]byte, error) {
	if t.SentGiftRegular != nil {
		return json.Marshal(t.SentGiftRegular)
	}
	if t.SentGiftUpgraded != nil {
		return json.Marshal(t.SentGiftUpgraded)
	}
	return nil, fmt.Errorf("no value set for SentGift")
}

// SessionType Represents the type of session
type SessionType struct {
	TypeStr            string              `json:"@type"`
	SessionTypeAndroid *SessionTypeAndroid `json:"sessionTypeAndroid,omitempty"`
	SessionTypeApple   *SessionTypeApple   `json:"sessionTypeApple,omitempty"`
	SessionTypeBrave   *SessionTypeBrave   `json:"sessionTypeBrave,omitempty"`
	SessionTypeChrome  *SessionTypeChrome  `json:"sessionTypeChrome,omitempty"`
	SessionTypeEdge    *SessionTypeEdge    `json:"sessionTypeEdge,omitempty"`
	SessionTypeFirefox *SessionTypeFirefox `json:"sessionTypeFirefox,omitempty"`
	SessionTypeIpad    *SessionTypeIpad    `json:"sessionTypeIpad,omitempty"`
	SessionTypeIphone  *SessionTypeIphone  `json:"sessionTypeIphone,omitempty"`
	SessionTypeLinux   *SessionTypeLinux   `json:"sessionTypeLinux,omitempty"`
	SessionTypeMac     *SessionTypeMac     `json:"sessionTypeMac,omitempty"`
	SessionTypeOpera   *SessionTypeOpera   `json:"sessionTypeOpera,omitempty"`
	SessionTypeSafari  *SessionTypeSafari  `json:"sessionTypeSafari,omitempty"`
	SessionTypeUbuntu  *SessionTypeUbuntu  `json:"sessionTypeUbuntu,omitempty"`
	SessionTypeUnknown *SessionTypeUnknown `json:"sessionTypeUnknown,omitempty"`
	SessionTypeVivaldi *SessionTypeVivaldi `json:"sessionTypeVivaldi,omitempty"`
	SessionTypeWindows *SessionTypeWindows `json:"sessionTypeWindows,omitempty"`
	SessionTypeXbox    *SessionTypeXbox    `json:"sessionTypeXbox,omitempty"`
}

func (t *SessionType) Type() string {
	return t.TypeStr
}

func (t *SessionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SessionType) GetExtra() string {
	return ""
}

func (t *SessionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "sessionTypeAndroid":
		t.SessionTypeAndroid = new(SessionTypeAndroid)
		return json.Unmarshal(b, t.SessionTypeAndroid)
	case "sessionTypeApple":
		t.SessionTypeApple = new(SessionTypeApple)
		return json.Unmarshal(b, t.SessionTypeApple)
	case "sessionTypeBrave":
		t.SessionTypeBrave = new(SessionTypeBrave)
		return json.Unmarshal(b, t.SessionTypeBrave)
	case "sessionTypeChrome":
		t.SessionTypeChrome = new(SessionTypeChrome)
		return json.Unmarshal(b, t.SessionTypeChrome)
	case "sessionTypeEdge":
		t.SessionTypeEdge = new(SessionTypeEdge)
		return json.Unmarshal(b, t.SessionTypeEdge)
	case "sessionTypeFirefox":
		t.SessionTypeFirefox = new(SessionTypeFirefox)
		return json.Unmarshal(b, t.SessionTypeFirefox)
	case "sessionTypeIpad":
		t.SessionTypeIpad = new(SessionTypeIpad)
		return json.Unmarshal(b, t.SessionTypeIpad)
	case "sessionTypeIphone":
		t.SessionTypeIphone = new(SessionTypeIphone)
		return json.Unmarshal(b, t.SessionTypeIphone)
	case "sessionTypeLinux":
		t.SessionTypeLinux = new(SessionTypeLinux)
		return json.Unmarshal(b, t.SessionTypeLinux)
	case "sessionTypeMac":
		t.SessionTypeMac = new(SessionTypeMac)
		return json.Unmarshal(b, t.SessionTypeMac)
	case "sessionTypeOpera":
		t.SessionTypeOpera = new(SessionTypeOpera)
		return json.Unmarshal(b, t.SessionTypeOpera)
	case "sessionTypeSafari":
		t.SessionTypeSafari = new(SessionTypeSafari)
		return json.Unmarshal(b, t.SessionTypeSafari)
	case "sessionTypeUbuntu":
		t.SessionTypeUbuntu = new(SessionTypeUbuntu)
		return json.Unmarshal(b, t.SessionTypeUbuntu)
	case "sessionTypeUnknown":
		t.SessionTypeUnknown = new(SessionTypeUnknown)
		return json.Unmarshal(b, t.SessionTypeUnknown)
	case "sessionTypeVivaldi":
		t.SessionTypeVivaldi = new(SessionTypeVivaldi)
		return json.Unmarshal(b, t.SessionTypeVivaldi)
	case "sessionTypeWindows":
		t.SessionTypeWindows = new(SessionTypeWindows)
		return json.Unmarshal(b, t.SessionTypeWindows)
	case "sessionTypeXbox":
		t.SessionTypeXbox = new(SessionTypeXbox)
		return json.Unmarshal(b, t.SessionTypeXbox)
	}
	return nil
}

func (t *SessionType) MarshalJSON() ([]byte, error) {
	if t.SessionTypeAndroid != nil {
		return json.Marshal(t.SessionTypeAndroid)
	}
	if t.SessionTypeApple != nil {
		return json.Marshal(t.SessionTypeApple)
	}
	if t.SessionTypeBrave != nil {
		return json.Marshal(t.SessionTypeBrave)
	}
	if t.SessionTypeChrome != nil {
		return json.Marshal(t.SessionTypeChrome)
	}
	if t.SessionTypeEdge != nil {
		return json.Marshal(t.SessionTypeEdge)
	}
	if t.SessionTypeFirefox != nil {
		return json.Marshal(t.SessionTypeFirefox)
	}
	if t.SessionTypeIpad != nil {
		return json.Marshal(t.SessionTypeIpad)
	}
	if t.SessionTypeIphone != nil {
		return json.Marshal(t.SessionTypeIphone)
	}
	if t.SessionTypeLinux != nil {
		return json.Marshal(t.SessionTypeLinux)
	}
	if t.SessionTypeMac != nil {
		return json.Marshal(t.SessionTypeMac)
	}
	if t.SessionTypeOpera != nil {
		return json.Marshal(t.SessionTypeOpera)
	}
	if t.SessionTypeSafari != nil {
		return json.Marshal(t.SessionTypeSafari)
	}
	if t.SessionTypeUbuntu != nil {
		return json.Marshal(t.SessionTypeUbuntu)
	}
	if t.SessionTypeUnknown != nil {
		return json.Marshal(t.SessionTypeUnknown)
	}
	if t.SessionTypeVivaldi != nil {
		return json.Marshal(t.SessionTypeVivaldi)
	}
	if t.SessionTypeWindows != nil {
		return json.Marshal(t.SessionTypeWindows)
	}
	if t.SessionTypeXbox != nil {
		return json.Marshal(t.SessionTypeXbox)
	}
	return nil, fmt.Errorf("no value set for SessionType")
}

// SpeechRecognitionResult Describes result of speech recognition in a voice note
type SpeechRecognitionResult struct {
	TypeStr                        string                          `json:"@type"`
	SpeechRecognitionResultError   *SpeechRecognitionResultError   `json:"speechRecognitionResultError,omitempty"`
	SpeechRecognitionResultPending *SpeechRecognitionResultPending `json:"speechRecognitionResultPending,omitempty"`
	SpeechRecognitionResultText    *SpeechRecognitionResultText    `json:"speechRecognitionResultText,omitempty"`
}

func (t *SpeechRecognitionResult) Type() string {
	return t.TypeStr
}

func (t *SpeechRecognitionResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SpeechRecognitionResult) GetExtra() string {
	return ""
}

func (t *SpeechRecognitionResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "speechRecognitionResultError":
		t.SpeechRecognitionResultError = new(SpeechRecognitionResultError)
		return json.Unmarshal(b, t.SpeechRecognitionResultError)
	case "speechRecognitionResultPending":
		t.SpeechRecognitionResultPending = new(SpeechRecognitionResultPending)
		return json.Unmarshal(b, t.SpeechRecognitionResultPending)
	case "speechRecognitionResultText":
		t.SpeechRecognitionResultText = new(SpeechRecognitionResultText)
		return json.Unmarshal(b, t.SpeechRecognitionResultText)
	}
	return nil
}

func (t *SpeechRecognitionResult) MarshalJSON() ([]byte, error) {
	if t.SpeechRecognitionResultError != nil {
		return json.Marshal(t.SpeechRecognitionResultError)
	}
	if t.SpeechRecognitionResultPending != nil {
		return json.Marshal(t.SpeechRecognitionResultPending)
	}
	if t.SpeechRecognitionResultText != nil {
		return json.Marshal(t.SpeechRecognitionResultText)
	}
	return nil, fmt.Errorf("no value set for SpeechRecognitionResult")
}

// StarSubscriptionType Describes type of subscription paid in Telegram Stars
type StarSubscriptionType struct {
	TypeStr                     string                       `json:"@type"`
	StarSubscriptionTypeBot     *StarSubscriptionTypeBot     `json:"starSubscriptionTypeBot,omitempty"`
	StarSubscriptionTypeChannel *StarSubscriptionTypeChannel `json:"starSubscriptionTypeChannel,omitempty"`
}

func (t *StarSubscriptionType) Type() string {
	return t.TypeStr
}

func (t *StarSubscriptionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StarSubscriptionType) GetExtra() string {
	return ""
}

func (t *StarSubscriptionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "starSubscriptionTypeBot":
		t.StarSubscriptionTypeBot = new(StarSubscriptionTypeBot)
		return json.Unmarshal(b, t.StarSubscriptionTypeBot)
	case "starSubscriptionTypeChannel":
		t.StarSubscriptionTypeChannel = new(StarSubscriptionTypeChannel)
		return json.Unmarshal(b, t.StarSubscriptionTypeChannel)
	}
	return nil
}

func (t *StarSubscriptionType) MarshalJSON() ([]byte, error) {
	if t.StarSubscriptionTypeBot != nil {
		return json.Marshal(t.StarSubscriptionTypeBot)
	}
	if t.StarSubscriptionTypeChannel != nil {
		return json.Marshal(t.StarSubscriptionTypeChannel)
	}
	return nil, fmt.Errorf("no value set for StarSubscriptionType")
}

// StartLiveStoryResult Represents result of starting a live story
type StartLiveStoryResult struct {
	TypeStr                  string                    `json:"@type"`
	StartLiveStoryResultFail *StartLiveStoryResultFail `json:"startLiveStoryResultFail,omitempty"`
	StartLiveStoryResultOk   *StartLiveStoryResultOk   `json:"startLiveStoryResultOk,omitempty"`
}

func (t *StartLiveStoryResult) Type() string {
	return t.TypeStr
}

func (t *StartLiveStoryResult) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StartLiveStoryResult) GetExtra() string {
	return ""
}

func (t *StartLiveStoryResult) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "startLiveStoryResultFail":
		t.StartLiveStoryResultFail = new(StartLiveStoryResultFail)
		return json.Unmarshal(b, t.StartLiveStoryResultFail)
	case "startLiveStoryResultOk":
		t.StartLiveStoryResultOk = new(StartLiveStoryResultOk)
		return json.Unmarshal(b, t.StartLiveStoryResultOk)
	}
	return nil
}

func (t *StartLiveStoryResult) MarshalJSON() ([]byte, error) {
	if t.StartLiveStoryResultFail != nil {
		return json.Marshal(t.StartLiveStoryResultFail)
	}
	if t.StartLiveStoryResultOk != nil {
		return json.Marshal(t.StartLiveStoryResultOk)
	}
	return nil, fmt.Errorf("no value set for StartLiveStoryResult")
}

// StarTransactionType Describes type of transaction with Telegram Stars
type StarTransactionType struct {
	TypeStr                                         string                                           `json:"@type"`
	StarTransactionTypeAffiliateProgramCommission   *StarTransactionTypeAffiliateProgramCommission   `json:"starTransactionTypeAffiliateProgramCommission,omitempty"`
	StarTransactionTypeAppStoreDeposit              *StarTransactionTypeAppStoreDeposit              `json:"starTransactionTypeAppStoreDeposit,omitempty"`
	StarTransactionTypeBotInvoicePurchase           *StarTransactionTypeBotInvoicePurchase           `json:"starTransactionTypeBotInvoicePurchase,omitempty"`
	StarTransactionTypeBotInvoiceSale               *StarTransactionTypeBotInvoiceSale               `json:"starTransactionTypeBotInvoiceSale,omitempty"`
	StarTransactionTypeBotPaidMediaPurchase         *StarTransactionTypeBotPaidMediaPurchase         `json:"starTransactionTypeBotPaidMediaPurchase,omitempty"`
	StarTransactionTypeBotPaidMediaSale             *StarTransactionTypeBotPaidMediaSale             `json:"starTransactionTypeBotPaidMediaSale,omitempty"`
	StarTransactionTypeBotSubscriptionPurchase      *StarTransactionTypeBotSubscriptionPurchase      `json:"starTransactionTypeBotSubscriptionPurchase,omitempty"`
	StarTransactionTypeBotSubscriptionSale          *StarTransactionTypeBotSubscriptionSale          `json:"starTransactionTypeBotSubscriptionSale,omitempty"`
	StarTransactionTypeBusinessBotTransferReceive   *StarTransactionTypeBusinessBotTransferReceive   `json:"starTransactionTypeBusinessBotTransferReceive,omitempty"`
	StarTransactionTypeBusinessBotTransferSend      *StarTransactionTypeBusinessBotTransferSend      `json:"starTransactionTypeBusinessBotTransferSend,omitempty"`
	StarTransactionTypeChannelPaidMediaPurchase     *StarTransactionTypeChannelPaidMediaPurchase     `json:"starTransactionTypeChannelPaidMediaPurchase,omitempty"`
	StarTransactionTypeChannelPaidMediaSale         *StarTransactionTypeChannelPaidMediaSale         `json:"starTransactionTypeChannelPaidMediaSale,omitempty"`
	StarTransactionTypeChannelPaidReactionReceive   *StarTransactionTypeChannelPaidReactionReceive   `json:"starTransactionTypeChannelPaidReactionReceive,omitempty"`
	StarTransactionTypeChannelPaidReactionSend      *StarTransactionTypeChannelPaidReactionSend      `json:"starTransactionTypeChannelPaidReactionSend,omitempty"`
	StarTransactionTypeChannelSubscriptionPurchase  *StarTransactionTypeChannelSubscriptionPurchase  `json:"starTransactionTypeChannelSubscriptionPurchase,omitempty"`
	StarTransactionTypeChannelSubscriptionSale      *StarTransactionTypeChannelSubscriptionSale      `json:"starTransactionTypeChannelSubscriptionSale,omitempty"`
	StarTransactionTypeFragmentDeposit              *StarTransactionTypeFragmentDeposit              `json:"starTransactionTypeFragmentDeposit,omitempty"`
	StarTransactionTypeFragmentWithdrawal           *StarTransactionTypeFragmentWithdrawal           `json:"starTransactionTypeFragmentWithdrawal,omitempty"`
	StarTransactionTypeGiftAuctionBid               *StarTransactionTypeGiftAuctionBid               `json:"starTransactionTypeGiftAuctionBid,omitempty"`
	StarTransactionTypeGiftOriginalDetailsDrop      *StarTransactionTypeGiftOriginalDetailsDrop      `json:"starTransactionTypeGiftOriginalDetailsDrop,omitempty"`
	StarTransactionTypeGiftPurchase                 *StarTransactionTypeGiftPurchase                 `json:"starTransactionTypeGiftPurchase,omitempty"`
	StarTransactionTypeGiftPurchaseOffer            *StarTransactionTypeGiftPurchaseOffer            `json:"starTransactionTypeGiftPurchaseOffer,omitempty"`
	StarTransactionTypeGiftSale                     *StarTransactionTypeGiftSale                     `json:"starTransactionTypeGiftSale,omitempty"`
	StarTransactionTypeGiftTransfer                 *StarTransactionTypeGiftTransfer                 `json:"starTransactionTypeGiftTransfer,omitempty"`
	StarTransactionTypeGiftUpgrade                  *StarTransactionTypeGiftUpgrade                  `json:"starTransactionTypeGiftUpgrade,omitempty"`
	StarTransactionTypeGiftUpgradePurchase          *StarTransactionTypeGiftUpgradePurchase          `json:"starTransactionTypeGiftUpgradePurchase,omitempty"`
	StarTransactionTypeGiveawayDeposit              *StarTransactionTypeGiveawayDeposit              `json:"starTransactionTypeGiveawayDeposit,omitempty"`
	StarTransactionTypeGooglePlayDeposit            *StarTransactionTypeGooglePlayDeposit            `json:"starTransactionTypeGooglePlayDeposit,omitempty"`
	StarTransactionTypePaidGroupCallMessageReceive  *StarTransactionTypePaidGroupCallMessageReceive  `json:"starTransactionTypePaidGroupCallMessageReceive,omitempty"`
	StarTransactionTypePaidGroupCallMessageSend     *StarTransactionTypePaidGroupCallMessageSend     `json:"starTransactionTypePaidGroupCallMessageSend,omitempty"`
	StarTransactionTypePaidGroupCallReactionReceive *StarTransactionTypePaidGroupCallReactionReceive `json:"starTransactionTypePaidGroupCallReactionReceive,omitempty"`
	StarTransactionTypePaidGroupCallReactionSend    *StarTransactionTypePaidGroupCallReactionSend    `json:"starTransactionTypePaidGroupCallReactionSend,omitempty"`
	StarTransactionTypePaidMessageReceive           *StarTransactionTypePaidMessageReceive           `json:"starTransactionTypePaidMessageReceive,omitempty"`
	StarTransactionTypePaidMessageSend              *StarTransactionTypePaidMessageSend              `json:"starTransactionTypePaidMessageSend,omitempty"`
	StarTransactionTypePremiumBotDeposit            *StarTransactionTypePremiumBotDeposit            `json:"starTransactionTypePremiumBotDeposit,omitempty"`
	StarTransactionTypePremiumPurchase              *StarTransactionTypePremiumPurchase              `json:"starTransactionTypePremiumPurchase,omitempty"`
	StarTransactionTypePublicPostSearch             *StarTransactionTypePublicPostSearch             `json:"starTransactionTypePublicPostSearch,omitempty"`
	StarTransactionTypeSuggestedPostPaymentReceive  *StarTransactionTypeSuggestedPostPaymentReceive  `json:"starTransactionTypeSuggestedPostPaymentReceive,omitempty"`
	StarTransactionTypeSuggestedPostPaymentSend     *StarTransactionTypeSuggestedPostPaymentSend     `json:"starTransactionTypeSuggestedPostPaymentSend,omitempty"`
	StarTransactionTypeTelegramAdsWithdrawal        *StarTransactionTypeTelegramAdsWithdrawal        `json:"starTransactionTypeTelegramAdsWithdrawal,omitempty"`
	StarTransactionTypeTelegramApiUsage             *StarTransactionTypeTelegramApiUsage             `json:"starTransactionTypeTelegramApiUsage,omitempty"`
	StarTransactionTypeUnsupported                  *StarTransactionTypeUnsupported                  `json:"starTransactionTypeUnsupported,omitempty"`
	StarTransactionTypeUpgradedGiftPurchase         *StarTransactionTypeUpgradedGiftPurchase         `json:"starTransactionTypeUpgradedGiftPurchase,omitempty"`
	StarTransactionTypeUpgradedGiftSale             *StarTransactionTypeUpgradedGiftSale             `json:"starTransactionTypeUpgradedGiftSale,omitempty"`
	StarTransactionTypeUserDeposit                  *StarTransactionTypeUserDeposit                  `json:"starTransactionTypeUserDeposit,omitempty"`
}

func (t *StarTransactionType) Type() string {
	return t.TypeStr
}

func (t *StarTransactionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StarTransactionType) GetExtra() string {
	return ""
}

func (t *StarTransactionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "starTransactionTypeAffiliateProgramCommission":
		t.StarTransactionTypeAffiliateProgramCommission = new(StarTransactionTypeAffiliateProgramCommission)
		return json.Unmarshal(b, t.StarTransactionTypeAffiliateProgramCommission)
	case "starTransactionTypeAppStoreDeposit":
		t.StarTransactionTypeAppStoreDeposit = new(StarTransactionTypeAppStoreDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeAppStoreDeposit)
	case "starTransactionTypeBotInvoicePurchase":
		t.StarTransactionTypeBotInvoicePurchase = new(StarTransactionTypeBotInvoicePurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotInvoicePurchase)
	case "starTransactionTypeBotInvoiceSale":
		t.StarTransactionTypeBotInvoiceSale = new(StarTransactionTypeBotInvoiceSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotInvoiceSale)
	case "starTransactionTypeBotPaidMediaPurchase":
		t.StarTransactionTypeBotPaidMediaPurchase = new(StarTransactionTypeBotPaidMediaPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotPaidMediaPurchase)
	case "starTransactionTypeBotPaidMediaSale":
		t.StarTransactionTypeBotPaidMediaSale = new(StarTransactionTypeBotPaidMediaSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotPaidMediaSale)
	case "starTransactionTypeBotSubscriptionPurchase":
		t.StarTransactionTypeBotSubscriptionPurchase = new(StarTransactionTypeBotSubscriptionPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotSubscriptionPurchase)
	case "starTransactionTypeBotSubscriptionSale":
		t.StarTransactionTypeBotSubscriptionSale = new(StarTransactionTypeBotSubscriptionSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotSubscriptionSale)
	case "starTransactionTypeBusinessBotTransferReceive":
		t.StarTransactionTypeBusinessBotTransferReceive = new(StarTransactionTypeBusinessBotTransferReceive)
		return json.Unmarshal(b, t.StarTransactionTypeBusinessBotTransferReceive)
	case "starTransactionTypeBusinessBotTransferSend":
		t.StarTransactionTypeBusinessBotTransferSend = new(StarTransactionTypeBusinessBotTransferSend)
		return json.Unmarshal(b, t.StarTransactionTypeBusinessBotTransferSend)
	case "starTransactionTypeChannelPaidMediaPurchase":
		t.StarTransactionTypeChannelPaidMediaPurchase = new(StarTransactionTypeChannelPaidMediaPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidMediaPurchase)
	case "starTransactionTypeChannelPaidMediaSale":
		t.StarTransactionTypeChannelPaidMediaSale = new(StarTransactionTypeChannelPaidMediaSale)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidMediaSale)
	case "starTransactionTypeChannelPaidReactionReceive":
		t.StarTransactionTypeChannelPaidReactionReceive = new(StarTransactionTypeChannelPaidReactionReceive)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidReactionReceive)
	case "starTransactionTypeChannelPaidReactionSend":
		t.StarTransactionTypeChannelPaidReactionSend = new(StarTransactionTypeChannelPaidReactionSend)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidReactionSend)
	case "starTransactionTypeChannelSubscriptionPurchase":
		t.StarTransactionTypeChannelSubscriptionPurchase = new(StarTransactionTypeChannelSubscriptionPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeChannelSubscriptionPurchase)
	case "starTransactionTypeChannelSubscriptionSale":
		t.StarTransactionTypeChannelSubscriptionSale = new(StarTransactionTypeChannelSubscriptionSale)
		return json.Unmarshal(b, t.StarTransactionTypeChannelSubscriptionSale)
	case "starTransactionTypeFragmentDeposit":
		t.StarTransactionTypeFragmentDeposit = new(StarTransactionTypeFragmentDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeFragmentDeposit)
	case "starTransactionTypeFragmentWithdrawal":
		t.StarTransactionTypeFragmentWithdrawal = new(StarTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.StarTransactionTypeFragmentWithdrawal)
	case "starTransactionTypeGiftAuctionBid":
		t.StarTransactionTypeGiftAuctionBid = new(StarTransactionTypeGiftAuctionBid)
		return json.Unmarshal(b, t.StarTransactionTypeGiftAuctionBid)
	case "starTransactionTypeGiftOriginalDetailsDrop":
		t.StarTransactionTypeGiftOriginalDetailsDrop = new(StarTransactionTypeGiftOriginalDetailsDrop)
		return json.Unmarshal(b, t.StarTransactionTypeGiftOriginalDetailsDrop)
	case "starTransactionTypeGiftPurchase":
		t.StarTransactionTypeGiftPurchase = new(StarTransactionTypeGiftPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeGiftPurchase)
	case "starTransactionTypeGiftPurchaseOffer":
		t.StarTransactionTypeGiftPurchaseOffer = new(StarTransactionTypeGiftPurchaseOffer)
		return json.Unmarshal(b, t.StarTransactionTypeGiftPurchaseOffer)
	case "starTransactionTypeGiftSale":
		t.StarTransactionTypeGiftSale = new(StarTransactionTypeGiftSale)
		return json.Unmarshal(b, t.StarTransactionTypeGiftSale)
	case "starTransactionTypeGiftTransfer":
		t.StarTransactionTypeGiftTransfer = new(StarTransactionTypeGiftTransfer)
		return json.Unmarshal(b, t.StarTransactionTypeGiftTransfer)
	case "starTransactionTypeGiftUpgrade":
		t.StarTransactionTypeGiftUpgrade = new(StarTransactionTypeGiftUpgrade)
		return json.Unmarshal(b, t.StarTransactionTypeGiftUpgrade)
	case "starTransactionTypeGiftUpgradePurchase":
		t.StarTransactionTypeGiftUpgradePurchase = new(StarTransactionTypeGiftUpgradePurchase)
		return json.Unmarshal(b, t.StarTransactionTypeGiftUpgradePurchase)
	case "starTransactionTypeGiveawayDeposit":
		t.StarTransactionTypeGiveawayDeposit = new(StarTransactionTypeGiveawayDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeGiveawayDeposit)
	case "starTransactionTypeGooglePlayDeposit":
		t.StarTransactionTypeGooglePlayDeposit = new(StarTransactionTypeGooglePlayDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeGooglePlayDeposit)
	case "starTransactionTypePaidGroupCallMessageReceive":
		t.StarTransactionTypePaidGroupCallMessageReceive = new(StarTransactionTypePaidGroupCallMessageReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallMessageReceive)
	case "starTransactionTypePaidGroupCallMessageSend":
		t.StarTransactionTypePaidGroupCallMessageSend = new(StarTransactionTypePaidGroupCallMessageSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallMessageSend)
	case "starTransactionTypePaidGroupCallReactionReceive":
		t.StarTransactionTypePaidGroupCallReactionReceive = new(StarTransactionTypePaidGroupCallReactionReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallReactionReceive)
	case "starTransactionTypePaidGroupCallReactionSend":
		t.StarTransactionTypePaidGroupCallReactionSend = new(StarTransactionTypePaidGroupCallReactionSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallReactionSend)
	case "starTransactionTypePaidMessageReceive":
		t.StarTransactionTypePaidMessageReceive = new(StarTransactionTypePaidMessageReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidMessageReceive)
	case "starTransactionTypePaidMessageSend":
		t.StarTransactionTypePaidMessageSend = new(StarTransactionTypePaidMessageSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidMessageSend)
	case "starTransactionTypePremiumBotDeposit":
		t.StarTransactionTypePremiumBotDeposit = new(StarTransactionTypePremiumBotDeposit)
		return json.Unmarshal(b, t.StarTransactionTypePremiumBotDeposit)
	case "starTransactionTypePremiumPurchase":
		t.StarTransactionTypePremiumPurchase = new(StarTransactionTypePremiumPurchase)
		return json.Unmarshal(b, t.StarTransactionTypePremiumPurchase)
	case "starTransactionTypePublicPostSearch":
		t.StarTransactionTypePublicPostSearch = new(StarTransactionTypePublicPostSearch)
		return json.Unmarshal(b, t.StarTransactionTypePublicPostSearch)
	case "starTransactionTypeSuggestedPostPaymentReceive":
		t.StarTransactionTypeSuggestedPostPaymentReceive = new(StarTransactionTypeSuggestedPostPaymentReceive)
		return json.Unmarshal(b, t.StarTransactionTypeSuggestedPostPaymentReceive)
	case "starTransactionTypeSuggestedPostPaymentSend":
		t.StarTransactionTypeSuggestedPostPaymentSend = new(StarTransactionTypeSuggestedPostPaymentSend)
		return json.Unmarshal(b, t.StarTransactionTypeSuggestedPostPaymentSend)
	case "starTransactionTypeTelegramAdsWithdrawal":
		t.StarTransactionTypeTelegramAdsWithdrawal = new(StarTransactionTypeTelegramAdsWithdrawal)
		return json.Unmarshal(b, t.StarTransactionTypeTelegramAdsWithdrawal)
	case "starTransactionTypeTelegramApiUsage":
		t.StarTransactionTypeTelegramApiUsage = new(StarTransactionTypeTelegramApiUsage)
		return json.Unmarshal(b, t.StarTransactionTypeTelegramApiUsage)
	case "starTransactionTypeUnsupported":
		t.StarTransactionTypeUnsupported = new(StarTransactionTypeUnsupported)
		return json.Unmarshal(b, t.StarTransactionTypeUnsupported)
	case "starTransactionTypeUpgradedGiftPurchase":
		t.StarTransactionTypeUpgradedGiftPurchase = new(StarTransactionTypeUpgradedGiftPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeUpgradedGiftPurchase)
	case "starTransactionTypeUpgradedGiftSale":
		t.StarTransactionTypeUpgradedGiftSale = new(StarTransactionTypeUpgradedGiftSale)
		return json.Unmarshal(b, t.StarTransactionTypeUpgradedGiftSale)
	case "starTransactionTypeUserDeposit":
		t.StarTransactionTypeUserDeposit = new(StarTransactionTypeUserDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeUserDeposit)
	}
	return nil
}

func (t *StarTransactionType) MarshalJSON() ([]byte, error) {
	if t.StarTransactionTypeAffiliateProgramCommission != nil {
		return json.Marshal(t.StarTransactionTypeAffiliateProgramCommission)
	}
	if t.StarTransactionTypeAppStoreDeposit != nil {
		return json.Marshal(t.StarTransactionTypeAppStoreDeposit)
	}
	if t.StarTransactionTypeBotInvoicePurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotInvoicePurchase)
	}
	if t.StarTransactionTypeBotInvoiceSale != nil {
		return json.Marshal(t.StarTransactionTypeBotInvoiceSale)
	}
	if t.StarTransactionTypeBotPaidMediaPurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotPaidMediaPurchase)
	}
	if t.StarTransactionTypeBotPaidMediaSale != nil {
		return json.Marshal(t.StarTransactionTypeBotPaidMediaSale)
	}
	if t.StarTransactionTypeBotSubscriptionPurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotSubscriptionPurchase)
	}
	if t.StarTransactionTypeBotSubscriptionSale != nil {
		return json.Marshal(t.StarTransactionTypeBotSubscriptionSale)
	}
	if t.StarTransactionTypeBusinessBotTransferReceive != nil {
		return json.Marshal(t.StarTransactionTypeBusinessBotTransferReceive)
	}
	if t.StarTransactionTypeBusinessBotTransferSend != nil {
		return json.Marshal(t.StarTransactionTypeBusinessBotTransferSend)
	}
	if t.StarTransactionTypeChannelPaidMediaPurchase != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidMediaPurchase)
	}
	if t.StarTransactionTypeChannelPaidMediaSale != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidMediaSale)
	}
	if t.StarTransactionTypeChannelPaidReactionReceive != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidReactionReceive)
	}
	if t.StarTransactionTypeChannelPaidReactionSend != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidReactionSend)
	}
	if t.StarTransactionTypeChannelSubscriptionPurchase != nil {
		return json.Marshal(t.StarTransactionTypeChannelSubscriptionPurchase)
	}
	if t.StarTransactionTypeChannelSubscriptionSale != nil {
		return json.Marshal(t.StarTransactionTypeChannelSubscriptionSale)
	}
	if t.StarTransactionTypeFragmentDeposit != nil {
		return json.Marshal(t.StarTransactionTypeFragmentDeposit)
	}
	if t.StarTransactionTypeFragmentWithdrawal != nil {
		return json.Marshal(t.StarTransactionTypeFragmentWithdrawal)
	}
	if t.StarTransactionTypeGiftAuctionBid != nil {
		return json.Marshal(t.StarTransactionTypeGiftAuctionBid)
	}
	if t.StarTransactionTypeGiftOriginalDetailsDrop != nil {
		return json.Marshal(t.StarTransactionTypeGiftOriginalDetailsDrop)
	}
	if t.StarTransactionTypeGiftPurchase != nil {
		return json.Marshal(t.StarTransactionTypeGiftPurchase)
	}
	if t.StarTransactionTypeGiftPurchaseOffer != nil {
		return json.Marshal(t.StarTransactionTypeGiftPurchaseOffer)
	}
	if t.StarTransactionTypeGiftSale != nil {
		return json.Marshal(t.StarTransactionTypeGiftSale)
	}
	if t.StarTransactionTypeGiftTransfer != nil {
		return json.Marshal(t.StarTransactionTypeGiftTransfer)
	}
	if t.StarTransactionTypeGiftUpgrade != nil {
		return json.Marshal(t.StarTransactionTypeGiftUpgrade)
	}
	if t.StarTransactionTypeGiftUpgradePurchase != nil {
		return json.Marshal(t.StarTransactionTypeGiftUpgradePurchase)
	}
	if t.StarTransactionTypeGiveawayDeposit != nil {
		return json.Marshal(t.StarTransactionTypeGiveawayDeposit)
	}
	if t.StarTransactionTypeGooglePlayDeposit != nil {
		return json.Marshal(t.StarTransactionTypeGooglePlayDeposit)
	}
	if t.StarTransactionTypePaidGroupCallMessageReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallMessageReceive)
	}
	if t.StarTransactionTypePaidGroupCallMessageSend != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallMessageSend)
	}
	if t.StarTransactionTypePaidGroupCallReactionReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallReactionReceive)
	}
	if t.StarTransactionTypePaidGroupCallReactionSend != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallReactionSend)
	}
	if t.StarTransactionTypePaidMessageReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidMessageReceive)
	}
	if t.StarTransactionTypePaidMessageSend != nil {
		return json.Marshal(t.StarTransactionTypePaidMessageSend)
	}
	if t.StarTransactionTypePremiumBotDeposit != nil {
		return json.Marshal(t.StarTransactionTypePremiumBotDeposit)
	}
	if t.StarTransactionTypePremiumPurchase != nil {
		return json.Marshal(t.StarTransactionTypePremiumPurchase)
	}
	if t.StarTransactionTypePublicPostSearch != nil {
		return json.Marshal(t.StarTransactionTypePublicPostSearch)
	}
	if t.StarTransactionTypeSuggestedPostPaymentReceive != nil {
		return json.Marshal(t.StarTransactionTypeSuggestedPostPaymentReceive)
	}
	if t.StarTransactionTypeSuggestedPostPaymentSend != nil {
		return json.Marshal(t.StarTransactionTypeSuggestedPostPaymentSend)
	}
	if t.StarTransactionTypeTelegramAdsWithdrawal != nil {
		return json.Marshal(t.StarTransactionTypeTelegramAdsWithdrawal)
	}
	if t.StarTransactionTypeTelegramApiUsage != nil {
		return json.Marshal(t.StarTransactionTypeTelegramApiUsage)
	}
	if t.StarTransactionTypeUnsupported != nil {
		return json.Marshal(t.StarTransactionTypeUnsupported)
	}
	if t.StarTransactionTypeUpgradedGiftPurchase != nil {
		return json.Marshal(t.StarTransactionTypeUpgradedGiftPurchase)
	}
	if t.StarTransactionTypeUpgradedGiftSale != nil {
		return json.Marshal(t.StarTransactionTypeUpgradedGiftSale)
	}
	if t.StarTransactionTypeUserDeposit != nil {
		return json.Marshal(t.StarTransactionTypeUserDeposit)
	}
	return nil, fmt.Errorf("no value set for StarTransactionType")
}

// StatisticalGraph Describes a statistical graph
type StatisticalGraph struct {
	TypeStr               string                 `json:"@type"`
	StatisticalGraphAsync *StatisticalGraphAsync `json:"statisticalGraphAsync,omitempty"`
	StatisticalGraphData  *StatisticalGraphData  `json:"statisticalGraphData,omitempty"`
	StatisticalGraphError *StatisticalGraphError `json:"statisticalGraphError,omitempty"`
}

func (t *StatisticalGraph) Type() string {
	return t.TypeStr
}

func (t *StatisticalGraph) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StatisticalGraph) GetExtra() string {
	return ""
}

func (t *StatisticalGraph) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "statisticalGraphAsync":
		t.StatisticalGraphAsync = new(StatisticalGraphAsync)
		return json.Unmarshal(b, t.StatisticalGraphAsync)
	case "statisticalGraphData":
		t.StatisticalGraphData = new(StatisticalGraphData)
		return json.Unmarshal(b, t.StatisticalGraphData)
	case "statisticalGraphError":
		t.StatisticalGraphError = new(StatisticalGraphError)
		return json.Unmarshal(b, t.StatisticalGraphError)
	}
	return nil
}

func (t *StatisticalGraph) MarshalJSON() ([]byte, error) {
	if t.StatisticalGraphAsync != nil {
		return json.Marshal(t.StatisticalGraphAsync)
	}
	if t.StatisticalGraphData != nil {
		return json.Marshal(t.StatisticalGraphData)
	}
	if t.StatisticalGraphError != nil {
		return json.Marshal(t.StatisticalGraphError)
	}
	return nil, fmt.Errorf("no value set for StatisticalGraph")
}

// StickerFormat Describes format of a sticker
type StickerFormat struct {
	TypeStr           string             `json:"@type"`
	StickerFormatTgs  *StickerFormatTgs  `json:"stickerFormatTgs,omitempty"`
	StickerFormatWebm *StickerFormatWebm `json:"stickerFormatWebm,omitempty"`
	StickerFormatWebp *StickerFormatWebp `json:"stickerFormatWebp,omitempty"`
}

func (t *StickerFormat) Type() string {
	return t.TypeStr
}

func (t *StickerFormat) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StickerFormat) GetExtra() string {
	return ""
}

func (t *StickerFormat) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "stickerFormatTgs":
		t.StickerFormatTgs = new(StickerFormatTgs)
		return json.Unmarshal(b, t.StickerFormatTgs)
	case "stickerFormatWebm":
		t.StickerFormatWebm = new(StickerFormatWebm)
		return json.Unmarshal(b, t.StickerFormatWebm)
	case "stickerFormatWebp":
		t.StickerFormatWebp = new(StickerFormatWebp)
		return json.Unmarshal(b, t.StickerFormatWebp)
	}
	return nil
}

func (t *StickerFormat) MarshalJSON() ([]byte, error) {
	if t.StickerFormatTgs != nil {
		return json.Marshal(t.StickerFormatTgs)
	}
	if t.StickerFormatWebm != nil {
		return json.Marshal(t.StickerFormatWebm)
	}
	if t.StickerFormatWebp != nil {
		return json.Marshal(t.StickerFormatWebp)
	}
	return nil, fmt.Errorf("no value set for StickerFormat")
}

// StickerFullType Contains full information about sticker type
type StickerFullType struct {
	TypeStr                    string                      `json:"@type"`
	StickerFullTypeCustomEmoji *StickerFullTypeCustomEmoji `json:"stickerFullTypeCustomEmoji,omitempty"`
	StickerFullTypeMask        *StickerFullTypeMask        `json:"stickerFullTypeMask,omitempty"`
	StickerFullTypeRegular     *StickerFullTypeRegular     `json:"stickerFullTypeRegular,omitempty"`
}

func (t *StickerFullType) Type() string {
	return t.TypeStr
}

func (t *StickerFullType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StickerFullType) GetExtra() string {
	return ""
}

func (t *StickerFullType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "stickerFullTypeCustomEmoji":
		t.StickerFullTypeCustomEmoji = new(StickerFullTypeCustomEmoji)
		return json.Unmarshal(b, t.StickerFullTypeCustomEmoji)
	case "stickerFullTypeMask":
		t.StickerFullTypeMask = new(StickerFullTypeMask)
		return json.Unmarshal(b, t.StickerFullTypeMask)
	case "stickerFullTypeRegular":
		t.StickerFullTypeRegular = new(StickerFullTypeRegular)
		return json.Unmarshal(b, t.StickerFullTypeRegular)
	}
	return nil
}

func (t *StickerFullType) MarshalJSON() ([]byte, error) {
	if t.StickerFullTypeCustomEmoji != nil {
		return json.Marshal(t.StickerFullTypeCustomEmoji)
	}
	if t.StickerFullTypeMask != nil {
		return json.Marshal(t.StickerFullTypeMask)
	}
	if t.StickerFullTypeRegular != nil {
		return json.Marshal(t.StickerFullTypeRegular)
	}
	return nil, fmt.Errorf("no value set for StickerFullType")
}

// StickerType Describes type of sticker
type StickerType struct {
	TypeStr                string                  `json:"@type"`
	StickerTypeCustomEmoji *StickerTypeCustomEmoji `json:"stickerTypeCustomEmoji,omitempty"`
	StickerTypeMask        *StickerTypeMask        `json:"stickerTypeMask,omitempty"`
	StickerTypeRegular     *StickerTypeRegular     `json:"stickerTypeRegular,omitempty"`
}

func (t *StickerType) Type() string {
	return t.TypeStr
}

func (t *StickerType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StickerType) GetExtra() string {
	return ""
}

func (t *StickerType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "stickerTypeCustomEmoji":
		t.StickerTypeCustomEmoji = new(StickerTypeCustomEmoji)
		return json.Unmarshal(b, t.StickerTypeCustomEmoji)
	case "stickerTypeMask":
		t.StickerTypeMask = new(StickerTypeMask)
		return json.Unmarshal(b, t.StickerTypeMask)
	case "stickerTypeRegular":
		t.StickerTypeRegular = new(StickerTypeRegular)
		return json.Unmarshal(b, t.StickerTypeRegular)
	}
	return nil
}

func (t *StickerType) MarshalJSON() ([]byte, error) {
	if t.StickerTypeCustomEmoji != nil {
		return json.Marshal(t.StickerTypeCustomEmoji)
	}
	if t.StickerTypeMask != nil {
		return json.Marshal(t.StickerTypeMask)
	}
	if t.StickerTypeRegular != nil {
		return json.Marshal(t.StickerTypeRegular)
	}
	return nil, fmt.Errorf("no value set for StickerType")
}

// StorePaymentPurpose Describes a purpose of an in-store payment
type StorePaymentPurpose struct {
	TypeStr                                string                                  `json:"@type"`
	StorePaymentPurposeGiftedStars         *StorePaymentPurposeGiftedStars         `json:"storePaymentPurposeGiftedStars,omitempty"`
	StorePaymentPurposePremiumGift         *StorePaymentPurposePremiumGift         `json:"storePaymentPurposePremiumGift,omitempty"`
	StorePaymentPurposePremiumGiftCodes    *StorePaymentPurposePremiumGiftCodes    `json:"storePaymentPurposePremiumGiftCodes,omitempty"`
	StorePaymentPurposePremiumGiveaway     *StorePaymentPurposePremiumGiveaway     `json:"storePaymentPurposePremiumGiveaway,omitempty"`
	StorePaymentPurposePremiumSubscription *StorePaymentPurposePremiumSubscription `json:"storePaymentPurposePremiumSubscription,omitempty"`
	StorePaymentPurposeStarGiveaway        *StorePaymentPurposeStarGiveaway        `json:"storePaymentPurposeStarGiveaway,omitempty"`
	StorePaymentPurposeStars               *StorePaymentPurposeStars               `json:"storePaymentPurposeStars,omitempty"`
}

func (t *StorePaymentPurpose) Type() string {
	return t.TypeStr
}

func (t *StorePaymentPurpose) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StorePaymentPurpose) GetExtra() string {
	return ""
}

func (t *StorePaymentPurpose) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storePaymentPurposeGiftedStars":
		t.StorePaymentPurposeGiftedStars = new(StorePaymentPurposeGiftedStars)
		return json.Unmarshal(b, t.StorePaymentPurposeGiftedStars)
	case "storePaymentPurposePremiumGift":
		t.StorePaymentPurposePremiumGift = new(StorePaymentPurposePremiumGift)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGift)
	case "storePaymentPurposePremiumGiftCodes":
		t.StorePaymentPurposePremiumGiftCodes = new(StorePaymentPurposePremiumGiftCodes)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGiftCodes)
	case "storePaymentPurposePremiumGiveaway":
		t.StorePaymentPurposePremiumGiveaway = new(StorePaymentPurposePremiumGiveaway)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGiveaway)
	case "storePaymentPurposePremiumSubscription":
		t.StorePaymentPurposePremiumSubscription = new(StorePaymentPurposePremiumSubscription)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumSubscription)
	case "storePaymentPurposeStarGiveaway":
		t.StorePaymentPurposeStarGiveaway = new(StorePaymentPurposeStarGiveaway)
		return json.Unmarshal(b, t.StorePaymentPurposeStarGiveaway)
	case "storePaymentPurposeStars":
		t.StorePaymentPurposeStars = new(StorePaymentPurposeStars)
		return json.Unmarshal(b, t.StorePaymentPurposeStars)
	}
	return nil
}

func (t *StorePaymentPurpose) MarshalJSON() ([]byte, error) {
	if t.StorePaymentPurposeGiftedStars != nil {
		return json.Marshal(t.StorePaymentPurposeGiftedStars)
	}
	if t.StorePaymentPurposePremiumGift != nil {
		return json.Marshal(t.StorePaymentPurposePremiumGift)
	}
	if t.StorePaymentPurposePremiumGiftCodes != nil {
		return json.Marshal(t.StorePaymentPurposePremiumGiftCodes)
	}
	if t.StorePaymentPurposePremiumGiveaway != nil {
		return json.Marshal(t.StorePaymentPurposePremiumGiveaway)
	}
	if t.StorePaymentPurposePremiumSubscription != nil {
		return json.Marshal(t.StorePaymentPurposePremiumSubscription)
	}
	if t.StorePaymentPurposeStarGiveaway != nil {
		return json.Marshal(t.StorePaymentPurposeStarGiveaway)
	}
	if t.StorePaymentPurposeStars != nil {
		return json.Marshal(t.StorePaymentPurposeStars)
	}
	return nil, fmt.Errorf("no value set for StorePaymentPurpose")
}

// StoreTransaction Describes an in-store transaction
type StoreTransaction struct {
	TypeStr                    string                      `json:"@type"`
	StoreTransactionAppStore   *StoreTransactionAppStore   `json:"storeTransactionAppStore,omitempty"`
	StoreTransactionGooglePlay *StoreTransactionGooglePlay `json:"storeTransactionGooglePlay,omitempty"`
}

func (t *StoreTransaction) Type() string {
	return t.TypeStr
}

func (t *StoreTransaction) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoreTransaction) GetExtra() string {
	return ""
}

func (t *StoreTransaction) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storeTransactionAppStore":
		t.StoreTransactionAppStore = new(StoreTransactionAppStore)
		return json.Unmarshal(b, t.StoreTransactionAppStore)
	case "storeTransactionGooglePlay":
		t.StoreTransactionGooglePlay = new(StoreTransactionGooglePlay)
		return json.Unmarshal(b, t.StoreTransactionGooglePlay)
	}
	return nil
}

func (t *StoreTransaction) MarshalJSON() ([]byte, error) {
	if t.StoreTransactionAppStore != nil {
		return json.Marshal(t.StoreTransactionAppStore)
	}
	if t.StoreTransactionGooglePlay != nil {
		return json.Marshal(t.StoreTransactionGooglePlay)
	}
	return nil, fmt.Errorf("no value set for StoreTransaction")
}

// StoryAreaType Describes type of clickable area on a story media
type StoryAreaType struct {
	TypeStr                        string                          `json:"@type"`
	StoryAreaTypeLink              *StoryAreaTypeLink              `json:"storyAreaTypeLink,omitempty"`
	StoryAreaTypeLocation          *StoryAreaTypeLocation          `json:"storyAreaTypeLocation,omitempty"`
	StoryAreaTypeMessage           *StoryAreaTypeMessage           `json:"storyAreaTypeMessage,omitempty"`
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction `json:"storyAreaTypeSuggestedReaction,omitempty"`
	StoryAreaTypeUpgradedGift      *StoryAreaTypeUpgradedGift      `json:"storyAreaTypeUpgradedGift,omitempty"`
	StoryAreaTypeVenue             *StoryAreaTypeVenue             `json:"storyAreaTypeVenue,omitempty"`
	StoryAreaTypeWeather           *StoryAreaTypeWeather           `json:"storyAreaTypeWeather,omitempty"`
}

func (t *StoryAreaType) Type() string {
	return t.TypeStr
}

func (t *StoryAreaType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryAreaType) GetExtra() string {
	return ""
}

func (t *StoryAreaType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyAreaTypeLink":
		t.StoryAreaTypeLink = new(StoryAreaTypeLink)
		return json.Unmarshal(b, t.StoryAreaTypeLink)
	case "storyAreaTypeLocation":
		t.StoryAreaTypeLocation = new(StoryAreaTypeLocation)
		return json.Unmarshal(b, t.StoryAreaTypeLocation)
	case "storyAreaTypeMessage":
		t.StoryAreaTypeMessage = new(StoryAreaTypeMessage)
		return json.Unmarshal(b, t.StoryAreaTypeMessage)
	case "storyAreaTypeSuggestedReaction":
		t.StoryAreaTypeSuggestedReaction = new(StoryAreaTypeSuggestedReaction)
		return json.Unmarshal(b, t.StoryAreaTypeSuggestedReaction)
	case "storyAreaTypeUpgradedGift":
		t.StoryAreaTypeUpgradedGift = new(StoryAreaTypeUpgradedGift)
		return json.Unmarshal(b, t.StoryAreaTypeUpgradedGift)
	case "storyAreaTypeVenue":
		t.StoryAreaTypeVenue = new(StoryAreaTypeVenue)
		return json.Unmarshal(b, t.StoryAreaTypeVenue)
	case "storyAreaTypeWeather":
		t.StoryAreaTypeWeather = new(StoryAreaTypeWeather)
		return json.Unmarshal(b, t.StoryAreaTypeWeather)
	}
	return nil
}

func (t *StoryAreaType) MarshalJSON() ([]byte, error) {
	if t.StoryAreaTypeLink != nil {
		return json.Marshal(t.StoryAreaTypeLink)
	}
	if t.StoryAreaTypeLocation != nil {
		return json.Marshal(t.StoryAreaTypeLocation)
	}
	if t.StoryAreaTypeMessage != nil {
		return json.Marshal(t.StoryAreaTypeMessage)
	}
	if t.StoryAreaTypeSuggestedReaction != nil {
		return json.Marshal(t.StoryAreaTypeSuggestedReaction)
	}
	if t.StoryAreaTypeUpgradedGift != nil {
		return json.Marshal(t.StoryAreaTypeUpgradedGift)
	}
	if t.StoryAreaTypeVenue != nil {
		return json.Marshal(t.StoryAreaTypeVenue)
	}
	if t.StoryAreaTypeWeather != nil {
		return json.Marshal(t.StoryAreaTypeWeather)
	}
	return nil, fmt.Errorf("no value set for StoryAreaType")
}

// StoryContent Contains the content of a story
type StoryContent struct {
	TypeStr                 string                   `json:"@type"`
	StoryContentLive        *StoryContentLive        `json:"storyContentLive,omitempty"`
	StoryContentPhoto       *StoryContentPhoto       `json:"storyContentPhoto,omitempty"`
	StoryContentUnsupported *StoryContentUnsupported `json:"storyContentUnsupported,omitempty"`
	StoryContentVideo       *StoryContentVideo       `json:"storyContentVideo,omitempty"`
}

func (t *StoryContent) Type() string {
	return t.TypeStr
}

func (t *StoryContent) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryContent) GetExtra() string {
	return ""
}

func (t *StoryContent) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyContentLive":
		t.StoryContentLive = new(StoryContentLive)
		return json.Unmarshal(b, t.StoryContentLive)
	case "storyContentPhoto":
		t.StoryContentPhoto = new(StoryContentPhoto)
		return json.Unmarshal(b, t.StoryContentPhoto)
	case "storyContentUnsupported":
		t.StoryContentUnsupported = new(StoryContentUnsupported)
		return json.Unmarshal(b, t.StoryContentUnsupported)
	case "storyContentVideo":
		t.StoryContentVideo = new(StoryContentVideo)
		return json.Unmarshal(b, t.StoryContentVideo)
	}
	return nil
}

func (t *StoryContent) MarshalJSON() ([]byte, error) {
	if t.StoryContentLive != nil {
		return json.Marshal(t.StoryContentLive)
	}
	if t.StoryContentPhoto != nil {
		return json.Marshal(t.StoryContentPhoto)
	}
	if t.StoryContentUnsupported != nil {
		return json.Marshal(t.StoryContentUnsupported)
	}
	if t.StoryContentVideo != nil {
		return json.Marshal(t.StoryContentVideo)
	}
	return nil, fmt.Errorf("no value set for StoryContent")
}

// StoryInteractionType Describes type of interaction with a story
type StoryInteractionType struct {
	TypeStr                     string                       `json:"@type"`
	StoryInteractionTypeForward *StoryInteractionTypeForward `json:"storyInteractionTypeForward,omitempty"`
	StoryInteractionTypeRepost  *StoryInteractionTypeRepost  `json:"storyInteractionTypeRepost,omitempty"`
	StoryInteractionTypeView    *StoryInteractionTypeView    `json:"storyInteractionTypeView,omitempty"`
}

func (t *StoryInteractionType) Type() string {
	return t.TypeStr
}

func (t *StoryInteractionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryInteractionType) GetExtra() string {
	return ""
}

func (t *StoryInteractionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyInteractionTypeForward":
		t.StoryInteractionTypeForward = new(StoryInteractionTypeForward)
		return json.Unmarshal(b, t.StoryInteractionTypeForward)
	case "storyInteractionTypeRepost":
		t.StoryInteractionTypeRepost = new(StoryInteractionTypeRepost)
		return json.Unmarshal(b, t.StoryInteractionTypeRepost)
	case "storyInteractionTypeView":
		t.StoryInteractionTypeView = new(StoryInteractionTypeView)
		return json.Unmarshal(b, t.StoryInteractionTypeView)
	}
	return nil
}

func (t *StoryInteractionType) MarshalJSON() ([]byte, error) {
	if t.StoryInteractionTypeForward != nil {
		return json.Marshal(t.StoryInteractionTypeForward)
	}
	if t.StoryInteractionTypeRepost != nil {
		return json.Marshal(t.StoryInteractionTypeRepost)
	}
	if t.StoryInteractionTypeView != nil {
		return json.Marshal(t.StoryInteractionTypeView)
	}
	return nil, fmt.Errorf("no value set for StoryInteractionType")
}

// StoryList Describes a list of stories
type StoryList struct {
	TypeStr          string            `json:"@type"`
	StoryListArchive *StoryListArchive `json:"storyListArchive,omitempty"`
	StoryListMain    *StoryListMain    `json:"storyListMain,omitempty"`
}

func (t *StoryList) Type() string {
	return t.TypeStr
}

func (t *StoryList) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryList) GetExtra() string {
	return ""
}

func (t *StoryList) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyListArchive":
		t.StoryListArchive = new(StoryListArchive)
		return json.Unmarshal(b, t.StoryListArchive)
	case "storyListMain":
		t.StoryListMain = new(StoryListMain)
		return json.Unmarshal(b, t.StoryListMain)
	}
	return nil
}

func (t *StoryList) MarshalJSON() ([]byte, error) {
	if t.StoryListArchive != nil {
		return json.Marshal(t.StoryListArchive)
	}
	if t.StoryListMain != nil {
		return json.Marshal(t.StoryListMain)
	}
	return nil, fmt.Errorf("no value set for StoryList")
}

// StoryOrigin Contains information about the origin of a story that was reposted
type StoryOrigin struct {
	TypeStr                string                  `json:"@type"`
	StoryOriginHiddenUser  *StoryOriginHiddenUser  `json:"storyOriginHiddenUser,omitempty"`
	StoryOriginPublicStory *StoryOriginPublicStory `json:"storyOriginPublicStory,omitempty"`
}

func (t *StoryOrigin) Type() string {
	return t.TypeStr
}

func (t *StoryOrigin) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryOrigin) GetExtra() string {
	return ""
}

func (t *StoryOrigin) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyOriginHiddenUser":
		t.StoryOriginHiddenUser = new(StoryOriginHiddenUser)
		return json.Unmarshal(b, t.StoryOriginHiddenUser)
	case "storyOriginPublicStory":
		t.StoryOriginPublicStory = new(StoryOriginPublicStory)
		return json.Unmarshal(b, t.StoryOriginPublicStory)
	}
	return nil
}

func (t *StoryOrigin) MarshalJSON() ([]byte, error) {
	if t.StoryOriginHiddenUser != nil {
		return json.Marshal(t.StoryOriginHiddenUser)
	}
	if t.StoryOriginPublicStory != nil {
		return json.Marshal(t.StoryOriginPublicStory)
	}
	return nil, fmt.Errorf("no value set for StoryOrigin")
}

// StoryPrivacySettings Describes privacy settings of a story
type StoryPrivacySettings struct {
	TypeStr                           string                             `json:"@type"`
	StoryPrivacySettingsCloseFriends  *StoryPrivacySettingsCloseFriends  `json:"storyPrivacySettingsCloseFriends,omitempty"`
	StoryPrivacySettingsContacts      *StoryPrivacySettingsContacts      `json:"storyPrivacySettingsContacts,omitempty"`
	StoryPrivacySettingsEveryone      *StoryPrivacySettingsEveryone      `json:"storyPrivacySettingsEveryone,omitempty"`
	StoryPrivacySettingsSelectedUsers *StoryPrivacySettingsSelectedUsers `json:"storyPrivacySettingsSelectedUsers,omitempty"`
}

func (t *StoryPrivacySettings) Type() string {
	return t.TypeStr
}

func (t *StoryPrivacySettings) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *StoryPrivacySettings) GetExtra() string {
	return ""
}

func (t *StoryPrivacySettings) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "storyPrivacySettingsCloseFriends":
		t.StoryPrivacySettingsCloseFriends = new(StoryPrivacySettingsCloseFriends)
		return json.Unmarshal(b, t.StoryPrivacySettingsCloseFriends)
	case "storyPrivacySettingsContacts":
		t.StoryPrivacySettingsContacts = new(StoryPrivacySettingsContacts)
		return json.Unmarshal(b, t.StoryPrivacySettingsContacts)
	case "storyPrivacySettingsEveryone":
		t.StoryPrivacySettingsEveryone = new(StoryPrivacySettingsEveryone)
		return json.Unmarshal(b, t.StoryPrivacySettingsEveryone)
	case "storyPrivacySettingsSelectedUsers":
		t.StoryPrivacySettingsSelectedUsers = new(StoryPrivacySettingsSelectedUsers)
		return json.Unmarshal(b, t.StoryPrivacySettingsSelectedUsers)
	}
	return nil
}

func (t *StoryPrivacySettings) MarshalJSON() ([]byte, error) {
	if t.StoryPrivacySettingsCloseFriends != nil {
		return json.Marshal(t.StoryPrivacySettingsCloseFriends)
	}
	if t.StoryPrivacySettingsContacts != nil {
		return json.Marshal(t.StoryPrivacySettingsContacts)
	}
	if t.StoryPrivacySettingsEveryone != nil {
		return json.Marshal(t.StoryPrivacySettingsEveryone)
	}
	if t.StoryPrivacySettingsSelectedUsers != nil {
		return json.Marshal(t.StoryPrivacySettingsSelectedUsers)
	}
	return nil, fmt.Errorf("no value set for StoryPrivacySettings")
}

// SuggestedAction Describes an action suggested to the current user
type SuggestedAction struct {
	TypeStr                                     string                                       `json:"@type"`
	SuggestedActionAddLoginPasskey              *SuggestedActionAddLoginPasskey              `json:"suggestedActionAddLoginPasskey,omitempty"`
	SuggestedActionCheckPassword                *SuggestedActionCheckPassword                `json:"suggestedActionCheckPassword,omitempty"`
	SuggestedActionCheckPhoneNumber             *SuggestedActionCheckPhoneNumber             `json:"suggestedActionCheckPhoneNumber,omitempty"`
	SuggestedActionConvertToBroadcastGroup      *SuggestedActionConvertToBroadcastGroup      `json:"suggestedActionConvertToBroadcastGroup,omitempty"`
	SuggestedActionCustom                       *SuggestedActionCustom                       `json:"suggestedActionCustom,omitempty"`
	SuggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats `json:"suggestedActionEnableArchiveAndMuteNewChats,omitempty"`
	SuggestedActionExtendPremium                *SuggestedActionExtendPremium                `json:"suggestedActionExtendPremium,omitempty"`
	SuggestedActionExtendStarSubscriptions      *SuggestedActionExtendStarSubscriptions      `json:"suggestedActionExtendStarSubscriptions,omitempty"`
	SuggestedActionGiftPremiumForChristmas      *SuggestedActionGiftPremiumForChristmas      `json:"suggestedActionGiftPremiumForChristmas,omitempty"`
	SuggestedActionRestorePremium               *SuggestedActionRestorePremium               `json:"suggestedActionRestorePremium,omitempty"`
	SuggestedActionSetBirthdate                 *SuggestedActionSetBirthdate                 `json:"suggestedActionSetBirthdate,omitempty"`
	SuggestedActionSetLoginEmailAddress         *SuggestedActionSetLoginEmailAddress         `json:"suggestedActionSetLoginEmailAddress,omitempty"`
	SuggestedActionSetPassword                  *SuggestedActionSetPassword                  `json:"suggestedActionSetPassword,omitempty"`
	SuggestedActionSetProfilePhoto              *SuggestedActionSetProfilePhoto              `json:"suggestedActionSetProfilePhoto,omitempty"`
	SuggestedActionSubscribeToAnnualPremium     *SuggestedActionSubscribeToAnnualPremium     `json:"suggestedActionSubscribeToAnnualPremium,omitempty"`
	SuggestedActionUpgradePremium               *SuggestedActionUpgradePremium               `json:"suggestedActionUpgradePremium,omitempty"`
	SuggestedActionViewChecksHint               *SuggestedActionViewChecksHint               `json:"suggestedActionViewChecksHint,omitempty"`
}

func (t *SuggestedAction) Type() string {
	return t.TypeStr
}

func (t *SuggestedAction) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SuggestedAction) GetExtra() string {
	return ""
}

func (t *SuggestedAction) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "suggestedActionAddLoginPasskey":
		t.SuggestedActionAddLoginPasskey = new(SuggestedActionAddLoginPasskey)
		return json.Unmarshal(b, t.SuggestedActionAddLoginPasskey)
	case "suggestedActionCheckPassword":
		t.SuggestedActionCheckPassword = new(SuggestedActionCheckPassword)
		return json.Unmarshal(b, t.SuggestedActionCheckPassword)
	case "suggestedActionCheckPhoneNumber":
		t.SuggestedActionCheckPhoneNumber = new(SuggestedActionCheckPhoneNumber)
		return json.Unmarshal(b, t.SuggestedActionCheckPhoneNumber)
	case "suggestedActionConvertToBroadcastGroup":
		t.SuggestedActionConvertToBroadcastGroup = new(SuggestedActionConvertToBroadcastGroup)
		return json.Unmarshal(b, t.SuggestedActionConvertToBroadcastGroup)
	case "suggestedActionCustom":
		t.SuggestedActionCustom = new(SuggestedActionCustom)
		return json.Unmarshal(b, t.SuggestedActionCustom)
	case "suggestedActionEnableArchiveAndMuteNewChats":
		t.SuggestedActionEnableArchiveAndMuteNewChats = new(SuggestedActionEnableArchiveAndMuteNewChats)
		return json.Unmarshal(b, t.SuggestedActionEnableArchiveAndMuteNewChats)
	case "suggestedActionExtendPremium":
		t.SuggestedActionExtendPremium = new(SuggestedActionExtendPremium)
		return json.Unmarshal(b, t.SuggestedActionExtendPremium)
	case "suggestedActionExtendStarSubscriptions":
		t.SuggestedActionExtendStarSubscriptions = new(SuggestedActionExtendStarSubscriptions)
		return json.Unmarshal(b, t.SuggestedActionExtendStarSubscriptions)
	case "suggestedActionGiftPremiumForChristmas":
		t.SuggestedActionGiftPremiumForChristmas = new(SuggestedActionGiftPremiumForChristmas)
		return json.Unmarshal(b, t.SuggestedActionGiftPremiumForChristmas)
	case "suggestedActionRestorePremium":
		t.SuggestedActionRestorePremium = new(SuggestedActionRestorePremium)
		return json.Unmarshal(b, t.SuggestedActionRestorePremium)
	case "suggestedActionSetBirthdate":
		t.SuggestedActionSetBirthdate = new(SuggestedActionSetBirthdate)
		return json.Unmarshal(b, t.SuggestedActionSetBirthdate)
	case "suggestedActionSetLoginEmailAddress":
		t.SuggestedActionSetLoginEmailAddress = new(SuggestedActionSetLoginEmailAddress)
		return json.Unmarshal(b, t.SuggestedActionSetLoginEmailAddress)
	case "suggestedActionSetPassword":
		t.SuggestedActionSetPassword = new(SuggestedActionSetPassword)
		return json.Unmarshal(b, t.SuggestedActionSetPassword)
	case "suggestedActionSetProfilePhoto":
		t.SuggestedActionSetProfilePhoto = new(SuggestedActionSetProfilePhoto)
		return json.Unmarshal(b, t.SuggestedActionSetProfilePhoto)
	case "suggestedActionSubscribeToAnnualPremium":
		t.SuggestedActionSubscribeToAnnualPremium = new(SuggestedActionSubscribeToAnnualPremium)
		return json.Unmarshal(b, t.SuggestedActionSubscribeToAnnualPremium)
	case "suggestedActionUpgradePremium":
		t.SuggestedActionUpgradePremium = new(SuggestedActionUpgradePremium)
		return json.Unmarshal(b, t.SuggestedActionUpgradePremium)
	case "suggestedActionViewChecksHint":
		t.SuggestedActionViewChecksHint = new(SuggestedActionViewChecksHint)
		return json.Unmarshal(b, t.SuggestedActionViewChecksHint)
	}
	return nil
}

func (t *SuggestedAction) MarshalJSON() ([]byte, error) {
	if t.SuggestedActionAddLoginPasskey != nil {
		return json.Marshal(t.SuggestedActionAddLoginPasskey)
	}
	if t.SuggestedActionCheckPassword != nil {
		return json.Marshal(t.SuggestedActionCheckPassword)
	}
	if t.SuggestedActionCheckPhoneNumber != nil {
		return json.Marshal(t.SuggestedActionCheckPhoneNumber)
	}
	if t.SuggestedActionConvertToBroadcastGroup != nil {
		return json.Marshal(t.SuggestedActionConvertToBroadcastGroup)
	}
	if t.SuggestedActionCustom != nil {
		return json.Marshal(t.SuggestedActionCustom)
	}
	if t.SuggestedActionEnableArchiveAndMuteNewChats != nil {
		return json.Marshal(t.SuggestedActionEnableArchiveAndMuteNewChats)
	}
	if t.SuggestedActionExtendPremium != nil {
		return json.Marshal(t.SuggestedActionExtendPremium)
	}
	if t.SuggestedActionExtendStarSubscriptions != nil {
		return json.Marshal(t.SuggestedActionExtendStarSubscriptions)
	}
	if t.SuggestedActionGiftPremiumForChristmas != nil {
		return json.Marshal(t.SuggestedActionGiftPremiumForChristmas)
	}
	if t.SuggestedActionRestorePremium != nil {
		return json.Marshal(t.SuggestedActionRestorePremium)
	}
	if t.SuggestedActionSetBirthdate != nil {
		return json.Marshal(t.SuggestedActionSetBirthdate)
	}
	if t.SuggestedActionSetLoginEmailAddress != nil {
		return json.Marshal(t.SuggestedActionSetLoginEmailAddress)
	}
	if t.SuggestedActionSetPassword != nil {
		return json.Marshal(t.SuggestedActionSetPassword)
	}
	if t.SuggestedActionSetProfilePhoto != nil {
		return json.Marshal(t.SuggestedActionSetProfilePhoto)
	}
	if t.SuggestedActionSubscribeToAnnualPremium != nil {
		return json.Marshal(t.SuggestedActionSubscribeToAnnualPremium)
	}
	if t.SuggestedActionUpgradePremium != nil {
		return json.Marshal(t.SuggestedActionUpgradePremium)
	}
	if t.SuggestedActionViewChecksHint != nil {
		return json.Marshal(t.SuggestedActionViewChecksHint)
	}
	return nil, fmt.Errorf("no value set for SuggestedAction")
}

// SuggestedPostPrice Describes price of a suggested post
type SuggestedPostPrice struct {
	TypeStr                string                  `json:"@type"`
	SuggestedPostPriceStar *SuggestedPostPriceStar `json:"suggestedPostPriceStar,omitempty"`
	SuggestedPostPriceTon  *SuggestedPostPriceTon  `json:"suggestedPostPriceTon,omitempty"`
}

func (t *SuggestedPostPrice) Type() string {
	return t.TypeStr
}

func (t *SuggestedPostPrice) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SuggestedPostPrice) GetExtra() string {
	return ""
}

func (t *SuggestedPostPrice) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "suggestedPostPriceStar":
		t.SuggestedPostPriceStar = new(SuggestedPostPriceStar)
		return json.Unmarshal(b, t.SuggestedPostPriceStar)
	case "suggestedPostPriceTon":
		t.SuggestedPostPriceTon = new(SuggestedPostPriceTon)
		return json.Unmarshal(b, t.SuggestedPostPriceTon)
	}
	return nil
}

func (t *SuggestedPostPrice) MarshalJSON() ([]byte, error) {
	if t.SuggestedPostPriceStar != nil {
		return json.Marshal(t.SuggestedPostPriceStar)
	}
	if t.SuggestedPostPriceTon != nil {
		return json.Marshal(t.SuggestedPostPriceTon)
	}
	return nil, fmt.Errorf("no value set for SuggestedPostPrice")
}

// SuggestedPostRefundReason Describes reason for refund of the payment for a suggested post
type SuggestedPostRefundReason struct {
	TypeStr                                  string                                    `json:"@type"`
	SuggestedPostRefundReasonPaymentRefunded *SuggestedPostRefundReasonPaymentRefunded `json:"suggestedPostRefundReasonPaymentRefunded,omitempty"`
	SuggestedPostRefundReasonPostDeleted     *SuggestedPostRefundReasonPostDeleted     `json:"suggestedPostRefundReasonPostDeleted,omitempty"`
}

func (t *SuggestedPostRefundReason) Type() string {
	return t.TypeStr
}

func (t *SuggestedPostRefundReason) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SuggestedPostRefundReason) GetExtra() string {
	return ""
}

func (t *SuggestedPostRefundReason) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "suggestedPostRefundReasonPaymentRefunded":
		t.SuggestedPostRefundReasonPaymentRefunded = new(SuggestedPostRefundReasonPaymentRefunded)
		return json.Unmarshal(b, t.SuggestedPostRefundReasonPaymentRefunded)
	case "suggestedPostRefundReasonPostDeleted":
		t.SuggestedPostRefundReasonPostDeleted = new(SuggestedPostRefundReasonPostDeleted)
		return json.Unmarshal(b, t.SuggestedPostRefundReasonPostDeleted)
	}
	return nil
}

func (t *SuggestedPostRefundReason) MarshalJSON() ([]byte, error) {
	if t.SuggestedPostRefundReasonPaymentRefunded != nil {
		return json.Marshal(t.SuggestedPostRefundReasonPaymentRefunded)
	}
	if t.SuggestedPostRefundReasonPostDeleted != nil {
		return json.Marshal(t.SuggestedPostRefundReasonPostDeleted)
	}
	return nil, fmt.Errorf("no value set for SuggestedPostRefundReason")
}

// SuggestedPostState Describes state of a suggested post
type SuggestedPostState struct {
	TypeStr                    string                      `json:"@type"`
	SuggestedPostStateApproved *SuggestedPostStateApproved `json:"suggestedPostStateApproved,omitempty"`
	SuggestedPostStateDeclined *SuggestedPostStateDeclined `json:"suggestedPostStateDeclined,omitempty"`
	SuggestedPostStatePending  *SuggestedPostStatePending  `json:"suggestedPostStatePending,omitempty"`
}

func (t *SuggestedPostState) Type() string {
	return t.TypeStr
}

func (t *SuggestedPostState) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SuggestedPostState) GetExtra() string {
	return ""
}

func (t *SuggestedPostState) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "suggestedPostStateApproved":
		t.SuggestedPostStateApproved = new(SuggestedPostStateApproved)
		return json.Unmarshal(b, t.SuggestedPostStateApproved)
	case "suggestedPostStateDeclined":
		t.SuggestedPostStateDeclined = new(SuggestedPostStateDeclined)
		return json.Unmarshal(b, t.SuggestedPostStateDeclined)
	case "suggestedPostStatePending":
		t.SuggestedPostStatePending = new(SuggestedPostStatePending)
		return json.Unmarshal(b, t.SuggestedPostStatePending)
	}
	return nil
}

func (t *SuggestedPostState) MarshalJSON() ([]byte, error) {
	if t.SuggestedPostStateApproved != nil {
		return json.Marshal(t.SuggestedPostStateApproved)
	}
	if t.SuggestedPostStateDeclined != nil {
		return json.Marshal(t.SuggestedPostStateDeclined)
	}
	if t.SuggestedPostStatePending != nil {
		return json.Marshal(t.SuggestedPostStatePending)
	}
	return nil, fmt.Errorf("no value set for SuggestedPostState")
}

// SupergroupMembersFilter Specifies the kind of chat members to return in getSupergroupMembers
type SupergroupMembersFilter struct {
	TypeStr                               string                                 `json:"@type"`
	SupergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators `json:"supergroupMembersFilterAdministrators,omitempty"`
	SupergroupMembersFilterBanned         *SupergroupMembersFilterBanned         `json:"supergroupMembersFilterBanned,omitempty"`
	SupergroupMembersFilterBots           *SupergroupMembersFilterBots           `json:"supergroupMembersFilterBots,omitempty"`
	SupergroupMembersFilterContacts       *SupergroupMembersFilterContacts       `json:"supergroupMembersFilterContacts,omitempty"`
	SupergroupMembersFilterMention        *SupergroupMembersFilterMention        `json:"supergroupMembersFilterMention,omitempty"`
	SupergroupMembersFilterRecent         *SupergroupMembersFilterRecent         `json:"supergroupMembersFilterRecent,omitempty"`
	SupergroupMembersFilterRestricted     *SupergroupMembersFilterRestricted     `json:"supergroupMembersFilterRestricted,omitempty"`
	SupergroupMembersFilterSearch         *SupergroupMembersFilterSearch         `json:"supergroupMembersFilterSearch,omitempty"`
}

func (t *SupergroupMembersFilter) Type() string {
	return t.TypeStr
}

func (t *SupergroupMembersFilter) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *SupergroupMembersFilter) GetExtra() string {
	return ""
}

func (t *SupergroupMembersFilter) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "supergroupMembersFilterAdministrators":
		t.SupergroupMembersFilterAdministrators = new(SupergroupMembersFilterAdministrators)
		return json.Unmarshal(b, t.SupergroupMembersFilterAdministrators)
	case "supergroupMembersFilterBanned":
		t.SupergroupMembersFilterBanned = new(SupergroupMembersFilterBanned)
		return json.Unmarshal(b, t.SupergroupMembersFilterBanned)
	case "supergroupMembersFilterBots":
		t.SupergroupMembersFilterBots = new(SupergroupMembersFilterBots)
		return json.Unmarshal(b, t.SupergroupMembersFilterBots)
	case "supergroupMembersFilterContacts":
		t.SupergroupMembersFilterContacts = new(SupergroupMembersFilterContacts)
		return json.Unmarshal(b, t.SupergroupMembersFilterContacts)
	case "supergroupMembersFilterMention":
		t.SupergroupMembersFilterMention = new(SupergroupMembersFilterMention)
		return json.Unmarshal(b, t.SupergroupMembersFilterMention)
	case "supergroupMembersFilterRecent":
		t.SupergroupMembersFilterRecent = new(SupergroupMembersFilterRecent)
		return json.Unmarshal(b, t.SupergroupMembersFilterRecent)
	case "supergroupMembersFilterRestricted":
		t.SupergroupMembersFilterRestricted = new(SupergroupMembersFilterRestricted)
		return json.Unmarshal(b, t.SupergroupMembersFilterRestricted)
	case "supergroupMembersFilterSearch":
		t.SupergroupMembersFilterSearch = new(SupergroupMembersFilterSearch)
		return json.Unmarshal(b, t.SupergroupMembersFilterSearch)
	}
	return nil
}

func (t *SupergroupMembersFilter) MarshalJSON() ([]byte, error) {
	if t.SupergroupMembersFilterAdministrators != nil {
		return json.Marshal(t.SupergroupMembersFilterAdministrators)
	}
	if t.SupergroupMembersFilterBanned != nil {
		return json.Marshal(t.SupergroupMembersFilterBanned)
	}
	if t.SupergroupMembersFilterBots != nil {
		return json.Marshal(t.SupergroupMembersFilterBots)
	}
	if t.SupergroupMembersFilterContacts != nil {
		return json.Marshal(t.SupergroupMembersFilterContacts)
	}
	if t.SupergroupMembersFilterMention != nil {
		return json.Marshal(t.SupergroupMembersFilterMention)
	}
	if t.SupergroupMembersFilterRecent != nil {
		return json.Marshal(t.SupergroupMembersFilterRecent)
	}
	if t.SupergroupMembersFilterRestricted != nil {
		return json.Marshal(t.SupergroupMembersFilterRestricted)
	}
	if t.SupergroupMembersFilterSearch != nil {
		return json.Marshal(t.SupergroupMembersFilterSearch)
	}
	return nil, fmt.Errorf("no value set for SupergroupMembersFilter")
}

// TargetChat Describes the target chat to be opened
type TargetChat struct {
	TypeStr                string                  `json:"@type"`
	TargetChatChosen       *TargetChatChosen       `json:"targetChatChosen,omitempty"`
	TargetChatCurrent      *TargetChatCurrent      `json:"targetChatCurrent,omitempty"`
	TargetChatInternalLink *TargetChatInternalLink `json:"targetChatInternalLink,omitempty"`
}

func (t *TargetChat) Type() string {
	return t.TypeStr
}

func (t *TargetChat) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TargetChat) GetExtra() string {
	return ""
}

func (t *TargetChat) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "targetChatChosen":
		t.TargetChatChosen = new(TargetChatChosen)
		return json.Unmarshal(b, t.TargetChatChosen)
	case "targetChatCurrent":
		t.TargetChatCurrent = new(TargetChatCurrent)
		return json.Unmarshal(b, t.TargetChatCurrent)
	case "targetChatInternalLink":
		t.TargetChatInternalLink = new(TargetChatInternalLink)
		return json.Unmarshal(b, t.TargetChatInternalLink)
	}
	return nil
}

func (t *TargetChat) MarshalJSON() ([]byte, error) {
	if t.TargetChatChosen != nil {
		return json.Marshal(t.TargetChatChosen)
	}
	if t.TargetChatCurrent != nil {
		return json.Marshal(t.TargetChatCurrent)
	}
	if t.TargetChatInternalLink != nil {
		return json.Marshal(t.TargetChatInternalLink)
	}
	return nil, fmt.Errorf("no value set for TargetChat")
}

// TelegramPaymentPurpose Describes a purpose of a payment toward Telegram
type TelegramPaymentPurpose struct {
	TypeStr                                string                                  `json:"@type"`
	TelegramPaymentPurposeGiftedStars      *TelegramPaymentPurposeGiftedStars      `json:"telegramPaymentPurposeGiftedStars,omitempty"`
	TelegramPaymentPurposeJoinChat         *TelegramPaymentPurposeJoinChat         `json:"telegramPaymentPurposeJoinChat,omitempty"`
	TelegramPaymentPurposePremiumGift      *TelegramPaymentPurposePremiumGift      `json:"telegramPaymentPurposePremiumGift,omitempty"`
	TelegramPaymentPurposePremiumGiftCodes *TelegramPaymentPurposePremiumGiftCodes `json:"telegramPaymentPurposePremiumGiftCodes,omitempty"`
	TelegramPaymentPurposePremiumGiveaway  *TelegramPaymentPurposePremiumGiveaway  `json:"telegramPaymentPurposePremiumGiveaway,omitempty"`
	TelegramPaymentPurposeStarGiveaway     *TelegramPaymentPurposeStarGiveaway     `json:"telegramPaymentPurposeStarGiveaway,omitempty"`
	TelegramPaymentPurposeStars            *TelegramPaymentPurposeStars            `json:"telegramPaymentPurposeStars,omitempty"`
}

func (t *TelegramPaymentPurpose) Type() string {
	return t.TypeStr
}

func (t *TelegramPaymentPurpose) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TelegramPaymentPurpose) GetExtra() string {
	return ""
}

func (t *TelegramPaymentPurpose) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "telegramPaymentPurposeGiftedStars":
		t.TelegramPaymentPurposeGiftedStars = new(TelegramPaymentPurposeGiftedStars)
		return json.Unmarshal(b, t.TelegramPaymentPurposeGiftedStars)
	case "telegramPaymentPurposeJoinChat":
		t.TelegramPaymentPurposeJoinChat = new(TelegramPaymentPurposeJoinChat)
		return json.Unmarshal(b, t.TelegramPaymentPurposeJoinChat)
	case "telegramPaymentPurposePremiumGift":
		t.TelegramPaymentPurposePremiumGift = new(TelegramPaymentPurposePremiumGift)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGift)
	case "telegramPaymentPurposePremiumGiftCodes":
		t.TelegramPaymentPurposePremiumGiftCodes = new(TelegramPaymentPurposePremiumGiftCodes)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGiftCodes)
	case "telegramPaymentPurposePremiumGiveaway":
		t.TelegramPaymentPurposePremiumGiveaway = new(TelegramPaymentPurposePremiumGiveaway)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGiveaway)
	case "telegramPaymentPurposeStarGiveaway":
		t.TelegramPaymentPurposeStarGiveaway = new(TelegramPaymentPurposeStarGiveaway)
		return json.Unmarshal(b, t.TelegramPaymentPurposeStarGiveaway)
	case "telegramPaymentPurposeStars":
		t.TelegramPaymentPurposeStars = new(TelegramPaymentPurposeStars)
		return json.Unmarshal(b, t.TelegramPaymentPurposeStars)
	}
	return nil
}

func (t *TelegramPaymentPurpose) MarshalJSON() ([]byte, error) {
	if t.TelegramPaymentPurposeGiftedStars != nil {
		return json.Marshal(t.TelegramPaymentPurposeGiftedStars)
	}
	if t.TelegramPaymentPurposeJoinChat != nil {
		return json.Marshal(t.TelegramPaymentPurposeJoinChat)
	}
	if t.TelegramPaymentPurposePremiumGift != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGift)
	}
	if t.TelegramPaymentPurposePremiumGiftCodes != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGiftCodes)
	}
	if t.TelegramPaymentPurposePremiumGiveaway != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGiveaway)
	}
	if t.TelegramPaymentPurposeStarGiveaway != nil {
		return json.Marshal(t.TelegramPaymentPurposeStarGiveaway)
	}
	if t.TelegramPaymentPurposeStars != nil {
		return json.Marshal(t.TelegramPaymentPurposeStars)
	}
	return nil, fmt.Errorf("no value set for TelegramPaymentPurpose")
}

// TextEntityType Represents a part of the text which must be formatted differently
type TextEntityType struct {
	TypeStr                            string                              `json:"@type"`
	TextEntityTypeBankCardNumber       *TextEntityTypeBankCardNumber       `json:"textEntityTypeBankCardNumber,omitempty"`
	TextEntityTypeBlockQuote           *TextEntityTypeBlockQuote           `json:"textEntityTypeBlockQuote,omitempty"`
	TextEntityTypeBold                 *TextEntityTypeBold                 `json:"textEntityTypeBold,omitempty"`
	TextEntityTypeBotCommand           *TextEntityTypeBotCommand           `json:"textEntityTypeBotCommand,omitempty"`
	TextEntityTypeCashtag              *TextEntityTypeCashtag              `json:"textEntityTypeCashtag,omitempty"`
	TextEntityTypeCode                 *TextEntityTypeCode                 `json:"textEntityTypeCode,omitempty"`
	TextEntityTypeCustomEmoji          *TextEntityTypeCustomEmoji          `json:"textEntityTypeCustomEmoji,omitempty"`
	TextEntityTypeEmailAddress         *TextEntityTypeEmailAddress         `json:"textEntityTypeEmailAddress,omitempty"`
	TextEntityTypeExpandableBlockQuote *TextEntityTypeExpandableBlockQuote `json:"textEntityTypeExpandableBlockQuote,omitempty"`
	TextEntityTypeHashtag              *TextEntityTypeHashtag              `json:"textEntityTypeHashtag,omitempty"`
	TextEntityTypeItalic               *TextEntityTypeItalic               `json:"textEntityTypeItalic,omitempty"`
	TextEntityTypeMediaTimestamp       *TextEntityTypeMediaTimestamp       `json:"textEntityTypeMediaTimestamp,omitempty"`
	TextEntityTypeMention              *TextEntityTypeMention              `json:"textEntityTypeMention,omitempty"`
	TextEntityTypeMentionName          *TextEntityTypeMentionName          `json:"textEntityTypeMentionName,omitempty"`
	TextEntityTypePhoneNumber          *TextEntityTypePhoneNumber          `json:"textEntityTypePhoneNumber,omitempty"`
	TextEntityTypePre                  *TextEntityTypePre                  `json:"textEntityTypePre,omitempty"`
	TextEntityTypePreCode              *TextEntityTypePreCode              `json:"textEntityTypePreCode,omitempty"`
	TextEntityTypeSpoiler              *TextEntityTypeSpoiler              `json:"textEntityTypeSpoiler,omitempty"`
	TextEntityTypeStrikethrough        *TextEntityTypeStrikethrough        `json:"textEntityTypeStrikethrough,omitempty"`
	TextEntityTypeTextUrl              *TextEntityTypeTextUrl              `json:"textEntityTypeTextUrl,omitempty"`
	TextEntityTypeUnderline            *TextEntityTypeUnderline            `json:"textEntityTypeUnderline,omitempty"`
	TextEntityTypeUrl                  *TextEntityTypeUrl                  `json:"textEntityTypeUrl,omitempty"`
}

func (t *TextEntityType) Type() string {
	return t.TypeStr
}

func (t *TextEntityType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TextEntityType) GetExtra() string {
	return ""
}

func (t *TextEntityType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "textEntityTypeBankCardNumber":
		t.TextEntityTypeBankCardNumber = new(TextEntityTypeBankCardNumber)
		return json.Unmarshal(b, t.TextEntityTypeBankCardNumber)
	case "textEntityTypeBlockQuote":
		t.TextEntityTypeBlockQuote = new(TextEntityTypeBlockQuote)
		return json.Unmarshal(b, t.TextEntityTypeBlockQuote)
	case "textEntityTypeBold":
		t.TextEntityTypeBold = new(TextEntityTypeBold)
		return json.Unmarshal(b, t.TextEntityTypeBold)
	case "textEntityTypeBotCommand":
		t.TextEntityTypeBotCommand = new(TextEntityTypeBotCommand)
		return json.Unmarshal(b, t.TextEntityTypeBotCommand)
	case "textEntityTypeCashtag":
		t.TextEntityTypeCashtag = new(TextEntityTypeCashtag)
		return json.Unmarshal(b, t.TextEntityTypeCashtag)
	case "textEntityTypeCode":
		t.TextEntityTypeCode = new(TextEntityTypeCode)
		return json.Unmarshal(b, t.TextEntityTypeCode)
	case "textEntityTypeCustomEmoji":
		t.TextEntityTypeCustomEmoji = new(TextEntityTypeCustomEmoji)
		return json.Unmarshal(b, t.TextEntityTypeCustomEmoji)
	case "textEntityTypeEmailAddress":
		t.TextEntityTypeEmailAddress = new(TextEntityTypeEmailAddress)
		return json.Unmarshal(b, t.TextEntityTypeEmailAddress)
	case "textEntityTypeExpandableBlockQuote":
		t.TextEntityTypeExpandableBlockQuote = new(TextEntityTypeExpandableBlockQuote)
		return json.Unmarshal(b, t.TextEntityTypeExpandableBlockQuote)
	case "textEntityTypeHashtag":
		t.TextEntityTypeHashtag = new(TextEntityTypeHashtag)
		return json.Unmarshal(b, t.TextEntityTypeHashtag)
	case "textEntityTypeItalic":
		t.TextEntityTypeItalic = new(TextEntityTypeItalic)
		return json.Unmarshal(b, t.TextEntityTypeItalic)
	case "textEntityTypeMediaTimestamp":
		t.TextEntityTypeMediaTimestamp = new(TextEntityTypeMediaTimestamp)
		return json.Unmarshal(b, t.TextEntityTypeMediaTimestamp)
	case "textEntityTypeMention":
		t.TextEntityTypeMention = new(TextEntityTypeMention)
		return json.Unmarshal(b, t.TextEntityTypeMention)
	case "textEntityTypeMentionName":
		t.TextEntityTypeMentionName = new(TextEntityTypeMentionName)
		return json.Unmarshal(b, t.TextEntityTypeMentionName)
	case "textEntityTypePhoneNumber":
		t.TextEntityTypePhoneNumber = new(TextEntityTypePhoneNumber)
		return json.Unmarshal(b, t.TextEntityTypePhoneNumber)
	case "textEntityTypePre":
		t.TextEntityTypePre = new(TextEntityTypePre)
		return json.Unmarshal(b, t.TextEntityTypePre)
	case "textEntityTypePreCode":
		t.TextEntityTypePreCode = new(TextEntityTypePreCode)
		return json.Unmarshal(b, t.TextEntityTypePreCode)
	case "textEntityTypeSpoiler":
		t.TextEntityTypeSpoiler = new(TextEntityTypeSpoiler)
		return json.Unmarshal(b, t.TextEntityTypeSpoiler)
	case "textEntityTypeStrikethrough":
		t.TextEntityTypeStrikethrough = new(TextEntityTypeStrikethrough)
		return json.Unmarshal(b, t.TextEntityTypeStrikethrough)
	case "textEntityTypeTextUrl":
		t.TextEntityTypeTextUrl = new(TextEntityTypeTextUrl)
		return json.Unmarshal(b, t.TextEntityTypeTextUrl)
	case "textEntityTypeUnderline":
		t.TextEntityTypeUnderline = new(TextEntityTypeUnderline)
		return json.Unmarshal(b, t.TextEntityTypeUnderline)
	case "textEntityTypeUrl":
		t.TextEntityTypeUrl = new(TextEntityTypeUrl)
		return json.Unmarshal(b, t.TextEntityTypeUrl)
	}
	return nil
}

func (t *TextEntityType) MarshalJSON() ([]byte, error) {
	if t.TextEntityTypeBankCardNumber != nil {
		return json.Marshal(t.TextEntityTypeBankCardNumber)
	}
	if t.TextEntityTypeBlockQuote != nil {
		return json.Marshal(t.TextEntityTypeBlockQuote)
	}
	if t.TextEntityTypeBold != nil {
		return json.Marshal(t.TextEntityTypeBold)
	}
	if t.TextEntityTypeBotCommand != nil {
		return json.Marshal(t.TextEntityTypeBotCommand)
	}
	if t.TextEntityTypeCashtag != nil {
		return json.Marshal(t.TextEntityTypeCashtag)
	}
	if t.TextEntityTypeCode != nil {
		return json.Marshal(t.TextEntityTypeCode)
	}
	if t.TextEntityTypeCustomEmoji != nil {
		return json.Marshal(t.TextEntityTypeCustomEmoji)
	}
	if t.TextEntityTypeEmailAddress != nil {
		return json.Marshal(t.TextEntityTypeEmailAddress)
	}
	if t.TextEntityTypeExpandableBlockQuote != nil {
		return json.Marshal(t.TextEntityTypeExpandableBlockQuote)
	}
	if t.TextEntityTypeHashtag != nil {
		return json.Marshal(t.TextEntityTypeHashtag)
	}
	if t.TextEntityTypeItalic != nil {
		return json.Marshal(t.TextEntityTypeItalic)
	}
	if t.TextEntityTypeMediaTimestamp != nil {
		return json.Marshal(t.TextEntityTypeMediaTimestamp)
	}
	if t.TextEntityTypeMention != nil {
		return json.Marshal(t.TextEntityTypeMention)
	}
	if t.TextEntityTypeMentionName != nil {
		return json.Marshal(t.TextEntityTypeMentionName)
	}
	if t.TextEntityTypePhoneNumber != nil {
		return json.Marshal(t.TextEntityTypePhoneNumber)
	}
	if t.TextEntityTypePre != nil {
		return json.Marshal(t.TextEntityTypePre)
	}
	if t.TextEntityTypePreCode != nil {
		return json.Marshal(t.TextEntityTypePreCode)
	}
	if t.TextEntityTypeSpoiler != nil {
		return json.Marshal(t.TextEntityTypeSpoiler)
	}
	if t.TextEntityTypeStrikethrough != nil {
		return json.Marshal(t.TextEntityTypeStrikethrough)
	}
	if t.TextEntityTypeTextUrl != nil {
		return json.Marshal(t.TextEntityTypeTextUrl)
	}
	if t.TextEntityTypeUnderline != nil {
		return json.Marshal(t.TextEntityTypeUnderline)
	}
	if t.TextEntityTypeUrl != nil {
		return json.Marshal(t.TextEntityTypeUrl)
	}
	return nil, fmt.Errorf("no value set for TextEntityType")
}

// TextParseMode Describes the way the text needs to be parsed for text entities
type TextParseMode struct {
	TypeStr               string                 `json:"@type"`
	TextParseModeHTML     *TextParseModeHTML     `json:"textParseModeHTML,omitempty"`
	TextParseModeMarkdown *TextParseModeMarkdown `json:"textParseModeMarkdown,omitempty"`
}

func (t *TextParseMode) Type() string {
	return t.TypeStr
}

func (t *TextParseMode) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TextParseMode) GetExtra() string {
	return ""
}

func (t *TextParseMode) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "textParseModeHTML":
		t.TextParseModeHTML = new(TextParseModeHTML)
		return json.Unmarshal(b, t.TextParseModeHTML)
	case "textParseModeMarkdown":
		t.TextParseModeMarkdown = new(TextParseModeMarkdown)
		return json.Unmarshal(b, t.TextParseModeMarkdown)
	}
	return nil
}

func (t *TextParseMode) MarshalJSON() ([]byte, error) {
	if t.TextParseModeHTML != nil {
		return json.Marshal(t.TextParseModeHTML)
	}
	if t.TextParseModeMarkdown != nil {
		return json.Marshal(t.TextParseModeMarkdown)
	}
	return nil, fmt.Errorf("no value set for TextParseMode")
}

// ThumbnailFormat Describes format of a thumbnail
type ThumbnailFormat struct {
	TypeStr              string                `json:"@type"`
	ThumbnailFormatGif   *ThumbnailFormatGif   `json:"thumbnailFormatGif,omitempty"`
	ThumbnailFormatJpeg  *ThumbnailFormatJpeg  `json:"thumbnailFormatJpeg,omitempty"`
	ThumbnailFormatMpeg4 *ThumbnailFormatMpeg4 `json:"thumbnailFormatMpeg4,omitempty"`
	ThumbnailFormatPng   *ThumbnailFormatPng   `json:"thumbnailFormatPng,omitempty"`
	ThumbnailFormatTgs   *ThumbnailFormatTgs   `json:"thumbnailFormatTgs,omitempty"`
	ThumbnailFormatWebm  *ThumbnailFormatWebm  `json:"thumbnailFormatWebm,omitempty"`
	ThumbnailFormatWebp  *ThumbnailFormatWebp  `json:"thumbnailFormatWebp,omitempty"`
}

func (t *ThumbnailFormat) Type() string {
	return t.TypeStr
}

func (t *ThumbnailFormat) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *ThumbnailFormat) GetExtra() string {
	return ""
}

func (t *ThumbnailFormat) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "thumbnailFormatGif":
		t.ThumbnailFormatGif = new(ThumbnailFormatGif)
		return json.Unmarshal(b, t.ThumbnailFormatGif)
	case "thumbnailFormatJpeg":
		t.ThumbnailFormatJpeg = new(ThumbnailFormatJpeg)
		return json.Unmarshal(b, t.ThumbnailFormatJpeg)
	case "thumbnailFormatMpeg4":
		t.ThumbnailFormatMpeg4 = new(ThumbnailFormatMpeg4)
		return json.Unmarshal(b, t.ThumbnailFormatMpeg4)
	case "thumbnailFormatPng":
		t.ThumbnailFormatPng = new(ThumbnailFormatPng)
		return json.Unmarshal(b, t.ThumbnailFormatPng)
	case "thumbnailFormatTgs":
		t.ThumbnailFormatTgs = new(ThumbnailFormatTgs)
		return json.Unmarshal(b, t.ThumbnailFormatTgs)
	case "thumbnailFormatWebm":
		t.ThumbnailFormatWebm = new(ThumbnailFormatWebm)
		return json.Unmarshal(b, t.ThumbnailFormatWebm)
	case "thumbnailFormatWebp":
		t.ThumbnailFormatWebp = new(ThumbnailFormatWebp)
		return json.Unmarshal(b, t.ThumbnailFormatWebp)
	}
	return nil
}

func (t *ThumbnailFormat) MarshalJSON() ([]byte, error) {
	if t.ThumbnailFormatGif != nil {
		return json.Marshal(t.ThumbnailFormatGif)
	}
	if t.ThumbnailFormatJpeg != nil {
		return json.Marshal(t.ThumbnailFormatJpeg)
	}
	if t.ThumbnailFormatMpeg4 != nil {
		return json.Marshal(t.ThumbnailFormatMpeg4)
	}
	if t.ThumbnailFormatPng != nil {
		return json.Marshal(t.ThumbnailFormatPng)
	}
	if t.ThumbnailFormatTgs != nil {
		return json.Marshal(t.ThumbnailFormatTgs)
	}
	if t.ThumbnailFormatWebm != nil {
		return json.Marshal(t.ThumbnailFormatWebm)
	}
	if t.ThumbnailFormatWebp != nil {
		return json.Marshal(t.ThumbnailFormatWebp)
	}
	return nil, fmt.Errorf("no value set for ThumbnailFormat")
}

// TMeUrlType Describes the type of URL linking to an internal Telegram entity
type TMeUrlType struct {
	TypeStr              string                `json:"@type"`
	TMeUrlTypeChatInvite *TMeUrlTypeChatInvite `json:"tMeUrlTypeChatInvite,omitempty"`
	TMeUrlTypeStickerSet *TMeUrlTypeStickerSet `json:"tMeUrlTypeStickerSet,omitempty"`
	TMeUrlTypeSupergroup *TMeUrlTypeSupergroup `json:"tMeUrlTypeSupergroup,omitempty"`
	TMeUrlTypeUser       *TMeUrlTypeUser       `json:"tMeUrlTypeUser,omitempty"`
}

func (t *TMeUrlType) Type() string {
	return t.TypeStr
}

func (t *TMeUrlType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TMeUrlType) GetExtra() string {
	return ""
}

func (t *TMeUrlType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "tMeUrlTypeChatInvite":
		t.TMeUrlTypeChatInvite = new(TMeUrlTypeChatInvite)
		return json.Unmarshal(b, t.TMeUrlTypeChatInvite)
	case "tMeUrlTypeStickerSet":
		t.TMeUrlTypeStickerSet = new(TMeUrlTypeStickerSet)
		return json.Unmarshal(b, t.TMeUrlTypeStickerSet)
	case "tMeUrlTypeSupergroup":
		t.TMeUrlTypeSupergroup = new(TMeUrlTypeSupergroup)
		return json.Unmarshal(b, t.TMeUrlTypeSupergroup)
	case "tMeUrlTypeUser":
		t.TMeUrlTypeUser = new(TMeUrlTypeUser)
		return json.Unmarshal(b, t.TMeUrlTypeUser)
	}
	return nil
}

func (t *TMeUrlType) MarshalJSON() ([]byte, error) {
	if t.TMeUrlTypeChatInvite != nil {
		return json.Marshal(t.TMeUrlTypeChatInvite)
	}
	if t.TMeUrlTypeStickerSet != nil {
		return json.Marshal(t.TMeUrlTypeStickerSet)
	}
	if t.TMeUrlTypeSupergroup != nil {
		return json.Marshal(t.TMeUrlTypeSupergroup)
	}
	if t.TMeUrlTypeUser != nil {
		return json.Marshal(t.TMeUrlTypeUser)
	}
	return nil, fmt.Errorf("no value set for TMeUrlType")
}

// TonTransactionType Describes type of transaction with Toncoins
type TonTransactionType struct {
	TypeStr                                string                                  `json:"@type"`
	TonTransactionTypeFragmentDeposit      *TonTransactionTypeFragmentDeposit      `json:"tonTransactionTypeFragmentDeposit,omitempty"`
	TonTransactionTypeFragmentWithdrawal   *TonTransactionTypeFragmentWithdrawal   `json:"tonTransactionTypeFragmentWithdrawal,omitempty"`
	TonTransactionTypeGiftPurchaseOffer    *TonTransactionTypeGiftPurchaseOffer    `json:"tonTransactionTypeGiftPurchaseOffer,omitempty"`
	TonTransactionTypeSuggestedPostPayment *TonTransactionTypeSuggestedPostPayment `json:"tonTransactionTypeSuggestedPostPayment,omitempty"`
	TonTransactionTypeUnsupported          *TonTransactionTypeUnsupported          `json:"tonTransactionTypeUnsupported,omitempty"`
	TonTransactionTypeUpgradedGiftPurchase *TonTransactionTypeUpgradedGiftPurchase `json:"tonTransactionTypeUpgradedGiftPurchase,omitempty"`
	TonTransactionTypeUpgradedGiftSale     *TonTransactionTypeUpgradedGiftSale     `json:"tonTransactionTypeUpgradedGiftSale,omitempty"`
}

func (t *TonTransactionType) Type() string {
	return t.TypeStr
}

func (t *TonTransactionType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TonTransactionType) GetExtra() string {
	return ""
}

func (t *TonTransactionType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "tonTransactionTypeFragmentDeposit":
		t.TonTransactionTypeFragmentDeposit = new(TonTransactionTypeFragmentDeposit)
		return json.Unmarshal(b, t.TonTransactionTypeFragmentDeposit)
	case "tonTransactionTypeFragmentWithdrawal":
		t.TonTransactionTypeFragmentWithdrawal = new(TonTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.TonTransactionTypeFragmentWithdrawal)
	case "tonTransactionTypeGiftPurchaseOffer":
		t.TonTransactionTypeGiftPurchaseOffer = new(TonTransactionTypeGiftPurchaseOffer)
		return json.Unmarshal(b, t.TonTransactionTypeGiftPurchaseOffer)
	case "tonTransactionTypeSuggestedPostPayment":
		t.TonTransactionTypeSuggestedPostPayment = new(TonTransactionTypeSuggestedPostPayment)
		return json.Unmarshal(b, t.TonTransactionTypeSuggestedPostPayment)
	case "tonTransactionTypeUnsupported":
		t.TonTransactionTypeUnsupported = new(TonTransactionTypeUnsupported)
		return json.Unmarshal(b, t.TonTransactionTypeUnsupported)
	case "tonTransactionTypeUpgradedGiftPurchase":
		t.TonTransactionTypeUpgradedGiftPurchase = new(TonTransactionTypeUpgradedGiftPurchase)
		return json.Unmarshal(b, t.TonTransactionTypeUpgradedGiftPurchase)
	case "tonTransactionTypeUpgradedGiftSale":
		t.TonTransactionTypeUpgradedGiftSale = new(TonTransactionTypeUpgradedGiftSale)
		return json.Unmarshal(b, t.TonTransactionTypeUpgradedGiftSale)
	}
	return nil
}

func (t *TonTransactionType) MarshalJSON() ([]byte, error) {
	if t.TonTransactionTypeFragmentDeposit != nil {
		return json.Marshal(t.TonTransactionTypeFragmentDeposit)
	}
	if t.TonTransactionTypeFragmentWithdrawal != nil {
		return json.Marshal(t.TonTransactionTypeFragmentWithdrawal)
	}
	if t.TonTransactionTypeGiftPurchaseOffer != nil {
		return json.Marshal(t.TonTransactionTypeGiftPurchaseOffer)
	}
	if t.TonTransactionTypeSuggestedPostPayment != nil {
		return json.Marshal(t.TonTransactionTypeSuggestedPostPayment)
	}
	if t.TonTransactionTypeUnsupported != nil {
		return json.Marshal(t.TonTransactionTypeUnsupported)
	}
	if t.TonTransactionTypeUpgradedGiftPurchase != nil {
		return json.Marshal(t.TonTransactionTypeUpgradedGiftPurchase)
	}
	if t.TonTransactionTypeUpgradedGiftSale != nil {
		return json.Marshal(t.TonTransactionTypeUpgradedGiftSale)
	}
	return nil, fmt.Errorf("no value set for TonTransactionType")
}

// TopChatCategory Represents the categories of chats for which a list of frequently used chats can be retrieved
type TopChatCategory struct {
	TypeStr                     string                       `json:"@type"`
	TopChatCategoryBots         *TopChatCategoryBots         `json:"topChatCategoryBots,omitempty"`
	TopChatCategoryCalls        *TopChatCategoryCalls        `json:"topChatCategoryCalls,omitempty"`
	TopChatCategoryChannels     *TopChatCategoryChannels     `json:"topChatCategoryChannels,omitempty"`
	TopChatCategoryForwardChats *TopChatCategoryForwardChats `json:"topChatCategoryForwardChats,omitempty"`
	TopChatCategoryGroups       *TopChatCategoryGroups       `json:"topChatCategoryGroups,omitempty"`
	TopChatCategoryInlineBots   *TopChatCategoryInlineBots   `json:"topChatCategoryInlineBots,omitempty"`
	TopChatCategoryUsers        *TopChatCategoryUsers        `json:"topChatCategoryUsers,omitempty"`
	TopChatCategoryWebAppBots   *TopChatCategoryWebAppBots   `json:"topChatCategoryWebAppBots,omitempty"`
}

func (t *TopChatCategory) Type() string {
	return t.TypeStr
}

func (t *TopChatCategory) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TopChatCategory) GetExtra() string {
	return ""
}

func (t *TopChatCategory) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "topChatCategoryBots":
		t.TopChatCategoryBots = new(TopChatCategoryBots)
		return json.Unmarshal(b, t.TopChatCategoryBots)
	case "topChatCategoryCalls":
		t.TopChatCategoryCalls = new(TopChatCategoryCalls)
		return json.Unmarshal(b, t.TopChatCategoryCalls)
	case "topChatCategoryChannels":
		t.TopChatCategoryChannels = new(TopChatCategoryChannels)
		return json.Unmarshal(b, t.TopChatCategoryChannels)
	case "topChatCategoryForwardChats":
		t.TopChatCategoryForwardChats = new(TopChatCategoryForwardChats)
		return json.Unmarshal(b, t.TopChatCategoryForwardChats)
	case "topChatCategoryGroups":
		t.TopChatCategoryGroups = new(TopChatCategoryGroups)
		return json.Unmarshal(b, t.TopChatCategoryGroups)
	case "topChatCategoryInlineBots":
		t.TopChatCategoryInlineBots = new(TopChatCategoryInlineBots)
		return json.Unmarshal(b, t.TopChatCategoryInlineBots)
	case "topChatCategoryUsers":
		t.TopChatCategoryUsers = new(TopChatCategoryUsers)
		return json.Unmarshal(b, t.TopChatCategoryUsers)
	case "topChatCategoryWebAppBots":
		t.TopChatCategoryWebAppBots = new(TopChatCategoryWebAppBots)
		return json.Unmarshal(b, t.TopChatCategoryWebAppBots)
	}
	return nil
}

func (t *TopChatCategory) MarshalJSON() ([]byte, error) {
	if t.TopChatCategoryBots != nil {
		return json.Marshal(t.TopChatCategoryBots)
	}
	if t.TopChatCategoryCalls != nil {
		return json.Marshal(t.TopChatCategoryCalls)
	}
	if t.TopChatCategoryChannels != nil {
		return json.Marshal(t.TopChatCategoryChannels)
	}
	if t.TopChatCategoryForwardChats != nil {
		return json.Marshal(t.TopChatCategoryForwardChats)
	}
	if t.TopChatCategoryGroups != nil {
		return json.Marshal(t.TopChatCategoryGroups)
	}
	if t.TopChatCategoryInlineBots != nil {
		return json.Marshal(t.TopChatCategoryInlineBots)
	}
	if t.TopChatCategoryUsers != nil {
		return json.Marshal(t.TopChatCategoryUsers)
	}
	if t.TopChatCategoryWebAppBots != nil {
		return json.Marshal(t.TopChatCategoryWebAppBots)
	}
	return nil, fmt.Errorf("no value set for TopChatCategory")
}

// TransactionDirection Describes direction of transactions in a transaction list
type TransactionDirection struct {
	TypeStr                      string                        `json:"@type"`
	TransactionDirectionIncoming *TransactionDirectionIncoming `json:"transactionDirectionIncoming,omitempty"`
	TransactionDirectionOutgoing *TransactionDirectionOutgoing `json:"transactionDirectionOutgoing,omitempty"`
}

func (t *TransactionDirection) Type() string {
	return t.TypeStr
}

func (t *TransactionDirection) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *TransactionDirection) GetExtra() string {
	return ""
}

func (t *TransactionDirection) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "transactionDirectionIncoming":
		t.TransactionDirectionIncoming = new(TransactionDirectionIncoming)
		return json.Unmarshal(b, t.TransactionDirectionIncoming)
	case "transactionDirectionOutgoing":
		t.TransactionDirectionOutgoing = new(TransactionDirectionOutgoing)
		return json.Unmarshal(b, t.TransactionDirectionOutgoing)
	}
	return nil
}

func (t *TransactionDirection) MarshalJSON() ([]byte, error) {
	if t.TransactionDirectionIncoming != nil {
		return json.Marshal(t.TransactionDirectionIncoming)
	}
	if t.TransactionDirectionOutgoing != nil {
		return json.Marshal(t.TransactionDirectionOutgoing)
	}
	return nil, fmt.Errorf("no value set for TransactionDirection")
}

// Update Contains notifications about data changes
type Update struct {
	TypeStr                                        string                                          `json:"@type"`
	UpdateAccentColors                             *UpdateAccentColors                             `json:"updateAccentColors,omitempty"`
	UpdateActiveEmojiReactions                     *UpdateActiveEmojiReactions                     `json:"updateActiveEmojiReactions,omitempty"`
	UpdateActiveGiftAuctions                       *UpdateActiveGiftAuctions                       `json:"updateActiveGiftAuctions,omitempty"`
	UpdateActiveLiveLocationMessages               *UpdateActiveLiveLocationMessages               `json:"updateActiveLiveLocationMessages,omitempty"`
	UpdateActiveNotifications                      *UpdateActiveNotifications                      `json:"updateActiveNotifications,omitempty"`
	UpdateAgeVerificationParameters                *UpdateAgeVerificationParameters                `json:"updateAgeVerificationParameters,omitempty"`
	UpdateAnimatedEmojiMessageClicked              *UpdateAnimatedEmojiMessageClicked              `json:"updateAnimatedEmojiMessageClicked,omitempty"`
	UpdateAnimationSearchParameters                *UpdateAnimationSearchParameters                `json:"updateAnimationSearchParameters,omitempty"`
	UpdateApplicationRecaptchaVerificationRequired *UpdateApplicationRecaptchaVerificationRequired `json:"updateApplicationRecaptchaVerificationRequired,omitempty"`
	UpdateApplicationVerificationRequired          *UpdateApplicationVerificationRequired          `json:"updateApplicationVerificationRequired,omitempty"`
	UpdateAttachmentMenuBots                       *UpdateAttachmentMenuBots                       `json:"updateAttachmentMenuBots,omitempty"`
	UpdateAuthorizationState                       *UpdateAuthorizationState                       `json:"updateAuthorizationState,omitempty"`
	UpdateAutosaveSettings                         *UpdateAutosaveSettings                         `json:"updateAutosaveSettings,omitempty"`
	UpdateAvailableMessageEffects                  *UpdateAvailableMessageEffects                  `json:"updateAvailableMessageEffects,omitempty"`
	UpdateBasicGroup                               *UpdateBasicGroup                               `json:"updateBasicGroup,omitempty"`
	UpdateBasicGroupFullInfo                       *UpdateBasicGroupFullInfo                       `json:"updateBasicGroupFullInfo,omitempty"`
	UpdateBusinessConnection                       *UpdateBusinessConnection                       `json:"updateBusinessConnection,omitempty"`
	UpdateBusinessMessageEdited                    *UpdateBusinessMessageEdited                    `json:"updateBusinessMessageEdited,omitempty"`
	UpdateBusinessMessagesDeleted                  *UpdateBusinessMessagesDeleted                  `json:"updateBusinessMessagesDeleted,omitempty"`
	UpdateCall                                     *UpdateCall                                     `json:"updateCall,omitempty"`
	UpdateChatAccentColors                         *UpdateChatAccentColors                         `json:"updateChatAccentColors,omitempty"`
	UpdateChatAction                               *UpdateChatAction                               `json:"updateChatAction,omitempty"`
	UpdateChatActionBar                            *UpdateChatActionBar                            `json:"updateChatActionBar,omitempty"`
	UpdateChatActiveStories                        *UpdateChatActiveStories                        `json:"updateChatActiveStories,omitempty"`
	UpdateChatAddedToList                          *UpdateChatAddedToList                          `json:"updateChatAddedToList,omitempty"`
	UpdateChatAvailableReactions                   *UpdateChatAvailableReactions                   `json:"updateChatAvailableReactions,omitempty"`
	UpdateChatBackground                           *UpdateChatBackground                           `json:"updateChatBackground,omitempty"`
	UpdateChatBlockList                            *UpdateChatBlockList                            `json:"updateChatBlockList,omitempty"`
	UpdateChatBoost                                *UpdateChatBoost                                `json:"updateChatBoost,omitempty"`
	UpdateChatBusinessBotManageBar                 *UpdateChatBusinessBotManageBar                 `json:"updateChatBusinessBotManageBar,omitempty"`
	UpdateChatDefaultDisableNotification           *UpdateChatDefaultDisableNotification           `json:"updateChatDefaultDisableNotification,omitempty"`
	UpdateChatDraftMessage                         *UpdateChatDraftMessage                         `json:"updateChatDraftMessage,omitempty"`
	UpdateChatEmojiStatus                          *UpdateChatEmojiStatus                          `json:"updateChatEmojiStatus,omitempty"`
	UpdateChatFolders                              *UpdateChatFolders                              `json:"updateChatFolders,omitempty"`
	UpdateChatHasProtectedContent                  *UpdateChatHasProtectedContent                  `json:"updateChatHasProtectedContent,omitempty"`
	UpdateChatHasScheduledMessages                 *UpdateChatHasScheduledMessages                 `json:"updateChatHasScheduledMessages,omitempty"`
	UpdateChatIsMarkedAsUnread                     *UpdateChatIsMarkedAsUnread                     `json:"updateChatIsMarkedAsUnread,omitempty"`
	UpdateChatIsTranslatable                       *UpdateChatIsTranslatable                       `json:"updateChatIsTranslatable,omitempty"`
	UpdateChatLastMessage                          *UpdateChatLastMessage                          `json:"updateChatLastMessage,omitempty"`
	UpdateChatMember                               *UpdateChatMember                               `json:"updateChatMember,omitempty"`
	UpdateChatMessageAutoDeleteTime                *UpdateChatMessageAutoDeleteTime                `json:"updateChatMessageAutoDeleteTime,omitempty"`
	UpdateChatMessageSender                        *UpdateChatMessageSender                        `json:"updateChatMessageSender,omitempty"`
	UpdateChatNotificationSettings                 *UpdateChatNotificationSettings                 `json:"updateChatNotificationSettings,omitempty"`
	UpdateChatOnlineMemberCount                    *UpdateChatOnlineMemberCount                    `json:"updateChatOnlineMemberCount,omitempty"`
	UpdateChatPendingJoinRequests                  *UpdateChatPendingJoinRequests                  `json:"updateChatPendingJoinRequests,omitempty"`
	UpdateChatPermissions                          *UpdateChatPermissions                          `json:"updateChatPermissions,omitempty"`
	UpdateChatPhoto                                *UpdateChatPhoto                                `json:"updateChatPhoto,omitempty"`
	UpdateChatPosition                             *UpdateChatPosition                             `json:"updateChatPosition,omitempty"`
	UpdateChatReadInbox                            *UpdateChatReadInbox                            `json:"updateChatReadInbox,omitempty"`
	UpdateChatReadOutbox                           *UpdateChatReadOutbox                           `json:"updateChatReadOutbox,omitempty"`
	UpdateChatRemovedFromList                      *UpdateChatRemovedFromList                      `json:"updateChatRemovedFromList,omitempty"`
	UpdateChatReplyMarkup                          *UpdateChatReplyMarkup                          `json:"updateChatReplyMarkup,omitempty"`
	UpdateChatRevenueAmount                        *UpdateChatRevenueAmount                        `json:"updateChatRevenueAmount,omitempty"`
	UpdateChatTheme                                *UpdateChatTheme                                `json:"updateChatTheme,omitempty"`
	UpdateChatTitle                                *UpdateChatTitle                                `json:"updateChatTitle,omitempty"`
	UpdateChatUnreadMentionCount                   *UpdateChatUnreadMentionCount                   `json:"updateChatUnreadMentionCount,omitempty"`
	UpdateChatUnreadReactionCount                  *UpdateChatUnreadReactionCount                  `json:"updateChatUnreadReactionCount,omitempty"`
	UpdateChatVideoChat                            *UpdateChatVideoChat                            `json:"updateChatVideoChat,omitempty"`
	UpdateChatViewAsTopics                         *UpdateChatViewAsTopics                         `json:"updateChatViewAsTopics,omitempty"`
	UpdateConnectionState                          *UpdateConnectionState                          `json:"updateConnectionState,omitempty"`
	UpdateContactCloseBirthdays                    *UpdateContactCloseBirthdays                    `json:"updateContactCloseBirthdays,omitempty"`
	UpdateDefaultBackground                        *UpdateDefaultBackground                        `json:"updateDefaultBackground,omitempty"`
	UpdateDefaultPaidReactionType                  *UpdateDefaultPaidReactionType                  `json:"updateDefaultPaidReactionType,omitempty"`
	UpdateDefaultReactionType                      *UpdateDefaultReactionType                      `json:"updateDefaultReactionType,omitempty"`
	UpdateDeleteMessages                           *UpdateDeleteMessages                           `json:"updateDeleteMessages,omitempty"`
	UpdateDiceEmojis                               *UpdateDiceEmojis                               `json:"updateDiceEmojis,omitempty"`
	UpdateDirectMessagesChatTopic                  *UpdateDirectMessagesChatTopic                  `json:"updateDirectMessagesChatTopic,omitempty"`
	UpdateEmojiChatThemes                          *UpdateEmojiChatThemes                          `json:"updateEmojiChatThemes,omitempty"`
	UpdateFavoriteStickers                         *UpdateFavoriteStickers                         `json:"updateFavoriteStickers,omitempty"`
	UpdateFile                                     *UpdateFile                                     `json:"updateFile,omitempty"`
	UpdateFileAddedToDownloads                     *UpdateFileAddedToDownloads                     `json:"updateFileAddedToDownloads,omitempty"`
	UpdateFileDownload                             *UpdateFileDownload                             `json:"updateFileDownload,omitempty"`
	UpdateFileDownloads                            *UpdateFileDownloads                            `json:"updateFileDownloads,omitempty"`
	UpdateFileGenerationStart                      *UpdateFileGenerationStart                      `json:"updateFileGenerationStart,omitempty"`
	UpdateFileGenerationStop                       *UpdateFileGenerationStop                       `json:"updateFileGenerationStop,omitempty"`
	UpdateFileRemovedFromDownloads                 *UpdateFileRemovedFromDownloads                 `json:"updateFileRemovedFromDownloads,omitempty"`
	UpdateForumTopic                               *UpdateForumTopic                               `json:"updateForumTopic,omitempty"`
	UpdateForumTopicInfo                           *UpdateForumTopicInfo                           `json:"updateForumTopicInfo,omitempty"`
	UpdateFreezeState                              *UpdateFreezeState                              `json:"updateFreezeState,omitempty"`
	UpdateGiftAuctionState                         *UpdateGiftAuctionState                         `json:"updateGiftAuctionState,omitempty"`
	UpdateGroupCall                                *UpdateGroupCall                                `json:"updateGroupCall,omitempty"`
	UpdateGroupCallMessageLevels                   *UpdateGroupCallMessageLevels                   `json:"updateGroupCallMessageLevels,omitempty"`
	UpdateGroupCallMessageSendFailed               *UpdateGroupCallMessageSendFailed               `json:"updateGroupCallMessageSendFailed,omitempty"`
	UpdateGroupCallMessagesDeleted                 *UpdateGroupCallMessagesDeleted                 `json:"updateGroupCallMessagesDeleted,omitempty"`
	UpdateGroupCallParticipant                     *UpdateGroupCallParticipant                     `json:"updateGroupCallParticipant,omitempty"`
	UpdateGroupCallParticipants                    *UpdateGroupCallParticipants                    `json:"updateGroupCallParticipants,omitempty"`
	UpdateGroupCallVerificationState               *UpdateGroupCallVerificationState               `json:"updateGroupCallVerificationState,omitempty"`
	UpdateHavePendingNotifications                 *UpdateHavePendingNotifications                 `json:"updateHavePendingNotifications,omitempty"`
	UpdateInstalledStickerSets                     *UpdateInstalledStickerSets                     `json:"updateInstalledStickerSets,omitempty"`
	UpdateLanguagePackStrings                      *UpdateLanguagePackStrings                      `json:"updateLanguagePackStrings,omitempty"`
	UpdateLiveStoryTopDonors                       *UpdateLiveStoryTopDonors                       `json:"updateLiveStoryTopDonors,omitempty"`
	UpdateMessageContent                           *UpdateMessageContent                           `json:"updateMessageContent,omitempty"`
	UpdateMessageContentOpened                     *UpdateMessageContentOpened                     `json:"updateMessageContentOpened,omitempty"`
	UpdateMessageEdited                            *UpdateMessageEdited                            `json:"updateMessageEdited,omitempty"`
	UpdateMessageFactCheck                         *UpdateMessageFactCheck                         `json:"updateMessageFactCheck,omitempty"`
	UpdateMessageInteractionInfo                   *UpdateMessageInteractionInfo                   `json:"updateMessageInteractionInfo,omitempty"`
	UpdateMessageIsPinned                          *UpdateMessageIsPinned                          `json:"updateMessageIsPinned,omitempty"`
	UpdateMessageLiveLocationViewed                *UpdateMessageLiveLocationViewed                `json:"updateMessageLiveLocationViewed,omitempty"`
	UpdateMessageMentionRead                       *UpdateMessageMentionRead                       `json:"updateMessageMentionRead,omitempty"`
	UpdateMessageReaction                          *UpdateMessageReaction                          `json:"updateMessageReaction,omitempty"`
	UpdateMessageReactions                         *UpdateMessageReactions                         `json:"updateMessageReactions,omitempty"`
	UpdateMessageSendAcknowledged                  *UpdateMessageSendAcknowledged                  `json:"updateMessageSendAcknowledged,omitempty"`
	UpdateMessageSendFailed                        *UpdateMessageSendFailed                        `json:"updateMessageSendFailed,omitempty"`
	UpdateMessageSendSucceeded                     *UpdateMessageSendSucceeded                     `json:"updateMessageSendSucceeded,omitempty"`
	UpdateMessageSuggestedPostInfo                 *UpdateMessageSuggestedPostInfo                 `json:"updateMessageSuggestedPostInfo,omitempty"`
	UpdateMessageUnreadReactions                   *UpdateMessageUnreadReactions                   `json:"updateMessageUnreadReactions,omitempty"`
	UpdateNewBusinessCallbackQuery                 *UpdateNewBusinessCallbackQuery                 `json:"updateNewBusinessCallbackQuery,omitempty"`
	UpdateNewBusinessMessage                       *UpdateNewBusinessMessage                       `json:"updateNewBusinessMessage,omitempty"`
	UpdateNewCallSignalingData                     *UpdateNewCallSignalingData                     `json:"updateNewCallSignalingData,omitempty"`
	UpdateNewCallbackQuery                         *UpdateNewCallbackQuery                         `json:"updateNewCallbackQuery,omitempty"`
	UpdateNewChat                                  *UpdateNewChat                                  `json:"updateNewChat,omitempty"`
	UpdateNewChatJoinRequest                       *UpdateNewChatJoinRequest                       `json:"updateNewChatJoinRequest,omitempty"`
	UpdateNewChosenInlineResult                    *UpdateNewChosenInlineResult                    `json:"updateNewChosenInlineResult,omitempty"`
	UpdateNewCustomEvent                           *UpdateNewCustomEvent                           `json:"updateNewCustomEvent,omitempty"`
	UpdateNewCustomQuery                           *UpdateNewCustomQuery                           `json:"updateNewCustomQuery,omitempty"`
	UpdateNewGroupCallMessage                      *UpdateNewGroupCallMessage                      `json:"updateNewGroupCallMessage,omitempty"`
	UpdateNewGroupCallPaidReaction                 *UpdateNewGroupCallPaidReaction                 `json:"updateNewGroupCallPaidReaction,omitempty"`
	UpdateNewInlineCallbackQuery                   *UpdateNewInlineCallbackQuery                   `json:"updateNewInlineCallbackQuery,omitempty"`
	UpdateNewInlineQuery                           *UpdateNewInlineQuery                           `json:"updateNewInlineQuery,omitempty"`
	UpdateNewMessage                               *UpdateNewMessage                               `json:"updateNewMessage,omitempty"`
	UpdateNewPreCheckoutQuery                      *UpdateNewPreCheckoutQuery                      `json:"updateNewPreCheckoutQuery,omitempty"`
	UpdateNewShippingQuery                         *UpdateNewShippingQuery                         `json:"updateNewShippingQuery,omitempty"`
	UpdateNotification                             *UpdateNotification                             `json:"updateNotification,omitempty"`
	UpdateNotificationGroup                        *UpdateNotificationGroup                        `json:"updateNotificationGroup,omitempty"`
	UpdateOption                                   *UpdateOption                                   `json:"updateOption,omitempty"`
	UpdateOwnedStarCount                           *UpdateOwnedStarCount                           `json:"updateOwnedStarCount,omitempty"`
	UpdateOwnedTonCount                            *UpdateOwnedTonCount                            `json:"updateOwnedTonCount,omitempty"`
	UpdatePaidMediaPurchased                       *UpdatePaidMediaPurchased                       `json:"updatePaidMediaPurchased,omitempty"`
	UpdatePendingTextMessage                       *UpdatePendingTextMessage                       `json:"updatePendingTextMessage,omitempty"`
	UpdatePoll                                     *UpdatePoll                                     `json:"updatePoll,omitempty"`
	UpdatePollAnswer                               *UpdatePollAnswer                               `json:"updatePollAnswer,omitempty"`
	UpdateProfileAccentColors                      *UpdateProfileAccentColors                      `json:"updateProfileAccentColors,omitempty"`
	UpdateQuickReplyShortcut                       *UpdateQuickReplyShortcut                       `json:"updateQuickReplyShortcut,omitempty"`
	UpdateQuickReplyShortcutDeleted                *UpdateQuickReplyShortcutDeleted                `json:"updateQuickReplyShortcutDeleted,omitempty"`
	UpdateQuickReplyShortcutMessages               *UpdateQuickReplyShortcutMessages               `json:"updateQuickReplyShortcutMessages,omitempty"`
	UpdateQuickReplyShortcuts                      *UpdateQuickReplyShortcuts                      `json:"updateQuickReplyShortcuts,omitempty"`
	UpdateReactionNotificationSettings             *UpdateReactionNotificationSettings             `json:"updateReactionNotificationSettings,omitempty"`
	UpdateRecentStickers                           *UpdateRecentStickers                           `json:"updateRecentStickers,omitempty"`
	UpdateSavedAnimations                          *UpdateSavedAnimations                          `json:"updateSavedAnimations,omitempty"`
	UpdateSavedMessagesTags                        *UpdateSavedMessagesTags                        `json:"updateSavedMessagesTags,omitempty"`
	UpdateSavedMessagesTopic                       *UpdateSavedMessagesTopic                       `json:"updateSavedMessagesTopic,omitempty"`
	UpdateSavedMessagesTopicCount                  *UpdateSavedMessagesTopicCount                  `json:"updateSavedMessagesTopicCount,omitempty"`
	UpdateSavedNotificationSounds                  *UpdateSavedNotificationSounds                  `json:"updateSavedNotificationSounds,omitempty"`
	UpdateScopeNotificationSettings                *UpdateScopeNotificationSettings                `json:"updateScopeNotificationSettings,omitempty"`
	UpdateSecretChat                               *UpdateSecretChat                               `json:"updateSecretChat,omitempty"`
	UpdateServiceNotification                      *UpdateServiceNotification                      `json:"updateServiceNotification,omitempty"`
	UpdateSpeechRecognitionTrial                   *UpdateSpeechRecognitionTrial                   `json:"updateSpeechRecognitionTrial,omitempty"`
	UpdateSpeedLimitNotification                   *UpdateSpeedLimitNotification                   `json:"updateSpeedLimitNotification,omitempty"`
	UpdateStakeDiceState                           *UpdateStakeDiceState                           `json:"updateStakeDiceState,omitempty"`
	UpdateStarRevenueStatus                        *UpdateStarRevenueStatus                        `json:"updateStarRevenueStatus,omitempty"`
	UpdateStickerSet                               *UpdateStickerSet                               `json:"updateStickerSet,omitempty"`
	UpdateStory                                    *UpdateStory                                    `json:"updateStory,omitempty"`
	UpdateStoryDeleted                             *UpdateStoryDeleted                             `json:"updateStoryDeleted,omitempty"`
	UpdateStoryListChatCount                       *UpdateStoryListChatCount                       `json:"updateStoryListChatCount,omitempty"`
	UpdateStoryPostFailed                          *UpdateStoryPostFailed                          `json:"updateStoryPostFailed,omitempty"`
	UpdateStoryPostSucceeded                       *UpdateStoryPostSucceeded                       `json:"updateStoryPostSucceeded,omitempty"`
	UpdateStoryStealthMode                         *UpdateStoryStealthMode                         `json:"updateStoryStealthMode,omitempty"`
	UpdateSuggestedActions                         *UpdateSuggestedActions                         `json:"updateSuggestedActions,omitempty"`
	UpdateSupergroup                               *UpdateSupergroup                               `json:"updateSupergroup,omitempty"`
	UpdateSupergroupFullInfo                       *UpdateSupergroupFullInfo                       `json:"updateSupergroupFullInfo,omitempty"`
	UpdateTermsOfService                           *UpdateTermsOfService                           `json:"updateTermsOfService,omitempty"`
	UpdateTonRevenueStatus                         *UpdateTonRevenueStatus                         `json:"updateTonRevenueStatus,omitempty"`
	UpdateTopicMessageCount                        *UpdateTopicMessageCount                        `json:"updateTopicMessageCount,omitempty"`
	UpdateTrendingStickerSets                      *UpdateTrendingStickerSets                      `json:"updateTrendingStickerSets,omitempty"`
	UpdateTrustedMiniAppBots                       *UpdateTrustedMiniAppBots                       `json:"updateTrustedMiniAppBots,omitempty"`
	UpdateUnconfirmedSession                       *UpdateUnconfirmedSession                       `json:"updateUnconfirmedSession,omitempty"`
	UpdateUnreadChatCount                          *UpdateUnreadChatCount                          `json:"updateUnreadChatCount,omitempty"`
	UpdateUnreadMessageCount                       *UpdateUnreadMessageCount                       `json:"updateUnreadMessageCount,omitempty"`
	UpdateUser                                     *UpdateUser                                     `json:"updateUser,omitempty"`
	UpdateUserFullInfo                             *UpdateUserFullInfo                             `json:"updateUserFullInfo,omitempty"`
	UpdateUserPrivacySettingRules                  *UpdateUserPrivacySettingRules                  `json:"updateUserPrivacySettingRules,omitempty"`
	UpdateUserStatus                               *UpdateUserStatus                               `json:"updateUserStatus,omitempty"`
	UpdateVideoPublished                           *UpdateVideoPublished                           `json:"updateVideoPublished,omitempty"`
	UpdateWebAppMessageSent                        *UpdateWebAppMessageSent                        `json:"updateWebAppMessageSent,omitempty"`
}

func (t *Update) Type() string {
	return t.TypeStr
}

func (t *Update) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *Update) GetExtra() string {
	return ""
}

func (t *Update) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "updateAccentColors":
		t.UpdateAccentColors = new(UpdateAccentColors)
		return json.Unmarshal(b, t.UpdateAccentColors)
	case "updateActiveEmojiReactions":
		t.UpdateActiveEmojiReactions = new(UpdateActiveEmojiReactions)
		return json.Unmarshal(b, t.UpdateActiveEmojiReactions)
	case "updateActiveGiftAuctions":
		t.UpdateActiveGiftAuctions = new(UpdateActiveGiftAuctions)
		return json.Unmarshal(b, t.UpdateActiveGiftAuctions)
	case "updateActiveLiveLocationMessages":
		t.UpdateActiveLiveLocationMessages = new(UpdateActiveLiveLocationMessages)
		return json.Unmarshal(b, t.UpdateActiveLiveLocationMessages)
	case "updateActiveNotifications":
		t.UpdateActiveNotifications = new(UpdateActiveNotifications)
		return json.Unmarshal(b, t.UpdateActiveNotifications)
	case "updateAgeVerificationParameters":
		t.UpdateAgeVerificationParameters = new(UpdateAgeVerificationParameters)
		return json.Unmarshal(b, t.UpdateAgeVerificationParameters)
	case "updateAnimatedEmojiMessageClicked":
		t.UpdateAnimatedEmojiMessageClicked = new(UpdateAnimatedEmojiMessageClicked)
		return json.Unmarshal(b, t.UpdateAnimatedEmojiMessageClicked)
	case "updateAnimationSearchParameters":
		t.UpdateAnimationSearchParameters = new(UpdateAnimationSearchParameters)
		return json.Unmarshal(b, t.UpdateAnimationSearchParameters)
	case "updateApplicationRecaptchaVerificationRequired":
		t.UpdateApplicationRecaptchaVerificationRequired = new(UpdateApplicationRecaptchaVerificationRequired)
		return json.Unmarshal(b, t.UpdateApplicationRecaptchaVerificationRequired)
	case "updateApplicationVerificationRequired":
		t.UpdateApplicationVerificationRequired = new(UpdateApplicationVerificationRequired)
		return json.Unmarshal(b, t.UpdateApplicationVerificationRequired)
	case "updateAttachmentMenuBots":
		t.UpdateAttachmentMenuBots = new(UpdateAttachmentMenuBots)
		return json.Unmarshal(b, t.UpdateAttachmentMenuBots)
	case "updateAuthorizationState":
		t.UpdateAuthorizationState = new(UpdateAuthorizationState)
		return json.Unmarshal(b, t.UpdateAuthorizationState)
	case "updateAutosaveSettings":
		t.UpdateAutosaveSettings = new(UpdateAutosaveSettings)
		return json.Unmarshal(b, t.UpdateAutosaveSettings)
	case "updateAvailableMessageEffects":
		t.UpdateAvailableMessageEffects = new(UpdateAvailableMessageEffects)
		return json.Unmarshal(b, t.UpdateAvailableMessageEffects)
	case "updateBasicGroup":
		t.UpdateBasicGroup = new(UpdateBasicGroup)
		return json.Unmarshal(b, t.UpdateBasicGroup)
	case "updateBasicGroupFullInfo":
		t.UpdateBasicGroupFullInfo = new(UpdateBasicGroupFullInfo)
		return json.Unmarshal(b, t.UpdateBasicGroupFullInfo)
	case "updateBusinessConnection":
		t.UpdateBusinessConnection = new(UpdateBusinessConnection)
		return json.Unmarshal(b, t.UpdateBusinessConnection)
	case "updateBusinessMessageEdited":
		t.UpdateBusinessMessageEdited = new(UpdateBusinessMessageEdited)
		return json.Unmarshal(b, t.UpdateBusinessMessageEdited)
	case "updateBusinessMessagesDeleted":
		t.UpdateBusinessMessagesDeleted = new(UpdateBusinessMessagesDeleted)
		return json.Unmarshal(b, t.UpdateBusinessMessagesDeleted)
	case "updateCall":
		t.UpdateCall = new(UpdateCall)
		return json.Unmarshal(b, t.UpdateCall)
	case "updateChatAccentColors":
		t.UpdateChatAccentColors = new(UpdateChatAccentColors)
		return json.Unmarshal(b, t.UpdateChatAccentColors)
	case "updateChatAction":
		t.UpdateChatAction = new(UpdateChatAction)
		return json.Unmarshal(b, t.UpdateChatAction)
	case "updateChatActionBar":
		t.UpdateChatActionBar = new(UpdateChatActionBar)
		return json.Unmarshal(b, t.UpdateChatActionBar)
	case "updateChatActiveStories":
		t.UpdateChatActiveStories = new(UpdateChatActiveStories)
		return json.Unmarshal(b, t.UpdateChatActiveStories)
	case "updateChatAddedToList":
		t.UpdateChatAddedToList = new(UpdateChatAddedToList)
		return json.Unmarshal(b, t.UpdateChatAddedToList)
	case "updateChatAvailableReactions":
		t.UpdateChatAvailableReactions = new(UpdateChatAvailableReactions)
		return json.Unmarshal(b, t.UpdateChatAvailableReactions)
	case "updateChatBackground":
		t.UpdateChatBackground = new(UpdateChatBackground)
		return json.Unmarshal(b, t.UpdateChatBackground)
	case "updateChatBlockList":
		t.UpdateChatBlockList = new(UpdateChatBlockList)
		return json.Unmarshal(b, t.UpdateChatBlockList)
	case "updateChatBoost":
		t.UpdateChatBoost = new(UpdateChatBoost)
		return json.Unmarshal(b, t.UpdateChatBoost)
	case "updateChatBusinessBotManageBar":
		t.UpdateChatBusinessBotManageBar = new(UpdateChatBusinessBotManageBar)
		return json.Unmarshal(b, t.UpdateChatBusinessBotManageBar)
	case "updateChatDefaultDisableNotification":
		t.UpdateChatDefaultDisableNotification = new(UpdateChatDefaultDisableNotification)
		return json.Unmarshal(b, t.UpdateChatDefaultDisableNotification)
	case "updateChatDraftMessage":
		t.UpdateChatDraftMessage = new(UpdateChatDraftMessage)
		return json.Unmarshal(b, t.UpdateChatDraftMessage)
	case "updateChatEmojiStatus":
		t.UpdateChatEmojiStatus = new(UpdateChatEmojiStatus)
		return json.Unmarshal(b, t.UpdateChatEmojiStatus)
	case "updateChatFolders":
		t.UpdateChatFolders = new(UpdateChatFolders)
		return json.Unmarshal(b, t.UpdateChatFolders)
	case "updateChatHasProtectedContent":
		t.UpdateChatHasProtectedContent = new(UpdateChatHasProtectedContent)
		return json.Unmarshal(b, t.UpdateChatHasProtectedContent)
	case "updateChatHasScheduledMessages":
		t.UpdateChatHasScheduledMessages = new(UpdateChatHasScheduledMessages)
		return json.Unmarshal(b, t.UpdateChatHasScheduledMessages)
	case "updateChatIsMarkedAsUnread":
		t.UpdateChatIsMarkedAsUnread = new(UpdateChatIsMarkedAsUnread)
		return json.Unmarshal(b, t.UpdateChatIsMarkedAsUnread)
	case "updateChatIsTranslatable":
		t.UpdateChatIsTranslatable = new(UpdateChatIsTranslatable)
		return json.Unmarshal(b, t.UpdateChatIsTranslatable)
	case "updateChatLastMessage":
		t.UpdateChatLastMessage = new(UpdateChatLastMessage)
		return json.Unmarshal(b, t.UpdateChatLastMessage)
	case "updateChatMember":
		t.UpdateChatMember = new(UpdateChatMember)
		return json.Unmarshal(b, t.UpdateChatMember)
	case "updateChatMessageAutoDeleteTime":
		t.UpdateChatMessageAutoDeleteTime = new(UpdateChatMessageAutoDeleteTime)
		return json.Unmarshal(b, t.UpdateChatMessageAutoDeleteTime)
	case "updateChatMessageSender":
		t.UpdateChatMessageSender = new(UpdateChatMessageSender)
		return json.Unmarshal(b, t.UpdateChatMessageSender)
	case "updateChatNotificationSettings":
		t.UpdateChatNotificationSettings = new(UpdateChatNotificationSettings)
		return json.Unmarshal(b, t.UpdateChatNotificationSettings)
	case "updateChatOnlineMemberCount":
		t.UpdateChatOnlineMemberCount = new(UpdateChatOnlineMemberCount)
		return json.Unmarshal(b, t.UpdateChatOnlineMemberCount)
	case "updateChatPendingJoinRequests":
		t.UpdateChatPendingJoinRequests = new(UpdateChatPendingJoinRequests)
		return json.Unmarshal(b, t.UpdateChatPendingJoinRequests)
	case "updateChatPermissions":
		t.UpdateChatPermissions = new(UpdateChatPermissions)
		return json.Unmarshal(b, t.UpdateChatPermissions)
	case "updateChatPhoto":
		t.UpdateChatPhoto = new(UpdateChatPhoto)
		return json.Unmarshal(b, t.UpdateChatPhoto)
	case "updateChatPosition":
		t.UpdateChatPosition = new(UpdateChatPosition)
		return json.Unmarshal(b, t.UpdateChatPosition)
	case "updateChatReadInbox":
		t.UpdateChatReadInbox = new(UpdateChatReadInbox)
		return json.Unmarshal(b, t.UpdateChatReadInbox)
	case "updateChatReadOutbox":
		t.UpdateChatReadOutbox = new(UpdateChatReadOutbox)
		return json.Unmarshal(b, t.UpdateChatReadOutbox)
	case "updateChatRemovedFromList":
		t.UpdateChatRemovedFromList = new(UpdateChatRemovedFromList)
		return json.Unmarshal(b, t.UpdateChatRemovedFromList)
	case "updateChatReplyMarkup":
		t.UpdateChatReplyMarkup = new(UpdateChatReplyMarkup)
		return json.Unmarshal(b, t.UpdateChatReplyMarkup)
	case "updateChatRevenueAmount":
		t.UpdateChatRevenueAmount = new(UpdateChatRevenueAmount)
		return json.Unmarshal(b, t.UpdateChatRevenueAmount)
	case "updateChatTheme":
		t.UpdateChatTheme = new(UpdateChatTheme)
		return json.Unmarshal(b, t.UpdateChatTheme)
	case "updateChatTitle":
		t.UpdateChatTitle = new(UpdateChatTitle)
		return json.Unmarshal(b, t.UpdateChatTitle)
	case "updateChatUnreadMentionCount":
		t.UpdateChatUnreadMentionCount = new(UpdateChatUnreadMentionCount)
		return json.Unmarshal(b, t.UpdateChatUnreadMentionCount)
	case "updateChatUnreadReactionCount":
		t.UpdateChatUnreadReactionCount = new(UpdateChatUnreadReactionCount)
		return json.Unmarshal(b, t.UpdateChatUnreadReactionCount)
	case "updateChatVideoChat":
		t.UpdateChatVideoChat = new(UpdateChatVideoChat)
		return json.Unmarshal(b, t.UpdateChatVideoChat)
	case "updateChatViewAsTopics":
		t.UpdateChatViewAsTopics = new(UpdateChatViewAsTopics)
		return json.Unmarshal(b, t.UpdateChatViewAsTopics)
	case "updateConnectionState":
		t.UpdateConnectionState = new(UpdateConnectionState)
		return json.Unmarshal(b, t.UpdateConnectionState)
	case "updateContactCloseBirthdays":
		t.UpdateContactCloseBirthdays = new(UpdateContactCloseBirthdays)
		return json.Unmarshal(b, t.UpdateContactCloseBirthdays)
	case "updateDefaultBackground":
		t.UpdateDefaultBackground = new(UpdateDefaultBackground)
		return json.Unmarshal(b, t.UpdateDefaultBackground)
	case "updateDefaultPaidReactionType":
		t.UpdateDefaultPaidReactionType = new(UpdateDefaultPaidReactionType)
		return json.Unmarshal(b, t.UpdateDefaultPaidReactionType)
	case "updateDefaultReactionType":
		t.UpdateDefaultReactionType = new(UpdateDefaultReactionType)
		return json.Unmarshal(b, t.UpdateDefaultReactionType)
	case "updateDeleteMessages":
		t.UpdateDeleteMessages = new(UpdateDeleteMessages)
		return json.Unmarshal(b, t.UpdateDeleteMessages)
	case "updateDiceEmojis":
		t.UpdateDiceEmojis = new(UpdateDiceEmojis)
		return json.Unmarshal(b, t.UpdateDiceEmojis)
	case "updateDirectMessagesChatTopic":
		t.UpdateDirectMessagesChatTopic = new(UpdateDirectMessagesChatTopic)
		return json.Unmarshal(b, t.UpdateDirectMessagesChatTopic)
	case "updateEmojiChatThemes":
		t.UpdateEmojiChatThemes = new(UpdateEmojiChatThemes)
		return json.Unmarshal(b, t.UpdateEmojiChatThemes)
	case "updateFavoriteStickers":
		t.UpdateFavoriteStickers = new(UpdateFavoriteStickers)
		return json.Unmarshal(b, t.UpdateFavoriteStickers)
	case "updateFile":
		t.UpdateFile = new(UpdateFile)
		return json.Unmarshal(b, t.UpdateFile)
	case "updateFileAddedToDownloads":
		t.UpdateFileAddedToDownloads = new(UpdateFileAddedToDownloads)
		return json.Unmarshal(b, t.UpdateFileAddedToDownloads)
	case "updateFileDownload":
		t.UpdateFileDownload = new(UpdateFileDownload)
		return json.Unmarshal(b, t.UpdateFileDownload)
	case "updateFileDownloads":
		t.UpdateFileDownloads = new(UpdateFileDownloads)
		return json.Unmarshal(b, t.UpdateFileDownloads)
	case "updateFileGenerationStart":
		t.UpdateFileGenerationStart = new(UpdateFileGenerationStart)
		return json.Unmarshal(b, t.UpdateFileGenerationStart)
	case "updateFileGenerationStop":
		t.UpdateFileGenerationStop = new(UpdateFileGenerationStop)
		return json.Unmarshal(b, t.UpdateFileGenerationStop)
	case "updateFileRemovedFromDownloads":
		t.UpdateFileRemovedFromDownloads = new(UpdateFileRemovedFromDownloads)
		return json.Unmarshal(b, t.UpdateFileRemovedFromDownloads)
	case "updateForumTopic":
		t.UpdateForumTopic = new(UpdateForumTopic)
		return json.Unmarshal(b, t.UpdateForumTopic)
	case "updateForumTopicInfo":
		t.UpdateForumTopicInfo = new(UpdateForumTopicInfo)
		return json.Unmarshal(b, t.UpdateForumTopicInfo)
	case "updateFreezeState":
		t.UpdateFreezeState = new(UpdateFreezeState)
		return json.Unmarshal(b, t.UpdateFreezeState)
	case "updateGiftAuctionState":
		t.UpdateGiftAuctionState = new(UpdateGiftAuctionState)
		return json.Unmarshal(b, t.UpdateGiftAuctionState)
	case "updateGroupCall":
		t.UpdateGroupCall = new(UpdateGroupCall)
		return json.Unmarshal(b, t.UpdateGroupCall)
	case "updateGroupCallMessageLevels":
		t.UpdateGroupCallMessageLevels = new(UpdateGroupCallMessageLevels)
		return json.Unmarshal(b, t.UpdateGroupCallMessageLevels)
	case "updateGroupCallMessageSendFailed":
		t.UpdateGroupCallMessageSendFailed = new(UpdateGroupCallMessageSendFailed)
		return json.Unmarshal(b, t.UpdateGroupCallMessageSendFailed)
	case "updateGroupCallMessagesDeleted":
		t.UpdateGroupCallMessagesDeleted = new(UpdateGroupCallMessagesDeleted)
		return json.Unmarshal(b, t.UpdateGroupCallMessagesDeleted)
	case "updateGroupCallParticipant":
		t.UpdateGroupCallParticipant = new(UpdateGroupCallParticipant)
		return json.Unmarshal(b, t.UpdateGroupCallParticipant)
	case "updateGroupCallParticipants":
		t.UpdateGroupCallParticipants = new(UpdateGroupCallParticipants)
		return json.Unmarshal(b, t.UpdateGroupCallParticipants)
	case "updateGroupCallVerificationState":
		t.UpdateGroupCallVerificationState = new(UpdateGroupCallVerificationState)
		return json.Unmarshal(b, t.UpdateGroupCallVerificationState)
	case "updateHavePendingNotifications":
		t.UpdateHavePendingNotifications = new(UpdateHavePendingNotifications)
		return json.Unmarshal(b, t.UpdateHavePendingNotifications)
	case "updateInstalledStickerSets":
		t.UpdateInstalledStickerSets = new(UpdateInstalledStickerSets)
		return json.Unmarshal(b, t.UpdateInstalledStickerSets)
	case "updateLanguagePackStrings":
		t.UpdateLanguagePackStrings = new(UpdateLanguagePackStrings)
		return json.Unmarshal(b, t.UpdateLanguagePackStrings)
	case "updateLiveStoryTopDonors":
		t.UpdateLiveStoryTopDonors = new(UpdateLiveStoryTopDonors)
		return json.Unmarshal(b, t.UpdateLiveStoryTopDonors)
	case "updateMessageContent":
		t.UpdateMessageContent = new(UpdateMessageContent)
		return json.Unmarshal(b, t.UpdateMessageContent)
	case "updateMessageContentOpened":
		t.UpdateMessageContentOpened = new(UpdateMessageContentOpened)
		return json.Unmarshal(b, t.UpdateMessageContentOpened)
	case "updateMessageEdited":
		t.UpdateMessageEdited = new(UpdateMessageEdited)
		return json.Unmarshal(b, t.UpdateMessageEdited)
	case "updateMessageFactCheck":
		t.UpdateMessageFactCheck = new(UpdateMessageFactCheck)
		return json.Unmarshal(b, t.UpdateMessageFactCheck)
	case "updateMessageInteractionInfo":
		t.UpdateMessageInteractionInfo = new(UpdateMessageInteractionInfo)
		return json.Unmarshal(b, t.UpdateMessageInteractionInfo)
	case "updateMessageIsPinned":
		t.UpdateMessageIsPinned = new(UpdateMessageIsPinned)
		return json.Unmarshal(b, t.UpdateMessageIsPinned)
	case "updateMessageLiveLocationViewed":
		t.UpdateMessageLiveLocationViewed = new(UpdateMessageLiveLocationViewed)
		return json.Unmarshal(b, t.UpdateMessageLiveLocationViewed)
	case "updateMessageMentionRead":
		t.UpdateMessageMentionRead = new(UpdateMessageMentionRead)
		return json.Unmarshal(b, t.UpdateMessageMentionRead)
	case "updateMessageReaction":
		t.UpdateMessageReaction = new(UpdateMessageReaction)
		return json.Unmarshal(b, t.UpdateMessageReaction)
	case "updateMessageReactions":
		t.UpdateMessageReactions = new(UpdateMessageReactions)
		return json.Unmarshal(b, t.UpdateMessageReactions)
	case "updateMessageSendAcknowledged":
		t.UpdateMessageSendAcknowledged = new(UpdateMessageSendAcknowledged)
		return json.Unmarshal(b, t.UpdateMessageSendAcknowledged)
	case "updateMessageSendFailed":
		t.UpdateMessageSendFailed = new(UpdateMessageSendFailed)
		return json.Unmarshal(b, t.UpdateMessageSendFailed)
	case "updateMessageSendSucceeded":
		t.UpdateMessageSendSucceeded = new(UpdateMessageSendSucceeded)
		return json.Unmarshal(b, t.UpdateMessageSendSucceeded)
	case "updateMessageSuggestedPostInfo":
		t.UpdateMessageSuggestedPostInfo = new(UpdateMessageSuggestedPostInfo)
		return json.Unmarshal(b, t.UpdateMessageSuggestedPostInfo)
	case "updateMessageUnreadReactions":
		t.UpdateMessageUnreadReactions = new(UpdateMessageUnreadReactions)
		return json.Unmarshal(b, t.UpdateMessageUnreadReactions)
	case "updateNewBusinessCallbackQuery":
		t.UpdateNewBusinessCallbackQuery = new(UpdateNewBusinessCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewBusinessCallbackQuery)
	case "updateNewBusinessMessage":
		t.UpdateNewBusinessMessage = new(UpdateNewBusinessMessage)
		return json.Unmarshal(b, t.UpdateNewBusinessMessage)
	case "updateNewCallSignalingData":
		t.UpdateNewCallSignalingData = new(UpdateNewCallSignalingData)
		return json.Unmarshal(b, t.UpdateNewCallSignalingData)
	case "updateNewCallbackQuery":
		t.UpdateNewCallbackQuery = new(UpdateNewCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewCallbackQuery)
	case "updateNewChat":
		t.UpdateNewChat = new(UpdateNewChat)
		return json.Unmarshal(b, t.UpdateNewChat)
	case "updateNewChatJoinRequest":
		t.UpdateNewChatJoinRequest = new(UpdateNewChatJoinRequest)
		return json.Unmarshal(b, t.UpdateNewChatJoinRequest)
	case "updateNewChosenInlineResult":
		t.UpdateNewChosenInlineResult = new(UpdateNewChosenInlineResult)
		return json.Unmarshal(b, t.UpdateNewChosenInlineResult)
	case "updateNewCustomEvent":
		t.UpdateNewCustomEvent = new(UpdateNewCustomEvent)
		return json.Unmarshal(b, t.UpdateNewCustomEvent)
	case "updateNewCustomQuery":
		t.UpdateNewCustomQuery = new(UpdateNewCustomQuery)
		return json.Unmarshal(b, t.UpdateNewCustomQuery)
	case "updateNewGroupCallMessage":
		t.UpdateNewGroupCallMessage = new(UpdateNewGroupCallMessage)
		return json.Unmarshal(b, t.UpdateNewGroupCallMessage)
	case "updateNewGroupCallPaidReaction":
		t.UpdateNewGroupCallPaidReaction = new(UpdateNewGroupCallPaidReaction)
		return json.Unmarshal(b, t.UpdateNewGroupCallPaidReaction)
	case "updateNewInlineCallbackQuery":
		t.UpdateNewInlineCallbackQuery = new(UpdateNewInlineCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewInlineCallbackQuery)
	case "updateNewInlineQuery":
		t.UpdateNewInlineQuery = new(UpdateNewInlineQuery)
		return json.Unmarshal(b, t.UpdateNewInlineQuery)
	case "updateNewMessage":
		t.UpdateNewMessage = new(UpdateNewMessage)
		return json.Unmarshal(b, t.UpdateNewMessage)
	case "updateNewPreCheckoutQuery":
		t.UpdateNewPreCheckoutQuery = new(UpdateNewPreCheckoutQuery)
		return json.Unmarshal(b, t.UpdateNewPreCheckoutQuery)
	case "updateNewShippingQuery":
		t.UpdateNewShippingQuery = new(UpdateNewShippingQuery)
		return json.Unmarshal(b, t.UpdateNewShippingQuery)
	case "updateNotification":
		t.UpdateNotification = new(UpdateNotification)
		return json.Unmarshal(b, t.UpdateNotification)
	case "updateNotificationGroup":
		t.UpdateNotificationGroup = new(UpdateNotificationGroup)
		return json.Unmarshal(b, t.UpdateNotificationGroup)
	case "updateOption":
		t.UpdateOption = new(UpdateOption)
		return json.Unmarshal(b, t.UpdateOption)
	case "updateOwnedStarCount":
		t.UpdateOwnedStarCount = new(UpdateOwnedStarCount)
		return json.Unmarshal(b, t.UpdateOwnedStarCount)
	case "updateOwnedTonCount":
		t.UpdateOwnedTonCount = new(UpdateOwnedTonCount)
		return json.Unmarshal(b, t.UpdateOwnedTonCount)
	case "updatePaidMediaPurchased":
		t.UpdatePaidMediaPurchased = new(UpdatePaidMediaPurchased)
		return json.Unmarshal(b, t.UpdatePaidMediaPurchased)
	case "updatePendingTextMessage":
		t.UpdatePendingTextMessage = new(UpdatePendingTextMessage)
		return json.Unmarshal(b, t.UpdatePendingTextMessage)
	case "updatePoll":
		t.UpdatePoll = new(UpdatePoll)
		return json.Unmarshal(b, t.UpdatePoll)
	case "updatePollAnswer":
		t.UpdatePollAnswer = new(UpdatePollAnswer)
		return json.Unmarshal(b, t.UpdatePollAnswer)
	case "updateProfileAccentColors":
		t.UpdateProfileAccentColors = new(UpdateProfileAccentColors)
		return json.Unmarshal(b, t.UpdateProfileAccentColors)
	case "updateQuickReplyShortcut":
		t.UpdateQuickReplyShortcut = new(UpdateQuickReplyShortcut)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcut)
	case "updateQuickReplyShortcutDeleted":
		t.UpdateQuickReplyShortcutDeleted = new(UpdateQuickReplyShortcutDeleted)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcutDeleted)
	case "updateQuickReplyShortcutMessages":
		t.UpdateQuickReplyShortcutMessages = new(UpdateQuickReplyShortcutMessages)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcutMessages)
	case "updateQuickReplyShortcuts":
		t.UpdateQuickReplyShortcuts = new(UpdateQuickReplyShortcuts)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcuts)
	case "updateReactionNotificationSettings":
		t.UpdateReactionNotificationSettings = new(UpdateReactionNotificationSettings)
		return json.Unmarshal(b, t.UpdateReactionNotificationSettings)
	case "updateRecentStickers":
		t.UpdateRecentStickers = new(UpdateRecentStickers)
		return json.Unmarshal(b, t.UpdateRecentStickers)
	case "updateSavedAnimations":
		t.UpdateSavedAnimations = new(UpdateSavedAnimations)
		return json.Unmarshal(b, t.UpdateSavedAnimations)
	case "updateSavedMessagesTags":
		t.UpdateSavedMessagesTags = new(UpdateSavedMessagesTags)
		return json.Unmarshal(b, t.UpdateSavedMessagesTags)
	case "updateSavedMessagesTopic":
		t.UpdateSavedMessagesTopic = new(UpdateSavedMessagesTopic)
		return json.Unmarshal(b, t.UpdateSavedMessagesTopic)
	case "updateSavedMessagesTopicCount":
		t.UpdateSavedMessagesTopicCount = new(UpdateSavedMessagesTopicCount)
		return json.Unmarshal(b, t.UpdateSavedMessagesTopicCount)
	case "updateSavedNotificationSounds":
		t.UpdateSavedNotificationSounds = new(UpdateSavedNotificationSounds)
		return json.Unmarshal(b, t.UpdateSavedNotificationSounds)
	case "updateScopeNotificationSettings":
		t.UpdateScopeNotificationSettings = new(UpdateScopeNotificationSettings)
		return json.Unmarshal(b, t.UpdateScopeNotificationSettings)
	case "updateSecretChat":
		t.UpdateSecretChat = new(UpdateSecretChat)
		return json.Unmarshal(b, t.UpdateSecretChat)
	case "updateServiceNotification":
		t.UpdateServiceNotification = new(UpdateServiceNotification)
		return json.Unmarshal(b, t.UpdateServiceNotification)
	case "updateSpeechRecognitionTrial":
		t.UpdateSpeechRecognitionTrial = new(UpdateSpeechRecognitionTrial)
		return json.Unmarshal(b, t.UpdateSpeechRecognitionTrial)
	case "updateSpeedLimitNotification":
		t.UpdateSpeedLimitNotification = new(UpdateSpeedLimitNotification)
		return json.Unmarshal(b, t.UpdateSpeedLimitNotification)
	case "updateStakeDiceState":
		t.UpdateStakeDiceState = new(UpdateStakeDiceState)
		return json.Unmarshal(b, t.UpdateStakeDiceState)
	case "updateStarRevenueStatus":
		t.UpdateStarRevenueStatus = new(UpdateStarRevenueStatus)
		return json.Unmarshal(b, t.UpdateStarRevenueStatus)
	case "updateStickerSet":
		t.UpdateStickerSet = new(UpdateStickerSet)
		return json.Unmarshal(b, t.UpdateStickerSet)
	case "updateStory":
		t.UpdateStory = new(UpdateStory)
		return json.Unmarshal(b, t.UpdateStory)
	case "updateStoryDeleted":
		t.UpdateStoryDeleted = new(UpdateStoryDeleted)
		return json.Unmarshal(b, t.UpdateStoryDeleted)
	case "updateStoryListChatCount":
		t.UpdateStoryListChatCount = new(UpdateStoryListChatCount)
		return json.Unmarshal(b, t.UpdateStoryListChatCount)
	case "updateStoryPostFailed":
		t.UpdateStoryPostFailed = new(UpdateStoryPostFailed)
		return json.Unmarshal(b, t.UpdateStoryPostFailed)
	case "updateStoryPostSucceeded":
		t.UpdateStoryPostSucceeded = new(UpdateStoryPostSucceeded)
		return json.Unmarshal(b, t.UpdateStoryPostSucceeded)
	case "updateStoryStealthMode":
		t.UpdateStoryStealthMode = new(UpdateStoryStealthMode)
		return json.Unmarshal(b, t.UpdateStoryStealthMode)
	case "updateSuggestedActions":
		t.UpdateSuggestedActions = new(UpdateSuggestedActions)
		return json.Unmarshal(b, t.UpdateSuggestedActions)
	case "updateSupergroup":
		t.UpdateSupergroup = new(UpdateSupergroup)
		return json.Unmarshal(b, t.UpdateSupergroup)
	case "updateSupergroupFullInfo":
		t.UpdateSupergroupFullInfo = new(UpdateSupergroupFullInfo)
		return json.Unmarshal(b, t.UpdateSupergroupFullInfo)
	case "updateTermsOfService":
		t.UpdateTermsOfService = new(UpdateTermsOfService)
		return json.Unmarshal(b, t.UpdateTermsOfService)
	case "updateTonRevenueStatus":
		t.UpdateTonRevenueStatus = new(UpdateTonRevenueStatus)
		return json.Unmarshal(b, t.UpdateTonRevenueStatus)
	case "updateTopicMessageCount":
		t.UpdateTopicMessageCount = new(UpdateTopicMessageCount)
		return json.Unmarshal(b, t.UpdateTopicMessageCount)
	case "updateTrendingStickerSets":
		t.UpdateTrendingStickerSets = new(UpdateTrendingStickerSets)
		return json.Unmarshal(b, t.UpdateTrendingStickerSets)
	case "updateTrustedMiniAppBots":
		t.UpdateTrustedMiniAppBots = new(UpdateTrustedMiniAppBots)
		return json.Unmarshal(b, t.UpdateTrustedMiniAppBots)
	case "updateUnconfirmedSession":
		t.UpdateUnconfirmedSession = new(UpdateUnconfirmedSession)
		return json.Unmarshal(b, t.UpdateUnconfirmedSession)
	case "updateUnreadChatCount":
		t.UpdateUnreadChatCount = new(UpdateUnreadChatCount)
		return json.Unmarshal(b, t.UpdateUnreadChatCount)
	case "updateUnreadMessageCount":
		t.UpdateUnreadMessageCount = new(UpdateUnreadMessageCount)
		return json.Unmarshal(b, t.UpdateUnreadMessageCount)
	case "updateUser":
		t.UpdateUser = new(UpdateUser)
		return json.Unmarshal(b, t.UpdateUser)
	case "updateUserFullInfo":
		t.UpdateUserFullInfo = new(UpdateUserFullInfo)
		return json.Unmarshal(b, t.UpdateUserFullInfo)
	case "updateUserPrivacySettingRules":
		t.UpdateUserPrivacySettingRules = new(UpdateUserPrivacySettingRules)
		return json.Unmarshal(b, t.UpdateUserPrivacySettingRules)
	case "updateUserStatus":
		t.UpdateUserStatus = new(UpdateUserStatus)
		return json.Unmarshal(b, t.UpdateUserStatus)
	case "updateVideoPublished":
		t.UpdateVideoPublished = new(UpdateVideoPublished)
		return json.Unmarshal(b, t.UpdateVideoPublished)
	case "updateWebAppMessageSent":
		t.UpdateWebAppMessageSent = new(UpdateWebAppMessageSent)
		return json.Unmarshal(b, t.UpdateWebAppMessageSent)
	}
	return nil
}

func (t *Update) MarshalJSON() ([]byte, error) {
	if t.UpdateAccentColors != nil {
		return json.Marshal(t.UpdateAccentColors)
	}
	if t.UpdateActiveEmojiReactions != nil {
		return json.Marshal(t.UpdateActiveEmojiReactions)
	}
	if t.UpdateActiveGiftAuctions != nil {
		return json.Marshal(t.UpdateActiveGiftAuctions)
	}
	if t.UpdateActiveLiveLocationMessages != nil {
		return json.Marshal(t.UpdateActiveLiveLocationMessages)
	}
	if t.UpdateActiveNotifications != nil {
		return json.Marshal(t.UpdateActiveNotifications)
	}
	if t.UpdateAgeVerificationParameters != nil {
		return json.Marshal(t.UpdateAgeVerificationParameters)
	}
	if t.UpdateAnimatedEmojiMessageClicked != nil {
		return json.Marshal(t.UpdateAnimatedEmojiMessageClicked)
	}
	if t.UpdateAnimationSearchParameters != nil {
		return json.Marshal(t.UpdateAnimationSearchParameters)
	}
	if t.UpdateApplicationRecaptchaVerificationRequired != nil {
		return json.Marshal(t.UpdateApplicationRecaptchaVerificationRequired)
	}
	if t.UpdateApplicationVerificationRequired != nil {
		return json.Marshal(t.UpdateApplicationVerificationRequired)
	}
	if t.UpdateAttachmentMenuBots != nil {
		return json.Marshal(t.UpdateAttachmentMenuBots)
	}
	if t.UpdateAuthorizationState != nil {
		return json.Marshal(t.UpdateAuthorizationState)
	}
	if t.UpdateAutosaveSettings != nil {
		return json.Marshal(t.UpdateAutosaveSettings)
	}
	if t.UpdateAvailableMessageEffects != nil {
		return json.Marshal(t.UpdateAvailableMessageEffects)
	}
	if t.UpdateBasicGroup != nil {
		return json.Marshal(t.UpdateBasicGroup)
	}
	if t.UpdateBasicGroupFullInfo != nil {
		return json.Marshal(t.UpdateBasicGroupFullInfo)
	}
	if t.UpdateBusinessConnection != nil {
		return json.Marshal(t.UpdateBusinessConnection)
	}
	if t.UpdateBusinessMessageEdited != nil {
		return json.Marshal(t.UpdateBusinessMessageEdited)
	}
	if t.UpdateBusinessMessagesDeleted != nil {
		return json.Marshal(t.UpdateBusinessMessagesDeleted)
	}
	if t.UpdateCall != nil {
		return json.Marshal(t.UpdateCall)
	}
	if t.UpdateChatAccentColors != nil {
		return json.Marshal(t.UpdateChatAccentColors)
	}
	if t.UpdateChatAction != nil {
		return json.Marshal(t.UpdateChatAction)
	}
	if t.UpdateChatActionBar != nil {
		return json.Marshal(t.UpdateChatActionBar)
	}
	if t.UpdateChatActiveStories != nil {
		return json.Marshal(t.UpdateChatActiveStories)
	}
	if t.UpdateChatAddedToList != nil {
		return json.Marshal(t.UpdateChatAddedToList)
	}
	if t.UpdateChatAvailableReactions != nil {
		return json.Marshal(t.UpdateChatAvailableReactions)
	}
	if t.UpdateChatBackground != nil {
		return json.Marshal(t.UpdateChatBackground)
	}
	if t.UpdateChatBlockList != nil {
		return json.Marshal(t.UpdateChatBlockList)
	}
	if t.UpdateChatBoost != nil {
		return json.Marshal(t.UpdateChatBoost)
	}
	if t.UpdateChatBusinessBotManageBar != nil {
		return json.Marshal(t.UpdateChatBusinessBotManageBar)
	}
	if t.UpdateChatDefaultDisableNotification != nil {
		return json.Marshal(t.UpdateChatDefaultDisableNotification)
	}
	if t.UpdateChatDraftMessage != nil {
		return json.Marshal(t.UpdateChatDraftMessage)
	}
	if t.UpdateChatEmojiStatus != nil {
		return json.Marshal(t.UpdateChatEmojiStatus)
	}
	if t.UpdateChatFolders != nil {
		return json.Marshal(t.UpdateChatFolders)
	}
	if t.UpdateChatHasProtectedContent != nil {
		return json.Marshal(t.UpdateChatHasProtectedContent)
	}
	if t.UpdateChatHasScheduledMessages != nil {
		return json.Marshal(t.UpdateChatHasScheduledMessages)
	}
	if t.UpdateChatIsMarkedAsUnread != nil {
		return json.Marshal(t.UpdateChatIsMarkedAsUnread)
	}
	if t.UpdateChatIsTranslatable != nil {
		return json.Marshal(t.UpdateChatIsTranslatable)
	}
	if t.UpdateChatLastMessage != nil {
		return json.Marshal(t.UpdateChatLastMessage)
	}
	if t.UpdateChatMember != nil {
		return json.Marshal(t.UpdateChatMember)
	}
	if t.UpdateChatMessageAutoDeleteTime != nil {
		return json.Marshal(t.UpdateChatMessageAutoDeleteTime)
	}
	if t.UpdateChatMessageSender != nil {
		return json.Marshal(t.UpdateChatMessageSender)
	}
	if t.UpdateChatNotificationSettings != nil {
		return json.Marshal(t.UpdateChatNotificationSettings)
	}
	if t.UpdateChatOnlineMemberCount != nil {
		return json.Marshal(t.UpdateChatOnlineMemberCount)
	}
	if t.UpdateChatPendingJoinRequests != nil {
		return json.Marshal(t.UpdateChatPendingJoinRequests)
	}
	if t.UpdateChatPermissions != nil {
		return json.Marshal(t.UpdateChatPermissions)
	}
	if t.UpdateChatPhoto != nil {
		return json.Marshal(t.UpdateChatPhoto)
	}
	if t.UpdateChatPosition != nil {
		return json.Marshal(t.UpdateChatPosition)
	}
	if t.UpdateChatReadInbox != nil {
		return json.Marshal(t.UpdateChatReadInbox)
	}
	if t.UpdateChatReadOutbox != nil {
		return json.Marshal(t.UpdateChatReadOutbox)
	}
	if t.UpdateChatRemovedFromList != nil {
		return json.Marshal(t.UpdateChatRemovedFromList)
	}
	if t.UpdateChatReplyMarkup != nil {
		return json.Marshal(t.UpdateChatReplyMarkup)
	}
	if t.UpdateChatRevenueAmount != nil {
		return json.Marshal(t.UpdateChatRevenueAmount)
	}
	if t.UpdateChatTheme != nil {
		return json.Marshal(t.UpdateChatTheme)
	}
	if t.UpdateChatTitle != nil {
		return json.Marshal(t.UpdateChatTitle)
	}
	if t.UpdateChatUnreadMentionCount != nil {
		return json.Marshal(t.UpdateChatUnreadMentionCount)
	}
	if t.UpdateChatUnreadReactionCount != nil {
		return json.Marshal(t.UpdateChatUnreadReactionCount)
	}
	if t.UpdateChatVideoChat != nil {
		return json.Marshal(t.UpdateChatVideoChat)
	}
	if t.UpdateChatViewAsTopics != nil {
		return json.Marshal(t.UpdateChatViewAsTopics)
	}
	if t.UpdateConnectionState != nil {
		return json.Marshal(t.UpdateConnectionState)
	}
	if t.UpdateContactCloseBirthdays != nil {
		return json.Marshal(t.UpdateContactCloseBirthdays)
	}
	if t.UpdateDefaultBackground != nil {
		return json.Marshal(t.UpdateDefaultBackground)
	}
	if t.UpdateDefaultPaidReactionType != nil {
		return json.Marshal(t.UpdateDefaultPaidReactionType)
	}
	if t.UpdateDefaultReactionType != nil {
		return json.Marshal(t.UpdateDefaultReactionType)
	}
	if t.UpdateDeleteMessages != nil {
		return json.Marshal(t.UpdateDeleteMessages)
	}
	if t.UpdateDiceEmojis != nil {
		return json.Marshal(t.UpdateDiceEmojis)
	}
	if t.UpdateDirectMessagesChatTopic != nil {
		return json.Marshal(t.UpdateDirectMessagesChatTopic)
	}
	if t.UpdateEmojiChatThemes != nil {
		return json.Marshal(t.UpdateEmojiChatThemes)
	}
	if t.UpdateFavoriteStickers != nil {
		return json.Marshal(t.UpdateFavoriteStickers)
	}
	if t.UpdateFile != nil {
		return json.Marshal(t.UpdateFile)
	}
	if t.UpdateFileAddedToDownloads != nil {
		return json.Marshal(t.UpdateFileAddedToDownloads)
	}
	if t.UpdateFileDownload != nil {
		return json.Marshal(t.UpdateFileDownload)
	}
	if t.UpdateFileDownloads != nil {
		return json.Marshal(t.UpdateFileDownloads)
	}
	if t.UpdateFileGenerationStart != nil {
		return json.Marshal(t.UpdateFileGenerationStart)
	}
	if t.UpdateFileGenerationStop != nil {
		return json.Marshal(t.UpdateFileGenerationStop)
	}
	if t.UpdateFileRemovedFromDownloads != nil {
		return json.Marshal(t.UpdateFileRemovedFromDownloads)
	}
	if t.UpdateForumTopic != nil {
		return json.Marshal(t.UpdateForumTopic)
	}
	if t.UpdateForumTopicInfo != nil {
		return json.Marshal(t.UpdateForumTopicInfo)
	}
	if t.UpdateFreezeState != nil {
		return json.Marshal(t.UpdateFreezeState)
	}
	if t.UpdateGiftAuctionState != nil {
		return json.Marshal(t.UpdateGiftAuctionState)
	}
	if t.UpdateGroupCall != nil {
		return json.Marshal(t.UpdateGroupCall)
	}
	if t.UpdateGroupCallMessageLevels != nil {
		return json.Marshal(t.UpdateGroupCallMessageLevels)
	}
	if t.UpdateGroupCallMessageSendFailed != nil {
		return json.Marshal(t.UpdateGroupCallMessageSendFailed)
	}
	if t.UpdateGroupCallMessagesDeleted != nil {
		return json.Marshal(t.UpdateGroupCallMessagesDeleted)
	}
	if t.UpdateGroupCallParticipant != nil {
		return json.Marshal(t.UpdateGroupCallParticipant)
	}
	if t.UpdateGroupCallParticipants != nil {
		return json.Marshal(t.UpdateGroupCallParticipants)
	}
	if t.UpdateGroupCallVerificationState != nil {
		return json.Marshal(t.UpdateGroupCallVerificationState)
	}
	if t.UpdateHavePendingNotifications != nil {
		return json.Marshal(t.UpdateHavePendingNotifications)
	}
	if t.UpdateInstalledStickerSets != nil {
		return json.Marshal(t.UpdateInstalledStickerSets)
	}
	if t.UpdateLanguagePackStrings != nil {
		return json.Marshal(t.UpdateLanguagePackStrings)
	}
	if t.UpdateLiveStoryTopDonors != nil {
		return json.Marshal(t.UpdateLiveStoryTopDonors)
	}
	if t.UpdateMessageContent != nil {
		return json.Marshal(t.UpdateMessageContent)
	}
	if t.UpdateMessageContentOpened != nil {
		return json.Marshal(t.UpdateMessageContentOpened)
	}
	if t.UpdateMessageEdited != nil {
		return json.Marshal(t.UpdateMessageEdited)
	}
	if t.UpdateMessageFactCheck != nil {
		return json.Marshal(t.UpdateMessageFactCheck)
	}
	if t.UpdateMessageInteractionInfo != nil {
		return json.Marshal(t.UpdateMessageInteractionInfo)
	}
	if t.UpdateMessageIsPinned != nil {
		return json.Marshal(t.UpdateMessageIsPinned)
	}
	if t.UpdateMessageLiveLocationViewed != nil {
		return json.Marshal(t.UpdateMessageLiveLocationViewed)
	}
	if t.UpdateMessageMentionRead != nil {
		return json.Marshal(t.UpdateMessageMentionRead)
	}
	if t.UpdateMessageReaction != nil {
		return json.Marshal(t.UpdateMessageReaction)
	}
	if t.UpdateMessageReactions != nil {
		return json.Marshal(t.UpdateMessageReactions)
	}
	if t.UpdateMessageSendAcknowledged != nil {
		return json.Marshal(t.UpdateMessageSendAcknowledged)
	}
	if t.UpdateMessageSendFailed != nil {
		return json.Marshal(t.UpdateMessageSendFailed)
	}
	if t.UpdateMessageSendSucceeded != nil {
		return json.Marshal(t.UpdateMessageSendSucceeded)
	}
	if t.UpdateMessageSuggestedPostInfo != nil {
		return json.Marshal(t.UpdateMessageSuggestedPostInfo)
	}
	if t.UpdateMessageUnreadReactions != nil {
		return json.Marshal(t.UpdateMessageUnreadReactions)
	}
	if t.UpdateNewBusinessCallbackQuery != nil {
		return json.Marshal(t.UpdateNewBusinessCallbackQuery)
	}
	if t.UpdateNewBusinessMessage != nil {
		return json.Marshal(t.UpdateNewBusinessMessage)
	}
	if t.UpdateNewCallSignalingData != nil {
		return json.Marshal(t.UpdateNewCallSignalingData)
	}
	if t.UpdateNewCallbackQuery != nil {
		return json.Marshal(t.UpdateNewCallbackQuery)
	}
	if t.UpdateNewChat != nil {
		return json.Marshal(t.UpdateNewChat)
	}
	if t.UpdateNewChatJoinRequest != nil {
		return json.Marshal(t.UpdateNewChatJoinRequest)
	}
	if t.UpdateNewChosenInlineResult != nil {
		return json.Marshal(t.UpdateNewChosenInlineResult)
	}
	if t.UpdateNewCustomEvent != nil {
		return json.Marshal(t.UpdateNewCustomEvent)
	}
	if t.UpdateNewCustomQuery != nil {
		return json.Marshal(t.UpdateNewCustomQuery)
	}
	if t.UpdateNewGroupCallMessage != nil {
		return json.Marshal(t.UpdateNewGroupCallMessage)
	}
	if t.UpdateNewGroupCallPaidReaction != nil {
		return json.Marshal(t.UpdateNewGroupCallPaidReaction)
	}
	if t.UpdateNewInlineCallbackQuery != nil {
		return json.Marshal(t.UpdateNewInlineCallbackQuery)
	}
	if t.UpdateNewInlineQuery != nil {
		return json.Marshal(t.UpdateNewInlineQuery)
	}
	if t.UpdateNewMessage != nil {
		return json.Marshal(t.UpdateNewMessage)
	}
	if t.UpdateNewPreCheckoutQuery != nil {
		return json.Marshal(t.UpdateNewPreCheckoutQuery)
	}
	if t.UpdateNewShippingQuery != nil {
		return json.Marshal(t.UpdateNewShippingQuery)
	}
	if t.UpdateNotification != nil {
		return json.Marshal(t.UpdateNotification)
	}
	if t.UpdateNotificationGroup != nil {
		return json.Marshal(t.UpdateNotificationGroup)
	}
	if t.UpdateOption != nil {
		return json.Marshal(t.UpdateOption)
	}
	if t.UpdateOwnedStarCount != nil {
		return json.Marshal(t.UpdateOwnedStarCount)
	}
	if t.UpdateOwnedTonCount != nil {
		return json.Marshal(t.UpdateOwnedTonCount)
	}
	if t.UpdatePaidMediaPurchased != nil {
		return json.Marshal(t.UpdatePaidMediaPurchased)
	}
	if t.UpdatePendingTextMessage != nil {
		return json.Marshal(t.UpdatePendingTextMessage)
	}
	if t.UpdatePoll != nil {
		return json.Marshal(t.UpdatePoll)
	}
	if t.UpdatePollAnswer != nil {
		return json.Marshal(t.UpdatePollAnswer)
	}
	if t.UpdateProfileAccentColors != nil {
		return json.Marshal(t.UpdateProfileAccentColors)
	}
	if t.UpdateQuickReplyShortcut != nil {
		return json.Marshal(t.UpdateQuickReplyShortcut)
	}
	if t.UpdateQuickReplyShortcutDeleted != nil {
		return json.Marshal(t.UpdateQuickReplyShortcutDeleted)
	}
	if t.UpdateQuickReplyShortcutMessages != nil {
		return json.Marshal(t.UpdateQuickReplyShortcutMessages)
	}
	if t.UpdateQuickReplyShortcuts != nil {
		return json.Marshal(t.UpdateQuickReplyShortcuts)
	}
	if t.UpdateReactionNotificationSettings != nil {
		return json.Marshal(t.UpdateReactionNotificationSettings)
	}
	if t.UpdateRecentStickers != nil {
		return json.Marshal(t.UpdateRecentStickers)
	}
	if t.UpdateSavedAnimations != nil {
		return json.Marshal(t.UpdateSavedAnimations)
	}
	if t.UpdateSavedMessagesTags != nil {
		return json.Marshal(t.UpdateSavedMessagesTags)
	}
	if t.UpdateSavedMessagesTopic != nil {
		return json.Marshal(t.UpdateSavedMessagesTopic)
	}
	if t.UpdateSavedMessagesTopicCount != nil {
		return json.Marshal(t.UpdateSavedMessagesTopicCount)
	}
	if t.UpdateSavedNotificationSounds != nil {
		return json.Marshal(t.UpdateSavedNotificationSounds)
	}
	if t.UpdateScopeNotificationSettings != nil {
		return json.Marshal(t.UpdateScopeNotificationSettings)
	}
	if t.UpdateSecretChat != nil {
		return json.Marshal(t.UpdateSecretChat)
	}
	if t.UpdateServiceNotification != nil {
		return json.Marshal(t.UpdateServiceNotification)
	}
	if t.UpdateSpeechRecognitionTrial != nil {
		return json.Marshal(t.UpdateSpeechRecognitionTrial)
	}
	if t.UpdateSpeedLimitNotification != nil {
		return json.Marshal(t.UpdateSpeedLimitNotification)
	}
	if t.UpdateStakeDiceState != nil {
		return json.Marshal(t.UpdateStakeDiceState)
	}
	if t.UpdateStarRevenueStatus != nil {
		return json.Marshal(t.UpdateStarRevenueStatus)
	}
	if t.UpdateStickerSet != nil {
		return json.Marshal(t.UpdateStickerSet)
	}
	if t.UpdateStory != nil {
		return json.Marshal(t.UpdateStory)
	}
	if t.UpdateStoryDeleted != nil {
		return json.Marshal(t.UpdateStoryDeleted)
	}
	if t.UpdateStoryListChatCount != nil {
		return json.Marshal(t.UpdateStoryListChatCount)
	}
	if t.UpdateStoryPostFailed != nil {
		return json.Marshal(t.UpdateStoryPostFailed)
	}
	if t.UpdateStoryPostSucceeded != nil {
		return json.Marshal(t.UpdateStoryPostSucceeded)
	}
	if t.UpdateStoryStealthMode != nil {
		return json.Marshal(t.UpdateStoryStealthMode)
	}
	if t.UpdateSuggestedActions != nil {
		return json.Marshal(t.UpdateSuggestedActions)
	}
	if t.UpdateSupergroup != nil {
		return json.Marshal(t.UpdateSupergroup)
	}
	if t.UpdateSupergroupFullInfo != nil {
		return json.Marshal(t.UpdateSupergroupFullInfo)
	}
	if t.UpdateTermsOfService != nil {
		return json.Marshal(t.UpdateTermsOfService)
	}
	if t.UpdateTonRevenueStatus != nil {
		return json.Marshal(t.UpdateTonRevenueStatus)
	}
	if t.UpdateTopicMessageCount != nil {
		return json.Marshal(t.UpdateTopicMessageCount)
	}
	if t.UpdateTrendingStickerSets != nil {
		return json.Marshal(t.UpdateTrendingStickerSets)
	}
	if t.UpdateTrustedMiniAppBots != nil {
		return json.Marshal(t.UpdateTrustedMiniAppBots)
	}
	if t.UpdateUnconfirmedSession != nil {
		return json.Marshal(t.UpdateUnconfirmedSession)
	}
	if t.UpdateUnreadChatCount != nil {
		return json.Marshal(t.UpdateUnreadChatCount)
	}
	if t.UpdateUnreadMessageCount != nil {
		return json.Marshal(t.UpdateUnreadMessageCount)
	}
	if t.UpdateUser != nil {
		return json.Marshal(t.UpdateUser)
	}
	if t.UpdateUserFullInfo != nil {
		return json.Marshal(t.UpdateUserFullInfo)
	}
	if t.UpdateUserPrivacySettingRules != nil {
		return json.Marshal(t.UpdateUserPrivacySettingRules)
	}
	if t.UpdateUserStatus != nil {
		return json.Marshal(t.UpdateUserStatus)
	}
	if t.UpdateVideoPublished != nil {
		return json.Marshal(t.UpdateVideoPublished)
	}
	if t.UpdateWebAppMessageSent != nil {
		return json.Marshal(t.UpdateWebAppMessageSent)
	}
	return nil, fmt.Errorf("no value set for Update")
}

// UpgradedGiftAttributeId Contains identifier of an upgraded gift attribute to search for
type UpgradedGiftAttributeId struct {
	TypeStr                         string                           `json:"@type"`
	UpgradedGiftAttributeIdBackdrop *UpgradedGiftAttributeIdBackdrop `json:"upgradedGiftAttributeIdBackdrop,omitempty"`
	UpgradedGiftAttributeIdModel    *UpgradedGiftAttributeIdModel    `json:"upgradedGiftAttributeIdModel,omitempty"`
	UpgradedGiftAttributeIdSymbol   *UpgradedGiftAttributeIdSymbol   `json:"upgradedGiftAttributeIdSymbol,omitempty"`
}

func (t *UpgradedGiftAttributeId) Type() string {
	return t.TypeStr
}

func (t *UpgradedGiftAttributeId) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UpgradedGiftAttributeId) GetExtra() string {
	return ""
}

func (t *UpgradedGiftAttributeId) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "upgradedGiftAttributeIdBackdrop":
		t.UpgradedGiftAttributeIdBackdrop = new(UpgradedGiftAttributeIdBackdrop)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdBackdrop)
	case "upgradedGiftAttributeIdModel":
		t.UpgradedGiftAttributeIdModel = new(UpgradedGiftAttributeIdModel)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdModel)
	case "upgradedGiftAttributeIdSymbol":
		t.UpgradedGiftAttributeIdSymbol = new(UpgradedGiftAttributeIdSymbol)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdSymbol)
	}
	return nil
}

func (t *UpgradedGiftAttributeId) MarshalJSON() ([]byte, error) {
	if t.UpgradedGiftAttributeIdBackdrop != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdBackdrop)
	}
	if t.UpgradedGiftAttributeIdModel != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdModel)
	}
	if t.UpgradedGiftAttributeIdSymbol != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdSymbol)
	}
	return nil, fmt.Errorf("no value set for UpgradedGiftAttributeId")
}

// UpgradedGiftOrigin Describes origin from which the upgraded gift was obtained
type UpgradedGiftOrigin struct {
	TypeStr                          string                            `json:"@type"`
	UpgradedGiftOriginBlockchain     *UpgradedGiftOriginBlockchain     `json:"upgradedGiftOriginBlockchain,omitempty"`
	UpgradedGiftOriginOffer          *UpgradedGiftOriginOffer          `json:"upgradedGiftOriginOffer,omitempty"`
	UpgradedGiftOriginPrepaidUpgrade *UpgradedGiftOriginPrepaidUpgrade `json:"upgradedGiftOriginPrepaidUpgrade,omitempty"`
	UpgradedGiftOriginResale         *UpgradedGiftOriginResale         `json:"upgradedGiftOriginResale,omitempty"`
	UpgradedGiftOriginTransfer       *UpgradedGiftOriginTransfer       `json:"upgradedGiftOriginTransfer,omitempty"`
	UpgradedGiftOriginUpgrade        *UpgradedGiftOriginUpgrade        `json:"upgradedGiftOriginUpgrade,omitempty"`
}

func (t *UpgradedGiftOrigin) Type() string {
	return t.TypeStr
}

func (t *UpgradedGiftOrigin) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UpgradedGiftOrigin) GetExtra() string {
	return ""
}

func (t *UpgradedGiftOrigin) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "upgradedGiftOriginBlockchain":
		t.UpgradedGiftOriginBlockchain = new(UpgradedGiftOriginBlockchain)
		return json.Unmarshal(b, t.UpgradedGiftOriginBlockchain)
	case "upgradedGiftOriginOffer":
		t.UpgradedGiftOriginOffer = new(UpgradedGiftOriginOffer)
		return json.Unmarshal(b, t.UpgradedGiftOriginOffer)
	case "upgradedGiftOriginPrepaidUpgrade":
		t.UpgradedGiftOriginPrepaidUpgrade = new(UpgradedGiftOriginPrepaidUpgrade)
		return json.Unmarshal(b, t.UpgradedGiftOriginPrepaidUpgrade)
	case "upgradedGiftOriginResale":
		t.UpgradedGiftOriginResale = new(UpgradedGiftOriginResale)
		return json.Unmarshal(b, t.UpgradedGiftOriginResale)
	case "upgradedGiftOriginTransfer":
		t.UpgradedGiftOriginTransfer = new(UpgradedGiftOriginTransfer)
		return json.Unmarshal(b, t.UpgradedGiftOriginTransfer)
	case "upgradedGiftOriginUpgrade":
		t.UpgradedGiftOriginUpgrade = new(UpgradedGiftOriginUpgrade)
		return json.Unmarshal(b, t.UpgradedGiftOriginUpgrade)
	}
	return nil
}

func (t *UpgradedGiftOrigin) MarshalJSON() ([]byte, error) {
	if t.UpgradedGiftOriginBlockchain != nil {
		return json.Marshal(t.UpgradedGiftOriginBlockchain)
	}
	if t.UpgradedGiftOriginOffer != nil {
		return json.Marshal(t.UpgradedGiftOriginOffer)
	}
	if t.UpgradedGiftOriginPrepaidUpgrade != nil {
		return json.Marshal(t.UpgradedGiftOriginPrepaidUpgrade)
	}
	if t.UpgradedGiftOriginResale != nil {
		return json.Marshal(t.UpgradedGiftOriginResale)
	}
	if t.UpgradedGiftOriginTransfer != nil {
		return json.Marshal(t.UpgradedGiftOriginTransfer)
	}
	if t.UpgradedGiftOriginUpgrade != nil {
		return json.Marshal(t.UpgradedGiftOriginUpgrade)
	}
	return nil, fmt.Errorf("no value set for UpgradedGiftOrigin")
}

// UserPrivacySetting Describes available user privacy settings
type UserPrivacySetting struct {
	TypeStr                                                 string                                                   `json:"@type"`
	UserPrivacySettingAllowCalls                            *UserPrivacySettingAllowCalls                            `json:"userPrivacySettingAllowCalls,omitempty"`
	UserPrivacySettingAllowChatInvites                      *UserPrivacySettingAllowChatInvites                      `json:"userPrivacySettingAllowChatInvites,omitempty"`
	UserPrivacySettingAllowFindingByPhoneNumber             *UserPrivacySettingAllowFindingByPhoneNumber             `json:"userPrivacySettingAllowFindingByPhoneNumber,omitempty"`
	UserPrivacySettingAllowPeerToPeerCalls                  *UserPrivacySettingAllowPeerToPeerCalls                  `json:"userPrivacySettingAllowPeerToPeerCalls,omitempty"`
	UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages *UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages `json:"userPrivacySettingAllowPrivateVoiceAndVideoNoteMessages,omitempty"`
	UserPrivacySettingAllowUnpaidMessages                   *UserPrivacySettingAllowUnpaidMessages                   `json:"userPrivacySettingAllowUnpaidMessages,omitempty"`
	UserPrivacySettingAutosaveGifts                         *UserPrivacySettingAutosaveGifts                         `json:"userPrivacySettingAutosaveGifts,omitempty"`
	UserPrivacySettingShowBio                               *UserPrivacySettingShowBio                               `json:"userPrivacySettingShowBio,omitempty"`
	UserPrivacySettingShowBirthdate                         *UserPrivacySettingShowBirthdate                         `json:"userPrivacySettingShowBirthdate,omitempty"`
	UserPrivacySettingShowLinkInForwardedMessages           *UserPrivacySettingShowLinkInForwardedMessages           `json:"userPrivacySettingShowLinkInForwardedMessages,omitempty"`
	UserPrivacySettingShowPhoneNumber                       *UserPrivacySettingShowPhoneNumber                       `json:"userPrivacySettingShowPhoneNumber,omitempty"`
	UserPrivacySettingShowProfileAudio                      *UserPrivacySettingShowProfileAudio                      `json:"userPrivacySettingShowProfileAudio,omitempty"`
	UserPrivacySettingShowProfilePhoto                      *UserPrivacySettingShowProfilePhoto                      `json:"userPrivacySettingShowProfilePhoto,omitempty"`
	UserPrivacySettingShowStatus                            *UserPrivacySettingShowStatus                            `json:"userPrivacySettingShowStatus,omitempty"`
}

func (t *UserPrivacySetting) Type() string {
	return t.TypeStr
}

func (t *UserPrivacySetting) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UserPrivacySetting) GetExtra() string {
	return ""
}

func (t *UserPrivacySetting) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "userPrivacySettingAllowCalls":
		t.UserPrivacySettingAllowCalls = new(UserPrivacySettingAllowCalls)
		return json.Unmarshal(b, t.UserPrivacySettingAllowCalls)
	case "userPrivacySettingAllowChatInvites":
		t.UserPrivacySettingAllowChatInvites = new(UserPrivacySettingAllowChatInvites)
		return json.Unmarshal(b, t.UserPrivacySettingAllowChatInvites)
	case "userPrivacySettingAllowFindingByPhoneNumber":
		t.UserPrivacySettingAllowFindingByPhoneNumber = new(UserPrivacySettingAllowFindingByPhoneNumber)
		return json.Unmarshal(b, t.UserPrivacySettingAllowFindingByPhoneNumber)
	case "userPrivacySettingAllowPeerToPeerCalls":
		t.UserPrivacySettingAllowPeerToPeerCalls = new(UserPrivacySettingAllowPeerToPeerCalls)
		return json.Unmarshal(b, t.UserPrivacySettingAllowPeerToPeerCalls)
	case "userPrivacySettingAllowPrivateVoiceAndVideoNoteMessages":
		t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages = new(UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
		return json.Unmarshal(b, t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
	case "userPrivacySettingAllowUnpaidMessages":
		t.UserPrivacySettingAllowUnpaidMessages = new(UserPrivacySettingAllowUnpaidMessages)
		return json.Unmarshal(b, t.UserPrivacySettingAllowUnpaidMessages)
	case "userPrivacySettingAutosaveGifts":
		t.UserPrivacySettingAutosaveGifts = new(UserPrivacySettingAutosaveGifts)
		return json.Unmarshal(b, t.UserPrivacySettingAutosaveGifts)
	case "userPrivacySettingShowBio":
		t.UserPrivacySettingShowBio = new(UserPrivacySettingShowBio)
		return json.Unmarshal(b, t.UserPrivacySettingShowBio)
	case "userPrivacySettingShowBirthdate":
		t.UserPrivacySettingShowBirthdate = new(UserPrivacySettingShowBirthdate)
		return json.Unmarshal(b, t.UserPrivacySettingShowBirthdate)
	case "userPrivacySettingShowLinkInForwardedMessages":
		t.UserPrivacySettingShowLinkInForwardedMessages = new(UserPrivacySettingShowLinkInForwardedMessages)
		return json.Unmarshal(b, t.UserPrivacySettingShowLinkInForwardedMessages)
	case "userPrivacySettingShowPhoneNumber":
		t.UserPrivacySettingShowPhoneNumber = new(UserPrivacySettingShowPhoneNumber)
		return json.Unmarshal(b, t.UserPrivacySettingShowPhoneNumber)
	case "userPrivacySettingShowProfileAudio":
		t.UserPrivacySettingShowProfileAudio = new(UserPrivacySettingShowProfileAudio)
		return json.Unmarshal(b, t.UserPrivacySettingShowProfileAudio)
	case "userPrivacySettingShowProfilePhoto":
		t.UserPrivacySettingShowProfilePhoto = new(UserPrivacySettingShowProfilePhoto)
		return json.Unmarshal(b, t.UserPrivacySettingShowProfilePhoto)
	case "userPrivacySettingShowStatus":
		t.UserPrivacySettingShowStatus = new(UserPrivacySettingShowStatus)
		return json.Unmarshal(b, t.UserPrivacySettingShowStatus)
	}
	return nil
}

func (t *UserPrivacySetting) MarshalJSON() ([]byte, error) {
	if t.UserPrivacySettingAllowCalls != nil {
		return json.Marshal(t.UserPrivacySettingAllowCalls)
	}
	if t.UserPrivacySettingAllowChatInvites != nil {
		return json.Marshal(t.UserPrivacySettingAllowChatInvites)
	}
	if t.UserPrivacySettingAllowFindingByPhoneNumber != nil {
		return json.Marshal(t.UserPrivacySettingAllowFindingByPhoneNumber)
	}
	if t.UserPrivacySettingAllowPeerToPeerCalls != nil {
		return json.Marshal(t.UserPrivacySettingAllowPeerToPeerCalls)
	}
	if t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages != nil {
		return json.Marshal(t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
	}
	if t.UserPrivacySettingAllowUnpaidMessages != nil {
		return json.Marshal(t.UserPrivacySettingAllowUnpaidMessages)
	}
	if t.UserPrivacySettingAutosaveGifts != nil {
		return json.Marshal(t.UserPrivacySettingAutosaveGifts)
	}
	if t.UserPrivacySettingShowBio != nil {
		return json.Marshal(t.UserPrivacySettingShowBio)
	}
	if t.UserPrivacySettingShowBirthdate != nil {
		return json.Marshal(t.UserPrivacySettingShowBirthdate)
	}
	if t.UserPrivacySettingShowLinkInForwardedMessages != nil {
		return json.Marshal(t.UserPrivacySettingShowLinkInForwardedMessages)
	}
	if t.UserPrivacySettingShowPhoneNumber != nil {
		return json.Marshal(t.UserPrivacySettingShowPhoneNumber)
	}
	if t.UserPrivacySettingShowProfileAudio != nil {
		return json.Marshal(t.UserPrivacySettingShowProfileAudio)
	}
	if t.UserPrivacySettingShowProfilePhoto != nil {
		return json.Marshal(t.UserPrivacySettingShowProfilePhoto)
	}
	if t.UserPrivacySettingShowStatus != nil {
		return json.Marshal(t.UserPrivacySettingShowStatus)
	}
	return nil, fmt.Errorf("no value set for UserPrivacySetting")
}

// UserPrivacySettingRule Represents a single rule for managing user privacy settings
type UserPrivacySettingRule struct {
	TypeStr                                   string                                     `json:"@type"`
	UserPrivacySettingRuleAllowAll            *UserPrivacySettingRuleAllowAll            `json:"userPrivacySettingRuleAllowAll,omitempty"`
	UserPrivacySettingRuleAllowBots           *UserPrivacySettingRuleAllowBots           `json:"userPrivacySettingRuleAllowBots,omitempty"`
	UserPrivacySettingRuleAllowChatMembers    *UserPrivacySettingRuleAllowChatMembers    `json:"userPrivacySettingRuleAllowChatMembers,omitempty"`
	UserPrivacySettingRuleAllowContacts       *UserPrivacySettingRuleAllowContacts       `json:"userPrivacySettingRuleAllowContacts,omitempty"`
	UserPrivacySettingRuleAllowPremiumUsers   *UserPrivacySettingRuleAllowPremiumUsers   `json:"userPrivacySettingRuleAllowPremiumUsers,omitempty"`
	UserPrivacySettingRuleAllowUsers          *UserPrivacySettingRuleAllowUsers          `json:"userPrivacySettingRuleAllowUsers,omitempty"`
	UserPrivacySettingRuleRestrictAll         *UserPrivacySettingRuleRestrictAll         `json:"userPrivacySettingRuleRestrictAll,omitempty"`
	UserPrivacySettingRuleRestrictBots        *UserPrivacySettingRuleRestrictBots        `json:"userPrivacySettingRuleRestrictBots,omitempty"`
	UserPrivacySettingRuleRestrictChatMembers *UserPrivacySettingRuleRestrictChatMembers `json:"userPrivacySettingRuleRestrictChatMembers,omitempty"`
	UserPrivacySettingRuleRestrictContacts    *UserPrivacySettingRuleRestrictContacts    `json:"userPrivacySettingRuleRestrictContacts,omitempty"`
	UserPrivacySettingRuleRestrictUsers       *UserPrivacySettingRuleRestrictUsers       `json:"userPrivacySettingRuleRestrictUsers,omitempty"`
}

func (t *UserPrivacySettingRule) Type() string {
	return t.TypeStr
}

func (t *UserPrivacySettingRule) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UserPrivacySettingRule) GetExtra() string {
	return ""
}

func (t *UserPrivacySettingRule) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "userPrivacySettingRuleAllowAll":
		t.UserPrivacySettingRuleAllowAll = new(UserPrivacySettingRuleAllowAll)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowAll)
	case "userPrivacySettingRuleAllowBots":
		t.UserPrivacySettingRuleAllowBots = new(UserPrivacySettingRuleAllowBots)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowBots)
	case "userPrivacySettingRuleAllowChatMembers":
		t.UserPrivacySettingRuleAllowChatMembers = new(UserPrivacySettingRuleAllowChatMembers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowChatMembers)
	case "userPrivacySettingRuleAllowContacts":
		t.UserPrivacySettingRuleAllowContacts = new(UserPrivacySettingRuleAllowContacts)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowContacts)
	case "userPrivacySettingRuleAllowPremiumUsers":
		t.UserPrivacySettingRuleAllowPremiumUsers = new(UserPrivacySettingRuleAllowPremiumUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowPremiumUsers)
	case "userPrivacySettingRuleAllowUsers":
		t.UserPrivacySettingRuleAllowUsers = new(UserPrivacySettingRuleAllowUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowUsers)
	case "userPrivacySettingRuleRestrictAll":
		t.UserPrivacySettingRuleRestrictAll = new(UserPrivacySettingRuleRestrictAll)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictAll)
	case "userPrivacySettingRuleRestrictBots":
		t.UserPrivacySettingRuleRestrictBots = new(UserPrivacySettingRuleRestrictBots)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictBots)
	case "userPrivacySettingRuleRestrictChatMembers":
		t.UserPrivacySettingRuleRestrictChatMembers = new(UserPrivacySettingRuleRestrictChatMembers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictChatMembers)
	case "userPrivacySettingRuleRestrictContacts":
		t.UserPrivacySettingRuleRestrictContacts = new(UserPrivacySettingRuleRestrictContacts)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictContacts)
	case "userPrivacySettingRuleRestrictUsers":
		t.UserPrivacySettingRuleRestrictUsers = new(UserPrivacySettingRuleRestrictUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictUsers)
	}
	return nil
}

func (t *UserPrivacySettingRule) MarshalJSON() ([]byte, error) {
	if t.UserPrivacySettingRuleAllowAll != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowAll)
	}
	if t.UserPrivacySettingRuleAllowBots != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowBots)
	}
	if t.UserPrivacySettingRuleAllowChatMembers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowChatMembers)
	}
	if t.UserPrivacySettingRuleAllowContacts != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowContacts)
	}
	if t.UserPrivacySettingRuleAllowPremiumUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowPremiumUsers)
	}
	if t.UserPrivacySettingRuleAllowUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowUsers)
	}
	if t.UserPrivacySettingRuleRestrictAll != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictAll)
	}
	if t.UserPrivacySettingRuleRestrictBots != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictBots)
	}
	if t.UserPrivacySettingRuleRestrictChatMembers != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictChatMembers)
	}
	if t.UserPrivacySettingRuleRestrictContacts != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictContacts)
	}
	if t.UserPrivacySettingRuleRestrictUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictUsers)
	}
	return nil, fmt.Errorf("no value set for UserPrivacySettingRule")
}

// UserStatus Describes the last time the user was online
type UserStatus struct {
	TypeStr             string               `json:"@type"`
	UserStatusEmpty     *UserStatusEmpty     `json:"userStatusEmpty,omitempty"`
	UserStatusLastMonth *UserStatusLastMonth `json:"userStatusLastMonth,omitempty"`
	UserStatusLastWeek  *UserStatusLastWeek  `json:"userStatusLastWeek,omitempty"`
	UserStatusOffline   *UserStatusOffline   `json:"userStatusOffline,omitempty"`
	UserStatusOnline    *UserStatusOnline    `json:"userStatusOnline,omitempty"`
	UserStatusRecently  *UserStatusRecently  `json:"userStatusRecently,omitempty"`
}

func (t *UserStatus) Type() string {
	return t.TypeStr
}

func (t *UserStatus) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UserStatus) GetExtra() string {
	return ""
}

func (t *UserStatus) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "userStatusEmpty":
		t.UserStatusEmpty = new(UserStatusEmpty)
		return json.Unmarshal(b, t.UserStatusEmpty)
	case "userStatusLastMonth":
		t.UserStatusLastMonth = new(UserStatusLastMonth)
		return json.Unmarshal(b, t.UserStatusLastMonth)
	case "userStatusLastWeek":
		t.UserStatusLastWeek = new(UserStatusLastWeek)
		return json.Unmarshal(b, t.UserStatusLastWeek)
	case "userStatusOffline":
		t.UserStatusOffline = new(UserStatusOffline)
		return json.Unmarshal(b, t.UserStatusOffline)
	case "userStatusOnline":
		t.UserStatusOnline = new(UserStatusOnline)
		return json.Unmarshal(b, t.UserStatusOnline)
	case "userStatusRecently":
		t.UserStatusRecently = new(UserStatusRecently)
		return json.Unmarshal(b, t.UserStatusRecently)
	}
	return nil
}

func (t *UserStatus) MarshalJSON() ([]byte, error) {
	if t.UserStatusEmpty != nil {
		return json.Marshal(t.UserStatusEmpty)
	}
	if t.UserStatusLastMonth != nil {
		return json.Marshal(t.UserStatusLastMonth)
	}
	if t.UserStatusLastWeek != nil {
		return json.Marshal(t.UserStatusLastWeek)
	}
	if t.UserStatusOffline != nil {
		return json.Marshal(t.UserStatusOffline)
	}
	if t.UserStatusOnline != nil {
		return json.Marshal(t.UserStatusOnline)
	}
	if t.UserStatusRecently != nil {
		return json.Marshal(t.UserStatusRecently)
	}
	return nil, fmt.Errorf("no value set for UserStatus")
}

// UserType Represents the type of user. The following types are possible: regular users, deleted users and bots
type UserType struct {
	TypeStr         string           `json:"@type"`
	UserTypeBot     *UserTypeBot     `json:"userTypeBot,omitempty"`
	UserTypeDeleted *UserTypeDeleted `json:"userTypeDeleted,omitempty"`
	UserTypeRegular *UserTypeRegular `json:"userTypeRegular,omitempty"`
	UserTypeUnknown *UserTypeUnknown `json:"userTypeUnknown,omitempty"`
}

func (t *UserType) Type() string {
	return t.TypeStr
}

func (t *UserType) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *UserType) GetExtra() string {
	return ""
}

func (t *UserType) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "userTypeBot":
		t.UserTypeBot = new(UserTypeBot)
		return json.Unmarshal(b, t.UserTypeBot)
	case "userTypeDeleted":
		t.UserTypeDeleted = new(UserTypeDeleted)
		return json.Unmarshal(b, t.UserTypeDeleted)
	case "userTypeRegular":
		t.UserTypeRegular = new(UserTypeRegular)
		return json.Unmarshal(b, t.UserTypeRegular)
	case "userTypeUnknown":
		t.UserTypeUnknown = new(UserTypeUnknown)
		return json.Unmarshal(b, t.UserTypeUnknown)
	}
	return nil
}

func (t *UserType) MarshalJSON() ([]byte, error) {
	if t.UserTypeBot != nil {
		return json.Marshal(t.UserTypeBot)
	}
	if t.UserTypeDeleted != nil {
		return json.Marshal(t.UserTypeDeleted)
	}
	if t.UserTypeRegular != nil {
		return json.Marshal(t.UserTypeRegular)
	}
	if t.UserTypeUnknown != nil {
		return json.Marshal(t.UserTypeUnknown)
	}
	return nil, fmt.Errorf("no value set for UserType")
}

// VectorPathCommand Represents a vector path command
type VectorPathCommand struct {
	TypeStr                           string                             `json:"@type"`
	VectorPathCommandCubicBezierCurve *VectorPathCommandCubicBezierCurve `json:"vectorPathCommandCubicBezierCurve,omitempty"`
	VectorPathCommandLine             *VectorPathCommandLine             `json:"vectorPathCommandLine,omitempty"`
}

func (t *VectorPathCommand) Type() string {
	return t.TypeStr
}

func (t *VectorPathCommand) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *VectorPathCommand) GetExtra() string {
	return ""
}

func (t *VectorPathCommand) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "vectorPathCommandCubicBezierCurve":
		t.VectorPathCommandCubicBezierCurve = new(VectorPathCommandCubicBezierCurve)
		return json.Unmarshal(b, t.VectorPathCommandCubicBezierCurve)
	case "vectorPathCommandLine":
		t.VectorPathCommandLine = new(VectorPathCommandLine)
		return json.Unmarshal(b, t.VectorPathCommandLine)
	}
	return nil
}

func (t *VectorPathCommand) MarshalJSON() ([]byte, error) {
	if t.VectorPathCommandCubicBezierCurve != nil {
		return json.Marshal(t.VectorPathCommandCubicBezierCurve)
	}
	if t.VectorPathCommandLine != nil {
		return json.Marshal(t.VectorPathCommandLine)
	}
	return nil, fmt.Errorf("no value set for VectorPathCommand")
}

// WebAppOpenMode Describes mode in which a Web App is opened
type WebAppOpenMode struct {
	TypeStr                  string                    `json:"@type"`
	WebAppOpenModeCompact    *WebAppOpenModeCompact    `json:"webAppOpenModeCompact,omitempty"`
	WebAppOpenModeFullScreen *WebAppOpenModeFullScreen `json:"webAppOpenModeFullScreen,omitempty"`
	WebAppOpenModeFullSize   *WebAppOpenModeFullSize   `json:"webAppOpenModeFullSize,omitempty"`
}

func (t *WebAppOpenMode) Type() string {
	return t.TypeStr
}

func (t *WebAppOpenMode) SetExtra(extra string) {
	// Extra not supported on abstract class directly, or no op
}

func (t *WebAppOpenMode) GetExtra() string {
	return ""
}

func (t *WebAppOpenMode) UnmarshalJSON(b []byte) error {
	var typeObj struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(b, &typeObj); err != nil {
		return err
	}
	t.TypeStr = typeObj.Type
	switch typeObj.Type {
	case "webAppOpenModeCompact":
		t.WebAppOpenModeCompact = new(WebAppOpenModeCompact)
		return json.Unmarshal(b, t.WebAppOpenModeCompact)
	case "webAppOpenModeFullScreen":
		t.WebAppOpenModeFullScreen = new(WebAppOpenModeFullScreen)
		return json.Unmarshal(b, t.WebAppOpenModeFullScreen)
	case "webAppOpenModeFullSize":
		t.WebAppOpenModeFullSize = new(WebAppOpenModeFullSize)
		return json.Unmarshal(b, t.WebAppOpenModeFullSize)
	}
	return nil
}

func (t *WebAppOpenMode) MarshalJSON() ([]byte, error) {
	if t.WebAppOpenModeCompact != nil {
		return json.Marshal(t.WebAppOpenModeCompact)
	}
	if t.WebAppOpenModeFullScreen != nil {
		return json.Marshal(t.WebAppOpenModeFullScreen)
	}
	if t.WebAppOpenModeFullSize != nil {
		return json.Marshal(t.WebAppOpenModeFullSize)
	}
	return nil, fmt.Errorf("no value set for WebAppOpenMode")
}
