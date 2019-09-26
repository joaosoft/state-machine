package state_machine

import (
	"errors"
	"fmt"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

var (
	stateMachineInstance, _ = New()
)

func New(options ...StateMachineOption) (*StateMachine, error) {
	config, _, err := NewConfig()

	newStateMachine := &StateMachine{
		stateMachineMap:     make(map[string]StateMap),
		userStateMachineMap: make(UserStateMachine),
		handlerMap: HandlerMap{
			Check:   make(map[string]CheckHandler),
			Execute: make(map[string]ExecuteHandler),
			Events: EventMap{
				Success: make(map[string]EventHandler),
				Error:   make(map[string]EventHandler),
			},
		},
		logger: logger.NewLogDefault("state_machine", logger.WarnLevel),
		config: config.StateMachine,
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

func (sm *StateMachine) Add(stateMachine string, file string) error {
	config := StateMachineCfg{}
	_, err := manager.NewSimpleConfig(file, &config)
	if err != nil {
		return err
	}

	// load states
	states := make(StateMap)
	for _, stateCfg := range config.StateMachine {
		state := &State{
			Id:            stateCfg.Id,
			Name:          stateCfg.Name,
			TransitionMap: make(TransitionMap),
		}

		for _, transitionCfg := range stateCfg.Transitions {
			transition := &Transition{
				Id: transitionCfg.Id,
			}

			// check
			for _, name := range transitionCfg.Check {
				if handler, ok := sm.handlerMap.Check[name]; ok {
					transition.Handler.Check = append(transition.Handler.Check, handler)
				}
			}

			// execute
			for _, name := range transitionCfg.Execute {
				if handler, ok := sm.handlerMap.Execute[name]; ok {
					transition.Handler.Execute = append(transition.Handler.Execute, handler)
				}
			}

			// events
			// -- success
			for _, name := range transitionCfg.Events.Success {
				if handler, ok := sm.handlerMap.Events.Success[name]; ok {
					transition.Handler.Events.Success = append(transition.Handler.Events.Success, handler)
				}
			}

			// -- error
			for _, name := range transitionCfg.Events.Error {
				if handler, ok := sm.handlerMap.Events.Error[name]; ok {
					transition.Handler.Events.Error = append(transition.Handler.Events.Error, handler)
				}
			}

			state.TransitionMap[transitionCfg.Id] = transition
		}

		states[stateCfg.Id] = state
	}

	// load all transition names
	for _, state := range states {
		for idTransition, transition := range state.TransitionMap {
			if s, ok := states[idTransition]; ok {
				transition.Name = s.Name
			} else {
				return errors.New(fmt.Sprintf("state not found %d", idTransition))
			}
		}
	}

	sm.stateMachineMap[stateMachine] = states

	// load users
	for user, statesCfg := range config.Users {
		stateMap := make(StateMap)

		for _, stateCfg := range statesCfg {
			var state *State
			if s, ok := states[stateCfg.Id]; ok {
				state = s
			} else {
				return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
			}

			userState := &State{
				Id:            stateCfg.Id,
				Name:          state.Name,
				TransitionMap: make(TransitionMap),
			}

			for _, idTransition := range stateCfg.Transitions {
				if transition, ok := state.TransitionMap[idTransition]; ok {
					userState.TransitionMap[idTransition] = transition
				} else {
					return errors.New(fmt.Sprintf("transition from %d to %d not found", stateCfg.Id, idTransition))
				}
			}

			stateMap[stateCfg.Id] = userState
		}
		newStateMachine := make(StateMachineMap)
		newStateMachine[stateMachine] = stateMap
		sm.userStateMachineMap[user] = newStateMachine
	}

	return nil
}

func (sm *StateMachine) AddCheckHandler(name string, handler CheckHandler) *StateMachine {
	sm.handlerMap.Check[name] = handler
	return sm
}

func (sm *StateMachine) AddExecuteHandler(name string, handler ExecuteHandler) *StateMachine {
	sm.handlerMap.Execute[name] = handler
	return sm
}

func (sm *StateMachine) AddEventOnSuccessHandler(name string, handler EventHandler) *StateMachine {
	sm.handlerMap.Events.Success[name] = handler
	return sm
}

func (sm *StateMachine) AddEventOnErrorHandler(name string, handler EventHandler) *StateMachine {
	sm.handlerMap.Events.Error[name] = handler
	return sm
}

func (sm *StateMachine) CheckTransition(stateMachine string, user string, from int, to int, args ...interface{}) (bool, error) {
	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if states, ok := stateM[stateMachine]; ok {
			if state, ok := states[from]; ok {
				if transition, ok := state.TransitionMap[to]; ok {
					var err error
					for _, handler := range transition.Handler.Check {
						if ok, err = handler(args); !ok || err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args); eventErr != nil {
										if eventErr != nil {
											return ok, eventErr
										}
									}
								}
								return ok, err
							}
							return false, nil
						}
					}
				}
			}
		}
	}
	return false, nil
}

func (sm *StateMachine) ExecuteTransition(stateMachine string, user string, from int, to int, args ...interface{}) (bool, error) {
	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if states, ok := stateM[stateMachine]; ok {
			if state, ok := states[from]; ok {
				if transition, ok := state.TransitionMap[to]; ok {
					var err error
					for _, handler := range transition.Handler.Check {
						if ok, err = handler(args); !ok || err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args); eventErr != nil {
										if eventErr != nil {
											return ok, eventErr
										}
									}
								}
								return ok, err
							}
							return false, nil
						}
					}

					for _, handler := range transition.Handler.Execute {
						if ok, err = handler(args); err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args); eventErr != nil {
										if eventErr != nil {
											return ok, eventErr
										}
									}
								}
								return ok, err
							}
							return false, nil
						}
					}

					for _, handler := range transition.Handler.Events.Success {
						if eventErr := handler(args); eventErr != nil {
							if eventErr != nil {
								return ok, eventErr
							}
						}
					}
				}
			}
		}
	}
	return true, nil
}

func (sm *StateMachine) GetTransitions(stateMachine string, user string, from int) (transitions []*Transition, err error) {
	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if states, ok := stateM[stateMachine]; ok {
			if state, ok := states[from]; ok {
				for _, transition := range state.TransitionMap {
					transitions = append(transitions, transition)
				}
			}
		}
	}
	return transitions, nil
}
