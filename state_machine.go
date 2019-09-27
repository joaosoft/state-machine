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
	var stateMachineMap StateMachineMap
	var stateMap StateMap

	if stateMachineMap, ok = sm.userStateMachineMap[user]; !ok {
		return false, errors.New(fmt.Sprintf("user [%s] not found", user))
	}

	if stateMap, ok = stateMachineMap[stateMachine]; !ok {
		return false, errors.New(fmt.Sprintf("state machine [%s] not found", stateMachine))
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

	stateMachineMap := make(StateMachineMap)
	states := make(StateMap)
	stateMachineMap[stateMachine] = states

	// load state machine
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
			transition.Handler.Check, err = transitionCfg.getCheckHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// execute
			transition.Handler.Execute, err = transitionCfg.getExecuteHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// events
			// -- success
			transition.Handler.Events.Success, err = transitionCfg.getEventSuccessHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// -- error
			transition.Handler.Events.Error, err = transitionCfg.getEventErrorHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			state.TransitionMap[transitionCfg.Id] = transition
		}

		states[stateCfg.Id] = state
	}

	// load users state machine
	for user, statesCfg := range config.Users {
		stateMap := make(StateMap)

		transitionStates := make(map[int]bool)
		for _, stateCfg := range statesCfg {
			var state *State
			var ok bool

			// check if state is valid !
			if state, ok = states[stateCfg.Id]; !ok {
				return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
			}

			userState := &State{
				Id:            state.Id,
				Name:          state.Name,
				TransitionMap: make(TransitionMap),
			}

			transitionStates[stateCfg.Id] = true

			for _, transitionCfg := range stateCfg.Transitions {
				transitionStates[transitionCfg.Id] = false

				// check if transition is valid !
				var stateTransition *State
				if stateTransition, ok = states[transitionCfg.Id]; !ok {
					return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
				}

				var transition *Transition
				if transition, ok = states[stateCfg.Id].TransitionMap[transitionCfg.Id]; !ok {
					return errors.New(fmt.Sprintf("transition from %d to %d not found", stateCfg.Id, transitionCfg.Id))
				}

				userTransition := &Transition{
					Id:   transition.Id,
					Name: stateTransition.Name,
					Handler: Handler{
						Check:   append([]CheckHandler{}, transition.Handler.Check...),
						Execute: append([]ExecuteHandler{}, transition.Handler.Execute...),
						Events: Events{
							Success: append([]EventHandler{}, transition.Handler.Events.Success...),
							Error:   append([]EventHandler{}, transition.Handler.Events.Error...),
						},
					},
				}

				// add specific handlers for the user

				// check
				checkHandlers, err := transitionCfg.getCheckHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.Handler.Check = append(userTransition.Handler.Check, checkHandlers...)

				// execute
				executeHandlers, err := transitionCfg.getExecuteHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.Handler.Execute = append(userTransition.Handler.Execute, executeHandlers...)

				// events
				// -- success
				eventSuccessHandlers, err := transitionCfg.getEventSuccessHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.Handler.Events.Success = append(userTransition.Handler.Events.Success, eventSuccessHandlers...)

				// -- error
				eventErrorHandlers, err := transitionCfg.getEventErrorHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.Handler.Events.Error = append(userTransition.Handler.Events.Error, eventErrorHandlers...)

				userState.TransitionMap[userTransition.Id] = userTransition
			}

			stateMap[stateCfg.Id] = userState

		}
		// add missing user transition states
		for idState, added := range transitionStates {
			if !added {
				var stateTransition *State
				var ok bool
				if stateTransition, ok = states[idState]; !ok {
					return errors.New(fmt.Sprintf("state not found %d", idState))
				}
				stateMap[idState] = &State{
					Id:            idState,
					Name:          stateTransition.Name,
					TransitionMap: make(TransitionMap),
				}
			}
		}

		if _, ok := sm.userStateMachineMap[UserType(user)]; !ok {
			sm.userStateMachineMap[UserType(user)] = make(StateMachineMap)
		}
		sm.userStateMachineMap[UserType(user)][stateMachine] = stateMap
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
