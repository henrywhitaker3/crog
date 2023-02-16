package event

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
)

type ResultLogger struct{}

func (a *ResultLogger) Handle(e events.Event) error {
	r := e.Data.(domain.Result)

	log.Log.Info(
		actionLogFormat(
			r.GetAction(),
			fmt.Sprintf("got exit code: %d", r.GetCode()),
		),
	)
	log.Log.Debug(
		actionLogFormat(
			r.GetAction(),
			fmt.Sprintf("got stdout:\n%s", r.GetStdout()),
		),
	)
	if r.GetErr() != nil {
		log.Log.Error(actionLogFormat(r.GetAction(), r.GetErr().Error()))
	} else {
		log.Log.Info(actionLogFormat(r.GetAction(), "Check passed"))
	}

	return nil
}
