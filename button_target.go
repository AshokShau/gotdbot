package gotdbot

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isButtonTarget(target string) bool {
	kind, _, _ := strings.Cut(target, ":")
	switch kind {
	case "buttonurl", "buttoncallback", "buttoncallback64", "buttonpassword", "buttonpassword64",
		"buttoncopy", "buttonbuy", "buttongame", "buttondisabled", "buttonlogin", "buttonswitch", "buttonuser", "buttonwebapp":
		return true
	}
	return false
}

func parseInlineButton(name, target string) (InlineKeyboardButton, bool, error) {
	button := InlineKeyboardButton{Text: name}
	if strings.TrimSpace(name) == "" || !utf8.ValidString(name) {
		return button, false, fmt.Errorf("button label must contain UTF-8 text")
	}
	same := strings.HasSuffix(target, ":same")
	if same {
		colon := len(target) - len(":same")
		for pos := colon - 1; pos >= 0 && target[pos] == '\\'; pos-- {
			same = !same
		}
		if same {
			target = target[:colon]
		}
	}
	fields := splitButtonFields(target)
	kind, value, _ := strings.Cut(fields[0], ":")
	value = unescapeButtonValue(value)
	options := make(map[string]string)
	for _, field := range fields[1:] {
		key, value, ok := strings.Cut(field, "=")
		if _, duplicate := options[key]; duplicate || !ok || key == "" {
			return button, false, fmt.Errorf("invalid button option %q", field)
		}
		options[key] = unescapeButtonValue(value)
	}

	if style, exists := takeButtonOption(options, "style"); exists {
		switch style {
		case "default":
			button.Style = &ButtonStyleDefault{}
		case "primary":
			button.Style = &ButtonStylePrimary{}
		case "success":
			button.Style = &ButtonStyleSuccess{}
		case "danger":
			button.Style = &ButtonStyleDanger{}
		case "link":
			button.Style = &ButtonStyleLink{}
		default:
			return button, false, fmt.Errorf("unknown button style %q", style)
		}
	}
	if icon, exists := takeButtonOption(options, "icon"); exists {
		id, err := strconv.ParseInt(icon, 10, 64)
		if err != nil {
			return button, false, fmt.Errorf("invalid button icon: %w", err)
		}
		button.IconCustomEmojiId = id
	}
	buttonType, err := parseInlineButtonType(kind, value, options)
	if err != nil {
		return button, false, err
	}
	button.Type = buttonType

	for key := range options {
		return button, false, fmt.Errorf("unsupported %s option %q", kind, key)
	}
	return button, same, nil
}

func takeButtonOption(options map[string]string, key string) (string, bool) {
	value, exists := options[key]
	delete(options, key)
	return value, exists
}

func parseInlineButtonType(kind, value string, options map[string]string) (InlineKeyboardButtonType, error) {
	switch kind {
	case "buttonurl", "buttonlogin", "buttonwebapp":
		value = strings.TrimPrefix(strings.TrimSpace(value), "//")
		parsed, err := url.Parse(value)
		if err == nil && parsed.Scheme == "" {
			value = "https://" + value
			parsed, err = url.Parse(value)
		}
		if err != nil || parsed.Hostname() == "" || (parsed.Scheme != "https" && parsed.Scheme != "http" && parsed.Scheme != "tg") {
			return nil, fmt.Errorf("invalid button URL %q", value)
		}
		if kind == "buttonurl" {
			return &InlineKeyboardButtonTypeUrl{Url: value}, nil
		}
		if kind == "buttonwebapp" {
			return &InlineKeyboardButtonTypeWebApp{Url: value}, nil
		}
		login := &InlineKeyboardButtonTypeLoginUrl{Url: value}
		if id, exists := takeButtonOption(options, "id"); exists {
			login.Id, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid login button ID: %w", err)
			}
		}
		login.ForwardText, _ = takeButtonOption(options, "forward")
		return login, nil
	case "buttoncallback", "buttoncallback64", "buttonpassword", "buttonpassword64":
		data := []byte(value)
		if strings.HasSuffix(kind, "64") {
			var err error
			data, err = base64.RawURLEncoding.Strict().DecodeString(value)
			if err != nil {
				return nil, fmt.Errorf("invalid callback data: %w", err)
			}
		}
		if len(data) < 1 || len(data) > 64 || (!strings.HasSuffix(kind, "64") && !utf8.Valid(data)) {
			return nil, fmt.Errorf("button callback must contain 1-64 bytes")
		}
		if strings.HasPrefix(kind, "buttonpassword") {
			return &InlineKeyboardButtonTypeCallbackWithPassword{Data: data}, nil
		}
		return &InlineKeyboardButtonTypeCallback{Data: data}, nil
	case "buttoncopy":
		if value == "" {
			return nil, fmt.Errorf("button copy text is empty")
		}
		return &InlineKeyboardButtonTypeCopyText{Text: value}, nil
	case "buttonbuy", "buttongame", "buttondisabled":
		if value != "" {
			return nil, fmt.Errorf("%s takes no value", kind)
		}
		if kind == "buttonbuy" {
			return &InlineKeyboardButtonTypeBuy{}, nil
		}
		if kind == "buttongame" {
			return &InlineKeyboardButtonTypeCallbackGame{}, nil
		}
		return &InlineKeyboardButtonTypeDisabled{}, nil
	case "buttonuser":
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid button user ID: %w", err)
		}
		return &InlineKeyboardButtonTypeUser{UserId: id}, nil
	case "buttonswitch":
		chat, err := parseButtonTargetChat(options)
		if err != nil {
			return nil, err
		}
		return &InlineKeyboardButtonTypeSwitchInline{Query: value, TargetChat: chat}, nil
	default:
		return nil, fmt.Errorf("unsupported button markup %q", kind)
	}
}

func parseButtonTargetChat(options map[string]string) (TargetChat, error) {
	target, _ := takeButtonOption(options, "target")
	switch target {
	case "":
		return nil, nil
	case "current":
		return &TargetChatCurrent{}, nil
	case "chosen":
		chosen := &TargetChatChosen{}
		if chats, exists := takeButtonOption(options, "chats"); exists {
			chosen.Types = &TargetChatTypes{}
			if chats != "" {
				for _, chat := range strings.Split(chats, ",") {
					switch chat {
					case "users":
						chosen.Types.AllowUserChats = true
					case "bots":
						chosen.Types.AllowBotChats = true
					case "groups":
						chosen.Types.AllowGroupChats = true
					case "channels":
						chosen.Types.AllowChannelChats = true
					default:
						return nil, fmt.Errorf("unknown button chat type %q", chat)
					}
				}
			}
		}
		return chosen, nil
	case "link":
		chat := &TargetChatInternalLink{}
		if link, exists := takeButtonOption(options, "link"); exists {
			var err error
			chat.Link, err = UnmarshalInternalLinkType([]byte(link))
			if err != nil || formattingValue(chat.Link) == nil {
				return nil, fmt.Errorf("invalid button internal link")
			}
		}
		return chat, nil
	default:
		return nil, fmt.Errorf("unknown button target %q", target)
	}
}

func inlineButtonTarget(button InlineKeyboardButton) (string, error) {
	if strings.TrimSpace(button.Text) == "" || !utf8.ValidString(button.Text) {
		return "", fmt.Errorf("button label must contain UTF-8 text")
	}
	var kind, value string
	var options []string
	add := func(key, value string) { options = append(options, key+"="+escapeButtonValue(value)) }
	switch typ := formattingValue(button.Type).(type) {
	case InlineKeyboardButtonTypeUrl:
		kind, value = "buttonurl", typ.Url
	case InlineKeyboardButtonTypeWebApp:
		kind, value = "buttonwebapp", typ.Url
	case InlineKeyboardButtonTypeCallback:
		kind, value = callbackButtonValue("buttoncallback", typ.Data)
	case InlineKeyboardButtonTypeCallbackWithPassword:
		kind, value = callbackButtonValue("buttonpassword", typ.Data)
	case InlineKeyboardButtonTypeCopyText:
		kind, value = "buttoncopy", typ.Text
	case InlineKeyboardButtonTypeBuy:
		kind = "buttonbuy"
	case InlineKeyboardButtonTypeCallbackGame:
		kind = "buttongame"
	case InlineKeyboardButtonTypeDisabled:
		kind = "buttondisabled"
	case InlineKeyboardButtonTypeUser:
		kind, value = "buttonuser", strconv.FormatInt(typ.UserId, 10)
	case InlineKeyboardButtonTypeLoginUrl:
		kind, value = "buttonlogin", typ.Url
		if typ.Id != 0 {
			add("id", strconv.FormatInt(typ.Id, 10))
		}
		if typ.ForwardText != "" {
			add("forward", typ.ForwardText)
		}
	case InlineKeyboardButtonTypeSwitchInline:
		kind, value = "buttonswitch", typ.Query
		switch chat := formattingValue(typ.TargetChat).(type) {
		case nil:
		case TargetChatCurrent:
			add("target", "current")
		case TargetChatChosen:
			add("target", "chosen")
			if chat.Types != nil {
				var chats []string
				if chat.Types.AllowUserChats {
					chats = append(chats, "users")
				}
				if chat.Types.AllowBotChats {
					chats = append(chats, "bots")
				}
				if chat.Types.AllowGroupChats {
					chats = append(chats, "groups")
				}
				if chat.Types.AllowChannelChats {
					chats = append(chats, "channels")
				}
				add("chats", strings.Join(chats, ","))
			}
		case TargetChatInternalLink:
			add("target", "link")
			if formattingValue(chat.Link) != nil {
				link, err := json.Marshal(chat.Link)
				if err != nil {
					return "", err
				}
				add("link", string(link))
			}
		default:
			return "", fmt.Errorf("unsupported button target %T", typ.TargetChat)
		}
	default:
		return "", fmt.Errorf("unsupported button type %T", button.Type)
	}
	switch formattingValue(button.Style).(type) {
	case nil:
	case ButtonStyleDefault:
		add("style", "default")
	case ButtonStylePrimary:
		add("style", "primary")
	case ButtonStyleSuccess:
		add("style", "success")
	case ButtonStyleDanger:
		add("style", "danger")
	case ButtonStyleLink:
		add("style", "link")
	default:
		return "", fmt.Errorf("unsupported button style %T", button.Style)
	}
	if button.IconCustomEmojiId != 0 {
		add("icon", strconv.FormatInt(button.IconCustomEmojiId, 10))
	}
	target := kind + ":" + escapeButtonValue(value)
	if len(options) > 0 {
		target += "|" + strings.Join(options, "|")
	}
	if _, _, err := parseInlineButton(button.Text, target); err != nil {
		return "", err
	}
	return target, nil
}

func callbackButtonValue(kind string, data []byte) (string, string) {
	if utf8.Valid(data) && strings.IndexFunc(string(data), unicode.IsControl) < 0 {
		return kind, string(data)
	}
	return kind + "64", base64.RawURLEncoding.EncodeToString(data)
}

func splitButtonFields(target string) []string {
	var fields []string
	start := 0
	for pos := 0; pos < len(target); pos++ {
		if target[pos] == '\\' {
			pos++
		} else if target[pos] == '|' {
			fields = append(fields, target[start:pos])
			start = pos + 1
		}
	}
	return append(fields, target[start:])
}

func escapeButtonValue(value string) string {
	value = strings.NewReplacer("\\", "\\\\", "|", "\\|", "(", "\\(", ")", "\\)", "\n", "\\n", "\r", "\\r").Replace(value)
	if strings.HasSuffix(value, ":same") {
		value = strings.TrimSuffix(value, ":same") + "\\:same"
	}
	return value
}

func unescapeButtonValue(value string) string {
	return strings.NewReplacer("\\\\", "\\", "\\|", "|", "\\(", "(", "\\)", ")", "\\:", ":", "\\n", "\n", "\\r", "\r").Replace(value)
}
