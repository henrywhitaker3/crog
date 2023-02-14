package events

import (
	"reflect"
	"sync"
	"testing"

	"github.com/henrywhitaker3/crog/internal/domain"
)

type testEvent struct{}

type testListener struct {
	hasRun bool
}

func (t *testListener) Handle(e domain.Event) error {
	t.hasRun = true
	return nil
}

func TestItRegistersEventsToListeners(t *testing.T) {
	eh := newHandler()
	e := testEvent{}

	if _, ok := eh.listeners[reflect.TypeOf(e)]; ok {
		t.Error("there is a listener setup for the event already")
	}

	eh.Register(e, &testListener{})

	if _, ok := eh.listeners[reflect.TypeOf(e)]; !ok {
		t.Error("there is no listener setup for the event already")
	}
}

func TestItDoesntErrorIfNoHandlerHasBeenRegistered(t *testing.T) {
	eh := newHandler()
	e := testEvent{}

	if _, ok := eh.listeners[reflect.TypeOf(e)]; ok {
		t.Error("there is a listener setup for the event already")
	}

	eh.Trigger(e)
}

func TestItCallsARegisteredListener(t *testing.T) {
	eh := newHandler()
	e := testEvent{}
	l := &testListener{}

	eh.Register(e, l)

	if l.hasRun {
		t.Error("the listener has already run")
	}

	eh.Trigger(e)
	eh.Wait()

	if !l.hasRun {
		t.Error("the listener has not run")
	}
}

func newHandler() eventHandler {
	return eventHandler{
		listeners: map[reflect.Type][]domain.Listener{},
		wg:        &sync.WaitGroup{},
	}
}