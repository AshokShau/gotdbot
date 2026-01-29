package types

import "encoding/json"
import "fmt"

// TlObject is the interface that all TDLib types satisfy
type TlObject interface {
	Type() string
	SetExtra(extra string)
	GetExtra() string
}

// NotificationGroupType Describes the type of notifications in a notification group
type NotificationGroupType struct {
	NotificationGroupTypeMessages   *NotificationGroupTypeMessages   `json:"notificationGroupTypeMessages,omitempty"`
	NotificationGroupTypeMentions   *NotificationGroupTypeMentions   `json:"notificationGroupTypeMentions,omitempty"`
	NotificationGroupTypeSecretChat *NotificationGroupTypeSecretChat `json:"notificationGroupTypeSecretChat,omitempty"`
	NotificationGroupTypeCalls      *NotificationGroupTypeCalls      `json:"notificationGroupTypeCalls,omitempty"`
}

func (t *NotificationGroupType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "notificationGroupTypeMessages":
		t.NotificationGroupTypeMessages = new(NotificationGroupTypeMessages)
		return json.Unmarshal(b, t.NotificationGroupTypeMessages)
	case "notificationGroupTypeMentions":
		t.NotificationGroupTypeMentions = new(NotificationGroupTypeMentions)
		return json.Unmarshal(b, t.NotificationGroupTypeMentions)
	case "notificationGroupTypeSecretChat":
		t.NotificationGroupTypeSecretChat = new(NotificationGroupTypeSecretChat)
		return json.Unmarshal(b, t.NotificationGroupTypeSecretChat)
	case "notificationGroupTypeCalls":
		t.NotificationGroupTypeCalls = new(NotificationGroupTypeCalls)
		return json.Unmarshal(b, t.NotificationGroupTypeCalls)
	}
	return nil
}

func (t *NotificationGroupType) MarshalJSON() ([]byte, error) {
	if t.NotificationGroupTypeMessages != nil {
		return json.Marshal(t.NotificationGroupTypeMessages)
	}
	if t.NotificationGroupTypeMentions != nil {
		return json.Marshal(t.NotificationGroupTypeMentions)
	}
	if t.NotificationGroupTypeSecretChat != nil {
		return json.Marshal(t.NotificationGroupTypeSecretChat)
	}
	if t.NotificationGroupTypeCalls != nil {
		return json.Marshal(t.NotificationGroupTypeCalls)
	}
	return nil, fmt.Errorf("no value set for NotificationGroupType")
}

// JsonValue Represents a JSON value
type JsonValue struct {
	JsonValueNull    *JsonValueNull    `json:"jsonValueNull,omitempty"`
	JsonValueBoolean *JsonValueBoolean `json:"jsonValueBoolean,omitempty"`
	JsonValueNumber  *JsonValueNumber  `json:"jsonValueNumber,omitempty"`
	JsonValueString  *JsonValueString  `json:"jsonValueString,omitempty"`
	JsonValueArray   *JsonValueArray   `json:"jsonValueArray,omitempty"`
	JsonValueObject  *JsonValueObject  `json:"jsonValueObject,omitempty"`
}

func (t *JsonValue) Type() string {
	return ""
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
	switch typeObj.Type {
	case "jsonValueNull":
		t.JsonValueNull = new(JsonValueNull)
		return json.Unmarshal(b, t.JsonValueNull)
	case "jsonValueBoolean":
		t.JsonValueBoolean = new(JsonValueBoolean)
		return json.Unmarshal(b, t.JsonValueBoolean)
	case "jsonValueNumber":
		t.JsonValueNumber = new(JsonValueNumber)
		return json.Unmarshal(b, t.JsonValueNumber)
	case "jsonValueString":
		t.JsonValueString = new(JsonValueString)
		return json.Unmarshal(b, t.JsonValueString)
	case "jsonValueArray":
		t.JsonValueArray = new(JsonValueArray)
		return json.Unmarshal(b, t.JsonValueArray)
	case "jsonValueObject":
		t.JsonValueObject = new(JsonValueObject)
		return json.Unmarshal(b, t.JsonValueObject)
	}
	return nil
}

func (t *JsonValue) MarshalJSON() ([]byte, error) {
	if t.JsonValueNull != nil {
		return json.Marshal(t.JsonValueNull)
	}
	if t.JsonValueBoolean != nil {
		return json.Marshal(t.JsonValueBoolean)
	}
	if t.JsonValueNumber != nil {
		return json.Marshal(t.JsonValueNumber)
	}
	if t.JsonValueString != nil {
		return json.Marshal(t.JsonValueString)
	}
	if t.JsonValueArray != nil {
		return json.Marshal(t.JsonValueArray)
	}
	if t.JsonValueObject != nil {
		return json.Marshal(t.JsonValueObject)
	}
	return nil, fmt.Errorf("no value set for JsonValue")
}

// CanSendMessageToUserResult Describes result of canSendMessageToUser
type CanSendMessageToUserResult struct {
	CanSendMessageToUserResultOk                    *CanSendMessageToUserResultOk                    `json:"canSendMessageToUserResultOk,omitempty"`
	CanSendMessageToUserResultUserHasPaidMessages   *CanSendMessageToUserResultUserHasPaidMessages   `json:"canSendMessageToUserResultUserHasPaidMessages,omitempty"`
	CanSendMessageToUserResultUserIsDeleted         *CanSendMessageToUserResultUserIsDeleted         `json:"canSendMessageToUserResultUserIsDeleted,omitempty"`
	CanSendMessageToUserResultUserRestrictsNewChats *CanSendMessageToUserResultUserRestrictsNewChats `json:"canSendMessageToUserResultUserRestrictsNewChats,omitempty"`
}

func (t *CanSendMessageToUserResult) Type() string {
	return ""
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

// ReactionNotificationSource Describes sources of reactions for which notifications will be shown
type ReactionNotificationSource struct {
	ReactionNotificationSourceNone     *ReactionNotificationSourceNone     `json:"reactionNotificationSourceNone,omitempty"`
	ReactionNotificationSourceContacts *ReactionNotificationSourceContacts `json:"reactionNotificationSourceContacts,omitempty"`
	ReactionNotificationSourceAll      *ReactionNotificationSourceAll      `json:"reactionNotificationSourceAll,omitempty"`
}

func (t *ReactionNotificationSource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "reactionNotificationSourceNone":
		t.ReactionNotificationSourceNone = new(ReactionNotificationSourceNone)
		return json.Unmarshal(b, t.ReactionNotificationSourceNone)
	case "reactionNotificationSourceContacts":
		t.ReactionNotificationSourceContacts = new(ReactionNotificationSourceContacts)
		return json.Unmarshal(b, t.ReactionNotificationSourceContacts)
	case "reactionNotificationSourceAll":
		t.ReactionNotificationSourceAll = new(ReactionNotificationSourceAll)
		return json.Unmarshal(b, t.ReactionNotificationSourceAll)
	}
	return nil
}

func (t *ReactionNotificationSource) MarshalJSON() ([]byte, error) {
	if t.ReactionNotificationSourceNone != nil {
		return json.Marshal(t.ReactionNotificationSourceNone)
	}
	if t.ReactionNotificationSourceContacts != nil {
		return json.Marshal(t.ReactionNotificationSourceContacts)
	}
	if t.ReactionNotificationSourceAll != nil {
		return json.Marshal(t.ReactionNotificationSourceAll)
	}
	return nil, fmt.Errorf("no value set for ReactionNotificationSource")
}

// InputMessageContent The content of a message to send
type InputMessageContent struct {
	InputMessageText      *InputMessageText      `json:"inputMessageText,omitempty"`
	InputMessageAnimation *InputMessageAnimation `json:"inputMessageAnimation,omitempty"`
	InputMessageAudio     *InputMessageAudio     `json:"inputMessageAudio,omitempty"`
	InputMessageDocument  *InputMessageDocument  `json:"inputMessageDocument,omitempty"`
	InputMessagePaidMedia *InputMessagePaidMedia `json:"inputMessagePaidMedia,omitempty"`
	InputMessagePhoto     *InputMessagePhoto     `json:"inputMessagePhoto,omitempty"`
	InputMessageSticker   *InputMessageSticker   `json:"inputMessageSticker,omitempty"`
	InputMessageVideo     *InputMessageVideo     `json:"inputMessageVideo,omitempty"`
	InputMessageVideoNote *InputMessageVideoNote `json:"inputMessageVideoNote,omitempty"`
	InputMessageVoiceNote *InputMessageVoiceNote `json:"inputMessageVoiceNote,omitempty"`
	InputMessageLocation  *InputMessageLocation  `json:"inputMessageLocation,omitempty"`
	InputMessageVenue     *InputMessageVenue     `json:"inputMessageVenue,omitempty"`
	InputMessageContact   *InputMessageContact   `json:"inputMessageContact,omitempty"`
	InputMessageDice      *InputMessageDice      `json:"inputMessageDice,omitempty"`
	InputMessageGame      *InputMessageGame      `json:"inputMessageGame,omitempty"`
	InputMessageInvoice   *InputMessageInvoice   `json:"inputMessageInvoice,omitempty"`
	InputMessagePoll      *InputMessagePoll      `json:"inputMessagePoll,omitempty"`
	InputMessageStakeDice *InputMessageStakeDice `json:"inputMessageStakeDice,omitempty"`
	InputMessageStory     *InputMessageStory     `json:"inputMessageStory,omitempty"`
	InputMessageChecklist *InputMessageChecklist `json:"inputMessageChecklist,omitempty"`
	InputMessageForwarded *InputMessageForwarded `json:"inputMessageForwarded,omitempty"`
}

func (t *InputMessageContent) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputMessageText":
		t.InputMessageText = new(InputMessageText)
		return json.Unmarshal(b, t.InputMessageText)
	case "inputMessageAnimation":
		t.InputMessageAnimation = new(InputMessageAnimation)
		return json.Unmarshal(b, t.InputMessageAnimation)
	case "inputMessageAudio":
		t.InputMessageAudio = new(InputMessageAudio)
		return json.Unmarshal(b, t.InputMessageAudio)
	case "inputMessageDocument":
		t.InputMessageDocument = new(InputMessageDocument)
		return json.Unmarshal(b, t.InputMessageDocument)
	case "inputMessagePaidMedia":
		t.InputMessagePaidMedia = new(InputMessagePaidMedia)
		return json.Unmarshal(b, t.InputMessagePaidMedia)
	case "inputMessagePhoto":
		t.InputMessagePhoto = new(InputMessagePhoto)
		return json.Unmarshal(b, t.InputMessagePhoto)
	case "inputMessageSticker":
		t.InputMessageSticker = new(InputMessageSticker)
		return json.Unmarshal(b, t.InputMessageSticker)
	case "inputMessageVideo":
		t.InputMessageVideo = new(InputMessageVideo)
		return json.Unmarshal(b, t.InputMessageVideo)
	case "inputMessageVideoNote":
		t.InputMessageVideoNote = new(InputMessageVideoNote)
		return json.Unmarshal(b, t.InputMessageVideoNote)
	case "inputMessageVoiceNote":
		t.InputMessageVoiceNote = new(InputMessageVoiceNote)
		return json.Unmarshal(b, t.InputMessageVoiceNote)
	case "inputMessageLocation":
		t.InputMessageLocation = new(InputMessageLocation)
		return json.Unmarshal(b, t.InputMessageLocation)
	case "inputMessageVenue":
		t.InputMessageVenue = new(InputMessageVenue)
		return json.Unmarshal(b, t.InputMessageVenue)
	case "inputMessageContact":
		t.InputMessageContact = new(InputMessageContact)
		return json.Unmarshal(b, t.InputMessageContact)
	case "inputMessageDice":
		t.InputMessageDice = new(InputMessageDice)
		return json.Unmarshal(b, t.InputMessageDice)
	case "inputMessageGame":
		t.InputMessageGame = new(InputMessageGame)
		return json.Unmarshal(b, t.InputMessageGame)
	case "inputMessageInvoice":
		t.InputMessageInvoice = new(InputMessageInvoice)
		return json.Unmarshal(b, t.InputMessageInvoice)
	case "inputMessagePoll":
		t.InputMessagePoll = new(InputMessagePoll)
		return json.Unmarshal(b, t.InputMessagePoll)
	case "inputMessageStakeDice":
		t.InputMessageStakeDice = new(InputMessageStakeDice)
		return json.Unmarshal(b, t.InputMessageStakeDice)
	case "inputMessageStory":
		t.InputMessageStory = new(InputMessageStory)
		return json.Unmarshal(b, t.InputMessageStory)
	case "inputMessageChecklist":
		t.InputMessageChecklist = new(InputMessageChecklist)
		return json.Unmarshal(b, t.InputMessageChecklist)
	case "inputMessageForwarded":
		t.InputMessageForwarded = new(InputMessageForwarded)
		return json.Unmarshal(b, t.InputMessageForwarded)
	}
	return nil
}

func (t *InputMessageContent) MarshalJSON() ([]byte, error) {
	if t.InputMessageText != nil {
		return json.Marshal(t.InputMessageText)
	}
	if t.InputMessageAnimation != nil {
		return json.Marshal(t.InputMessageAnimation)
	}
	if t.InputMessageAudio != nil {
		return json.Marshal(t.InputMessageAudio)
	}
	if t.InputMessageDocument != nil {
		return json.Marshal(t.InputMessageDocument)
	}
	if t.InputMessagePaidMedia != nil {
		return json.Marshal(t.InputMessagePaidMedia)
	}
	if t.InputMessagePhoto != nil {
		return json.Marshal(t.InputMessagePhoto)
	}
	if t.InputMessageSticker != nil {
		return json.Marshal(t.InputMessageSticker)
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
	if t.InputMessageLocation != nil {
		return json.Marshal(t.InputMessageLocation)
	}
	if t.InputMessageVenue != nil {
		return json.Marshal(t.InputMessageVenue)
	}
	if t.InputMessageContact != nil {
		return json.Marshal(t.InputMessageContact)
	}
	if t.InputMessageDice != nil {
		return json.Marshal(t.InputMessageDice)
	}
	if t.InputMessageGame != nil {
		return json.Marshal(t.InputMessageGame)
	}
	if t.InputMessageInvoice != nil {
		return json.Marshal(t.InputMessageInvoice)
	}
	if t.InputMessagePoll != nil {
		return json.Marshal(t.InputMessagePoll)
	}
	if t.InputMessageStakeDice != nil {
		return json.Marshal(t.InputMessageStakeDice)
	}
	if t.InputMessageStory != nil {
		return json.Marshal(t.InputMessageStory)
	}
	if t.InputMessageChecklist != nil {
		return json.Marshal(t.InputMessageChecklist)
	}
	if t.InputMessageForwarded != nil {
		return json.Marshal(t.InputMessageForwarded)
	}
	return nil, fmt.Errorf("no value set for InputMessageContent")
}

// StoryList Describes a list of stories
type StoryList struct {
	StoryListMain    *StoryListMain    `json:"storyListMain,omitempty"`
	StoryListArchive *StoryListArchive `json:"storyListArchive,omitempty"`
}

func (t *StoryList) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyListMain":
		t.StoryListMain = new(StoryListMain)
		return json.Unmarshal(b, t.StoryListMain)
	case "storyListArchive":
		t.StoryListArchive = new(StoryListArchive)
		return json.Unmarshal(b, t.StoryListArchive)
	}
	return nil
}

func (t *StoryList) MarshalJSON() ([]byte, error) {
	if t.StoryListMain != nil {
		return json.Marshal(t.StoryListMain)
	}
	if t.StoryListArchive != nil {
		return json.Marshal(t.StoryListArchive)
	}
	return nil, fmt.Errorf("no value set for StoryList")
}

// ChatStatisticsObjectType Describes type of object, for which statistics are provided
type ChatStatisticsObjectType struct {
	ChatStatisticsObjectTypeMessage *ChatStatisticsObjectTypeMessage `json:"chatStatisticsObjectTypeMessage,omitempty"`
	ChatStatisticsObjectTypeStory   *ChatStatisticsObjectTypeStory   `json:"chatStatisticsObjectTypeStory,omitempty"`
}

func (t *ChatStatisticsObjectType) Type() string {
	return ""
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

// RevenueWithdrawalState Describes state of a revenue withdrawal
type RevenueWithdrawalState struct {
	RevenueWithdrawalStatePending   *RevenueWithdrawalStatePending   `json:"revenueWithdrawalStatePending,omitempty"`
	RevenueWithdrawalStateSucceeded *RevenueWithdrawalStateSucceeded `json:"revenueWithdrawalStateSucceeded,omitempty"`
	RevenueWithdrawalStateFailed    *RevenueWithdrawalStateFailed    `json:"revenueWithdrawalStateFailed,omitempty"`
}

func (t *RevenueWithdrawalState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "revenueWithdrawalStatePending":
		t.RevenueWithdrawalStatePending = new(RevenueWithdrawalStatePending)
		return json.Unmarshal(b, t.RevenueWithdrawalStatePending)
	case "revenueWithdrawalStateSucceeded":
		t.RevenueWithdrawalStateSucceeded = new(RevenueWithdrawalStateSucceeded)
		return json.Unmarshal(b, t.RevenueWithdrawalStateSucceeded)
	case "revenueWithdrawalStateFailed":
		t.RevenueWithdrawalStateFailed = new(RevenueWithdrawalStateFailed)
		return json.Unmarshal(b, t.RevenueWithdrawalStateFailed)
	}
	return nil
}

func (t *RevenueWithdrawalState) MarshalJSON() ([]byte, error) {
	if t.RevenueWithdrawalStatePending != nil {
		return json.Marshal(t.RevenueWithdrawalStatePending)
	}
	if t.RevenueWithdrawalStateSucceeded != nil {
		return json.Marshal(t.RevenueWithdrawalStateSucceeded)
	}
	if t.RevenueWithdrawalStateFailed != nil {
		return json.Marshal(t.RevenueWithdrawalStateFailed)
	}
	return nil, fmt.Errorf("no value set for RevenueWithdrawalState")
}

// LogStream Describes a stream to which TDLib internal log is written
type LogStream struct {
	LogStreamDefault *LogStreamDefault `json:"logStreamDefault,omitempty"`
	LogStreamFile    *LogStreamFile    `json:"logStreamFile,omitempty"`
	LogStreamEmpty   *LogStreamEmpty   `json:"logStreamEmpty,omitempty"`
}

func (t *LogStream) Type() string {
	return ""
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
	switch typeObj.Type {
	case "logStreamDefault":
		t.LogStreamDefault = new(LogStreamDefault)
		return json.Unmarshal(b, t.LogStreamDefault)
	case "logStreamFile":
		t.LogStreamFile = new(LogStreamFile)
		return json.Unmarshal(b, t.LogStreamFile)
	case "logStreamEmpty":
		t.LogStreamEmpty = new(LogStreamEmpty)
		return json.Unmarshal(b, t.LogStreamEmpty)
	}
	return nil
}

func (t *LogStream) MarshalJSON() ([]byte, error) {
	if t.LogStreamDefault != nil {
		return json.Marshal(t.LogStreamDefault)
	}
	if t.LogStreamFile != nil {
		return json.Marshal(t.LogStreamFile)
	}
	if t.LogStreamEmpty != nil {
		return json.Marshal(t.LogStreamEmpty)
	}
	return nil, fmt.Errorf("no value set for LogStream")
}

// BuiltInTheme Describes a built-in theme of an official app
type BuiltInTheme struct {
	BuiltInThemeClassic *BuiltInThemeClassic `json:"builtInThemeClassic,omitempty"`
	BuiltInThemeDay     *BuiltInThemeDay     `json:"builtInThemeDay,omitempty"`
	BuiltInThemeNight   *BuiltInThemeNight   `json:"builtInThemeNight,omitempty"`
	BuiltInThemeTinted  *BuiltInThemeTinted  `json:"builtInThemeTinted,omitempty"`
	BuiltInThemeArctic  *BuiltInThemeArctic  `json:"builtInThemeArctic,omitempty"`
}

func (t *BuiltInTheme) Type() string {
	return ""
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
	switch typeObj.Type {
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
	case "builtInThemeArctic":
		t.BuiltInThemeArctic = new(BuiltInThemeArctic)
		return json.Unmarshal(b, t.BuiltInThemeArctic)
	}
	return nil
}

func (t *BuiltInTheme) MarshalJSON() ([]byte, error) {
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
	if t.BuiltInThemeArctic != nil {
		return json.Marshal(t.BuiltInThemeArctic)
	}
	return nil, fmt.Errorf("no value set for BuiltInTheme")
}

// UserType Represents the type of user. The following types are possible: regular users, deleted users and bots
type UserType struct {
	UserTypeRegular *UserTypeRegular `json:"userTypeRegular,omitempty"`
	UserTypeDeleted *UserTypeDeleted `json:"userTypeDeleted,omitempty"`
	UserTypeBot     *UserTypeBot     `json:"userTypeBot,omitempty"`
	UserTypeUnknown *UserTypeUnknown `json:"userTypeUnknown,omitempty"`
}

func (t *UserType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "userTypeRegular":
		t.UserTypeRegular = new(UserTypeRegular)
		return json.Unmarshal(b, t.UserTypeRegular)
	case "userTypeDeleted":
		t.UserTypeDeleted = new(UserTypeDeleted)
		return json.Unmarshal(b, t.UserTypeDeleted)
	case "userTypeBot":
		t.UserTypeBot = new(UserTypeBot)
		return json.Unmarshal(b, t.UserTypeBot)
	case "userTypeUnknown":
		t.UserTypeUnknown = new(UserTypeUnknown)
		return json.Unmarshal(b, t.UserTypeUnknown)
	}
	return nil
}

func (t *UserType) MarshalJSON() ([]byte, error) {
	if t.UserTypeRegular != nil {
		return json.Marshal(t.UserTypeRegular)
	}
	if t.UserTypeDeleted != nil {
		return json.Marshal(t.UserTypeDeleted)
	}
	if t.UserTypeBot != nil {
		return json.Marshal(t.UserTypeBot)
	}
	if t.UserTypeUnknown != nil {
		return json.Marshal(t.UserTypeUnknown)
	}
	return nil, fmt.Errorf("no value set for UserType")
}

// BusinessAwayMessageSchedule Describes conditions for sending of away messages by a Telegram Business account
type BusinessAwayMessageSchedule struct {
	BusinessAwayMessageScheduleAlways                *BusinessAwayMessageScheduleAlways                `json:"businessAwayMessageScheduleAlways,omitempty"`
	BusinessAwayMessageScheduleOutsideOfOpeningHours *BusinessAwayMessageScheduleOutsideOfOpeningHours `json:"businessAwayMessageScheduleOutsideOfOpeningHours,omitempty"`
	BusinessAwayMessageScheduleCustom                *BusinessAwayMessageScheduleCustom                `json:"businessAwayMessageScheduleCustom,omitempty"`
}

func (t *BusinessAwayMessageSchedule) Type() string {
	return ""
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
	switch typeObj.Type {
	case "businessAwayMessageScheduleAlways":
		t.BusinessAwayMessageScheduleAlways = new(BusinessAwayMessageScheduleAlways)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleAlways)
	case "businessAwayMessageScheduleOutsideOfOpeningHours":
		t.BusinessAwayMessageScheduleOutsideOfOpeningHours = new(BusinessAwayMessageScheduleOutsideOfOpeningHours)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleOutsideOfOpeningHours)
	case "businessAwayMessageScheduleCustom":
		t.BusinessAwayMessageScheduleCustom = new(BusinessAwayMessageScheduleCustom)
		return json.Unmarshal(b, t.BusinessAwayMessageScheduleCustom)
	}
	return nil
}

func (t *BusinessAwayMessageSchedule) MarshalJSON() ([]byte, error) {
	if t.BusinessAwayMessageScheduleAlways != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleAlways)
	}
	if t.BusinessAwayMessageScheduleOutsideOfOpeningHours != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleOutsideOfOpeningHours)
	}
	if t.BusinessAwayMessageScheduleCustom != nil {
		return json.Marshal(t.BusinessAwayMessageScheduleCustom)
	}
	return nil, fmt.Errorf("no value set for BusinessAwayMessageSchedule")
}

// AffiliateProgramSortOrder Describes the order of the found affiliate programs
type AffiliateProgramSortOrder struct {
	AffiliateProgramSortOrderProfitability *AffiliateProgramSortOrderProfitability `json:"affiliateProgramSortOrderProfitability,omitempty"`
	AffiliateProgramSortOrderCreationDate  *AffiliateProgramSortOrderCreationDate  `json:"affiliateProgramSortOrderCreationDate,omitempty"`
	AffiliateProgramSortOrderRevenue       *AffiliateProgramSortOrderRevenue       `json:"affiliateProgramSortOrderRevenue,omitempty"`
}

func (t *AffiliateProgramSortOrder) Type() string {
	return ""
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
	switch typeObj.Type {
	case "affiliateProgramSortOrderProfitability":
		t.AffiliateProgramSortOrderProfitability = new(AffiliateProgramSortOrderProfitability)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderProfitability)
	case "affiliateProgramSortOrderCreationDate":
		t.AffiliateProgramSortOrderCreationDate = new(AffiliateProgramSortOrderCreationDate)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderCreationDate)
	case "affiliateProgramSortOrderRevenue":
		t.AffiliateProgramSortOrderRevenue = new(AffiliateProgramSortOrderRevenue)
		return json.Unmarshal(b, t.AffiliateProgramSortOrderRevenue)
	}
	return nil
}

func (t *AffiliateProgramSortOrder) MarshalJSON() ([]byte, error) {
	if t.AffiliateProgramSortOrderProfitability != nil {
		return json.Marshal(t.AffiliateProgramSortOrderProfitability)
	}
	if t.AffiliateProgramSortOrderCreationDate != nil {
		return json.Marshal(t.AffiliateProgramSortOrderCreationDate)
	}
	if t.AffiliateProgramSortOrderRevenue != nil {
		return json.Marshal(t.AffiliateProgramSortOrderRevenue)
	}
	return nil, fmt.Errorf("no value set for AffiliateProgramSortOrder")
}

// MessageContent Contains the content of a message
type MessageContent struct {
	MessageText                              *MessageText                              `json:"messageText,omitempty"`
	MessageAnimation                         *MessageAnimation                         `json:"messageAnimation,omitempty"`
	MessageAudio                             *MessageAudio                             `json:"messageAudio,omitempty"`
	MessageDocument                          *MessageDocument                          `json:"messageDocument,omitempty"`
	MessagePaidMedia                         *MessagePaidMedia                         `json:"messagePaidMedia,omitempty"`
	MessagePhoto                             *MessagePhoto                             `json:"messagePhoto,omitempty"`
	MessageSticker                           *MessageSticker                           `json:"messageSticker,omitempty"`
	MessageVideo                             *MessageVideo                             `json:"messageVideo,omitempty"`
	MessageVideoNote                         *MessageVideoNote                         `json:"messageVideoNote,omitempty"`
	MessageVoiceNote                         *MessageVoiceNote                         `json:"messageVoiceNote,omitempty"`
	MessageExpiredPhoto                      *MessageExpiredPhoto                      `json:"messageExpiredPhoto,omitempty"`
	MessageExpiredVideo                      *MessageExpiredVideo                      `json:"messageExpiredVideo,omitempty"`
	MessageExpiredVideoNote                  *MessageExpiredVideoNote                  `json:"messageExpiredVideoNote,omitempty"`
	MessageExpiredVoiceNote                  *MessageExpiredVoiceNote                  `json:"messageExpiredVoiceNote,omitempty"`
	MessageLocation                          *MessageLocation                          `json:"messageLocation,omitempty"`
	MessageVenue                             *MessageVenue                             `json:"messageVenue,omitempty"`
	MessageContact                           *MessageContact                           `json:"messageContact,omitempty"`
	MessageAnimatedEmoji                     *MessageAnimatedEmoji                     `json:"messageAnimatedEmoji,omitempty"`
	MessageDice                              *MessageDice                              `json:"messageDice,omitempty"`
	MessageGame                              *MessageGame                              `json:"messageGame,omitempty"`
	MessagePoll                              *MessagePoll                              `json:"messagePoll,omitempty"`
	MessageStakeDice                         *MessageStakeDice                         `json:"messageStakeDice,omitempty"`
	MessageStory                             *MessageStory                             `json:"messageStory,omitempty"`
	MessageChecklist                         *MessageChecklist                         `json:"messageChecklist,omitempty"`
	MessageInvoice                           *MessageInvoice                           `json:"messageInvoice,omitempty"`
	MessageCall                              *MessageCall                              `json:"messageCall,omitempty"`
	MessageGroupCall                         *MessageGroupCall                         `json:"messageGroupCall,omitempty"`
	MessageVideoChatScheduled                *MessageVideoChatScheduled                `json:"messageVideoChatScheduled,omitempty"`
	MessageVideoChatStarted                  *MessageVideoChatStarted                  `json:"messageVideoChatStarted,omitempty"`
	MessageVideoChatEnded                    *MessageVideoChatEnded                    `json:"messageVideoChatEnded,omitempty"`
	MessageInviteVideoChatParticipants       *MessageInviteVideoChatParticipants       `json:"messageInviteVideoChatParticipants,omitempty"`
	MessageBasicGroupChatCreate              *MessageBasicGroupChatCreate              `json:"messageBasicGroupChatCreate,omitempty"`
	MessageSupergroupChatCreate              *MessageSupergroupChatCreate              `json:"messageSupergroupChatCreate,omitempty"`
	MessageChatChangeTitle                   *MessageChatChangeTitle                   `json:"messageChatChangeTitle,omitempty"`
	MessageChatChangePhoto                   *MessageChatChangePhoto                   `json:"messageChatChangePhoto,omitempty"`
	MessageChatDeletePhoto                   *MessageChatDeletePhoto                   `json:"messageChatDeletePhoto,omitempty"`
	MessageChatAddMembers                    *MessageChatAddMembers                    `json:"messageChatAddMembers,omitempty"`
	MessageChatJoinByLink                    *MessageChatJoinByLink                    `json:"messageChatJoinByLink,omitempty"`
	MessageChatJoinByRequest                 *MessageChatJoinByRequest                 `json:"messageChatJoinByRequest,omitempty"`
	MessageChatDeleteMember                  *MessageChatDeleteMember                  `json:"messageChatDeleteMember,omitempty"`
	MessageChatUpgradeTo                     *MessageChatUpgradeTo                     `json:"messageChatUpgradeTo,omitempty"`
	MessageChatUpgradeFrom                   *MessageChatUpgradeFrom                   `json:"messageChatUpgradeFrom,omitempty"`
	MessagePinMessage                        *MessagePinMessage                        `json:"messagePinMessage,omitempty"`
	MessageScreenshotTaken                   *MessageScreenshotTaken                   `json:"messageScreenshotTaken,omitempty"`
	MessageChatSetBackground                 *MessageChatSetBackground                 `json:"messageChatSetBackground,omitempty"`
	MessageChatSetTheme                      *MessageChatSetTheme                      `json:"messageChatSetTheme,omitempty"`
	MessageChatSetMessageAutoDeleteTime      *MessageChatSetMessageAutoDeleteTime      `json:"messageChatSetMessageAutoDeleteTime,omitempty"`
	MessageChatBoost                         *MessageChatBoost                         `json:"messageChatBoost,omitempty"`
	MessageForumTopicCreated                 *MessageForumTopicCreated                 `json:"messageForumTopicCreated,omitempty"`
	MessageForumTopicEdited                  *MessageForumTopicEdited                  `json:"messageForumTopicEdited,omitempty"`
	MessageForumTopicIsClosedToggled         *MessageForumTopicIsClosedToggled         `json:"messageForumTopicIsClosedToggled,omitempty"`
	MessageForumTopicIsHiddenToggled         *MessageForumTopicIsHiddenToggled         `json:"messageForumTopicIsHiddenToggled,omitempty"`
	MessageSuggestProfilePhoto               *MessageSuggestProfilePhoto               `json:"messageSuggestProfilePhoto,omitempty"`
	MessageSuggestBirthdate                  *MessageSuggestBirthdate                  `json:"messageSuggestBirthdate,omitempty"`
	MessageCustomServiceAction               *MessageCustomServiceAction               `json:"messageCustomServiceAction,omitempty"`
	MessageGameScore                         *MessageGameScore                         `json:"messageGameScore,omitempty"`
	MessagePaymentSuccessful                 *MessagePaymentSuccessful                 `json:"messagePaymentSuccessful,omitempty"`
	MessagePaymentSuccessfulBot              *MessagePaymentSuccessfulBot              `json:"messagePaymentSuccessfulBot,omitempty"`
	MessagePaymentRefunded                   *MessagePaymentRefunded                   `json:"messagePaymentRefunded,omitempty"`
	MessageGiftedPremium                     *MessageGiftedPremium                     `json:"messageGiftedPremium,omitempty"`
	MessagePremiumGiftCode                   *MessagePremiumGiftCode                   `json:"messagePremiumGiftCode,omitempty"`
	MessageGiveawayCreated                   *MessageGiveawayCreated                   `json:"messageGiveawayCreated,omitempty"`
	MessageGiveaway                          *MessageGiveaway                          `json:"messageGiveaway,omitempty"`
	MessageGiveawayCompleted                 *MessageGiveawayCompleted                 `json:"messageGiveawayCompleted,omitempty"`
	MessageGiveawayWinners                   *MessageGiveawayWinners                   `json:"messageGiveawayWinners,omitempty"`
	MessageGiftedStars                       *MessageGiftedStars                       `json:"messageGiftedStars,omitempty"`
	MessageGiftedTon                         *MessageGiftedTon                         `json:"messageGiftedTon,omitempty"`
	MessageGiveawayPrizeStars                *MessageGiveawayPrizeStars                `json:"messageGiveawayPrizeStars,omitempty"`
	MessageGift                              *MessageGift                              `json:"messageGift,omitempty"`
	MessageUpgradedGift                      *MessageUpgradedGift                      `json:"messageUpgradedGift,omitempty"`
	MessageRefundedUpgradedGift              *MessageRefundedUpgradedGift              `json:"messageRefundedUpgradedGift,omitempty"`
	MessageUpgradedGiftPurchaseOffer         *MessageUpgradedGiftPurchaseOffer         `json:"messageUpgradedGiftPurchaseOffer,omitempty"`
	MessageUpgradedGiftPurchaseOfferRejected *MessageUpgradedGiftPurchaseOfferRejected `json:"messageUpgradedGiftPurchaseOfferRejected,omitempty"`
	MessagePaidMessagesRefunded              *MessagePaidMessagesRefunded              `json:"messagePaidMessagesRefunded,omitempty"`
	MessagePaidMessagePriceChanged           *MessagePaidMessagePriceChanged           `json:"messagePaidMessagePriceChanged,omitempty"`
	MessageDirectMessagePriceChanged         *MessageDirectMessagePriceChanged         `json:"messageDirectMessagePriceChanged,omitempty"`
	MessageChecklistTasksDone                *MessageChecklistTasksDone                `json:"messageChecklistTasksDone,omitempty"`
	MessageChecklistTasksAdded               *MessageChecklistTasksAdded               `json:"messageChecklistTasksAdded,omitempty"`
	MessageSuggestedPostApprovalFailed       *MessageSuggestedPostApprovalFailed       `json:"messageSuggestedPostApprovalFailed,omitempty"`
	MessageSuggestedPostApproved             *MessageSuggestedPostApproved             `json:"messageSuggestedPostApproved,omitempty"`
	MessageSuggestedPostDeclined             *MessageSuggestedPostDeclined             `json:"messageSuggestedPostDeclined,omitempty"`
	MessageSuggestedPostPaid                 *MessageSuggestedPostPaid                 `json:"messageSuggestedPostPaid,omitempty"`
	MessageSuggestedPostRefunded             *MessageSuggestedPostRefunded             `json:"messageSuggestedPostRefunded,omitempty"`
	MessageContactRegistered                 *MessageContactRegistered                 `json:"messageContactRegistered,omitempty"`
	MessageUsersShared                       *MessageUsersShared                       `json:"messageUsersShared,omitempty"`
	MessageChatShared                        *MessageChatShared                        `json:"messageChatShared,omitempty"`
	MessageBotWriteAccessAllowed             *MessageBotWriteAccessAllowed             `json:"messageBotWriteAccessAllowed,omitempty"`
	MessageWebAppDataSent                    *MessageWebAppDataSent                    `json:"messageWebAppDataSent,omitempty"`
	MessageWebAppDataReceived                *MessageWebAppDataReceived                `json:"messageWebAppDataReceived,omitempty"`
	MessagePassportDataSent                  *MessagePassportDataSent                  `json:"messagePassportDataSent,omitempty"`
	MessagePassportDataReceived              *MessagePassportDataReceived              `json:"messagePassportDataReceived,omitempty"`
	MessageProximityAlertTriggered           *MessageProximityAlertTriggered           `json:"messageProximityAlertTriggered,omitempty"`
	MessageUnsupported                       *MessageUnsupported                       `json:"messageUnsupported,omitempty"`
}

func (t *MessageContent) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageText":
		t.MessageText = new(MessageText)
		return json.Unmarshal(b, t.MessageText)
	case "messageAnimation":
		t.MessageAnimation = new(MessageAnimation)
		return json.Unmarshal(b, t.MessageAnimation)
	case "messageAudio":
		t.MessageAudio = new(MessageAudio)
		return json.Unmarshal(b, t.MessageAudio)
	case "messageDocument":
		t.MessageDocument = new(MessageDocument)
		return json.Unmarshal(b, t.MessageDocument)
	case "messagePaidMedia":
		t.MessagePaidMedia = new(MessagePaidMedia)
		return json.Unmarshal(b, t.MessagePaidMedia)
	case "messagePhoto":
		t.MessagePhoto = new(MessagePhoto)
		return json.Unmarshal(b, t.MessagePhoto)
	case "messageSticker":
		t.MessageSticker = new(MessageSticker)
		return json.Unmarshal(b, t.MessageSticker)
	case "messageVideo":
		t.MessageVideo = new(MessageVideo)
		return json.Unmarshal(b, t.MessageVideo)
	case "messageVideoNote":
		t.MessageVideoNote = new(MessageVideoNote)
		return json.Unmarshal(b, t.MessageVideoNote)
	case "messageVoiceNote":
		t.MessageVoiceNote = new(MessageVoiceNote)
		return json.Unmarshal(b, t.MessageVoiceNote)
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
	case "messageLocation":
		t.MessageLocation = new(MessageLocation)
		return json.Unmarshal(b, t.MessageLocation)
	case "messageVenue":
		t.MessageVenue = new(MessageVenue)
		return json.Unmarshal(b, t.MessageVenue)
	case "messageContact":
		t.MessageContact = new(MessageContact)
		return json.Unmarshal(b, t.MessageContact)
	case "messageAnimatedEmoji":
		t.MessageAnimatedEmoji = new(MessageAnimatedEmoji)
		return json.Unmarshal(b, t.MessageAnimatedEmoji)
	case "messageDice":
		t.MessageDice = new(MessageDice)
		return json.Unmarshal(b, t.MessageDice)
	case "messageGame":
		t.MessageGame = new(MessageGame)
		return json.Unmarshal(b, t.MessageGame)
	case "messagePoll":
		t.MessagePoll = new(MessagePoll)
		return json.Unmarshal(b, t.MessagePoll)
	case "messageStakeDice":
		t.MessageStakeDice = new(MessageStakeDice)
		return json.Unmarshal(b, t.MessageStakeDice)
	case "messageStory":
		t.MessageStory = new(MessageStory)
		return json.Unmarshal(b, t.MessageStory)
	case "messageChecklist":
		t.MessageChecklist = new(MessageChecklist)
		return json.Unmarshal(b, t.MessageChecklist)
	case "messageInvoice":
		t.MessageInvoice = new(MessageInvoice)
		return json.Unmarshal(b, t.MessageInvoice)
	case "messageCall":
		t.MessageCall = new(MessageCall)
		return json.Unmarshal(b, t.MessageCall)
	case "messageGroupCall":
		t.MessageGroupCall = new(MessageGroupCall)
		return json.Unmarshal(b, t.MessageGroupCall)
	case "messageVideoChatScheduled":
		t.MessageVideoChatScheduled = new(MessageVideoChatScheduled)
		return json.Unmarshal(b, t.MessageVideoChatScheduled)
	case "messageVideoChatStarted":
		t.MessageVideoChatStarted = new(MessageVideoChatStarted)
		return json.Unmarshal(b, t.MessageVideoChatStarted)
	case "messageVideoChatEnded":
		t.MessageVideoChatEnded = new(MessageVideoChatEnded)
		return json.Unmarshal(b, t.MessageVideoChatEnded)
	case "messageInviteVideoChatParticipants":
		t.MessageInviteVideoChatParticipants = new(MessageInviteVideoChatParticipants)
		return json.Unmarshal(b, t.MessageInviteVideoChatParticipants)
	case "messageBasicGroupChatCreate":
		t.MessageBasicGroupChatCreate = new(MessageBasicGroupChatCreate)
		return json.Unmarshal(b, t.MessageBasicGroupChatCreate)
	case "messageSupergroupChatCreate":
		t.MessageSupergroupChatCreate = new(MessageSupergroupChatCreate)
		return json.Unmarshal(b, t.MessageSupergroupChatCreate)
	case "messageChatChangeTitle":
		t.MessageChatChangeTitle = new(MessageChatChangeTitle)
		return json.Unmarshal(b, t.MessageChatChangeTitle)
	case "messageChatChangePhoto":
		t.MessageChatChangePhoto = new(MessageChatChangePhoto)
		return json.Unmarshal(b, t.MessageChatChangePhoto)
	case "messageChatDeletePhoto":
		t.MessageChatDeletePhoto = new(MessageChatDeletePhoto)
		return json.Unmarshal(b, t.MessageChatDeletePhoto)
	case "messageChatAddMembers":
		t.MessageChatAddMembers = new(MessageChatAddMembers)
		return json.Unmarshal(b, t.MessageChatAddMembers)
	case "messageChatJoinByLink":
		t.MessageChatJoinByLink = new(MessageChatJoinByLink)
		return json.Unmarshal(b, t.MessageChatJoinByLink)
	case "messageChatJoinByRequest":
		t.MessageChatJoinByRequest = new(MessageChatJoinByRequest)
		return json.Unmarshal(b, t.MessageChatJoinByRequest)
	case "messageChatDeleteMember":
		t.MessageChatDeleteMember = new(MessageChatDeleteMember)
		return json.Unmarshal(b, t.MessageChatDeleteMember)
	case "messageChatUpgradeTo":
		t.MessageChatUpgradeTo = new(MessageChatUpgradeTo)
		return json.Unmarshal(b, t.MessageChatUpgradeTo)
	case "messageChatUpgradeFrom":
		t.MessageChatUpgradeFrom = new(MessageChatUpgradeFrom)
		return json.Unmarshal(b, t.MessageChatUpgradeFrom)
	case "messagePinMessage":
		t.MessagePinMessage = new(MessagePinMessage)
		return json.Unmarshal(b, t.MessagePinMessage)
	case "messageScreenshotTaken":
		t.MessageScreenshotTaken = new(MessageScreenshotTaken)
		return json.Unmarshal(b, t.MessageScreenshotTaken)
	case "messageChatSetBackground":
		t.MessageChatSetBackground = new(MessageChatSetBackground)
		return json.Unmarshal(b, t.MessageChatSetBackground)
	case "messageChatSetTheme":
		t.MessageChatSetTheme = new(MessageChatSetTheme)
		return json.Unmarshal(b, t.MessageChatSetTheme)
	case "messageChatSetMessageAutoDeleteTime":
		t.MessageChatSetMessageAutoDeleteTime = new(MessageChatSetMessageAutoDeleteTime)
		return json.Unmarshal(b, t.MessageChatSetMessageAutoDeleteTime)
	case "messageChatBoost":
		t.MessageChatBoost = new(MessageChatBoost)
		return json.Unmarshal(b, t.MessageChatBoost)
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
	case "messageSuggestProfilePhoto":
		t.MessageSuggestProfilePhoto = new(MessageSuggestProfilePhoto)
		return json.Unmarshal(b, t.MessageSuggestProfilePhoto)
	case "messageSuggestBirthdate":
		t.MessageSuggestBirthdate = new(MessageSuggestBirthdate)
		return json.Unmarshal(b, t.MessageSuggestBirthdate)
	case "messageCustomServiceAction":
		t.MessageCustomServiceAction = new(MessageCustomServiceAction)
		return json.Unmarshal(b, t.MessageCustomServiceAction)
	case "messageGameScore":
		t.MessageGameScore = new(MessageGameScore)
		return json.Unmarshal(b, t.MessageGameScore)
	case "messagePaymentSuccessful":
		t.MessagePaymentSuccessful = new(MessagePaymentSuccessful)
		return json.Unmarshal(b, t.MessagePaymentSuccessful)
	case "messagePaymentSuccessfulBot":
		t.MessagePaymentSuccessfulBot = new(MessagePaymentSuccessfulBot)
		return json.Unmarshal(b, t.MessagePaymentSuccessfulBot)
	case "messagePaymentRefunded":
		t.MessagePaymentRefunded = new(MessagePaymentRefunded)
		return json.Unmarshal(b, t.MessagePaymentRefunded)
	case "messageGiftedPremium":
		t.MessageGiftedPremium = new(MessageGiftedPremium)
		return json.Unmarshal(b, t.MessageGiftedPremium)
	case "messagePremiumGiftCode":
		t.MessagePremiumGiftCode = new(MessagePremiumGiftCode)
		return json.Unmarshal(b, t.MessagePremiumGiftCode)
	case "messageGiveawayCreated":
		t.MessageGiveawayCreated = new(MessageGiveawayCreated)
		return json.Unmarshal(b, t.MessageGiveawayCreated)
	case "messageGiveaway":
		t.MessageGiveaway = new(MessageGiveaway)
		return json.Unmarshal(b, t.MessageGiveaway)
	case "messageGiveawayCompleted":
		t.MessageGiveawayCompleted = new(MessageGiveawayCompleted)
		return json.Unmarshal(b, t.MessageGiveawayCompleted)
	case "messageGiveawayWinners":
		t.MessageGiveawayWinners = new(MessageGiveawayWinners)
		return json.Unmarshal(b, t.MessageGiveawayWinners)
	case "messageGiftedStars":
		t.MessageGiftedStars = new(MessageGiftedStars)
		return json.Unmarshal(b, t.MessageGiftedStars)
	case "messageGiftedTon":
		t.MessageGiftedTon = new(MessageGiftedTon)
		return json.Unmarshal(b, t.MessageGiftedTon)
	case "messageGiveawayPrizeStars":
		t.MessageGiveawayPrizeStars = new(MessageGiveawayPrizeStars)
		return json.Unmarshal(b, t.MessageGiveawayPrizeStars)
	case "messageGift":
		t.MessageGift = new(MessageGift)
		return json.Unmarshal(b, t.MessageGift)
	case "messageUpgradedGift":
		t.MessageUpgradedGift = new(MessageUpgradedGift)
		return json.Unmarshal(b, t.MessageUpgradedGift)
	case "messageRefundedUpgradedGift":
		t.MessageRefundedUpgradedGift = new(MessageRefundedUpgradedGift)
		return json.Unmarshal(b, t.MessageRefundedUpgradedGift)
	case "messageUpgradedGiftPurchaseOffer":
		t.MessageUpgradedGiftPurchaseOffer = new(MessageUpgradedGiftPurchaseOffer)
		return json.Unmarshal(b, t.MessageUpgradedGiftPurchaseOffer)
	case "messageUpgradedGiftPurchaseOfferRejected":
		t.MessageUpgradedGiftPurchaseOfferRejected = new(MessageUpgradedGiftPurchaseOfferRejected)
		return json.Unmarshal(b, t.MessageUpgradedGiftPurchaseOfferRejected)
	case "messagePaidMessagesRefunded":
		t.MessagePaidMessagesRefunded = new(MessagePaidMessagesRefunded)
		return json.Unmarshal(b, t.MessagePaidMessagesRefunded)
	case "messagePaidMessagePriceChanged":
		t.MessagePaidMessagePriceChanged = new(MessagePaidMessagePriceChanged)
		return json.Unmarshal(b, t.MessagePaidMessagePriceChanged)
	case "messageDirectMessagePriceChanged":
		t.MessageDirectMessagePriceChanged = new(MessageDirectMessagePriceChanged)
		return json.Unmarshal(b, t.MessageDirectMessagePriceChanged)
	case "messageChecklistTasksDone":
		t.MessageChecklistTasksDone = new(MessageChecklistTasksDone)
		return json.Unmarshal(b, t.MessageChecklistTasksDone)
	case "messageChecklistTasksAdded":
		t.MessageChecklistTasksAdded = new(MessageChecklistTasksAdded)
		return json.Unmarshal(b, t.MessageChecklistTasksAdded)
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
	case "messageContactRegistered":
		t.MessageContactRegistered = new(MessageContactRegistered)
		return json.Unmarshal(b, t.MessageContactRegistered)
	case "messageUsersShared":
		t.MessageUsersShared = new(MessageUsersShared)
		return json.Unmarshal(b, t.MessageUsersShared)
	case "messageChatShared":
		t.MessageChatShared = new(MessageChatShared)
		return json.Unmarshal(b, t.MessageChatShared)
	case "messageBotWriteAccessAllowed":
		t.MessageBotWriteAccessAllowed = new(MessageBotWriteAccessAllowed)
		return json.Unmarshal(b, t.MessageBotWriteAccessAllowed)
	case "messageWebAppDataSent":
		t.MessageWebAppDataSent = new(MessageWebAppDataSent)
		return json.Unmarshal(b, t.MessageWebAppDataSent)
	case "messageWebAppDataReceived":
		t.MessageWebAppDataReceived = new(MessageWebAppDataReceived)
		return json.Unmarshal(b, t.MessageWebAppDataReceived)
	case "messagePassportDataSent":
		t.MessagePassportDataSent = new(MessagePassportDataSent)
		return json.Unmarshal(b, t.MessagePassportDataSent)
	case "messagePassportDataReceived":
		t.MessagePassportDataReceived = new(MessagePassportDataReceived)
		return json.Unmarshal(b, t.MessagePassportDataReceived)
	case "messageProximityAlertTriggered":
		t.MessageProximityAlertTriggered = new(MessageProximityAlertTriggered)
		return json.Unmarshal(b, t.MessageProximityAlertTriggered)
	case "messageUnsupported":
		t.MessageUnsupported = new(MessageUnsupported)
		return json.Unmarshal(b, t.MessageUnsupported)
	}
	return nil
}

func (t *MessageContent) MarshalJSON() ([]byte, error) {
	if t.MessageText != nil {
		return json.Marshal(t.MessageText)
	}
	if t.MessageAnimation != nil {
		return json.Marshal(t.MessageAnimation)
	}
	if t.MessageAudio != nil {
		return json.Marshal(t.MessageAudio)
	}
	if t.MessageDocument != nil {
		return json.Marshal(t.MessageDocument)
	}
	if t.MessagePaidMedia != nil {
		return json.Marshal(t.MessagePaidMedia)
	}
	if t.MessagePhoto != nil {
		return json.Marshal(t.MessagePhoto)
	}
	if t.MessageSticker != nil {
		return json.Marshal(t.MessageSticker)
	}
	if t.MessageVideo != nil {
		return json.Marshal(t.MessageVideo)
	}
	if t.MessageVideoNote != nil {
		return json.Marshal(t.MessageVideoNote)
	}
	if t.MessageVoiceNote != nil {
		return json.Marshal(t.MessageVoiceNote)
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
	if t.MessageLocation != nil {
		return json.Marshal(t.MessageLocation)
	}
	if t.MessageVenue != nil {
		return json.Marshal(t.MessageVenue)
	}
	if t.MessageContact != nil {
		return json.Marshal(t.MessageContact)
	}
	if t.MessageAnimatedEmoji != nil {
		return json.Marshal(t.MessageAnimatedEmoji)
	}
	if t.MessageDice != nil {
		return json.Marshal(t.MessageDice)
	}
	if t.MessageGame != nil {
		return json.Marshal(t.MessageGame)
	}
	if t.MessagePoll != nil {
		return json.Marshal(t.MessagePoll)
	}
	if t.MessageStakeDice != nil {
		return json.Marshal(t.MessageStakeDice)
	}
	if t.MessageStory != nil {
		return json.Marshal(t.MessageStory)
	}
	if t.MessageChecklist != nil {
		return json.Marshal(t.MessageChecklist)
	}
	if t.MessageInvoice != nil {
		return json.Marshal(t.MessageInvoice)
	}
	if t.MessageCall != nil {
		return json.Marshal(t.MessageCall)
	}
	if t.MessageGroupCall != nil {
		return json.Marshal(t.MessageGroupCall)
	}
	if t.MessageVideoChatScheduled != nil {
		return json.Marshal(t.MessageVideoChatScheduled)
	}
	if t.MessageVideoChatStarted != nil {
		return json.Marshal(t.MessageVideoChatStarted)
	}
	if t.MessageVideoChatEnded != nil {
		return json.Marshal(t.MessageVideoChatEnded)
	}
	if t.MessageInviteVideoChatParticipants != nil {
		return json.Marshal(t.MessageInviteVideoChatParticipants)
	}
	if t.MessageBasicGroupChatCreate != nil {
		return json.Marshal(t.MessageBasicGroupChatCreate)
	}
	if t.MessageSupergroupChatCreate != nil {
		return json.Marshal(t.MessageSupergroupChatCreate)
	}
	if t.MessageChatChangeTitle != nil {
		return json.Marshal(t.MessageChatChangeTitle)
	}
	if t.MessageChatChangePhoto != nil {
		return json.Marshal(t.MessageChatChangePhoto)
	}
	if t.MessageChatDeletePhoto != nil {
		return json.Marshal(t.MessageChatDeletePhoto)
	}
	if t.MessageChatAddMembers != nil {
		return json.Marshal(t.MessageChatAddMembers)
	}
	if t.MessageChatJoinByLink != nil {
		return json.Marshal(t.MessageChatJoinByLink)
	}
	if t.MessageChatJoinByRequest != nil {
		return json.Marshal(t.MessageChatJoinByRequest)
	}
	if t.MessageChatDeleteMember != nil {
		return json.Marshal(t.MessageChatDeleteMember)
	}
	if t.MessageChatUpgradeTo != nil {
		return json.Marshal(t.MessageChatUpgradeTo)
	}
	if t.MessageChatUpgradeFrom != nil {
		return json.Marshal(t.MessageChatUpgradeFrom)
	}
	if t.MessagePinMessage != nil {
		return json.Marshal(t.MessagePinMessage)
	}
	if t.MessageScreenshotTaken != nil {
		return json.Marshal(t.MessageScreenshotTaken)
	}
	if t.MessageChatSetBackground != nil {
		return json.Marshal(t.MessageChatSetBackground)
	}
	if t.MessageChatSetTheme != nil {
		return json.Marshal(t.MessageChatSetTheme)
	}
	if t.MessageChatSetMessageAutoDeleteTime != nil {
		return json.Marshal(t.MessageChatSetMessageAutoDeleteTime)
	}
	if t.MessageChatBoost != nil {
		return json.Marshal(t.MessageChatBoost)
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
	if t.MessageSuggestProfilePhoto != nil {
		return json.Marshal(t.MessageSuggestProfilePhoto)
	}
	if t.MessageSuggestBirthdate != nil {
		return json.Marshal(t.MessageSuggestBirthdate)
	}
	if t.MessageCustomServiceAction != nil {
		return json.Marshal(t.MessageCustomServiceAction)
	}
	if t.MessageGameScore != nil {
		return json.Marshal(t.MessageGameScore)
	}
	if t.MessagePaymentSuccessful != nil {
		return json.Marshal(t.MessagePaymentSuccessful)
	}
	if t.MessagePaymentSuccessfulBot != nil {
		return json.Marshal(t.MessagePaymentSuccessfulBot)
	}
	if t.MessagePaymentRefunded != nil {
		return json.Marshal(t.MessagePaymentRefunded)
	}
	if t.MessageGiftedPremium != nil {
		return json.Marshal(t.MessageGiftedPremium)
	}
	if t.MessagePremiumGiftCode != nil {
		return json.Marshal(t.MessagePremiumGiftCode)
	}
	if t.MessageGiveawayCreated != nil {
		return json.Marshal(t.MessageGiveawayCreated)
	}
	if t.MessageGiveaway != nil {
		return json.Marshal(t.MessageGiveaway)
	}
	if t.MessageGiveawayCompleted != nil {
		return json.Marshal(t.MessageGiveawayCompleted)
	}
	if t.MessageGiveawayWinners != nil {
		return json.Marshal(t.MessageGiveawayWinners)
	}
	if t.MessageGiftedStars != nil {
		return json.Marshal(t.MessageGiftedStars)
	}
	if t.MessageGiftedTon != nil {
		return json.Marshal(t.MessageGiftedTon)
	}
	if t.MessageGiveawayPrizeStars != nil {
		return json.Marshal(t.MessageGiveawayPrizeStars)
	}
	if t.MessageGift != nil {
		return json.Marshal(t.MessageGift)
	}
	if t.MessageUpgradedGift != nil {
		return json.Marshal(t.MessageUpgradedGift)
	}
	if t.MessageRefundedUpgradedGift != nil {
		return json.Marshal(t.MessageRefundedUpgradedGift)
	}
	if t.MessageUpgradedGiftPurchaseOffer != nil {
		return json.Marshal(t.MessageUpgradedGiftPurchaseOffer)
	}
	if t.MessageUpgradedGiftPurchaseOfferRejected != nil {
		return json.Marshal(t.MessageUpgradedGiftPurchaseOfferRejected)
	}
	if t.MessagePaidMessagesRefunded != nil {
		return json.Marshal(t.MessagePaidMessagesRefunded)
	}
	if t.MessagePaidMessagePriceChanged != nil {
		return json.Marshal(t.MessagePaidMessagePriceChanged)
	}
	if t.MessageDirectMessagePriceChanged != nil {
		return json.Marshal(t.MessageDirectMessagePriceChanged)
	}
	if t.MessageChecklistTasksDone != nil {
		return json.Marshal(t.MessageChecklistTasksDone)
	}
	if t.MessageChecklistTasksAdded != nil {
		return json.Marshal(t.MessageChecklistTasksAdded)
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
	if t.MessageContactRegistered != nil {
		return json.Marshal(t.MessageContactRegistered)
	}
	if t.MessageUsersShared != nil {
		return json.Marshal(t.MessageUsersShared)
	}
	if t.MessageChatShared != nil {
		return json.Marshal(t.MessageChatShared)
	}
	if t.MessageBotWriteAccessAllowed != nil {
		return json.Marshal(t.MessageBotWriteAccessAllowed)
	}
	if t.MessageWebAppDataSent != nil {
		return json.Marshal(t.MessageWebAppDataSent)
	}
	if t.MessageWebAppDataReceived != nil {
		return json.Marshal(t.MessageWebAppDataReceived)
	}
	if t.MessagePassportDataSent != nil {
		return json.Marshal(t.MessagePassportDataSent)
	}
	if t.MessagePassportDataReceived != nil {
		return json.Marshal(t.MessagePassportDataReceived)
	}
	if t.MessageProximityAlertTriggered != nil {
		return json.Marshal(t.MessageProximityAlertTriggered)
	}
	if t.MessageUnsupported != nil {
		return json.Marshal(t.MessageUnsupported)
	}
	return nil, fmt.Errorf("no value set for MessageContent")
}

// StoryContent Contains the content of a story
type StoryContent struct {
	StoryContentPhoto       *StoryContentPhoto       `json:"storyContentPhoto,omitempty"`
	StoryContentVideo       *StoryContentVideo       `json:"storyContentVideo,omitempty"`
	StoryContentLive        *StoryContentLive        `json:"storyContentLive,omitempty"`
	StoryContentUnsupported *StoryContentUnsupported `json:"storyContentUnsupported,omitempty"`
}

func (t *StoryContent) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyContentPhoto":
		t.StoryContentPhoto = new(StoryContentPhoto)
		return json.Unmarshal(b, t.StoryContentPhoto)
	case "storyContentVideo":
		t.StoryContentVideo = new(StoryContentVideo)
		return json.Unmarshal(b, t.StoryContentVideo)
	case "storyContentLive":
		t.StoryContentLive = new(StoryContentLive)
		return json.Unmarshal(b, t.StoryContentLive)
	case "storyContentUnsupported":
		t.StoryContentUnsupported = new(StoryContentUnsupported)
		return json.Unmarshal(b, t.StoryContentUnsupported)
	}
	return nil
}

func (t *StoryContent) MarshalJSON() ([]byte, error) {
	if t.StoryContentPhoto != nil {
		return json.Marshal(t.StoryContentPhoto)
	}
	if t.StoryContentVideo != nil {
		return json.Marshal(t.StoryContentVideo)
	}
	if t.StoryContentLive != nil {
		return json.Marshal(t.StoryContentLive)
	}
	if t.StoryContentUnsupported != nil {
		return json.Marshal(t.StoryContentUnsupported)
	}
	return nil, fmt.Errorf("no value set for StoryContent")
}

// ChatEventAction Represents a chat event
type ChatEventAction struct {
	ChatEventMessageEdited                          *ChatEventMessageEdited                          `json:"chatEventMessageEdited,omitempty"`
	ChatEventMessageDeleted                         *ChatEventMessageDeleted                         `json:"chatEventMessageDeleted,omitempty"`
	ChatEventMessagePinned                          *ChatEventMessagePinned                          `json:"chatEventMessagePinned,omitempty"`
	ChatEventMessageUnpinned                        *ChatEventMessageUnpinned                        `json:"chatEventMessageUnpinned,omitempty"`
	ChatEventPollStopped                            *ChatEventPollStopped                            `json:"chatEventPollStopped,omitempty"`
	ChatEventMemberJoined                           *ChatEventMemberJoined                           `json:"chatEventMemberJoined,omitempty"`
	ChatEventMemberJoinedByInviteLink               *ChatEventMemberJoinedByInviteLink               `json:"chatEventMemberJoinedByInviteLink,omitempty"`
	ChatEventMemberJoinedByRequest                  *ChatEventMemberJoinedByRequest                  `json:"chatEventMemberJoinedByRequest,omitempty"`
	ChatEventMemberInvited                          *ChatEventMemberInvited                          `json:"chatEventMemberInvited,omitempty"`
	ChatEventMemberLeft                             *ChatEventMemberLeft                             `json:"chatEventMemberLeft,omitempty"`
	ChatEventMemberPromoted                         *ChatEventMemberPromoted                         `json:"chatEventMemberPromoted,omitempty"`
	ChatEventMemberRestricted                       *ChatEventMemberRestricted                       `json:"chatEventMemberRestricted,omitempty"`
	ChatEventMemberSubscriptionExtended             *ChatEventMemberSubscriptionExtended             `json:"chatEventMemberSubscriptionExtended,omitempty"`
	ChatEventAvailableReactionsChanged              *ChatEventAvailableReactionsChanged              `json:"chatEventAvailableReactionsChanged,omitempty"`
	ChatEventBackgroundChanged                      *ChatEventBackgroundChanged                      `json:"chatEventBackgroundChanged,omitempty"`
	ChatEventDescriptionChanged                     *ChatEventDescriptionChanged                     `json:"chatEventDescriptionChanged,omitempty"`
	ChatEventEmojiStatusChanged                     *ChatEventEmojiStatusChanged                     `json:"chatEventEmojiStatusChanged,omitempty"`
	ChatEventLinkedChatChanged                      *ChatEventLinkedChatChanged                      `json:"chatEventLinkedChatChanged,omitempty"`
	ChatEventLocationChanged                        *ChatEventLocationChanged                        `json:"chatEventLocationChanged,omitempty"`
	ChatEventMessageAutoDeleteTimeChanged           *ChatEventMessageAutoDeleteTimeChanged           `json:"chatEventMessageAutoDeleteTimeChanged,omitempty"`
	ChatEventPermissionsChanged                     *ChatEventPermissionsChanged                     `json:"chatEventPermissionsChanged,omitempty"`
	ChatEventPhotoChanged                           *ChatEventPhotoChanged                           `json:"chatEventPhotoChanged,omitempty"`
	ChatEventSlowModeDelayChanged                   *ChatEventSlowModeDelayChanged                   `json:"chatEventSlowModeDelayChanged,omitempty"`
	ChatEventStickerSetChanged                      *ChatEventStickerSetChanged                      `json:"chatEventStickerSetChanged,omitempty"`
	ChatEventCustomEmojiStickerSetChanged           *ChatEventCustomEmojiStickerSetChanged           `json:"chatEventCustomEmojiStickerSetChanged,omitempty"`
	ChatEventTitleChanged                           *ChatEventTitleChanged                           `json:"chatEventTitleChanged,omitempty"`
	ChatEventUsernameChanged                        *ChatEventUsernameChanged                        `json:"chatEventUsernameChanged,omitempty"`
	ChatEventActiveUsernamesChanged                 *ChatEventActiveUsernamesChanged                 `json:"chatEventActiveUsernamesChanged,omitempty"`
	ChatEventAccentColorChanged                     *ChatEventAccentColorChanged                     `json:"chatEventAccentColorChanged,omitempty"`
	ChatEventProfileAccentColorChanged              *ChatEventProfileAccentColorChanged              `json:"chatEventProfileAccentColorChanged,omitempty"`
	ChatEventHasProtectedContentToggled             *ChatEventHasProtectedContentToggled             `json:"chatEventHasProtectedContentToggled,omitempty"`
	ChatEventInvitesToggled                         *ChatEventInvitesToggled                         `json:"chatEventInvitesToggled,omitempty"`
	ChatEventIsAllHistoryAvailableToggled           *ChatEventIsAllHistoryAvailableToggled           `json:"chatEventIsAllHistoryAvailableToggled,omitempty"`
	ChatEventHasAggressiveAntiSpamEnabledToggled    *ChatEventHasAggressiveAntiSpamEnabledToggled    `json:"chatEventHasAggressiveAntiSpamEnabledToggled,omitempty"`
	ChatEventSignMessagesToggled                    *ChatEventSignMessagesToggled                    `json:"chatEventSignMessagesToggled,omitempty"`
	ChatEventShowMessageSenderToggled               *ChatEventShowMessageSenderToggled               `json:"chatEventShowMessageSenderToggled,omitempty"`
	ChatEventAutomaticTranslationToggled            *ChatEventAutomaticTranslationToggled            `json:"chatEventAutomaticTranslationToggled,omitempty"`
	ChatEventInviteLinkEdited                       *ChatEventInviteLinkEdited                       `json:"chatEventInviteLinkEdited,omitempty"`
	ChatEventInviteLinkRevoked                      *ChatEventInviteLinkRevoked                      `json:"chatEventInviteLinkRevoked,omitempty"`
	ChatEventInviteLinkDeleted                      *ChatEventInviteLinkDeleted                      `json:"chatEventInviteLinkDeleted,omitempty"`
	ChatEventVideoChatCreated                       *ChatEventVideoChatCreated                       `json:"chatEventVideoChatCreated,omitempty"`
	ChatEventVideoChatEnded                         *ChatEventVideoChatEnded                         `json:"chatEventVideoChatEnded,omitempty"`
	ChatEventVideoChatMuteNewParticipantsToggled    *ChatEventVideoChatMuteNewParticipantsToggled    `json:"chatEventVideoChatMuteNewParticipantsToggled,omitempty"`
	ChatEventVideoChatParticipantIsMutedToggled     *ChatEventVideoChatParticipantIsMutedToggled     `json:"chatEventVideoChatParticipantIsMutedToggled,omitempty"`
	ChatEventVideoChatParticipantVolumeLevelChanged *ChatEventVideoChatParticipantVolumeLevelChanged `json:"chatEventVideoChatParticipantVolumeLevelChanged,omitempty"`
	ChatEventIsForumToggled                         *ChatEventIsForumToggled                         `json:"chatEventIsForumToggled,omitempty"`
	ChatEventForumTopicCreated                      *ChatEventForumTopicCreated                      `json:"chatEventForumTopicCreated,omitempty"`
	ChatEventForumTopicEdited                       *ChatEventForumTopicEdited                       `json:"chatEventForumTopicEdited,omitempty"`
	ChatEventForumTopicToggleIsClosed               *ChatEventForumTopicToggleIsClosed               `json:"chatEventForumTopicToggleIsClosed,omitempty"`
	ChatEventForumTopicToggleIsHidden               *ChatEventForumTopicToggleIsHidden               `json:"chatEventForumTopicToggleIsHidden,omitempty"`
	ChatEventForumTopicDeleted                      *ChatEventForumTopicDeleted                      `json:"chatEventForumTopicDeleted,omitempty"`
	ChatEventForumTopicPinned                       *ChatEventForumTopicPinned                       `json:"chatEventForumTopicPinned,omitempty"`
}

func (t *ChatEventAction) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatEventMessageEdited":
		t.ChatEventMessageEdited = new(ChatEventMessageEdited)
		return json.Unmarshal(b, t.ChatEventMessageEdited)
	case "chatEventMessageDeleted":
		t.ChatEventMessageDeleted = new(ChatEventMessageDeleted)
		return json.Unmarshal(b, t.ChatEventMessageDeleted)
	case "chatEventMessagePinned":
		t.ChatEventMessagePinned = new(ChatEventMessagePinned)
		return json.Unmarshal(b, t.ChatEventMessagePinned)
	case "chatEventMessageUnpinned":
		t.ChatEventMessageUnpinned = new(ChatEventMessageUnpinned)
		return json.Unmarshal(b, t.ChatEventMessageUnpinned)
	case "chatEventPollStopped":
		t.ChatEventPollStopped = new(ChatEventPollStopped)
		return json.Unmarshal(b, t.ChatEventPollStopped)
	case "chatEventMemberJoined":
		t.ChatEventMemberJoined = new(ChatEventMemberJoined)
		return json.Unmarshal(b, t.ChatEventMemberJoined)
	case "chatEventMemberJoinedByInviteLink":
		t.ChatEventMemberJoinedByInviteLink = new(ChatEventMemberJoinedByInviteLink)
		return json.Unmarshal(b, t.ChatEventMemberJoinedByInviteLink)
	case "chatEventMemberJoinedByRequest":
		t.ChatEventMemberJoinedByRequest = new(ChatEventMemberJoinedByRequest)
		return json.Unmarshal(b, t.ChatEventMemberJoinedByRequest)
	case "chatEventMemberInvited":
		t.ChatEventMemberInvited = new(ChatEventMemberInvited)
		return json.Unmarshal(b, t.ChatEventMemberInvited)
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
	case "chatEventAvailableReactionsChanged":
		t.ChatEventAvailableReactionsChanged = new(ChatEventAvailableReactionsChanged)
		return json.Unmarshal(b, t.ChatEventAvailableReactionsChanged)
	case "chatEventBackgroundChanged":
		t.ChatEventBackgroundChanged = new(ChatEventBackgroundChanged)
		return json.Unmarshal(b, t.ChatEventBackgroundChanged)
	case "chatEventDescriptionChanged":
		t.ChatEventDescriptionChanged = new(ChatEventDescriptionChanged)
		return json.Unmarshal(b, t.ChatEventDescriptionChanged)
	case "chatEventEmojiStatusChanged":
		t.ChatEventEmojiStatusChanged = new(ChatEventEmojiStatusChanged)
		return json.Unmarshal(b, t.ChatEventEmojiStatusChanged)
	case "chatEventLinkedChatChanged":
		t.ChatEventLinkedChatChanged = new(ChatEventLinkedChatChanged)
		return json.Unmarshal(b, t.ChatEventLinkedChatChanged)
	case "chatEventLocationChanged":
		t.ChatEventLocationChanged = new(ChatEventLocationChanged)
		return json.Unmarshal(b, t.ChatEventLocationChanged)
	case "chatEventMessageAutoDeleteTimeChanged":
		t.ChatEventMessageAutoDeleteTimeChanged = new(ChatEventMessageAutoDeleteTimeChanged)
		return json.Unmarshal(b, t.ChatEventMessageAutoDeleteTimeChanged)
	case "chatEventPermissionsChanged":
		t.ChatEventPermissionsChanged = new(ChatEventPermissionsChanged)
		return json.Unmarshal(b, t.ChatEventPermissionsChanged)
	case "chatEventPhotoChanged":
		t.ChatEventPhotoChanged = new(ChatEventPhotoChanged)
		return json.Unmarshal(b, t.ChatEventPhotoChanged)
	case "chatEventSlowModeDelayChanged":
		t.ChatEventSlowModeDelayChanged = new(ChatEventSlowModeDelayChanged)
		return json.Unmarshal(b, t.ChatEventSlowModeDelayChanged)
	case "chatEventStickerSetChanged":
		t.ChatEventStickerSetChanged = new(ChatEventStickerSetChanged)
		return json.Unmarshal(b, t.ChatEventStickerSetChanged)
	case "chatEventCustomEmojiStickerSetChanged":
		t.ChatEventCustomEmojiStickerSetChanged = new(ChatEventCustomEmojiStickerSetChanged)
		return json.Unmarshal(b, t.ChatEventCustomEmojiStickerSetChanged)
	case "chatEventTitleChanged":
		t.ChatEventTitleChanged = new(ChatEventTitleChanged)
		return json.Unmarshal(b, t.ChatEventTitleChanged)
	case "chatEventUsernameChanged":
		t.ChatEventUsernameChanged = new(ChatEventUsernameChanged)
		return json.Unmarshal(b, t.ChatEventUsernameChanged)
	case "chatEventActiveUsernamesChanged":
		t.ChatEventActiveUsernamesChanged = new(ChatEventActiveUsernamesChanged)
		return json.Unmarshal(b, t.ChatEventActiveUsernamesChanged)
	case "chatEventAccentColorChanged":
		t.ChatEventAccentColorChanged = new(ChatEventAccentColorChanged)
		return json.Unmarshal(b, t.ChatEventAccentColorChanged)
	case "chatEventProfileAccentColorChanged":
		t.ChatEventProfileAccentColorChanged = new(ChatEventProfileAccentColorChanged)
		return json.Unmarshal(b, t.ChatEventProfileAccentColorChanged)
	case "chatEventHasProtectedContentToggled":
		t.ChatEventHasProtectedContentToggled = new(ChatEventHasProtectedContentToggled)
		return json.Unmarshal(b, t.ChatEventHasProtectedContentToggled)
	case "chatEventInvitesToggled":
		t.ChatEventInvitesToggled = new(ChatEventInvitesToggled)
		return json.Unmarshal(b, t.ChatEventInvitesToggled)
	case "chatEventIsAllHistoryAvailableToggled":
		t.ChatEventIsAllHistoryAvailableToggled = new(ChatEventIsAllHistoryAvailableToggled)
		return json.Unmarshal(b, t.ChatEventIsAllHistoryAvailableToggled)
	case "chatEventHasAggressiveAntiSpamEnabledToggled":
		t.ChatEventHasAggressiveAntiSpamEnabledToggled = new(ChatEventHasAggressiveAntiSpamEnabledToggled)
		return json.Unmarshal(b, t.ChatEventHasAggressiveAntiSpamEnabledToggled)
	case "chatEventSignMessagesToggled":
		t.ChatEventSignMessagesToggled = new(ChatEventSignMessagesToggled)
		return json.Unmarshal(b, t.ChatEventSignMessagesToggled)
	case "chatEventShowMessageSenderToggled":
		t.ChatEventShowMessageSenderToggled = new(ChatEventShowMessageSenderToggled)
		return json.Unmarshal(b, t.ChatEventShowMessageSenderToggled)
	case "chatEventAutomaticTranslationToggled":
		t.ChatEventAutomaticTranslationToggled = new(ChatEventAutomaticTranslationToggled)
		return json.Unmarshal(b, t.ChatEventAutomaticTranslationToggled)
	case "chatEventInviteLinkEdited":
		t.ChatEventInviteLinkEdited = new(ChatEventInviteLinkEdited)
		return json.Unmarshal(b, t.ChatEventInviteLinkEdited)
	case "chatEventInviteLinkRevoked":
		t.ChatEventInviteLinkRevoked = new(ChatEventInviteLinkRevoked)
		return json.Unmarshal(b, t.ChatEventInviteLinkRevoked)
	case "chatEventInviteLinkDeleted":
		t.ChatEventInviteLinkDeleted = new(ChatEventInviteLinkDeleted)
		return json.Unmarshal(b, t.ChatEventInviteLinkDeleted)
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
	case "chatEventIsForumToggled":
		t.ChatEventIsForumToggled = new(ChatEventIsForumToggled)
		return json.Unmarshal(b, t.ChatEventIsForumToggled)
	case "chatEventForumTopicCreated":
		t.ChatEventForumTopicCreated = new(ChatEventForumTopicCreated)
		return json.Unmarshal(b, t.ChatEventForumTopicCreated)
	case "chatEventForumTopicEdited":
		t.ChatEventForumTopicEdited = new(ChatEventForumTopicEdited)
		return json.Unmarshal(b, t.ChatEventForumTopicEdited)
	case "chatEventForumTopicToggleIsClosed":
		t.ChatEventForumTopicToggleIsClosed = new(ChatEventForumTopicToggleIsClosed)
		return json.Unmarshal(b, t.ChatEventForumTopicToggleIsClosed)
	case "chatEventForumTopicToggleIsHidden":
		t.ChatEventForumTopicToggleIsHidden = new(ChatEventForumTopicToggleIsHidden)
		return json.Unmarshal(b, t.ChatEventForumTopicToggleIsHidden)
	case "chatEventForumTopicDeleted":
		t.ChatEventForumTopicDeleted = new(ChatEventForumTopicDeleted)
		return json.Unmarshal(b, t.ChatEventForumTopicDeleted)
	case "chatEventForumTopicPinned":
		t.ChatEventForumTopicPinned = new(ChatEventForumTopicPinned)
		return json.Unmarshal(b, t.ChatEventForumTopicPinned)
	}
	return nil
}

func (t *ChatEventAction) MarshalJSON() ([]byte, error) {
	if t.ChatEventMessageEdited != nil {
		return json.Marshal(t.ChatEventMessageEdited)
	}
	if t.ChatEventMessageDeleted != nil {
		return json.Marshal(t.ChatEventMessageDeleted)
	}
	if t.ChatEventMessagePinned != nil {
		return json.Marshal(t.ChatEventMessagePinned)
	}
	if t.ChatEventMessageUnpinned != nil {
		return json.Marshal(t.ChatEventMessageUnpinned)
	}
	if t.ChatEventPollStopped != nil {
		return json.Marshal(t.ChatEventPollStopped)
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
	if t.ChatEventMemberInvited != nil {
		return json.Marshal(t.ChatEventMemberInvited)
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
	if t.ChatEventAvailableReactionsChanged != nil {
		return json.Marshal(t.ChatEventAvailableReactionsChanged)
	}
	if t.ChatEventBackgroundChanged != nil {
		return json.Marshal(t.ChatEventBackgroundChanged)
	}
	if t.ChatEventDescriptionChanged != nil {
		return json.Marshal(t.ChatEventDescriptionChanged)
	}
	if t.ChatEventEmojiStatusChanged != nil {
		return json.Marshal(t.ChatEventEmojiStatusChanged)
	}
	if t.ChatEventLinkedChatChanged != nil {
		return json.Marshal(t.ChatEventLinkedChatChanged)
	}
	if t.ChatEventLocationChanged != nil {
		return json.Marshal(t.ChatEventLocationChanged)
	}
	if t.ChatEventMessageAutoDeleteTimeChanged != nil {
		return json.Marshal(t.ChatEventMessageAutoDeleteTimeChanged)
	}
	if t.ChatEventPermissionsChanged != nil {
		return json.Marshal(t.ChatEventPermissionsChanged)
	}
	if t.ChatEventPhotoChanged != nil {
		return json.Marshal(t.ChatEventPhotoChanged)
	}
	if t.ChatEventSlowModeDelayChanged != nil {
		return json.Marshal(t.ChatEventSlowModeDelayChanged)
	}
	if t.ChatEventStickerSetChanged != nil {
		return json.Marshal(t.ChatEventStickerSetChanged)
	}
	if t.ChatEventCustomEmojiStickerSetChanged != nil {
		return json.Marshal(t.ChatEventCustomEmojiStickerSetChanged)
	}
	if t.ChatEventTitleChanged != nil {
		return json.Marshal(t.ChatEventTitleChanged)
	}
	if t.ChatEventUsernameChanged != nil {
		return json.Marshal(t.ChatEventUsernameChanged)
	}
	if t.ChatEventActiveUsernamesChanged != nil {
		return json.Marshal(t.ChatEventActiveUsernamesChanged)
	}
	if t.ChatEventAccentColorChanged != nil {
		return json.Marshal(t.ChatEventAccentColorChanged)
	}
	if t.ChatEventProfileAccentColorChanged != nil {
		return json.Marshal(t.ChatEventProfileAccentColorChanged)
	}
	if t.ChatEventHasProtectedContentToggled != nil {
		return json.Marshal(t.ChatEventHasProtectedContentToggled)
	}
	if t.ChatEventInvitesToggled != nil {
		return json.Marshal(t.ChatEventInvitesToggled)
	}
	if t.ChatEventIsAllHistoryAvailableToggled != nil {
		return json.Marshal(t.ChatEventIsAllHistoryAvailableToggled)
	}
	if t.ChatEventHasAggressiveAntiSpamEnabledToggled != nil {
		return json.Marshal(t.ChatEventHasAggressiveAntiSpamEnabledToggled)
	}
	if t.ChatEventSignMessagesToggled != nil {
		return json.Marshal(t.ChatEventSignMessagesToggled)
	}
	if t.ChatEventShowMessageSenderToggled != nil {
		return json.Marshal(t.ChatEventShowMessageSenderToggled)
	}
	if t.ChatEventAutomaticTranslationToggled != nil {
		return json.Marshal(t.ChatEventAutomaticTranslationToggled)
	}
	if t.ChatEventInviteLinkEdited != nil {
		return json.Marshal(t.ChatEventInviteLinkEdited)
	}
	if t.ChatEventInviteLinkRevoked != nil {
		return json.Marshal(t.ChatEventInviteLinkRevoked)
	}
	if t.ChatEventInviteLinkDeleted != nil {
		return json.Marshal(t.ChatEventInviteLinkDeleted)
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
	if t.ChatEventIsForumToggled != nil {
		return json.Marshal(t.ChatEventIsForumToggled)
	}
	if t.ChatEventForumTopicCreated != nil {
		return json.Marshal(t.ChatEventForumTopicCreated)
	}
	if t.ChatEventForumTopicEdited != nil {
		return json.Marshal(t.ChatEventForumTopicEdited)
	}
	if t.ChatEventForumTopicToggleIsClosed != nil {
		return json.Marshal(t.ChatEventForumTopicToggleIsClosed)
	}
	if t.ChatEventForumTopicToggleIsHidden != nil {
		return json.Marshal(t.ChatEventForumTopicToggleIsHidden)
	}
	if t.ChatEventForumTopicDeleted != nil {
		return json.Marshal(t.ChatEventForumTopicDeleted)
	}
	if t.ChatEventForumTopicPinned != nil {
		return json.Marshal(t.ChatEventForumTopicPinned)
	}
	return nil, fmt.Errorf("no value set for ChatEventAction")
}

// PremiumFeature Describes a feature available to Premium users
type PremiumFeature struct {
	PremiumFeatureIncreasedLimits         *PremiumFeatureIncreasedLimits         `json:"premiumFeatureIncreasedLimits,omitempty"`
	PremiumFeatureIncreasedUploadFileSize *PremiumFeatureIncreasedUploadFileSize `json:"premiumFeatureIncreasedUploadFileSize,omitempty"`
	PremiumFeatureImprovedDownloadSpeed   *PremiumFeatureImprovedDownloadSpeed   `json:"premiumFeatureImprovedDownloadSpeed,omitempty"`
	PremiumFeatureVoiceRecognition        *PremiumFeatureVoiceRecognition        `json:"premiumFeatureVoiceRecognition,omitempty"`
	PremiumFeatureDisabledAds             *PremiumFeatureDisabledAds             `json:"premiumFeatureDisabledAds,omitempty"`
	PremiumFeatureUniqueReactions         *PremiumFeatureUniqueReactions         `json:"premiumFeatureUniqueReactions,omitempty"`
	PremiumFeatureUniqueStickers          *PremiumFeatureUniqueStickers          `json:"premiumFeatureUniqueStickers,omitempty"`
	PremiumFeatureCustomEmoji             *PremiumFeatureCustomEmoji             `json:"premiumFeatureCustomEmoji,omitempty"`
	PremiumFeatureAdvancedChatManagement  *PremiumFeatureAdvancedChatManagement  `json:"premiumFeatureAdvancedChatManagement,omitempty"`
	PremiumFeatureProfileBadge            *PremiumFeatureProfileBadge            `json:"premiumFeatureProfileBadge,omitempty"`
	PremiumFeatureEmojiStatus             *PremiumFeatureEmojiStatus             `json:"premiumFeatureEmojiStatus,omitempty"`
	PremiumFeatureAnimatedProfilePhoto    *PremiumFeatureAnimatedProfilePhoto    `json:"premiumFeatureAnimatedProfilePhoto,omitempty"`
	PremiumFeatureForumTopicIcon          *PremiumFeatureForumTopicIcon          `json:"premiumFeatureForumTopicIcon,omitempty"`
	PremiumFeatureAppIcons                *PremiumFeatureAppIcons                `json:"premiumFeatureAppIcons,omitempty"`
	PremiumFeatureRealTimeChatTranslation *PremiumFeatureRealTimeChatTranslation `json:"premiumFeatureRealTimeChatTranslation,omitempty"`
	PremiumFeatureUpgradedStories         *PremiumFeatureUpgradedStories         `json:"premiumFeatureUpgradedStories,omitempty"`
	PremiumFeatureChatBoost               *PremiumFeatureChatBoost               `json:"premiumFeatureChatBoost,omitempty"`
	PremiumFeatureAccentColor             *PremiumFeatureAccentColor             `json:"premiumFeatureAccentColor,omitempty"`
	PremiumFeatureBackgroundForBoth       *PremiumFeatureBackgroundForBoth       `json:"premiumFeatureBackgroundForBoth,omitempty"`
	PremiumFeatureSavedMessagesTags       *PremiumFeatureSavedMessagesTags       `json:"premiumFeatureSavedMessagesTags,omitempty"`
	PremiumFeatureMessagePrivacy          *PremiumFeatureMessagePrivacy          `json:"premiumFeatureMessagePrivacy,omitempty"`
	PremiumFeatureLastSeenTimes           *PremiumFeatureLastSeenTimes           `json:"premiumFeatureLastSeenTimes,omitempty"`
	PremiumFeatureBusiness                *PremiumFeatureBusiness                `json:"premiumFeatureBusiness,omitempty"`
	PremiumFeatureMessageEffects          *PremiumFeatureMessageEffects          `json:"premiumFeatureMessageEffects,omitempty"`
	PremiumFeatureChecklists              *PremiumFeatureChecklists              `json:"premiumFeatureChecklists,omitempty"`
	PremiumFeaturePaidMessages            *PremiumFeaturePaidMessages            `json:"premiumFeaturePaidMessages,omitempty"`
}

func (t *PremiumFeature) Type() string {
	return ""
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
	switch typeObj.Type {
	case "premiumFeatureIncreasedLimits":
		t.PremiumFeatureIncreasedLimits = new(PremiumFeatureIncreasedLimits)
		return json.Unmarshal(b, t.PremiumFeatureIncreasedLimits)
	case "premiumFeatureIncreasedUploadFileSize":
		t.PremiumFeatureIncreasedUploadFileSize = new(PremiumFeatureIncreasedUploadFileSize)
		return json.Unmarshal(b, t.PremiumFeatureIncreasedUploadFileSize)
	case "premiumFeatureImprovedDownloadSpeed":
		t.PremiumFeatureImprovedDownloadSpeed = new(PremiumFeatureImprovedDownloadSpeed)
		return json.Unmarshal(b, t.PremiumFeatureImprovedDownloadSpeed)
	case "premiumFeatureVoiceRecognition":
		t.PremiumFeatureVoiceRecognition = new(PremiumFeatureVoiceRecognition)
		return json.Unmarshal(b, t.PremiumFeatureVoiceRecognition)
	case "premiumFeatureDisabledAds":
		t.PremiumFeatureDisabledAds = new(PremiumFeatureDisabledAds)
		return json.Unmarshal(b, t.PremiumFeatureDisabledAds)
	case "premiumFeatureUniqueReactions":
		t.PremiumFeatureUniqueReactions = new(PremiumFeatureUniqueReactions)
		return json.Unmarshal(b, t.PremiumFeatureUniqueReactions)
	case "premiumFeatureUniqueStickers":
		t.PremiumFeatureUniqueStickers = new(PremiumFeatureUniqueStickers)
		return json.Unmarshal(b, t.PremiumFeatureUniqueStickers)
	case "premiumFeatureCustomEmoji":
		t.PremiumFeatureCustomEmoji = new(PremiumFeatureCustomEmoji)
		return json.Unmarshal(b, t.PremiumFeatureCustomEmoji)
	case "premiumFeatureAdvancedChatManagement":
		t.PremiumFeatureAdvancedChatManagement = new(PremiumFeatureAdvancedChatManagement)
		return json.Unmarshal(b, t.PremiumFeatureAdvancedChatManagement)
	case "premiumFeatureProfileBadge":
		t.PremiumFeatureProfileBadge = new(PremiumFeatureProfileBadge)
		return json.Unmarshal(b, t.PremiumFeatureProfileBadge)
	case "premiumFeatureEmojiStatus":
		t.PremiumFeatureEmojiStatus = new(PremiumFeatureEmojiStatus)
		return json.Unmarshal(b, t.PremiumFeatureEmojiStatus)
	case "premiumFeatureAnimatedProfilePhoto":
		t.PremiumFeatureAnimatedProfilePhoto = new(PremiumFeatureAnimatedProfilePhoto)
		return json.Unmarshal(b, t.PremiumFeatureAnimatedProfilePhoto)
	case "premiumFeatureForumTopicIcon":
		t.PremiumFeatureForumTopicIcon = new(PremiumFeatureForumTopicIcon)
		return json.Unmarshal(b, t.PremiumFeatureForumTopicIcon)
	case "premiumFeatureAppIcons":
		t.PremiumFeatureAppIcons = new(PremiumFeatureAppIcons)
		return json.Unmarshal(b, t.PremiumFeatureAppIcons)
	case "premiumFeatureRealTimeChatTranslation":
		t.PremiumFeatureRealTimeChatTranslation = new(PremiumFeatureRealTimeChatTranslation)
		return json.Unmarshal(b, t.PremiumFeatureRealTimeChatTranslation)
	case "premiumFeatureUpgradedStories":
		t.PremiumFeatureUpgradedStories = new(PremiumFeatureUpgradedStories)
		return json.Unmarshal(b, t.PremiumFeatureUpgradedStories)
	case "premiumFeatureChatBoost":
		t.PremiumFeatureChatBoost = new(PremiumFeatureChatBoost)
		return json.Unmarshal(b, t.PremiumFeatureChatBoost)
	case "premiumFeatureAccentColor":
		t.PremiumFeatureAccentColor = new(PremiumFeatureAccentColor)
		return json.Unmarshal(b, t.PremiumFeatureAccentColor)
	case "premiumFeatureBackgroundForBoth":
		t.PremiumFeatureBackgroundForBoth = new(PremiumFeatureBackgroundForBoth)
		return json.Unmarshal(b, t.PremiumFeatureBackgroundForBoth)
	case "premiumFeatureSavedMessagesTags":
		t.PremiumFeatureSavedMessagesTags = new(PremiumFeatureSavedMessagesTags)
		return json.Unmarshal(b, t.PremiumFeatureSavedMessagesTags)
	case "premiumFeatureMessagePrivacy":
		t.PremiumFeatureMessagePrivacy = new(PremiumFeatureMessagePrivacy)
		return json.Unmarshal(b, t.PremiumFeatureMessagePrivacy)
	case "premiumFeatureLastSeenTimes":
		t.PremiumFeatureLastSeenTimes = new(PremiumFeatureLastSeenTimes)
		return json.Unmarshal(b, t.PremiumFeatureLastSeenTimes)
	case "premiumFeatureBusiness":
		t.PremiumFeatureBusiness = new(PremiumFeatureBusiness)
		return json.Unmarshal(b, t.PremiumFeatureBusiness)
	case "premiumFeatureMessageEffects":
		t.PremiumFeatureMessageEffects = new(PremiumFeatureMessageEffects)
		return json.Unmarshal(b, t.PremiumFeatureMessageEffects)
	case "premiumFeatureChecklists":
		t.PremiumFeatureChecklists = new(PremiumFeatureChecklists)
		return json.Unmarshal(b, t.PremiumFeatureChecklists)
	case "premiumFeaturePaidMessages":
		t.PremiumFeaturePaidMessages = new(PremiumFeaturePaidMessages)
		return json.Unmarshal(b, t.PremiumFeaturePaidMessages)
	}
	return nil
}

func (t *PremiumFeature) MarshalJSON() ([]byte, error) {
	if t.PremiumFeatureIncreasedLimits != nil {
		return json.Marshal(t.PremiumFeatureIncreasedLimits)
	}
	if t.PremiumFeatureIncreasedUploadFileSize != nil {
		return json.Marshal(t.PremiumFeatureIncreasedUploadFileSize)
	}
	if t.PremiumFeatureImprovedDownloadSpeed != nil {
		return json.Marshal(t.PremiumFeatureImprovedDownloadSpeed)
	}
	if t.PremiumFeatureVoiceRecognition != nil {
		return json.Marshal(t.PremiumFeatureVoiceRecognition)
	}
	if t.PremiumFeatureDisabledAds != nil {
		return json.Marshal(t.PremiumFeatureDisabledAds)
	}
	if t.PremiumFeatureUniqueReactions != nil {
		return json.Marshal(t.PremiumFeatureUniqueReactions)
	}
	if t.PremiumFeatureUniqueStickers != nil {
		return json.Marshal(t.PremiumFeatureUniqueStickers)
	}
	if t.PremiumFeatureCustomEmoji != nil {
		return json.Marshal(t.PremiumFeatureCustomEmoji)
	}
	if t.PremiumFeatureAdvancedChatManagement != nil {
		return json.Marshal(t.PremiumFeatureAdvancedChatManagement)
	}
	if t.PremiumFeatureProfileBadge != nil {
		return json.Marshal(t.PremiumFeatureProfileBadge)
	}
	if t.PremiumFeatureEmojiStatus != nil {
		return json.Marshal(t.PremiumFeatureEmojiStatus)
	}
	if t.PremiumFeatureAnimatedProfilePhoto != nil {
		return json.Marshal(t.PremiumFeatureAnimatedProfilePhoto)
	}
	if t.PremiumFeatureForumTopicIcon != nil {
		return json.Marshal(t.PremiumFeatureForumTopicIcon)
	}
	if t.PremiumFeatureAppIcons != nil {
		return json.Marshal(t.PremiumFeatureAppIcons)
	}
	if t.PremiumFeatureRealTimeChatTranslation != nil {
		return json.Marshal(t.PremiumFeatureRealTimeChatTranslation)
	}
	if t.PremiumFeatureUpgradedStories != nil {
		return json.Marshal(t.PremiumFeatureUpgradedStories)
	}
	if t.PremiumFeatureChatBoost != nil {
		return json.Marshal(t.PremiumFeatureChatBoost)
	}
	if t.PremiumFeatureAccentColor != nil {
		return json.Marshal(t.PremiumFeatureAccentColor)
	}
	if t.PremiumFeatureBackgroundForBoth != nil {
		return json.Marshal(t.PremiumFeatureBackgroundForBoth)
	}
	if t.PremiumFeatureSavedMessagesTags != nil {
		return json.Marshal(t.PremiumFeatureSavedMessagesTags)
	}
	if t.PremiumFeatureMessagePrivacy != nil {
		return json.Marshal(t.PremiumFeatureMessagePrivacy)
	}
	if t.PremiumFeatureLastSeenTimes != nil {
		return json.Marshal(t.PremiumFeatureLastSeenTimes)
	}
	if t.PremiumFeatureBusiness != nil {
		return json.Marshal(t.PremiumFeatureBusiness)
	}
	if t.PremiumFeatureMessageEffects != nil {
		return json.Marshal(t.PremiumFeatureMessageEffects)
	}
	if t.PremiumFeatureChecklists != nil {
		return json.Marshal(t.PremiumFeatureChecklists)
	}
	if t.PremiumFeaturePaidMessages != nil {
		return json.Marshal(t.PremiumFeaturePaidMessages)
	}
	return nil, fmt.Errorf("no value set for PremiumFeature")
}

// StoreTransaction Describes an in-store transaction
type StoreTransaction struct {
	StoreTransactionAppStore   *StoreTransactionAppStore   `json:"storeTransactionAppStore,omitempty"`
	StoreTransactionGooglePlay *StoreTransactionGooglePlay `json:"storeTransactionGooglePlay,omitempty"`
}

func (t *StoreTransaction) Type() string {
	return ""
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

// TonTransactionType Describes type of transaction with Toncoins
type TonTransactionType struct {
	TonTransactionTypeFragmentDeposit      *TonTransactionTypeFragmentDeposit      `json:"tonTransactionTypeFragmentDeposit,omitempty"`
	TonTransactionTypeFragmentWithdrawal   *TonTransactionTypeFragmentWithdrawal   `json:"tonTransactionTypeFragmentWithdrawal,omitempty"`
	TonTransactionTypeSuggestedPostPayment *TonTransactionTypeSuggestedPostPayment `json:"tonTransactionTypeSuggestedPostPayment,omitempty"`
	TonTransactionTypeGiftPurchaseOffer    *TonTransactionTypeGiftPurchaseOffer    `json:"tonTransactionTypeGiftPurchaseOffer,omitempty"`
	TonTransactionTypeUpgradedGiftPurchase *TonTransactionTypeUpgradedGiftPurchase `json:"tonTransactionTypeUpgradedGiftPurchase,omitempty"`
	TonTransactionTypeUpgradedGiftSale     *TonTransactionTypeUpgradedGiftSale     `json:"tonTransactionTypeUpgradedGiftSale,omitempty"`
	TonTransactionTypeUnsupported          *TonTransactionTypeUnsupported          `json:"tonTransactionTypeUnsupported,omitempty"`
}

func (t *TonTransactionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "tonTransactionTypeFragmentDeposit":
		t.TonTransactionTypeFragmentDeposit = new(TonTransactionTypeFragmentDeposit)
		return json.Unmarshal(b, t.TonTransactionTypeFragmentDeposit)
	case "tonTransactionTypeFragmentWithdrawal":
		t.TonTransactionTypeFragmentWithdrawal = new(TonTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.TonTransactionTypeFragmentWithdrawal)
	case "tonTransactionTypeSuggestedPostPayment":
		t.TonTransactionTypeSuggestedPostPayment = new(TonTransactionTypeSuggestedPostPayment)
		return json.Unmarshal(b, t.TonTransactionTypeSuggestedPostPayment)
	case "tonTransactionTypeGiftPurchaseOffer":
		t.TonTransactionTypeGiftPurchaseOffer = new(TonTransactionTypeGiftPurchaseOffer)
		return json.Unmarshal(b, t.TonTransactionTypeGiftPurchaseOffer)
	case "tonTransactionTypeUpgradedGiftPurchase":
		t.TonTransactionTypeUpgradedGiftPurchase = new(TonTransactionTypeUpgradedGiftPurchase)
		return json.Unmarshal(b, t.TonTransactionTypeUpgradedGiftPurchase)
	case "tonTransactionTypeUpgradedGiftSale":
		t.TonTransactionTypeUpgradedGiftSale = new(TonTransactionTypeUpgradedGiftSale)
		return json.Unmarshal(b, t.TonTransactionTypeUpgradedGiftSale)
	case "tonTransactionTypeUnsupported":
		t.TonTransactionTypeUnsupported = new(TonTransactionTypeUnsupported)
		return json.Unmarshal(b, t.TonTransactionTypeUnsupported)
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
	if t.TonTransactionTypeSuggestedPostPayment != nil {
		return json.Marshal(t.TonTransactionTypeSuggestedPostPayment)
	}
	if t.TonTransactionTypeGiftPurchaseOffer != nil {
		return json.Marshal(t.TonTransactionTypeGiftPurchaseOffer)
	}
	if t.TonTransactionTypeUpgradedGiftPurchase != nil {
		return json.Marshal(t.TonTransactionTypeUpgradedGiftPurchase)
	}
	if t.TonTransactionTypeUpgradedGiftSale != nil {
		return json.Marshal(t.TonTransactionTypeUpgradedGiftSale)
	}
	if t.TonTransactionTypeUnsupported != nil {
		return json.Marshal(t.TonTransactionTypeUnsupported)
	}
	return nil, fmt.Errorf("no value set for TonTransactionType")
}

// PaidMedia Describes a paid media
type PaidMedia struct {
	PaidMediaPreview     *PaidMediaPreview     `json:"paidMediaPreview,omitempty"`
	PaidMediaPhoto       *PaidMediaPhoto       `json:"paidMediaPhoto,omitempty"`
	PaidMediaVideo       *PaidMediaVideo       `json:"paidMediaVideo,omitempty"`
	PaidMediaUnsupported *PaidMediaUnsupported `json:"paidMediaUnsupported,omitempty"`
}

func (t *PaidMedia) Type() string {
	return ""
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
	switch typeObj.Type {
	case "paidMediaPreview":
		t.PaidMediaPreview = new(PaidMediaPreview)
		return json.Unmarshal(b, t.PaidMediaPreview)
	case "paidMediaPhoto":
		t.PaidMediaPhoto = new(PaidMediaPhoto)
		return json.Unmarshal(b, t.PaidMediaPhoto)
	case "paidMediaVideo":
		t.PaidMediaVideo = new(PaidMediaVideo)
		return json.Unmarshal(b, t.PaidMediaVideo)
	case "paidMediaUnsupported":
		t.PaidMediaUnsupported = new(PaidMediaUnsupported)
		return json.Unmarshal(b, t.PaidMediaUnsupported)
	}
	return nil
}

func (t *PaidMedia) MarshalJSON() ([]byte, error) {
	if t.PaidMediaPreview != nil {
		return json.Marshal(t.PaidMediaPreview)
	}
	if t.PaidMediaPhoto != nil {
		return json.Marshal(t.PaidMediaPhoto)
	}
	if t.PaidMediaVideo != nil {
		return json.Marshal(t.PaidMediaVideo)
	}
	if t.PaidMediaUnsupported != nil {
		return json.Marshal(t.PaidMediaUnsupported)
	}
	return nil, fmt.Errorf("no value set for PaidMedia")
}

// PublicForward Describes a public forward or repost of a story
type PublicForward struct {
	PublicForwardMessage *PublicForwardMessage `json:"publicForwardMessage,omitempty"`
	PublicForwardStory   *PublicForwardStory   `json:"publicForwardStory,omitempty"`
}

func (t *PublicForward) Type() string {
	return ""
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

// PremiumStoryFeature Describes a story feature available to Premium users
type PremiumStoryFeature struct {
	PremiumStoryFeaturePriorityOrder            *PremiumStoryFeaturePriorityOrder            `json:"premiumStoryFeaturePriorityOrder,omitempty"`
	PremiumStoryFeatureStealthMode              *PremiumStoryFeatureStealthMode              `json:"premiumStoryFeatureStealthMode,omitempty"`
	PremiumStoryFeaturePermanentViewsHistory    *PremiumStoryFeaturePermanentViewsHistory    `json:"premiumStoryFeaturePermanentViewsHistory,omitempty"`
	PremiumStoryFeatureCustomExpirationDuration *PremiumStoryFeatureCustomExpirationDuration `json:"premiumStoryFeatureCustomExpirationDuration,omitempty"`
	PremiumStoryFeatureSaveStories              *PremiumStoryFeatureSaveStories              `json:"premiumStoryFeatureSaveStories,omitempty"`
	PremiumStoryFeatureLinksAndFormatting       *PremiumStoryFeatureLinksAndFormatting       `json:"premiumStoryFeatureLinksAndFormatting,omitempty"`
	PremiumStoryFeatureVideoQuality             *PremiumStoryFeatureVideoQuality             `json:"premiumStoryFeatureVideoQuality,omitempty"`
}

func (t *PremiumStoryFeature) Type() string {
	return ""
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
	switch typeObj.Type {
	case "premiumStoryFeaturePriorityOrder":
		t.PremiumStoryFeaturePriorityOrder = new(PremiumStoryFeaturePriorityOrder)
		return json.Unmarshal(b, t.PremiumStoryFeaturePriorityOrder)
	case "premiumStoryFeatureStealthMode":
		t.PremiumStoryFeatureStealthMode = new(PremiumStoryFeatureStealthMode)
		return json.Unmarshal(b, t.PremiumStoryFeatureStealthMode)
	case "premiumStoryFeaturePermanentViewsHistory":
		t.PremiumStoryFeaturePermanentViewsHistory = new(PremiumStoryFeaturePermanentViewsHistory)
		return json.Unmarshal(b, t.PremiumStoryFeaturePermanentViewsHistory)
	case "premiumStoryFeatureCustomExpirationDuration":
		t.PremiumStoryFeatureCustomExpirationDuration = new(PremiumStoryFeatureCustomExpirationDuration)
		return json.Unmarshal(b, t.PremiumStoryFeatureCustomExpirationDuration)
	case "premiumStoryFeatureSaveStories":
		t.PremiumStoryFeatureSaveStories = new(PremiumStoryFeatureSaveStories)
		return json.Unmarshal(b, t.PremiumStoryFeatureSaveStories)
	case "premiumStoryFeatureLinksAndFormatting":
		t.PremiumStoryFeatureLinksAndFormatting = new(PremiumStoryFeatureLinksAndFormatting)
		return json.Unmarshal(b, t.PremiumStoryFeatureLinksAndFormatting)
	case "premiumStoryFeatureVideoQuality":
		t.PremiumStoryFeatureVideoQuality = new(PremiumStoryFeatureVideoQuality)
		return json.Unmarshal(b, t.PremiumStoryFeatureVideoQuality)
	}
	return nil
}

func (t *PremiumStoryFeature) MarshalJSON() ([]byte, error) {
	if t.PremiumStoryFeaturePriorityOrder != nil {
		return json.Marshal(t.PremiumStoryFeaturePriorityOrder)
	}
	if t.PremiumStoryFeatureStealthMode != nil {
		return json.Marshal(t.PremiumStoryFeatureStealthMode)
	}
	if t.PremiumStoryFeaturePermanentViewsHistory != nil {
		return json.Marshal(t.PremiumStoryFeaturePermanentViewsHistory)
	}
	if t.PremiumStoryFeatureCustomExpirationDuration != nil {
		return json.Marshal(t.PremiumStoryFeatureCustomExpirationDuration)
	}
	if t.PremiumStoryFeatureSaveStories != nil {
		return json.Marshal(t.PremiumStoryFeatureSaveStories)
	}
	if t.PremiumStoryFeatureLinksAndFormatting != nil {
		return json.Marshal(t.PremiumStoryFeatureLinksAndFormatting)
	}
	if t.PremiumStoryFeatureVideoQuality != nil {
		return json.Marshal(t.PremiumStoryFeatureVideoQuality)
	}
	return nil, fmt.Errorf("no value set for PremiumStoryFeature")
}

// UserPrivacySetting Describes available user privacy settings
type UserPrivacySetting struct {
	UserPrivacySettingShowStatus                            *UserPrivacySettingShowStatus                            `json:"userPrivacySettingShowStatus,omitempty"`
	UserPrivacySettingShowProfilePhoto                      *UserPrivacySettingShowProfilePhoto                      `json:"userPrivacySettingShowProfilePhoto,omitempty"`
	UserPrivacySettingShowLinkInForwardedMessages           *UserPrivacySettingShowLinkInForwardedMessages           `json:"userPrivacySettingShowLinkInForwardedMessages,omitempty"`
	UserPrivacySettingShowPhoneNumber                       *UserPrivacySettingShowPhoneNumber                       `json:"userPrivacySettingShowPhoneNumber,omitempty"`
	UserPrivacySettingShowBio                               *UserPrivacySettingShowBio                               `json:"userPrivacySettingShowBio,omitempty"`
	UserPrivacySettingShowBirthdate                         *UserPrivacySettingShowBirthdate                         `json:"userPrivacySettingShowBirthdate,omitempty"`
	UserPrivacySettingShowProfileAudio                      *UserPrivacySettingShowProfileAudio                      `json:"userPrivacySettingShowProfileAudio,omitempty"`
	UserPrivacySettingAllowChatInvites                      *UserPrivacySettingAllowChatInvites                      `json:"userPrivacySettingAllowChatInvites,omitempty"`
	UserPrivacySettingAllowCalls                            *UserPrivacySettingAllowCalls                            `json:"userPrivacySettingAllowCalls,omitempty"`
	UserPrivacySettingAllowPeerToPeerCalls                  *UserPrivacySettingAllowPeerToPeerCalls                  `json:"userPrivacySettingAllowPeerToPeerCalls,omitempty"`
	UserPrivacySettingAllowFindingByPhoneNumber             *UserPrivacySettingAllowFindingByPhoneNumber             `json:"userPrivacySettingAllowFindingByPhoneNumber,omitempty"`
	UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages *UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages `json:"userPrivacySettingAllowPrivateVoiceAndVideoNoteMessages,omitempty"`
	UserPrivacySettingAutosaveGifts                         *UserPrivacySettingAutosaveGifts                         `json:"userPrivacySettingAutosaveGifts,omitempty"`
	UserPrivacySettingAllowUnpaidMessages                   *UserPrivacySettingAllowUnpaidMessages                   `json:"userPrivacySettingAllowUnpaidMessages,omitempty"`
}

func (t *UserPrivacySetting) Type() string {
	return ""
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
	switch typeObj.Type {
	case "userPrivacySettingShowStatus":
		t.UserPrivacySettingShowStatus = new(UserPrivacySettingShowStatus)
		return json.Unmarshal(b, t.UserPrivacySettingShowStatus)
	case "userPrivacySettingShowProfilePhoto":
		t.UserPrivacySettingShowProfilePhoto = new(UserPrivacySettingShowProfilePhoto)
		return json.Unmarshal(b, t.UserPrivacySettingShowProfilePhoto)
	case "userPrivacySettingShowLinkInForwardedMessages":
		t.UserPrivacySettingShowLinkInForwardedMessages = new(UserPrivacySettingShowLinkInForwardedMessages)
		return json.Unmarshal(b, t.UserPrivacySettingShowLinkInForwardedMessages)
	case "userPrivacySettingShowPhoneNumber":
		t.UserPrivacySettingShowPhoneNumber = new(UserPrivacySettingShowPhoneNumber)
		return json.Unmarshal(b, t.UserPrivacySettingShowPhoneNumber)
	case "userPrivacySettingShowBio":
		t.UserPrivacySettingShowBio = new(UserPrivacySettingShowBio)
		return json.Unmarshal(b, t.UserPrivacySettingShowBio)
	case "userPrivacySettingShowBirthdate":
		t.UserPrivacySettingShowBirthdate = new(UserPrivacySettingShowBirthdate)
		return json.Unmarshal(b, t.UserPrivacySettingShowBirthdate)
	case "userPrivacySettingShowProfileAudio":
		t.UserPrivacySettingShowProfileAudio = new(UserPrivacySettingShowProfileAudio)
		return json.Unmarshal(b, t.UserPrivacySettingShowProfileAudio)
	case "userPrivacySettingAllowChatInvites":
		t.UserPrivacySettingAllowChatInvites = new(UserPrivacySettingAllowChatInvites)
		return json.Unmarshal(b, t.UserPrivacySettingAllowChatInvites)
	case "userPrivacySettingAllowCalls":
		t.UserPrivacySettingAllowCalls = new(UserPrivacySettingAllowCalls)
		return json.Unmarshal(b, t.UserPrivacySettingAllowCalls)
	case "userPrivacySettingAllowPeerToPeerCalls":
		t.UserPrivacySettingAllowPeerToPeerCalls = new(UserPrivacySettingAllowPeerToPeerCalls)
		return json.Unmarshal(b, t.UserPrivacySettingAllowPeerToPeerCalls)
	case "userPrivacySettingAllowFindingByPhoneNumber":
		t.UserPrivacySettingAllowFindingByPhoneNumber = new(UserPrivacySettingAllowFindingByPhoneNumber)
		return json.Unmarshal(b, t.UserPrivacySettingAllowFindingByPhoneNumber)
	case "userPrivacySettingAllowPrivateVoiceAndVideoNoteMessages":
		t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages = new(UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
		return json.Unmarshal(b, t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
	case "userPrivacySettingAutosaveGifts":
		t.UserPrivacySettingAutosaveGifts = new(UserPrivacySettingAutosaveGifts)
		return json.Unmarshal(b, t.UserPrivacySettingAutosaveGifts)
	case "userPrivacySettingAllowUnpaidMessages":
		t.UserPrivacySettingAllowUnpaidMessages = new(UserPrivacySettingAllowUnpaidMessages)
		return json.Unmarshal(b, t.UserPrivacySettingAllowUnpaidMessages)
	}
	return nil
}

func (t *UserPrivacySetting) MarshalJSON() ([]byte, error) {
	if t.UserPrivacySettingShowStatus != nil {
		return json.Marshal(t.UserPrivacySettingShowStatus)
	}
	if t.UserPrivacySettingShowProfilePhoto != nil {
		return json.Marshal(t.UserPrivacySettingShowProfilePhoto)
	}
	if t.UserPrivacySettingShowLinkInForwardedMessages != nil {
		return json.Marshal(t.UserPrivacySettingShowLinkInForwardedMessages)
	}
	if t.UserPrivacySettingShowPhoneNumber != nil {
		return json.Marshal(t.UserPrivacySettingShowPhoneNumber)
	}
	if t.UserPrivacySettingShowBio != nil {
		return json.Marshal(t.UserPrivacySettingShowBio)
	}
	if t.UserPrivacySettingShowBirthdate != nil {
		return json.Marshal(t.UserPrivacySettingShowBirthdate)
	}
	if t.UserPrivacySettingShowProfileAudio != nil {
		return json.Marshal(t.UserPrivacySettingShowProfileAudio)
	}
	if t.UserPrivacySettingAllowChatInvites != nil {
		return json.Marshal(t.UserPrivacySettingAllowChatInvites)
	}
	if t.UserPrivacySettingAllowCalls != nil {
		return json.Marshal(t.UserPrivacySettingAllowCalls)
	}
	if t.UserPrivacySettingAllowPeerToPeerCalls != nil {
		return json.Marshal(t.UserPrivacySettingAllowPeerToPeerCalls)
	}
	if t.UserPrivacySettingAllowFindingByPhoneNumber != nil {
		return json.Marshal(t.UserPrivacySettingAllowFindingByPhoneNumber)
	}
	if t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages != nil {
		return json.Marshal(t.UserPrivacySettingAllowPrivateVoiceAndVideoNoteMessages)
	}
	if t.UserPrivacySettingAutosaveGifts != nil {
		return json.Marshal(t.UserPrivacySettingAutosaveGifts)
	}
	if t.UserPrivacySettingAllowUnpaidMessages != nil {
		return json.Marshal(t.UserPrivacySettingAllowUnpaidMessages)
	}
	return nil, fmt.Errorf("no value set for UserPrivacySetting")
}

// ChatStatistics Contains a detailed statistics about a chat
type ChatStatistics struct {
	ChatStatisticsSupergroup *ChatStatisticsSupergroup `json:"chatStatisticsSupergroup,omitempty"`
	ChatStatisticsChannel    *ChatStatisticsChannel    `json:"chatStatisticsChannel,omitempty"`
}

func (t *ChatStatistics) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatStatisticsSupergroup":
		t.ChatStatisticsSupergroup = new(ChatStatisticsSupergroup)
		return json.Unmarshal(b, t.ChatStatisticsSupergroup)
	case "chatStatisticsChannel":
		t.ChatStatisticsChannel = new(ChatStatisticsChannel)
		return json.Unmarshal(b, t.ChatStatisticsChannel)
	}
	return nil
}

func (t *ChatStatistics) MarshalJSON() ([]byte, error) {
	if t.ChatStatisticsSupergroup != nil {
		return json.Marshal(t.ChatStatisticsSupergroup)
	}
	if t.ChatStatisticsChannel != nil {
		return json.Marshal(t.ChatStatisticsChannel)
	}
	return nil, fmt.Errorf("no value set for ChatStatistics")
}

// VectorPathCommand Represents a vector path command
type VectorPathCommand struct {
	VectorPathCommandLine             *VectorPathCommandLine             `json:"vectorPathCommandLine,omitempty"`
	VectorPathCommandCubicBezierCurve *VectorPathCommandCubicBezierCurve `json:"vectorPathCommandCubicBezierCurve,omitempty"`
}

func (t *VectorPathCommand) Type() string {
	return ""
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
	switch typeObj.Type {
	case "vectorPathCommandLine":
		t.VectorPathCommandLine = new(VectorPathCommandLine)
		return json.Unmarshal(b, t.VectorPathCommandLine)
	case "vectorPathCommandCubicBezierCurve":
		t.VectorPathCommandCubicBezierCurve = new(VectorPathCommandCubicBezierCurve)
		return json.Unmarshal(b, t.VectorPathCommandCubicBezierCurve)
	}
	return nil
}

func (t *VectorPathCommand) MarshalJSON() ([]byte, error) {
	if t.VectorPathCommandLine != nil {
		return json.Marshal(t.VectorPathCommandLine)
	}
	if t.VectorPathCommandCubicBezierCurve != nil {
		return json.Marshal(t.VectorPathCommandCubicBezierCurve)
	}
	return nil, fmt.Errorf("no value set for VectorPathCommand")
}

// ThumbnailFormat Describes format of a thumbnail
type ThumbnailFormat struct {
	ThumbnailFormatJpeg  *ThumbnailFormatJpeg  `json:"thumbnailFormatJpeg,omitempty"`
	ThumbnailFormatGif   *ThumbnailFormatGif   `json:"thumbnailFormatGif,omitempty"`
	ThumbnailFormatMpeg4 *ThumbnailFormatMpeg4 `json:"thumbnailFormatMpeg4,omitempty"`
	ThumbnailFormatPng   *ThumbnailFormatPng   `json:"thumbnailFormatPng,omitempty"`
	ThumbnailFormatTgs   *ThumbnailFormatTgs   `json:"thumbnailFormatTgs,omitempty"`
	ThumbnailFormatWebm  *ThumbnailFormatWebm  `json:"thumbnailFormatWebm,omitempty"`
	ThumbnailFormatWebp  *ThumbnailFormatWebp  `json:"thumbnailFormatWebp,omitempty"`
}

func (t *ThumbnailFormat) Type() string {
	return ""
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
	switch typeObj.Type {
	case "thumbnailFormatJpeg":
		t.ThumbnailFormatJpeg = new(ThumbnailFormatJpeg)
		return json.Unmarshal(b, t.ThumbnailFormatJpeg)
	case "thumbnailFormatGif":
		t.ThumbnailFormatGif = new(ThumbnailFormatGif)
		return json.Unmarshal(b, t.ThumbnailFormatGif)
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
	if t.ThumbnailFormatJpeg != nil {
		return json.Marshal(t.ThumbnailFormatJpeg)
	}
	if t.ThumbnailFormatGif != nil {
		return json.Marshal(t.ThumbnailFormatGif)
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

// GiveawayPrize Contains information about a giveaway prize
type GiveawayPrize struct {
	GiveawayPrizePremium *GiveawayPrizePremium `json:"giveawayPrizePremium,omitempty"`
	GiveawayPrizeStars   *GiveawayPrizeStars   `json:"giveawayPrizeStars,omitempty"`
}

func (t *GiveawayPrize) Type() string {
	return ""
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

// MessageReadDate Describes read date of a recent outgoing message in a private chat
type MessageReadDate struct {
	MessageReadDateRead                  *MessageReadDateRead                  `json:"messageReadDateRead,omitempty"`
	MessageReadDateUnread                *MessageReadDateUnread                `json:"messageReadDateUnread,omitempty"`
	MessageReadDateTooOld                *MessageReadDateTooOld                `json:"messageReadDateTooOld,omitempty"`
	MessageReadDateUserPrivacyRestricted *MessageReadDateUserPrivacyRestricted `json:"messageReadDateUserPrivacyRestricted,omitempty"`
	MessageReadDateMyPrivacyRestricted   *MessageReadDateMyPrivacyRestricted   `json:"messageReadDateMyPrivacyRestricted,omitempty"`
}

func (t *MessageReadDate) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageReadDateRead":
		t.MessageReadDateRead = new(MessageReadDateRead)
		return json.Unmarshal(b, t.MessageReadDateRead)
	case "messageReadDateUnread":
		t.MessageReadDateUnread = new(MessageReadDateUnread)
		return json.Unmarshal(b, t.MessageReadDateUnread)
	case "messageReadDateTooOld":
		t.MessageReadDateTooOld = new(MessageReadDateTooOld)
		return json.Unmarshal(b, t.MessageReadDateTooOld)
	case "messageReadDateUserPrivacyRestricted":
		t.MessageReadDateUserPrivacyRestricted = new(MessageReadDateUserPrivacyRestricted)
		return json.Unmarshal(b, t.MessageReadDateUserPrivacyRestricted)
	case "messageReadDateMyPrivacyRestricted":
		t.MessageReadDateMyPrivacyRestricted = new(MessageReadDateMyPrivacyRestricted)
		return json.Unmarshal(b, t.MessageReadDateMyPrivacyRestricted)
	}
	return nil
}

func (t *MessageReadDate) MarshalJSON() ([]byte, error) {
	if t.MessageReadDateRead != nil {
		return json.Marshal(t.MessageReadDateRead)
	}
	if t.MessageReadDateUnread != nil {
		return json.Marshal(t.MessageReadDateUnread)
	}
	if t.MessageReadDateTooOld != nil {
		return json.Marshal(t.MessageReadDateTooOld)
	}
	if t.MessageReadDateUserPrivacyRestricted != nil {
		return json.Marshal(t.MessageReadDateUserPrivacyRestricted)
	}
	if t.MessageReadDateMyPrivacyRestricted != nil {
		return json.Marshal(t.MessageReadDateMyPrivacyRestricted)
	}
	return nil, fmt.Errorf("no value set for MessageReadDate")
}

// PassportElementType Contains the type of Telegram Passport element
type PassportElementType struct {
	PassportElementTypePersonalDetails       *PassportElementTypePersonalDetails       `json:"passportElementTypePersonalDetails,omitempty"`
	PassportElementTypePassport              *PassportElementTypePassport              `json:"passportElementTypePassport,omitempty"`
	PassportElementTypeDriverLicense         *PassportElementTypeDriverLicense         `json:"passportElementTypeDriverLicense,omitempty"`
	PassportElementTypeIdentityCard          *PassportElementTypeIdentityCard          `json:"passportElementTypeIdentityCard,omitempty"`
	PassportElementTypeInternalPassport      *PassportElementTypeInternalPassport      `json:"passportElementTypeInternalPassport,omitempty"`
	PassportElementTypeAddress               *PassportElementTypeAddress               `json:"passportElementTypeAddress,omitempty"`
	PassportElementTypeUtilityBill           *PassportElementTypeUtilityBill           `json:"passportElementTypeUtilityBill,omitempty"`
	PassportElementTypeBankStatement         *PassportElementTypeBankStatement         `json:"passportElementTypeBankStatement,omitempty"`
	PassportElementTypeRentalAgreement       *PassportElementTypeRentalAgreement       `json:"passportElementTypeRentalAgreement,omitempty"`
	PassportElementTypePassportRegistration  *PassportElementTypePassportRegistration  `json:"passportElementTypePassportRegistration,omitempty"`
	PassportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration `json:"passportElementTypeTemporaryRegistration,omitempty"`
	PassportElementTypePhoneNumber           *PassportElementTypePhoneNumber           `json:"passportElementTypePhoneNumber,omitempty"`
	PassportElementTypeEmailAddress          *PassportElementTypeEmailAddress          `json:"passportElementTypeEmailAddress,omitempty"`
}

func (t *PassportElementType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "passportElementTypePersonalDetails":
		t.PassportElementTypePersonalDetails = new(PassportElementTypePersonalDetails)
		return json.Unmarshal(b, t.PassportElementTypePersonalDetails)
	case "passportElementTypePassport":
		t.PassportElementTypePassport = new(PassportElementTypePassport)
		return json.Unmarshal(b, t.PassportElementTypePassport)
	case "passportElementTypeDriverLicense":
		t.PassportElementTypeDriverLicense = new(PassportElementTypeDriverLicense)
		return json.Unmarshal(b, t.PassportElementTypeDriverLicense)
	case "passportElementTypeIdentityCard":
		t.PassportElementTypeIdentityCard = new(PassportElementTypeIdentityCard)
		return json.Unmarshal(b, t.PassportElementTypeIdentityCard)
	case "passportElementTypeInternalPassport":
		t.PassportElementTypeInternalPassport = new(PassportElementTypeInternalPassport)
		return json.Unmarshal(b, t.PassportElementTypeInternalPassport)
	case "passportElementTypeAddress":
		t.PassportElementTypeAddress = new(PassportElementTypeAddress)
		return json.Unmarshal(b, t.PassportElementTypeAddress)
	case "passportElementTypeUtilityBill":
		t.PassportElementTypeUtilityBill = new(PassportElementTypeUtilityBill)
		return json.Unmarshal(b, t.PassportElementTypeUtilityBill)
	case "passportElementTypeBankStatement":
		t.PassportElementTypeBankStatement = new(PassportElementTypeBankStatement)
		return json.Unmarshal(b, t.PassportElementTypeBankStatement)
	case "passportElementTypeRentalAgreement":
		t.PassportElementTypeRentalAgreement = new(PassportElementTypeRentalAgreement)
		return json.Unmarshal(b, t.PassportElementTypeRentalAgreement)
	case "passportElementTypePassportRegistration":
		t.PassportElementTypePassportRegistration = new(PassportElementTypePassportRegistration)
		return json.Unmarshal(b, t.PassportElementTypePassportRegistration)
	case "passportElementTypeTemporaryRegistration":
		t.PassportElementTypeTemporaryRegistration = new(PassportElementTypeTemporaryRegistration)
		return json.Unmarshal(b, t.PassportElementTypeTemporaryRegistration)
	case "passportElementTypePhoneNumber":
		t.PassportElementTypePhoneNumber = new(PassportElementTypePhoneNumber)
		return json.Unmarshal(b, t.PassportElementTypePhoneNumber)
	case "passportElementTypeEmailAddress":
		t.PassportElementTypeEmailAddress = new(PassportElementTypeEmailAddress)
		return json.Unmarshal(b, t.PassportElementTypeEmailAddress)
	}
	return nil
}

func (t *PassportElementType) MarshalJSON() ([]byte, error) {
	if t.PassportElementTypePersonalDetails != nil {
		return json.Marshal(t.PassportElementTypePersonalDetails)
	}
	if t.PassportElementTypePassport != nil {
		return json.Marshal(t.PassportElementTypePassport)
	}
	if t.PassportElementTypeDriverLicense != nil {
		return json.Marshal(t.PassportElementTypeDriverLicense)
	}
	if t.PassportElementTypeIdentityCard != nil {
		return json.Marshal(t.PassportElementTypeIdentityCard)
	}
	if t.PassportElementTypeInternalPassport != nil {
		return json.Marshal(t.PassportElementTypeInternalPassport)
	}
	if t.PassportElementTypeAddress != nil {
		return json.Marshal(t.PassportElementTypeAddress)
	}
	if t.PassportElementTypeUtilityBill != nil {
		return json.Marshal(t.PassportElementTypeUtilityBill)
	}
	if t.PassportElementTypeBankStatement != nil {
		return json.Marshal(t.PassportElementTypeBankStatement)
	}
	if t.PassportElementTypeRentalAgreement != nil {
		return json.Marshal(t.PassportElementTypeRentalAgreement)
	}
	if t.PassportElementTypePassportRegistration != nil {
		return json.Marshal(t.PassportElementTypePassportRegistration)
	}
	if t.PassportElementTypeTemporaryRegistration != nil {
		return json.Marshal(t.PassportElementTypeTemporaryRegistration)
	}
	if t.PassportElementTypePhoneNumber != nil {
		return json.Marshal(t.PassportElementTypePhoneNumber)
	}
	if t.PassportElementTypeEmailAddress != nil {
		return json.Marshal(t.PassportElementTypeEmailAddress)
	}
	return nil, fmt.Errorf("no value set for PassportElementType")
}

// InputPassportElementErrorSource Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementErrorSource struct {
	InputPassportElementErrorSourceUnspecified      *InputPassportElementErrorSourceUnspecified      `json:"inputPassportElementErrorSourceUnspecified,omitempty"`
	InputPassportElementErrorSourceDataField        *InputPassportElementErrorSourceDataField        `json:"inputPassportElementErrorSourceDataField,omitempty"`
	InputPassportElementErrorSourceFrontSide        *InputPassportElementErrorSourceFrontSide        `json:"inputPassportElementErrorSourceFrontSide,omitempty"`
	InputPassportElementErrorSourceReverseSide      *InputPassportElementErrorSourceReverseSide      `json:"inputPassportElementErrorSourceReverseSide,omitempty"`
	InputPassportElementErrorSourceSelfie           *InputPassportElementErrorSourceSelfie           `json:"inputPassportElementErrorSourceSelfie,omitempty"`
	InputPassportElementErrorSourceTranslationFile  *InputPassportElementErrorSourceTranslationFile  `json:"inputPassportElementErrorSourceTranslationFile,omitempty"`
	InputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles `json:"inputPassportElementErrorSourceTranslationFiles,omitempty"`
	InputPassportElementErrorSourceFile             *InputPassportElementErrorSourceFile             `json:"inputPassportElementErrorSourceFile,omitempty"`
	InputPassportElementErrorSourceFiles            *InputPassportElementErrorSourceFiles            `json:"inputPassportElementErrorSourceFiles,omitempty"`
}

func (t *InputPassportElementErrorSource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputPassportElementErrorSourceUnspecified":
		t.InputPassportElementErrorSourceUnspecified = new(InputPassportElementErrorSourceUnspecified)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceUnspecified)
	case "inputPassportElementErrorSourceDataField":
		t.InputPassportElementErrorSourceDataField = new(InputPassportElementErrorSourceDataField)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceDataField)
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
	case "inputPassportElementErrorSourceFile":
		t.InputPassportElementErrorSourceFile = new(InputPassportElementErrorSourceFile)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceFile)
	case "inputPassportElementErrorSourceFiles":
		t.InputPassportElementErrorSourceFiles = new(InputPassportElementErrorSourceFiles)
		return json.Unmarshal(b, t.InputPassportElementErrorSourceFiles)
	}
	return nil
}

func (t *InputPassportElementErrorSource) MarshalJSON() ([]byte, error) {
	if t.InputPassportElementErrorSourceUnspecified != nil {
		return json.Marshal(t.InputPassportElementErrorSourceUnspecified)
	}
	if t.InputPassportElementErrorSourceDataField != nil {
		return json.Marshal(t.InputPassportElementErrorSourceDataField)
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
	if t.InputPassportElementErrorSourceFile != nil {
		return json.Marshal(t.InputPassportElementErrorSourceFile)
	}
	if t.InputPassportElementErrorSourceFiles != nil {
		return json.Marshal(t.InputPassportElementErrorSourceFiles)
	}
	return nil, fmt.Errorf("no value set for InputPassportElementErrorSource")
}

// PremiumLimitType Describes type of limit, increased for Premium users
type PremiumLimitType struct {
	PremiumLimitTypeSupergroupCount                 *PremiumLimitTypeSupergroupCount                 `json:"premiumLimitTypeSupergroupCount,omitempty"`
	PremiumLimitTypePinnedChatCount                 *PremiumLimitTypePinnedChatCount                 `json:"premiumLimitTypePinnedChatCount,omitempty"`
	PremiumLimitTypeCreatedPublicChatCount          *PremiumLimitTypeCreatedPublicChatCount          `json:"premiumLimitTypeCreatedPublicChatCount,omitempty"`
	PremiumLimitTypeSavedAnimationCount             *PremiumLimitTypeSavedAnimationCount             `json:"premiumLimitTypeSavedAnimationCount,omitempty"`
	PremiumLimitTypeFavoriteStickerCount            *PremiumLimitTypeFavoriteStickerCount            `json:"premiumLimitTypeFavoriteStickerCount,omitempty"`
	PremiumLimitTypeChatFolderCount                 *PremiumLimitTypeChatFolderCount                 `json:"premiumLimitTypeChatFolderCount,omitempty"`
	PremiumLimitTypeChatFolderChosenChatCount       *PremiumLimitTypeChatFolderChosenChatCount       `json:"premiumLimitTypeChatFolderChosenChatCount,omitempty"`
	PremiumLimitTypePinnedArchivedChatCount         *PremiumLimitTypePinnedArchivedChatCount         `json:"premiumLimitTypePinnedArchivedChatCount,omitempty"`
	PremiumLimitTypePinnedSavedMessagesTopicCount   *PremiumLimitTypePinnedSavedMessagesTopicCount   `json:"premiumLimitTypePinnedSavedMessagesTopicCount,omitempty"`
	PremiumLimitTypeCaptionLength                   *PremiumLimitTypeCaptionLength                   `json:"premiumLimitTypeCaptionLength,omitempty"`
	PremiumLimitTypeBioLength                       *PremiumLimitTypeBioLength                       `json:"premiumLimitTypeBioLength,omitempty"`
	PremiumLimitTypeChatFolderInviteLinkCount       *PremiumLimitTypeChatFolderInviteLinkCount       `json:"premiumLimitTypeChatFolderInviteLinkCount,omitempty"`
	PremiumLimitTypeShareableChatFolderCount        *PremiumLimitTypeShareableChatFolderCount        `json:"premiumLimitTypeShareableChatFolderCount,omitempty"`
	PremiumLimitTypeActiveStoryCount                *PremiumLimitTypeActiveStoryCount                `json:"premiumLimitTypeActiveStoryCount,omitempty"`
	PremiumLimitTypeWeeklyPostedStoryCount          *PremiumLimitTypeWeeklyPostedStoryCount          `json:"premiumLimitTypeWeeklyPostedStoryCount,omitempty"`
	PremiumLimitTypeMonthlyPostedStoryCount         *PremiumLimitTypeMonthlyPostedStoryCount         `json:"premiumLimitTypeMonthlyPostedStoryCount,omitempty"`
	PremiumLimitTypeStoryCaptionLength              *PremiumLimitTypeStoryCaptionLength              `json:"premiumLimitTypeStoryCaptionLength,omitempty"`
	PremiumLimitTypeStorySuggestedReactionAreaCount *PremiumLimitTypeStorySuggestedReactionAreaCount `json:"premiumLimitTypeStorySuggestedReactionAreaCount,omitempty"`
	PremiumLimitTypeSimilarChatCount                *PremiumLimitTypeSimilarChatCount                `json:"premiumLimitTypeSimilarChatCount,omitempty"`
}

func (t *PremiumLimitType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "premiumLimitTypeSupergroupCount":
		t.PremiumLimitTypeSupergroupCount = new(PremiumLimitTypeSupergroupCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSupergroupCount)
	case "premiumLimitTypePinnedChatCount":
		t.PremiumLimitTypePinnedChatCount = new(PremiumLimitTypePinnedChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedChatCount)
	case "premiumLimitTypeCreatedPublicChatCount":
		t.PremiumLimitTypeCreatedPublicChatCount = new(PremiumLimitTypeCreatedPublicChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeCreatedPublicChatCount)
	case "premiumLimitTypeSavedAnimationCount":
		t.PremiumLimitTypeSavedAnimationCount = new(PremiumLimitTypeSavedAnimationCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSavedAnimationCount)
	case "premiumLimitTypeFavoriteStickerCount":
		t.PremiumLimitTypeFavoriteStickerCount = new(PremiumLimitTypeFavoriteStickerCount)
		return json.Unmarshal(b, t.PremiumLimitTypeFavoriteStickerCount)
	case "premiumLimitTypeChatFolderCount":
		t.PremiumLimitTypeChatFolderCount = new(PremiumLimitTypeChatFolderCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderCount)
	case "premiumLimitTypeChatFolderChosenChatCount":
		t.PremiumLimitTypeChatFolderChosenChatCount = new(PremiumLimitTypeChatFolderChosenChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderChosenChatCount)
	case "premiumLimitTypePinnedArchivedChatCount":
		t.PremiumLimitTypePinnedArchivedChatCount = new(PremiumLimitTypePinnedArchivedChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedArchivedChatCount)
	case "premiumLimitTypePinnedSavedMessagesTopicCount":
		t.PremiumLimitTypePinnedSavedMessagesTopicCount = new(PremiumLimitTypePinnedSavedMessagesTopicCount)
		return json.Unmarshal(b, t.PremiumLimitTypePinnedSavedMessagesTopicCount)
	case "premiumLimitTypeCaptionLength":
		t.PremiumLimitTypeCaptionLength = new(PremiumLimitTypeCaptionLength)
		return json.Unmarshal(b, t.PremiumLimitTypeCaptionLength)
	case "premiumLimitTypeBioLength":
		t.PremiumLimitTypeBioLength = new(PremiumLimitTypeBioLength)
		return json.Unmarshal(b, t.PremiumLimitTypeBioLength)
	case "premiumLimitTypeChatFolderInviteLinkCount":
		t.PremiumLimitTypeChatFolderInviteLinkCount = new(PremiumLimitTypeChatFolderInviteLinkCount)
		return json.Unmarshal(b, t.PremiumLimitTypeChatFolderInviteLinkCount)
	case "premiumLimitTypeShareableChatFolderCount":
		t.PremiumLimitTypeShareableChatFolderCount = new(PremiumLimitTypeShareableChatFolderCount)
		return json.Unmarshal(b, t.PremiumLimitTypeShareableChatFolderCount)
	case "premiumLimitTypeActiveStoryCount":
		t.PremiumLimitTypeActiveStoryCount = new(PremiumLimitTypeActiveStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeActiveStoryCount)
	case "premiumLimitTypeWeeklyPostedStoryCount":
		t.PremiumLimitTypeWeeklyPostedStoryCount = new(PremiumLimitTypeWeeklyPostedStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeWeeklyPostedStoryCount)
	case "premiumLimitTypeMonthlyPostedStoryCount":
		t.PremiumLimitTypeMonthlyPostedStoryCount = new(PremiumLimitTypeMonthlyPostedStoryCount)
		return json.Unmarshal(b, t.PremiumLimitTypeMonthlyPostedStoryCount)
	case "premiumLimitTypeStoryCaptionLength":
		t.PremiumLimitTypeStoryCaptionLength = new(PremiumLimitTypeStoryCaptionLength)
		return json.Unmarshal(b, t.PremiumLimitTypeStoryCaptionLength)
	case "premiumLimitTypeStorySuggestedReactionAreaCount":
		t.PremiumLimitTypeStorySuggestedReactionAreaCount = new(PremiumLimitTypeStorySuggestedReactionAreaCount)
		return json.Unmarshal(b, t.PremiumLimitTypeStorySuggestedReactionAreaCount)
	case "premiumLimitTypeSimilarChatCount":
		t.PremiumLimitTypeSimilarChatCount = new(PremiumLimitTypeSimilarChatCount)
		return json.Unmarshal(b, t.PremiumLimitTypeSimilarChatCount)
	}
	return nil
}

func (t *PremiumLimitType) MarshalJSON() ([]byte, error) {
	if t.PremiumLimitTypeSupergroupCount != nil {
		return json.Marshal(t.PremiumLimitTypeSupergroupCount)
	}
	if t.PremiumLimitTypePinnedChatCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedChatCount)
	}
	if t.PremiumLimitTypeCreatedPublicChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeCreatedPublicChatCount)
	}
	if t.PremiumLimitTypeSavedAnimationCount != nil {
		return json.Marshal(t.PremiumLimitTypeSavedAnimationCount)
	}
	if t.PremiumLimitTypeFavoriteStickerCount != nil {
		return json.Marshal(t.PremiumLimitTypeFavoriteStickerCount)
	}
	if t.PremiumLimitTypeChatFolderCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderCount)
	}
	if t.PremiumLimitTypeChatFolderChosenChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderChosenChatCount)
	}
	if t.PremiumLimitTypePinnedArchivedChatCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedArchivedChatCount)
	}
	if t.PremiumLimitTypePinnedSavedMessagesTopicCount != nil {
		return json.Marshal(t.PremiumLimitTypePinnedSavedMessagesTopicCount)
	}
	if t.PremiumLimitTypeCaptionLength != nil {
		return json.Marshal(t.PremiumLimitTypeCaptionLength)
	}
	if t.PremiumLimitTypeBioLength != nil {
		return json.Marshal(t.PremiumLimitTypeBioLength)
	}
	if t.PremiumLimitTypeChatFolderInviteLinkCount != nil {
		return json.Marshal(t.PremiumLimitTypeChatFolderInviteLinkCount)
	}
	if t.PremiumLimitTypeShareableChatFolderCount != nil {
		return json.Marshal(t.PremiumLimitTypeShareableChatFolderCount)
	}
	if t.PremiumLimitTypeActiveStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeActiveStoryCount)
	}
	if t.PremiumLimitTypeWeeklyPostedStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeWeeklyPostedStoryCount)
	}
	if t.PremiumLimitTypeMonthlyPostedStoryCount != nil {
		return json.Marshal(t.PremiumLimitTypeMonthlyPostedStoryCount)
	}
	if t.PremiumLimitTypeStoryCaptionLength != nil {
		return json.Marshal(t.PremiumLimitTypeStoryCaptionLength)
	}
	if t.PremiumLimitTypeStorySuggestedReactionAreaCount != nil {
		return json.Marshal(t.PremiumLimitTypeStorySuggestedReactionAreaCount)
	}
	if t.PremiumLimitTypeSimilarChatCount != nil {
		return json.Marshal(t.PremiumLimitTypeSimilarChatCount)
	}
	return nil, fmt.Errorf("no value set for PremiumLimitType")
}

// PushMessageContent Contains content of a push message notification
type PushMessageContent struct {
	PushMessageContentHidden                      *PushMessageContentHidden                      `json:"pushMessageContentHidden,omitempty"`
	PushMessageContentAnimation                   *PushMessageContentAnimation                   `json:"pushMessageContentAnimation,omitempty"`
	PushMessageContentAudio                       *PushMessageContentAudio                       `json:"pushMessageContentAudio,omitempty"`
	PushMessageContentContact                     *PushMessageContentContact                     `json:"pushMessageContentContact,omitempty"`
	PushMessageContentContactRegistered           *PushMessageContentContactRegistered           `json:"pushMessageContentContactRegistered,omitempty"`
	PushMessageContentDocument                    *PushMessageContentDocument                    `json:"pushMessageContentDocument,omitempty"`
	PushMessageContentGame                        *PushMessageContentGame                        `json:"pushMessageContentGame,omitempty"`
	PushMessageContentGameScore                   *PushMessageContentGameScore                   `json:"pushMessageContentGameScore,omitempty"`
	PushMessageContentInvoice                     *PushMessageContentInvoice                     `json:"pushMessageContentInvoice,omitempty"`
	PushMessageContentLocation                    *PushMessageContentLocation                    `json:"pushMessageContentLocation,omitempty"`
	PushMessageContentPaidMedia                   *PushMessageContentPaidMedia                   `json:"pushMessageContentPaidMedia,omitempty"`
	PushMessageContentPhoto                       *PushMessageContentPhoto                       `json:"pushMessageContentPhoto,omitempty"`
	PushMessageContentPoll                        *PushMessageContentPoll                        `json:"pushMessageContentPoll,omitempty"`
	PushMessageContentPremiumGiftCode             *PushMessageContentPremiumGiftCode             `json:"pushMessageContentPremiumGiftCode,omitempty"`
	PushMessageContentGiveaway                    *PushMessageContentGiveaway                    `json:"pushMessageContentGiveaway,omitempty"`
	PushMessageContentGift                        *PushMessageContentGift                        `json:"pushMessageContentGift,omitempty"`
	PushMessageContentUpgradedGift                *PushMessageContentUpgradedGift                `json:"pushMessageContentUpgradedGift,omitempty"`
	PushMessageContentScreenshotTaken             *PushMessageContentScreenshotTaken             `json:"pushMessageContentScreenshotTaken,omitempty"`
	PushMessageContentSticker                     *PushMessageContentSticker                     `json:"pushMessageContentSticker,omitempty"`
	PushMessageContentStory                       *PushMessageContentStory                       `json:"pushMessageContentStory,omitempty"`
	PushMessageContentText                        *PushMessageContentText                        `json:"pushMessageContentText,omitempty"`
	PushMessageContentChecklist                   *PushMessageContentChecklist                   `json:"pushMessageContentChecklist,omitempty"`
	PushMessageContentVideo                       *PushMessageContentVideo                       `json:"pushMessageContentVideo,omitempty"`
	PushMessageContentVideoNote                   *PushMessageContentVideoNote                   `json:"pushMessageContentVideoNote,omitempty"`
	PushMessageContentVoiceNote                   *PushMessageContentVoiceNote                   `json:"pushMessageContentVoiceNote,omitempty"`
	PushMessageContentBasicGroupChatCreate        *PushMessageContentBasicGroupChatCreate        `json:"pushMessageContentBasicGroupChatCreate,omitempty"`
	PushMessageContentVideoChatStarted            *PushMessageContentVideoChatStarted            `json:"pushMessageContentVideoChatStarted,omitempty"`
	PushMessageContentVideoChatEnded              *PushMessageContentVideoChatEnded              `json:"pushMessageContentVideoChatEnded,omitempty"`
	PushMessageContentInviteVideoChatParticipants *PushMessageContentInviteVideoChatParticipants `json:"pushMessageContentInviteVideoChatParticipants,omitempty"`
	PushMessageContentChatAddMembers              *PushMessageContentChatAddMembers              `json:"pushMessageContentChatAddMembers,omitempty"`
	PushMessageContentChatChangePhoto             *PushMessageContentChatChangePhoto             `json:"pushMessageContentChatChangePhoto,omitempty"`
	PushMessageContentChatChangeTitle             *PushMessageContentChatChangeTitle             `json:"pushMessageContentChatChangeTitle,omitempty"`
	PushMessageContentChatSetBackground           *PushMessageContentChatSetBackground           `json:"pushMessageContentChatSetBackground,omitempty"`
	PushMessageContentChatSetTheme                *PushMessageContentChatSetTheme                `json:"pushMessageContentChatSetTheme,omitempty"`
	PushMessageContentChatDeleteMember            *PushMessageContentChatDeleteMember            `json:"pushMessageContentChatDeleteMember,omitempty"`
	PushMessageContentChatJoinByLink              *PushMessageContentChatJoinByLink              `json:"pushMessageContentChatJoinByLink,omitempty"`
	PushMessageContentChatJoinByRequest           *PushMessageContentChatJoinByRequest           `json:"pushMessageContentChatJoinByRequest,omitempty"`
	PushMessageContentRecurringPayment            *PushMessageContentRecurringPayment            `json:"pushMessageContentRecurringPayment,omitempty"`
	PushMessageContentSuggestProfilePhoto         *PushMessageContentSuggestProfilePhoto         `json:"pushMessageContentSuggestProfilePhoto,omitempty"`
	PushMessageContentSuggestBirthdate            *PushMessageContentSuggestBirthdate            `json:"pushMessageContentSuggestBirthdate,omitempty"`
	PushMessageContentProximityAlertTriggered     *PushMessageContentProximityAlertTriggered     `json:"pushMessageContentProximityAlertTriggered,omitempty"`
	PushMessageContentChecklistTasksAdded         *PushMessageContentChecklistTasksAdded         `json:"pushMessageContentChecklistTasksAdded,omitempty"`
	PushMessageContentChecklistTasksDone          *PushMessageContentChecklistTasksDone          `json:"pushMessageContentChecklistTasksDone,omitempty"`
	PushMessageContentMessageForwards             *PushMessageContentMessageForwards             `json:"pushMessageContentMessageForwards,omitempty"`
	PushMessageContentMediaAlbum                  *PushMessageContentMediaAlbum                  `json:"pushMessageContentMediaAlbum,omitempty"`
}

func (t *PushMessageContent) Type() string {
	return ""
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
	switch typeObj.Type {
	case "pushMessageContentHidden":
		t.PushMessageContentHidden = new(PushMessageContentHidden)
		return json.Unmarshal(b, t.PushMessageContentHidden)
	case "pushMessageContentAnimation":
		t.PushMessageContentAnimation = new(PushMessageContentAnimation)
		return json.Unmarshal(b, t.PushMessageContentAnimation)
	case "pushMessageContentAudio":
		t.PushMessageContentAudio = new(PushMessageContentAudio)
		return json.Unmarshal(b, t.PushMessageContentAudio)
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
	case "pushMessageContentInvoice":
		t.PushMessageContentInvoice = new(PushMessageContentInvoice)
		return json.Unmarshal(b, t.PushMessageContentInvoice)
	case "pushMessageContentLocation":
		t.PushMessageContentLocation = new(PushMessageContentLocation)
		return json.Unmarshal(b, t.PushMessageContentLocation)
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
	case "pushMessageContentGiveaway":
		t.PushMessageContentGiveaway = new(PushMessageContentGiveaway)
		return json.Unmarshal(b, t.PushMessageContentGiveaway)
	case "pushMessageContentGift":
		t.PushMessageContentGift = new(PushMessageContentGift)
		return json.Unmarshal(b, t.PushMessageContentGift)
	case "pushMessageContentUpgradedGift":
		t.PushMessageContentUpgradedGift = new(PushMessageContentUpgradedGift)
		return json.Unmarshal(b, t.PushMessageContentUpgradedGift)
	case "pushMessageContentScreenshotTaken":
		t.PushMessageContentScreenshotTaken = new(PushMessageContentScreenshotTaken)
		return json.Unmarshal(b, t.PushMessageContentScreenshotTaken)
	case "pushMessageContentSticker":
		t.PushMessageContentSticker = new(PushMessageContentSticker)
		return json.Unmarshal(b, t.PushMessageContentSticker)
	case "pushMessageContentStory":
		t.PushMessageContentStory = new(PushMessageContentStory)
		return json.Unmarshal(b, t.PushMessageContentStory)
	case "pushMessageContentText":
		t.PushMessageContentText = new(PushMessageContentText)
		return json.Unmarshal(b, t.PushMessageContentText)
	case "pushMessageContentChecklist":
		t.PushMessageContentChecklist = new(PushMessageContentChecklist)
		return json.Unmarshal(b, t.PushMessageContentChecklist)
	case "pushMessageContentVideo":
		t.PushMessageContentVideo = new(PushMessageContentVideo)
		return json.Unmarshal(b, t.PushMessageContentVideo)
	case "pushMessageContentVideoNote":
		t.PushMessageContentVideoNote = new(PushMessageContentVideoNote)
		return json.Unmarshal(b, t.PushMessageContentVideoNote)
	case "pushMessageContentVoiceNote":
		t.PushMessageContentVoiceNote = new(PushMessageContentVoiceNote)
		return json.Unmarshal(b, t.PushMessageContentVoiceNote)
	case "pushMessageContentBasicGroupChatCreate":
		t.PushMessageContentBasicGroupChatCreate = new(PushMessageContentBasicGroupChatCreate)
		return json.Unmarshal(b, t.PushMessageContentBasicGroupChatCreate)
	case "pushMessageContentVideoChatStarted":
		t.PushMessageContentVideoChatStarted = new(PushMessageContentVideoChatStarted)
		return json.Unmarshal(b, t.PushMessageContentVideoChatStarted)
	case "pushMessageContentVideoChatEnded":
		t.PushMessageContentVideoChatEnded = new(PushMessageContentVideoChatEnded)
		return json.Unmarshal(b, t.PushMessageContentVideoChatEnded)
	case "pushMessageContentInviteVideoChatParticipants":
		t.PushMessageContentInviteVideoChatParticipants = new(PushMessageContentInviteVideoChatParticipants)
		return json.Unmarshal(b, t.PushMessageContentInviteVideoChatParticipants)
	case "pushMessageContentChatAddMembers":
		t.PushMessageContentChatAddMembers = new(PushMessageContentChatAddMembers)
		return json.Unmarshal(b, t.PushMessageContentChatAddMembers)
	case "pushMessageContentChatChangePhoto":
		t.PushMessageContentChatChangePhoto = new(PushMessageContentChatChangePhoto)
		return json.Unmarshal(b, t.PushMessageContentChatChangePhoto)
	case "pushMessageContentChatChangeTitle":
		t.PushMessageContentChatChangeTitle = new(PushMessageContentChatChangeTitle)
		return json.Unmarshal(b, t.PushMessageContentChatChangeTitle)
	case "pushMessageContentChatSetBackground":
		t.PushMessageContentChatSetBackground = new(PushMessageContentChatSetBackground)
		return json.Unmarshal(b, t.PushMessageContentChatSetBackground)
	case "pushMessageContentChatSetTheme":
		t.PushMessageContentChatSetTheme = new(PushMessageContentChatSetTheme)
		return json.Unmarshal(b, t.PushMessageContentChatSetTheme)
	case "pushMessageContentChatDeleteMember":
		t.PushMessageContentChatDeleteMember = new(PushMessageContentChatDeleteMember)
		return json.Unmarshal(b, t.PushMessageContentChatDeleteMember)
	case "pushMessageContentChatJoinByLink":
		t.PushMessageContentChatJoinByLink = new(PushMessageContentChatJoinByLink)
		return json.Unmarshal(b, t.PushMessageContentChatJoinByLink)
	case "pushMessageContentChatJoinByRequest":
		t.PushMessageContentChatJoinByRequest = new(PushMessageContentChatJoinByRequest)
		return json.Unmarshal(b, t.PushMessageContentChatJoinByRequest)
	case "pushMessageContentRecurringPayment":
		t.PushMessageContentRecurringPayment = new(PushMessageContentRecurringPayment)
		return json.Unmarshal(b, t.PushMessageContentRecurringPayment)
	case "pushMessageContentSuggestProfilePhoto":
		t.PushMessageContentSuggestProfilePhoto = new(PushMessageContentSuggestProfilePhoto)
		return json.Unmarshal(b, t.PushMessageContentSuggestProfilePhoto)
	case "pushMessageContentSuggestBirthdate":
		t.PushMessageContentSuggestBirthdate = new(PushMessageContentSuggestBirthdate)
		return json.Unmarshal(b, t.PushMessageContentSuggestBirthdate)
	case "pushMessageContentProximityAlertTriggered":
		t.PushMessageContentProximityAlertTriggered = new(PushMessageContentProximityAlertTriggered)
		return json.Unmarshal(b, t.PushMessageContentProximityAlertTriggered)
	case "pushMessageContentChecklistTasksAdded":
		t.PushMessageContentChecklistTasksAdded = new(PushMessageContentChecklistTasksAdded)
		return json.Unmarshal(b, t.PushMessageContentChecklistTasksAdded)
	case "pushMessageContentChecklistTasksDone":
		t.PushMessageContentChecklistTasksDone = new(PushMessageContentChecklistTasksDone)
		return json.Unmarshal(b, t.PushMessageContentChecklistTasksDone)
	case "pushMessageContentMessageForwards":
		t.PushMessageContentMessageForwards = new(PushMessageContentMessageForwards)
		return json.Unmarshal(b, t.PushMessageContentMessageForwards)
	case "pushMessageContentMediaAlbum":
		t.PushMessageContentMediaAlbum = new(PushMessageContentMediaAlbum)
		return json.Unmarshal(b, t.PushMessageContentMediaAlbum)
	}
	return nil
}

func (t *PushMessageContent) MarshalJSON() ([]byte, error) {
	if t.PushMessageContentHidden != nil {
		return json.Marshal(t.PushMessageContentHidden)
	}
	if t.PushMessageContentAnimation != nil {
		return json.Marshal(t.PushMessageContentAnimation)
	}
	if t.PushMessageContentAudio != nil {
		return json.Marshal(t.PushMessageContentAudio)
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
	if t.PushMessageContentInvoice != nil {
		return json.Marshal(t.PushMessageContentInvoice)
	}
	if t.PushMessageContentLocation != nil {
		return json.Marshal(t.PushMessageContentLocation)
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
	if t.PushMessageContentGiveaway != nil {
		return json.Marshal(t.PushMessageContentGiveaway)
	}
	if t.PushMessageContentGift != nil {
		return json.Marshal(t.PushMessageContentGift)
	}
	if t.PushMessageContentUpgradedGift != nil {
		return json.Marshal(t.PushMessageContentUpgradedGift)
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
	if t.PushMessageContentText != nil {
		return json.Marshal(t.PushMessageContentText)
	}
	if t.PushMessageContentChecklist != nil {
		return json.Marshal(t.PushMessageContentChecklist)
	}
	if t.PushMessageContentVideo != nil {
		return json.Marshal(t.PushMessageContentVideo)
	}
	if t.PushMessageContentVideoNote != nil {
		return json.Marshal(t.PushMessageContentVideoNote)
	}
	if t.PushMessageContentVoiceNote != nil {
		return json.Marshal(t.PushMessageContentVoiceNote)
	}
	if t.PushMessageContentBasicGroupChatCreate != nil {
		return json.Marshal(t.PushMessageContentBasicGroupChatCreate)
	}
	if t.PushMessageContentVideoChatStarted != nil {
		return json.Marshal(t.PushMessageContentVideoChatStarted)
	}
	if t.PushMessageContentVideoChatEnded != nil {
		return json.Marshal(t.PushMessageContentVideoChatEnded)
	}
	if t.PushMessageContentInviteVideoChatParticipants != nil {
		return json.Marshal(t.PushMessageContentInviteVideoChatParticipants)
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
	if t.PushMessageContentChatSetBackground != nil {
		return json.Marshal(t.PushMessageContentChatSetBackground)
	}
	if t.PushMessageContentChatSetTheme != nil {
		return json.Marshal(t.PushMessageContentChatSetTheme)
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
	if t.PushMessageContentRecurringPayment != nil {
		return json.Marshal(t.PushMessageContentRecurringPayment)
	}
	if t.PushMessageContentSuggestProfilePhoto != nil {
		return json.Marshal(t.PushMessageContentSuggestProfilePhoto)
	}
	if t.PushMessageContentSuggestBirthdate != nil {
		return json.Marshal(t.PushMessageContentSuggestBirthdate)
	}
	if t.PushMessageContentProximityAlertTriggered != nil {
		return json.Marshal(t.PushMessageContentProximityAlertTriggered)
	}
	if t.PushMessageContentChecklistTasksAdded != nil {
		return json.Marshal(t.PushMessageContentChecklistTasksAdded)
	}
	if t.PushMessageContentChecklistTasksDone != nil {
		return json.Marshal(t.PushMessageContentChecklistTasksDone)
	}
	if t.PushMessageContentMessageForwards != nil {
		return json.Marshal(t.PushMessageContentMessageForwards)
	}
	if t.PushMessageContentMediaAlbum != nil {
		return json.Marshal(t.PushMessageContentMediaAlbum)
	}
	return nil, fmt.Errorf("no value set for PushMessageContent")
}

// BotCommandScope Represents the scope to which bot commands are relevant
type BotCommandScope struct {
	BotCommandScopeDefault               *BotCommandScopeDefault               `json:"botCommandScopeDefault,omitempty"`
	BotCommandScopeAllPrivateChats       *BotCommandScopeAllPrivateChats       `json:"botCommandScopeAllPrivateChats,omitempty"`
	BotCommandScopeAllGroupChats         *BotCommandScopeAllGroupChats         `json:"botCommandScopeAllGroupChats,omitempty"`
	BotCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators `json:"botCommandScopeAllChatAdministrators,omitempty"`
	BotCommandScopeChat                  *BotCommandScopeChat                  `json:"botCommandScopeChat,omitempty"`
	BotCommandScopeChatAdministrators    *BotCommandScopeChatAdministrators    `json:"botCommandScopeChatAdministrators,omitempty"`
	BotCommandScopeChatMember            *BotCommandScopeChatMember            `json:"botCommandScopeChatMember,omitempty"`
}

func (t *BotCommandScope) Type() string {
	return ""
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
	switch typeObj.Type {
	case "botCommandScopeDefault":
		t.BotCommandScopeDefault = new(BotCommandScopeDefault)
		return json.Unmarshal(b, t.BotCommandScopeDefault)
	case "botCommandScopeAllPrivateChats":
		t.BotCommandScopeAllPrivateChats = new(BotCommandScopeAllPrivateChats)
		return json.Unmarshal(b, t.BotCommandScopeAllPrivateChats)
	case "botCommandScopeAllGroupChats":
		t.BotCommandScopeAllGroupChats = new(BotCommandScopeAllGroupChats)
		return json.Unmarshal(b, t.BotCommandScopeAllGroupChats)
	case "botCommandScopeAllChatAdministrators":
		t.BotCommandScopeAllChatAdministrators = new(BotCommandScopeAllChatAdministrators)
		return json.Unmarshal(b, t.BotCommandScopeAllChatAdministrators)
	case "botCommandScopeChat":
		t.BotCommandScopeChat = new(BotCommandScopeChat)
		return json.Unmarshal(b, t.BotCommandScopeChat)
	case "botCommandScopeChatAdministrators":
		t.BotCommandScopeChatAdministrators = new(BotCommandScopeChatAdministrators)
		return json.Unmarshal(b, t.BotCommandScopeChatAdministrators)
	case "botCommandScopeChatMember":
		t.BotCommandScopeChatMember = new(BotCommandScopeChatMember)
		return json.Unmarshal(b, t.BotCommandScopeChatMember)
	}
	return nil
}

func (t *BotCommandScope) MarshalJSON() ([]byte, error) {
	if t.BotCommandScopeDefault != nil {
		return json.Marshal(t.BotCommandScopeDefault)
	}
	if t.BotCommandScopeAllPrivateChats != nil {
		return json.Marshal(t.BotCommandScopeAllPrivateChats)
	}
	if t.BotCommandScopeAllGroupChats != nil {
		return json.Marshal(t.BotCommandScopeAllGroupChats)
	}
	if t.BotCommandScopeAllChatAdministrators != nil {
		return json.Marshal(t.BotCommandScopeAllChatAdministrators)
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
	return nil, fmt.Errorf("no value set for BotCommandScope")
}

// InputInlineQueryResult Represents a single result of an inline query; for bots only
type InputInlineQueryResult struct {
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
	return ""
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

// TelegramPaymentPurpose Describes a purpose of a payment toward Telegram
type TelegramPaymentPurpose struct {
	TelegramPaymentPurposePremiumGift      *TelegramPaymentPurposePremiumGift      `json:"telegramPaymentPurposePremiumGift,omitempty"`
	TelegramPaymentPurposePremiumGiftCodes *TelegramPaymentPurposePremiumGiftCodes `json:"telegramPaymentPurposePremiumGiftCodes,omitempty"`
	TelegramPaymentPurposePremiumGiveaway  *TelegramPaymentPurposePremiumGiveaway  `json:"telegramPaymentPurposePremiumGiveaway,omitempty"`
	TelegramPaymentPurposeStars            *TelegramPaymentPurposeStars            `json:"telegramPaymentPurposeStars,omitempty"`
	TelegramPaymentPurposeGiftedStars      *TelegramPaymentPurposeGiftedStars      `json:"telegramPaymentPurposeGiftedStars,omitempty"`
	TelegramPaymentPurposeStarGiveaway     *TelegramPaymentPurposeStarGiveaway     `json:"telegramPaymentPurposeStarGiveaway,omitempty"`
	TelegramPaymentPurposeJoinChat         *TelegramPaymentPurposeJoinChat         `json:"telegramPaymentPurposeJoinChat,omitempty"`
}

func (t *TelegramPaymentPurpose) Type() string {
	return ""
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
	switch typeObj.Type {
	case "telegramPaymentPurposePremiumGift":
		t.TelegramPaymentPurposePremiumGift = new(TelegramPaymentPurposePremiumGift)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGift)
	case "telegramPaymentPurposePremiumGiftCodes":
		t.TelegramPaymentPurposePremiumGiftCodes = new(TelegramPaymentPurposePremiumGiftCodes)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGiftCodes)
	case "telegramPaymentPurposePremiumGiveaway":
		t.TelegramPaymentPurposePremiumGiveaway = new(TelegramPaymentPurposePremiumGiveaway)
		return json.Unmarshal(b, t.TelegramPaymentPurposePremiumGiveaway)
	case "telegramPaymentPurposeStars":
		t.TelegramPaymentPurposeStars = new(TelegramPaymentPurposeStars)
		return json.Unmarshal(b, t.TelegramPaymentPurposeStars)
	case "telegramPaymentPurposeGiftedStars":
		t.TelegramPaymentPurposeGiftedStars = new(TelegramPaymentPurposeGiftedStars)
		return json.Unmarshal(b, t.TelegramPaymentPurposeGiftedStars)
	case "telegramPaymentPurposeStarGiveaway":
		t.TelegramPaymentPurposeStarGiveaway = new(TelegramPaymentPurposeStarGiveaway)
		return json.Unmarshal(b, t.TelegramPaymentPurposeStarGiveaway)
	case "telegramPaymentPurposeJoinChat":
		t.TelegramPaymentPurposeJoinChat = new(TelegramPaymentPurposeJoinChat)
		return json.Unmarshal(b, t.TelegramPaymentPurposeJoinChat)
	}
	return nil
}

func (t *TelegramPaymentPurpose) MarshalJSON() ([]byte, error) {
	if t.TelegramPaymentPurposePremiumGift != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGift)
	}
	if t.TelegramPaymentPurposePremiumGiftCodes != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGiftCodes)
	}
	if t.TelegramPaymentPurposePremiumGiveaway != nil {
		return json.Marshal(t.TelegramPaymentPurposePremiumGiveaway)
	}
	if t.TelegramPaymentPurposeStars != nil {
		return json.Marshal(t.TelegramPaymentPurposeStars)
	}
	if t.TelegramPaymentPurposeGiftedStars != nil {
		return json.Marshal(t.TelegramPaymentPurposeGiftedStars)
	}
	if t.TelegramPaymentPurposeStarGiveaway != nil {
		return json.Marshal(t.TelegramPaymentPurposeStarGiveaway)
	}
	if t.TelegramPaymentPurposeJoinChat != nil {
		return json.Marshal(t.TelegramPaymentPurposeJoinChat)
	}
	return nil, fmt.Errorf("no value set for TelegramPaymentPurpose")
}

// BlockList Describes type of block list
type BlockList struct {
	BlockListMain    *BlockListMain    `json:"blockListMain,omitempty"`
	BlockListStories *BlockListStories `json:"blockListStories,omitempty"`
}

func (t *BlockList) Type() string {
	return ""
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

// NetworkStatisticsEntry Contains statistics about network usage
type NetworkStatisticsEntry struct {
	NetworkStatisticsEntryFile *NetworkStatisticsEntryFile `json:"networkStatisticsEntryFile,omitempty"`
	NetworkStatisticsEntryCall *NetworkStatisticsEntryCall `json:"networkStatisticsEntryCall,omitempty"`
}

func (t *NetworkStatisticsEntry) Type() string {
	return ""
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
	switch typeObj.Type {
	case "networkStatisticsEntryFile":
		t.NetworkStatisticsEntryFile = new(NetworkStatisticsEntryFile)
		return json.Unmarshal(b, t.NetworkStatisticsEntryFile)
	case "networkStatisticsEntryCall":
		t.NetworkStatisticsEntryCall = new(NetworkStatisticsEntryCall)
		return json.Unmarshal(b, t.NetworkStatisticsEntryCall)
	}
	return nil
}

func (t *NetworkStatisticsEntry) MarshalJSON() ([]byte, error) {
	if t.NetworkStatisticsEntryFile != nil {
		return json.Marshal(t.NetworkStatisticsEntryFile)
	}
	if t.NetworkStatisticsEntryCall != nil {
		return json.Marshal(t.NetworkStatisticsEntryCall)
	}
	return nil, fmt.Errorf("no value set for NetworkStatisticsEntry")
}

// GiveawayInfo Contains information about a giveaway
type GiveawayInfo struct {
	GiveawayInfoOngoing   *GiveawayInfoOngoing   `json:"giveawayInfoOngoing,omitempty"`
	GiveawayInfoCompleted *GiveawayInfoCompleted `json:"giveawayInfoCompleted,omitempty"`
}

func (t *GiveawayInfo) Type() string {
	return ""
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
	switch typeObj.Type {
	case "giveawayInfoOngoing":
		t.GiveawayInfoOngoing = new(GiveawayInfoOngoing)
		return json.Unmarshal(b, t.GiveawayInfoOngoing)
	case "giveawayInfoCompleted":
		t.GiveawayInfoCompleted = new(GiveawayInfoCompleted)
		return json.Unmarshal(b, t.GiveawayInfoCompleted)
	}
	return nil
}

func (t *GiveawayInfo) MarshalJSON() ([]byte, error) {
	if t.GiveawayInfoOngoing != nil {
		return json.Marshal(t.GiveawayInfoOngoing)
	}
	if t.GiveawayInfoCompleted != nil {
		return json.Marshal(t.GiveawayInfoCompleted)
	}
	return nil, fmt.Errorf("no value set for GiveawayInfo")
}

// NotificationSettingsScope Describes the types of chats to which notification settings are relevant
type NotificationSettingsScope struct {
	NotificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats `json:"notificationSettingsScopePrivateChats,omitempty"`
	NotificationSettingsScopeGroupChats   *NotificationSettingsScopeGroupChats   `json:"notificationSettingsScopeGroupChats,omitempty"`
	NotificationSettingsScopeChannelChats *NotificationSettingsScopeChannelChats `json:"notificationSettingsScopeChannelChats,omitempty"`
}

func (t *NotificationSettingsScope) Type() string {
	return ""
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
	switch typeObj.Type {
	case "notificationSettingsScopePrivateChats":
		t.NotificationSettingsScopePrivateChats = new(NotificationSettingsScopePrivateChats)
		return json.Unmarshal(b, t.NotificationSettingsScopePrivateChats)
	case "notificationSettingsScopeGroupChats":
		t.NotificationSettingsScopeGroupChats = new(NotificationSettingsScopeGroupChats)
		return json.Unmarshal(b, t.NotificationSettingsScopeGroupChats)
	case "notificationSettingsScopeChannelChats":
		t.NotificationSettingsScopeChannelChats = new(NotificationSettingsScopeChannelChats)
		return json.Unmarshal(b, t.NotificationSettingsScopeChannelChats)
	}
	return nil
}

func (t *NotificationSettingsScope) MarshalJSON() ([]byte, error) {
	if t.NotificationSettingsScopePrivateChats != nil {
		return json.Marshal(t.NotificationSettingsScopePrivateChats)
	}
	if t.NotificationSettingsScopeGroupChats != nil {
		return json.Marshal(t.NotificationSettingsScopeGroupChats)
	}
	if t.NotificationSettingsScopeChannelChats != nil {
		return json.Marshal(t.NotificationSettingsScopeChannelChats)
	}
	return nil, fmt.Errorf("no value set for NotificationSettingsScope")
}

// ChatTheme Describes a chat theme
type ChatTheme struct {
	ChatThemeEmoji *ChatThemeEmoji `json:"chatThemeEmoji,omitempty"`
	ChatThemeGift  *ChatThemeGift  `json:"chatThemeGift,omitempty"`
}

func (t *ChatTheme) Type() string {
	return ""
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

// StartLiveStoryResult Represents result of starting a live story
type StartLiveStoryResult struct {
	StartLiveStoryResultOk   *StartLiveStoryResultOk   `json:"startLiveStoryResultOk,omitempty"`
	StartLiveStoryResultFail *StartLiveStoryResultFail `json:"startLiveStoryResultFail,omitempty"`
}

func (t *StartLiveStoryResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "startLiveStoryResultOk":
		t.StartLiveStoryResultOk = new(StartLiveStoryResultOk)
		return json.Unmarshal(b, t.StartLiveStoryResultOk)
	case "startLiveStoryResultFail":
		t.StartLiveStoryResultFail = new(StartLiveStoryResultFail)
		return json.Unmarshal(b, t.StartLiveStoryResultFail)
	}
	return nil
}

func (t *StartLiveStoryResult) MarshalJSON() ([]byte, error) {
	if t.StartLiveStoryResultOk != nil {
		return json.Marshal(t.StartLiveStoryResultOk)
	}
	if t.StartLiveStoryResultFail != nil {
		return json.Marshal(t.StartLiveStoryResultFail)
	}
	return nil, fmt.Errorf("no value set for StartLiveStoryResult")
}

// ConnectionState Describes the current state of the connection to Telegram servers
type ConnectionState struct {
	ConnectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork `json:"connectionStateWaitingForNetwork,omitempty"`
	ConnectionStateConnectingToProxy *ConnectionStateConnectingToProxy `json:"connectionStateConnectingToProxy,omitempty"`
	ConnectionStateConnecting        *ConnectionStateConnecting        `json:"connectionStateConnecting,omitempty"`
	ConnectionStateUpdating          *ConnectionStateUpdating          `json:"connectionStateUpdating,omitempty"`
	ConnectionStateReady             *ConnectionStateReady             `json:"connectionStateReady,omitempty"`
}

func (t *ConnectionState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "connectionStateWaitingForNetwork":
		t.ConnectionStateWaitingForNetwork = new(ConnectionStateWaitingForNetwork)
		return json.Unmarshal(b, t.ConnectionStateWaitingForNetwork)
	case "connectionStateConnectingToProxy":
		t.ConnectionStateConnectingToProxy = new(ConnectionStateConnectingToProxy)
		return json.Unmarshal(b, t.ConnectionStateConnectingToProxy)
	case "connectionStateConnecting":
		t.ConnectionStateConnecting = new(ConnectionStateConnecting)
		return json.Unmarshal(b, t.ConnectionStateConnecting)
	case "connectionStateUpdating":
		t.ConnectionStateUpdating = new(ConnectionStateUpdating)
		return json.Unmarshal(b, t.ConnectionStateUpdating)
	case "connectionStateReady":
		t.ConnectionStateReady = new(ConnectionStateReady)
		return json.Unmarshal(b, t.ConnectionStateReady)
	}
	return nil
}

func (t *ConnectionState) MarshalJSON() ([]byte, error) {
	if t.ConnectionStateWaitingForNetwork != nil {
		return json.Marshal(t.ConnectionStateWaitingForNetwork)
	}
	if t.ConnectionStateConnectingToProxy != nil {
		return json.Marshal(t.ConnectionStateConnectingToProxy)
	}
	if t.ConnectionStateConnecting != nil {
		return json.Marshal(t.ConnectionStateConnecting)
	}
	if t.ConnectionStateUpdating != nil {
		return json.Marshal(t.ConnectionStateUpdating)
	}
	if t.ConnectionStateReady != nil {
		return json.Marshal(t.ConnectionStateReady)
	}
	return nil, fmt.Errorf("no value set for ConnectionState")
}

// InputCredentials Contains information about the payment method chosen by the user
type InputCredentials struct {
	InputCredentialsSaved     *InputCredentialsSaved     `json:"inputCredentialsSaved,omitempty"`
	InputCredentialsNew       *InputCredentialsNew       `json:"inputCredentialsNew,omitempty"`
	InputCredentialsApplePay  *InputCredentialsApplePay  `json:"inputCredentialsApplePay,omitempty"`
	InputCredentialsGooglePay *InputCredentialsGooglePay `json:"inputCredentialsGooglePay,omitempty"`
}

func (t *InputCredentials) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputCredentialsSaved":
		t.InputCredentialsSaved = new(InputCredentialsSaved)
		return json.Unmarshal(b, t.InputCredentialsSaved)
	case "inputCredentialsNew":
		t.InputCredentialsNew = new(InputCredentialsNew)
		return json.Unmarshal(b, t.InputCredentialsNew)
	case "inputCredentialsApplePay":
		t.InputCredentialsApplePay = new(InputCredentialsApplePay)
		return json.Unmarshal(b, t.InputCredentialsApplePay)
	case "inputCredentialsGooglePay":
		t.InputCredentialsGooglePay = new(InputCredentialsGooglePay)
		return json.Unmarshal(b, t.InputCredentialsGooglePay)
	}
	return nil
}

func (t *InputCredentials) MarshalJSON() ([]byte, error) {
	if t.InputCredentialsSaved != nil {
		return json.Marshal(t.InputCredentialsSaved)
	}
	if t.InputCredentialsNew != nil {
		return json.Marshal(t.InputCredentialsNew)
	}
	if t.InputCredentialsApplePay != nil {
		return json.Marshal(t.InputCredentialsApplePay)
	}
	if t.InputCredentialsGooglePay != nil {
		return json.Marshal(t.InputCredentialsGooglePay)
	}
	return nil, fmt.Errorf("no value set for InputCredentials")
}

// StoryAreaType Describes type of clickable area on a story media
type StoryAreaType struct {
	StoryAreaTypeLocation          *StoryAreaTypeLocation          `json:"storyAreaTypeLocation,omitempty"`
	StoryAreaTypeVenue             *StoryAreaTypeVenue             `json:"storyAreaTypeVenue,omitempty"`
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction `json:"storyAreaTypeSuggestedReaction,omitempty"`
	StoryAreaTypeMessage           *StoryAreaTypeMessage           `json:"storyAreaTypeMessage,omitempty"`
	StoryAreaTypeLink              *StoryAreaTypeLink              `json:"storyAreaTypeLink,omitempty"`
	StoryAreaTypeWeather           *StoryAreaTypeWeather           `json:"storyAreaTypeWeather,omitempty"`
	StoryAreaTypeUpgradedGift      *StoryAreaTypeUpgradedGift      `json:"storyAreaTypeUpgradedGift,omitempty"`
}

func (t *StoryAreaType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyAreaTypeLocation":
		t.StoryAreaTypeLocation = new(StoryAreaTypeLocation)
		return json.Unmarshal(b, t.StoryAreaTypeLocation)
	case "storyAreaTypeVenue":
		t.StoryAreaTypeVenue = new(StoryAreaTypeVenue)
		return json.Unmarshal(b, t.StoryAreaTypeVenue)
	case "storyAreaTypeSuggestedReaction":
		t.StoryAreaTypeSuggestedReaction = new(StoryAreaTypeSuggestedReaction)
		return json.Unmarshal(b, t.StoryAreaTypeSuggestedReaction)
	case "storyAreaTypeMessage":
		t.StoryAreaTypeMessage = new(StoryAreaTypeMessage)
		return json.Unmarshal(b, t.StoryAreaTypeMessage)
	case "storyAreaTypeLink":
		t.StoryAreaTypeLink = new(StoryAreaTypeLink)
		return json.Unmarshal(b, t.StoryAreaTypeLink)
	case "storyAreaTypeWeather":
		t.StoryAreaTypeWeather = new(StoryAreaTypeWeather)
		return json.Unmarshal(b, t.StoryAreaTypeWeather)
	case "storyAreaTypeUpgradedGift":
		t.StoryAreaTypeUpgradedGift = new(StoryAreaTypeUpgradedGift)
		return json.Unmarshal(b, t.StoryAreaTypeUpgradedGift)
	}
	return nil
}

func (t *StoryAreaType) MarshalJSON() ([]byte, error) {
	if t.StoryAreaTypeLocation != nil {
		return json.Marshal(t.StoryAreaTypeLocation)
	}
	if t.StoryAreaTypeVenue != nil {
		return json.Marshal(t.StoryAreaTypeVenue)
	}
	if t.StoryAreaTypeSuggestedReaction != nil {
		return json.Marshal(t.StoryAreaTypeSuggestedReaction)
	}
	if t.StoryAreaTypeMessage != nil {
		return json.Marshal(t.StoryAreaTypeMessage)
	}
	if t.StoryAreaTypeLink != nil {
		return json.Marshal(t.StoryAreaTypeLink)
	}
	if t.StoryAreaTypeWeather != nil {
		return json.Marshal(t.StoryAreaTypeWeather)
	}
	if t.StoryAreaTypeUpgradedGift != nil {
		return json.Marshal(t.StoryAreaTypeUpgradedGift)
	}
	return nil, fmt.Errorf("no value set for StoryAreaType")
}

// GiftPurchaseOfferState Describes state of a gift purchase offer
type GiftPurchaseOfferState struct {
	GiftPurchaseOfferStatePending  *GiftPurchaseOfferStatePending  `json:"giftPurchaseOfferStatePending,omitempty"`
	GiftPurchaseOfferStateAccepted *GiftPurchaseOfferStateAccepted `json:"giftPurchaseOfferStateAccepted,omitempty"`
	GiftPurchaseOfferStateRejected *GiftPurchaseOfferStateRejected `json:"giftPurchaseOfferStateRejected,omitempty"`
}

func (t *GiftPurchaseOfferState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "giftPurchaseOfferStatePending":
		t.GiftPurchaseOfferStatePending = new(GiftPurchaseOfferStatePending)
		return json.Unmarshal(b, t.GiftPurchaseOfferStatePending)
	case "giftPurchaseOfferStateAccepted":
		t.GiftPurchaseOfferStateAccepted = new(GiftPurchaseOfferStateAccepted)
		return json.Unmarshal(b, t.GiftPurchaseOfferStateAccepted)
	case "giftPurchaseOfferStateRejected":
		t.GiftPurchaseOfferStateRejected = new(GiftPurchaseOfferStateRejected)
		return json.Unmarshal(b, t.GiftPurchaseOfferStateRejected)
	}
	return nil
}

func (t *GiftPurchaseOfferState) MarshalJSON() ([]byte, error) {
	if t.GiftPurchaseOfferStatePending != nil {
		return json.Marshal(t.GiftPurchaseOfferStatePending)
	}
	if t.GiftPurchaseOfferStateAccepted != nil {
		return json.Marshal(t.GiftPurchaseOfferStateAccepted)
	}
	if t.GiftPurchaseOfferStateRejected != nil {
		return json.Marshal(t.GiftPurchaseOfferStateRejected)
	}
	return nil, fmt.Errorf("no value set for GiftPurchaseOfferState")
}

// SuggestedPostPrice Describes price of a suggested post
type SuggestedPostPrice struct {
	SuggestedPostPriceStar *SuggestedPostPriceStar `json:"suggestedPostPriceStar,omitempty"`
	SuggestedPostPriceTon  *SuggestedPostPriceTon  `json:"suggestedPostPriceTon,omitempty"`
}

func (t *SuggestedPostPrice) Type() string {
	return ""
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

// UpgradedGiftOrigin Describes origin from which the upgraded gift was obtained
type UpgradedGiftOrigin struct {
	UpgradedGiftOriginUpgrade        *UpgradedGiftOriginUpgrade        `json:"upgradedGiftOriginUpgrade,omitempty"`
	UpgradedGiftOriginTransfer       *UpgradedGiftOriginTransfer       `json:"upgradedGiftOriginTransfer,omitempty"`
	UpgradedGiftOriginResale         *UpgradedGiftOriginResale         `json:"upgradedGiftOriginResale,omitempty"`
	UpgradedGiftOriginBlockchain     *UpgradedGiftOriginBlockchain     `json:"upgradedGiftOriginBlockchain,omitempty"`
	UpgradedGiftOriginPrepaidUpgrade *UpgradedGiftOriginPrepaidUpgrade `json:"upgradedGiftOriginPrepaidUpgrade,omitempty"`
	UpgradedGiftOriginOffer          *UpgradedGiftOriginOffer          `json:"upgradedGiftOriginOffer,omitempty"`
}

func (t *UpgradedGiftOrigin) Type() string {
	return ""
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
	switch typeObj.Type {
	case "upgradedGiftOriginUpgrade":
		t.UpgradedGiftOriginUpgrade = new(UpgradedGiftOriginUpgrade)
		return json.Unmarshal(b, t.UpgradedGiftOriginUpgrade)
	case "upgradedGiftOriginTransfer":
		t.UpgradedGiftOriginTransfer = new(UpgradedGiftOriginTransfer)
		return json.Unmarshal(b, t.UpgradedGiftOriginTransfer)
	case "upgradedGiftOriginResale":
		t.UpgradedGiftOriginResale = new(UpgradedGiftOriginResale)
		return json.Unmarshal(b, t.UpgradedGiftOriginResale)
	case "upgradedGiftOriginBlockchain":
		t.UpgradedGiftOriginBlockchain = new(UpgradedGiftOriginBlockchain)
		return json.Unmarshal(b, t.UpgradedGiftOriginBlockchain)
	case "upgradedGiftOriginPrepaidUpgrade":
		t.UpgradedGiftOriginPrepaidUpgrade = new(UpgradedGiftOriginPrepaidUpgrade)
		return json.Unmarshal(b, t.UpgradedGiftOriginPrepaidUpgrade)
	case "upgradedGiftOriginOffer":
		t.UpgradedGiftOriginOffer = new(UpgradedGiftOriginOffer)
		return json.Unmarshal(b, t.UpgradedGiftOriginOffer)
	}
	return nil
}

func (t *UpgradedGiftOrigin) MarshalJSON() ([]byte, error) {
	if t.UpgradedGiftOriginUpgrade != nil {
		return json.Marshal(t.UpgradedGiftOriginUpgrade)
	}
	if t.UpgradedGiftOriginTransfer != nil {
		return json.Marshal(t.UpgradedGiftOriginTransfer)
	}
	if t.UpgradedGiftOriginResale != nil {
		return json.Marshal(t.UpgradedGiftOriginResale)
	}
	if t.UpgradedGiftOriginBlockchain != nil {
		return json.Marshal(t.UpgradedGiftOriginBlockchain)
	}
	if t.UpgradedGiftOriginPrepaidUpgrade != nil {
		return json.Marshal(t.UpgradedGiftOriginPrepaidUpgrade)
	}
	if t.UpgradedGiftOriginOffer != nil {
		return json.Marshal(t.UpgradedGiftOriginOffer)
	}
	return nil, fmt.Errorf("no value set for UpgradedGiftOrigin")
}

// SupergroupMembersFilter Specifies the kind of chat members to return in getSupergroupMembers
type SupergroupMembersFilter struct {
	SupergroupMembersFilterRecent         *SupergroupMembersFilterRecent         `json:"supergroupMembersFilterRecent,omitempty"`
	SupergroupMembersFilterContacts       *SupergroupMembersFilterContacts       `json:"supergroupMembersFilterContacts,omitempty"`
	SupergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators `json:"supergroupMembersFilterAdministrators,omitempty"`
	SupergroupMembersFilterSearch         *SupergroupMembersFilterSearch         `json:"supergroupMembersFilterSearch,omitempty"`
	SupergroupMembersFilterRestricted     *SupergroupMembersFilterRestricted     `json:"supergroupMembersFilterRestricted,omitempty"`
	SupergroupMembersFilterBanned         *SupergroupMembersFilterBanned         `json:"supergroupMembersFilterBanned,omitempty"`
	SupergroupMembersFilterMention        *SupergroupMembersFilterMention        `json:"supergroupMembersFilterMention,omitempty"`
	SupergroupMembersFilterBots           *SupergroupMembersFilterBots           `json:"supergroupMembersFilterBots,omitempty"`
}

func (t *SupergroupMembersFilter) Type() string {
	return ""
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
	switch typeObj.Type {
	case "supergroupMembersFilterRecent":
		t.SupergroupMembersFilterRecent = new(SupergroupMembersFilterRecent)
		return json.Unmarshal(b, t.SupergroupMembersFilterRecent)
	case "supergroupMembersFilterContacts":
		t.SupergroupMembersFilterContacts = new(SupergroupMembersFilterContacts)
		return json.Unmarshal(b, t.SupergroupMembersFilterContacts)
	case "supergroupMembersFilterAdministrators":
		t.SupergroupMembersFilterAdministrators = new(SupergroupMembersFilterAdministrators)
		return json.Unmarshal(b, t.SupergroupMembersFilterAdministrators)
	case "supergroupMembersFilterSearch":
		t.SupergroupMembersFilterSearch = new(SupergroupMembersFilterSearch)
		return json.Unmarshal(b, t.SupergroupMembersFilterSearch)
	case "supergroupMembersFilterRestricted":
		t.SupergroupMembersFilterRestricted = new(SupergroupMembersFilterRestricted)
		return json.Unmarshal(b, t.SupergroupMembersFilterRestricted)
	case "supergroupMembersFilterBanned":
		t.SupergroupMembersFilterBanned = new(SupergroupMembersFilterBanned)
		return json.Unmarshal(b, t.SupergroupMembersFilterBanned)
	case "supergroupMembersFilterMention":
		t.SupergroupMembersFilterMention = new(SupergroupMembersFilterMention)
		return json.Unmarshal(b, t.SupergroupMembersFilterMention)
	case "supergroupMembersFilterBots":
		t.SupergroupMembersFilterBots = new(SupergroupMembersFilterBots)
		return json.Unmarshal(b, t.SupergroupMembersFilterBots)
	}
	return nil
}

func (t *SupergroupMembersFilter) MarshalJSON() ([]byte, error) {
	if t.SupergroupMembersFilterRecent != nil {
		return json.Marshal(t.SupergroupMembersFilterRecent)
	}
	if t.SupergroupMembersFilterContacts != nil {
		return json.Marshal(t.SupergroupMembersFilterContacts)
	}
	if t.SupergroupMembersFilterAdministrators != nil {
		return json.Marshal(t.SupergroupMembersFilterAdministrators)
	}
	if t.SupergroupMembersFilterSearch != nil {
		return json.Marshal(t.SupergroupMembersFilterSearch)
	}
	if t.SupergroupMembersFilterRestricted != nil {
		return json.Marshal(t.SupergroupMembersFilterRestricted)
	}
	if t.SupergroupMembersFilterBanned != nil {
		return json.Marshal(t.SupergroupMembersFilterBanned)
	}
	if t.SupergroupMembersFilterMention != nil {
		return json.Marshal(t.SupergroupMembersFilterMention)
	}
	if t.SupergroupMembersFilterBots != nil {
		return json.Marshal(t.SupergroupMembersFilterBots)
	}
	return nil, fmt.Errorf("no value set for SupergroupMembersFilter")
}

// PageBlockVerticalAlignment Describes a Vertical alignment of a table cell content
type PageBlockVerticalAlignment struct {
	PageBlockVerticalAlignmentTop    *PageBlockVerticalAlignmentTop    `json:"pageBlockVerticalAlignmentTop,omitempty"`
	PageBlockVerticalAlignmentMiddle *PageBlockVerticalAlignmentMiddle `json:"pageBlockVerticalAlignmentMiddle,omitempty"`
	PageBlockVerticalAlignmentBottom *PageBlockVerticalAlignmentBottom `json:"pageBlockVerticalAlignmentBottom,omitempty"`
}

func (t *PageBlockVerticalAlignment) Type() string {
	return ""
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
	switch typeObj.Type {
	case "pageBlockVerticalAlignmentTop":
		t.PageBlockVerticalAlignmentTop = new(PageBlockVerticalAlignmentTop)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentTop)
	case "pageBlockVerticalAlignmentMiddle":
		t.PageBlockVerticalAlignmentMiddle = new(PageBlockVerticalAlignmentMiddle)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentMiddle)
	case "pageBlockVerticalAlignmentBottom":
		t.PageBlockVerticalAlignmentBottom = new(PageBlockVerticalAlignmentBottom)
		return json.Unmarshal(b, t.PageBlockVerticalAlignmentBottom)
	}
	return nil
}

func (t *PageBlockVerticalAlignment) MarshalJSON() ([]byte, error) {
	if t.PageBlockVerticalAlignmentTop != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentTop)
	}
	if t.PageBlockVerticalAlignmentMiddle != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentMiddle)
	}
	if t.PageBlockVerticalAlignmentBottom != nil {
		return json.Marshal(t.PageBlockVerticalAlignmentBottom)
	}
	return nil, fmt.Errorf("no value set for PageBlockVerticalAlignment")
}

// PassportElementErrorSource Contains the description of an error in a Telegram Passport element
type PassportElementErrorSource struct {
	PassportElementErrorSourceUnspecified      *PassportElementErrorSourceUnspecified      `json:"passportElementErrorSourceUnspecified,omitempty"`
	PassportElementErrorSourceDataField        *PassportElementErrorSourceDataField        `json:"passportElementErrorSourceDataField,omitempty"`
	PassportElementErrorSourceFrontSide        *PassportElementErrorSourceFrontSide        `json:"passportElementErrorSourceFrontSide,omitempty"`
	PassportElementErrorSourceReverseSide      *PassportElementErrorSourceReverseSide      `json:"passportElementErrorSourceReverseSide,omitempty"`
	PassportElementErrorSourceSelfie           *PassportElementErrorSourceSelfie           `json:"passportElementErrorSourceSelfie,omitempty"`
	PassportElementErrorSourceTranslationFile  *PassportElementErrorSourceTranslationFile  `json:"passportElementErrorSourceTranslationFile,omitempty"`
	PassportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles `json:"passportElementErrorSourceTranslationFiles,omitempty"`
	PassportElementErrorSourceFile             *PassportElementErrorSourceFile             `json:"passportElementErrorSourceFile,omitempty"`
	PassportElementErrorSourceFiles            *PassportElementErrorSourceFiles            `json:"passportElementErrorSourceFiles,omitempty"`
}

func (t *PassportElementErrorSource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "passportElementErrorSourceUnspecified":
		t.PassportElementErrorSourceUnspecified = new(PassportElementErrorSourceUnspecified)
		return json.Unmarshal(b, t.PassportElementErrorSourceUnspecified)
	case "passportElementErrorSourceDataField":
		t.PassportElementErrorSourceDataField = new(PassportElementErrorSourceDataField)
		return json.Unmarshal(b, t.PassportElementErrorSourceDataField)
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
	case "passportElementErrorSourceFile":
		t.PassportElementErrorSourceFile = new(PassportElementErrorSourceFile)
		return json.Unmarshal(b, t.PassportElementErrorSourceFile)
	case "passportElementErrorSourceFiles":
		t.PassportElementErrorSourceFiles = new(PassportElementErrorSourceFiles)
		return json.Unmarshal(b, t.PassportElementErrorSourceFiles)
	}
	return nil
}

func (t *PassportElementErrorSource) MarshalJSON() ([]byte, error) {
	if t.PassportElementErrorSourceUnspecified != nil {
		return json.Marshal(t.PassportElementErrorSourceUnspecified)
	}
	if t.PassportElementErrorSourceDataField != nil {
		return json.Marshal(t.PassportElementErrorSourceDataField)
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
	if t.PassportElementErrorSourceFile != nil {
		return json.Marshal(t.PassportElementErrorSourceFile)
	}
	if t.PassportElementErrorSourceFiles != nil {
		return json.Marshal(t.PassportElementErrorSourceFiles)
	}
	return nil, fmt.Errorf("no value set for PassportElementErrorSource")
}

// InputStoryContent The content of a story to post
type InputStoryContent struct {
	InputStoryContentPhoto *InputStoryContentPhoto `json:"inputStoryContentPhoto,omitempty"`
	InputStoryContentVideo *InputStoryContentVideo `json:"inputStoryContentVideo,omitempty"`
}

func (t *InputStoryContent) Type() string {
	return ""
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

// CallProblem Describes the exact type of problem with a call
type CallProblem struct {
	CallProblemEcho            *CallProblemEcho            `json:"callProblemEcho,omitempty"`
	CallProblemNoise           *CallProblemNoise           `json:"callProblemNoise,omitempty"`
	CallProblemInterruptions   *CallProblemInterruptions   `json:"callProblemInterruptions,omitempty"`
	CallProblemDistortedSpeech *CallProblemDistortedSpeech `json:"callProblemDistortedSpeech,omitempty"`
	CallProblemSilentLocal     *CallProblemSilentLocal     `json:"callProblemSilentLocal,omitempty"`
	CallProblemSilentRemote    *CallProblemSilentRemote    `json:"callProblemSilentRemote,omitempty"`
	CallProblemDropped         *CallProblemDropped         `json:"callProblemDropped,omitempty"`
	CallProblemDistortedVideo  *CallProblemDistortedVideo  `json:"callProblemDistortedVideo,omitempty"`
	CallProblemPixelatedVideo  *CallProblemPixelatedVideo  `json:"callProblemPixelatedVideo,omitempty"`
}

func (t *CallProblem) Type() string {
	return ""
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
	switch typeObj.Type {
	case "callProblemEcho":
		t.CallProblemEcho = new(CallProblemEcho)
		return json.Unmarshal(b, t.CallProblemEcho)
	case "callProblemNoise":
		t.CallProblemNoise = new(CallProblemNoise)
		return json.Unmarshal(b, t.CallProblemNoise)
	case "callProblemInterruptions":
		t.CallProblemInterruptions = new(CallProblemInterruptions)
		return json.Unmarshal(b, t.CallProblemInterruptions)
	case "callProblemDistortedSpeech":
		t.CallProblemDistortedSpeech = new(CallProblemDistortedSpeech)
		return json.Unmarshal(b, t.CallProblemDistortedSpeech)
	case "callProblemSilentLocal":
		t.CallProblemSilentLocal = new(CallProblemSilentLocal)
		return json.Unmarshal(b, t.CallProblemSilentLocal)
	case "callProblemSilentRemote":
		t.CallProblemSilentRemote = new(CallProblemSilentRemote)
		return json.Unmarshal(b, t.CallProblemSilentRemote)
	case "callProblemDropped":
		t.CallProblemDropped = new(CallProblemDropped)
		return json.Unmarshal(b, t.CallProblemDropped)
	case "callProblemDistortedVideo":
		t.CallProblemDistortedVideo = new(CallProblemDistortedVideo)
		return json.Unmarshal(b, t.CallProblemDistortedVideo)
	case "callProblemPixelatedVideo":
		t.CallProblemPixelatedVideo = new(CallProblemPixelatedVideo)
		return json.Unmarshal(b, t.CallProblemPixelatedVideo)
	}
	return nil
}

func (t *CallProblem) MarshalJSON() ([]byte, error) {
	if t.CallProblemEcho != nil {
		return json.Marshal(t.CallProblemEcho)
	}
	if t.CallProblemNoise != nil {
		return json.Marshal(t.CallProblemNoise)
	}
	if t.CallProblemInterruptions != nil {
		return json.Marshal(t.CallProblemInterruptions)
	}
	if t.CallProblemDistortedSpeech != nil {
		return json.Marshal(t.CallProblemDistortedSpeech)
	}
	if t.CallProblemSilentLocal != nil {
		return json.Marshal(t.CallProblemSilentLocal)
	}
	if t.CallProblemSilentRemote != nil {
		return json.Marshal(t.CallProblemSilentRemote)
	}
	if t.CallProblemDropped != nil {
		return json.Marshal(t.CallProblemDropped)
	}
	if t.CallProblemDistortedVideo != nil {
		return json.Marshal(t.CallProblemDistortedVideo)
	}
	if t.CallProblemPixelatedVideo != nil {
		return json.Marshal(t.CallProblemPixelatedVideo)
	}
	return nil, fmt.Errorf("no value set for CallProblem")
}

// StickerType Describes type of sticker
type StickerType struct {
	StickerTypeRegular     *StickerTypeRegular     `json:"stickerTypeRegular,omitempty"`
	StickerTypeMask        *StickerTypeMask        `json:"stickerTypeMask,omitempty"`
	StickerTypeCustomEmoji *StickerTypeCustomEmoji `json:"stickerTypeCustomEmoji,omitempty"`
}

func (t *StickerType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "stickerTypeRegular":
		t.StickerTypeRegular = new(StickerTypeRegular)
		return json.Unmarshal(b, t.StickerTypeRegular)
	case "stickerTypeMask":
		t.StickerTypeMask = new(StickerTypeMask)
		return json.Unmarshal(b, t.StickerTypeMask)
	case "stickerTypeCustomEmoji":
		t.StickerTypeCustomEmoji = new(StickerTypeCustomEmoji)
		return json.Unmarshal(b, t.StickerTypeCustomEmoji)
	}
	return nil
}

func (t *StickerType) MarshalJSON() ([]byte, error) {
	if t.StickerTypeRegular != nil {
		return json.Marshal(t.StickerTypeRegular)
	}
	if t.StickerTypeMask != nil {
		return json.Marshal(t.StickerTypeMask)
	}
	if t.StickerTypeCustomEmoji != nil {
		return json.Marshal(t.StickerTypeCustomEmoji)
	}
	return nil, fmt.Errorf("no value set for StickerType")
}

// ReplyMarkup Contains a description of a custom keyboard and actions that can be done with it to quickly reply to bots
type ReplyMarkup struct {
	ReplyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard `json:"replyMarkupRemoveKeyboard,omitempty"`
	ReplyMarkupForceReply     *ReplyMarkupForceReply     `json:"replyMarkupForceReply,omitempty"`
	ReplyMarkupShowKeyboard   *ReplyMarkupShowKeyboard   `json:"replyMarkupShowKeyboard,omitempty"`
	ReplyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard `json:"replyMarkupInlineKeyboard,omitempty"`
}

func (t *ReplyMarkup) Type() string {
	return ""
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
	switch typeObj.Type {
	case "replyMarkupRemoveKeyboard":
		t.ReplyMarkupRemoveKeyboard = new(ReplyMarkupRemoveKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupRemoveKeyboard)
	case "replyMarkupForceReply":
		t.ReplyMarkupForceReply = new(ReplyMarkupForceReply)
		return json.Unmarshal(b, t.ReplyMarkupForceReply)
	case "replyMarkupShowKeyboard":
		t.ReplyMarkupShowKeyboard = new(ReplyMarkupShowKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupShowKeyboard)
	case "replyMarkupInlineKeyboard":
		t.ReplyMarkupInlineKeyboard = new(ReplyMarkupInlineKeyboard)
		return json.Unmarshal(b, t.ReplyMarkupInlineKeyboard)
	}
	return nil
}

func (t *ReplyMarkup) MarshalJSON() ([]byte, error) {
	if t.ReplyMarkupRemoveKeyboard != nil {
		return json.Marshal(t.ReplyMarkupRemoveKeyboard)
	}
	if t.ReplyMarkupForceReply != nil {
		return json.Marshal(t.ReplyMarkupForceReply)
	}
	if t.ReplyMarkupShowKeyboard != nil {
		return json.Marshal(t.ReplyMarkupShowKeyboard)
	}
	if t.ReplyMarkupInlineKeyboard != nil {
		return json.Marshal(t.ReplyMarkupInlineKeyboard)
	}
	return nil, fmt.Errorf("no value set for ReplyMarkup")
}

// TargetChat Describes the target chat to be opened
type TargetChat struct {
	TargetChatCurrent      *TargetChatCurrent      `json:"targetChatCurrent,omitempty"`
	TargetChatChosen       *TargetChatChosen       `json:"targetChatChosen,omitempty"`
	TargetChatInternalLink *TargetChatInternalLink `json:"targetChatInternalLink,omitempty"`
}

func (t *TargetChat) Type() string {
	return ""
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
	switch typeObj.Type {
	case "targetChatCurrent":
		t.TargetChatCurrent = new(TargetChatCurrent)
		return json.Unmarshal(b, t.TargetChatCurrent)
	case "targetChatChosen":
		t.TargetChatChosen = new(TargetChatChosen)
		return json.Unmarshal(b, t.TargetChatChosen)
	case "targetChatInternalLink":
		t.TargetChatInternalLink = new(TargetChatInternalLink)
		return json.Unmarshal(b, t.TargetChatInternalLink)
	}
	return nil
}

func (t *TargetChat) MarshalJSON() ([]byte, error) {
	if t.TargetChatCurrent != nil {
		return json.Marshal(t.TargetChatCurrent)
	}
	if t.TargetChatChosen != nil {
		return json.Marshal(t.TargetChatChosen)
	}
	if t.TargetChatInternalLink != nil {
		return json.Marshal(t.TargetChatInternalLink)
	}
	return nil, fmt.Errorf("no value set for TargetChat")
}

// CheckStickerSetNameResult Represents result of checking whether a name can be used for a new sticker set
type CheckStickerSetNameResult struct {
	CheckStickerSetNameResultOk           *CheckStickerSetNameResultOk           `json:"checkStickerSetNameResultOk,omitempty"`
	CheckStickerSetNameResultNameInvalid  *CheckStickerSetNameResultNameInvalid  `json:"checkStickerSetNameResultNameInvalid,omitempty"`
	CheckStickerSetNameResultNameOccupied *CheckStickerSetNameResultNameOccupied `json:"checkStickerSetNameResultNameOccupied,omitempty"`
}

func (t *CheckStickerSetNameResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "checkStickerSetNameResultOk":
		t.CheckStickerSetNameResultOk = new(CheckStickerSetNameResultOk)
		return json.Unmarshal(b, t.CheckStickerSetNameResultOk)
	case "checkStickerSetNameResultNameInvalid":
		t.CheckStickerSetNameResultNameInvalid = new(CheckStickerSetNameResultNameInvalid)
		return json.Unmarshal(b, t.CheckStickerSetNameResultNameInvalid)
	case "checkStickerSetNameResultNameOccupied":
		t.CheckStickerSetNameResultNameOccupied = new(CheckStickerSetNameResultNameOccupied)
		return json.Unmarshal(b, t.CheckStickerSetNameResultNameOccupied)
	}
	return nil
}

func (t *CheckStickerSetNameResult) MarshalJSON() ([]byte, error) {
	if t.CheckStickerSetNameResultOk != nil {
		return json.Marshal(t.CheckStickerSetNameResultOk)
	}
	if t.CheckStickerSetNameResultNameInvalid != nil {
		return json.Marshal(t.CheckStickerSetNameResultNameInvalid)
	}
	if t.CheckStickerSetNameResultNameOccupied != nil {
		return json.Marshal(t.CheckStickerSetNameResultNameOccupied)
	}
	return nil, fmt.Errorf("no value set for CheckStickerSetNameResult")
}

// ReportStoryResult Describes result of story report
type ReportStoryResult struct {
	ReportStoryResultOk             *ReportStoryResultOk             `json:"reportStoryResultOk,omitempty"`
	ReportStoryResultOptionRequired *ReportStoryResultOptionRequired `json:"reportStoryResultOptionRequired,omitempty"`
	ReportStoryResultTextRequired   *ReportStoryResultTextRequired   `json:"reportStoryResultTextRequired,omitempty"`
}

func (t *ReportStoryResult) Type() string {
	return ""
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

// TransactionDirection Describes direction of transactions in a transaction list
type TransactionDirection struct {
	TransactionDirectionIncoming *TransactionDirectionIncoming `json:"transactionDirectionIncoming,omitempty"`
	TransactionDirectionOutgoing *TransactionDirectionOutgoing `json:"transactionDirectionOutgoing,omitempty"`
}

func (t *TransactionDirection) Type() string {
	return ""
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

// ReactionType Describes type of message reaction
type ReactionType struct {
	ReactionTypeEmoji       *ReactionTypeEmoji       `json:"reactionTypeEmoji,omitempty"`
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji `json:"reactionTypeCustomEmoji,omitempty"`
	ReactionTypePaid        *ReactionTypePaid        `json:"reactionTypePaid,omitempty"`
}

func (t *ReactionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "reactionTypeEmoji":
		t.ReactionTypeEmoji = new(ReactionTypeEmoji)
		return json.Unmarshal(b, t.ReactionTypeEmoji)
	case "reactionTypeCustomEmoji":
		t.ReactionTypeCustomEmoji = new(ReactionTypeCustomEmoji)
		return json.Unmarshal(b, t.ReactionTypeCustomEmoji)
	case "reactionTypePaid":
		t.ReactionTypePaid = new(ReactionTypePaid)
		return json.Unmarshal(b, t.ReactionTypePaid)
	}
	return nil
}

func (t *ReactionType) MarshalJSON() ([]byte, error) {
	if t.ReactionTypeEmoji != nil {
		return json.Marshal(t.ReactionTypeEmoji)
	}
	if t.ReactionTypeCustomEmoji != nil {
		return json.Marshal(t.ReactionTypeCustomEmoji)
	}
	if t.ReactionTypePaid != nil {
		return json.Marshal(t.ReactionTypePaid)
	}
	return nil, fmt.Errorf("no value set for ReactionType")
}

// PaidReactionType Describes type of paid message reaction
type PaidReactionType struct {
	PaidReactionTypeRegular   *PaidReactionTypeRegular   `json:"paidReactionTypeRegular,omitempty"`
	PaidReactionTypeAnonymous *PaidReactionTypeAnonymous `json:"paidReactionTypeAnonymous,omitempty"`
	PaidReactionTypeChat      *PaidReactionTypeChat      `json:"paidReactionTypeChat,omitempty"`
}

func (t *PaidReactionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "paidReactionTypeRegular":
		t.PaidReactionTypeRegular = new(PaidReactionTypeRegular)
		return json.Unmarshal(b, t.PaidReactionTypeRegular)
	case "paidReactionTypeAnonymous":
		t.PaidReactionTypeAnonymous = new(PaidReactionTypeAnonymous)
		return json.Unmarshal(b, t.PaidReactionTypeAnonymous)
	case "paidReactionTypeChat":
		t.PaidReactionTypeChat = new(PaidReactionTypeChat)
		return json.Unmarshal(b, t.PaidReactionTypeChat)
	}
	return nil
}

func (t *PaidReactionType) MarshalJSON() ([]byte, error) {
	if t.PaidReactionTypeRegular != nil {
		return json.Marshal(t.PaidReactionTypeRegular)
	}
	if t.PaidReactionTypeAnonymous != nil {
		return json.Marshal(t.PaidReactionTypeAnonymous)
	}
	if t.PaidReactionTypeChat != nil {
		return json.Marshal(t.PaidReactionTypeChat)
	}
	return nil, fmt.Errorf("no value set for PaidReactionType")
}

// WebAppOpenMode Describes mode in which a Web App is opened
type WebAppOpenMode struct {
	WebAppOpenModeCompact    *WebAppOpenModeCompact    `json:"webAppOpenModeCompact,omitempty"`
	WebAppOpenModeFullSize   *WebAppOpenModeFullSize   `json:"webAppOpenModeFullSize,omitempty"`
	WebAppOpenModeFullScreen *WebAppOpenModeFullScreen `json:"webAppOpenModeFullScreen,omitempty"`
}

func (t *WebAppOpenMode) Type() string {
	return ""
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
	switch typeObj.Type {
	case "webAppOpenModeCompact":
		t.WebAppOpenModeCompact = new(WebAppOpenModeCompact)
		return json.Unmarshal(b, t.WebAppOpenModeCompact)
	case "webAppOpenModeFullSize":
		t.WebAppOpenModeFullSize = new(WebAppOpenModeFullSize)
		return json.Unmarshal(b, t.WebAppOpenModeFullSize)
	case "webAppOpenModeFullScreen":
		t.WebAppOpenModeFullScreen = new(WebAppOpenModeFullScreen)
		return json.Unmarshal(b, t.WebAppOpenModeFullScreen)
	}
	return nil
}

func (t *WebAppOpenMode) MarshalJSON() ([]byte, error) {
	if t.WebAppOpenModeCompact != nil {
		return json.Marshal(t.WebAppOpenModeCompact)
	}
	if t.WebAppOpenModeFullSize != nil {
		return json.Marshal(t.WebAppOpenModeFullSize)
	}
	if t.WebAppOpenModeFullScreen != nil {
		return json.Marshal(t.WebAppOpenModeFullScreen)
	}
	return nil, fmt.Errorf("no value set for WebAppOpenMode")
}

// PaymentReceiptType Describes type of successful payment
type PaymentReceiptType struct {
	PaymentReceiptTypeRegular *PaymentReceiptTypeRegular `json:"paymentReceiptTypeRegular,omitempty"`
	PaymentReceiptTypeStars   *PaymentReceiptTypeStars   `json:"paymentReceiptTypeStars,omitempty"`
}

func (t *PaymentReceiptType) Type() string {
	return ""
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

// InputBackground Contains information about background to set
type InputBackground struct {
	InputBackgroundLocal    *InputBackgroundLocal    `json:"inputBackgroundLocal,omitempty"`
	InputBackgroundRemote   *InputBackgroundRemote   `json:"inputBackgroundRemote,omitempty"`
	InputBackgroundPrevious *InputBackgroundPrevious `json:"inputBackgroundPrevious,omitempty"`
}

func (t *InputBackground) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputBackgroundLocal":
		t.InputBackgroundLocal = new(InputBackgroundLocal)
		return json.Unmarshal(b, t.InputBackgroundLocal)
	case "inputBackgroundRemote":
		t.InputBackgroundRemote = new(InputBackgroundRemote)
		return json.Unmarshal(b, t.InputBackgroundRemote)
	case "inputBackgroundPrevious":
		t.InputBackgroundPrevious = new(InputBackgroundPrevious)
		return json.Unmarshal(b, t.InputBackgroundPrevious)
	}
	return nil
}

func (t *InputBackground) MarshalJSON() ([]byte, error) {
	if t.InputBackgroundLocal != nil {
		return json.Marshal(t.InputBackgroundLocal)
	}
	if t.InputBackgroundRemote != nil {
		return json.Marshal(t.InputBackgroundRemote)
	}
	if t.InputBackgroundPrevious != nil {
		return json.Marshal(t.InputBackgroundPrevious)
	}
	return nil, fmt.Errorf("no value set for InputBackground")
}

// MessageSender Contains information about the sender of a message
type MessageSender struct {
	MessageSenderUser *MessageSenderUser `json:"messageSenderUser,omitempty"`
	MessageSenderChat *MessageSenderChat `json:"messageSenderChat,omitempty"`
}

func (t *MessageSender) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageSenderUser":
		t.MessageSenderUser = new(MessageSenderUser)
		return json.Unmarshal(b, t.MessageSenderUser)
	case "messageSenderChat":
		t.MessageSenderChat = new(MessageSenderChat)
		return json.Unmarshal(b, t.MessageSenderChat)
	}
	return nil
}

func (t *MessageSender) MarshalJSON() ([]byte, error) {
	if t.MessageSenderUser != nil {
		return json.Marshal(t.MessageSenderUser)
	}
	if t.MessageSenderChat != nil {
		return json.Marshal(t.MessageSenderChat)
	}
	return nil, fmt.Errorf("no value set for MessageSender")
}

// UpgradedGiftAttributeId Contains identifier of an upgraded gift attribute to search for
type UpgradedGiftAttributeId struct {
	UpgradedGiftAttributeIdModel    *UpgradedGiftAttributeIdModel    `json:"upgradedGiftAttributeIdModel,omitempty"`
	UpgradedGiftAttributeIdSymbol   *UpgradedGiftAttributeIdSymbol   `json:"upgradedGiftAttributeIdSymbol,omitempty"`
	UpgradedGiftAttributeIdBackdrop *UpgradedGiftAttributeIdBackdrop `json:"upgradedGiftAttributeIdBackdrop,omitempty"`
}

func (t *UpgradedGiftAttributeId) Type() string {
	return ""
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
	switch typeObj.Type {
	case "upgradedGiftAttributeIdModel":
		t.UpgradedGiftAttributeIdModel = new(UpgradedGiftAttributeIdModel)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdModel)
	case "upgradedGiftAttributeIdSymbol":
		t.UpgradedGiftAttributeIdSymbol = new(UpgradedGiftAttributeIdSymbol)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdSymbol)
	case "upgradedGiftAttributeIdBackdrop":
		t.UpgradedGiftAttributeIdBackdrop = new(UpgradedGiftAttributeIdBackdrop)
		return json.Unmarshal(b, t.UpgradedGiftAttributeIdBackdrop)
	}
	return nil
}

func (t *UpgradedGiftAttributeId) MarshalJSON() ([]byte, error) {
	if t.UpgradedGiftAttributeIdModel != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdModel)
	}
	if t.UpgradedGiftAttributeIdSymbol != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdSymbol)
	}
	if t.UpgradedGiftAttributeIdBackdrop != nil {
		return json.Marshal(t.UpgradedGiftAttributeIdBackdrop)
	}
	return nil, fmt.Errorf("no value set for UpgradedGiftAttributeId")
}

// GiveawayParticipantStatus Contains information about status of a user in a giveaway
type GiveawayParticipantStatus struct {
	GiveawayParticipantStatusEligible          *GiveawayParticipantStatusEligible          `json:"giveawayParticipantStatusEligible,omitempty"`
	GiveawayParticipantStatusParticipating     *GiveawayParticipantStatusParticipating     `json:"giveawayParticipantStatusParticipating,omitempty"`
	GiveawayParticipantStatusAlreadyWasMember  *GiveawayParticipantStatusAlreadyWasMember  `json:"giveawayParticipantStatusAlreadyWasMember,omitempty"`
	GiveawayParticipantStatusAdministrator     *GiveawayParticipantStatusAdministrator     `json:"giveawayParticipantStatusAdministrator,omitempty"`
	GiveawayParticipantStatusDisallowedCountry *GiveawayParticipantStatusDisallowedCountry `json:"giveawayParticipantStatusDisallowedCountry,omitempty"`
}

func (t *GiveawayParticipantStatus) Type() string {
	return ""
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
	switch typeObj.Type {
	case "giveawayParticipantStatusEligible":
		t.GiveawayParticipantStatusEligible = new(GiveawayParticipantStatusEligible)
		return json.Unmarshal(b, t.GiveawayParticipantStatusEligible)
	case "giveawayParticipantStatusParticipating":
		t.GiveawayParticipantStatusParticipating = new(GiveawayParticipantStatusParticipating)
		return json.Unmarshal(b, t.GiveawayParticipantStatusParticipating)
	case "giveawayParticipantStatusAlreadyWasMember":
		t.GiveawayParticipantStatusAlreadyWasMember = new(GiveawayParticipantStatusAlreadyWasMember)
		return json.Unmarshal(b, t.GiveawayParticipantStatusAlreadyWasMember)
	case "giveawayParticipantStatusAdministrator":
		t.GiveawayParticipantStatusAdministrator = new(GiveawayParticipantStatusAdministrator)
		return json.Unmarshal(b, t.GiveawayParticipantStatusAdministrator)
	case "giveawayParticipantStatusDisallowedCountry":
		t.GiveawayParticipantStatusDisallowedCountry = new(GiveawayParticipantStatusDisallowedCountry)
		return json.Unmarshal(b, t.GiveawayParticipantStatusDisallowedCountry)
	}
	return nil
}

func (t *GiveawayParticipantStatus) MarshalJSON() ([]byte, error) {
	if t.GiveawayParticipantStatusEligible != nil {
		return json.Marshal(t.GiveawayParticipantStatusEligible)
	}
	if t.GiveawayParticipantStatusParticipating != nil {
		return json.Marshal(t.GiveawayParticipantStatusParticipating)
	}
	if t.GiveawayParticipantStatusAlreadyWasMember != nil {
		return json.Marshal(t.GiveawayParticipantStatusAlreadyWasMember)
	}
	if t.GiveawayParticipantStatusAdministrator != nil {
		return json.Marshal(t.GiveawayParticipantStatusAdministrator)
	}
	if t.GiveawayParticipantStatusDisallowedCountry != nil {
		return json.Marshal(t.GiveawayParticipantStatusDisallowedCountry)
	}
	return nil, fmt.Errorf("no value set for GiveawayParticipantStatus")
}

// SavedMessagesTopicType Describes type of Saved Messages topic
type SavedMessagesTopicType struct {
	SavedMessagesTopicTypeMyNotes       *SavedMessagesTopicTypeMyNotes       `json:"savedMessagesTopicTypeMyNotes,omitempty"`
	SavedMessagesTopicTypeAuthorHidden  *SavedMessagesTopicTypeAuthorHidden  `json:"savedMessagesTopicTypeAuthorHidden,omitempty"`
	SavedMessagesTopicTypeSavedFromChat *SavedMessagesTopicTypeSavedFromChat `json:"savedMessagesTopicTypeSavedFromChat,omitempty"`
}

func (t *SavedMessagesTopicType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "savedMessagesTopicTypeMyNotes":
		t.SavedMessagesTopicTypeMyNotes = new(SavedMessagesTopicTypeMyNotes)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeMyNotes)
	case "savedMessagesTopicTypeAuthorHidden":
		t.SavedMessagesTopicTypeAuthorHidden = new(SavedMessagesTopicTypeAuthorHidden)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeAuthorHidden)
	case "savedMessagesTopicTypeSavedFromChat":
		t.SavedMessagesTopicTypeSavedFromChat = new(SavedMessagesTopicTypeSavedFromChat)
		return json.Unmarshal(b, t.SavedMessagesTopicTypeSavedFromChat)
	}
	return nil
}

func (t *SavedMessagesTopicType) MarshalJSON() ([]byte, error) {
	if t.SavedMessagesTopicTypeMyNotes != nil {
		return json.Marshal(t.SavedMessagesTopicTypeMyNotes)
	}
	if t.SavedMessagesTopicTypeAuthorHidden != nil {
		return json.Marshal(t.SavedMessagesTopicTypeAuthorHidden)
	}
	if t.SavedMessagesTopicTypeSavedFromChat != nil {
		return json.Marshal(t.SavedMessagesTopicTypeSavedFromChat)
	}
	return nil, fmt.Errorf("no value set for SavedMessagesTopicType")
}

// ResetPasswordResult Represents result of 2-step verification password reset
type ResetPasswordResult struct {
	ResetPasswordResultOk       *ResetPasswordResultOk       `json:"resetPasswordResultOk,omitempty"`
	ResetPasswordResultPending  *ResetPasswordResultPending  `json:"resetPasswordResultPending,omitempty"`
	ResetPasswordResultDeclined *ResetPasswordResultDeclined `json:"resetPasswordResultDeclined,omitempty"`
}

func (t *ResetPasswordResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "resetPasswordResultOk":
		t.ResetPasswordResultOk = new(ResetPasswordResultOk)
		return json.Unmarshal(b, t.ResetPasswordResultOk)
	case "resetPasswordResultPending":
		t.ResetPasswordResultPending = new(ResetPasswordResultPending)
		return json.Unmarshal(b, t.ResetPasswordResultPending)
	case "resetPasswordResultDeclined":
		t.ResetPasswordResultDeclined = new(ResetPasswordResultDeclined)
		return json.Unmarshal(b, t.ResetPasswordResultDeclined)
	}
	return nil
}

func (t *ResetPasswordResult) MarshalJSON() ([]byte, error) {
	if t.ResetPasswordResultOk != nil {
		return json.Marshal(t.ResetPasswordResultOk)
	}
	if t.ResetPasswordResultPending != nil {
		return json.Marshal(t.ResetPasswordResultPending)
	}
	if t.ResetPasswordResultDeclined != nil {
		return json.Marshal(t.ResetPasswordResultDeclined)
	}
	return nil, fmt.Errorf("no value set for ResetPasswordResult")
}

// NotificationType Contains detailed information about a notification
type NotificationType struct {
	NotificationTypeNewMessage     *NotificationTypeNewMessage     `json:"notificationTypeNewMessage,omitempty"`
	NotificationTypeNewSecretChat  *NotificationTypeNewSecretChat  `json:"notificationTypeNewSecretChat,omitempty"`
	NotificationTypeNewCall        *NotificationTypeNewCall        `json:"notificationTypeNewCall,omitempty"`
	NotificationTypeNewPushMessage *NotificationTypeNewPushMessage `json:"notificationTypeNewPushMessage,omitempty"`
}

func (t *NotificationType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "notificationTypeNewMessage":
		t.NotificationTypeNewMessage = new(NotificationTypeNewMessage)
		return json.Unmarshal(b, t.NotificationTypeNewMessage)
	case "notificationTypeNewSecretChat":
		t.NotificationTypeNewSecretChat = new(NotificationTypeNewSecretChat)
		return json.Unmarshal(b, t.NotificationTypeNewSecretChat)
	case "notificationTypeNewCall":
		t.NotificationTypeNewCall = new(NotificationTypeNewCall)
		return json.Unmarshal(b, t.NotificationTypeNewCall)
	case "notificationTypeNewPushMessage":
		t.NotificationTypeNewPushMessage = new(NotificationTypeNewPushMessage)
		return json.Unmarshal(b, t.NotificationTypeNewPushMessage)
	}
	return nil
}

func (t *NotificationType) MarshalJSON() ([]byte, error) {
	if t.NotificationTypeNewMessage != nil {
		return json.Marshal(t.NotificationTypeNewMessage)
	}
	if t.NotificationTypeNewSecretChat != nil {
		return json.Marshal(t.NotificationTypeNewSecretChat)
	}
	if t.NotificationTypeNewCall != nil {
		return json.Marshal(t.NotificationTypeNewCall)
	}
	if t.NotificationTypeNewPushMessage != nil {
		return json.Marshal(t.NotificationTypeNewPushMessage)
	}
	return nil, fmt.Errorf("no value set for NotificationType")
}

// OptionValue Represents the value of an option
type OptionValue struct {
	OptionValueBoolean *OptionValueBoolean `json:"optionValueBoolean,omitempty"`
	OptionValueEmpty   *OptionValueEmpty   `json:"optionValueEmpty,omitempty"`
	OptionValueInteger *OptionValueInteger `json:"optionValueInteger,omitempty"`
	OptionValueString  *OptionValueString  `json:"optionValueString,omitempty"`
}

func (t *OptionValue) Type() string {
	return ""
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

// ChatMembersFilter Specifies the kind of chat members to return in searchChatMembers
type ChatMembersFilter struct {
	ChatMembersFilterContacts       *ChatMembersFilterContacts       `json:"chatMembersFilterContacts,omitempty"`
	ChatMembersFilterAdministrators *ChatMembersFilterAdministrators `json:"chatMembersFilterAdministrators,omitempty"`
	ChatMembersFilterMembers        *ChatMembersFilterMembers        `json:"chatMembersFilterMembers,omitempty"`
	ChatMembersFilterMention        *ChatMembersFilterMention        `json:"chatMembersFilterMention,omitempty"`
	ChatMembersFilterRestricted     *ChatMembersFilterRestricted     `json:"chatMembersFilterRestricted,omitempty"`
	ChatMembersFilterBanned         *ChatMembersFilterBanned         `json:"chatMembersFilterBanned,omitempty"`
	ChatMembersFilterBots           *ChatMembersFilterBots           `json:"chatMembersFilterBots,omitempty"`
}

func (t *ChatMembersFilter) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatMembersFilterContacts":
		t.ChatMembersFilterContacts = new(ChatMembersFilterContacts)
		return json.Unmarshal(b, t.ChatMembersFilterContacts)
	case "chatMembersFilterAdministrators":
		t.ChatMembersFilterAdministrators = new(ChatMembersFilterAdministrators)
		return json.Unmarshal(b, t.ChatMembersFilterAdministrators)
	case "chatMembersFilterMembers":
		t.ChatMembersFilterMembers = new(ChatMembersFilterMembers)
		return json.Unmarshal(b, t.ChatMembersFilterMembers)
	case "chatMembersFilterMention":
		t.ChatMembersFilterMention = new(ChatMembersFilterMention)
		return json.Unmarshal(b, t.ChatMembersFilterMention)
	case "chatMembersFilterRestricted":
		t.ChatMembersFilterRestricted = new(ChatMembersFilterRestricted)
		return json.Unmarshal(b, t.ChatMembersFilterRestricted)
	case "chatMembersFilterBanned":
		t.ChatMembersFilterBanned = new(ChatMembersFilterBanned)
		return json.Unmarshal(b, t.ChatMembersFilterBanned)
	case "chatMembersFilterBots":
		t.ChatMembersFilterBots = new(ChatMembersFilterBots)
		return json.Unmarshal(b, t.ChatMembersFilterBots)
	}
	return nil
}

func (t *ChatMembersFilter) MarshalJSON() ([]byte, error) {
	if t.ChatMembersFilterContacts != nil {
		return json.Marshal(t.ChatMembersFilterContacts)
	}
	if t.ChatMembersFilterAdministrators != nil {
		return json.Marshal(t.ChatMembersFilterAdministrators)
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
	if t.ChatMembersFilterBanned != nil {
		return json.Marshal(t.ChatMembersFilterBanned)
	}
	if t.ChatMembersFilterBots != nil {
		return json.Marshal(t.ChatMembersFilterBots)
	}
	return nil, fmt.Errorf("no value set for ChatMembersFilter")
}

// InputInvoice Describes an invoice to process
type InputInvoice struct {
	InputInvoiceMessage  *InputInvoiceMessage  `json:"inputInvoiceMessage,omitempty"`
	InputInvoiceName     *InputInvoiceName     `json:"inputInvoiceName,omitempty"`
	InputInvoiceTelegram *InputInvoiceTelegram `json:"inputInvoiceTelegram,omitempty"`
}

func (t *InputInvoice) Type() string {
	return ""
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

// InternalLinkType Describes an internal https://t.me or tg: link, which must be processed by the application in a special way
type InternalLinkType struct {
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
	return ""
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

// StoryInteractionType Describes type of interaction with a story
type StoryInteractionType struct {
	StoryInteractionTypeView    *StoryInteractionTypeView    `json:"storyInteractionTypeView,omitempty"`
	StoryInteractionTypeForward *StoryInteractionTypeForward `json:"storyInteractionTypeForward,omitempty"`
	StoryInteractionTypeRepost  *StoryInteractionTypeRepost  `json:"storyInteractionTypeRepost,omitempty"`
}

func (t *StoryInteractionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyInteractionTypeView":
		t.StoryInteractionTypeView = new(StoryInteractionTypeView)
		return json.Unmarshal(b, t.StoryInteractionTypeView)
	case "storyInteractionTypeForward":
		t.StoryInteractionTypeForward = new(StoryInteractionTypeForward)
		return json.Unmarshal(b, t.StoryInteractionTypeForward)
	case "storyInteractionTypeRepost":
		t.StoryInteractionTypeRepost = new(StoryInteractionTypeRepost)
		return json.Unmarshal(b, t.StoryInteractionTypeRepost)
	}
	return nil
}

func (t *StoryInteractionType) MarshalJSON() ([]byte, error) {
	if t.StoryInteractionTypeView != nil {
		return json.Marshal(t.StoryInteractionTypeView)
	}
	if t.StoryInteractionTypeForward != nil {
		return json.Marshal(t.StoryInteractionTypeForward)
	}
	if t.StoryInteractionTypeRepost != nil {
		return json.Marshal(t.StoryInteractionTypeRepost)
	}
	return nil, fmt.Errorf("no value set for StoryInteractionType")
}

// ChatPhotoStickerType Describes type of sticker, which was used to create a chat photo
type ChatPhotoStickerType struct {
	ChatPhotoStickerTypeRegularOrMask *ChatPhotoStickerTypeRegularOrMask `json:"chatPhotoStickerTypeRegularOrMask,omitempty"`
	ChatPhotoStickerTypeCustomEmoji   *ChatPhotoStickerTypeCustomEmoji   `json:"chatPhotoStickerTypeCustomEmoji,omitempty"`
}

func (t *ChatPhotoStickerType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatPhotoStickerTypeRegularOrMask":
		t.ChatPhotoStickerTypeRegularOrMask = new(ChatPhotoStickerTypeRegularOrMask)
		return json.Unmarshal(b, t.ChatPhotoStickerTypeRegularOrMask)
	case "chatPhotoStickerTypeCustomEmoji":
		t.ChatPhotoStickerTypeCustomEmoji = new(ChatPhotoStickerTypeCustomEmoji)
		return json.Unmarshal(b, t.ChatPhotoStickerTypeCustomEmoji)
	}
	return nil
}

func (t *ChatPhotoStickerType) MarshalJSON() ([]byte, error) {
	if t.ChatPhotoStickerTypeRegularOrMask != nil {
		return json.Marshal(t.ChatPhotoStickerTypeRegularOrMask)
	}
	if t.ChatPhotoStickerTypeCustomEmoji != nil {
		return json.Marshal(t.ChatPhotoStickerTypeCustomEmoji)
	}
	return nil, fmt.Errorf("no value set for ChatPhotoStickerType")
}

// InputChatPhoto Describes a photo to be set as a user profile or chat photo
type InputChatPhoto struct {
	InputChatPhotoPrevious  *InputChatPhotoPrevious  `json:"inputChatPhotoPrevious,omitempty"`
	InputChatPhotoStatic    *InputChatPhotoStatic    `json:"inputChatPhotoStatic,omitempty"`
	InputChatPhotoAnimation *InputChatPhotoAnimation `json:"inputChatPhotoAnimation,omitempty"`
	InputChatPhotoSticker   *InputChatPhotoSticker   `json:"inputChatPhotoSticker,omitempty"`
}

func (t *InputChatPhoto) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputChatPhotoPrevious":
		t.InputChatPhotoPrevious = new(InputChatPhotoPrevious)
		return json.Unmarshal(b, t.InputChatPhotoPrevious)
	case "inputChatPhotoStatic":
		t.InputChatPhotoStatic = new(InputChatPhotoStatic)
		return json.Unmarshal(b, t.InputChatPhotoStatic)
	case "inputChatPhotoAnimation":
		t.InputChatPhotoAnimation = new(InputChatPhotoAnimation)
		return json.Unmarshal(b, t.InputChatPhotoAnimation)
	case "inputChatPhotoSticker":
		t.InputChatPhotoSticker = new(InputChatPhotoSticker)
		return json.Unmarshal(b, t.InputChatPhotoSticker)
	}
	return nil
}

func (t *InputChatPhoto) MarshalJSON() ([]byte, error) {
	if t.InputChatPhotoPrevious != nil {
		return json.Marshal(t.InputChatPhotoPrevious)
	}
	if t.InputChatPhotoStatic != nil {
		return json.Marshal(t.InputChatPhotoStatic)
	}
	if t.InputChatPhotoAnimation != nil {
		return json.Marshal(t.InputChatPhotoAnimation)
	}
	if t.InputChatPhotoSticker != nil {
		return json.Marshal(t.InputChatPhotoSticker)
	}
	return nil, fmt.Errorf("no value set for InputChatPhoto")
}

// MessageEffectType Describes type of emoji effect
type MessageEffectType struct {
	MessageEffectTypeEmojiReaction  *MessageEffectTypeEmojiReaction  `json:"messageEffectTypeEmojiReaction,omitempty"`
	MessageEffectTypePremiumSticker *MessageEffectTypePremiumSticker `json:"messageEffectTypePremiumSticker,omitempty"`
}

func (t *MessageEffectType) Type() string {
	return ""
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

// InlineKeyboardButtonType Describes the type of inline keyboard button
type InlineKeyboardButtonType struct {
	InlineKeyboardButtonTypeUrl                  *InlineKeyboardButtonTypeUrl                  `json:"inlineKeyboardButtonTypeUrl,omitempty"`
	InlineKeyboardButtonTypeLoginUrl             *InlineKeyboardButtonTypeLoginUrl             `json:"inlineKeyboardButtonTypeLoginUrl,omitempty"`
	InlineKeyboardButtonTypeWebApp               *InlineKeyboardButtonTypeWebApp               `json:"inlineKeyboardButtonTypeWebApp,omitempty"`
	InlineKeyboardButtonTypeCallback             *InlineKeyboardButtonTypeCallback             `json:"inlineKeyboardButtonTypeCallback,omitempty"`
	InlineKeyboardButtonTypeCallbackWithPassword *InlineKeyboardButtonTypeCallbackWithPassword `json:"inlineKeyboardButtonTypeCallbackWithPassword,omitempty"`
	InlineKeyboardButtonTypeCallbackGame         *InlineKeyboardButtonTypeCallbackGame         `json:"inlineKeyboardButtonTypeCallbackGame,omitempty"`
	InlineKeyboardButtonTypeSwitchInline         *InlineKeyboardButtonTypeSwitchInline         `json:"inlineKeyboardButtonTypeSwitchInline,omitempty"`
	InlineKeyboardButtonTypeBuy                  *InlineKeyboardButtonTypeBuy                  `json:"inlineKeyboardButtonTypeBuy,omitempty"`
	InlineKeyboardButtonTypeUser                 *InlineKeyboardButtonTypeUser                 `json:"inlineKeyboardButtonTypeUser,omitempty"`
	InlineKeyboardButtonTypeCopyText             *InlineKeyboardButtonTypeCopyText             `json:"inlineKeyboardButtonTypeCopyText,omitempty"`
}

func (t *InlineKeyboardButtonType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inlineKeyboardButtonTypeUrl":
		t.InlineKeyboardButtonTypeUrl = new(InlineKeyboardButtonTypeUrl)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeUrl)
	case "inlineKeyboardButtonTypeLoginUrl":
		t.InlineKeyboardButtonTypeLoginUrl = new(InlineKeyboardButtonTypeLoginUrl)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeLoginUrl)
	case "inlineKeyboardButtonTypeWebApp":
		t.InlineKeyboardButtonTypeWebApp = new(InlineKeyboardButtonTypeWebApp)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeWebApp)
	case "inlineKeyboardButtonTypeCallback":
		t.InlineKeyboardButtonTypeCallback = new(InlineKeyboardButtonTypeCallback)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallback)
	case "inlineKeyboardButtonTypeCallbackWithPassword":
		t.InlineKeyboardButtonTypeCallbackWithPassword = new(InlineKeyboardButtonTypeCallbackWithPassword)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallbackWithPassword)
	case "inlineKeyboardButtonTypeCallbackGame":
		t.InlineKeyboardButtonTypeCallbackGame = new(InlineKeyboardButtonTypeCallbackGame)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCallbackGame)
	case "inlineKeyboardButtonTypeSwitchInline":
		t.InlineKeyboardButtonTypeSwitchInline = new(InlineKeyboardButtonTypeSwitchInline)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeSwitchInline)
	case "inlineKeyboardButtonTypeBuy":
		t.InlineKeyboardButtonTypeBuy = new(InlineKeyboardButtonTypeBuy)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeBuy)
	case "inlineKeyboardButtonTypeUser":
		t.InlineKeyboardButtonTypeUser = new(InlineKeyboardButtonTypeUser)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeUser)
	case "inlineKeyboardButtonTypeCopyText":
		t.InlineKeyboardButtonTypeCopyText = new(InlineKeyboardButtonTypeCopyText)
		return json.Unmarshal(b, t.InlineKeyboardButtonTypeCopyText)
	}
	return nil
}

func (t *InlineKeyboardButtonType) MarshalJSON() ([]byte, error) {
	if t.InlineKeyboardButtonTypeUrl != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeUrl)
	}
	if t.InlineKeyboardButtonTypeLoginUrl != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeLoginUrl)
	}
	if t.InlineKeyboardButtonTypeWebApp != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeWebApp)
	}
	if t.InlineKeyboardButtonTypeCallback != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallback)
	}
	if t.InlineKeyboardButtonTypeCallbackWithPassword != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallbackWithPassword)
	}
	if t.InlineKeyboardButtonTypeCallbackGame != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCallbackGame)
	}
	if t.InlineKeyboardButtonTypeSwitchInline != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeSwitchInline)
	}
	if t.InlineKeyboardButtonTypeBuy != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeBuy)
	}
	if t.InlineKeyboardButtonTypeUser != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeUser)
	}
	if t.InlineKeyboardButtonTypeCopyText != nil {
		return json.Marshal(t.InlineKeyboardButtonTypeCopyText)
	}
	return nil, fmt.Errorf("no value set for InlineKeyboardButtonType")
}

// SessionType Represents the type of session
type SessionType struct {
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
	return ""
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

// SecretChatState Describes the current secret chat state
type SecretChatState struct {
	SecretChatStatePending *SecretChatStatePending `json:"secretChatStatePending,omitempty"`
	SecretChatStateReady   *SecretChatStateReady   `json:"secretChatStateReady,omitempty"`
	SecretChatStateClosed  *SecretChatStateClosed  `json:"secretChatStateClosed,omitempty"`
}

func (t *SecretChatState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "secretChatStatePending":
		t.SecretChatStatePending = new(SecretChatStatePending)
		return json.Unmarshal(b, t.SecretChatStatePending)
	case "secretChatStateReady":
		t.SecretChatStateReady = new(SecretChatStateReady)
		return json.Unmarshal(b, t.SecretChatStateReady)
	case "secretChatStateClosed":
		t.SecretChatStateClosed = new(SecretChatStateClosed)
		return json.Unmarshal(b, t.SecretChatStateClosed)
	}
	return nil
}

func (t *SecretChatState) MarshalJSON() ([]byte, error) {
	if t.SecretChatStatePending != nil {
		return json.Marshal(t.SecretChatStatePending)
	}
	if t.SecretChatStateReady != nil {
		return json.Marshal(t.SecretChatStateReady)
	}
	if t.SecretChatStateClosed != nil {
		return json.Marshal(t.SecretChatStateClosed)
	}
	return nil, fmt.Errorf("no value set for SecretChatState")
}

// LoginUrlInfo Contains information about an inline button of type inlineKeyboardButtonTypeLoginUrl
type LoginUrlInfo struct {
	LoginUrlInfoOpen                *LoginUrlInfoOpen                `json:"loginUrlInfoOpen,omitempty"`
	LoginUrlInfoRequestConfirmation *LoginUrlInfoRequestConfirmation `json:"loginUrlInfoRequestConfirmation,omitempty"`
}

func (t *LoginUrlInfo) Type() string {
	return ""
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

// PaymentProvider Contains information about a payment provider
type PaymentProvider struct {
	PaymentProviderSmartGlocal *PaymentProviderSmartGlocal `json:"paymentProviderSmartGlocal,omitempty"`
	PaymentProviderStripe      *PaymentProviderStripe      `json:"paymentProviderStripe,omitempty"`
	PaymentProviderOther       *PaymentProviderOther       `json:"paymentProviderOther,omitempty"`
}

func (t *PaymentProvider) Type() string {
	return ""
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
	switch typeObj.Type {
	case "paymentProviderSmartGlocal":
		t.PaymentProviderSmartGlocal = new(PaymentProviderSmartGlocal)
		return json.Unmarshal(b, t.PaymentProviderSmartGlocal)
	case "paymentProviderStripe":
		t.PaymentProviderStripe = new(PaymentProviderStripe)
		return json.Unmarshal(b, t.PaymentProviderStripe)
	case "paymentProviderOther":
		t.PaymentProviderOther = new(PaymentProviderOther)
		return json.Unmarshal(b, t.PaymentProviderOther)
	}
	return nil
}

func (t *PaymentProvider) MarshalJSON() ([]byte, error) {
	if t.PaymentProviderSmartGlocal != nil {
		return json.Marshal(t.PaymentProviderSmartGlocal)
	}
	if t.PaymentProviderStripe != nil {
		return json.Marshal(t.PaymentProviderStripe)
	}
	if t.PaymentProviderOther != nil {
		return json.Marshal(t.PaymentProviderOther)
	}
	return nil, fmt.Errorf("no value set for PaymentProvider")
}

// SpeechRecognitionResult Describes result of speech recognition in a voice note
type SpeechRecognitionResult struct {
	SpeechRecognitionResultPending *SpeechRecognitionResultPending `json:"speechRecognitionResultPending,omitempty"`
	SpeechRecognitionResultText    *SpeechRecognitionResultText    `json:"speechRecognitionResultText,omitempty"`
	SpeechRecognitionResultError   *SpeechRecognitionResultError   `json:"speechRecognitionResultError,omitempty"`
}

func (t *SpeechRecognitionResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "speechRecognitionResultPending":
		t.SpeechRecognitionResultPending = new(SpeechRecognitionResultPending)
		return json.Unmarshal(b, t.SpeechRecognitionResultPending)
	case "speechRecognitionResultText":
		t.SpeechRecognitionResultText = new(SpeechRecognitionResultText)
		return json.Unmarshal(b, t.SpeechRecognitionResultText)
	case "speechRecognitionResultError":
		t.SpeechRecognitionResultError = new(SpeechRecognitionResultError)
		return json.Unmarshal(b, t.SpeechRecognitionResultError)
	}
	return nil
}

func (t *SpeechRecognitionResult) MarshalJSON() ([]byte, error) {
	if t.SpeechRecognitionResultPending != nil {
		return json.Marshal(t.SpeechRecognitionResultPending)
	}
	if t.SpeechRecognitionResultText != nil {
		return json.Marshal(t.SpeechRecognitionResultText)
	}
	if t.SpeechRecognitionResultError != nil {
		return json.Marshal(t.SpeechRecognitionResultError)
	}
	return nil, fmt.Errorf("no value set for SpeechRecognitionResult")
}

// EmailAddressAuthentication Contains authentication data for an email address
type EmailAddressAuthentication struct {
	EmailAddressAuthenticationCode     *EmailAddressAuthenticationCode     `json:"emailAddressAuthenticationCode,omitempty"`
	EmailAddressAuthenticationAppleId  *EmailAddressAuthenticationAppleId  `json:"emailAddressAuthenticationAppleId,omitempty"`
	EmailAddressAuthenticationGoogleId *EmailAddressAuthenticationGoogleId `json:"emailAddressAuthenticationGoogleId,omitempty"`
}

func (t *EmailAddressAuthentication) Type() string {
	return ""
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
	switch typeObj.Type {
	case "emailAddressAuthenticationCode":
		t.EmailAddressAuthenticationCode = new(EmailAddressAuthenticationCode)
		return json.Unmarshal(b, t.EmailAddressAuthenticationCode)
	case "emailAddressAuthenticationAppleId":
		t.EmailAddressAuthenticationAppleId = new(EmailAddressAuthenticationAppleId)
		return json.Unmarshal(b, t.EmailAddressAuthenticationAppleId)
	case "emailAddressAuthenticationGoogleId":
		t.EmailAddressAuthenticationGoogleId = new(EmailAddressAuthenticationGoogleId)
		return json.Unmarshal(b, t.EmailAddressAuthenticationGoogleId)
	}
	return nil
}

func (t *EmailAddressAuthentication) MarshalJSON() ([]byte, error) {
	if t.EmailAddressAuthenticationCode != nil {
		return json.Marshal(t.EmailAddressAuthenticationCode)
	}
	if t.EmailAddressAuthenticationAppleId != nil {
		return json.Marshal(t.EmailAddressAuthenticationAppleId)
	}
	if t.EmailAddressAuthenticationGoogleId != nil {
		return json.Marshal(t.EmailAddressAuthenticationGoogleId)
	}
	return nil, fmt.Errorf("no value set for EmailAddressAuthentication")
}

// StickerFullType Contains full information about sticker type
type StickerFullType struct {
	StickerFullTypeRegular     *StickerFullTypeRegular     `json:"stickerFullTypeRegular,omitempty"`
	StickerFullTypeMask        *StickerFullTypeMask        `json:"stickerFullTypeMask,omitempty"`
	StickerFullTypeCustomEmoji *StickerFullTypeCustomEmoji `json:"stickerFullTypeCustomEmoji,omitempty"`
}

func (t *StickerFullType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "stickerFullTypeRegular":
		t.StickerFullTypeRegular = new(StickerFullTypeRegular)
		return json.Unmarshal(b, t.StickerFullTypeRegular)
	case "stickerFullTypeMask":
		t.StickerFullTypeMask = new(StickerFullTypeMask)
		return json.Unmarshal(b, t.StickerFullTypeMask)
	case "stickerFullTypeCustomEmoji":
		t.StickerFullTypeCustomEmoji = new(StickerFullTypeCustomEmoji)
		return json.Unmarshal(b, t.StickerFullTypeCustomEmoji)
	}
	return nil
}

func (t *StickerFullType) MarshalJSON() ([]byte, error) {
	if t.StickerFullTypeRegular != nil {
		return json.Marshal(t.StickerFullTypeRegular)
	}
	if t.StickerFullTypeMask != nil {
		return json.Marshal(t.StickerFullTypeMask)
	}
	if t.StickerFullTypeCustomEmoji != nil {
		return json.Marshal(t.StickerFullTypeCustomEmoji)
	}
	return nil, fmt.Errorf("no value set for StickerFullType")
}

// CallbackQueryPayload Represents a payload of a callback query
type CallbackQueryPayload struct {
	CallbackQueryPayloadData             *CallbackQueryPayloadData             `json:"callbackQueryPayloadData,omitempty"`
	CallbackQueryPayloadDataWithPassword *CallbackQueryPayloadDataWithPassword `json:"callbackQueryPayloadDataWithPassword,omitempty"`
	CallbackQueryPayloadGame             *CallbackQueryPayloadGame             `json:"callbackQueryPayloadGame,omitempty"`
}

func (t *CallbackQueryPayload) Type() string {
	return ""
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

// ReportChatResult Describes result of chat report
type ReportChatResult struct {
	ReportChatResultOk               *ReportChatResultOk               `json:"reportChatResultOk,omitempty"`
	ReportChatResultOptionRequired   *ReportChatResultOptionRequired   `json:"reportChatResultOptionRequired,omitempty"`
	ReportChatResultTextRequired     *ReportChatResultTextRequired     `json:"reportChatResultTextRequired,omitempty"`
	ReportChatResultMessagesRequired *ReportChatResultMessagesRequired `json:"reportChatResultMessagesRequired,omitempty"`
}

func (t *ReportChatResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "reportChatResultOk":
		t.ReportChatResultOk = new(ReportChatResultOk)
		return json.Unmarshal(b, t.ReportChatResultOk)
	case "reportChatResultOptionRequired":
		t.ReportChatResultOptionRequired = new(ReportChatResultOptionRequired)
		return json.Unmarshal(b, t.ReportChatResultOptionRequired)
	case "reportChatResultTextRequired":
		t.ReportChatResultTextRequired = new(ReportChatResultTextRequired)
		return json.Unmarshal(b, t.ReportChatResultTextRequired)
	case "reportChatResultMessagesRequired":
		t.ReportChatResultMessagesRequired = new(ReportChatResultMessagesRequired)
		return json.Unmarshal(b, t.ReportChatResultMessagesRequired)
	}
	return nil
}

func (t *ReportChatResult) MarshalJSON() ([]byte, error) {
	if t.ReportChatResultOk != nil {
		return json.Marshal(t.ReportChatResultOk)
	}
	if t.ReportChatResultOptionRequired != nil {
		return json.Marshal(t.ReportChatResultOptionRequired)
	}
	if t.ReportChatResultTextRequired != nil {
		return json.Marshal(t.ReportChatResultTextRequired)
	}
	if t.ReportChatResultMessagesRequired != nil {
		return json.Marshal(t.ReportChatResultMessagesRequired)
	}
	return nil, fmt.Errorf("no value set for ReportChatResult")
}

// CallState Describes the current call state
type CallState struct {
	CallStatePending        *CallStatePending        `json:"callStatePending,omitempty"`
	CallStateExchangingKeys *CallStateExchangingKeys `json:"callStateExchangingKeys,omitempty"`
	CallStateReady          *CallStateReady          `json:"callStateReady,omitempty"`
	CallStateHangingUp      *CallStateHangingUp      `json:"callStateHangingUp,omitempty"`
	CallStateDiscarded      *CallStateDiscarded      `json:"callStateDiscarded,omitempty"`
	CallStateError          *CallStateError          `json:"callStateError,omitempty"`
}

func (t *CallState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "callStatePending":
		t.CallStatePending = new(CallStatePending)
		return json.Unmarshal(b, t.CallStatePending)
	case "callStateExchangingKeys":
		t.CallStateExchangingKeys = new(CallStateExchangingKeys)
		return json.Unmarshal(b, t.CallStateExchangingKeys)
	case "callStateReady":
		t.CallStateReady = new(CallStateReady)
		return json.Unmarshal(b, t.CallStateReady)
	case "callStateHangingUp":
		t.CallStateHangingUp = new(CallStateHangingUp)
		return json.Unmarshal(b, t.CallStateHangingUp)
	case "callStateDiscarded":
		t.CallStateDiscarded = new(CallStateDiscarded)
		return json.Unmarshal(b, t.CallStateDiscarded)
	case "callStateError":
		t.CallStateError = new(CallStateError)
		return json.Unmarshal(b, t.CallStateError)
	}
	return nil
}

func (t *CallState) MarshalJSON() ([]byte, error) {
	if t.CallStatePending != nil {
		return json.Marshal(t.CallStatePending)
	}
	if t.CallStateExchangingKeys != nil {
		return json.Marshal(t.CallStateExchangingKeys)
	}
	if t.CallStateReady != nil {
		return json.Marshal(t.CallStateReady)
	}
	if t.CallStateHangingUp != nil {
		return json.Marshal(t.CallStateHangingUp)
	}
	if t.CallStateDiscarded != nil {
		return json.Marshal(t.CallStateDiscarded)
	}
	if t.CallStateError != nil {
		return json.Marshal(t.CallStateError)
	}
	return nil, fmt.Errorf("no value set for CallState")
}

// InputGroupCall Describes a non-joined group call that isn't bound to a chat
type InputGroupCall struct {
	InputGroupCallLink    *InputGroupCallLink    `json:"inputGroupCallLink,omitempty"`
	InputGroupCallMessage *InputGroupCallMessage `json:"inputGroupCallMessage,omitempty"`
}

func (t *InputGroupCall) Type() string {
	return ""
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

// LanguagePackStringValue Represents the value of a string in a language pack
type LanguagePackStringValue struct {
	LanguagePackStringValueOrdinary   *LanguagePackStringValueOrdinary   `json:"languagePackStringValueOrdinary,omitempty"`
	LanguagePackStringValuePluralized *LanguagePackStringValuePluralized `json:"languagePackStringValuePluralized,omitempty"`
	LanguagePackStringValueDeleted    *LanguagePackStringValueDeleted    `json:"languagePackStringValueDeleted,omitempty"`
}

func (t *LanguagePackStringValue) Type() string {
	return ""
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
	switch typeObj.Type {
	case "languagePackStringValueOrdinary":
		t.LanguagePackStringValueOrdinary = new(LanguagePackStringValueOrdinary)
		return json.Unmarshal(b, t.LanguagePackStringValueOrdinary)
	case "languagePackStringValuePluralized":
		t.LanguagePackStringValuePluralized = new(LanguagePackStringValuePluralized)
		return json.Unmarshal(b, t.LanguagePackStringValuePluralized)
	case "languagePackStringValueDeleted":
		t.LanguagePackStringValueDeleted = new(LanguagePackStringValueDeleted)
		return json.Unmarshal(b, t.LanguagePackStringValueDeleted)
	}
	return nil
}

func (t *LanguagePackStringValue) MarshalJSON() ([]byte, error) {
	if t.LanguagePackStringValueOrdinary != nil {
		return json.Marshal(t.LanguagePackStringValueOrdinary)
	}
	if t.LanguagePackStringValuePluralized != nil {
		return json.Marshal(t.LanguagePackStringValuePluralized)
	}
	if t.LanguagePackStringValueDeleted != nil {
		return json.Marshal(t.LanguagePackStringValueDeleted)
	}
	return nil, fmt.Errorf("no value set for LanguagePackStringValue")
}

// AutosaveSettingsScope Describes scope of autosave settings
type AutosaveSettingsScope struct {
	AutosaveSettingsScopePrivateChats *AutosaveSettingsScopePrivateChats `json:"autosaveSettingsScopePrivateChats,omitempty"`
	AutosaveSettingsScopeGroupChats   *AutosaveSettingsScopeGroupChats   `json:"autosaveSettingsScopeGroupChats,omitempty"`
	AutosaveSettingsScopeChannelChats *AutosaveSettingsScopeChannelChats `json:"autosaveSettingsScopeChannelChats,omitempty"`
	AutosaveSettingsScopeChat         *AutosaveSettingsScopeChat         `json:"autosaveSettingsScopeChat,omitempty"`
}

func (t *AutosaveSettingsScope) Type() string {
	return ""
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
	switch typeObj.Type {
	case "autosaveSettingsScopePrivateChats":
		t.AutosaveSettingsScopePrivateChats = new(AutosaveSettingsScopePrivateChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopePrivateChats)
	case "autosaveSettingsScopeGroupChats":
		t.AutosaveSettingsScopeGroupChats = new(AutosaveSettingsScopeGroupChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopeGroupChats)
	case "autosaveSettingsScopeChannelChats":
		t.AutosaveSettingsScopeChannelChats = new(AutosaveSettingsScopeChannelChats)
		return json.Unmarshal(b, t.AutosaveSettingsScopeChannelChats)
	case "autosaveSettingsScopeChat":
		t.AutosaveSettingsScopeChat = new(AutosaveSettingsScopeChat)
		return json.Unmarshal(b, t.AutosaveSettingsScopeChat)
	}
	return nil
}

func (t *AutosaveSettingsScope) MarshalJSON() ([]byte, error) {
	if t.AutosaveSettingsScopePrivateChats != nil {
		return json.Marshal(t.AutosaveSettingsScopePrivateChats)
	}
	if t.AutosaveSettingsScopeGroupChats != nil {
		return json.Marshal(t.AutosaveSettingsScopeGroupChats)
	}
	if t.AutosaveSettingsScopeChannelChats != nil {
		return json.Marshal(t.AutosaveSettingsScopeChannelChats)
	}
	if t.AutosaveSettingsScopeChat != nil {
		return json.Marshal(t.AutosaveSettingsScopeChat)
	}
	return nil, fmt.Errorf("no value set for AutosaveSettingsScope")
}

// FirebaseDeviceVerificationParameters Describes parameters to be used for device verification
type FirebaseDeviceVerificationParameters struct {
	FirebaseDeviceVerificationParametersSafetyNet     *FirebaseDeviceVerificationParametersSafetyNet     `json:"firebaseDeviceVerificationParametersSafetyNet,omitempty"`
	FirebaseDeviceVerificationParametersPlayIntegrity *FirebaseDeviceVerificationParametersPlayIntegrity `json:"firebaseDeviceVerificationParametersPlayIntegrity,omitempty"`
}

func (t *FirebaseDeviceVerificationParameters) Type() string {
	return ""
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
	switch typeObj.Type {
	case "firebaseDeviceVerificationParametersSafetyNet":
		t.FirebaseDeviceVerificationParametersSafetyNet = new(FirebaseDeviceVerificationParametersSafetyNet)
		return json.Unmarshal(b, t.FirebaseDeviceVerificationParametersSafetyNet)
	case "firebaseDeviceVerificationParametersPlayIntegrity":
		t.FirebaseDeviceVerificationParametersPlayIntegrity = new(FirebaseDeviceVerificationParametersPlayIntegrity)
		return json.Unmarshal(b, t.FirebaseDeviceVerificationParametersPlayIntegrity)
	}
	return nil
}

func (t *FirebaseDeviceVerificationParameters) MarshalJSON() ([]byte, error) {
	if t.FirebaseDeviceVerificationParametersSafetyNet != nil {
		return json.Marshal(t.FirebaseDeviceVerificationParametersSafetyNet)
	}
	if t.FirebaseDeviceVerificationParametersPlayIntegrity != nil {
		return json.Marshal(t.FirebaseDeviceVerificationParametersPlayIntegrity)
	}
	return nil, fmt.Errorf("no value set for FirebaseDeviceVerificationParameters")
}

// SuggestedPostRefundReason Describes reason for refund of the payment for a suggested post
type SuggestedPostRefundReason struct {
	SuggestedPostRefundReasonPostDeleted     *SuggestedPostRefundReasonPostDeleted     `json:"suggestedPostRefundReasonPostDeleted,omitempty"`
	SuggestedPostRefundReasonPaymentRefunded *SuggestedPostRefundReasonPaymentRefunded `json:"suggestedPostRefundReasonPaymentRefunded,omitempty"`
}

func (t *SuggestedPostRefundReason) Type() string {
	return ""
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
	switch typeObj.Type {
	case "suggestedPostRefundReasonPostDeleted":
		t.SuggestedPostRefundReasonPostDeleted = new(SuggestedPostRefundReasonPostDeleted)
		return json.Unmarshal(b, t.SuggestedPostRefundReasonPostDeleted)
	case "suggestedPostRefundReasonPaymentRefunded":
		t.SuggestedPostRefundReasonPaymentRefunded = new(SuggestedPostRefundReasonPaymentRefunded)
		return json.Unmarshal(b, t.SuggestedPostRefundReasonPaymentRefunded)
	}
	return nil
}

func (t *SuggestedPostRefundReason) MarshalJSON() ([]byte, error) {
	if t.SuggestedPostRefundReasonPostDeleted != nil {
		return json.Marshal(t.SuggestedPostRefundReasonPostDeleted)
	}
	if t.SuggestedPostRefundReasonPaymentRefunded != nil {
		return json.Marshal(t.SuggestedPostRefundReasonPaymentRefunded)
	}
	return nil, fmt.Errorf("no value set for SuggestedPostRefundReason")
}

// MessageTopic Describes a topic of messages in a chat
type MessageTopic struct {
	MessageTopicThread         *MessageTopicThread         `json:"messageTopicThread,omitempty"`
	MessageTopicForum          *MessageTopicForum          `json:"messageTopicForum,omitempty"`
	MessageTopicDirectMessages *MessageTopicDirectMessages `json:"messageTopicDirectMessages,omitempty"`
	MessageTopicSavedMessages  *MessageTopicSavedMessages  `json:"messageTopicSavedMessages,omitempty"`
}

func (t *MessageTopic) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageTopicThread":
		t.MessageTopicThread = new(MessageTopicThread)
		return json.Unmarshal(b, t.MessageTopicThread)
	case "messageTopicForum":
		t.MessageTopicForum = new(MessageTopicForum)
		return json.Unmarshal(b, t.MessageTopicForum)
	case "messageTopicDirectMessages":
		t.MessageTopicDirectMessages = new(MessageTopicDirectMessages)
		return json.Unmarshal(b, t.MessageTopicDirectMessages)
	case "messageTopicSavedMessages":
		t.MessageTopicSavedMessages = new(MessageTopicSavedMessages)
		return json.Unmarshal(b, t.MessageTopicSavedMessages)
	}
	return nil
}

func (t *MessageTopic) MarshalJSON() ([]byte, error) {
	if t.MessageTopicThread != nil {
		return json.Marshal(t.MessageTopicThread)
	}
	if t.MessageTopicForum != nil {
		return json.Marshal(t.MessageTopicForum)
	}
	if t.MessageTopicDirectMessages != nil {
		return json.Marshal(t.MessageTopicDirectMessages)
	}
	if t.MessageTopicSavedMessages != nil {
		return json.Marshal(t.MessageTopicSavedMessages)
	}
	return nil, fmt.Errorf("no value set for MessageTopic")
}

// InputPaidMediaType Describes type of paid media to sent
type InputPaidMediaType struct {
	InputPaidMediaTypePhoto *InputPaidMediaTypePhoto `json:"inputPaidMediaTypePhoto,omitempty"`
	InputPaidMediaTypeVideo *InputPaidMediaTypeVideo `json:"inputPaidMediaTypeVideo,omitempty"`
}

func (t *InputPaidMediaType) Type() string {
	return ""
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

// DeviceToken Represents a data needed to subscribe for push notifications through registerDevice method.
type DeviceToken struct {
	DeviceTokenFirebaseCloudMessaging *DeviceTokenFirebaseCloudMessaging `json:"deviceTokenFirebaseCloudMessaging,omitempty"`
	DeviceTokenApplePush              *DeviceTokenApplePush              `json:"deviceTokenApplePush,omitempty"`
	DeviceTokenApplePushVoIP          *DeviceTokenApplePushVoIP          `json:"deviceTokenApplePushVoIP,omitempty"`
	DeviceTokenWindowsPush            *DeviceTokenWindowsPush            `json:"deviceTokenWindowsPush,omitempty"`
	DeviceTokenMicrosoftPush          *DeviceTokenMicrosoftPush          `json:"deviceTokenMicrosoftPush,omitempty"`
	DeviceTokenMicrosoftPushVoIP      *DeviceTokenMicrosoftPushVoIP      `json:"deviceTokenMicrosoftPushVoIP,omitempty"`
	DeviceTokenWebPush                *DeviceTokenWebPush                `json:"deviceTokenWebPush,omitempty"`
	DeviceTokenSimplePush             *DeviceTokenSimplePush             `json:"deviceTokenSimplePush,omitempty"`
	DeviceTokenUbuntuPush             *DeviceTokenUbuntuPush             `json:"deviceTokenUbuntuPush,omitempty"`
	DeviceTokenBlackBerryPush         *DeviceTokenBlackBerryPush         `json:"deviceTokenBlackBerryPush,omitempty"`
	DeviceTokenTizenPush              *DeviceTokenTizenPush              `json:"deviceTokenTizenPush,omitempty"`
	DeviceTokenHuaweiPush             *DeviceTokenHuaweiPush             `json:"deviceTokenHuaweiPush,omitempty"`
}

func (t *DeviceToken) Type() string {
	return ""
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
	switch typeObj.Type {
	case "deviceTokenFirebaseCloudMessaging":
		t.DeviceTokenFirebaseCloudMessaging = new(DeviceTokenFirebaseCloudMessaging)
		return json.Unmarshal(b, t.DeviceTokenFirebaseCloudMessaging)
	case "deviceTokenApplePush":
		t.DeviceTokenApplePush = new(DeviceTokenApplePush)
		return json.Unmarshal(b, t.DeviceTokenApplePush)
	case "deviceTokenApplePushVoIP":
		t.DeviceTokenApplePushVoIP = new(DeviceTokenApplePushVoIP)
		return json.Unmarshal(b, t.DeviceTokenApplePushVoIP)
	case "deviceTokenWindowsPush":
		t.DeviceTokenWindowsPush = new(DeviceTokenWindowsPush)
		return json.Unmarshal(b, t.DeviceTokenWindowsPush)
	case "deviceTokenMicrosoftPush":
		t.DeviceTokenMicrosoftPush = new(DeviceTokenMicrosoftPush)
		return json.Unmarshal(b, t.DeviceTokenMicrosoftPush)
	case "deviceTokenMicrosoftPushVoIP":
		t.DeviceTokenMicrosoftPushVoIP = new(DeviceTokenMicrosoftPushVoIP)
		return json.Unmarshal(b, t.DeviceTokenMicrosoftPushVoIP)
	case "deviceTokenWebPush":
		t.DeviceTokenWebPush = new(DeviceTokenWebPush)
		return json.Unmarshal(b, t.DeviceTokenWebPush)
	case "deviceTokenSimplePush":
		t.DeviceTokenSimplePush = new(DeviceTokenSimplePush)
		return json.Unmarshal(b, t.DeviceTokenSimplePush)
	case "deviceTokenUbuntuPush":
		t.DeviceTokenUbuntuPush = new(DeviceTokenUbuntuPush)
		return json.Unmarshal(b, t.DeviceTokenUbuntuPush)
	case "deviceTokenBlackBerryPush":
		t.DeviceTokenBlackBerryPush = new(DeviceTokenBlackBerryPush)
		return json.Unmarshal(b, t.DeviceTokenBlackBerryPush)
	case "deviceTokenTizenPush":
		t.DeviceTokenTizenPush = new(DeviceTokenTizenPush)
		return json.Unmarshal(b, t.DeviceTokenTizenPush)
	case "deviceTokenHuaweiPush":
		t.DeviceTokenHuaweiPush = new(DeviceTokenHuaweiPush)
		return json.Unmarshal(b, t.DeviceTokenHuaweiPush)
	}
	return nil
}

func (t *DeviceToken) MarshalJSON() ([]byte, error) {
	if t.DeviceTokenFirebaseCloudMessaging != nil {
		return json.Marshal(t.DeviceTokenFirebaseCloudMessaging)
	}
	if t.DeviceTokenApplePush != nil {
		return json.Marshal(t.DeviceTokenApplePush)
	}
	if t.DeviceTokenApplePushVoIP != nil {
		return json.Marshal(t.DeviceTokenApplePushVoIP)
	}
	if t.DeviceTokenWindowsPush != nil {
		return json.Marshal(t.DeviceTokenWindowsPush)
	}
	if t.DeviceTokenMicrosoftPush != nil {
		return json.Marshal(t.DeviceTokenMicrosoftPush)
	}
	if t.DeviceTokenMicrosoftPushVoIP != nil {
		return json.Marshal(t.DeviceTokenMicrosoftPushVoIP)
	}
	if t.DeviceTokenWebPush != nil {
		return json.Marshal(t.DeviceTokenWebPush)
	}
	if t.DeviceTokenSimplePush != nil {
		return json.Marshal(t.DeviceTokenSimplePush)
	}
	if t.DeviceTokenUbuntuPush != nil {
		return json.Marshal(t.DeviceTokenUbuntuPush)
	}
	if t.DeviceTokenBlackBerryPush != nil {
		return json.Marshal(t.DeviceTokenBlackBerryPush)
	}
	if t.DeviceTokenTizenPush != nil {
		return json.Marshal(t.DeviceTokenTizenPush)
	}
	if t.DeviceTokenHuaweiPush != nil {
		return json.Marshal(t.DeviceTokenHuaweiPush)
	}
	return nil, fmt.Errorf("no value set for DeviceToken")
}

// TMeUrlType Describes the type of URL linking to an internal Telegram entity
type TMeUrlType struct {
	TMeUrlTypeUser       *TMeUrlTypeUser       `json:"tMeUrlTypeUser,omitempty"`
	TMeUrlTypeSupergroup *TMeUrlTypeSupergroup `json:"tMeUrlTypeSupergroup,omitempty"`
	TMeUrlTypeChatInvite *TMeUrlTypeChatInvite `json:"tMeUrlTypeChatInvite,omitempty"`
	TMeUrlTypeStickerSet *TMeUrlTypeStickerSet `json:"tMeUrlTypeStickerSet,omitempty"`
}

func (t *TMeUrlType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "tMeUrlTypeUser":
		t.TMeUrlTypeUser = new(TMeUrlTypeUser)
		return json.Unmarshal(b, t.TMeUrlTypeUser)
	case "tMeUrlTypeSupergroup":
		t.TMeUrlTypeSupergroup = new(TMeUrlTypeSupergroup)
		return json.Unmarshal(b, t.TMeUrlTypeSupergroup)
	case "tMeUrlTypeChatInvite":
		t.TMeUrlTypeChatInvite = new(TMeUrlTypeChatInvite)
		return json.Unmarshal(b, t.TMeUrlTypeChatInvite)
	case "tMeUrlTypeStickerSet":
		t.TMeUrlTypeStickerSet = new(TMeUrlTypeStickerSet)
		return json.Unmarshal(b, t.TMeUrlTypeStickerSet)
	}
	return nil
}

func (t *TMeUrlType) MarshalJSON() ([]byte, error) {
	if t.TMeUrlTypeUser != nil {
		return json.Marshal(t.TMeUrlTypeUser)
	}
	if t.TMeUrlTypeSupergroup != nil {
		return json.Marshal(t.TMeUrlTypeSupergroup)
	}
	if t.TMeUrlTypeChatInvite != nil {
		return json.Marshal(t.TMeUrlTypeChatInvite)
	}
	if t.TMeUrlTypeStickerSet != nil {
		return json.Marshal(t.TMeUrlTypeStickerSet)
	}
	return nil, fmt.Errorf("no value set for TMeUrlType")
}

// PhoneNumberCodeType Describes type of the request for which a code is sent to a phone number
type PhoneNumberCodeType struct {
	PhoneNumberCodeTypeChange           *PhoneNumberCodeTypeChange           `json:"phoneNumberCodeTypeChange,omitempty"`
	PhoneNumberCodeTypeVerify           *PhoneNumberCodeTypeVerify           `json:"phoneNumberCodeTypeVerify,omitempty"`
	PhoneNumberCodeTypeConfirmOwnership *PhoneNumberCodeTypeConfirmOwnership `json:"phoneNumberCodeTypeConfirmOwnership,omitempty"`
}

func (t *PhoneNumberCodeType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "phoneNumberCodeTypeChange":
		t.PhoneNumberCodeTypeChange = new(PhoneNumberCodeTypeChange)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeChange)
	case "phoneNumberCodeTypeVerify":
		t.PhoneNumberCodeTypeVerify = new(PhoneNumberCodeTypeVerify)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeVerify)
	case "phoneNumberCodeTypeConfirmOwnership":
		t.PhoneNumberCodeTypeConfirmOwnership = new(PhoneNumberCodeTypeConfirmOwnership)
		return json.Unmarshal(b, t.PhoneNumberCodeTypeConfirmOwnership)
	}
	return nil
}

func (t *PhoneNumberCodeType) MarshalJSON() ([]byte, error) {
	if t.PhoneNumberCodeTypeChange != nil {
		return json.Marshal(t.PhoneNumberCodeTypeChange)
	}
	if t.PhoneNumberCodeTypeVerify != nil {
		return json.Marshal(t.PhoneNumberCodeTypeVerify)
	}
	if t.PhoneNumberCodeTypeConfirmOwnership != nil {
		return json.Marshal(t.PhoneNumberCodeTypeConfirmOwnership)
	}
	return nil, fmt.Errorf("no value set for PhoneNumberCodeType")
}

// Update Contains notifications about data changes
type Update struct {
	UpdateAuthorizationState                       *UpdateAuthorizationState                       `json:"updateAuthorizationState,omitempty"`
	UpdateNewMessage                               *UpdateNewMessage                               `json:"updateNewMessage,omitempty"`
	UpdateMessageSendAcknowledged                  *UpdateMessageSendAcknowledged                  `json:"updateMessageSendAcknowledged,omitempty"`
	UpdateMessageSendSucceeded                     *UpdateMessageSendSucceeded                     `json:"updateMessageSendSucceeded,omitempty"`
	UpdateMessageSendFailed                        *UpdateMessageSendFailed                        `json:"updateMessageSendFailed,omitempty"`
	UpdateMessageContent                           *UpdateMessageContent                           `json:"updateMessageContent,omitempty"`
	UpdateMessageEdited                            *UpdateMessageEdited                            `json:"updateMessageEdited,omitempty"`
	UpdateMessageIsPinned                          *UpdateMessageIsPinned                          `json:"updateMessageIsPinned,omitempty"`
	UpdateMessageInteractionInfo                   *UpdateMessageInteractionInfo                   `json:"updateMessageInteractionInfo,omitempty"`
	UpdateMessageContentOpened                     *UpdateMessageContentOpened                     `json:"updateMessageContentOpened,omitempty"`
	UpdateMessageMentionRead                       *UpdateMessageMentionRead                       `json:"updateMessageMentionRead,omitempty"`
	UpdateMessageUnreadReactions                   *UpdateMessageUnreadReactions                   `json:"updateMessageUnreadReactions,omitempty"`
	UpdateMessageFactCheck                         *UpdateMessageFactCheck                         `json:"updateMessageFactCheck,omitempty"`
	UpdateMessageSuggestedPostInfo                 *UpdateMessageSuggestedPostInfo                 `json:"updateMessageSuggestedPostInfo,omitempty"`
	UpdateMessageLiveLocationViewed                *UpdateMessageLiveLocationViewed                `json:"updateMessageLiveLocationViewed,omitempty"`
	UpdateVideoPublished                           *UpdateVideoPublished                           `json:"updateVideoPublished,omitempty"`
	UpdateNewChat                                  *UpdateNewChat                                  `json:"updateNewChat,omitempty"`
	UpdateChatTitle                                *UpdateChatTitle                                `json:"updateChatTitle,omitempty"`
	UpdateChatPhoto                                *UpdateChatPhoto                                `json:"updateChatPhoto,omitempty"`
	UpdateChatAccentColors                         *UpdateChatAccentColors                         `json:"updateChatAccentColors,omitempty"`
	UpdateChatPermissions                          *UpdateChatPermissions                          `json:"updateChatPermissions,omitempty"`
	UpdateChatLastMessage                          *UpdateChatLastMessage                          `json:"updateChatLastMessage,omitempty"`
	UpdateChatPosition                             *UpdateChatPosition                             `json:"updateChatPosition,omitempty"`
	UpdateChatAddedToList                          *UpdateChatAddedToList                          `json:"updateChatAddedToList,omitempty"`
	UpdateChatRemovedFromList                      *UpdateChatRemovedFromList                      `json:"updateChatRemovedFromList,omitempty"`
	UpdateChatReadInbox                            *UpdateChatReadInbox                            `json:"updateChatReadInbox,omitempty"`
	UpdateChatReadOutbox                           *UpdateChatReadOutbox                           `json:"updateChatReadOutbox,omitempty"`
	UpdateChatActionBar                            *UpdateChatActionBar                            `json:"updateChatActionBar,omitempty"`
	UpdateChatBusinessBotManageBar                 *UpdateChatBusinessBotManageBar                 `json:"updateChatBusinessBotManageBar,omitempty"`
	UpdateChatAvailableReactions                   *UpdateChatAvailableReactions                   `json:"updateChatAvailableReactions,omitempty"`
	UpdateChatDraftMessage                         *UpdateChatDraftMessage                         `json:"updateChatDraftMessage,omitempty"`
	UpdateChatEmojiStatus                          *UpdateChatEmojiStatus                          `json:"updateChatEmojiStatus,omitempty"`
	UpdateChatMessageSender                        *UpdateChatMessageSender                        `json:"updateChatMessageSender,omitempty"`
	UpdateChatMessageAutoDeleteTime                *UpdateChatMessageAutoDeleteTime                `json:"updateChatMessageAutoDeleteTime,omitempty"`
	UpdateChatNotificationSettings                 *UpdateChatNotificationSettings                 `json:"updateChatNotificationSettings,omitempty"`
	UpdateChatPendingJoinRequests                  *UpdateChatPendingJoinRequests                  `json:"updateChatPendingJoinRequests,omitempty"`
	UpdateChatReplyMarkup                          *UpdateChatReplyMarkup                          `json:"updateChatReplyMarkup,omitempty"`
	UpdateChatBackground                           *UpdateChatBackground                           `json:"updateChatBackground,omitempty"`
	UpdateChatTheme                                *UpdateChatTheme                                `json:"updateChatTheme,omitempty"`
	UpdateChatUnreadMentionCount                   *UpdateChatUnreadMentionCount                   `json:"updateChatUnreadMentionCount,omitempty"`
	UpdateChatUnreadReactionCount                  *UpdateChatUnreadReactionCount                  `json:"updateChatUnreadReactionCount,omitempty"`
	UpdateChatVideoChat                            *UpdateChatVideoChat                            `json:"updateChatVideoChat,omitempty"`
	UpdateChatDefaultDisableNotification           *UpdateChatDefaultDisableNotification           `json:"updateChatDefaultDisableNotification,omitempty"`
	UpdateChatHasProtectedContent                  *UpdateChatHasProtectedContent                  `json:"updateChatHasProtectedContent,omitempty"`
	UpdateChatIsTranslatable                       *UpdateChatIsTranslatable                       `json:"updateChatIsTranslatable,omitempty"`
	UpdateChatIsMarkedAsUnread                     *UpdateChatIsMarkedAsUnread                     `json:"updateChatIsMarkedAsUnread,omitempty"`
	UpdateChatViewAsTopics                         *UpdateChatViewAsTopics                         `json:"updateChatViewAsTopics,omitempty"`
	UpdateChatBlockList                            *UpdateChatBlockList                            `json:"updateChatBlockList,omitempty"`
	UpdateChatHasScheduledMessages                 *UpdateChatHasScheduledMessages                 `json:"updateChatHasScheduledMessages,omitempty"`
	UpdateChatFolders                              *UpdateChatFolders                              `json:"updateChatFolders,omitempty"`
	UpdateChatOnlineMemberCount                    *UpdateChatOnlineMemberCount                    `json:"updateChatOnlineMemberCount,omitempty"`
	UpdateSavedMessagesTopic                       *UpdateSavedMessagesTopic                       `json:"updateSavedMessagesTopic,omitempty"`
	UpdateSavedMessagesTopicCount                  *UpdateSavedMessagesTopicCount                  `json:"updateSavedMessagesTopicCount,omitempty"`
	UpdateDirectMessagesChatTopic                  *UpdateDirectMessagesChatTopic                  `json:"updateDirectMessagesChatTopic,omitempty"`
	UpdateTopicMessageCount                        *UpdateTopicMessageCount                        `json:"updateTopicMessageCount,omitempty"`
	UpdateQuickReplyShortcut                       *UpdateQuickReplyShortcut                       `json:"updateQuickReplyShortcut,omitempty"`
	UpdateQuickReplyShortcutDeleted                *UpdateQuickReplyShortcutDeleted                `json:"updateQuickReplyShortcutDeleted,omitempty"`
	UpdateQuickReplyShortcuts                      *UpdateQuickReplyShortcuts                      `json:"updateQuickReplyShortcuts,omitempty"`
	UpdateQuickReplyShortcutMessages               *UpdateQuickReplyShortcutMessages               `json:"updateQuickReplyShortcutMessages,omitempty"`
	UpdateForumTopicInfo                           *UpdateForumTopicInfo                           `json:"updateForumTopicInfo,omitempty"`
	UpdateForumTopic                               *UpdateForumTopic                               `json:"updateForumTopic,omitempty"`
	UpdateScopeNotificationSettings                *UpdateScopeNotificationSettings                `json:"updateScopeNotificationSettings,omitempty"`
	UpdateReactionNotificationSettings             *UpdateReactionNotificationSettings             `json:"updateReactionNotificationSettings,omitempty"`
	UpdateNotification                             *UpdateNotification                             `json:"updateNotification,omitempty"`
	UpdateNotificationGroup                        *UpdateNotificationGroup                        `json:"updateNotificationGroup,omitempty"`
	UpdateActiveNotifications                      *UpdateActiveNotifications                      `json:"updateActiveNotifications,omitempty"`
	UpdateHavePendingNotifications                 *UpdateHavePendingNotifications                 `json:"updateHavePendingNotifications,omitempty"`
	UpdateDeleteMessages                           *UpdateDeleteMessages                           `json:"updateDeleteMessages,omitempty"`
	UpdateChatAction                               *UpdateChatAction                               `json:"updateChatAction,omitempty"`
	UpdatePendingTextMessage                       *UpdatePendingTextMessage                       `json:"updatePendingTextMessage,omitempty"`
	UpdateUserStatus                               *UpdateUserStatus                               `json:"updateUserStatus,omitempty"`
	UpdateUser                                     *UpdateUser                                     `json:"updateUser,omitempty"`
	UpdateBasicGroup                               *UpdateBasicGroup                               `json:"updateBasicGroup,omitempty"`
	UpdateSupergroup                               *UpdateSupergroup                               `json:"updateSupergroup,omitempty"`
	UpdateSecretChat                               *UpdateSecretChat                               `json:"updateSecretChat,omitempty"`
	UpdateUserFullInfo                             *UpdateUserFullInfo                             `json:"updateUserFullInfo,omitempty"`
	UpdateBasicGroupFullInfo                       *UpdateBasicGroupFullInfo                       `json:"updateBasicGroupFullInfo,omitempty"`
	UpdateSupergroupFullInfo                       *UpdateSupergroupFullInfo                       `json:"updateSupergroupFullInfo,omitempty"`
	UpdateServiceNotification                      *UpdateServiceNotification                      `json:"updateServiceNotification,omitempty"`
	UpdateFile                                     *UpdateFile                                     `json:"updateFile,omitempty"`
	UpdateFileGenerationStart                      *UpdateFileGenerationStart                      `json:"updateFileGenerationStart,omitempty"`
	UpdateFileGenerationStop                       *UpdateFileGenerationStop                       `json:"updateFileGenerationStop,omitempty"`
	UpdateFileDownloads                            *UpdateFileDownloads                            `json:"updateFileDownloads,omitempty"`
	UpdateFileAddedToDownloads                     *UpdateFileAddedToDownloads                     `json:"updateFileAddedToDownloads,omitempty"`
	UpdateFileDownload                             *UpdateFileDownload                             `json:"updateFileDownload,omitempty"`
	UpdateFileRemovedFromDownloads                 *UpdateFileRemovedFromDownloads                 `json:"updateFileRemovedFromDownloads,omitempty"`
	UpdateApplicationVerificationRequired          *UpdateApplicationVerificationRequired          `json:"updateApplicationVerificationRequired,omitempty"`
	UpdateApplicationRecaptchaVerificationRequired *UpdateApplicationRecaptchaVerificationRequired `json:"updateApplicationRecaptchaVerificationRequired,omitempty"`
	UpdateCall                                     *UpdateCall                                     `json:"updateCall,omitempty"`
	UpdateGroupCall                                *UpdateGroupCall                                `json:"updateGroupCall,omitempty"`
	UpdateGroupCallParticipant                     *UpdateGroupCallParticipant                     `json:"updateGroupCallParticipant,omitempty"`
	UpdateGroupCallParticipants                    *UpdateGroupCallParticipants                    `json:"updateGroupCallParticipants,omitempty"`
	UpdateGroupCallVerificationState               *UpdateGroupCallVerificationState               `json:"updateGroupCallVerificationState,omitempty"`
	UpdateNewGroupCallMessage                      *UpdateNewGroupCallMessage                      `json:"updateNewGroupCallMessage,omitempty"`
	UpdateNewGroupCallPaidReaction                 *UpdateNewGroupCallPaidReaction                 `json:"updateNewGroupCallPaidReaction,omitempty"`
	UpdateGroupCallMessageSendFailed               *UpdateGroupCallMessageSendFailed               `json:"updateGroupCallMessageSendFailed,omitempty"`
	UpdateGroupCallMessagesDeleted                 *UpdateGroupCallMessagesDeleted                 `json:"updateGroupCallMessagesDeleted,omitempty"`
	UpdateLiveStoryTopDonors                       *UpdateLiveStoryTopDonors                       `json:"updateLiveStoryTopDonors,omitempty"`
	UpdateNewCallSignalingData                     *UpdateNewCallSignalingData                     `json:"updateNewCallSignalingData,omitempty"`
	UpdateGiftAuctionState                         *UpdateGiftAuctionState                         `json:"updateGiftAuctionState,omitempty"`
	UpdateActiveGiftAuctions                       *UpdateActiveGiftAuctions                       `json:"updateActiveGiftAuctions,omitempty"`
	UpdateUserPrivacySettingRules                  *UpdateUserPrivacySettingRules                  `json:"updateUserPrivacySettingRules,omitempty"`
	UpdateUnreadMessageCount                       *UpdateUnreadMessageCount                       `json:"updateUnreadMessageCount,omitempty"`
	UpdateUnreadChatCount                          *UpdateUnreadChatCount                          `json:"updateUnreadChatCount,omitempty"`
	UpdateStory                                    *UpdateStory                                    `json:"updateStory,omitempty"`
	UpdateStoryDeleted                             *UpdateStoryDeleted                             `json:"updateStoryDeleted,omitempty"`
	UpdateStoryPostSucceeded                       *UpdateStoryPostSucceeded                       `json:"updateStoryPostSucceeded,omitempty"`
	UpdateStoryPostFailed                          *UpdateStoryPostFailed                          `json:"updateStoryPostFailed,omitempty"`
	UpdateChatActiveStories                        *UpdateChatActiveStories                        `json:"updateChatActiveStories,omitempty"`
	UpdateStoryListChatCount                       *UpdateStoryListChatCount                       `json:"updateStoryListChatCount,omitempty"`
	UpdateStoryStealthMode                         *UpdateStoryStealthMode                         `json:"updateStoryStealthMode,omitempty"`
	UpdateTrustedMiniAppBots                       *UpdateTrustedMiniAppBots                       `json:"updateTrustedMiniAppBots,omitempty"`
	UpdateOption                                   *UpdateOption                                   `json:"updateOption,omitempty"`
	UpdateStickerSet                               *UpdateStickerSet                               `json:"updateStickerSet,omitempty"`
	UpdateInstalledStickerSets                     *UpdateInstalledStickerSets                     `json:"updateInstalledStickerSets,omitempty"`
	UpdateTrendingStickerSets                      *UpdateTrendingStickerSets                      `json:"updateTrendingStickerSets,omitempty"`
	UpdateRecentStickers                           *UpdateRecentStickers                           `json:"updateRecentStickers,omitempty"`
	UpdateFavoriteStickers                         *UpdateFavoriteStickers                         `json:"updateFavoriteStickers,omitempty"`
	UpdateSavedAnimations                          *UpdateSavedAnimations                          `json:"updateSavedAnimations,omitempty"`
	UpdateSavedNotificationSounds                  *UpdateSavedNotificationSounds                  `json:"updateSavedNotificationSounds,omitempty"`
	UpdateDefaultBackground                        *UpdateDefaultBackground                        `json:"updateDefaultBackground,omitempty"`
	UpdateEmojiChatThemes                          *UpdateEmojiChatThemes                          `json:"updateEmojiChatThemes,omitempty"`
	UpdateAccentColors                             *UpdateAccentColors                             `json:"updateAccentColors,omitempty"`
	UpdateProfileAccentColors                      *UpdateProfileAccentColors                      `json:"updateProfileAccentColors,omitempty"`
	UpdateLanguagePackStrings                      *UpdateLanguagePackStrings                      `json:"updateLanguagePackStrings,omitempty"`
	UpdateConnectionState                          *UpdateConnectionState                          `json:"updateConnectionState,omitempty"`
	UpdateFreezeState                              *UpdateFreezeState                              `json:"updateFreezeState,omitempty"`
	UpdateAgeVerificationParameters                *UpdateAgeVerificationParameters                `json:"updateAgeVerificationParameters,omitempty"`
	UpdateTermsOfService                           *UpdateTermsOfService                           `json:"updateTermsOfService,omitempty"`
	UpdateUnconfirmedSession                       *UpdateUnconfirmedSession                       `json:"updateUnconfirmedSession,omitempty"`
	UpdateAttachmentMenuBots                       *UpdateAttachmentMenuBots                       `json:"updateAttachmentMenuBots,omitempty"`
	UpdateWebAppMessageSent                        *UpdateWebAppMessageSent                        `json:"updateWebAppMessageSent,omitempty"`
	UpdateActiveEmojiReactions                     *UpdateActiveEmojiReactions                     `json:"updateActiveEmojiReactions,omitempty"`
	UpdateAvailableMessageEffects                  *UpdateAvailableMessageEffects                  `json:"updateAvailableMessageEffects,omitempty"`
	UpdateDefaultReactionType                      *UpdateDefaultReactionType                      `json:"updateDefaultReactionType,omitempty"`
	UpdateDefaultPaidReactionType                  *UpdateDefaultPaidReactionType                  `json:"updateDefaultPaidReactionType,omitempty"`
	UpdateSavedMessagesTags                        *UpdateSavedMessagesTags                        `json:"updateSavedMessagesTags,omitempty"`
	UpdateActiveLiveLocationMessages               *UpdateActiveLiveLocationMessages               `json:"updateActiveLiveLocationMessages,omitempty"`
	UpdateOwnedStarCount                           *UpdateOwnedStarCount                           `json:"updateOwnedStarCount,omitempty"`
	UpdateOwnedTonCount                            *UpdateOwnedTonCount                            `json:"updateOwnedTonCount,omitempty"`
	UpdateChatRevenueAmount                        *UpdateChatRevenueAmount                        `json:"updateChatRevenueAmount,omitempty"`
	UpdateStarRevenueStatus                        *UpdateStarRevenueStatus                        `json:"updateStarRevenueStatus,omitempty"`
	UpdateTonRevenueStatus                         *UpdateTonRevenueStatus                         `json:"updateTonRevenueStatus,omitempty"`
	UpdateSpeechRecognitionTrial                   *UpdateSpeechRecognitionTrial                   `json:"updateSpeechRecognitionTrial,omitempty"`
	UpdateGroupCallMessageLevels                   *UpdateGroupCallMessageLevels                   `json:"updateGroupCallMessageLevels,omitempty"`
	UpdateDiceEmojis                               *UpdateDiceEmojis                               `json:"updateDiceEmojis,omitempty"`
	UpdateStakeDiceState                           *UpdateStakeDiceState                           `json:"updateStakeDiceState,omitempty"`
	UpdateAnimatedEmojiMessageClicked              *UpdateAnimatedEmojiMessageClicked              `json:"updateAnimatedEmojiMessageClicked,omitempty"`
	UpdateAnimationSearchParameters                *UpdateAnimationSearchParameters                `json:"updateAnimationSearchParameters,omitempty"`
	UpdateSuggestedActions                         *UpdateSuggestedActions                         `json:"updateSuggestedActions,omitempty"`
	UpdateSpeedLimitNotification                   *UpdateSpeedLimitNotification                   `json:"updateSpeedLimitNotification,omitempty"`
	UpdateContactCloseBirthdays                    *UpdateContactCloseBirthdays                    `json:"updateContactCloseBirthdays,omitempty"`
	UpdateAutosaveSettings                         *UpdateAutosaveSettings                         `json:"updateAutosaveSettings,omitempty"`
	UpdateBusinessConnection                       *UpdateBusinessConnection                       `json:"updateBusinessConnection,omitempty"`
	UpdateNewBusinessMessage                       *UpdateNewBusinessMessage                       `json:"updateNewBusinessMessage,omitempty"`
	UpdateBusinessMessageEdited                    *UpdateBusinessMessageEdited                    `json:"updateBusinessMessageEdited,omitempty"`
	UpdateBusinessMessagesDeleted                  *UpdateBusinessMessagesDeleted                  `json:"updateBusinessMessagesDeleted,omitempty"`
	UpdateNewInlineQuery                           *UpdateNewInlineQuery                           `json:"updateNewInlineQuery,omitempty"`
	UpdateNewChosenInlineResult                    *UpdateNewChosenInlineResult                    `json:"updateNewChosenInlineResult,omitempty"`
	UpdateNewCallbackQuery                         *UpdateNewCallbackQuery                         `json:"updateNewCallbackQuery,omitempty"`
	UpdateNewInlineCallbackQuery                   *UpdateNewInlineCallbackQuery                   `json:"updateNewInlineCallbackQuery,omitempty"`
	UpdateNewBusinessCallbackQuery                 *UpdateNewBusinessCallbackQuery                 `json:"updateNewBusinessCallbackQuery,omitempty"`
	UpdateNewShippingQuery                         *UpdateNewShippingQuery                         `json:"updateNewShippingQuery,omitempty"`
	UpdateNewPreCheckoutQuery                      *UpdateNewPreCheckoutQuery                      `json:"updateNewPreCheckoutQuery,omitempty"`
	UpdateNewCustomEvent                           *UpdateNewCustomEvent                           `json:"updateNewCustomEvent,omitempty"`
	UpdateNewCustomQuery                           *UpdateNewCustomQuery                           `json:"updateNewCustomQuery,omitempty"`
	UpdatePoll                                     *UpdatePoll                                     `json:"updatePoll,omitempty"`
	UpdatePollAnswer                               *UpdatePollAnswer                               `json:"updatePollAnswer,omitempty"`
	UpdateChatMember                               *UpdateChatMember                               `json:"updateChatMember,omitempty"`
	UpdateNewChatJoinRequest                       *UpdateNewChatJoinRequest                       `json:"updateNewChatJoinRequest,omitempty"`
	UpdateChatBoost                                *UpdateChatBoost                                `json:"updateChatBoost,omitempty"`
	UpdateMessageReaction                          *UpdateMessageReaction                          `json:"updateMessageReaction,omitempty"`
	UpdateMessageReactions                         *UpdateMessageReactions                         `json:"updateMessageReactions,omitempty"`
	UpdatePaidMediaPurchased                       *UpdatePaidMediaPurchased                       `json:"updatePaidMediaPurchased,omitempty"`
}

func (t *Update) Type() string {
	return ""
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
	switch typeObj.Type {
	case "updateAuthorizationState":
		t.UpdateAuthorizationState = new(UpdateAuthorizationState)
		return json.Unmarshal(b, t.UpdateAuthorizationState)
	case "updateNewMessage":
		t.UpdateNewMessage = new(UpdateNewMessage)
		return json.Unmarshal(b, t.UpdateNewMessage)
	case "updateMessageSendAcknowledged":
		t.UpdateMessageSendAcknowledged = new(UpdateMessageSendAcknowledged)
		return json.Unmarshal(b, t.UpdateMessageSendAcknowledged)
	case "updateMessageSendSucceeded":
		t.UpdateMessageSendSucceeded = new(UpdateMessageSendSucceeded)
		return json.Unmarshal(b, t.UpdateMessageSendSucceeded)
	case "updateMessageSendFailed":
		t.UpdateMessageSendFailed = new(UpdateMessageSendFailed)
		return json.Unmarshal(b, t.UpdateMessageSendFailed)
	case "updateMessageContent":
		t.UpdateMessageContent = new(UpdateMessageContent)
		return json.Unmarshal(b, t.UpdateMessageContent)
	case "updateMessageEdited":
		t.UpdateMessageEdited = new(UpdateMessageEdited)
		return json.Unmarshal(b, t.UpdateMessageEdited)
	case "updateMessageIsPinned":
		t.UpdateMessageIsPinned = new(UpdateMessageIsPinned)
		return json.Unmarshal(b, t.UpdateMessageIsPinned)
	case "updateMessageInteractionInfo":
		t.UpdateMessageInteractionInfo = new(UpdateMessageInteractionInfo)
		return json.Unmarshal(b, t.UpdateMessageInteractionInfo)
	case "updateMessageContentOpened":
		t.UpdateMessageContentOpened = new(UpdateMessageContentOpened)
		return json.Unmarshal(b, t.UpdateMessageContentOpened)
	case "updateMessageMentionRead":
		t.UpdateMessageMentionRead = new(UpdateMessageMentionRead)
		return json.Unmarshal(b, t.UpdateMessageMentionRead)
	case "updateMessageUnreadReactions":
		t.UpdateMessageUnreadReactions = new(UpdateMessageUnreadReactions)
		return json.Unmarshal(b, t.UpdateMessageUnreadReactions)
	case "updateMessageFactCheck":
		t.UpdateMessageFactCheck = new(UpdateMessageFactCheck)
		return json.Unmarshal(b, t.UpdateMessageFactCheck)
	case "updateMessageSuggestedPostInfo":
		t.UpdateMessageSuggestedPostInfo = new(UpdateMessageSuggestedPostInfo)
		return json.Unmarshal(b, t.UpdateMessageSuggestedPostInfo)
	case "updateMessageLiveLocationViewed":
		t.UpdateMessageLiveLocationViewed = new(UpdateMessageLiveLocationViewed)
		return json.Unmarshal(b, t.UpdateMessageLiveLocationViewed)
	case "updateVideoPublished":
		t.UpdateVideoPublished = new(UpdateVideoPublished)
		return json.Unmarshal(b, t.UpdateVideoPublished)
	case "updateNewChat":
		t.UpdateNewChat = new(UpdateNewChat)
		return json.Unmarshal(b, t.UpdateNewChat)
	case "updateChatTitle":
		t.UpdateChatTitle = new(UpdateChatTitle)
		return json.Unmarshal(b, t.UpdateChatTitle)
	case "updateChatPhoto":
		t.UpdateChatPhoto = new(UpdateChatPhoto)
		return json.Unmarshal(b, t.UpdateChatPhoto)
	case "updateChatAccentColors":
		t.UpdateChatAccentColors = new(UpdateChatAccentColors)
		return json.Unmarshal(b, t.UpdateChatAccentColors)
	case "updateChatPermissions":
		t.UpdateChatPermissions = new(UpdateChatPermissions)
		return json.Unmarshal(b, t.UpdateChatPermissions)
	case "updateChatLastMessage":
		t.UpdateChatLastMessage = new(UpdateChatLastMessage)
		return json.Unmarshal(b, t.UpdateChatLastMessage)
	case "updateChatPosition":
		t.UpdateChatPosition = new(UpdateChatPosition)
		return json.Unmarshal(b, t.UpdateChatPosition)
	case "updateChatAddedToList":
		t.UpdateChatAddedToList = new(UpdateChatAddedToList)
		return json.Unmarshal(b, t.UpdateChatAddedToList)
	case "updateChatRemovedFromList":
		t.UpdateChatRemovedFromList = new(UpdateChatRemovedFromList)
		return json.Unmarshal(b, t.UpdateChatRemovedFromList)
	case "updateChatReadInbox":
		t.UpdateChatReadInbox = new(UpdateChatReadInbox)
		return json.Unmarshal(b, t.UpdateChatReadInbox)
	case "updateChatReadOutbox":
		t.UpdateChatReadOutbox = new(UpdateChatReadOutbox)
		return json.Unmarshal(b, t.UpdateChatReadOutbox)
	case "updateChatActionBar":
		t.UpdateChatActionBar = new(UpdateChatActionBar)
		return json.Unmarshal(b, t.UpdateChatActionBar)
	case "updateChatBusinessBotManageBar":
		t.UpdateChatBusinessBotManageBar = new(UpdateChatBusinessBotManageBar)
		return json.Unmarshal(b, t.UpdateChatBusinessBotManageBar)
	case "updateChatAvailableReactions":
		t.UpdateChatAvailableReactions = new(UpdateChatAvailableReactions)
		return json.Unmarshal(b, t.UpdateChatAvailableReactions)
	case "updateChatDraftMessage":
		t.UpdateChatDraftMessage = new(UpdateChatDraftMessage)
		return json.Unmarshal(b, t.UpdateChatDraftMessage)
	case "updateChatEmojiStatus":
		t.UpdateChatEmojiStatus = new(UpdateChatEmojiStatus)
		return json.Unmarshal(b, t.UpdateChatEmojiStatus)
	case "updateChatMessageSender":
		t.UpdateChatMessageSender = new(UpdateChatMessageSender)
		return json.Unmarshal(b, t.UpdateChatMessageSender)
	case "updateChatMessageAutoDeleteTime":
		t.UpdateChatMessageAutoDeleteTime = new(UpdateChatMessageAutoDeleteTime)
		return json.Unmarshal(b, t.UpdateChatMessageAutoDeleteTime)
	case "updateChatNotificationSettings":
		t.UpdateChatNotificationSettings = new(UpdateChatNotificationSettings)
		return json.Unmarshal(b, t.UpdateChatNotificationSettings)
	case "updateChatPendingJoinRequests":
		t.UpdateChatPendingJoinRequests = new(UpdateChatPendingJoinRequests)
		return json.Unmarshal(b, t.UpdateChatPendingJoinRequests)
	case "updateChatReplyMarkup":
		t.UpdateChatReplyMarkup = new(UpdateChatReplyMarkup)
		return json.Unmarshal(b, t.UpdateChatReplyMarkup)
	case "updateChatBackground":
		t.UpdateChatBackground = new(UpdateChatBackground)
		return json.Unmarshal(b, t.UpdateChatBackground)
	case "updateChatTheme":
		t.UpdateChatTheme = new(UpdateChatTheme)
		return json.Unmarshal(b, t.UpdateChatTheme)
	case "updateChatUnreadMentionCount":
		t.UpdateChatUnreadMentionCount = new(UpdateChatUnreadMentionCount)
		return json.Unmarshal(b, t.UpdateChatUnreadMentionCount)
	case "updateChatUnreadReactionCount":
		t.UpdateChatUnreadReactionCount = new(UpdateChatUnreadReactionCount)
		return json.Unmarshal(b, t.UpdateChatUnreadReactionCount)
	case "updateChatVideoChat":
		t.UpdateChatVideoChat = new(UpdateChatVideoChat)
		return json.Unmarshal(b, t.UpdateChatVideoChat)
	case "updateChatDefaultDisableNotification":
		t.UpdateChatDefaultDisableNotification = new(UpdateChatDefaultDisableNotification)
		return json.Unmarshal(b, t.UpdateChatDefaultDisableNotification)
	case "updateChatHasProtectedContent":
		t.UpdateChatHasProtectedContent = new(UpdateChatHasProtectedContent)
		return json.Unmarshal(b, t.UpdateChatHasProtectedContent)
	case "updateChatIsTranslatable":
		t.UpdateChatIsTranslatable = new(UpdateChatIsTranslatable)
		return json.Unmarshal(b, t.UpdateChatIsTranslatable)
	case "updateChatIsMarkedAsUnread":
		t.UpdateChatIsMarkedAsUnread = new(UpdateChatIsMarkedAsUnread)
		return json.Unmarshal(b, t.UpdateChatIsMarkedAsUnread)
	case "updateChatViewAsTopics":
		t.UpdateChatViewAsTopics = new(UpdateChatViewAsTopics)
		return json.Unmarshal(b, t.UpdateChatViewAsTopics)
	case "updateChatBlockList":
		t.UpdateChatBlockList = new(UpdateChatBlockList)
		return json.Unmarshal(b, t.UpdateChatBlockList)
	case "updateChatHasScheduledMessages":
		t.UpdateChatHasScheduledMessages = new(UpdateChatHasScheduledMessages)
		return json.Unmarshal(b, t.UpdateChatHasScheduledMessages)
	case "updateChatFolders":
		t.UpdateChatFolders = new(UpdateChatFolders)
		return json.Unmarshal(b, t.UpdateChatFolders)
	case "updateChatOnlineMemberCount":
		t.UpdateChatOnlineMemberCount = new(UpdateChatOnlineMemberCount)
		return json.Unmarshal(b, t.UpdateChatOnlineMemberCount)
	case "updateSavedMessagesTopic":
		t.UpdateSavedMessagesTopic = new(UpdateSavedMessagesTopic)
		return json.Unmarshal(b, t.UpdateSavedMessagesTopic)
	case "updateSavedMessagesTopicCount":
		t.UpdateSavedMessagesTopicCount = new(UpdateSavedMessagesTopicCount)
		return json.Unmarshal(b, t.UpdateSavedMessagesTopicCount)
	case "updateDirectMessagesChatTopic":
		t.UpdateDirectMessagesChatTopic = new(UpdateDirectMessagesChatTopic)
		return json.Unmarshal(b, t.UpdateDirectMessagesChatTopic)
	case "updateTopicMessageCount":
		t.UpdateTopicMessageCount = new(UpdateTopicMessageCount)
		return json.Unmarshal(b, t.UpdateTopicMessageCount)
	case "updateQuickReplyShortcut":
		t.UpdateQuickReplyShortcut = new(UpdateQuickReplyShortcut)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcut)
	case "updateQuickReplyShortcutDeleted":
		t.UpdateQuickReplyShortcutDeleted = new(UpdateQuickReplyShortcutDeleted)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcutDeleted)
	case "updateQuickReplyShortcuts":
		t.UpdateQuickReplyShortcuts = new(UpdateQuickReplyShortcuts)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcuts)
	case "updateQuickReplyShortcutMessages":
		t.UpdateQuickReplyShortcutMessages = new(UpdateQuickReplyShortcutMessages)
		return json.Unmarshal(b, t.UpdateQuickReplyShortcutMessages)
	case "updateForumTopicInfo":
		t.UpdateForumTopicInfo = new(UpdateForumTopicInfo)
		return json.Unmarshal(b, t.UpdateForumTopicInfo)
	case "updateForumTopic":
		t.UpdateForumTopic = new(UpdateForumTopic)
		return json.Unmarshal(b, t.UpdateForumTopic)
	case "updateScopeNotificationSettings":
		t.UpdateScopeNotificationSettings = new(UpdateScopeNotificationSettings)
		return json.Unmarshal(b, t.UpdateScopeNotificationSettings)
	case "updateReactionNotificationSettings":
		t.UpdateReactionNotificationSettings = new(UpdateReactionNotificationSettings)
		return json.Unmarshal(b, t.UpdateReactionNotificationSettings)
	case "updateNotification":
		t.UpdateNotification = new(UpdateNotification)
		return json.Unmarshal(b, t.UpdateNotification)
	case "updateNotificationGroup":
		t.UpdateNotificationGroup = new(UpdateNotificationGroup)
		return json.Unmarshal(b, t.UpdateNotificationGroup)
	case "updateActiveNotifications":
		t.UpdateActiveNotifications = new(UpdateActiveNotifications)
		return json.Unmarshal(b, t.UpdateActiveNotifications)
	case "updateHavePendingNotifications":
		t.UpdateHavePendingNotifications = new(UpdateHavePendingNotifications)
		return json.Unmarshal(b, t.UpdateHavePendingNotifications)
	case "updateDeleteMessages":
		t.UpdateDeleteMessages = new(UpdateDeleteMessages)
		return json.Unmarshal(b, t.UpdateDeleteMessages)
	case "updateChatAction":
		t.UpdateChatAction = new(UpdateChatAction)
		return json.Unmarshal(b, t.UpdateChatAction)
	case "updatePendingTextMessage":
		t.UpdatePendingTextMessage = new(UpdatePendingTextMessage)
		return json.Unmarshal(b, t.UpdatePendingTextMessage)
	case "updateUserStatus":
		t.UpdateUserStatus = new(UpdateUserStatus)
		return json.Unmarshal(b, t.UpdateUserStatus)
	case "updateUser":
		t.UpdateUser = new(UpdateUser)
		return json.Unmarshal(b, t.UpdateUser)
	case "updateBasicGroup":
		t.UpdateBasicGroup = new(UpdateBasicGroup)
		return json.Unmarshal(b, t.UpdateBasicGroup)
	case "updateSupergroup":
		t.UpdateSupergroup = new(UpdateSupergroup)
		return json.Unmarshal(b, t.UpdateSupergroup)
	case "updateSecretChat":
		t.UpdateSecretChat = new(UpdateSecretChat)
		return json.Unmarshal(b, t.UpdateSecretChat)
	case "updateUserFullInfo":
		t.UpdateUserFullInfo = new(UpdateUserFullInfo)
		return json.Unmarshal(b, t.UpdateUserFullInfo)
	case "updateBasicGroupFullInfo":
		t.UpdateBasicGroupFullInfo = new(UpdateBasicGroupFullInfo)
		return json.Unmarshal(b, t.UpdateBasicGroupFullInfo)
	case "updateSupergroupFullInfo":
		t.UpdateSupergroupFullInfo = new(UpdateSupergroupFullInfo)
		return json.Unmarshal(b, t.UpdateSupergroupFullInfo)
	case "updateServiceNotification":
		t.UpdateServiceNotification = new(UpdateServiceNotification)
		return json.Unmarshal(b, t.UpdateServiceNotification)
	case "updateFile":
		t.UpdateFile = new(UpdateFile)
		return json.Unmarshal(b, t.UpdateFile)
	case "updateFileGenerationStart":
		t.UpdateFileGenerationStart = new(UpdateFileGenerationStart)
		return json.Unmarshal(b, t.UpdateFileGenerationStart)
	case "updateFileGenerationStop":
		t.UpdateFileGenerationStop = new(UpdateFileGenerationStop)
		return json.Unmarshal(b, t.UpdateFileGenerationStop)
	case "updateFileDownloads":
		t.UpdateFileDownloads = new(UpdateFileDownloads)
		return json.Unmarshal(b, t.UpdateFileDownloads)
	case "updateFileAddedToDownloads":
		t.UpdateFileAddedToDownloads = new(UpdateFileAddedToDownloads)
		return json.Unmarshal(b, t.UpdateFileAddedToDownloads)
	case "updateFileDownload":
		t.UpdateFileDownload = new(UpdateFileDownload)
		return json.Unmarshal(b, t.UpdateFileDownload)
	case "updateFileRemovedFromDownloads":
		t.UpdateFileRemovedFromDownloads = new(UpdateFileRemovedFromDownloads)
		return json.Unmarshal(b, t.UpdateFileRemovedFromDownloads)
	case "updateApplicationVerificationRequired":
		t.UpdateApplicationVerificationRequired = new(UpdateApplicationVerificationRequired)
		return json.Unmarshal(b, t.UpdateApplicationVerificationRequired)
	case "updateApplicationRecaptchaVerificationRequired":
		t.UpdateApplicationRecaptchaVerificationRequired = new(UpdateApplicationRecaptchaVerificationRequired)
		return json.Unmarshal(b, t.UpdateApplicationRecaptchaVerificationRequired)
	case "updateCall":
		t.UpdateCall = new(UpdateCall)
		return json.Unmarshal(b, t.UpdateCall)
	case "updateGroupCall":
		t.UpdateGroupCall = new(UpdateGroupCall)
		return json.Unmarshal(b, t.UpdateGroupCall)
	case "updateGroupCallParticipant":
		t.UpdateGroupCallParticipant = new(UpdateGroupCallParticipant)
		return json.Unmarshal(b, t.UpdateGroupCallParticipant)
	case "updateGroupCallParticipants":
		t.UpdateGroupCallParticipants = new(UpdateGroupCallParticipants)
		return json.Unmarshal(b, t.UpdateGroupCallParticipants)
	case "updateGroupCallVerificationState":
		t.UpdateGroupCallVerificationState = new(UpdateGroupCallVerificationState)
		return json.Unmarshal(b, t.UpdateGroupCallVerificationState)
	case "updateNewGroupCallMessage":
		t.UpdateNewGroupCallMessage = new(UpdateNewGroupCallMessage)
		return json.Unmarshal(b, t.UpdateNewGroupCallMessage)
	case "updateNewGroupCallPaidReaction":
		t.UpdateNewGroupCallPaidReaction = new(UpdateNewGroupCallPaidReaction)
		return json.Unmarshal(b, t.UpdateNewGroupCallPaidReaction)
	case "updateGroupCallMessageSendFailed":
		t.UpdateGroupCallMessageSendFailed = new(UpdateGroupCallMessageSendFailed)
		return json.Unmarshal(b, t.UpdateGroupCallMessageSendFailed)
	case "updateGroupCallMessagesDeleted":
		t.UpdateGroupCallMessagesDeleted = new(UpdateGroupCallMessagesDeleted)
		return json.Unmarshal(b, t.UpdateGroupCallMessagesDeleted)
	case "updateLiveStoryTopDonors":
		t.UpdateLiveStoryTopDonors = new(UpdateLiveStoryTopDonors)
		return json.Unmarshal(b, t.UpdateLiveStoryTopDonors)
	case "updateNewCallSignalingData":
		t.UpdateNewCallSignalingData = new(UpdateNewCallSignalingData)
		return json.Unmarshal(b, t.UpdateNewCallSignalingData)
	case "updateGiftAuctionState":
		t.UpdateGiftAuctionState = new(UpdateGiftAuctionState)
		return json.Unmarshal(b, t.UpdateGiftAuctionState)
	case "updateActiveGiftAuctions":
		t.UpdateActiveGiftAuctions = new(UpdateActiveGiftAuctions)
		return json.Unmarshal(b, t.UpdateActiveGiftAuctions)
	case "updateUserPrivacySettingRules":
		t.UpdateUserPrivacySettingRules = new(UpdateUserPrivacySettingRules)
		return json.Unmarshal(b, t.UpdateUserPrivacySettingRules)
	case "updateUnreadMessageCount":
		t.UpdateUnreadMessageCount = new(UpdateUnreadMessageCount)
		return json.Unmarshal(b, t.UpdateUnreadMessageCount)
	case "updateUnreadChatCount":
		t.UpdateUnreadChatCount = new(UpdateUnreadChatCount)
		return json.Unmarshal(b, t.UpdateUnreadChatCount)
	case "updateStory":
		t.UpdateStory = new(UpdateStory)
		return json.Unmarshal(b, t.UpdateStory)
	case "updateStoryDeleted":
		t.UpdateStoryDeleted = new(UpdateStoryDeleted)
		return json.Unmarshal(b, t.UpdateStoryDeleted)
	case "updateStoryPostSucceeded":
		t.UpdateStoryPostSucceeded = new(UpdateStoryPostSucceeded)
		return json.Unmarshal(b, t.UpdateStoryPostSucceeded)
	case "updateStoryPostFailed":
		t.UpdateStoryPostFailed = new(UpdateStoryPostFailed)
		return json.Unmarshal(b, t.UpdateStoryPostFailed)
	case "updateChatActiveStories":
		t.UpdateChatActiveStories = new(UpdateChatActiveStories)
		return json.Unmarshal(b, t.UpdateChatActiveStories)
	case "updateStoryListChatCount":
		t.UpdateStoryListChatCount = new(UpdateStoryListChatCount)
		return json.Unmarshal(b, t.UpdateStoryListChatCount)
	case "updateStoryStealthMode":
		t.UpdateStoryStealthMode = new(UpdateStoryStealthMode)
		return json.Unmarshal(b, t.UpdateStoryStealthMode)
	case "updateTrustedMiniAppBots":
		t.UpdateTrustedMiniAppBots = new(UpdateTrustedMiniAppBots)
		return json.Unmarshal(b, t.UpdateTrustedMiniAppBots)
	case "updateOption":
		t.UpdateOption = new(UpdateOption)
		return json.Unmarshal(b, t.UpdateOption)
	case "updateStickerSet":
		t.UpdateStickerSet = new(UpdateStickerSet)
		return json.Unmarshal(b, t.UpdateStickerSet)
	case "updateInstalledStickerSets":
		t.UpdateInstalledStickerSets = new(UpdateInstalledStickerSets)
		return json.Unmarshal(b, t.UpdateInstalledStickerSets)
	case "updateTrendingStickerSets":
		t.UpdateTrendingStickerSets = new(UpdateTrendingStickerSets)
		return json.Unmarshal(b, t.UpdateTrendingStickerSets)
	case "updateRecentStickers":
		t.UpdateRecentStickers = new(UpdateRecentStickers)
		return json.Unmarshal(b, t.UpdateRecentStickers)
	case "updateFavoriteStickers":
		t.UpdateFavoriteStickers = new(UpdateFavoriteStickers)
		return json.Unmarshal(b, t.UpdateFavoriteStickers)
	case "updateSavedAnimations":
		t.UpdateSavedAnimations = new(UpdateSavedAnimations)
		return json.Unmarshal(b, t.UpdateSavedAnimations)
	case "updateSavedNotificationSounds":
		t.UpdateSavedNotificationSounds = new(UpdateSavedNotificationSounds)
		return json.Unmarshal(b, t.UpdateSavedNotificationSounds)
	case "updateDefaultBackground":
		t.UpdateDefaultBackground = new(UpdateDefaultBackground)
		return json.Unmarshal(b, t.UpdateDefaultBackground)
	case "updateEmojiChatThemes":
		t.UpdateEmojiChatThemes = new(UpdateEmojiChatThemes)
		return json.Unmarshal(b, t.UpdateEmojiChatThemes)
	case "updateAccentColors":
		t.UpdateAccentColors = new(UpdateAccentColors)
		return json.Unmarshal(b, t.UpdateAccentColors)
	case "updateProfileAccentColors":
		t.UpdateProfileAccentColors = new(UpdateProfileAccentColors)
		return json.Unmarshal(b, t.UpdateProfileAccentColors)
	case "updateLanguagePackStrings":
		t.UpdateLanguagePackStrings = new(UpdateLanguagePackStrings)
		return json.Unmarshal(b, t.UpdateLanguagePackStrings)
	case "updateConnectionState":
		t.UpdateConnectionState = new(UpdateConnectionState)
		return json.Unmarshal(b, t.UpdateConnectionState)
	case "updateFreezeState":
		t.UpdateFreezeState = new(UpdateFreezeState)
		return json.Unmarshal(b, t.UpdateFreezeState)
	case "updateAgeVerificationParameters":
		t.UpdateAgeVerificationParameters = new(UpdateAgeVerificationParameters)
		return json.Unmarshal(b, t.UpdateAgeVerificationParameters)
	case "updateTermsOfService":
		t.UpdateTermsOfService = new(UpdateTermsOfService)
		return json.Unmarshal(b, t.UpdateTermsOfService)
	case "updateUnconfirmedSession":
		t.UpdateUnconfirmedSession = new(UpdateUnconfirmedSession)
		return json.Unmarshal(b, t.UpdateUnconfirmedSession)
	case "updateAttachmentMenuBots":
		t.UpdateAttachmentMenuBots = new(UpdateAttachmentMenuBots)
		return json.Unmarshal(b, t.UpdateAttachmentMenuBots)
	case "updateWebAppMessageSent":
		t.UpdateWebAppMessageSent = new(UpdateWebAppMessageSent)
		return json.Unmarshal(b, t.UpdateWebAppMessageSent)
	case "updateActiveEmojiReactions":
		t.UpdateActiveEmojiReactions = new(UpdateActiveEmojiReactions)
		return json.Unmarshal(b, t.UpdateActiveEmojiReactions)
	case "updateAvailableMessageEffects":
		t.UpdateAvailableMessageEffects = new(UpdateAvailableMessageEffects)
		return json.Unmarshal(b, t.UpdateAvailableMessageEffects)
	case "updateDefaultReactionType":
		t.UpdateDefaultReactionType = new(UpdateDefaultReactionType)
		return json.Unmarshal(b, t.UpdateDefaultReactionType)
	case "updateDefaultPaidReactionType":
		t.UpdateDefaultPaidReactionType = new(UpdateDefaultPaidReactionType)
		return json.Unmarshal(b, t.UpdateDefaultPaidReactionType)
	case "updateSavedMessagesTags":
		t.UpdateSavedMessagesTags = new(UpdateSavedMessagesTags)
		return json.Unmarshal(b, t.UpdateSavedMessagesTags)
	case "updateActiveLiveLocationMessages":
		t.UpdateActiveLiveLocationMessages = new(UpdateActiveLiveLocationMessages)
		return json.Unmarshal(b, t.UpdateActiveLiveLocationMessages)
	case "updateOwnedStarCount":
		t.UpdateOwnedStarCount = new(UpdateOwnedStarCount)
		return json.Unmarshal(b, t.UpdateOwnedStarCount)
	case "updateOwnedTonCount":
		t.UpdateOwnedTonCount = new(UpdateOwnedTonCount)
		return json.Unmarshal(b, t.UpdateOwnedTonCount)
	case "updateChatRevenueAmount":
		t.UpdateChatRevenueAmount = new(UpdateChatRevenueAmount)
		return json.Unmarshal(b, t.UpdateChatRevenueAmount)
	case "updateStarRevenueStatus":
		t.UpdateStarRevenueStatus = new(UpdateStarRevenueStatus)
		return json.Unmarshal(b, t.UpdateStarRevenueStatus)
	case "updateTonRevenueStatus":
		t.UpdateTonRevenueStatus = new(UpdateTonRevenueStatus)
		return json.Unmarshal(b, t.UpdateTonRevenueStatus)
	case "updateSpeechRecognitionTrial":
		t.UpdateSpeechRecognitionTrial = new(UpdateSpeechRecognitionTrial)
		return json.Unmarshal(b, t.UpdateSpeechRecognitionTrial)
	case "updateGroupCallMessageLevels":
		t.UpdateGroupCallMessageLevels = new(UpdateGroupCallMessageLevels)
		return json.Unmarshal(b, t.UpdateGroupCallMessageLevels)
	case "updateDiceEmojis":
		t.UpdateDiceEmojis = new(UpdateDiceEmojis)
		return json.Unmarshal(b, t.UpdateDiceEmojis)
	case "updateStakeDiceState":
		t.UpdateStakeDiceState = new(UpdateStakeDiceState)
		return json.Unmarshal(b, t.UpdateStakeDiceState)
	case "updateAnimatedEmojiMessageClicked":
		t.UpdateAnimatedEmojiMessageClicked = new(UpdateAnimatedEmojiMessageClicked)
		return json.Unmarshal(b, t.UpdateAnimatedEmojiMessageClicked)
	case "updateAnimationSearchParameters":
		t.UpdateAnimationSearchParameters = new(UpdateAnimationSearchParameters)
		return json.Unmarshal(b, t.UpdateAnimationSearchParameters)
	case "updateSuggestedActions":
		t.UpdateSuggestedActions = new(UpdateSuggestedActions)
		return json.Unmarshal(b, t.UpdateSuggestedActions)
	case "updateSpeedLimitNotification":
		t.UpdateSpeedLimitNotification = new(UpdateSpeedLimitNotification)
		return json.Unmarshal(b, t.UpdateSpeedLimitNotification)
	case "updateContactCloseBirthdays":
		t.UpdateContactCloseBirthdays = new(UpdateContactCloseBirthdays)
		return json.Unmarshal(b, t.UpdateContactCloseBirthdays)
	case "updateAutosaveSettings":
		t.UpdateAutosaveSettings = new(UpdateAutosaveSettings)
		return json.Unmarshal(b, t.UpdateAutosaveSettings)
	case "updateBusinessConnection":
		t.UpdateBusinessConnection = new(UpdateBusinessConnection)
		return json.Unmarshal(b, t.UpdateBusinessConnection)
	case "updateNewBusinessMessage":
		t.UpdateNewBusinessMessage = new(UpdateNewBusinessMessage)
		return json.Unmarshal(b, t.UpdateNewBusinessMessage)
	case "updateBusinessMessageEdited":
		t.UpdateBusinessMessageEdited = new(UpdateBusinessMessageEdited)
		return json.Unmarshal(b, t.UpdateBusinessMessageEdited)
	case "updateBusinessMessagesDeleted":
		t.UpdateBusinessMessagesDeleted = new(UpdateBusinessMessagesDeleted)
		return json.Unmarshal(b, t.UpdateBusinessMessagesDeleted)
	case "updateNewInlineQuery":
		t.UpdateNewInlineQuery = new(UpdateNewInlineQuery)
		return json.Unmarshal(b, t.UpdateNewInlineQuery)
	case "updateNewChosenInlineResult":
		t.UpdateNewChosenInlineResult = new(UpdateNewChosenInlineResult)
		return json.Unmarshal(b, t.UpdateNewChosenInlineResult)
	case "updateNewCallbackQuery":
		t.UpdateNewCallbackQuery = new(UpdateNewCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewCallbackQuery)
	case "updateNewInlineCallbackQuery":
		t.UpdateNewInlineCallbackQuery = new(UpdateNewInlineCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewInlineCallbackQuery)
	case "updateNewBusinessCallbackQuery":
		t.UpdateNewBusinessCallbackQuery = new(UpdateNewBusinessCallbackQuery)
		return json.Unmarshal(b, t.UpdateNewBusinessCallbackQuery)
	case "updateNewShippingQuery":
		t.UpdateNewShippingQuery = new(UpdateNewShippingQuery)
		return json.Unmarshal(b, t.UpdateNewShippingQuery)
	case "updateNewPreCheckoutQuery":
		t.UpdateNewPreCheckoutQuery = new(UpdateNewPreCheckoutQuery)
		return json.Unmarshal(b, t.UpdateNewPreCheckoutQuery)
	case "updateNewCustomEvent":
		t.UpdateNewCustomEvent = new(UpdateNewCustomEvent)
		return json.Unmarshal(b, t.UpdateNewCustomEvent)
	case "updateNewCustomQuery":
		t.UpdateNewCustomQuery = new(UpdateNewCustomQuery)
		return json.Unmarshal(b, t.UpdateNewCustomQuery)
	case "updatePoll":
		t.UpdatePoll = new(UpdatePoll)
		return json.Unmarshal(b, t.UpdatePoll)
	case "updatePollAnswer":
		t.UpdatePollAnswer = new(UpdatePollAnswer)
		return json.Unmarshal(b, t.UpdatePollAnswer)
	case "updateChatMember":
		t.UpdateChatMember = new(UpdateChatMember)
		return json.Unmarshal(b, t.UpdateChatMember)
	case "updateNewChatJoinRequest":
		t.UpdateNewChatJoinRequest = new(UpdateNewChatJoinRequest)
		return json.Unmarshal(b, t.UpdateNewChatJoinRequest)
	case "updateChatBoost":
		t.UpdateChatBoost = new(UpdateChatBoost)
		return json.Unmarshal(b, t.UpdateChatBoost)
	case "updateMessageReaction":
		t.UpdateMessageReaction = new(UpdateMessageReaction)
		return json.Unmarshal(b, t.UpdateMessageReaction)
	case "updateMessageReactions":
		t.UpdateMessageReactions = new(UpdateMessageReactions)
		return json.Unmarshal(b, t.UpdateMessageReactions)
	case "updatePaidMediaPurchased":
		t.UpdatePaidMediaPurchased = new(UpdatePaidMediaPurchased)
		return json.Unmarshal(b, t.UpdatePaidMediaPurchased)
	}
	return nil
}

func (t *Update) MarshalJSON() ([]byte, error) {
	if t.UpdateAuthorizationState != nil {
		return json.Marshal(t.UpdateAuthorizationState)
	}
	if t.UpdateNewMessage != nil {
		return json.Marshal(t.UpdateNewMessage)
	}
	if t.UpdateMessageSendAcknowledged != nil {
		return json.Marshal(t.UpdateMessageSendAcknowledged)
	}
	if t.UpdateMessageSendSucceeded != nil {
		return json.Marshal(t.UpdateMessageSendSucceeded)
	}
	if t.UpdateMessageSendFailed != nil {
		return json.Marshal(t.UpdateMessageSendFailed)
	}
	if t.UpdateMessageContent != nil {
		return json.Marshal(t.UpdateMessageContent)
	}
	if t.UpdateMessageEdited != nil {
		return json.Marshal(t.UpdateMessageEdited)
	}
	if t.UpdateMessageIsPinned != nil {
		return json.Marshal(t.UpdateMessageIsPinned)
	}
	if t.UpdateMessageInteractionInfo != nil {
		return json.Marshal(t.UpdateMessageInteractionInfo)
	}
	if t.UpdateMessageContentOpened != nil {
		return json.Marshal(t.UpdateMessageContentOpened)
	}
	if t.UpdateMessageMentionRead != nil {
		return json.Marshal(t.UpdateMessageMentionRead)
	}
	if t.UpdateMessageUnreadReactions != nil {
		return json.Marshal(t.UpdateMessageUnreadReactions)
	}
	if t.UpdateMessageFactCheck != nil {
		return json.Marshal(t.UpdateMessageFactCheck)
	}
	if t.UpdateMessageSuggestedPostInfo != nil {
		return json.Marshal(t.UpdateMessageSuggestedPostInfo)
	}
	if t.UpdateMessageLiveLocationViewed != nil {
		return json.Marshal(t.UpdateMessageLiveLocationViewed)
	}
	if t.UpdateVideoPublished != nil {
		return json.Marshal(t.UpdateVideoPublished)
	}
	if t.UpdateNewChat != nil {
		return json.Marshal(t.UpdateNewChat)
	}
	if t.UpdateChatTitle != nil {
		return json.Marshal(t.UpdateChatTitle)
	}
	if t.UpdateChatPhoto != nil {
		return json.Marshal(t.UpdateChatPhoto)
	}
	if t.UpdateChatAccentColors != nil {
		return json.Marshal(t.UpdateChatAccentColors)
	}
	if t.UpdateChatPermissions != nil {
		return json.Marshal(t.UpdateChatPermissions)
	}
	if t.UpdateChatLastMessage != nil {
		return json.Marshal(t.UpdateChatLastMessage)
	}
	if t.UpdateChatPosition != nil {
		return json.Marshal(t.UpdateChatPosition)
	}
	if t.UpdateChatAddedToList != nil {
		return json.Marshal(t.UpdateChatAddedToList)
	}
	if t.UpdateChatRemovedFromList != nil {
		return json.Marshal(t.UpdateChatRemovedFromList)
	}
	if t.UpdateChatReadInbox != nil {
		return json.Marshal(t.UpdateChatReadInbox)
	}
	if t.UpdateChatReadOutbox != nil {
		return json.Marshal(t.UpdateChatReadOutbox)
	}
	if t.UpdateChatActionBar != nil {
		return json.Marshal(t.UpdateChatActionBar)
	}
	if t.UpdateChatBusinessBotManageBar != nil {
		return json.Marshal(t.UpdateChatBusinessBotManageBar)
	}
	if t.UpdateChatAvailableReactions != nil {
		return json.Marshal(t.UpdateChatAvailableReactions)
	}
	if t.UpdateChatDraftMessage != nil {
		return json.Marshal(t.UpdateChatDraftMessage)
	}
	if t.UpdateChatEmojiStatus != nil {
		return json.Marshal(t.UpdateChatEmojiStatus)
	}
	if t.UpdateChatMessageSender != nil {
		return json.Marshal(t.UpdateChatMessageSender)
	}
	if t.UpdateChatMessageAutoDeleteTime != nil {
		return json.Marshal(t.UpdateChatMessageAutoDeleteTime)
	}
	if t.UpdateChatNotificationSettings != nil {
		return json.Marshal(t.UpdateChatNotificationSettings)
	}
	if t.UpdateChatPendingJoinRequests != nil {
		return json.Marshal(t.UpdateChatPendingJoinRequests)
	}
	if t.UpdateChatReplyMarkup != nil {
		return json.Marshal(t.UpdateChatReplyMarkup)
	}
	if t.UpdateChatBackground != nil {
		return json.Marshal(t.UpdateChatBackground)
	}
	if t.UpdateChatTheme != nil {
		return json.Marshal(t.UpdateChatTheme)
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
	if t.UpdateChatDefaultDisableNotification != nil {
		return json.Marshal(t.UpdateChatDefaultDisableNotification)
	}
	if t.UpdateChatHasProtectedContent != nil {
		return json.Marshal(t.UpdateChatHasProtectedContent)
	}
	if t.UpdateChatIsTranslatable != nil {
		return json.Marshal(t.UpdateChatIsTranslatable)
	}
	if t.UpdateChatIsMarkedAsUnread != nil {
		return json.Marshal(t.UpdateChatIsMarkedAsUnread)
	}
	if t.UpdateChatViewAsTopics != nil {
		return json.Marshal(t.UpdateChatViewAsTopics)
	}
	if t.UpdateChatBlockList != nil {
		return json.Marshal(t.UpdateChatBlockList)
	}
	if t.UpdateChatHasScheduledMessages != nil {
		return json.Marshal(t.UpdateChatHasScheduledMessages)
	}
	if t.UpdateChatFolders != nil {
		return json.Marshal(t.UpdateChatFolders)
	}
	if t.UpdateChatOnlineMemberCount != nil {
		return json.Marshal(t.UpdateChatOnlineMemberCount)
	}
	if t.UpdateSavedMessagesTopic != nil {
		return json.Marshal(t.UpdateSavedMessagesTopic)
	}
	if t.UpdateSavedMessagesTopicCount != nil {
		return json.Marshal(t.UpdateSavedMessagesTopicCount)
	}
	if t.UpdateDirectMessagesChatTopic != nil {
		return json.Marshal(t.UpdateDirectMessagesChatTopic)
	}
	if t.UpdateTopicMessageCount != nil {
		return json.Marshal(t.UpdateTopicMessageCount)
	}
	if t.UpdateQuickReplyShortcut != nil {
		return json.Marshal(t.UpdateQuickReplyShortcut)
	}
	if t.UpdateQuickReplyShortcutDeleted != nil {
		return json.Marshal(t.UpdateQuickReplyShortcutDeleted)
	}
	if t.UpdateQuickReplyShortcuts != nil {
		return json.Marshal(t.UpdateQuickReplyShortcuts)
	}
	if t.UpdateQuickReplyShortcutMessages != nil {
		return json.Marshal(t.UpdateQuickReplyShortcutMessages)
	}
	if t.UpdateForumTopicInfo != nil {
		return json.Marshal(t.UpdateForumTopicInfo)
	}
	if t.UpdateForumTopic != nil {
		return json.Marshal(t.UpdateForumTopic)
	}
	if t.UpdateScopeNotificationSettings != nil {
		return json.Marshal(t.UpdateScopeNotificationSettings)
	}
	if t.UpdateReactionNotificationSettings != nil {
		return json.Marshal(t.UpdateReactionNotificationSettings)
	}
	if t.UpdateNotification != nil {
		return json.Marshal(t.UpdateNotification)
	}
	if t.UpdateNotificationGroup != nil {
		return json.Marshal(t.UpdateNotificationGroup)
	}
	if t.UpdateActiveNotifications != nil {
		return json.Marshal(t.UpdateActiveNotifications)
	}
	if t.UpdateHavePendingNotifications != nil {
		return json.Marshal(t.UpdateHavePendingNotifications)
	}
	if t.UpdateDeleteMessages != nil {
		return json.Marshal(t.UpdateDeleteMessages)
	}
	if t.UpdateChatAction != nil {
		return json.Marshal(t.UpdateChatAction)
	}
	if t.UpdatePendingTextMessage != nil {
		return json.Marshal(t.UpdatePendingTextMessage)
	}
	if t.UpdateUserStatus != nil {
		return json.Marshal(t.UpdateUserStatus)
	}
	if t.UpdateUser != nil {
		return json.Marshal(t.UpdateUser)
	}
	if t.UpdateBasicGroup != nil {
		return json.Marshal(t.UpdateBasicGroup)
	}
	if t.UpdateSupergroup != nil {
		return json.Marshal(t.UpdateSupergroup)
	}
	if t.UpdateSecretChat != nil {
		return json.Marshal(t.UpdateSecretChat)
	}
	if t.UpdateUserFullInfo != nil {
		return json.Marshal(t.UpdateUserFullInfo)
	}
	if t.UpdateBasicGroupFullInfo != nil {
		return json.Marshal(t.UpdateBasicGroupFullInfo)
	}
	if t.UpdateSupergroupFullInfo != nil {
		return json.Marshal(t.UpdateSupergroupFullInfo)
	}
	if t.UpdateServiceNotification != nil {
		return json.Marshal(t.UpdateServiceNotification)
	}
	if t.UpdateFile != nil {
		return json.Marshal(t.UpdateFile)
	}
	if t.UpdateFileGenerationStart != nil {
		return json.Marshal(t.UpdateFileGenerationStart)
	}
	if t.UpdateFileGenerationStop != nil {
		return json.Marshal(t.UpdateFileGenerationStop)
	}
	if t.UpdateFileDownloads != nil {
		return json.Marshal(t.UpdateFileDownloads)
	}
	if t.UpdateFileAddedToDownloads != nil {
		return json.Marshal(t.UpdateFileAddedToDownloads)
	}
	if t.UpdateFileDownload != nil {
		return json.Marshal(t.UpdateFileDownload)
	}
	if t.UpdateFileRemovedFromDownloads != nil {
		return json.Marshal(t.UpdateFileRemovedFromDownloads)
	}
	if t.UpdateApplicationVerificationRequired != nil {
		return json.Marshal(t.UpdateApplicationVerificationRequired)
	}
	if t.UpdateApplicationRecaptchaVerificationRequired != nil {
		return json.Marshal(t.UpdateApplicationRecaptchaVerificationRequired)
	}
	if t.UpdateCall != nil {
		return json.Marshal(t.UpdateCall)
	}
	if t.UpdateGroupCall != nil {
		return json.Marshal(t.UpdateGroupCall)
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
	if t.UpdateNewGroupCallMessage != nil {
		return json.Marshal(t.UpdateNewGroupCallMessage)
	}
	if t.UpdateNewGroupCallPaidReaction != nil {
		return json.Marshal(t.UpdateNewGroupCallPaidReaction)
	}
	if t.UpdateGroupCallMessageSendFailed != nil {
		return json.Marshal(t.UpdateGroupCallMessageSendFailed)
	}
	if t.UpdateGroupCallMessagesDeleted != nil {
		return json.Marshal(t.UpdateGroupCallMessagesDeleted)
	}
	if t.UpdateLiveStoryTopDonors != nil {
		return json.Marshal(t.UpdateLiveStoryTopDonors)
	}
	if t.UpdateNewCallSignalingData != nil {
		return json.Marshal(t.UpdateNewCallSignalingData)
	}
	if t.UpdateGiftAuctionState != nil {
		return json.Marshal(t.UpdateGiftAuctionState)
	}
	if t.UpdateActiveGiftAuctions != nil {
		return json.Marshal(t.UpdateActiveGiftAuctions)
	}
	if t.UpdateUserPrivacySettingRules != nil {
		return json.Marshal(t.UpdateUserPrivacySettingRules)
	}
	if t.UpdateUnreadMessageCount != nil {
		return json.Marshal(t.UpdateUnreadMessageCount)
	}
	if t.UpdateUnreadChatCount != nil {
		return json.Marshal(t.UpdateUnreadChatCount)
	}
	if t.UpdateStory != nil {
		return json.Marshal(t.UpdateStory)
	}
	if t.UpdateStoryDeleted != nil {
		return json.Marshal(t.UpdateStoryDeleted)
	}
	if t.UpdateStoryPostSucceeded != nil {
		return json.Marshal(t.UpdateStoryPostSucceeded)
	}
	if t.UpdateStoryPostFailed != nil {
		return json.Marshal(t.UpdateStoryPostFailed)
	}
	if t.UpdateChatActiveStories != nil {
		return json.Marshal(t.UpdateChatActiveStories)
	}
	if t.UpdateStoryListChatCount != nil {
		return json.Marshal(t.UpdateStoryListChatCount)
	}
	if t.UpdateStoryStealthMode != nil {
		return json.Marshal(t.UpdateStoryStealthMode)
	}
	if t.UpdateTrustedMiniAppBots != nil {
		return json.Marshal(t.UpdateTrustedMiniAppBots)
	}
	if t.UpdateOption != nil {
		return json.Marshal(t.UpdateOption)
	}
	if t.UpdateStickerSet != nil {
		return json.Marshal(t.UpdateStickerSet)
	}
	if t.UpdateInstalledStickerSets != nil {
		return json.Marshal(t.UpdateInstalledStickerSets)
	}
	if t.UpdateTrendingStickerSets != nil {
		return json.Marshal(t.UpdateTrendingStickerSets)
	}
	if t.UpdateRecentStickers != nil {
		return json.Marshal(t.UpdateRecentStickers)
	}
	if t.UpdateFavoriteStickers != nil {
		return json.Marshal(t.UpdateFavoriteStickers)
	}
	if t.UpdateSavedAnimations != nil {
		return json.Marshal(t.UpdateSavedAnimations)
	}
	if t.UpdateSavedNotificationSounds != nil {
		return json.Marshal(t.UpdateSavedNotificationSounds)
	}
	if t.UpdateDefaultBackground != nil {
		return json.Marshal(t.UpdateDefaultBackground)
	}
	if t.UpdateEmojiChatThemes != nil {
		return json.Marshal(t.UpdateEmojiChatThemes)
	}
	if t.UpdateAccentColors != nil {
		return json.Marshal(t.UpdateAccentColors)
	}
	if t.UpdateProfileAccentColors != nil {
		return json.Marshal(t.UpdateProfileAccentColors)
	}
	if t.UpdateLanguagePackStrings != nil {
		return json.Marshal(t.UpdateLanguagePackStrings)
	}
	if t.UpdateConnectionState != nil {
		return json.Marshal(t.UpdateConnectionState)
	}
	if t.UpdateFreezeState != nil {
		return json.Marshal(t.UpdateFreezeState)
	}
	if t.UpdateAgeVerificationParameters != nil {
		return json.Marshal(t.UpdateAgeVerificationParameters)
	}
	if t.UpdateTermsOfService != nil {
		return json.Marshal(t.UpdateTermsOfService)
	}
	if t.UpdateUnconfirmedSession != nil {
		return json.Marshal(t.UpdateUnconfirmedSession)
	}
	if t.UpdateAttachmentMenuBots != nil {
		return json.Marshal(t.UpdateAttachmentMenuBots)
	}
	if t.UpdateWebAppMessageSent != nil {
		return json.Marshal(t.UpdateWebAppMessageSent)
	}
	if t.UpdateActiveEmojiReactions != nil {
		return json.Marshal(t.UpdateActiveEmojiReactions)
	}
	if t.UpdateAvailableMessageEffects != nil {
		return json.Marshal(t.UpdateAvailableMessageEffects)
	}
	if t.UpdateDefaultReactionType != nil {
		return json.Marshal(t.UpdateDefaultReactionType)
	}
	if t.UpdateDefaultPaidReactionType != nil {
		return json.Marshal(t.UpdateDefaultPaidReactionType)
	}
	if t.UpdateSavedMessagesTags != nil {
		return json.Marshal(t.UpdateSavedMessagesTags)
	}
	if t.UpdateActiveLiveLocationMessages != nil {
		return json.Marshal(t.UpdateActiveLiveLocationMessages)
	}
	if t.UpdateOwnedStarCount != nil {
		return json.Marshal(t.UpdateOwnedStarCount)
	}
	if t.UpdateOwnedTonCount != nil {
		return json.Marshal(t.UpdateOwnedTonCount)
	}
	if t.UpdateChatRevenueAmount != nil {
		return json.Marshal(t.UpdateChatRevenueAmount)
	}
	if t.UpdateStarRevenueStatus != nil {
		return json.Marshal(t.UpdateStarRevenueStatus)
	}
	if t.UpdateTonRevenueStatus != nil {
		return json.Marshal(t.UpdateTonRevenueStatus)
	}
	if t.UpdateSpeechRecognitionTrial != nil {
		return json.Marshal(t.UpdateSpeechRecognitionTrial)
	}
	if t.UpdateGroupCallMessageLevels != nil {
		return json.Marshal(t.UpdateGroupCallMessageLevels)
	}
	if t.UpdateDiceEmojis != nil {
		return json.Marshal(t.UpdateDiceEmojis)
	}
	if t.UpdateStakeDiceState != nil {
		return json.Marshal(t.UpdateStakeDiceState)
	}
	if t.UpdateAnimatedEmojiMessageClicked != nil {
		return json.Marshal(t.UpdateAnimatedEmojiMessageClicked)
	}
	if t.UpdateAnimationSearchParameters != nil {
		return json.Marshal(t.UpdateAnimationSearchParameters)
	}
	if t.UpdateSuggestedActions != nil {
		return json.Marshal(t.UpdateSuggestedActions)
	}
	if t.UpdateSpeedLimitNotification != nil {
		return json.Marshal(t.UpdateSpeedLimitNotification)
	}
	if t.UpdateContactCloseBirthdays != nil {
		return json.Marshal(t.UpdateContactCloseBirthdays)
	}
	if t.UpdateAutosaveSettings != nil {
		return json.Marshal(t.UpdateAutosaveSettings)
	}
	if t.UpdateBusinessConnection != nil {
		return json.Marshal(t.UpdateBusinessConnection)
	}
	if t.UpdateNewBusinessMessage != nil {
		return json.Marshal(t.UpdateNewBusinessMessage)
	}
	if t.UpdateBusinessMessageEdited != nil {
		return json.Marshal(t.UpdateBusinessMessageEdited)
	}
	if t.UpdateBusinessMessagesDeleted != nil {
		return json.Marshal(t.UpdateBusinessMessagesDeleted)
	}
	if t.UpdateNewInlineQuery != nil {
		return json.Marshal(t.UpdateNewInlineQuery)
	}
	if t.UpdateNewChosenInlineResult != nil {
		return json.Marshal(t.UpdateNewChosenInlineResult)
	}
	if t.UpdateNewCallbackQuery != nil {
		return json.Marshal(t.UpdateNewCallbackQuery)
	}
	if t.UpdateNewInlineCallbackQuery != nil {
		return json.Marshal(t.UpdateNewInlineCallbackQuery)
	}
	if t.UpdateNewBusinessCallbackQuery != nil {
		return json.Marshal(t.UpdateNewBusinessCallbackQuery)
	}
	if t.UpdateNewShippingQuery != nil {
		return json.Marshal(t.UpdateNewShippingQuery)
	}
	if t.UpdateNewPreCheckoutQuery != nil {
		return json.Marshal(t.UpdateNewPreCheckoutQuery)
	}
	if t.UpdateNewCustomEvent != nil {
		return json.Marshal(t.UpdateNewCustomEvent)
	}
	if t.UpdateNewCustomQuery != nil {
		return json.Marshal(t.UpdateNewCustomQuery)
	}
	if t.UpdatePoll != nil {
		return json.Marshal(t.UpdatePoll)
	}
	if t.UpdatePollAnswer != nil {
		return json.Marshal(t.UpdatePollAnswer)
	}
	if t.UpdateChatMember != nil {
		return json.Marshal(t.UpdateChatMember)
	}
	if t.UpdateNewChatJoinRequest != nil {
		return json.Marshal(t.UpdateNewChatJoinRequest)
	}
	if t.UpdateChatBoost != nil {
		return json.Marshal(t.UpdateChatBoost)
	}
	if t.UpdateMessageReaction != nil {
		return json.Marshal(t.UpdateMessageReaction)
	}
	if t.UpdateMessageReactions != nil {
		return json.Marshal(t.UpdateMessageReactions)
	}
	if t.UpdatePaidMediaPurchased != nil {
		return json.Marshal(t.UpdatePaidMediaPurchased)
	}
	return nil, fmt.Errorf("no value set for Update")
}

// MaskPoint Part of the face, relative to which a mask is placed
type MaskPoint struct {
	MaskPointForehead *MaskPointForehead `json:"maskPointForehead,omitempty"`
	MaskPointEyes     *MaskPointEyes     `json:"maskPointEyes,omitempty"`
	MaskPointMouth    *MaskPointMouth    `json:"maskPointMouth,omitempty"`
	MaskPointChin     *MaskPointChin     `json:"maskPointChin,omitempty"`
}

func (t *MaskPoint) Type() string {
	return ""
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
	switch typeObj.Type {
	case "maskPointForehead":
		t.MaskPointForehead = new(MaskPointForehead)
		return json.Unmarshal(b, t.MaskPointForehead)
	case "maskPointEyes":
		t.MaskPointEyes = new(MaskPointEyes)
		return json.Unmarshal(b, t.MaskPointEyes)
	case "maskPointMouth":
		t.MaskPointMouth = new(MaskPointMouth)
		return json.Unmarshal(b, t.MaskPointMouth)
	case "maskPointChin":
		t.MaskPointChin = new(MaskPointChin)
		return json.Unmarshal(b, t.MaskPointChin)
	}
	return nil
}

func (t *MaskPoint) MarshalJSON() ([]byte, error) {
	if t.MaskPointForehead != nil {
		return json.Marshal(t.MaskPointForehead)
	}
	if t.MaskPointEyes != nil {
		return json.Marshal(t.MaskPointEyes)
	}
	if t.MaskPointMouth != nil {
		return json.Marshal(t.MaskPointMouth)
	}
	if t.MaskPointChin != nil {
		return json.Marshal(t.MaskPointChin)
	}
	return nil, fmt.Errorf("no value set for MaskPoint")
}

// GiftForResaleOrder Describes order in which upgraded gifts for resale will be sorted
type GiftForResaleOrder struct {
	GiftForResaleOrderPrice           *GiftForResaleOrderPrice           `json:"giftForResaleOrderPrice,omitempty"`
	GiftForResaleOrderPriceChangeDate *GiftForResaleOrderPriceChangeDate `json:"giftForResaleOrderPriceChangeDate,omitempty"`
	GiftForResaleOrderNumber          *GiftForResaleOrderNumber          `json:"giftForResaleOrderNumber,omitempty"`
}

func (t *GiftForResaleOrder) Type() string {
	return ""
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
	switch typeObj.Type {
	case "giftForResaleOrderPrice":
		t.GiftForResaleOrderPrice = new(GiftForResaleOrderPrice)
		return json.Unmarshal(b, t.GiftForResaleOrderPrice)
	case "giftForResaleOrderPriceChangeDate":
		t.GiftForResaleOrderPriceChangeDate = new(GiftForResaleOrderPriceChangeDate)
		return json.Unmarshal(b, t.GiftForResaleOrderPriceChangeDate)
	case "giftForResaleOrderNumber":
		t.GiftForResaleOrderNumber = new(GiftForResaleOrderNumber)
		return json.Unmarshal(b, t.GiftForResaleOrderNumber)
	}
	return nil
}

func (t *GiftForResaleOrder) MarshalJSON() ([]byte, error) {
	if t.GiftForResaleOrderPrice != nil {
		return json.Marshal(t.GiftForResaleOrderPrice)
	}
	if t.GiftForResaleOrderPriceChangeDate != nil {
		return json.Marshal(t.GiftForResaleOrderPriceChangeDate)
	}
	if t.GiftForResaleOrderNumber != nil {
		return json.Marshal(t.GiftForResaleOrderNumber)
	}
	return nil, fmt.Errorf("no value set for GiftForResaleOrder")
}

// ChatMemberStatus Provides information about the status of a member in a chat
type ChatMemberStatus struct {
	ChatMemberStatusCreator       *ChatMemberStatusCreator       `json:"chatMemberStatusCreator,omitempty"`
	ChatMemberStatusAdministrator *ChatMemberStatusAdministrator `json:"chatMemberStatusAdministrator,omitempty"`
	ChatMemberStatusMember        *ChatMemberStatusMember        `json:"chatMemberStatusMember,omitempty"`
	ChatMemberStatusRestricted    *ChatMemberStatusRestricted    `json:"chatMemberStatusRestricted,omitempty"`
	ChatMemberStatusLeft          *ChatMemberStatusLeft          `json:"chatMemberStatusLeft,omitempty"`
	ChatMemberStatusBanned        *ChatMemberStatusBanned        `json:"chatMemberStatusBanned,omitempty"`
}

func (t *ChatMemberStatus) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatMemberStatusCreator":
		t.ChatMemberStatusCreator = new(ChatMemberStatusCreator)
		return json.Unmarshal(b, t.ChatMemberStatusCreator)
	case "chatMemberStatusAdministrator":
		t.ChatMemberStatusAdministrator = new(ChatMemberStatusAdministrator)
		return json.Unmarshal(b, t.ChatMemberStatusAdministrator)
	case "chatMemberStatusMember":
		t.ChatMemberStatusMember = new(ChatMemberStatusMember)
		return json.Unmarshal(b, t.ChatMemberStatusMember)
	case "chatMemberStatusRestricted":
		t.ChatMemberStatusRestricted = new(ChatMemberStatusRestricted)
		return json.Unmarshal(b, t.ChatMemberStatusRestricted)
	case "chatMemberStatusLeft":
		t.ChatMemberStatusLeft = new(ChatMemberStatusLeft)
		return json.Unmarshal(b, t.ChatMemberStatusLeft)
	case "chatMemberStatusBanned":
		t.ChatMemberStatusBanned = new(ChatMemberStatusBanned)
		return json.Unmarshal(b, t.ChatMemberStatusBanned)
	}
	return nil
}

func (t *ChatMemberStatus) MarshalJSON() ([]byte, error) {
	if t.ChatMemberStatusCreator != nil {
		return json.Marshal(t.ChatMemberStatusCreator)
	}
	if t.ChatMemberStatusAdministrator != nil {
		return json.Marshal(t.ChatMemberStatusAdministrator)
	}
	if t.ChatMemberStatusMember != nil {
		return json.Marshal(t.ChatMemberStatusMember)
	}
	if t.ChatMemberStatusRestricted != nil {
		return json.Marshal(t.ChatMemberStatusRestricted)
	}
	if t.ChatMemberStatusLeft != nil {
		return json.Marshal(t.ChatMemberStatusLeft)
	}
	if t.ChatMemberStatusBanned != nil {
		return json.Marshal(t.ChatMemberStatusBanned)
	}
	return nil, fmt.Errorf("no value set for ChatMemberStatus")
}

// MessageOrigin Contains information about the origin of a message
type MessageOrigin struct {
	MessageOriginUser       *MessageOriginUser       `json:"messageOriginUser,omitempty"`
	MessageOriginHiddenUser *MessageOriginHiddenUser `json:"messageOriginHiddenUser,omitempty"`
	MessageOriginChat       *MessageOriginChat       `json:"messageOriginChat,omitempty"`
	MessageOriginChannel    *MessageOriginChannel    `json:"messageOriginChannel,omitempty"`
}

func (t *MessageOrigin) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageOriginUser":
		t.MessageOriginUser = new(MessageOriginUser)
		return json.Unmarshal(b, t.MessageOriginUser)
	case "messageOriginHiddenUser":
		t.MessageOriginHiddenUser = new(MessageOriginHiddenUser)
		return json.Unmarshal(b, t.MessageOriginHiddenUser)
	case "messageOriginChat":
		t.MessageOriginChat = new(MessageOriginChat)
		return json.Unmarshal(b, t.MessageOriginChat)
	case "messageOriginChannel":
		t.MessageOriginChannel = new(MessageOriginChannel)
		return json.Unmarshal(b, t.MessageOriginChannel)
	}
	return nil
}

func (t *MessageOrigin) MarshalJSON() ([]byte, error) {
	if t.MessageOriginUser != nil {
		return json.Marshal(t.MessageOriginUser)
	}
	if t.MessageOriginHiddenUser != nil {
		return json.Marshal(t.MessageOriginHiddenUser)
	}
	if t.MessageOriginChat != nil {
		return json.Marshal(t.MessageOriginChat)
	}
	if t.MessageOriginChannel != nil {
		return json.Marshal(t.MessageOriginChannel)
	}
	return nil, fmt.Errorf("no value set for MessageOrigin")
}

// ChatList Describes a list of chats
type ChatList struct {
	ChatListMain    *ChatListMain    `json:"chatListMain,omitempty"`
	ChatListArchive *ChatListArchive `json:"chatListArchive,omitempty"`
	ChatListFolder  *ChatListFolder  `json:"chatListFolder,omitempty"`
}

func (t *ChatList) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatListMain":
		t.ChatListMain = new(ChatListMain)
		return json.Unmarshal(b, t.ChatListMain)
	case "chatListArchive":
		t.ChatListArchive = new(ChatListArchive)
		return json.Unmarshal(b, t.ChatListArchive)
	case "chatListFolder":
		t.ChatListFolder = new(ChatListFolder)
		return json.Unmarshal(b, t.ChatListFolder)
	}
	return nil
}

func (t *ChatList) MarshalJSON() ([]byte, error) {
	if t.ChatListMain != nil {
		return json.Marshal(t.ChatListMain)
	}
	if t.ChatListArchive != nil {
		return json.Marshal(t.ChatListArchive)
	}
	if t.ChatListFolder != nil {
		return json.Marshal(t.ChatListFolder)
	}
	return nil, fmt.Errorf("no value set for ChatList")
}

// KeyboardButtonType Describes a keyboard button type
type KeyboardButtonType struct {
	KeyboardButtonTypeText               *KeyboardButtonTypeText               `json:"keyboardButtonTypeText,omitempty"`
	KeyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber `json:"keyboardButtonTypeRequestPhoneNumber,omitempty"`
	KeyboardButtonTypeRequestLocation    *KeyboardButtonTypeRequestLocation    `json:"keyboardButtonTypeRequestLocation,omitempty"`
	KeyboardButtonTypeRequestPoll        *KeyboardButtonTypeRequestPoll        `json:"keyboardButtonTypeRequestPoll,omitempty"`
	KeyboardButtonTypeRequestUsers       *KeyboardButtonTypeRequestUsers       `json:"keyboardButtonTypeRequestUsers,omitempty"`
	KeyboardButtonTypeRequestChat        *KeyboardButtonTypeRequestChat        `json:"keyboardButtonTypeRequestChat,omitempty"`
	KeyboardButtonTypeWebApp             *KeyboardButtonTypeWebApp             `json:"keyboardButtonTypeWebApp,omitempty"`
}

func (t *KeyboardButtonType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "keyboardButtonTypeText":
		t.KeyboardButtonTypeText = new(KeyboardButtonTypeText)
		return json.Unmarshal(b, t.KeyboardButtonTypeText)
	case "keyboardButtonTypeRequestPhoneNumber":
		t.KeyboardButtonTypeRequestPhoneNumber = new(KeyboardButtonTypeRequestPhoneNumber)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestPhoneNumber)
	case "keyboardButtonTypeRequestLocation":
		t.KeyboardButtonTypeRequestLocation = new(KeyboardButtonTypeRequestLocation)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestLocation)
	case "keyboardButtonTypeRequestPoll":
		t.KeyboardButtonTypeRequestPoll = new(KeyboardButtonTypeRequestPoll)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestPoll)
	case "keyboardButtonTypeRequestUsers":
		t.KeyboardButtonTypeRequestUsers = new(KeyboardButtonTypeRequestUsers)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestUsers)
	case "keyboardButtonTypeRequestChat":
		t.KeyboardButtonTypeRequestChat = new(KeyboardButtonTypeRequestChat)
		return json.Unmarshal(b, t.KeyboardButtonTypeRequestChat)
	case "keyboardButtonTypeWebApp":
		t.KeyboardButtonTypeWebApp = new(KeyboardButtonTypeWebApp)
		return json.Unmarshal(b, t.KeyboardButtonTypeWebApp)
	}
	return nil
}

func (t *KeyboardButtonType) MarshalJSON() ([]byte, error) {
	if t.KeyboardButtonTypeText != nil {
		return json.Marshal(t.KeyboardButtonTypeText)
	}
	if t.KeyboardButtonTypeRequestPhoneNumber != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestPhoneNumber)
	}
	if t.KeyboardButtonTypeRequestLocation != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestLocation)
	}
	if t.KeyboardButtonTypeRequestPoll != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestPoll)
	}
	if t.KeyboardButtonTypeRequestUsers != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestUsers)
	}
	if t.KeyboardButtonTypeRequestChat != nil {
		return json.Marshal(t.KeyboardButtonTypeRequestChat)
	}
	if t.KeyboardButtonTypeWebApp != nil {
		return json.Marshal(t.KeyboardButtonTypeWebApp)
	}
	return nil, fmt.Errorf("no value set for KeyboardButtonType")
}

// PassportElement Contains information about a Telegram Passport element
type PassportElement struct {
	PassportElementPersonalDetails       *PassportElementPersonalDetails       `json:"passportElementPersonalDetails,omitempty"`
	PassportElementPassport              *PassportElementPassport              `json:"passportElementPassport,omitempty"`
	PassportElementDriverLicense         *PassportElementDriverLicense         `json:"passportElementDriverLicense,omitempty"`
	PassportElementIdentityCard          *PassportElementIdentityCard          `json:"passportElementIdentityCard,omitempty"`
	PassportElementInternalPassport      *PassportElementInternalPassport      `json:"passportElementInternalPassport,omitempty"`
	PassportElementAddress               *PassportElementAddress               `json:"passportElementAddress,omitempty"`
	PassportElementUtilityBill           *PassportElementUtilityBill           `json:"passportElementUtilityBill,omitempty"`
	PassportElementBankStatement         *PassportElementBankStatement         `json:"passportElementBankStatement,omitempty"`
	PassportElementRentalAgreement       *PassportElementRentalAgreement       `json:"passportElementRentalAgreement,omitempty"`
	PassportElementPassportRegistration  *PassportElementPassportRegistration  `json:"passportElementPassportRegistration,omitempty"`
	PassportElementTemporaryRegistration *PassportElementTemporaryRegistration `json:"passportElementTemporaryRegistration,omitempty"`
	PassportElementPhoneNumber           *PassportElementPhoneNumber           `json:"passportElementPhoneNumber,omitempty"`
	PassportElementEmailAddress          *PassportElementEmailAddress          `json:"passportElementEmailAddress,omitempty"`
}

func (t *PassportElement) Type() string {
	return ""
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
	switch typeObj.Type {
	case "passportElementPersonalDetails":
		t.PassportElementPersonalDetails = new(PassportElementPersonalDetails)
		return json.Unmarshal(b, t.PassportElementPersonalDetails)
	case "passportElementPassport":
		t.PassportElementPassport = new(PassportElementPassport)
		return json.Unmarshal(b, t.PassportElementPassport)
	case "passportElementDriverLicense":
		t.PassportElementDriverLicense = new(PassportElementDriverLicense)
		return json.Unmarshal(b, t.PassportElementDriverLicense)
	case "passportElementIdentityCard":
		t.PassportElementIdentityCard = new(PassportElementIdentityCard)
		return json.Unmarshal(b, t.PassportElementIdentityCard)
	case "passportElementInternalPassport":
		t.PassportElementInternalPassport = new(PassportElementInternalPassport)
		return json.Unmarshal(b, t.PassportElementInternalPassport)
	case "passportElementAddress":
		t.PassportElementAddress = new(PassportElementAddress)
		return json.Unmarshal(b, t.PassportElementAddress)
	case "passportElementUtilityBill":
		t.PassportElementUtilityBill = new(PassportElementUtilityBill)
		return json.Unmarshal(b, t.PassportElementUtilityBill)
	case "passportElementBankStatement":
		t.PassportElementBankStatement = new(PassportElementBankStatement)
		return json.Unmarshal(b, t.PassportElementBankStatement)
	case "passportElementRentalAgreement":
		t.PassportElementRentalAgreement = new(PassportElementRentalAgreement)
		return json.Unmarshal(b, t.PassportElementRentalAgreement)
	case "passportElementPassportRegistration":
		t.PassportElementPassportRegistration = new(PassportElementPassportRegistration)
		return json.Unmarshal(b, t.PassportElementPassportRegistration)
	case "passportElementTemporaryRegistration":
		t.PassportElementTemporaryRegistration = new(PassportElementTemporaryRegistration)
		return json.Unmarshal(b, t.PassportElementTemporaryRegistration)
	case "passportElementPhoneNumber":
		t.PassportElementPhoneNumber = new(PassportElementPhoneNumber)
		return json.Unmarshal(b, t.PassportElementPhoneNumber)
	case "passportElementEmailAddress":
		t.PassportElementEmailAddress = new(PassportElementEmailAddress)
		return json.Unmarshal(b, t.PassportElementEmailAddress)
	}
	return nil
}

func (t *PassportElement) MarshalJSON() ([]byte, error) {
	if t.PassportElementPersonalDetails != nil {
		return json.Marshal(t.PassportElementPersonalDetails)
	}
	if t.PassportElementPassport != nil {
		return json.Marshal(t.PassportElementPassport)
	}
	if t.PassportElementDriverLicense != nil {
		return json.Marshal(t.PassportElementDriverLicense)
	}
	if t.PassportElementIdentityCard != nil {
		return json.Marshal(t.PassportElementIdentityCard)
	}
	if t.PassportElementInternalPassport != nil {
		return json.Marshal(t.PassportElementInternalPassport)
	}
	if t.PassportElementAddress != nil {
		return json.Marshal(t.PassportElementAddress)
	}
	if t.PassportElementUtilityBill != nil {
		return json.Marshal(t.PassportElementUtilityBill)
	}
	if t.PassportElementBankStatement != nil {
		return json.Marshal(t.PassportElementBankStatement)
	}
	if t.PassportElementRentalAgreement != nil {
		return json.Marshal(t.PassportElementRentalAgreement)
	}
	if t.PassportElementPassportRegistration != nil {
		return json.Marshal(t.PassportElementPassportRegistration)
	}
	if t.PassportElementTemporaryRegistration != nil {
		return json.Marshal(t.PassportElementTemporaryRegistration)
	}
	if t.PassportElementPhoneNumber != nil {
		return json.Marshal(t.PassportElementPhoneNumber)
	}
	if t.PassportElementEmailAddress != nil {
		return json.Marshal(t.PassportElementEmailAddress)
	}
	return nil, fmt.Errorf("no value set for PassportElement")
}

// ChatAction Describes the different types of activity in a chat
type ChatAction struct {
	ChatActionTyping             *ChatActionTyping             `json:"chatActionTyping,omitempty"`
	ChatActionRecordingVideo     *ChatActionRecordingVideo     `json:"chatActionRecordingVideo,omitempty"`
	ChatActionUploadingVideo     *ChatActionUploadingVideo     `json:"chatActionUploadingVideo,omitempty"`
	ChatActionRecordingVoiceNote *ChatActionRecordingVoiceNote `json:"chatActionRecordingVoiceNote,omitempty"`
	ChatActionUploadingVoiceNote *ChatActionUploadingVoiceNote `json:"chatActionUploadingVoiceNote,omitempty"`
	ChatActionUploadingPhoto     *ChatActionUploadingPhoto     `json:"chatActionUploadingPhoto,omitempty"`
	ChatActionUploadingDocument  *ChatActionUploadingDocument  `json:"chatActionUploadingDocument,omitempty"`
	ChatActionChoosingSticker    *ChatActionChoosingSticker    `json:"chatActionChoosingSticker,omitempty"`
	ChatActionChoosingLocation   *ChatActionChoosingLocation   `json:"chatActionChoosingLocation,omitempty"`
	ChatActionChoosingContact    *ChatActionChoosingContact    `json:"chatActionChoosingContact,omitempty"`
	ChatActionStartPlayingGame   *ChatActionStartPlayingGame   `json:"chatActionStartPlayingGame,omitempty"`
	ChatActionRecordingVideoNote *ChatActionRecordingVideoNote `json:"chatActionRecordingVideoNote,omitempty"`
	ChatActionUploadingVideoNote *ChatActionUploadingVideoNote `json:"chatActionUploadingVideoNote,omitempty"`
	ChatActionWatchingAnimations *ChatActionWatchingAnimations `json:"chatActionWatchingAnimations,omitempty"`
	ChatActionCancel             *ChatActionCancel             `json:"chatActionCancel,omitempty"`
}

func (t *ChatAction) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatActionTyping":
		t.ChatActionTyping = new(ChatActionTyping)
		return json.Unmarshal(b, t.ChatActionTyping)
	case "chatActionRecordingVideo":
		t.ChatActionRecordingVideo = new(ChatActionRecordingVideo)
		return json.Unmarshal(b, t.ChatActionRecordingVideo)
	case "chatActionUploadingVideo":
		t.ChatActionUploadingVideo = new(ChatActionUploadingVideo)
		return json.Unmarshal(b, t.ChatActionUploadingVideo)
	case "chatActionRecordingVoiceNote":
		t.ChatActionRecordingVoiceNote = new(ChatActionRecordingVoiceNote)
		return json.Unmarshal(b, t.ChatActionRecordingVoiceNote)
	case "chatActionUploadingVoiceNote":
		t.ChatActionUploadingVoiceNote = new(ChatActionUploadingVoiceNote)
		return json.Unmarshal(b, t.ChatActionUploadingVoiceNote)
	case "chatActionUploadingPhoto":
		t.ChatActionUploadingPhoto = new(ChatActionUploadingPhoto)
		return json.Unmarshal(b, t.ChatActionUploadingPhoto)
	case "chatActionUploadingDocument":
		t.ChatActionUploadingDocument = new(ChatActionUploadingDocument)
		return json.Unmarshal(b, t.ChatActionUploadingDocument)
	case "chatActionChoosingSticker":
		t.ChatActionChoosingSticker = new(ChatActionChoosingSticker)
		return json.Unmarshal(b, t.ChatActionChoosingSticker)
	case "chatActionChoosingLocation":
		t.ChatActionChoosingLocation = new(ChatActionChoosingLocation)
		return json.Unmarshal(b, t.ChatActionChoosingLocation)
	case "chatActionChoosingContact":
		t.ChatActionChoosingContact = new(ChatActionChoosingContact)
		return json.Unmarshal(b, t.ChatActionChoosingContact)
	case "chatActionStartPlayingGame":
		t.ChatActionStartPlayingGame = new(ChatActionStartPlayingGame)
		return json.Unmarshal(b, t.ChatActionStartPlayingGame)
	case "chatActionRecordingVideoNote":
		t.ChatActionRecordingVideoNote = new(ChatActionRecordingVideoNote)
		return json.Unmarshal(b, t.ChatActionRecordingVideoNote)
	case "chatActionUploadingVideoNote":
		t.ChatActionUploadingVideoNote = new(ChatActionUploadingVideoNote)
		return json.Unmarshal(b, t.ChatActionUploadingVideoNote)
	case "chatActionWatchingAnimations":
		t.ChatActionWatchingAnimations = new(ChatActionWatchingAnimations)
		return json.Unmarshal(b, t.ChatActionWatchingAnimations)
	case "chatActionCancel":
		t.ChatActionCancel = new(ChatActionCancel)
		return json.Unmarshal(b, t.ChatActionCancel)
	}
	return nil
}

func (t *ChatAction) MarshalJSON() ([]byte, error) {
	if t.ChatActionTyping != nil {
		return json.Marshal(t.ChatActionTyping)
	}
	if t.ChatActionRecordingVideo != nil {
		return json.Marshal(t.ChatActionRecordingVideo)
	}
	if t.ChatActionUploadingVideo != nil {
		return json.Marshal(t.ChatActionUploadingVideo)
	}
	if t.ChatActionRecordingVoiceNote != nil {
		return json.Marshal(t.ChatActionRecordingVoiceNote)
	}
	if t.ChatActionUploadingVoiceNote != nil {
		return json.Marshal(t.ChatActionUploadingVoiceNote)
	}
	if t.ChatActionUploadingPhoto != nil {
		return json.Marshal(t.ChatActionUploadingPhoto)
	}
	if t.ChatActionUploadingDocument != nil {
		return json.Marshal(t.ChatActionUploadingDocument)
	}
	if t.ChatActionChoosingSticker != nil {
		return json.Marshal(t.ChatActionChoosingSticker)
	}
	if t.ChatActionChoosingLocation != nil {
		return json.Marshal(t.ChatActionChoosingLocation)
	}
	if t.ChatActionChoosingContact != nil {
		return json.Marshal(t.ChatActionChoosingContact)
	}
	if t.ChatActionStartPlayingGame != nil {
		return json.Marshal(t.ChatActionStartPlayingGame)
	}
	if t.ChatActionRecordingVideoNote != nil {
		return json.Marshal(t.ChatActionRecordingVideoNote)
	}
	if t.ChatActionUploadingVideoNote != nil {
		return json.Marshal(t.ChatActionUploadingVideoNote)
	}
	if t.ChatActionWatchingAnimations != nil {
		return json.Marshal(t.ChatActionWatchingAnimations)
	}
	if t.ChatActionCancel != nil {
		return json.Marshal(t.ChatActionCancel)
	}
	return nil, fmt.Errorf("no value set for ChatAction")
}

// ChatBoostSource Describes source of a chat boost
type ChatBoostSource struct {
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode `json:"chatBoostSourceGiftCode,omitempty"`
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway `json:"chatBoostSourceGiveaway,omitempty"`
	ChatBoostSourcePremium  *ChatBoostSourcePremium  `json:"chatBoostSourcePremium,omitempty"`
}

func (t *ChatBoostSource) Type() string {
	return ""
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

// StarSubscriptionType Describes type of subscription paid in Telegram Stars
type StarSubscriptionType struct {
	StarSubscriptionTypeChannel *StarSubscriptionTypeChannel `json:"starSubscriptionTypeChannel,omitempty"`
	StarSubscriptionTypeBot     *StarSubscriptionTypeBot     `json:"starSubscriptionTypeBot,omitempty"`
}

func (t *StarSubscriptionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "starSubscriptionTypeChannel":
		t.StarSubscriptionTypeChannel = new(StarSubscriptionTypeChannel)
		return json.Unmarshal(b, t.StarSubscriptionTypeChannel)
	case "starSubscriptionTypeBot":
		t.StarSubscriptionTypeBot = new(StarSubscriptionTypeBot)
		return json.Unmarshal(b, t.StarSubscriptionTypeBot)
	}
	return nil
}

func (t *StarSubscriptionType) MarshalJSON() ([]byte, error) {
	if t.StarSubscriptionTypeChannel != nil {
		return json.Marshal(t.StarSubscriptionTypeChannel)
	}
	if t.StarSubscriptionTypeBot != nil {
		return json.Marshal(t.StarSubscriptionTypeBot)
	}
	return nil, fmt.Errorf("no value set for StarSubscriptionType")
}

// GroupCallDataChannel Describes data channel for a group call
type GroupCallDataChannel struct {
	GroupCallDataChannelMain          *GroupCallDataChannelMain          `json:"groupCallDataChannelMain,omitempty"`
	GroupCallDataChannelScreenSharing *GroupCallDataChannelScreenSharing `json:"groupCallDataChannelScreenSharing,omitempty"`
}

func (t *GroupCallDataChannel) Type() string {
	return ""
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

// ReactionUnavailabilityReason Describes why the current user can't add reactions to the message, despite some other users can
type ReactionUnavailabilityReason struct {
	ReactionUnavailabilityReasonAnonymousAdministrator *ReactionUnavailabilityReasonAnonymousAdministrator `json:"reactionUnavailabilityReasonAnonymousAdministrator,omitempty"`
	ReactionUnavailabilityReasonGuest                  *ReactionUnavailabilityReasonGuest                  `json:"reactionUnavailabilityReasonGuest,omitempty"`
}

func (t *ReactionUnavailabilityReason) Type() string {
	return ""
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

// BusinessFeature Describes a feature available to Business user accounts
type BusinessFeature struct {
	BusinessFeatureLocation        *BusinessFeatureLocation        `json:"businessFeatureLocation,omitempty"`
	BusinessFeatureOpeningHours    *BusinessFeatureOpeningHours    `json:"businessFeatureOpeningHours,omitempty"`
	BusinessFeatureQuickReplies    *BusinessFeatureQuickReplies    `json:"businessFeatureQuickReplies,omitempty"`
	BusinessFeatureGreetingMessage *BusinessFeatureGreetingMessage `json:"businessFeatureGreetingMessage,omitempty"`
	BusinessFeatureAwayMessage     *BusinessFeatureAwayMessage     `json:"businessFeatureAwayMessage,omitempty"`
	BusinessFeatureAccountLinks    *BusinessFeatureAccountLinks    `json:"businessFeatureAccountLinks,omitempty"`
	BusinessFeatureStartPage       *BusinessFeatureStartPage       `json:"businessFeatureStartPage,omitempty"`
	BusinessFeatureBots            *BusinessFeatureBots            `json:"businessFeatureBots,omitempty"`
	BusinessFeatureEmojiStatus     *BusinessFeatureEmojiStatus     `json:"businessFeatureEmojiStatus,omitempty"`
	BusinessFeatureChatFolderTags  *BusinessFeatureChatFolderTags  `json:"businessFeatureChatFolderTags,omitempty"`
	BusinessFeatureUpgradedStories *BusinessFeatureUpgradedStories `json:"businessFeatureUpgradedStories,omitempty"`
}

func (t *BusinessFeature) Type() string {
	return ""
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
	switch typeObj.Type {
	case "businessFeatureLocation":
		t.BusinessFeatureLocation = new(BusinessFeatureLocation)
		return json.Unmarshal(b, t.BusinessFeatureLocation)
	case "businessFeatureOpeningHours":
		t.BusinessFeatureOpeningHours = new(BusinessFeatureOpeningHours)
		return json.Unmarshal(b, t.BusinessFeatureOpeningHours)
	case "businessFeatureQuickReplies":
		t.BusinessFeatureQuickReplies = new(BusinessFeatureQuickReplies)
		return json.Unmarshal(b, t.BusinessFeatureQuickReplies)
	case "businessFeatureGreetingMessage":
		t.BusinessFeatureGreetingMessage = new(BusinessFeatureGreetingMessage)
		return json.Unmarshal(b, t.BusinessFeatureGreetingMessage)
	case "businessFeatureAwayMessage":
		t.BusinessFeatureAwayMessage = new(BusinessFeatureAwayMessage)
		return json.Unmarshal(b, t.BusinessFeatureAwayMessage)
	case "businessFeatureAccountLinks":
		t.BusinessFeatureAccountLinks = new(BusinessFeatureAccountLinks)
		return json.Unmarshal(b, t.BusinessFeatureAccountLinks)
	case "businessFeatureStartPage":
		t.BusinessFeatureStartPage = new(BusinessFeatureStartPage)
		return json.Unmarshal(b, t.BusinessFeatureStartPage)
	case "businessFeatureBots":
		t.BusinessFeatureBots = new(BusinessFeatureBots)
		return json.Unmarshal(b, t.BusinessFeatureBots)
	case "businessFeatureEmojiStatus":
		t.BusinessFeatureEmojiStatus = new(BusinessFeatureEmojiStatus)
		return json.Unmarshal(b, t.BusinessFeatureEmojiStatus)
	case "businessFeatureChatFolderTags":
		t.BusinessFeatureChatFolderTags = new(BusinessFeatureChatFolderTags)
		return json.Unmarshal(b, t.BusinessFeatureChatFolderTags)
	case "businessFeatureUpgradedStories":
		t.BusinessFeatureUpgradedStories = new(BusinessFeatureUpgradedStories)
		return json.Unmarshal(b, t.BusinessFeatureUpgradedStories)
	}
	return nil
}

func (t *BusinessFeature) MarshalJSON() ([]byte, error) {
	if t.BusinessFeatureLocation != nil {
		return json.Marshal(t.BusinessFeatureLocation)
	}
	if t.BusinessFeatureOpeningHours != nil {
		return json.Marshal(t.BusinessFeatureOpeningHours)
	}
	if t.BusinessFeatureQuickReplies != nil {
		return json.Marshal(t.BusinessFeatureQuickReplies)
	}
	if t.BusinessFeatureGreetingMessage != nil {
		return json.Marshal(t.BusinessFeatureGreetingMessage)
	}
	if t.BusinessFeatureAwayMessage != nil {
		return json.Marshal(t.BusinessFeatureAwayMessage)
	}
	if t.BusinessFeatureAccountLinks != nil {
		return json.Marshal(t.BusinessFeatureAccountLinks)
	}
	if t.BusinessFeatureStartPage != nil {
		return json.Marshal(t.BusinessFeatureStartPage)
	}
	if t.BusinessFeatureBots != nil {
		return json.Marshal(t.BusinessFeatureBots)
	}
	if t.BusinessFeatureEmojiStatus != nil {
		return json.Marshal(t.BusinessFeatureEmojiStatus)
	}
	if t.BusinessFeatureChatFolderTags != nil {
		return json.Marshal(t.BusinessFeatureChatFolderTags)
	}
	if t.BusinessFeatureUpgradedStories != nil {
		return json.Marshal(t.BusinessFeatureUpgradedStories)
	}
	return nil, fmt.Errorf("no value set for BusinessFeature")
}

// PremiumSource Describes a source from which the Premium features screen is opened
type PremiumSource struct {
	PremiumSourceLimitExceeded   *PremiumSourceLimitExceeded   `json:"premiumSourceLimitExceeded,omitempty"`
	PremiumSourceFeature         *PremiumSourceFeature         `json:"premiumSourceFeature,omitempty"`
	PremiumSourceBusinessFeature *PremiumSourceBusinessFeature `json:"premiumSourceBusinessFeature,omitempty"`
	PremiumSourceStoryFeature    *PremiumSourceStoryFeature    `json:"premiumSourceStoryFeature,omitempty"`
	PremiumSourceLink            *PremiumSourceLink            `json:"premiumSourceLink,omitempty"`
	PremiumSourceSettings        *PremiumSourceSettings        `json:"premiumSourceSettings,omitempty"`
}

func (t *PremiumSource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "premiumSourceLimitExceeded":
		t.PremiumSourceLimitExceeded = new(PremiumSourceLimitExceeded)
		return json.Unmarshal(b, t.PremiumSourceLimitExceeded)
	case "premiumSourceFeature":
		t.PremiumSourceFeature = new(PremiumSourceFeature)
		return json.Unmarshal(b, t.PremiumSourceFeature)
	case "premiumSourceBusinessFeature":
		t.PremiumSourceBusinessFeature = new(PremiumSourceBusinessFeature)
		return json.Unmarshal(b, t.PremiumSourceBusinessFeature)
	case "premiumSourceStoryFeature":
		t.PremiumSourceStoryFeature = new(PremiumSourceStoryFeature)
		return json.Unmarshal(b, t.PremiumSourceStoryFeature)
	case "premiumSourceLink":
		t.PremiumSourceLink = new(PremiumSourceLink)
		return json.Unmarshal(b, t.PremiumSourceLink)
	case "premiumSourceSettings":
		t.PremiumSourceSettings = new(PremiumSourceSettings)
		return json.Unmarshal(b, t.PremiumSourceSettings)
	}
	return nil
}

func (t *PremiumSource) MarshalJSON() ([]byte, error) {
	if t.PremiumSourceLimitExceeded != nil {
		return json.Marshal(t.PremiumSourceLimitExceeded)
	}
	if t.PremiumSourceFeature != nil {
		return json.Marshal(t.PremiumSourceFeature)
	}
	if t.PremiumSourceBusinessFeature != nil {
		return json.Marshal(t.PremiumSourceBusinessFeature)
	}
	if t.PremiumSourceStoryFeature != nil {
		return json.Marshal(t.PremiumSourceStoryFeature)
	}
	if t.PremiumSourceLink != nil {
		return json.Marshal(t.PremiumSourceLink)
	}
	if t.PremiumSourceSettings != nil {
		return json.Marshal(t.PremiumSourceSettings)
	}
	return nil, fmt.Errorf("no value set for PremiumSource")
}

// MessageFileType Contains information about a file with messages exported from another app
type MessageFileType struct {
	MessageFileTypePrivate *MessageFileTypePrivate `json:"messageFileTypePrivate,omitempty"`
	MessageFileTypeGroup   *MessageFileTypeGroup   `json:"messageFileTypeGroup,omitempty"`
	MessageFileTypeUnknown *MessageFileTypeUnknown `json:"messageFileTypeUnknown,omitempty"`
}

func (t *MessageFileType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageFileTypePrivate":
		t.MessageFileTypePrivate = new(MessageFileTypePrivate)
		return json.Unmarshal(b, t.MessageFileTypePrivate)
	case "messageFileTypeGroup":
		t.MessageFileTypeGroup = new(MessageFileTypeGroup)
		return json.Unmarshal(b, t.MessageFileTypeGroup)
	case "messageFileTypeUnknown":
		t.MessageFileTypeUnknown = new(MessageFileTypeUnknown)
		return json.Unmarshal(b, t.MessageFileTypeUnknown)
	}
	return nil
}

func (t *MessageFileType) MarshalJSON() ([]byte, error) {
	if t.MessageFileTypePrivate != nil {
		return json.Marshal(t.MessageFileTypePrivate)
	}
	if t.MessageFileTypeGroup != nil {
		return json.Marshal(t.MessageFileTypeGroup)
	}
	if t.MessageFileTypeUnknown != nil {
		return json.Marshal(t.MessageFileTypeUnknown)
	}
	return nil, fmt.Errorf("no value set for MessageFileType")
}

// TopChatCategory Represents the categories of chats for which a list of frequently used chats can be retrieved
type TopChatCategory struct {
	TopChatCategoryUsers        *TopChatCategoryUsers        `json:"topChatCategoryUsers,omitempty"`
	TopChatCategoryBots         *TopChatCategoryBots         `json:"topChatCategoryBots,omitempty"`
	TopChatCategoryGroups       *TopChatCategoryGroups       `json:"topChatCategoryGroups,omitempty"`
	TopChatCategoryChannels     *TopChatCategoryChannels     `json:"topChatCategoryChannels,omitempty"`
	TopChatCategoryInlineBots   *TopChatCategoryInlineBots   `json:"topChatCategoryInlineBots,omitempty"`
	TopChatCategoryWebAppBots   *TopChatCategoryWebAppBots   `json:"topChatCategoryWebAppBots,omitempty"`
	TopChatCategoryCalls        *TopChatCategoryCalls        `json:"topChatCategoryCalls,omitempty"`
	TopChatCategoryForwardChats *TopChatCategoryForwardChats `json:"topChatCategoryForwardChats,omitempty"`
}

func (t *TopChatCategory) Type() string {
	return ""
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
	switch typeObj.Type {
	case "topChatCategoryUsers":
		t.TopChatCategoryUsers = new(TopChatCategoryUsers)
		return json.Unmarshal(b, t.TopChatCategoryUsers)
	case "topChatCategoryBots":
		t.TopChatCategoryBots = new(TopChatCategoryBots)
		return json.Unmarshal(b, t.TopChatCategoryBots)
	case "topChatCategoryGroups":
		t.TopChatCategoryGroups = new(TopChatCategoryGroups)
		return json.Unmarshal(b, t.TopChatCategoryGroups)
	case "topChatCategoryChannels":
		t.TopChatCategoryChannels = new(TopChatCategoryChannels)
		return json.Unmarshal(b, t.TopChatCategoryChannels)
	case "topChatCategoryInlineBots":
		t.TopChatCategoryInlineBots = new(TopChatCategoryInlineBots)
		return json.Unmarshal(b, t.TopChatCategoryInlineBots)
	case "topChatCategoryWebAppBots":
		t.TopChatCategoryWebAppBots = new(TopChatCategoryWebAppBots)
		return json.Unmarshal(b, t.TopChatCategoryWebAppBots)
	case "topChatCategoryCalls":
		t.TopChatCategoryCalls = new(TopChatCategoryCalls)
		return json.Unmarshal(b, t.TopChatCategoryCalls)
	case "topChatCategoryForwardChats":
		t.TopChatCategoryForwardChats = new(TopChatCategoryForwardChats)
		return json.Unmarshal(b, t.TopChatCategoryForwardChats)
	}
	return nil
}

func (t *TopChatCategory) MarshalJSON() ([]byte, error) {
	if t.TopChatCategoryUsers != nil {
		return json.Marshal(t.TopChatCategoryUsers)
	}
	if t.TopChatCategoryBots != nil {
		return json.Marshal(t.TopChatCategoryBots)
	}
	if t.TopChatCategoryGroups != nil {
		return json.Marshal(t.TopChatCategoryGroups)
	}
	if t.TopChatCategoryChannels != nil {
		return json.Marshal(t.TopChatCategoryChannels)
	}
	if t.TopChatCategoryInlineBots != nil {
		return json.Marshal(t.TopChatCategoryInlineBots)
	}
	if t.TopChatCategoryWebAppBots != nil {
		return json.Marshal(t.TopChatCategoryWebAppBots)
	}
	if t.TopChatCategoryCalls != nil {
		return json.Marshal(t.TopChatCategoryCalls)
	}
	if t.TopChatCategoryForwardChats != nil {
		return json.Marshal(t.TopChatCategoryForwardChats)
	}
	return nil, fmt.Errorf("no value set for TopChatCategory")
}

// CanSendGiftResult Describes whether a gift can be sent now by the current user
type CanSendGiftResult struct {
	CanSendGiftResultOk   *CanSendGiftResultOk   `json:"canSendGiftResultOk,omitempty"`
	CanSendGiftResultFail *CanSendGiftResultFail `json:"canSendGiftResultFail,omitempty"`
}

func (t *CanSendGiftResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "canSendGiftResultOk":
		t.CanSendGiftResultOk = new(CanSendGiftResultOk)
		return json.Unmarshal(b, t.CanSendGiftResultOk)
	case "canSendGiftResultFail":
		t.CanSendGiftResultFail = new(CanSendGiftResultFail)
		return json.Unmarshal(b, t.CanSendGiftResultFail)
	}
	return nil
}

func (t *CanSendGiftResult) MarshalJSON() ([]byte, error) {
	if t.CanSendGiftResultOk != nil {
		return json.Marshal(t.CanSendGiftResultOk)
	}
	if t.CanSendGiftResultFail != nil {
		return json.Marshal(t.CanSendGiftResultFail)
	}
	return nil, fmt.Errorf("no value set for CanSendGiftResult")
}

// PageBlockHorizontalAlignment Describes a horizontal alignment of a table cell content
type PageBlockHorizontalAlignment struct {
	PageBlockHorizontalAlignmentLeft   *PageBlockHorizontalAlignmentLeft   `json:"pageBlockHorizontalAlignmentLeft,omitempty"`
	PageBlockHorizontalAlignmentCenter *PageBlockHorizontalAlignmentCenter `json:"pageBlockHorizontalAlignmentCenter,omitempty"`
	PageBlockHorizontalAlignmentRight  *PageBlockHorizontalAlignmentRight  `json:"pageBlockHorizontalAlignmentRight,omitempty"`
}

func (t *PageBlockHorizontalAlignment) Type() string {
	return ""
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
	switch typeObj.Type {
	case "pageBlockHorizontalAlignmentLeft":
		t.PageBlockHorizontalAlignmentLeft = new(PageBlockHorizontalAlignmentLeft)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentLeft)
	case "pageBlockHorizontalAlignmentCenter":
		t.PageBlockHorizontalAlignmentCenter = new(PageBlockHorizontalAlignmentCenter)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentCenter)
	case "pageBlockHorizontalAlignmentRight":
		t.PageBlockHorizontalAlignmentRight = new(PageBlockHorizontalAlignmentRight)
		return json.Unmarshal(b, t.PageBlockHorizontalAlignmentRight)
	}
	return nil
}

func (t *PageBlockHorizontalAlignment) MarshalJSON() ([]byte, error) {
	if t.PageBlockHorizontalAlignmentLeft != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentLeft)
	}
	if t.PageBlockHorizontalAlignmentCenter != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentCenter)
	}
	if t.PageBlockHorizontalAlignmentRight != nil {
		return json.Marshal(t.PageBlockHorizontalAlignmentRight)
	}
	return nil, fmt.Errorf("no value set for PageBlockHorizontalAlignment")
}

// InputPassportElement Contains information about a Telegram Passport element to be saved
type InputPassportElement struct {
	InputPassportElementPersonalDetails       *InputPassportElementPersonalDetails       `json:"inputPassportElementPersonalDetails,omitempty"`
	InputPassportElementPassport              *InputPassportElementPassport              `json:"inputPassportElementPassport,omitempty"`
	InputPassportElementDriverLicense         *InputPassportElementDriverLicense         `json:"inputPassportElementDriverLicense,omitempty"`
	InputPassportElementIdentityCard          *InputPassportElementIdentityCard          `json:"inputPassportElementIdentityCard,omitempty"`
	InputPassportElementInternalPassport      *InputPassportElementInternalPassport      `json:"inputPassportElementInternalPassport,omitempty"`
	InputPassportElementAddress               *InputPassportElementAddress               `json:"inputPassportElementAddress,omitempty"`
	InputPassportElementUtilityBill           *InputPassportElementUtilityBill           `json:"inputPassportElementUtilityBill,omitempty"`
	InputPassportElementBankStatement         *InputPassportElementBankStatement         `json:"inputPassportElementBankStatement,omitempty"`
	InputPassportElementRentalAgreement       *InputPassportElementRentalAgreement       `json:"inputPassportElementRentalAgreement,omitempty"`
	InputPassportElementPassportRegistration  *InputPassportElementPassportRegistration  `json:"inputPassportElementPassportRegistration,omitempty"`
	InputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration `json:"inputPassportElementTemporaryRegistration,omitempty"`
	InputPassportElementPhoneNumber           *InputPassportElementPhoneNumber           `json:"inputPassportElementPhoneNumber,omitempty"`
	InputPassportElementEmailAddress          *InputPassportElementEmailAddress          `json:"inputPassportElementEmailAddress,omitempty"`
}

func (t *InputPassportElement) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputPassportElementPersonalDetails":
		t.InputPassportElementPersonalDetails = new(InputPassportElementPersonalDetails)
		return json.Unmarshal(b, t.InputPassportElementPersonalDetails)
	case "inputPassportElementPassport":
		t.InputPassportElementPassport = new(InputPassportElementPassport)
		return json.Unmarshal(b, t.InputPassportElementPassport)
	case "inputPassportElementDriverLicense":
		t.InputPassportElementDriverLicense = new(InputPassportElementDriverLicense)
		return json.Unmarshal(b, t.InputPassportElementDriverLicense)
	case "inputPassportElementIdentityCard":
		t.InputPassportElementIdentityCard = new(InputPassportElementIdentityCard)
		return json.Unmarshal(b, t.InputPassportElementIdentityCard)
	case "inputPassportElementInternalPassport":
		t.InputPassportElementInternalPassport = new(InputPassportElementInternalPassport)
		return json.Unmarshal(b, t.InputPassportElementInternalPassport)
	case "inputPassportElementAddress":
		t.InputPassportElementAddress = new(InputPassportElementAddress)
		return json.Unmarshal(b, t.InputPassportElementAddress)
	case "inputPassportElementUtilityBill":
		t.InputPassportElementUtilityBill = new(InputPassportElementUtilityBill)
		return json.Unmarshal(b, t.InputPassportElementUtilityBill)
	case "inputPassportElementBankStatement":
		t.InputPassportElementBankStatement = new(InputPassportElementBankStatement)
		return json.Unmarshal(b, t.InputPassportElementBankStatement)
	case "inputPassportElementRentalAgreement":
		t.InputPassportElementRentalAgreement = new(InputPassportElementRentalAgreement)
		return json.Unmarshal(b, t.InputPassportElementRentalAgreement)
	case "inputPassportElementPassportRegistration":
		t.InputPassportElementPassportRegistration = new(InputPassportElementPassportRegistration)
		return json.Unmarshal(b, t.InputPassportElementPassportRegistration)
	case "inputPassportElementTemporaryRegistration":
		t.InputPassportElementTemporaryRegistration = new(InputPassportElementTemporaryRegistration)
		return json.Unmarshal(b, t.InputPassportElementTemporaryRegistration)
	case "inputPassportElementPhoneNumber":
		t.InputPassportElementPhoneNumber = new(InputPassportElementPhoneNumber)
		return json.Unmarshal(b, t.InputPassportElementPhoneNumber)
	case "inputPassportElementEmailAddress":
		t.InputPassportElementEmailAddress = new(InputPassportElementEmailAddress)
		return json.Unmarshal(b, t.InputPassportElementEmailAddress)
	}
	return nil
}

func (t *InputPassportElement) MarshalJSON() ([]byte, error) {
	if t.InputPassportElementPersonalDetails != nil {
		return json.Marshal(t.InputPassportElementPersonalDetails)
	}
	if t.InputPassportElementPassport != nil {
		return json.Marshal(t.InputPassportElementPassport)
	}
	if t.InputPassportElementDriverLicense != nil {
		return json.Marshal(t.InputPassportElementDriverLicense)
	}
	if t.InputPassportElementIdentityCard != nil {
		return json.Marshal(t.InputPassportElementIdentityCard)
	}
	if t.InputPassportElementInternalPassport != nil {
		return json.Marshal(t.InputPassportElementInternalPassport)
	}
	if t.InputPassportElementAddress != nil {
		return json.Marshal(t.InputPassportElementAddress)
	}
	if t.InputPassportElementUtilityBill != nil {
		return json.Marshal(t.InputPassportElementUtilityBill)
	}
	if t.InputPassportElementBankStatement != nil {
		return json.Marshal(t.InputPassportElementBankStatement)
	}
	if t.InputPassportElementRentalAgreement != nil {
		return json.Marshal(t.InputPassportElementRentalAgreement)
	}
	if t.InputPassportElementPassportRegistration != nil {
		return json.Marshal(t.InputPassportElementPassportRegistration)
	}
	if t.InputPassportElementTemporaryRegistration != nil {
		return json.Marshal(t.InputPassportElementTemporaryRegistration)
	}
	if t.InputPassportElementPhoneNumber != nil {
		return json.Marshal(t.InputPassportElementPhoneNumber)
	}
	if t.InputPassportElementEmailAddress != nil {
		return json.Marshal(t.InputPassportElementEmailAddress)
	}
	return nil, fmt.Errorf("no value set for InputPassportElement")
}

// SearchMessagesChatTypeFilter Represents a filter for type of the chats in which to search messages
type SearchMessagesChatTypeFilter struct {
	SearchMessagesChatTypeFilterPrivate *SearchMessagesChatTypeFilterPrivate `json:"searchMessagesChatTypeFilterPrivate,omitempty"`
	SearchMessagesChatTypeFilterGroup   *SearchMessagesChatTypeFilterGroup   `json:"searchMessagesChatTypeFilterGroup,omitempty"`
	SearchMessagesChatTypeFilterChannel *SearchMessagesChatTypeFilterChannel `json:"searchMessagesChatTypeFilterChannel,omitempty"`
}

func (t *SearchMessagesChatTypeFilter) Type() string {
	return ""
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
	switch typeObj.Type {
	case "searchMessagesChatTypeFilterPrivate":
		t.SearchMessagesChatTypeFilterPrivate = new(SearchMessagesChatTypeFilterPrivate)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterPrivate)
	case "searchMessagesChatTypeFilterGroup":
		t.SearchMessagesChatTypeFilterGroup = new(SearchMessagesChatTypeFilterGroup)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterGroup)
	case "searchMessagesChatTypeFilterChannel":
		t.SearchMessagesChatTypeFilterChannel = new(SearchMessagesChatTypeFilterChannel)
		return json.Unmarshal(b, t.SearchMessagesChatTypeFilterChannel)
	}
	return nil
}

func (t *SearchMessagesChatTypeFilter) MarshalJSON() ([]byte, error) {
	if t.SearchMessagesChatTypeFilterPrivate != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterPrivate)
	}
	if t.SearchMessagesChatTypeFilterGroup != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterGroup)
	}
	if t.SearchMessagesChatTypeFilterChannel != nil {
		return json.Marshal(t.SearchMessagesChatTypeFilterChannel)
	}
	return nil, fmt.Errorf("no value set for SearchMessagesChatTypeFilter")
}

// StoryPrivacySettings Describes privacy settings of a story
type StoryPrivacySettings struct {
	StoryPrivacySettingsEveryone      *StoryPrivacySettingsEveryone      `json:"storyPrivacySettingsEveryone,omitempty"`
	StoryPrivacySettingsContacts      *StoryPrivacySettingsContacts      `json:"storyPrivacySettingsContacts,omitempty"`
	StoryPrivacySettingsCloseFriends  *StoryPrivacySettingsCloseFriends  `json:"storyPrivacySettingsCloseFriends,omitempty"`
	StoryPrivacySettingsSelectedUsers *StoryPrivacySettingsSelectedUsers `json:"storyPrivacySettingsSelectedUsers,omitempty"`
}

func (t *StoryPrivacySettings) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyPrivacySettingsEveryone":
		t.StoryPrivacySettingsEveryone = new(StoryPrivacySettingsEveryone)
		return json.Unmarshal(b, t.StoryPrivacySettingsEveryone)
	case "storyPrivacySettingsContacts":
		t.StoryPrivacySettingsContacts = new(StoryPrivacySettingsContacts)
		return json.Unmarshal(b, t.StoryPrivacySettingsContacts)
	case "storyPrivacySettingsCloseFriends":
		t.StoryPrivacySettingsCloseFriends = new(StoryPrivacySettingsCloseFriends)
		return json.Unmarshal(b, t.StoryPrivacySettingsCloseFriends)
	case "storyPrivacySettingsSelectedUsers":
		t.StoryPrivacySettingsSelectedUsers = new(StoryPrivacySettingsSelectedUsers)
		return json.Unmarshal(b, t.StoryPrivacySettingsSelectedUsers)
	}
	return nil
}

func (t *StoryPrivacySettings) MarshalJSON() ([]byte, error) {
	if t.StoryPrivacySettingsEveryone != nil {
		return json.Marshal(t.StoryPrivacySettingsEveryone)
	}
	if t.StoryPrivacySettingsContacts != nil {
		return json.Marshal(t.StoryPrivacySettingsContacts)
	}
	if t.StoryPrivacySettingsCloseFriends != nil {
		return json.Marshal(t.StoryPrivacySettingsCloseFriends)
	}
	if t.StoryPrivacySettingsSelectedUsers != nil {
		return json.Marshal(t.StoryPrivacySettingsSelectedUsers)
	}
	return nil, fmt.Errorf("no value set for StoryPrivacySettings")
}

// MessageReplyTo Contains information about the message or the story a message is replying to
type MessageReplyTo struct {
	MessageReplyToMessage *MessageReplyToMessage `json:"messageReplyToMessage,omitempty"`
	MessageReplyToStory   *MessageReplyToStory   `json:"messageReplyToStory,omitempty"`
}

func (t *MessageReplyTo) Type() string {
	return ""
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

// MessageSendingState Contains information about the sending state of the message
type MessageSendingState struct {
	MessageSendingStatePending *MessageSendingStatePending `json:"messageSendingStatePending,omitempty"`
	MessageSendingStateFailed  *MessageSendingStateFailed  `json:"messageSendingStateFailed,omitempty"`
}

func (t *MessageSendingState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageSendingStatePending":
		t.MessageSendingStatePending = new(MessageSendingStatePending)
		return json.Unmarshal(b, t.MessageSendingStatePending)
	case "messageSendingStateFailed":
		t.MessageSendingStateFailed = new(MessageSendingStateFailed)
		return json.Unmarshal(b, t.MessageSendingStateFailed)
	}
	return nil
}

func (t *MessageSendingState) MarshalJSON() ([]byte, error) {
	if t.MessageSendingStatePending != nil {
		return json.Marshal(t.MessageSendingStatePending)
	}
	if t.MessageSendingStateFailed != nil {
		return json.Marshal(t.MessageSendingStateFailed)
	}
	return nil, fmt.Errorf("no value set for MessageSendingState")
}

// RichText Describes a formatted text object
type RichText struct {
	RichTextPlain         *RichTextPlain         `json:"richTextPlain,omitempty"`
	RichTextBold          *RichTextBold          `json:"richTextBold,omitempty"`
	RichTextItalic        *RichTextItalic        `json:"richTextItalic,omitempty"`
	RichTextUnderline     *RichTextUnderline     `json:"richTextUnderline,omitempty"`
	RichTextStrikethrough *RichTextStrikethrough `json:"richTextStrikethrough,omitempty"`
	RichTextFixed         *RichTextFixed         `json:"richTextFixed,omitempty"`
	RichTextUrl           *RichTextUrl           `json:"richTextUrl,omitempty"`
	RichTextEmailAddress  *RichTextEmailAddress  `json:"richTextEmailAddress,omitempty"`
	RichTextSubscript     *RichTextSubscript     `json:"richTextSubscript,omitempty"`
	RichTextSuperscript   *RichTextSuperscript   `json:"richTextSuperscript,omitempty"`
	RichTextMarked        *RichTextMarked        `json:"richTextMarked,omitempty"`
	RichTextPhoneNumber   *RichTextPhoneNumber   `json:"richTextPhoneNumber,omitempty"`
	RichTextIcon          *RichTextIcon          `json:"richTextIcon,omitempty"`
	RichTextReference     *RichTextReference     `json:"richTextReference,omitempty"`
	RichTextAnchor        *RichTextAnchor        `json:"richTextAnchor,omitempty"`
	RichTextAnchorLink    *RichTextAnchorLink    `json:"richTextAnchorLink,omitempty"`
	RichTexts             *RichTexts             `json:"richTexts,omitempty"`
}

func (t *RichText) Type() string {
	return ""
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
	switch typeObj.Type {
	case "richTextPlain":
		t.RichTextPlain = new(RichTextPlain)
		return json.Unmarshal(b, t.RichTextPlain)
	case "richTextBold":
		t.RichTextBold = new(RichTextBold)
		return json.Unmarshal(b, t.RichTextBold)
	case "richTextItalic":
		t.RichTextItalic = new(RichTextItalic)
		return json.Unmarshal(b, t.RichTextItalic)
	case "richTextUnderline":
		t.RichTextUnderline = new(RichTextUnderline)
		return json.Unmarshal(b, t.RichTextUnderline)
	case "richTextStrikethrough":
		t.RichTextStrikethrough = new(RichTextStrikethrough)
		return json.Unmarshal(b, t.RichTextStrikethrough)
	case "richTextFixed":
		t.RichTextFixed = new(RichTextFixed)
		return json.Unmarshal(b, t.RichTextFixed)
	case "richTextUrl":
		t.RichTextUrl = new(RichTextUrl)
		return json.Unmarshal(b, t.RichTextUrl)
	case "richTextEmailAddress":
		t.RichTextEmailAddress = new(RichTextEmailAddress)
		return json.Unmarshal(b, t.RichTextEmailAddress)
	case "richTextSubscript":
		t.RichTextSubscript = new(RichTextSubscript)
		return json.Unmarshal(b, t.RichTextSubscript)
	case "richTextSuperscript":
		t.RichTextSuperscript = new(RichTextSuperscript)
		return json.Unmarshal(b, t.RichTextSuperscript)
	case "richTextMarked":
		t.RichTextMarked = new(RichTextMarked)
		return json.Unmarshal(b, t.RichTextMarked)
	case "richTextPhoneNumber":
		t.RichTextPhoneNumber = new(RichTextPhoneNumber)
		return json.Unmarshal(b, t.RichTextPhoneNumber)
	case "richTextIcon":
		t.RichTextIcon = new(RichTextIcon)
		return json.Unmarshal(b, t.RichTextIcon)
	case "richTextReference":
		t.RichTextReference = new(RichTextReference)
		return json.Unmarshal(b, t.RichTextReference)
	case "richTextAnchor":
		t.RichTextAnchor = new(RichTextAnchor)
		return json.Unmarshal(b, t.RichTextAnchor)
	case "richTextAnchorLink":
		t.RichTextAnchorLink = new(RichTextAnchorLink)
		return json.Unmarshal(b, t.RichTextAnchorLink)
	case "richTexts":
		t.RichTexts = new(RichTexts)
		return json.Unmarshal(b, t.RichTexts)
	}
	return nil
}

func (t *RichText) MarshalJSON() ([]byte, error) {
	if t.RichTextPlain != nil {
		return json.Marshal(t.RichTextPlain)
	}
	if t.RichTextBold != nil {
		return json.Marshal(t.RichTextBold)
	}
	if t.RichTextItalic != nil {
		return json.Marshal(t.RichTextItalic)
	}
	if t.RichTextUnderline != nil {
		return json.Marshal(t.RichTextUnderline)
	}
	if t.RichTextStrikethrough != nil {
		return json.Marshal(t.RichTextStrikethrough)
	}
	if t.RichTextFixed != nil {
		return json.Marshal(t.RichTextFixed)
	}
	if t.RichTextUrl != nil {
		return json.Marshal(t.RichTextUrl)
	}
	if t.RichTextEmailAddress != nil {
		return json.Marshal(t.RichTextEmailAddress)
	}
	if t.RichTextSubscript != nil {
		return json.Marshal(t.RichTextSubscript)
	}
	if t.RichTextSuperscript != nil {
		return json.Marshal(t.RichTextSuperscript)
	}
	if t.RichTextMarked != nil {
		return json.Marshal(t.RichTextMarked)
	}
	if t.RichTextPhoneNumber != nil {
		return json.Marshal(t.RichTextPhoneNumber)
	}
	if t.RichTextIcon != nil {
		return json.Marshal(t.RichTextIcon)
	}
	if t.RichTextReference != nil {
		return json.Marshal(t.RichTextReference)
	}
	if t.RichTextAnchor != nil {
		return json.Marshal(t.RichTextAnchor)
	}
	if t.RichTextAnchorLink != nil {
		return json.Marshal(t.RichTextAnchorLink)
	}
	if t.RichTexts != nil {
		return json.Marshal(t.RichTexts)
	}
	return nil, fmt.Errorf("no value set for RichText")
}

// PaymentFormType Describes type of payment form
type PaymentFormType struct {
	PaymentFormTypeRegular          *PaymentFormTypeRegular          `json:"paymentFormTypeRegular,omitempty"`
	PaymentFormTypeStars            *PaymentFormTypeStars            `json:"paymentFormTypeStars,omitempty"`
	PaymentFormTypeStarSubscription *PaymentFormTypeStarSubscription `json:"paymentFormTypeStarSubscription,omitempty"`
}

func (t *PaymentFormType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "paymentFormTypeRegular":
		t.PaymentFormTypeRegular = new(PaymentFormTypeRegular)
		return json.Unmarshal(b, t.PaymentFormTypeRegular)
	case "paymentFormTypeStars":
		t.PaymentFormTypeStars = new(PaymentFormTypeStars)
		return json.Unmarshal(b, t.PaymentFormTypeStars)
	case "paymentFormTypeStarSubscription":
		t.PaymentFormTypeStarSubscription = new(PaymentFormTypeStarSubscription)
		return json.Unmarshal(b, t.PaymentFormTypeStarSubscription)
	}
	return nil
}

func (t *PaymentFormType) MarshalJSON() ([]byte, error) {
	if t.PaymentFormTypeRegular != nil {
		return json.Marshal(t.PaymentFormTypeRegular)
	}
	if t.PaymentFormTypeStars != nil {
		return json.Marshal(t.PaymentFormTypeStars)
	}
	if t.PaymentFormTypeStarSubscription != nil {
		return json.Marshal(t.PaymentFormTypeStarSubscription)
	}
	return nil, fmt.Errorf("no value set for PaymentFormType")
}

// CallDiscardReason Describes the reason why a call was discarded
type CallDiscardReason struct {
	CallDiscardReasonEmpty              *CallDiscardReasonEmpty              `json:"callDiscardReasonEmpty,omitempty"`
	CallDiscardReasonMissed             *CallDiscardReasonMissed             `json:"callDiscardReasonMissed,omitempty"`
	CallDiscardReasonDeclined           *CallDiscardReasonDeclined           `json:"callDiscardReasonDeclined,omitempty"`
	CallDiscardReasonDisconnected       *CallDiscardReasonDisconnected       `json:"callDiscardReasonDisconnected,omitempty"`
	CallDiscardReasonHungUp             *CallDiscardReasonHungUp             `json:"callDiscardReasonHungUp,omitempty"`
	CallDiscardReasonUpgradeToGroupCall *CallDiscardReasonUpgradeToGroupCall `json:"callDiscardReasonUpgradeToGroupCall,omitempty"`
}

func (t *CallDiscardReason) Type() string {
	return ""
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
	switch typeObj.Type {
	case "callDiscardReasonEmpty":
		t.CallDiscardReasonEmpty = new(CallDiscardReasonEmpty)
		return json.Unmarshal(b, t.CallDiscardReasonEmpty)
	case "callDiscardReasonMissed":
		t.CallDiscardReasonMissed = new(CallDiscardReasonMissed)
		return json.Unmarshal(b, t.CallDiscardReasonMissed)
	case "callDiscardReasonDeclined":
		t.CallDiscardReasonDeclined = new(CallDiscardReasonDeclined)
		return json.Unmarshal(b, t.CallDiscardReasonDeclined)
	case "callDiscardReasonDisconnected":
		t.CallDiscardReasonDisconnected = new(CallDiscardReasonDisconnected)
		return json.Unmarshal(b, t.CallDiscardReasonDisconnected)
	case "callDiscardReasonHungUp":
		t.CallDiscardReasonHungUp = new(CallDiscardReasonHungUp)
		return json.Unmarshal(b, t.CallDiscardReasonHungUp)
	case "callDiscardReasonUpgradeToGroupCall":
		t.CallDiscardReasonUpgradeToGroupCall = new(CallDiscardReasonUpgradeToGroupCall)
		return json.Unmarshal(b, t.CallDiscardReasonUpgradeToGroupCall)
	}
	return nil
}

func (t *CallDiscardReason) MarshalJSON() ([]byte, error) {
	if t.CallDiscardReasonEmpty != nil {
		return json.Marshal(t.CallDiscardReasonEmpty)
	}
	if t.CallDiscardReasonMissed != nil {
		return json.Marshal(t.CallDiscardReasonMissed)
	}
	if t.CallDiscardReasonDeclined != nil {
		return json.Marshal(t.CallDiscardReasonDeclined)
	}
	if t.CallDiscardReasonDisconnected != nil {
		return json.Marshal(t.CallDiscardReasonDisconnected)
	}
	if t.CallDiscardReasonHungUp != nil {
		return json.Marshal(t.CallDiscardReasonHungUp)
	}
	if t.CallDiscardReasonUpgradeToGroupCall != nil {
		return json.Marshal(t.CallDiscardReasonUpgradeToGroupCall)
	}
	return nil, fmt.Errorf("no value set for CallDiscardReason")
}

// BackgroundFill Describes a fill of a background
type BackgroundFill struct {
	BackgroundFillSolid            *BackgroundFillSolid            `json:"backgroundFillSolid,omitempty"`
	BackgroundFillGradient         *BackgroundFillGradient         `json:"backgroundFillGradient,omitempty"`
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient `json:"backgroundFillFreeformGradient,omitempty"`
}

func (t *BackgroundFill) Type() string {
	return ""
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
	switch typeObj.Type {
	case "backgroundFillSolid":
		t.BackgroundFillSolid = new(BackgroundFillSolid)
		return json.Unmarshal(b, t.BackgroundFillSolid)
	case "backgroundFillGradient":
		t.BackgroundFillGradient = new(BackgroundFillGradient)
		return json.Unmarshal(b, t.BackgroundFillGradient)
	case "backgroundFillFreeformGradient":
		t.BackgroundFillFreeformGradient = new(BackgroundFillFreeformGradient)
		return json.Unmarshal(b, t.BackgroundFillFreeformGradient)
	}
	return nil
}

func (t *BackgroundFill) MarshalJSON() ([]byte, error) {
	if t.BackgroundFillSolid != nil {
		return json.Marshal(t.BackgroundFillSolid)
	}
	if t.BackgroundFillGradient != nil {
		return json.Marshal(t.BackgroundFillGradient)
	}
	if t.BackgroundFillFreeformGradient != nil {
		return json.Marshal(t.BackgroundFillFreeformGradient)
	}
	return nil, fmt.Errorf("no value set for BackgroundFill")
}

// CheckChatUsernameResult Represents result of checking whether a username can be set for a chat
type CheckChatUsernameResult struct {
	CheckChatUsernameResultOk                      *CheckChatUsernameResultOk                      `json:"checkChatUsernameResultOk,omitempty"`
	CheckChatUsernameResultUsernameInvalid         *CheckChatUsernameResultUsernameInvalid         `json:"checkChatUsernameResultUsernameInvalid,omitempty"`
	CheckChatUsernameResultUsernameOccupied        *CheckChatUsernameResultUsernameOccupied        `json:"checkChatUsernameResultUsernameOccupied,omitempty"`
	CheckChatUsernameResultUsernamePurchasable     *CheckChatUsernameResultUsernamePurchasable     `json:"checkChatUsernameResultUsernamePurchasable,omitempty"`
	CheckChatUsernameResultPublicChatsTooMany      *CheckChatUsernameResultPublicChatsTooMany      `json:"checkChatUsernameResultPublicChatsTooMany,omitempty"`
	CheckChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable `json:"checkChatUsernameResultPublicGroupsUnavailable,omitempty"`
}

func (t *CheckChatUsernameResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "checkChatUsernameResultOk":
		t.CheckChatUsernameResultOk = new(CheckChatUsernameResultOk)
		return json.Unmarshal(b, t.CheckChatUsernameResultOk)
	case "checkChatUsernameResultUsernameInvalid":
		t.CheckChatUsernameResultUsernameInvalid = new(CheckChatUsernameResultUsernameInvalid)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernameInvalid)
	case "checkChatUsernameResultUsernameOccupied":
		t.CheckChatUsernameResultUsernameOccupied = new(CheckChatUsernameResultUsernameOccupied)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernameOccupied)
	case "checkChatUsernameResultUsernamePurchasable":
		t.CheckChatUsernameResultUsernamePurchasable = new(CheckChatUsernameResultUsernamePurchasable)
		return json.Unmarshal(b, t.CheckChatUsernameResultUsernamePurchasable)
	case "checkChatUsernameResultPublicChatsTooMany":
		t.CheckChatUsernameResultPublicChatsTooMany = new(CheckChatUsernameResultPublicChatsTooMany)
		return json.Unmarshal(b, t.CheckChatUsernameResultPublicChatsTooMany)
	case "checkChatUsernameResultPublicGroupsUnavailable":
		t.CheckChatUsernameResultPublicGroupsUnavailable = new(CheckChatUsernameResultPublicGroupsUnavailable)
		return json.Unmarshal(b, t.CheckChatUsernameResultPublicGroupsUnavailable)
	}
	return nil
}

func (t *CheckChatUsernameResult) MarshalJSON() ([]byte, error) {
	if t.CheckChatUsernameResultOk != nil {
		return json.Marshal(t.CheckChatUsernameResultOk)
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
	if t.CheckChatUsernameResultPublicChatsTooMany != nil {
		return json.Marshal(t.CheckChatUsernameResultPublicChatsTooMany)
	}
	if t.CheckChatUsernameResultPublicGroupsUnavailable != nil {
		return json.Marshal(t.CheckChatUsernameResultPublicGroupsUnavailable)
	}
	return nil, fmt.Errorf("no value set for CheckChatUsernameResult")
}

// StatisticalGraph Describes a statistical graph
type StatisticalGraph struct {
	StatisticalGraphData  *StatisticalGraphData  `json:"statisticalGraphData,omitempty"`
	StatisticalGraphAsync *StatisticalGraphAsync `json:"statisticalGraphAsync,omitempty"`
	StatisticalGraphError *StatisticalGraphError `json:"statisticalGraphError,omitempty"`
}

func (t *StatisticalGraph) Type() string {
	return ""
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
	switch typeObj.Type {
	case "statisticalGraphData":
		t.StatisticalGraphData = new(StatisticalGraphData)
		return json.Unmarshal(b, t.StatisticalGraphData)
	case "statisticalGraphAsync":
		t.StatisticalGraphAsync = new(StatisticalGraphAsync)
		return json.Unmarshal(b, t.StatisticalGraphAsync)
	case "statisticalGraphError":
		t.StatisticalGraphError = new(StatisticalGraphError)
		return json.Unmarshal(b, t.StatisticalGraphError)
	}
	return nil
}

func (t *StatisticalGraph) MarshalJSON() ([]byte, error) {
	if t.StatisticalGraphData != nil {
		return json.Marshal(t.StatisticalGraphData)
	}
	if t.StatisticalGraphAsync != nil {
		return json.Marshal(t.StatisticalGraphAsync)
	}
	if t.StatisticalGraphError != nil {
		return json.Marshal(t.StatisticalGraphError)
	}
	return nil, fmt.Errorf("no value set for StatisticalGraph")
}

// PollType Describes the type of poll
type PollType struct {
	PollTypeRegular *PollTypeRegular `json:"pollTypeRegular,omitempty"`
	PollTypeQuiz    *PollTypeQuiz    `json:"pollTypeQuiz,omitempty"`
}

func (t *PollType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "pollTypeRegular":
		t.PollTypeRegular = new(PollTypeRegular)
		return json.Unmarshal(b, t.PollTypeRegular)
	case "pollTypeQuiz":
		t.PollTypeQuiz = new(PollTypeQuiz)
		return json.Unmarshal(b, t.PollTypeQuiz)
	}
	return nil
}

func (t *PollType) MarshalJSON() ([]byte, error) {
	if t.PollTypeRegular != nil {
		return json.Marshal(t.PollTypeRegular)
	}
	if t.PollTypeQuiz != nil {
		return json.Marshal(t.PollTypeQuiz)
	}
	return nil, fmt.Errorf("no value set for PollType")
}

// GiftResalePrice Describes price of a resold gift
type GiftResalePrice struct {
	GiftResalePriceStar *GiftResalePriceStar `json:"giftResalePriceStar,omitempty"`
	GiftResalePriceTon  *GiftResalePriceTon  `json:"giftResalePriceTon,omitempty"`
}

func (t *GiftResalePrice) Type() string {
	return ""
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

// ReportSponsoredResult Describes result of sponsored message or chat report
type ReportSponsoredResult struct {
	ReportSponsoredResultOk              *ReportSponsoredResultOk              `json:"reportSponsoredResultOk,omitempty"`
	ReportSponsoredResultFailed          *ReportSponsoredResultFailed          `json:"reportSponsoredResultFailed,omitempty"`
	ReportSponsoredResultOptionRequired  *ReportSponsoredResultOptionRequired  `json:"reportSponsoredResultOptionRequired,omitempty"`
	ReportSponsoredResultAdsHidden       *ReportSponsoredResultAdsHidden       `json:"reportSponsoredResultAdsHidden,omitempty"`
	ReportSponsoredResultPremiumRequired *ReportSponsoredResultPremiumRequired `json:"reportSponsoredResultPremiumRequired,omitempty"`
}

func (t *ReportSponsoredResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "reportSponsoredResultOk":
		t.ReportSponsoredResultOk = new(ReportSponsoredResultOk)
		return json.Unmarshal(b, t.ReportSponsoredResultOk)
	case "reportSponsoredResultFailed":
		t.ReportSponsoredResultFailed = new(ReportSponsoredResultFailed)
		return json.Unmarshal(b, t.ReportSponsoredResultFailed)
	case "reportSponsoredResultOptionRequired":
		t.ReportSponsoredResultOptionRequired = new(ReportSponsoredResultOptionRequired)
		return json.Unmarshal(b, t.ReportSponsoredResultOptionRequired)
	case "reportSponsoredResultAdsHidden":
		t.ReportSponsoredResultAdsHidden = new(ReportSponsoredResultAdsHidden)
		return json.Unmarshal(b, t.ReportSponsoredResultAdsHidden)
	case "reportSponsoredResultPremiumRequired":
		t.ReportSponsoredResultPremiumRequired = new(ReportSponsoredResultPremiumRequired)
		return json.Unmarshal(b, t.ReportSponsoredResultPremiumRequired)
	}
	return nil
}

func (t *ReportSponsoredResult) MarshalJSON() ([]byte, error) {
	if t.ReportSponsoredResultOk != nil {
		return json.Marshal(t.ReportSponsoredResultOk)
	}
	if t.ReportSponsoredResultFailed != nil {
		return json.Marshal(t.ReportSponsoredResultFailed)
	}
	if t.ReportSponsoredResultOptionRequired != nil {
		return json.Marshal(t.ReportSponsoredResultOptionRequired)
	}
	if t.ReportSponsoredResultAdsHidden != nil {
		return json.Marshal(t.ReportSponsoredResultAdsHidden)
	}
	if t.ReportSponsoredResultPremiumRequired != nil {
		return json.Marshal(t.ReportSponsoredResultPremiumRequired)
	}
	return nil, fmt.Errorf("no value set for ReportSponsoredResult")
}

// LinkPreviewType Describes type of link preview
type LinkPreviewType struct {
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
	return ""
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

// EmojiCategoryType Describes type of emoji category
type EmojiCategoryType struct {
	EmojiCategoryTypeDefault         *EmojiCategoryTypeDefault         `json:"emojiCategoryTypeDefault,omitempty"`
	EmojiCategoryTypeRegularStickers *EmojiCategoryTypeRegularStickers `json:"emojiCategoryTypeRegularStickers,omitempty"`
	EmojiCategoryTypeEmojiStatus     *EmojiCategoryTypeEmojiStatus     `json:"emojiCategoryTypeEmojiStatus,omitempty"`
	EmojiCategoryTypeChatPhoto       *EmojiCategoryTypeChatPhoto       `json:"emojiCategoryTypeChatPhoto,omitempty"`
}

func (t *EmojiCategoryType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "emojiCategoryTypeDefault":
		t.EmojiCategoryTypeDefault = new(EmojiCategoryTypeDefault)
		return json.Unmarshal(b, t.EmojiCategoryTypeDefault)
	case "emojiCategoryTypeRegularStickers":
		t.EmojiCategoryTypeRegularStickers = new(EmojiCategoryTypeRegularStickers)
		return json.Unmarshal(b, t.EmojiCategoryTypeRegularStickers)
	case "emojiCategoryTypeEmojiStatus":
		t.EmojiCategoryTypeEmojiStatus = new(EmojiCategoryTypeEmojiStatus)
		return json.Unmarshal(b, t.EmojiCategoryTypeEmojiStatus)
	case "emojiCategoryTypeChatPhoto":
		t.EmojiCategoryTypeChatPhoto = new(EmojiCategoryTypeChatPhoto)
		return json.Unmarshal(b, t.EmojiCategoryTypeChatPhoto)
	}
	return nil
}

func (t *EmojiCategoryType) MarshalJSON() ([]byte, error) {
	if t.EmojiCategoryTypeDefault != nil {
		return json.Marshal(t.EmojiCategoryTypeDefault)
	}
	if t.EmojiCategoryTypeRegularStickers != nil {
		return json.Marshal(t.EmojiCategoryTypeRegularStickers)
	}
	if t.EmojiCategoryTypeEmojiStatus != nil {
		return json.Marshal(t.EmojiCategoryTypeEmojiStatus)
	}
	if t.EmojiCategoryTypeChatPhoto != nil {
		return json.Marshal(t.EmojiCategoryTypeChatPhoto)
	}
	return nil, fmt.Errorf("no value set for EmojiCategoryType")
}

// InviteGroupCallParticipantResult Describes result of group call participant invitation
type InviteGroupCallParticipantResult struct {
	InviteGroupCallParticipantResultUserPrivacyRestricted  *InviteGroupCallParticipantResultUserPrivacyRestricted  `json:"inviteGroupCallParticipantResultUserPrivacyRestricted,omitempty"`
	InviteGroupCallParticipantResultUserAlreadyParticipant *InviteGroupCallParticipantResultUserAlreadyParticipant `json:"inviteGroupCallParticipantResultUserAlreadyParticipant,omitempty"`
	InviteGroupCallParticipantResultUserWasBanned          *InviteGroupCallParticipantResultUserWasBanned          `json:"inviteGroupCallParticipantResultUserWasBanned,omitempty"`
	InviteGroupCallParticipantResultSuccess                *InviteGroupCallParticipantResultSuccess                `json:"inviteGroupCallParticipantResultSuccess,omitempty"`
}

func (t *InviteGroupCallParticipantResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inviteGroupCallParticipantResultUserPrivacyRestricted":
		t.InviteGroupCallParticipantResultUserPrivacyRestricted = new(InviteGroupCallParticipantResultUserPrivacyRestricted)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserPrivacyRestricted)
	case "inviteGroupCallParticipantResultUserAlreadyParticipant":
		t.InviteGroupCallParticipantResultUserAlreadyParticipant = new(InviteGroupCallParticipantResultUserAlreadyParticipant)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserAlreadyParticipant)
	case "inviteGroupCallParticipantResultUserWasBanned":
		t.InviteGroupCallParticipantResultUserWasBanned = new(InviteGroupCallParticipantResultUserWasBanned)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultUserWasBanned)
	case "inviteGroupCallParticipantResultSuccess":
		t.InviteGroupCallParticipantResultSuccess = new(InviteGroupCallParticipantResultSuccess)
		return json.Unmarshal(b, t.InviteGroupCallParticipantResultSuccess)
	}
	return nil
}

func (t *InviteGroupCallParticipantResult) MarshalJSON() ([]byte, error) {
	if t.InviteGroupCallParticipantResultUserPrivacyRestricted != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserPrivacyRestricted)
	}
	if t.InviteGroupCallParticipantResultUserAlreadyParticipant != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserAlreadyParticipant)
	}
	if t.InviteGroupCallParticipantResultUserWasBanned != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultUserWasBanned)
	}
	if t.InviteGroupCallParticipantResultSuccess != nil {
		return json.Marshal(t.InviteGroupCallParticipantResultSuccess)
	}
	return nil, fmt.Errorf("no value set for InviteGroupCallParticipantResult")
}

// InputChatTheme Describes a chat theme to set
type InputChatTheme struct {
	InputChatThemeEmoji *InputChatThemeEmoji `json:"inputChatThemeEmoji,omitempty"`
	InputChatThemeGift  *InputChatThemeGift  `json:"inputChatThemeGift,omitempty"`
}

func (t *InputChatTheme) Type() string {
	return ""
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

// FileType Represents the type of file
type FileType struct {
	FileTypeNone                     *FileTypeNone                     `json:"fileTypeNone,omitempty"`
	FileTypeAnimation                *FileTypeAnimation                `json:"fileTypeAnimation,omitempty"`
	FileTypeAudio                    *FileTypeAudio                    `json:"fileTypeAudio,omitempty"`
	FileTypeDocument                 *FileTypeDocument                 `json:"fileTypeDocument,omitempty"`
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
	return ""
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
	switch typeObj.Type {
	case "fileTypeNone":
		t.FileTypeNone = new(FileTypeNone)
		return json.Unmarshal(b, t.FileTypeNone)
	case "fileTypeAnimation":
		t.FileTypeAnimation = new(FileTypeAnimation)
		return json.Unmarshal(b, t.FileTypeAnimation)
	case "fileTypeAudio":
		t.FileTypeAudio = new(FileTypeAudio)
		return json.Unmarshal(b, t.FileTypeAudio)
	case "fileTypeDocument":
		t.FileTypeDocument = new(FileTypeDocument)
		return json.Unmarshal(b, t.FileTypeDocument)
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
	if t.FileTypeNone != nil {
		return json.Marshal(t.FileTypeNone)
	}
	if t.FileTypeAnimation != nil {
		return json.Marshal(t.FileTypeAnimation)
	}
	if t.FileTypeAudio != nil {
		return json.Marshal(t.FileTypeAudio)
	}
	if t.FileTypeDocument != nil {
		return json.Marshal(t.FileTypeDocument)
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

// ChatType Describes the type of chat
type ChatType struct {
	ChatTypePrivate    *ChatTypePrivate    `json:"chatTypePrivate,omitempty"`
	ChatTypeBasicGroup *ChatTypeBasicGroup `json:"chatTypeBasicGroup,omitempty"`
	ChatTypeSupergroup *ChatTypeSupergroup `json:"chatTypeSupergroup,omitempty"`
	ChatTypeSecret     *ChatTypeSecret     `json:"chatTypeSecret,omitempty"`
}

func (t *ChatType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatTypePrivate":
		t.ChatTypePrivate = new(ChatTypePrivate)
		return json.Unmarshal(b, t.ChatTypePrivate)
	case "chatTypeBasicGroup":
		t.ChatTypeBasicGroup = new(ChatTypeBasicGroup)
		return json.Unmarshal(b, t.ChatTypeBasicGroup)
	case "chatTypeSupergroup":
		t.ChatTypeSupergroup = new(ChatTypeSupergroup)
		return json.Unmarshal(b, t.ChatTypeSupergroup)
	case "chatTypeSecret":
		t.ChatTypeSecret = new(ChatTypeSecret)
		return json.Unmarshal(b, t.ChatTypeSecret)
	}
	return nil
}

func (t *ChatType) MarshalJSON() ([]byte, error) {
	if t.ChatTypePrivate != nil {
		return json.Marshal(t.ChatTypePrivate)
	}
	if t.ChatTypeBasicGroup != nil {
		return json.Marshal(t.ChatTypeBasicGroup)
	}
	if t.ChatTypeSupergroup != nil {
		return json.Marshal(t.ChatTypeSupergroup)
	}
	if t.ChatTypeSecret != nil {
		return json.Marshal(t.ChatTypeSecret)
	}
	return nil, fmt.Errorf("no value set for ChatType")
}

// ChatAvailableReactions Describes reactions available in the chat
type ChatAvailableReactions struct {
	ChatAvailableReactionsAll  *ChatAvailableReactionsAll  `json:"chatAvailableReactionsAll,omitempty"`
	ChatAvailableReactionsSome *ChatAvailableReactionsSome `json:"chatAvailableReactionsSome,omitempty"`
}

func (t *ChatAvailableReactions) Type() string {
	return ""
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

// ResendCodeReason Describes the reason why a code needs to be re-sent
type ResendCodeReason struct {
	ResendCodeReasonUserRequest        *ResendCodeReasonUserRequest        `json:"resendCodeReasonUserRequest,omitempty"`
	ResendCodeReasonVerificationFailed *ResendCodeReasonVerificationFailed `json:"resendCodeReasonVerificationFailed,omitempty"`
}

func (t *ResendCodeReason) Type() string {
	return ""
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

// FirebaseAuthenticationSettings Contains settings for Firebase Authentication in the official applications
type FirebaseAuthenticationSettings struct {
	FirebaseAuthenticationSettingsAndroid *FirebaseAuthenticationSettingsAndroid `json:"firebaseAuthenticationSettingsAndroid,omitempty"`
	FirebaseAuthenticationSettingsIos     *FirebaseAuthenticationSettingsIos     `json:"firebaseAuthenticationSettingsIos,omitempty"`
}

func (t *FirebaseAuthenticationSettings) Type() string {
	return ""
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

// InlineQueryResult Represents a single result of an inline query
type InlineQueryResult struct {
	InlineQueryResultArticle   *InlineQueryResultArticle   `json:"inlineQueryResultArticle,omitempty"`
	InlineQueryResultContact   *InlineQueryResultContact   `json:"inlineQueryResultContact,omitempty"`
	InlineQueryResultLocation  *InlineQueryResultLocation  `json:"inlineQueryResultLocation,omitempty"`
	InlineQueryResultVenue     *InlineQueryResultVenue     `json:"inlineQueryResultVenue,omitempty"`
	InlineQueryResultGame      *InlineQueryResultGame      `json:"inlineQueryResultGame,omitempty"`
	InlineQueryResultAnimation *InlineQueryResultAnimation `json:"inlineQueryResultAnimation,omitempty"`
	InlineQueryResultAudio     *InlineQueryResultAudio     `json:"inlineQueryResultAudio,omitempty"`
	InlineQueryResultDocument  *InlineQueryResultDocument  `json:"inlineQueryResultDocument,omitempty"`
	InlineQueryResultPhoto     *InlineQueryResultPhoto     `json:"inlineQueryResultPhoto,omitempty"`
	InlineQueryResultSticker   *InlineQueryResultSticker   `json:"inlineQueryResultSticker,omitempty"`
	InlineQueryResultVideo     *InlineQueryResultVideo     `json:"inlineQueryResultVideo,omitempty"`
	InlineQueryResultVoiceNote *InlineQueryResultVoiceNote `json:"inlineQueryResultVoiceNote,omitempty"`
}

func (t *InlineQueryResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inlineQueryResultArticle":
		t.InlineQueryResultArticle = new(InlineQueryResultArticle)
		return json.Unmarshal(b, t.InlineQueryResultArticle)
	case "inlineQueryResultContact":
		t.InlineQueryResultContact = new(InlineQueryResultContact)
		return json.Unmarshal(b, t.InlineQueryResultContact)
	case "inlineQueryResultLocation":
		t.InlineQueryResultLocation = new(InlineQueryResultLocation)
		return json.Unmarshal(b, t.InlineQueryResultLocation)
	case "inlineQueryResultVenue":
		t.InlineQueryResultVenue = new(InlineQueryResultVenue)
		return json.Unmarshal(b, t.InlineQueryResultVenue)
	case "inlineQueryResultGame":
		t.InlineQueryResultGame = new(InlineQueryResultGame)
		return json.Unmarshal(b, t.InlineQueryResultGame)
	case "inlineQueryResultAnimation":
		t.InlineQueryResultAnimation = new(InlineQueryResultAnimation)
		return json.Unmarshal(b, t.InlineQueryResultAnimation)
	case "inlineQueryResultAudio":
		t.InlineQueryResultAudio = new(InlineQueryResultAudio)
		return json.Unmarshal(b, t.InlineQueryResultAudio)
	case "inlineQueryResultDocument":
		t.InlineQueryResultDocument = new(InlineQueryResultDocument)
		return json.Unmarshal(b, t.InlineQueryResultDocument)
	case "inlineQueryResultPhoto":
		t.InlineQueryResultPhoto = new(InlineQueryResultPhoto)
		return json.Unmarshal(b, t.InlineQueryResultPhoto)
	case "inlineQueryResultSticker":
		t.InlineQueryResultSticker = new(InlineQueryResultSticker)
		return json.Unmarshal(b, t.InlineQueryResultSticker)
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
	if t.InlineQueryResultArticle != nil {
		return json.Marshal(t.InlineQueryResultArticle)
	}
	if t.InlineQueryResultContact != nil {
		return json.Marshal(t.InlineQueryResultContact)
	}
	if t.InlineQueryResultLocation != nil {
		return json.Marshal(t.InlineQueryResultLocation)
	}
	if t.InlineQueryResultVenue != nil {
		return json.Marshal(t.InlineQueryResultVenue)
	}
	if t.InlineQueryResultGame != nil {
		return json.Marshal(t.InlineQueryResultGame)
	}
	if t.InlineQueryResultAnimation != nil {
		return json.Marshal(t.InlineQueryResultAnimation)
	}
	if t.InlineQueryResultAudio != nil {
		return json.Marshal(t.InlineQueryResultAudio)
	}
	if t.InlineQueryResultDocument != nil {
		return json.Marshal(t.InlineQueryResultDocument)
	}
	if t.InlineQueryResultPhoto != nil {
		return json.Marshal(t.InlineQueryResultPhoto)
	}
	if t.InlineQueryResultSticker != nil {
		return json.Marshal(t.InlineQueryResultSticker)
	}
	if t.InlineQueryResultVideo != nil {
		return json.Marshal(t.InlineQueryResultVideo)
	}
	if t.InlineQueryResultVoiceNote != nil {
		return json.Marshal(t.InlineQueryResultVoiceNote)
	}
	return nil, fmt.Errorf("no value set for InlineQueryResult")
}

// CanTransferOwnershipResult Represents result of checking whether the current session can be used to transfer a chat ownership to another user
type CanTransferOwnershipResult struct {
	CanTransferOwnershipResultOk               *CanTransferOwnershipResultOk               `json:"canTransferOwnershipResultOk,omitempty"`
	CanTransferOwnershipResultPasswordNeeded   *CanTransferOwnershipResultPasswordNeeded   `json:"canTransferOwnershipResultPasswordNeeded,omitempty"`
	CanTransferOwnershipResultPasswordTooFresh *CanTransferOwnershipResultPasswordTooFresh `json:"canTransferOwnershipResultPasswordTooFresh,omitempty"`
	CanTransferOwnershipResultSessionTooFresh  *CanTransferOwnershipResultSessionTooFresh  `json:"canTransferOwnershipResultSessionTooFresh,omitempty"`
}

func (t *CanTransferOwnershipResult) Type() string {
	return ""
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

// NetworkType Represents the type of network
type NetworkType struct {
	NetworkTypeNone          *NetworkTypeNone          `json:"networkTypeNone,omitempty"`
	NetworkTypeMobile        *NetworkTypeMobile        `json:"networkTypeMobile,omitempty"`
	NetworkTypeMobileRoaming *NetworkTypeMobileRoaming `json:"networkTypeMobileRoaming,omitempty"`
	NetworkTypeWiFi          *NetworkTypeWiFi          `json:"networkTypeWiFi,omitempty"`
	NetworkTypeOther         *NetworkTypeOther         `json:"networkTypeOther,omitempty"`
}

func (t *NetworkType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "networkTypeNone":
		t.NetworkTypeNone = new(NetworkTypeNone)
		return json.Unmarshal(b, t.NetworkTypeNone)
	case "networkTypeMobile":
		t.NetworkTypeMobile = new(NetworkTypeMobile)
		return json.Unmarshal(b, t.NetworkTypeMobile)
	case "networkTypeMobileRoaming":
		t.NetworkTypeMobileRoaming = new(NetworkTypeMobileRoaming)
		return json.Unmarshal(b, t.NetworkTypeMobileRoaming)
	case "networkTypeWiFi":
		t.NetworkTypeWiFi = new(NetworkTypeWiFi)
		return json.Unmarshal(b, t.NetworkTypeWiFi)
	case "networkTypeOther":
		t.NetworkTypeOther = new(NetworkTypeOther)
		return json.Unmarshal(b, t.NetworkTypeOther)
	}
	return nil
}

func (t *NetworkType) MarshalJSON() ([]byte, error) {
	if t.NetworkTypeNone != nil {
		return json.Marshal(t.NetworkTypeNone)
	}
	if t.NetworkTypeMobile != nil {
		return json.Marshal(t.NetworkTypeMobile)
	}
	if t.NetworkTypeMobileRoaming != nil {
		return json.Marshal(t.NetworkTypeMobileRoaming)
	}
	if t.NetworkTypeWiFi != nil {
		return json.Marshal(t.NetworkTypeWiFi)
	}
	if t.NetworkTypeOther != nil {
		return json.Marshal(t.NetworkTypeOther)
	}
	return nil, fmt.Errorf("no value set for NetworkType")
}

// StoryOrigin Contains information about the origin of a story that was reposted
type StoryOrigin struct {
	StoryOriginPublicStory *StoryOriginPublicStory `json:"storyOriginPublicStory,omitempty"`
	StoryOriginHiddenUser  *StoryOriginHiddenUser  `json:"storyOriginHiddenUser,omitempty"`
}

func (t *StoryOrigin) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storyOriginPublicStory":
		t.StoryOriginPublicStory = new(StoryOriginPublicStory)
		return json.Unmarshal(b, t.StoryOriginPublicStory)
	case "storyOriginHiddenUser":
		t.StoryOriginHiddenUser = new(StoryOriginHiddenUser)
		return json.Unmarshal(b, t.StoryOriginHiddenUser)
	}
	return nil
}

func (t *StoryOrigin) MarshalJSON() ([]byte, error) {
	if t.StoryOriginPublicStory != nil {
		return json.Marshal(t.StoryOriginPublicStory)
	}
	if t.StoryOriginHiddenUser != nil {
		return json.Marshal(t.StoryOriginHiddenUser)
	}
	return nil, fmt.Errorf("no value set for StoryOrigin")
}

// ProfileTab Describes a tab shown in a user or a chat profile
type ProfileTab struct {
	ProfileTabPosts *ProfileTabPosts `json:"profileTabPosts,omitempty"`
	ProfileTabGifts *ProfileTabGifts `json:"profileTabGifts,omitempty"`
	ProfileTabMedia *ProfileTabMedia `json:"profileTabMedia,omitempty"`
	ProfileTabFiles *ProfileTabFiles `json:"profileTabFiles,omitempty"`
	ProfileTabLinks *ProfileTabLinks `json:"profileTabLinks,omitempty"`
	ProfileTabMusic *ProfileTabMusic `json:"profileTabMusic,omitempty"`
	ProfileTabVoice *ProfileTabVoice `json:"profileTabVoice,omitempty"`
	ProfileTabGifs  *ProfileTabGifs  `json:"profileTabGifs,omitempty"`
}

func (t *ProfileTab) Type() string {
	return ""
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
	switch typeObj.Type {
	case "profileTabPosts":
		t.ProfileTabPosts = new(ProfileTabPosts)
		return json.Unmarshal(b, t.ProfileTabPosts)
	case "profileTabGifts":
		t.ProfileTabGifts = new(ProfileTabGifts)
		return json.Unmarshal(b, t.ProfileTabGifts)
	case "profileTabMedia":
		t.ProfileTabMedia = new(ProfileTabMedia)
		return json.Unmarshal(b, t.ProfileTabMedia)
	case "profileTabFiles":
		t.ProfileTabFiles = new(ProfileTabFiles)
		return json.Unmarshal(b, t.ProfileTabFiles)
	case "profileTabLinks":
		t.ProfileTabLinks = new(ProfileTabLinks)
		return json.Unmarshal(b, t.ProfileTabLinks)
	case "profileTabMusic":
		t.ProfileTabMusic = new(ProfileTabMusic)
		return json.Unmarshal(b, t.ProfileTabMusic)
	case "profileTabVoice":
		t.ProfileTabVoice = new(ProfileTabVoice)
		return json.Unmarshal(b, t.ProfileTabVoice)
	case "profileTabGifs":
		t.ProfileTabGifs = new(ProfileTabGifs)
		return json.Unmarshal(b, t.ProfileTabGifs)
	}
	return nil
}

func (t *ProfileTab) MarshalJSON() ([]byte, error) {
	if t.ProfileTabPosts != nil {
		return json.Marshal(t.ProfileTabPosts)
	}
	if t.ProfileTabGifts != nil {
		return json.Marshal(t.ProfileTabGifts)
	}
	if t.ProfileTabMedia != nil {
		return json.Marshal(t.ProfileTabMedia)
	}
	if t.ProfileTabFiles != nil {
		return json.Marshal(t.ProfileTabFiles)
	}
	if t.ProfileTabLinks != nil {
		return json.Marshal(t.ProfileTabLinks)
	}
	if t.ProfileTabMusic != nil {
		return json.Marshal(t.ProfileTabMusic)
	}
	if t.ProfileTabVoice != nil {
		return json.Marshal(t.ProfileTabVoice)
	}
	if t.ProfileTabGifs != nil {
		return json.Marshal(t.ProfileTabGifs)
	}
	return nil, fmt.Errorf("no value set for ProfileTab")
}

// ChatActionBar Describes actions which must be possible to do through a chat action bar
type ChatActionBar struct {
	ChatActionBarReportSpam       *ChatActionBarReportSpam       `json:"chatActionBarReportSpam,omitempty"`
	ChatActionBarInviteMembers    *ChatActionBarInviteMembers    `json:"chatActionBarInviteMembers,omitempty"`
	ChatActionBarReportAddBlock   *ChatActionBarReportAddBlock   `json:"chatActionBarReportAddBlock,omitempty"`
	ChatActionBarAddContact       *ChatActionBarAddContact       `json:"chatActionBarAddContact,omitempty"`
	ChatActionBarSharePhoneNumber *ChatActionBarSharePhoneNumber `json:"chatActionBarSharePhoneNumber,omitempty"`
	ChatActionBarJoinRequest      *ChatActionBarJoinRequest      `json:"chatActionBarJoinRequest,omitempty"`
}

func (t *ChatActionBar) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatActionBarReportSpam":
		t.ChatActionBarReportSpam = new(ChatActionBarReportSpam)
		return json.Unmarshal(b, t.ChatActionBarReportSpam)
	case "chatActionBarInviteMembers":
		t.ChatActionBarInviteMembers = new(ChatActionBarInviteMembers)
		return json.Unmarshal(b, t.ChatActionBarInviteMembers)
	case "chatActionBarReportAddBlock":
		t.ChatActionBarReportAddBlock = new(ChatActionBarReportAddBlock)
		return json.Unmarshal(b, t.ChatActionBarReportAddBlock)
	case "chatActionBarAddContact":
		t.ChatActionBarAddContact = new(ChatActionBarAddContact)
		return json.Unmarshal(b, t.ChatActionBarAddContact)
	case "chatActionBarSharePhoneNumber":
		t.ChatActionBarSharePhoneNumber = new(ChatActionBarSharePhoneNumber)
		return json.Unmarshal(b, t.ChatActionBarSharePhoneNumber)
	case "chatActionBarJoinRequest":
		t.ChatActionBarJoinRequest = new(ChatActionBarJoinRequest)
		return json.Unmarshal(b, t.ChatActionBarJoinRequest)
	}
	return nil
}

func (t *ChatActionBar) MarshalJSON() ([]byte, error) {
	if t.ChatActionBarReportSpam != nil {
		return json.Marshal(t.ChatActionBarReportSpam)
	}
	if t.ChatActionBarInviteMembers != nil {
		return json.Marshal(t.ChatActionBarInviteMembers)
	}
	if t.ChatActionBarReportAddBlock != nil {
		return json.Marshal(t.ChatActionBarReportAddBlock)
	}
	if t.ChatActionBarAddContact != nil {
		return json.Marshal(t.ChatActionBarAddContact)
	}
	if t.ChatActionBarSharePhoneNumber != nil {
		return json.Marshal(t.ChatActionBarSharePhoneNumber)
	}
	if t.ChatActionBarJoinRequest != nil {
		return json.Marshal(t.ChatActionBarJoinRequest)
	}
	return nil, fmt.Errorf("no value set for ChatActionBar")
}

// CollectibleItemType Describes a collectible item that can be purchased at https://fragment.com
type CollectibleItemType struct {
	CollectibleItemTypeUsername    *CollectibleItemTypeUsername    `json:"collectibleItemTypeUsername,omitempty"`
	CollectibleItemTypePhoneNumber *CollectibleItemTypePhoneNumber `json:"collectibleItemTypePhoneNumber,omitempty"`
}

func (t *CollectibleItemType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "collectibleItemTypeUsername":
		t.CollectibleItemTypeUsername = new(CollectibleItemTypeUsername)
		return json.Unmarshal(b, t.CollectibleItemTypeUsername)
	case "collectibleItemTypePhoneNumber":
		t.CollectibleItemTypePhoneNumber = new(CollectibleItemTypePhoneNumber)
		return json.Unmarshal(b, t.CollectibleItemTypePhoneNumber)
	}
	return nil
}

func (t *CollectibleItemType) MarshalJSON() ([]byte, error) {
	if t.CollectibleItemTypeUsername != nil {
		return json.Marshal(t.CollectibleItemTypeUsername)
	}
	if t.CollectibleItemTypePhoneNumber != nil {
		return json.Marshal(t.CollectibleItemTypePhoneNumber)
	}
	return nil, fmt.Errorf("no value set for CollectibleItemType")
}

// MessageSelfDestructType Describes when a message will be self-destructed
type MessageSelfDestructType struct {
	MessageSelfDestructTypeTimer       *MessageSelfDestructTypeTimer       `json:"messageSelfDestructTypeTimer,omitempty"`
	MessageSelfDestructTypeImmediately *MessageSelfDestructTypeImmediately `json:"messageSelfDestructTypeImmediately,omitempty"`
}

func (t *MessageSelfDestructType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageSelfDestructTypeTimer":
		t.MessageSelfDestructTypeTimer = new(MessageSelfDestructTypeTimer)
		return json.Unmarshal(b, t.MessageSelfDestructTypeTimer)
	case "messageSelfDestructTypeImmediately":
		t.MessageSelfDestructTypeImmediately = new(MessageSelfDestructTypeImmediately)
		return json.Unmarshal(b, t.MessageSelfDestructTypeImmediately)
	}
	return nil
}

func (t *MessageSelfDestructType) MarshalJSON() ([]byte, error) {
	if t.MessageSelfDestructTypeTimer != nil {
		return json.Marshal(t.MessageSelfDestructTypeTimer)
	}
	if t.MessageSelfDestructTypeImmediately != nil {
		return json.Marshal(t.MessageSelfDestructTypeImmediately)
	}
	return nil, fmt.Errorf("no value set for MessageSelfDestructType")
}

// CallServerType Describes the type of call server
type CallServerType struct {
	CallServerTypeTelegramReflector *CallServerTypeTelegramReflector `json:"callServerTypeTelegramReflector,omitempty"`
	CallServerTypeWebrtc            *CallServerTypeWebrtc            `json:"callServerTypeWebrtc,omitempty"`
}

func (t *CallServerType) Type() string {
	return ""
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

// StorePaymentPurpose Describes a purpose of an in-store payment
type StorePaymentPurpose struct {
	StorePaymentPurposePremiumSubscription *StorePaymentPurposePremiumSubscription `json:"storePaymentPurposePremiumSubscription,omitempty"`
	StorePaymentPurposePremiumGift         *StorePaymentPurposePremiumGift         `json:"storePaymentPurposePremiumGift,omitempty"`
	StorePaymentPurposePremiumGiftCodes    *StorePaymentPurposePremiumGiftCodes    `json:"storePaymentPurposePremiumGiftCodes,omitempty"`
	StorePaymentPurposePremiumGiveaway     *StorePaymentPurposePremiumGiveaway     `json:"storePaymentPurposePremiumGiveaway,omitempty"`
	StorePaymentPurposeStarGiveaway        *StorePaymentPurposeStarGiveaway        `json:"storePaymentPurposeStarGiveaway,omitempty"`
	StorePaymentPurposeStars               *StorePaymentPurposeStars               `json:"storePaymentPurposeStars,omitempty"`
	StorePaymentPurposeGiftedStars         *StorePaymentPurposeGiftedStars         `json:"storePaymentPurposeGiftedStars,omitempty"`
}

func (t *StorePaymentPurpose) Type() string {
	return ""
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
	switch typeObj.Type {
	case "storePaymentPurposePremiumSubscription":
		t.StorePaymentPurposePremiumSubscription = new(StorePaymentPurposePremiumSubscription)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumSubscription)
	case "storePaymentPurposePremiumGift":
		t.StorePaymentPurposePremiumGift = new(StorePaymentPurposePremiumGift)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGift)
	case "storePaymentPurposePremiumGiftCodes":
		t.StorePaymentPurposePremiumGiftCodes = new(StorePaymentPurposePremiumGiftCodes)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGiftCodes)
	case "storePaymentPurposePremiumGiveaway":
		t.StorePaymentPurposePremiumGiveaway = new(StorePaymentPurposePremiumGiveaway)
		return json.Unmarshal(b, t.StorePaymentPurposePremiumGiveaway)
	case "storePaymentPurposeStarGiveaway":
		t.StorePaymentPurposeStarGiveaway = new(StorePaymentPurposeStarGiveaway)
		return json.Unmarshal(b, t.StorePaymentPurposeStarGiveaway)
	case "storePaymentPurposeStars":
		t.StorePaymentPurposeStars = new(StorePaymentPurposeStars)
		return json.Unmarshal(b, t.StorePaymentPurposeStars)
	case "storePaymentPurposeGiftedStars":
		t.StorePaymentPurposeGiftedStars = new(StorePaymentPurposeGiftedStars)
		return json.Unmarshal(b, t.StorePaymentPurposeGiftedStars)
	}
	return nil
}

func (t *StorePaymentPurpose) MarshalJSON() ([]byte, error) {
	if t.StorePaymentPurposePremiumSubscription != nil {
		return json.Marshal(t.StorePaymentPurposePremiumSubscription)
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
	if t.StorePaymentPurposeStarGiveaway != nil {
		return json.Marshal(t.StorePaymentPurposeStarGiveaway)
	}
	if t.StorePaymentPurposeStars != nil {
		return json.Marshal(t.StorePaymentPurposeStars)
	}
	if t.StorePaymentPurposeGiftedStars != nil {
		return json.Marshal(t.StorePaymentPurposeGiftedStars)
	}
	return nil, fmt.Errorf("no value set for StorePaymentPurpose")
}

// BackgroundType Describes the type of background
type BackgroundType struct {
	BackgroundTypeWallpaper *BackgroundTypeWallpaper `json:"backgroundTypeWallpaper,omitempty"`
	BackgroundTypePattern   *BackgroundTypePattern   `json:"backgroundTypePattern,omitempty"`
	BackgroundTypeFill      *BackgroundTypeFill      `json:"backgroundTypeFill,omitempty"`
	BackgroundTypeChatTheme *BackgroundTypeChatTheme `json:"backgroundTypeChatTheme,omitempty"`
}

func (t *BackgroundType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "backgroundTypeWallpaper":
		t.BackgroundTypeWallpaper = new(BackgroundTypeWallpaper)
		return json.Unmarshal(b, t.BackgroundTypeWallpaper)
	case "backgroundTypePattern":
		t.BackgroundTypePattern = new(BackgroundTypePattern)
		return json.Unmarshal(b, t.BackgroundTypePattern)
	case "backgroundTypeFill":
		t.BackgroundTypeFill = new(BackgroundTypeFill)
		return json.Unmarshal(b, t.BackgroundTypeFill)
	case "backgroundTypeChatTheme":
		t.BackgroundTypeChatTheme = new(BackgroundTypeChatTheme)
		return json.Unmarshal(b, t.BackgroundTypeChatTheme)
	}
	return nil
}

func (t *BackgroundType) MarshalJSON() ([]byte, error) {
	if t.BackgroundTypeWallpaper != nil {
		return json.Marshal(t.BackgroundTypeWallpaper)
	}
	if t.BackgroundTypePattern != nil {
		return json.Marshal(t.BackgroundTypePattern)
	}
	if t.BackgroundTypeFill != nil {
		return json.Marshal(t.BackgroundTypeFill)
	}
	if t.BackgroundTypeChatTheme != nil {
		return json.Marshal(t.BackgroundTypeChatTheme)
	}
	return nil, fmt.Errorf("no value set for BackgroundType")
}

// AuthenticationCodeType Provides information about the method by which an authentication code is delivered to the user
type AuthenticationCodeType struct {
	AuthenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage `json:"authenticationCodeTypeTelegramMessage,omitempty"`
	AuthenticationCodeTypeSms             *AuthenticationCodeTypeSms             `json:"authenticationCodeTypeSms,omitempty"`
	AuthenticationCodeTypeSmsWord         *AuthenticationCodeTypeSmsWord         `json:"authenticationCodeTypeSmsWord,omitempty"`
	AuthenticationCodeTypeSmsPhrase       *AuthenticationCodeTypeSmsPhrase       `json:"authenticationCodeTypeSmsPhrase,omitempty"`
	AuthenticationCodeTypeCall            *AuthenticationCodeTypeCall            `json:"authenticationCodeTypeCall,omitempty"`
	AuthenticationCodeTypeFlashCall       *AuthenticationCodeTypeFlashCall       `json:"authenticationCodeTypeFlashCall,omitempty"`
	AuthenticationCodeTypeMissedCall      *AuthenticationCodeTypeMissedCall      `json:"authenticationCodeTypeMissedCall,omitempty"`
	AuthenticationCodeTypeFragment        *AuthenticationCodeTypeFragment        `json:"authenticationCodeTypeFragment,omitempty"`
	AuthenticationCodeTypeFirebaseAndroid *AuthenticationCodeTypeFirebaseAndroid `json:"authenticationCodeTypeFirebaseAndroid,omitempty"`
	AuthenticationCodeTypeFirebaseIos     *AuthenticationCodeTypeFirebaseIos     `json:"authenticationCodeTypeFirebaseIos,omitempty"`
}

func (t *AuthenticationCodeType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "authenticationCodeTypeTelegramMessage":
		t.AuthenticationCodeTypeTelegramMessage = new(AuthenticationCodeTypeTelegramMessage)
		return json.Unmarshal(b, t.AuthenticationCodeTypeTelegramMessage)
	case "authenticationCodeTypeSms":
		t.AuthenticationCodeTypeSms = new(AuthenticationCodeTypeSms)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSms)
	case "authenticationCodeTypeSmsWord":
		t.AuthenticationCodeTypeSmsWord = new(AuthenticationCodeTypeSmsWord)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSmsWord)
	case "authenticationCodeTypeSmsPhrase":
		t.AuthenticationCodeTypeSmsPhrase = new(AuthenticationCodeTypeSmsPhrase)
		return json.Unmarshal(b, t.AuthenticationCodeTypeSmsPhrase)
	case "authenticationCodeTypeCall":
		t.AuthenticationCodeTypeCall = new(AuthenticationCodeTypeCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeCall)
	case "authenticationCodeTypeFlashCall":
		t.AuthenticationCodeTypeFlashCall = new(AuthenticationCodeTypeFlashCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFlashCall)
	case "authenticationCodeTypeMissedCall":
		t.AuthenticationCodeTypeMissedCall = new(AuthenticationCodeTypeMissedCall)
		return json.Unmarshal(b, t.AuthenticationCodeTypeMissedCall)
	case "authenticationCodeTypeFragment":
		t.AuthenticationCodeTypeFragment = new(AuthenticationCodeTypeFragment)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFragment)
	case "authenticationCodeTypeFirebaseAndroid":
		t.AuthenticationCodeTypeFirebaseAndroid = new(AuthenticationCodeTypeFirebaseAndroid)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFirebaseAndroid)
	case "authenticationCodeTypeFirebaseIos":
		t.AuthenticationCodeTypeFirebaseIos = new(AuthenticationCodeTypeFirebaseIos)
		return json.Unmarshal(b, t.AuthenticationCodeTypeFirebaseIos)
	}
	return nil
}

func (t *AuthenticationCodeType) MarshalJSON() ([]byte, error) {
	if t.AuthenticationCodeTypeTelegramMessage != nil {
		return json.Marshal(t.AuthenticationCodeTypeTelegramMessage)
	}
	if t.AuthenticationCodeTypeSms != nil {
		return json.Marshal(t.AuthenticationCodeTypeSms)
	}
	if t.AuthenticationCodeTypeSmsWord != nil {
		return json.Marshal(t.AuthenticationCodeTypeSmsWord)
	}
	if t.AuthenticationCodeTypeSmsPhrase != nil {
		return json.Marshal(t.AuthenticationCodeTypeSmsPhrase)
	}
	if t.AuthenticationCodeTypeCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeCall)
	}
	if t.AuthenticationCodeTypeFlashCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeFlashCall)
	}
	if t.AuthenticationCodeTypeMissedCall != nil {
		return json.Marshal(t.AuthenticationCodeTypeMissedCall)
	}
	if t.AuthenticationCodeTypeFragment != nil {
		return json.Marshal(t.AuthenticationCodeTypeFragment)
	}
	if t.AuthenticationCodeTypeFirebaseAndroid != nil {
		return json.Marshal(t.AuthenticationCodeTypeFirebaseAndroid)
	}
	if t.AuthenticationCodeTypeFirebaseIos != nil {
		return json.Marshal(t.AuthenticationCodeTypeFirebaseIos)
	}
	return nil, fmt.Errorf("no value set for AuthenticationCodeType")
}

// StarTransactionType Describes type of transaction with Telegram Stars
type StarTransactionType struct {
	StarTransactionTypePremiumBotDeposit            *StarTransactionTypePremiumBotDeposit            `json:"starTransactionTypePremiumBotDeposit,omitempty"`
	StarTransactionTypeAppStoreDeposit              *StarTransactionTypeAppStoreDeposit              `json:"starTransactionTypeAppStoreDeposit,omitempty"`
	StarTransactionTypeGooglePlayDeposit            *StarTransactionTypeGooglePlayDeposit            `json:"starTransactionTypeGooglePlayDeposit,omitempty"`
	StarTransactionTypeFragmentDeposit              *StarTransactionTypeFragmentDeposit              `json:"starTransactionTypeFragmentDeposit,omitempty"`
	StarTransactionTypeUserDeposit                  *StarTransactionTypeUserDeposit                  `json:"starTransactionTypeUserDeposit,omitempty"`
	StarTransactionTypeGiveawayDeposit              *StarTransactionTypeGiveawayDeposit              `json:"starTransactionTypeGiveawayDeposit,omitempty"`
	StarTransactionTypeFragmentWithdrawal           *StarTransactionTypeFragmentWithdrawal           `json:"starTransactionTypeFragmentWithdrawal,omitempty"`
	StarTransactionTypeTelegramAdsWithdrawal        *StarTransactionTypeTelegramAdsWithdrawal        `json:"starTransactionTypeTelegramAdsWithdrawal,omitempty"`
	StarTransactionTypeTelegramApiUsage             *StarTransactionTypeTelegramApiUsage             `json:"starTransactionTypeTelegramApiUsage,omitempty"`
	StarTransactionTypeBotPaidMediaPurchase         *StarTransactionTypeBotPaidMediaPurchase         `json:"starTransactionTypeBotPaidMediaPurchase,omitempty"`
	StarTransactionTypeBotPaidMediaSale             *StarTransactionTypeBotPaidMediaSale             `json:"starTransactionTypeBotPaidMediaSale,omitempty"`
	StarTransactionTypeChannelPaidMediaPurchase     *StarTransactionTypeChannelPaidMediaPurchase     `json:"starTransactionTypeChannelPaidMediaPurchase,omitempty"`
	StarTransactionTypeChannelPaidMediaSale         *StarTransactionTypeChannelPaidMediaSale         `json:"starTransactionTypeChannelPaidMediaSale,omitempty"`
	StarTransactionTypeBotInvoicePurchase           *StarTransactionTypeBotInvoicePurchase           `json:"starTransactionTypeBotInvoicePurchase,omitempty"`
	StarTransactionTypeBotInvoiceSale               *StarTransactionTypeBotInvoiceSale               `json:"starTransactionTypeBotInvoiceSale,omitempty"`
	StarTransactionTypeBotSubscriptionPurchase      *StarTransactionTypeBotSubscriptionPurchase      `json:"starTransactionTypeBotSubscriptionPurchase,omitempty"`
	StarTransactionTypeBotSubscriptionSale          *StarTransactionTypeBotSubscriptionSale          `json:"starTransactionTypeBotSubscriptionSale,omitempty"`
	StarTransactionTypeChannelSubscriptionPurchase  *StarTransactionTypeChannelSubscriptionPurchase  `json:"starTransactionTypeChannelSubscriptionPurchase,omitempty"`
	StarTransactionTypeChannelSubscriptionSale      *StarTransactionTypeChannelSubscriptionSale      `json:"starTransactionTypeChannelSubscriptionSale,omitempty"`
	StarTransactionTypeGiftAuctionBid               *StarTransactionTypeGiftAuctionBid               `json:"starTransactionTypeGiftAuctionBid,omitempty"`
	StarTransactionTypeGiftPurchase                 *StarTransactionTypeGiftPurchase                 `json:"starTransactionTypeGiftPurchase,omitempty"`
	StarTransactionTypeGiftPurchaseOffer            *StarTransactionTypeGiftPurchaseOffer            `json:"starTransactionTypeGiftPurchaseOffer,omitempty"`
	StarTransactionTypeGiftTransfer                 *StarTransactionTypeGiftTransfer                 `json:"starTransactionTypeGiftTransfer,omitempty"`
	StarTransactionTypeGiftOriginalDetailsDrop      *StarTransactionTypeGiftOriginalDetailsDrop      `json:"starTransactionTypeGiftOriginalDetailsDrop,omitempty"`
	StarTransactionTypeGiftSale                     *StarTransactionTypeGiftSale                     `json:"starTransactionTypeGiftSale,omitempty"`
	StarTransactionTypeGiftUpgrade                  *StarTransactionTypeGiftUpgrade                  `json:"starTransactionTypeGiftUpgrade,omitempty"`
	StarTransactionTypeGiftUpgradePurchase          *StarTransactionTypeGiftUpgradePurchase          `json:"starTransactionTypeGiftUpgradePurchase,omitempty"`
	StarTransactionTypeUpgradedGiftPurchase         *StarTransactionTypeUpgradedGiftPurchase         `json:"starTransactionTypeUpgradedGiftPurchase,omitempty"`
	StarTransactionTypeUpgradedGiftSale             *StarTransactionTypeUpgradedGiftSale             `json:"starTransactionTypeUpgradedGiftSale,omitempty"`
	StarTransactionTypeChannelPaidReactionSend      *StarTransactionTypeChannelPaidReactionSend      `json:"starTransactionTypeChannelPaidReactionSend,omitempty"`
	StarTransactionTypeChannelPaidReactionReceive   *StarTransactionTypeChannelPaidReactionReceive   `json:"starTransactionTypeChannelPaidReactionReceive,omitempty"`
	StarTransactionTypeAffiliateProgramCommission   *StarTransactionTypeAffiliateProgramCommission   `json:"starTransactionTypeAffiliateProgramCommission,omitempty"`
	StarTransactionTypePaidMessageSend              *StarTransactionTypePaidMessageSend              `json:"starTransactionTypePaidMessageSend,omitempty"`
	StarTransactionTypePaidMessageReceive           *StarTransactionTypePaidMessageReceive           `json:"starTransactionTypePaidMessageReceive,omitempty"`
	StarTransactionTypePaidGroupCallMessageSend     *StarTransactionTypePaidGroupCallMessageSend     `json:"starTransactionTypePaidGroupCallMessageSend,omitempty"`
	StarTransactionTypePaidGroupCallMessageReceive  *StarTransactionTypePaidGroupCallMessageReceive  `json:"starTransactionTypePaidGroupCallMessageReceive,omitempty"`
	StarTransactionTypePaidGroupCallReactionSend    *StarTransactionTypePaidGroupCallReactionSend    `json:"starTransactionTypePaidGroupCallReactionSend,omitempty"`
	StarTransactionTypePaidGroupCallReactionReceive *StarTransactionTypePaidGroupCallReactionReceive `json:"starTransactionTypePaidGroupCallReactionReceive,omitempty"`
	StarTransactionTypeSuggestedPostPaymentSend     *StarTransactionTypeSuggestedPostPaymentSend     `json:"starTransactionTypeSuggestedPostPaymentSend,omitempty"`
	StarTransactionTypeSuggestedPostPaymentReceive  *StarTransactionTypeSuggestedPostPaymentReceive  `json:"starTransactionTypeSuggestedPostPaymentReceive,omitempty"`
	StarTransactionTypePremiumPurchase              *StarTransactionTypePremiumPurchase              `json:"starTransactionTypePremiumPurchase,omitempty"`
	StarTransactionTypeBusinessBotTransferSend      *StarTransactionTypeBusinessBotTransferSend      `json:"starTransactionTypeBusinessBotTransferSend,omitempty"`
	StarTransactionTypeBusinessBotTransferReceive   *StarTransactionTypeBusinessBotTransferReceive   `json:"starTransactionTypeBusinessBotTransferReceive,omitempty"`
	StarTransactionTypePublicPostSearch             *StarTransactionTypePublicPostSearch             `json:"starTransactionTypePublicPostSearch,omitempty"`
	StarTransactionTypeUnsupported                  *StarTransactionTypeUnsupported                  `json:"starTransactionTypeUnsupported,omitempty"`
}

func (t *StarTransactionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "starTransactionTypePremiumBotDeposit":
		t.StarTransactionTypePremiumBotDeposit = new(StarTransactionTypePremiumBotDeposit)
		return json.Unmarshal(b, t.StarTransactionTypePremiumBotDeposit)
	case "starTransactionTypeAppStoreDeposit":
		t.StarTransactionTypeAppStoreDeposit = new(StarTransactionTypeAppStoreDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeAppStoreDeposit)
	case "starTransactionTypeGooglePlayDeposit":
		t.StarTransactionTypeGooglePlayDeposit = new(StarTransactionTypeGooglePlayDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeGooglePlayDeposit)
	case "starTransactionTypeFragmentDeposit":
		t.StarTransactionTypeFragmentDeposit = new(StarTransactionTypeFragmentDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeFragmentDeposit)
	case "starTransactionTypeUserDeposit":
		t.StarTransactionTypeUserDeposit = new(StarTransactionTypeUserDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeUserDeposit)
	case "starTransactionTypeGiveawayDeposit":
		t.StarTransactionTypeGiveawayDeposit = new(StarTransactionTypeGiveawayDeposit)
		return json.Unmarshal(b, t.StarTransactionTypeGiveawayDeposit)
	case "starTransactionTypeFragmentWithdrawal":
		t.StarTransactionTypeFragmentWithdrawal = new(StarTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.StarTransactionTypeFragmentWithdrawal)
	case "starTransactionTypeTelegramAdsWithdrawal":
		t.StarTransactionTypeTelegramAdsWithdrawal = new(StarTransactionTypeTelegramAdsWithdrawal)
		return json.Unmarshal(b, t.StarTransactionTypeTelegramAdsWithdrawal)
	case "starTransactionTypeTelegramApiUsage":
		t.StarTransactionTypeTelegramApiUsage = new(StarTransactionTypeTelegramApiUsage)
		return json.Unmarshal(b, t.StarTransactionTypeTelegramApiUsage)
	case "starTransactionTypeBotPaidMediaPurchase":
		t.StarTransactionTypeBotPaidMediaPurchase = new(StarTransactionTypeBotPaidMediaPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotPaidMediaPurchase)
	case "starTransactionTypeBotPaidMediaSale":
		t.StarTransactionTypeBotPaidMediaSale = new(StarTransactionTypeBotPaidMediaSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotPaidMediaSale)
	case "starTransactionTypeChannelPaidMediaPurchase":
		t.StarTransactionTypeChannelPaidMediaPurchase = new(StarTransactionTypeChannelPaidMediaPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidMediaPurchase)
	case "starTransactionTypeChannelPaidMediaSale":
		t.StarTransactionTypeChannelPaidMediaSale = new(StarTransactionTypeChannelPaidMediaSale)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidMediaSale)
	case "starTransactionTypeBotInvoicePurchase":
		t.StarTransactionTypeBotInvoicePurchase = new(StarTransactionTypeBotInvoicePurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotInvoicePurchase)
	case "starTransactionTypeBotInvoiceSale":
		t.StarTransactionTypeBotInvoiceSale = new(StarTransactionTypeBotInvoiceSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotInvoiceSale)
	case "starTransactionTypeBotSubscriptionPurchase":
		t.StarTransactionTypeBotSubscriptionPurchase = new(StarTransactionTypeBotSubscriptionPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeBotSubscriptionPurchase)
	case "starTransactionTypeBotSubscriptionSale":
		t.StarTransactionTypeBotSubscriptionSale = new(StarTransactionTypeBotSubscriptionSale)
		return json.Unmarshal(b, t.StarTransactionTypeBotSubscriptionSale)
	case "starTransactionTypeChannelSubscriptionPurchase":
		t.StarTransactionTypeChannelSubscriptionPurchase = new(StarTransactionTypeChannelSubscriptionPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeChannelSubscriptionPurchase)
	case "starTransactionTypeChannelSubscriptionSale":
		t.StarTransactionTypeChannelSubscriptionSale = new(StarTransactionTypeChannelSubscriptionSale)
		return json.Unmarshal(b, t.StarTransactionTypeChannelSubscriptionSale)
	case "starTransactionTypeGiftAuctionBid":
		t.StarTransactionTypeGiftAuctionBid = new(StarTransactionTypeGiftAuctionBid)
		return json.Unmarshal(b, t.StarTransactionTypeGiftAuctionBid)
	case "starTransactionTypeGiftPurchase":
		t.StarTransactionTypeGiftPurchase = new(StarTransactionTypeGiftPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeGiftPurchase)
	case "starTransactionTypeGiftPurchaseOffer":
		t.StarTransactionTypeGiftPurchaseOffer = new(StarTransactionTypeGiftPurchaseOffer)
		return json.Unmarshal(b, t.StarTransactionTypeGiftPurchaseOffer)
	case "starTransactionTypeGiftTransfer":
		t.StarTransactionTypeGiftTransfer = new(StarTransactionTypeGiftTransfer)
		return json.Unmarshal(b, t.StarTransactionTypeGiftTransfer)
	case "starTransactionTypeGiftOriginalDetailsDrop":
		t.StarTransactionTypeGiftOriginalDetailsDrop = new(StarTransactionTypeGiftOriginalDetailsDrop)
		return json.Unmarshal(b, t.StarTransactionTypeGiftOriginalDetailsDrop)
	case "starTransactionTypeGiftSale":
		t.StarTransactionTypeGiftSale = new(StarTransactionTypeGiftSale)
		return json.Unmarshal(b, t.StarTransactionTypeGiftSale)
	case "starTransactionTypeGiftUpgrade":
		t.StarTransactionTypeGiftUpgrade = new(StarTransactionTypeGiftUpgrade)
		return json.Unmarshal(b, t.StarTransactionTypeGiftUpgrade)
	case "starTransactionTypeGiftUpgradePurchase":
		t.StarTransactionTypeGiftUpgradePurchase = new(StarTransactionTypeGiftUpgradePurchase)
		return json.Unmarshal(b, t.StarTransactionTypeGiftUpgradePurchase)
	case "starTransactionTypeUpgradedGiftPurchase":
		t.StarTransactionTypeUpgradedGiftPurchase = new(StarTransactionTypeUpgradedGiftPurchase)
		return json.Unmarshal(b, t.StarTransactionTypeUpgradedGiftPurchase)
	case "starTransactionTypeUpgradedGiftSale":
		t.StarTransactionTypeUpgradedGiftSale = new(StarTransactionTypeUpgradedGiftSale)
		return json.Unmarshal(b, t.StarTransactionTypeUpgradedGiftSale)
	case "starTransactionTypeChannelPaidReactionSend":
		t.StarTransactionTypeChannelPaidReactionSend = new(StarTransactionTypeChannelPaidReactionSend)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidReactionSend)
	case "starTransactionTypeChannelPaidReactionReceive":
		t.StarTransactionTypeChannelPaidReactionReceive = new(StarTransactionTypeChannelPaidReactionReceive)
		return json.Unmarshal(b, t.StarTransactionTypeChannelPaidReactionReceive)
	case "starTransactionTypeAffiliateProgramCommission":
		t.StarTransactionTypeAffiliateProgramCommission = new(StarTransactionTypeAffiliateProgramCommission)
		return json.Unmarshal(b, t.StarTransactionTypeAffiliateProgramCommission)
	case "starTransactionTypePaidMessageSend":
		t.StarTransactionTypePaidMessageSend = new(StarTransactionTypePaidMessageSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidMessageSend)
	case "starTransactionTypePaidMessageReceive":
		t.StarTransactionTypePaidMessageReceive = new(StarTransactionTypePaidMessageReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidMessageReceive)
	case "starTransactionTypePaidGroupCallMessageSend":
		t.StarTransactionTypePaidGroupCallMessageSend = new(StarTransactionTypePaidGroupCallMessageSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallMessageSend)
	case "starTransactionTypePaidGroupCallMessageReceive":
		t.StarTransactionTypePaidGroupCallMessageReceive = new(StarTransactionTypePaidGroupCallMessageReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallMessageReceive)
	case "starTransactionTypePaidGroupCallReactionSend":
		t.StarTransactionTypePaidGroupCallReactionSend = new(StarTransactionTypePaidGroupCallReactionSend)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallReactionSend)
	case "starTransactionTypePaidGroupCallReactionReceive":
		t.StarTransactionTypePaidGroupCallReactionReceive = new(StarTransactionTypePaidGroupCallReactionReceive)
		return json.Unmarshal(b, t.StarTransactionTypePaidGroupCallReactionReceive)
	case "starTransactionTypeSuggestedPostPaymentSend":
		t.StarTransactionTypeSuggestedPostPaymentSend = new(StarTransactionTypeSuggestedPostPaymentSend)
		return json.Unmarshal(b, t.StarTransactionTypeSuggestedPostPaymentSend)
	case "starTransactionTypeSuggestedPostPaymentReceive":
		t.StarTransactionTypeSuggestedPostPaymentReceive = new(StarTransactionTypeSuggestedPostPaymentReceive)
		return json.Unmarshal(b, t.StarTransactionTypeSuggestedPostPaymentReceive)
	case "starTransactionTypePremiumPurchase":
		t.StarTransactionTypePremiumPurchase = new(StarTransactionTypePremiumPurchase)
		return json.Unmarshal(b, t.StarTransactionTypePremiumPurchase)
	case "starTransactionTypeBusinessBotTransferSend":
		t.StarTransactionTypeBusinessBotTransferSend = new(StarTransactionTypeBusinessBotTransferSend)
		return json.Unmarshal(b, t.StarTransactionTypeBusinessBotTransferSend)
	case "starTransactionTypeBusinessBotTransferReceive":
		t.StarTransactionTypeBusinessBotTransferReceive = new(StarTransactionTypeBusinessBotTransferReceive)
		return json.Unmarshal(b, t.StarTransactionTypeBusinessBotTransferReceive)
	case "starTransactionTypePublicPostSearch":
		t.StarTransactionTypePublicPostSearch = new(StarTransactionTypePublicPostSearch)
		return json.Unmarshal(b, t.StarTransactionTypePublicPostSearch)
	case "starTransactionTypeUnsupported":
		t.StarTransactionTypeUnsupported = new(StarTransactionTypeUnsupported)
		return json.Unmarshal(b, t.StarTransactionTypeUnsupported)
	}
	return nil
}

func (t *StarTransactionType) MarshalJSON() ([]byte, error) {
	if t.StarTransactionTypePremiumBotDeposit != nil {
		return json.Marshal(t.StarTransactionTypePremiumBotDeposit)
	}
	if t.StarTransactionTypeAppStoreDeposit != nil {
		return json.Marshal(t.StarTransactionTypeAppStoreDeposit)
	}
	if t.StarTransactionTypeGooglePlayDeposit != nil {
		return json.Marshal(t.StarTransactionTypeGooglePlayDeposit)
	}
	if t.StarTransactionTypeFragmentDeposit != nil {
		return json.Marshal(t.StarTransactionTypeFragmentDeposit)
	}
	if t.StarTransactionTypeUserDeposit != nil {
		return json.Marshal(t.StarTransactionTypeUserDeposit)
	}
	if t.StarTransactionTypeGiveawayDeposit != nil {
		return json.Marshal(t.StarTransactionTypeGiveawayDeposit)
	}
	if t.StarTransactionTypeFragmentWithdrawal != nil {
		return json.Marshal(t.StarTransactionTypeFragmentWithdrawal)
	}
	if t.StarTransactionTypeTelegramAdsWithdrawal != nil {
		return json.Marshal(t.StarTransactionTypeTelegramAdsWithdrawal)
	}
	if t.StarTransactionTypeTelegramApiUsage != nil {
		return json.Marshal(t.StarTransactionTypeTelegramApiUsage)
	}
	if t.StarTransactionTypeBotPaidMediaPurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotPaidMediaPurchase)
	}
	if t.StarTransactionTypeBotPaidMediaSale != nil {
		return json.Marshal(t.StarTransactionTypeBotPaidMediaSale)
	}
	if t.StarTransactionTypeChannelPaidMediaPurchase != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidMediaPurchase)
	}
	if t.StarTransactionTypeChannelPaidMediaSale != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidMediaSale)
	}
	if t.StarTransactionTypeBotInvoicePurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotInvoicePurchase)
	}
	if t.StarTransactionTypeBotInvoiceSale != nil {
		return json.Marshal(t.StarTransactionTypeBotInvoiceSale)
	}
	if t.StarTransactionTypeBotSubscriptionPurchase != nil {
		return json.Marshal(t.StarTransactionTypeBotSubscriptionPurchase)
	}
	if t.StarTransactionTypeBotSubscriptionSale != nil {
		return json.Marshal(t.StarTransactionTypeBotSubscriptionSale)
	}
	if t.StarTransactionTypeChannelSubscriptionPurchase != nil {
		return json.Marshal(t.StarTransactionTypeChannelSubscriptionPurchase)
	}
	if t.StarTransactionTypeChannelSubscriptionSale != nil {
		return json.Marshal(t.StarTransactionTypeChannelSubscriptionSale)
	}
	if t.StarTransactionTypeGiftAuctionBid != nil {
		return json.Marshal(t.StarTransactionTypeGiftAuctionBid)
	}
	if t.StarTransactionTypeGiftPurchase != nil {
		return json.Marshal(t.StarTransactionTypeGiftPurchase)
	}
	if t.StarTransactionTypeGiftPurchaseOffer != nil {
		return json.Marshal(t.StarTransactionTypeGiftPurchaseOffer)
	}
	if t.StarTransactionTypeGiftTransfer != nil {
		return json.Marshal(t.StarTransactionTypeGiftTransfer)
	}
	if t.StarTransactionTypeGiftOriginalDetailsDrop != nil {
		return json.Marshal(t.StarTransactionTypeGiftOriginalDetailsDrop)
	}
	if t.StarTransactionTypeGiftSale != nil {
		return json.Marshal(t.StarTransactionTypeGiftSale)
	}
	if t.StarTransactionTypeGiftUpgrade != nil {
		return json.Marshal(t.StarTransactionTypeGiftUpgrade)
	}
	if t.StarTransactionTypeGiftUpgradePurchase != nil {
		return json.Marshal(t.StarTransactionTypeGiftUpgradePurchase)
	}
	if t.StarTransactionTypeUpgradedGiftPurchase != nil {
		return json.Marshal(t.StarTransactionTypeUpgradedGiftPurchase)
	}
	if t.StarTransactionTypeUpgradedGiftSale != nil {
		return json.Marshal(t.StarTransactionTypeUpgradedGiftSale)
	}
	if t.StarTransactionTypeChannelPaidReactionSend != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidReactionSend)
	}
	if t.StarTransactionTypeChannelPaidReactionReceive != nil {
		return json.Marshal(t.StarTransactionTypeChannelPaidReactionReceive)
	}
	if t.StarTransactionTypeAffiliateProgramCommission != nil {
		return json.Marshal(t.StarTransactionTypeAffiliateProgramCommission)
	}
	if t.StarTransactionTypePaidMessageSend != nil {
		return json.Marshal(t.StarTransactionTypePaidMessageSend)
	}
	if t.StarTransactionTypePaidMessageReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidMessageReceive)
	}
	if t.StarTransactionTypePaidGroupCallMessageSend != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallMessageSend)
	}
	if t.StarTransactionTypePaidGroupCallMessageReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallMessageReceive)
	}
	if t.StarTransactionTypePaidGroupCallReactionSend != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallReactionSend)
	}
	if t.StarTransactionTypePaidGroupCallReactionReceive != nil {
		return json.Marshal(t.StarTransactionTypePaidGroupCallReactionReceive)
	}
	if t.StarTransactionTypeSuggestedPostPaymentSend != nil {
		return json.Marshal(t.StarTransactionTypeSuggestedPostPaymentSend)
	}
	if t.StarTransactionTypeSuggestedPostPaymentReceive != nil {
		return json.Marshal(t.StarTransactionTypeSuggestedPostPaymentReceive)
	}
	if t.StarTransactionTypePremiumPurchase != nil {
		return json.Marshal(t.StarTransactionTypePremiumPurchase)
	}
	if t.StarTransactionTypeBusinessBotTransferSend != nil {
		return json.Marshal(t.StarTransactionTypeBusinessBotTransferSend)
	}
	if t.StarTransactionTypeBusinessBotTransferReceive != nil {
		return json.Marshal(t.StarTransactionTypeBusinessBotTransferReceive)
	}
	if t.StarTransactionTypePublicPostSearch != nil {
		return json.Marshal(t.StarTransactionTypePublicPostSearch)
	}
	if t.StarTransactionTypeUnsupported != nil {
		return json.Marshal(t.StarTransactionTypeUnsupported)
	}
	return nil, fmt.Errorf("no value set for StarTransactionType")
}

// ChatSource Describes a reason why an external chat is shown in a chat list
type ChatSource struct {
	ChatSourceMtprotoProxy              *ChatSourceMtprotoProxy              `json:"chatSourceMtprotoProxy,omitempty"`
	ChatSourcePublicServiceAnnouncement *ChatSourcePublicServiceAnnouncement `json:"chatSourcePublicServiceAnnouncement,omitempty"`
}

func (t *ChatSource) Type() string {
	return ""
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

// CanPostStoryResult Represents result of checking whether the current user can post a story on behalf of the specific chat
type CanPostStoryResult struct {
	CanPostStoryResultOk                       *CanPostStoryResultOk                       `json:"canPostStoryResultOk,omitempty"`
	CanPostStoryResultPremiumNeeded            *CanPostStoryResultPremiumNeeded            `json:"canPostStoryResultPremiumNeeded,omitempty"`
	CanPostStoryResultBoostNeeded              *CanPostStoryResultBoostNeeded              `json:"canPostStoryResultBoostNeeded,omitempty"`
	CanPostStoryResultActiveStoryLimitExceeded *CanPostStoryResultActiveStoryLimitExceeded `json:"canPostStoryResultActiveStoryLimitExceeded,omitempty"`
	CanPostStoryResultWeeklyLimitExceeded      *CanPostStoryResultWeeklyLimitExceeded      `json:"canPostStoryResultWeeklyLimitExceeded,omitempty"`
	CanPostStoryResultMonthlyLimitExceeded     *CanPostStoryResultMonthlyLimitExceeded     `json:"canPostStoryResultMonthlyLimitExceeded,omitempty"`
	CanPostStoryResultLiveStoryIsActive        *CanPostStoryResultLiveStoryIsActive        `json:"canPostStoryResultLiveStoryIsActive,omitempty"`
}

func (t *CanPostStoryResult) Type() string {
	return ""
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
	switch typeObj.Type {
	case "canPostStoryResultOk":
		t.CanPostStoryResultOk = new(CanPostStoryResultOk)
		return json.Unmarshal(b, t.CanPostStoryResultOk)
	case "canPostStoryResultPremiumNeeded":
		t.CanPostStoryResultPremiumNeeded = new(CanPostStoryResultPremiumNeeded)
		return json.Unmarshal(b, t.CanPostStoryResultPremiumNeeded)
	case "canPostStoryResultBoostNeeded":
		t.CanPostStoryResultBoostNeeded = new(CanPostStoryResultBoostNeeded)
		return json.Unmarshal(b, t.CanPostStoryResultBoostNeeded)
	case "canPostStoryResultActiveStoryLimitExceeded":
		t.CanPostStoryResultActiveStoryLimitExceeded = new(CanPostStoryResultActiveStoryLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultActiveStoryLimitExceeded)
	case "canPostStoryResultWeeklyLimitExceeded":
		t.CanPostStoryResultWeeklyLimitExceeded = new(CanPostStoryResultWeeklyLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultWeeklyLimitExceeded)
	case "canPostStoryResultMonthlyLimitExceeded":
		t.CanPostStoryResultMonthlyLimitExceeded = new(CanPostStoryResultMonthlyLimitExceeded)
		return json.Unmarshal(b, t.CanPostStoryResultMonthlyLimitExceeded)
	case "canPostStoryResultLiveStoryIsActive":
		t.CanPostStoryResultLiveStoryIsActive = new(CanPostStoryResultLiveStoryIsActive)
		return json.Unmarshal(b, t.CanPostStoryResultLiveStoryIsActive)
	}
	return nil
}

func (t *CanPostStoryResult) MarshalJSON() ([]byte, error) {
	if t.CanPostStoryResultOk != nil {
		return json.Marshal(t.CanPostStoryResultOk)
	}
	if t.CanPostStoryResultPremiumNeeded != nil {
		return json.Marshal(t.CanPostStoryResultPremiumNeeded)
	}
	if t.CanPostStoryResultBoostNeeded != nil {
		return json.Marshal(t.CanPostStoryResultBoostNeeded)
	}
	if t.CanPostStoryResultActiveStoryLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultActiveStoryLimitExceeded)
	}
	if t.CanPostStoryResultWeeklyLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultWeeklyLimitExceeded)
	}
	if t.CanPostStoryResultMonthlyLimitExceeded != nil {
		return json.Marshal(t.CanPostStoryResultMonthlyLimitExceeded)
	}
	if t.CanPostStoryResultLiveStoryIsActive != nil {
		return json.Marshal(t.CanPostStoryResultLiveStoryIsActive)
	}
	return nil, fmt.Errorf("no value set for CanPostStoryResult")
}

// SuggestedAction Describes an action suggested to the current user
type SuggestedAction struct {
	SuggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats `json:"suggestedActionEnableArchiveAndMuteNewChats,omitempty"`
	SuggestedActionCheckPassword                *SuggestedActionCheckPassword                `json:"suggestedActionCheckPassword,omitempty"`
	SuggestedActionCheckPhoneNumber             *SuggestedActionCheckPhoneNumber             `json:"suggestedActionCheckPhoneNumber,omitempty"`
	SuggestedActionViewChecksHint               *SuggestedActionViewChecksHint               `json:"suggestedActionViewChecksHint,omitempty"`
	SuggestedActionConvertToBroadcastGroup      *SuggestedActionConvertToBroadcastGroup      `json:"suggestedActionConvertToBroadcastGroup,omitempty"`
	SuggestedActionSetPassword                  *SuggestedActionSetPassword                  `json:"suggestedActionSetPassword,omitempty"`
	SuggestedActionUpgradePremium               *SuggestedActionUpgradePremium               `json:"suggestedActionUpgradePremium,omitempty"`
	SuggestedActionRestorePremium               *SuggestedActionRestorePremium               `json:"suggestedActionRestorePremium,omitempty"`
	SuggestedActionSubscribeToAnnualPremium     *SuggestedActionSubscribeToAnnualPremium     `json:"suggestedActionSubscribeToAnnualPremium,omitempty"`
	SuggestedActionGiftPremiumForChristmas      *SuggestedActionGiftPremiumForChristmas      `json:"suggestedActionGiftPremiumForChristmas,omitempty"`
	SuggestedActionSetBirthdate                 *SuggestedActionSetBirthdate                 `json:"suggestedActionSetBirthdate,omitempty"`
	SuggestedActionSetProfilePhoto              *SuggestedActionSetProfilePhoto              `json:"suggestedActionSetProfilePhoto,omitempty"`
	SuggestedActionExtendPremium                *SuggestedActionExtendPremium                `json:"suggestedActionExtendPremium,omitempty"`
	SuggestedActionExtendStarSubscriptions      *SuggestedActionExtendStarSubscriptions      `json:"suggestedActionExtendStarSubscriptions,omitempty"`
	SuggestedActionCustom                       *SuggestedActionCustom                       `json:"suggestedActionCustom,omitempty"`
	SuggestedActionSetLoginEmailAddress         *SuggestedActionSetLoginEmailAddress         `json:"suggestedActionSetLoginEmailAddress,omitempty"`
	SuggestedActionAddLoginPasskey              *SuggestedActionAddLoginPasskey              `json:"suggestedActionAddLoginPasskey,omitempty"`
}

func (t *SuggestedAction) Type() string {
	return ""
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
	switch typeObj.Type {
	case "suggestedActionEnableArchiveAndMuteNewChats":
		t.SuggestedActionEnableArchiveAndMuteNewChats = new(SuggestedActionEnableArchiveAndMuteNewChats)
		return json.Unmarshal(b, t.SuggestedActionEnableArchiveAndMuteNewChats)
	case "suggestedActionCheckPassword":
		t.SuggestedActionCheckPassword = new(SuggestedActionCheckPassword)
		return json.Unmarshal(b, t.SuggestedActionCheckPassword)
	case "suggestedActionCheckPhoneNumber":
		t.SuggestedActionCheckPhoneNumber = new(SuggestedActionCheckPhoneNumber)
		return json.Unmarshal(b, t.SuggestedActionCheckPhoneNumber)
	case "suggestedActionViewChecksHint":
		t.SuggestedActionViewChecksHint = new(SuggestedActionViewChecksHint)
		return json.Unmarshal(b, t.SuggestedActionViewChecksHint)
	case "suggestedActionConvertToBroadcastGroup":
		t.SuggestedActionConvertToBroadcastGroup = new(SuggestedActionConvertToBroadcastGroup)
		return json.Unmarshal(b, t.SuggestedActionConvertToBroadcastGroup)
	case "suggestedActionSetPassword":
		t.SuggestedActionSetPassword = new(SuggestedActionSetPassword)
		return json.Unmarshal(b, t.SuggestedActionSetPassword)
	case "suggestedActionUpgradePremium":
		t.SuggestedActionUpgradePremium = new(SuggestedActionUpgradePremium)
		return json.Unmarshal(b, t.SuggestedActionUpgradePremium)
	case "suggestedActionRestorePremium":
		t.SuggestedActionRestorePremium = new(SuggestedActionRestorePremium)
		return json.Unmarshal(b, t.SuggestedActionRestorePremium)
	case "suggestedActionSubscribeToAnnualPremium":
		t.SuggestedActionSubscribeToAnnualPremium = new(SuggestedActionSubscribeToAnnualPremium)
		return json.Unmarshal(b, t.SuggestedActionSubscribeToAnnualPremium)
	case "suggestedActionGiftPremiumForChristmas":
		t.SuggestedActionGiftPremiumForChristmas = new(SuggestedActionGiftPremiumForChristmas)
		return json.Unmarshal(b, t.SuggestedActionGiftPremiumForChristmas)
	case "suggestedActionSetBirthdate":
		t.SuggestedActionSetBirthdate = new(SuggestedActionSetBirthdate)
		return json.Unmarshal(b, t.SuggestedActionSetBirthdate)
	case "suggestedActionSetProfilePhoto":
		t.SuggestedActionSetProfilePhoto = new(SuggestedActionSetProfilePhoto)
		return json.Unmarshal(b, t.SuggestedActionSetProfilePhoto)
	case "suggestedActionExtendPremium":
		t.SuggestedActionExtendPremium = new(SuggestedActionExtendPremium)
		return json.Unmarshal(b, t.SuggestedActionExtendPremium)
	case "suggestedActionExtendStarSubscriptions":
		t.SuggestedActionExtendStarSubscriptions = new(SuggestedActionExtendStarSubscriptions)
		return json.Unmarshal(b, t.SuggestedActionExtendStarSubscriptions)
	case "suggestedActionCustom":
		t.SuggestedActionCustom = new(SuggestedActionCustom)
		return json.Unmarshal(b, t.SuggestedActionCustom)
	case "suggestedActionSetLoginEmailAddress":
		t.SuggestedActionSetLoginEmailAddress = new(SuggestedActionSetLoginEmailAddress)
		return json.Unmarshal(b, t.SuggestedActionSetLoginEmailAddress)
	case "suggestedActionAddLoginPasskey":
		t.SuggestedActionAddLoginPasskey = new(SuggestedActionAddLoginPasskey)
		return json.Unmarshal(b, t.SuggestedActionAddLoginPasskey)
	}
	return nil
}

func (t *SuggestedAction) MarshalJSON() ([]byte, error) {
	if t.SuggestedActionEnableArchiveAndMuteNewChats != nil {
		return json.Marshal(t.SuggestedActionEnableArchiveAndMuteNewChats)
	}
	if t.SuggestedActionCheckPassword != nil {
		return json.Marshal(t.SuggestedActionCheckPassword)
	}
	if t.SuggestedActionCheckPhoneNumber != nil {
		return json.Marshal(t.SuggestedActionCheckPhoneNumber)
	}
	if t.SuggestedActionViewChecksHint != nil {
		return json.Marshal(t.SuggestedActionViewChecksHint)
	}
	if t.SuggestedActionConvertToBroadcastGroup != nil {
		return json.Marshal(t.SuggestedActionConvertToBroadcastGroup)
	}
	if t.SuggestedActionSetPassword != nil {
		return json.Marshal(t.SuggestedActionSetPassword)
	}
	if t.SuggestedActionUpgradePremium != nil {
		return json.Marshal(t.SuggestedActionUpgradePremium)
	}
	if t.SuggestedActionRestorePremium != nil {
		return json.Marshal(t.SuggestedActionRestorePremium)
	}
	if t.SuggestedActionSubscribeToAnnualPremium != nil {
		return json.Marshal(t.SuggestedActionSubscribeToAnnualPremium)
	}
	if t.SuggestedActionGiftPremiumForChristmas != nil {
		return json.Marshal(t.SuggestedActionGiftPremiumForChristmas)
	}
	if t.SuggestedActionSetBirthdate != nil {
		return json.Marshal(t.SuggestedActionSetBirthdate)
	}
	if t.SuggestedActionSetProfilePhoto != nil {
		return json.Marshal(t.SuggestedActionSetProfilePhoto)
	}
	if t.SuggestedActionExtendPremium != nil {
		return json.Marshal(t.SuggestedActionExtendPremium)
	}
	if t.SuggestedActionExtendStarSubscriptions != nil {
		return json.Marshal(t.SuggestedActionExtendStarSubscriptions)
	}
	if t.SuggestedActionCustom != nil {
		return json.Marshal(t.SuggestedActionCustom)
	}
	if t.SuggestedActionSetLoginEmailAddress != nil {
		return json.Marshal(t.SuggestedActionSetLoginEmailAddress)
	}
	if t.SuggestedActionAddLoginPasskey != nil {
		return json.Marshal(t.SuggestedActionAddLoginPasskey)
	}
	return nil, fmt.Errorf("no value set for SuggestedAction")
}

// ProxyType Describes the type of proxy server
type ProxyType struct {
	ProxyTypeSocks5  *ProxyTypeSocks5  `json:"proxyTypeSocks5,omitempty"`
	ProxyTypeHttp    *ProxyTypeHttp    `json:"proxyTypeHttp,omitempty"`
	ProxyTypeMtproto *ProxyTypeMtproto `json:"proxyTypeMtproto,omitempty"`
}

func (t *ProxyType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "proxyTypeSocks5":
		t.ProxyTypeSocks5 = new(ProxyTypeSocks5)
		return json.Unmarshal(b, t.ProxyTypeSocks5)
	case "proxyTypeHttp":
		t.ProxyTypeHttp = new(ProxyTypeHttp)
		return json.Unmarshal(b, t.ProxyTypeHttp)
	case "proxyTypeMtproto":
		t.ProxyTypeMtproto = new(ProxyTypeMtproto)
		return json.Unmarshal(b, t.ProxyTypeMtproto)
	}
	return nil
}

func (t *ProxyType) MarshalJSON() ([]byte, error) {
	if t.ProxyTypeSocks5 != nil {
		return json.Marshal(t.ProxyTypeSocks5)
	}
	if t.ProxyTypeHttp != nil {
		return json.Marshal(t.ProxyTypeHttp)
	}
	if t.ProxyTypeMtproto != nil {
		return json.Marshal(t.ProxyTypeMtproto)
	}
	return nil, fmt.Errorf("no value set for ProxyType")
}

// ChatRevenueTransactionType Describes type of transaction for revenue earned from sponsored messages in a chat
type ChatRevenueTransactionType struct {
	ChatRevenueTransactionTypeUnsupported              *ChatRevenueTransactionTypeUnsupported              `json:"chatRevenueTransactionTypeUnsupported,omitempty"`
	ChatRevenueTransactionTypeSponsoredMessageEarnings *ChatRevenueTransactionTypeSponsoredMessageEarnings `json:"chatRevenueTransactionTypeSponsoredMessageEarnings,omitempty"`
	ChatRevenueTransactionTypeSuggestedPostEarnings    *ChatRevenueTransactionTypeSuggestedPostEarnings    `json:"chatRevenueTransactionTypeSuggestedPostEarnings,omitempty"`
	ChatRevenueTransactionTypeFragmentWithdrawal       *ChatRevenueTransactionTypeFragmentWithdrawal       `json:"chatRevenueTransactionTypeFragmentWithdrawal,omitempty"`
	ChatRevenueTransactionTypeFragmentRefund           *ChatRevenueTransactionTypeFragmentRefund           `json:"chatRevenueTransactionTypeFragmentRefund,omitempty"`
}

func (t *ChatRevenueTransactionType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "chatRevenueTransactionTypeUnsupported":
		t.ChatRevenueTransactionTypeUnsupported = new(ChatRevenueTransactionTypeUnsupported)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeUnsupported)
	case "chatRevenueTransactionTypeSponsoredMessageEarnings":
		t.ChatRevenueTransactionTypeSponsoredMessageEarnings = new(ChatRevenueTransactionTypeSponsoredMessageEarnings)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeSponsoredMessageEarnings)
	case "chatRevenueTransactionTypeSuggestedPostEarnings":
		t.ChatRevenueTransactionTypeSuggestedPostEarnings = new(ChatRevenueTransactionTypeSuggestedPostEarnings)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeSuggestedPostEarnings)
	case "chatRevenueTransactionTypeFragmentWithdrawal":
		t.ChatRevenueTransactionTypeFragmentWithdrawal = new(ChatRevenueTransactionTypeFragmentWithdrawal)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeFragmentWithdrawal)
	case "chatRevenueTransactionTypeFragmentRefund":
		t.ChatRevenueTransactionTypeFragmentRefund = new(ChatRevenueTransactionTypeFragmentRefund)
		return json.Unmarshal(b, t.ChatRevenueTransactionTypeFragmentRefund)
	}
	return nil
}

func (t *ChatRevenueTransactionType) MarshalJSON() ([]byte, error) {
	if t.ChatRevenueTransactionTypeUnsupported != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeUnsupported)
	}
	if t.ChatRevenueTransactionTypeSponsoredMessageEarnings != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeSponsoredMessageEarnings)
	}
	if t.ChatRevenueTransactionTypeSuggestedPostEarnings != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeSuggestedPostEarnings)
	}
	if t.ChatRevenueTransactionTypeFragmentWithdrawal != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeFragmentWithdrawal)
	}
	if t.ChatRevenueTransactionTypeFragmentRefund != nil {
		return json.Marshal(t.ChatRevenueTransactionTypeFragmentRefund)
	}
	return nil, fmt.Errorf("no value set for ChatRevenueTransactionType")
}

// InputMessageReplyTo Contains information about the message or the story to be replied
type InputMessageReplyTo struct {
	InputMessageReplyToMessage         *InputMessageReplyToMessage         `json:"inputMessageReplyToMessage,omitempty"`
	InputMessageReplyToExternalMessage *InputMessageReplyToExternalMessage `json:"inputMessageReplyToExternalMessage,omitempty"`
	InputMessageReplyToStory           *InputMessageReplyToStory           `json:"inputMessageReplyToStory,omitempty"`
}

func (t *InputMessageReplyTo) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputMessageReplyToMessage":
		t.InputMessageReplyToMessage = new(InputMessageReplyToMessage)
		return json.Unmarshal(b, t.InputMessageReplyToMessage)
	case "inputMessageReplyToExternalMessage":
		t.InputMessageReplyToExternalMessage = new(InputMessageReplyToExternalMessage)
		return json.Unmarshal(b, t.InputMessageReplyToExternalMessage)
	case "inputMessageReplyToStory":
		t.InputMessageReplyToStory = new(InputMessageReplyToStory)
		return json.Unmarshal(b, t.InputMessageReplyToStory)
	}
	return nil
}

func (t *InputMessageReplyTo) MarshalJSON() ([]byte, error) {
	if t.InputMessageReplyToMessage != nil {
		return json.Marshal(t.InputMessageReplyToMessage)
	}
	if t.InputMessageReplyToExternalMessage != nil {
		return json.Marshal(t.InputMessageReplyToExternalMessage)
	}
	if t.InputMessageReplyToStory != nil {
		return json.Marshal(t.InputMessageReplyToStory)
	}
	return nil, fmt.Errorf("no value set for InputMessageReplyTo")
}

// LinkPreviewAlbumMedia Describes a media from a link preview album
type LinkPreviewAlbumMedia struct {
	LinkPreviewAlbumMediaPhoto *LinkPreviewAlbumMediaPhoto `json:"linkPreviewAlbumMediaPhoto,omitempty"`
	LinkPreviewAlbumMediaVideo *LinkPreviewAlbumMediaVideo `json:"linkPreviewAlbumMediaVideo,omitempty"`
}

func (t *LinkPreviewAlbumMedia) Type() string {
	return ""
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

// MessageSchedulingState Contains information about the time when a scheduled message will be sent
type MessageSchedulingState struct {
	MessageSchedulingStateSendAtDate             *MessageSchedulingStateSendAtDate             `json:"messageSchedulingStateSendAtDate,omitempty"`
	MessageSchedulingStateSendWhenOnline         *MessageSchedulingStateSendWhenOnline         `json:"messageSchedulingStateSendWhenOnline,omitempty"`
	MessageSchedulingStateSendWhenVideoProcessed *MessageSchedulingStateSendWhenVideoProcessed `json:"messageSchedulingStateSendWhenVideoProcessed,omitempty"`
}

func (t *MessageSchedulingState) Type() string {
	return ""
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

// InputStoryAreaType Describes type of clickable area on a story media to be added
type InputStoryAreaType struct {
	InputStoryAreaTypeLocation          *InputStoryAreaTypeLocation          `json:"inputStoryAreaTypeLocation,omitempty"`
	InputStoryAreaTypeFoundVenue        *InputStoryAreaTypeFoundVenue        `json:"inputStoryAreaTypeFoundVenue,omitempty"`
	InputStoryAreaTypePreviousVenue     *InputStoryAreaTypePreviousVenue     `json:"inputStoryAreaTypePreviousVenue,omitempty"`
	InputStoryAreaTypeSuggestedReaction *InputStoryAreaTypeSuggestedReaction `json:"inputStoryAreaTypeSuggestedReaction,omitempty"`
	InputStoryAreaTypeMessage           *InputStoryAreaTypeMessage           `json:"inputStoryAreaTypeMessage,omitempty"`
	InputStoryAreaTypeLink              *InputStoryAreaTypeLink              `json:"inputStoryAreaTypeLink,omitempty"`
	InputStoryAreaTypeWeather           *InputStoryAreaTypeWeather           `json:"inputStoryAreaTypeWeather,omitempty"`
	InputStoryAreaTypeUpgradedGift      *InputStoryAreaTypeUpgradedGift      `json:"inputStoryAreaTypeUpgradedGift,omitempty"`
}

func (t *InputStoryAreaType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputStoryAreaTypeLocation":
		t.InputStoryAreaTypeLocation = new(InputStoryAreaTypeLocation)
		return json.Unmarshal(b, t.InputStoryAreaTypeLocation)
	case "inputStoryAreaTypeFoundVenue":
		t.InputStoryAreaTypeFoundVenue = new(InputStoryAreaTypeFoundVenue)
		return json.Unmarshal(b, t.InputStoryAreaTypeFoundVenue)
	case "inputStoryAreaTypePreviousVenue":
		t.InputStoryAreaTypePreviousVenue = new(InputStoryAreaTypePreviousVenue)
		return json.Unmarshal(b, t.InputStoryAreaTypePreviousVenue)
	case "inputStoryAreaTypeSuggestedReaction":
		t.InputStoryAreaTypeSuggestedReaction = new(InputStoryAreaTypeSuggestedReaction)
		return json.Unmarshal(b, t.InputStoryAreaTypeSuggestedReaction)
	case "inputStoryAreaTypeMessage":
		t.InputStoryAreaTypeMessage = new(InputStoryAreaTypeMessage)
		return json.Unmarshal(b, t.InputStoryAreaTypeMessage)
	case "inputStoryAreaTypeLink":
		t.InputStoryAreaTypeLink = new(InputStoryAreaTypeLink)
		return json.Unmarshal(b, t.InputStoryAreaTypeLink)
	case "inputStoryAreaTypeWeather":
		t.InputStoryAreaTypeWeather = new(InputStoryAreaTypeWeather)
		return json.Unmarshal(b, t.InputStoryAreaTypeWeather)
	case "inputStoryAreaTypeUpgradedGift":
		t.InputStoryAreaTypeUpgradedGift = new(InputStoryAreaTypeUpgradedGift)
		return json.Unmarshal(b, t.InputStoryAreaTypeUpgradedGift)
	}
	return nil
}

func (t *InputStoryAreaType) MarshalJSON() ([]byte, error) {
	if t.InputStoryAreaTypeLocation != nil {
		return json.Marshal(t.InputStoryAreaTypeLocation)
	}
	if t.InputStoryAreaTypeFoundVenue != nil {
		return json.Marshal(t.InputStoryAreaTypeFoundVenue)
	}
	if t.InputStoryAreaTypePreviousVenue != nil {
		return json.Marshal(t.InputStoryAreaTypePreviousVenue)
	}
	if t.InputStoryAreaTypeSuggestedReaction != nil {
		return json.Marshal(t.InputStoryAreaTypeSuggestedReaction)
	}
	if t.InputStoryAreaTypeMessage != nil {
		return json.Marshal(t.InputStoryAreaTypeMessage)
	}
	if t.InputStoryAreaTypeLink != nil {
		return json.Marshal(t.InputStoryAreaTypeLink)
	}
	if t.InputStoryAreaTypeWeather != nil {
		return json.Marshal(t.InputStoryAreaTypeWeather)
	}
	if t.InputStoryAreaTypeUpgradedGift != nil {
		return json.Marshal(t.InputStoryAreaTypeUpgradedGift)
	}
	return nil, fmt.Errorf("no value set for InputStoryAreaType")
}

// StickerFormat Describes format of a sticker
type StickerFormat struct {
	StickerFormatWebp *StickerFormatWebp `json:"stickerFormatWebp,omitempty"`
	StickerFormatTgs  *StickerFormatTgs  `json:"stickerFormatTgs,omitempty"`
	StickerFormatWebm *StickerFormatWebm `json:"stickerFormatWebm,omitempty"`
}

func (t *StickerFormat) Type() string {
	return ""
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
	switch typeObj.Type {
	case "stickerFormatWebp":
		t.StickerFormatWebp = new(StickerFormatWebp)
		return json.Unmarshal(b, t.StickerFormatWebp)
	case "stickerFormatTgs":
		t.StickerFormatTgs = new(StickerFormatTgs)
		return json.Unmarshal(b, t.StickerFormatTgs)
	case "stickerFormatWebm":
		t.StickerFormatWebm = new(StickerFormatWebm)
		return json.Unmarshal(b, t.StickerFormatWebm)
	}
	return nil
}

func (t *StickerFormat) MarshalJSON() ([]byte, error) {
	if t.StickerFormatWebp != nil {
		return json.Marshal(t.StickerFormatWebp)
	}
	if t.StickerFormatTgs != nil {
		return json.Marshal(t.StickerFormatTgs)
	}
	if t.StickerFormatWebm != nil {
		return json.Marshal(t.StickerFormatWebm)
	}
	return nil, fmt.Errorf("no value set for StickerFormat")
}

// AuctionState Describes state of an auction
type AuctionState struct {
	AuctionStateActive   *AuctionStateActive   `json:"auctionStateActive,omitempty"`
	AuctionStateFinished *AuctionStateFinished `json:"auctionStateFinished,omitempty"`
}

func (t *AuctionState) Type() string {
	return ""
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

// ActiveStoryState Describes state of active stories posted by a chat
type ActiveStoryState struct {
	ActiveStoryStateLive   *ActiveStoryStateLive   `json:"activeStoryStateLive,omitempty"`
	ActiveStoryStateUnread *ActiveStoryStateUnread `json:"activeStoryStateUnread,omitempty"`
	ActiveStoryStateRead   *ActiveStoryStateRead   `json:"activeStoryStateRead,omitempty"`
}

func (t *ActiveStoryState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "activeStoryStateLive":
		t.ActiveStoryStateLive = new(ActiveStoryStateLive)
		return json.Unmarshal(b, t.ActiveStoryStateLive)
	case "activeStoryStateUnread":
		t.ActiveStoryStateUnread = new(ActiveStoryStateUnread)
		return json.Unmarshal(b, t.ActiveStoryStateUnread)
	case "activeStoryStateRead":
		t.ActiveStoryStateRead = new(ActiveStoryStateRead)
		return json.Unmarshal(b, t.ActiveStoryStateRead)
	}
	return nil
}

func (t *ActiveStoryState) MarshalJSON() ([]byte, error) {
	if t.ActiveStoryStateLive != nil {
		return json.Marshal(t.ActiveStoryStateLive)
	}
	if t.ActiveStoryStateUnread != nil {
		return json.Marshal(t.ActiveStoryStateUnread)
	}
	if t.ActiveStoryStateRead != nil {
		return json.Marshal(t.ActiveStoryStateRead)
	}
	return nil, fmt.Errorf("no value set for ActiveStoryState")
}

// EmojiStatusType Describes type of emoji status
type EmojiStatusType struct {
	EmojiStatusTypeCustomEmoji  *EmojiStatusTypeCustomEmoji  `json:"emojiStatusTypeCustomEmoji,omitempty"`
	EmojiStatusTypeUpgradedGift *EmojiStatusTypeUpgradedGift `json:"emojiStatusTypeUpgradedGift,omitempty"`
}

func (t *EmojiStatusType) Type() string {
	return ""
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

// InviteLinkChatType Describes the type of chat to which points an invite link
type InviteLinkChatType struct {
	InviteLinkChatTypeBasicGroup *InviteLinkChatTypeBasicGroup `json:"inviteLinkChatTypeBasicGroup,omitempty"`
	InviteLinkChatTypeSupergroup *InviteLinkChatTypeSupergroup `json:"inviteLinkChatTypeSupergroup,omitempty"`
	InviteLinkChatTypeChannel    *InviteLinkChatTypeChannel    `json:"inviteLinkChatTypeChannel,omitempty"`
}

func (t *InviteLinkChatType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inviteLinkChatTypeBasicGroup":
		t.InviteLinkChatTypeBasicGroup = new(InviteLinkChatTypeBasicGroup)
		return json.Unmarshal(b, t.InviteLinkChatTypeBasicGroup)
	case "inviteLinkChatTypeSupergroup":
		t.InviteLinkChatTypeSupergroup = new(InviteLinkChatTypeSupergroup)
		return json.Unmarshal(b, t.InviteLinkChatTypeSupergroup)
	case "inviteLinkChatTypeChannel":
		t.InviteLinkChatTypeChannel = new(InviteLinkChatTypeChannel)
		return json.Unmarshal(b, t.InviteLinkChatTypeChannel)
	}
	return nil
}

func (t *InviteLinkChatType) MarshalJSON() ([]byte, error) {
	if t.InviteLinkChatTypeBasicGroup != nil {
		return json.Marshal(t.InviteLinkChatTypeBasicGroup)
	}
	if t.InviteLinkChatTypeSupergroup != nil {
		return json.Marshal(t.InviteLinkChatTypeSupergroup)
	}
	if t.InviteLinkChatTypeChannel != nil {
		return json.Marshal(t.InviteLinkChatTypeChannel)
	}
	return nil, fmt.Errorf("no value set for InviteLinkChatType")
}

// PublicChatType Describes type of public chat
type PublicChatType struct {
	PublicChatTypeHasUsername     *PublicChatTypeHasUsername     `json:"publicChatTypeHasUsername,omitempty"`
	PublicChatTypeIsLocationBased *PublicChatTypeIsLocationBased `json:"publicChatTypeIsLocationBased,omitempty"`
}

func (t *PublicChatType) Type() string {
	return ""
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

// PageBlock Describes a block of an instant view for a web page
type PageBlock struct {
	PageBlockTitle           *PageBlockTitle           `json:"pageBlockTitle,omitempty"`
	PageBlockSubtitle        *PageBlockSubtitle        `json:"pageBlockSubtitle,omitempty"`
	PageBlockAuthorDate      *PageBlockAuthorDate      `json:"pageBlockAuthorDate,omitempty"`
	PageBlockHeader          *PageBlockHeader          `json:"pageBlockHeader,omitempty"`
	PageBlockSubheader       *PageBlockSubheader       `json:"pageBlockSubheader,omitempty"`
	PageBlockKicker          *PageBlockKicker          `json:"pageBlockKicker,omitempty"`
	PageBlockParagraph       *PageBlockParagraph       `json:"pageBlockParagraph,omitempty"`
	PageBlockPreformatted    *PageBlockPreformatted    `json:"pageBlockPreformatted,omitempty"`
	PageBlockFooter          *PageBlockFooter          `json:"pageBlockFooter,omitempty"`
	PageBlockDivider         *PageBlockDivider         `json:"pageBlockDivider,omitempty"`
	PageBlockAnchor          *PageBlockAnchor          `json:"pageBlockAnchor,omitempty"`
	PageBlockList            *PageBlockList            `json:"pageBlockList,omitempty"`
	PageBlockBlockQuote      *PageBlockBlockQuote      `json:"pageBlockBlockQuote,omitempty"`
	PageBlockPullQuote       *PageBlockPullQuote       `json:"pageBlockPullQuote,omitempty"`
	PageBlockAnimation       *PageBlockAnimation       `json:"pageBlockAnimation,omitempty"`
	PageBlockAudio           *PageBlockAudio           `json:"pageBlockAudio,omitempty"`
	PageBlockPhoto           *PageBlockPhoto           `json:"pageBlockPhoto,omitempty"`
	PageBlockVideo           *PageBlockVideo           `json:"pageBlockVideo,omitempty"`
	PageBlockVoiceNote       *PageBlockVoiceNote       `json:"pageBlockVoiceNote,omitempty"`
	PageBlockCover           *PageBlockCover           `json:"pageBlockCover,omitempty"`
	PageBlockEmbedded        *PageBlockEmbedded        `json:"pageBlockEmbedded,omitempty"`
	PageBlockEmbeddedPost    *PageBlockEmbeddedPost    `json:"pageBlockEmbeddedPost,omitempty"`
	PageBlockCollage         *PageBlockCollage         `json:"pageBlockCollage,omitempty"`
	PageBlockSlideshow       *PageBlockSlideshow       `json:"pageBlockSlideshow,omitempty"`
	PageBlockChatLink        *PageBlockChatLink        `json:"pageBlockChatLink,omitempty"`
	PageBlockTable           *PageBlockTable           `json:"pageBlockTable,omitempty"`
	PageBlockDetails         *PageBlockDetails         `json:"pageBlockDetails,omitempty"`
	PageBlockRelatedArticles *PageBlockRelatedArticles `json:"pageBlockRelatedArticles,omitempty"`
	PageBlockMap             *PageBlockMap             `json:"pageBlockMap,omitempty"`
}

func (t *PageBlock) Type() string {
	return ""
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
	switch typeObj.Type {
	case "pageBlockTitle":
		t.PageBlockTitle = new(PageBlockTitle)
		return json.Unmarshal(b, t.PageBlockTitle)
	case "pageBlockSubtitle":
		t.PageBlockSubtitle = new(PageBlockSubtitle)
		return json.Unmarshal(b, t.PageBlockSubtitle)
	case "pageBlockAuthorDate":
		t.PageBlockAuthorDate = new(PageBlockAuthorDate)
		return json.Unmarshal(b, t.PageBlockAuthorDate)
	case "pageBlockHeader":
		t.PageBlockHeader = new(PageBlockHeader)
		return json.Unmarshal(b, t.PageBlockHeader)
	case "pageBlockSubheader":
		t.PageBlockSubheader = new(PageBlockSubheader)
		return json.Unmarshal(b, t.PageBlockSubheader)
	case "pageBlockKicker":
		t.PageBlockKicker = new(PageBlockKicker)
		return json.Unmarshal(b, t.PageBlockKicker)
	case "pageBlockParagraph":
		t.PageBlockParagraph = new(PageBlockParagraph)
		return json.Unmarshal(b, t.PageBlockParagraph)
	case "pageBlockPreformatted":
		t.PageBlockPreformatted = new(PageBlockPreformatted)
		return json.Unmarshal(b, t.PageBlockPreformatted)
	case "pageBlockFooter":
		t.PageBlockFooter = new(PageBlockFooter)
		return json.Unmarshal(b, t.PageBlockFooter)
	case "pageBlockDivider":
		t.PageBlockDivider = new(PageBlockDivider)
		return json.Unmarshal(b, t.PageBlockDivider)
	case "pageBlockAnchor":
		t.PageBlockAnchor = new(PageBlockAnchor)
		return json.Unmarshal(b, t.PageBlockAnchor)
	case "pageBlockList":
		t.PageBlockList = new(PageBlockList)
		return json.Unmarshal(b, t.PageBlockList)
	case "pageBlockBlockQuote":
		t.PageBlockBlockQuote = new(PageBlockBlockQuote)
		return json.Unmarshal(b, t.PageBlockBlockQuote)
	case "pageBlockPullQuote":
		t.PageBlockPullQuote = new(PageBlockPullQuote)
		return json.Unmarshal(b, t.PageBlockPullQuote)
	case "pageBlockAnimation":
		t.PageBlockAnimation = new(PageBlockAnimation)
		return json.Unmarshal(b, t.PageBlockAnimation)
	case "pageBlockAudio":
		t.PageBlockAudio = new(PageBlockAudio)
		return json.Unmarshal(b, t.PageBlockAudio)
	case "pageBlockPhoto":
		t.PageBlockPhoto = new(PageBlockPhoto)
		return json.Unmarshal(b, t.PageBlockPhoto)
	case "pageBlockVideo":
		t.PageBlockVideo = new(PageBlockVideo)
		return json.Unmarshal(b, t.PageBlockVideo)
	case "pageBlockVoiceNote":
		t.PageBlockVoiceNote = new(PageBlockVoiceNote)
		return json.Unmarshal(b, t.PageBlockVoiceNote)
	case "pageBlockCover":
		t.PageBlockCover = new(PageBlockCover)
		return json.Unmarshal(b, t.PageBlockCover)
	case "pageBlockEmbedded":
		t.PageBlockEmbedded = new(PageBlockEmbedded)
		return json.Unmarshal(b, t.PageBlockEmbedded)
	case "pageBlockEmbeddedPost":
		t.PageBlockEmbeddedPost = new(PageBlockEmbeddedPost)
		return json.Unmarshal(b, t.PageBlockEmbeddedPost)
	case "pageBlockCollage":
		t.PageBlockCollage = new(PageBlockCollage)
		return json.Unmarshal(b, t.PageBlockCollage)
	case "pageBlockSlideshow":
		t.PageBlockSlideshow = new(PageBlockSlideshow)
		return json.Unmarshal(b, t.PageBlockSlideshow)
	case "pageBlockChatLink":
		t.PageBlockChatLink = new(PageBlockChatLink)
		return json.Unmarshal(b, t.PageBlockChatLink)
	case "pageBlockTable":
		t.PageBlockTable = new(PageBlockTable)
		return json.Unmarshal(b, t.PageBlockTable)
	case "pageBlockDetails":
		t.PageBlockDetails = new(PageBlockDetails)
		return json.Unmarshal(b, t.PageBlockDetails)
	case "pageBlockRelatedArticles":
		t.PageBlockRelatedArticles = new(PageBlockRelatedArticles)
		return json.Unmarshal(b, t.PageBlockRelatedArticles)
	case "pageBlockMap":
		t.PageBlockMap = new(PageBlockMap)
		return json.Unmarshal(b, t.PageBlockMap)
	}
	return nil
}

func (t *PageBlock) MarshalJSON() ([]byte, error) {
	if t.PageBlockTitle != nil {
		return json.Marshal(t.PageBlockTitle)
	}
	if t.PageBlockSubtitle != nil {
		return json.Marshal(t.PageBlockSubtitle)
	}
	if t.PageBlockAuthorDate != nil {
		return json.Marshal(t.PageBlockAuthorDate)
	}
	if t.PageBlockHeader != nil {
		return json.Marshal(t.PageBlockHeader)
	}
	if t.PageBlockSubheader != nil {
		return json.Marshal(t.PageBlockSubheader)
	}
	if t.PageBlockKicker != nil {
		return json.Marshal(t.PageBlockKicker)
	}
	if t.PageBlockParagraph != nil {
		return json.Marshal(t.PageBlockParagraph)
	}
	if t.PageBlockPreformatted != nil {
		return json.Marshal(t.PageBlockPreformatted)
	}
	if t.PageBlockFooter != nil {
		return json.Marshal(t.PageBlockFooter)
	}
	if t.PageBlockDivider != nil {
		return json.Marshal(t.PageBlockDivider)
	}
	if t.PageBlockAnchor != nil {
		return json.Marshal(t.PageBlockAnchor)
	}
	if t.PageBlockList != nil {
		return json.Marshal(t.PageBlockList)
	}
	if t.PageBlockBlockQuote != nil {
		return json.Marshal(t.PageBlockBlockQuote)
	}
	if t.PageBlockPullQuote != nil {
		return json.Marshal(t.PageBlockPullQuote)
	}
	if t.PageBlockAnimation != nil {
		return json.Marshal(t.PageBlockAnimation)
	}
	if t.PageBlockAudio != nil {
		return json.Marshal(t.PageBlockAudio)
	}
	if t.PageBlockPhoto != nil {
		return json.Marshal(t.PageBlockPhoto)
	}
	if t.PageBlockVideo != nil {
		return json.Marshal(t.PageBlockVideo)
	}
	if t.PageBlockVoiceNote != nil {
		return json.Marshal(t.PageBlockVoiceNote)
	}
	if t.PageBlockCover != nil {
		return json.Marshal(t.PageBlockCover)
	}
	if t.PageBlockEmbedded != nil {
		return json.Marshal(t.PageBlockEmbedded)
	}
	if t.PageBlockEmbeddedPost != nil {
		return json.Marshal(t.PageBlockEmbeddedPost)
	}
	if t.PageBlockCollage != nil {
		return json.Marshal(t.PageBlockCollage)
	}
	if t.PageBlockSlideshow != nil {
		return json.Marshal(t.PageBlockSlideshow)
	}
	if t.PageBlockChatLink != nil {
		return json.Marshal(t.PageBlockChatLink)
	}
	if t.PageBlockTable != nil {
		return json.Marshal(t.PageBlockTable)
	}
	if t.PageBlockDetails != nil {
		return json.Marshal(t.PageBlockDetails)
	}
	if t.PageBlockRelatedArticles != nil {
		return json.Marshal(t.PageBlockRelatedArticles)
	}
	if t.PageBlockMap != nil {
		return json.Marshal(t.PageBlockMap)
	}
	return nil, fmt.Errorf("no value set for PageBlock")
}

// InlineQueryResultsButtonType Represents type of button in results of inline query
type InlineQueryResultsButtonType struct {
	InlineQueryResultsButtonTypeStartBot *InlineQueryResultsButtonTypeStartBot `json:"inlineQueryResultsButtonTypeStartBot,omitempty"`
	InlineQueryResultsButtonTypeWebApp   *InlineQueryResultsButtonTypeWebApp   `json:"inlineQueryResultsButtonTypeWebApp,omitempty"`
}

func (t *InlineQueryResultsButtonType) Type() string {
	return ""
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

// EmailAddressResetState Describes reset state of an email address
type EmailAddressResetState struct {
	EmailAddressResetStateAvailable *EmailAddressResetStateAvailable `json:"emailAddressResetStateAvailable,omitempty"`
	EmailAddressResetStatePending   *EmailAddressResetStatePending   `json:"emailAddressResetStatePending,omitempty"`
}

func (t *EmailAddressResetState) Type() string {
	return ""
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

// AuthorizationState Represents the current authorization state of the TDLib client
type AuthorizationState struct {
	AuthorizationStateWaitTdlibParameters         *AuthorizationStateWaitTdlibParameters         `json:"authorizationStateWaitTdlibParameters,omitempty"`
	AuthorizationStateWaitPhoneNumber             *AuthorizationStateWaitPhoneNumber             `json:"authorizationStateWaitPhoneNumber,omitempty"`
	AuthorizationStateWaitPremiumPurchase         *AuthorizationStateWaitPremiumPurchase         `json:"authorizationStateWaitPremiumPurchase,omitempty"`
	AuthorizationStateWaitEmailAddress            *AuthorizationStateWaitEmailAddress            `json:"authorizationStateWaitEmailAddress,omitempty"`
	AuthorizationStateWaitEmailCode               *AuthorizationStateWaitEmailCode               `json:"authorizationStateWaitEmailCode,omitempty"`
	AuthorizationStateWaitCode                    *AuthorizationStateWaitCode                    `json:"authorizationStateWaitCode,omitempty"`
	AuthorizationStateWaitOtherDeviceConfirmation *AuthorizationStateWaitOtherDeviceConfirmation `json:"authorizationStateWaitOtherDeviceConfirmation,omitempty"`
	AuthorizationStateWaitRegistration            *AuthorizationStateWaitRegistration            `json:"authorizationStateWaitRegistration,omitempty"`
	AuthorizationStateWaitPassword                *AuthorizationStateWaitPassword                `json:"authorizationStateWaitPassword,omitempty"`
	AuthorizationStateReady                       *AuthorizationStateReady                       `json:"authorizationStateReady,omitempty"`
	AuthorizationStateLoggingOut                  *AuthorizationStateLoggingOut                  `json:"authorizationStateLoggingOut,omitempty"`
	AuthorizationStateClosing                     *AuthorizationStateClosing                     `json:"authorizationStateClosing,omitempty"`
	AuthorizationStateClosed                      *AuthorizationStateClosed                      `json:"authorizationStateClosed,omitempty"`
}

func (t *AuthorizationState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "authorizationStateWaitTdlibParameters":
		t.AuthorizationStateWaitTdlibParameters = new(AuthorizationStateWaitTdlibParameters)
		return json.Unmarshal(b, t.AuthorizationStateWaitTdlibParameters)
	case "authorizationStateWaitPhoneNumber":
		t.AuthorizationStateWaitPhoneNumber = new(AuthorizationStateWaitPhoneNumber)
		return json.Unmarshal(b, t.AuthorizationStateWaitPhoneNumber)
	case "authorizationStateWaitPremiumPurchase":
		t.AuthorizationStateWaitPremiumPurchase = new(AuthorizationStateWaitPremiumPurchase)
		return json.Unmarshal(b, t.AuthorizationStateWaitPremiumPurchase)
	case "authorizationStateWaitEmailAddress":
		t.AuthorizationStateWaitEmailAddress = new(AuthorizationStateWaitEmailAddress)
		return json.Unmarshal(b, t.AuthorizationStateWaitEmailAddress)
	case "authorizationStateWaitEmailCode":
		t.AuthorizationStateWaitEmailCode = new(AuthorizationStateWaitEmailCode)
		return json.Unmarshal(b, t.AuthorizationStateWaitEmailCode)
	case "authorizationStateWaitCode":
		t.AuthorizationStateWaitCode = new(AuthorizationStateWaitCode)
		return json.Unmarshal(b, t.AuthorizationStateWaitCode)
	case "authorizationStateWaitOtherDeviceConfirmation":
		t.AuthorizationStateWaitOtherDeviceConfirmation = new(AuthorizationStateWaitOtherDeviceConfirmation)
		return json.Unmarshal(b, t.AuthorizationStateWaitOtherDeviceConfirmation)
	case "authorizationStateWaitRegistration":
		t.AuthorizationStateWaitRegistration = new(AuthorizationStateWaitRegistration)
		return json.Unmarshal(b, t.AuthorizationStateWaitRegistration)
	case "authorizationStateWaitPassword":
		t.AuthorizationStateWaitPassword = new(AuthorizationStateWaitPassword)
		return json.Unmarshal(b, t.AuthorizationStateWaitPassword)
	case "authorizationStateReady":
		t.AuthorizationStateReady = new(AuthorizationStateReady)
		return json.Unmarshal(b, t.AuthorizationStateReady)
	case "authorizationStateLoggingOut":
		t.AuthorizationStateLoggingOut = new(AuthorizationStateLoggingOut)
		return json.Unmarshal(b, t.AuthorizationStateLoggingOut)
	case "authorizationStateClosing":
		t.AuthorizationStateClosing = new(AuthorizationStateClosing)
		return json.Unmarshal(b, t.AuthorizationStateClosing)
	case "authorizationStateClosed":
		t.AuthorizationStateClosed = new(AuthorizationStateClosed)
		return json.Unmarshal(b, t.AuthorizationStateClosed)
	}
	return nil
}

func (t *AuthorizationState) MarshalJSON() ([]byte, error) {
	if t.AuthorizationStateWaitTdlibParameters != nil {
		return json.Marshal(t.AuthorizationStateWaitTdlibParameters)
	}
	if t.AuthorizationStateWaitPhoneNumber != nil {
		return json.Marshal(t.AuthorizationStateWaitPhoneNumber)
	}
	if t.AuthorizationStateWaitPremiumPurchase != nil {
		return json.Marshal(t.AuthorizationStateWaitPremiumPurchase)
	}
	if t.AuthorizationStateWaitEmailAddress != nil {
		return json.Marshal(t.AuthorizationStateWaitEmailAddress)
	}
	if t.AuthorizationStateWaitEmailCode != nil {
		return json.Marshal(t.AuthorizationStateWaitEmailCode)
	}
	if t.AuthorizationStateWaitCode != nil {
		return json.Marshal(t.AuthorizationStateWaitCode)
	}
	if t.AuthorizationStateWaitOtherDeviceConfirmation != nil {
		return json.Marshal(t.AuthorizationStateWaitOtherDeviceConfirmation)
	}
	if t.AuthorizationStateWaitRegistration != nil {
		return json.Marshal(t.AuthorizationStateWaitRegistration)
	}
	if t.AuthorizationStateWaitPassword != nil {
		return json.Marshal(t.AuthorizationStateWaitPassword)
	}
	if t.AuthorizationStateReady != nil {
		return json.Marshal(t.AuthorizationStateReady)
	}
	if t.AuthorizationStateLoggingOut != nil {
		return json.Marshal(t.AuthorizationStateLoggingOut)
	}
	if t.AuthorizationStateClosing != nil {
		return json.Marshal(t.AuthorizationStateClosing)
	}
	if t.AuthorizationStateClosed != nil {
		return json.Marshal(t.AuthorizationStateClosed)
	}
	return nil, fmt.Errorf("no value set for AuthorizationState")
}

// AffiliateType Describes type of affiliate for an affiliate program
type AffiliateType struct {
	AffiliateTypeCurrentUser *AffiliateTypeCurrentUser `json:"affiliateTypeCurrentUser,omitempty"`
	AffiliateTypeBot         *AffiliateTypeBot         `json:"affiliateTypeBot,omitempty"`
	AffiliateTypeChannel     *AffiliateTypeChannel     `json:"affiliateTypeChannel,omitempty"`
}

func (t *AffiliateType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "affiliateTypeCurrentUser":
		t.AffiliateTypeCurrentUser = new(AffiliateTypeCurrentUser)
		return json.Unmarshal(b, t.AffiliateTypeCurrentUser)
	case "affiliateTypeBot":
		t.AffiliateTypeBot = new(AffiliateTypeBot)
		return json.Unmarshal(b, t.AffiliateTypeBot)
	case "affiliateTypeChannel":
		t.AffiliateTypeChannel = new(AffiliateTypeChannel)
		return json.Unmarshal(b, t.AffiliateTypeChannel)
	}
	return nil
}

func (t *AffiliateType) MarshalJSON() ([]byte, error) {
	if t.AffiliateTypeCurrentUser != nil {
		return json.Marshal(t.AffiliateTypeCurrentUser)
	}
	if t.AffiliateTypeBot != nil {
		return json.Marshal(t.AffiliateTypeBot)
	}
	if t.AffiliateTypeChannel != nil {
		return json.Marshal(t.AffiliateTypeChannel)
	}
	return nil, fmt.Errorf("no value set for AffiliateType")
}

// BotWriteAccessAllowReason Describes a reason why a bot was allowed to write messages to the current user
type BotWriteAccessAllowReason struct {
	BotWriteAccessAllowReasonConnectedWebsite      *BotWriteAccessAllowReasonConnectedWebsite      `json:"botWriteAccessAllowReasonConnectedWebsite,omitempty"`
	BotWriteAccessAllowReasonAddedToAttachmentMenu *BotWriteAccessAllowReasonAddedToAttachmentMenu `json:"botWriteAccessAllowReasonAddedToAttachmentMenu,omitempty"`
	BotWriteAccessAllowReasonLaunchedWebApp        *BotWriteAccessAllowReasonLaunchedWebApp        `json:"botWriteAccessAllowReasonLaunchedWebApp,omitempty"`
	BotWriteAccessAllowReasonAcceptedRequest       *BotWriteAccessAllowReasonAcceptedRequest       `json:"botWriteAccessAllowReasonAcceptedRequest,omitempty"`
}

func (t *BotWriteAccessAllowReason) Type() string {
	return ""
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
	switch typeObj.Type {
	case "botWriteAccessAllowReasonConnectedWebsite":
		t.BotWriteAccessAllowReasonConnectedWebsite = new(BotWriteAccessAllowReasonConnectedWebsite)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonConnectedWebsite)
	case "botWriteAccessAllowReasonAddedToAttachmentMenu":
		t.BotWriteAccessAllowReasonAddedToAttachmentMenu = new(BotWriteAccessAllowReasonAddedToAttachmentMenu)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonAddedToAttachmentMenu)
	case "botWriteAccessAllowReasonLaunchedWebApp":
		t.BotWriteAccessAllowReasonLaunchedWebApp = new(BotWriteAccessAllowReasonLaunchedWebApp)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonLaunchedWebApp)
	case "botWriteAccessAllowReasonAcceptedRequest":
		t.BotWriteAccessAllowReasonAcceptedRequest = new(BotWriteAccessAllowReasonAcceptedRequest)
		return json.Unmarshal(b, t.BotWriteAccessAllowReasonAcceptedRequest)
	}
	return nil
}

func (t *BotWriteAccessAllowReason) MarshalJSON() ([]byte, error) {
	if t.BotWriteAccessAllowReasonConnectedWebsite != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonConnectedWebsite)
	}
	if t.BotWriteAccessAllowReasonAddedToAttachmentMenu != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonAddedToAttachmentMenu)
	}
	if t.BotWriteAccessAllowReasonLaunchedWebApp != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonLaunchedWebApp)
	}
	if t.BotWriteAccessAllowReasonAcceptedRequest != nil {
		return json.Marshal(t.BotWriteAccessAllowReasonAcceptedRequest)
	}
	return nil, fmt.Errorf("no value set for BotWriteAccessAllowReason")
}

// MessageSource Describes source of a message
type MessageSource struct {
	MessageSourceChatHistory                    *MessageSourceChatHistory                    `json:"messageSourceChatHistory,omitempty"`
	MessageSourceMessageThreadHistory           *MessageSourceMessageThreadHistory           `json:"messageSourceMessageThreadHistory,omitempty"`
	MessageSourceForumTopicHistory              *MessageSourceForumTopicHistory              `json:"messageSourceForumTopicHistory,omitempty"`
	MessageSourceDirectMessagesChatTopicHistory *MessageSourceDirectMessagesChatTopicHistory `json:"messageSourceDirectMessagesChatTopicHistory,omitempty"`
	MessageSourceHistoryPreview                 *MessageSourceHistoryPreview                 `json:"messageSourceHistoryPreview,omitempty"`
	MessageSourceChatList                       *MessageSourceChatList                       `json:"messageSourceChatList,omitempty"`
	MessageSourceSearch                         *MessageSourceSearch                         `json:"messageSourceSearch,omitempty"`
	MessageSourceChatEventLog                   *MessageSourceChatEventLog                   `json:"messageSourceChatEventLog,omitempty"`
	MessageSourceNotification                   *MessageSourceNotification                   `json:"messageSourceNotification,omitempty"`
	MessageSourceScreenshot                     *MessageSourceScreenshot                     `json:"messageSourceScreenshot,omitempty"`
	MessageSourceOther                          *MessageSourceOther                          `json:"messageSourceOther,omitempty"`
}

func (t *MessageSource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "messageSourceChatHistory":
		t.MessageSourceChatHistory = new(MessageSourceChatHistory)
		return json.Unmarshal(b, t.MessageSourceChatHistory)
	case "messageSourceMessageThreadHistory":
		t.MessageSourceMessageThreadHistory = new(MessageSourceMessageThreadHistory)
		return json.Unmarshal(b, t.MessageSourceMessageThreadHistory)
	case "messageSourceForumTopicHistory":
		t.MessageSourceForumTopicHistory = new(MessageSourceForumTopicHistory)
		return json.Unmarshal(b, t.MessageSourceForumTopicHistory)
	case "messageSourceDirectMessagesChatTopicHistory":
		t.MessageSourceDirectMessagesChatTopicHistory = new(MessageSourceDirectMessagesChatTopicHistory)
		return json.Unmarshal(b, t.MessageSourceDirectMessagesChatTopicHistory)
	case "messageSourceHistoryPreview":
		t.MessageSourceHistoryPreview = new(MessageSourceHistoryPreview)
		return json.Unmarshal(b, t.MessageSourceHistoryPreview)
	case "messageSourceChatList":
		t.MessageSourceChatList = new(MessageSourceChatList)
		return json.Unmarshal(b, t.MessageSourceChatList)
	case "messageSourceSearch":
		t.MessageSourceSearch = new(MessageSourceSearch)
		return json.Unmarshal(b, t.MessageSourceSearch)
	case "messageSourceChatEventLog":
		t.MessageSourceChatEventLog = new(MessageSourceChatEventLog)
		return json.Unmarshal(b, t.MessageSourceChatEventLog)
	case "messageSourceNotification":
		t.MessageSourceNotification = new(MessageSourceNotification)
		return json.Unmarshal(b, t.MessageSourceNotification)
	case "messageSourceScreenshot":
		t.MessageSourceScreenshot = new(MessageSourceScreenshot)
		return json.Unmarshal(b, t.MessageSourceScreenshot)
	case "messageSourceOther":
		t.MessageSourceOther = new(MessageSourceOther)
		return json.Unmarshal(b, t.MessageSourceOther)
	}
	return nil
}

func (t *MessageSource) MarshalJSON() ([]byte, error) {
	if t.MessageSourceChatHistory != nil {
		return json.Marshal(t.MessageSourceChatHistory)
	}
	if t.MessageSourceMessageThreadHistory != nil {
		return json.Marshal(t.MessageSourceMessageThreadHistory)
	}
	if t.MessageSourceForumTopicHistory != nil {
		return json.Marshal(t.MessageSourceForumTopicHistory)
	}
	if t.MessageSourceDirectMessagesChatTopicHistory != nil {
		return json.Marshal(t.MessageSourceDirectMessagesChatTopicHistory)
	}
	if t.MessageSourceHistoryPreview != nil {
		return json.Marshal(t.MessageSourceHistoryPreview)
	}
	if t.MessageSourceChatList != nil {
		return json.Marshal(t.MessageSourceChatList)
	}
	if t.MessageSourceSearch != nil {
		return json.Marshal(t.MessageSourceSearch)
	}
	if t.MessageSourceChatEventLog != nil {
		return json.Marshal(t.MessageSourceChatEventLog)
	}
	if t.MessageSourceNotification != nil {
		return json.Marshal(t.MessageSourceNotification)
	}
	if t.MessageSourceScreenshot != nil {
		return json.Marshal(t.MessageSourceScreenshot)
	}
	if t.MessageSourceOther != nil {
		return json.Marshal(t.MessageSourceOther)
	}
	return nil, fmt.Errorf("no value set for MessageSource")
}

// SentGift Represents content of a gift received by a user or a channel chat
type SentGift struct {
	SentGiftRegular  *SentGiftRegular  `json:"sentGiftRegular,omitempty"`
	SentGiftUpgraded *SentGiftUpgraded `json:"sentGiftUpgraded,omitempty"`
}

func (t *SentGift) Type() string {
	return ""
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

// TextEntityType Represents a part of the text which must be formatted differently
type TextEntityType struct {
	TextEntityTypeMention              *TextEntityTypeMention              `json:"textEntityTypeMention,omitempty"`
	TextEntityTypeHashtag              *TextEntityTypeHashtag              `json:"textEntityTypeHashtag,omitempty"`
	TextEntityTypeCashtag              *TextEntityTypeCashtag              `json:"textEntityTypeCashtag,omitempty"`
	TextEntityTypeBotCommand           *TextEntityTypeBotCommand           `json:"textEntityTypeBotCommand,omitempty"`
	TextEntityTypeUrl                  *TextEntityTypeUrl                  `json:"textEntityTypeUrl,omitempty"`
	TextEntityTypeEmailAddress         *TextEntityTypeEmailAddress         `json:"textEntityTypeEmailAddress,omitempty"`
	TextEntityTypePhoneNumber          *TextEntityTypePhoneNumber          `json:"textEntityTypePhoneNumber,omitempty"`
	TextEntityTypeBankCardNumber       *TextEntityTypeBankCardNumber       `json:"textEntityTypeBankCardNumber,omitempty"`
	TextEntityTypeBold                 *TextEntityTypeBold                 `json:"textEntityTypeBold,omitempty"`
	TextEntityTypeItalic               *TextEntityTypeItalic               `json:"textEntityTypeItalic,omitempty"`
	TextEntityTypeUnderline            *TextEntityTypeUnderline            `json:"textEntityTypeUnderline,omitempty"`
	TextEntityTypeStrikethrough        *TextEntityTypeStrikethrough        `json:"textEntityTypeStrikethrough,omitempty"`
	TextEntityTypeSpoiler              *TextEntityTypeSpoiler              `json:"textEntityTypeSpoiler,omitempty"`
	TextEntityTypeCode                 *TextEntityTypeCode                 `json:"textEntityTypeCode,omitempty"`
	TextEntityTypePre                  *TextEntityTypePre                  `json:"textEntityTypePre,omitempty"`
	TextEntityTypePreCode              *TextEntityTypePreCode              `json:"textEntityTypePreCode,omitempty"`
	TextEntityTypeBlockQuote           *TextEntityTypeBlockQuote           `json:"textEntityTypeBlockQuote,omitempty"`
	TextEntityTypeExpandableBlockQuote *TextEntityTypeExpandableBlockQuote `json:"textEntityTypeExpandableBlockQuote,omitempty"`
	TextEntityTypeTextUrl              *TextEntityTypeTextUrl              `json:"textEntityTypeTextUrl,omitempty"`
	TextEntityTypeMentionName          *TextEntityTypeMentionName          `json:"textEntityTypeMentionName,omitempty"`
	TextEntityTypeCustomEmoji          *TextEntityTypeCustomEmoji          `json:"textEntityTypeCustomEmoji,omitempty"`
	TextEntityTypeMediaTimestamp       *TextEntityTypeMediaTimestamp       `json:"textEntityTypeMediaTimestamp,omitempty"`
}

func (t *TextEntityType) Type() string {
	return ""
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
	switch typeObj.Type {
	case "textEntityTypeMention":
		t.TextEntityTypeMention = new(TextEntityTypeMention)
		return json.Unmarshal(b, t.TextEntityTypeMention)
	case "textEntityTypeHashtag":
		t.TextEntityTypeHashtag = new(TextEntityTypeHashtag)
		return json.Unmarshal(b, t.TextEntityTypeHashtag)
	case "textEntityTypeCashtag":
		t.TextEntityTypeCashtag = new(TextEntityTypeCashtag)
		return json.Unmarshal(b, t.TextEntityTypeCashtag)
	case "textEntityTypeBotCommand":
		t.TextEntityTypeBotCommand = new(TextEntityTypeBotCommand)
		return json.Unmarshal(b, t.TextEntityTypeBotCommand)
	case "textEntityTypeUrl":
		t.TextEntityTypeUrl = new(TextEntityTypeUrl)
		return json.Unmarshal(b, t.TextEntityTypeUrl)
	case "textEntityTypeEmailAddress":
		t.TextEntityTypeEmailAddress = new(TextEntityTypeEmailAddress)
		return json.Unmarshal(b, t.TextEntityTypeEmailAddress)
	case "textEntityTypePhoneNumber":
		t.TextEntityTypePhoneNumber = new(TextEntityTypePhoneNumber)
		return json.Unmarshal(b, t.TextEntityTypePhoneNumber)
	case "textEntityTypeBankCardNumber":
		t.TextEntityTypeBankCardNumber = new(TextEntityTypeBankCardNumber)
		return json.Unmarshal(b, t.TextEntityTypeBankCardNumber)
	case "textEntityTypeBold":
		t.TextEntityTypeBold = new(TextEntityTypeBold)
		return json.Unmarshal(b, t.TextEntityTypeBold)
	case "textEntityTypeItalic":
		t.TextEntityTypeItalic = new(TextEntityTypeItalic)
		return json.Unmarshal(b, t.TextEntityTypeItalic)
	case "textEntityTypeUnderline":
		t.TextEntityTypeUnderline = new(TextEntityTypeUnderline)
		return json.Unmarshal(b, t.TextEntityTypeUnderline)
	case "textEntityTypeStrikethrough":
		t.TextEntityTypeStrikethrough = new(TextEntityTypeStrikethrough)
		return json.Unmarshal(b, t.TextEntityTypeStrikethrough)
	case "textEntityTypeSpoiler":
		t.TextEntityTypeSpoiler = new(TextEntityTypeSpoiler)
		return json.Unmarshal(b, t.TextEntityTypeSpoiler)
	case "textEntityTypeCode":
		t.TextEntityTypeCode = new(TextEntityTypeCode)
		return json.Unmarshal(b, t.TextEntityTypeCode)
	case "textEntityTypePre":
		t.TextEntityTypePre = new(TextEntityTypePre)
		return json.Unmarshal(b, t.TextEntityTypePre)
	case "textEntityTypePreCode":
		t.TextEntityTypePreCode = new(TextEntityTypePreCode)
		return json.Unmarshal(b, t.TextEntityTypePreCode)
	case "textEntityTypeBlockQuote":
		t.TextEntityTypeBlockQuote = new(TextEntityTypeBlockQuote)
		return json.Unmarshal(b, t.TextEntityTypeBlockQuote)
	case "textEntityTypeExpandableBlockQuote":
		t.TextEntityTypeExpandableBlockQuote = new(TextEntityTypeExpandableBlockQuote)
		return json.Unmarshal(b, t.TextEntityTypeExpandableBlockQuote)
	case "textEntityTypeTextUrl":
		t.TextEntityTypeTextUrl = new(TextEntityTypeTextUrl)
		return json.Unmarshal(b, t.TextEntityTypeTextUrl)
	case "textEntityTypeMentionName":
		t.TextEntityTypeMentionName = new(TextEntityTypeMentionName)
		return json.Unmarshal(b, t.TextEntityTypeMentionName)
	case "textEntityTypeCustomEmoji":
		t.TextEntityTypeCustomEmoji = new(TextEntityTypeCustomEmoji)
		return json.Unmarshal(b, t.TextEntityTypeCustomEmoji)
	case "textEntityTypeMediaTimestamp":
		t.TextEntityTypeMediaTimestamp = new(TextEntityTypeMediaTimestamp)
		return json.Unmarshal(b, t.TextEntityTypeMediaTimestamp)
	}
	return nil
}

func (t *TextEntityType) MarshalJSON() ([]byte, error) {
	if t.TextEntityTypeMention != nil {
		return json.Marshal(t.TextEntityTypeMention)
	}
	if t.TextEntityTypeHashtag != nil {
		return json.Marshal(t.TextEntityTypeHashtag)
	}
	if t.TextEntityTypeCashtag != nil {
		return json.Marshal(t.TextEntityTypeCashtag)
	}
	if t.TextEntityTypeBotCommand != nil {
		return json.Marshal(t.TextEntityTypeBotCommand)
	}
	if t.TextEntityTypeUrl != nil {
		return json.Marshal(t.TextEntityTypeUrl)
	}
	if t.TextEntityTypeEmailAddress != nil {
		return json.Marshal(t.TextEntityTypeEmailAddress)
	}
	if t.TextEntityTypePhoneNumber != nil {
		return json.Marshal(t.TextEntityTypePhoneNumber)
	}
	if t.TextEntityTypeBankCardNumber != nil {
		return json.Marshal(t.TextEntityTypeBankCardNumber)
	}
	if t.TextEntityTypeBold != nil {
		return json.Marshal(t.TextEntityTypeBold)
	}
	if t.TextEntityTypeItalic != nil {
		return json.Marshal(t.TextEntityTypeItalic)
	}
	if t.TextEntityTypeUnderline != nil {
		return json.Marshal(t.TextEntityTypeUnderline)
	}
	if t.TextEntityTypeStrikethrough != nil {
		return json.Marshal(t.TextEntityTypeStrikethrough)
	}
	if t.TextEntityTypeSpoiler != nil {
		return json.Marshal(t.TextEntityTypeSpoiler)
	}
	if t.TextEntityTypeCode != nil {
		return json.Marshal(t.TextEntityTypeCode)
	}
	if t.TextEntityTypePre != nil {
		return json.Marshal(t.TextEntityTypePre)
	}
	if t.TextEntityTypePreCode != nil {
		return json.Marshal(t.TextEntityTypePreCode)
	}
	if t.TextEntityTypeBlockQuote != nil {
		return json.Marshal(t.TextEntityTypeBlockQuote)
	}
	if t.TextEntityTypeExpandableBlockQuote != nil {
		return json.Marshal(t.TextEntityTypeExpandableBlockQuote)
	}
	if t.TextEntityTypeTextUrl != nil {
		return json.Marshal(t.TextEntityTypeTextUrl)
	}
	if t.TextEntityTypeMentionName != nil {
		return json.Marshal(t.TextEntityTypeMentionName)
	}
	if t.TextEntityTypeCustomEmoji != nil {
		return json.Marshal(t.TextEntityTypeCustomEmoji)
	}
	if t.TextEntityTypeMediaTimestamp != nil {
		return json.Marshal(t.TextEntityTypeMediaTimestamp)
	}
	return nil, fmt.Errorf("no value set for TextEntityType")
}

// UserPrivacySettingRule Represents a single rule for managing user privacy settings
type UserPrivacySettingRule struct {
	UserPrivacySettingRuleAllowAll            *UserPrivacySettingRuleAllowAll            `json:"userPrivacySettingRuleAllowAll,omitempty"`
	UserPrivacySettingRuleAllowContacts       *UserPrivacySettingRuleAllowContacts       `json:"userPrivacySettingRuleAllowContacts,omitempty"`
	UserPrivacySettingRuleAllowBots           *UserPrivacySettingRuleAllowBots           `json:"userPrivacySettingRuleAllowBots,omitempty"`
	UserPrivacySettingRuleAllowPremiumUsers   *UserPrivacySettingRuleAllowPremiumUsers   `json:"userPrivacySettingRuleAllowPremiumUsers,omitempty"`
	UserPrivacySettingRuleAllowUsers          *UserPrivacySettingRuleAllowUsers          `json:"userPrivacySettingRuleAllowUsers,omitempty"`
	UserPrivacySettingRuleAllowChatMembers    *UserPrivacySettingRuleAllowChatMembers    `json:"userPrivacySettingRuleAllowChatMembers,omitempty"`
	UserPrivacySettingRuleRestrictAll         *UserPrivacySettingRuleRestrictAll         `json:"userPrivacySettingRuleRestrictAll,omitempty"`
	UserPrivacySettingRuleRestrictContacts    *UserPrivacySettingRuleRestrictContacts    `json:"userPrivacySettingRuleRestrictContacts,omitempty"`
	UserPrivacySettingRuleRestrictBots        *UserPrivacySettingRuleRestrictBots        `json:"userPrivacySettingRuleRestrictBots,omitempty"`
	UserPrivacySettingRuleRestrictUsers       *UserPrivacySettingRuleRestrictUsers       `json:"userPrivacySettingRuleRestrictUsers,omitempty"`
	UserPrivacySettingRuleRestrictChatMembers *UserPrivacySettingRuleRestrictChatMembers `json:"userPrivacySettingRuleRestrictChatMembers,omitempty"`
}

func (t *UserPrivacySettingRule) Type() string {
	return ""
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
	switch typeObj.Type {
	case "userPrivacySettingRuleAllowAll":
		t.UserPrivacySettingRuleAllowAll = new(UserPrivacySettingRuleAllowAll)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowAll)
	case "userPrivacySettingRuleAllowContacts":
		t.UserPrivacySettingRuleAllowContacts = new(UserPrivacySettingRuleAllowContacts)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowContacts)
	case "userPrivacySettingRuleAllowBots":
		t.UserPrivacySettingRuleAllowBots = new(UserPrivacySettingRuleAllowBots)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowBots)
	case "userPrivacySettingRuleAllowPremiumUsers":
		t.UserPrivacySettingRuleAllowPremiumUsers = new(UserPrivacySettingRuleAllowPremiumUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowPremiumUsers)
	case "userPrivacySettingRuleAllowUsers":
		t.UserPrivacySettingRuleAllowUsers = new(UserPrivacySettingRuleAllowUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowUsers)
	case "userPrivacySettingRuleAllowChatMembers":
		t.UserPrivacySettingRuleAllowChatMembers = new(UserPrivacySettingRuleAllowChatMembers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleAllowChatMembers)
	case "userPrivacySettingRuleRestrictAll":
		t.UserPrivacySettingRuleRestrictAll = new(UserPrivacySettingRuleRestrictAll)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictAll)
	case "userPrivacySettingRuleRestrictContacts":
		t.UserPrivacySettingRuleRestrictContacts = new(UserPrivacySettingRuleRestrictContacts)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictContacts)
	case "userPrivacySettingRuleRestrictBots":
		t.UserPrivacySettingRuleRestrictBots = new(UserPrivacySettingRuleRestrictBots)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictBots)
	case "userPrivacySettingRuleRestrictUsers":
		t.UserPrivacySettingRuleRestrictUsers = new(UserPrivacySettingRuleRestrictUsers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictUsers)
	case "userPrivacySettingRuleRestrictChatMembers":
		t.UserPrivacySettingRuleRestrictChatMembers = new(UserPrivacySettingRuleRestrictChatMembers)
		return json.Unmarshal(b, t.UserPrivacySettingRuleRestrictChatMembers)
	}
	return nil
}

func (t *UserPrivacySettingRule) MarshalJSON() ([]byte, error) {
	if t.UserPrivacySettingRuleAllowAll != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowAll)
	}
	if t.UserPrivacySettingRuleAllowContacts != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowContacts)
	}
	if t.UserPrivacySettingRuleAllowBots != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowBots)
	}
	if t.UserPrivacySettingRuleAllowPremiumUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowPremiumUsers)
	}
	if t.UserPrivacySettingRuleAllowUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowUsers)
	}
	if t.UserPrivacySettingRuleAllowChatMembers != nil {
		return json.Marshal(t.UserPrivacySettingRuleAllowChatMembers)
	}
	if t.UserPrivacySettingRuleRestrictAll != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictAll)
	}
	if t.UserPrivacySettingRuleRestrictContacts != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictContacts)
	}
	if t.UserPrivacySettingRuleRestrictBots != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictBots)
	}
	if t.UserPrivacySettingRuleRestrictUsers != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictUsers)
	}
	if t.UserPrivacySettingRuleRestrictChatMembers != nil {
		return json.Marshal(t.UserPrivacySettingRuleRestrictChatMembers)
	}
	return nil, fmt.Errorf("no value set for UserPrivacySettingRule")
}

// TextParseMode Describes the way the text needs to be parsed for text entities
type TextParseMode struct {
	TextParseModeMarkdown *TextParseModeMarkdown `json:"textParseModeMarkdown,omitempty"`
	TextParseModeHTML     *TextParseModeHTML     `json:"textParseModeHTML,omitempty"`
}

func (t *TextParseMode) Type() string {
	return ""
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
	switch typeObj.Type {
	case "textParseModeMarkdown":
		t.TextParseModeMarkdown = new(TextParseModeMarkdown)
		return json.Unmarshal(b, t.TextParseModeMarkdown)
	case "textParseModeHTML":
		t.TextParseModeHTML = new(TextParseModeHTML)
		return json.Unmarshal(b, t.TextParseModeHTML)
	}
	return nil
}

func (t *TextParseMode) MarshalJSON() ([]byte, error) {
	if t.TextParseModeMarkdown != nil {
		return json.Marshal(t.TextParseModeMarkdown)
	}
	if t.TextParseModeHTML != nil {
		return json.Marshal(t.TextParseModeHTML)
	}
	return nil, fmt.Errorf("no value set for TextParseMode")
}

// InputFile Points to a file
type InputFile struct {
	InputFileId        *InputFileId        `json:"inputFileId,omitempty"`
	InputFileRemote    *InputFileRemote    `json:"inputFileRemote,omitempty"`
	InputFileLocal     *InputFileLocal     `json:"inputFileLocal,omitempty"`
	InputFileGenerated *InputFileGenerated `json:"inputFileGenerated,omitempty"`
}

func (t *InputFile) Type() string {
	return ""
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
	switch typeObj.Type {
	case "inputFileId":
		t.InputFileId = new(InputFileId)
		return json.Unmarshal(b, t.InputFileId)
	case "inputFileRemote":
		t.InputFileRemote = new(InputFileRemote)
		return json.Unmarshal(b, t.InputFileRemote)
	case "inputFileLocal":
		t.InputFileLocal = new(InputFileLocal)
		return json.Unmarshal(b, t.InputFileLocal)
	case "inputFileGenerated":
		t.InputFileGenerated = new(InputFileGenerated)
		return json.Unmarshal(b, t.InputFileGenerated)
	}
	return nil
}

func (t *InputFile) MarshalJSON() ([]byte, error) {
	if t.InputFileId != nil {
		return json.Marshal(t.InputFileId)
	}
	if t.InputFileRemote != nil {
		return json.Marshal(t.InputFileRemote)
	}
	if t.InputFileLocal != nil {
		return json.Marshal(t.InputFileLocal)
	}
	if t.InputFileGenerated != nil {
		return json.Marshal(t.InputFileGenerated)
	}
	return nil, fmt.Errorf("no value set for InputFile")
}

// SearchMessagesFilter Represents a filter for message search results
type SearchMessagesFilter struct {
	SearchMessagesFilterEmpty             *SearchMessagesFilterEmpty             `json:"searchMessagesFilterEmpty,omitempty"`
	SearchMessagesFilterAnimation         *SearchMessagesFilterAnimation         `json:"searchMessagesFilterAnimation,omitempty"`
	SearchMessagesFilterAudio             *SearchMessagesFilterAudio             `json:"searchMessagesFilterAudio,omitempty"`
	SearchMessagesFilterDocument          *SearchMessagesFilterDocument          `json:"searchMessagesFilterDocument,omitempty"`
	SearchMessagesFilterPhoto             *SearchMessagesFilterPhoto             `json:"searchMessagesFilterPhoto,omitempty"`
	SearchMessagesFilterVideo             *SearchMessagesFilterVideo             `json:"searchMessagesFilterVideo,omitempty"`
	SearchMessagesFilterVoiceNote         *SearchMessagesFilterVoiceNote         `json:"searchMessagesFilterVoiceNote,omitempty"`
	SearchMessagesFilterPhotoAndVideo     *SearchMessagesFilterPhotoAndVideo     `json:"searchMessagesFilterPhotoAndVideo,omitempty"`
	SearchMessagesFilterUrl               *SearchMessagesFilterUrl               `json:"searchMessagesFilterUrl,omitempty"`
	SearchMessagesFilterChatPhoto         *SearchMessagesFilterChatPhoto         `json:"searchMessagesFilterChatPhoto,omitempty"`
	SearchMessagesFilterVideoNote         *SearchMessagesFilterVideoNote         `json:"searchMessagesFilterVideoNote,omitempty"`
	SearchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote `json:"searchMessagesFilterVoiceAndVideoNote,omitempty"`
	SearchMessagesFilterMention           *SearchMessagesFilterMention           `json:"searchMessagesFilterMention,omitempty"`
	SearchMessagesFilterUnreadMention     *SearchMessagesFilterUnreadMention     `json:"searchMessagesFilterUnreadMention,omitempty"`
	SearchMessagesFilterUnreadReaction    *SearchMessagesFilterUnreadReaction    `json:"searchMessagesFilterUnreadReaction,omitempty"`
	SearchMessagesFilterFailedToSend      *SearchMessagesFilterFailedToSend      `json:"searchMessagesFilterFailedToSend,omitempty"`
	SearchMessagesFilterPinned            *SearchMessagesFilterPinned            `json:"searchMessagesFilterPinned,omitempty"`
}

func (t *SearchMessagesFilter) Type() string {
	return ""
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
	switch typeObj.Type {
	case "searchMessagesFilterEmpty":
		t.SearchMessagesFilterEmpty = new(SearchMessagesFilterEmpty)
		return json.Unmarshal(b, t.SearchMessagesFilterEmpty)
	case "searchMessagesFilterAnimation":
		t.SearchMessagesFilterAnimation = new(SearchMessagesFilterAnimation)
		return json.Unmarshal(b, t.SearchMessagesFilterAnimation)
	case "searchMessagesFilterAudio":
		t.SearchMessagesFilterAudio = new(SearchMessagesFilterAudio)
		return json.Unmarshal(b, t.SearchMessagesFilterAudio)
	case "searchMessagesFilterDocument":
		t.SearchMessagesFilterDocument = new(SearchMessagesFilterDocument)
		return json.Unmarshal(b, t.SearchMessagesFilterDocument)
	case "searchMessagesFilterPhoto":
		t.SearchMessagesFilterPhoto = new(SearchMessagesFilterPhoto)
		return json.Unmarshal(b, t.SearchMessagesFilterPhoto)
	case "searchMessagesFilterVideo":
		t.SearchMessagesFilterVideo = new(SearchMessagesFilterVideo)
		return json.Unmarshal(b, t.SearchMessagesFilterVideo)
	case "searchMessagesFilterVoiceNote":
		t.SearchMessagesFilterVoiceNote = new(SearchMessagesFilterVoiceNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVoiceNote)
	case "searchMessagesFilterPhotoAndVideo":
		t.SearchMessagesFilterPhotoAndVideo = new(SearchMessagesFilterPhotoAndVideo)
		return json.Unmarshal(b, t.SearchMessagesFilterPhotoAndVideo)
	case "searchMessagesFilterUrl":
		t.SearchMessagesFilterUrl = new(SearchMessagesFilterUrl)
		return json.Unmarshal(b, t.SearchMessagesFilterUrl)
	case "searchMessagesFilterChatPhoto":
		t.SearchMessagesFilterChatPhoto = new(SearchMessagesFilterChatPhoto)
		return json.Unmarshal(b, t.SearchMessagesFilterChatPhoto)
	case "searchMessagesFilterVideoNote":
		t.SearchMessagesFilterVideoNote = new(SearchMessagesFilterVideoNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVideoNote)
	case "searchMessagesFilterVoiceAndVideoNote":
		t.SearchMessagesFilterVoiceAndVideoNote = new(SearchMessagesFilterVoiceAndVideoNote)
		return json.Unmarshal(b, t.SearchMessagesFilterVoiceAndVideoNote)
	case "searchMessagesFilterMention":
		t.SearchMessagesFilterMention = new(SearchMessagesFilterMention)
		return json.Unmarshal(b, t.SearchMessagesFilterMention)
	case "searchMessagesFilterUnreadMention":
		t.SearchMessagesFilterUnreadMention = new(SearchMessagesFilterUnreadMention)
		return json.Unmarshal(b, t.SearchMessagesFilterUnreadMention)
	case "searchMessagesFilterUnreadReaction":
		t.SearchMessagesFilterUnreadReaction = new(SearchMessagesFilterUnreadReaction)
		return json.Unmarshal(b, t.SearchMessagesFilterUnreadReaction)
	case "searchMessagesFilterFailedToSend":
		t.SearchMessagesFilterFailedToSend = new(SearchMessagesFilterFailedToSend)
		return json.Unmarshal(b, t.SearchMessagesFilterFailedToSend)
	case "searchMessagesFilterPinned":
		t.SearchMessagesFilterPinned = new(SearchMessagesFilterPinned)
		return json.Unmarshal(b, t.SearchMessagesFilterPinned)
	}
	return nil
}

func (t *SearchMessagesFilter) MarshalJSON() ([]byte, error) {
	if t.SearchMessagesFilterEmpty != nil {
		return json.Marshal(t.SearchMessagesFilterEmpty)
	}
	if t.SearchMessagesFilterAnimation != nil {
		return json.Marshal(t.SearchMessagesFilterAnimation)
	}
	if t.SearchMessagesFilterAudio != nil {
		return json.Marshal(t.SearchMessagesFilterAudio)
	}
	if t.SearchMessagesFilterDocument != nil {
		return json.Marshal(t.SearchMessagesFilterDocument)
	}
	if t.SearchMessagesFilterPhoto != nil {
		return json.Marshal(t.SearchMessagesFilterPhoto)
	}
	if t.SearchMessagesFilterVideo != nil {
		return json.Marshal(t.SearchMessagesFilterVideo)
	}
	if t.SearchMessagesFilterVoiceNote != nil {
		return json.Marshal(t.SearchMessagesFilterVoiceNote)
	}
	if t.SearchMessagesFilterPhotoAndVideo != nil {
		return json.Marshal(t.SearchMessagesFilterPhotoAndVideo)
	}
	if t.SearchMessagesFilterUrl != nil {
		return json.Marshal(t.SearchMessagesFilterUrl)
	}
	if t.SearchMessagesFilterChatPhoto != nil {
		return json.Marshal(t.SearchMessagesFilterChatPhoto)
	}
	if t.SearchMessagesFilterVideoNote != nil {
		return json.Marshal(t.SearchMessagesFilterVideoNote)
	}
	if t.SearchMessagesFilterVoiceAndVideoNote != nil {
		return json.Marshal(t.SearchMessagesFilterVoiceAndVideoNote)
	}
	if t.SearchMessagesFilterMention != nil {
		return json.Marshal(t.SearchMessagesFilterMention)
	}
	if t.SearchMessagesFilterUnreadMention != nil {
		return json.Marshal(t.SearchMessagesFilterUnreadMention)
	}
	if t.SearchMessagesFilterUnreadReaction != nil {
		return json.Marshal(t.SearchMessagesFilterUnreadReaction)
	}
	if t.SearchMessagesFilterFailedToSend != nil {
		return json.Marshal(t.SearchMessagesFilterFailedToSend)
	}
	if t.SearchMessagesFilterPinned != nil {
		return json.Marshal(t.SearchMessagesFilterPinned)
	}
	return nil, fmt.Errorf("no value set for SearchMessagesFilter")
}

// DiceStickers Contains animated stickers which must be used for dice animation rendering
type DiceStickers struct {
	DiceStickersRegular     *DiceStickersRegular     `json:"diceStickersRegular,omitempty"`
	DiceStickersSlotMachine *DiceStickersSlotMachine `json:"diceStickersSlotMachine,omitempty"`
}

func (t *DiceStickers) Type() string {
	return ""
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

// ReportReason Describes the reason why a chat is reported
type ReportReason struct {
	ReportReasonSpam              *ReportReasonSpam              `json:"reportReasonSpam,omitempty"`
	ReportReasonViolence          *ReportReasonViolence          `json:"reportReasonViolence,omitempty"`
	ReportReasonPornography       *ReportReasonPornography       `json:"reportReasonPornography,omitempty"`
	ReportReasonChildAbuse        *ReportReasonChildAbuse        `json:"reportReasonChildAbuse,omitempty"`
	ReportReasonCopyright         *ReportReasonCopyright         `json:"reportReasonCopyright,omitempty"`
	ReportReasonUnrelatedLocation *ReportReasonUnrelatedLocation `json:"reportReasonUnrelatedLocation,omitempty"`
	ReportReasonFake              *ReportReasonFake              `json:"reportReasonFake,omitempty"`
	ReportReasonIllegalDrugs      *ReportReasonIllegalDrugs      `json:"reportReasonIllegalDrugs,omitempty"`
	ReportReasonPersonalDetails   *ReportReasonPersonalDetails   `json:"reportReasonPersonalDetails,omitempty"`
	ReportReasonCustom            *ReportReasonCustom            `json:"reportReasonCustom,omitempty"`
}

func (t *ReportReason) Type() string {
	return ""
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
	switch typeObj.Type {
	case "reportReasonSpam":
		t.ReportReasonSpam = new(ReportReasonSpam)
		return json.Unmarshal(b, t.ReportReasonSpam)
	case "reportReasonViolence":
		t.ReportReasonViolence = new(ReportReasonViolence)
		return json.Unmarshal(b, t.ReportReasonViolence)
	case "reportReasonPornography":
		t.ReportReasonPornography = new(ReportReasonPornography)
		return json.Unmarshal(b, t.ReportReasonPornography)
	case "reportReasonChildAbuse":
		t.ReportReasonChildAbuse = new(ReportReasonChildAbuse)
		return json.Unmarshal(b, t.ReportReasonChildAbuse)
	case "reportReasonCopyright":
		t.ReportReasonCopyright = new(ReportReasonCopyright)
		return json.Unmarshal(b, t.ReportReasonCopyright)
	case "reportReasonUnrelatedLocation":
		t.ReportReasonUnrelatedLocation = new(ReportReasonUnrelatedLocation)
		return json.Unmarshal(b, t.ReportReasonUnrelatedLocation)
	case "reportReasonFake":
		t.ReportReasonFake = new(ReportReasonFake)
		return json.Unmarshal(b, t.ReportReasonFake)
	case "reportReasonIllegalDrugs":
		t.ReportReasonIllegalDrugs = new(ReportReasonIllegalDrugs)
		return json.Unmarshal(b, t.ReportReasonIllegalDrugs)
	case "reportReasonPersonalDetails":
		t.ReportReasonPersonalDetails = new(ReportReasonPersonalDetails)
		return json.Unmarshal(b, t.ReportReasonPersonalDetails)
	case "reportReasonCustom":
		t.ReportReasonCustom = new(ReportReasonCustom)
		return json.Unmarshal(b, t.ReportReasonCustom)
	}
	return nil
}

func (t *ReportReason) MarshalJSON() ([]byte, error) {
	if t.ReportReasonSpam != nil {
		return json.Marshal(t.ReportReasonSpam)
	}
	if t.ReportReasonViolence != nil {
		return json.Marshal(t.ReportReasonViolence)
	}
	if t.ReportReasonPornography != nil {
		return json.Marshal(t.ReportReasonPornography)
	}
	if t.ReportReasonChildAbuse != nil {
		return json.Marshal(t.ReportReasonChildAbuse)
	}
	if t.ReportReasonCopyright != nil {
		return json.Marshal(t.ReportReasonCopyright)
	}
	if t.ReportReasonUnrelatedLocation != nil {
		return json.Marshal(t.ReportReasonUnrelatedLocation)
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
	if t.ReportReasonCustom != nil {
		return json.Marshal(t.ReportReasonCustom)
	}
	return nil, fmt.Errorf("no value set for ReportReason")
}

// SuggestedPostState Describes state of a suggested post
type SuggestedPostState struct {
	SuggestedPostStatePending  *SuggestedPostStatePending  `json:"suggestedPostStatePending,omitempty"`
	SuggestedPostStateApproved *SuggestedPostStateApproved `json:"suggestedPostStateApproved,omitempty"`
	SuggestedPostStateDeclined *SuggestedPostStateDeclined `json:"suggestedPostStateDeclined,omitempty"`
}

func (t *SuggestedPostState) Type() string {
	return ""
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
	switch typeObj.Type {
	case "suggestedPostStatePending":
		t.SuggestedPostStatePending = new(SuggestedPostStatePending)
		return json.Unmarshal(b, t.SuggestedPostStatePending)
	case "suggestedPostStateApproved":
		t.SuggestedPostStateApproved = new(SuggestedPostStateApproved)
		return json.Unmarshal(b, t.SuggestedPostStateApproved)
	case "suggestedPostStateDeclined":
		t.SuggestedPostStateDeclined = new(SuggestedPostStateDeclined)
		return json.Unmarshal(b, t.SuggestedPostStateDeclined)
	}
	return nil
}

func (t *SuggestedPostState) MarshalJSON() ([]byte, error) {
	if t.SuggestedPostStatePending != nil {
		return json.Marshal(t.SuggestedPostStatePending)
	}
	if t.SuggestedPostStateApproved != nil {
		return json.Marshal(t.SuggestedPostStateApproved)
	}
	if t.SuggestedPostStateDeclined != nil {
		return json.Marshal(t.SuggestedPostStateDeclined)
	}
	return nil, fmt.Errorf("no value set for SuggestedPostState")
}

// GiftResaleResult Describes result of sending a resold gift
type GiftResaleResult struct {
	GiftResaleResultOk             *GiftResaleResultOk             `json:"giftResaleResultOk,omitempty"`
	GiftResaleResultPriceIncreased *GiftResaleResultPriceIncreased `json:"giftResaleResultPriceIncreased,omitempty"`
}

func (t *GiftResaleResult) Type() string {
	return ""
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

// UserStatus Describes the last time the user was online
type UserStatus struct {
	UserStatusEmpty     *UserStatusEmpty     `json:"userStatusEmpty,omitempty"`
	UserStatusOnline    *UserStatusOnline    `json:"userStatusOnline,omitempty"`
	UserStatusOffline   *UserStatusOffline   `json:"userStatusOffline,omitempty"`
	UserStatusRecently  *UserStatusRecently  `json:"userStatusRecently,omitempty"`
	UserStatusLastWeek  *UserStatusLastWeek  `json:"userStatusLastWeek,omitempty"`
	UserStatusLastMonth *UserStatusLastMonth `json:"userStatusLastMonth,omitempty"`
}

func (t *UserStatus) Type() string {
	return ""
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
	switch typeObj.Type {
	case "userStatusEmpty":
		t.UserStatusEmpty = new(UserStatusEmpty)
		return json.Unmarshal(b, t.UserStatusEmpty)
	case "userStatusOnline":
		t.UserStatusOnline = new(UserStatusOnline)
		return json.Unmarshal(b, t.UserStatusOnline)
	case "userStatusOffline":
		t.UserStatusOffline = new(UserStatusOffline)
		return json.Unmarshal(b, t.UserStatusOffline)
	case "userStatusRecently":
		t.UserStatusRecently = new(UserStatusRecently)
		return json.Unmarshal(b, t.UserStatusRecently)
	case "userStatusLastWeek":
		t.UserStatusLastWeek = new(UserStatusLastWeek)
		return json.Unmarshal(b, t.UserStatusLastWeek)
	case "userStatusLastMonth":
		t.UserStatusLastMonth = new(UserStatusLastMonth)
		return json.Unmarshal(b, t.UserStatusLastMonth)
	}
	return nil
}

func (t *UserStatus) MarshalJSON() ([]byte, error) {
	if t.UserStatusEmpty != nil {
		return json.Marshal(t.UserStatusEmpty)
	}
	if t.UserStatusOnline != nil {
		return json.Marshal(t.UserStatusOnline)
	}
	if t.UserStatusOffline != nil {
		return json.Marshal(t.UserStatusOffline)
	}
	if t.UserStatusRecently != nil {
		return json.Marshal(t.UserStatusRecently)
	}
	if t.UserStatusLastWeek != nil {
		return json.Marshal(t.UserStatusLastWeek)
	}
	if t.UserStatusLastMonth != nil {
		return json.Marshal(t.UserStatusLastMonth)
	}
	return nil, fmt.Errorf("no value set for UserStatus")
}

// EmojiCategorySource Describes source of stickers for an emoji category
type EmojiCategorySource struct {
	EmojiCategorySourceSearch  *EmojiCategorySourceSearch  `json:"emojiCategorySourceSearch,omitempty"`
	EmojiCategorySourcePremium *EmojiCategorySourcePremium `json:"emojiCategorySourcePremium,omitempty"`
}

func (t *EmojiCategorySource) Type() string {
	return ""
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
	switch typeObj.Type {
	case "emojiCategorySourceSearch":
		t.EmojiCategorySourceSearch = new(EmojiCategorySourceSearch)
		return json.Unmarshal(b, t.EmojiCategorySourceSearch)
	case "emojiCategorySourcePremium":
		t.EmojiCategorySourcePremium = new(EmojiCategorySourcePremium)
		return json.Unmarshal(b, t.EmojiCategorySourcePremium)
	}
	return nil
}

func (t *EmojiCategorySource) MarshalJSON() ([]byte, error) {
	if t.EmojiCategorySourceSearch != nil {
		return json.Marshal(t.EmojiCategorySourceSearch)
	}
	if t.EmojiCategorySourcePremium != nil {
		return json.Marshal(t.EmojiCategorySourcePremium)
	}
	return nil, fmt.Errorf("no value set for EmojiCategorySource")
}

// GroupCallVideoQuality Describes the quality of a group call video
type GroupCallVideoQuality struct {
	GroupCallVideoQualityThumbnail *GroupCallVideoQualityThumbnail `json:"groupCallVideoQualityThumbnail,omitempty"`
	GroupCallVideoQualityMedium    *GroupCallVideoQualityMedium    `json:"groupCallVideoQualityMedium,omitempty"`
	GroupCallVideoQualityFull      *GroupCallVideoQualityFull      `json:"groupCallVideoQualityFull,omitempty"`
}

func (t *GroupCallVideoQuality) Type() string {
	return ""
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
	switch typeObj.Type {
	case "groupCallVideoQualityThumbnail":
		t.GroupCallVideoQualityThumbnail = new(GroupCallVideoQualityThumbnail)
		return json.Unmarshal(b, t.GroupCallVideoQualityThumbnail)
	case "groupCallVideoQualityMedium":
		t.GroupCallVideoQualityMedium = new(GroupCallVideoQualityMedium)
		return json.Unmarshal(b, t.GroupCallVideoQualityMedium)
	case "groupCallVideoQualityFull":
		t.GroupCallVideoQualityFull = new(GroupCallVideoQualityFull)
		return json.Unmarshal(b, t.GroupCallVideoQualityFull)
	}
	return nil
}

func (t *GroupCallVideoQuality) MarshalJSON() ([]byte, error) {
	if t.GroupCallVideoQualityThumbnail != nil {
		return json.Marshal(t.GroupCallVideoQualityThumbnail)
	}
	if t.GroupCallVideoQualityMedium != nil {
		return json.Marshal(t.GroupCallVideoQualityMedium)
	}
	if t.GroupCallVideoQualityFull != nil {
		return json.Marshal(t.GroupCallVideoQualityFull)
	}
	return nil, fmt.Errorf("no value set for GroupCallVideoQuality")
}
