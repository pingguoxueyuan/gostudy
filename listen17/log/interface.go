package log

type LogInterface interface {
	LogDebug(msg string)
	LogWarn(msg string)
}
