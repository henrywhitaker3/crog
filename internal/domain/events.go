package domain

type Event interface{}

type Listener interface {
	Handle(Event) error
}

type EventHandler interface {
	Register(Listener)
	Trigger(Event)
	Wait()
}
