package triggy

// Trigger represents a Triggy trigger
type Trigger struct {
	Provider   string
	Trigger    string
	Parameters map[string]interface{}
	Actions    []string
}
