package config

import (
	"encoding/json"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the configuration structure
type Config struct {
	Debug bool `default:"false"`

	LoggerLevel  string `default:"info" envconfig:"logger_level"`
	LoggerFormat string `default:"json" envconfig:"logger_format"`

	ConfigPath string `default:"./" envconfig:"config_path"`
	ConfigFile string `default:"triggy.yml" envconfig:"config_file"`
}

// JSON exports configuration as JSON
func (c *Config) JSON() []byte {
	b, _ := json.Marshal(c)
	return b
}

// New returns the configuration
func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
