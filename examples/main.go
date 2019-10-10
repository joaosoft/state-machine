package main

import (
	"fmt"
	state_machine "state-machine"
	"strconv"
)

const (
	StateMachineA     state_machine.StateMachineType = "A"
	RoleStateMachineA state_machine.RoleType         = "operator"

	StateMachineB     state_machine.StateMachineType = "B"
	RoleStateMachineB state_machine.RoleType         = "worker"
)

func init() {
	// :: add handlers

	// state machine A
	fmt.Println(":: State Machine: A - Adding handlers")
	state_machine.NewAddHandlers(StateMachineA).
		Load("load_dummy", loadDummy).
		//
		Check("check_new_to_in-progress", checkNewToInProgress).
		Check("check_in-progress_to_approved", checkInProgressToApproved).
		Check("check_in-progress_to_denied", checkInProgressToDenied).
		//
		Execute("execute_new_to_in-progress", executeNewToInProgress).
		Execute("execute_new_to_in-progress_role", executeNewToInProgressByRole).
		Execute("execute_in-progress_to_approved", executeInProgressToApproved).
		Execute("execute_in-progress_to_denied", executeInProgressToDenied).
		//
		EventSuccess("event_success_new_to_in-progress_role", eventOnSuccessNewToInProgressByRole).
		EventSuccess("event_success_new_to_in-progress", eventOnSuccessNewToInProgress).
		EventSuccess("event_success_in-progress_to_approved", eventOnSuccessInProgressToApproved).
		EventSuccess("event_success_in-progress_to_denied", eventOnSuccessInProgressToDenied).
		//
		EventError("event_error_new_to_in-progress", eventOnErrorNewToInProgress).
		EventError("event_error_in-progress_to_approved", eventOnErrorInProgressToApproved).
		EventError("event_error_in-progress_to_denied", eventOnErrorInProgressToDenied)

	// state machine B
	fmt.Println(":: State Machine: B - Adding handlers")
	state_machine.NewAddHandlers(StateMachineB).
		Manual(beforeExecuteLoadFromState, state_machine.BeforeCheck, state_machine.BeforeExecute).
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
	fmt.Println(":: State Machine: A - Adding state machine")
	if err := state_machine.NewStateMachine().
		Key(StateMachineA).
		File("/config/state_machines/state_machine_a.yaml").
		TransitionHandler(StateMachineATransitionHandler).
		Load(); err != nil {
		panic(err)
	}

	// B
	fmt.Println(":: State Machine: B - Adding state machine")
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
	stateMachinesRoles := []state_machine.RoleType{RoleStateMachineA, RoleStateMachineB}
	maxLen := 4
	ok := false

	// get all transitions of state machine A
	fmt.Println("\n:: State Machine: A - get all transition from 1 to 2")
	transitions, err := state_machine.NewGetTransitions().
		Role(RoleStateMachineA).
		StateMachine(StateMachineA).
		From("1").
		Execute()
	if err != nil {
		panic(err)
	}
	for _, transition := range transitions {
		fmt.Printf("can make transition to %s\n", transition.Name)
	}

	// check transitions of state machines
	for index, stateMachine := range stateMachines {
		fmt.Printf("\n:: State Machine: %s - check transitions\n", stateMachine)
		for i := 1; i <= maxLen; i++ {
			for j := maxLen; j >= 1; j-- {
				ok, err := state_machine.NewCheckTransition().
					Role(stateMachinesRoles[index]).
					StateMachine(stateMachine).
					From(strconv.Itoa(i)).
					To(strconv.Itoa(j)).
					Execute(1, "text", true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("transition from %d to %d  with role %s ? %t\n", i, j, stateMachinesRoles[index], ok)
			}
		}
	}

	// check transaction - state machine B - from the state loaded by method 'beforeExecuteLoadFromState' to state 2
	fmt.Println("\n:: State Machine: B - check transition from state 1 (loaded) to state 2")
	ok, err = state_machine.NewTransition().
		Role(RoleStateMachineB).
		StateMachine(StateMachineB).
		To("2").
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}

	// execute transaction - state machine B - from the state loaded by method 'beforeExecuteLoadFromState' to state 2
	fmt.Println("\n:: State Machine: B - making transition from state 1 (loaded) to state 2")
	ok, err = state_machine.NewTransition().
		Role(RoleStateMachineB).
		StateMachine(StateMachineB).
		To("2").
		Execute(1, "text", true)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("transition !ok")
	}
}
