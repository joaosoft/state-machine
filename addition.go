package logger

import (
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
func (addition *Addition) ToError() errors.IErr {
	return errors.New("0", addition.message)
}