package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/event"
	"github.com/henrywhitaker3/events"
)

type Scheduler struct {
	actions   []action.Action
	scheduler *gocron.Scheduler
}

func New(time *time.Location, actions []action.Action) *Scheduler {
	return &Scheduler{
		actions:   actions,
		scheduler: gocron.NewScheduler(time),
	}
}

func (s *Scheduler) Start() error {
	s.scheduler = gocron.NewScheduler(time.UTC)

	for _, action := range s.actions {
		s.scheduler.Cron(action.Cron).Do(runAction, action)
		event.Emit(events.Event{Tag: "ActionScheduled", Data: action})
	}

	s.scheduler.StartAsync()
	event.Emit(events.Event{Tag: "SchedulerStarted", Data: nil})

	return nil
}

func (s *Scheduler) Stop() error {
	event.Emit(events.Event{Tag: "SchedulerStopped", Data: nil})
	s.scheduler.Stop()
	return nil
}

func runAction(action action.Action) {
	event.Emit(events.Event{Tag: "RunScheduledAction", Data: action})
	event.Emit(events.Event{Tag: "ActionPreflight", Data: action})

	res := action.Execute()

	event.Emit(events.Event{Tag: "Result", Data: res})
}
