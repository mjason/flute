package flute

// 中间件类, 是一个类似controller的东西, 可以用户在发生组restful路由的时候进行定制
// 比如插入日子, 进行权限判断
type Middleware struct {
	*Context
}

// 路由发生前发生的操作
func (plugin *Middleware) Before() bool { return true }

// 路由结束前发生的操作
func (plugin *Middleware) After() {}

// 初始化
func (plugin *Middleware) Init(c *Context) {
	plugin.Context = c
}

type MiddlewareInterface interface {
	Init(context *Context)
	Before() bool
	After()
}
