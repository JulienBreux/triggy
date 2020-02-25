package action

// Action represents the action interface
type Action interface {
	Name() string

	Call(parameters map[string]interface{}) error
}
