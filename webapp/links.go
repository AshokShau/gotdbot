package webapp

import (
	"fmt"
	"net/url"
	"strings"
)

type LaunchMode string

const (
	LaunchModeCompact    LaunchMode = "compact"
	LaunchModeFullscreen LaunchMode = "fullscreen"
)

type MiniAppLinkOptions struct {
	StartParam string
	Mode       LaunchMode
}

type ChooseChatType string

const (
	ChooseUsers    ChooseChatType = "users"
	ChooseBots     ChooseChatType = "bots"
	ChooseGroups   ChooseChatType = "groups"
	ChooseChannels ChooseChatType = "channels"
)

type AttachmentMenuLinkOptions struct {
	StartParam string
	Target     string
	Choose     []ChooseChatType
}

// MiniAppLink opens the main Mini App when appName is empty, or a named Mini App.
func MiniAppLink(botUsername, appName string, opts *MiniAppLinkOptions) (string, error) {
	botUsername = strings.TrimPrefix(botUsername, "@")
	if !validLinkName(botUsername) || (appName != "" && !validLinkName(appName)) {
		return "", fmt.Errorf("invalid Mini App bot username or app name")
	}
	var options MiniAppLinkOptions
	if opts != nil {
		options = *opts
	}
	if options.Mode != "" && options.Mode != LaunchModeCompact && options.Mode != LaunchModeFullscreen {
		return "", fmt.Errorf("invalid Mini App launch mode")
	}
	path := botUsername
	if appName != "" {
		path += "/" + appName
	}
	query := url.Values{"startapp": {options.StartParam}}
	if options.Mode != "" {
		query.Set("mode", string(options.Mode))
	}
	return (&url.URL{Scheme: "https", Host: "t.me", Path: "/" + path, RawQuery: query.Encode()}).String(), nil
}

func AttachmentMenuLink(botUsername string, opts *AttachmentMenuLinkOptions) (string, error) {
	botUsername = strings.TrimPrefix(botUsername, "@")
	if !validLinkName(botUsername) {
		return "", fmt.Errorf("invalid attachment menu bot username")
	}
	var options AttachmentMenuLinkOptions
	if opts != nil {
		options = *opts
	}
	query := url.Values{"startattach": {options.StartParam}}
	path := botUsername
	if options.Target != "" {
		if len(options.Choose) != 0 {
			return "", fmt.Errorf("attachment menu target and chat choices are mutually exclusive")
		}
		target := strings.TrimPrefix(options.Target, "@")
		valid := validLinkName(target)
		if strings.HasPrefix(target, "+") {
			valid = len(target) > 1
			for _, digit := range target[1:] {
				valid = valid && digit >= '0' && digit <= '9'
			}
		}
		if !valid {
			return "", fmt.Errorf("invalid attachment menu target")
		}
		path = target
		query.Set("attach", botUsername)
	}
	if len(options.Choose) > 0 {
		choices := make([]string, 0, len(options.Choose))
		seen := make(map[ChooseChatType]bool)
		for _, choice := range options.Choose {
			switch choice {
			case ChooseUsers, ChooseBots, ChooseGroups, ChooseChannels:
				if !seen[choice] {
					choices = append(choices, string(choice))
					seen[choice] = true
				}
			default:
				return "", fmt.Errorf("invalid attachment menu chat choice")
			}
		}
		query.Set("choose", strings.Join(choices, " "))
	}
	return (&url.URL{Scheme: "https", Host: "t.me", Path: "/" + path, RawQuery: query.Encode()}).String(), nil
}

func validLinkName(name string) bool {
	if name == "" {
		return false
	}
	for _, r := range name {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_') {
			return false
		}
	}
	return true
}
