package types

type Logger interface {
	Init()
	Debug(v string)
	Info(v string)
	Error(err error)
	Warning(v string)
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}
