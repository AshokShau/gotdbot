package gotdbot

// HandlerFunc is a function that handles updates.
type HandlerFunc func(client *Client, update TlObject) error

// Handler represents an update handler.
type Handler struct {
	Func       HandlerFunc
	UpdateType string
	Position   int
}
