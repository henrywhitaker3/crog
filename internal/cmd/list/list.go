package list

import (
	"github.com/henrywhitaker3/crog/internal/cli"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/spf13/cobra"
)

func NewListCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List the configured actions",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cli.PrintActionTable(cfg.Actions)

			return err
		},
	}
}
