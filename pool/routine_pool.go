package pool

import (
	//"container/list"
	"errors"
	//"fmt"
	//"github.com/alenstar/nanoweb/log"
	"runtime"
	"sync"
	//"time"
)

type coroutine struct {
	exit    chan struct{}
	notify  chan struct{}
	runner  func()
	_runing bool
}

func newcoroutine() *coroutine {
	return &coroutine{
		exit:    make(chan struct{}),
		notify:  make(chan struct{}),
		runner:  nil,
		_runing: false,
	}
}

func (c *coroutine) run() {
	c._runing = true
	for c._runing {
		select {
		case <-c.notify:
			if c.runner != nil {
				c.runner()
				c.runner = nil
			}
		//case <-time.After(time.Second * 3):
		//	log.Debug("running ...")
		case <-c.exit:
			c._runing = false
		}
	}
}

func (c *coroutine) shutdown() {
	var e struct{}
	select {
	case c.exit <- e:
	}
}

func (c *coroutine) running() bool {
	if c._runing {
		return true
	}
	return false
}

func (c *coroutine) empty() bool {
	if c.runner == nil {
		return true
	}
	return false
}

func (c *coroutine) do(f func()) bool {
	if c.runner != nil {
		return false
	}
	c.runner = f
	var n struct{}
	c.notify <- n
	return true
}

type Worker struct {
	sync.Mutex
	routines []*coroutine
}

func NewWorker(num uint) *Worker {
	routines := make([]*coroutine, num)
	for k := range routines {
		routines[k] = newcoroutine()
	}
	return &Worker{
		routines: routines,
	}
}

type CoroutinePool struct {
	cur     uint
	maxsize uint
	exit    chan struct{}
	notify  chan struct{}
	wg      *sync.WaitGroup
	groups  []*Worker
	//lock     *sync.Mutex
	//routines []*Worker //*list.List
}

func NewCoroutinePool(num uint) *CoroutinePool {
	groups := make([]*Worker, runtime.NumCPU()*4)
	for i := 0; i < len(groups); i++ {
		groups[i] = NewWorker(num)
	}
	return &CoroutinePool{
		cur:     0,
		maxsize: num,
		exit:    make(chan struct{}),
		notify:  make(chan struct{}),
		wg:      &sync.WaitGroup{},
		groups:  groups,
	}
}

func (c *CoroutinePool) Run() {
	for {
		select {
		case <-c.exit:
			break
		case <-c.notify:
		}
	}
	c.wg.Wait()
}

func (c *CoroutinePool) Shutdown() {
	for _, g := range c.groups {
		for _, v := range g.routines {
			if v.running() {
				v.shutdown()
			}
		}
	}

	//var e struct{}
	//sc.exit <- e
}

func (c *CoroutinePool) Add(f func()) error {
	if f != nil {
		c := c.findcoroutine()
		if c != nil {
			if !c.running() {
				//wg.Add(1)
				go c.run()
			}
			c.do(f)
		} else {
			return errors.New("coroutine pool is full")
		}
	}
	return nil
}

func (c *CoroutinePool) findcoroutine() *coroutine {
	for _, g := range c.groups {
		for _, v := range g.routines {
			if v.empty() {
				return v
			}
		}
	}
	return nil
}

// func (g *Coroutine) full() bool {
// 	for v := range g.routines {
// 		if v.runner == nil {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (g *Coroutine) empty() bool {
// 	return false
// }
