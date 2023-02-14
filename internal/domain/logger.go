package domain

type Verbosity int

type Logger interface {
	Info(string)
	Infof(string, ...any)

	Debug(string)
	Debugf(string, ...any)

	Error(string)
	Errorf(string, ...any)

	SetVerbosity(Verbosity)
	GetVerbosity() Verbosity
}
