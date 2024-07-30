package logger

type LogLevel int

//go:generate enumer -trimprefix=LogLevel -type=LogLevel -json -output log_level_enum.go
const (
	LogLevelUnknown LogLevel = iota
	LogLevelInfo
	LogLevelDebug
	LogLevelWarning
	LogLevelError
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string)
	Error(msg string)
}
