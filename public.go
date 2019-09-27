package state_machine

type ManualHandlerTag string

// manual handlers
func AddManualHandler(tag ManualHandlerTag, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addManualHandler(tag, handler, stateMachine...)
}

// handlers
func AddLoadHandler(name string, handler LoadHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addHandler(name, handler, stateMachine...)
}

func AddCheckHandler(name string, handler CheckHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addHandler(name, handler, stateMachine...)
}

func AddExecuteHandler(name string, handler ExecuteHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addHandler(name, handler, stateMachine...)
}

func AddEventOnSuccessHandler(name string, handler EventSuccessHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addHandler(name, handler, stateMachine...)
}

func AddEventOnErrorHandler(name string, handler EventSuccessHandler, stateMachine ...StateMachineType) *stateMachine {
	return stateMachineInstance.addHandler(name, handler, stateMachine...)
}

// tools

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
func (sm *stateMachine) AddManualHandler(tag ManualHandlerTag, handler ManualHandler, stateMachine ...StateMachineType) *stateMachine {
	return sm.addManualHandler(tag, handler, stateMachine...)
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
