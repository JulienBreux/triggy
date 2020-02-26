package config_test

import (
	"os"
	"testing"

	"github.com/julienbreux/triggy/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "debug")
	c, err := config.New()

	assert.Equal(t, "debug", c.LoggerLevel)
	assert.False(t, c.Debug)
	assert.NoError(t, err)
}

func TestJSON(t *testing.T) {
	c, err := config.New()

	assert.Contains(t, string(c.JSON()), "LoggerLevel\":\"info")
	assert.NoError(t, err)
}
