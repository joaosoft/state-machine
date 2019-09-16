package main

import (
	"fmt"
	state_machine "state-machine"
)

const (
	StateMachineA = "A"
	StateMachineB = "B"
)

func main() {
	var err error

	// add transition check handlers
	state_machine.
		AddTransitionCheckHandler("check_in-progress", CheckInProgress).
		AddTransitionCheckHandler("check_in-development", CheckInDevelopment)

	// add state machines
	if err = state_machine.AddStateMachine(StateMachineA, "/config/state_machine_a.json"); err != nil {
		panic(err)
	}
	if err = state_machine.AddStateMachine(StateMachineB, "/config/state_machine_b.json"); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []string{StateMachineA, StateMachineB}
	maxLen := 5
	ok := false

	for _, stateMachine := range stateMachines {
		fmt.Printf("\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err = state_machine.CheckTransition(stateMachine, i, j, 1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d ? %t", i, j, ok)
			}
		}
	}

	// Get all transitions
	transitions, err := state_machine.GetTransitions(StateMachineA, 1)
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}
}

func CheckInProgress(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check in-progress handler with %+v", args)
	return true, nil
}

func CheckInDevelopment(args ...interface{}) (bool, error) {
	fmt.Printf("\nexecuting check in-development handler with %+v", args)
	return true, nil
}
