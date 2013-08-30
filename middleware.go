package flute

type Middleware struct {
	*Context
}

func (plugin *Middleware) Before() bool { return true }
func (plugin *Middleware) After()       {}
func (plugin *Middleware) Init(c *Context) {
	plugin.Context = c
}

type MiddlewareInterface interface {
	Init(context *Context)
	Before() bool
	After()
}
