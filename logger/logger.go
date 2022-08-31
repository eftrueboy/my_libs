package logger

import (
	"fmt"
	"my_libs/logger/log_config"
	"my_libs/logger/log_config_manager"
	"my_libs/logger/log_level"
	"my_libs/logger/log_manager"
)

const config = "./conf/mylogger.json" //配置文件路径

var initialized bool

func InitDefault() {
	InitByConfig(nil)
}

func InitByConfig(config *log_config.LogConfig) {
	log_level.InitLogCaption()
	if config == nil {
		log_config_manager.InitDefaultConfig()
	} else {
		log_config_manager.InitNewConfig(config)
	}
	log_manager.Init()
	initialized = true
}

func Init(logBufferSize int, fileFullName string, maxFileSize int, callerSkip int, logLevel log_level.LogLevel) {
	config := log_config.LogConfig{BufferSize: logBufferSize, FileFullName: fileFullName, MaxFileSize: maxFileSize,
		CallerSkip: callerSkip, LogLevel: logLevel}
	InitByConfig(&config)
}

func assert() {
	if !initialized {
		panic("my_logger not initialized.")
	}
}

func Trace(args ...interface{}) {
	traceLog(args...)
}

func Tracef(format string, args ...interface{}) {
	traceLog(fmt.Sprintf(format, args...))
}

func traceLog(args ...interface{}) {
	assert()
	log_manager.Write(log_level.TRACE, args...)
}

func Debug(args ...interface{}) {
	assert()
	debugLog(args...)
}

func Debugf(format string, args ...interface{}) {
	assert()
	debugLog(fmt.Sprintf(format, args...))
}

func debugLog(args ...interface{}) {
	log_manager.Write(log_level.DEBUG, args...)
}

func Info(args ...interface{}) {
	assert()
	infoLog(args...)
}

func Infof(format string, args ...interface{}) {
	assert()
	infoLog(fmt.Sprintf(format, args...))
}

func infoLog(args ...interface{}) {
	log_manager.Write(log_level.INFO, args...)
}

func Warn(args ...interface{}) {
	assert()
	warnLog(args...)
}

func Warnf(format string, args ...interface{}) {
	assert()
	warnLog(fmt.Sprintf(format, args...))
}

func warnLog(args ...interface{}) {
	log_manager.Write(log_level.WARN, args...)
}

func Error(args ...interface{}) {
	assert()
	errorLog(args...)
}

func Errorf(format string, args ...interface{}) {
	assert()
	errorLog(fmt.Sprintf(format, args...))
}

func errorLog(args ...interface{}) {
	log_manager.Write(log_level.ERROR, args...)
}

func Fatal(args ...interface{}) {
	assert()
	fatalLog(args...)
}

func Fatalf(format string, args ...interface{}) {
	assert()
	fatalLog(fmt.Sprintf(format, args...))
}

func fatalLog(args ...interface{}) {
	log_manager.Write(log_level.FATAL, args...)
}
