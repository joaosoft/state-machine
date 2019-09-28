package state_machine

import (
	"github.com/joaosoft/logger"
)

// stateMachineOption ...
type stateMachineOption func(stateMachine *stateMachine)

// Reconfigure ...
func (sm *stateMachine) Reconfigure(options ...stateMachineOption) {
	for _, option := range options {
		option(sm)
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) stateMachineOption {
	return func(stateMachine *stateMachine) {
		stateMachine.logger = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) stateMachineOption {
	return func(stateMachine *stateMachine) {
		stateMachine.logger.SetLevel(level)
	}
}

