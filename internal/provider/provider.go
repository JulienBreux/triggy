package provider

import (
	"github.com/julienbreux/triggy/internal/adapter"
	"github.com/julienbreux/triggy/internal/provider/action"
	"github.com/julienbreux/triggy/internal/provider/trigger"
)

// Provider represents the provider interface
type Provider interface {
	Name() string

	Adapter() adapter.Adapter

	Actions() map[string]action.Action
	Triggers() map[string]trigger.Trigger
}
