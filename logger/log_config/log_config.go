package log_config

import "my_libs/logger/log_level"

const (
	MAX_BUFFER_SIZE     = 500_0000                //日志缓冲区的最大大小, 500万个队列内容.
	DEFAULT_BUFFER_SIZE = 100_0000                //日志缓冲区的默认大小,
	MAX_FILE_SIZE       = 1024 * 1024 * 1024 * 10 //日志文件最大大小, 10GB
	DEFAULT_FILE_SIZE   = 1024 * 1024 * 100       //日志文件默认大小, 100MB
	DEFAULT_LOG_FILE    = "./logs/mylogger.log"   //默认日志文件
	MAX_CALLER_SKIP     = 5                       //最大调用深度
	MIN_CALLER_SKIP     = 3                       //最小调用深度
	DEFAULT_CALLER_SKIP = 3                       //默认调用深度
	DEFAULT_LOG_LEVEL   = log_level.ERROR         //默认日志级别
)

type LogConfig struct {
	BufferSize   int
	FileFullName string
	MaxFileSize  int
	CallerSkip   int
	LogLevel     log_level.LogLevel
}
