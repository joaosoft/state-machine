package state_machine

type stateMachineCfg struct {
	StateMachine []struct {
		Id          int             `json:"id" yaml:"id"`
		Name        string          `json:"name" yaml:"name"`
		Transitions []transitionCfg `json:"transitions" yaml:"transitions"`
	} `json:"state_machine" yaml:"state_machine"`
	Roles map[string][]struct {
		Id          int             `json:"id" yaml:"id"`
		Transitions []transitionCfg `json:"transitions" yaml:"transitions"`
	} `json:"roles" yaml:"roles"`
}

type transitionCfg struct {
	Id      int      `json:"id" yaml:"id"`
	Load    []string `json:"load" yaml:"load"`
	Check   []string `json:"check" yaml:"check"`
	Execute []string `json:"execute" yaml:"execute"`
	Events  struct {
		Success []string `json:"success" yaml:"success"`
		Error   []string `json:"error" yaml:"error"`
	} `json:"events" yaml:"events"`
}

func (t transitionCfg) getLoadHandlers(stateMachine StateMachineType, handlers *handlers) (loadHandlers loadHandlerList, err error) {
	for _, name := range t.Load {
		handler, err := handlers.getLoadHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		loadHandlers = append(loadHandlers, handler)
	}

	return loadHandlers, nil
}

func (t transitionCfg) getCheckHandlers(stateMachine StateMachineType, handlers *handlers) (checkHandlers checkHandlerList, err error) {
	for _, name := range t.Check {
		handler, err := handlers.getCheckHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		checkHandlers = append(checkHandlers, handler)
	}

	return checkHandlers, nil
}

func (t transitionCfg) getExecuteHandlers(stateMachine StateMachineType, handlers *handlers) (executeHandlers executeHandlerList, err error) {
	for _, name := range t.Execute {
		handler, err := handlers.getExecuteHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		executeHandlers = append(executeHandlers, handler)
	}

	return executeHandlers, nil
}

func (t transitionCfg) getEventSuccessHandlers(stateMachine StateMachineType, handlers *handlers) (eventSuccessHandlers eventSuccessHandlerList, err error) {
	for _, name := range t.Events.Success {
		handler, err := handlers.getEventSuccessHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		eventSuccessHandlers = append(eventSuccessHandlers, handler)
	}

	return eventSuccessHandlers, nil
}

func (t transitionCfg) getEventErrorHandlers(stateMachine StateMachineType, handlers *handlers) (eventErrorHandlers eventErrorHandlerList, err error) {
	for _, name := range t.Events.Error {
		handler, err := handlers.getEventErrorHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		eventErrorHandlers = append(eventErrorHandlers, handler)
	}

	return eventErrorHandlers, nil
}
