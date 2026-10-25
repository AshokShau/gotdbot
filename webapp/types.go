package webapp

import (
	"encoding/json"
	"net/url"
)

type WebAppInitData struct {
	QueryID                string      `json:"query_id,omitempty"`
	ChatJoinRequestQueryID string      `json:"chat_join_request_query_id,omitempty"`
	User                   *WebAppUser `json:"user,omitempty"`
	Receiver               *WebAppUser `json:"receiver,omitempty"`
	Chat                   *WebAppChat `json:"chat,omitempty"`
	ChatType               ChatType    `json:"chat_type,omitempty"`
	ChatInstance           string      `json:"chat_instance,omitempty"`
	StartParam             string      `json:"start_param,omitempty"`
	CanSendAfter           int64       `json:"can_send_after,omitempty"`
	AuthDate               int64       `json:"auth_date"`
	Hash                   string      `json:"hash"`
	Signature              string      `json:"signature,omitempty"`
	Raw                    string      `json:"-"`
	Values                 url.Values  `json:"-"`
}

type WebAppUser struct {
	ID                    int64  `json:"id"`
	IsBot                 bool   `json:"is_bot,omitempty"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name,omitempty"`
	Username              string `json:"username,omitempty"`
	LanguageCode          string `json:"language_code,omitempty"`
	IsPremium             bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu,omitempty"`
	AllowsWriteToPM       bool   `json:"allows_write_to_pm,omitempty"`
	PhotoURL              string `json:"photo_url,omitempty"`
}

type WebAppChat struct {
	ID       int64    `json:"id"`
	Type     ChatType `json:"type"`
	Title    string   `json:"title"`
	Username string   `json:"username,omitempty"`
	PhotoURL string   `json:"photo_url,omitempty"`
}

type ChatType string

const (
	ChatTypeSender     ChatType = "sender"
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type ThemeParams struct {
	BGColor                string `json:"bg_color,omitempty"`
	TextColor              string `json:"text_color,omitempty"`
	HintColor              string `json:"hint_color,omitempty"`
	LinkColor              string `json:"link_color,omitempty"`
	ButtonColor            string `json:"button_color,omitempty"`
	ButtonTextColor        string `json:"button_text_color,omitempty"`
	SecondaryBGColor       string `json:"secondary_bg_color,omitempty"`
	HeaderBGColor          string `json:"header_bg_color,omitempty"`
	BottomBarBGColor       string `json:"bottom_bar_bg_color,omitempty"`
	AccentTextColor        string `json:"accent_text_color,omitempty"`
	SectionBGColor         string `json:"section_bg_color,omitempty"`
	SectionHeaderTextColor string `json:"section_header_text_color,omitempty"`
	SectionSeparatorColor  string `json:"section_separator_color,omitempty"`
	SubtitleTextColor      string `json:"subtitle_text_color,omitempty"`
	DestructiveTextColor   string `json:"destructive_text_color,omitempty"`
}

type StoryShareParams struct {
	Text       string           `json:"text,omitempty"`
	WidgetLink *StoryWidgetLink `json:"widget_link,omitempty"`
}

type StoryWidgetLink struct {
	URL  string `json:"url"`
	Name string `json:"name,omitempty"`
}

type ScanQrPopupParams struct {
	Text string `json:"text,omitempty"`
}

type PopupParams struct {
	Title   string        `json:"title,omitempty"`
	Message string        `json:"message"`
	Buttons []PopupButton `json:"buttons,omitempty"`
}

type PopupButton struct {
	ID   string          `json:"id,omitempty"`
	Type PopupButtonType `json:"type,omitempty"`
	Text string          `json:"text,omitempty"`
}

type PopupButtonType string

const (
	PopupButtonDefault     PopupButtonType = "default"
	PopupButtonOK          PopupButtonType = "ok"
	PopupButtonClose       PopupButtonType = "close"
	PopupButtonCancel      PopupButtonType = "cancel"
	PopupButtonDestructive PopupButtonType = "destructive"
)

type EmojiStatusParams struct {
	Duration int64 `json:"duration,omitempty"`
}

type DownloadFileParams struct {
	URL      string `json:"url"`
	FileName string `json:"file_name"`
}

type SafeAreaInset struct {
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
}

type ContentSafeAreaInset SafeAreaInset

type BiometricRequestAccessParams struct {
	Reason string `json:"reason,omitempty"`
}

type BiometricAuthenticateParams struct {
	Reason string `json:"reason,omitempty"`
}

type AccelerometerStartParams struct {
	RefreshRate int `json:"refresh_rate,omitempty"`
}

type DeviceOrientationStartParams struct {
	RefreshRate  int  `json:"refresh_rate,omitempty"`
	NeedAbsolute bool `json:"need_absolute,omitempty"`
}

type GyroscopeStartParams struct {
	RefreshRate int `json:"refresh_rate,omitempty"`
}

type LocationData struct {
	Latitude           float64  `json:"latitude"`
	Longitude          float64  `json:"longitude"`
	Altitude           *float64 `json:"altitude"`
	Course             *float64 `json:"course"`
	Speed              *float64 `json:"speed"`
	HorizontalAccuracy *float64 `json:"horizontal_accuracy"`
	VerticalAccuracy   *float64 `json:"vertical_accuracy"`
	CourseAccuracy     *float64 `json:"course_accuracy"`
	SpeedAccuracy      *float64 `json:"speed_accuracy"`
}

type OpenLinkOptions struct {
	TryInstantView bool `json:"try_instant_view,omitempty"`
}

type BottomButtonParams struct {
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
	Text              *string `json:"text,omitempty"`
	Color             *string `json:"color,omitempty"`
	TextColor         *string `json:"text_color,omitempty"`
	HasShineEffect    *bool   `json:"has_shine_effect,omitempty"`
	Position          *string `json:"position,omitempty"`
	IsActive          *bool   `json:"is_active,omitempty"`
	IsVisible         *bool   `json:"is_visible,omitempty"`
}

type ServerlessError struct {
	Message    string          `json:"message"`
	Type       string          `json:"type,omitempty"`
	Status     int             `json:"status"`
	Parameters json.RawMessage `json:"parameters,omitempty"`
}

func (e *ServerlessError) Error() string { return e.Message }
