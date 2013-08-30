package flute

import (
	paramsJson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mime/multipart"
)

// controller 接口, 和rails基本类似
type ControllerInterface interface {
	Init(context *Context)
	Index()
	Show()
	Create()
	Update()
	Delete()
}

// controller 基类
type Controller struct {
	*Context
	vars map[string]string
}

func (controller *Controller) Init(context *Context) {
	controller.Context = context
	controller.vars = mux.Vars(controller.Request)
}

// 获取地址上的参数
func (controller *Controller) GetVars(key string) string {
	return controller.vars[key]
}

// 获取表单的内容
func (controller *Controller) ParamsFrom(key string) string {
	return controller.Request.FormValue(key)
}

// 获取文件
func (controller *Controller) ParamsFile(key string) (f multipart.File, h *multipart.FileHeader, e error) {
	f, h, e = controller.Request.FormFile(key)
	return
}

// 获取body内容, 并解析成json, 详细看github.com/bitly/go-simplejson的用法
func (controller *Controller) ParamsJson() (json *paramsJson.Json, e error) {
	body, e := ioutil.ReadAll(controller.Request.Body)
	if e != nil {
		return
	}
	json, e = paramsJson.NewJson(body)
	return
}

func (controller *Controller) Index() {}

func (controller *Controller) Show() {}

func (controller *Controller) Create() {}

func (controller *Controller) Update() {}

func (controller *Controller) Delete() {}
