package gotdbot

// FilterFunc is a function that filters updates.
type FilterFunc func(TlObject) bool

// HandlerFunc is a function that handles updates.
type HandlerFunc func(client *Client, update TlObject) error

// Handler represents an update handler.
type Handler struct {
	Func       HandlerFunc
	UpdateType string
	Filter     FilterFunc
	Position   int
}
