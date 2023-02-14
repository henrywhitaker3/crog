package root

import (
	"github.com/henrywhitaker3/crog/internal/cmd/list"
	"github.com/henrywhitaker3/crog/internal/cmd/remote"
	"github.com/henrywhitaker3/crog/internal/cmd/run"
	"github.com/henrywhitaker3/crog/internal/cmd/work"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/spf13/cobra"
)

func NewRootCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "crog",
		Version: "0.2.7",
		Short:   "A CLI tool to setup scheduled tasks and call URLs based on the result, configured in yaml.",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cmd.PersistentFlags().BoolVarP(&log.Log.Verbose, "verbose", "v", false, "Turn on verbose logging")
	cmd.PersistentFlags().StringP("config", "c", "crog.yaml", "Config file path")

	cmd.AddCommand(
		list.NewListCmd(cfg),
		run.NewRunCmd(cfg),
		work.NewWorkCmd(cfg),
		remote.NewRootCmd(cfg),
	)

	return cmd
}
