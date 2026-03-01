package gotdbot

import (
	"errors"
)

var (
	// EndGroups stops the iteration of subsequent groups.
	EndGroups = errors.New("group iteration ended")
	// ContinueGroups continues iterating over the current group.
	ContinueGroups = errors.New("group iteration continued")

	ConversationCancelled = errors.New("conversation cancelled")
	ConversationTimeout   = errors.New("conversation wait timeout")

	SendTimeout         = errors.New("TDLib send timeout")
	WaitPremiumPurchase = errors.New("account requires Telegram Premium")
)
