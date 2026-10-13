package gotdbot

import (
	"errors"
	"strings"
)

// Protocol completion must not depend on an available application handler.
func (c *Client) acceptResponse(obj TlObject, extra string, managerStop <-chan struct{}) {
	if obj == nil {
		return
	}
	if extra != "" {
		c.deliverResponse(extra, obj)
		return
	}
	switch u := obj.(type) {
	case *UpdateUser:
		if u == nil {
			return
		}
		if u.User != nil {
			c.meMu.Lock()
			if c.Me != nil && u.User.Id == c.Me.Id {
				c.Me = u.User
			}
			c.meMu.Unlock()
		}
	case *UpdateMessageSendSucceeded:
		if u.Message != nil {
			c.completeMessage(u.Message.ChatId, u.OldMessageId, u)
		}
	case *UpdateMessageSendFailed:
		if u.Message != nil {
			c.completeMessage(u.Message.ChatId, u.OldMessageId, u)
		}
	case *UpdateAuthorizationState:
		if _, ok := u.AuthorizationState.(*AuthorizationStateClosed); ok {
			c.closing.Store(true)
			if !c.isAuthorized.Load() {
				c.sendAuthError(ErrClientClosed)
			}
			c.closedOnce.Do(func() { close(c.closed) })
		}
	}

	if !strings.HasPrefix(obj.GetType(), "update") {
		if err, ok := obj.(*Error); ok {
			c.Logger.Error("Uncorrelated TDLib error", "error", err)
		}
		return
	}
	select {
	case <-c.stop:
		return
	case <-managerStop:
		return
	default:
	}
	if err := c.updates.push(obj); err != nil && !errors.Is(err, ErrClientClosed) {
		c.Logger.Error("Cannot queue TDLib update", "type", obj.GetType(), "error", err)
	}
}
