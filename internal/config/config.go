package config

import (
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/henrywhitaker3/go-healthcheck/internal/check"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string        `yaml:"version"`
	Checks  []check.Check `yaml:"checks"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	if structs.HasZero(config) {
		return nil, fmt.Errorf("invalid configuration file")
	}

	return config, nil
}

func (cfg *Config) GetCheck(name string) (*check.Check, error) {
	var check *check.Check

	for _, chk := range cfg.Checks {
		if chk.Name == name {
			check = &chk
		}
	}

	if check == nil {
		return nil, fmt.Errorf("could not find check")
	}

	return check, nil
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
