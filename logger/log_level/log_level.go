package log_level

type LogLevel uint32

const (
	NONE LogLevel = iota
	ALL
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

var logCaption []string

func InitLogCaption() {
	logCaption = make([]string, LogLevel(OFF)+1)
	logCaption[ALL] = "ALL"
	logCaption[TRACE] = "TRACE"
	logCaption[DEBUG] = "DEBUG"
	logCaption[INFO] = "INFO"
	logCaption[WARN] = "WARN"
	logCaption[ERROR] = "ERROR"
	logCaption[FATAL] = "FATAL"
	logCaption[OFF] = "OFF"
}

func GetLogCaption(logLevel LogLevel) string {
	return logCaption[logLevel]
}
