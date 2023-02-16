package event

import (
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
)

var EventHandler events.EventHandler = events.NewHandler()

func Boot() {
	EventHandler.Register("ActionPreflight", &ActionPreflightLogger{})
	EventHandler.Register("Result", &ResultLogger{})
	EventHandler.Register("ServerStarted", &ServerStartedLogger{})
	EventHandler.Register("SchedulerStarted", &SchedulerStartedLogger{})
	EventHandler.Register("SchedulerStopped", &SchedulerStoppedLogger{})
	EventHandler.Register("ActionScheduled", &ActionScheduledLogger{})
	EventHandler.Register("RunScheduledAction", &RunScheduledActionLogger{})
}

func Emit(event events.Event) {
	log.Log.Debugf("Emitting event %s", event.Tag)
	EventHandler.Trigger(event)
}
