package triggy

import (
	"encoding/json"
)

// Triggy represents the Triggy configuration file
type Triggy struct {
	Version string `json:",omitempty"`

	Providers map[string]Provider `json:",omitempty"`
	Triggers  map[string]Trigger  `json:",omitempty"`
	Actions   map[string]Action   `json:",omitempty"`
}

// JSON exports Triggy as JSON
func (t *Triggy) JSON() []byte {
	b, _ := json.Marshal(t)
	return b
}
