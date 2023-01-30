package cli

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/pterm/pterm"
)

func PrintActionTable(actions []action.Action) error {
	lines := [][]string{
		{"Name", "Command", "Code", "Cron"},
	}

	for _, action := range actions {
		lines = append(lines, []string{action.Name, action.Command, fmt.Sprintf("%d", action.Code), action.Cron})
	}
	return pterm.DefaultTable.WithHasHeader().WithData(lines).Render()
}

func GetActionNames(actions []action.Action) []string {
	names := []string{}

	for _, action := range actions {
		names = append(names, action.Name)
	}

	return names
}

func GetRemoteNames(remotes []config.Remote) []string {
	names := []string{}

	for _, remote := range remotes {
		names = append(names, remote.Name)
	}

	return names
}