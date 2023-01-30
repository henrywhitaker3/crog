/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/server"
	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Run the check daemon",
	Long: `Run the program as a daemon.

	For each check where a cron value is specified, the daemon will run them when accordingly.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := gocron.NewScheduler(time.UTC)

		for _, action := range cfg.Actions {
			log.Infof("Adding action '%s' to scheduler", action.Name)
			s.Cron(action.Cron).Do(runAction, action)
		}

		log.ForceInfo("Starting schduler")
		s.StartAsync()

		if cfg.Server.Enabled {
			log.ForceInfo("Starting grpc server")
			gs, err := server.New(cfg)
			if err != nil {
				return err
			}
			go gs.Listen()

			defer gs.Close()
			defer log.ForceInfo("Stopping grpc server")
		}

		log.Info("Registering signal handlers for graceful shutdown")

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-sig

		log.ForceInfo("Stopping scheduler")
		s.Stop()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}

func runAction(action action.Action) {
	log.ForceInfof("Running schduled command '%s'", action.Name)
	action.Execute()
}
