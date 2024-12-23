package logger

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type EmptyLogger struct{}

func (*EmptyLogger) Info(args ...interface{})                  {}
func (*EmptyLogger) Infof(format string, args ...interface{})  {}
func (*EmptyLogger) Debug(args ...interface{})                 {}
func (*EmptyLogger) Debugf(format string, args ...interface{}) {}
func (*EmptyLogger) Error(args ...interface{})                 {}
func (*EmptyLogger) Errorf(format string, args ...interface{}) {}
func (*EmptyLogger) Fatal(args ...interface{})                 {}
func (*EmptyLogger) Fatalf(format string, args ...interface{}) {}
