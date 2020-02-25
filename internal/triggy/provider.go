package triggy

import (
	"errors"

	adapter_error "github.com/julienbreux/triggy/internal/adapter/error"
	"github.com/julienbreux/triggy/pkg/slice"
)

var (
	adapters = []string{"rabbitmq", "redis"}

	// ErrAdapterNotFound is returned for not found adapters
	ErrAdapterNotFound = errors.New("adapter: not found adapter")
)

// Provider represents a triggy provider
type Provider struct {
	Enabled bool
	Adapter string

	Parameters map[string]interface{}
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (pvr *Provider) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var provider struct {
		Adapter    string
		Parameters map[string]interface{}
	}

	if err := unmarshal(&provider); err != nil {
		return err
	}

	if !slice.Contains(provider.Adapter, adapters) {
		return adapter_error.NewNotFoundError(provider.Adapter)
	}

	pvr.Adapter = provider.Adapter
	pvr.Parameters = provider.Parameters

	return nil
}
