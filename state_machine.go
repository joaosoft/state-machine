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
		roleStateMachineMap: make(roleStateMachineMap),
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

	if stateMachineMap, ok = sm.roleStateMachineMap[ctx.Role]; !ok {
		return false, errors.New(fmt.Sprintf("role [%s] not found", ctx.Role))
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

	// load roles state machine
	for role, statesCfg := range config.Roles {
		stateMap := make(stateMap)

		transitionStates := make(map[int]bool)
		for _, stateCfg := range statesCfg {
			var fromState *state
			var ok bool

			// check if state is valid !
			if fromState, ok = states[stateCfg.Id]; !ok {
				return errors.New(fmt.Sprintf("state not found %d", stateCfg.Id))
			}

			roleState := &state{
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

				roleTransition := &Transition{
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

				// add specific handlers for the role

				// load
				loadHandlers, err := transitionCfg.getLoadHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				roleTransition.handler.load = append(roleTransition.handler.load, loadHandlers...)

				// check
				checkHandlers, err := transitionCfg.getCheckHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				roleTransition.handler.check = append(roleTransition.handler.check, checkHandlers...)

				// execute
				executeHandlers, err := transitionCfg.getExecuteHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				roleTransition.handler.execute = append(roleTransition.handler.execute, executeHandlers...)

				// events
				// -- success
				eventSuccessHandlers, err := transitionCfg.getEventSuccessHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				roleTransition.handler.events.success = append(roleTransition.handler.events.success, eventSuccessHandlers...)

				// -- error
				eventErrorHandlers, err := transitionCfg.getEventErrorHandlers(stateMachine, sm.handlers)
				if err != nil {
					return err
				}
				roleTransition.handler.events.error = append(roleTransition.handler.events.error, eventErrorHandlers...)

				roleState.transitionMap[roleTransition.Id] = roleTransition
			}

			stateMap[stateCfg.Id] = roleState

		}
		// add missing role transition states
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

		if _, ok := sm.roleStateMachineMap[RoleType(role)]; !ok {
			sm.roleStateMachineMap[RoleType(role)] = make(stateMachineMap)
		}
		sm.roleStateMachineMap[RoleType(role)][stateMachine] = &stateMachineData{
			stateMap:          stateMap,
			transitionHandler: transitionHandler,
		}
	}

	return nil
}

func (sm *stateMachine) addManualHandler(tags []manualHandlerKey, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	if len(stateMachine) > 0 {
		for _, stateMachine := range stateMachine {
			sm.handlers.initStateMachineHandlers(stateMachine)
			for _, tag := range tags {
				if _, ok := sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag]; !ok {
					sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag] = make(manualHandlerList, 0)
				}
				sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag] = append(sm.handlers.stateMachineHandlersMap[stateMachine].manual[tag], handler)
			}
		}
	} else {
		for _, tag := range tags {
			if _, ok := sm.handlers.handlersMap.manual[tag]; !ok {
				sm.handlers.handlersMap.manual[tag] = make(manualHandlerList, 0)
			}
			sm.handlers.handlersMap.manual[tag] = append(sm.handlers.handlersMap.manual[tag], handler)
		}
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

	// manual - before
	if err := sm.handlers.RunManual(BeforeCheck, ctx); err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return false, err
	}

	if ok, err := sm.validate(ctx); err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return ok, err
	}

	stateM, ok := sm.roleStateMachineMap[ctx.Role]
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
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return false, err
	}

	// check
	allowed, err := transition.handler.check.Run(ctx)
	if err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return false, err
	}

	return allowed, nil
}

func (sm *stateMachine) executeTransition(ctx *Context) (bool, error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	// manual - before
	if err := sm.handlers.RunManual(BeforeExecute, ctx); err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return false, err
	}

	if ok, err := sm.validate(ctx, ctx.From, ctx.To); err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return ok, err
	}

	// role
	roleStateMachine, ok := sm.roleStateMachineMap[ctx.Role]
	if !ok {
		return false, nil
	}

	// state machine
	states, ok := roleStateMachine[ctx.StateMachine]
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

	if allowed, err := transition.handler.Run(ctx, states.transitionHandler, sm.logger); !allowed || err != nil {
		if err != nil {
			if err := sm.handlers.RunManual(OnError, ctx); err != nil {
				sm.logger.Error(err)
			}
		}
		return allowed, err
	}

	// manual - after
	if err := sm.handlers.RunManual(AfterExecute, ctx); err != nil {
		if err := sm.handlers.RunManual(OnError, ctx); err != nil {
			sm.logger.Error(err)
		}
		return false, err
	}

	if err := sm.handlers.RunManual(OnSuccess, ctx); err != nil {
		sm.logger.Error(err)
	}

	return true, nil
}

func (sm *stateMachine) getTransitions(ctx *Context) (transitions []*Transition, err error) {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	if _, err := sm.validate(ctx, ctx.From); err != nil {
		return nil, err
	}

	var ok bool

	var roleStateMachine stateMachineMap
	if roleStateMachine, ok = sm.roleStateMachineMap[ctx.Role]; !ok {
		return nil, nil
	}

	var stateMachineData *stateMachineData
	if stateMachineData, ok = roleStateMachine[ctx.StateMachine]; !ok {
		return nil, nil
	}

	var state *state
	if state, ok = stateMachineData.stateMap[ctx.From]; !ok {
		return nil, nil
	}

	newCtx := &Context{
		Role:         ctx.Role,
		StateMachine: ctx.StateMachine,
		From:         ctx.From,
		Resource:     ctx.Resource,
		Args:         ctx.Args,
	}

	var allowed bool

	for _, transition := range state.transitionMap {
		newCtx.To = transition.Id
		if allowed, err = sm.checkTransition(newCtx); err != nil {
			return nil, err
		}

		if allowed {
			transitions = append(transitions, transition)
		}
	}

	return transitions, nil
}
