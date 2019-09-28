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

func new(options ...stateMachineOption) (*stateMachine, error) {
	config, _, err := newConfig()

	newStateMachine := &stateMachine{
		stateMachineMap:     make(stateMachineMap),
		userStateMachineMap: make(userStateMachineMap),
		handlers: &handlers{
			handlersMap: &handlersMap{
				manual:  make(manualHandlerMap),
				check:   make(checkHandlerMap),
				execute: make(executeHandlerMap),
				events: &eventMap{
					success: make(eventSuccessHandlerMap),
					error:   make(eventErrorHandlerMap),
				},
			},
			stateMachineHandlersMap: make(stateMachineHandlersMap),
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
	var stateMachineMap stateMachineMap
	var stateMachineData *stateMachineData

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

	config := stateMachineCfg{}
	_, err := manager.NewSimpleConfig(file, &config)
	if err != nil {
		return err
	}

	states := make(stateMap)
	sm.stateMachineMap[stateMachine] = &stateMachineData{
		stateMap:          states,
		transitionHandler: transitionHandler,
	}

	// load state machine
	for _, stateCfg := range config.StateMachine {

		state := &state{
			id:            stateCfg.Id,
			name:          stateCfg.Name,
			transitionMap: make(transitionMap),
		}

		for _, transitionCfg := range stateCfg.Transitions {
			transition := &Transition{
				Id: transitionCfg.Id,
			}

			// load
			transition.handler.load, err = transitionCfg.getLoadHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// check
			transition.handler.check, err = transitionCfg.getCheckHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// execute
			transition.handler.execute, err = transitionCfg.getExecuteHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// events
			// -- success
			transition.handler.events.success, err = transitionCfg.getEventSuccessHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			// -- error
			transition.handler.events.error, err = transitionCfg.getEventErrorHandlers(stateMachine, sm.handlers)
			if err != nil {
				return err
			}

			state.transitionMap[transitionCfg.Id] = transition
		}

		states[stateCfg.Id] = state
	}

	// load users state machine
	for user, statesCfg := range config.Users {
		stateMap := make(stateMap)

		transitionStates := make(map[int]bool)
		for _, stateCfg := range statesCfg {
			var fromState *state
			var ok bool

			// check if state is valid !
			if fromState, ok = states[stateCfg.Id]; !ok {
				return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
			}

			userState := &state{
				id:            fromState.id,
				name:          fromState.name,
				transitionMap: make(transitionMap),
			}

			transitionStates[stateCfg.Id] = true

			for _, transitionCfg := range stateCfg.Transitions {
				transitionStates[transitionCfg.Id] = false

				// check if transition is valid !
				var toState *state
				if toState, ok = states[transitionCfg.Id]; !ok {
					return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
				}

				var transition *Transition
				if transition, ok = states[stateCfg.Id].transitionMap[transitionCfg.Id]; !ok {
					return errors.New(fmt.Sprintf("transition From %d To %d not found", stateCfg.Id, transitionCfg.Id))
				}

				userTransition := &Transition{
					Id:   transition.Id,
					Name: toState.name,
					handler: handler{
						check:   append([]CheckHandler{}, transition.handler.check...),
						execute: append([]ExecuteHandler{}, transition.handler.execute...),
						events: events{
							success: append([]EventSuccessHandler{}, transition.handler.events.success...),
							error:   append([]EventErrorHandler{}, transition.handler.events.error...),
						},
					},
				}

				// add specific handlers for the user

				// load
				loadHandlers, err := transitionCfg.getLoadHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.handler.load = append(userTransition.handler.load, loadHandlers...)

				// check
				checkHandlers, err := transitionCfg.getCheckHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.handler.check = append(userTransition.handler.check, checkHandlers...)

				// execute
				executeHandlers, err := transitionCfg.getExecuteHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.handler.execute = append(userTransition.handler.execute, executeHandlers...)

				// events
				// -- success
				eventSuccessHandlers, err := transitionCfg.getEventSuccessHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.handler.events.success = append(userTransition.handler.events.success, eventSuccessHandlers...)

				// -- error
				eventErrorHandlers, err := transitionCfg.getEventErrorHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				userTransition.handler.events.error = append(userTransition.handler.events.error, eventErrorHandlers...)

				userState.transitionMap[userTransition.Id] = userTransition
			}

			stateMap[stateCfg.Id] = userState

		}
		// add missing user transition states
		for idState, added := range transitionStates {
			if !added {
				var toState *state
				var ok bool
				if toState, ok = states[idState]; !ok {
					return errors.New(fmt.Sprintf("state not found %d", idState))
				}
				stateMap[idState] = &state{
					id:            idState,
					name:          toState.name,
					transitionMap: make(transitionMap),
				}
			}
		}

		if _, ok := sm.userStateMachineMap[UserType(user)]; !ok {
			sm.userStateMachineMap[UserType(user)] = make(stateMachineMap)
		}
		sm.userStateMachineMap[UserType(user)][stateMachine] = &stateMachineData{
			stateMap:          stateMap,
			transitionHandler: transitionHandler,
		}
	}

	return nil
}

func (sm *stateMachine) addManualHandler(tag manualHandlerTag, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, stateMachine := range stateMachine {
			sm.handlers.initStateMachineHandlers(stateMachine)
			if _, ok := sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag]; !ok {
				sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag] = make(manualHandlerList, 0)
			}
			sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag] = append(sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag], handler)
		}
	} else {
		if _, ok := sm.handlers.handlersMap.manual[tag]; !ok {
			sm.handlers.handlersMap.manual[tag] = make(manualHandlerList, 0)
		}
		sm.handlers.handlersMap.manual[tag] = append(sm.handlers.handlersMap.manual[tag], handler)
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
				sm.handlers.stateMachineHandlersMap[stateMachine].check[name] = h
			case ExecuteHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].execute[name] = h
			case EventSuccessHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].events.success[name] = h
			case EventErrorHandler:
				sm.handlers.stateMachineHandlersMap[stateMachine].events.error[name] = h
			}
		}
	} else {
		switch h := handler.(type) {
		case LoadHandler:
			sm.handlers.handlersMap.Load[name] = h
		case CheckHandler:
			sm.handlers.handlersMap.check[name] = h
		case ExecuteHandler:
			sm.handlers.handlersMap.execute[name] = h
		case EventSuccessHandler:
			sm.handlers.handlersMap.events.success[name] = h
		case EventErrorHandler:
			sm.handlers.handlersMap.events.error[name] = h
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

	transition, ok := state.transitionMap[ctx.To]
	if !ok {
		return false, nil
	}

	// load
	err := transition.handler.load.Run(ctx)
	if err != nil {
		return false, err
	}

	// check
	allowed, err := transition.handler.check.Run(ctx)
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
	transition, ok := state.transitionMap[ctx.To]
	if !ok {
		return false, nil
	}

	return transition.handler.Run(ctx, states.transitionHandler, sm.handlers)
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
				for _, transition := range state.transitionMap {
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
