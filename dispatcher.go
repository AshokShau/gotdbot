package gotdbot

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type Waiter struct {
	Client *Client
	Filter UpdateFilter
	C      chan TlObject
}

type Dispatcher struct {
	// Handlers are grouped by integer index. Low indices are processed first.
	handlers map[int][]Handler
	// keys stores sorted keys for iteration
	groups []int
	mu     sync.RWMutex

	// Waiters
	waiters     map[string]*Waiter
	waiterCount int64
	wMu         sync.RWMutex

	// PanicHandler handles panics during update processing.
	PanicHandler func(c *Context, r interface{})

	// ErrorHandler handles errors returned by handlers.
	// It should only return EndGroups, ContinueGroups, or nil.
	// Any other error returned will be logged and treated as nil (success).
	ErrorHandler func(client *Client, ctx *Context, err error) error
}

type DispatcherOpts struct {
	PanicHandler func(c *Context, r interface{})
	// ErrorHandler handles errors returned by handlers.
	// It should only return EndGroups, ContinueGroups, or nil.
	// Any other error returned will be logged and treated as nil (success).
	ErrorHandler func(client *Client, ctx *Context, err error) error
}

func NewDispatcher(opts *DispatcherOpts) *Dispatcher {
	d := &Dispatcher{
		handlers: make(map[int][]Handler),
		waiters:  make(map[string]*Waiter),
	}
	if opts != nil {
		d.PanicHandler = opts.PanicHandler
		d.ErrorHandler = opts.ErrorHandler
	}
	return d
}

// AddHandler adds a handler to group 0.
func (d *Dispatcher) AddHandler(h Handler) {
	d.AddHandlerToGroup(h, 0)
}

// AddHandlerToGroup adds a handler to a specific group.
func (d *Dispatcher) AddHandlerToGroup(h Handler, group int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.handlers[group]; !ok {
		d.handlers[group] = []Handler{}
		d.groups = append(d.groups, group)
		sort.Ints(d.groups)
	}
	d.handlers[group] = append(d.handlers[group], h)
}

// WaitFor registers a waiter for a specific client and blocks until a matching update arrives or timeout occurs.
func (d *Dispatcher) WaitFor(client *Client, filter UpdateFilter, timeout time.Duration) (TlObject, error) {
	ch := make(chan TlObject, 1)
	idNum := atomic.AddInt64(&d.waiterCount, 1)
	id := fmt.Sprintf("%d", idNum)

	d.wMu.Lock()
	d.waiters[id] = &Waiter{Client: client, Filter: filter, C: ch}
	d.wMu.Unlock()

	defer func() {
		d.wMu.Lock()
		delete(d.waiters, id)
		d.wMu.Unlock()
	}()

	select {
	case u := <-ch:
		return u, nil
	case <-time.After(timeout):
		return nil, ConversationTimeout
	}
}

// ProcessUpdate processes a single update through the handlers for a specific client.
// It launches a goroutine to prevent blocking the main update loop.
func (d *Dispatcher) ProcessUpdate(client *Client, update TlObject) {
	go func() {
		ctx := NewContext(client, update, d)

		defer func() {
			if r := recover(); r != nil {
				if d.PanicHandler != nil {
					d.PanicHandler(ctx, r)
				} else {
					log.Printf("Panic in dispatcher: %v\n%s", r, debug.Stack())
				}
			}
		}()

		d.wMu.RLock()
		var matchedWaiters []*Waiter
		for _, w := range d.waiters {
			if w.Client == client && w.Filter(client, ctx) {
				matchedWaiters = append(matchedWaiters, w)
			}
		}
		d.wMu.RUnlock()

		for _, w := range matchedWaiters {
			select {
			case w.C <- update:
			default:
			}
		}

		d.mu.RLock()
		groups := make([]int, len(d.groups))
		copy(groups, d.groups)
		handlers := make(map[int][]Handler, len(groups))
		for _, g := range groups {
			h := make([]Handler, len(d.handlers[g]))
			copy(h, d.handlers[g])
			handlers[g] = h
		}
		d.mu.RUnlock()

		handleError := func(err error) error {
			if err == nil {
				return nil
			}
			if errors.Is(err, EndGroups) {
				return EndGroups
			}
			if errors.Is(err, ContinueGroups) {
				return ContinueGroups
			}
			if d.ErrorHandler != nil {
				action := d.ErrorHandler(client, ctx, err)
				if action != nil && !errors.Is(action, EndGroups) && !errors.Is(action, ContinueGroups) {
					log.Printf("ErrorHandler returned unexpected error: %v", action)
					return nil
				}
				return action
			}
			log.Printf("Handler error: %v", err)
			return nil
		}

		// Grouped Handlers
		for _, group := range groups {
			groupHandlers := handlers[group]
			for _, h := range groupHandlers {
				if h.CheckUpdate(client, ctx) {
					err := h.HandleUpdate(client, ctx)
					action := handleError(err)

					if errors.Is(action, EndGroups) {
						return
					}
					if errors.Is(action, ContinueGroups) {
						continue
					}
					// If handler executed successfully (or action returned nil),
					// stop iterating this group and move to the next group.
					break
				}
			}
		}
	}()
}
