package gotdbot

import (
	"context"
	"fmt"
	"time"
)

type waiter struct {
	filter   func(client *Client, update TlObject) bool
	response chan TlObject
}

func (c *Client) snapshotWaiters(update TlObject) []*waiter {
	c.wMu.RLock()
	defer c.wMu.RUnlock()
	if len(c.waiters) == 0 {
		return nil
	}
	global := c.waiters[0]
	chatID := c.UpdateChatID(update)
	var chat map[string]*waiter
	if chatID != 0 {
		chat = c.waiters[chatID]
	}
	candidates := make([]*waiter, 0, len(global)+len(chat))
	for _, w := range global {
		candidates = append(candidates, w)
	}
	for _, w := range chat {
		candidates = append(candidates, w)
	}
	return candidates
}

// WaitFor registers a waiter and blocks until a matching update arrives or timeout occurs.
func (c *Client) WaitFor(filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	return c.WaitForChat(0, filter, timeout)
}

// WaitForContext registers a waiter and blocks until a matching update arrives,
// timeout occurs, or ctx is cancelled.
func (c *Client) WaitForContext(ctx context.Context, chatId int64, filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	if ctx == nil || filter == nil {
		return nil, fmt.Errorf("context and filter required")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c.closing.Load() {
		return nil, ErrClientClosed
	}
	ch := make(chan TlObject, 1)
	idNum := c.waiterCount.Add(1)
	id := fmt.Sprintf("%d", idNum)

	c.wMu.Lock()
	if c.waiters == nil {
		c.waiters = make(map[int64]map[string]*waiter)
	}
	if c.waiters[chatId] == nil {
		c.waiters[chatId] = make(map[string]*waiter)
	}
	c.waiters[chatId][id] = &waiter{filter: filter, response: ch}
	c.wMu.Unlock()

	defer func() {
		c.wMu.Lock()
		delete(c.waiters[chatId], id)
		if len(c.waiters[chatId]) == 0 {
			delete(c.waiters, chatId)
		}
		c.wMu.Unlock()
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case u := <-ch:
		return u, nil
	case <-c.closed:
		return nil, ErrClientClosed
	case <-c.stop:
		return nil, ErrClientClosed
	case <-timer.C:
		return nil, ConversationTimeout
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// WaitForChat registers a waiter for a specific chat and blocks until a matching
// update arrives or timeout occurs.
func (c *Client) WaitForChat(chatId int64, filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	return c.WaitForContext(context.Background(), chatId, filter, timeout)
}
