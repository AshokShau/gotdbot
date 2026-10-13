package gotdbot

// ReplyQuiz replies with a quiz. Explicit reply destinations take precedence.
func (m *Message) ReplyQuiz(c *Client, question string, options []InputPollOption, correctOptionIds []int32, opts *SendQuizOpts) (*Message, error) {
	copyOpts := SendQuizOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyToMessageID == 0 && copyOpts.ReplyTo == nil {
		copyOpts.ReplyToMessageID = m.Id
	}
	return c.SendQuiz(m.ChatId, question, options, correctOptionIds, &copyOpts)
}

// ReplyRichMessageMarkdown replies with rich Markdown; for bots only.
func (m *Message) ReplyRichMessageMarkdown(c *Client, text string, opts *SendRichTextOpts) (*Message, error) {
	copyOpts := richReplyOptions(m.Id, opts)
	return c.SendRichMessageMarkdown(m.ChatId, text, &copyOpts)
}

// ReplyRichMessageHTML replies with rich HTML; for bots only.
func (m *Message) ReplyRichMessageHTML(c *Client, text string, opts *SendRichTextOpts) (*Message, error) {
	copyOpts := richReplyOptions(m.Id, opts)
	return c.SendRichMessageHTML(m.ChatId, text, &copyOpts)
}

func richReplyOptions(messageId int64, opts *SendRichTextOpts) SendRichTextOpts {
	copyOpts := SendRichTextOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyToMessageID == 0 && copyOpts.ReplyTo == nil {
		copyOpts.ReplyToMessageID = messageId
	}
	return copyOpts
}

// ReplyLivePhoto replies with a live photo, preserving caller-owned options.
func (m *Message) ReplyLivePhoto(c *Client, photo, video InputFile, opts *SendPhotoOpts) (*Message, error) {
	copyOpts := SendPhotoOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyToMessageID == 0 && copyOpts.ReplyTo == nil {
		copyOpts.ReplyToMessageID = m.Id
	}
	return c.SendLivePhoto(m.ChatId, photo, video, &copyOpts)
}

// ReplyLiveLocation replies with a live location, preserving caller-owned options.
func (m *Message) ReplyLiveLocation(c *Client, location *LiveLocation, opts *SendLocationOpts) (*Message, error) {
	copyOpts := SendLocationOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyToMessageID == 0 && copyOpts.ReplyTo == nil {
		copyOpts.ReplyToMessageID = m.Id
	}
	return c.SendLiveLocation(m.ChatId, location, &copyOpts)
}

// ReplyText replies to the message with text.
func (m *Message) ReplyText(c *Client, text string, opts *SendMessageOpts) (*Message, error) {
	copyOpts := SendMessageOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts

	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendMessage(m.ChatId, text, opts)
}

// ReplyAlbum replies to the message with an album.
func (m *Message) ReplyAlbum(c *Client, inputMessageContents []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	copyOpts := SendMessageAlbumOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyTo == nil {
		opts.ReplyTo = &InputMessageReplyToMessage{
			MessageId: m.Id,
		}
	}
	return c.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// ReplyAnimation replies to the message with animation.
func (m *Message) ReplyAnimation(c *Client, animation InputFile, opts *SendAnimationOpts) (*Message, error) {
	copyOpts := SendAnimationOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAnimation(m.ChatId, animation, opts)
}

// ReplyAudio replies to the message with audio.
func (m *Message) ReplyAudio(c *Client, audio InputFile, opts *SendAudioOpts) (*Message, error) {
	copyOpts := SendAudioOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAudio(m.ChatId, audio, opts)
}

// ReplyDocument replies to the message with a document.
func (m *Message) ReplyDocument(c *Client, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	copyOpts := SendDocumentOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendDocument(m.ChatId, document, opts)
}

// ReplyPhoto replies to the message with a photo.
func (m *Message) ReplyPhoto(c *Client, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	copyOpts := SendPhotoOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPhoto(m.ChatId, photo, opts)
}

// ReplyVideo replies to the message with a video.
func (m *Message) ReplyVideo(c *Client, video InputFile, opts *SendVideoOpts) (*Message, error) {
	copyOpts := SendVideoOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideo(m.ChatId, video, opts)
}

// ReplyVideoNote replies to the message with a video note.
func (m *Message) ReplyVideoNote(c *Client, videoNote InputFile, opts *SendVideoNoteOpts) (*Message, error) {
	copyOpts := SendVideoNoteOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideoNote(m.ChatId, videoNote, opts)
}

// ReplyVoice replies to the message with a voice note.
func (m *Message) ReplyVoice(c *Client, voice InputFile, opts *SendVoiceOpts) (*Message, error) {
	copyOpts := SendVoiceOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVoice(m.ChatId, voice, opts)
}

// ReplySticker replies to the message with a sticker.
func (m *Message) ReplySticker(c *Client, sticker InputFile, opts *SendStickerOpts) (*Message, error) {
	copyOpts := SendStickerOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendSticker(m.ChatId, sticker, opts)
}

// ReplyMediaGroup replies with a validated native TDLib album.
func (m *Message) ReplyMediaGroup(c *Client, media []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	copyOpts := SendMessageAlbumOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyTo == nil {
		copyOpts.ReplyTo = &InputMessageReplyToMessage{MessageId: m.Id}
	}
	return c.SendMediaGroup(m.ChatId, media, &copyOpts)
}

// ReplyChecklist replies to the message with a checklist.
func (m *Message) ReplyChecklist(c *Client, checklist *InputChecklist, opts *SendChecklistOpts) (*Message, error) {
	copyOpts := SendChecklistOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendChecklist(m.ChatId, checklist, opts)
}

// ReplyContact replies to the message with a contact.
func (m *Message) ReplyContact(c *Client, contact *Contact, opts *SendContactOpts) (*Message, error) {
	copyOpts := SendContactOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendContact(m.ChatId, contact, opts)
}

// ReplyDice replies to the message with a dice.
func (m *Message) ReplyDice(c *Client, emoji string, opts *SendDiceOpts) (*Message, error) {
	copyOpts := SendDiceOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendDice(m.ChatId, emoji, opts)
}

// ReplyGame replies to the message with a game.
func (m *Message) ReplyGame(c *Client, botUserId int64, gameShortName string, opts *SendGameOpts) (*Message, error) {
	copyOpts := SendGameOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendGame(m.ChatId, botUserId, gameShortName, opts)
}

// ReplyInvoice replies to the message with an invoice.
func (m *Message) ReplyInvoice(c *Client, invoice *Invoice, title string, description string, payload []byte, opts *SendInvoiceOpts) (*Message, error) {
	copyOpts := SendInvoiceOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendInvoice(m.ChatId, invoice, title, description, payload, opts)
}

// ReplyLocation replies to the message with a location.
func (m *Message) ReplyLocation(c *Client, location *Location, opts *SendLocationOpts) (*Message, error) {
	copyOpts := SendLocationOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendLocation(m.ChatId, location, opts)
}

// ReplyPaidMedia replies to the message with paid media.
func (m *Message) ReplyPaidMedia(c *Client, starCount int64, paidMedia []InputPaidMedia, opts *SendPaidMediaOpts) (*Message, error) {
	copyOpts := SendPaidMediaOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPaidMedia(m.ChatId, starCount, paidMedia, opts)
}

// ReplyPoll replies to the message with a poll.
func (m *Message) ReplyPoll(c *Client, question string, options []InputPollOption, opts *SendPollOpts) (*Message, error) {
	copyOpts := SendPollOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPoll(m.ChatId, question, options, opts)
}

// ReplyStakeDice replies to the message with a stake dice.
func (m *Message) ReplyStakeDice(c *Client, stakeGramAmount int64, stateHash string, opts *SendStakeDiceOpts) (*Message, error) {
	copyOpts := SendStakeDiceOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendStakeDice(m.ChatId, stakeGramAmount, stateHash, opts)
}

// ReplyStory replies to the message with a story.
func (m *Message) ReplyStory(c *Client, storyPosterChatId int64, storyId int32, opts *SendStoryOpts) (*Message, error) {
	copyOpts := SendStoryOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendStory(m.ChatId, storyPosterChatId, storyId, opts)
}

// ReplyVenue replies to the message with a venue.
func (m *Message) ReplyVenue(c *Client, venue *Venue, opts *SendVenueOpts) (*Message, error) {
	copyOpts := SendVenueOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVenue(m.ChatId, venue, opts)
}

// ReplyCopy replies to the message by copying another message.
func (m *Message) ReplyCopy(c *Client, fromChatId int64, messageId int64, opts *CopyMessageOpts) (*Message, error) {
	copyOpts := CopyMessageOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.CopyMessage(m.ChatId, fromChatId, messageId, opts)
}

// ReplyContent replies to the message with the given content.
func (m *Message) ReplyContent(c *Client, content InputMessageContent, opts *SendMessageContentOpts) (*Message, error) {
	copyOpts := SendMessageContentOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyTo == nil {
		opts.ReplyTo = &InputMessageReplyToMessage{
			MessageId: m.Id,
		}
	}
	return c.SendMessageContent(m.ChatId, content, opts)
}

// ReplyRichMessage replies to the message with a rich message.
func (m *Message) ReplyRichMessage(c *Client, richMessage *InputRichMessage, opts *SendMessageOpts) (*Message, error) {
	copyOpts := SendMessageOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendRichMessage(m.ChatId, richMessage, opts)
}

// ReplyForwarded replies to the message by forwarding another message.
func (m *Message) ReplyForwarded(c *Client, fromChatId int64, messageId int64, opts *SendForwardedOpts) (*Message, error) {
	copyOpts := SendForwardedOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	opts = &copyOpts
	if opts.ReplyToMessageID == 0 && opts.ReplyTo == nil {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendForwarded(m.ChatId, fromChatId, messageId, opts)
}
