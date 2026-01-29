package client

import "github.com/AshokShau/gotdbot/types"

// FilterFunc is a function that filters updates.
type FilterFunc func(types.TlObject) bool

// HandlerFunc is a function that handles updates.
type HandlerFunc func(client *Client, update types.TlObject) error

// Handler represents an update handler.
type Handler struct {
	Func       HandlerFunc
	UpdateType string
	Filter     FilterFunc
	Position   int
}
