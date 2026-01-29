package client

import (
	"sync"
	"time"

	"github.com/AshokShau/gotdbot/types"
)

// ChatAction is a helper to send chat actions periodically.
type ChatAction struct {
	client  *Client
	chatID  int64
	action  *types.ChatAction
	topicID *types.MessageTopic
	ticker  *time.Ticker
	stopCh  chan struct{}
	mu      sync.Mutex
}

// NewChatAction creates a new ChatAction helper.
func NewChatAction(client *Client, chatID int64, actionStr string, topicID *types.MessageTopic) *ChatAction {
	var action *types.ChatAction

	switch actionStr {
	case "typing", "chatActionTyping":
		action = &types.ChatAction{ChatActionTyping: &types.ChatActionTyping{}}
	case "upload_photo", "chatActionUploadingPhoto":
		action = &types.ChatAction{ChatActionUploadingPhoto: &types.ChatActionUploadingPhoto{Progress: 50}}
	case "record_video", "chatActionRecordingVideo":
		action = &types.ChatAction{ChatActionRecordingVideo: &types.ChatActionRecordingVideo{}}
	case "upload_video", "chatActionUploadingVideo":
		action = &types.ChatAction{ChatActionUploadingVideo: &types.ChatActionUploadingVideo{Progress: 50}}
	case "record_voice", "chatActionRecordingVoiceNote":
		action = &types.ChatAction{ChatActionRecordingVoiceNote: &types.ChatActionRecordingVoiceNote{}}
	case "upload_voice", "chatActionUploadingVoiceNote":
		action = &types.ChatAction{ChatActionUploadingVoiceNote: &types.ChatActionUploadingVoiceNote{Progress: 50}}
	case "upload_document", "chatActionUploadingDocument":
		action = &types.ChatAction{ChatActionUploadingDocument: &types.ChatActionUploadingDocument{Progress: 50}}
	case "choose_sticker", "chatActionChoosingSticker":
		action = &types.ChatAction{ChatActionChoosingSticker: &types.ChatActionChoosingSticker{}}
	case "find_location", "chatActionChoosingLocation":
		action = &types.ChatAction{ChatActionChoosingLocation: &types.ChatActionChoosingLocation{}}
	case "record_video_note", "chatActionRecordingVideoNote":
		action = &types.ChatAction{ChatActionRecordingVideoNote: &types.ChatActionRecordingVideoNote{}}
	case "upload_video_note", "chatActionUploadingVideoNote":
		action = &types.ChatAction{ChatActionUploadingVideoNote: &types.ChatActionUploadingVideoNote{Progress: 50}}
	case "cancel", "chatActionCancel":
		action = &types.ChatAction{ChatActionCancel: &types.ChatActionCancel{}}
	default:
		action = &types.ChatAction{ChatActionTyping: &types.ChatActionTyping{}}
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
		cancelAction := &types.ChatAction{ChatActionCancel: &types.ChatActionCancel{}}
		c.client.SendChatAction(c.chatID, c.topicID, "", &types.SendChatActionOpts{Action: cancelAction})
	}
}

func (c *ChatAction) send() {
	c.client.SendChatAction(c.chatID, c.topicID, "", &types.SendChatActionOpts{Action: c.action})
}
