package log_config_manager

import (
	"my_libs/logger/log_config"
	"my_libs/logger/log_level"
)

var currentConfig *log_config.LogConfig

func InitDefaultConfig() {
	currentConfig = newDefaultConfig()
}

func newDefaultConfig() *log_config.LogConfig {
	return &log_config.LogConfig{
		BufferSize:   log_config.DEFAULT_BUFFER_SIZE,
		FileFullName: log_config.DEFAULT_LOG_FILE,
		MaxFileSize:  log_config.DEFAULT_FILE_SIZE,
		CallerSkip:   log_config.DEFAULT_CALLER_SKIP,
		LogLevel:     log_config.DEFAULT_LOG_LEVEL}
}

func InitNewConfig(config *log_config.LogConfig) {
	currentConfig = config
	if currentConfig.LogLevel <= log_level.NONE || currentConfig.LogLevel > log_level.OFF {
		currentConfig.LogLevel = log_config.DEFAULT_LOG_LEVEL
	}
	if currentConfig.BufferSize <= 0 || currentConfig.BufferSize > log_config.MAX_BUFFER_SIZE {
		currentConfig.BufferSize = log_config.DEFAULT_BUFFER_SIZE
	}
	if currentConfig.MaxFileSize <= 0 || currentConfig.MaxFileSize > log_config.MAX_FILE_SIZE {
		currentConfig.MaxFileSize = log_config.DEFAULT_FILE_SIZE
	}
	if currentConfig.CallerSkip < log_config.MIN_CALLER_SKIP || currentConfig.CallerSkip > log_config.MAX_CALLER_SKIP {
		currentConfig.CallerSkip = log_config.DEFAULT_CALLER_SKIP
	}
	if len(currentConfig.FileFullName) < 1 {
		currentConfig.FileFullName = log_config.DEFAULT_LOG_FILE
	}
}

func CurrentConfig() *log_config.LogConfig {
	return currentConfig
}
