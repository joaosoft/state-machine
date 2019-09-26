package main

import (
	"fmt"
	state_machine "state-machine"
)

const (
	StateMachineA                        = "A"
	UserStateMachineA state_machine.User = "operator"

	StateMachineB                        = "B"
	UserStateMachineB state_machine.User = "worker"
)

func main() {
	var err error

	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_in-progress", CheckInProgress).
		AddCheckHandler("check_in-progress_to_approved", ExecuteInProgressToApproved).
		AddCheckHandler("check_in-progress_to_denied", ExecuteInProgressToDenied).
		AddCheckHandler("check_in-development_to_done", ExecuteInProgressToDone).
		AddCheckHandler("check_in-development_to_canceled", ExecuteInProgressToCanceled).
		AddExecuteHandler("execute_in-progress", ExecuteInProgress).
		AddExecuteHandler("execute_approved", ExecuteApproved).
		AddExecuteHandler("execute_denied", ExecuteDenied).
		AddExecuteHandler("execute_canceled", ExecuteCanceled).
		AddEventOnSuccessHandler("event_success_in-progress", EventOnSuccessInProgress).
		AddEventOnErrorHandler("event_error_in-progress", EventOnErrorInProgress).
		AddEventOnSuccessHandler("event_success_approved", EventOnSuccessApproved).
		AddEventOnErrorHandler("event_error_approved", EventOnErrorApproved).
		AddEventOnSuccessHandler("event_success_denied", EventOnSuccessDenied).
		AddEventOnErrorHandler("event_error_denied", EventOnErrorDenied).
		AddEventOnSuccessHandler("event_success_done", EventOnSuccessDone).
		AddEventOnErrorHandler("event_error_done", EventOnErrorDone).
		AddEventOnSuccessHandler("event_success_canceled", EventOnSuccessCanceled).
		AddEventOnErrorHandler("event_error_canceled", EventOnErrorCanceled).

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
	stateMachinesUsers := []state_machine.User{UserStateMachineA, UserStateMachineB}
	maxLen := 4
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
