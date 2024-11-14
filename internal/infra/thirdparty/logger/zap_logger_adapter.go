package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLoggerAdapter struct {
	logger *zap.Logger
}

func NewZapLoggerAdapter() *ZapLoggerAdapter {
	logger := &ZapLoggerAdapter{logger: &zap.Logger{}}
	logger.init()
	return logger
}

func (la *ZapLoggerAdapter) init() {
	logConfig := zap.Config{
		OutputPaths: []string{la.getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(la.getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	la.logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func (la *ZapLoggerAdapter) Info(message string) {
	la.logger.Info(message)
	la.logger.Sync()
}

func (la *ZapLoggerAdapter) Error(message string) {
	la.logger.Error(message)
	la.logger.Sync()
}

func (la *ZapLoggerAdapter) getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
	if output == "" {
		la.logger.Warn("No log output provided, defaulting to stdout")
		return "stdout"
	}

	return output
}

func (la *ZapLoggerAdapter) getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		la.logger.Warn("Invalid log level, defaulting to InfoLevel", zap.String("provided_level", os.Getenv("LOG_LEVEL")))
		return zapcore.InfoLevel
	}
}
