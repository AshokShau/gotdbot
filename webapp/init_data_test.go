package webapp

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strings"
	"testing"
	"time"
)

const fixtureToken = "123456:TEST_token"
const fixtureInitData = `auth_date=1700000000&query_id=AAE%2B%2F_%3D%3F&chat_join_request_query_id=9223372036854775807&user=%7B%22id%22%3A4503599627370495%2C%22first_name%22%3A%22A%2BB+%26+C%22%2C%22last_name%22%3A%22%5Cu00e9%22%2C%22username%22%3A%22tester%22%2C%22language_code%22%3A%22en%22%2C%22is_premium%22%3Atrue%2C%22added_to_attachment_menu%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Fexample.org%5C%2Fa%2Bb.svg%22%7D&receiver=%7B%22id%22%3A42%2C%22first_name%22%3A%22Bot%22%2C%22is_bot%22%3Atrue%7D&chat=%7B%22id%22%3A-1001234567890%2C%22type%22%3A%22supergroup%22%2C%22title%22%3A%22Test+%26+chat%22%2C%22username%22%3A%22testchat%22%2C%22photo_url%22%3A%22https%3A%2F%2Fexample.org%2Fchat.svg%22%7D&chat_type=supergroup&chat_instance=18446744073709551615&start_param=plus%2B+percent%25+equals%3D+ampersand%26&can_send_after=5&signature=AA_-&future_field=keep+%2B+%2526+%26+%3D&hash=8d94da2f6cc95403aa3000f0cbdcfc9dfc5acd6e01b9615f2e3baedc1da8c0de`

func TestValidateTelegramInitDataFixture(t *testing.T) {
	values, err := url.ParseQuery(fixtureInitData)
	if err != nil {
		t.Fatal(err)
	}
	for _, raw := range []string{fixtureInitData, values.Encode(), strings.ReplaceAll(values.Encode(), "+", "%20")} {
		data, err := ValidateTelegramInitData(raw, fixtureToken, &ValidationOptions{Now: time.Unix(1700000010, 0), MaxAge: time.Minute})
		if err != nil {
			t.Fatal(err)
		}
		if data.Raw != raw || data.QueryID != "AAE+/_=?" || data.ChatJoinRequestQueryID != "9223372036854775807" || data.ChatInstance != "18446744073709551615" {
			t.Fatalf("lost session data: %+v", data)
		}
		if data.User == nil || data.User.ID != 4503599627370495 || data.User.FirstName != "A+B & C" || data.User.LastName != "é" || !data.User.IsPremium || !data.User.AddedToAttachmentMenu || !data.User.AllowsWriteToPM || data.User.PhotoURL != "https://example.org/a+b.svg" {
			t.Fatalf("lost user data: %+v", data.User)
		}
		if data.Receiver == nil || !data.Receiver.IsBot || data.Receiver.ID != 42 || data.Chat == nil || data.Chat.ID != -1001234567890 || data.Chat.Type != ChatTypeSupergroup {
			t.Fatalf("lost chat/receiver: %+v", data)
		}
		if data.StartParam != "plus+ percent% equals= ampersand&" || data.Values.Get("future_field") != "keep + %26 & =" || data.Signature != "AA_-" {
			t.Fatalf("re-encoded signed values: %+v", data)
		}
		if data.AuthTime().Unix() != 1700000000 || data.CanSendAt().Unix() != 1700000005 {
			t.Fatal("wrong auth/send time")
		}
		encoded, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(encoded), "future_field") || strings.Contains(string(encoded), "Values") || strings.Contains(string(encoded), "Raw") {
			t.Fatalf("raw verification data leaked into JSON: %s", encoded)
		}
	}
}

func TestInitDataRejectsTampering(t *testing.T) {
	values, _ := url.ParseQuery(fixtureInitData)
	for _, key := range []string{"user", "receiver", "chat", "signature", "future_field", "auth_date", "start_param", "chat_join_request_query_id"} {
		t.Run(key, func(t *testing.T) {
			modified, _ := url.ParseQuery(values.Encode())
			modified.Set(key, modified.Get(key)+"x")
			if data, err := ValidateTelegramInitData(modified.Encode(), fixtureToken, nil); data != nil || !errors.Is(err, ErrInvalidHash) {
				t.Fatalf("tampering accepted: %v, %v", data, err)
			}
		})
	}
	for _, raw := range []string{
		fixtureInitData + "&user=%7B%7D", fixtureInitData + "&%61uth_date=1700000000", fixtureInitData + "&hash=x",
		fixtureInitData + "&future_field=second", fixtureInitData + "&=empty", fixtureInitData + "&a%3Db=c",
		fixtureInitData + "&x=%0Aa=b", fixtureInitData + "&x=%0D", fixtureInitData + "&x=%FF",
		fixtureInitData + "&%FF=x", fixtureInitData + "&bad=%zz", fixtureInitData + "&x=a;b",
	} {
		if data, err := ValidateTelegramInitData(raw, fixtureToken, nil); data != nil || !errors.Is(err, ErrInvalidInitData) {
			t.Fatalf("malformed data accepted: %v, %v", data, err)
		}
	}
	if _, err := ValidateTelegramInitData(fixtureInitData, "wrong-token", nil); !errors.Is(err, ErrInvalidHash) {
		t.Fatalf("wrong token: %v", err)
	}
	if _, err := ValidateTelegramInitData(fixtureInitData, "", nil); !errors.Is(err, ErrBotTokenRequired) {
		t.Fatalf("empty token: %v", err)
	}
	for _, hash := range []string{"", "xx", strings.Repeat("00", 31), strings.Repeat("00", 33)} {
		modified, _ := url.ParseQuery(values.Encode())
		modified.Set("hash", hash)
		if _, err := ValidateTelegramInitData(modified.Encode(), fixtureToken, nil); !errors.Is(err, ErrInvalidHash) {
			t.Fatalf("bad hash %q: %v", hash, err)
		}
	}
}

func signedQuery(fields url.Values) string {
	var entries []string
	for key, values := range fields {
		if key != "hash" {
			entries = append(entries, key+"="+values[0])
		}
	}
	sort.Strings(entries)
	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(fixtureToken))
	mac := hmac.New(sha256.New, secret.Sum(nil))
	mac.Write([]byte(strings.Join(entries, "\n")))
	fields.Set("hash", hex.EncodeToString(mac.Sum(nil)))
	return fields.Encode()
}

func TestInitDataDatesAndMalformedSignedPayloads(t *testing.T) {
	now := time.Unix(1700000000, 0)
	for _, tc := range []struct {
		name string
		date string
		opts ValidationOptions
		want error
	}{
		{"valid", "1699999999", ValidationOptions{MaxAge: time.Minute}, nil},
		{"age boundary", "1699999940", ValidationOptions{MaxAge: time.Minute}, nil},
		{"expired", "1699999939", ValidationOptions{MaxAge: time.Minute}, ErrExpired},
		{"age disabled", "1", ValidationOptions{}, nil},
		{"future", "1700000001", ValidationOptions{}, ErrFutureAuthDate},
		{"future allowed", "1700000001", ValidationOptions{ClockSkew: time.Second}, nil},
		{"age skew allowed", "1699999939", ValidationOptions{ClockSkew: time.Second, MaxAge: time.Minute}, nil},
		{"huge future", "9223372036854775807", ValidationOptions{}, ErrFutureAuthDate},
		{"missing", "", ValidationOptions{}, ErrAuthDate},
		{"invalid", "abc", ValidationOptions{}, ErrAuthDate},
		{"zero", "0", ValidationOptions{}, ErrAuthDate},
		{"negative", "-1", ValidationOptions{}, ErrAuthDate},
		{"overflow", "9223372036854775808", ValidationOptions{}, ErrAuthDate},
		{"negative age limit", "1700000000", ValidationOptions{MaxAge: -1}, ErrInvalidInitData},
		{"negative skew", "1700000000", ValidationOptions{ClockSkew: -1}, ErrInvalidInitData},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.opts.Now = now
			values := url.Values{}
			if tc.date != "" {
				values.Set("auth_date", tc.date)
			}
			data, err := ValidateTelegramInitData(signedQuery(values), fixtureToken, &tc.opts)
			if !errors.Is(err, tc.want) || (err != nil && data != nil) {
				t.Fatalf("validation = %v, %v; want %v", data, err, tc.want)
			}
		})
	}
	for key, value := range map[string]string{"user": "{", "receiver": "[]", "chat": `{"id":"invalid"}`, "can_send_after": "-1"} {
		raw := signedQuery(url.Values{"auth_date": {"1700000000"}, key: {value}})
		if data, err := ValidateTelegramInitData(raw, fixtureToken, &ValidationOptions{Now: now}); data != nil || !errors.Is(err, ErrInvalidInitData) {
			t.Fatalf("invalid %s accepted: %v, %v", key, data, err)
		}
	}
	unsigned, err := ParseTelegramInitData("user=%7B%22id%22%3A42%2C%22first_name%22%3A%22test%22%7D")
	if err != nil || unsigned.User.ID != 42 {
		t.Fatalf("parsing unsigned data: %v, %v", unsigned, err)
	}
	if _, err := ValidateTelegramInitData("user=%7B%22id%22%3A42%7D", fixtureToken, nil); !errors.Is(err, ErrInvalidHash) {
		t.Fatalf("unsigned data validated: %v", err)
	}
	if data, err := ParseTelegramInitData(""); err != nil || data.User != nil || !data.AuthTime().IsZero() || !data.CanSendAt().IsZero() {
		t.Fatalf("empty launch: %v, %v", data, err)
	}
	var absent *WebAppInitData
	if !absent.AuthTime().IsZero() || !absent.CanSendAt().IsZero() {
		t.Fatal("nil init data has timestamps")
	}
	if !(&WebAppInitData{AuthDate: 1, CanSendAfter: 1<<63 - 1}).CanSendAt().IsZero() {
		t.Fatal("overflowed send timestamp")
	}
}

func TestEd25519InitData(t *testing.T) {
	key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
	publicKey := key.Public().(ed25519.PublicKey)
	user := `{"id":42,"first_name":"A+B & C"}`
	check := "123456:WebAppData\nauth_date=1700000000\nfuture_field=keep + %26 & =\nuser=" + user
	signature := ed25519.Sign(key, []byte(check))
	opts := &ValidationOptions{Now: time.Unix(1700000001, 0), MaxAge: time.Minute}
	for _, encoding := range []*base64.Encoding{base64.RawURLEncoding, base64.URLEncoding} {
		values := url.Values{
			"user": {user}, "auth_date": {"1700000000"}, "future_field": {"keep + %26 & ="},
			"signature": {encoding.EncodeToString(signature)}, "hash": {"ignored by Ed25519"},
		}
		for _, includeHash := range []bool{true, false} {
			if !includeHash {
				values.Del("hash")
			}
			data, err := validateInitDataSignature(values.Encode(), 123456, publicKey, opts)
			if err != nil || data.User == nil || data.User.ID != 42 {
				t.Fatalf("signature validation: %v, %v", data, err)
			}
		}
		if _, err := validateInitDataSignature(values.Encode(), 123457, publicKey, opts); !errors.Is(err, ErrInvalidSignature) {
			t.Fatalf("signature accepted for another bot: %v", err)
		}
		for _, testEnvironment := range []bool{false, true} {
			if _, err := ValidateTelegramInitDataSignature(values.Encode(), 123456, &ValidationOptions{TestEnvironment: testEnvironment}); !errors.Is(err, ErrInvalidSignature) {
				t.Fatalf("non-Telegram signature accepted: %v", err)
			}
		}
		values.Set("future_field", "changed")
		if data, err := validateInitDataSignature(values.Encode(), 123456, publicKey, opts); data != nil || !errors.Is(err, ErrInvalidSignature) {
			t.Fatalf("tampered signature accepted: %v, %v", data, err)
		}
	}
	for _, signature := range []string{"", "a", "???", base64.RawURLEncoding.EncodeToString(make([]byte, 63))} {
		if _, err := validateInitDataSignature(url.Values{"signature": {signature}}.Encode(), 123456, publicKey, opts); !errors.Is(err, ErrInvalidSignature) {
			t.Fatalf("invalid signature %q accepted: %v", signature, err)
		}
	}
	for _, botID := range []int64{0, -1} {
		if _, err := ValidateTelegramInitDataSignature("", botID, nil); !errors.Is(err, ErrInvalidBotID) {
			t.Fatalf("invalid bot ID accepted: %v", err)
		}
	}
}

func FuzzParseTelegramInitData(f *testing.F) {
	for _, input := range []string{fixtureInitData, "", "a=%", "a=b&a=c", "user=null", "auth_date=9223372036854775807"} {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, raw string) {
		data, err := ParseTelegramInitData(raw)
		if err == nil {
			data.AuthTime()
			data.CanSendAt()
		}
		ValidateTelegramInitData(raw, fixtureToken, &ValidationOptions{Now: time.Unix(1700000000, 0)})
		ValidateTelegramInitDataSignature(raw, 123456, nil)
	})
}
