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
	// add handlers
	state_machine.
		// state machine A
		AddCheckHandler("check_new_to_in-progress", checkNewToInProgress, StateMachineA).
		AddCheckHandler("check_in-progress_to_approved", checkInProgressToApproved, StateMachineA).
		AddCheckHandler("check_in-progress_to_denied", checkInProgressToDenied, StateMachineA).
		//
		AddExecuteHandler("execute_new_to_in-progress", executeNewToInProgress, StateMachineA).
		AddExecuteHandler("execute_new_to_in-progress_user", executeNewToInProgressByUser, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_approved", executeInProgressToApproved, StateMachineA).
		AddExecuteHandler("execute_in-progress_to_denied", executeInProgressToDenied, StateMachineA).
		//
		AddEventOnSuccessHandler("event_success_new_to_in-progress_user", eventOnSuccessNewToInProgressByUser, StateMachineA).
		AddEventOnSuccessHandler("event_success_new_to_in-progress", eventOnSuccessNewToInProgress, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_approved", eventOnSuccessInProgressToApproved, StateMachineA).
		AddEventOnSuccessHandler("event_success_in-progress_to_denied", eventOnSuccessInProgressToDenied, StateMachineA).
		//
		AddEventOnErrorHandler("event_error_new_to_in-progress", eventOnErrorNewToInProgress, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_approved", eventOnErrorInProgressToApproved, StateMachineA).
		AddEventOnErrorHandler("event_error_in-progress_to_denied", eventOnErrorInProgressToDenied, StateMachineA).

		// state machine B
		AddCheckHandler("check_todo_to_in-development", checkTodoToInDevelopment, StateMachineB).
		AddCheckHandler("check_in-development_to_done", checkInDevelopmentToDone, StateMachineB).
		AddCheckHandler("check_in-development_to_canceled", checkInDevelopmentToCanceled, StateMachineB).
		//
		AddExecuteHandler("execute_todo_to_in-development", executeTodoToInDevelopment, StateMachineB).
		AddExecuteHandler("execute_in-development_to_canceled", executeInDevelopmentToCanceled, StateMachineB).
		AddExecuteHandler("execute_in-development_to_done", executeInDevelopmentToDone, StateMachineB).
		//
		AddEventOnSuccessHandler("event_success_todo_to_in-development", eventOnSuccessTodoToInDevelopment, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_done", eventOnSuccessInDevelopmentToDone, StateMachineB).
		AddEventOnSuccessHandler("event_success_in-development_to_canceled", eventOnSuccessInDevelopmentToCanceled, StateMachineB).
		//
		AddEventOnErrorHandler("event_error_todo_to_in-development", eventOnErrorTodoToInDevelopment, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_done", eventOnErrorInDevelopmentToDone, StateMachineB).
		AddEventOnErrorHandler("event_error_in-development_to_canceled", eventOnErrorInDevelopmentToCanceled, StateMachineB)

	// add state machines
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

	// execute transaction
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
}
