package gotdbot

import (
	"context"
	"errors"
	"fmt"
)

type pendingCall struct {
	response    chan TlObject
	messages    map[[2]int64]chan TlObject
	waitForSend bool
}

// Only outgoing sends await updateMessageSendSucceeded/Failed. Reading or
// editing a pending message must not wait for somebody else's send to finish.
func waitsForSend(req TlObject) bool {
	switch req.(type) {
	case *SendMessage, *SendMessageAlbum, *SendEphemeralMessage, *ForwardMessages,
		*SendBotStartMessage, *SendInlineQueryResultMessage, *SendQuickReplyShortcutMessages,
		*ResendMessages, *AddOffer:
		return true
	default:
		return false
	}
}

func (c *Client) deliverResponse(extra string, obj TlObject) {
	c.requestMu.Lock()
	defer c.requestMu.Unlock()
	p := c.pendingRequests[extra]
	if p == nil {
		return // The caller already timed out or cancelled.
	}

	if p.waitForSend {
		p.waitForSend = false
		p.messages = make(map[[2]int64]chan TlObject)
		register := func(m *Message) {
			if m == nil {
				return
			}
			if _, pending := m.SendingState.(*MessageSendingStatePending); !pending {
				return
			}
			key := [2]int64{m.ChatId, m.Id}
			ch := make(chan TlObject, 1)
			p.messages[key] = ch
			c.pendingMessages[key] = ch
		}
		switch r := obj.(type) {
		case *Message:
			register(r)
		case *Messages:
			if r != nil {
				for i := range r.Messages {
					register(&r.Messages[i])
				}
			}
		}
	}
	// Registration precedes delivery, so even an immediate completion is kept.
	select {
	case p.response <- obj:
	default:
	}
}

func (c *Client) finishCall(extra string) {
	c.requestMu.Lock()
	defer c.requestMu.Unlock()
	if p := c.pendingRequests[extra]; p != nil {
		delete(c.pendingRequests, extra)
		for key := range p.messages {
			delete(c.pendingMessages, key)
		}
	}
}

func (c *Client) completeMessage(chatID, oldID int64, update TlObject) {
	key := [2]int64{chatID, oldID}
	c.requestMu.Lock()
	ch := c.pendingMessages[key]
	delete(c.pendingMessages, key)
	c.requestMu.Unlock()
	if ch != nil {
		ch <- update
	}
}

func (c *Client) waitSend(ctx context.Context, result TlObject, p *pendingCall) (TlObject, error) {
	wait := func(m *Message) (*Message, error) {
		if m == nil {
			return nil, fmt.Errorf("nil message response")
		}
		if failed, ok := m.SendingState.(*MessageSendingStateFailed); ok {
			return m, failed.Error
		}
		if _, pending := m.SendingState.(*MessageSendingStatePending); !pending {
			return m, nil
		}
		select {
		case u := <-p.messages[[2]int64{m.ChatId, m.Id}]:
			switch completion := u.(type) {
			case *UpdateMessageSendSucceeded:
				return completion.Message, nil
			case *UpdateMessageSendFailed:
				return completion.Message, completion.Error
			default:
				return m, fmt.Errorf("unexpected send completion %T", u)
			}
		case <-ctx.Done():
			return m, ctx.Err()
		case <-c.closed:
			return m, ErrClientClosed
		case <-c.stop:
			return m, ErrClientClosed
		}
	}
	switch r := result.(type) {
	case *Message:
		return wait(r)
	case *Messages:
		if r == nil {
			return nil, fmt.Errorf("nil messages response")
		}
		var firstError error
		for i := range r.Messages {
			m, err := wait(&r.Messages[i])
			if m != nil {
				r.Messages[i] = *m
			}
			if firstError == nil {
				firstError = err
			}
			if ctx.Err() != nil || c.closing.Load() {
				return r, errors.Join(firstError, err)
			}
		}
		return r, firstError
	default:
		return result, nil
	}
}

func responseAs[T TlObject](obj TlObject) (T, error) {
	var zero T
	v, ok := obj.(T)
	if !ok {
		return zero, fmt.Errorf("unexpected TDLib response %T", obj)
	}
	return v, nil
}
