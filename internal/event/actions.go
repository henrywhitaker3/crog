package event

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
)

type ActionPreflightLogger struct{}

func (al *ActionPreflightLogger) Handle(e events.Event) error {
	a := e.Data.(domain.Action)

	log.Log.Info(actionLogFormat(a, "Running action"))
	log.Log.Debug(actionLogFormat(a, fmt.Sprintf("command: %s", a.GetCommand())))

	return nil
}

func actionLogFormat(a domain.Action, value string) string {
	return fmt.Sprintf("[%s] %s", a.GetName(), value)
}
