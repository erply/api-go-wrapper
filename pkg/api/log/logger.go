package log

import "log"

var Log Logger

func init() {
	Log = NullLogger{}
}

type Type int

const (
	Debug Type = iota + 1
	Info
	Warn
	Error
)

type Logger interface {
	Log(t Type, message string, arguments ...interface{})
}

type NullLogger struct{}

func (nl NullLogger) Log(t Type, message string, arguments ...interface{}) {}

type StdLogger struct{}

func (sl StdLogger) Log(t Type, message string, arguments ...interface{}) {
	log.Printf(message, arguments...)
}
