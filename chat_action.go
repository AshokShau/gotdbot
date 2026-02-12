package gotdbot

import (
	"fmt"
	"sync"
	"time"
)

// ChatActionSender is a context manager for sending chat actions.
type ChatActionSender struct {
	client   *Client
	chatId   int64
	action   ChatAction
	topicId  MessageTopic
	stopChan chan struct{}
	mu       sync.Mutex
}

// NewChatAction creates a new ChatActionSender instance.
func NewChatAction(client *Client, chatId int64, action string, topicId MessageTopic) (*ChatActionSender, error) {
	ca := &ChatActionSender{
		client:   client,
		chatId:   chatId,
		topicId:  topicId,
		stopChan: make(chan struct{}),
	}

	if err := ca.setAction(action); err != nil {
		return nil, err
	}

	return ca, nil
}

// Send sends the action once.
func (ca *ChatActionSender) Send() {
	ca.sendAction()
}

// Start sends the action immediately and starts a loop to resend it every 4 seconds.
func (ca *ChatActionSender) Start() {
	ca.sendAction()
	go ca.loopAction()
}

// Stop stops the action loop and sends a cancel action.
func (ca *ChatActionSender) Stop() {
	close(ca.stopChan)
	ca.setAction("cancel")
	ca.sendAction()
}

func (ca *ChatActionSender) sendAction() {
	ca.mu.Lock()
	action := ca.action
	ca.mu.Unlock()
	ca.client.SendChatAction("", ca.chatId, &SendChatActionOpts{Action: action, TopicId: ca.topicId})
}

func (ca *ChatActionSender) setAction(action string) error {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	switch action {
	case "typing", "chatActionTyping":
		ca.action = &ChatActionTyping{}
	case "upload_photo", "chatActionUploadingPhoto":
		ca.action = &ChatActionUploadingPhoto{}
	case "record_video", "chatActionRecordingVideo":
		ca.action = &ChatActionRecordingVideo{}
	case "upload_video", "chatActionUploadingVideo":
		ca.action = &ChatActionUploadingVideo{}
	case "record_voice", "chatActionRecordingVoiceNote":
		ca.action = &ChatActionRecordingVoiceNote{}
	case "upload_voice", "chatActionUploadingVoiceNote":
		ca.action = &ChatActionUploadingVoiceNote{}
	case "upload_document", "chatActionUploadingDocument":
		ca.action = &ChatActionUploadingDocument{}
	case "choose_sticker", "chatActionChoosingSticker":
		ca.action = &ChatActionChoosingSticker{}
	case "find_location", "chatActionChoosingLocation":
		ca.action = &ChatActionChoosingLocation{}
	case "record_video_note", "chatActionRecordingVideoNote":
		ca.action = &ChatActionRecordingVideoNote{}
	case "upload_video_note", "chatActionUploadingVideoNote":
		ca.action = &ChatActionUploadingVideoNote{}
	case "cancel", "chatActionCancel":
		ca.action = &ChatActionCancel{}
	default:
		return fmt.Errorf("unknown action type %s", action)
	}
	return nil
}

func (ca *ChatActionSender) loopAction() {
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ca.stopChan:
			return
		case <-ticker.C:
			ca.sendAction()
		}
	}
}
