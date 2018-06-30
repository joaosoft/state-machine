package logger

import (
	"errors"
	goerror "github.com/joaosoft/go-error/app"
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
	*err = errors.New(addition.message)
	return addition
}

// ToErrorData
func (addition *Addition) ToErrorData(err *goerror.ErrorData) IAddition {
	newErr := errors.New(addition.message)
	err.AddError(newErr)

	return addition
}
