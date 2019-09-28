package main

import (
	"fmt"
	state_machine "state-machine"
)

// ::  state machine A
// load
func loadDummy(ctx *state_machine.Context) error {
	fmt.Printf("\nload dummy handler with %+v", ctx.Args)
	return nil
}

// check
func checkNewToInProgress(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-progress handler with %+v", ctx.Args)
	return true, nil
}

func checkInProgressToApproved(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-progress to approved handler with %+v", ctx.Args)
	return true, nil
}

func checkInProgressToDenied(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-progress to denied handler with %+v", ctx.Args)
	return true, nil
}

// execute
func executeNewToInProgress(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute in-progress handler with %+v", ctx.Args)
	return nil
}

func executeNewToInProgressByUser(ctx *state_machine.Context) error {
	fmt.Printf("\nby user: execute in-progress handler with %+v", ctx.Args)
	return nil
}

func executeInProgressToApproved(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute in-progress to approved handler with %+v", ctx.Args)
	return nil
}

func executeInProgressToDenied(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute in-progress to denied handler with %+v", ctx.Args)
	return nil
}

// event success
func eventOnSuccessNewToInProgress(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event in-progress handler with %+v", ctx.Args)
	return nil
}

func eventOnSuccessNewToInProgressByUser(ctx *state_machine.Context) error {
	fmt.Printf("\nby user: success event in-progress handler with %+v", ctx.Args)
	return nil
}

func eventOnSuccessInProgressToApproved(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event approved handler with %+v", ctx.Args)
	return nil
}

func eventOnSuccessInProgressToDenied(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event denied handler with %+v", ctx.Args)
	return nil
}

// event error
func eventOnErrorNewToInProgress(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event in-progress handler with %+v", err, ctx.Args)
	return nil
}

func eventOnErrorInProgressToApproved(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event approved handler with %+v", err, ctx.Args)
	return nil
}

func eventOnErrorInProgressToDenied(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event denied handler with %+v", err, ctx.Args)
	return nil
}

// transition handler
func StateMachineATransitionHandler(ctx *state_machine.Context) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, ctx.Args)
	return nil
}

// :: state machine B
// manual - init
func beforeExecuteLoadFromState(ctx *state_machine.Context) error {
	fmt.Printf("\nload 'from' state handler with %+v", ctx.Args)
	ctx.From = 1
	return nil
}

// check
func checkTodoToInDevelopment(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-development handler with %+v", ctx.Args)
	return true, nil
}

func checkInDevelopmentToDone(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-development to done handler with %+v", ctx.Args)
	return true, nil
}

func checkInDevelopmentToCanceled(ctx *state_machine.Context) (bool, error) {
	fmt.Printf("\ncheck in-development to canceled handler with %+v", ctx.Args)
	return true, nil
}

// execute
func executeTodoToInDevelopment(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute in-development handler with %+v", ctx.Args)
	return nil
}

func executeInDevelopmentToCanceled(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute canceled handler with %+v", ctx.Args)
	return nil
}

func executeInDevelopmentToDone(ctx *state_machine.Context) error {
	fmt.Printf("\nexecute in-progress to done handler with %+v", ctx.Args)
	return nil
}

// event success
func eventOnSuccessInDevelopmentToDone(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event done handler with %+v", ctx.Args)
	return nil
}

func eventOnSuccessInDevelopmentToCanceled(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event canceled handler with %+v", ctx.Args)
	return nil
}

func eventOnSuccessTodoToInDevelopment(ctx *state_machine.Context) error {
	fmt.Printf("\nsuccess event in-development handler with %+v", ctx.Args)
	return nil
}

// event error
func eventOnErrorInDevelopmentToDone(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event done handler with %+v", err, ctx.Args)
	return nil
}

func eventOnErrorInDevelopmentToCanceled(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event canceled handler with %+v", err, ctx.Args)
	return nil
}

func eventOnErrorTodoToInDevelopment(ctx *state_machine.Context, err error) error {
	fmt.Printf("\nerror %s, event in-development handler with %+v", err, ctx.Args)
	return nil
}

// transition handler
func StateMachineBTransitionHandler(ctx *state_machine.Context) error {
	fmt.Printf("\nstate machine: %s, transition handler with %+v", ctx.StateMachine, ctx.Args)
	return nil
}
