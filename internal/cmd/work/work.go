package work

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/events"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/scheduler"
	"github.com/henrywhitaker3/crog/internal/server"
	"github.com/spf13/cobra"
)

func NewWorkCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "work",
		Short: "Run the check daemon",
		Long: `Run the program as a daemon.

	For each check where a cron value is specified, the daemon will run them when accordingly.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			workers, err := getWorkers(cfg)
			if err != nil {
				return err
			}

			for _, worker := range workers {
				if err := worker.Start(); err != nil {
					return err
				}
			}

			log.Log.Info("Registering signal handlers for graceful shutdown")

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
			<-sig

			for _, worker := range workers {
				if err := worker.Stop(); err != nil {
					return err
				}
			}

			events.EventHandler.Wait()

			return nil
		},
	}
}

func getWorkers(cfg *config.Config) ([]domain.Worker, error) {
	workers := []domain.Worker{}

	if cfg.Server.Enabled {
		serv, err := server.New(cfg)
		if err != nil {
			return workers, err
		}
		workers = append(workers, serv)
	}

	workers = append(workers, scheduler.New(time.UTC, cfg.Actions))

	return workers, nil
}
