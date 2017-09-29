package utils

import (
	"sync/atomic"
	"unsafe"
)

type element struct {
	value interface{}
	next  unsafe.Pointer
}

type Queue struct {
	size uint32
	root unsafe.Pointer
}

func NewQueue() *Queue {
	return &Queue{
		size: 0,
		root: unsafe.Pointer(nil),
	}
}

func (q *Queue) Push(value interface{}) {
	if value == nil {
		return
	}
	v := new(element)
	v.value = value
	elem := unsafe.Pointer(v)
	// FIXME
	// this not goroutine safe
	cursor := &q.root
	for {
		if atomic.CompareAndSwapPointer(cursor, nil, elem) {
			atomic.AddUint32(&q.size, 1)
			return
		}
		cursor = &(*element)(*cursor).next
	}
}

func (q *Queue) Pop() interface{} {
	var value interface{}
	cursor := &q.root
	for atomic.LoadPointer(cursor) != nil {
		value = (*element)(*cursor).value
		atomic.StorePointer(cursor, (*element)(*cursor).next)
		atomic.StoreUint32(&q.size, atomic.LoadUint32(&q.size) - 1)
		return value
	}
	return value
}

func (q *Queue) Reset() {
	atomic.StoreUint32(&q.size, 0)
	atomic.StorePointer(&q.root, nil)
}

func (q *Queue) Size() uint {
	size := atomic.LoadUint32(&q.size)
	return uint(size)
}
