package state_machine

import (
	"errors"
	"fmt"
	"github.com/joaosoft/logger"
)

func newHandlersMaps() *handlersMap {
	return &handlersMap{
		manual:  make(manualHandlerMap),
		Load:    make(loadHandlerMap),
		check:   make(checkHandlerMap),
		execute: make(executeHandlerMap),
		events: &eventMap{
			success: make(eventSuccessHandlerMap),
			error:   make(eventErrorHandlerMap),
		},
	}
}

func (h *handlers) initStateMachineHandlers(stateMachine StateMachineType) {
	if _, ok := h.stateMachineHandlersMap[stateMachine]; !ok {
		h.stateMachineHandlersMap[stateMachine] = newHandlersMaps()
	}
}

func (h *handlers) RunManual(tag manualHandlerKey, ctx *Context) error {
	if err := h.handlersMap.manual.Run(tag, ctx); err != nil {
		return err
	}

	if err := h.stateMachineHandlersMap[ctx.StateMachine].manual.Run(tag, ctx); err != nil {
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
		if handler, ok := stateM.check[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.check[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("check handler [%s] not found", name))
	}
}

func (h *handlers) getExecuteHandler(stateMachine StateMachineType, name string) (ExecuteHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.execute[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.execute[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("execute handler [%s] not found", name))
	}
}

func (h *handlers) getEventSuccessHandler(stateMachine StateMachineType, name string) (EventSuccessHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.events.success[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.events.success[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("event success handler [%s] not found", name))
	}
}

func (h *handlers) getEventErrorHandler(stateMachine StateMachineType, name string) (EventErrorHandler, error) {
	if stateM, ok := h.stateMachineHandlersMap[stateMachine]; ok {
		if handler, ok := stateM.events.error[name]; ok {
			return handler, nil
		}
	}

	if handler, ok := h.handlersMap.events.error[name]; ok {
		return handler, nil
	} else {
		return nil, errors.New(fmt.Sprintf("event error handler [%s] not found", name))
	}
}

func (h *handler) Run(ctx *Context, transitionHandler TransitionHandler, logger logger.ILogger) (bool, error) {
	var err error

	// load
	err = h.load.Run(ctx)
	if err != nil {
		// on error
		h.events.error.Run(ctx, err, logger)
		return false, err
	}

	// check
	var allowed bool
	allowed, err = h.check.Run(ctx)
	if err != nil {
		// on error
		h.events.error.Run(ctx, err, logger)
		return false, err
	}

	if !allowed {
		return false, nil
	}

	// execute
	err = h.execute.Run(ctx)
	if err != nil {
		// on error
		h.events.error.Run(ctx, err, logger)
		return false, err
	}

	// transition handler
	if transitionHandler != nil {
		err = transitionHandler(ctx)
		if err != nil {
			// on error
			h.events.error.Run(ctx, err, logger)
			return false, err
		}
	}

	// on success
	h.events.success.Run(ctx, logger)

	return true, nil
}

func (h loadHandlerList) Run(ctx *Context) error {
	for _, handler := range h {
		if err := handler(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h checkHandlerList) Run(ctx *Context) (bool, error) {
	for _, handler := range h {
		if ok, err := handler(ctx); !ok || err != nil {
			return false, err
		}
	}
	return true, nil
}

func (h executeHandlerList) Run(ctx *Context) error {
	for _, handler := range h {
		if err := handler(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h eventSuccessHandlerList) Run(ctx *Context, logger logger.ILogger) {
	var err error
	for _, handler := range h {
		if err = handler(ctx); err != nil {
			logger.Error(err)
		}
	}
}

func (h eventErrorHandlerList) Run(ctx *Context, err error, logger logger.ILogger) {
	var eventErr error
	for _, handler := range h {
		if eventErr = handler(ctx, err); err != nil {
			logger.Error(eventErr)
		}
	}
}

func (m manualHandlerMap) Run(tag manualHandlerKey, ctx *Context) error {
	if list, ok := m[tag]; ok {
		for _, handler := range list {
			if err := handler(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
