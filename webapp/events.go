package webapp

type EventType string

const (
	EventActivated                  EventType = "activated"
	EventDeactivated                EventType = "deactivated"
	EventThemeChanged               EventType = "themeChanged"
	EventViewportChanged            EventType = "viewportChanged"
	EventSafeAreaChanged            EventType = "safeAreaChanged"
	EventContentSafeAreaChanged     EventType = "contentSafeAreaChanged"
	EventMainButtonClicked          EventType = "mainButtonClicked"
	EventSecondaryButtonClicked     EventType = "secondaryButtonClicked"
	EventBackButtonClicked          EventType = "backButtonClicked"
	EventSettingsButtonClicked      EventType = "settingsButtonClicked"
	EventInvoiceClosed              EventType = "invoiceClosed"
	EventPopupClosed                EventType = "popupClosed"
	EventQrTextReceived             EventType = "qrTextReceived"
	EventScanQrPopupClosed          EventType = "scanQrPopupClosed"
	EventClipboardTextReceived      EventType = "clipboardTextReceived"
	EventWriteAccessRequested       EventType = "writeAccessRequested"
	EventContactRequested           EventType = "contactRequested"
	EventBiometricManagerUpdated    EventType = "biometricManagerUpdated"
	EventBiometricAuthRequested     EventType = "biometricAuthRequested"
	EventBiometricTokenUpdated      EventType = "biometricTokenUpdated"
	EventFullscreenChanged          EventType = "fullscreenChanged"
	EventFullscreenFailed           EventType = "fullscreenFailed"
	EventHomeScreenAdded            EventType = "homeScreenAdded"
	EventHomeScreenChecked          EventType = "homeScreenChecked"
	EventAccelerometerStarted       EventType = "accelerometerStarted"
	EventAccelerometerStopped       EventType = "accelerometerStopped"
	EventAccelerometerChanged       EventType = "accelerometerChanged"
	EventAccelerometerFailed        EventType = "accelerometerFailed"
	EventDeviceOrientationStarted   EventType = "deviceOrientationStarted"
	EventDeviceOrientationStopped   EventType = "deviceOrientationStopped"
	EventDeviceOrientationChanged   EventType = "deviceOrientationChanged"
	EventDeviceOrientationFailed    EventType = "deviceOrientationFailed"
	EventGyroscopeStarted           EventType = "gyroscopeStarted"
	EventGyroscopeStopped           EventType = "gyroscopeStopped"
	EventGyroscopeChanged           EventType = "gyroscopeChanged"
	EventGyroscopeFailed            EventType = "gyroscopeFailed"
	EventLocationManagerUpdated     EventType = "locationManagerUpdated"
	EventLocationRequested          EventType = "locationRequested"
	EventShareMessageSent           EventType = "shareMessageSent"
	EventShareMessageFailed         EventType = "shareMessageFailed"
	EventEmojiStatusSet             EventType = "emojiStatusSet"
	EventEmojiStatusFailed          EventType = "emojiStatusFailed"
	EventEmojiStatusAccessRequested EventType = "emojiStatusAccessRequested"
	EventFileDownloadRequested      EventType = "fileDownloadRequested"
)

type InvoiceStatus string

const (
	InvoicePaid      InvoiceStatus = "paid"
	InvoiceCancelled InvoiceStatus = "cancelled"
	InvoiceFailed    InvoiceStatus = "failed"
	InvoicePending   InvoiceStatus = "pending"
)

type PermissionStatus string

const (
	PermissionAllowed   PermissionStatus = "allowed"
	PermissionCancelled PermissionStatus = "cancelled"
)

type HomeScreenStatus string

const (
	HomeScreenUnsupported HomeScreenStatus = "unsupported"
	HomeScreenUnknown     HomeScreenStatus = "unknown"
	HomeScreenAdded       HomeScreenStatus = "added"
	HomeScreenMissed      HomeScreenStatus = "missed"
)

const (
	ErrorUnsupported           = "UNSUPPORTED"
	ErrorAlreadyFullscreen     = "ALREADY_FULLSCREEN"
	ErrorMessageExpired        = "MESSAGE_EXPIRED"
	ErrorMessageSendFailed     = "MESSAGE_SEND_FAILED"
	ErrorUserDeclined          = "USER_DECLINED"
	ErrorUnknown               = "UNKNOWN_ERROR"
	ErrorSuggestedEmojiInvalid = "SUGGESTED_EMOJI_INVALID"
	ErrorDurationInvalid       = "DURATION_INVALID"
	ErrorServer                = "SERVER_ERROR"
	ContactSent                = "sent"
	ContactCancelled           = "cancelled"
	FileDownloading            = "downloading"
	FileDownloadCancelled      = "cancelled"
)

type ViewportChangedEvent struct {
	IsStateStable bool `json:"isStateStable"`
}

type InvoiceClosedEvent struct {
	URL    string        `json:"url"`
	Status InvoiceStatus `json:"status"`
}

type PopupClosedEvent struct {
	ButtonID *string `json:"button_id"`
}

type QrTextReceivedEvent struct {
	Data string `json:"data"`
}

type ClipboardTextReceivedEvent struct {
	Data *string `json:"data"`
}

type WriteAccessRequestedEvent struct {
	Status PermissionStatus `json:"status"`
}

type ContactRequestedEvent struct {
	Status string `json:"status"`
}

type BiometricAuthRequestedEvent struct {
	IsAuthenticated bool   `json:"isAuthenticated"`
	BiometricToken  string `json:"biometricToken,omitempty"`
}

type BiometricTokenUpdatedEvent struct {
	IsUpdated bool `json:"isUpdated"`
}

type ErrorEvent struct {
	Error string `json:"error"`
}

type FullscreenFailedEvent = ErrorEvent
type AccelerometerFailedEvent = ErrorEvent
type DeviceOrientationFailedEvent = ErrorEvent
type GyroscopeFailedEvent = ErrorEvent
type ShareMessageFailedEvent = ErrorEvent
type EmojiStatusFailedEvent = ErrorEvent

type HomeScreenCheckedEvent struct {
	Status HomeScreenStatus `json:"status"`
}

type LocationRequestedEvent struct {
	LocationData *LocationData `json:"locationData"`
}

type EmojiStatusAccessRequestedEvent = WriteAccessRequestedEvent

type FileDownloadRequestedEvent struct {
	Status string `json:"status"`
}
