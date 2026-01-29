package ext

import (
	"sync"
	"time"

	"github.com/AshokShau/gotdbot"
)

// ChatAction is a helper to send chat actions periodically.
type ChatAction struct {
	client  *gotdbot.Client
	chatID  int64
	action  *gotdbot.ChatAction
	topicID *gotdbot.MessageTopic
	ticker  *time.Ticker
	stopCh  chan struct{}
	mu      sync.Mutex
}

// NewChatAction creates a new ChatAction helper.
func NewChatAction(client *gotdbot.Client, chatID int64, actionStr string, topicID *gotdbot.MessageTopic) *ChatAction {
	var action *gotdbot.ChatAction

	switch actionStr {
	case "typing", "chatActionTyping":
		action = &gotdbot.ChatAction{ChatActionTyping: &gotdbot.ChatActionTyping{}}
	case "upload_photo", "chatActionUploadingPhoto":
		action = &gotdbot.ChatAction{ChatActionUploadingPhoto: &gotdbot.ChatActionUploadingPhoto{Progress: 50}}
	case "record_video", "chatActionRecordingVideo":
		action = &gotdbot.ChatAction{ChatActionRecordingVideo: &gotdbot.ChatActionRecordingVideo{}}
	case "upload_video", "chatActionUploadingVideo":
		action = &gotdbot.ChatAction{ChatActionUploadingVideo: &gotdbot.ChatActionUploadingVideo{Progress: 50}}
	case "record_voice", "chatActionRecordingVoiceNote":
		action = &gotdbot.ChatAction{ChatActionRecordingVoiceNote: &gotdbot.ChatActionRecordingVoiceNote{}}
	case "upload_voice", "chatActionUploadingVoiceNote":
		action = &gotdbot.ChatAction{ChatActionUploadingVoiceNote: &gotdbot.ChatActionUploadingVoiceNote{Progress: 50}}
	case "upload_document", "chatActionUploadingDocument":
		action = &gotdbot.ChatAction{ChatActionUploadingDocument: &gotdbot.ChatActionUploadingDocument{Progress: 50}}
	case "choose_sticker", "chatActionChoosingSticker":
		action = &gotdbot.ChatAction{ChatActionChoosingSticker: &gotdbot.ChatActionChoosingSticker{}}
	case "find_location", "chatActionChoosingLocation":
		action = &gotdbot.ChatAction{ChatActionChoosingLocation: &gotdbot.ChatActionChoosingLocation{}}
	case "record_video_note", "chatActionRecordingVideoNote":
		action = &gotdbot.ChatAction{ChatActionRecordingVideoNote: &gotdbot.ChatActionRecordingVideoNote{}}
	case "upload_video_note", "chatActionUploadingVideoNote":
		action = &gotdbot.ChatAction{ChatActionUploadingVideoNote: &gotdbot.ChatActionUploadingVideoNote{Progress: 50}}
	case "cancel", "chatActionCancel":
		action = &gotdbot.ChatAction{ChatActionCancel: &gotdbot.ChatActionCancel{}}
	default:
		action = &gotdbot.ChatAction{ChatActionTyping: &gotdbot.ChatActionTyping{}}
	}

	return &ChatAction{
		client:  client,
		chatID:  chatID,
		action:  action,
		topicID: topicID,
		stopCh:  make(chan struct{}),
	}
}

// Start starts sending the chat action every 4 seconds.
func (c *ChatAction) Start() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.ticker != nil {
		return
	}
	c.ticker = time.NewTicker(4 * time.Second)

	// Send immediately first
	c.send()

	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.send()
			case <-c.stopCh:
				c.ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops sending the chat action.
func (c *ChatAction) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.ticker != nil {
		close(c.stopCh)
		c.ticker = nil
		// Send cancel action
		cancelAction := &gotdbot.ChatAction{ChatActionCancel: &gotdbot.ChatActionCancel{}}
		c.client.SendChatAction(c.chatID, c.topicID, "", &gotdbot.SendChatActionOpts{Action: cancelAction})
	}
}

func (c *ChatAction) send() {
	c.client.SendChatAction(c.chatID, c.topicID, "", &gotdbot.SendChatActionOpts{Action: c.action})
}
