package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Environment *DefaultEnv
	Tasks       []*ConfigTask
}

type DefaultEnv map[string]string

type ConfigTask struct {
	Id          int
	Name        string
	Command     string
	Environment map[string]string
	Service     bool
	Enabled     bool
}

func LoadConfig(file string) (*Config, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg *Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
