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

	log.Log.Info(actionLogFormat(ev.Action, fmt.Sprintf("executable: %s", ev.Action.GetExecutable())))
	log.Log.Info(actionLogFormat(ev.Action, fmt.Sprintf("args: %s", ev.Action.GetArguments())))

	return nil
}

func actionLogFormat(a domain.Action, value string) string {
	return fmt.Sprintf("[%s] %s", a.GetName(), value)
}
