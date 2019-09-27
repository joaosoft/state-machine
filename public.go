package state_machine

// handlers
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

// state machine
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
