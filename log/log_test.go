package log

import (
	"testing"
)

func TestLog(t *testing.T) {

	Info("TestLog Info")

	Warn("TestLog Warn")

	Notice("TestLog Notice")

	Error("TestLog Error")

	Alert("TestLog Alert")

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
	Debug("Debug", *a)

	Shutdown()
}

/*
func TestFileWriter(t *testing.T) {
	fw := NewFileWriter("test.log")
	l := New(fw)

	time.Sleep(time.Second * 1)

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
	l.Debug("Debug", *a)

	fw.Close()
}

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
	l.Debug("Debug", *a)

	fw.Close()
}
func TestMonGoWriter(t *testing.T) {
	fw := NewMonGoWriter("127.0.0.1:27117, 127.0.0.1:27217, 127.0.0.1:27317", "log", "test")
	l := New(fw)

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

	l.Debug("Debug", *a)

	fw.Close()
}
*/