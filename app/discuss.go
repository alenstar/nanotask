package app

import (
	_ "github.com/alenstar/nanoweb/config"
	"github.com/alenstar/nanoweb/controller"
	"github.com/alenstar/nanoweb/log"
	_ "github.com/alenstar/nanoweb/modules"
)

type DiscussController struct {
	controller.CommonController
}

func (d *DiscussController) Get() {
	log.Debug("Discuss Get")
}

func (d *DiscussController) Post() {
	log.Debug("Discuss Post")
}

func (d *DiscussController) Delete() {
	log.Debug("Discuss Delete")
}

func (d *DiscussController) Put() {
	log.Debug("Discuss Put", d.Ctx.Params)
}
