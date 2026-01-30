package ext

type Handler interface {
	CheckUpdate(ctx *Context) bool
	HandleUpdate(ctx *Context) error
}
