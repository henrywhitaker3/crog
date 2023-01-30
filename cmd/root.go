/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "crog",
	Version: "0.2.2",
	Short:   "A CLI tool to setup scheduled tasks and call URLs based on the result, configured in yaml.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var cfgPath string
var cfg *config.Config

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "crog.yaml", "config file (default is crog.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&log.Verbose, "verbose", "v", false, "Turn on verbose logging")
}

func initConfig() {
	var err error
	cfg, err = config.LoadConfig(cfgPath)
	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}
	if cfg.Verbose {
		log.Verbose = true
	}
}
