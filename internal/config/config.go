package config

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/check"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/validation"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string        `yaml:"version" required:"true"`
	Checks   []check.Check `yaml:"checks"`
	Verbose  bool          `yaml:"verbose" default:"false"`
	Timezone string        `yaml:"timezone" default:"UTC"`
}

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

	if err := config.Validate(); err != nil {
		return nil, err
	}

	// if structs.HasZero(config) {
	// 	return nil, fmt.Errorf("invalid configuration file")
	// }

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

func (cfg *Config) Validate() error {
	err := validation.Validate(cfg)
	if err != nil {
		return err
	}

	for i, check := range cfg.Checks {
		if err := check.Validate(); err != nil {
			return fmt.Errorf("checks[%d]: %s", i, err)
		}
	}

	return nil
}
