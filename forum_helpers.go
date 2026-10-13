package gotdbot

// CloseForumTopic closes a native TDLib forum topic.
func (c *Client) CloseForumTopic(chatId int64, topicId int32) error {
	return c.ToggleForumTopicIsClosed(chatId, topicId, &ToggleForumTopicIsClosedOpts{IsClosed: true})
}

// ReopenForumTopic reopens a native TDLib forum topic.
func (c *Client) ReopenForumTopic(chatId int64, topicId int32) error {
	return c.ToggleForumTopicIsClosed(chatId, topicId, &ToggleForumTopicIsClosedOpts{IsClosed: false})
}

// CloseGeneralForumTopic closes the General topic.
func (c *Client) CloseGeneralForumTopic(chatId int64) error { return c.CloseForumTopic(chatId, 1) }

// ReopenGeneralForumTopic reopens the General topic.
func (c *Client) ReopenGeneralForumTopic(chatId int64) error { return c.ReopenForumTopic(chatId, 1) }

// HideGeneralForumTopic hides the General topic.
func (c *Client) HideGeneralForumTopic(chatId int64) error {
	return c.ToggleGeneralForumTopicIsHidden(chatId, &ToggleGeneralForumTopicIsHiddenOpts{IsHidden: true})
}

// UnhideGeneralForumTopic unhides the General topic.
func (c *Client) UnhideGeneralForumTopic(chatId int64) error {
	return c.ToggleGeneralForumTopicIsHidden(chatId, &ToggleGeneralForumTopicIsHiddenOpts{IsHidden: false})
}

// UnpinAllGeneralForumTopicMessages unpins messages in the General topic.
func (c *Client) UnpinAllGeneralForumTopicMessages(chatId int64) error {
	return c.UnpinAllForumTopicMessages(chatId, 1)
}
