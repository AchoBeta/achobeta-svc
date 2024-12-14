package tlog

import "go.uber.org/zap/zapcore"

type LogConfig struct {
	logLevel zapcore.Level
}

func NewConfig(opts ...func(config *LogConfig)) *LogConfig {
	config := &LogConfig{
		logLevel: zapcore.InfoLevel,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func WithLogLevel(logLevel zapcore.Level) func(config *LogConfig) {
	return func(config *LogConfig) {
		config.logLevel = logLevel
	}
}
