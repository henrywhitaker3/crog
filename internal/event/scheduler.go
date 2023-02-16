package event

import (
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
)

type SchedulerStartedLogger struct{}

func (s *SchedulerStartedLogger) Handle(e events.Event) error {
	log.Log.Info("Scheduler started")
	return nil
}

type SchedulerStoppedLogger struct{}

func (s *SchedulerStoppedLogger) Handle(e events.Event) error {
	log.Log.Info("Scheduler stopped")
	return nil
}

type ActionScheduledLogger struct{}

func (s *ActionScheduledLogger) Handle(e events.Event) error {
	a := e.Data.(domain.Action)
	log.Log.Debugf("Scheduling action '%s' at '%s'", a.GetName(), a.GetCron())
	return nil
}

type RunScheduledActionLogger struct{}

func (s *RunScheduledActionLogger) Handle(e events.Event) error {
	a := e.Data.(domain.Action)
	log.Log.Infof("Running scheduled action %s", a.GetName())
	return nil
}
