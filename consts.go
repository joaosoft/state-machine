package state_machine

const (
	ManualInit   manualHandlerTag = "init" // do not run events
	ManualBefore manualHandlerTag = "before"
	ManualAfter  manualHandlerTag = "after"
)
