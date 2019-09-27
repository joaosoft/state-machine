package main

import (
	"fmt"
	state_machine "state-machine"
)

// ::  state machine A
// check
func checkNewToInProgress(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress handler with %+v", args)
	return true, nil
}

func checkInProgressToApproved(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to approved handler with %+v", args)
	return true, nil
}

func checkInProgressToDenied(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to denied handler with %+v", args)
	return true, nil
}

// execute
func executeNewToInProgress(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress handler with %+v", args)
	return nil
}

func executeNewToInProgressByUser(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nby user: execute in-progress handler with %+v", args)
	return nil
}

func executeInProgressToApproved(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to approved handler with %+v", args)
	return nil
}

func executeInProgressToDenied(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to denied handler with %+v", args)
	return nil
}

// event success
func eventOnSuccessNewToInProgress(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event in-progress handler with %+v", args)
}

func eventOnSuccessNewToInProgressByUser(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nby user: success event in-progress handler with %+v", args)
}

func eventOnSuccessInProgressToApproved(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event approved handler with %+v", args)
}

func eventOnSuccessInProgressToDenied(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event denied handler with %+v", args)
}

// event error
func eventOnErrorNewToInProgress(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event in-progress handler with %+v", err, args)
}

func eventOnErrorInProgressToApproved(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event approved handler with %+v", err, args)
}

func eventOnErrorInProgressToDenied(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event denied handler with %+v", err, args)
}

// transition handler
func StateMachineATransitionHandler(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, args)
	return nil
}

// :: state machine B
// check
func checkTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development handler with %+v", args)
	return true, nil
}

func checkInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to done handler with %+v", args)
	return true, nil
}

func checkInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to canceled handler with %+v", args)
	return true, nil
}

// execute
func executeTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-development handler with %+v", args)
	return nil
}

func executeInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute canceled handler with %+v", args)
	return nil
}

func executeInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to done handler with %+v", args)
	return nil
}

// event success
func eventOnSuccessInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event done handler with %+v", args)
}

func eventOnSuccessInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event canceled handler with %+v", args)
}

func eventOnSuccessTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event in-development handler with %+v", args)
}

// event error
func eventOnErrorInDevelopmentToDone(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event done handler with %+v", err, args)
}

func eventOnErrorInDevelopmentToCanceled(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event canceled handler with %+v", err, args)
}

func eventOnErrorTodoToInDevelopment(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event in-development handler with %+v", err, args)
}

// transition handler
func StateMachineBTransitionHandler(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, args)
	return nil
}
