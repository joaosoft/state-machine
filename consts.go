package state_machine

const (
	ManualInit   ManualHandlerTag = "init" // do not run events
	ManualBefore ManualHandlerTag = "before"
	ManualAfter  ManualHandlerTag = "after"
)
