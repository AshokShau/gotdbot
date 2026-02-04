package filters

import "github.com/AshokShau/gotdbot"

type (
	Message func(msg *gotdbot.Message) bool
)
