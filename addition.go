package logger

import (
	errors "github.com/joaosoft/errors"
)

type Addition struct {
	message string
}

// NewAddition ...
func NewAddition(message string) IAddition {
	addition := &Addition{
		message: message,
	}

	return addition
}

// ToError
func (addition *Addition) ToError(err *error) IAddition {
	*err = errors.New("0", addition.message)

	return addition
}

// ToErr
func (addition *Addition) ToErr(err *errors.Err) IAddition {
	err.Add(err)

	return addition
}
