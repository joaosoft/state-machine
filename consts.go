package state_machine

const (
	BeforeExecute manualHandlerKey = iota
	AfterExecute
	BeforeCheck
	OnSuccess
	OnError
)
