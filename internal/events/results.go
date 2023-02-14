package events

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

type Result struct {
	Result domain.Result
}

type ResultLogger struct{}

func (a *ResultLogger) Handle(e domain.Event) error {
	ev := e.(Result)

	log.Log.Info(
		actionLogFormat(
			ev.Result.GetAction(),
			fmt.Sprintf("got exit code: %d", ev.Result.GetCode()),
		),
	)
	log.Log.Debug(
		actionLogFormat(
			ev.Result.GetAction(),
			fmt.Sprintf("got stdout:\n%s", ev.Result.GetStdout()),
		),
	)
	if ev.Result.GetErr() != nil {
		log.Log.Error(actionLogFormat(ev.Result.GetAction(), ev.Result.GetErr().Error()))
	} else {
		log.Log.Info(actionLogFormat(ev.Result.GetAction(), "Check passed"))
	}

	return nil
}
