package ext

import (
	"log"
	"runtime/debug"
	"sort"
	"sync"

	"github.com/AshokShau/gotdbot"
)

type Dispatcher struct {
	Client *gotdbot.Client

	// Handlers are grouped by integer index. Low indices are processed first.
	handlers map[int][]Handler
	// keys stores sorted keys for iteration
	groups []int
	mu     sync.RWMutex

	// PanicHandler handles panics during update processing.
	PanicHandler func(c *Context, r interface{})
}

func NewDispatcher(client *gotdbot.Client) *Dispatcher {
	return &Dispatcher{
		Client:   client,
		handlers: make(map[int][]Handler),
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

// ProcessUpdate processes a single update through the handlers.
func (d *Dispatcher) ProcessUpdate(update gotdbot.TlObject) {
	ctx := NewContext(d.Client, update)

	defer func() {
		if r := recover(); r != nil {
			if d.PanicHandler != nil {
				d.PanicHandler(ctx, r)
			} else {
				log.Printf("Panic in dispatcher: %v\n%s", r, debug.Stack())
			}
		}
	}()

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
