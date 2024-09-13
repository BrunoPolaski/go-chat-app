package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerAdapter struct {
	Logger *zap.Logger
}

func (la *LoggerAdapter) init() {
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
	la.Logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func (la *LoggerAdapter) Info(message string, tags ...zap.Field) {
	la.Logger.Info(message, tags...)
	la.Logger.Sync()
}

func (la *LoggerAdapter) Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	la.Logger.Error(message, tags...)
	la.Logger.Sync()
}

func (la *LoggerAdapter) getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
	if output == "" {
		la.Logger.Warn("No log output provided, defaulting to stdout")
		return "stdout"
	}

	return output
}

func (la *LoggerAdapter) getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		la.Logger.Warn("Invalid log level, defaulting to InfoLevel", zap.String("provided_level", os.Getenv("LOG_LEVEL")))
		return zapcore.InfoLevel
	}
}
