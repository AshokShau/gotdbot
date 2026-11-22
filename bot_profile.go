package gotdbot

import (
	"fmt"
)

// currentBotID reads the identity cached during authorization.
func (c *Client) currentBotID() (int64, error) {
	me := c.Me()
	if me.Id == 0 {
		return 0, fmt.Errorf("client identity unavailable; call Start first")
	}
	if bot, ok := me.Type.(*UserTypeBot); !ok || bot == nil {
		return 0, fmt.Errorf("operation requires a bot account")
	}
	return me.Id, nil
}

// GetMyName returns the bot's name in the requested language.
func (c *Client) GetMyName(languageCode string) (*Text, error) {
	id, err := c.currentBotID()
	if err != nil {
		return nil, err
	}
	return c.GetBotName(id, &GetBotNameOpts{LanguageCode: languageCode})
}

// SetMyName sets the bot's name in the requested language.
func (c *Client) SetMyName(name string, opts *SetBotNameOpts) error {
	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotName(id, name, opts)
}

// GetMyDescription returns the bot's description in the requested language.
func (c *Client) GetMyDescription(languageCode string) (*Text, error) {
	id, err := c.currentBotID()
	if err != nil {
		return nil, err
	}
	return c.GetBotInfoDescription(id, &GetBotInfoDescriptionOpts{LanguageCode: languageCode})
}

// SetMyDescription sets the bot's description in the requested language.
func (c *Client) SetMyDescription(description string, opts *SetBotInfoDescriptionOpts) error {
	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotInfoDescription(id, description, opts)
}

// GetMyShortDescription returns the bot's short description.
func (c *Client) GetMyShortDescription(languageCode string) (*Text, error) {
	id, err := c.currentBotID()
	if err != nil {
		return nil, err
	}
	return c.GetBotInfoShortDescription(id, &GetBotInfoShortDescriptionOpts{LanguageCode: languageCode})
}

// SetMyShortDescription sets the bot's short description.
func (c *Client) SetMyShortDescription(description string, opts *SetBotInfoShortDescriptionOpts) error {
	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotInfoShortDescription(id, description, opts)
}

// SetMySettingsOpts contains the optional settings for SetMySettings.
// A nil pointer leaves the corresponding setting unchanged.
type SetMySettingsOpts struct {
	// SupportsInlineQueries enables or disables inline mode.
	SupportsInlineQueries *bool
	// InlineQueryPlaceholder is shown after the bot's username in the input field.
	// An empty string restores the default placeholder.
	InlineQueryPlaceholder *string
	// NeedsInlineQueryLocation requests the user's location with each inline query.
	NeedsInlineQueryLocation *bool
	// InlineFeedbackShareDenominator is the share of chosen inline results received
	// by the bot: 1 in this many results. Use 1, 10, 100, or 1000, or 0 to stop.
	InlineFeedbackShareDenominator *int32
	// CanJoinGroups allows the bot to be added to groups.
	CanJoinGroups *bool
	// CanReadAllGroupMessages disables privacy mode when true.
	CanReadAllGroupMessages *bool
	// CanConnectToBusiness allows the bot to manage connected user accounts.
	CanConnectToBusiness *bool
	// CanManageBots allows the bot to create and control other bots.
	CanManageBots *bool
	// SupportsGuestQueries allows queries from chats the bot isn't a member of.
	SupportsGuestQueries *bool
	// SupportsJoinRequestQueries allows the bot to process join request queries.
	SupportsJoinRequestQueries *bool
	// HasTopicsEnabled enables forum topics in private chats with the bot.
	HasTopicsEnabled *bool
	// AllowsUsersToCreateTopics lets users create and delete private chat topics.
	AllowsUsersToCreateTopics *bool
	// HasTopicSidebar shows private chat topics in a sidebar.
	HasTopicSidebar *bool
	// CanCommunicateWithBots allows communication with other bots.
	CanCommunicateWithBots *bool
	// CallbackQueryMode controls how queries from the bot's buttons are received:
	// "default" uses message, "inline" uses inline_message_id, and "auto" uses
	// inline_message_id when the message is inaccessible.
	CallbackQueryMode *string
}

// SetMySettings changes settings of the bot. Settings that aren't specified
// remain unchanged. It returns nil on success, corresponding to True in the Bot API.
func (c *Client) SetMySettings(opts *SetMySettingsOpts) error {
	if opts == nil {
		return fmt.Errorf("settings required")
	}

	settings := make([]BotSetting, 0, 15)
	if v := opts.SupportsInlineQueries; v != nil {
		settings = append(settings, BotSettingIsInline{IsInline: *v})
	}
	if v := opts.InlineQueryPlaceholder; v != nil {
		settings = append(settings, BotSettingInlineQueryPlaceholder{Placeholder: *v})
	}
	if v := opts.NeedsInlineQueryLocation; v != nil {
		settings = append(settings, BotSettingNeedLocation{NeedLocation: *v})
	}
	if v := opts.InlineFeedbackShareDenominator; v != nil {
		switch *v {
		case 0, 1, 10, 100, 1000:
		default:
			return fmt.Errorf("inline_feedback_share_denominator must be 0, 1, 10, 100, or 1000")
		}
		settings = append(settings, BotSettingInlineFeedback{ShareDenominator: *v})
	}
	if v := opts.CanJoinGroups; v != nil {
		settings = append(settings, BotSettingCanJoinGroups{CanJoinGroups: *v})
	}
	if v := opts.CanReadAllGroupMessages; v != nil {
		settings = append(settings, BotSettingCanReadAllGroupMessages{CanReadAllGroupMessages: *v})
	}
	if v := opts.CanConnectToBusiness; v != nil {
		settings = append(settings, BotSettingCanConnectToBusiness{CanConnectToBusiness: *v})
	}
	if v := opts.CanManageBots; v != nil {
		settings = append(settings, BotSettingCanManageBots{CanManageBots: *v})
	}
	if v := opts.SupportsGuestQueries; v != nil {
		settings = append(settings, BotSettingSupportsGuestQueries{SupportsGuestQueries: *v})
	}
	if v := opts.SupportsJoinRequestQueries; v != nil {
		settings = append(settings, BotSettingIsGuard{IsGuard: *v})
	}
	if v := opts.HasTopicsEnabled; v != nil {
		settings = append(settings, BotSettingHasTopics{HasTopics: *v})
	}
	if v := opts.AllowsUsersToCreateTopics; v != nil {
		settings = append(settings, BotSettingAllowsUsersToCreateTopics{AllowsUsersToCreateTopics: *v})
	}
	if v := opts.HasTopicSidebar; v != nil {
		settings = append(settings, BotSettingHasTopicSidebar{HasTopicSidebar: *v})
	}
	if v := opts.CanCommunicateWithBots; v != nil {
		settings = append(settings, BotSettingCanCommunicateWithBots{CanCommunicateWithBots: *v})
	}
	if v := opts.CallbackQueryMode; v != nil {
		var mode BotCallbackQueryMode
		switch *v {
		case "default":
			mode = BotCallbackQueryModeDefault{}
		case "auto":
			mode = BotCallbackQueryModeAuto{}
		case "inline":
			mode = BotCallbackQueryModeInline{}
		default:
			return fmt.Errorf("callback_query_mode must be default, auto, or inline")
		}
		settings = append(settings, BotSettingCallbackQueryMode{Mode: mode})
	}

	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotSettings(id, settings)
}

// SetMyProfilePhoto sets the bot's native profile photo.
func (c *Client) SetMyProfilePhoto(photo InputChatPhoto) error {
	if photo == nil {
		return fmt.Errorf("profile photo required; use RemoveMyProfilePhoto")
	}
	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotProfilePhoto(id, &SetBotProfilePhotoOpts{Photo: photo})
}

// RemoveMyProfilePhoto removes the bot's profile photo.
func (c *Client) RemoveMyProfilePhoto() error {
	id, err := c.currentBotID()
	if err != nil {
		return err
	}
	return c.SetBotProfilePhoto(id, nil)
}

// SetMyDefaultAdministratorRights sets rights requested when the bot is added.
// Pass nil to remove defaults. forChannels selects channel rather than group rights.
func (c *Client) SetMyDefaultAdministratorRights(rights *ChatAdministratorRights, forChannels bool) error {
	if forChannels {
		return c.SetDefaultChannelAdministratorRights(rights)
	}
	return c.SetDefaultGroupAdministratorRights(rights)
}

// GetMyDefaultAdministratorRights gets the bot's default channel or group rights.
func (c *Client) GetMyDefaultAdministratorRights(forChannels bool) (*ChatAdministratorRights, error) {
	id, err := c.currentBotID()
	if err != nil {
		return nil, err
	}
	info, err := c.GetUserFullInfo(id)
	if err != nil {
		return nil, err
	}
	if info.BotInfo == nil {
		return nil, fmt.Errorf("bot information unavailable")
	}
	if forChannels {
		return info.BotInfo.DefaultChannelAdministratorRights, nil
	}
	return info.BotInfo.DefaultGroupAdministratorRights, nil
}
