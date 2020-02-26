package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents the configuration structure
type Config struct {
	Debug bool `default:"false"`

	LoggerLevel string `default:"info" envconfig:"logger_level"`
}

// New returns the configuration
func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
