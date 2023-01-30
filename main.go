/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package main

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/cmd/root"
	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/log"
	flag "github.com/spf13/pflag"
)

func main() {
	cfgPath := flag.StringP("config", "c", "crog.yaml", "config file (default is crog.yaml)")

	cfg, err := config.LoadConfig(*cfgPath)
	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}
	if cfg.Verbose {
		log.Verbose = true
	}

	if err := root.NewRootCmd(cfg).Execute(); err != nil {
		os.Exit(1)
	}
}
