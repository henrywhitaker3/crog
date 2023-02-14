package events

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

type ActionPreflight struct {
	Action domain.Action
}

type ActionPreflightLogger struct{}

func (a *ActionPreflightLogger) Handle(e domain.Event) error {
	ev := e.(ActionPreflight)

	log.Log.Info(actionLogFormat(ev.Action, "Running action"))
	log.Log.Debug(actionLogFormat(ev.Action, fmt.Sprintf("command: %s", ev.Action.GetCommand())))

	return nil
}

func actionLogFormat(a domain.Action, value string) string {
	return fmt.Sprintf("[%s] %s", a.GetName(), value)
}
