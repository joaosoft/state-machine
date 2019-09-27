package main

import (
	"fmt"
	state_machine "state-machine"
)

// ::  state machine A
// check
func CheckNewToInProgress(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress handler with %+v", args)
	return true, nil
}

func CheckInProgressToApproved(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to approved handler with %+v", args)
	return true, nil
}

func CheckInProgressToDenied(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-progress to denied handler with %+v", args)
	return true, nil
}

// execute
func ExecuteNewToInProgress(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress handler with %+v", args)
	return nil
}

func ExecuteNewToInProgressByUser(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nby user: execute in-progress handler with %+v", args)
	return nil
}

func ExecuteInProgressToApproved(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to approved handler with %+v", args)
	return nil
}

func ExecuteInProgressToDenied(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to denied handler with %+v", args)
	return nil
}

// event success
func EventOnSuccessNewToInProgress(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event in-progress handler with %+v", args)
}

func EventOnSuccessNewToInProgressByUser(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nby user: success event in-progress handler with %+v", args)
}

func EventOnSuccessInProgressToApproved(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event approved handler with %+v", args)
}

func EventOnSuccessInProgressToDenied(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event denied handler with %+v", args)
}

// event error
func EventOnErrorNewToInProgress(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event in-progress handler with %+v", err, args)
}

func EventOnErrorInProgressToApproved(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event approved handler with %+v", err, args)
}

func EventOnErrorInProgressToDenied(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event denied handler with %+v", err, args)
}

// transition handler
func StateMachineATransitionHandler(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, args)
	return nil
}

// :: state machine B
// check
func CheckTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development handler with %+v", args)
	return true, nil
}

func CheckInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to done handler with %+v", args)
	return true, nil
}

func CheckInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) (bool, error) {
	fmt.Printf("\ncheck in-development to canceled handler with %+v", args)
	return true, nil
}

// execute
func ExecuteTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-development handler with %+v", args)
	return nil
}

func ExecuteInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute canceled handler with %+v", args)
	return nil
}

func ExecuteInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nexecute in-progress to done handler with %+v", args)
	return nil
}

// event success
func EventOnSuccessInDevelopmentToDone(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event done handler with %+v", args)
}

func EventOnSuccessInDevelopmentToCanceled(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event canceled handler with %+v", args)
}

func EventOnSuccessTodoToInDevelopment(ctx *state_machine.Context, args ...interface{}) {
	fmt.Printf("\nsuccess event in-development handler with %+v", args)
}

// event error
func EventOnErrorInDevelopmentToDone(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event done handler with %+v", err, args)
}

func EventOnErrorInDevelopmentToCanceled(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event canceled handler with %+v", err, args)
}

func EventOnErrorTodoToInDevelopment(ctx *state_machine.Context, err error, args ...interface{}) {
	fmt.Printf("\nerror %s, event in-development handler with %+v", err, args)
}

// transition handler
func StateMachineBTransitionHandler(ctx *state_machine.Context, args ...interface{}) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, args)
	return nil
}
