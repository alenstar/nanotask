package log

import (
	"fmt"
	"goworker/log/queue"
	"goworker/sockets"
	"time"
)

type TcpWriter struct {
	addr string
	q    *queue.Queue
	tcp  *sockets.TcpSocket
}

func NewTcpWriter(addr string) *TcpWriter {
	t := &TcpWriter{
		addr: addr,
		tcp:  sockets.NewTcpSocket(addr),
	}

	t.q = queue.New(func(val interface{}) {
		v := val.([]byte)
		if v != nil {
			_, err := t.tcp.WriteMessage(v)
			if err != nil {
				fmt.Println("Write:", err)
			}
		}
	})

	go func(t *sockets.TcpSocket, q *queue.Queue) {
		buf := make([]byte, 4096+32)
		for true {
			err := t.Connect(addr)
			if err != nil {
				// wait for try to do
				fmt.Println("Connect: ", err.Error())
				time.Sleep(time.Second * 3)
			} else {
				for true {
					n, err := t.ReadMessage(buf)
					if err != nil {
						fmt.Println("ReadMessage:", err)
						break
					} else if n > 0 {
						q.Put(buf[:n])
					}
				}
			}
		}
	}(t.tcp, t.q)

	go t.q.Run()
	return t
}

func (t *TcpWriter) Write(p []byte) (n int, err error) {
	buf := make([]byte, len(p))
	copy(buf, p)
	t.q.Put(buf)
	return len(p), nil
}

func (t *TcpWriter) Close() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	t.q.Done()
}
