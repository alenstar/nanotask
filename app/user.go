package app

import (
	_ "github.com/alenstar/nanoweb/config"
	"github.com/alenstar/nanoweb/controller"
	"github.com/alenstar/nanoweb/log"
	_ "github.com/alenstar/nanoweb/modules"
)

type UserController struct {
	controller.CommonController
}

func (u *UserController) Get() {
	if id, ok := u.Ctx.Params["id"]; ok {
		log.Info("User ", id)
	} else {
		log.Debug("User Get bad id", id)
	}
}

func (u *UserController) Post() {
	log.Debug("User Post")
}

func (u *UserController) Delete() {
	log.Debug("User Delete")
}

func (u *UserController) Put() {
	log.Debug("User Put", u.Ctx.Params)
}
