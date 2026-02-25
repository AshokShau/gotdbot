package gotdbot

// UpdateFilter is a function that returns true if the update matches the filter.
type UpdateFilter func(ctx *Context) bool

type Handler interface {
	CheckUpdate(ctx *Context) bool
	HandleUpdate(ctx *Context) error
}
