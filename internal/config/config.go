package config

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/validation"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string          `yaml:"version" required:"true"`
	Actions  []action.Action `yaml:"actions"`
	Verbose  bool            `yaml:"verbose" default:"false"`
	Timezone string          `yaml:"timezone" default:"UTC"`
	Server   ServerConfig    `yaml:"server"`
	Remotes  []Remote        `yaml:"remotes"`
}

type ServerConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Listen  string `yaml:"listen" default:":9399"`
}

func (sc *ServerConfig) Validate() error {
	return validation.Validate(sc)
}

type Remote struct {
	Name string `yaml:"name" required:"true"`
	Url  string `yaml:"url" required:"true"`
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

func (cfg *Config) GetRemote(name string) (*Remote, error) {
	for _, chk := range cfg.Remotes {
		if chk.Name == name {
			return &chk, nil
		}
	}

	return nil, fmt.Errorf("could not find remote")
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

	for i, remote := range cfg.Remotes {
		if err := validation.Validate(remote); err != nil {
			return fmt.Errorf("remotes[%d]: %s", i, err)
		}
	}

	if err := cfg.Server.Validate(); err != nil {
		return err
	}

	return nil
}
