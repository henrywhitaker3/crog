package remote

import (
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remote",
		Short: "Run crog actions on a remote server",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cmd.AddCommand(
		NewRunCmd(cfg),
	)

	return cmd
}
