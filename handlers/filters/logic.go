package filters

import "github.com/AshokShau/gotdbot"

func (f Message) And(next Message) Message {
	return func(msg *gotdbot.Message) bool {
		return f(msg) && next(msg)
	}
}

func (f Message) Or(next Message) Message {
	return func(msg *gotdbot.Message) bool {
		return f(msg) || next(msg)
	}
}

func Not(f Message) Message {
	return func(msg *gotdbot.Message) bool {
		return !f(msg)
	}
}
