package error

import "fmt"

// NotFoundError represents the not found error
type NotFoundError struct {
	Adapter string
}

// NewNotFoundError creates not found error
func NewNotFoundError(adapter string) error {
	return &NotFoundError{adapter}
}

// Error returns the error stringified
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("adapter: %s not found", e.Adapter)
}
