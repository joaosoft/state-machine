package state_machine

func AddStateMachine(name, file string) error {
	return stateMachineInstance.Add(name, file)
}

func CheckTransition(stateMachine string, from int, to int, args ...interface{}) (bool, error) {
	return stateMachineInstance.CheckTransition(stateMachine, from, to, args)
}

func GetTransitions(stateMachine string, from int) ([]*Transition, error) {
	return stateMachineInstance.GetTransitions(stateMachine, from)
}

func AddTransitionCheckHandler(name string, handler TransitionCheckHandler) *StateMachine {
	return stateMachineInstance.AddTransitionCheckHandler(name, handler)
}
