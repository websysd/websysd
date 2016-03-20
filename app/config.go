package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ian-kent/go-log/log"
)

// GlobalConfigWorkspace is the global workspace config
var GlobalConfigWorkspace *ConfigWorkspace

// ConfigWorkspaces is the user workspace configs
var ConfigWorkspaces = make(map[string]*ConfigWorkspace)

// ConfigWorkspace is the config for a workspace
type ConfigWorkspace struct {
	Functions          map[string]*ConfigFunction
	Environment        map[string]string
	Name               string
	Tasks              []*ConfigTask
	IsLocked           bool
	Columns            map[string]map[string][]string
	InheritEnvironment bool
}

// ConfigFunction is the config for a function
type ConfigFunction struct {
	Args     []string
	Command  string
	Executor []string
}

// ConfigTask is the config for a task
type ConfigTask struct {
	ID          int
	Name        string
	Command     string
	Environment map[string]string
	Service     bool
	Executor    []string
	Stdout      string
	Stderr      string
	Metadata    map[string]string
	Pwd         string
}

// LoadConfigFile loads a config file and returns a ConfigWorkspace
func LoadConfigFile(file string) (*ConfigWorkspace, error) {
	var b []byte
	var err error
	var locked bool

	if strings.HasPrefix(file, "http://") ||
		strings.HasPrefix(file, "https://") {
		b, err = readHTTP(file)
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

func readHTTP(url string) ([]byte, error) {
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

// LoadConfig loads a set of config files
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
