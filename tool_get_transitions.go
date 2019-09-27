package state_machine

type newGetTransitions struct {
	stateMachine *stateMachine
	ctx          *Context
}

func (sm *stateMachine) newGetTransitions() *newGetTransitions {
	return &newGetTransitions{
		stateMachine: sm,
		ctx: &Context{},
	}
}

func (gt *newGetTransitions) User(name UserType) *newGetTransitions {
	gt.ctx.User = name
	return gt
}

func (gt *newGetTransitions) StateMachine(name StateMachineType) *newGetTransitions {
	gt.ctx.StateMachine = name
	return gt
}

func (gt *newGetTransitions) From(idStatus int) *newGetTransitions {
	gt.ctx.From = idStatus
	return gt
}

func (gt *newGetTransitions) Execute() ([]*Transition, error) {
	return gt.stateMachine.getTransitions(gt.ctx)
}
