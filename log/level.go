package log

// AllLevels is an array of all log levels, for easier registering of all levels to a handler
var AllLevels = []Level{
	DebugLevel,
	TraceLevel,
	InfoLevel,
	NoticeLevel,
	WarnLevel,
	ErrorLevel,
	PanicLevel,
	AlertLevel,
	FatalLevel,
}

// Level of the log
type Level uint8

// Log levels.
const (
	DebugLevel Level = iota
	TraceLevel
	InfoLevel
	NoticeLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	AlertLevel
	FatalLevel // same as syslog CRITICAL
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "\t[D] " // debug
	case TraceLevel:
		return "\t[T] " // trace
	case InfoLevel:
		return "\t[I] " // info
	case NoticeLevel:
		return "\t[N] " // notice
	case WarnLevel:
		return "\t[W] " // warn
	case ErrorLevel:
		return "\t[E] " // error
	case PanicLevel:
		return "\t[P] " // panic
	case AlertLevel:
		return "\t[A] " // alert
	case FatalLevel:
		return "\t[F] " // fatal
	default:
		return "Unknow Level "
	}
}

func (l Level) Reset() string {
	return string(Reset)
}

func (l Level) Color() string {
	switch l {
	case DebugLevel:
		return string(LightBlue) // debug
	case TraceLevel:
		return string(Black) // trace
	case InfoLevel:
		return string(Green) // info
	case NoticeLevel:
		return string(Cyan) // notice
	case WarnLevel:
		return string(LightYellow) // warn
	case ErrorLevel:
		return string(LightRed) // error
	case PanicLevel:
		return string(White) // panic
	case AlertLevel:
		return string(Yellow) // alert
	case FatalLevel:
		return string(Gray) // fatal
	default:
		return ""
	}
}

func (l Level) ColorString() string {
	return l.Color() + l.String() + l.Reset()
}

func (l Level) Background() string {
	switch l {
	case DebugLevel:
		return "" // debug
	case TraceLevel:
		return "" // trace
	case InfoLevel:
		return "" // info
	case NoticeLevel:
		return "" // notice
	case WarnLevel:
		return "" // warn
	case ErrorLevel:
		return "" // error
	case PanicLevel:
		return "" // panic
	case AlertLevel:
		return "" // alert
	case FatalLevel:
		return "" // fatal
	default:
		return ""
	}
}

// TODO: Add a bytes method along with string
