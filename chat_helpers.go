package gotdbot

// SendQuiz sends a quiz to this chat.
func (chat *Chat) SendQuiz(c *Client, question string, options []InputPollOption, correctOptionIds []int32, opts *SendQuizOpts) (*Message, error) {
	return c.SendQuiz(chat.Id, question, options, correctOptionIds, opts)
}

// SendRichMessage sends rich content to this chat.
func (chat *Chat) SendRichMessage(c *Client, message *InputRichMessage, opts *SendMessageOpts) (*Message, error) {
	return c.SendRichMessage(chat.Id, message, opts)
}

// SendRichMessageMarkdown sends rich Markdown to this chat.
func (chat *Chat) SendRichMessageMarkdown(c *Client, text string, opts *SendRichTextOpts) (*Message, error) {
	return c.SendRichMessageMarkdown(chat.Id, text, opts)
}

// SendRichMessageHTML sends rich HTML to this chat.
func (chat *Chat) SendRichMessageHTML(c *Client, text string, opts *SendRichTextOpts) (*Message, error) {
	return c.SendRichMessageHTML(chat.Id, text, opts)
}

// SendLivePhoto sends a live photo to this chat.
func (chat *Chat) SendLivePhoto(c *Client, photo, video InputFile, opts *SendPhotoOpts) (*Message, error) {
	return c.SendLivePhoto(chat.Id, photo, video, opts)
}

// SendLiveLocation sends a live location to this chat.
func (chat *Chat) SendLiveLocation(c *Client, location *LiveLocation, opts *SendLocationOpts) (*Message, error) {
	return c.SendLiveLocation(chat.Id, location, opts)
}

// SendMessage sends text to this chat with optional formatting and delivery settings.
func (chat *Chat) SendMessage(c *Client, text string, opts *SendMessageOpts) (*Message, error) {
	return c.SendMessage(chat.Id, text, opts)
}

// SendPhoto sends a photo to this chat.
func (chat *Chat) SendPhoto(c *Client, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	return c.SendPhoto(chat.Id, photo, opts)
}

// SendDocument sends a document to this chat.
func (chat *Chat) SendDocument(c *Client, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	return c.SendDocument(chat.Id, document, opts)
}

// SendVideo sends a video to this chat.
func (chat *Chat) SendVideo(c *Client, video InputFile, opts *SendVideoOpts) (*Message, error) {
	return c.SendVideo(chat.Id, video, opts)
}

// SendMediaGroup sends a validated media group to this chat.
func (chat *Chat) SendMediaGroup(c *Client, media []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return c.SendMediaGroup(chat.Id, media, opts)
}

// SendMessageDraft streams a partial text message to this chat.
func (chat *Chat) SendMessageDraft(c *Client, draftId int64, text string, opts *SendMessageDraftOpts) error {
	return c.SendMessageDraft(chat.Id, draftId, text, opts)
}

// SendRichMessageDraftMarkdown streams a partial rich Markdown message to this chat.
func (chat *Chat) SendRichMessageDraftMarkdown(c *Client, draftId int64, text string, opts *SendRichMessageDraftTextOpts) error {
	return c.SendRichMessageDraftMarkdown(chat.Id, draftId, text, opts)
}

// SendRichMessageDraftHTML streams a partial rich HTML message to this chat.
func (chat *Chat) SendRichMessageDraftHTML(c *Client, draftId int64, text string, opts *SendRichMessageDraftTextOpts) error {
	return c.SendRichMessageDraftHTML(chat.Id, draftId, text, opts)
}

// SendForwarded forwards a message to this chat.
func (chat *Chat) SendForwarded(c *Client, fromChatId int64, messageId int64, opts *SendForwardedOpts) (*Message, error) {
	return c.SendForwarded(chat.Id, fromChatId, messageId, opts)
}

// SendChecklist sends a checklist to this chat.
func (chat *Chat) SendChecklist(c *Client, checklist *InputChecklist, opts *SendChecklistOpts) (*Message, error) {
	return c.SendChecklist(chat.Id, checklist, opts)
}

// SendDice sends a dice message to this chat.
func (chat *Chat) SendDice(c *Client, emoji string, opts *SendDiceOpts) (*Message, error) {
	return c.SendDice(chat.Id, emoji, opts)
}

// SendGame sends a game to this chat.
func (chat *Chat) SendGame(c *Client, botUserId int64, gameShortName string, opts *SendGameOpts) (*Message, error) {
	return c.SendGame(chat.Id, botUserId, gameShortName, opts)
}

// SendPoll sends a poll to this chat.
func (chat *Chat) SendPoll(c *Client, question string, options []InputPollOption, opts *SendPollOpts) (*Message, error) {
	return c.SendPoll(chat.Id, question, options, opts)
}

// SendStakeDice sends a stake dice message to this chat.
func (chat *Chat) SendStakeDice(c *Client, stakeGramAmount int64, stateHash string, opts *SendStakeDiceOpts) (*Message, error) {
	return c.SendStakeDice(chat.Id, stakeGramAmount, stateHash, opts)
}

// SendStory sends a story to this chat.
func (chat *Chat) SendStory(c *Client, storyPosterChatId int64, storyId int32, opts *SendStoryOpts) (*Message, error) {
	return c.SendStory(chat.Id, storyPosterChatId, storyId, opts)
}

// SendContact sends a contact to this chat.
func (chat *Chat) SendContact(c *Client, contact *Contact, opts *SendContactOpts) (*Message, error) {
	return c.SendContact(chat.Id, contact, opts)
}

// SendLocation sends a location to this chat.
func (chat *Chat) SendLocation(c *Client, location *Location, opts *SendLocationOpts) (*Message, error) {
	return c.SendLocation(chat.Id, location, opts)
}

// SendVenue sends a venue to this chat.
func (chat *Chat) SendVenue(c *Client, venue *Venue, opts *SendVenueOpts) (*Message, error) {
	return c.SendVenue(chat.Id, venue, opts)
}

// SendAnimation sends an animation to this chat.
func (chat *Chat) SendAnimation(c *Client, animation InputFile, opts *SendAnimationOpts) (*Message, error) {
	return c.SendAnimation(chat.Id, animation, opts)
}

// SendAudio sends audio to this chat.
func (chat *Chat) SendAudio(c *Client, audio InputFile, opts *SendAudioOpts) (*Message, error) {
	return c.SendAudio(chat.Id, audio, opts)
}

// SendVoice sends a voice note to this chat.
func (chat *Chat) SendVoice(c *Client, voice InputFile, opts *SendVoiceOpts) (*Message, error) {
	return c.SendVoice(chat.Id, voice, opts)
}

// SendVideoNote sends a video note to this chat.
func (chat *Chat) SendVideoNote(c *Client, videoNote InputFile, opts *SendVideoNoteOpts) (*Message, error) {
	return c.SendVideoNote(chat.Id, videoNote, opts)
}

// SendSticker sends a sticker to this chat.
func (chat *Chat) SendSticker(c *Client, sticker InputFile, opts *SendStickerOpts) (*Message, error) {
	return c.SendSticker(chat.Id, sticker, opts)
}

// SendInvoice sends an invoice to this chat.
func (chat *Chat) SendInvoice(c *Client, invoice *Invoice, title string, description string, payload []byte, opts *SendInvoiceOpts) (*Message, error) {
	return c.SendInvoice(chat.Id, invoice, title, description, payload, opts)
}

// SendPaidMedia sends paid media to this chat.
func (chat *Chat) SendPaidMedia(c *Client, starCount int64, paidMedia []InputPaidMedia, opts *SendPaidMediaOpts) (*Message, error) {
	return c.SendPaidMedia(chat.Id, starCount, paidMedia, opts)
}
