/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package cmd

import (
	"github.com/henrywhitaker3/crog/internal/client"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Run crog actions on a remote server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(cfg.Remotes) == 0 {
			pterm.Error.Println("No remotes configured")
			return
		}

		remotes := []string{}

		for _, remote := range cfg.Remotes {
			remotes = append(remotes, remote.Name)
		}

		selectedRemote, _ := pterm.DefaultInteractiveSelect.WithOptions(remotes).Show()

		remote, err := cfg.GetRemote(selectedRemote)
		if err != nil {
			pterm.Error.PrintOnErrorf("%s", err)
			return
		}
		pterm.Info.Printfln("Listing actions on %s", pterm.Green(remote.Name))

		cl, err := client.New(remote.Url)
		if err != nil {
			pterm.Error.PrintOnErrorf("%s", err)
			return
		}
		defer cl.Close()

		resp, err := cl.GetActions()
		if err != nil {
			pterm.Error.PrintOnErrorf("%s", err)
			return
		}

		selectedAction, _ := pterm.DefaultInteractiveSelect.WithOptions(resp).Show()

		pterm.Info.Printfln("Running '%s' action on remote", selectedAction)

		res, err := cl.RunAction(selectedAction)
		if err != nil {
			pterm.Error.PrintOnErrorf("%s", err)
			return
		}
		pterm.Info.Printfln("Command: %s", res.Command)
		pterm.Info.Printfln("Code: %d", res.Code)
		pterm.Info.Printfln("Stdout:\n%s", res.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
