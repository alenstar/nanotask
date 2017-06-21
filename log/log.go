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
	calldepth int
	logs      []*log.Logger
}

func init() {
	if defaultLogger == nil {
		defaultLogger = New(nil)
	}
}

func New(output io.Writer) *Logger {
	color := false
	if output == nil {
		output = os.Stderr
		color = true
	}
	logs = make([]*log.Logger, int(FatalLevel))
	for i := 0; i < int(FatalLevel); i++ {
		logs[i] = log.New(output, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return &Logger{
		color:     color,
		calldepth: defaultCalldepth,
		logs:      logs,
	}
}
func (l *Logger) output(level Level, v ...interface{}) *Logger {
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
	l.output(FatalLevel, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.output(PanicLevel, v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.output(TraceLevel, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(DebugLevel, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.output(InfoLevel, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.output(NoticeLevel, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(WarnLevel, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(ErrorLevel, v...)
}

func (l *Logger) Alert(v ...interface{}) {
	l.output(AlertLevel, v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.output(FatalLevel, v...)
}

func Panic(v ...interface{}) {
	defaultLogger.output(PanicLevel, v...)
}

func Trace(v ...interface{}) {
	defaultLogger.output(TraceLevel, v...)
}

func Debug(v ...interface{}) {
	defaultLogger.output(DebugLevel, v...)
}

func Info(v ...interface{}) {
	defaultLogger.output(InfoLevel, v...)
}

func Notice(v ...interface{}) {
	defaultLogger.output(NoticeLevel, v...)
}

func Warn(v ...interface{}) {
	defaultLogger.output(WarnLevel, v...)
}

func Error(v ...interface{}) {
	defaultLogger.output(ErrorLevel, v...)
}

func Alert(v ...interface{}) {
	defaultLogger.output(AlertLevel, v...)
}
