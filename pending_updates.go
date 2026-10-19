package gotdbot

import "sync"

type pendingUpdateState struct {
	mu         sync.Mutex
	authorized bool
	connected  bool
	done       bool
}

func (c *Client) acceptPendingUpdate(update TlObject) bool {
	s := &c.startupUpdates
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done {
		return true
	}

	switch u := update.(type) {
	case *UpdateAuthorizationState:
		if u != nil {
			_, s.authorized = u.AuthorizationState.(*AuthorizationStateReady)
		}
		s.done = s.authorized && s.connected
		return true
	case *UpdateConnectionState:
		if u != nil {
			_, s.connected = u.State.(*ConnectionStateReady)
		}
		s.done = s.authorized && s.connected
		return true
	case *UpdateMessageSendSucceeded, *UpdateMessageSendFailed, *UpdateMessageSendAcknowledged:
		return true
	case *UpdateNewMessage:
		if u != nil && u.Message != nil {
			message := u.Message
			return message.IsOutgoing || message.SendingState != nil || message.SchedulingState != nil
		}
	}
	return false
}
