package log

import (
	"fmt"
	"github.com/alenstar/nanotask/log/queue"
	"gopkg.in/mgo.v2"
	"regexp"
	"time"
)

type MonGoWriter struct {
	c       *mgo.Collection
	session *mgo.Session
	q       *queue.Queue
}

func findTime(b []byte) time.Time {
	size := len(b)
	if size > 20 {
		size = 20
	}
	reg := regexp.MustCompile(`(\d{4})/(\d{1,2})/(\d{1,2}) (\d{1,2}):(\d{1,2}):(\d{1,2})`)
	t := reg.FindAllString(string(b[:size-1]), 1)
	if len(t) > 0 {
		tm, err := time.Parse("2006/01/02 15:04:05", t[0])
		if err != nil {
			fmt.Println("time.Parse:", err)
			return time.Now()
		}
		return tm
	}
	return time.Now()
}
func NewMonGoWriter(urls, dbname, colname string) *MonGoWriter {
	session, err := mgo.Dial(urls)
	if err != nil {
		panic(err)
	}
	out := &MonGoWriter{
		c:       session.DB(dbname).C(colname),
		session: session,
	}
	out.q = queue.New(func(val interface{}) {
		v := val.([]byte)

		if v != nil {
			size := len(v)
			if size > 21 {
				size = 21
			}
			vv := struct {
				Time time.Time
				Msg  string
			}{
				Time: findTime(v),
				Msg:  string(v[size-1:]),
			}

			err = out.c.Insert(vv)
			if err != nil {
				fmt.Println("Write:", err)
			}
		}
	})
	go func() {
		out.q.Run()
		defer session.Close()
	}()
	return out
}

func (m *MonGoWriter) Close() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	m.q.Done()
}

func (m *MonGoWriter) Write(p []byte) (n int, err error) {
	buf := make([]byte, len(p))
	copy(buf, p)
	m.q.Put(buf)
	return len(p), nil
}
