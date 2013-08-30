package flute

var Flute *Router

func init() {
	Flute = NewRouter()
}

// 启动服务器, 需要传入一个地址
func Start(addr string) error {
	return Flute.Start(addr)
}

// 基础路由
func HandleFunc(path string, method string, f func(c *Context)) {
	Flute.AddFunc(path, method, f)
}

// restful 路由风格
func Resources(path string, controller ControllerInterface, middlewares ...MiddlewareInterface) {
	Flute.Resources(path, controller, middlewares)
}
