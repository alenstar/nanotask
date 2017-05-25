package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const defaultCalldepth = 3

var logs []*log.Logger

func init() {
	defaultOutput := os.Stderr

	logs = make([]*log.Logger, int(FatalLevel))
	for i := 0; i < int(FatalLevel); i++ {
		logs[i] = log.New(defaultOutput, Level(i).Color(), log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func output(l Level, v ...interface{}) {
	var str string
	for _, s := range v {
		if strings.HasSuffix(str, " ") || len(str) == 0 {
			str = str + fmt.Sprint(s)
		} else {
			str = str + " " + fmt.Sprint(s)
		}
	}
	logs[l].Output(defaultCalldepth, l.String()+l.Reset()+str)
}

func Shutdown() {
	//logs[Trace].Output("shutdown").End()
}

func Fatal(v ...interface{}) {
	output(FatalLevel, v...)
	//logs[FatalLevel].Output(defaultCalldepth, FatalLevel.String()+FatalLevel.Reset()+fmt.Sprint(v...))
}

func Panic(v ...interface{}) {
	output(PanicLevel, v...)
	//logs[PanicLevel].Output(defaultCalldepth, PanicLevel.String()+PanicLevel.Reset()+fmt.Sprint(v...))
}

func Trace(v ...interface{}) {
	output(TraceLevel, v...)
	//logs[TraceLevel].Output(defaultCalldepth, TraceLevel.String()+TraceLevel.Reset()+fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	output(DebugLevel, v...)
	//logs[DebugLevel].Output(defaultCalldepth, DebugLevel.String()+DebugLevel.Reset()+fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	// var str string
	// for _, s := range v {
	// 	if strings.HasSuffix(str, " ") {
	// 		str = str + fmt.Sprint(s)
	// 	} else {
	// 		str = str + " " + fmt.Sprint(s)
	// 	}
	// }
	// logs[InfoLevel].Output(defaultCalldepth, InfoLevel.String()+InfoLevel.Reset()+str)
	output(InfoLevel, v...)
}

func Notice(v ...interface{}) {
	output(NoticeLevel, v...)
	//logs[NoticeLevel].Output(defaultCalldepth, NoticeLevel.String()+NoticeLevel.Reset()+fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	output(WarnLevel, v...)
	//logs[WarnLevel].Output(defaultCalldepth, WarnLevel.String()+WarnLevel.Reset()+fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	output(ErrorLevel, v...)
	//logs[ErrorLevel].Output(defaultCalldepth, ErrorLevel.String()+ErrorLevel.Reset()+fmt.Sprint(v...))
}

func Alert(v ...interface{}) {
	output(AlertLevel, v...)
	//logs[AlertLevel].Output(defaultCalldepth, AlertLevel.String()+AlertLevel.Reset()+fmt.Sprint(v...))
}
