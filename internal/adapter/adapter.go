package adapter

import (
	"github.com/julienbreux/triggy/internal/provider/action"
	"github.com/julienbreux/triggy/internal/provider/trigger"
)

// Adapter represents the adapter interface
type Adapter interface {
	Name() string
	Version() string

	Configure(map[string]interface{}) error

	Connect() error
	Disconnect() error

	HealthCheck() error

	Actions() map[string]action.Action
	Trigger() map[string]trigger.Trigger
}
