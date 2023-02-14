package log

import (
	"fmt"
	"io"
	"time"

	"github.com/henrywhitaker3/crog/internal/domain"
)

var (
	Log domain.Logger
)

type Logger struct {
	Verbosity domain.Verbosity
	Output    io.Writer
}

func (l *Logger) Info(value string) {
	if l.Verbosity >= Info {
		l.print("INFO", value)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.Info(value)
}

func (l *Logger) Error(value string) {
	if l.Verbosity >= Error {
		l.print("ERROR", value)
	}
}

func (l *Logger) Errorf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.Error(value)
}

func (l *Logger) Debug(value string) {
	if l.Verbosity >= Debug {
		l.print("DEBUG", value)
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.Debug(value)
}

func (l *Logger) print(level, content string) {
	t := time.Now()

	fmt.Fprintf(l.Output, "[%s] %s: %s\n", t.Format(time.RFC3339), level, content)
}

func (l *Logger) SetVerbosity(v domain.Verbosity) {
	l.Verbosity = v
}

func (l *Logger) GetVerbosity() domain.Verbosity {
	return l.Verbosity
}
