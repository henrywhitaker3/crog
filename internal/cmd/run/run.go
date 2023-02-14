package run

import (
	"github.com/henrywhitaker3/crog/internal/cli"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/events"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewRunCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run a check",
		Run: func(cmd *cobra.Command, args []string) {
			selectedOption, _ := cli.SingleChoice(cli.GetActionNames(cfg.Actions))

			action, err := cfg.GetAction(selectedOption)
			if err != nil {
				cli.ErrorExit(err)
			}
			cli.Printfln("Running: %s", pterm.Green(action.Name))

			if log.Log.GetVerbosity() >= log.Debug {
				events.Emit(events.ActionPreflight{Action: action})
			}

			res := action.Execute()

			if log.Log.GetVerbosity() >= log.Debug {
				events.Emit(events.Result{Result: res})
			}

			if res.GetErr() != nil {
				cli.ErrorExit(res.GetErr())
			}

			events.EventHandler.Wait()
		},
	}
}
