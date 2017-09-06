package pool

import (
	"sync"
)

type item struct {

}

type ObjectPool struct {
	sync.Mutex
	notify   chan struct{}
	objs 	map[interface{}]bool
	running  bool
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		notify:   make(chan struct{}),
		objs: make(map[interface{}]bool),
		running:false,
	}
}

func (o *ObjectPool) Release(obj interface{}){
	o.Lock()
	if used, ok:= o.objs[obj]; ok {
		if used {
			o.objs[used] = false
		}
	}
}

func (o *ObjectPool) Obtain() interface{} {


}
