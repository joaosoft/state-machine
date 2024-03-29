package state_machine

type newCheckTransition struct {
	stateMachine *stateMachine
	ctx          *Context
}

func (sm *stateMachine) newCheckTransition() *newCheckTransition {
	return &newCheckTransition{
		stateMachine: sm,
		ctx:          &Context{},
	}
}

func (ct *newCheckTransition) Role(name RoleType) *newCheckTransition {
	ct.ctx.Role = name
	return ct
}

func (ct *newCheckTransition) StateMachine(name StateMachineType) *newCheckTransition {
	ct.ctx.StateMachine = name
	return ct
}

func (ct *newCheckTransition) From(idStatus string) *newCheckTransition {
	ct.ctx.From = idStatus
	return ct
}

func (ct *newCheckTransition) To(idStatus string) *newCheckTransition {
	ct.ctx.To = idStatus
	return ct
}

func (ct *newCheckTransition) Resource(id int) *newCheckTransition {
	ct.ctx.Resource = id
	return ct
}

func (ct *newCheckTransition) Execute(args ...interface{}) (bool, error) {
	ct.ctx.Args = args
	return ct.stateMachine.checkTransition(ct.ctx)
}
