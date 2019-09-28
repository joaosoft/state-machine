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

func init() {
	// :: add handlers

	// state machine A
	state_machine.NewAddHandlers(StateMachineA).
		Check("check_new_to_in-progress", checkNewToInProgress).
		Check("check_in-progress_to_approved", checkInProgressToApproved).
		Check("check_in-progress_to_denied", checkInProgressToDenied).
		//
		Execute("execute_new_to_in-progress", executeNewToInProgress).
		Execute("execute_new_to_in-progress_user", executeNewToInProgressByUser).
		Execute("execute_in-progress_to_approved", executeInProgressToApproved).
		Execute("execute_in-progress_to_denied", executeInProgressToDenied).
		//
		EventSuccess("event_success_new_to_in-progress_user", eventOnSuccessNewToInProgressByUser).
		EventSuccess("event_success_new_to_in-progress", eventOnSuccessNewToInProgress).
		EventSuccess("event_success_in-progress_to_approved", eventOnSuccessInProgressToApproved).
		EventSuccess("event_success_in-progress_to_denied", eventOnSuccessInProgressToDenied).
		//
		EventError("event_error_new_to_in-progress", eventOnErrorNewToInProgress).
		EventError("event_error_in-progress_to_approved", eventOnErrorInProgressToApproved).
		EventError("event_error_in-progress_to_denied", eventOnErrorInProgressToDenied)

	// state machine B
	state_machine.NewAddHandlers(StateMachineB).
		Manual(state_machine.ManualInit, loadFromState).
		//
		Check("check_todo_to_in-development", checkTodoToInDevelopment).
		Check("check_in-development_to_done", checkInDevelopmentToDone).
		Check("check_in-development_to_canceled", checkInDevelopmentToCanceled).
		//
		Execute("execute_todo_to_in-development", executeTodoToInDevelopment).
		Execute("execute_in-development_to_canceled", executeInDevelopmentToCanceled).
		Execute("execute_in-development_to_done", executeInDevelopmentToDone).
		//
		EventSuccess("event_success_todo_to_in-development", eventOnSuccessTodoToInDevelopment).
		EventSuccess("event_success_in-development_to_done", eventOnSuccessInDevelopmentToDone).
		EventSuccess("event_success_in-development_to_canceled", eventOnSuccessInDevelopmentToCanceled).
		//
		EventError("event_error_todo_to_in-development", eventOnErrorTodoToInDevelopment).
		EventError("event_error_in-development_to_done", eventOnErrorInDevelopmentToDone).
		EventError("event_error_in-development_to_canceled", eventOnErrorInDevelopmentToCanceled)

	// :: add state machines

	// A
	if err := state_machine.NewStateMachine().
		Key(StateMachineA).
		File("/config/state_machines/state_machine_a.yaml").
		TransitionHandler(StateMachineATransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// B
	if err := state_machine.NewStateMachine().
		Key(StateMachineB).
		File("/config/state_machines/state_machine_b.json").
		TransitionHandler(StateMachineBTransitionHandler).
		Load(); err != nil {
		panic(err)
	}
}

func main() {
	stateMachines := []state_machine.StateMachineType{StateMachineA, StateMachineB}
	stateMachinesUsers := []state_machine.UserType{UserStateMachineA, UserStateMachineB}
	maxLen := 4
	ok := false

	// check transitions of state machines
	for index, stateMachine := range stateMachines {
		fmt.Printf("\n\n\nState Machine: %s\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err := state_machine.NewCheckTransition().
					User(stateMachinesUsers[index]).
					StateMachine(stateMachine).
					From(i).
					To(j).
					Execute(1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\ntransition from %d to %d  with user %s ? %t", i, j, stateMachinesUsers[index], ok)
			}
		}
	}

	// get all transitions of state machine A
	transitions, err := state_machine.NewGetTransitions().
		User(UserStateMachineA).
		StateMachine(StateMachineA).
		From(1).
		Execute()
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("\ncan make transition to %s", transition.Name)
	}

	// execute transaction - state machine A - from state 1 to state 2
	ok, err = state_machine.NewTransition().
		User(UserStateMachineA).
		StateMachine(StateMachineA).
		From(1).
		To(2).
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}

	// execute transaction - state machine B - from the state loaded by method 'loadFromState' to state 2
	ok, err = state_machine.NewTransition().
		User(UserStateMachineB).
		StateMachine(StateMachineB).
		To(2).
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}
}
