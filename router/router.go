package router

import (
	"goworker/controller"
	"net/http"
)

type Router struct {
}

func (r *Router) Add(parten string, c controller.IController) func(http.ResponseWriter, *http.Request) {
	return nil
}

func (r *Router) NameSpace(ns string) *Router {
	return nil
}
