package state_machine

type newTransition struct {
	stateMachine *stateMachine
	ctx          *Context
}

func (sm *stateMachine) newTransition() *newTransition {
	return &newTransition{
		stateMachine: sm,
		ctx:          &Context{},
	}
}

func (nt *newTransition) User(name UserType) *newTransition {
	nt.ctx.User = name
	return nt
}

func (nt *newTransition) StateMachine(name StateMachineType) *newTransition {
	nt.ctx.StateMachine = name
	return nt
}

func (nt *newTransition) From(idStatus int) *newTransition {
	nt.ctx.From = idStatus
	return nt
}

func (nt *newTransition) To(idStatus int) *newTransition {
	nt.ctx.To = idStatus
	return nt
}

func (nt *newTransition) Execute(args ...interface{}) (bool, error) {
	return nt.stateMachine.executeTransition(nt.ctx, args...)
}
