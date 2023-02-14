package log

import (
	"fmt"
	"io"
	"time"
)

var (
	Log *Logger
)

type Logger struct {
	Verbose bool
	Output  io.Writer
}

func (l *Logger) Info(value string) {
	l.printLog("INFO", value)
}

func (l *Logger) ForceInfo(value string) {
	l.print("INFO", value)
}

func (l *Logger) Infof(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.Info(value)
}

func (l *Logger) ForceInfof(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.ForceInfo(value)
}

func (l *Logger) Error(value string) {
	l.printLog("ERROR", value)
}

func (l *Logger) ForceError(value string) {
	l.print("ERROR", value)
}

func (l *Logger) Errorf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.Error(value)
}

func (l *Logger) ForceErrorf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	l.ForceError(value)
}

func (l *Logger) printLog(level, content string) {
	if l.Verbose {
		l.print(level, content)
	}
}

func (l *Logger) print(level, content string) {
	t := time.Now()

	fmt.Fprintf(l.Output, "[%s] %s: %s\n", t.Format(time.RFC3339), level, content)
}
