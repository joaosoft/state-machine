package state_machine

import (
	"errors"
	"fmt"
)

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

func (h *Handlers) getEventSuccessHandler(stateMachine StateMachineType, name string) (EventHandler, error) {
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

func (h *Handlers) getEventErrorHandler(stateMachine StateMachineType, name string) (EventHandler, error) {
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