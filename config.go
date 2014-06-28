package main

import (
	"encoding/json"
	"github.com/ian-kent/go-log/log"
	"io/ioutil"
)

var ConfigEnvironment = make(map[string]string)
var ConfigWorkspaces = make(map[string]*ConfigWorkspace)

type ConfigWorkspace struct {
	Environment map[string]string
	Name        string
	Tasks       []*ConfigTask
}

type ConfigTask struct {
	Id          int
	Name        string
	Command     string
	Environment map[string]string
	Service     bool
	Executor    []string
	Stdout      string
	Stderr      string
}

func LoadConfigFile(file string) (*ConfigWorkspace, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg *ConfigWorkspace
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadConfig(global string, workspaces []string) {
	// Load global environment
	log.Info("Loading global environment file: %s", global)
	cfg, err := LoadConfigFile(global)
	if err != nil {
		log.Error("Error loading global configuration: %s", err.Error())
	}
	if cfg != nil {
		for k, v := range cfg.Environment {
			ConfigEnvironment[k] = v
		}
	}

	// Load workspaces
	for _, conf := range workspaces {
		log.Info("Loading workspace file: %s", conf)
		cfg, err := LoadConfigFile(conf)
		if err != nil {
			log.Error("Error loading global configuration: %s", err.Error())
		}
		if cfg != nil {
			ConfigWorkspaces[cfg.Name] = cfg
		}
	}
}
