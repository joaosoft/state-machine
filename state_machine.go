package state_machine

import (
	"strings"

	"github.com/joaosoft/manager"

	"github.com/joaosoft/logger"
)

var (
	stateMachineInstance, _ = New()
)

type StateConfigList []StateConfig

type StateConfig struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Transitions []TransitionConfig `json:"transitions"`
}

type TransitionConfig struct {
	Id      int    `json:"id"`
	Handler string `json:"handler"`
}

type StateMap map[int]State

type TransitionCheckHandler func(args ...interface{}) (bool, error)
type StateMachine struct {
	config                  *StateMachineConfig
	stateMachineMap         map[string]StateMap
	transitionCheckHandlers map[string]TransitionCheckHandler
	logger                  logger.ILogger
}

type State struct {
	Name        string                         `json:"name"`
	Transitions map[int]TransitionCheckHandler `json:"transitions"`
}

func New(options ...StateMachineOption) (*StateMachine, error) {
	config, _, err := NewConfig()

	newStateMachine := &StateMachine{
		stateMachineMap:         make(map[string]StateMap),
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

	return newStateMachine, nil
}

func (stateMachine *StateMachine) Add(name, file string) error {
	config := StateConfigList{}
	_, err := manager.NewSimpleConfig(file, &config)
	if err != nil {
		return err
	}

	states := make(StateMap)
	for _, stateConfig := range config {
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

		states[stateConfig.Id] = state
	}
	stateMachine.stateMachineMap[name] = states

	return nil
}

func (stateMachine *StateMachine) AddTransitionCheckHandler(name string, handler TransitionCheckHandler) *StateMachine {
	stateMachine.transitionCheckHandlers[name] = handler
	return stateMachine
}

func (stateMachine *StateMachine) CheckTransition(name string, from int, to int, args ...interface{}) (bool, error) {
	if states, ok := stateMachine.stateMachineMap[name]; ok {
		if state, ok := states[from]; ok {
			if handler, ok := state.Transitions[to]; ok {
				var err error
				if handler != nil {
					ok, err = handler(args)
				}
				return ok && err == nil, err
			}
		}
	}
	return false, nil
}
