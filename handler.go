package state_machine

import (
	"errors"
	"fmt"
)

func newHandlersMaps() *HandlersMap {
	return &HandlersMap{
		Manual:  make(ManualHandlerMap),
		Load:    make(LoadHandlerMap),
		Check:   make(CheckHandlerMap),
		Execute: make(ExecuteHandlerMap),
		Events: &EventMap{
			Success: make(EventSuccessHandlerMap),
			Error:   make(EventErrorHandlerMap),
		},
	}
}

func (h *handlers) initStateMachineHandlers(stateMachine StateMachineType) {
	if _, ok := h.stateMachineHandlersMap[stateMachine]; !ok {
		h.stateMachineHandlersMap[stateMachine] = newHandlersMaps()
	}
}

func (h *handlers) RunManual(tag ManualHandlerTag, ctx *Context) error {
	if err := h.handlersMap.Manual.Run(tag, ctx); err != nil {
		return err
	}

	if err := h.stateMachineHandlersMap[ctx.StateMachine].Manual.Run(tag, ctx); err != nil {
		return err
	}

	return nil
}
func (h *handlers) getLoadHandler(stateMachine StateMachineType, name string) (LoadHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.Load[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.Load[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("load handler [%s] not found", name))
	}
}

func (h *handlers) getCheckHandler(stateMachine StateMachineType, name string) (CheckHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.Check[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.Check[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("check handler [%s] not found", name))
	}
}

func (h *handlers) getExecuteHandler(stateMachine StateMachineType, name string) (ExecuteHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.Execute[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.Execute[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("execute handler [%s] not found", name))
	}
}

func (h *handlers) getEventSuccessHandler(stateMachine StateMachineType, name string) (EventSuccessHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.Events.Success[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.Events.Success[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("event success handler [%s] not found", name))
	}
}

func (h *handlers) getEventErrorHandler(stateMachine StateMachineType, name string) (EventErrorHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.Events.Error[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.Events.Error[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("event error handler [%s] not found", name))
	}
}

func (h *Handler) Run(ctx *Context, transitionHandler TransitionHandler, handlers *handlers) (bool, error) {

	// manual - before
	if err := handlers.RunManual(ManualBefore, ctx); err != nil {
		// on error
		h.Events.Error.Run(ctx, err)
		return false, err
	}

	// check
	allowed, err := h.Check.Run(ctx)
	if err != nil {
		// on error
		h.Events.Error.Run(ctx, err)
		return false, err
	}

	if !allowed {
		return false, nil
	}

	// execute
	err = h.Execute.Run(ctx)
	if err != nil {
		// on error
		h.Events.Error.Run(ctx, err)
		return false, err
	}

	// transition handler
	if transitionHandler != nil {
		err = transitionHandler(ctx)
		if err != nil {
			// on error
			h.Events.Error.Run(ctx, err)
			return false, err
		}
	}

	// on success
	h.Events.Success.Run(ctx)

	// manual - after
	if err := handlers.RunManual(ManualAfter, ctx); err != nil {
		// on error
		h.Events.Error.Run(ctx, err)
		return false, err
	}

	return true, nil
}

func (h LoadHandlerList) Run(ctx *Context) error {
	for _, handler := range h {
		if err := handler(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h CheckHandlerList) Run(ctx *Context) (bool, error) {
	for _, handler := range h {
		if ok, err := handler(ctx); !ok || err != nil {
			return false, err
		}
	}
	return true, nil
}

func (h ExecuteHandlerList) Run(ctx *Context) error {
	for _, handler := range h {
		if err := handler(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h EventSuccessHandlerList) Run(ctx *Context) {
	for _, handler := range h {
		handler(ctx)
	}
}

func (h EventErrorHandlerList) Run(ctx *Context, err error) error {
	for _, handler := range h {
		handler(ctx, err)
	}
	return nil
}

func (m ManualHandlerMap) Run(tag ManualHandlerTag, ctx *Context) error {
	if list, ok := m[tag]; ok {
		for _, handler := range list {
			if err := handler(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
