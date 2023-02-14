package events

import (
	"reflect"
	"sync"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

var EventHandler eventHandler = eventHandler{
	listeners: map[reflect.Type][]domain.Listener{},
	wg:        &sync.WaitGroup{},
}

func Boot() {
	EventHandler.Register(ServerStarted{}, &ServerStartedLogger{})
	EventHandler.Register(Result{}, &ResultLogger{})
	EventHandler.Register(ActionPreflight{}, &ActionPreflightLogger{})
	EventHandler.Register(ActionScheduled{}, &ActionScheduled{})
	EventHandler.Register(RunScheduledAction{}, &RunScheduledActionLogger{})
	EventHandler.Register(SchedulerStarted{}, &SchedulerStartedLogger{})
	EventHandler.Register(SchedulerStopped{}, &SchedulerStoppedLogger{})
}

func Emit(event domain.Event) {
	log.Log.Infof("Emitting event %s", reflect.TypeOf(event))
	EventHandler.Trigger(event)
}

type eventHandler struct {
	listeners map[reflect.Type][]domain.Listener
	wg        *sync.WaitGroup
}

func (a *eventHandler) Register(e domain.Event, l domain.Listener) {
	log.Log.Infof("Registering event %s to listener %s", reflect.TypeOf(e), reflect.TypeOf(l))
	a.listeners[reflect.TypeOf(e)] = append(a.listeners[reflect.TypeOf(e)], l)
}

func (a eventHandler) Trigger(e domain.Event) {
	for _, listener := range a.getListenersForEvent(e) {
		a.wg.Add(1)
		go func(l domain.Listener) {
			l.Handle(e)
			a.wg.Done()
		}(listener)
	}
}

func (a *eventHandler) Wait() {
	a.wg.Wait()
}

func (a eventHandler) getListenersForEvent(e domain.Event) []domain.Listener {
	return a.listeners[reflect.TypeOf(e)]
}
