package state_machine

import (
	"errors"
	"fmt"
)

func newHandlersMaps() *HandlersMap {
	return &HandlersMap{
		Check:   make(CheckHandlerMap),
		Execute: make(ExecuteHandlerMap),
		Events: &EventMap{
			Success: make(EventSuccessHandlerMap),
			Error:   make(EventErrorHandlerMap),
		},
	}
}

func (h *Handlers) initStateMachineHandlers(stateMachine StateMachineType) {
	if _, ok := h.stateMachineHandlersMap[stateMachine]; !ok {
		h.stateMachineHandlersMap[stateMachine] = newHandlersMaps()
	}
}

func (h *Handlers) getCheckHandler(stateMachine StateMachineType, name string) (CheckHandler, error) {
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

func (h *Handlers) getExecuteHandler(stateMachine StateMachineType, name string) (ExecuteHandler, error) {
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

func (h *Handlers) getEventSuccessHandler(stateMachine StateMachineType, name string) (EventSuccessHandler, error) {
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

func (h *Handlers) getEventErrorHandler(stateMachine StateMachineType, name string) (EventErrorHandler, error) {
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

func (h CheckHandlerList) Run(ctx *Context, args ...interface{}) (bool, error) {
	for _, handler := range h {
		if ok, err := handler(ctx, args...); !ok || err != nil {
			return false, err
		}
	}
	return true, nil
}

func (h ExecuteHandlerList) Run(ctx *Context, args ...interface{}) error {
	for _, handler := range h {
		if err := handler(ctx, args...); err != nil {
			return err
		}
	}
	return nil
}

func (h EventSuccessHandlerList) Run(ctx *Context, args ...interface{}) {
	for _, handler := range h {
		handler(ctx, args...)
	}
}

func (h EventErrorHandlerList) Run(ctx *Context, err error, args ...interface{}) error {
	for _, handler := range h {
		handler(ctx, err, args...)
	}
	return nil
}

func (h *Handler) Run(ctx *Context, transitionHandler TransitionHandler, args ...interface{}) (bool, error) {
	// check
	allowed, err := h.Check.Run(ctx, args...)
	if err != nil {
		h.Events.Error.Run(ctx, err, args...)
		return false, err
	}

	if !allowed {
		return false, nil
	}

	err = h.Execute.Run(ctx, args...)
	if err != nil {
		h.Events.Error.Run(ctx, err, args...)
		return false, err
	}
	// transition handler
	err = transitionHandler(ctx, args...)
	if err != nil {
		h.Events.Error.Run(ctx, err, args...)
		return false, err
	}

	// on success
	h.Events.Success.Run(ctx, args...)

	return true, nil
}
