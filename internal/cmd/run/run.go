package run

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/cli"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/event"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
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
			fmt.Println(action.GetTries())

			if log.Log.GetVerbosity() >= log.Debug {
				event.Emit(events.Event{Tag: "ActionPreflight", Data: action})
			}

			res := action.Execute()

			if log.Log.GetVerbosity() >= log.Debug {
				event.Emit(events.Event{Tag: "Result", Data: res})
			}

			if res.GetErr() != nil {
				cli.ErrorExit(res.GetErr())
			}
		},
	}
}
