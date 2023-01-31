package remote

import (
	"errors"

	"github.com/henrywhitaker3/crog/internal/cli"
	"github.com/henrywhitaker3/crog/internal/client"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/pb"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewRunCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "run [<remote_server> <remote_action>]",
		Short: "Run crog actions on a remote server",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch len(args) {
			case 0:
				return runInteractive(cfg)
			case 2:
				// TODO: sort out remote/action autocompletion
				return runManual(cfg, args[0], args[1])
			default:
				return errors.New("invalid args")
			}
		},
	}
}

func runManual(cfg *config.Config, remote, action string) error {
	rem, err := cfg.GetRemote(remote)
	if err != nil {
		return err
	}

	cli.Printfln("Running '%s' action on remote %s", action, remote)

	cl, err := client.New(rem.Url)
	if err != nil {
		return err
	}
	resp, err := cl.RunAction(action)
	return handleActionResponse(resp, err)
}

func runInteractive(cfg *config.Config) error {
	if len(cfg.Remotes) == 0 {
		return errors.New("no remotes configured")
	}

	selectedRemote, _ := cli.SingleChoice(cli.GetRemoteNames(cfg.Remotes))

	remote, err := cfg.GetRemote(selectedRemote)
	if err != nil {
		return err
	}
	cli.Printfln("Listing actions on %s", pterm.Green(remote.Name))

	cl, err := client.New(remote.Url)
	if err != nil {
		return err
	}
	defer cl.Close()

	resp, err := cl.GetActions()
	if err != nil {
		return err
	}

	selectedAction, _ := pterm.DefaultInteractiveSelect.WithOptions(resp).Show()

	cli.Printfln("Running '%s' action on remote", selectedAction)

	res, err := cl.RunAction(selectedAction)
	return handleActionResponse(res, err)
}

func handleActionResponse(res *pb.RunActionResponse, err error) error {
	if err != nil {
		return err
	}

	cli.Printfln("Success: %v", pterm.Green(res.Success))

	if log.Verbose {
		cli.Printfln("Command: %s", res.Action.Command)
		cli.Printfln("Code: %d", res.Code)
		cli.Printfln("Stdout:\n%s", res.Stdout)
	}

	return nil
}
