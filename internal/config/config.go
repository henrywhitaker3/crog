package config

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/validation"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string          `yaml:"version" required:"true"`
	Actions  []action.Action `yaml:"actions"`
	Verbose  bool            `yaml:"verbose" default:"false"`
	Timezone string          `yaml:"timezone" default:"UTC"`
	Server   ServerConfig    `yaml:"server"`
}

type ServerConfig struct {
	Enabled bool   `yaml:"enabled" default:"true"`
	Listen  string `yaml:"listen" default:":9399"`
}

func (sc *ServerConfig) Validate() error {
	return validation.Validate(sc)
}

func LoadConfig(path string) (*Config, error) {
	// TODO: periodically check the config for changes
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

func (cfg *Config) GetAction(name string) (*action.Action, error) {
	for _, chk := range cfg.Actions {
		if chk.Name == name {
			return &chk, nil
		}
	}

	return nil, fmt.Errorf("could not find action")
}

func (cfg *Config) PrintActionTable() error {
	lines := [][]string{
		{"Name", "Command", "Code", "Cron"},
	}

	for _, action := range cfg.Actions {
		lines = append(lines, []string{action.Name, action.Command, fmt.Sprintf("%d", action.Code), action.Cron})
	}
	return pterm.DefaultTable.WithHasHeader().WithData(lines).Render()
}

func (cfg *Config) Validate() error {
	err := validation.Validate(cfg)
	if err != nil {
		return err
	}

	for i, action := range cfg.Actions {
		if err := action.Validate(); err != nil {
			return fmt.Errorf("actions[%d]: %s", i, err)
		}
	}

	if err := cfg.Server.Validate(); err != nil {
		return err
	}

	return nil
}
