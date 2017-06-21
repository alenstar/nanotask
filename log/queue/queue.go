package queue

import (
	_list "container/list"
	"runtime"
	"sync"
)

type Queue struct {
	list    *_list.List
	lock    *sync.Mutex
	done    chan bool
	sem     chan int
	running bool
	handler func(interface{})
}

func New(fn func(interface{})) *Queue {
	return &Queue{
		list:    _list.New(),
		lock:    new(sync.Mutex),
		done:    make(chan bool),
		sem:     make(chan int, runtime.NumCPU()*4),
		running: false,
		handler: fn,
	}
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) Put(val interface{}) {

	q.lock.Lock()
	_ = q.list.PushFront(val)
	q.lock.Unlock()

	if q.running {
		select {
		case q.sem <- 1:
			// chan send success
			break
		default:
			// chan is full
			break
		}
	}
}

func (q *Queue) Get() interface{} {
	q.lock.Lock()
	e := q.list.Back()
	if e != nil {
		q.list.Remove(e)
	}
	q.lock.Unlock()
	if e == nil {
		return nil
	}
	return e.Value
}

func (q *Queue) Done() {
	if !q.running {
		return
	}
	q.running = false
	q.sem <- 0
	<-q.done
}

func (q *Queue) Run() {
	q.running = true
	for {
		select {
		case state := <-q.sem:
			for q.Size() > 0 {
				n := q.Get()
				if n == nil {
					break
				} else {
					q.handler(n)
				}
			}
			if state == 0 {
				q.done <- true
				return
			}
		}
	}
}
