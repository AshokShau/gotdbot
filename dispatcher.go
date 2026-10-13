package gotdbot

import (
	"errors"
	"time"
)

func (c *Client) processor() {
	defer c.wg.Done()
	for {
		update, err := c.updates.next(c.stop)
		if errors.Is(err, ErrClientClosed) {
			return
		}
		if err != nil {
			c.Logger.Error("Cannot read queued TDLib update", "error", err)
			if update == nil {
				timer := time.NewTimer(time.Second)
				select {
				case <-c.stop:
					timer.Stop()
					return
				case <-timer.C:
				}
				continue
			}
		}
		select {
		case <-c.stop:
			return
		default:
		}
		func() {
			defer func() {
				if r := recover(); r != nil {
					if c.panicHandler != nil {
						c.panicHandler(c, update, r)
					} else {
						c.Logger.Error("Panic in handler", "panic", r)
					}
				}
			}()

			// Internal state handling is independent of application handler groups.
			switch u := update.(type) {
			case *UpdateAuthorizationState:
				if u != nil && u.AuthorizationState != nil {
					c.authHandler(u)
				}
			case *UpdateConnectionState:
				if u != nil && u.State != nil {
					c.connectionStateHandler(u)
				}
			}

			// Snapshot waiters before invoking application filters. Filters may
			// register/remove waiters and must never execute under wMu.
			c.wMu.RLock()
			var candidates []*waiter
			chatID := c.UpdateChatID(update)
			for id, inner := range c.waiters {
				if id == 0 || id == chatID {
					for _, w := range inner {
						candidates = append(candidates, w)
					}
				}
			}
			c.wMu.RUnlock()
			for _, w := range candidates {
				if w.filter(c, update) {
					select {
					case w.response <- update:
					default:
					}
				}
			}

			hData := c.handlers.Load()
			if hData == nil {
				return
			}
			handlersMap := hData.handlers
			groups := hData.groups

			handleError := func(err error) error {
				if err == nil {
					return nil
				}
				if errors.Is(err, EndGroups) || errors.Is(err, ContinueGroups) || errors.Is(err, ContinueHandlers) {
					return err
				}
				if c.errorHandler != nil {
					return c.errorHandler(c, update, err)
				}
				c.Logger.Error("Handler error", "error", err, "type", update.GetType())
				return nil
			}

			for _, group := range groups {
				groupHandlers := handlersMap[group]
				for _, h := range groupHandlers {
					if h.CheckUpdate(c, update) {
						err := h.HandleUpdate(c, update)
						action := handleError(err)

						if errors.Is(action, EndGroups) {
							return
						}
						if errors.Is(action, ContinueGroups) {
							break // Move to next group
						}
						if errors.Is(action, ContinueHandlers) {
							continue // Move to next handler in same group
						}
						// Run at most one matching handler in each group.
						break
					}
				}
			}
		}()
	}
}

type handlersData struct {
	handlers map[int][]Handler
	groups   []int
}
