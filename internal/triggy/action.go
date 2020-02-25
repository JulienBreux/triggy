package triggy

// Action represents a Triggy action
type Action struct {
	Provider string
	Action   string

	Parameters map[string]interface{}
}
