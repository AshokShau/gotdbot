package gotdbot

// UpdateFilter is a function that returns true if the update matches the filter.
type UpdateFilter func(client *Client, ctx *Context) bool

type Handler interface {
	CheckUpdate(client *Client, ctx *Context) bool
	HandleUpdate(client *Client, ctx *Context) error
}
