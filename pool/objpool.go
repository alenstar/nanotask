package pool

import (
	"container/list"
	"sync"
)

type ObjectPool struct {
	notify   chan struct{}
	lock     *sync.Mutex
	routines *list.List
	running  bool
}

func NewObjectPool() *ObjectPool {
	r := list.New()
	r.Init()
	return &ObjectPool{
		notify:   make(chan struct{}),
		lock:     &sync.Mutex{},
		routines: r,
	}
}
