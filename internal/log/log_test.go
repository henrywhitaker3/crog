package log

import (
	"testing"
)

var (
	logger *Logger
	writer *mockWriter
)

func setup() {
	writer = &mockWriter{lines: []string{}}

	logger = &Logger{
		Verbosity: Debug,
		Output:    writer,
	}
}

func TestItDoesntLogDebugWhenLevelSetToInfo(t *testing.T) {
	setup()

	logger.SetVerbosity(Info)

	logger.Debug("bongo")

	if len(writer.lines) != 0 {
		t.Error("wrote more than 0 bytes")
	}
}

func TestItDoesntLogDebugWhenLevelSetToError(t *testing.T) {
	setup()

	logger.SetVerbosity(Error)

	logger.Debug("bongo")

	if len(writer.lines) != 0 {
		t.Error("wrote more than 0 bytes")
	}
}

func TestItLogsDebugWhenLevelSetToDebug(t *testing.T) {
	setup()

	logger.SetVerbosity(Debug)

	logger.Debug("bongo")

	if len(writer.lines) != 1 {
		t.Error("didn't write 1 line to log")
	}
}

func TestItDoesntLogInfoWhenLevelSetToError(t *testing.T) {
	setup()

	logger.SetVerbosity(Error)

	logger.Info("bongo")

	if len(writer.lines) != 0 {
		t.Error("wrote more than 0 bytes")
	}
}

func TestItLogsInfoWhenLevelSetToInfo(t *testing.T) {
	setup()

	logger.SetVerbosity(Info)

	logger.Info("bongo")

	if len(writer.lines) != 1 {
		t.Error("didn't write 1 line to log")
	}
}

func TestItLogsInfoWhenLevelSetToDebug(t *testing.T) {
	setup()

	logger.SetVerbosity(Debug)

	logger.Info("bongo")

	if len(writer.lines) != 1 {
		t.Error("didn't write 1 line to log")
	}
}

func TestItLogsInfoErrorWhenLevelSetToError(t *testing.T) {
	setup()

	logger.SetVerbosity(Error)

	logger.Error("bongo")

	if len(writer.lines) == 0 {
		t.Error("didn't write 1 line to log")
	}
}

func TestItLogsErrorWhenLevelSetToInfo(t *testing.T) {
	setup()

	logger.SetVerbosity(Info)

	logger.Error("bongo")

	if len(writer.lines) != 1 {
		t.Error("didn't write 1 line to log")
	}
}

func TestItLogsErrorWhenLevelSetToDebug(t *testing.T) {
	setup()

	logger.SetVerbosity(Debug)

	logger.Error("bongo")

	if len(writer.lines) != 1 {
		t.Error("didn't write 1 line to log")
	}
}

type mockWriter struct {
	lines []string
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.lines = append(m.lines, string(p))
	return 0, nil
}
