package logger

import (
	errors "github.com/joaosoft/errors"
	goerrors "errors"
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
	*err = goerrors.New(addition.message)

	return addition
}

// ToErrorData
func (addition *Addition) ToErrorData(err *errors.ErrorData) IAddition {
	err.Add(err)

	return addition
}
