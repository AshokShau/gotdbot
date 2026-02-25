package gotdbot

// HandlerFunc is a function that handles updates.
type HandlerFunc func(client *Client, update TlObject) error

// LegacyHandler represents an update handler.
// Deprecated: Use Dispatcher and its handlers instead.
type LegacyHandler struct {
	Func       HandlerFunc
	UpdateType string
	Position   int
}

func (h *LegacyHandler) CheckUpdate(ctx *Context) bool {
	return ctx.RawUpdate.Type() == h.UpdateType
}

func (h *LegacyHandler) HandleUpdate(ctx *Context) error {
	return h.Func(ctx.Client, ctx.RawUpdate)
}
