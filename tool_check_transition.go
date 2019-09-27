package state_machine

type newCheckTransition struct {
	stateMachine *stateMachine
	ctx          *Context
}

func (sm *stateMachine) newCheckTransition() *newCheckTransition {
	return &newCheckTransition{
		stateMachine: sm,
		ctx: &Context{},
	}
}

func (nt *newCheckTransition) User(name UserType) *newCheckTransition {
	nt.ctx.User = name
	return nt
}

func (nt *newCheckTransition) StateMachine(name StateMachineType) *newCheckTransition {
	nt.ctx.StateMachine = name
	return nt
}

func (nt *newCheckTransition) From(idStatus int) *newCheckTransition {
	nt.ctx.From = idStatus
	return nt
}

func (nt *newCheckTransition) To(idStatus int) *newCheckTransition {
	nt.ctx.To = idStatus
	return nt
}

func (nt *newCheckTransition) Execute(args ...interface{}) (bool, error) {
	return nt.stateMachine.checkTransition(nt.ctx, args...)
}
