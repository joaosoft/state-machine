package state_machine

import (
	"github.com/joaosoft/logger"
)

// StateMachineOption ...
type StateMachineOption func(stateMachine *stateMachine)

// Reconfigure ...
func (sm *stateMachine) Reconfigure(options ...StateMachineOption) {
	for _, option := range options {
		option(sm)
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) StateMachineOption {
	return func(stateMachine *stateMachine) {
		stateMachine.logger = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) StateMachineOption {
	return func(stateMachine *stateMachine) {
		stateMachine.logger.SetLevel(level)
	}
}

