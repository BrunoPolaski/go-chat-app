package contract

type LoggerContract interface {
	Info(message string)
	Error(message string)
}
