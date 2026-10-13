package gotdbot

import (
	"fmt"
)

// currentBotID reads the identity cached during authorization.
func (c *Client) currentBotID() (int64, error) {
	c.meMu.RLock()
	defer c.meMu.RUnlock()
	me := c.Me
	if me == nil {
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
