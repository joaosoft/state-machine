package state_machine

import (
	"strings"

	"github.com/joaosoft/logger"
)

type TransitionCheckHandler func(args ...interface{}) (bool, error)
type StateMachine struct {
	config                  *StateMachineConfig
	states                  map[int]State
	transitionCheckHandlers map[string]TransitionCheckHandler
	logger                  logger.ILogger
}

type State struct {
	Name        string                         `json:"name"`
	Transitions map[int]TransitionCheckHandler `json:"transitions"`
}

// New ...
func New(options ...StateMachineOption) (*StateMachine, error) {
	config, _, err := NewConfig()

	newStateMachine := &StateMachine{
		states:                  make(map[int]State),
		transitionCheckHandlers: make(map[string]TransitionCheckHandler),
		logger:                  logger.NewLogDefault("state_machine", logger.WarnLevel),
		config:                  config.StateMachine,
	}

	if err != nil {
		newStateMachine.logger.Error(err.Error())
	} else if config.StateMachine != nil {
		level, _ := logger.ParseLevel(config.StateMachine.Log.Level)
		newStateMachine.logger.Debugf("setting log level to %s", level)
		newStateMachine.logger.Reconfigure(logger.WithLevel(level))
	}

	newStateMachine.Reconfigure(options...)

	return newStateMachine, newStateMachine.init()
}

func (stateMachine *StateMachine) init() error {
	for _, stateConfig := range stateMachine.config.States {
		state := State{
			Name:        stateConfig.Name,
			Transitions: make(map[int]TransitionCheckHandler),
		}

		for _, transition := range stateConfig.Transitions {
			var handler TransitionCheckHandler
			if len(strings.TrimSpace(transition.Handler)) > 0 {
				if h, ok := stateMachine.transitionCheckHandlers[transition.Handler]; ok {
					handler = h
				}
			}
			state.Transitions[transition.Id] = handler
		}

		stateMachine.states[stateConfig.Id] = state
	}

	return nil
}

func (stateMachine *StateMachine) CheckTransition(from int, to int, args ...interface{}) (bool, error) {
	if state, ok := stateMachine.states[from]; ok {
		if handler, ok := state.Transitions[to]; ok {
			var err error
			if handler != nil {
				ok, err = handler(args)
			}
			return ok && err == nil, err
		}
	}
	return false, nil
}
