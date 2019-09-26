package state_machine

func AddCheckHandler(name string, handler CheckHandler, stateMachine ...StateMachineType) *StateMachine {
	return stateMachineInstance.AddCheckHandler(name, handler, stateMachine...)
}

func AddExecuteHandler(name string, handler ExecuteHandler, stateMachine ...StateMachineType) *StateMachine {
	return stateMachineInstance.AddExecuteHandler(name, handler, stateMachine...)
}

func AddEventOnSuccessHandler(name string, handler EventHandler, stateMachine ...StateMachineType) *StateMachine {
	return stateMachineInstance.AddEventOnSuccessHandler(name, handler, stateMachine...)
}

func AddEventOnErrorHandler(name string, handler EventHandler, stateMachine ...StateMachineType) *StateMachine {
	return stateMachineInstance.AddEventOnErrorHandler(name, handler, stateMachine...)
}

func AddStateMachine(stateMachine StateMachineType, file string) error {
	return stateMachineInstance.Add(stateMachine, file)
}

func CheckTransition(stateMachine StateMachineType, user UserType, from int, to int, args ...interface{}) (bool, error) {
	return stateMachineInstance.CheckTransition(stateMachine, user, from, to, args...)
}

func ExecuteTransition(stateMachine StateMachineType, user UserType, from int, to int, args ...interface{}) (bool, error) {
	return stateMachineInstance.ExecuteTransition(stateMachine, user, from, to, args...)
}

func GetTransitions(stateMachine StateMachineType, user UserType, from int) (transitions []*Transition, err error) {
	return stateMachineInstance.GetTransitions(stateMachine, user, from)
}
