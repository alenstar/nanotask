package log

import (
	"testing"
	"time"
)

func TestTcpWriter(t *testing.T) {
	fw := NewTcpWriter("127.0.0.1:8888")
	l := New(fw)

	// wait sockets connecting
	time.Sleep(time.Second * 3)

	l.Info("TestLog Info")

	l.Warn("TestLog Warn")

	l.Notice("TestLog Notice")

	l.Error("TestLog Error")

	l.Alert("TestLog Alert")

	a := &struct {
		A1 string
		A2 string
		A3 int
		A4 bool
	}{
		A1: "hello",
		A2: "world",
		A3: 4,
		A4: false,
	}
	cnt := 10
	for cnt > 0 {
		l.Debug("Debug", cnt, time.Now().String(), *a)
		cnt = cnt - 1
	}

	fw.Close()
}
