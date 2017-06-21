package queue

import (
	_ "errors"
	"fmt"
	_ "strings"
	"testing"
	"time"
)

type UserData struct {
	Id   int64
	Idx  int
	From string
	To   string
	Data string
}

func callback(val interface{}) {
	u := val.(*UserData)
	fmt.Printf("#%05d %s #%8d from %s to %s : %s\n", u.Idx, time.Now().String(), u.Id, u.From, u.To, u.Data)
}

func insert(q *Queue, u *UserData) {
	cnt := 20000
	for cnt > 0 {
		q.Put(u)
		cnt = cnt - 1
	}
}

func TestQueue(t *testing.T) {
	cnt := 20
	fmt.Println("Queue begin ")
	q := New(callback)
	go q.Run()
	time.Sleep(time.Second * 1) // wait queue run
	for cnt > 0 {
		//time.Sleep(time.Second * 3)
		u := &UserData{}
		u.Id = time.Now().Unix()
		u.Idx = cnt
		u.From = "hello"
		u.To = "world"
		u.Data = time.Now().String()
		go insert(q, u)
		cnt = cnt - 1
	}
	q.Done()
	fmt.Println("Queue end")
}
