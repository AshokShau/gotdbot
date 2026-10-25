package webapp

import (
	"net/url"
	"strings"
	"testing"
)

func TestMiniAppLinks(t *testing.T) {
	for _, tc := range []struct {
		appName string
		mode    LaunchMode
	}{
		{"", ""}, {"shop", LaunchModeCompact}, {"shop", LaunchModeFullscreen},
	} {
		link, err := MiniAppLink("@ExampleBot", tc.appName, &MiniAppLinkOptions{StartParam: "a+b & c=é%", Mode: tc.mode})
		if err != nil {
			t.Fatal(err)
		}
		parsed, err := url.Parse(link)
		if err != nil {
			t.Fatal(err)
		}
		wantPath := "/ExampleBot"
		if tc.appName != "" {
			wantPath += "/" + tc.appName
		}
		if parsed.Scheme != "https" || parsed.Host != "t.me" || parsed.Path != wantPath || parsed.Query().Get("startapp") != "a+b & c=é%" || parsed.Query().Get("mode") != string(tc.mode) || len(parsed.Query()["startapp"]) != 1 {
			t.Fatalf("incorrect link: %s", link)
		}
	}
	if link, err := MiniAppLink("ExampleBot", "", nil); err != nil || link != "https://t.me/ExampleBot?startapp=" {
		t.Fatalf("main Mini App: %q, %v", link, err)
	}
	for _, tc := range []struct{ bot, app string }{{"", ""}, {"example.org/bot", ""}, {"bot", "../other"}, {"bot?x=y", "app"}} {
		if link, err := MiniAppLink(tc.bot, tc.app, nil); link != "" || err == nil {
			t.Fatalf("invalid path accepted: %q", link)
		}
	}
	if _, err := MiniAppLink("ExampleBot", "app", &MiniAppLinkOptions{Mode: "invalid"}); err == nil {
		t.Fatal("invalid mode accepted")
	}
}

func TestAttachmentMenuLinks(t *testing.T) {
	for _, target := range []string{"", "@partner", "+1234567890"} {
		link, err := AttachmentMenuLink("@ExampleBot", &AttachmentMenuLinkOptions{Target: target, StartParam: "a+b&c"})
		if err != nil {
			t.Fatal(err)
		}
		parsed, err := url.Parse(link)
		if err != nil {
			t.Fatal(err)
		}
		if parsed.Host != "t.me" || parsed.Query().Get("startattach") != "a+b&c" {
			t.Fatalf("incorrect attachment link: %s", link)
		}
		if target == "" {
			if parsed.Path != "/ExampleBot" || parsed.Query().Has("attach") {
				t.Fatalf("current chat attachment link: %s", link)
			}
		} else if parsed.Path != "/"+strings.TrimPrefix(target, "@") || parsed.Query().Get("attach") != "ExampleBot" {
			t.Fatalf("target attachment link: %s", link)
		}
	}
	link, err := AttachmentMenuLink("ExampleBot", &AttachmentMenuLinkOptions{Choose: []ChooseChatType{ChooseUsers, ChooseBots, ChooseGroups, ChooseChannels, ChooseUsers}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(link, "choose=users+bots+groups+channels") {
		t.Fatalf("incorrect choose encoding: %s", link)
	}
	for _, opts := range []*AttachmentMenuLinkOptions{
		{Target: "partner", Choose: []ChooseChatType{ChooseUsers}}, {Target: "+123/other"}, {Target: "+"},
		{Target: "example.org/path"}, {Choose: []ChooseChatType{"invalid"}},
	} {
		if _, err := AttachmentMenuLink("ExampleBot", opts); err == nil {
			t.Fatalf("invalid attachment options accepted: %+v", opts)
		}
	}
}

func TestAndroidUserAgent(t *testing.T) {
	raw := "Mozilla/5.0 (Linux; Android 14; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.136 Mobile Safari/537.36 Telegram-Android/11.3.3 (Google sdk_gphone64_arm64; Android 14; SDK 34; LOW)"
	data, err := ParseAndroidUserAgent(raw)
	if err != nil || data.AppVersion != "11.3.3" || data.DeviceModel != "Google sdk_gphone64_arm64" || data.AndroidVersion != "14" || data.SDKVersion != 34 || data.PerformanceClass != PerformanceLow {
		t.Fatalf("user agent = %+v, %v", data, err)
	}
	for _, raw := range []string{"", "Mozilla/5.0", "Telegram-Android/11.3.3 (Device; Android 14; SDK nope; LOW)"} {
		if _, err := ParseAndroidUserAgent(raw); err == nil {
			t.Fatalf("invalid user agent accepted: %q", raw)
		}
	}
}
