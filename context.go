package flute

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) Vars() map[string]string {
	return mux.Vars(c.Request)
}

func (c *Context) RenderText(status int, data string) {
	c.ResponseWriter.WriteHeader(status)
	c.ResponseWriter.Write([]byte(data))
}
