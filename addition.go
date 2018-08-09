package logger

import (
	errs "errors"

	"github.com/joaosoft/errors"
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
	*err = errs.New(addition.message)

	return addition
}

// ToErr
func (addition *Addition) ToErr(err errors.IErr) IAddition {
	err.Add(errors.New("0", addition.message))

	return addition
}
