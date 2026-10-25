package webapp

// WebApp describes serializable client state; client methods run in Telegram's JavaScript SDK.
type WebApp struct {
	InitData                     string               `json:"initData"`
	InitDataUnsafe               *WebAppInitData      `json:"initDataUnsafe"`
	Version                      string               `json:"version"`
	Platform                     string               `json:"platform"`
	ColorScheme                  string               `json:"colorScheme"`
	ThemeParams                  ThemeParams          `json:"themeParams"`
	IsActive                     bool                 `json:"isActive"`
	IsExpanded                   bool                 `json:"isExpanded"`
	ViewportHeight               float64              `json:"viewportHeight"`
	ViewportStableHeight         float64              `json:"viewportStableHeight"`
	HeaderColor                  string               `json:"headerColor"`
	BackgroundColor              string               `json:"backgroundColor"`
	BottomBarColor               string               `json:"bottomBarColor"`
	IsClosingConfirmationEnabled bool                 `json:"isClosingConfirmationEnabled"`
	IsVerticalSwipesEnabled      bool                 `json:"isVerticalSwipesEnabled"`
	IsFullscreen                 bool                 `json:"isFullscreen"`
	IsOrientationLocked          bool                 `json:"isOrientationLocked"`
	SafeAreaInset                SafeAreaInset        `json:"safeAreaInset"`
	ContentSafeAreaInset         ContentSafeAreaInset `json:"contentSafeAreaInset"`
	BackButton                   *BackButton          `json:"BackButton,omitempty"`
	MainButton                   *BottomButton        `json:"MainButton,omitempty"`
	SecondaryButton              *BottomButton        `json:"SecondaryButton,omitempty"`
	SettingsButton               *SettingsButton      `json:"SettingsButton,omitempty"`
	BiometricManager             *BiometricManager    `json:"BiometricManager,omitempty"`
	Accelerometer                *Accelerometer       `json:"Accelerometer,omitempty"`
	DeviceOrientation            *DeviceOrientation   `json:"DeviceOrientation,omitempty"`
	Gyroscope                    *Gyroscope           `json:"Gyroscope,omitempty"`
	LocationManager              *LocationManager     `json:"LocationManager,omitempty"`
}

type BackButton struct {
	IsVisible bool `json:"isVisible"`
}

type SettingsButton struct {
	IsVisible bool `json:"isVisible"`
}

type BottomButton struct {
	Type              string `json:"type"`
	IconCustomEmojiID string `json:"iconCustomEmojiId,omitempty"`
	Text              string `json:"text"`
	Color             string `json:"color"`
	TextColor         string `json:"textColor"`
	IsVisible         bool   `json:"isVisible"`
	IsActive          bool   `json:"isActive"`
	HasShineEffect    bool   `json:"hasShineEffect"`
	Position          string `json:"position,omitempty"`
	IsProgressVisible bool   `json:"isProgressVisible"`
}

type BiometricManager struct {
	IsInited              bool   `json:"isInited"`
	IsBiometricAvailable  bool   `json:"isBiometricAvailable"`
	BiometricType         string `json:"biometricType"`
	IsAccessRequested     bool   `json:"isAccessRequested"`
	IsAccessGranted       bool   `json:"isAccessGranted"`
	IsBiometricTokenSaved bool   `json:"isBiometricTokenSaved"`
	DeviceID              string `json:"deviceId"`
}

type Accelerometer struct {
	IsStarted bool    `json:"isStarted"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type DeviceOrientation struct {
	IsStarted bool    `json:"isStarted"`
	Absolute  bool    `json:"absolute"`
	Alpha     float64 `json:"alpha"`
	Beta      float64 `json:"beta"`
	Gamma     float64 `json:"gamma"`
}

type Gyroscope struct {
	IsStarted bool    `json:"isStarted"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type LocationManager struct {
	IsInited            bool `json:"isInited"`
	IsLocationAvailable bool `json:"isLocationAvailable"`
	IsAccessRequested   bool `json:"isAccessRequested"`
	IsAccessGranted     bool `json:"isAccessGranted"`
}
