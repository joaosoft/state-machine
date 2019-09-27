package state_machine

type newStateMachine struct {
	stateMachine      *stateMachine
	key               StateMachineType
	file              string
	transitionHandler TransitionHandler
}

func (sm *stateMachine) newStateMachine() *newStateMachine {
	return &newStateMachine{
		stateMachine: sm,
	}
}

func (nsm *newStateMachine) Key(key StateMachineType) *newStateMachine {
	nsm.key = key
	return nsm
}

func (nsm *newStateMachine) File(path string) *newStateMachine {
	nsm.file = path
	return nsm
}

func (nsm *newStateMachine) TransitionHandler(handler TransitionHandler) *newStateMachine {
	nsm.transitionHandler = handler
	return nsm
}

func (nsm *newStateMachine) Load() error {
	return nsm.stateMachine.add(nsm.key, nsm.file, nsm.transitionHandler)
}
