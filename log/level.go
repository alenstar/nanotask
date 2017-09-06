package log

// AllLevels is an array of all log levels, for easier registering of all levels to a handler

// Level of the log
type LogLevel uint8

// Log levels.
const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	NOTICE
	WARN
	ALERT
	ERROR
	FATAL // same as syslog CRITICAL
	PANIC
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "\t[D] " // debug
	case TRACE:
		return "\t[T] " // trace
	case INFO:
		return "\t[I] " // info
	case NOTICE:
		return "\t[N] " // notice
	case WARN:
		return "\t[W] " // warn
	case ALERT:
		return "\t[A] " // alert
	case ERROR:
		return "\t[E] " // error
	case FATAL:
		return "\t[F] " // fatal
	case PANIC:
		return "\t[P] " // panic
	default:
		return "Unknow Level "
	}
}

func (l LogLevel) Reset() string {
	return string(Reset)
}

func (l LogLevel) Color() string {
	switch l {
	case DEBUG:
		return string(LightBlue) // debug
	case TRACE:
		return string(Black) // trace
	case INFO:
		return string(Green) // info
	case NOTICE:
		return string(Cyan) // notice
	case WARN:
		return string(LightYellow) // warn
	case ERROR:
		return string(LightRed) // error
	case PANIC:
		return string(White) // panic
	case ALERT:
		return string(Yellow) // alert
	case FATAL:
		return string(Gray) // fatal
	default:
		return ""
	}
}

func (l LogLevel) ColorString() string {
	return l.Color() + l.String() + l.Reset()
}

func (l LogLevel) Background() string {
	switch l {
	case DEBUG:
		return "" // debug
	case TRACE:
		return "" // trace
	case INFO:
		return "" // info
	case NOTICE:
		return "" // notice
	case WARN:
		return "" // warn
	case ERROR:
		return "" // error
	case PANIC:
		return "" // panic
	case ALERT:
		return "" // alert
	case FATAL:
		return "" // fatal
	default:
		return ""
	}
}

// TODO: Add a bytes method along with string
