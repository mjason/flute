package flute

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	r := Router{mux.NewRouter()}
	return &r
}

func (r *Router) Start(addr string) error {
	http.Handle("/", r)
	return http.ListenAndServe(addr, nil)
}

func (r *Router) AddFunc(path string, method string, f func(c *Context)) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context := Context{w, r}
		f(&context)
	}).Methods(method)
}

func (r *Router) Resources(path string, controller ControllerInterface) {
	path1 = "/" + path
	path2 = path1 + "/{id}"
	r.AddFunc(path1, "GET", func(c *Context) {
		controller.Init(c)
		controller.Index()
	})
	r.AddFunc(path2, "GET", func(c *Context) {
		controller.Init(c)
		controller.Show()
	})
	r.AddFunc(path1, "POST", func(c *Context) {
		controller.Init(c)
		controller.Create()
	})
	r.AddFunc(path2, "PUT", func(c *Context) {
		controller.Init(c)
		controller.Update()
	})
	r.AddFunc(path2, "DELETE", func(c *Context) {
		controller.Init(c)
		controller.Delete()
	})
}
