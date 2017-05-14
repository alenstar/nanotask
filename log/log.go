package log

import (
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

var defaultClog *console.Console

func init() {
	cLog := console.New()
	log.RegisterHandler(cLog, log.AllLevels...)
}

func Shutdown() {
	log.Trace("shutdown").End()
}

func Debug(v ...interface{}) {
	log.Debug(v)
}

func Info(v ...interface{}) {
	log.Info(v)
}

func Notice(v ...interface{}) {
	log.Notice(v)
}

func Warn(v ...interface{}) {
	log.Warn(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Alert(v ...interface{}) {
	log.Alert(v)
}
