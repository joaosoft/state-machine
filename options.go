package state_machine

import (
	"github.com/joaosoft/logger"
)

// StateMachineOption ...
type StateMachineOption func(stateMachine *StateMachine)

// Reconfigure ...
func (stateMachine *StateMachine) Reconfigure(options ...StateMachineOption) {
	for _, option := range options {
		option(stateMachine)
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) StateMachineOption {
	return func(stateMachine *StateMachine) {
		stateMachine.logger = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) StateMachineOption {
	return func(stateMachine *StateMachine) {
		stateMachine.logger.SetLevel(level)
	}
}

// WithTransitionCheckHandler ...
func WithTransitionCheckHandler(name string, handler TransitionCheckHandler) StateMachineOption {
	return func(stateMachine *StateMachine) {
		stateMachine.transitionCheckHandlers[name] = handler
	}
}
