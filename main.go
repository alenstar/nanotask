package main

import (
	"github.com/alenstar/nanoweb/controller"
	"github.com/alenstar/nanoweb/http/server"
	"github.com/alenstar/nanoweb/log"
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
	s.Handle("/object", controller.Create(&MyController{}))
	s.Run()
}
