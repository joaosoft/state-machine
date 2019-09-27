package state_machine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

var (
	stateMachineInstance, _ = new()
)

func new(options ...StateMachineOption) (*stateMachine, error) {
	config, _, err := NewConfig()

	newStateMachine := &stateMachine{
		stateMachineMap:     make(StateMachineMap),
		userStateMachineMap: make(UserStateMachineMap),
		handlers: &handlers{
			handlersMap: &HandlersMap{
				Manual:  make(ManualHandlerMap),
				Check:   make(CheckHandlerMap),
				Execute: make(ExecuteHandlerMap),
				Events: &EventMap{
					Success: make(EventSuccessHandlerMap),
					Error:   make(EventErrorHandlerMap),
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
		newStateMachine.logger.Debugf("setting log level To %s", level)
		newStateMachine.logger.Reconfigure(logger.WithLevel(level))
	}

	newStateMachine.Reconfigure(options...)

	return newStateMachine, nil
}

func (sm *stateMachine) validate(ctx *Context, states ...int) (bool, error) {
	var ok bool
	var stateMachineMap StateMachineMap
	var stateMachineData *StateMachineData

	if stateMachineMap, ok = sm.userStateMachineMap[ctx.User]; !ok {
		return false, errors.New(fmt.Sprintf("user [%s] not found", ctx.User))
	}

	if stateMachineData, ok = stateMachineMap[ctx.StateMachine]; !ok {
		return false, errors.New(fmt.Sprintf("state machine [%s] not found", ctx.StateMachine))
	} else {
		for _, state := range states {
			if _, ok = stateMachineData.stateMap[state]; !ok {
				return false, errors.New(fmt.Sprintf("state [%d] not found", state))
			}
		}
	}

	return true, nil
}

func (sm *stateMachine) add(stateMachine StateMachineType, file string, transitionHandler TransitionHandler) error {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	config := StateMachineCfg{}
	_, err := manager.NewSimpleConfig(file, &config)
	if err != nil {
		return err
	}

	states := make(StateMap)
	sm.stateMachineMap[stateMachine] = &StateMachineData{
		stateMap:          states,
		transitionHandler: transitionHandler,
	}

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

			// load
			transition.Handler.Load, err = transitionCfg.getLoadHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
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
					return errors.New(fmt.Sprintf("transition From %d To %d not found", stateCfg.Id, transitionCfg.Id))
				}

				userTransition := &Transition{
					Id:   transition.Id,
					Name: stateTransition.Name,
					Handler: Handler{
						Check:   append([]CheckHandler{}, transition.Handler.Check...),
						Execute: append([]ExecuteHandler{}, transition.Handler.Execute...),
						Events: Events{
							Success: append([]EventSuccessHandler{}, transition.Handler.Events.Success...),
							Error:   append([]EventErrorHandler{}, transition.Handler.Events.Error...),
						},
					},
				}

				// add specific handlers for the user

				// load
				loadHandlers, err := transitionCfg.getLoadHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.Handler.Load = append(userTransition.Handler.Load, loadHandlers...)

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
		sm.userStateMachineMap[UserType(user)][stateMachine] = &StateMachineData{
			stateMap:          stateMap,
			transitionHandler: transitionHandler,
		}
	}

	return nil
}

func (sm *stateMachine) addManualHandler(tag ManualHandlerTag, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, stateMachine := range stateMachine {
			sm.handlers.initStateMachineHandlers(stateMachine)
			if _, ok := sm.handlers.stateMachineHandlersMap[stateMachine].Manual[tag]; !ok {
				sm.handlers.stateMachineHandlersMap[stateMachine].Manual[tag] = make(ManualHandlerList, 0)
			}
			sm.handlers.stateMachineHandlersMap[stateMachine].Manual[tag] = append(sm.handlers.stateMachineHandlersMap[stateMachine].Manual[tag], handler)
		}
	} else {
		if _, ok := sm.handlers.handlersMap.Manual[tag]; !ok {
			sm.handlers.handlersMap.Manual[tag] = make(ManualHandlerList, 0)
		}
		sm.handlers.handlersMap.Manual[tag] = append(sm.handlers.handlersMap.Manual[tag], handler)
	}

	return sm
}

func (sm *stateMachine) addHandler(name string, handler interface{}, stateMachine ...StateMachineType) *stateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, stateMachine := range stateMachine {
			sm.handlers.initStateMachineHandlers(stateMachine)

			switch h := handler.(type) {
			case LoadHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].Load[name] = h
			case CheckHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].Check[name] = h
			case ExecuteHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].Execute[name] = h
			case EventSuccessHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].Events.Success[name] = h
			case EventErrorHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].Events.Error[name] = h
			}
		}
	} else {
		switch h := handler.(type) {
		case LoadHandler:
			sm.handlers.handlersMap.Load[name] = h
		case CheckHandler:
			sm.handlers.handlersMap.Check[name] = h
		case ExecuteHandler:
			sm.handlers.handlersMap.Execute[name] = h
		case EventSuccessHandler:
			sm.handlers.handlersMap.Events.Success[name] = h
		case EventErrorHandler:
			sm.handlers.handlersMap.Events.Error[name] = h
		}
	}

	return sm
}

func (sm *stateMachine) checkTransition(ctx *Context) (bool, error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	// manual - init
	if err := sm.handlers.RunManual(ManualInit, ctx); err != nil {
		return false, err
	}

	if ok, err := sm.validate(ctx); err != nil {
		return ok, err
	}

	stateM, ok := sm.userStateMachineMap[ctx.User]
	if !ok {
		return false, nil
	}

	stateMachineData, ok := stateM[ctx.StateMachine]
	if !ok {
		return false, nil
	}

	state, ok := stateMachineData.stateMap[ctx.From]
	if !ok {
		return false, nil
	}

	transition, ok := state.TransitionMap[ctx.To]
	if !ok {
		return false, nil
	}

	// load
	err := transition.Handler.Load.Run(ctx)
	if err != nil {
		return false, err
	}

	// check
	allowed, err := transition.Handler.Check.Run(ctx)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

func (sm *stateMachine) executeTransition(ctx *Context) (bool, error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	// manual - init
	if err := sm.handlers.RunManual(ManualInit, ctx); err != nil {
		return false, err
	}

	if ok, err := sm.validate(ctx, ctx.From, ctx.To); err != nil {
		return ok, err
	}

	// user
	userStateMachine, ok := sm.userStateMachineMap[ctx.User]
	if !ok {
		return false, nil
	}

	// state machine
	states, ok := userStateMachine[ctx.StateMachine]
	if !ok {
		return false, nil
	}

	// from
	state, ok := states.stateMap[ctx.From]
	if !ok {
		return false, nil
	}

	// to
	transition, ok := state.TransitionMap[ctx.To]
	if !ok {
		return false, nil
	}

	return transition.Handler.Run(ctx, states.transitionHandler, sm.handlers)
}

func (sm *stateMachine) getTransitions(ctx *Context) (transitions []*Transition, err error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	if _, err := sm.validate(ctx, ctx.From); err != nil {
		return nil, err
	}

	if userStateMachine, ok := sm.userStateMachineMap[ctx.User]; ok {
		if stateMachineData, ok := userStateMachine[ctx.StateMachine]; ok {
			if state, ok := stateMachineData.stateMap[ctx.From]; ok {
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
