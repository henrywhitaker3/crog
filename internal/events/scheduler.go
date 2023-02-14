package events

import (
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

type SchedulerStarted struct {
	Action domain.Action
}

type SchedulerStartedLogger struct{}

func (s *SchedulerStartedLogger) Handle(e domain.Event) error {
	log.Log.Info("Scheduler started")
	return nil
}

type SchedulerStopped struct {
	Action domain.Action
}

type SchedulerStoppedLogger struct{}

func (s *SchedulerStoppedLogger) Handle(e domain.Event) error {
	log.Log.Info("Scheduler stopped")
	return nil
}

type ActionScheduled struct {
	Action domain.Action
}

type ActionScheduledLogger struct{}

func (s *ActionScheduled) Handle(e domain.Event) error {
	ev := e.(ActionScheduled)
	log.Log.Infof("Scheduling action '%s' at '%s'", ev.Action.GetName(), ev.Action.GetCron())
	return nil
}

type RunScheduledAction struct {
	Action domain.Action
}

type RunScheduledActionLogger struct{}

func (s *RunScheduledActionLogger) Handle(e domain.Event) error {
	ev := e.(RunScheduledAction)
	log.Log.Infof("Running scheduled action %s", ev.Action.GetName())
	return nil
}
