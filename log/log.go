package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const defaultCalldepth = 3

var logs []*log.Logger

var defaultLogger *Logger

type Logger struct {
	color     bool
	callDepth int
	logs      []*log.Logger
}

func init() {
	if defaultLogger == nil {
		defaultLogger = New(nil)
	}
}

func SetDefaultLog(l *Logger) {
	if l != nil {
		defaultLogger = l
	}
}

func New(output io.Writer) *Logger {
	color := false
	if output == nil {
		output = os.Stderr
		color = true
	}
	logs = make([]*log.Logger, int(FATAL) + 1)
	for i := 0; i <= int(FATAL); i++ {
		logs[i] = log.New(output, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return &Logger{
		color:     color,
		callDepth: defaultCalldepth,
		logs:      logs,
	}
}
func (l *Logger) output(level LogLevel, v ...interface{}) *Logger {
	var str string
	for _, s := range v {
		if strings.HasSuffix(str, " ") || len(str) == 0 {
			str = str + fmt.Sprint(s)
		} else {
			str = str + " " + fmt.Sprint(s)
		}
	}
	if l.color {
		l.logs[level].Output(defaultCalldepth, level.ColorString()+str)
	} else {
		l.logs[level].Output(defaultCalldepth, level.String()+str)
	}
	return l
}

func Shutdown() {
	// TODO
}

func (l *Logger) UseColor() *Logger {
	l.color = true
	return l
}

func (l *Logger) Fatal(v ...interface{}) {
	l.output(FATAL, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.output(PANIC, v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.output(TRACE, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(DEBUG, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.output(INFO, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.output(NOTICE, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(WARN, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(ERROR, v...)
}

func (l *Logger) Alert(v ...interface{}) {
	l.output(ALERT, v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.output(FATAL, v...)
}

func Panic(v ...interface{}) {
	defaultLogger.output(PANIC, v...)
}

func Trace(v ...interface{}) {
	defaultLogger.output(TRACE, v...)
}

func Debug(v ...interface{}) {
	defaultLogger.output(DEBUG, v...)
}

func Info(v ...interface{}) {
	defaultLogger.output(INFO, v...)
}

func Notice(v ...interface{}) {
	defaultLogger.output(NOTICE, v...)
}

func Warn(v ...interface{}) {
	defaultLogger.output(WARN, v...)
}

func Error(v ...interface{}) {
	defaultLogger.output(ERROR, v...)
}

func Alert(v ...interface{}) {
	defaultLogger.output(ALERT, v...)
}

// output to stdout
func Printf(format string, v... interface{}){
	fmt.Printf(format, v...)
}
func Println(v ...interface{}){
	fmt.Println(v...)
}
func Print(v ...interface{}) {
	fmt.Print(v...)
}