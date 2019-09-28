package state_machine

type addHandler struct {
	stateMachine  *stateMachine
	stateMachines []StateMachineType
}

func (sm *stateMachine) newAddHandler(stateMachine ...StateMachineType) *addHandler {
	return &addHandler{
		stateMachine:  sm,
		stateMachines: stateMachine,
	}
}

func (ah *addHandler) Manual(handler ManualHandler, tag manualHandlerKey, tags ...manualHandlerKey) *addHandler {
	ah.stateMachine.AddManualHandler(append(tags, tag), handler, ah.stateMachines...)
	return ah
}

func (ah *addHandler) Load(name string, handler LoadHandler) *addHandler {
	ah.stateMachine.addHandler(name, handler, ah.stateMachines...)
	return ah
}

func (ah *addHandler) Check(name string, handler CheckHandler) *addHandler {
	ah.stateMachine.addHandler(name, handler, ah.stateMachines...)
	return ah
}

func (ah *addHandler) Execute(name string, handler ExecuteHandler) *addHandler {
	ah.stateMachine.addHandler(name, handler, ah.stateMachines...)
	return ah
}

func (ah *addHandler) EventSuccess(name string, handler EventSuccessHandler) *addHandler {
	ah.stateMachine.addHandler(name, handler, ah.stateMachines...)
	return ah
}

func (ah *addHandler) EventError(name string, handler EventErrorHandler) *addHandler {
	ah.stateMachine.addHandler(name, handler, ah.stateMachines...)
	return ah
}
