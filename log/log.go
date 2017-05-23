package log

import (
	"fmt"
	"log"
	"os"
)

const defaultCalldepth = 2

var logs []*log.Logger

func init() {
	defaultOutput := os.Stderr

	logs = make([]*log.Logger, int(FatalLevel))
	for i := 0; i < int(FatalLevel); i++ {
		logs[i] = log.New(defaultOutput, Level(i).Color(), log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func Shutdown() {
	//logs[Trace].Output("shutdown").End()
}

func Fatal(v ...interface{}) {
	logs[FatalLevel].Output(defaultCalldepth, FatalLevel.String()+FatalLevel.Reset()+fmt.Sprint(v...))
}

func Panic(v ...interface{}) {
	logs[PanicLevel].Output(defaultCalldepth, PanicLevel.String()+PanicLevel.Reset()+fmt.Sprint(v...))
}

func Trace(v ...interface{}) {
	logs[TraceLevel].Output(defaultCalldepth, TraceLevel.String()+TraceLevel.Reset()+fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	//log.Debug(v)
	logs[DebugLevel].Output(defaultCalldepth, DebugLevel.String()+DebugLevel.Reset()+fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	//log.Info(v)
	logs[InfoLevel].Output(defaultCalldepth, InfoLevel.String()+InfoLevel.Reset()+fmt.Sprint(v...))
}

func Notice(v ...interface{}) {
	//log.Notice(v)
	logs[NoticeLevel].Output(defaultCalldepth, NoticeLevel.String()+NoticeLevel.Reset()+fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	// log.Warn(v)
	logs[WarnLevel].Output(defaultCalldepth, WarnLevel.String()+WarnLevel.Reset()+fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	// log.Error(v)
	logs[ErrorLevel].Output(defaultCalldepth, ErrorLevel.String()+ErrorLevel.Reset()+fmt.Sprint(v...))
}

func Alert(v ...interface{}) {
	// log.Alert(v)
	logs[AlertLevel].Output(defaultCalldepth, AlertLevel.String()+AlertLevel.Reset()+fmt.Sprint(v...))
}
