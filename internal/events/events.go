package events

import (
	"reflect"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

var EventHandler eventHandler = eventHandler{listeners: map[reflect.Type][]domain.Listener{}}

func Boot() {
	EventHandler.Register(ServerStarted{}, &ServerStartedLogger{})
	EventHandler.Register(Result{}, &ResultLogger{})
	EventHandler.Register(ActionPreflight{}, &ActionPreflightLogger{})
}

func Emit(event domain.Event) {
	log.Infof("Emitting event %s", reflect.TypeOf(event))
	EventHandler.Trigger(event)
}

type eventHandler struct {
	listeners map[reflect.Type][]domain.Listener
}

func (a *eventHandler) Register(e domain.Event, l domain.Listener) {
	log.Infof("Registering event %s to listener %s", reflect.TypeOf(e), reflect.TypeOf(l))
	a.listeners[reflect.TypeOf(e)] = append(a.listeners[reflect.TypeOf(e)], l)
}

func (a eventHandler) Trigger(e domain.Event) {
	for _, listener := range a.getListeners(e) {
		go listener.Handle(e)
	}
}

func (a eventHandler) getListeners(e domain.Event) []domain.Listener {
	return a.listeners[reflect.TypeOf(e)]
}
