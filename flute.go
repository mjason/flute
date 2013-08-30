package flute

var Flute *Router

func init() {
	Flute = NewRouter()
}

func Start(addr string) error {
	return Flute.Start(addr)
}

func HandleFunc(path string, method string, f func(c *Context)) {
	Flute.AddFunc(path, method, f)
}

func Resources(path string, controller ControllerInterface) {
	Flute.Resources(path, controller)
}
