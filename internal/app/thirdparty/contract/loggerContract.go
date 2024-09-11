package contract

import "go.uber.org/zap"

type LoggerContract interface {
	Info(message string, tags ...zap.Field)
	Error(message string, err error, tags ...zap.Field)
}
