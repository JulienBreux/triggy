package manager

import (
	"github.com/julienbreux/triggy/internal/config"
	"github.com/julienbreux/triggy/internal/triggy"
)

// Manager represents the manager interface
type Manager interface {
	Start() error
	Stop() error
}

type mngr struct {
	t *triggy.Triggy
	c *config.Config
}

// New creates a new manager instance
func New(t *triggy.Triggy, c *config.Config) Manager {
	return &mngr{}
}

// Start starts the manager
func (m *mngr) Start() error {
	return nil
}

// Stop stops the manager
func (m *mngr) Stop() error {
	return nil
}
