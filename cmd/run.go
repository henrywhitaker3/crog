/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a check",
	RunE: func(cmd *cobra.Command, args []string) error {
		options := []string{}

		for _, check := range cfg.Checks {
			options = append(options, check.Name)
		}

		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		check, err := cfg.GetCheck(selectedOption)
		if err != nil {
			return err
		}
		pterm.Info.Printfln("Running: %s", pterm.Green(check.Name))

		err = check.Execute()

		pterm.Error.PrintOnErrorf("%s", err)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
