package log

import "testing"

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
