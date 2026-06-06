package gotdbot

// Private filters messages sent in private chats
func Private(msg *Message) bool {
	return msg.IsPrivate()
}

// Group filters messages sent in group or supergroup chats
func Group(msg *Message) bool {
	return msg.IsGroup()
}

// Channel filters messages sent in channel chats
func Channel(msg *Message) bool {
	return msg.IsChannel()
}

// Incoming filters incoming messages
func Incoming(msg *Message) bool {
	return !msg.IsOutgoingMessage()
}

// Outgoing filters outgoing messages
func Outgoing(msg *Message) bool {
	return msg.IsOutgoingMessage()
}

// Edited filters edited messages
func Edited(msg *Message) bool {
	return msg.IsEditedMessage()
}

// And combines multiple filters, all must pass
func And(f1, f2 func(msg *Message) bool) func(msg *Message) bool {
	return func(msg *Message) bool {
		return f1(msg) && f2(msg)
	}
}

// Or combines multiple filters, at least one must pass
func Or(f1, f2 func(msg *Message) bool) func(msg *Message) bool {
	return func(msg *Message) bool {
		return f1(msg) || f2(msg)
	}
}

// Not negates a filter
func Not(f func(msg *Message) bool) func(msg *Message) bool {
	return func(msg *Message) bool {
		return !f(msg)
	}
}
