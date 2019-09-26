package main

import (
	"fmt"
	state_machine "state-machine"
)

const (
	StateMachineA     state_machine.StateMachineType = "A"
	UserStateMachineA state_machine.UserType         = "operator"

	StateMachineB     state_machine.StateMachineType = "B"
	UserStateMachineB state_machine.UserType         = "worker"
)

func main() {
	var err error

	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_new_to_in-progress", CheckNewToInProgress, StateMachineA).
		AddCheckHandler("check_in-progress_to_approved", CheckInProgressToApproved, StateMachineA).
		AddCheckHandler("check_in-progress_to_denied", CheckInProgressToDenied, StateMachineA).
		//
		AddExecuteHandler("execute_new_to_in-progress", ExecuteNewToInProgress, StateMachineA).
		AddExecuteHandler("execute_new_to_in-progress_user", ExecuteNewToInProgressUser, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_approved", ExecuteInProgressToApproved, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_denied", ExecuteInProgressToDenied, StateMachineA).
		//
		AddEventOnSuccessHandler("event_success_new_to_in-progress", EventOnSuccessNewToInProgress, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_approved", EventOnSuccessInProgressToApproved, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_denied", EventOnSuccessInProgressToDenied, StateMachineA).
		//
		AddEventOnErrorHandler("event_error_new_to_in-progress", EventOnErrorNewToInProgress, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_approved", EventOnErrorInProgressToApproved, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_denied", EventOnErrorInProgressToDenied, StateMachineA).

		// state machine B
		AddCheckHandler("check_todo_to_in-development", CheckTodoToInDevelopment, StateMachineB).
		AddCheckHandler("check_in-development_to_done", CheckInDevelopmentToDone, StateMachineB).
		AddCheckHandler("check_in-development_to_canceled", CheckInDevelopmentToCanceled, StateMachineB).
		//
		AddExecuteHandler("execute_todo_to_in-development", ExecuteTodoToInDevelopment, StateMachineB).
		AddExecuteHandler("execute_in-development_to_canceled", ExecuteInDevelopmentToCanceled, StateMachineB).
		AddExecuteHandler("execute_in-development_to_done", ExecuteInDevelopmentToDone, StateMachineB).
		//
		AddEventOnSuccessHandler("event_success_todo_to_in-development", EventOnSuccessTodoToInDevelopment, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_done", EventOnSuccessInDevelopmentToDone, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_canceled", EventOnSuccessInDevelopmentToCanceled, StateMachineB).
		//
		AddEventOnErrorHandler("event_error_todo_to_in-development", EventOnErrorTodoToInDevelopment, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_done", EventOnErrorInDevelopmentToDone, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_canceled", EventOnErrorInDevelopmentToCanceled, StateMachineB)

	// add state machines
	if err = state_machine.AddStateMachine(StateMachineA, "/config/state_machine_a.yaml"); err != nil {
		panic(err)
	}
	if err = state_machine.AddStateMachine(StateMachineB, "/config/state_machine_b.json"); err != nil {
		panic(err)
	}

	// check transitions of state machines
	stateMachines := []state_machine.StateMachineType{StateMachineA, StateMachineB}
	stateMachinesUsers := []state_machine.UserType{UserStateMachineA, UserStateMachineB}
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
