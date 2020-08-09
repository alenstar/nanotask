package goworker

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type Worker struct {
	router *httprouter.Router
	srv *http.Server
}
func New(addr string) *Worker {
	n:= &Worker {
		httprouter.New(),
		nil,
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: n.router,
	}
	n.srv = srv

	return n
}

func (n *Worker) Run() error {
	return n.srv.ListenAndServe()
}

func (n *Worker) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return n.srv.Shutdown(ctx)
}


func (n *Worker) GET(path string, handler func(c *Context)) *Worker{
	n.router.GET(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) POST(path string, handler func(c *Context)) *Worker{
	n.router.POST(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) PUT(path string, handler func(c *Context)) *Worker{
	n.router.PUT(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) PATCH(path string, handler func(c *Context)) *Worker{
	n.router.PATCH(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) DELETE(path string, handler func(c *Context)) *Worker{
	n.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) HEAD(path string, handler func(c *Context)) *Worker{
	n.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}
func (n *Worker) OPTIONS(path string, handler func(c *Context)) *Worker{
	n.router.OPTIONS(path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		c := &Context{
			w,
			r,
			ps,
		}
		handler(c)
	})
	return n
}

func (n *Worker) UseStaticDir(path, dir string) *Worker{
	n.router.Handler("GET", path, http.FileServer(http.Dir(dir)))
	return n
}