package trigger

// Trigger represents the trigger interface
type Trigger interface {
	Name() string

	Trigger(parameters map[string]interface{}) error
}
