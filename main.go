/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package main

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/cmd/root"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/events"
	"github.com/henrywhitaker3/crog/internal/log"
)

func main() {
	cfgPath := getConfigFilePath(os.Args[1:])
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}
	if cfg.Verbose || isVerboseFlagSet(os.Args[1:]) {
		log.Verbose = true
	}

	events.Boot()

	if err := root.NewRootCmd(cfg).Execute(); err != nil {
		os.Exit(1)
	}
}

func getConfigFilePath(args []string) string {
	path := "crog.yaml"

	for i, arg := range args {
		if arg == "-c" || arg == "--config" {
			path = args[i+1]
			break
		}
	}

	return path
}

func isVerboseFlagSet(args []string) bool {
	for _, arg := range args {
		if arg == "-v" || arg == "--verbose" {
			return true
		}
	}

	return false
}
