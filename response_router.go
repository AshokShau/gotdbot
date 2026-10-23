package gotdbot

import (
	"errors"
	"strings"
	"time"
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
			if c.me != nil && u.User.Id == c.me.Id {
				c.me = u.User
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
			c.Logger.Debug("TDLib client closed", "user_id", c.Me().Id)
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
	// Track membership and filter before concurrent handlers run.
	if c.botToken != "" && !c.config.ReceiveOldMessages && !c.acceptBotUpdate(obj, time.Now()) {
		return
	}
	if c.config.DropPendingUpdates && !c.acceptPendingUpdate(obj) {
		return
	}
	if err := c.updates.push(obj); err != nil && !errors.Is(err, ErrClientClosed) {
		c.Logger.Error("Cannot queue TDLib update", "type", obj.GetType(), "error", err)
	}
}
