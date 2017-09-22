package nano

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type Nano struct {
	router *httprouter.Router
	srv *http.Server
}
func New(addr string) *Nano {
	n:= &Nano {
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

func (n *Nano) Run() error {
	return n.srv.ListenAndServe()
}

func (n *Nano) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return n.srv.Shutdown(ctx)
}


func (n *Nano) GET(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) POST(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) PUT(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) PATCH(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) DELETE(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) HEAD(path string, handler func(c *Context)) *Nano{
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
func (n *Nano) OPTIONS(path string, handler func(c *Context)) *Nano{
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

func (n *Nano) UseStaticDir(path, dir string) *Nano{
	n.router.Handler("GET", path, http.FileServer(http.Dir(dir)))
	return n
}