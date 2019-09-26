package state_machine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

var (
	stateMachineInstance, _ = New()
)

func New(options ...StateMachineOption) (*StateMachine, error) {
	config, _, err := NewConfig()

	newStateMachine := &StateMachine{
		stateMachineMap:     make(StateMachineMap),
		userStateMachineMap: make(UserStateMachineMap),
		handlers: &Handlers{
			handlersMap: &HandlersMap{
				Check:   make(CheckHandlerMap),
				Execute: make(ExecuteHandlerMap),
				Events: &EventMap{
					Success: make(EventHandlerMap),
					Error:   make(EventHandlerMap),
				},
			},
			stateMachineHandlersMap: make(StateMachineHandlersMap),
		},
		logger: logger.NewLogDefault("state_machine", logger.WarnLevel),
		config: config.StateMachine,
		mux:    &sync.RWMutex{},
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

func newHandlersMaps() *HandlersMap {
	return &HandlersMap{
		Check:   make(CheckHandlerMap),
		Execute: make(ExecuteHandlerMap),
		Events: &EventMap{
			Success: make(EventHandlerMap),
			Error:   make(EventHandlerMap),
		},
	}
}

func (h *Handlers) initStateMachineHandlers(stateMachine StateMachineType) {
	if _, ok := h.stateMachineHandlersMap[stateMachine]; !ok {
		h.stateMachineHandlersMap[stateMachine] = newHandlersMaps()
	}
}

func (sm *StateMachine) validate(stateMachine StateMachineType, user UserType, states ...int) (bool, error) {
	var ok bool
	var stateMap StateMap

	if _, ok = sm.userStateMachineMap[user]; !ok {
		return false, errors.New(fmt.Sprintf("user [%s] not found", user))
	}

	if stateMap, ok = sm.stateMachineMap[stateMachine]; !ok {
		return false, errors.New(fmt.Sprintf("state machine [%s] not found", user))
	} else {
		for _, state := range states {
			if _, ok = stateMap[state]; !ok {
				return false, errors.New(fmt.Sprintf("state [%d] not found", state))
			}
		}
	}

	return true, nil
}

func (sm *StateMachine) Add(stateMachine StateMachineType, file string) error {
	sm.mux.Lock()
	defer sm.mux.Unlock()

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
				enc := false
				if stateM, ok := sm.handlers.stateMachineHandlersMap[stateMachine]; ok {
					if handler, ok := stateM.Check[name]; ok {
						enc = true
						transition.Handler.Check = append(transition.Handler.Check, handler)
					}
				}

				if !enc {
					if handler, ok := sm.handlers.handlersMap.Check[name]; ok {
						transition.Handler.Check = append(transition.Handler.Check, handler)
					} else {
						return errors.New(fmt.Sprintf("check handler [%s] not found", name))
					}
				}
			}

			// execute
			for _, name := range transitionCfg.Execute {
				enc := false
				if stateM, ok := sm.handlers.stateMachineHandlersMap[stateMachine]; ok {
					if handler, ok := stateM.Execute[name]; ok {
						enc = true
						transition.Handler.Execute = append(transition.Handler.Execute, handler)
					}
				}

				if !enc {
					if handler, ok := sm.handlers.handlersMap.Execute[name]; ok {
						transition.Handler.Execute = append(transition.Handler.Execute, handler)
					} else {
						return errors.New(fmt.Sprintf("execute handler [%s] not found", name))
					}
				}
			}

			// events
			// -- success
			for _, name := range transitionCfg.Events.Success {
				enc := false
				if stateM, ok := sm.handlers.stateMachineHandlersMap[stateMachine]; ok {
					if handler, ok := stateM.Events.Success[name]; ok {
						enc = true
						transition.Handler.Events.Success = append(transition.Handler.Events.Success, handler)
					}
				}

				if !enc {
					if handler, ok := sm.handlers.handlersMap.Events.Success[name]; ok {
						transition.Handler.Events.Success = append(transition.Handler.Events.Success, handler)
					} else {
						return errors.New(fmt.Sprintf("event success handler [%s] not found", name))
					}
				}
			}

			// -- error
			for _, name := range transitionCfg.Events.Error {
				enc := false
				if stateM, ok := sm.handlers.stateMachineHandlersMap[stateMachine]; ok {
					if handler, ok := stateM.Events.Error[name]; ok {
						enc = true
						transition.Handler.Events.Error = append(transition.Handler.Events.Error, handler)
					}
				}

				if !enc {
					if handler, ok := sm.handlers.handlersMap.Events.Error[name]; ok {
						transition.Handler.Events.Error = append(transition.Handler.Events.Error, handler)
					} else {
						return errors.New(fmt.Sprintf("event error handler [%s] not found", name))
					}
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
		sm.userStateMachineMap[UserType(user)] = newStateMachine
	}

	return nil
}

func (sm *StateMachine) AddCheckHandler(name string, handler CheckHandler, stateMachine ...StateMachineType) *StateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, smName := range stateMachine {
			sm.handlers.initStateMachineHandlers(smName)
			sm.handlers.stateMachineHandlersMap[smName].Check[name] = handler
		}
	} else {
		sm.handlers.handlersMap.Check[name] = handler
	}

	return sm
}

func (sm *StateMachine) AddExecuteHandler(name string, handler ExecuteHandler, stateMachine ...StateMachineType) *StateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, smName := range stateMachine {
			sm.handlers.initStateMachineHandlers(smName)
			sm.handlers.stateMachineHandlersMap[smName].Execute[name] = handler
		}
	} else {
		sm.handlers.handlersMap.Execute[name] = handler
	}

	return sm
}

func (sm *StateMachine) AddEventOnSuccessHandler(name string, handler EventHandler, stateMachine ...StateMachineType) *StateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, smName := range stateMachine {
			sm.handlers.initStateMachineHandlers(smName)
			sm.handlers.stateMachineHandlersMap[smName].Events.Success[name] = handler
		}
	} else {
		sm.handlers.handlersMap.Events.Success[name] = handler
	}

	return sm
}

func (sm *StateMachine) AddEventOnErrorHandler(name string, handler EventHandler, stateMachine ...StateMachineType) *StateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, smName := range stateMachine {
			sm.handlers.initStateMachineHandlers(smName)
			sm.handlers.stateMachineHandlersMap[smName].Events.Error[name] = handler
		}
	} else {
		sm.handlers.handlersMap.Events.Error[name] = handler
	}

	return sm
}

func (sm *StateMachine) CheckTransition(stateMachine StateMachineType, user UserType, from int, to int, args ...interface{}) (bool, error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	if ok, err := sm.validate(stateMachine, user, from, to); err != nil {
		return ok, err
	}

	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if stateMap, ok := stateM[stateMachine]; ok {
			if state, ok := stateMap[from]; ok {
				if transition, ok := state.TransitionMap[to]; ok {
					var err error
					for _, handler := range transition.Handler.Check {
						if ok, err = handler(args...); !ok || err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args...); eventErr != nil {
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
				} else {
					return false, nil
				}
			} else {
				return false, nil
			}
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}

	return false, nil
}

func (sm *StateMachine) ExecuteTransition(stateMachine StateMachineType, user UserType, from int, to int, args ...interface{}) (bool, error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	if ok, err := sm.validate(stateMachine, user, from, to); err != nil {
		return ok, err
	}

	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if stateMap, ok := stateM[stateMachine]; ok {
			if state, ok := stateMap[from]; ok {
				if transition, ok := state.TransitionMap[to]; ok {
					var err error
					for _, handler := range transition.Handler.Check {
						if ok, err = handler(args...); !ok || err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args...); eventErr != nil {
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
						if ok, err = handler(args...); err != nil {
							if err != nil {
								for _, handler := range transition.Handler.Events.Error {
									if eventErr := handler(args...); eventErr != nil {
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
						if eventErr := handler(args...); eventErr != nil {
							if eventErr != nil {
								return ok, eventErr
							}
						}
					}
				} else {
					return false, nil
				}
			} else {
				return false, nil
			}
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
	return true, nil
}

func (sm *StateMachine) GetTransitions(stateMachine StateMachineType, user UserType, from int) (transitions []*Transition, err error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	if _, err := sm.validate(stateMachine, user, from); err != nil {
		return nil, err
	}

	if stateM, ok := sm.userStateMachineMap[user]; ok {
		if stateMap, ok := stateM[stateMachine]; ok {
			if state, ok := stateMap[from]; ok {
				for _, transition := range state.TransitionMap {
					transitions = append(transitions, transition)
				}
			} else {
				return nil, nil
			}
		} else {
			return nil, nil
		}
	} else {
		return nil, nil
	}

	return transitions, nil
}
