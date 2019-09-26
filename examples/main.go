package main

import (
	"fmt"
	state_machine "state-machine"
)

const (
	StateMachineA     = "A"
	UserStateMachineA = "operator"

	StateMachineB     = "B"
	UserStateMachineB = "worker"
)

func main() {
	var err error

	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_in-progress", CheckInProgress).
		AddExecuteHandler("execute_in-progress", ExecuteInProgress).
		AddEventOnSuccessHandler("event_success_in-progress", EventOnSuccessInProgress).
		AddEventOnErrorHandler("event_error_in-progress", EventOnErrorInProgress).

		// state machine B
		AddCheckHandler("check_in-development", CheckInDevelopment).
		AddExecuteHandler("execute_in-development", ExecuteInDevelopment).
		AddEventOnSuccessHandler("event_success_in-development", EventOnSuccessInDevelopment).
		AddEventOnErrorHandler("event_error_in-development", EventOnErrorInDevelopment)

	// add state machines
	if err = state_machine.AddStateMachine(StateMachineA, "/config/state_machine_a.yaml"); err != nil {
		panic(err)
	}
	if err = state_machine.AddStateMachine(StateMachineB, "/config/state_machine_b.json"); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []string{StateMachineA, StateMachineB}
	stateMachinesUsers := []string{UserStateMachineA, UserStateMachineB}
	maxLen := 5
	ok := false

	for index, stateMachine := range stateMachines {
		fmt.Printf("\n\n\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err = state_machine.CheckTransition(stateMachine, stateMachinesUsers[index], i, j, 1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d  with user %s ? %t", i, j, stateMachinesUsers[index], ok)
			}
		}
	}

	// get all transitions of state machine A
	transitions, err := state_machine.GetTransitions(StateMachineA, UserStateMachineA, 1)
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}

	// execute transaction
	ok, err = state_machine.ExecuteTransition(StateMachineA, UserStateMachineA, 1, 2)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}
}
