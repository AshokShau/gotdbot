package webapp

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	TelegramPublicKey     = "e7bf03a2fa4602af4580703d88dda5bb59f32ed8b02a56c187fe7d34caed242d"
	TelegramTestPublicKey = "40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec"
)

var (
	ErrInvalidInitData  = errors.New("invalid Telegram init data")
	ErrInvalidHash      = errors.New("invalid Telegram init data hash")
	ErrInvalidSignature = errors.New("invalid Telegram init data signature")
	ErrAuthDate         = errors.New("missing or invalid Telegram init data auth_date")
	ErrExpired          = errors.New("Telegram init data expired")
	ErrFutureAuthDate   = errors.New("Telegram init data auth_date is in the future")
	ErrBotTokenRequired = errors.New("bot token required")
	ErrInvalidBotID     = errors.New("bot ID must be positive")
)

type ValidationOptions struct {
	// MaxAge is the age limit; zero disables expiration checks.
	MaxAge time.Duration
	// ClockSkew allows this much clock difference when checking auth_date.
	ClockSkew time.Duration
	// Now uses time.Now when zero.
	Now             time.Time
	TestEnvironment bool
}

// ParseTelegramInitData decodes init data without verifying its authenticity.
func ParseTelegramInitData(raw string) (*WebAppInitData, error) {
	values, err := parseInitDataValues(raw)
	if err != nil {
		return nil, err
	}
	return decodeInitData(raw, values)
}

// ValidateTelegramInitData verifies the bot-token HMAC before decoding the data.
// All fields except hash, including signature, participate in the HMAC.
func ValidateTelegramInitData(raw, botToken string, opts *ValidationOptions) (*WebAppInitData, error) {
	if botToken == "" {
		return nil, ErrBotTokenRequired
	}
	values, err := parseInitDataValues(raw)
	if err != nil {
		return nil, err
	}
	received, err := hex.DecodeString(values.Get("hash"))
	if err != nil || len(received) != sha256.Size {
		return nil, ErrInvalidHash
	}
	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(botToken))
	mac := hmac.New(sha256.New, secret.Sum(nil))
	mac.Write([]byte(dataCheckString(values, false)))
	if !hmac.Equal(received, mac.Sum(nil)) {
		return nil, ErrInvalidHash
	}
	return validatedInitData(raw, values, opts)
}

// ValidateTelegramInitDataSignature verifies Telegram's Ed25519 signature using
// the production key, or the test key when TestEnvironment is true.
func ValidateTelegramInitDataSignature(raw string, botID int64, opts *ValidationOptions) (*WebAppInitData, error) {
	key := TelegramPublicKey
	if opts != nil && opts.TestEnvironment {
		key = TelegramTestPublicKey
	}
	publicKey, _ := hex.DecodeString(key)
	return validateInitDataSignature(raw, botID, ed25519.PublicKey(publicKey), opts)
}

func validateInitDataSignature(raw string, botID int64, publicKey ed25519.PublicKey, opts *ValidationOptions) (*WebAppInitData, error) {
	if botID <= 0 {
		return nil, ErrInvalidBotID
	}
	values, err := parseInitDataValues(raw)
	if err != nil {
		return nil, err
	}
	signature := values.Get("signature")
	encoding := base64.RawURLEncoding.Strict()
	if strings.HasSuffix(signature, "=") {
		encoding = base64.URLEncoding.Strict()
	}
	decoded, err := encoding.DecodeString(signature)
	if err != nil || len(decoded) != ed25519.SignatureSize {
		return nil, ErrInvalidSignature
	}
	check := strconv.FormatInt(botID, 10) + ":WebAppData\n" + dataCheckString(values, true)
	if len(publicKey) != ed25519.PublicKeySize || !ed25519.Verify(publicKey, []byte(check), decoded) {
		return nil, ErrInvalidSignature
	}
	return validatedInitData(raw, values, opts)
}

func parseInitDataValues(raw string) (url.Values, error) {
	values, err := url.ParseQuery(raw)
	if err != nil {
		return nil, fmt.Errorf("%w: malformed query string", ErrInvalidInitData)
	}
	for key, entries := range values {
		if key == "" || !utf8.ValidString(key) || strings.ContainsAny(key, "=\r\n") || len(entries) != 1 {
			return nil, fmt.Errorf("%w: invalid or duplicate field", ErrInvalidInitData)
		}
		if !utf8.ValidString(entries[0]) || strings.ContainsAny(entries[0], "\r\n") {
			return nil, fmt.Errorf("%w: invalid field value", ErrInvalidInitData)
		}
	}
	return values, nil
}

func dataCheckString(values url.Values, excludeSignature bool) string {
	keys := make([]string, 0, len(values))
	for key := range values {
		if key != "hash" && (!excludeSignature || key != "signature") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	fields := make([]string, 0, len(keys))
	for _, key := range keys {
		fields = append(fields, key+"="+values.Get(key))
	}
	return strings.Join(fields, "\n")
}

func validatedInitData(raw string, values url.Values, opts *ValidationOptions) (*WebAppInitData, error) {
	var options ValidationOptions
	if opts != nil {
		options = *opts
	}
	if options.MaxAge < 0 || options.ClockSkew < 0 {
		return nil, fmt.Errorf("%w: negative validation duration", ErrInvalidInitData)
	}
	authDate, err := strconv.ParseInt(values.Get("auth_date"), 10, 64)
	if err != nil || authDate <= 0 {
		return nil, ErrAuthDate
	}
	now := options.Now
	if now.IsZero() {
		now = time.Now()
	}
	issued := time.Unix(authDate, 0)
	if authDate > now.Add(options.ClockSkew).Unix() {
		return nil, ErrFutureAuthDate
	}
	if options.MaxAge > 0 && issued.Before(now.Add(-options.MaxAge).Add(-options.ClockSkew)) {
		return nil, ErrExpired
	}
	return decodeInitData(raw, values)
}

func decodeInitData(raw string, values url.Values) (*WebAppInitData, error) {
	data := &WebAppInitData{
		QueryID: values.Get("query_id"), ChatJoinRequestQueryID: values.Get("chat_join_request_query_id"),
		ChatType: ChatType(values.Get("chat_type")), ChatInstance: values.Get("chat_instance"),
		StartParam: values.Get("start_param"), Hash: values.Get("hash"), Signature: values.Get("signature"),
		Raw: raw, Values: values,
	}
	for _, object := range []struct {
		key string
		dst any
	}{{"user", &data.User}, {"receiver", &data.Receiver}, {"chat", &data.Chat}} {
		if entries, exists := values[object.key]; exists {
			if err := json.Unmarshal([]byte(entries[0]), object.dst); err != nil {
				return nil, fmt.Errorf("%w: invalid %s JSON", ErrInvalidInitData, object.key)
			}
		}
	}
	if entries, exists := values["auth_date"]; exists {
		var err error
		data.AuthDate, err = strconv.ParseInt(entries[0], 10, 64)
		if err != nil || data.AuthDate <= 0 {
			return nil, ErrAuthDate
		}
	}
	if entries, exists := values["can_send_after"]; exists {
		seconds, err := strconv.ParseInt(entries[0], 10, 64)
		if err != nil || seconds < 0 {
			return nil, fmt.Errorf("%w: invalid can_send_after", ErrInvalidInitData)
		}
		data.CanSendAfter = seconds
	}
	return data, nil
}

func (d *WebAppInitData) AuthTime() time.Time {
	if d == nil || d.AuthDate <= 0 {
		return time.Time{}
	}
	return time.Unix(d.AuthDate, 0)
}

func (d *WebAppInitData) CanSendAt() time.Time {
	issued := d.AuthTime()
	if issued.IsZero() || d.CanSendAfter < 0 || d.AuthDate > (1<<63-1)-d.CanSendAfter {
		return time.Time{}
	}
	return time.Unix(d.AuthDate+d.CanSendAfter, 0)
}
