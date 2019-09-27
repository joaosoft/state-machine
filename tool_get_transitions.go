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

func (nt *newGetTransitions) User(name UserType) *newGetTransitions {
	nt.ctx.User = name
	return nt
}

func (nt *newGetTransitions) StateMachine(name StateMachineType) *newGetTransitions {
	nt.ctx.StateMachine = name
	return nt
}

func (nt *newGetTransitions) From(idStatus int) *newGetTransitions {
	nt.ctx.From = idStatus
	return nt
}

func (nt *newGetTransitions) Execute() ([]*Transition, error) {
	return nt.stateMachine.getTransitions(nt.ctx)
}
