package triggy

// Triggy represents the triggy configuration file
type Triggy struct {
	Version string

	Providers map[string]Provider
	Triggers  map[string]Trigger
	Actions   map[string]Action
}
