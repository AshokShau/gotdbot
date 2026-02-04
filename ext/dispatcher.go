package ext

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/AshokShau/gotdbot"
)

var ErrWaitTimeout = errors.New("wait timeout")

type Waiter struct {
	Filter UpdateFilter
	C      chan gotdbot.TlObject
}

type Dispatcher struct {
	Client *gotdbot.Client

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
}

func NewDispatcher(client *gotdbot.Client) *Dispatcher {
	return &Dispatcher{
		Client:   client,
		handlers: make(map[int][]Handler),
		waiters:  make(map[string]*Waiter),
	}
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

// WaitFor registers a waiter and blocks until a matching update arrives or timeout occurs.
func (d *Dispatcher) WaitFor(filter UpdateFilter, timeout time.Duration) (gotdbot.TlObject, error) {
	ch := make(chan gotdbot.TlObject, 1)
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
		return nil, ErrWaitTimeout
	}
}

// ProcessUpdate processes a single update through the handlers.
func (d *Dispatcher) ProcessUpdate(update gotdbot.TlObject) {
	ctx := NewContext(d.Client, update, d)

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
		if w.Filter(ctx) {
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

	for _, group := range d.groups {
		groupHandlers := d.handlers[group]
		for _, h := range groupHandlers {
			if h.CheckUpdate(ctx) {
				err := h.HandleUpdate(ctx)
				if err != nil {
					log.Printf("Handler error: %v", err)
				}
				break
			}
		}
	}
}

// Start registers the dispatcher with the client to receive updates.
// It adds itself as an "initializer" to the client, which ensures it sees all updates.
func (d *Dispatcher) Start() {
	// The client's AddHandler for "initializer" expects a HandlerFunc and FilterFunc.
	d.Client.AddHandler("initializer", func(c *gotdbot.Client, u gotdbot.TlObject) error {
		d.ProcessUpdate(u)
		return nil
	}, nil, 0)
}
