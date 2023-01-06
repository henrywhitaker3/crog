package log

import (
	"fmt"
	"time"
)

var Verbose bool = false

func Info(value string) {
	printLog("INFO", value)
}

func ForceInfo(value string) {
	print("INFO", value)
}

func Infof(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	Info(value)
}

func ForceInfof(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	ForceInfo(value)
}

func Error(value string) {
	printLog("Error", value)
}

func ForceError(value string) {
	print("Error", value)
}

func Errorf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	Error(value)
}

func ForceErrorf(format string, args ...any) {
	value := fmt.Sprintf(format, args...)
	ForceError(value)
}

func printLog(level, content string) {
	if Verbose {
		print(level, content)
	}
}

func print(level, content string) {
	t := time.Now()

	fmt.Printf("[%s] %s: %s\n", t.Format(time.RFC3339), level, content)
}
