package flute

import (
	"github.com/gorilla/mux"
	"net/http"
)

// 内部路由
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

// 内部的restful路由实现, 还有中间件的调用, 总是感觉不优雅, 如果你有更好的思路,请告诉我,谢谢
func (r *Router) Resources(path string, controller ControllerInterface, middlewares []MiddlewareInterface) {

	path1 := "/" + path
	path2 := path1 + "/{id}"

	beforeFunc := func(c *Context) bool {
		for _, v := range middlewares {
			v.Init(c)
			if v.Before() == false {
				return false
			}
		}
		return true
	}

	afterFunc := func(c *Context) {
		for _, v := range middlewares {
			v.Init(c)
			v.After()
		}
	}

	r.AddFunc(path1, "GET", func(c *Context) {
		if beforeFunc(c) {
			controller.Init(c)
			controller.Index()
		}
		afterFunc(c)

	})

	r.AddFunc(path2, "GET", func(c *Context) {
		if beforeFunc(c) {
			controller.Init(c)
			controller.Show()
		}
		afterFunc(c)
	})

	r.AddFunc(path1, "POST", func(c *Context) {
		if beforeFunc(c) {
			controller.Init(c)
			controller.Create()
		}
		afterFunc(c)
	})

	r.AddFunc(path2, "PUT", func(c *Context) {
		if beforeFunc(c) {
			controller.Init(c)
			controller.Update()
		}
		afterFunc(c)
	})

	r.AddFunc(path2, "DELETE", func(c *Context) {
		if beforeFunc(c) {
			controller.Init(c)
			controller.Delete()
		}
		afterFunc(c)
	})

}
