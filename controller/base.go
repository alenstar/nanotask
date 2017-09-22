package controller

import (
	"bytes"
	"encoding/json"
	_ "errors"
	"github.com/alenstar/nanotask/context"
	"github.com/alenstar/nanotask/log"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
)

var (
	DEBUG bool
)

func init() {
	DEBUG = true
}

type IController interface {
	Get()
	Put()
	Post()
	Delete()
	Patch()
	Options()
	NotFound()
	Any()
	Head()

	Prepare()
	Finish()
	Init(*context.Context)
}

type BaseController struct {
	Ctx            *context.Context
	Render         *render.Render
	controllerType reflect.Type
}

func (b *BaseController) Init(ctx *context.Context) {
	b.Ctx = ctx
	b.Render = render.New()
}

func Create(controller IController) func(http.ResponseWriter, *http.Request) {
	return func(rsp http.ResponseWriter, req *http.Request) {
		vc := reflect.New(reflect.Indirect(reflect.ValueOf(controller)).Type())
		execController, ok := vc.Interface().(IController)
		if !ok {
			panic("controller is not IController")
		}

		ctx := context.New(rsp, req, "")

		execController.Init(ctx)
		execController.Prepare()

		switch req.Method {
		case "GET":
			execController.Get()
		case "PATCH":
			execController.Patch()
		case "POST":
			execController.Post()
		case "PUT":
			execController.Put()
		case "DELETE":
			execController.Delete()
		case "OPTIONS":
			execController.Options()
		case "HEAD":
			execController.Head()
		default:
			execController.NotFound()
		}

		execController.Finish()
	}
}

func (b *BaseController) Get() {
	log.Debug("BaseControoler.GET:", b.controllerType)
}

func (b *BaseController) Put() {

}

func (b *BaseController) Post() {

}

func (b *BaseController) Delete() {

}

func (b *BaseController) Patch() {

}

func (b *BaseController) Head() {

}

func (b *BaseController) Options() {

}

func (b *BaseController) Any() {

}

func (b *BaseController) NotFound() {

}

func (b *BaseController) Prepare() {

}

func (b *BaseController) Finish() {

}

func (b *BaseController) WriteString(str string) {
	if DEBUG {
		b.Ctx.WriteString(str + "\r\n")
	} else {
		b.Ctx.WriteString(str)
	}
}

func (b *BaseController) Write(str []byte) {
	b.Ctx.Write(str)
}

func (b *BaseController) WriteJSON(data interface{}) {
	b.Ctx.Header().Set("Content-Type", "application/json; charset=utf-8")
	b.Ctx.Header().Set("Access-Control-Allow-Origin", "*")
	var content []byte
	var err error

	if DEBUG {
		content, err = json.MarshalIndent(data, "", "  ")
		content = append(content, '\r')
		content = append(content, '\n')
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(b.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	b.Ctx.Write(content)
}

func (b *BaseController) WriteJSONP(data interface{}) {
	b.Ctx.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	b.Ctx.Header().Set("Access-Control-Allow-Origin", "*")
	var content []byte
	var err error
	if DEBUG {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(b.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	var callback string
	var ok bool
	if callback, ok = b.Ctx.Params["callback"]; ok {
		if callback == "" {
			log.Error("WriteJSONP:", `"callback" parameter required`)
			return
		}
	}
	callbackContent := bytes.NewBufferString(" " + template.JSEscapeString(callback))
	callbackContent.WriteString("(")
	callbackContent.Write(content)
	callbackContent.WriteString(");\r\n")
	b.Ctx.Write(callbackContent.Bytes())
}

func stringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}
