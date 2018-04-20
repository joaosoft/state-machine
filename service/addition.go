package golog

import (
	"errors"

	"github.com/joaosoft/go-manager/service"
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

// ToErrorData
func (addition *Addition) ToErrorData(err *gomanager.ErrorData) IAddition {
	newErr := errors.New(addition.message)
	err.Add(newErr)

	return addition
}
