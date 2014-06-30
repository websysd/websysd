package main

import (
	"encoding/json"
	"github.com/ian-kent/go-log/log"
	"io/ioutil"
	"net/http"
	"strings"
)

var GlobalConfigWorkspace *ConfigWorkspace
var ConfigWorkspaces = make(map[string]*ConfigWorkspace)

type ConfigWorkspace struct {
	Functions          map[string]*ConfigFunction
	Environment        map[string]string
	Name               string
	Tasks              []*ConfigTask
	IsLocked           bool
	Columns            map[string]map[string][]string
	InheritEnvironment bool
}

type ConfigFunction struct {
	Args     []string
	Command  string
	Executor []string
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
	Metadata    map[string]string
}

func LoadConfigFile(file string) (*ConfigWorkspace, error) {
	var b []byte
	var err error
	var locked bool

	if strings.HasPrefix(file, "http://") ||
		strings.HasPrefix(file, "https://") {
		b, err = ReadHttp(file)
		locked = true
	} else {
		b, err = ioutil.ReadFile(file)
	}

	if err != nil {
		return nil, err
	}

	var cfg *ConfigWorkspace
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	cfg.IsLocked = locked

	return cfg, nil
}

func ReadHttp(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()
	return b, nil
}

func LoadConfig(global string, workspaces []string) {
	// Load global environment
	log.Info("Loading global environment file: %s", global)
	cfg, err := LoadConfigFile(global)
	if err != nil {
		log.Error("Error loading global configuration: %s", err.Error())
	}
	if cfg != nil {
		GlobalConfigWorkspace = cfg
	}

	// Load workspaces
	for _, conf := range workspaces {
		log.Info("Loading workspace file: %s", conf)
		cfg, err := LoadConfigFile(conf)
		if err != nil {
			log.Error("Error loading workspace: %s", err.Error())
		}
		if cfg != nil {
			ConfigWorkspaces[cfg.Name] = cfg
		}
	}
}
