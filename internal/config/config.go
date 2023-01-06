package config

import (
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/henrywhitaker3/go-healthcheck/internal/check"
	"github.com/henrywhitaker3/go-healthcheck/internal/log"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string        `yaml:"version"`
	Checks   []check.Check `yaml:"checks"`
	Verbose  *bool         `yaml:"verbose"`
	Timezone string        `yaml:"timezone"`
}

// TODO: unmarshall to set defaults e.g. utc timezone

func LoadConfig(path string) (*Config, error) {
	log.Infof("Loading config file from %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	if config.Verbose == nil {
		v := false
		config.Verbose = &v
	}

	if structs.HasZero(config) {
		return nil, fmt.Errorf("invalid configuration file")
	}

	return config, nil
}

func (cfg *Config) GetCheck(name string) (*check.Check, error) {
	for _, chk := range cfg.Checks {
		if chk.Name == name {
			return &chk, nil
		}
	}

	return nil, fmt.Errorf("could not find check")
}

func (cfg *Config) PrintCheckTable() error {
	lines := [][]string{
		{"Name", "ID", "Command", "Code"},
	}

	for _, check := range cfg.Checks {
		lines = append(lines, []string{check.Name, check.Id, check.Command, fmt.Sprintf("%d", check.Code)})
	}
	return pterm.DefaultTable.WithHasHeader().WithData(lines).Render()
}

// TODO: unmarshal function for check
