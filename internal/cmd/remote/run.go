package remote

import (
	"errors"

	"github.com/henrywhitaker3/crog/internal/cli"
	"github.com/henrywhitaker3/crog/internal/client"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewRunCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run crog actions on a remote server",
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg.Remotes) == 0 {
				cli.ErrorExit(errors.New("no remotes configured"))
			}

			selectedRemote, _ := cli.SingleChoice(cli.GetRemoteNames(cfg.Remotes))

			remote, err := cfg.GetRemote(selectedRemote)
			if err != nil {
				cli.ErrorExit(err)
			}
			cli.Printfln("Listing actions on %s", pterm.Green(remote.Name))

			cl, err := client.New(remote.Url)
			if err != nil {
				cli.ErrorExit(err)
			}
			defer cl.Close()

			resp, err := cl.GetActions()
			if err != nil {
				cli.ErrorExit(err)
			}

			selectedAction, _ := pterm.DefaultInteractiveSelect.WithOptions(resp).Show()

			cli.Printfln("Running '%s' action on remote", selectedAction)

			res, err := cl.RunAction(selectedAction)
			if err != nil {
				cli.ErrorExit(err)
			}
			cli.Printfln("Success: %v", pterm.Green(res.Success))

			if log.Verbose {
				cli.Printfln("Command: %s", res.Action.Command)
				cli.Printfln("Code: %d", res.Code)
				cli.Printfln("Stdout:\n%s", res.Stdout)
			}
		},
	}
}
