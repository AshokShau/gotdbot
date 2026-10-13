package gotdbot

import (
	"reflect"

	"github.com/AshokShau/gotdbot/filters"
)

// UpdateFilter filters any native Telegram update.
type UpdateFilter func(Update) bool

// UpdateHandler handles Telegram updates.
type UpdateHandler struct {
	Filter   UpdateFilter
	Response func(*Client, Update) error
}

// NewUpdateHandler creates a handler for any native update. A nil filter accepts all updates.
func NewUpdateHandler(response func(*Client, Update) error, filter UpdateFilter) *UpdateHandler {
	return &UpdateHandler{Response: response, Filter: filter}
}

func nativeUpdate(object TlObject) (Update, bool) {
	update, ok := object.(Update)
	if !ok || update == nil {
		return nil, false
	}
	value := reflect.ValueOf(update)
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return nil, false
	}
	return update, true
}

// CheckUpdate reports whether this handler accepts an update.
func (h *UpdateHandler) CheckUpdate(client *Client, object TlObject) bool {
	if h == nil || h.Response == nil {
		return false
	}
	update, ok := nativeUpdate(object)
	return ok && (h.Filter == nil || h.Filter(update))
}

// HandleUpdate calls the handler for a native update.
func (h *UpdateHandler) HandleUpdate(client *Client, object TlObject) error {
	if h == nil || h.Response == nil {
		return nil
	}
	update, ok := nativeUpdate(object)
	if !ok {
		return nil
	}
	return h.Response(client, update)
}

// AddFilter combines a filter with the handler's existing filters using AND.
// Configure handlers before registration or starting the client.
func (h *UpdateHandler) AddFilter(filter UpdateFilter) *UpdateHandler {
	if filter != nil {
		h.Filter = filters.AllFilters[Update](h.Filter, filter)
	}
	return h
}

// OnUpdate registers a general update handler in group 0.
func (c *Client) OnUpdate(response func(*Client, Update) error, filter UpdateFilter) *UpdateHandler {
	return c.OnUpdateGroup(response, filter, 0)
}

// OnUpdateGroup registers a general update handler in a specific group.
func (c *Client) OnUpdateGroup(response func(*Client, Update) error, filter UpdateFilter, group int) *UpdateHandler {
	handler := NewUpdateHandler(response, filter)
	if response != nil {
		c.AddHandlerGroup(handler, group)
	}
	return handler
}
