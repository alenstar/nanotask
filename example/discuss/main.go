package main

import (
	_ "goworker/config"
	"goworker/controller"
	"goworker/http/server"
	"goworker/log"
	_ "goworker/modules"
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
