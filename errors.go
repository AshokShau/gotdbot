package gotdbot

import (
	"errors"
)

var (
	// EndGroups stops processing all remaining groups.
	EndGroups = errors.New("group iteration ended")
	// ContinueGroups skips remaining handlers in the current group and continues with the next group.
	ContinueGroups = errors.New("group iteration continued")
	// ContinueHandlers continues processing the remaining handlers in the same group.
	ContinueHandlers = errors.New("handler iteration continued")

	ConversationCancelled = errors.New("conversation cancelled")
	ConversationTimeout   = errors.New("conversation wait timeout")

	SendTimeout         = errors.New("TDLib send timeout")
	ErrClientClosed     = errors.New("TDLib client closed")
	WaitPremiumPurchase = errors.New("account requires Telegram Premium")

	// ErrNoFormattedText is returned when a message content does not contain text or caption.
	ErrNoFormattedText = errors.New("message does not contain formatted text or caption")

	// ErrNoRichMessage is returned when a message doesn't contain rich content.
	ErrNoRichMessage = errors.New("message does not contain a rich message")

	// ErrNilUpdate is returned when an action is called on a nil update.
	ErrNilUpdate = errors.New("telegram update is nil")

	// ErrUpdateHasNoMessage is returned when an update has no single message to use.
	ErrUpdateHasNoMessage = errors.New("update does not identify a single message")

	// ErrNoMedia is returned when a message has no downloadable media file.
	ErrNoMedia = errors.New("message has no downloadable media")
)
