package flute

import (
	"github.com/gorilla/mux"
	"net/http"
)

// 这里把http的两个请求和成一个新的数据结构,加入一些方法
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// 获取url上面的参数,比如{id}
func (c *Context) Vars() map[string]string {
	return mux.Vars(c.Request)
}

// 返回字符串
func (c *Context) RenderText(status int, data string) {
	c.ResponseWriter.WriteHeader(status)
	c.ResponseWriter.Write([]byte(data))
}
