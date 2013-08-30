package flute

import (
	paramsJson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mime/multipart"
)

type ControllerInterface interface {
	Init(context *Context)
	Index()
	Show()
	Create()
	Update()
	Delete()
}

type Controller struct {
	*Context
	vars map[string]string
}

func (controller *Controller) Init(context *Context) {
	controller.Context = context
	controller.vars = mux.Vars(controller.Request)
}

func (controller *Controller) GetVars(key string) string {
	return controller.vars[key]
}

func (controller *Controller) ParamsFrom(key string) string {
	return controller.Request.FormValue(key)
}

func (controller *Controller) ParamsFile(key string) (f multipart.File, h *multipart.FileHeader, e error) {
	f, h, e = controller.Request.FormFile(key)
	return
}

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
