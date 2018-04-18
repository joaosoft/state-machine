package golog

import (
	"errors"
)

type Addition struct {
	message string
}

// newAddition ...
func newAddition(message string) IAddition {
	addition := &Addition{
		message: message,
	}

	return addition
}

// ToError
func (addition *Addition) ToError(err *error) IAddition {
	*err = errors.New(addition.message)
	return addition
}
