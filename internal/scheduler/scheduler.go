package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/log"
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
		log.Infof("Adding action '%s' to scheduler", action.Name)
		s.scheduler.Cron(action.Cron).Do(runAction, action)
	}

	log.ForceInfo("Starting schduler")
	s.scheduler.StartAsync()

	return nil
}

func (s *Scheduler) Stop() error {
	log.ForceInfo("Stopping scheduler")
	s.scheduler.Stop()
	return nil
}

func runAction(action action.Action) {
	log.ForceInfof("Running schduled command '%s'", action.Name)
	log.ActionPreflight(&action)
	log.LogResult(action.Execute())
}
