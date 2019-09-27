package state_machine

type StateMachineCfg struct {
	StateMachine []struct {
		Id          int             `json:"id" yaml:"id"`
		Name        string          `json:"name" yaml:"name"`
		Transitions []TransitionCfg `json:"transitions" yaml:"transitions"`
	} `json:"state_machine" yaml:"state_machine"`
	Users map[string][]struct {
		Id          int             `json:"id" yaml:"id"`
		Transitions []TransitionCfg `json:"transitions" yaml:"transitions"`
	} `json:"users" yaml:"users"`
}

type TransitionCfg struct {
	Id      int      `json:"id" yaml:"id"`
	Check   []string `json:"check" yaml:"check"`
	Execute []string `json:"execute" yaml:"execute"`
	Events  struct {
		Success []string `json:"success" yaml:"success"`
		Error   []string `json:"error" yaml:"error"`
	} `json:"events" yaml:"events"`
}

func (t TransitionCfg) getCheckHandlers(stateMachine StateMachineType, handlers *Handlers) (checkHandlers []CheckHandler, err error) {
	for _, name := range t.Check {
		handler, err := handlers.getCheckHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		checkHandlers = append(checkHandlers, handler)
	}

	return checkHandlers, nil
}

func (t TransitionCfg) getExecuteHandlers(stateMachine StateMachineType, handlers *Handlers) (executeHandlers []ExecuteHandler, err error) {
	for _, name := range t.Execute {
		handler, err := handlers.getExecuteHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		executeHandlers = append(executeHandlers, handler)
	}

	return executeHandlers, nil
}

func (t TransitionCfg) getEventSuccessHandlers(stateMachine StateMachineType, handlers *Handlers) (eventSuccessHandlers []EventSuccessHandler, err error) {
	for _, name := range t.Events.Success {
		handler, err := handlers.getEventSuccessHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		eventSuccessHandlers = append(eventSuccessHandlers, handler)
	}

	return eventSuccessHandlers, nil
}

func (t TransitionCfg) getEventErrorHandlers(stateMachine StateMachineType, handlers *Handlers) (eventErrorHandlers []EventErrorHandler, err error) {
	for _, name := range t.Events.Error {
		handler, err := handlers.getEventErrorHandler(stateMachine, name)
		if err != nil {
			return nil, err
		}
		eventErrorHandlers = append(eventErrorHandlers, handler)
	}

	return eventErrorHandlers, nil
}
