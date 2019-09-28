package state_machine

type manualHandlerKey int

func NewAddHandlers(stateMachine StateMachineType) *addHandler {
	return stateMachineInstance.newAddHandler(stateMachine)
}

func NewStateMachine() *newStateMachine {
	return stateMachineInstance.newStateMachine()
}

func NewCheckTransition() *newCheckTransition {
	return stateMachineInstance.newCheckTransition()
}

func NewTransition() *newTransition {
	return stateMachineInstance.newTransition()
}

func NewGetTransitions() *newGetTransitions {
	return stateMachineInstance.newGetTransitions()
}

// manual handlers
func (sm *stateMachine) AddManualHandler(tags []manualHandlerKey, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addManualHandler(tags, handler, stateMachine...)
}

// state machine handlers
func (sm *stateMachine) AddLoadHandler(name string, handler LoadHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addHandler(name, handler, stateMachine...)
}

func (sm *stateMachine) AddCheckHandler(name string, handler CheckHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addHandler(name, handler, stateMachine...)
}

func (sm *stateMachine) AddExecuteHandler(name string, handler ExecuteHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addHandler(name, handler, stateMachine...)
}

func (sm *stateMachine) AddEventOnSuccessHandler(name string, handler EventSuccessHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addHandler(name, handler, stateMachine...)
}

func (sm *stateMachine) AddEventOnErrorHandler(name string, handler EventErrorHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addHandler(name, handler, stateMachine...)
}
