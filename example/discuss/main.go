package main

import (
	_ "github.com/alenstar/nanotask/config"
	"github.com/alenstar/nanotask/controller"
	"github.com/alenstar/nanotask/http/server"
	"github.com/alenstar/nanotask/log"
	_ "github.com/alenstar/nanotask/modules"
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
	s.Handle("/user", controller.Create(&UserController{}))
	s.Handle("/article", controller.Create(&ArticleController{}))
	s.Handle("/discuss", controller.Create(&DiscussController{}))
	s.Run()
}
