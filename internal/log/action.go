package log

import (
	"fmt"

	"github.com/google/shlex"
	"github.com/henrywhitaker3/crog/internal/action"
)

func ActionPreflight(a *action.Action) {
	args, _ := shlex.Split(a.Command)
	bin := args[0]
	args = args[1:]
	Info(ActionLogFormat(a, fmt.Sprintf("executable: %s", bin)))
	Info(ActionLogFormat(a, fmt.Sprintf("args: %s", args)))
}

func LogResult(res *action.Result) {
	Info(
		ActionLogFormat(
			res.Action,
			fmt.Sprintf("got exit code: %d", res.Code),
		),
	)
	Info(
		ActionLogFormat(
			res.Action,
			fmt.Sprintf("got stdout:\n%s", res.Stdout),
		),
	)
	if res.Err != nil {
		logResultFailure(res)
		return
	}
	logResultSuccess(res)
}

func logResultSuccess(res *action.Result) {
	ForceInfo(ActionLogFormat(res.Action, "Check passed"))
}

func logResultFailure(res *action.Result) {
	ForceError(ActionLogFormat(res.Action, res.Err.Error()))
}

func ActionLogFormat(a *action.Action, value string) string {
	return fmt.Sprintf("[%s] %s", a.Name, value)
}
