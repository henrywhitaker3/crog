package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/events"
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
		events.Emit(events.ActionScheduled{Action: &action})
	}

	s.scheduler.StartAsync()
	events.Emit(events.SchedulerStarted{})

	return nil
}

func (s *Scheduler) Stop() error {
	events.Emit(events.SchedulerStopped{})
	s.scheduler.Stop()
	return nil
}

func runAction(action action.Action) {
	events.Emit(events.RunScheduledAction{Action: &action})
	events.Emit(events.ActionPreflight{Action: &action})

	res := action.Execute()

	events.Emit(events.Result{Result: res})
}
