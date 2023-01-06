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
	"github.com/henrywhitaker3/go-healthcheck/internal/log"
	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Run the check daemon",
	Long: `Run the program as a daemon.

	For each check where a cron value is specified, the daemon will run them when accordingly.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := gocron.NewScheduler(time.UTC)

		for _, check := range cfg.Checks {
			log.Infof("Adding check '%s' to scheduler", check.Name)
			s.Cron(check.Cron).Do(func() {
				log.ForceInfof("Running schduled command '%s'", check.Name)
				check.Execute()
			})
		}

		log.ForceInfo("Starting schduler")
		s.StartAsync()

		log.Info("Registering signal handlers for graceful shutdown")

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-sig

		log.ForceInfo("Stopping scheduler")
		s.Stop()
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
