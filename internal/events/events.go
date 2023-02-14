package events

import (
	"reflect"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

func Boot() {
	ActionPreflightHandler.Register(&ActionPreflightLogger{})
	ResultHandler.Register(&ResultLogger{})
	ServerStartedHandler.Register(&ServerStartedLogger{})
}

func Emit(handler domain.EventHandler, event domain.Event) {
	log.Infof("Emitting event %s using handler %s", reflect.TypeOf(event), reflect.TypeOf(handler))
	handler.Trigger(event)
}

type eventHandler struct {
	listeners []domain.Listener
}

func (a *eventHandler) Register(l domain.Listener) {
	log.Infof("Registering listener %s to handler %s", reflect.TypeOf(l), reflect.TypeOf(a))
	a.listeners = append(a.listeners, l)
}

func (a eventHandler) Trigger(e domain.Event) {
	for _, listener := range a.listeners {
		go listener.Handle(e)
	}
}
