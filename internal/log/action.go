package log

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
)

func ActionPreflight(a domain.Action) {
	Info(ActionLogFormat(a, fmt.Sprintf("executable: %s", a.GetExecutable())))
	Info(ActionLogFormat(a, fmt.Sprintf("args: %s", a.GetArguments())))
}

func LogResult(res domain.Result) {
	Info(
		ActionLogFormat(
			res.GetAction(),
			fmt.Sprintf("got exit code: %d", res.GetCode()),
		),
	)
	Info(
		ActionLogFormat(
			res.GetAction(),
			fmt.Sprintf("got stdout:\n%s", res.GetStdout()),
		),
	)
	if res.GetErr() != nil {
		logResultFailure(res)
		return
	}
	logResultSuccess(res)
}

func logResultSuccess(res domain.Result) {
	ForceInfo(ActionLogFormat(res.GetAction(), "Check passed"))
}

func logResultFailure(res domain.Result) {
	ForceError(ActionLogFormat(res.GetAction(), res.GetErr().Error()))
}

func ActionLogFormat(a domain.Action, value string) string {
	return fmt.Sprintf("[%s] %s", a.GetName(), value)
}
