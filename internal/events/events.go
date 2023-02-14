package events

import "github.com/henrywhitaker3/crog/internal/domain"

type eventHandler struct {
	listeners []domain.Listener
}

func (a *eventHandler) Register(l domain.Listener) {
	a.listeners = append(a.listeners, l)
}

func (a *eventHandler) Trigger(e domain.Event) {
	for _, listener := range a.listeners {
		go listener.Handle(e)
	}
}
