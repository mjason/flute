package flute

import (
	"github.com/robfig/config"
	"os"
)

var Flute *Router
var Message map[string]string
var Config *config.Config

func init() {
	Flute = NewRouter()
	Message = make(map[string]string)
}

// 启动服务器, 需要传入一个地址
func Start() error {
	addr, e := Config.String(Message["env"], "addr")
	if e != nil {
		return e
	}
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

// 加载配置文件到Message里面去
func LoadConfig(path string) error {
	if env := os.Getenv("FLUTEENV"); env != "" {
		Message["env"] = env
	} else {
		Message["env"] = "development"
	}
	c, e := config.ReadDefault(path)
	if e != nil {
		return e
	}
	Config = c
	return nil
}
