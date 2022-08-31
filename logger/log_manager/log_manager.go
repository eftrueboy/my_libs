package log_manager

import (
	"fmt"
	"log"
	"my_libs/logger/log_buffer"
	"my_libs/logger/log_config_manager"
	"my_libs/logger/log_level"
	"my_libs/logger/log_task"
	"my_libs/utils/date_util"
	"runtime"
	"strings"
)

var buffer *log_buffer.LogBuffer
var logTask *log_task.LogTask
var initialized = false

func Init() {
	buffer = log_buffer.NewLogBuffer(log_config_manager.CurrentConfig().BufferSize)
	logTask = log_task.NewLogTask(log_config_manager.CurrentConfig().FileFullName)
	go runLogTask()
	initialized = true
}

func assert() {
	if !initialized {
		panic("log_manager not initialized.")
	}
}

func Write(logLevel log_level.LogLevel, args ...interface{}) {
	assert()

	if log_config_manager.CurrentConfig().LogLevel > logLevel {
		return
	}

	logCaption := log_level.GetLogCaption(logLevel)
	_, filepath, line, ok := runtime.Caller(log_config_manager.CurrentConfig().CallerSkip)
	if !ok {
		log.Println("runtime.Caller() failed!")
		return
	}
	lastIndex := strings.LastIndex(filepath, "/")
	filename := filepath[lastIndex+1:]
	text := fmt.Sprintf("%s [%s] %s:%d: %s", date_util.NowDateMicrosecond(), logCaption, filename, line, fmt.Sprint(args...))
	buffer.Write(text)
}

func runLogTask() {
	defer func() {
		if err := recover(); err != nil {
			go runLogTask()
		}
	}()
	writeLog()
}

func writeLog() {
	for {
		text := buffer.Read()
		logTask.Write(text)
		fmt.Println(text)
	}
}
