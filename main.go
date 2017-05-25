package main

import (
	"github.com/alenstar/nanoweb/app"
	_ "github.com/alenstar/nanoweb/config"
	"github.com/alenstar/nanoweb/controller"
	"github.com/alenstar/nanoweb/http/server"
	"github.com/alenstar/nanoweb/log"
	_ "github.com/alenstar/nanoweb/modules"
)

type MyController struct {
	controller.CommonController
}

func (m *MyController) Get() {
	m.Ctx.WriteString("Hello, world !\n")
}

func main() {
	s := server.New(":8888")
	log.Info("Http Server Start ... ... ", *s)
	s.UseStaticDir("/", "admin/dist/")
	s.Handle("/object", controller.Create(&MyController{}))
	s.Handle("/user", controller.Create(&app.UserController{}))
	s.Handle("/article", controller.Create(&app.ArticleController{}))
	s.Handle("/discuss", controller.Create(&app.DiscussController{}))
	s.Run()
}
