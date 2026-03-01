package gotdbot

// UpdateFilter is a function that returns true if the update matches the filter.
type UpdateFilter func(client *Client, ctx *Context) bool

type Handler interface {
	CheckUpdate(client *Client, ctx *Context) bool
	HandleUpdate(client *Client, ctx *Context) error
}

type internalHandler struct {
	handleFunc func(client *Client, update TlObject) error
	updateType string
}

func (h *internalHandler) CheckUpdate(client *Client, ctx *Context) bool {
	return ctx.RawUpdate.Type() == h.updateType
}

func (h *internalHandler) HandleUpdate(client *Client, ctx *Context) error {
	return h.handleFunc(client, ctx.RawUpdate)
}
