package events

import (
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

var ActionPreflightHandler eventHandler

type ActionPreflight struct {
	Action domain.Action
}

type ActionPreflightLogger struct{}

func (a *ActionPreflightLogger) Handle(e domain.Event) error {
	ev := e.(ActionPreflight)
	log.ActionPreflight(ev.Action)
	return nil
}
