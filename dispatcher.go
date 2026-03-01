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
	Filter UpdateFilter
	C      chan TlObject
}

type Dispatcher struct {
	Client *Client

	// Handlers are grouped by integer index. Low indices are processed first.
	handlers map[int][]Handler
	// keys stores sorted keys for iteration
	groups []int
	mu     sync.RWMutex

	// Initializers are handlers that run before grouped handlers.
	Initializers []Handler
	// Finalizers are handlers that run after grouped handlers.
	Finalizers []Handler

	// Waiters
	waiters     map[string]*Waiter
	waiterCount int64
	wMu         sync.RWMutex

	// PanicHandler handles panics during update processing.
	PanicHandler func(c *Context, r interface{})

	// ErrorHandler handles errors returned by handlers.
	ErrorHandler func(client *Client, ctx *Context, err error) error
}

type DispatcherOpts struct {
	PanicHandler func(c *Context, r interface{})
	ErrorHandler func(client *Client, ctx *Context, err error) error
}

func NewDispatcher(client *Client, opts *DispatcherOpts) *Dispatcher {
	d := &Dispatcher{
		Client:   client,
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

// AddInitializer adds a handler to the initializers list.
func (d *Dispatcher) AddInitializer(h Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Initializers = append(d.Initializers, h)
}

// AddFinalizer adds a handler to the finalizers list.
func (d *Dispatcher) AddFinalizer(h Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Finalizers = append(d.Finalizers, h)
}

// WaitFor registers a waiter and blocks until a matching update arrives or timeout occurs.
func (d *Dispatcher) WaitFor(filter UpdateFilter, timeout time.Duration) (TlObject, error) {
	ch := make(chan TlObject, 1)
	idNum := atomic.AddInt64(&d.waiterCount, 1)
	id := fmt.Sprintf("%d", idNum)

	d.wMu.Lock()
	d.waiters[id] = &Waiter{Filter: filter, C: ch}
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

// ProcessUpdate processes a single update through the handlers.
// It launches a goroutine to prevent blocking the main update loop.
func (d *Dispatcher) ProcessUpdate(update TlObject) {
	go func() {
		ctx := NewContext(update, d)

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
			if w.Filter(d.Client, ctx) {
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
		defer d.mu.RUnlock()

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
				return d.ErrorHandler(d.Client, ctx, err)
			}
			log.Printf("Handler error: %v", err)
			return nil
		}

		// Initializers
		for _, h := range d.Initializers {
			if h.CheckUpdate(d.Client, ctx) {
				if err := h.HandleUpdate(d.Client, ctx); err != nil {
					action := handleError(err)
					if errors.Is(action, EndGroups) {
						return
					}
				}
			}
		}

		// Grouped Handlers
		for _, group := range d.groups {
			groupHandlers := d.handlers[group]
			for _, h := range groupHandlers {
				if h.CheckUpdate(d.Client, ctx) {
					err := h.HandleUpdate(d.Client, ctx)
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

		// Finalizers
		for _, h := range d.Finalizers {
			if h.CheckUpdate(d.Client, ctx) {
				if err := h.HandleUpdate(d.Client, ctx); err != nil {
					action := handleError(err)
					if errors.Is(action, EndGroups) {
						return
					}
				}
			}
		}
	}()
}
